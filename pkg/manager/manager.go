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

package manager

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type Manager struct {
	manager.Manager
	triggers   ReconcilationTrigger
	ownerCache *OwnerCache
}

func NewManager(config *rest.Config, options manager.Options) (*Manager, error) {
	mgr, err := manager.New(config, options)
	if err != nil {
		return nil, err
	}
	trig := NewReconcilationTrigger()
	oc := NewOwnerCache(mgr, trig)
	return &Manager{
		Manager:    mgr,
		ownerCache: oc,
		triggers:   trig,
	}, nil
}

func (m *Manager) GetOwnerCache() *OwnerCache {
	return m.ownerCache
}
