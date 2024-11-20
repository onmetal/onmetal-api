/*
 * Copyright (c) 2021 by the OnMetal authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package compute

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	computev1alpha1 "github.com/onmetal/onmetal-api/api/compute/v1alpha1"
	corev1alpha1 "github.com/onmetal/onmetal-api/api/core/v1alpha1"
	ipamv1alpha1 "github.com/onmetal/onmetal-api/api/ipam/v1alpha1"
	networkingv1alpha1 "github.com/onmetal/onmetal-api/api/networking/v1alpha1"
	storagev1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	computeclient "github.com/onmetal/onmetal-api/internal/client/compute"
	networkingclient "github.com/onmetal/onmetal-api/internal/client/networking"
	"github.com/onmetal/onmetal-api/internal/controllers/networking"
	"github.com/onmetal/onmetal-api/internal/controllers/storage"
	utilsenvtest "github.com/onmetal/onmetal-api/utils/envtest"
	"github.com/onmetal/onmetal-api/utils/envtest/apiserver"

	"github.com/onmetal/controller-utils/buildutils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/komega"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	metricserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

const (
	pollingInterval      = 50 * time.Millisecond
	eventuallyTimeout    = 3 * time.Second
	consistentlyDuration = 1 * time.Second
	apiServiceTimeout    = 5 * time.Minute
)

var (
	cfg        *rest.Config
	k8sClient  client.Client
	testEnv    *envtest.Environment
	testEnvExt *utilsenvtest.EnvironmentExtensions
)

func TestAPIs(t *testing.T) {
	SetDefaultConsistentlyPollingInterval(pollingInterval)
	SetDefaultEventuallyPollingInterval(pollingInterval)
	SetDefaultEventuallyTimeout(eventuallyTimeout)
	SetDefaultConsistentlyDuration(consistentlyDuration)

	RegisterFailHandler(Fail)

	RunSpecs(t, "Compute Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	var err error

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{}
	testEnvExt = &utilsenvtest.EnvironmentExtensions{
		APIServiceDirectoryPaths:       []string{filepath.Join("..", "..", "..", "config", "apiserver", "apiservice", "bases")},
		ErrorIfAPIServicePathIsMissing: true,
	}

	cfg, err = utilsenvtest.StartWithExtensions(testEnv, testEnvExt)
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())
	DeferCleanup(utilsenvtest.StopWithExtensions, testEnv, testEnvExt)

	Expect(computev1alpha1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(storagev1alpha1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(ipamv1alpha1.AddToScheme(scheme.Scheme)).To(Succeed())
	Expect(networkingv1alpha1.AddToScheme(scheme.Scheme)).To(Succeed())

	//+kubebuilder:scaffold:scheme

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).NotTo(HaveOccurred())
	Expect(k8sClient).NotTo(BeNil())

	komega.SetClient(k8sClient)

	apiSrv, err := apiserver.New(cfg, apiserver.Options{
		MainPath:     "github.com/onmetal/onmetal-api/cmd/onmetal-apiserver",
		BuildOptions: []buildutils.BuildOption{buildutils.ModModeMod},
		ETCDServers:  []string{testEnv.ControlPlane.Etcd.URL.String()},
		Host:         testEnvExt.APIServiceInstallOptions.LocalServingHost,
		Port:         testEnvExt.APIServiceInstallOptions.LocalServingPort,
		CertDir:      testEnvExt.APIServiceInstallOptions.LocalServingCertDir,
	})
	Expect(err).NotTo(HaveOccurred())

	Expect(apiSrv.Start()).To(Succeed())
	DeferCleanup(apiSrv.Stop)

	err = utilsenvtest.WaitUntilAPIServicesReadyWithTimeout(apiServiceTimeout, testEnvExt, k8sClient, scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())
})

// SetupTest returns a namespace which will be created before each ginkgo `It` block and deleted at the end of `It`
// so that each test case can run in an independent way
func SetupTest(ctx context.Context) (*corev1.Namespace, *computev1alpha1.MachineClass) {
	var (
		cancel       context.CancelFunc
		ns           = &corev1.Namespace{}
		machineClass = &computev1alpha1.MachineClass{}
	)

	BeforeEach(func() {
		var mgrCtx context.Context
		mgrCtx, cancel = context.WithCancel(ctx)
		*ns = corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "testns-",
			},
		}
		Expect(k8sClient.Create(ctx, ns)).To(Succeed(), "failed to create test namespace")

		*machineClass = computev1alpha1.MachineClass{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "machine-class-",
			},
			Capabilities: corev1alpha1.ResourceList{
				corev1alpha1.ResourceCPU:    resource.MustParse("1"),
				corev1alpha1.ResourceMemory: resource.MustParse("1Gi"),
			},
		}
		Expect(k8sClient.Create(ctx, machineClass)).To(Succeed(), "failed to create test machine class")

		k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
			Scheme: scheme.Scheme,
			Metrics: metricserver.Options{
				BindAddress: "0",
			},
		})
		Expect(err).ToNot(HaveOccurred())

		Expect(computeclient.SetupMachineSpecNetworkInterfaceNamesFieldIndexer(ctx, k8sManager.GetFieldIndexer())).To(Succeed())
		Expect(computeclient.SetupMachineSpecMachineClassRefNameFieldIndexer(ctx, k8sManager.GetFieldIndexer())).To(Succeed())
		Expect(computeclient.SetupMachineSpecVolumeNamesFieldIndexer(ctx, k8sManager.GetFieldIndexer())).To(Succeed())
		Expect(computeclient.SetupMachinePoolAvailableMachineClassesFieldIndexer(ctx, k8sManager.GetFieldIndexer())).To(Succeed())
		Expect(computeclient.SetupMachineSpecMachinePoolRefNameFieldIndexer(ctx, k8sManager.GetFieldIndexer())).To(Succeed())

		Expect(networkingclient.SetupNetworkInterfaceNetworkNameFieldIndexer(ctx, k8sManager.GetFieldIndexer())).To(Succeed())
		Expect(networkingclient.SetupNetworkInterfaceVirtualIPNameFieldIndexer(ctx, k8sManager.GetFieldIndexer())).To(Succeed())

		// register reconciler here
		Expect((&MachineScheduler{
			Client:        k8sManager.GetClient(),
			EventRecorder: &record.FakeRecorder{},
		}).SetupWithManager(k8sManager)).To(Succeed())

		Expect((&MachineReconciler{
			EventRecorder: &record.FakeRecorder{},
			Client:        k8sManager.GetClient(),
			Scheme:        k8sManager.GetScheme(),
		}).SetupWithManager(k8sManager)).To(Succeed())

		Expect((&MachineClassReconciler{
			Client:    k8sManager.GetClient(),
			APIReader: k8sManager.GetAPIReader(),
		}).SetupWithManager(k8sManager)).To(Succeed())

		Expect((&networking.NetworkInterfaceReconciler{
			EventRecorder: &record.FakeRecorder{},
			Client:        k8sManager.GetClient(),
			Scheme:        k8sManager.GetScheme(),
		}).SetupWithManager(k8sManager)).To(Succeed())

		Expect((&storage.VolumeReconciler{
			EventRecorder: &record.FakeRecorder{},
			Client:        k8sManager.GetClient(),
			APIReader:     k8sManager.GetAPIReader(),
			Scheme:        k8sManager.GetScheme(),
			BindTimeout:   1 * time.Second,
		}).SetupWithManager(k8sManager)).To(Succeed())

		go func() {
			defer GinkgoRecover()
			Expect(k8sManager.Start(mgrCtx)).To(Succeed(), "failed to start manager")
		}()
	})

	AfterEach(func() {
		cancel()
		Expect(k8sClient.Delete(ctx, ns)).To(Succeed(), "failed to delete test namespace")
		Expect(k8sClient.DeleteAllOf(ctx, &computev1alpha1.MachinePool{})).To(Succeed())
		Expect(k8sClient.DeleteAllOf(ctx, &computev1alpha1.MachineClass{})).To(Succeed())
	})

	return ns, machineClass
}
