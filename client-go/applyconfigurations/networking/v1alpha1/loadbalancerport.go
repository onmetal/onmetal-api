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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// LoadBalancerPortApplyConfiguration represents an declarative configuration of the LoadBalancerPort type for use
// with apply.
type LoadBalancerPortApplyConfiguration struct {
	Protocol *v1.Protocol `json:"protocol,omitempty"`
	Port     *int32       `json:"port,omitempty"`
	EndPort  *int32       `json:"endPort,omitempty"`
}

// LoadBalancerPortApplyConfiguration constructs an declarative configuration of the LoadBalancerPort type for use with
// apply.
func LoadBalancerPort() *LoadBalancerPortApplyConfiguration {
	return &LoadBalancerPortApplyConfiguration{}
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *LoadBalancerPortApplyConfiguration) WithProtocol(value v1.Protocol) *LoadBalancerPortApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *LoadBalancerPortApplyConfiguration) WithPort(value int32) *LoadBalancerPortApplyConfiguration {
	b.Port = &value
	return b
}

// WithEndPort sets the EndPort field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndPort field is set to the value of the last call.
func (b *LoadBalancerPortApplyConfiguration) WithEndPort(value int32) *LoadBalancerPortApplyConfiguration {
	b.EndPort = &value
	return b
}
