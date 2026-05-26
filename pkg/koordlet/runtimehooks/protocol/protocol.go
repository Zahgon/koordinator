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

	runtimeapi "github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/audit"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type HooksProtocol interface {
	ReconcilerDone(executor resourceexecutor.ResourceUpdateExecutor)
	Update()
	GetUpdaters() []resourceexecutor.ResourceUpdater
	RecordEvent(r record.EventRecorder, pod *corev1.Pod)
}

type hooksProtocolBuilder struct {
	KubeQOS   func(kubeQOS corev1.PodQOSClass) HooksProtocol
	Pod       func(podMeta *statesinformer.PodMeta) HooksProtocol
	Sandbox   func(podMeta *statesinformer.PodMeta) HooksProtocol
	Container func(podMeta *statesinformer.PodMeta, containerName string) HooksProtocol
	HostApp   func(hostAppSpec *slov1alpha1.HostApplicationSpec) HooksProtocol
}

var HooksProtocolBuilder = hooksProtocolBuilder{
	KubeQOS: func(kubeQOS corev1.PodQOSClass) HooksProtocol {
		k := &KubeQOSContext{}
		k.FromReconciler(kubeQOS)
		return k
	},
	Pod: func(podMeta *statesinformer.PodMeta) HooksProtocol {
		p := &PodContext{}
		p.FromReconciler(podMeta)
		return p
	},
	Sandbox: func(podMeta *statesinformer.PodMeta) HooksProtocol {
		p := &ContainerContext{}
		p.FromReconciler(podMeta, "", true)
		return p
	},
	Container: func(podMeta *statesinformer.PodMeta, containerName string) HooksProtocol {
		c := &ContainerContext{}
		c.FromReconciler(podMeta, containerName, false)
		return c
	},
	HostApp: func(hostAppSpec *slov1alpha1.HostApplicationSpec) HooksProtocol {
		c := &HostAppContext{}
		c.FromReconciler(hostAppSpec)
		return c
	},
}

type Resctrl struct {
	Schemata   string
	Closid     string
	NewTaskIds []int32
}

type Resources struct {
	// origin resources
	CPUShares     *int64
	CFSQuota      *int64
	CPUSet        *string
	MemoryLimit   *int64
	NetClsClassId *uint32

	// extended resources
	CPUBvt  *int64
	CPUIdle *int64
	Resctrl *Resctrl
}

func (r *Resources) IsOriginResSet() bool { _ = "STUB: not implemented"; return false }

func (r *Resources) FromPod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (r *Resources) FromContainer(container *corev1.Container) { _ = "STUB: not implemented"; return }

// FromLinuxContainerResources extracts resource information from LinuxContainerResources
func (r *Resources) FromLinuxContainerResources(resources *runtimeapi.LinuxContainerResources) {
	_ = "STUB: not implemented"
	return
}

// MemoryLimitInBytes: 0 means not specified, negative values mean unlimited (-1)

// FromNriLinuxResources extracts resource information from NRI's LinuxResources
func (r *Resources) FromNriLinuxResources(resources *api.LinuxResources) {
	_ = "STUB: not implemented"
	return
}

// Extract CPU resources

// Extract Memory resources

type Mount struct {
	Destination string   `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	Type        string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Source      string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Options     []string `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
}

func injectCPUShares(cgroupParent string, cpuShares int64, a *audit.EventHelper, e resourceexecutor.ResourceUpdateExecutor) (resourceexecutor.ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(resourceexecutor.ResourceUpdater), nil
}

func injectCPUSet(cgroupParent string, cpuset string, a *audit.EventHelper, e resourceexecutor.ResourceUpdateExecutor) (resourceexecutor.ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(resourceexecutor.ResourceUpdater), nil
}

func injectCPUQuota(cgroupParent string, cpuQuota int64, a *audit.EventHelper, e resourceexecutor.ResourceUpdateExecutor) (resourceexecutor.ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(resourceexecutor.ResourceUpdater), nil
}

func injectMemoryLimit(cgroupParent string, memoryLimit int64, a *audit.EventHelper, e resourceexecutor.ResourceUpdateExecutor) (resourceexecutor.ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(resourceexecutor.ResourceUpdater), nil
}

func injectCPUBvt(cgroupParent string, bvtValue int64, a *audit.EventHelper, e resourceexecutor.ResourceUpdateExecutor) (resourceexecutor.ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(resourceexecutor.ResourceUpdater), nil
}

func injectCPUIdle(cgroupParent string, idleValue int64, a *audit.EventHelper, e resourceexecutor.ResourceUpdateExecutor) (resourceexecutor.ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(resourceexecutor.ResourceUpdater), nil
}

func injectNetClsClassId(cgroupParent string, classId uint32, a *audit.EventHelper, e resourceexecutor.ResourceUpdateExecutor) (resourceexecutor.ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(resourceexecutor.ResourceUpdater), nil
}

func createCatGroup(closid string, a *audit.EventHelper, e resourceexecutor.ResourceUpdateExecutor) (resourceexecutor.ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(resourceexecutor.ResourceUpdater), nil
}

func injectResctrl(closid string, schemata string, e *audit.EventHelper, executor resourceexecutor.ResourceUpdateExecutor) (resourceexecutor.ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(resourceexecutor.ResourceUpdater), nil
}
