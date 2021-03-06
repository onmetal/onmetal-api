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
	clientset "github.com/onmetal/onmetal-api/generated/clientset/internalversion"
	computeinternalversion "github.com/onmetal/onmetal-api/generated/clientset/internalversion/typed/compute/internalversion"
	fakecomputeinternalversion "github.com/onmetal/onmetal-api/generated/clientset/internalversion/typed/compute/internalversion/fake"
	ipaminternalversion "github.com/onmetal/onmetal-api/generated/clientset/internalversion/typed/ipam/internalversion"
	fakeipaminternalversion "github.com/onmetal/onmetal-api/generated/clientset/internalversion/typed/ipam/internalversion/fake"
	networkinginternalversion "github.com/onmetal/onmetal-api/generated/clientset/internalversion/typed/networking/internalversion"
	fakenetworkinginternalversion "github.com/onmetal/onmetal-api/generated/clientset/internalversion/typed/networking/internalversion/fake"
	storageinternalversion "github.com/onmetal/onmetal-api/generated/clientset/internalversion/typed/storage/internalversion"
	fakestorageinternalversion "github.com/onmetal/onmetal-api/generated/clientset/internalversion/typed/storage/internalversion/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// Compute retrieves the ComputeClient
func (c *Clientset) Compute() computeinternalversion.ComputeInterface {
	return &fakecomputeinternalversion.FakeCompute{Fake: &c.Fake}
}

// Ipam retrieves the IpamClient
func (c *Clientset) Ipam() ipaminternalversion.IpamInterface {
	return &fakeipaminternalversion.FakeIpam{Fake: &c.Fake}
}

// Networking retrieves the NetworkingClient
func (c *Clientset) Networking() networkinginternalversion.NetworkingInterface {
	return &fakenetworkinginternalversion.FakeNetworking{Fake: &c.Fake}
}

// Storage retrieves the StorageClient
func (c *Clientset) Storage() storageinternalversion.StorageInterface {
	return &fakestorageinternalversion.FakeStorage{Fake: &c.Fake}
}
