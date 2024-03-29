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

// BucketAccessApplyConfiguration represents an declarative configuration of the BucketAccess type for use
// with apply.
type BucketAccessApplyConfiguration struct {
	SecretRef *v1.LocalObjectReference `json:"secretRef,omitempty"`
	Endpoint  *string                  `json:"endpoint,omitempty"`
}

// BucketAccessApplyConfiguration constructs an declarative configuration of the BucketAccess type for use with
// apply.
func BucketAccess() *BucketAccessApplyConfiguration {
	return &BucketAccessApplyConfiguration{}
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *BucketAccessApplyConfiguration) WithSecretRef(value v1.LocalObjectReference) *BucketAccessApplyConfiguration {
	b.SecretRef = &value
	return b
}

// WithEndpoint sets the Endpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Endpoint field is set to the value of the last call.
func (b *BucketAccessApplyConfiguration) WithEndpoint(value string) *BucketAccessApplyConfiguration {
	b.Endpoint = &value
	return b
}
