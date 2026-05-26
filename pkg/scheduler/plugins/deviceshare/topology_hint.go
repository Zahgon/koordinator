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

	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/topologymanager"
	"github.com/koordinator-sh/koordinator/pkg/util/bitmask"
)

const (
	ErrInsufficientNUMAScopedDevices = "Insufficient NUMA Scoped Devices"
	ErrDesignatedAllocationLackNUMA  = "Designated Allocation Lack NUMA Node ID"

	defaultNUMAScore = 500
)

func (p *Plugin) GetPodTopologyHints(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, node *corev1.Node) (map[string][]topologymanager.NUMATopologyHint, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateDesignatedHints(allocations apiext.DeviceAllocations, topology *NUMATopology) (map[string][]topologymanager.NUMATopologyHint, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) Allocate(ctx context.Context, cycleState fwktype.CycleState, affinity topologymanager.NUMATopologyHint, pod *corev1.Pod, node *corev1.Node) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) generateTopologyHints(cycleState fwktype.CycleState, state *preFilterState, nodeDevice *nodeDevice, node *corev1.Node, pod *corev1.Pod) (map[string][]topologymanager.NUMATopologyHint, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we just use a score bigger than 100 to make that device numa preference take precedence over cpu

// update hints preferred according to multiNUMAGroups, in case when it wasn't provided, the default
// behavior to prefer the minimal amount of NUMA nodes will be used

// no possible NUMA affinities for resource, just return status

type numaScopedAllocation struct {
	mask             bitmask.BitMask
	allocationResult apiext.DeviceAllocations
}

func hashAllocateResult(allocations apiext.DeviceAllocations) int {
	_ = "STUB: not implemented"
	return 0
}

func calcTotalDevicesByNUMA(nd *nodeDevice, numaNodes []int) map[schedulingv1alpha1.DeviceType]int {
	_ = "STUB: not implemented"
	return nil
}
