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

	v1alpha1 "github.com/onmetal/onmetal-api/apis/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkInterfaceBindings implements NetworkInterfaceBindingInterface
type FakeNetworkInterfaceBindings struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var networkinterfacebindingsResource = schema.GroupVersionResource{Group: "networking.api.onmetal.de", Version: "v1alpha1", Resource: "networkinterfacebindings"}

var networkinterfacebindingsKind = schema.GroupVersionKind{Group: "networking.api.onmetal.de", Version: "v1alpha1", Kind: "NetworkInterfaceBinding"}

// Get takes name of the networkInterfaceBinding, and returns the corresponding networkInterfaceBinding object, and an error if there is any.
func (c *FakeNetworkInterfaceBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkInterfaceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkinterfacebindingsResource, c.ns, name), &v1alpha1.NetworkInterfaceBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceBinding), err
}

// List takes label and field selectors, and returns the list of NetworkInterfaceBindings that match those selectors.
func (c *FakeNetworkInterfaceBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkinterfacebindingsResource, networkinterfacebindingsKind, c.ns, opts), &v1alpha1.NetworkInterfaceBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkInterfaceBindingList{ListMeta: obj.(*v1alpha1.NetworkInterfaceBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkInterfaceBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkInterfaceBindings.
func (c *FakeNetworkInterfaceBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkinterfacebindingsResource, c.ns, opts))

}

// Create takes the representation of a networkInterfaceBinding and creates it.  Returns the server's representation of the networkInterfaceBinding, and an error, if there is any.
func (c *FakeNetworkInterfaceBindings) Create(ctx context.Context, networkInterfaceBinding *v1alpha1.NetworkInterfaceBinding, opts v1.CreateOptions) (result *v1alpha1.NetworkInterfaceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkinterfacebindingsResource, c.ns, networkInterfaceBinding), &v1alpha1.NetworkInterfaceBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceBinding), err
}

// Update takes the representation of a networkInterfaceBinding and updates it. Returns the server's representation of the networkInterfaceBinding, and an error, if there is any.
func (c *FakeNetworkInterfaceBindings) Update(ctx context.Context, networkInterfaceBinding *v1alpha1.NetworkInterfaceBinding, opts v1.UpdateOptions) (result *v1alpha1.NetworkInterfaceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkinterfacebindingsResource, c.ns, networkInterfaceBinding), &v1alpha1.NetworkInterfaceBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceBinding), err
}

// Delete takes name of the networkInterfaceBinding and deletes it. Returns an error if one occurs.
func (c *FakeNetworkInterfaceBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkinterfacebindingsResource, c.ns, name, opts), &v1alpha1.NetworkInterfaceBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkInterfaceBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkinterfacebindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkInterfaceBindingList{})
	return err
}

// Patch applies the patch and returns the patched networkInterfaceBinding.
func (c *FakeNetworkInterfaceBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkInterfaceBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkinterfacebindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkInterfaceBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterfaceBinding), err
}