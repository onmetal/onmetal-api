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

import (
	v1alpha1 "github.com/ironcore-dev/ironcore/api/core/v1alpha1"
)

// ResourceQuotaStatusApplyConfiguration represents an declarative configuration of the ResourceQuotaStatus type for use
// with apply.
type ResourceQuotaStatusApplyConfiguration struct {
	Hard *v1alpha1.ResourceList `json:"hard,omitempty"`
	Used *v1alpha1.ResourceList `json:"used,omitempty"`
}

// ResourceQuotaStatusApplyConfiguration constructs an declarative configuration of the ResourceQuotaStatus type for use with
// apply.
func ResourceQuotaStatus() *ResourceQuotaStatusApplyConfiguration {
	return &ResourceQuotaStatusApplyConfiguration{}
}

// WithHard sets the Hard field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hard field is set to the value of the last call.
func (b *ResourceQuotaStatusApplyConfiguration) WithHard(value v1alpha1.ResourceList) *ResourceQuotaStatusApplyConfiguration {
	b.Hard = &value
	return b
}

// WithUsed sets the Used field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Used field is set to the value of the last call.
func (b *ResourceQuotaStatusApplyConfiguration) WithUsed(value v1alpha1.ResourceList) *ResourceQuotaStatusApplyConfiguration {
	b.Used = &value
	return b
}
