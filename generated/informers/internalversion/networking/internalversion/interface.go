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
// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	internalinterfaces "github.com/onmetal/onmetal-api/generated/informers/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Networks returns a NetworkInformer.
	Networks() NetworkInformer
	// NetworkInterfaces returns a NetworkInterfaceInformer.
	NetworkInterfaces() NetworkInterfaceInformer
	// VirtualIPs returns a VirtualIPInformer.
	VirtualIPs() VirtualIPInformer
	// VirtualIPRoutings returns a VirtualIPRoutingInformer.
	VirtualIPRoutings() VirtualIPRoutingInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Networks returns a NetworkInformer.
func (v *version) Networks() NetworkInformer {
	return &networkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkInterfaces returns a NetworkInterfaceInformer.
func (v *version) NetworkInterfaces() NetworkInterfaceInformer {
	return &networkInterfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualIPs returns a VirtualIPInformer.
func (v *version) VirtualIPs() VirtualIPInformer {
	return &virtualIPInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualIPRoutings returns a VirtualIPRoutingInformer.
func (v *version) VirtualIPRoutings() VirtualIPRoutingInformer {
	return &virtualIPRoutingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
