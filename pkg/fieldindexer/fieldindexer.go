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

package fieldindexer

import (
	"context"
	storagev1alpha1 "github.com/onmetal/onmetal-api/apis/storage/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const (
	VolumeSpecVolumeClaimNameRefField = ".spec.claimRef.name"
	VolumeClaimSpecVolumeRefNameField = ".spec.volumeRef.name"
)

type FieldIndexer struct {
	manager.Manager
}

func NewIndexer(mgr manager.Manager) *FieldIndexer {
	return &FieldIndexer{mgr}
}

func (i *FieldIndexer) IndexFieldForVolumeClaim() error {
	return i.indexField(&storagev1alpha1.VolumeClaim{}, VolumeClaimSpecVolumeRefNameField, func(object client.Object) []string {
		claim := object.(*storagev1alpha1.VolumeClaim)
		if claim.Spec.VolumeRef.Name == "" {
			return nil
		}
		return []string{claim.Spec.VolumeRef.Name}
	})
}

func (i *FieldIndexer) IndexFieldForVolume() error {
	return i.indexField(&storagev1alpha1.Volume{}, VolumeSpecVolumeClaimNameRefField, func(object client.Object) []string {
		volume := object.(*storagev1alpha1.Volume)
		if volume.Spec.ClaimRef.Name == "" {
			return nil
		}
		return []string{volume.Spec.ClaimRef.Name}
	})
}

func (i *FieldIndexer) indexField(object client.Object, field string, indexFunc func(object client.Object) []string) error {
	ctx := context.Background()
	if err := i.Manager.GetFieldIndexer().IndexField(ctx, object, field, indexFunc); err != nil {
		return err
	}
	return nil
}
