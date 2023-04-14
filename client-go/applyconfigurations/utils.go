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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1alpha1 "github.com/onmetal/onmetal-api/api/common/v1alpha1"
	computev1alpha1 "github.com/onmetal/onmetal-api/api/compute/v1alpha1"
	corev1alpha1 "github.com/onmetal/onmetal-api/api/core/v1alpha1"
	ipamv1alpha1 "github.com/onmetal/onmetal-api/api/ipam/v1alpha1"
	networkingv1alpha1 "github.com/onmetal/onmetal-api/api/networking/v1alpha1"
	storagev1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	commonv1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/common/v1alpha1"
	applyconfigurationscomputev1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/compute/v1alpha1"
	applyconfigurationscorev1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/core/v1alpha1"
	applyconfigurationsipamv1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/ipam/v1alpha1"
	metav1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/meta/v1"
	applyconfigurationsnetworkingv1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/networking/v1alpha1"
	applyconfigurationsstoragev1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=common.api.onmetal.de, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("LocalUIDReference"):
		return &commonv1alpha1.LocalUIDReferenceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SecretKeySelector"):
		return &commonv1alpha1.SecretKeySelectorApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Taint"):
		return &commonv1alpha1.TaintApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Toleration"):
		return &commonv1alpha1.TolerationApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("UIDReference"):
		return &commonv1alpha1.UIDReferenceApplyConfiguration{}

		// Group=compute.api.onmetal.de, Version=v1alpha1
	case computev1alpha1.SchemeGroupVersion.WithKind("DaemonEndpoint"):
		return &applyconfigurationscomputev1alpha1.DaemonEndpointApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("EFIVar"):
		return &applyconfigurationscomputev1alpha1.EFIVarApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("EmptyDiskVolumeSource"):
		return &applyconfigurationscomputev1alpha1.EmptyDiskVolumeSourceApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("EphemeralNetworkInterfaceSource"):
		return &applyconfigurationscomputev1alpha1.EphemeralNetworkInterfaceSourceApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("EphemeralVolumeSource"):
		return &applyconfigurationscomputev1alpha1.EphemeralVolumeSourceApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("Machine"):
		return &applyconfigurationscomputev1alpha1.MachineApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("MachineClass"):
		return &applyconfigurationscomputev1alpha1.MachineClassApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("MachinePool"):
		return &applyconfigurationscomputev1alpha1.MachinePoolApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("MachinePoolAddress"):
		return &applyconfigurationscomputev1alpha1.MachinePoolAddressApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("MachinePoolCondition"):
		return &applyconfigurationscomputev1alpha1.MachinePoolConditionApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("MachinePoolDaemonEndpoints"):
		return &applyconfigurationscomputev1alpha1.MachinePoolDaemonEndpointsApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("MachinePoolSpec"):
		return &applyconfigurationscomputev1alpha1.MachinePoolSpecApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("MachinePoolStatus"):
		return &applyconfigurationscomputev1alpha1.MachinePoolStatusApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("MachineSpec"):
		return &applyconfigurationscomputev1alpha1.MachineSpecApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("MachineStatus"):
		return &applyconfigurationscomputev1alpha1.MachineStatusApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("NetworkInterface"):
		return &applyconfigurationscomputev1alpha1.NetworkInterfaceApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("NetworkInterfaceSource"):
		return &applyconfigurationscomputev1alpha1.NetworkInterfaceSourceApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("NetworkInterfaceStatus"):
		return &applyconfigurationscomputev1alpha1.NetworkInterfaceStatusApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("Volume"):
		return &applyconfigurationscomputev1alpha1.VolumeApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("VolumeSource"):
		return &applyconfigurationscomputev1alpha1.VolumeSourceApplyConfiguration{}
	case computev1alpha1.SchemeGroupVersion.WithKind("VolumeStatus"):
		return &applyconfigurationscomputev1alpha1.VolumeStatusApplyConfiguration{}

		// Group=core.api.onmetal.de, Version=v1alpha1
	case corev1alpha1.SchemeGroupVersion.WithKind("ObjectSelector"):
		return &applyconfigurationscorev1alpha1.ObjectSelectorApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("ResourceQuota"):
		return &applyconfigurationscorev1alpha1.ResourceQuotaApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("ResourceQuotaSpec"):
		return &applyconfigurationscorev1alpha1.ResourceQuotaSpecApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("ResourceQuotaStatus"):
		return &applyconfigurationscorev1alpha1.ResourceQuotaStatusApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("ResourceScopeSelector"):
		return &applyconfigurationscorev1alpha1.ResourceScopeSelectorApplyConfiguration{}
	case corev1alpha1.SchemeGroupVersion.WithKind("ResourceScopeSelectorRequirement"):
		return &applyconfigurationscorev1alpha1.ResourceScopeSelectorRequirementApplyConfiguration{}

		// Group=ipam.api.onmetal.de, Version=v1alpha1
	case ipamv1alpha1.SchemeGroupVersion.WithKind("Prefix"):
		return &applyconfigurationsipamv1alpha1.PrefixApplyConfiguration{}
	case ipamv1alpha1.SchemeGroupVersion.WithKind("PrefixAllocation"):
		return &applyconfigurationsipamv1alpha1.PrefixAllocationApplyConfiguration{}
	case ipamv1alpha1.SchemeGroupVersion.WithKind("PrefixAllocationSpec"):
		return &applyconfigurationsipamv1alpha1.PrefixAllocationSpecApplyConfiguration{}
	case ipamv1alpha1.SchemeGroupVersion.WithKind("PrefixAllocationStatus"):
		return &applyconfigurationsipamv1alpha1.PrefixAllocationStatusApplyConfiguration{}
	case ipamv1alpha1.SchemeGroupVersion.WithKind("PrefixSpec"):
		return &applyconfigurationsipamv1alpha1.PrefixSpecApplyConfiguration{}
	case ipamv1alpha1.SchemeGroupVersion.WithKind("PrefixStatus"):
		return &applyconfigurationsipamv1alpha1.PrefixStatusApplyConfiguration{}
	case ipamv1alpha1.SchemeGroupVersion.WithKind("PrefixTemplateSpec"):
		return &applyconfigurationsipamv1alpha1.PrefixTemplateSpecApplyConfiguration{}

		// Group=meta.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("LabelSelector"):
		return &metav1.LabelSelectorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("LabelSelectorRequirement"):
		return &metav1.LabelSelectorRequirementApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ManagedFieldsEntry"):
		return &metav1.ManagedFieldsEntryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ObjectMeta"):
		return &metav1.ObjectMetaApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("OwnerReference"):
		return &metav1.OwnerReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TypeMeta"):
		return &metav1.TypeMetaApplyConfiguration{}

		// Group=networking.api.onmetal.de, Version=v1alpha1
	case networkingv1alpha1.SchemeGroupVersion.WithKind("AliasPrefix"):
		return &applyconfigurationsnetworkingv1alpha1.AliasPrefixApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("AliasPrefixRouting"):
		return &applyconfigurationsnetworkingv1alpha1.AliasPrefixRoutingApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("AliasPrefixSpec"):
		return &applyconfigurationsnetworkingv1alpha1.AliasPrefixSpecApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("AliasPrefixStatus"):
		return &applyconfigurationsnetworkingv1alpha1.AliasPrefixStatusApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("EphemeralPrefixSource"):
		return &applyconfigurationsnetworkingv1alpha1.EphemeralPrefixSourceApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("EphemeralVirtualIPSource"):
		return &applyconfigurationsnetworkingv1alpha1.EphemeralVirtualIPSourceApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("IPBlock"):
		return &applyconfigurationsnetworkingv1alpha1.IPBlockApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("IPSource"):
		return &applyconfigurationsnetworkingv1alpha1.IPSourceApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("LoadBalancer"):
		return &applyconfigurationsnetworkingv1alpha1.LoadBalancerApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("LoadBalancerPort"):
		return &applyconfigurationsnetworkingv1alpha1.LoadBalancerPortApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("LoadBalancerRouting"):
		return &applyconfigurationsnetworkingv1alpha1.LoadBalancerRoutingApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("LoadBalancerSpec"):
		return &applyconfigurationsnetworkingv1alpha1.LoadBalancerSpecApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("LoadBalancerStatus"):
		return &applyconfigurationsnetworkingv1alpha1.LoadBalancerStatusApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NATGateway"):
		return &applyconfigurationsnetworkingv1alpha1.NATGatewayApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NATGatewayDestination"):
		return &applyconfigurationsnetworkingv1alpha1.NATGatewayDestinationApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NATGatewayDestinationIP"):
		return &applyconfigurationsnetworkingv1alpha1.NATGatewayDestinationIPApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NATGatewayIP"):
		return &applyconfigurationsnetworkingv1alpha1.NATGatewayIPApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NATGatewayIPStatus"):
		return &applyconfigurationsnetworkingv1alpha1.NATGatewayIPStatusApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NATGatewayRouting"):
		return &applyconfigurationsnetworkingv1alpha1.NATGatewayRoutingApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NATGatewaySpec"):
		return &applyconfigurationsnetworkingv1alpha1.NATGatewaySpecApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NATGatewayStatus"):
		return &applyconfigurationsnetworkingv1alpha1.NATGatewayStatusApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("Network"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkInterface"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkInterfaceApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkInterfaceSpec"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkInterfaceSpecApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkInterfaceStatus"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkInterfaceStatusApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkInterfaceTemplateSpec"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkInterfaceTemplateSpecApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPeering"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPeeringApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPeeringStatus"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPeeringStatusApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPolicy"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPolicyApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPolicyCondition"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPolicyConditionApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPolicyEgressRule"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPolicyEgressRuleApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPolicyIngressRule"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPolicyIngressRuleApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPolicyPeer"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPolicyPeerApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPolicyPort"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPolicyPortApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPolicySpec"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPolicySpecApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkPolicyStatus"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkPolicyStatusApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkSpec"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkSpecApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("NetworkStatus"):
		return &applyconfigurationsnetworkingv1alpha1.NetworkStatusApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("PrefixSource"):
		return &applyconfigurationsnetworkingv1alpha1.PrefixSourceApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("VirtualIP"):
		return &applyconfigurationsnetworkingv1alpha1.VirtualIPApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("VirtualIPSource"):
		return &applyconfigurationsnetworkingv1alpha1.VirtualIPSourceApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("VirtualIPSpec"):
		return &applyconfigurationsnetworkingv1alpha1.VirtualIPSpecApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("VirtualIPStatus"):
		return &applyconfigurationsnetworkingv1alpha1.VirtualIPStatusApplyConfiguration{}
	case networkingv1alpha1.SchemeGroupVersion.WithKind("VirtualIPTemplateSpec"):
		return &applyconfigurationsnetworkingv1alpha1.VirtualIPTemplateSpecApplyConfiguration{}

		// Group=storage.api.onmetal.de, Version=v1alpha1
	case storagev1alpha1.SchemeGroupVersion.WithKind("Bucket"):
		return &applyconfigurationsstoragev1alpha1.BucketApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("BucketAccess"):
		return &applyconfigurationsstoragev1alpha1.BucketAccessApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("BucketClass"):
		return &applyconfigurationsstoragev1alpha1.BucketClassApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("BucketCondition"):
		return &applyconfigurationsstoragev1alpha1.BucketConditionApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("BucketPool"):
		return &applyconfigurationsstoragev1alpha1.BucketPoolApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("BucketPoolSpec"):
		return &applyconfigurationsstoragev1alpha1.BucketPoolSpecApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("BucketPoolStatus"):
		return &applyconfigurationsstoragev1alpha1.BucketPoolStatusApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("BucketSpec"):
		return &applyconfigurationsstoragev1alpha1.BucketSpecApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("BucketStatus"):
		return &applyconfigurationsstoragev1alpha1.BucketStatusApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("Volume"):
		return &applyconfigurationsstoragev1alpha1.VolumeApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumeAccess"):
		return &applyconfigurationsstoragev1alpha1.VolumeAccessApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumeClass"):
		return &applyconfigurationsstoragev1alpha1.VolumeClassApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumeCondition"):
		return &applyconfigurationsstoragev1alpha1.VolumeConditionApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumeEncryption"):
		return &applyconfigurationsstoragev1alpha1.VolumeEncryptionApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumePool"):
		return &applyconfigurationsstoragev1alpha1.VolumePoolApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumePoolCondition"):
		return &applyconfigurationsstoragev1alpha1.VolumePoolConditionApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumePoolSpec"):
		return &applyconfigurationsstoragev1alpha1.VolumePoolSpecApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumePoolStatus"):
		return &applyconfigurationsstoragev1alpha1.VolumePoolStatusApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumeSpec"):
		return &applyconfigurationsstoragev1alpha1.VolumeSpecApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumeStatus"):
		return &applyconfigurationsstoragev1alpha1.VolumeStatusApplyConfiguration{}
	case storagev1alpha1.SchemeGroupVersion.WithKind("VolumeTemplateSpec"):
		return &applyconfigurationsstoragev1alpha1.VolumeTemplateSpecApplyConfiguration{}

	}
	return nil
}
