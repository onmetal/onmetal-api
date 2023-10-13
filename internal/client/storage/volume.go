// Copyright 2023 OnMetal authors
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

package storage

import (
	"context"

	storagev1beta1 "github.com/onmetal/onmetal-api/api/storage/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	VolumeSpecVolumeClassRefNameField = storagev1beta1.VolumeVolumeClassRefNameField
	VolumeSpecVolumePoolRefNameField  = storagev1beta1.VolumeVolumePoolRefNameField
)

func SetupVolumeSpecVolumeClassRefNameFieldIndexer(ctx context.Context, indexer client.FieldIndexer) error {
	return indexer.IndexField(ctx, &storagev1beta1.Volume{}, VolumeSpecVolumeClassRefNameField, func(obj client.Object) []string {
		volume := obj.(*storagev1beta1.Volume)
		volumeClassRef := volume.Spec.VolumeClassRef
		if volumeClassRef == nil {
			return []string{""}
		}
		return []string{volumeClassRef.Name}
	})
}

func SetupVolumeSpecVolumePoolRefNameFieldIndexer(ctx context.Context, indexer client.FieldIndexer) error {
	return indexer.IndexField(ctx, &storagev1beta1.Volume{}, VolumeSpecVolumePoolRefNameField, func(obj client.Object) []string {
		volume := obj.(*storagev1beta1.Volume)
		volumePoolRef := volume.Spec.VolumePoolRef
		if volumePoolRef == nil {
			return []string{""}
		}
		return []string{volumePoolRef.Name}
	})
}
