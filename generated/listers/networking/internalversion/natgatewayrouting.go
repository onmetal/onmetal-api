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
// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	networking "github.com/onmetal/onmetal-api/apis/networking"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NATGatewayRoutingLister helps list NATGatewayRoutings.
// All objects returned here must be treated as read-only.
type NATGatewayRoutingLister interface {
	// List lists all NATGatewayRoutings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networking.NATGatewayRouting, err error)
	// NATGatewayRoutings returns an object that can list and get NATGatewayRoutings.
	NATGatewayRoutings(namespace string) NATGatewayRoutingNamespaceLister
	NATGatewayRoutingListerExpansion
}

// nATGatewayRoutingLister implements the NATGatewayRoutingLister interface.
type nATGatewayRoutingLister struct {
	indexer cache.Indexer
}

// NewNATGatewayRoutingLister returns a new NATGatewayRoutingLister.
func NewNATGatewayRoutingLister(indexer cache.Indexer) NATGatewayRoutingLister {
	return &nATGatewayRoutingLister{indexer: indexer}
}

// List lists all NATGatewayRoutings in the indexer.
func (s *nATGatewayRoutingLister) List(selector labels.Selector) (ret []*networking.NATGatewayRouting, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*networking.NATGatewayRouting))
	})
	return ret, err
}

// NATGatewayRoutings returns an object that can list and get NATGatewayRoutings.
func (s *nATGatewayRoutingLister) NATGatewayRoutings(namespace string) NATGatewayRoutingNamespaceLister {
	return nATGatewayRoutingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NATGatewayRoutingNamespaceLister helps list and get NATGatewayRoutings.
// All objects returned here must be treated as read-only.
type NATGatewayRoutingNamespaceLister interface {
	// List lists all NATGatewayRoutings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networking.NATGatewayRouting, err error)
	// Get retrieves the NATGatewayRouting from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*networking.NATGatewayRouting, error)
	NATGatewayRoutingNamespaceListerExpansion
}

// nATGatewayRoutingNamespaceLister implements the NATGatewayRoutingNamespaceLister
// interface.
type nATGatewayRoutingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NATGatewayRoutings in the indexer for a given namespace.
func (s nATGatewayRoutingNamespaceLister) List(selector labels.Selector) (ret []*networking.NATGatewayRouting, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*networking.NATGatewayRouting))
	})
	return ret, err
}

// Get retrieves the NATGatewayRouting from the indexer for a given namespace and name.
func (s nATGatewayRoutingNamespaceLister) Get(name string) (*networking.NATGatewayRouting, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(networking.Resource("natgatewayrouting"), name)
	}
	return obj.(*networking.NATGatewayRouting), nil
}