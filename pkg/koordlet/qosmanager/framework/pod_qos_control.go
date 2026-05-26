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

package framework

import (
	corev1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
)

type QOSPolicyType string

const (
	QOSPolicyCPUBurst  QOSPolicyType = "CPUBurst"
	QOSPolicyMemoryQOS QOSPolicyType = "MemoryQOS"
)

var globalQOSControlPlugins = map[string]QOSGreyControlPlugin{}

func RegisterQOSGreyCtrlPlugin(name string, plugin QOSGreyControlPlugin) error {
	_ = "STUB: not implemented"
	return nil
}

func RunQOSGreyCtrlPlugins(client clientset.Interface, stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func InjectQOSGreyCtrlPlugins(pod *corev1.Pod, policyType QOSPolicyType, policy *interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func UnregisterQOSGreyCtrlPlugin(name string) { _ = "STUB: not implemented"; return }

func ClearQOSGreyCtrlPlugin() { _ = "STUB: not implemented"; return }

type QOSGreyControlPlugin interface {
	Setup(kubeClient clientset.Interface) error
	Run(stopCh <-chan struct{})
	InjectPodPolicy(pod *corev1.Pod, policyType QOSPolicyType, greyCtlPolicy *interface{}) (bool, error)
}
