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

package networking

import (
	"github.com/ironcore-dev/ironcore/internal/apis/ipam"
)

const (
	// NetworkPluginsGroup is the system rbac group all network plugins are in.
	NetworkPluginsGroup = "networking.ironcore.dev:system:networkplugins"

	// NetworkPluginUserNamePrefix is the prefix all network plugin users should have.
	NetworkPluginUserNamePrefix = "networking.ironcore.dev:system:networkplugin:"
)

// NetworkPluginCommonName constructs the common name for a certificate of a network plugin user.
func NetworkPluginCommonName(name string) string {
	return NetworkPluginUserNamePrefix + name
}

// EphemeralPrefixSource contains the definition to create an ephemeral (i.e. coupled to the lifetime of the
// surrounding object) Prefix.
type EphemeralPrefixSource struct {
	// PrefixTemplate is the template for the Prefix.
	PrefixTemplate *ipam.PrefixTemplateSpec
}

// EphemeralVirtualIPSource contains the definition to create an ephemeral (i.e. coupled to the lifetime of the
// surrounding object) VirtualIP.
type EphemeralVirtualIPSource struct {
	// VirtualIPTemplate is the template for the VirtualIP.
	VirtualIPTemplate *VirtualIPTemplateSpec
}
