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

	networking "github.com/onmetal/onmetal-api/apis/networking"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualIPRoutings implements VirtualIPRoutingInterface
type FakeVirtualIPRoutings struct {
	Fake *FakeNetworking
	ns   string
}

var virtualiproutingsResource = schema.GroupVersionResource{Group: "networking.api.onmetal.de", Version: "", Resource: "virtualiproutings"}

var virtualiproutingsKind = schema.GroupVersionKind{Group: "networking.api.onmetal.de", Version: "", Kind: "VirtualIPRouting"}

// Get takes name of the virtualIPRouting, and returns the corresponding virtualIPRouting object, and an error if there is any.
func (c *FakeVirtualIPRoutings) Get(ctx context.Context, name string, options v1.GetOptions) (result *networking.VirtualIPRouting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualiproutingsResource, c.ns, name), &networking.VirtualIPRouting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networking.VirtualIPRouting), err
}

// List takes label and field selectors, and returns the list of VirtualIPRoutings that match those selectors.
func (c *FakeVirtualIPRoutings) List(ctx context.Context, opts v1.ListOptions) (result *networking.VirtualIPRoutingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualiproutingsResource, virtualiproutingsKind, c.ns, opts), &networking.VirtualIPRoutingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networking.VirtualIPRoutingList{ListMeta: obj.(*networking.VirtualIPRoutingList).ListMeta}
	for _, item := range obj.(*networking.VirtualIPRoutingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualIPRoutings.
func (c *FakeVirtualIPRoutings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualiproutingsResource, c.ns, opts))

}

// Create takes the representation of a virtualIPRouting and creates it.  Returns the server's representation of the virtualIPRouting, and an error, if there is any.
func (c *FakeVirtualIPRoutings) Create(ctx context.Context, virtualIPRouting *networking.VirtualIPRouting, opts v1.CreateOptions) (result *networking.VirtualIPRouting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualiproutingsResource, c.ns, virtualIPRouting), &networking.VirtualIPRouting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networking.VirtualIPRouting), err
}

// Update takes the representation of a virtualIPRouting and updates it. Returns the server's representation of the virtualIPRouting, and an error, if there is any.
func (c *FakeVirtualIPRoutings) Update(ctx context.Context, virtualIPRouting *networking.VirtualIPRouting, opts v1.UpdateOptions) (result *networking.VirtualIPRouting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualiproutingsResource, c.ns, virtualIPRouting), &networking.VirtualIPRouting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networking.VirtualIPRouting), err
}

// Delete takes name of the virtualIPRouting and deletes it. Returns an error if one occurs.
func (c *FakeVirtualIPRoutings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualiproutingsResource, c.ns, name, opts), &networking.VirtualIPRouting{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualIPRoutings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualiproutingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &networking.VirtualIPRoutingList{})
	return err
}

// Patch applies the patch and returns the patched virtualIPRouting.
func (c *FakeVirtualIPRoutings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networking.VirtualIPRouting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualiproutingsResource, c.ns, name, pt, data, subresources...), &networking.VirtualIPRouting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networking.VirtualIPRouting), err
}
