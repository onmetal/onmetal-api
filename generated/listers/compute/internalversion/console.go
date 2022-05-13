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
	compute "github.com/onmetal/onmetal-api/apis/compute"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConsoleLister helps list Consoles.
// All objects returned here must be treated as read-only.
type ConsoleLister interface {
	// List lists all Consoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*compute.Console, err error)
	// Consoles returns an object that can list and get Consoles.
	Consoles(namespace string) ConsoleNamespaceLister
	ConsoleListerExpansion
}

// consoleLister implements the ConsoleLister interface.
type consoleLister struct {
	indexer cache.Indexer
}

// NewConsoleLister returns a new ConsoleLister.
func NewConsoleLister(indexer cache.Indexer) ConsoleLister {
	return &consoleLister{indexer: indexer}
}

// List lists all Consoles in the indexer.
func (s *consoleLister) List(selector labels.Selector) (ret []*compute.Console, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*compute.Console))
	})
	return ret, err
}

// Consoles returns an object that can list and get Consoles.
func (s *consoleLister) Consoles(namespace string) ConsoleNamespaceLister {
	return consoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConsoleNamespaceLister helps list and get Consoles.
// All objects returned here must be treated as read-only.
type ConsoleNamespaceLister interface {
	// List lists all Consoles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*compute.Console, err error)
	// Get retrieves the Console from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*compute.Console, error)
	ConsoleNamespaceListerExpansion
}

// consoleNamespaceLister implements the ConsoleNamespaceLister
// interface.
type consoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Consoles in the indexer for a given namespace.
func (s consoleNamespaceLister) List(selector labels.Selector) (ret []*compute.Console, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*compute.Console))
	})
	return ret, err
}

// Get retrieves the Console from the indexer for a given namespace and name.
func (s consoleNamespaceLister) Get(name string) (*compute.Console, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(compute.Resource("console"), name)
	}
	return obj.(*compute.Console), nil
}
