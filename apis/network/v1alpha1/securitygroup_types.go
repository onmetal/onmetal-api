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

package v1alpha1

import (
	common "github.com/onmetal/onmetal-api/apis/common/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	// Ingress is a list of inbound rules
	Ingress []IngressSecurityGroupRule `json:"ingress,omitempty"`
	// Egress is a list of outbound rules
	Egress []EgressSecurityGroupRule `json:"egress,omitempty"`
}

// IngressSecurityGroupRule is an ingress rule of a security group
type IngressSecurityGroupRule struct {
	SecurityGroupRule `json:",inline"`
	// Source is either the cird or a reference to another security group
	Source IPSetSpec `json:"source,omitempty"`
}

// EgressSecurityGroupRule is an egress rule of a security group
type EgressSecurityGroupRule struct {
	SecurityGroupRule `json:",inline"`
	// Destination is either the cird or a reference to another security group
	Destination IPSetSpec `json:"destination,omitempty"`
}

// SecurityGroupRule is a single access rule
type SecurityGroupRule struct {
	// Name is the name of the SecurityGroupRule
	Name string `json:"name"`
	// SecurityGroupRef is a scoped reference to an existing SecurityGroup
	SecurityGroupRef common.ScopeReference `json:"securitygroupref,omitempty"`
	// Action defines the action type of a SecurityGroupRule
	Action ActionType `json:"action,omitempty"`
	// Protocol defines the protocol of a SecurityGroupRule
	Protocol string `json:"protocol,omitempty"`
	// PortRange is the port range of the SecurityGroupRule
	PortRange PortRange `json:"portrange,omitempty"`
}

// IPSetSpec defines either a cidr or a security group reference
type IPSetSpec struct {
	// CIDR block for source/destination
	CIDR common.Cidr `json:"cidr,omitempty"`
	// SecurityGroupRef references a security group
	SecurityGroupRef common.ScopeReference `json:"securitygroupref,omitempty"`
}

// PortRange defines the start and end of a port range
type PortRange struct {
	// StartPort is the start port of the port range
	StartPort int `json:"startport"`
	// EndPort is the end port of the port range
	EndPort int `json:"endport"`
}

// ActionType describes the action type of a SecurityGroupRule
type ActionType string

// SecurityGroupStatus defines the observed state of SecurityGroup
type SecurityGroupStatus struct {
	common.StateFields `json:",inline"`
}

const (
	SecurityGroupActionTypeAllowed ActionType = "allowed"
	SecurityGroupActionTypeDeny    ActionType = "deny"
	SecurityGroupStateUsed                    = "Used"
	SecurityGroupStateUnused                  = "Unused"
	SecurityGroupStateInvalid                 = "Invalid"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="StateFields",type=string,JSONPath=`.status.state`
//+kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// SecurityGroup is the Schema for the securitygroups API
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityGroupSpec   `json:"spec,omitempty"`
	Status SecurityGroupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroup
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
