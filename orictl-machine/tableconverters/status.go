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

package tableconverters

import (
	ori "github.com/onmetal/onmetal-api/ori/apis/machine/v1alpha1"
	"github.com/onmetal/onmetal-api/orictl/api"
	"github.com/onmetal/onmetal-api/orictl/tableconverter"
	"k8s.io/apimachinery/pkg/api/resource"
)

var (
	machineClassHeaders = []api.Header{
		{Name: "Name"},
		{Name: "CPU"},
		{Name: "Memory"},
		{Name: "Quantity"},
	}

	MachineClassStatus = tableconverter.Funcs[*ori.MachineClassStatus]{
		Headers: tableconverter.Headers(machineClassHeaders),
		Rows: tableconverter.SingleRowFrom(func(status *ori.MachineClassStatus) (api.Row, error) {
			return api.Row{
				status.MachineClass.Name,
				resource.NewMilliQuantity(status.MachineClass.Capabilities.CpuMillis, resource.DecimalSI).String(),
				resource.NewQuantity(int64(status.MachineClass.Capabilities.MemoryBytes), resource.DecimalSI).String(),
				resource.NewQuantity(status.Quantity, resource.DecimalSI).String(),
			}, nil
		}),
	}

	MachineClassStatusSlice = tableconverter.SliceFuncs[*ori.MachineClassStatus](MachineClassStatus)
)

func init() {
	RegistryBuilder.Register(
		tableconverter.ToTagAndTypedAny[*ori.MachineClassStatus](MachineClassStatus),
		tableconverter.ToTagAndTypedAny[[]*ori.MachineClassStatus](MachineClassStatusSlice),
	)
}
