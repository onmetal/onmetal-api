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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MachineClassSpec defines the desired state of MachineClass
type MachineClassSpec struct {
	// Description is a short description of size constraint set
	// +kubebuilder:validation:Optional
	Description string `json:"description,omitempty"`
	// Capabilities describes the features of the MachineClass
	Capabilities []Capability `json:"capabilities"`
}

// Capability describes a single feature of a MachineClass
type Capability struct {
	// Name is the name of the capability
	Name string `json:"name"`
	// Type defines the type of the capability
	Type string `json:"type"`
	// Value is the effective value of the capability
	Value string `json:"value"`
}
type ConstraintValSpec struct {
	Literal *string            `json:"-"`
	Numeric *resource.Quantity `json:"-"`
}

type AggregateType string

type ConstraintSpec struct {
	// Path is a path to the struct field constraint will be applied to
	// +kubebuilder:validation:Optional
	Path string `json:"path,omitempty"`
	// Aggregate defines whether collection values should be aggregated
	// for constraint checks, in case if path defines selector for collection
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=avg;sum
	Aggregate AggregateType `json:"agg,omitempty"`
	// Equal contains an exact expected value
	// +kubebuilder:validation:Optional
	Equal *ConstraintValSpec `json:"eq,omitempty"`
	// NotEqual contains an exact not expected value
	// +kubebuilder:validation:Optional
	NotEqual *ConstraintValSpec `json:"neq,omitempty"`
	// LessThan contains an highest expected value, exclusive
	// +kubebuilder:validation:Optional
	LessThan *resource.Quantity `json:"lt,omitempty"`
	// LessThan contains an highest expected value, inclusive
	// +kubebuilder:validation:Optional
	LessThanOrEqual *resource.Quantity `json:"lte,omitempty"`
	// LessThan contains an lowest expected value, exclusive
	// +kubebuilder:validation:Optional
	GreaterThan *resource.Quantity `json:"gt,omitempty"`
	// GreaterThanOrEqual contains an lowest expected value, inclusive
	// +kubebuilder:validation:Optional
	GreaterThanOrEqual *resource.Quantity `json:"gte,omitempty"`
}

// MachineClassStatus defines the observed state of MachineClass
type MachineClassStatus struct {
	// Availability describes the regions and zones where this MachineClass is available
	Availability common.Availability `json:"availability,omitempty"`
	Constraints  []ConstraintSpec    `json:"constraints,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// MachineClass is the Schema for the machineclasses API
type MachineClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MachineClassSpec   `json:"spec,omitempty"`
	Status MachineClassStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MachineClassList contains a list of MachineClass
type MachineClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachineClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MachineClass{}, &MachineClassList{})
}
