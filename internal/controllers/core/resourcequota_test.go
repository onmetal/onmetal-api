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
	computev1alpha1 "github.com/onmetal/onmetal-api/api/compute/v1alpha1"
	corev1alpha1 "github.com/onmetal/onmetal-api/api/core/v1alpha1"
	. "github.com/onmetal/onmetal-api/utils/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	. "sigs.k8s.io/controller-runtime/pkg/envtest/komega"
)

var _ = Describe("ResourceNamespaceController", func() {
	var (
		ctx          = SetupContext()
		ns           = SetupTest(ctx)
		machineClass = &computev1alpha1.MachineClass{}
	)

	BeforeEach(func() {
		*machineClass = computev1alpha1.MachineClass{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "machine-class-",
			},
			Capabilities: corev1alpha1.ResourceList{
				corev1alpha1.ResourceCPU:    resource.MustParse("1"),
				corev1alpha1.ResourceMemory: resource.MustParse("1Gi"),
			},
		}
		Expect(k8sClient.Create(ctx, machineClass)).To(Succeed())
	})

	It("should mark the namespace with the replenish quota annotation", MustPassRepeatedly(3), func() {
		By("creating a machine")
		machine := &computev1alpha1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "machine-",
			},
			Spec: computev1alpha1.MachineSpec{
				MachineClassRef: corev1.LocalObjectReference{Name: machineClass.Name},
			},
		}
		Expect(k8sClient.Create(ctx, machine)).To(Succeed())

		By("creating a resource quota")
		resourceQuota := &corev1alpha1.ResourceQuota{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "resource-quota-",
			},
			Spec: corev1alpha1.ResourceQuotaSpec{
				Hard: corev1alpha1.ResourceList{
					corev1alpha1.ResourceRequestsCPU:    resource.MustParse("2"),
					corev1alpha1.ResourceRequestsMemory: resource.MustParse("2Gi"),
				},
			},
		}
		Expect(k8sClient.Create(ctx, resourceQuota)).To(Succeed())

		By("waiting for the resource quota to report status")
		Eventually(Object(resourceQuota)).Should(HaveField("Status", SatisfyAll(
			HaveField("Hard", resourceQuota.Spec.Hard),
			HaveField("Used", corev1alpha1.ResourceList{
				corev1alpha1.ResourceRequestsCPU:    resource.MustParse("1"),
				corev1alpha1.ResourceRequestsMemory: resource.MustParse("1Gi"),
			}),
		)))

		By("getting the resource version of the namespace pre machine deletion")
		Expect(k8sClient.Get(ctx, client.ObjectKeyFromObject(ns), ns)).Should(Succeed())
		preMachineDeletionNamespaceResourceVersion := ns.ResourceVersion

		By("deleting the machine")
		Expect(k8sClient.Delete(ctx, machine)).To(Succeed())

		By("waiting for the resource quota to be updated")
		Eventually(Object(resourceQuota)).Should(HaveField("Status", SatisfyAll(
			HaveField("Hard", resourceQuota.Spec.Hard),
			HaveField("Used", corev1alpha1.ResourceList{
				corev1alpha1.ResourceRequestsCPU:    resource.MustParse("0"),
				corev1alpha1.ResourceRequestsMemory: resource.MustParse("0"),
			}),
		)))

		By("fetching the namespace and compare the resource version")
		Expect(k8sClient.Get(ctx, client.ObjectKeyFromObject(ns), ns)).Should(Succeed())
		Expect(ns.ResourceVersion).NotTo(Equal(preMachineDeletionNamespaceResourceVersion))
	})
})
