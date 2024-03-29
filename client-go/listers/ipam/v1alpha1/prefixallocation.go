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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/onmetal/onmetal-api/api/ipam/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PrefixAllocationLister helps list PrefixAllocations.
// All objects returned here must be treated as read-only.
type PrefixAllocationLister interface {
	// List lists all PrefixAllocations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PrefixAllocation, err error)
	// PrefixAllocations returns an object that can list and get PrefixAllocations.
	PrefixAllocations(namespace string) PrefixAllocationNamespaceLister
	PrefixAllocationListerExpansion
}

// prefixAllocationLister implements the PrefixAllocationLister interface.
type prefixAllocationLister struct {
	indexer cache.Indexer
}

// NewPrefixAllocationLister returns a new PrefixAllocationLister.
func NewPrefixAllocationLister(indexer cache.Indexer) PrefixAllocationLister {
	return &prefixAllocationLister{indexer: indexer}
}

// List lists all PrefixAllocations in the indexer.
func (s *prefixAllocationLister) List(selector labels.Selector) (ret []*v1alpha1.PrefixAllocation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrefixAllocation))
	})
	return ret, err
}

// PrefixAllocations returns an object that can list and get PrefixAllocations.
func (s *prefixAllocationLister) PrefixAllocations(namespace string) PrefixAllocationNamespaceLister {
	return prefixAllocationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrefixAllocationNamespaceLister helps list and get PrefixAllocations.
// All objects returned here must be treated as read-only.
type PrefixAllocationNamespaceLister interface {
	// List lists all PrefixAllocations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PrefixAllocation, err error)
	// Get retrieves the PrefixAllocation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PrefixAllocation, error)
	PrefixAllocationNamespaceListerExpansion
}

// prefixAllocationNamespaceLister implements the PrefixAllocationNamespaceLister
// interface.
type prefixAllocationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrefixAllocations in the indexer for a given namespace.
func (s prefixAllocationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PrefixAllocation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrefixAllocation))
	})
	return ret, err
}

// Get retrieves the PrefixAllocation from the indexer for a given namespace and name.
func (s prefixAllocationNamespaceLister) Get(name string) (*v1alpha1.PrefixAllocation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("prefixallocation"), name)
	}
	return obj.(*v1alpha1.PrefixAllocation), nil
}
