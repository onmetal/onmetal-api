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
	"fmt"

	"github.com/go-logr/logr"
	computev1alpha1 "github.com/onmetal/onmetal-api/api/compute/v1alpha1"
	"github.com/onmetal/onmetal-api/poollet/machinepoollet/mcm"
	"github.com/onmetal/onmetal-api/poollet/orievent"
	onmetalapiclient "github.com/onmetal/onmetal-api/utils/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

type MachinePoolAnnotatorReconciler struct {
	client.Client

	MachinePoolName    string
	MachineClassMapper mcm.MachineClassMapper
}

func (r *MachinePoolAnnotatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	machinePool := &computev1alpha1.MachinePool{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
	}

	if err := onmetalapiclient.PatchAddReconcileAnnotation(ctx, r.Client, machinePool); client.IgnoreNotFound(err) != nil {
		return ctrl.Result{}, fmt.Errorf("error patching machine pool: %w", err)
	}
	return ctrl.Result{}, nil
}

func (r *MachinePoolAnnotatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	c, err := controller.New("machinepoolannotator", mgr, controller.Options{
		Reconciler: r,
	})
	if err != nil {
		return err
	}

	src, err := r.oriClassEventSource(mgr)
	if err != nil {
		return err
	}

	if err := c.Watch(src, handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []ctrl.Request {
		return []ctrl.Request{{NamespacedName: client.ObjectKey{Name: r.MachinePoolName}}}
	})); err != nil {
		return err
	}

	return nil
}

func (r *MachinePoolAnnotatorReconciler) machinePoolAnnotatorEventHandler(log logr.Logger, c chan<- event.GenericEvent) orievent.EnqueueFunc {
	handleEvent := func() {
		select {
		case c <- event.GenericEvent{Object: &computev1alpha1.MachinePool{ObjectMeta: metav1.ObjectMeta{
			Name: r.MachinePoolName,
		}}}:
			log.V(1).Info("Added item to queue")
		default:
			log.V(5).Info("Channel full, discarding event")
		}
	}

	return orievent.EnqueueFunc{EnqueueFunc: handleEvent}
}

func (r *MachinePoolAnnotatorReconciler) oriClassEventSource(mgr ctrl.Manager) (source.Source, error) {
	ch := make(chan event.GenericEvent, 1024)

	if err := mgr.Add(manager.RunnableFunc(func(ctx context.Context) error {
		log := ctrl.LoggerFrom(ctx).WithName("machinepool").WithName("orieventhandlers")

		notifierFuncs := []func() (orievent.ListenerRegistration, error){
			func() (orievent.ListenerRegistration, error) {
				return r.MachineClassMapper.AddListener(r.machinePoolAnnotatorEventHandler(log, ch))
			},
		}

		var notifier []orievent.ListenerRegistration
		defer func() {
			log.V(1).Info("Removing notifier")
			for _, n := range notifier {
				if err := r.MachineClassMapper.RemoveListener(n); err != nil {
					log.Error(err, "Error removing handle")
				}
			}
		}()

		for _, notifierFunc := range notifierFuncs {
			ntf, err := notifierFunc()
			if err != nil {
				return err
			}

			notifier = append(notifier, ntf)
		}

		<-ctx.Done()
		return nil
	})); err != nil {
		return nil, err
	}

	return &source.Channel{Source: ch}, nil
}
