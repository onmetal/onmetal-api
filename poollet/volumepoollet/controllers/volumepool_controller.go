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

package controllers

import (
	"context"
	"errors"
	"fmt"

	corev1alpha1 "github.com/onmetal/onmetal-api/api/core/v1alpha1"
	storageclient "github.com/onmetal/onmetal-api/internal/client/storage"
	"github.com/onmetal/onmetal-api/utils/quota"
	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/go-logr/logr"
	storagev1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	ori "github.com/onmetal/onmetal-api/ori/apis/volume/v1alpha1"
	"github.com/onmetal/onmetal-api/poollet/volumepoollet/vcm"
	onmetalapiclient "github.com/onmetal/onmetal-api/utils/client"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type VolumePoolReconciler struct {
	client.Client
	VolumePoolName    string
	VolumeRuntime     ori.VolumeRuntimeClient
	VolumeClassMapper vcm.VolumeClassMapper
}

//+kubebuilder:rbac:groups=storage.api.onmetal.de,resources=volumepools,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=storage.api.onmetal.de,resources=volumepools/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=storage.api.onmetal.de,resources=volumeclasses,verbs=get;list;watch

func (r *VolumePoolReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx)
	volumePool := &storagev1alpha1.VolumePool{}
	if err := r.Get(ctx, req.NamespacedName, volumePool); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return r.reconcileExists(ctx, log, volumePool)
}

func (r *VolumePoolReconciler) reconcileExists(ctx context.Context, log logr.Logger, volumePool *storagev1alpha1.VolumePool) (ctrl.Result, error) {
	if !volumePool.DeletionTimestamp.IsZero() {
		return r.delete(ctx, log, volumePool)
	}
	return r.reconcile(ctx, log, volumePool)
}

func (r *VolumePoolReconciler) delete(ctx context.Context, log logr.Logger, volumePool *storagev1alpha1.VolumePool) (ctrl.Result, error) {
	log.V(1).Info("Delete")
	log.V(1).Info("Deleted")
	return ctrl.Result{}, nil
}

func (r *VolumePoolReconciler) supportsVolumeClass(ctx context.Context, log logr.Logger, volumeClass *storagev1alpha1.VolumeClass) (*ori.VolumeClass, *resource.Quantity, error) {
	oriCapabilities, err := getORIVolumeClassCapabilities(volumeClass)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting ori mahchine class capabilities: %w", err)
	}

	class, quantity, err := r.VolumeClassMapper.GetVolumeClassFor(ctx, volumeClass.Name, oriCapabilities)
	if err != nil {
		if !errors.Is(err, vcm.ErrNoMatchingVolumeClass) && !errors.Is(err, vcm.ErrAmbiguousMatchingVolumeClass) {
			return nil, nil, fmt.Errorf("error getting volume class for %s: %w", volumeClass.Name, err)
		}
		return nil, nil, nil
	}
	return class, quantity, nil
}

func (r *VolumePoolReconciler) calculateCapacity(
	ctx context.Context,
	log logr.Logger,
	volumes []storagev1alpha1.Volume,
	volumeClassList []storagev1alpha1.VolumeClass,
) (capacity, allocatable corev1alpha1.ResourceList, supported []corev1.LocalObjectReference, err error) {
	log.V(1).Info("Determining supported volume classes, capacity and allocatable")

	capacity = corev1alpha1.ResourceList{}
	for _, volumeClass := range volumeClassList {
		class, quantity, err := r.supportsVolumeClass(ctx, log, &volumeClass)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error checking whether volume class %s is supported: %w", volumeClass.Name, err)
		}
		if class == nil {
			continue
		}

		supported = append(supported, corev1.LocalObjectReference{Name: volumeClass.Name})
		capacity[corev1alpha1.ClassCountFor(corev1alpha1.ClassTypeVolumeClass, volumeClass.Name)] = *quantity
	}

	usedResources := corev1alpha1.ResourceList{}
	for _, volume := range volumes {
		className := volume.Spec.VolumeClassRef.Name
		res, ok := usedResources[corev1alpha1.ClassCountFor(corev1alpha1.ClassTypeVolumeClass, className)]
		if !ok {
			usedResources[corev1alpha1.ClassCountFor(corev1alpha1.ClassTypeVolumeClass, className)] = *volume.Spec.Resources.Storage()
			continue
		}

		res.Add(*volume.Spec.Resources.Storage())
	}

	return capacity, quota.SubtractWithNonNegativeResult(capacity, usedResources), supported, nil
}

func (r *VolumePoolReconciler) updateStatus(ctx context.Context, log logr.Logger, volumePool *storagev1alpha1.VolumePool, volumes []storagev1alpha1.Volume, volumeClassList []storagev1alpha1.VolumeClass) error {
	capacity, allocatable, supported, err := r.calculateCapacity(ctx, log, volumes, volumeClassList)
	if err != nil {
		return fmt.Errorf("error calculating pool resources:%w", err)
	}

	base := volumePool.DeepCopy()
	volumePool.Status.State = storagev1alpha1.VolumePoolStateAvailable
	volumePool.Status.AvailableVolumeClasses = supported
	volumePool.Status.Capacity = capacity
	volumePool.Status.Allocatable = allocatable

	if err := r.Status().Patch(ctx, volumePool, client.MergeFrom(base)); err != nil {
		return fmt.Errorf("error patching volume pool status: %w", err)
	}

	return nil
}

func (r *VolumePoolReconciler) reconcile(ctx context.Context, log logr.Logger, volumePool *storagev1alpha1.VolumePool) (ctrl.Result, error) {
	log.V(1).Info("Reconcile")

	log.V(1).Info("Ensuring no reconcile annotation")
	modified, err := onmetalapiclient.PatchEnsureNoReconcileAnnotation(ctx, r.Client, volumePool)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("error ensuring no reconcile annotation: %w", err)
	}
	if modified {
		log.V(1).Info("Removed reconcile annotation, requeueing")
		return ctrl.Result{Requeue: true}, nil
	}

	log.V(1).Info("Listing volume classes")
	volumeClassList := &storagev1alpha1.VolumeClassList{}
	if err := r.List(ctx, volumeClassList); err != nil {
		return ctrl.Result{}, fmt.Errorf("error listing volume classes: %w", err)
	}

	log.V(1).Info("Listing volumes in pool")
	volumeList := &storagev1alpha1.VolumeList{}
	if err := r.List(ctx, volumeList, client.MatchingFields{
		storageclient.VolumeSpecVolumePoolRefNameField: r.VolumePoolName,
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to list volumes in pool: %w", err)
	}

	log.V(1).Info("Updating volume pool status")
	if err := r.updateStatus(ctx, log, volumePool, volumeList.Items, volumeClassList.Items); err != nil {
		return ctrl.Result{}, fmt.Errorf("error updating status: %w", err)
	}

	log.V(1).Info("Reconciled")
	return ctrl.Result{}, nil
}

func (r *VolumePoolReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(
			&storagev1alpha1.VolumePool{},
			builder.WithPredicates(
				predicate.NewPredicateFuncs(func(obj client.Object) bool {
					return obj.GetName() == r.VolumePoolName
				}),
			),
		).
		Watches(
			&storagev1alpha1.VolumeClass{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []ctrl.Request {
				return []ctrl.Request{{NamespacedName: client.ObjectKey{Name: r.VolumePoolName}}}
			}),
		).
		Complete(r)
}
