/*
 * Copyright (c) 2022 by the IronCore authors.
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

package v1alpha1

// NetworkPolicyEgressRuleApplyConfiguration represents an declarative configuration of the NetworkPolicyEgressRule type for use
// with apply.
type NetworkPolicyEgressRuleApplyConfiguration struct {
	Ports []NetworkPolicyPortApplyConfiguration `json:"ports,omitempty"`
	To    []NetworkPolicyPeerApplyConfiguration `json:"to,omitempty"`
}

// NetworkPolicyEgressRuleApplyConfiguration constructs an declarative configuration of the NetworkPolicyEgressRule type for use with
// apply.
func NetworkPolicyEgressRule() *NetworkPolicyEgressRuleApplyConfiguration {
	return &NetworkPolicyEgressRuleApplyConfiguration{}
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *NetworkPolicyEgressRuleApplyConfiguration) WithPorts(values ...*NetworkPolicyPortApplyConfiguration) *NetworkPolicyEgressRuleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}

// WithTo adds the given value to the To field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the To field.
func (b *NetworkPolicyEgressRuleApplyConfiguration) WithTo(values ...*NetworkPolicyPeerApplyConfiguration) *NetworkPolicyEgressRuleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTo")
		}
		b.To = append(b.To, *values[i])
	}
	return b
}
