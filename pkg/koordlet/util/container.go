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

	"github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

// GetContainerCgroupPath gets the file path of the given container's cgroup.
// @parentDir kubepods-burstable.slice/kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/
// @return /sys/fs/cgroup/cpu/kubepods.slice/kubepods-burstable.slice/kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/cgroup.procs
func GetContainerCgroupPath(podParentDir string, c *corev1.ContainerStatus, resourceType system.ResourceType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// @parentDir kubepods-burstable.slice/kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/
// @return /sys/fs/cgroup/cpu/kubepods.slice/kubepods-burstable.slice/kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/cgroup.procs
func GetContainerCgroupCPUProcsPath(podParentDir string, c *corev1.ContainerStatus) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetContainerCgroupPerfPath(podParentDir string, c *corev1.ContainerStatus) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func GetContainerBaseCFSQuota(container *corev1.Container) int64 {
	_ = "STUB: not implemented"
	return 0
}

// ParseContainerID parse container ID from the container base path.
// e.g. 7712555c_ce62_454a_9e18_9ff0217b8941 from docker-7712555c_ce62_454a_9e18_9ff0217b8941.scope
func ParseContainerID(basename string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func IsValidContainerCgroupDir(containerParentDir string) bool {
	_ = "STUB: not implemented"
	return false
}

func GetPIDsInContainer(podParentDir string, c *corev1.ContainerStatus) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
