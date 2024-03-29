// Copyright 2023 OnMetal authors
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

package core_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/onmetal/controller-utils/buildutils"
	computev1alpha1 "github.com/onmetal/onmetal-api/api/compute/v1alpha1"
	corev1alpha1 "github.com/onmetal/onmetal-api/api/core/v1alpha1"
	storagev1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	"github.com/onmetal/onmetal-api/internal/controllers/core"
	certificateonmetal "github.com/onmetal/onmetal-api/internal/controllers/core/certificate/onmetal"
	quotacontrollergeneric "github.com/onmetal/onmetal-api/internal/controllers/core/quota/generic"
	quotacontrolleronmetal "github.com/onmetal/onmetal-api/internal/controllers/core/quota/onmetal"
	quotaevaluatoronmetal "github.com/onmetal/onmetal-api/internal/quota/evaluator/onmetal"
	utilsenvtest "github.com/onmetal/onmetal-api/utils/envtest"
	"github.com/onmetal/onmetal-api/utils/envtest/apiserver"
	"github.com/onmetal/onmetal-api/utils/quota"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/komega"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	metricserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

const (
	pollingInterval      = 50 * time.Millisecond
	eventuallyTimeout    = 3 * time.Second
	consistentlyDuration = 1 * time.Second
	apiServiceTimeout    = 5 * time.Minute
)

var (
	k8sClient  client.Client
	testEnv    *envtest.Environment
	testEnvExt *utilsenvtest.EnvironmentExtensions
)

func TestCore(t *testing.T) {
	SetDefaultConsistentlyPollingInterval(pollingInterval)
	SetDefaultEventuallyPollingInterval(pollingInterval)
	SetDefaultEventuallyTimeout(eventuallyTimeout)
	SetDefaultConsistentlyDuration(consistentlyDuration)
	RegisterFailHandler(Fail)

	RunSpecs(t, "Core Controller Suite")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{}
	testEnvExt = &utilsenvtest.EnvironmentExtensions{
		APIServiceDirectoryPaths:       []string{filepath.Join("..", "..", "..", "config", "apiserver", "apiservice", "bases")},
		ErrorIfAPIServicePathIsMissing: true,
	}

	cfg, err := utilsenvtest.StartWithExtensions(testEnv, testEnvExt)
	Expect(err).NotTo(HaveOccurred())
	Expect(cfg).NotTo(BeNil())

	DeferCleanup(utilsenvtest.StopWithExtensions, testEnv, testEnvExt)

	Expect(corev1alpha1.AddToScheme(scheme.Scheme)).Should(Succeed())
	Expect(computev1alpha1.AddToScheme(scheme.Scheme)).Should(Succeed())
	Expect(storagev1alpha1.AddToScheme(scheme.Scheme)).Should(Succeed())

	// Init package-level k8sClient
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

	Expect(utilsenvtest.WaitUntilAPIServicesReadyWithTimeout(apiServiceTimeout, testEnvExt, k8sClient, scheme.Scheme)).To(Succeed())

	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
		Metrics: metricserver.Options{
			BindAddress: "0",
		},
	})
	Expect(err).ToNot(HaveOccurred())

	registry := quota.NewRegistry(scheme.Scheme)
	Expect(quota.AddAllToRegistry(registry, quotaevaluatoronmetal.NewEvaluatorsForControllers(k8sManager.GetClient()))).To(Succeed())

	replenishReconcilers, err := quotacontrolleronmetal.NewReplenishReconcilers(k8sManager.GetClient(), registry)
	Expect(err).NotTo(HaveOccurred())
	Expect(quotacontrollergeneric.SetupReplenishReconcilersWithManager(k8sManager, replenishReconcilers)).To(Succeed())

	Expect((&core.ResourceQuotaReconciler{
		Client:    k8sManager.GetClient(),
		APIReader: k8sManager.GetAPIReader(),
		Scheme:    scheme.Scheme,
		Registry:  registry,
	}).SetupWithManager(k8sManager)).To(Succeed())

	Expect((&core.CertificateApprovalReconciler{
		Client:      k8sManager.GetClient(),
		Recognizers: certificateonmetal.Recognizers,
	}).SetupWithManager(k8sManager)).To(Succeed())

	mgrCtx, cancel := context.WithCancel(context.Background())
	DeferCleanup(cancel)

	go func() {
		defer GinkgoRecover()
		err = k8sManager.Start(mgrCtx)
		Expect(err).ToNot(HaveOccurred())
	}()
})

func SetupTest(ctx context.Context) *corev1.Namespace {
	ns := &corev1.Namespace{}

	BeforeEach(func() {
		*ns = corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{GenerateName: "testns-"},
		}
		Expect(k8sClient.Create(ctx, ns)).NotTo(HaveOccurred(), "failed to create test namespace")

		DeferCleanup(k8sClient.Delete, ctx, ns)
	})

	return ns
}
