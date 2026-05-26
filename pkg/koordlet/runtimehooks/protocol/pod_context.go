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
	"github.com/containerd/nri/pkg/api"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	runtimeapi "github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type PodMeta struct {
	Namespace string
	Name      string
	UID       string
}

func (p *PodMeta) String() string { _ = "STUB: not implemented"; return "" }

func (p *PodMeta) FromNri(pod *api.PodSandbox) { _ = "STUB: not implemented"; return }

func (p *PodMeta) FromProxy(meta *runtimeapi.PodSandboxMetadata) { _ = "STUB: not implemented"; return }

func (p *PodMeta) FromReconciler(meta metav1.ObjectMeta) { _ = "STUB: not implemented"; return }

type PodRequest struct {
	PodMeta           PodMeta
	Labels            map[string]string
	Annotations       map[string]string
	CgroupParent      string
	Resources         *Resources // TODO: support proxy & nri mode
	ExtendedResources *apiext.ExtendedResourceSpec
	ContainerTaskIds  map[string][]int32
}

func (p *PodRequest) FromNri(pod *api.PodSandbox) { _ = "STUB: not implemented"; return }

// retrieve ExtendedResources from pod annotations

func (p *PodRequest) FromProxy(req *runtimeapi.PodSandboxHookRequest) {
	_ = "STUB: not implemented"
	return
}

// retrieve ExtendedResources from pod annotations

func (p *PodRequest) FromReconciler(podMeta *statesinformer.PodMeta) {
	_ = "STUB: not implemented"
	return
}

// retrieve ExtendedResources from pod spec and pod annotations (prefer pod spec)

// specFromPod == nil

type RecorderEvent struct {
	HookName  string
	MsgFmt    string
	Reason    string
	EventType string
}

type PodResponse struct {
	Resources Resources
}

type PodContext struct {
	Request        PodRequest
	Response       PodResponse
	executor       resourceexecutor.ResourceUpdateExecutor
	updaters       []resourceexecutor.ResourceUpdater
	RecorderEvents []RecorderEvent
}

func (p *PodContext) RecordEvent(r record.EventRecorder, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	// Noraml, Warning => RecordEvent
	return
}

func (p *PodResponse) ProxyDone(resp *runtimeapi.PodSandboxHookResponse) {
	_ = "STUB: not implemented"
	return
}

// resource value is injected but origin request is nil, init resource response

func (p *PodContext) FromNri(pod *api.PodSandbox) { _ = "STUB: not implemented"; return }

func (p *PodContext) FromProxy(req *runtimeapi.PodSandboxHookRequest) {
	_ = "STUB: not implemented"
	return
}

func (p *PodContext) ProxyDone(resp *runtimeapi.PodSandboxHookResponse, executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (p *PodContext) NriDone(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (p *PodContext) NriRemoveDone(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (p *PodContext) FromReconciler(podMeta *statesinformer.PodMeta) {
	_ = "STUB: not implemented"
	return
}

// ReconcilerProcess generate the resource updaters but not do the update until the Update() is called.
func (p *PodContext) ReconcilerProcess(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (p *PodContext) ReconcilerDone(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (p *PodContext) GetUpdaters() []resourceexecutor.ResourceUpdater {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodContext) Update() { _ = "STUB: not implemented"; return }

func (p *PodContext) injectForOrigin() {
	_ = "STUB: not implemented"
	// TODO
	return
}

func (p *PodContext) injectForExt() { _ = "STUB: not implemented"; return }

// some of pod-level cgroups are manually updated since pod-stage hooks do not support it;
// kubelet may set the cgroups when pod is created or restarted, so we need to update the cgroups repeatedly

func (p *PodContext) removeForExt() { _ = "STUB: not implemented"; return }
