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
	v1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/common/v1alpha1"
	v1 "k8s.io/api/core/v1"
)

// BucketSpecApplyConfiguration represents an declarative configuration of the BucketSpec type for use
// with apply.
type BucketSpecApplyConfiguration struct {
	BucketClassRef     *v1.LocalObjectReference                `json:"bucketClassRef,omitempty"`
	BucketPoolSelector map[string]string                       `json:"bucketPoolSelector,omitempty"`
	BucketPoolRef      *v1.LocalObjectReference                `json:"bucketPoolRef,omitempty"`
	Tolerations        []v1alpha1.TolerationApplyConfiguration `json:"tolerations,omitempty"`
}

// BucketSpecApplyConfiguration constructs an declarative configuration of the BucketSpec type for use with
// apply.
func BucketSpec() *BucketSpecApplyConfiguration {
	return &BucketSpecApplyConfiguration{}
}

// WithBucketClassRef sets the BucketClassRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BucketClassRef field is set to the value of the last call.
func (b *BucketSpecApplyConfiguration) WithBucketClassRef(value v1.LocalObjectReference) *BucketSpecApplyConfiguration {
	b.BucketClassRef = &value
	return b
}

// WithBucketPoolSelector puts the entries into the BucketPoolSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the BucketPoolSelector field,
// overwriting an existing map entries in BucketPoolSelector field with the same key.
func (b *BucketSpecApplyConfiguration) WithBucketPoolSelector(entries map[string]string) *BucketSpecApplyConfiguration {
	if b.BucketPoolSelector == nil && len(entries) > 0 {
		b.BucketPoolSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.BucketPoolSelector[k] = v
	}
	return b
}

// WithBucketPoolRef sets the BucketPoolRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BucketPoolRef field is set to the value of the last call.
func (b *BucketSpecApplyConfiguration) WithBucketPoolRef(value v1.LocalObjectReference) *BucketSpecApplyConfiguration {
	b.BucketPoolRef = &value
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *BucketSpecApplyConfiguration) WithTolerations(values ...*v1alpha1.TolerationApplyConfiguration) *BucketSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTolerations")
		}
		b.Tolerations = append(b.Tolerations, *values[i])
	}
	return b
}
