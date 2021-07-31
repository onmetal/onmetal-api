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

//+kubebuilder:object:generate=true

type Availability []RegionAvailability

//+kubebuilder:object:generate=true

// RegionAvailability defines a region with its availability zones
type RegionAvailability struct {
	// Region is the name of the region
	Region string `json:"region"`
	// Zones is a list of zones in this region
	Zones []ZoneAvailability `json:"availabilityZone"`
}

//+kubebuilder:object:generate=true

// Location describes the location of a resource
type Location struct {
	// Region defines the region of a resource
	Region string `json:"region"`
	// AvailabilityZone is the availability zone of a resource
	AvailabilityZone string `json:"availabilityZone"`
}

//+kubebuilder:object:generate=true

// ZoneAvailability defines the name of a zone
type ZoneAvailability struct {
	// Name is the name of the availability zone
	Name string `json:"name"`
}

//+kubebuilder:object:generate=true

// ScopeReference refers to a scope and the scopes name
type ScopeReference struct {
	// Name is the name of the scope
	Name string `json:"name"`
	// Scope is the absolute scope path
	Scope string `json:"scope"`
}

// KindReference defines an object with its kind and API group and its scope reference
type KindReference struct {
	// Kind is the kind of the object
	Kind string `json:"kind"`
	// APIGroup is the API group of the object
	APIGroup       string `json:"apigroup"`
	ScopeReference `json:",inline"`
}

// TODO: create marshal/unmarshal functions
type IPAddr string

// TODO: create marshal/unmarshal functions
type Cidr string
