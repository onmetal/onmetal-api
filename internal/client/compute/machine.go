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

package compute

import (
	"context"

	computev1beta1 "github.com/onmetal/onmetal-api/api/compute/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	MachineSpecMachinePoolRefNameField    = computev1beta1.MachineMachinePoolRefNameField
	MachineSpecMachineClassRefNameField   = computev1beta1.MachineMachineClassRefNameField
	MachineSpecNetworkInterfaceNamesField = "machine-spec-network-interface-names"
	MachineSpecVolumeNamesField           = "machine-spec-volume-names"
)

func SetupMachineSpecMachinePoolRefNameFieldIndexer(ctx context.Context, indexer client.FieldIndexer) error {
	return indexer.IndexField(ctx, &computev1beta1.Machine{}, MachineSpecMachinePoolRefNameField, func(obj client.Object) []string {
		machine := obj.(*computev1beta1.Machine)
		machinePoolRef := machine.Spec.MachinePoolRef
		if machinePoolRef == nil {
			return []string{""}
		}
		return []string{machinePoolRef.Name}
	})
}

func SetupMachineSpecMachineClassRefNameFieldIndexer(ctx context.Context, indexer client.FieldIndexer) error {
	return indexer.IndexField(ctx, &computev1beta1.Machine{}, MachineSpecMachineClassRefNameField, func(obj client.Object) []string {
		machine := obj.(*computev1beta1.Machine)
		return []string{machine.Spec.MachineClassRef.Name}
	})
}

func SetupMachineSpecNetworkInterfaceNamesFieldIndexer(ctx context.Context, indexer client.FieldIndexer) error {
	return indexer.IndexField(ctx, &computev1beta1.Machine{}, MachineSpecNetworkInterfaceNamesField, func(obj client.Object) []string {
		machine := obj.(*computev1beta1.Machine)
		if names := computev1beta1.MachineNetworkInterfaceNames(machine); len(names) > 0 {
			return names
		}
		return []string{""}
	})
}

func SetupMachineSpecVolumeNamesFieldIndexer(ctx context.Context, indexer client.FieldIndexer) error {
	return indexer.IndexField(ctx, &computev1beta1.Machine{}, MachineSpecVolumeNamesField, func(obj client.Object) []string {
		machine := obj.(*computev1beta1.Machine)
		if names := computev1beta1.MachineVolumeNames(machine); len(names) > 0 {
			return names
		}
		return []string{""}
	})
}
