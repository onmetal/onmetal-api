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

package utils

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ObjectId struct {
	client.ObjectKey
	schema.GroupKind
}

func NewObjectIdForRequest(req ctrl.Request, gk schema.GroupKind) ObjectId {
	return ObjectId{
		ObjectKey: client.ObjectKey{
			Namespace: req.Namespace,
			Name:      req.Name,
		},
		GroupKind: gk,
	}
}

func NewObjectId(object client.Object) ObjectId {
	gvk := object.GetObjectKind().GroupVersionKind()
	return ObjectId{
		ObjectKey: client.ObjectKey{
			Namespace: object.GetNamespace(),
			Name:      object.GetName()},
		GroupKind: schema.GroupKind{
			Group: gvk.Group,
			Kind:  gvk.Kind,
		},
	}
}

func GetOwnerIdsFor(object client.Object) ObjectIds {
	ids := ObjectIds{}
	for _, o := range object.GetOwnerReferences() {
		gv, err := schema.ParseGroupVersion(o.APIVersion)
		if err == nil {
			id := ObjectId{
				ObjectKey: client.ObjectKey{
					Namespace: object.GetNamespace(),
					Name:      o.Name,
				},
				GroupKind: schema.GroupKind{
					Group: gv.Group,
					Kind:  o.Kind,
				},
			}
			ids.Add(id)
		}
	}
	return ids
}

func (o ObjectId) String() string {
	return fmt.Sprintf("%s/%s", o.GroupKind, o.ObjectKey)
}

type ObjectIds map[ObjectId]struct{}

func (o ObjectIds) Add(id ObjectId) {
	o[id] = struct{}{}
}

func (o ObjectIds) Remove(id ObjectId) {
	delete(o, id)
}

func (o ObjectIds) Copy() ObjectIds {
	if o == nil {
		return nil
	}
	new := ObjectIds{}
	for id := range o {
		new.Add(id)
	}
	return new
}

func (o ObjectIds) String() string {
	s := "["
	sep := ""
	for id := range o {
		s = fmt.Sprintf("%s%s%s", s, sep, id)
		sep = ","
	}
	return s + "]"
}

func (o ObjectIds) Equal(ids ObjectIds) bool {
	if o == nil && ids == nil {
		return true
	}

	if len(o) != len(ids) {
		return false
	}

	for id := range ids {
		if _, ok := o[id]; !ok {
			return false
		}
	}
	return true
}
