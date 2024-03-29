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
	v1alpha1 "github.com/onmetal/onmetal-api/api/networking/v1alpha1"
)

// NetworkStatusApplyConfiguration represents an declarative configuration of the NetworkStatus type for use
// with apply.
type NetworkStatusApplyConfiguration struct {
	State    *v1alpha1.NetworkState                   `json:"state,omitempty"`
	Peerings []NetworkPeeringStatusApplyConfiguration `json:"peerings,omitempty"`
}

// NetworkStatusApplyConfiguration constructs an declarative configuration of the NetworkStatus type for use with
// apply.
func NetworkStatus() *NetworkStatusApplyConfiguration {
	return &NetworkStatusApplyConfiguration{}
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *NetworkStatusApplyConfiguration) WithState(value v1alpha1.NetworkState) *NetworkStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithPeerings adds the given value to the Peerings field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Peerings field.
func (b *NetworkStatusApplyConfiguration) WithPeerings(values ...*NetworkPeeringStatusApplyConfiguration) *NetworkStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPeerings")
		}
		b.Peerings = append(b.Peerings, *values[i])
	}
	return b
}
