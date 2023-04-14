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
	v1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/core/v1alpha1"
)

// BucketPoolSpecApplyConfiguration represents an declarative configuration of the BucketPoolSpec type for use
// with apply.
type BucketPoolSpecApplyConfiguration struct {
	ProviderID *string                            `json:"providerID,omitempty"`
	Taints     []v1alpha1.TaintApplyConfiguration `json:"taints,omitempty"`
}

// BucketPoolSpecApplyConfiguration constructs an declarative configuration of the BucketPoolSpec type for use with
// apply.
func BucketPoolSpec() *BucketPoolSpecApplyConfiguration {
	return &BucketPoolSpecApplyConfiguration{}
}

// WithProviderID sets the ProviderID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProviderID field is set to the value of the last call.
func (b *BucketPoolSpecApplyConfiguration) WithProviderID(value string) *BucketPoolSpecApplyConfiguration {
	b.ProviderID = &value
	return b
}

// WithTaints adds the given value to the Taints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Taints field.
func (b *BucketPoolSpecApplyConfiguration) WithTaints(values ...*v1alpha1.TaintApplyConfiguration) *BucketPoolSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTaints")
		}
		b.Taints = append(b.Taints, *values[i])
	}
	return b
}
