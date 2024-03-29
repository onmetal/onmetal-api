// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import corev1 "k8s.io/api/core/v1"

const (
	VolumeVolumePoolRefNameField  = "spec.volumePoolRef.name"
	VolumeVolumeClassRefNameField = "spec.volumeClassRef.name"

	BucketBucketPoolRefNameField  = "spec.bucketPoolRef.name"
	BucketBucketClassRefNameField = "spec.bucketClassRef.name"

	// VolumePoolsGroup is the system rbac group all volume pools are in.
	VolumePoolsGroup = "storage.api.onmetal.de:system:volumepools"

	// VolumePoolUserNamePrefix is the prefix all volume pool users should have.
	VolumePoolUserNamePrefix = "storage.api.onmetal.de:system:volumepool:"

	// BucketPoolsGroup is the system rbac group all bucket pools are in.
	BucketPoolsGroup = "storage.api.onmetal.de:system:bucketpools"

	// BucketPoolUserNamePrefix is the prefix all bucket pool users should have.
	BucketPoolUserNamePrefix = "storage.api.onmetal.de:system:bucketpool:"

	SecretTypeVolumeAuth = corev1.SecretType("storage.api.onmetal.de/volume-auth")
)

// VolumePoolCommonName constructs the common name for a certificate of a volume pool user.
func VolumePoolCommonName(name string) string {
	return VolumePoolUserNamePrefix + name
}

// BucketPoolCommonName constructs the common name for a certificate of a bucket pool user.
func BucketPoolCommonName(name string) string {
	return BucketPoolUserNamePrefix + name
}
