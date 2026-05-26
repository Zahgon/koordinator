/*
Copyright 2022 The Koordinator Authors.
Copyright 2017 The Kubernetes Authors.

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

package kubelet

import (
	corev1 "k8s.io/api/core/v1"
	kubeletconfiginternal "k8s.io/kubernetes/pkg/kubelet/apis/config"
	"k8s.io/kubernetes/pkg/kubelet/cm/cpumanager/topology"
	evictionapi "k8s.io/kubernetes/pkg/kubelet/eviction/api"
	"k8s.io/utils/cpuset"

	koordletutil "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

func NewCPUTopology(cpuInfo *koordletutil.LocalCPUInfo) *topology.CPUTopology {
	_ = "STUB: not implemented"
	return nil
}

func GetStaticCPUManagerPolicyReservedCPUs(topology *topology.CPUTopology, kubeletConfiguration *kubeletconfiginternal.KubeletConfiguration) (cpuset.CPUSet, error) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), nil
}

// takeByTopology allocates CPUs associated with low-numbered cores from
// allCPUs.
//
// For example: Given a system with 8 CPUs available and HT enabled,
// if numReservedCPUs=2, then reserved={0,4}

func getNumReservedCPUs(nodeAllocatableReservation corev1.ResourceList) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// The static policy cannot initialize without this information.

// The static policy requires this to be nonzero. Zero CPU reservation
// would allow the shared pool to be completely exhausted. At that point
// either we would violate our guarantee of exclusivity or need to evict
// any pod that has at least one container that requires zero CPUs.
// See the comments in policy_static.go for more details.

// Take the ceiling of the reservation, since fractional CPUs cannot be
// exclusively allocated.

func GetKubeletReservedOptions(kubeletConfiguration *kubeletconfiginternal.KubeletConfiguration, topology *topology.CPUTopology) (reservedSystemCPUs cpuset.CPUSet, kubeReserved, systemReserved corev1.ResourceList, err error) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), *new(corev1.ResourceList), *new(corev1.ResourceList), nil
}

// at cmd option validation phase it is tested either --system-reserved-cgroup or --kube-reserved-cgroup is specified, so overwrite should be ok

func getReservedCPUs(topology *topology.CPUTopology, cpus string) (cpuset.CPUSet, error) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), nil
}

// parseResourceList parses the given configuration map into an API
// ResourceList or returns an error.
func parseResourceList(m map[string]string) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

// CPU, memory, local storage, and PID resources are supported.

func GetNodeAllocatableReservation(numCPUs int, totalMemoryInBytes int64, evictionHard map[string]string, systemReserved, kubeReserved corev1.ResourceList, experimentalNodeAllocatableIgnoreEvictionThreshold bool) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func getKubeletHardEvictionThresholds(evictionHard map[string]string, experimentalNodeAllocatableIgnoreEvictionThreshold bool) ([]evictionapi.Threshold, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the user requested to ignore eviction thresholds, then do not set valid values for hardEvictionThresholds here.

// hardEvictionReservation returns a resourcelist that includes reservation of resources based on hard eviction thresholds.
func hardEvictionReservation(thresholds []evictionapi.Threshold, capacity corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func GetCPUManagerStateFilePath(rootDirectory string) string { _ = "STUB: not implemented"; return "" }
