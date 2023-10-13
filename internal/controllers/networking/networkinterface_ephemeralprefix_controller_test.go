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
package networking

import (
	commonv1beta1 "github.com/onmetal/onmetal-api/api/common/v1beta1"
	ipamv1beta1 "github.com/onmetal/onmetal-api/api/ipam/v1beta1"
	networkingv1beta1 "github.com/onmetal/onmetal-api/api/networking/v1beta1"
	"github.com/onmetal/onmetal-api/utils/annotations"
	. "github.com/onmetal/onmetal-api/utils/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	. "sigs.k8s.io/controller-runtime/pkg/envtest/komega"
)

var _ = Describe("NetworkInterfaceEphemeralPrefix", func() {
	ns := SetupNamespace(&k8sClient)

	It("should manage ephemeral IP prefixes for a network interface", func(ctx SpecContext) {
		By("creating a network interface that requires a prefix")
		nic := &networkingv1beta1.NetworkInterface{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "nic-",
			},
			Spec: networkingv1beta1.NetworkInterfaceSpec{
				NetworkRef: corev1.LocalObjectReference{Name: "my-network"},
				IPs: []networkingv1beta1.IPSource{
					{
						Ephemeral: &networkingv1beta1.EphemeralPrefixSource{
							PrefixTemplate: &ipamv1beta1.PrefixTemplateSpec{
								Spec: ipamv1beta1.PrefixSpec{
									IPFamily: corev1.IPv4Protocol,
									Prefix:   commonv1beta1.MustParseNewIPPrefix("10.0.0.1/32"),
								},
							},
						},
					},
				},
			},
		}
		Expect(k8sClient.Create(ctx, nic)).To(Succeed())

		By("waiting for the prefix to exist")
		prefix := &ipamv1beta1.Prefix{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: ns.Name,
				Name:      networkingv1beta1.NetworkInterfaceIPIPAMPrefixName(nic.Name, 0),
			},
		}
		Eventually(Object(prefix)).Should(SatisfyAll(
			BeControlledBy(nic),
			HaveField("Spec", ipamv1beta1.PrefixSpec{
				IPFamily: corev1.IPv4Protocol,
				Prefix:   commonv1beta1.MustParseNewIPPrefix("10.0.0.1/32"),
			}),
		))
	})

	It("should delete undesired prefixes for a network interface", func(ctx SpecContext) {
		By("creating a network interface")
		nic := &networkingv1beta1.NetworkInterface{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "nic-",
			},
			Spec: networkingv1beta1.NetworkInterfaceSpec{
				NetworkRef: corev1.LocalObjectReference{Name: "my-network"},
				IPs:        []networkingv1beta1.IPSource{{Value: commonv1beta1.MustParseNewIP("10.0.0.1")}},
			},
		}
		Expect(k8sClient.Create(ctx, nic)).To(Succeed())

		By("creating an undesired prefix")
		prefix := &ipamv1beta1.Prefix{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "undesired-prefix-",
			},
			Spec: ipamv1beta1.PrefixSpec{
				IPFamily: corev1.IPv4Protocol,
				Prefix:   commonv1beta1.MustParseNewIPPrefix("10.0.0.1/32"),
			},
		}
		annotations.SetDefaultEphemeralManagedBy(prefix)
		Expect(ctrl.SetControllerReference(nic, prefix, k8sClient.Scheme())).To(Succeed())
		Expect(k8sClient.Create(ctx, prefix)).To(Succeed())

		By("waiting for the prefix to be marked for deletion")
		Eventually(Get(prefix)).Should(Satisfy(apierrors.IsNotFound))
	})

	It("should not delete externally managed prefix for a network interface", func(ctx SpecContext) {
		By("creating a network interface")
		nic := &networkingv1beta1.NetworkInterface{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "nic-",
			},
			Spec: networkingv1beta1.NetworkInterfaceSpec{
				NetworkRef: corev1.LocalObjectReference{Name: "my-network"},
				IPs:        []networkingv1beta1.IPSource{{Value: commonv1beta1.MustParseNewIP("10.0.0.1")}},
			},
		}
		Expect(k8sClient.Create(ctx, nic)).To(Succeed())

		By("creating an undesired prefix")
		externalPrefix := &ipamv1beta1.Prefix{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:    ns.Name,
				GenerateName: "external-prefix-",
			},
			Spec: ipamv1beta1.PrefixSpec{
				IPFamily: corev1.IPv4Protocol,
				Prefix:   commonv1beta1.MustParseNewIPPrefix("10.0.0.1/32"),
			},
		}
		Expect(ctrl.SetControllerReference(nic, externalPrefix, k8sClient.Scheme())).To(Succeed())
		Expect(k8sClient.Create(ctx, externalPrefix)).To(Succeed())

		By("asserting that the prefix is not being deleted")
		Eventually(Object(externalPrefix)).Should(HaveField("DeletionTimestamp", BeNil()))
	})
})
