/*
 * Copyright (c) 2022 by the OnMetal authors.
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

package validation

import (
	"github.com/onmetal/onmetal-api/internal/apis/networking"
	. "github.com/onmetal/onmetal-api/internal/testutils/validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Network", func() {
	DescribeTable("ValidateNetwork",
		func(network *networking.Network, match types.GomegaMatcher) {
			errList := ValidateNetwork(network)
			Expect(errList).To(match)
		},
		Entry("missing name",
			&networking.Network{},
			ContainElement(RequiredField("metadata.name")),
		),
		Entry("missing namespace",
			&networking.Network{ObjectMeta: metav1.ObjectMeta{Name: "foo"}},
			ContainElement(RequiredField("metadata.namespace")),
		),
		Entry("bad name",
			&networking.Network{ObjectMeta: metav1.ObjectMeta{Name: "foo*"}},
			ContainElement(InvalidField("metadata.name")),
		),
		Entry("peering references itself",
			&networking.Network{
				ObjectMeta: metav1.ObjectMeta{Namespace: "ns", Name: "foo"},
				Spec: networking.NetworkSpec{
					Peerings: []networking.NetworkPeering{
						{
							Name: "peering",
							NetworkRef: networking.NetworkPeeringNetworkRef{
								Name: "foo",
							},
						},
					},
				},
			},
			ContainElement(ForbiddenField("spec.peerings[0].networkRef")),
		),
		Entry("duplicate peering name",
			&networking.Network{
				Spec: networking.NetworkSpec{
					Peerings: []networking.NetworkPeering{
						{Name: "peering"},
						{Name: "peering"},
					},
				},
			},
			ContainElement(DuplicateField("spec.peerings[1].name")),
		),
		Entry("duplicate network ref",
			&networking.Network{
				Spec: networking.NetworkSpec{
					Peerings: []networking.NetworkPeering{
						{NetworkRef: networking.NetworkPeeringNetworkRef{Name: "bar"}},
						{NetworkRef: networking.NetworkPeeringNetworkRef{Name: "bar"}},
					},
				},
			},
			ContainElement(DuplicateField("spec.peerings[1].networkRef")),
		),
	)

	DescribeTable("ValidateNetworkUpdate",
		func(newNetwork, oldNetwork *networking.Network, match types.GomegaMatcher) {
			errList := ValidateNetworkUpdate(newNetwork, oldNetwork)
			Expect(errList).To(match)
		},
		Entry("immutable providerID if set",
			&networking.Network{
				Spec: networking.NetworkSpec{
					ProviderID: "foo",
				},
			},
			&networking.Network{
				Spec: networking.NetworkSpec{
					ProviderID: "bar",
				},
			},
			ContainElement(ImmutableField("spec.providerID")),
		),
		Entry("mutable providerID if not set",
			&networking.Network{
				Spec: networking.NetworkSpec{
					ProviderID: "foo",
				},
			},
			&networking.Network{},
			Not(ContainElement(ImmutableField("spec.providerID"))),
		),
	)
})
