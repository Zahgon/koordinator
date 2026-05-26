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

// ParsePodID parse pod ID from the pod base path.
// e.g. 7712555c_ce62_454a_9e18_9ff0217b8941 from kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice
func ParsePodID(basename string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func GetPIDsInPod(podParentDir string, cs []corev1.ContainerStatus) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPodSandboxContainerID lists all dirs of pod cgroup(cpuset), exclude containers' dir, get sandbox hash id from the
// remaining dir.
// e.g. return "containerd://91cf0413ee0e6745335e9043b261a829ce07d28a5a66b5ec39b06811ef75a1ff"
func GetPodSandboxContainerID(pod *corev1.Pod) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// container not created, skip until container is created because container runtime is unknown

// get runtime type and dir names of known containers
