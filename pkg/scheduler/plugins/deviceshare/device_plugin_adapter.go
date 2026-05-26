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

package deviceshare

import (
	"context"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corelister "k8s.io/client-go/listers/core/v1"
	"k8s.io/utils/clock"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

const (
	// AnnotationBindTimestamp represents the bind time (unix nano) of the pod which can be used by
	// custom device plugins which are adapted for koord-scheduler to determine pod.
	//
	// Background: Device plugins cannot get pod manifest (including annotations) from kubelet so they have to
	// use the bind time to determine the target pod when multiple pods are assigned to the same node at the same time.
	AnnotationBindTimestamp = apiext.SchedulingDomainPrefix + "/bind-timestamp"

	// AnnotationGPUMinors represents a comma separated minor list of GPU(s) allocated for the pod.
	// A typical use case is used as env field ref to override the built-in NVIDIA_VISIBLE_DEVICES=all env of CUDA images.
	//
	// Background: The envs injected by device plugin would be overridden by the ones of pod image, so sometimes users
	// need to explicitly declare the envs in pod template to override the ones of pod image.
	AnnotationGPUMinors = apiext.SchedulingDomainPrefix + "/gpu-minors"

	// AnnotationHAMiLock is annotated to node by scheduler to prevent scheduling multiple pods
	// with devices of the same vendor to the same node which would make the vendor's device plugin unable to determine pod.
	// The vendor's device plugin will automatically remove it after allocation of a pod.
	// This annotation key was originally defined in HAMi and then be adopted by some vendors.
	// Vendors that currently need this lock: metax
	AnnotationHAMiLock = "hami.io/mutex.lock"

	// AnnotationPredicateTime represents the bind time (unix nano) of the pod which is used by
	// Huawei NPU device plugins to determine pod.
	AnnotationPredicateTime = "predicate-time"
	// AnnotationHuaweiNPUCore represents the NPU/vNPU allocation result for the pod which is used by
	// Huawei NPU device plugins to allocate NPU(s)/vNPU for pod.
	AnnotationHuaweiNPUCore = "huawei.com/npu-core"
	// AnnotationHuaweiAscend310P represents the NPU allocation result for the pod which is used by
	// Huawei Ascend 310P device plugins to allocate NPU(s) for pod.
	AnnotationHuaweiAscend310P = "huawei.com/Ascend310P"

	// AnnotationCambriconDsmluAssigned is used by Cambricon dynamic sMLU device plugins to determine pod.
	AnnotationCambriconDsmluAssigned = "CAMBRICON_DSMLU_ASSIGHED"
	// AnnotationCambriconDsmluProfile represents the dynamic sMLU allocation result for the pod which is used by
	// Cambricon dynamic sMLU device plugins to allocate sMLU for pod.
	AnnotationCambriconDsmluProfile = "CAMBRICON_DSMLU_PROFILE"
	// AnnotationCambriconDsmluLock is annotated to node by scheduler to prevent scheduling multiple pods
	// with dynamic sMLU to the same node which would make Cambricon device plugin unable to determine pod.
	// Cambricon dynamic sMLU device plugin will automatically remove it after allocation of a pod.
	AnnotationCambriconDsmluLock = "cambricon.com/dsmlu.lock"
	// cambriconVMemoryUnit is the minial virtual memory unit of Cambricon dynamic sMLU currently supported.
	// This should always match the min-dsmlu-unit arg of device plugin.
	cambriconVMemoryUnit = 256 * 1024 * 1024

	// AnnotationMetaXGPUDevicesAllocated represents the GPU/sGPU allocation result for the pod which is used by
	// MetaX GPU device plugins to allocate GPU(s)/sGPU(s) for pod.
	AnnotationMetaXGPUDevicesAllocated = "metax-tech.com/gpu-devices-allocated"
	// metaxVRamUnit is the minial virtual memory unit of MetaX sGPU currently supported.
	metaxVRamUnit = 1 * 1024 * 1024
)

const (
	nodeLockTimeout = 5 * time.Minute
)

// DevicePluginAdapter adapt koord-scheduler's device allocation result to the format that third party vendor's
// device plugin recognizes, so that user can directly use them as allocators with koord-scheduler without modification.
// This is especially useful when third party device plugin natively supports fine-grained device allocation or virtualization.
type DevicePluginAdapter interface {
	Adapt(ctx *DevicePluginAdaptContext, object metav1.Object, allocation []*apiext.DeviceAllocation) error
}

// DevicePluginAdapterWithNodeLock is DevicePluginAdapter which utilize node lock to prevent device plugin
// from unable to determine pod.
type DevicePluginAdapterWithNodeLock interface {
	DevicePluginAdapter
	NodeLockKey() string
}

type DevicePluginAdaptContext struct {
	context.Context
	nodeLister   corelister.NodeLister
	podLister    corelister.PodLister
	node         *corev1.Node
	modifiedNode *corev1.Node
	nodeLock     *sync.Mutex
	gpuModel     string
}

var (
	defaultDevicePluginAdapter    DevicePluginAdapter = &generalDevicePluginAdapter{}
	defaultDevicePluginAdapterMap                     = map[schedulingv1alpha1.DeviceType]DevicePluginAdapter{
		schedulingv1alpha1.GPU: &generalGPUDevicePluginAdapter{},
	}
	gpuDevicePluginAdapterMap = map[string]DevicePluginAdapter{
		apiext.GPUVendorHuawei:    &huaweiGPUDevicePluginAdapter{},
		apiext.GPUVendorCambricon: &cambriconGPUDevicePluginAdapter{},
		apiext.GPUVendorMetaX:     &metaxDevicePluginAdapter{},
	}
)

var (
	dpAdapterClock clock.Clock = clock.RealClock{} // for testing

	nodeLockMap     = make(map[string]*sync.Mutex)
	nodeLockMapLock sync.Mutex
)

func (p *Plugin) adaptForDevicePlugin(ctx context.Context, object metav1.Object, allocationResult apiext.DeviceAllocations, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func adapt(ctx *DevicePluginAdaptContext, adapter DevicePluginAdapter, object metav1.Object, allocation []*apiext.DeviceAllocation, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

// always get latest node here to avoid lock conflict

func getNodeLock(nodeName string) *sync.Mutex { _ = "STUB: not implemented"; return nil }

func deleteNodeLock(nodeName string) { _ = "STUB: not implemented"; return }

func lockNode(ctx *DevicePluginAdaptContext, object metav1.Object, lockKey string, lockTime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// this should never happen in normal cases, and we don't want to lock the node forever

// original locker pod has been deleted, safe to unlock

func unlockNode(node *corev1.Node, lockKey string) { _ = "STUB: not implemented"; return }

// generalDevicePluginAdapter annotates the bind timestamp to pod which enables users to write their own device plugins
// that can be used with koord-scheduler.
type generalDevicePluginAdapter struct{}

func (a *generalDevicePluginAdapter) Adapt(_ *DevicePluginAdaptContext, object metav1.Object, _ []*apiext.DeviceAllocation) error {
	_ = "STUB: not implemented"
	return nil
}

// generalGPUDevicePluginAdapter annotates necessary information to pod which resolves problems when using device plugin
// with koord-scheduler.
type generalGPUDevicePluginAdapter struct {
}

func (a *generalGPUDevicePluginAdapter) Adapt(ctx *DevicePluginAdaptContext, object metav1.Object, allocation []*apiext.DeviceAllocation) error {
	_ = "STUB: not implemented"
	return nil
}

type huaweiGPUDevicePluginAdapter struct{}

func (a *huaweiGPUDevicePluginAdapter) Adapt(ctx *DevicePluginAdaptContext, object metav1.Object, allocation []*apiext.DeviceAllocation) error {
	_ = "STUB: not implemented"
	return nil
}

// vNPU

// full NPU

type cambriconGPUDevicePluginAdapter struct{}

func (a *cambriconGPUDevicePluginAdapter) Adapt(_ *DevicePluginAdaptContext, object metav1.Object, allocation []*apiext.DeviceAllocation) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *cambriconGPUDevicePluginAdapter) NodeLockKey() string {
	_ = "STUB: not implemented"
	return ""
}

type metaxDevicePluginAdapter struct{}

type metaxContainerDeviceRequest struct {
	UUID    string `json:"uuid"`
	Compute int32  `json:"compute"`
	VRam    int32  `json:"vRam"`
}

func (a *metaxDevicePluginAdapter) Adapt(_ *DevicePluginAdaptContext, object metav1.Object, allocation []*apiext.DeviceAllocation) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *metaxDevicePluginAdapter) NodeLockKey() string { _ = "STUB: not implemented"; return "" }

func buildGPUMinorsStr(allocation []*apiext.DeviceAllocation, prefix string) string {
	_ = "STUB: not implemented"
	return ""
}
