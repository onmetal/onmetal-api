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

	v1alpha1 "github.com/onmetal/onmetal-api/api/storage/v1alpha1"
	storagev1alpha1 "github.com/onmetal/onmetal-api/client-go/applyconfigurations/storage/v1alpha1"
	scheme "github.com/onmetal/onmetal-api/client-go/onmetalapi/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumePoolsGetter has a method to return a VolumePoolInterface.
// A group's client should implement this interface.
type VolumePoolsGetter interface {
	VolumePools() VolumePoolInterface
}

// VolumePoolInterface has methods to work with VolumePool resources.
type VolumePoolInterface interface {
	Create(ctx context.Context, volumePool *v1alpha1.VolumePool, opts v1.CreateOptions) (*v1alpha1.VolumePool, error)
	Update(ctx context.Context, volumePool *v1alpha1.VolumePool, opts v1.UpdateOptions) (*v1alpha1.VolumePool, error)
	UpdateStatus(ctx context.Context, volumePool *v1alpha1.VolumePool, opts v1.UpdateOptions) (*v1alpha1.VolumePool, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VolumePool, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VolumePoolList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumePool, err error)
	Apply(ctx context.Context, volumePool *storagev1alpha1.VolumePoolApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumePool, err error)
	ApplyStatus(ctx context.Context, volumePool *storagev1alpha1.VolumePoolApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumePool, err error)
	VolumePoolExpansion
}

// volumePools implements VolumePoolInterface
type volumePools struct {
	client rest.Interface
}

// newVolumePools returns a VolumePools
func newVolumePools(c *StorageV1alpha1Client) *volumePools {
	return &volumePools{
		client: c.RESTClient(),
	}
}

// Get takes name of the volumePool, and returns the corresponding volumePool object, and an error if there is any.
func (c *volumePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumePool, err error) {
	result = &v1alpha1.VolumePool{}
	err = c.client.Get().
		Resource("volumepools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumePools that match those selectors.
func (c *volumePools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumePoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VolumePoolList{}
	err = c.client.Get().
		Resource("volumepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumePools.
func (c *volumePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("volumepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a volumePool and creates it.  Returns the server's representation of the volumePool, and an error, if there is any.
func (c *volumePools) Create(ctx context.Context, volumePool *v1alpha1.VolumePool, opts v1.CreateOptions) (result *v1alpha1.VolumePool, err error) {
	result = &v1alpha1.VolumePool{}
	err = c.client.Post().
		Resource("volumepools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumePool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a volumePool and updates it. Returns the server's representation of the volumePool, and an error, if there is any.
func (c *volumePools) Update(ctx context.Context, volumePool *v1alpha1.VolumePool, opts v1.UpdateOptions) (result *v1alpha1.VolumePool, err error) {
	result = &v1alpha1.VolumePool{}
	err = c.client.Put().
		Resource("volumepools").
		Name(volumePool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumePool).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *volumePools) UpdateStatus(ctx context.Context, volumePool *v1alpha1.VolumePool, opts v1.UpdateOptions) (result *v1alpha1.VolumePool, err error) {
	result = &v1alpha1.VolumePool{}
	err = c.client.Put().
		Resource("volumepools").
		Name(volumePool.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumePool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the volumePool and deletes it. Returns an error if one occurs.
func (c *volumePools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("volumepools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumePools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("volumepools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched volumePool.
func (c *volumePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumePool, err error) {
	result = &v1alpha1.VolumePool{}
	err = c.client.Patch(pt).
		Resource("volumepools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied volumePool.
func (c *volumePools) Apply(ctx context.Context, volumePool *storagev1alpha1.VolumePoolApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumePool, err error) {
	if volumePool == nil {
		return nil, fmt.Errorf("volumePool provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(volumePool)
	if err != nil {
		return nil, err
	}
	name := volumePool.Name
	if name == nil {
		return nil, fmt.Errorf("volumePool.Name must be provided to Apply")
	}
	result = &v1alpha1.VolumePool{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("volumepools").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *volumePools) ApplyStatus(ctx context.Context, volumePool *storagev1alpha1.VolumePoolApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumePool, err error) {
	if volumePool == nil {
		return nil, fmt.Errorf("volumePool provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(volumePool)
	if err != nil {
		return nil, err
	}

	name := volumePool.Name
	if name == nil {
		return nil, fmt.Errorf("volumePool.Name must be provided to Apply")
	}

	result = &v1alpha1.VolumePool{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("volumepools").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
