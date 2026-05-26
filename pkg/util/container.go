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

package util

import (
	corev1 "k8s.io/api/core/v1"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
)

func GetEmptyContainerExtendedResources() *apiext.ExtendedResourceContainerSpec {
	_ = "STUB: not implemented"
	return nil
}

func GetContainerExtendedResources(container *corev1.Container) *apiext.ExtendedResourceContainerSpec {
	_ = "STUB: not implemented"
	return nil
}

// GetContainerTargetExtendedResources gets the resource requirements of a container with given extended resources.
// It returns nil if container is nil or specifies no extended resource.
func GetContainerTargetExtendedResources(container *corev1.Container, resourceNames ...corev1.ResourceName) *apiext.ExtendedResourceContainerSpec {
	_ = "STUB: not implemented"
	return nil
}

// assert container.Resources.Requests != nil && container.Resources.Limits != nil

// no requirement of specified extended resources

func GetContainerMilliCPULimit(c *corev1.Container) int64 { _ = "STUB: not implemented"; return 0 }

func GetContainerMemoryByteLimit(c *corev1.Container) int64 { _ = "STUB: not implemented"; return 0 }

func FindContainerIdAndStatusByName(status *corev1.PodStatus, name string) (string, *corev1.ContainerStatus, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func ParseContainerId(data string) (cType, cID string, err error) {
	_ = "STUB: not implemented"
	// Trim the quotes and split the type and ID.
	return "", "", nil
}
