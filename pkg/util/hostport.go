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
	"k8s.io/kube-scheduler/framework"
)

func RequestedHostPorts(pod *corev1.Pod) framework.HostPortInfo {
	_ = "STUB: not implemented"
	return *new(framework.HostPortInfo)
}

func ResetHostPorts(pod *corev1.Pod, ports framework.HostPortInfo) {
	_ = "STUB: not implemented"
	return
}

func CloneHostPorts(ports framework.HostPortInfo) framework.HostPortInfo {
	_ = "STUB: not implemented"
	return *new(framework.HostPortInfo)
}

func AppendHostPorts(ports framework.HostPortInfo, r framework.HostPortInfo) framework.HostPortInfo {
	_ = "STUB: not implemented"
	return *new(framework.HostPortInfo)
}

func RemoveHostPorts(ports framework.HostPortInfo, r framework.HostPortInfo) {
	_ = "STUB: not implemented"
	return
}
