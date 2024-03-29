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

	"github.com/go-logr/logr"
	storagev1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	ori "github.com/onmetal/onmetal-api/ori/apis/bucket/v1alpha1"
	"github.com/onmetal/onmetal-api/poollet/bucketpoollet/bcm"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type BucketPoolReconciler struct {
	client.Client
	BucketPoolName    string
	BucketRuntime     ori.BucketRuntimeClient
	BucketClassMapper bcm.BucketClassMapper
}

//+kubebuilder:rbac:groups=storage.api.onmetal.de,resources=bucketpools,verbs=get;list;watch;update;patch
//+kubebuilder:rbac:groups=storage.api.onmetal.de,resources=bucketpools/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=storage.api.onmetal.de,resources=bucketclasses,verbs=get;list;watch

func (r *BucketPoolReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx)
	bucketPool := &storagev1alpha1.BucketPool{}
	if err := r.Get(ctx, req.NamespacedName, bucketPool); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return r.reconcileExists(ctx, log, bucketPool)
}

func (r *BucketPoolReconciler) reconcileExists(ctx context.Context, log logr.Logger, bucketPool *storagev1alpha1.BucketPool) (ctrl.Result, error) {
	if !bucketPool.DeletionTimestamp.IsZero() {
		return r.delete(ctx, log, bucketPool)
	}
	return r.reconcile(ctx, log, bucketPool)
}

func (r *BucketPoolReconciler) delete(ctx context.Context, log logr.Logger, bucketPool *storagev1alpha1.BucketPool) (ctrl.Result, error) {
	log.V(1).Info("Delete")
	log.V(1).Info("Deleted")
	return ctrl.Result{}, nil
}

func (r *BucketPoolReconciler) supportsBucketClass(ctx context.Context, log logr.Logger, bucketClass *storagev1alpha1.BucketClass) (bool, error) {
	oriCapabilities, err := getORIBucketClassCapabilities(bucketClass)
	if err != nil {
		return false, fmt.Errorf("error getting ori mahchine class capabilities: %w", err)
	}

	_, err = r.BucketClassMapper.GetBucketClassFor(ctx, bucketClass.Name, oriCapabilities)
	if err != nil {
		if !errors.Is(err, bcm.ErrNoMatchingBucketClass) && !errors.Is(err, bcm.ErrAmbiguousMatchingBucketClass) {
			return false, fmt.Errorf("error getting bucket class for %s: %w", bucketClass.Name, err)
		}
		return false, nil
	}
	return true, nil
}

func (r *BucketPoolReconciler) reconcile(ctx context.Context, log logr.Logger, bucketPool *storagev1alpha1.BucketPool) (ctrl.Result, error) {
	log.V(1).Info("Reconcile")

	log.V(1).Info("Listing bucket classes")
	bucketClassList := &storagev1alpha1.BucketClassList{}
	if err := r.List(ctx, bucketClassList); err != nil {
		return ctrl.Result{}, fmt.Errorf("error listing bucket classes: %w", err)
	}

	log.V(1).Info("Determining supported bucket classes")
	var supported []corev1.LocalObjectReference
	for _, bucketClass := range bucketClassList.Items {
		ok, err := r.supportsBucketClass(ctx, log, &bucketClass)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error checking whether bucket class %s is supported: %w", bucketClass.Name, err)
		}
		if !ok {
			continue
		}

		supported = append(supported, corev1.LocalObjectReference{Name: bucketClass.Name})
	}

	log.V(1).Info("Updating bucket pool status")
	base := bucketPool.DeepCopy()
	bucketPool.Status.AvailableBucketClasses = supported
	if err := r.Status().Patch(ctx, bucketPool, client.MergeFrom(base)); err != nil {
		return ctrl.Result{}, fmt.Errorf("error patchign bucket pool status: %w", err)
	}

	log.V(1).Info("Reconciled")
	return ctrl.Result{}, nil
}

func (r *BucketPoolReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(
			&storagev1alpha1.BucketPool{},
			builder.WithPredicates(
				predicate.NewPredicateFuncs(func(obj client.Object) bool {
					return obj.GetName() == r.BucketPoolName
				}),
			),
		).
		Watches(
			&storagev1alpha1.BucketClass{},
			handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []ctrl.Request {
				return []ctrl.Request{{NamespacedName: client.ObjectKey{Name: r.BucketPoolName}}}
			}),
		).
		Complete(r)
}
