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
	topologyv1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	"github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/noderesource/framework"
)

var (
	updateNRTResourceSet = sets.NewString(string(extension.BatchCPU), string(extension.BatchMemory))
)

const (
	MidCPUThreshold                = "midCPUThreshold"
	MidMemoryThreshold             = "midMemoryThreshold"
	MidUnallocatedPercent          = "midUnallocatedPercent"
	MidStaticCPUReservedPercent    = "midStaticCPUReservedPercent"
	MidStaticMemoryReservedPercent = "midStaticMemoryReservedPercent"
	BatchCPUThreshold              = "batchCPUThreshold"
	BatchMemoryThreshold           = "batchMemoryThreshold"
)

func CalculateBatchResourceByPolicy(strategy *configuration.ColocationStrategy, nodeCapacity, nodeSafetyMargin, nodeReserved,
	systemUsed, podHPReq, podHPUsed, podHPMaxUsedReq corev1.ResourceList) (corev1.ResourceList, string, string) {
	_ = "STUB: not implemented"
	// Node(Batch).Alloc[usage] := Node.Total - Node.SafetyMargin - System.Used - sum(Pod(Prod/Mid).Used)
	// System.Used = max(Node.Used - Pod(All).Used, Node.Anno.Reserved, Node.Kubelet.Reserved)
	return *new(corev1.ResourceList), "", ""
}

// Node(Batch).Alloc[request] := Node.Total - Node.SafetyMargin - System.Reserved - sum(Pod(Prod/Mid).Request)
// System.Reserved = max(Node.Anno.Reserved, Node.Kubelet.Reserved)

// Node(Batch).Alloc[maxUsageRequest] := Node.Total - Node.SafetyMargin - System.Used - sum(max(Pod(Prod/Mid).Request, Pod(Prod/Mid).Used))

// while batchCPU/MEMThresholdPercent != nil, use BatchXXXThresholdPercent of nodeCapacity to limit batchAllocatable
//       Node(Batch).Alloc = min(nodeCapacity * BatchXXXThresholdPercent, Node(Batch).Alloc[Policy])
// while batchCPU/MEMThresholdPercent == nil
//       Node(Batch).Alloc = Node(Batch).Alloc[Policy]

// batch cpu support policy "usage" and "maxUsageRequest"

// use CalculatePolicy "usage" by default

// batch memory support policy "usage", "request" and "maxUsageRequest"

// use CalculatePolicy "usage" by default

func CalculateMidResourceByStaticMode(strategy *configuration.ColocationStrategy, nodeCapacity corev1.ResourceList, nodeName string) (*resource.Quantity, *resource.Quantity, string, string) {
	_ = "STUB: not implemented"
	return nil, nil, "", ""
}

func CalculateMidResourceByPolicy(strategy *configuration.ColocationStrategy, nodeCapacity, unallocated, nodeUnused corev1.ResourceList, allocatableMilliCPU, allocatableMemory int64,
	prodReclaimableCPU, prodReclaimableMemory *resource.Quantity, nodeName string) (*resource.Quantity, *resource.Quantity, string, string) {
	_ = "STUB: not implemented"
	return nil, nil, "", ""
}

// CPU need turn into milli value

func PrepareNodeForResource(node *corev1.Node, nr *framework.NodeResource, name corev1.ResourceName) {
	_ = "STUB: not implemented"
	return
}

// if the specified resource has no quantity

// TODO mv to post-calculate stage for merging multiple calculate results
// amplify batch cpu according to cpu normalization ratio

// skip for invalid ratio

// NOTE: extended resource would be validated as an integer, so it should be checked before the update

// GetPodMetricUsage gets pod usage from the PodMetricInfo
func GetPodMetricUsage(info *slov1alpha1.PodMetricInfo) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func GetHostAppHPUsed(resourceMetrics *framework.ResourceMetrics, resPriority extension.PriorityClass) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// consider higher priority usage for mid or batch allocatable
// now only support product and batch(hadoop-yarn) priority for host application

// GetHostAppMetricUsage gets host application usage from HostApplicationMetricInfo
func GetHostAppMetricUsage(info *slov1alpha1.HostApplicationMetricInfo) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// GetPodNUMARequestAndUsage returns the pod request and usage on each NUMA nodes.
// It averages the metrics over all sharepools when the pod does not allocate any sharepool or use all sharepools.
func GetPodNUMARequestAndUsage(pod *corev1.Pod, podRequest, podUsage corev1.ResourceList, numaNum int) ([]corev1.ResourceList, []corev1.ResourceList) {
	_ = "STUB: not implemented"
	// get pod NUMA allocation
	return nil, nil
}

// NOTE: For the pod which does not set NUMA-aware allocation policy, it has set particular cpuset cpus but may not
//       have NUMA allocation information in annotations. In this case, it can be inaccurate to average the
//       request/usage over all NUMA nodes.

// the number of allocated NUMA node

// The invalid allocated NUMA ids will be ignored since it cannot be successfully bind on the node either.

// share all NUMAs

func GetPodUnknownNUMAUsage(podUsage corev1.ResourceList, numaNum int) []corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

// GetNodeCapacity gets node capacity and filters out non-CPU and non-Mem resources
func GetNodeCapacity(node *corev1.Node) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// GetNodeSafetyMargin gets node-level safe-guarding reservation with the node's allocatable
func GetNodeSafetyMargin(strategy *configuration.ColocationStrategy, nodeCapacity corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func DivideResourceList(rl corev1.ResourceList, divisor float64) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func zoneResourceListHandler(a, b []corev1.ResourceList, zoneNum int,
	handleFn func(a corev1.ResourceList, b corev1.ResourceList) corev1.ResourceList) []corev1.ResourceList {
	_ = "STUB: not implemented"
	// assert len(a) == len(b) == zoneNum
	return nil
}

func AddZoneResourceList(a, b []corev1.ResourceList, zoneNum int) []corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

func MaxZoneResourceList(a, b []corev1.ResourceList, zoneNum int) []corev1.ResourceList {
	_ = "STUB: not implemented"
	return nil
}

func GetResourceListForCPUAndMemory(rl corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func MixResourceListCPUAndMemory(resourcesForCPU, resourcesForMemory corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func MinxZoneResourceListCPUAndMemory(resourcesForCPU, resourcesForMemory []corev1.ResourceList, zoneNum int) []corev1.ResourceList {
	_ = "STUB: not implemented"
	// assert len(a) == len(b) == zoneNum
	return nil
}

// getReserveRatio returns resource reserved ratio
func getReserveRatio(reclaimThreshold int64) float64 { _ = "STUB: not implemented"; return 0 }

func UpdateNRTZoneListIfNeeded(node *corev1.Node, zoneList topologyv1alpha1.ZoneList, nr *framework.NodeResource, diffThreshold float64) bool {
	_ = "STUB: not implemented"
	return false
}

// the resources of the zone should be reset

// FIXME: currently we set value to zero instead of deleting resource

// already reset

// amplify batch cpu according to cpu normalization ratio

// old has the resource key

// old has no resource key

// keep the resources order

func getCPUNormalizationRatio(nr *framework.NodeResource) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getPercentFromStrategy(strategy, defaultStrategy *configuration.ColocationStrategy, strategyType string) float64 {
	_ = "STUB: not implemented"
	return 0
}

func IsValidNodeUsage(nodeMetric *slov1alpha1.NodeMetric) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}
