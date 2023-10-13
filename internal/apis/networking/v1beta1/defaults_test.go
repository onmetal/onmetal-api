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

package v1beta1_test

import (
	ipamv1beta1 "github.com/onmetal/onmetal-api/api/ipam/v1beta1"
	networkingv1beta1 "github.com/onmetal/onmetal-api/api/networking/v1beta1"
	. "github.com/onmetal/onmetal-api/internal/apis/networking/v1beta1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
)

var _ = Describe("Defaults", func() {
	Describe("SetDefaults_NetworkInterfaceSpec", func() {
		It("should default the prefix length of ephemeral ips depending on the ip family", func() {
			spec := &networkingv1beta1.NetworkInterfaceSpec{
				IPFamilies: []corev1.IPFamily{corev1.IPv4Protocol, corev1.IPv6Protocol},
				IPs: []networkingv1beta1.IPSource{
					{
						Ephemeral: &networkingv1beta1.EphemeralPrefixSource{
							PrefixTemplate: &ipamv1beta1.PrefixTemplateSpec{
								Spec: ipamv1beta1.PrefixSpec{
									IPFamily:  corev1.IPv4Protocol,
									ParentRef: &corev1.LocalObjectReference{Name: "parent-v4"},
								},
							},
						},
					},
					{
						Ephemeral: &networkingv1beta1.EphemeralPrefixSource{
							PrefixTemplate: &ipamv1beta1.PrefixTemplateSpec{
								Spec: ipamv1beta1.PrefixSpec{
									IPFamily:  corev1.IPv6Protocol,
									ParentRef: &corev1.LocalObjectReference{Name: "parent-v6"},
								},
							},
						},
					},
				},
			}
			SetDefaults_NetworkInterfaceSpec(spec)

			Expect(spec.IPs).To(Equal([]networkingv1beta1.IPSource{
				{
					Ephemeral: &networkingv1beta1.EphemeralPrefixSource{
						PrefixTemplate: &ipamv1beta1.PrefixTemplateSpec{
							Spec: ipamv1beta1.PrefixSpec{
								IPFamily:     corev1.IPv4Protocol,
								ParentRef:    &corev1.LocalObjectReference{Name: "parent-v4"},
								PrefixLength: 32,
							},
						},
					},
				},
				{
					Ephemeral: &networkingv1beta1.EphemeralPrefixSource{
						PrefixTemplate: &ipamv1beta1.PrefixTemplateSpec{
							Spec: ipamv1beta1.PrefixSpec{
								IPFamily:     corev1.IPv6Protocol,
								ParentRef:    &corev1.LocalObjectReference{Name: "parent-v6"},
								PrefixLength: 128,
							},
						},
					},
				},
			}))
		})
	})
})
