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

const (
	PodCgroupPathRelativeDepth       = 1
	ContainerCgroupPathRelativeDepth = 2
)

// GetRootCgroupCPUSetDir gets the cpuset parent directory of the specified podQos' root cgroup
// @output /sys/fs/cgroup/cpuset/kubepods.slice/kubepods-besteffort.slice
func GetRootCgroupCPUSetDir(qosClass corev1.PodQOSClass) string {
	_ = "STUB: not implemented"
	return ""
}

// GetBECgroupCurCPUSet gets the current cpuset of besteffort podQoS' cgroup.
func GetBECgroupCurCPUSet() ([]int32, error) { _ = "STUB: not implemented"; return nil, nil }

// find the minimal length of container's cpuset to avoid the interference of sandbox's cpuset
// since the cpuset's value of sandbox could always be set to all cpu ids of the machine by kubelet

// GetBECPUSetPathsByMaxDepth gets all the be cpuset groups' paths recursively from upper to lower
func GetBECPUSetPathsByMaxDepth(relativeDepth int) ([]string, error) {
	_ = "STUB: not implemented"
	// walk from root path to lower nodes
	return nil, nil
}

// make sure the rootCgroupPath is available

// get the path of parentDir

func GetCgroupPathsByTargetDepth(resourceType system.ResourceType, cgroupParent string, relativeDepth int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// make sure the rootCgroupPath is available

// get the path of parentDir

// GetBECPUSetPathsByTargetDepth only gets the be containers' cpuset groups' paths
func GetBECPUSetPathsByTargetDepth(relativeDepth int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCgroupRootBlkIOAbsoluteDir gets the root blkio directory
// @output /sys/fs/cgroup/blkio
func GetCgroupRootBlkIOAbsoluteDir() string { _ = "STUB: not implemented"; return "" }

// GetPodCgroupBlkIOAbsoluteDir gets the blkio parent directory of the specified podQos' root cgroup
// @output /sys/fs/cgroup/blkio/kubepods.slice/kubepods-besteffort.slice
func GetPodCgroupBlkIOAbsoluteDir(qosClass corev1.PodQOSClass) string {
	_ = "STUB: not implemented"
	return ""
}
