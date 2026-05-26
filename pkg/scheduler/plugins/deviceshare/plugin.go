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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	fwktype "k8s.io/kube-scheduler/framework"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

const (
	// Name is the name of the plugin used in the plugin registry and configurations.
	Name = "DeviceShare"

	// stateKey is the key in CycleState to pre-computed data.
	stateKey = Name
)

var (
	_ fwktype.EnqueueExtensions = &Plugin{}

	_ fwktype.PreFilterPlugin = &Plugin{}
	_ fwktype.FilterPlugin    = &Plugin{}
	_ fwktype.PreScorePlugin  = &Plugin{}
	_ fwktype.ScorePlugin     = &Plugin{}
	_ fwktype.ScoreExtensions = &Plugin{}
	_ fwktype.ReservePlugin   = &Plugin{}
	_ fwktype.PreBindPlugin   = &Plugin{}

	_ frameworkext.ResizePodPlugin                       = &Plugin{}
	_ frameworkext.ReservationRestorePlugin              = &Plugin{}
	_ frameworkext.ReservationPreAllocationRestorePlugin = &Plugin{}
	_ frameworkext.ReservationFilterPlugin               = &Plugin{}
	_ frameworkext.ReservationScorePlugin                = &Plugin{}
	_ frameworkext.ReservationScoreExtensions            = &Plugin{}
	_ frameworkext.ReservationPreBindPlugin              = &Plugin{}
)

type Plugin struct {
	disableDeviceNUMATopologyAlignment         bool
	handle                                     frameworkext.ExtendedHandle
	nodeDeviceCache                            *nodeDeviceCache
	gpuSharedResourceTemplatesCache            *gpuSharedResourceTemplatesCache
	gpuSharedResourceTemplatesMatchedResources []corev1.ResourceName
	gpuShareUnsupportedModels                  map[string]sets.Set[string]
	scorer                                     *resourceAllocationScorer
}

type preFilterState struct {
	skip                              bool
	podRequests                       map[schedulingv1alpha1.DeviceType]corev1.ResourceList
	hints                             apiext.DeviceAllocateHints
	hintSelectors                     map[schedulingv1alpha1.DeviceType][2]labels.Selector
	hasSelectors                      bool
	jointAllocate                     *apiext.DeviceJointAllocate
	primaryDeviceType                 schedulingv1alpha1.DeviceType
	podFitsSecondaryDeviceWellPlanned bool
	gpuRequirements                   *GPURequirements
	allocationResult                  apiext.DeviceAllocations
	preemptibleDevices                map[string]map[schedulingv1alpha1.DeviceType]deviceResources
	preemptibleInRRs                  map[string]map[types.UID]map[schedulingv1alpha1.DeviceType]deviceResources

	isReservationRequired bool

	// designatedAllocation is parsed from the Pod Annotation during the PreFilter phase. In this case, we should assume that the Node has already been selected. That is, all Score-related plug-ins will not be executed. Instead, we only need to call Filter to confirm whether the designatedAllocation is still valid and use it as the allocation result during actual allocation.
	designatedAllocation apiext.DeviceAllocations
	designatedVF         map[schedulingv1alpha1.DeviceType]map[int32]sets.Set[string]
}

type GPURequirements struct {
	numberOfGPUs                        int
	requestsPerGPU                      corev1.ResourceList
	gpuShared                           bool
	honorGPUPartition                   bool
	restrictedGPUPartition              bool
	rindBusBandwidth                    *resource.Quantity
	requiredTopologyScopeLevel          int
	requiredTopologyScope               apiext.DeviceTopologyScope
	enforceGPUSharedResourceTemplate    bool
	candidateGPUSharedResourceTemplates map[string]apiext.GPUSharedResourceTemplates
}

func (s *preFilterState) Clone() fwktype.StateData {
	_ = "STUB: not implemented"
	return *new(fwktype.StateData)
}

func (p *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

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

func (p *Plugin) AddPod(ctx context.Context, cycleState fwktype.CycleState, podToSchedule *corev1.Pod, podInfoToAdd fwktype.PodInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) RemovePod(ctx context.Context, cycleState fwktype.CycleState, podToSchedule *corev1.Pod, podInfoToRemove fwktype.PodInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Filter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// 当 Pod 在该节点上的 NUMA Affinity 已经算好，表明：Pod 在该节点上开启了 NUMA，且以算好的 NUMA Affinity 资源可分配，所以后续 Filter 逻辑可以忽略

// TODO 这里应该表示从节点剩余资源分，但是这里看起来不是这个意思

func (p *Plugin) FilterReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, reservationInfo *frameworkext.ReservationInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) FilterNominateReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, reservationInfo *frameworkext.ReservationInfo, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// if pre-allocation, we filter the reserve pod with the pre-allocatable pod

func (p *Plugin) Reserve(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"

	// ReservationRestoreState is O(n) complexity of node number of the cluster.
	// cleanReservationRestoreState clears ReservationRestoreState in the stateData to reduce memory cost before entering
	// the binding cycle.
	return nil
}

func logAllocationContext(pod *corev1.Pod, nodeName string, nodeDeviceInfo *nodeDevice, designatedAllocation, allocationResult apiext.DeviceAllocations) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *Plugin) allocate(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, node *corev1.Node) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// TODO: de-duplicate logic done by the Filter phase and move head the pre-process of the resource options

func (p *Plugin) Unreserve(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

func (p *Plugin) ResizePod(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
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

// indicates we already skip our device allocation logic, leave it to kubelet

func (p *Plugin) fillID(allocationResult apiext.DeviceAllocations, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func updateReservationAllocatable(cycleState fwktype.CycleState, reservation *schedulingv1alpha1.Reservation) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) getNodeDeviceSummary(nodeName string) (*NodeDeviceSummary, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *Plugin) getAllNodeDeviceSummary() map[string]*NodeDeviceSummary {
	_ = "STUB: not implemented"
	return nil
}

func New(ctx context.Context, obj runtime.Object, handle fwktype.Handle) (fwktype.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.Plugin), nil
}

// Register the node informer synchronously during New so that the registration
// is visible to the framework's WaitForHandlersSync. If we registered it inside
// the gcNodeDevice goroutine, the framework might start scheduling before the
// handler registration is collected. Using a no-op handler because gcNodeDevice
// only reads from the node lister.
