// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"context"
	goflag "flag"
	"fmt"
	"os"
	"time"

	"github.com/go-logr/logr"
	ipamv1alpha1 "github.com/onmetal/onmetal-api/api/ipam/v1alpha1"
	networkingv1alpha1 "github.com/onmetal/onmetal-api/api/networking/v1alpha1"
	storagev1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	ori "github.com/onmetal/onmetal-api/ori/apis/storage/v1alpha1"
	"github.com/onmetal/onmetal-api/volumepoollet/controllers"
	"github.com/onmetal/onmetal-api/volumepoollet/vcm"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(storagev1alpha1.AddToScheme(scheme))
	utilruntime.Must(storagev1alpha1.AddToScheme(scheme))
	utilruntime.Must(networkingv1alpha1.AddToScheme(scheme))
	utilruntime.Must(ipamv1alpha1.AddToScheme(scheme))
}

type Options struct {
	MetricsAddr          string
	EnableLeaderElection bool
	ProbeAddr            string

	VolumePoolName               string
	ProviderID                   string
	VolumeRuntimeEndpoint        string
	DialTimeout                  time.Duration
	VolumeClassMapperSyncTimeout time.Duration

	WatchFilterValue string
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.MetricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	fs.StringVar(&o.ProbeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	fs.BoolVar(&o.EnableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")

	fs.StringVar(&o.VolumePoolName, "volume-pool-name", o.VolumePoolName, "Name of the volume pool to announce / watch")
	fs.StringVar(&o.ProviderID, "provider-id", "", "Provider id to announce on the volume pool.")
	fs.StringVar(&o.VolumeRuntimeEndpoint, "volume-runtime-endpoint", o.VolumeRuntimeEndpoint, "Endpoint of the remote volume runtime service.")
	fs.DurationVar(&o.DialTimeout, "dial-timeout", 1*time.Second, "Timeout for dialing to the volume runtime endpoint.")
	fs.DurationVar(&o.VolumeClassMapperSyncTimeout, "vcm-sync-timeout", 10*time.Second, "Timeout waiting for the volume class mapper to sync.")

	fs.StringVar(&o.WatchFilterValue, "watch-filter", "", "Value to filter for while watching.")
}

func (o *Options) MarkFlagsRequired(cmd *cobra.Command) {
	_ = cmd.MarkFlagRequired("volume-pool-name")
	_ = cmd.MarkFlagRequired("provider-id")
}

func Command() *cobra.Command {
	var (
		zapOpts = zap.Options{Development: true}
		opts    Options
	)

	cmd := &cobra.Command{
		Use: "volumepoollet",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger := zap.New(zap.UseFlagOptions(&zapOpts))
			ctrl.SetLogger(logger)
			cmd.SetContext(ctrl.LoggerInto(cmd.Context(), ctrl.Log))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			return Run(ctx, opts)
		},
	}

	goFlags := goflag.NewFlagSet("", 0)
	zapOpts.BindFlags(goFlags)
	cmd.PersistentFlags().AddGoFlagSet(goFlags)

	opts.AddFlags(cmd.Flags())
	opts.MarkFlagsRequired(cmd)

	return cmd
}

var wellKnownVolumeRuntimeEndpoints = []string{
	"/var/run/ori-volumebroker.sock",
	"/var/run/ori-cephd.sock",
}

func DetectVolumeRuntimeEndpoint(log logr.Logger, endpoint string) (string, error) {
	if endpoint != "" {
		log.V(1).Info("Using explicitly defined endpoint", "Endpoint", endpoint)
		return endpoint, nil
	}

	endpoint = os.Getenv("ORI_VOLUME_RUNTIME")
	if endpoint != "" {
		log.V(1).Info("Using volume runtime endpoint from environment variable", "Endpoint", endpoint)
		return endpoint, nil
	}

	for _, wellKnownEndpoint := range wellKnownVolumeRuntimeEndpoints {
		if stat, err := os.Stat(wellKnownEndpoint); err == nil && stat.Mode().Type()&os.ModeSocket != 0 {
			log.V(1).Info("Detected and using well-known endpoint", "Endpoint", wellKnownEndpoint)
			return wellKnownEndpoint, nil
		}
	}

	return "", fmt.Errorf("error detecting volume runtime endpoint")
}

func Run(ctx context.Context, opts Options) error {
	logger := ctrl.LoggerFrom(ctx)
	setupLog := ctrl.Log.WithName("setup")

	endpoint, err := DetectVolumeRuntimeEndpoint(setupLog, opts.VolumeRuntimeEndpoint)
	if err != nil {
		return fmt.Errorf("error detecting volume runtime endpoint: %w", err)
	}

	dialCtx, cancel := context.WithTimeout(ctx, opts.DialTimeout)
	defer cancel()
	conn, err := grpc.DialContext(dialCtx, endpoint,
		grpc.WithReturnConnectionError(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("error dialing: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			setupLog.Error(err, "Error closing volume runtime connection")
		}
	}()

	volumeRuntime := ori.NewVolumeRuntimeClient(conn)

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Logger:                 logger,
		Scheme:                 scheme,
		MetricsBindAddress:     opts.MetricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: opts.ProbeAddr,
		LeaderElection:         opts.EnableLeaderElection,
		LeaderElectionID:       "bfafcebe.api.onmetal.de",
	})
	if err != nil {
		return fmt.Errorf("error creating manager: %w", err)
	}

	onInitialized := func(ctx context.Context) error {
		volumeClassMapper := vcm.NewGeneric(volumeRuntime, vcm.GenericOptions{})
		if err := mgr.Add(volumeClassMapper); err != nil {
			return fmt.Errorf("error adding volume class mapper: %w", err)
		}

		volumeClassMapperSyncCtx, cancel := context.WithTimeout(ctx, opts.VolumeClassMapperSyncTimeout)
		defer cancel()

		if err := volumeClassMapper.WaitForSync(volumeClassMapperSyncCtx); err != nil {
			return fmt.Errorf("error waiting for volume class mapper to sync: %w", err)
		}

		if err := (&controllers.VolumeReconciler{
			EventRecorder:     mgr.GetEventRecorderFor("volumes"),
			Client:            mgr.GetClient(),
			VolumeRuntime:     volumeRuntime,
			VolumeClassMapper: volumeClassMapper,
			VolumePoolName:    opts.VolumePoolName,
			WatchFilterValue:  opts.WatchFilterValue,
		}).SetupWithManager(mgr); err != nil {
			return fmt.Errorf("error setting up volume reconciler with manager: %w", err)
		}

		if err := (&controllers.VolumePoolReconciler{
			Client:            mgr.GetClient(),
			VolumePoolName:    opts.VolumePoolName,
			VolumeClassMapper: volumeClassMapper,
			VolumeRuntime:     volumeRuntime,
		}).SetupWithManager(mgr); err != nil {
			return fmt.Errorf("error setting up volume pool reconciler with manager: %w", err)
		}

		return nil
	}

	if err := (&controllers.VolumePoolInit{
		Client:         mgr.GetClient(),
		VolumePoolName: opts.VolumePoolName,
		ProviderID:     opts.ProviderID,
		OnInitialized:  onInitialized,
	}).SetupWithManager(mgr); err != nil {
		return fmt.Errorf("error setting up volume pool init with manager: %w", err)
	}

	// TODO: Add vleg healthz check
	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}

	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctx); err != nil {
		return fmt.Errorf("error running manager: %w", err)
	}
	return nil
}