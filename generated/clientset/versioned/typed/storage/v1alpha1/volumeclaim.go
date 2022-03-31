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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/onmetal/onmetal-api/apis/storage/v1alpha1"
	scheme "github.com/onmetal/onmetal-api/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumeClaimsGetter has a method to return a VolumeClaimInterface.
// A group's client should implement this interface.
type VolumeClaimsGetter interface {
	VolumeClaims(namespace string) VolumeClaimInterface
}

// VolumeClaimInterface has methods to work with VolumeClaim resources.
type VolumeClaimInterface interface {
	Create(ctx context.Context, volumeClaim *v1alpha1.VolumeClaim, opts v1.CreateOptions) (*v1alpha1.VolumeClaim, error)
	Update(ctx context.Context, volumeClaim *v1alpha1.VolumeClaim, opts v1.UpdateOptions) (*v1alpha1.VolumeClaim, error)
	UpdateStatus(ctx context.Context, volumeClaim *v1alpha1.VolumeClaim, opts v1.UpdateOptions) (*v1alpha1.VolumeClaim, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VolumeClaim, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VolumeClaimList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeClaim, err error)
	VolumeClaimExpansion
}

// volumeClaims implements VolumeClaimInterface
type volumeClaims struct {
	client rest.Interface
	ns     string
}

// newVolumeClaims returns a VolumeClaims
func newVolumeClaims(c *StorageV1alpha1Client, namespace string) *volumeClaims {
	return &volumeClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the volumeClaim, and returns the corresponding volumeClaim object, and an error if there is any.
func (c *volumeClaims) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumeClaim, err error) {
	result = &v1alpha1.VolumeClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumeclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeClaims that match those selectors.
func (c *volumeClaims) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeClaimList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VolumeClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeClaims.
func (c *volumeClaims) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumeClaim and creates it.  Returns the server's representation of the volumeClaim, and an error, if there is any.
func (c *volumeClaims) Create(ctx context.Context, volumeClaim *v1alpha1.VolumeClaim, opts v1.CreateOptions) (result *v1alpha1.VolumeClaim, err error) {
	result = &v1alpha1.VolumeClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumeclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeClaim).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumeClaim and updates it. Returns the server's representation of the volumeClaim, and an error, if there is any.
func (c *volumeClaims) Update(ctx context.Context, volumeClaim *v1alpha1.VolumeClaim, opts v1.UpdateOptions) (result *v1alpha1.VolumeClaim, err error) {
	result = &v1alpha1.VolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumeclaims").
		Name(volumeClaim.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeClaim).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *volumeClaims) UpdateStatus(ctx context.Context, volumeClaim *v1alpha1.VolumeClaim, opts v1.UpdateOptions) (result *v1alpha1.VolumeClaim, err error) {
	result = &v1alpha1.VolumeClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumeclaims").
		Name(volumeClaim.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeClaim).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumeClaim and deletes it. Returns an error if one occurs.
func (c *volumeClaims) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumeclaims").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeClaims) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumeclaims").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumeClaim.
func (c *volumeClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeClaim, err error) {
	result = &v1alpha1.VolumeClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumeclaims").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}