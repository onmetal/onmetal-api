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

package networking

import (
	commonv1alpha1 "github.com/onmetal/onmetal-api/api/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	// Handle is the identifier of the network provider.
	Handle string
	// InternetGateway is a flag that indicates whether the network has an internet gateway.
	InternetGateway bool `json:"internetGateway,omitempty"`
	// Peerings are the network peerings with this network.
	// ProviderID is the provider-internal ID of the network.
	ProviderID string
	// Peerings are the network peerings with this network.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	Peerings []NetworkPeering

	// PeeringClaimRefs are the peering claim references of other networks.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	PeeringClaimRefs []NetworkPeeringClaimRef
}

type NetworkPeeringClaimRef struct {
	// Namespace is the namespace of the referenced entity. If empty,
	// the same namespace as the referring resource is implied.
	Namespace string
	// Name is the name of the referenced entity.
	Name string
	// UID is the UID of the referenced entity.
	UID types.UID
}

// NetworkPeeringNetworkRef is a reference to a network to peer with.
type NetworkPeeringNetworkRef struct {
	// Namespace is the namespace of the referenced entity. If empty,
	// the same namespace as the referring resource is implied.
	Namespace string
	// Name is the name of the referenced entity.
	Name string
}

// NetworkPeering defines a network peering with another network.
type NetworkPeering struct {
	// Name is the semantical name of the network peering.
	Name string
	// Prefixes is a list of CIDRs that we want only to be exposed to the peered network, if no prefixes are specified no filtering will be done.
	Prefixes *[]commonv1alpha1.IPPrefix
	// TODO this will explode
	// THIS IS THE OLD WAY

	// NetworkRef is the reference to the network to peer with.
	// If the UID is empty, it will be populated once when the peering is successfully bound.
	// If namespace is empty it is implied that the target network resides in the same network.
	NetworkRef commonv1alpha1.UIDReference
}

/*
	// THIS IS THE NEW WAY
	// NetworkRef is the reference to the network to peer with.
	// An empty namespace indicates that the target network resides in the same namespace as the source network.
	NetworkRef NetworkPeeringNetworkRef
}
*/

// NetworkStatus defines the observed state of Network
type NetworkStatus struct {
	// State is the state of the machine.
	State NetworkState
	// Peerings contains the states of the network peerings for the network.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge,retainKeys
	Peerings []NetworkPeeringStatus
}

// NetworkState is the state of a network.
// +enum
type NetworkState string

// NetworkPeeringStatus is the status of a network peering.
type NetworkPeeringStatus struct {
	// Name is the name of the network peering.
	Name string
	// NetworkHandle is the handle of the peered network.
	NetworkHandle string
	// Prefixes is a list of CIDRs that we want only to be exposed to the peered network, if no prefixes are specified no filtering will be done.
	Prefixes *[]commonv1alpha1.IPPrefix
	// Phase represents the binding phase of a network peering.
	Phase NetworkPeeringPhase
	// LastPhaseTransitionTime is the last time the Phase transitioned.
	LastPhaseTransitionTime *metav1.Time
}

const (
	// NetworkStatePending means the network is being provisioned.
	NetworkStatePending NetworkState = "Pending"
	// NetworkStateAvailable means the network is ready to use.
	NetworkStateAvailable NetworkState = "Available"
	// NetworkStateError means the network is in an error state.
	NetworkStateError NetworkState = "Error"
)

// NetworkPeeringPhase is the phase a NetworkPeering can be in.
type NetworkPeeringPhase string

const (
	// NetworkPeeringPhasePending signals that the network peering is not bound.
	NetworkPeeringPhasePending NetworkPeeringPhase = "Pending"
	// NetworkPeeringPhaseBound signals that the network peering is bound.
	NetworkPeeringPhaseBound NetworkPeeringPhase = "Bound"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Network is the Schema for the network API
type Network struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   NetworkSpec
	Status NetworkStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkList contains a list of Network
type NetworkList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Network
}
