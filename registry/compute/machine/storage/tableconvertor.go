package storage

import (
	"context"

	"github.com/onmetal/onmetal-api/apis/compute"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/meta/table"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type convertor struct{}

var (
	objectMetaSwaggerDoc = metav1.ObjectMeta{}.SwaggerDoc()

	headers = []metav1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: objectMetaSwaggerDoc["name"]},
		{Name: "MachineClass", Type: "string", Description: "The machine class of this machine"},
		{Name: "Image", Type: "string", Description: "The image the machine shall use"},
		{Name: "MachinePool", Type: "string", Description: "The machine pool the machine is running on"},
		{Name: "State", Type: "string", Description: "The current state of the machine"},
		{Name: "Age", Type: "string", Format: "date", Description: objectMetaSwaggerDoc["creationTimestamp"]},
	}
)

func newTableConvertor() *convertor {
	return &convertor{}
}

func (c *convertor) ConvertToTable(ctx context.Context, obj runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	tab := &metav1.Table{
		ColumnDefinitions: headers,
	}

	if m, err := meta.ListAccessor(obj); err == nil {
		tab.ResourceVersion = m.GetResourceVersion()
		tab.SelfLink = m.GetSelfLink()
		tab.Continue = m.GetContinue()
	} else {
		if m, err := meta.CommonAccessor(obj); err == nil {
			tab.ResourceVersion = m.GetResourceVersion()
			tab.SelfLink = m.GetSelfLink()
		}
	}

	var err error
	tab.Rows, err = table.MetaToTableRow(obj, func(obj runtime.Object, m metav1.Object, name, age string) (cells []interface{}, err error) {
		machine := obj.(*compute.Machine)

		cells = append(cells, name)
		cells = append(cells, machine.Spec.MachineClass.Name)
		cells = append(cells, machine.Spec.Image)
		if machinePoolName := machine.Spec.MachinePool.Name; machinePoolName != "" {
			cells = append(cells, machinePoolName)
		} else {
			cells = append(cells, "<none>")
		}
		if state := machine.Status.State; state != "" {
			cells = append(cells, machine.Status.State)
		} else {
			cells = append(cells, "<unknown>")
		}
		cells = append(cells, age)

		return cells, nil
	})
	return tab, err
}