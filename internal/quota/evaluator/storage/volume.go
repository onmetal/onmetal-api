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
	"fmt"

	corev1beta1 "github.com/onmetal/onmetal-api/api/core/v1beta1"
	storagev1beta1 "github.com/onmetal/onmetal-api/api/storage/v1beta1"
	"github.com/onmetal/onmetal-api/internal/apis/storage"
	internalstoragev1beta1 "github.com/onmetal/onmetal-api/internal/apis/storage/v1beta1"
	"github.com/onmetal/onmetal-api/internal/quota/evaluator/generic"
	"github.com/onmetal/onmetal-api/utils/quota"
	"golang.org/x/exp/slices"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	volumeResource          = storagev1beta1.Resource("volumes")
	volumeCountResourceName = corev1beta1.ObjectCountQuotaResourceNameFor(volumeResource)

	VolumeResourceNames = sets.New(
		volumeCountResourceName,
		corev1beta1.ResourceRequestsStorage,
	)
)

type volumeEvaluator struct {
	capabilities generic.CapabilitiesReader
}

func NewVolumeEvaluator(capabilities generic.CapabilitiesReader) quota.Evaluator {
	return &volumeEvaluator{
		capabilities: capabilities,
	}
}

func (m *volumeEvaluator) Type() client.Object {
	return &storagev1beta1.Volume{}
}

func (m *volumeEvaluator) MatchesResourceName(name corev1beta1.ResourceName) bool {
	return VolumeResourceNames.Has(name)
}

func (m *volumeEvaluator) MatchesResourceScopeSelectorRequirement(item client.Object, req corev1beta1.ResourceScopeSelectorRequirement) (bool, error) {
	volume := item.(*storagev1beta1.Volume)

	switch req.ScopeName {
	case corev1beta1.ResourceScopeVolumeClass:
		return volumeMatchesVolumeClassScope(volume, req.Operator, req.Values), nil
	default:
		return false, nil
	}
}

func volumeMatchesVolumeClassScope(volume *storagev1beta1.Volume, op corev1beta1.ResourceScopeSelectorOperator, values []string) bool {
	volumeClassRef := volume.Spec.VolumeClassRef

	switch op {
	case corev1beta1.ResourceScopeSelectorOperatorExists:
		return volumeClassRef != nil
	case corev1beta1.ResourceScopeSelectorOperatorDoesNotExist:
		return volumeClassRef == nil
	case corev1beta1.ResourceScopeSelectorOperatorIn:
		return slices.Contains(values, volumeClassRef.Name)
	case corev1beta1.ResourceScopeSelectorOperatorNotIn:
		if volumeClassRef == nil {
			return false
		}
		return !slices.Contains(values, volumeClassRef.Name)
	default:
		return false
	}
}

func toExternalVolumeOrError(obj client.Object) (*storagev1beta1.Volume, error) {
	switch t := obj.(type) {
	case *storagev1beta1.Volume:
		return t, nil
	case *storage.Volume:
		volume := &storagev1beta1.Volume{}
		if err := internalstoragev1beta1.Convert_storage_Volume_To_v1beta1_Volume(t, volume, nil); err != nil {
			return nil, err
		}
		return volume, nil
	default:
		return nil, fmt.Errorf("expect *storage.Volume or *storagev1beta1.Volume but got %v", t)
	}
}

func (m *volumeEvaluator) Usage(ctx context.Context, item client.Object) (corev1beta1.ResourceList, error) {
	volume, err := toExternalVolumeOrError(item)
	if err != nil {
		return nil, err
	}

	return corev1beta1.ResourceList{
		volumeCountResourceName:             resource.MustParse("1"),
		corev1beta1.ResourceRequestsStorage: *volume.Spec.Resources.Storage(),
	}, nil
}
