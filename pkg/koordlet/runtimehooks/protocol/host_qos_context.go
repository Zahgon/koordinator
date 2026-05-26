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

	ext "github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
)

type HostAppRequest struct {
	Name         string
	QOSClass     ext.QoSClass
	CgroupParent string
}

func (r *HostAppRequest) FromReconciler(hostAppSpec *slov1alpha1.HostApplicationSpec) {
	_ = "STUB: not implemented"
	return
}

type HostAppResponse struct {
	Resources Resources
}

type HostAppContext struct {
	Request  HostAppRequest
	Response HostAppResponse
	executor resourceexecutor.ResourceUpdateExecutor
	updaters []resourceexecutor.ResourceUpdater
}

func (c *HostAppContext) RecordEvent(r record.EventRecorder, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	//TODO: don't support record pod by host level
	return
}

func (c *HostAppContext) FromReconciler(hostAppSpec *slov1alpha1.HostApplicationSpec) {
	_ = "STUB: not implemented"
	return
}

func (c *HostAppContext) ReconcilerProcess(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (c *HostAppContext) ReconcilerDone(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (c *HostAppContext) GetUpdaters() []resourceexecutor.ResourceUpdater {
	_ = "STUB: not implemented"
	return nil
}

func (c *HostAppContext) Update() { _ = "STUB: not implemented"; return }

func (c *HostAppContext) injectForOrigin() {
	_ = "STUB: not implemented"
	// If CPUSet is not nil and is not an empty string, set cpuset
	return
}

func (c *HostAppContext) injectForExt() { _ = "STUB: not implemented"; return }
