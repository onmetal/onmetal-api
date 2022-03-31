package storage

import (
	"context"

	"github.com/onmetal/onmetal-api/apis/storage"
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
		{Name: "StoragePool", Type: "string", Description: "The storage pool this volume is hosted on"},
		{Name: "StorageClass", Type: "string", Description: "The storage class of this volume"},
		{Name: "State", Type: "string", Description: "The state of the volume on the underlying storage pool"},
		{Name: "Phase", Type: "string", Description: "The binding phase of the volume"},
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
		volume := obj.(*storage.Volume)

		cells = append(cells, name)
		if storagePoolName := volume.Spec.StoragePool.Name; storagePoolName != "" {
			cells = append(cells, storagePoolName)
		} else {
			cells = append(cells, "<none>")
		}
		cells = append(cells, volume.Spec.StorageClassRef.Name)
		if state := volume.Status.Phase; state != "" {
			cells = append(cells, state)
		} else {
			cells = append(cells, "<unknown>")
		}
		if phase := volume.Status.Phase; phase != "" {
			cells = append(cells, phase)
		} else {
			cells = append(cells, "<unknown>")
		}
		cells = append(cells, age)

		return cells, nil
	})
	return tab, err
}