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
	"k8s.io/client-go/tools/record"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	runtimeapi "github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type ContainerMeta struct {
	Name string
	ID   string // docker://xxx; containerd://
	// is sandbox container
	Sandbox bool
}

func (c *ContainerMeta) FromNri(container *api.Container, podAnnotations map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (c *ContainerMeta) FromProxy(containerMeta *runtimeapi.ContainerMetadata, podAnnotations map[string]string) {
	_ = "STUB: not implemented"
	return
}

type ContainerRequest struct {
	PodMeta           PodMeta
	ContainerMeta     ContainerMeta
	PodLabels         map[string]string
	PodAnnotations    map[string]string
	CgroupParent      string
	ContainerEnvs     map[string]string
	Resources         *Resources // supports proxy, reconciler & nri mode
	ExtendedResources *apiext.ExtendedResourceContainerSpec
}

func splitEnvVar(s string) (string, string) { _ = "STUB: not implemented"; return "", "" }

func (c *ContainerRequest) FromNri(pod *api.PodSandbox, container *api.Container) {
	_ = "STUB: not implemented"
	return
}

// extract Resources from NRI container

func (c *ContainerRequest) FromProxy(req *runtimeapi.ContainerResourceHookRequest) {
	_ = "STUB: not implemented"
	return
}

// extract Resources from proxy request

// retrieve ExtendedResources from pod annotations

func (c *ContainerRequest) FromReconciler(podMeta *statesinformer.PodMeta, containerName string, sandbox bool) {
	_ = "STUB: not implemented"
	return
}

// retrieve ExtendedResources from container spec and pod annotations (prefer container spec)

// specFromContainer == nil

type ContainerResponse struct {
	Resources           Resources
	AddContainerEnvs    map[string]string
	AddContainerMounts  []*Mount
	AddContainerDevices []*LinuxDevice
}

type LinuxDevice struct {
	Path          string
	Type          string
	Major         int64
	Minor         int64
	FileModeValue uint32
}

func (c *ContainerResponse) ProxyDone(resp *runtimeapi.ContainerResourceHookResponse) {
	_ = "STUB: not implemented"
	return
}

// resource value is injected but origin request is nil, init resource response

type ContainerContext struct {
	Request  ContainerRequest
	Response ContainerResponse
	executor resourceexecutor.ResourceUpdateExecutor
	updaters []resourceexecutor.ResourceUpdater
}

func (c *ContainerContext) RecordEvent(r record.EventRecorder, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	//TODO: Don't record pod by container level
	return
}

func (c *ContainerContext) FromNri(pod *api.PodSandbox, container *api.Container) {
	_ = "STUB: not implemented"
	return
}

func (c *ContainerContext) FromProxy(req *runtimeapi.ContainerResourceHookRequest) {
	_ = "STUB: not implemented"
	return
}

func (c *ContainerContext) ProxyDone(resp *runtimeapi.ContainerResourceHookResponse, executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (c *ContainerContext) NriDone(executor resourceexecutor.ResourceUpdateExecutor) (*api.ContainerAdjustment, *api.ContainerUpdate, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// todo: add more fields conversions

func (c *ContainerContext) FromReconciler(podMeta *statesinformer.PodMeta, containerName string, sandbox bool) {
	_ = "STUB: not implemented"
	return
}

// ReconcilerProcess generate the resource updaters but not do the update until the Update() is called.
func (c *ContainerContext) ReconcilerProcess(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (c *ContainerContext) ReconcilerDone(executor resourceexecutor.ResourceUpdateExecutor) {
	_ = "STUB: not implemented"
	return
}

func (c *ContainerContext) GetUpdaters() []resourceexecutor.ResourceUpdater {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContainerContext) Update() { _ = "STUB: not implemented"; return }

// Inject valid parameters in ContainerContext.Response.Resources,
// such as CPUShares, CPUSet, CFSQuota, MemoryLimit...
func (c *ContainerContext) injectForOrigin() {
	_ = "STUB: not implemented"
	// If CPUShares is not nil, set container cpu share
	return
}

// If CPUSet is not nil and is not an empty string, set container cpuset

// If CFSQuota is not nil, set container cfs quota

// If MemoryLimit is not nil, set container memory limit

// TODO other fields

func (c *ContainerContext) injectForExt() {
	_ = "STUB: not implemented"
	// TODO
	return
}

func getContainerID(podAnnotations map[string]string, containerUID string) string {
	_ = "STUB: not implemented"
	// TODO parse from runtime hook request directly such as cgroup path format
	return ""
}
