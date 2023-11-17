/*
 * Copyright (c) 2022 by the IronCore authors.
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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	storagev1alpha1 "github.com/ironcore-dev/ironcore/api/storage/v1alpha1"
	internalinterfaces "github.com/ironcore-dev/ironcore/client-go/informers/internalinterfaces"
	ironcore "github.com/ironcore-dev/ironcore/client-go/ironcore"
	v1alpha1 "github.com/ironcore-dev/ironcore/client-go/listers/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VolumePoolInformer provides access to a shared informer and lister for
// VolumePools.
type VolumePoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VolumePoolLister
}

type volumePoolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVolumePoolInformer constructs a new informer for VolumePool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVolumePoolInformer(client ironcore.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVolumePoolInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVolumePoolInformer constructs a new informer for VolumePool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVolumePoolInformer(client ironcore.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1alpha1().VolumePools().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1alpha1().VolumePools().Watch(context.TODO(), options)
			},
		},
		&storagev1alpha1.VolumePool{},
		resyncPeriod,
		indexers,
	)
}

func (f *volumePoolInformer) defaultInformer(client ironcore.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVolumePoolInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *volumePoolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1alpha1.VolumePool{}, f.defaultInformer)
}

func (f *volumePoolInformer) Lister() v1alpha1.VolumePoolLister {
	return v1alpha1.NewVolumePoolLister(f.Informer().GetIndexer())
}
