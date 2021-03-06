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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	compute "github.com/onmetal/onmetal-api/apis/compute"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachinePools implements MachinePoolInterface
type FakeMachinePools struct {
	Fake *FakeCompute
}

var machinepoolsResource = schema.GroupVersionResource{Group: "compute.api.onmetal.de", Version: "", Resource: "machinepools"}

var machinepoolsKind = schema.GroupVersionKind{Group: "compute.api.onmetal.de", Version: "", Kind: "MachinePool"}

// Get takes name of the machinePool, and returns the corresponding machinePool object, and an error if there is any.
func (c *FakeMachinePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *compute.MachinePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(machinepoolsResource, name), &compute.MachinePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*compute.MachinePool), err
}

// List takes label and field selectors, and returns the list of MachinePools that match those selectors.
func (c *FakeMachinePools) List(ctx context.Context, opts v1.ListOptions) (result *compute.MachinePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(machinepoolsResource, machinepoolsKind, opts), &compute.MachinePoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &compute.MachinePoolList{ListMeta: obj.(*compute.MachinePoolList).ListMeta}
	for _, item := range obj.(*compute.MachinePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machinePools.
func (c *FakeMachinePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(machinepoolsResource, opts))
}

// Create takes the representation of a machinePool and creates it.  Returns the server's representation of the machinePool, and an error, if there is any.
func (c *FakeMachinePools) Create(ctx context.Context, machinePool *compute.MachinePool, opts v1.CreateOptions) (result *compute.MachinePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(machinepoolsResource, machinePool), &compute.MachinePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*compute.MachinePool), err
}

// Update takes the representation of a machinePool and updates it. Returns the server's representation of the machinePool, and an error, if there is any.
func (c *FakeMachinePools) Update(ctx context.Context, machinePool *compute.MachinePool, opts v1.UpdateOptions) (result *compute.MachinePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(machinepoolsResource, machinePool), &compute.MachinePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*compute.MachinePool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachinePools) UpdateStatus(ctx context.Context, machinePool *compute.MachinePool, opts v1.UpdateOptions) (*compute.MachinePool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(machinepoolsResource, "status", machinePool), &compute.MachinePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*compute.MachinePool), err
}

// Delete takes name of the machinePool and deletes it. Returns an error if one occurs.
func (c *FakeMachinePools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(machinepoolsResource, name, opts), &compute.MachinePool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachinePools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(machinepoolsResource, listOpts)

	_, err := c.Fake.Invokes(action, &compute.MachinePoolList{})
	return err
}

// Patch applies the patch and returns the patched machinePool.
func (c *FakeMachinePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *compute.MachinePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(machinepoolsResource, name, pt, data, subresources...), &compute.MachinePool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*compute.MachinePool), err
}
