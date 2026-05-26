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
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/informers"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

const (
	defaultGCPeriod = 3 * time.Second
)

type nodeDevice struct {
	lock          sync.RWMutex
	deviceTotal   map[schedulingv1alpha1.DeviceType]deviceResources
	deviceFree    map[schedulingv1alpha1.DeviceType]deviceResources
	deviceUsed    map[schedulingv1alpha1.DeviceType]deviceResources
	vfAllocations map[schedulingv1alpha1.DeviceType]*VFAllocation
	allocateSet   map[schedulingv1alpha1.DeviceType]map[types.NamespacedName]deviceResources
	deviceInfos   map[schedulingv1alpha1.DeviceType][]*schedulingv1alpha1.DeviceInfo

	numaTopology               *NUMATopology
	secondaryDeviceWellPlanned bool

	nodeHonorGPUPartition bool
	gpuPartitionIndexer   GPUPartitionIndexer
	gpuTopologyScope      *GPUTopologyScope
}

type VFAllocation struct {
	allocatedVFs map[int]sets.String
}

func newNodeDevice() *nodeDevice { _ = "STUB: not implemented"; return nil }

func (n *nodeDevice) getNodeDeviceSummary() *NodeDeviceSummary {
	_ = "STUB: not implemented"
	return nil
}

func (n *nodeDevice) resetDeviceTotal(resources map[schedulingv1alpha1.DeviceType]deviceResources) {
	_ = "STUB: not implemented"
	return
}

// updateCacheUsed is used to update deviceUsed when there is a new pod created/deleted
func (n *nodeDevice) updateCacheUsed(deviceAllocations apiext.DeviceAllocations, pod *corev1.Pod, add bool) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeDevice) getUsed(namespace, name string) map[schedulingv1alpha1.DeviceType]deviceResources {
	_ = "STUB: not implemented"
	return nil
}

// use a shallow copy to reduce overhead

func (n *nodeDevice) resetDeviceFree(deviceType schedulingv1alpha1.DeviceType) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeDevice) updateDeviceUsed(deviceType schedulingv1alpha1.DeviceType, allocations []*apiext.DeviceAllocation, add bool) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeDevice) isValid(deviceType schedulingv1alpha1.DeviceType, namespace string, name string, add bool) bool {
	_ = "STUB: not implemented"
	return false
}

// for non-failover scenario, pod might already exist in cache after Reserve step.

func (n *nodeDevice) updateAllocateSet(deviceType schedulingv1alpha1.DeviceType, allocations []*apiext.DeviceAllocation, pod *corev1.Pod, add bool) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeDevice) updateCacheVFAllocations(deviceType schedulingv1alpha1.DeviceType, allocations []*apiext.DeviceAllocation, add bool) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeDevice) updateVFAllocations(deviceType schedulingv1alpha1.DeviceType, vfAllocations *VFAllocation) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeDevice) removeVFAllocations(deviceType schedulingv1alpha1.DeviceType, vfAllocations *VFAllocation) {
	_ = "STUB: not implemented"
	return
}

func getVFAllocations(allocations []*apiext.DeviceAllocation) *VFAllocation {
	_ = "STUB: not implemented"
	return nil
}

func (n *nodeDevice) calcFreeWithPreemptible(deviceType schedulingv1alpha1.DeviceType, preemptible, requiredDeviceResources deviceResources) deviceResources {
	_ = "STUB: not implemented"
	return *new(deviceResources)
}

// The merging logic is executed only when there is a device that can be preempted,
// and the remaining idle devices are merged together to participate in the allocation

// If allocating from a required resources, e.g. a reservation, the free should be no larger than the reserved free.

func (n *nodeDevice) filter(
	devices map[schedulingv1alpha1.DeviceType][]int,
	hints apiext.DeviceAllocateHints,
	requiredDeviceResources, preemptibleDeviceResources map[schedulingv1alpha1.DeviceType]deviceResources,
) *nodeDevice {
	_ = "STUB: not implemented"
	return nil
}

func filterFreeDevicesByPCIe(n *nodeDevice, freeDeviceResources deviceResources, deviceType schedulingv1alpha1.DeviceType) sets.Int {
	_ = "STUB: not implemented"
	return *new(sets.Int)
}

func (n *nodeDevice) split(requestsPerInstance corev1.ResourceList, deviceType schedulingv1alpha1.DeviceType) deviceResources {
	_ = "STUB: not implemented"
	return *new(deviceResources)
}

type nodeDeviceCache struct {
	lock sync.RWMutex
	// nodeDeviceInfos stores nodeDevice for each node.
	nodeDeviceInfos map[string]*nodeDevice
}

func newNodeDeviceCache() *nodeDeviceCache { _ = "STUB: not implemented"; return nil }

func (n *nodeDeviceCache) getNodeDevice(nodeName string, needInit bool) *nodeDevice {
	_ = "STUB: not implemented"
	return nil
}

// getNodeDevice will create new `nodeDevice` if needInit is true and nodeDeviceInfos[nodeName] is nil

func (n *nodeDeviceCache) removeNodeDevice(nodeName string) { _ = "STUB: not implemented"; return }

func (n *nodeDeviceCache) invalidateNodeDevice(device *schedulingv1alpha1.Device) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeDeviceCache) updateNodeDevice(nodeName string, device *schedulingv1alpha1.Device) {
	_ = "STUB: not implemented"
	return
}

func buildDeviceResources(device *schedulingv1alpha1.Device) map[schedulingv1alpha1.DeviceType]deviceResources {
	_ = "STUB: not implemented"
	return nil
}

func (n *nodeDeviceCache) getNodeDeviceSummary(nodeName string) (*NodeDeviceSummary, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (n *nodeDeviceCache) getAllNodeDeviceSummary() map[string]*NodeDeviceSummary {
	_ = "STUB: not implemented"
	return nil
}

func (n *nodeDeviceCache) gcNodeDevice(ctx context.Context, informerFactory informers.SharedInformerFactory, period time.Duration) {
	_ = "STUB: not implemented"
	// Wait for all koord plugin event handlers (pod / device / reservation /
	// node, etc.) to finish processing their initial list before starting GC.
	// A bare cache.WaitForCacheSync(nodeInformer.HasSynced) only guarantees the
	// underlying store is populated; it does NOT guarantee that OnAdd callbacks
	// (which build nodeDeviceInfos) have run to completion. Without this wait
	// the first GC tick may operate on a half-populated nodeDeviceInfos map
	// (or, worse, see a warm nodeDeviceInfos but an empty node lister) and
	// mistakenly evict entries that are about to be recreated.
	// WaitForHandlersSync defers on ResourceEventHandlerRegistration.HasSynced
	// for every registration collected by ForceSyncFromInformer, which is the
	// same signal the scheduler's Run() waits on before accepting scheduling
	// cycles -- effectively delaying GC until "after scheduler Run".
	return
}
