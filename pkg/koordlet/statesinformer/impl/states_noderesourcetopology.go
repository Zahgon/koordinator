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

package impl

import (
	"sync"

	"github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	topologyclientset "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/clientset/versioned"
	topologylister "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/listers/topology/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/kubelet/cm/cpumanager/topology"

	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/util"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

const (
	nodeTopoInformerName PluginName = "nodeTopoInformer"
)

var (
	// ResourceName list managed by the nodeTopoInformer
	managedNRTResources = []corev1.ResourceName{
		corev1.ResourceCPU,
		corev1.ResourceMemory,
		extension.ResourceGPU,
		corev1.ResourceHugePagesPrefix + "2Mi",
		corev1.ResourceHugePagesPrefix + "1Gi",
	}
	// Annotation keys list managed by the nodeTopoInformer
	managedNRTAnnotationKeys = []string{
		extension.AnnotationCPUBasicInfo,
		extension.AnnotationKubeletCPUManagerPolicy,
		extension.AnnotationNodeCPUSharedPools,
		extension.AnnotationNodeBECPUSharedPools,
		extension.AnnotationNodeCPUTopology,
		extension.AnnotationNodeCPUAllocs,
		extension.AnnotationNodeReservation,
		extension.AnnotationNodeSystemQOSResource,
	}

	// Label keys list managed by the nodeTopoInformer
	managedNRTLabelKeys = []string{
		extension.LabelNodeEnableNUMAReservation,
	}
	// Zone resources merging function
	mergeNRTZoneFn = func(nrt *v1alpha1.NodeResourceTopology, zoneList v1alpha1.ZoneList) v1alpha1.ZoneList {
		return util.MergeZoneList(util.TrimDifferentZone(nrt.Zones, zoneList), zoneList)
	}
	// NRT status extension function
	extendNRTStatusFn = func(_ *nodeTopologyStatus, _ *corev1.Node, _ int) error {
		return nil
	}
	// NRT status comparison function
	extendNRTEqualFn = func(_ *nodeTopologyStatus, _ *v1alpha1.NodeResourceTopology) (bool, string) {
		return true, ""
	}
)

type nodeTopologyStatus struct {
	Annotations    map[string]string
	Labels         map[string]string
	TopologyPolicy v1alpha1.TopologyManagerPolicy
	Zones          v1alpha1.ZoneList
}

func (n *nodeTopologyStatus) isChanged(oldNRT *v1alpha1.NodeResourceTopology) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// check TopologyPolicies

// check annotations

// check labels

// check Zones

// hook to compare extended status

func (n *nodeTopologyStatus) updateNRT(nrt *v1alpha1.NodeResourceTopology) {
	_ = "STUB: not implemented"
	return
}

// trim useless zone name and merge with the existing zone list

type nodeTopoInformer struct {
	config         *Config
	topologyClient topologyclientset.Interface
	nodeTopoMutex  sync.RWMutex
	nodeTopology   *v1alpha1.NodeResourceTopology

	metricCache    metriccache.MetricCache
	callbackRunner *callbackRunner

	nodeResourceTopologyInformer cache.SharedIndexInformer
	nodeResourceTopologyLister   topologylister.NodeResourceTopologyLister

	kubelet      KubeletStub
	nodeInformer *nodeInformer
	podsInformer *podsInformer
}

func NewNodeTopoInformer() *nodeTopoInformer { _ = "STUB: not implemented"; return nil }

func (s *nodeTopoInformer) GetNodeTopo() *v1alpha1.NodeResourceTopology {
	_ = "STUB: not implemented"
	return nil
}

func (s *nodeTopoInformer) Setup(ctx *PluginOption, state *PluginState) {
	_ = "STUB: not implemented"
	return
}

func (s *nodeTopoInformer) Start(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (s *nodeTopoInformer) HasSynced() bool {
	_ = "STUB: not implemented"
	// TODO only node cpu info collector relies on node topo informer
	return false
}

func newNodeResourceTopologyInformer(client topologyclientset.Interface, nodeName string) cache.SharedIndexInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedIndexInformer)
}

func (s *nodeTopoInformer) createNodeTopoIfNotExist() { _ = "STUB: not implemented"; return }

// TODO: add retry if create fail

// calcNodeTopo returns the calculated annotations, zone list, topology policy, error.
func (s *nodeTopoInformer) calcNodeTopo() (*nodeTopologyStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get CPUBasicInfo

// NOTE: The Koordlet is compatible with the Kubelet static cpu manager. Users can move to the Koordinator's CPU
// orchestration strategy by the following steps:
// 1. When the users want to keep the kubelet static cpu manager in use until the node is offline or removed, they
//    can disable the awareness and reporting of the Kubelet cpu manager by setting the DisableQueryKubeletConfig
//    to true. So the scheduler can allocate the cpuset cpus including the cpus managed by the kubelet static cpu
//    manager. After the node is offline or ready to go to the step 2, the users can reset the
//    DisableQueryKubeletConfig to false.
// 2. By default, the koordlet takes over the cpuset cpus for the new pods when DisableQueryKubeletConfig = false.
//    The remain cpuset pods managed by the kubelet static cpu manager are reported according to the /configz
//    and cpu_manager_state, so the related cpuset cpus are excluded from the cpu allocation of the scheduler.
//    For the newly-scheduled cpuset pods, the koordlet follows their cpuset allocation results on the annotations
//    that are exclusive to the remaining cpuset cpus managed by the kubelet static cpu manager. The users should
//    no longer use the kubelet static cpu manager anymore and should set the policy to "none". After the last pod
//    of the static cpu manager policy is terminated, the cpuset cpus will be fully managed by the koordlet.

// default policy is none

// NOTE: We should not remove reservedCPUs from sharedPoolCPUs to
//  ensure that Burstable Pods (e.g. Pods request 0C but are limited to 4C)
//  at least there are reservedCPUs available when nodes are allocated

// handle cpus allocated by the Kubelet cpu manager

// get NRT topology policy

// handle cpus reserved by annotation of node.

// handle cpus allocated for system qos of node

// TODO consider define in NodeSLO for system qos, annotation on node is provided as "Syntactic Sugar", which overlaps the NodeSLO for custom-definition

// "null" when the podAllocs is empty

// remove cpus that already reserved by node.annotation.

// remove cpus that exclusive for system qos from annotation

// set optional annotations

// sync managed labels

// hook to set extra status

// removeNodeReservedCPUs filter out cpus that reserved by annotation of node.
func removeNodeReservedCPUs(cpuSharePools []extension.CPUSharedPool, reservedCPUs cpuset.CPUSet) []extension.CPUSharedPool {
	_ = "STUB: not implemented"
	return nil
}

// removeSystemQOSCPUs filter out cpus that for system qos.
func removeSystemQOSCPUs(cpuSharePools []extension.CPUSharedPool, sysQOSRes *extension.SystemQOSResource) []extension.CPUSharedPool {
	_ = "STUB: not implemented"
	return nil
}

// system QoS resource not specified, or cpu is not exclusive

func getNodeReserved(cpuTopology *topology.CPUTopology, nodeAnnotations map[string]string) *extension.NodeReservation {
	_ = "STUB: not implemented"
	return nil
}

func (s *nodeTopoInformer) calGuaranteedCpu(usedCPUs map[int32]*extension.CPUInfo, stateJSON string) ([]extension.PodCPUAlloc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// entries can be empty when the kubelet cpu manager policy is none

// TODO: It is possible that the data in the checkpoint file is invalid
//  and should be checked with the data in the cgroup to determine whether it is consistent

func (s *nodeTopoInformer) reportNodeTopology() { _ = "STUB: not implemented"; return }

// do not CREATE if reporting is disabled,
// but update the node topo object internally

// TODO: merge the create and update

// avoid overwrite the cache

// update fields

// TODO need to separate the local update and remote update
// do local update

// do remote update

func isEqualNRTZones(oldZones, newZones v1alpha1.ZoneList) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// Zone name and type are maintained by the agent, while the resources field can be updated by other components.

// isEqualNRTAnnotations returns whether the new topology annotations has difference with the old one or not
func isEqualNRTAnnotations(oldAnno, newAnno map[string]string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// both not exist, no need to compare this key

// (oldExist = true, newExist = false) OR (oldExist = false, newExist = true), node topo not equal

// else both exist in new and old, compare value

func isEqualNRTLabels(oldLabels, newLabels map[string]string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (s *nodeTopoInformer) calCPUSharePools(lsSharedPoolCPUs map[int32]*extension.CPUInfo) (lsSharePools []extension.CPUSharedPool, beSharePools []extension.CPUSharedPool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func covertCPUsToSharePool(cpuIDMap map[int32]*extension.CPUInfo) (sharePools []extension.CPUSharedPool) {
	_ = "STUB: not implemented"
	// nodeID -> cpulist
	return nil
}

func (s *nodeTopoInformer) calCPUTopology() (*metriccache.NodeCPUInfo, *extension.CPUTopology, map[int32]*extension.CPUInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (s *nodeTopoInformer) calTopologyZoneList(nodeCPUInfo *metriccache.NodeCPUInfo) (v1alpha1.ZoneList, error) {
	_ = "STUB: not implemented"
	return *new(v1alpha1.ZoneList), nil
}

// Check if NUMA reservation is enabled

// Calculate kubelet reserved resources for NUMA-level deduction

// Build zone resources with capacity and allocatable

// Initialize allocatable as a copy of capacity

// Apply NUMA-level reservation if enabled

func (s *nodeTopoInformer) calKubeletAllocatedCPUs(sharePoolCPUs map[int32]*extension.CPUInfo) ([]extension.PodCPUAlloc, error) {
	_ = "STUB: not implemented"
	// Users can specify the kubelet RootDirectory on the host in the koordlet DaemonSet,
	// inside koordlet it is mounted to the path /var/lib/kubelet by default.
	return nil, nil
}

// TODO: report lse/lsr pod from cgroup

func (s *nodeTopoInformer) updateNodeTopo(newTopo *v1alpha1.NodeResourceTopology) {
	_ = "STUB: not implemented"
	return
}

func (s *nodeTopoInformer) setNodeTopo(newTopo *v1alpha1.NodeResourceTopology) {
	_ = "STUB: not implemented"
	return
}

func newNodeTopo(node *corev1.Node) *v1alpha1.NodeResourceTopology {
	_ = "STUB: not implemented"
	return nil
}

// fields are required

// getTopologyPolicy gets the NRT topology policy with the kubelet topology manager policy and scope.
func getTopologyPolicy(topologyManagerPolicy string, topologyManagerScope string) v1alpha1.TopologyManagerPolicy {
	_ = "STUB: not implemented"
	return *new(v1alpha1.TopologyManagerPolicy)
}

// getNodeKubeletReservedResources calculates the kubelet reserved resources on the node.
// It returns the difference between node capacity and allocatable resources.
func (s *nodeTopoInformer) getNodeKubeletReservedResources() corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// Calculate reserved resources: capacity - allocatable

// applyNUMAReservationToZoneResources applies kubelet reserved resources to NUMA-level zones evenly.
// It deducts ceil(reserved / numaCount) from each NUMA node's allocatable resources only.
func applyNUMAReservationToZoneResources(zoneResources map[string]util.ZoneResources, nodeReserved corev1.ResourceList, numaCount int) map[string]util.ZoneResources {
	_ = "STUB: not implemented"
	return nil
}

// Calculate per-NUMA reservation: ceil(reserved / numaCount)

// Apply reservation to each zone's allocatable (capacity remains unchanged)

// Deduct reservation from allocatable

// Keep the same format but set to zero

// Ensure capacity is not affected by explicitly copying it

// divideResourceByCount divides a resource quantity by count using ceiling division.
// This ensures that the total reserved across all NUMA nodes is at least the node-level reservation.
func divideResourceByCount(quantity resource.Quantity, count int) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// Use MilliValue for better precision with CPU resources

// Ceiling division: (milliValue + count - 1) / count
