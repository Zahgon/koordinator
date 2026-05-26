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

package protocol

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
)

type KubeQOSRequet struct {
	KubeQOSClass corev1.PodQOSClass
	CgroupParent string
}

func (r *KubeQOSRequet) FromReconciler(kubeQOS corev1.PodQOSClass) {
	_ = "STUB: not implemented"
	return
}

type KubeQOSResponse struct {
	Resources Resources
}

type KubeQOSContext struct {
	Request  KubeQOSRequet
	Response KubeQOSResponse
	executor resourceexecutor.ResourceUpdateExecutor
	updaters []resourceexecutor.ResourceUpdater
}

func (k *KubeQOSContext) RecordEvent(r record.EventRecorder, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	//TODO: Don't record pods by QoS
	return
}

func (k *KubeQOSContext) FromReconciler(kubeQOS corev1.PodQOSClass) {
	_ = "STUB: not implemented"
	return
}

// ReconcilerProcess generate the resource updaters but not do the update until the Update() is called.
func (k *KubeQOSContext) ReconcilerProcess(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (k *KubeQOSContext) ReconcilerDone(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (k *KubeQOSContext) GetUpdaters() []resourceexecutor.ResourceUpdater {
	_ = "STUB: not implemented"
	return nil
}

func (k *KubeQOSContext) Update() { _ = "STUB: not implemented"; return }

func (k *KubeQOSContext) injectForOrigin() {
	_ = "STUB: not implemented"
	// TODO
	return
}

func (k *KubeQOSContext) injectForExt() { _ = "STUB: not implemented"; return }
