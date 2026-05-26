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
)

// NOTE: functions in this file can be overwritten for extension

// GetPodCgroupParentDir gets the full pod cgroup parent with the pod info.
// @podKubeRelativeDir kubepods-burstable.slice/kubepods-burstable-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/
// @return kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/
func GetPodCgroupParentDir(pod *corev1.Pod) string { _ = "STUB: not implemented"; return "" }

func GetKubeQoSByCgroupParent(cgroupDir string) corev1.PodQOSClass {
	_ = "STUB: not implemented"
	return *new(corev1.PodQOSClass)
}

// @return like kubepods-burstable.slice/kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/
// /sys/fs/cgroup/blkio/kubepods.slice/kubepods-burstable.slice/kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice
func GetPodCgroupBlkIOAbsolutePath(podParentDir string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetPodQoSRelativePath gets the relative parent directory of a pod's qos class.
// @qosClass corev1.PodQOSBurstable
// @return kubepods.slice/kubepods-burstable.slice/
func GetPodQoSRelativePath(qosClass corev1.PodQOSClass) string {
	_ = "STUB: not implemented"
	return ""
}

// GetContainerCgroupParentDir gets the full container cgroup parent with the pod parent dir and the containerStatus.
// @parentDir kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/
// @return kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/****.scope
func GetContainerCgroupParentDir(podParentDir string, c *corev1.ContainerStatus) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetContainerCgroupParentDirByID gets the full container cgroup parent dir with the podParentDir and the container ID.
// @parentDir kubepods.slice/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/
// @return kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/****.scope
func GetContainerCgroupParentDirByID(podParentDir string, containerID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
