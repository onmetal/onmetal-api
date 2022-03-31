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

// MachinePoolLister helps list MachinePools.
// All objects returned here must be treated as read-only.
type MachinePoolLister interface {
	// List lists all MachinePools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*compute.MachinePool, err error)
	// Get retrieves the MachinePool from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*compute.MachinePool, error)
	MachinePoolListerExpansion
}

// machinePoolLister implements the MachinePoolLister interface.
type machinePoolLister struct {
	indexer cache.Indexer
}

// NewMachinePoolLister returns a new MachinePoolLister.
func NewMachinePoolLister(indexer cache.Indexer) MachinePoolLister {
	return &machinePoolLister{indexer: indexer}
}

// List lists all MachinePools in the indexer.
func (s *machinePoolLister) List(selector labels.Selector) (ret []*compute.MachinePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*compute.MachinePool))
	})
	return ret, err
}

// Get retrieves the MachinePool from the index for a given name.
func (s *machinePoolLister) Get(name string) (*compute.MachinePool, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(compute.Resource("machinepool"), name)
	}
	return obj.(*compute.MachinePool), nil
}