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

package server

import (
	"context"
	"fmt"

	corev1alpha1 "github.com/onmetal/onmetal-api/api/core/v1alpha1"
	storagev1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	ori "github.com/onmetal/onmetal-api/ori/apis/volume/v1alpha1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (s *Server) getTargetOnmetalVolumePools(ctx context.Context) ([]storagev1alpha1.VolumePool, error) {
	if s.volumePoolName != "" {
		onmetalVolumePool := &storagev1alpha1.VolumePool{}
		onmetalVolumePoolKey := client.ObjectKey{Name: s.volumePoolName}
		if err := s.client.Get(ctx, onmetalVolumePoolKey, onmetalVolumePool); err != nil {
			if !apierrors.IsNotFound(err) {
				return nil, fmt.Errorf("error getting volume pool %s: %w", s.volumePoolName, err)
			}
			return nil, nil
		}
	}

	volumePoolList := &storagev1alpha1.VolumePoolList{}
	if err := s.client.List(ctx, volumePoolList,
		client.MatchingLabels(s.volumePoolSelector),
	); err != nil {
		return nil, fmt.Errorf("error listing volume pools: %w", err)
	}
	return volumePoolList.Items, nil
}

func (s *Server) gatherAvailableVolumeClassNames(onmetalVolumePools []storagev1alpha1.VolumePool) sets.Set[string] {
	res := sets.New[string]()
	for _, onmetalVolumePool := range onmetalVolumePools {
		for _, availableVolumeClass := range onmetalVolumePool.Status.AvailableVolumeClasses {
			res.Insert(availableVolumeClass.Name)
		}
	}
	return res
}

func (s *Server) gatherVolumeClassQuantity(onmetalVolumePools []storagev1alpha1.VolumePool) map[string]*resource.Quantity {
	res := map[string]*resource.Quantity{}
	for _, onmetalVolumePool := range onmetalVolumePools {
		for resourceName, resourceQuantity := range onmetalVolumePool.Status.Capacity {
			if corev1alpha1.IsClassCountResource(resourceName) {
				if _, ok := res[string(resourceName)]; !ok {
					res[string(resourceName)] = resource.NewQuantity(0, resource.BinarySI)
				}
				res[string(resourceName)].Add(resourceQuantity)
			}
		}
	}
	return res
}

func (s *Server) filterOnmetalVolumeClasses(
	availableVolumeClassNames sets.Set[string],
	volumeClasses []storagev1alpha1.VolumeClass,
) []storagev1alpha1.VolumeClass {
	var filtered []storagev1alpha1.VolumeClass
	for _, volumeClass := range volumeClasses {
		if !availableVolumeClassNames.Has(volumeClass.Name) {
			continue
		}

		filtered = append(filtered, volumeClass)
	}
	return filtered
}

func (s *Server) convertOnmetalVolumeClassStatus(volumeClass *storagev1alpha1.VolumeClass, quantity *resource.Quantity) (*ori.VolumeClassStatus, error) {
	tps := volumeClass.Capabilities.TPS()
	iops := volumeClass.Capabilities.IOPS()

	return &ori.VolumeClassStatus{
		VolumeClass: &ori.VolumeClass{
			Name: volumeClass.Name,
			Capabilities: &ori.VolumeClassCapabilities{
				Tps:  tps.Value(),
				Iops: iops.Value(),
			},
		},
		Quantity: quantity.Value(),
	}, nil
}

func (s *Server) Status(ctx context.Context, req *ori.StatusRequest) (*ori.StatusResponse, error) {
	log := s.loggerFrom(ctx)

	log.V(1).Info("Getting target onmetal volume pools")
	onmetalVolumePools, err := s.getTargetOnmetalVolumePools(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting target onmetal volume pools: %w", err)
	}

	log.V(1).Info("Gathering available volume class names")
	availableOnmetalVolumeClassNames := s.gatherAvailableVolumeClassNames(onmetalVolumePools)

	if len(availableOnmetalVolumeClassNames) == 0 {
		log.V(1).Info("No available volume classes")
		return &ori.StatusResponse{VolumeClassStatus: []*ori.VolumeClassStatus{}}, nil
	}

	log.V(1).Info("Gathering volume class quantity")
	volumeClassQuantity := s.gatherVolumeClassQuantity(onmetalVolumePools)

	log.V(1).Info("Listing onmetal volume classes")
	onmetalVolumeClassList := &storagev1alpha1.VolumeClassList{}
	if err := s.client.List(ctx, onmetalVolumeClassList); err != nil {
		return nil, fmt.Errorf("error listing onmetal volume classes: %w", err)
	}

	availableOnmetalVolumeClasses := s.filterOnmetalVolumeClasses(availableOnmetalVolumeClassNames, onmetalVolumeClassList.Items)
	volumeClassStatus := make([]*ori.VolumeClassStatus, 0, len(availableOnmetalVolumeClasses))
	for _, onmetalVolumeClass := range availableOnmetalVolumeClasses {
		quantity, ok := volumeClassQuantity[string(corev1alpha1.ClassCountFor(corev1alpha1.ClassTypeVolumeClass, onmetalVolumeClass.Name))]
		if !ok {
			log.V(1).Info("Ignored class - missing quantity", "VolumeClass", onmetalVolumeClass.Name)
			continue
		}

		volumeClass, err := s.convertOnmetalVolumeClassStatus(&onmetalVolumeClass, quantity)
		if err != nil {
			return nil, fmt.Errorf("error converting onmetal volume class %s: %w", onmetalVolumeClass.Name, err)
		}

		volumeClassStatus = append(volumeClassStatus, volumeClass)
	}

	log.V(1).Info("Returning volume classes")
	return &ori.StatusResponse{
		VolumeClassStatus: volumeClassStatus,
	}, nil
}
