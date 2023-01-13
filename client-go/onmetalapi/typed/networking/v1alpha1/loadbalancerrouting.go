/*
 * Copyright (c) 2022 by the OnMetal authors.
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

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/onmetal/onmetal-api/api/networking/v1alpha1"
	networkingv1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/networking/v1alpha1"
	scheme "github.com/onmetal/onmetal-api/client-go/onmetalapi/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LoadBalancerRoutingsGetter has a method to return a LoadBalancerRoutingInterface.
// A group's client should implement this interface.
type LoadBalancerRoutingsGetter interface {
	LoadBalancerRoutings(namespace string) LoadBalancerRoutingInterface
}

// LoadBalancerRoutingInterface has methods to work with LoadBalancerRouting resources.
type LoadBalancerRoutingInterface interface {
	Create(ctx context.Context, loadBalancerRouting *v1alpha1.LoadBalancerRouting, opts v1.CreateOptions) (*v1alpha1.LoadBalancerRouting, error)
	Update(ctx context.Context, loadBalancerRouting *v1alpha1.LoadBalancerRouting, opts v1.UpdateOptions) (*v1alpha1.LoadBalancerRouting, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.LoadBalancerRouting, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.LoadBalancerRoutingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LoadBalancerRouting, err error)
	Apply(ctx context.Context, loadBalancerRouting *networkingv1alpha1.LoadBalancerRoutingApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.LoadBalancerRouting, err error)
	LoadBalancerRoutingExpansion
}

// loadBalancerRoutings implements LoadBalancerRoutingInterface
type loadBalancerRoutings struct {
	client rest.Interface
	ns     string
}

// newLoadBalancerRoutings returns a LoadBalancerRoutings
func newLoadBalancerRoutings(c *NetworkingV1alpha1Client, namespace string) *loadBalancerRoutings {
	return &loadBalancerRoutings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the loadBalancerRouting, and returns the corresponding loadBalancerRouting object, and an error if there is any.
func (c *loadBalancerRoutings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LoadBalancerRouting, err error) {
	result = &v1alpha1.LoadBalancerRouting{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loadbalancerroutings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LoadBalancerRoutings that match those selectors.
func (c *loadBalancerRoutings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LoadBalancerRoutingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LoadBalancerRoutingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loadbalancerroutings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loadBalancerRoutings.
func (c *loadBalancerRoutings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loadbalancerroutings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a loadBalancerRouting and creates it.  Returns the server's representation of the loadBalancerRouting, and an error, if there is any.
func (c *loadBalancerRoutings) Create(ctx context.Context, loadBalancerRouting *v1alpha1.LoadBalancerRouting, opts v1.CreateOptions) (result *v1alpha1.LoadBalancerRouting, err error) {
	result = &v1alpha1.LoadBalancerRouting{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loadbalancerroutings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loadBalancerRouting).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a loadBalancerRouting and updates it. Returns the server's representation of the loadBalancerRouting, and an error, if there is any.
func (c *loadBalancerRoutings) Update(ctx context.Context, loadBalancerRouting *v1alpha1.LoadBalancerRouting, opts v1.UpdateOptions) (result *v1alpha1.LoadBalancerRouting, err error) {
	result = &v1alpha1.LoadBalancerRouting{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loadbalancerroutings").
		Name(loadBalancerRouting.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loadBalancerRouting).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the loadBalancerRouting and deletes it. Returns an error if one occurs.
func (c *loadBalancerRoutings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loadbalancerroutings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loadBalancerRoutings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loadbalancerroutings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched loadBalancerRouting.
func (c *loadBalancerRoutings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LoadBalancerRouting, err error) {
	result = &v1alpha1.LoadBalancerRouting{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loadbalancerroutings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied loadBalancerRouting.
func (c *loadBalancerRoutings) Apply(ctx context.Context, loadBalancerRouting *networkingv1alpha1.LoadBalancerRoutingApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.LoadBalancerRouting, err error) {
	if loadBalancerRouting == nil {
		return nil, fmt.Errorf("loadBalancerRouting provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(loadBalancerRouting)
	if err != nil {
		return nil, err
	}
	name := loadBalancerRouting.Name
	if name == nil {
		return nil, fmt.Errorf("loadBalancerRouting.Name must be provided to Apply")
	}
	result = &v1alpha1.LoadBalancerRouting{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("loadbalancerroutings").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
