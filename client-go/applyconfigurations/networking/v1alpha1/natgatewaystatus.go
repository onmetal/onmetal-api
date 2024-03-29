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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/onmetal/onmetal-api/api/common/v1alpha1"
)

// NATGatewayStatusApplyConfiguration represents an declarative configuration of the NATGatewayStatus type for use
// with apply.
type NATGatewayStatusApplyConfiguration struct {
	IPs []v1alpha1.IP `json:"ips,omitempty"`
}

// NATGatewayStatusApplyConfiguration constructs an declarative configuration of the NATGatewayStatus type for use with
// apply.
func NATGatewayStatus() *NATGatewayStatusApplyConfiguration {
	return &NATGatewayStatusApplyConfiguration{}
}

// WithIPs adds the given value to the IPs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the IPs field.
func (b *NATGatewayStatusApplyConfiguration) WithIPs(values ...v1alpha1.IP) *NATGatewayStatusApplyConfiguration {
	for i := range values {
		b.IPs = append(b.IPs, values[i])
	}
	return b
}
