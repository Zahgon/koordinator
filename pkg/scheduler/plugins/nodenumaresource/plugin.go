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

package nodenumaresource

import (
	"context"
	"sync"

	nrtinformers "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions"
	topologylister "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/listers/topology/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	schedulingconfig "github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/topologymanager"
	"github.com/koordinator-sh/koordinator/pkg/util/bitmask"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

const (
	Name     = "NodeNUMAResource"
	stateKey = Name
)

const (
	ErrNotMatchNUMATopology         = "node(s) NUMA Topology policy not match"
	ErrInvalidRequestedCPUs         = "the requested CPUs must be integer"
	ErrInvalidCPUTopology           = "node(s) invalid CPU Topology"
	ErrSMTAlignmentError            = "node(s) requested cpus not multiple cpus per core"
	ErrCPUBindPolicyConflict        = "node(s) cpu bind policy conflicts with pod's required cpu bind policy"
	ErrInvalidCPUAmplificationRatio = "node(s) invalid CPU amplification ratio"
	ErrInsufficientAmplifiedCPU     = "Insufficient amplified cpu"
	ErrNotEnoughCPUs                = "not enough cpus available to satisfy request"
)

var (
	_ fwktype.EnqueueExtensions = &Plugin{}

	_ fwktype.PreFilterPlugin = &Plugin{}
	_ fwktype.FilterPlugin    = &Plugin{}
	_ fwktype.PreScorePlugin  = &Plugin{}
	_ fwktype.ScorePlugin     = &Plugin{}
	_ fwktype.ReservePlugin   = &Plugin{}
	_ fwktype.PreBindPlugin   = &Plugin{}

	_ frameworkext.ReservationRestorePlugin              = &Plugin{}
	_ frameworkext.ReservationPreAllocationRestorePlugin = &Plugin{}
	_ frameworkext.ReservationFilterPlugin               = &Plugin{}
	_ frameworkext.ReservationPreBindPlugin              = &Plugin{}
	_ topologymanager.NUMATopologyHintProvider           = &Plugin{}
)

type Plugin struct {
	handle                 frameworkext.ExtendedHandle
	pluginArgs             *schedulingconfig.NodeNUMAResourceArgs
	nrtInformerFactory     nrtinformers.SharedInformerFactory
	nrtLister              topologylister.NodeResourceTopologyLister
	scorer                 *resourceAllocationScorer
	numaScorer             *resourceAllocationScorer
	resourceManager        ResourceManager
	topologyOptionsManager TopologyOptionsManager
}

type Option func(*pluginOptions)

type pluginOptions struct {
	topologyOptionsManager TopologyOptionsManager
	resourceManager        ResourceManager
}

func WithTopologyOptionsManager(topologyOptionsManager TopologyOptionsManager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithResourceManager(resourceManager ResourceManager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NewWithOptions(args runtime.Object, handle fwktype.Handle, opts ...Option) (fwktype.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.Plugin), nil
}

func New(_ context.Context, args runtime.Object, handle fwktype.Handle) (fwktype.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.Plugin), nil
}

func (p *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) GetResourceManager() ResourceManager {
	_ = "STUB: not implemented"
	return *new(ResourceManager)
}

func (p *Plugin) GetTopologyOptionsManager() TopologyOptionsManager {
	_ = "STUB: not implemented"
	return *new(TopologyOptionsManager)
}

// schedulingStateData is the data only kept in the scheduling cycle. It could be cleaned up
// before entering the binding cycle to reduce memory cost.
type schedulingStateData struct {
	lock             sync.RWMutex
	preemptibleState map[string]*preemptibleNodeState
}

type preFilterState struct {
	schedulingStateData
	skip                        bool                                // whether the pod should skip the scheduling by this plugin
	requestCPUBind              bool                                // whether the pod requires cpu binding (e.g. qos=LSE and requests.cpu > 0)
	requests                    corev1.ResourceList                 // the resource requests of the pod
	requiredCPUBindPolicy       schedulingconfig.CPUBindPolicy      // the required binding policy if specified
	preferredCPUBindPolicy      schedulingconfig.CPUBindPolicy      // the preferred binding policy if specified
	preferredCPUExclusivePolicy schedulingconfig.CPUExclusivePolicy // the preferred exclusive policy if specified
	podNUMATopologyPolicy       extension.NUMATopologyPolicy        // the pod-level NUMA topology policy if specified
	podNUMAExclusive            extension.NumaTopologyExclusive     // the pod-level NUMA exclusive policy if specified
	numCPUsNeeded               int                                 // the number of requested CPUs
	allocation                  *PodAllocation                      // the CPU allocation reserved for the pod
	hasReservationAffinity      bool                                // whether the pod has a required reservation affinity
	isPreAllocationRequired     bool                                // whether the reserve pod requires pre-allocation

	// designatedAllocation is parsed from the Pod Annotation during the PreFilter phase. In this case, we should assume that the Node has already been selected. That is, all Score-related plug-ins will not be executed. Instead, we only need to call Filter to confirm whether the designatedAllocation is still valid and use it as the allocation result during actual allocation.
	designatedAllocation *allocation
}

type allocation struct {
	numaNodeResource map[int]corev1.ResourceList
	cpuset           cpuset.CPUSet
}

func (s *preFilterState) Clone() fwktype.StateData {
	_ = "STUB: not implemented"
	return *new(fwktype.StateData)
}

// CleanSchedulingData clears the scheduling cycle data in the stateData to reduce memory cost before entering
// the binding cycle.
func (s *preFilterState) CleanSchedulingData() { _ = "STUB: not implemented"; return }

func getPreFilterState(cycleState fwktype.CycleState) (*preFilterState, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) EventsToRegister(_ context.Context) ([]fwktype.ClusterEventWithHint, error) {
	_ = "STUB: not implemented"
	// To register a custom event, follow the naming convention at:
	// https://github.com/kubernetes/kubernetes/blob/e1ad9bee5bba8fbe85a6bf6201379ce8b1a611b1/pkg/scheduler/eventhandlers.go#L415-L422
	return nil, nil
}

func (p *Plugin) PreFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) (*fwktype.PreFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) PreFilterExtensions() fwktype.PreFilterExtensions {
	_ = "STUB: not implemented"
	return *new(fwktype.PreFilterExtensions)
}

func (p *Plugin) Filter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// It's necessary to force node to have NodeResourceTopology and CPUTopology
// We must satisfy the user's CPUSet request. Even if some nodes in the cluster have resources,
// they cannot be allocated without valid CPU topology.

// FIXME: move it ahead the resourceManager.Allocate so that we can check with NUMA hints almost in the Filter

// when numa topology policy is set on node, we should maintain the same behavior as before, so we only
// set default podNUMAExclusive when podNUMATopologyPolicy is not none

func (p *Plugin) filterAmplifiedCPUs(podRequestMilliCPU int64, nodeInfo fwktype.NodeInfo, requestCPUBind bool) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// TODO(joseph): Reservations and preemption should be considered here.
// TODO: support allocate reserved cpus with amplified ratios

func (p *Plugin) FilterReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, reservationInfo *frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) FilterNominateReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, reservationInfo *frameworkext.ReservationInfo, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// we have checked in filter, so we will not get error in reserve

// Here we want to filter if the pod can fit the reservation by this plugin, no matter the reservation is scheduled or not.
// But the tryAllocateFromReusable restricts that the pod should allocate from an scheduled reservation.
// So for pre-allocation, we keep the FilterNominateReservation but swap the allocate func inside for pre-allocation.
// use the cycle state to avoid misunderstanding, e.g. a normal pod allocate from a scheduled pre-allocation reservation

// tryAllocateFromReusable returned nil/nil ("no reservation satisfied, but fallback allowed").
// In the FilterNominateReservation context, this means the reservation's NUMA scope is
// incompatible with the current best topology hint; treat it as unschedulable so this
// reservation is excluded from the nominator candidate pool.

func (p *Plugin) Reserve(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// ReservationRestoreState is O(n) complexity of node number of the cluster.
// clearData clears all nodes' data in the cycleState to reduce memory cost before entering the binding cycle.

// we have check in filter, so we will not get error in reserve

// when numa topology policy is set on node, we should maintain the same behavior as before, so we only
// set default podNUMAExclusive when podNUMATopologyPolicy is not none

func (p *Plugin) allocate(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, node *corev1.Node, numaTopologyPolicy extension.NUMATopologyPolicy) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// ReservationRestoreState is O(n) complexity of node number of the cluster.
// clearData clears all nodes' data in the cycleState to reduce memory cost before entering the binding cycle.

// TODO: de-duplicate logic done by the Filter phase and move head the pre-process of the resource options

// Finally, try to allocate from node remaining resources

func getDesignatedNUMAHints(alloc *allocation) bitmask.BitMask {
	_ = "STUB: not implemented"
	return *new(bitmask.BitMask)
}

func (p *Plugin) Unreserve(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

func (p *Plugin) PreBindPreFlight(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) PreBind(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) PreBindReservation(ctx context.Context, cycleState fwktype.CycleState, reservation *schedulingv1alpha1.Reservation, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) preBindObject(ctx context.Context, cycleState fwktype.CycleState, object metav1.Object, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) getResourceOptions(state *preFilterState, node *corev1.Node, requestCPUBind bool, affinity topologymanager.NUMATopologyHint, topologyOptions TopologyOptions) (*ResourceOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tryAllocateFromNode(
	manager ResourceManager,
	preFilterState *preFilterState,
	restoreState *nodeReservationRestoreStateData,
	resourceOptions *ResourceOptions,
	pod *corev1.Pod,
	node *corev1.Node,
) (allocation *PodAllocation, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return both unmatched and matched reservation owner usage to eliminate double accounting.
// For each reservation, the fake reservation pod occupies R.Spec.Resources while its owner
// pods are independently accounted in the node cache. Only the second-layer owner usage
// (already consumed inside the reservation) should be returned here; R.Spec.Resources remain
// deducted as they are still reserved from the pod's perspective when going through the
// node-free path, so the pod cannot steal any reservation capacity.

// update with node preemption state

func appendResourceSpecIfMissed(object metav1.Object, state *preFilterState, node *corev1.Node, topologyOpts *TopologyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Write back ResourceSpec annotation if the Pod hasn't specified CPUBindPolicy
