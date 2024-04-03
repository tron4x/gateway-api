/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// GRPCMethodMatchApplyConfiguration represents an declarative configuration of the GRPCMethodMatch type for use
// with apply.
type GRPCMethodMatchApplyConfiguration struct {
	Type    *v1alpha2.GRPCMethodMatchType `json:"type,omitempty"`
	Service *string                       `json:"service,omitempty"`
	Method  *string                       `json:"method,omitempty"`
}

// GRPCMethodMatchApplyConfiguration constructs an declarative configuration of the GRPCMethodMatch type for use with
// apply.
func GRPCMethodMatch() *GRPCMethodMatchApplyConfiguration {
	return &GRPCMethodMatchApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *GRPCMethodMatchApplyConfiguration) WithType(value v1alpha2.GRPCMethodMatchType) *GRPCMethodMatchApplyConfiguration {
	b.Type = &value
	return b
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *GRPCMethodMatchApplyConfiguration) WithService(value string) *GRPCMethodMatchApplyConfiguration {
	b.Service = &value
	return b
}

// WithMethod sets the Method field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Method field is set to the value of the last call.
func (b *GRPCMethodMatchApplyConfiguration) WithMethod(value string) *GRPCMethodMatchApplyConfiguration {
	b.Method = &value
	return b
}
