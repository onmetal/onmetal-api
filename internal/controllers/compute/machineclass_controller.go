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

package compute

import (
	"context"
	"fmt"
	"sort"

	"github.com/go-logr/logr"
	"github.com/onmetal/controller-utils/clientutils"
	computev1beta1 "github.com/onmetal/onmetal-api/api/compute/v1beta1"
	computeclient "github.com/onmetal/onmetal-api/internal/client/compute"
	"github.com/onmetal/onmetal-api/utils/slices"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
)

// MachineClassReconciler reconciles a MachineClassRef object
type MachineClassReconciler struct {
	client.Client
	APIReader client.Reader
}

//+kubebuilder:rbac:groups=compute.api.onmetal.de,resources=machineclasses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=compute.api.onmetal.de,resources=machineclasses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=compute.api.onmetal.de,resources=machineclasses/finalizers,verbs=update
//+kubebuilder:rbac:groups=compute.api.onmetal.de,resources=machines,verbs=get;list;watch

// Reconcile moves the current state of the cluster closer to the desired state
func (r *MachineClassReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx)
	machineClass := &computev1beta1.MachineClass{}
	if err := r.Get(ctx, req.NamespacedName, machineClass); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return r.reconcileExists(ctx, log, machineClass)
}

func (r *MachineClassReconciler) listReferencingMachinesWithReader(
	ctx context.Context,
	rd client.Reader,
	machineClass *computev1beta1.MachineClass,
) ([]computev1beta1.Machine, error) {
	machineList := &computev1beta1.MachineList{}
	if err := rd.List(ctx, machineList,
		client.InNamespace(machineClass.Namespace),
		client.MatchingFields{computeclient.MachineSpecMachineClassRefNameField: machineClass.Name},
	); err != nil {
		return nil, fmt.Errorf("error listing the machines using the machine class: %w", err)
	}

	return machineList.Items, nil
}

func (r *MachineClassReconciler) collectMachineNames(machines []computev1beta1.Machine) []string {
	machineNames := slices.MapRef(machines, func(machine *computev1beta1.Machine) string {
		return machine.Name
	})
	sort.Strings(machineNames)
	return machineNames
}

func (r *MachineClassReconciler) delete(ctx context.Context, log logr.Logger, machineClass *computev1beta1.MachineClass) (ctrl.Result, error) {
	if !controllerutil.ContainsFinalizer(machineClass, computev1beta1.MachineClassFinalizer) {
		return ctrl.Result{}, nil
	}

	machines, err := r.listReferencingMachinesWithReader(ctx, r.Client, machineClass)
	if err != nil {
		return ctrl.Result{}, err
	}
	if len(machines) > 0 {
		log.V(1).Info("Machine class is still in use", "ReferencingMachineNames", r.collectMachineNames(machines))
		return ctrl.Result{Requeue: true}, nil
	}

	machines, err = r.listReferencingMachinesWithReader(ctx, r.APIReader, machineClass)
	if err != nil {
		return ctrl.Result{}, err
	}
	if len(machines) > 0 {
		log.V(1).Info("Machine class is still in use", "ReferencingMachineNames", r.collectMachineNames(machines))
		return ctrl.Result{Requeue: true}, nil
	}

	log.V(1).Info("Machine class is not in use anymore, removing finalizer")
	if err := clientutils.PatchRemoveFinalizer(ctx, r.Client, machineClass, computev1beta1.MachineClassFinalizer); err != nil {
		return ctrl.Result{}, err
	}

	log.V(1).Info("Successfully removed finalizer")
	return ctrl.Result{}, nil
}

func (r *MachineClassReconciler) reconcile(ctx context.Context, log logr.Logger, machineClass *computev1beta1.MachineClass) (ctrl.Result, error) {
	log.V(1).Info("Ensuring finalizer")
	if modified, err := clientutils.PatchEnsureFinalizer(ctx, r.Client, machineClass, computev1beta1.MachineClassFinalizer); err != nil || modified {
		return ctrl.Result{}, err
	}

	log.V(1).Info("Finalizer is present")
	return ctrl.Result{}, nil
}

func (r *MachineClassReconciler) reconcileExists(ctx context.Context, log logr.Logger, machineClass *computev1beta1.MachineClass) (ctrl.Result, error) {
	if !machineClass.DeletionTimestamp.IsZero() {
		return r.delete(ctx, log, machineClass)
	}
	return r.reconcile(ctx, log, machineClass)
}

// SetupWithManager sets up the controller with the Manager.
func (r *MachineClassReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&computev1beta1.MachineClass{}).
		Watches(
			&computev1beta1.Machine{},
			handler.Funcs{
				DeleteFunc: func(ctx context.Context, event event.DeleteEvent, queue workqueue.RateLimitingInterface) {
					machine := event.Object.(*computev1beta1.Machine)
					queue.Add(ctrl.Request{NamespacedName: types.NamespacedName{Name: machine.Spec.MachineClassRef.Name}})
				},
			},
		).
		Complete(r)
}
