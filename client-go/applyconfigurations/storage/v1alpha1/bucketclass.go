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
	v1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	internal "github.com/onmetal/onmetal-api/client-go/applyconfigurations/internal"
	v1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/meta/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// BucketClassApplyConfiguration represents an declarative configuration of the BucketClass type for use
// with apply.
type BucketClassApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Capabilities                     *corev1.ResourceList `json:"capabilities,omitempty"`
}

// BucketClass constructs an declarative configuration of the BucketClass type for use with
// apply.
func BucketClass(name string) *BucketClassApplyConfiguration {
	b := &BucketClassApplyConfiguration{}
	b.WithName(name)
	b.WithKind("BucketClass")
	b.WithAPIVersion("storage.api.onmetal.de/v1alpha1")
	return b
}

// ExtractBucketClass extracts the applied configuration owned by fieldManager from
// bucketClass. If no managedFields are found in bucketClass for fieldManager, a
// BucketClassApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// bucketClass must be a unmodified BucketClass API object that was retrieved from the Kubernetes API.
// ExtractBucketClass provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractBucketClass(bucketClass *v1alpha1.BucketClass, fieldManager string) (*BucketClassApplyConfiguration, error) {
	return extractBucketClass(bucketClass, fieldManager, "")
}

// ExtractBucketClassStatus is the same as ExtractBucketClass except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractBucketClassStatus(bucketClass *v1alpha1.BucketClass, fieldManager string) (*BucketClassApplyConfiguration, error) {
	return extractBucketClass(bucketClass, fieldManager, "status")
}

func extractBucketClass(bucketClass *v1alpha1.BucketClass, fieldManager string, subresource string) (*BucketClassApplyConfiguration, error) {
	b := &BucketClassApplyConfiguration{}
	err := managedfields.ExtractInto(bucketClass, internal.Parser().Type("com.github.onmetal.onmetal-api.api.storage.v1alpha1.BucketClass"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(bucketClass.Name)

	b.WithKind("BucketClass")
	b.WithAPIVersion("storage.api.onmetal.de/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithKind(value string) *BucketClassApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithAPIVersion(value string) *BucketClassApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithName(value string) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithGenerateName(value string) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithNamespace(value string) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithUID(value types.UID) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithResourceVersion(value string) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithGeneration(value int64) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithCreationTimestamp(value metav1.Time) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *BucketClassApplyConfiguration) WithLabels(entries map[string]string) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *BucketClassApplyConfiguration) WithAnnotations(entries map[string]string) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *BucketClassApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *BucketClassApplyConfiguration) WithFinalizers(values ...string) *BucketClassApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *BucketClassApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithCapabilities sets the Capabilities field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Capabilities field is set to the value of the last call.
func (b *BucketClassApplyConfiguration) WithCapabilities(value corev1.ResourceList) *BucketClassApplyConfiguration {
	b.Capabilities = &value
	return b
}