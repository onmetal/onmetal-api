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

	"github.com/go-logr/logr"
	commonv1beta1 "github.com/onmetal/onmetal-api/api/common/v1beta1"
	computev1beta1 "github.com/onmetal/onmetal-api/api/compute/v1beta1"
	storagev1beta1 "github.com/onmetal/onmetal-api/api/storage/v1beta1"
	"github.com/onmetal/onmetal-api/broker/common/cleaner"
	machinebrokerv1beta1 "github.com/onmetal/onmetal-api/broker/machinebroker/api/v1beta1"
	ori "github.com/onmetal/onmetal-api/ori/apis/machine/v1beta1"
	metautils "github.com/onmetal/onmetal-api/utils/meta"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type OnmetalVolumeConfig struct {
	Name      string
	Device    string
	EmptyDisk *OnmetalVolumeEmptyDiskConfig
	Remote    *OnmetalVolumeRemoteConfig
}

type OnmetalVolumeEmptyDiskConfig struct {
	SizeLimit *resource.Quantity
}

type OnmetalVolumeRemoteConfig struct {
	Driver     string
	Handle     string
	Attributes map[string]string
	SecretData map[string][]byte
}

func (s *Server) getOnmetalVolumeConfig(volume *ori.Volume) (*OnmetalVolumeConfig, error) {
	var (
		emptyDisk *OnmetalVolumeEmptyDiskConfig
		remote    *OnmetalVolumeRemoteConfig
	)
	switch {
	case volume.EmptyDisk != nil:
		var sizeLimit *resource.Quantity
		if sizeBytes := volume.EmptyDisk.SizeBytes; sizeBytes > 0 {
			sizeLimit = resource.NewQuantity(int64(sizeBytes), resource.DecimalSI)
		}
		emptyDisk = &OnmetalVolumeEmptyDiskConfig{
			SizeLimit: sizeLimit,
		}
	case volume.Connection != nil:
		remote = &OnmetalVolumeRemoteConfig{
			Driver:     volume.Connection.Driver,
			Handle:     volume.Connection.Handle,
			Attributes: volume.Connection.Attributes,
			SecretData: volume.Connection.SecretData,
		}
	default:
		return nil, fmt.Errorf("unrecognized volume %#v", volume)
	}

	return &OnmetalVolumeConfig{
		Name:      volume.Name,
		Device:    volume.Device,
		EmptyDisk: emptyDisk,
		Remote:    remote,
	}, nil
}

func (s *Server) createOnmetalVolume(
	ctx context.Context,
	log logr.Logger,
	c *cleaner.Cleaner,
	optOnmetalMachine client.Object,
	cfg *OnmetalVolumeConfig,
) (onmetalMachineVolume *computev1beta1.Volume, aggOnmetalVolume *AggregateOnmetalVolume, retErr error) {
	var onmetalVolumeSrc computev1beta1.VolumeSource
	switch {
	case cfg.Remote != nil:
		log.V(1).Info("Creating onmetal volume")
		remote := cfg.Remote
		onmetalVolume := &storagev1beta1.Volume{
			ObjectMeta: metav1.ObjectMeta{
				Namespace:       s.cluster.Namespace(),
				Name:            s.cluster.IDGen().Generate(),
				OwnerReferences: s.optionalOwnerReferences(onmetalMachineGVK, optOnmetalMachine),
				Annotations: map[string]string{
					commonv1beta1.ManagedByAnnotation: machinebrokerv1beta1.MachineBrokerManager,
				},
				Labels: map[string]string{
					machinebrokerv1beta1.ManagerLabel: machinebrokerv1beta1.MachineBrokerManager,
				},
			},
			Spec: storagev1beta1.VolumeSpec{
				ClaimRef: s.optionalLocalUIDReference(optOnmetalMachine),
			},
		}
		if err := s.cluster.Client().Create(ctx, onmetalVolume); err != nil {
			return nil, nil, fmt.Errorf("error creating onmetal volume: %w", err)
		}
		c.Add(cleaner.CleanupObject(s.cluster.Client(), onmetalVolume))

		var (
			secretRef    *corev1.LocalObjectReference
			accessSecret *corev1.Secret
		)
		if secretData := remote.SecretData; secretData != nil {
			log.V(1).Info("Creating onmetal volume secret")
			accessSecret = &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: s.cluster.Namespace(),
					Name:      s.cluster.IDGen().Generate(),
					OwnerReferences: []metav1.OwnerReference{
						metautils.MakeControllerRef(
							storagev1beta1.SchemeGroupVersion.WithKind("Volume"),
							onmetalVolume,
						),
					},
					Labels: map[string]string{
						machinebrokerv1beta1.ManagerLabel: machinebrokerv1beta1.MachineBrokerManager,
					},
				},
				Type: storagev1beta1.SecretTypeVolumeAuth,
				Data: secretData,
			}
			if err := s.cluster.Client().Create(ctx, accessSecret); err != nil {
				return nil, nil, fmt.Errorf("error creating onmetal volume secret: %w", err)
			}
			c.Add(cleaner.CleanupObject(s.cluster.Client(), accessSecret))
			secretRef = &corev1.LocalObjectReference{Name: accessSecret.Name}
		}

		log.V(1).Info("Patching onmetal volume status")
		baseOnmetalVolume := onmetalVolume.DeepCopy()
		onmetalVolume.Status.State = storagev1beta1.VolumeStateAvailable
		onmetalVolume.Status.Access = &storagev1beta1.VolumeAccess{
			SecretRef:        secretRef,
			Driver:           remote.Driver,
			Handle:           remote.Handle,
			VolumeAttributes: remote.Attributes,
		}
		if err := s.cluster.Client().Status().Patch(ctx, onmetalVolume, client.MergeFrom(baseOnmetalVolume)); err != nil {
			return nil, nil, fmt.Errorf("error patching onmetal volume status: %w", err)
		}

		aggOnmetalVolume = &AggregateOnmetalVolume{
			Volume:       onmetalVolume,
			AccessSecret: accessSecret,
		}
		onmetalVolumeSrc.VolumeRef = &corev1.LocalObjectReference{Name: onmetalVolume.Name}
	case cfg.EmptyDisk != nil:
		onmetalVolumeSrc.EmptyDisk = &computev1beta1.EmptyDiskVolumeSource{
			SizeLimit: cfg.EmptyDisk.SizeLimit,
		}
	}
	return &computev1beta1.Volume{
		Name:         cfg.Name,
		Device:       &cfg.Device,
		VolumeSource: onmetalVolumeSrc,
	}, aggOnmetalVolume, nil
}

func (s *Server) attachOnmetalVolume(
	ctx context.Context,
	log logr.Logger,
	onmetalMachine *computev1beta1.Machine,
	onmetalMachineVolume *computev1beta1.Volume,
) error {
	baseOnmetalMachine := onmetalMachine.DeepCopy()
	onmetalMachine.Spec.Volumes = append(onmetalMachine.Spec.Volumes, *onmetalMachineVolume)
	if err := s.cluster.Client().Patch(ctx, onmetalMachine, client.StrategicMergeFrom(baseOnmetalMachine)); err != nil {
		return fmt.Errorf("error patching onmetal machine volumes: %w", err)
	}
	return nil
}

func (s *Server) AttachVolume(ctx context.Context, req *ori.AttachVolumeRequest) (res *ori.AttachVolumeResponse, retErr error) {
	machineID := req.MachineId
	volumeName := req.Volume.Name
	log := s.loggerFrom(ctx, "MachineID", machineID, "VolumeName", volumeName)

	log.V(1).Info("Getting onmetal machine")
	onmetalMachine, err := s.getOnmetalMachine(ctx, machineID)
	if err != nil {
		return nil, err
	}

	log.V(1).Info("Getting onmetal volume config")
	cfg, err := s.getOnmetalVolumeConfig(req.Volume)
	if err != nil {
		return nil, err
	}

	c, cleanup := s.setupCleaner(ctx, log, &retErr)
	defer cleanup()

	log.V(1).Info("Creating onmetal volume")
	onmetalMachineVolume, _, err := s.createOnmetalVolume(ctx, log, c, onmetalMachine, cfg)
	if err != nil {
		return nil, err
	}

	log.V(1).Info("Attaching onmetal volume")
	if err := s.attachOnmetalVolume(ctx, log, onmetalMachine, onmetalMachineVolume); err != nil {
		return nil, err
	}

	return &ori.AttachVolumeResponse{}, nil
}
