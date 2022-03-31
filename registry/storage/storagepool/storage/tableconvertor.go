package storage

import (
	"context"
	"sort"
	"strings"

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
		{Name: "StorageClasses", Type: "string", Description: "Storage classes available on this storage pool."},
		{Name: "Age", Type: "string", Format: "date", Description: objectMetaSwaggerDoc["creationTimestamp"]},
	}
)

func newTableConvertor() *convertor {
	return &convertor{}
}

func storagePoolStorageClassNames(storagePool *storage.StoragePool) []string {
	names := make([]string, 0, len(storagePool.Status.AvailableStorageClasses))
	for _, storageClassRef := range storagePool.Status.AvailableStorageClasses {
		names = append(names, storageClassRef.Name)
	}
	sort.Strings(names)
	return names
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
		storagePool := obj.(*storage.StoragePool)

		cells = append(cells, name)
		storageClassNames := storagePoolStorageClassNames(storagePool)
		if len(storageClassNames) == 0 {
			cells = append(cells, "<none>")
		} else {
			cells = append(cells, strings.Join(storageClassNames, ","))
		}
		cells = append(cells, age)

		return cells, nil
	})
	return tab, err
}