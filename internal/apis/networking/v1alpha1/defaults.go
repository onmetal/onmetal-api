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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/utils/pointer"
)

var (
	ipFamilyToPrefixLength = map[corev1.IPFamily]int32{
		corev1.IPv4Protocol: 32,
		corev1.IPv6Protocol: 128,
	}
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_NetworkPolicySpec(spec *v1beta1.NetworkPolicySpec) {
	policyTypes := sets.New[v1beta1.PolicyType](spec.PolicyTypes...)
	if len(spec.Ingress) > 0 {
		policyTypes.Insert(v1beta1.PolicyTypeIngress)
	}
	if len(spec.Egress) > 0 {
		policyTypes.Insert(v1beta1.PolicyTypeEgress)
	}
	spec.PolicyTypes = sets.List(policyTypes)
}

func SetDefaults_NetworkInterfaceSpec(spec *v1beta1.NetworkInterfaceSpec) {
	setDefaults_IPFamiliesIPSources(&spec.IPFamilies, &spec.IPs)
}

func SetDefaults_LoadBalancerSpec(spec *v1beta1.LoadBalancerSpec) {
	setDefaults_IPFamiliesIPSources(&spec.IPFamilies, &spec.IPs)
}

func setDefaults_IPFamiliesIPSources(ipFamilies *[]corev1.IPFamily, ipSources *[]v1beta1.IPSource) {
	if len(*ipFamilies) > 0 {
		if len(*ipFamilies) == len(*ipSources) {
			for i, ip := range *ipSources {
				if ip.Ephemeral != nil {
					if ip.Ephemeral.PrefixTemplate != nil {
						ephemeralPrefixSpec := &ip.Ephemeral.PrefixTemplate.Spec

						if ephemeralPrefixSpec.IPFamily == "" {
							ephemeralPrefixSpec.IPFamily = (*ipFamilies)[i]
						}
					}
				}
			}
		}
	} else if len(*ipSources) > 0 {
		for _, ip := range *ipSources {
			switch {
			case ip.Value != nil:
				*ipFamilies = append(*ipFamilies, ip.Value.Family())
			case ip.Ephemeral != nil && ip.Ephemeral.PrefixTemplate != nil:
				*ipFamilies = append(*ipFamilies, ip.Ephemeral.PrefixTemplate.Spec.IPFamily)
			}
		}
	}

	for _, ip := range *ipSources {
		if ip.Ephemeral != nil && ip.Ephemeral.PrefixTemplate != nil {
			templateSpec := &ip.Ephemeral.PrefixTemplate.Spec
			if templateSpec.Prefix == nil && templateSpec.PrefixLength == 0 {
				templateSpec.PrefixLength = ipFamilyToPrefixLength[templateSpec.IPFamily]
			}
		}
	}
}

func SetDefaults_NetworkInterfaceStatus(status *v1beta1.NetworkInterfaceStatus) {
	if status.State == "" {
		status.State = v1beta1.NetworkInterfaceStatePending
	}
}

func SetDefaults_NATGatewaySpec(spec *v1beta1.NATGatewaySpec) {
	if spec.PortsPerNetworkInterface == nil {
		spec.PortsPerNetworkInterface = pointer.Int32(v1beta1.DefaultPortsPerNetworkInterface)
	}
}
