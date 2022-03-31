package storage

import (
	"context"

	"github.com/onmetal/controller-utils/conditionutils"
	"github.com/onmetal/onmetal-api/apis/ipam"
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
		{Name: "Prefix", Type: "string", Description: "The managed prefix, if any"},
		{Name: "Parent", Type: "string", Description: "The parent, if any"},
		{Name: "State", Type: "string", Description: "The allocation of the prefix"},
		{Name: "Age", Type: "string", Format: "date", Description: objectMetaSwaggerDoc["creationTimestamp"]},
	}
)

func newTableConvertor() *convertor {
	return &convertor{}
}

func prefixReadyState(prefix *ipam.Prefix) string {
	readyCond := ipam.PrefixCondition{}
	conditionutils.MustFindSlice(prefix.Status.Conditions, string(ipam.PrefixReady), &readyCond)
	return readyCond.Reason
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
		prefix := obj.(*ipam.Prefix)

		cells = append(cells, name)
		if prefix := prefix.Spec.Prefix; prefix.IsValid() {
			cells = append(cells, prefix.String())
		} else {
			cells = append(cells, "<none>")
		}
		if parentRef := prefix.Spec.ParentRef; parentRef != nil {
			cells = append(cells, parentRef.String())
		} else {
			cells = append(cells, "<none>")
		}
		if readyState := prefixReadyState(prefix); readyState != "" {
			cells = append(cells, readyState)
		} else {
			cells = append(cells, "<unknown>")
		}
		cells = append(cells, age)

		return cells, nil
	})
	return tab, err
}