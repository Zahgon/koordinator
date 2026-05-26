/*
Copyright 2022 The Koordinator Authors.

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

package client

import (
	"k8s.io/client-go/rest"
)

var (
	cfg *rest.Config

	defaultGenericClient *GenericClientset
)

// NewRegistry creates clientset by client-go
func NewRegistry(c *rest.Config) error { _ = "STUB: not implemented"; return nil }

// GetGenericClient returns default clientset
func GetGenericClient() *GenericClientset { _ = "STUB: not implemented"; return nil }

// GetGenericClientWithName returns clientset with given name as user-agent
func GetGenericClientWithName(name string) *GenericClientset { _ = "STUB: not implemented"; return nil }
