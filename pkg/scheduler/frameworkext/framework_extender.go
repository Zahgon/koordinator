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

package frameworkext

import (
	"context"
	"fmt"

	nrtinformers "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/generated/informers/externalversions"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	listerscorev1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog/v2"
	fwktype "k8s.io/kube-scheduler/framework"
	schedconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/metrics"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	koordinatorclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	koordinatorinformers "github.com/koordinator-sh/koordinator/pkg/client/informers/externalversions"
	listerschedulingv1alpha1 "github.com/koordinator-sh/koordinator/pkg/client/listers/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/networktopology"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/topologymanager"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/workloadauditor"
)

var (
	_ FrameworkExtender                               = &frameworkExtenderImpl{}
	_ topologymanager.NUMATopologyHintProviderFactory = &frameworkExtenderImpl{}

	ErrSchedulerNameUnmatched = fmt.Errorf("schedulerName unmatched")
	ErrPodIsBeingDeleted      = fmt.Errorf("pod is being deleted")
)

type frameworkExtenderImpl struct {
	framework.Framework
	*errorHandlerDispatcher
	forgetPodHandlers []ForgetPodHandler

	schedulerFn       func() Scheduler
	configuredPlugins *schedconfig.Plugins
	monitor           *SchedulerMonitor

	koordinatorClientSet                koordinatorclientset.Interface
	koordinatorSharedInformerFactory    koordinatorinformers.SharedInformerFactory
	nodeResourceTopologyInformerFactory nrtinformers.SharedInformerFactory
	podLister                           listerscorev1.PodLister
	reservationLister                   listerschedulingv1alpha1.ReservationLister

	networkTopologyTreeManager networktopology.TreeManager
	workloadAuditor            workloadauditor.WorkloadAuditor

	preFilterTransformers         map[string]PreFilterTransformer
	filterTransformers            map[string]FilterTransformer
	scoreTransformers             map[string]ScoreTransformer
	postFilterTransformers        map[string]PostFilterTransformer
	preFilterTransformersEnabled  []PreFilterTransformer
	filterTransformersEnabled     []FilterTransformer
	scoreTransformersEnabled      []ScoreTransformer
	postFilterTransformersEnabled []PostFilterTransformer

	findOneNodePlugin FindOneNodePlugin
	preferNodesPlugin PreferNodesPlugin

	reservationCache                       ReservationCache
	reservationNominator                   ReservationNominator
	reservationFilterPlugins               []ReservationFilterPlugin
	reservationScorePlugins                []ReservationScorePlugin
	reservationPreBindPlugins              map[string]ReservationPreBindPlugin
	reservationRestorePlugins              []ReservationRestorePlugin
	reservationPreAllocationRestorePlugins []ReservationPreAllocationRestorePlugin

	resizePodPlugins         []ResizePodPlugin
	preBindExtensionsPlugins map[string]PreBindExtensions

	numaTopologyHintProviders []topologymanager.NUMATopologyHintProvider
	topologyManager           topologymanager.Interface

	metricsRecorder *metrics.MetricAsyncRecorder

	crossSchedulerNominator *CrossSchedulerPodNominator
}

func NewFrameworkExtender(f *FrameworkExtenderFactory, fw framework.Framework) FrameworkExtender {
	_ = "STUB: not implemented"
	return *new(FrameworkExtender)
}

// Register the profile name to CrossSchedulerPodNominator so that pods from
// this profile are excluded from cross-scheduler nomination tracking.

func (ext *frameworkExtenderImpl) updateTransformer(transformers ...SchedulingTransformer) {
	_ = "STUB: not implemented"
	return
}

func (ext *frameworkExtenderImpl) updatePlugins(pl fwktype.Plugin) {
	_ = "STUB: not implemented"
	return
}

// TODO(joseph): In the future, use only the default ReservationNominator

func (ext *frameworkExtenderImpl) SetConfiguredPlugins(plugins *schedconfig.Plugins) {
	_ = "STUB: not implemented"
	return
}

func (ext *frameworkExtenderImpl) KoordinatorClientSet() koordinatorclientset.Interface {
	_ = "STUB: not implemented"
	return *new(koordinatorclientset.Interface)
}

func (ext *frameworkExtenderImpl) KoordinatorSharedInformerFactory() koordinatorinformers.SharedInformerFactory {
	_ = "STUB: not implemented"
	return *new(koordinatorinformers.SharedInformerFactory)
}

func (ext *frameworkExtenderImpl) NodeResourceTopologyInformerFactory() nrtinformers.SharedInformerFactory {
	_ = "STUB: not implemented"
	return *new(nrtinformers.SharedInformerFactory)
}

// Scheduler return the scheduler adapter to support operating with cache and schedulingQueue.
// NOTE: Plugins do not acquire a dispatcher instance during plugin initialization,
// nor are they allowed to hold the object within the plugin object.
func (ext *frameworkExtenderImpl) Scheduler() Scheduler {
	_ = "STUB: not implemented"
	return *new(Scheduler)
}

func (ext *frameworkExtenderImpl) GetReservationCache() ReservationCache {
	_ = "STUB: not implemented"
	return *new(ReservationCache)
}

func (ext *frameworkExtenderImpl) GetReservationNominator() ReservationNominator {
	_ = "STUB: not implemented"
	return *new(ReservationNominator)
}

func (ext *frameworkExtenderImpl) GetNetworkTopologyTreeManager() networktopology.TreeManager {
	_ = "STUB: not implemented"
	return *new(networktopology.TreeManager)
}

func (ext *frameworkExtenderImpl) GetCrossSchedulerPodNominator() *CrossSchedulerPodNominator {
	_ = "STUB: not implemented"
	return nil
}

func (ext *frameworkExtenderImpl) GetWorkloadAuditor() workloadauditor.WorkloadAuditor {
	_ = "STUB: not implemented"
	return *new(workloadauditor.WorkloadAuditor)
}

// RunPreFilterPlugins transforms the PreFilter phase of framework with pre-filter transformers.
func (ext *frameworkExtenderImpl) RunPreFilterPlugins(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod) (*fwktype.PreFilterResult, *fwktype.Status, sets.Set[string]) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// FindOneNode

// skip

// PreferNodes

// skip

func (ext *frameworkExtenderImpl) RunFindOneNodePlugin(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, result *fwktype.PreFilterResult) (string, *fwktype.Status) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ext *frameworkExtenderImpl) RunPreferNodesPlugin(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, result *fwktype.PreFilterResult) (string, *fwktype.Status) {
	_ = "STUB: not implemented"
	return "", nil
}

// If the PreferNodes plugin returns Success, try to filter the preferred nodes.
// Then if any preferred node passes the predicate, return the node as the PreFilterResult.
// Otherwise, fallback to the original PreFilterResult.

// pick first suitable node

// RunFilterPluginsWithNominatedPods transforms the Filter phase of framework with filter transformers.
// We don't transform RunFilterPlugins since framework's RunFilterPluginsWithNominatedPods just calls its RunFilterPlugins.
func (ext *frameworkExtenderImpl) RunFilterPluginsWithNominatedPods(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (ext *frameworkExtenderImpl) RunScorePlugins(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodeInfos []fwktype.NodeInfo) ([]fwktype.NodePluginScores, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ext *frameworkExtenderImpl) RunPostFilterPlugins(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, filteredNodeStatusMap fwktype.NodeToStatusReader) (_ *fwktype.PostFilterResult, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunPreBindPlugins supports PreBindReservation for Reservation
func (ext *frameworkExtenderImpl) RunPreBindPlugins(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// check if schedulerName matched

func (ext *frameworkExtenderImpl) runPreBindExtensionPlugins(ctx context.Context, cycleState fwktype.CycleState, originalObj, modifiedObj metav1.Object) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (ext *frameworkExtenderImpl) RunPostBindPlugins(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

func (ext *frameworkExtenderImpl) RunReservationExtensionPreRestoreReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// RunReservationExtensionRestoreReservation restores the Reservation during PreFilter phase
func (ext *frameworkExtenderImpl) RunReservationExtensionRestoreReservation(ctx context.Context, cycleState fwktype.CycleState, podToSchedule *corev1.Pod, matched []*ReservationInfo, unmatched []*ReservationInfo, nodeInfo fwktype.NodeInfo) (PluginToReservationRestoreStates, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(PluginToReservationRestoreStates), nil
}

func (ext *frameworkExtenderImpl) RunReservationExtensionFinalRestoreReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, states PluginToNodeReservationRestoreStates) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (ext *frameworkExtenderImpl) RunReservationExtensionPreRestoreReservationPreAllocation(ctx context.Context, cycleState fwktype.CycleState, rInfo *ReservationInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (ext *frameworkExtenderImpl) RunReservationExtensionRestoreReservationPreAllocation(ctx context.Context, cycleState fwktype.CycleState, rInfo *ReservationInfo, preAllocatable []*corev1.Pod, nodeInfo fwktype.NodeInfo) (PluginToReservationRestoreStates, *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(PluginToReservationRestoreStates), nil
}

// RunReservationFilterPlugins determines whether the Reservation can participate in the Reserve
func (ext *frameworkExtenderImpl) RunReservationFilterPlugins(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, reservationInfo *ReservationInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// RunNominateReservationFilterPlugins determines whether the Reservation can participate in the Reserve.
func (ext *frameworkExtenderImpl) RunNominateReservationFilterPlugins(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, reservationInfo *ReservationInfo, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// RunReservationScorePlugins ranks the Reservations
func (ext *frameworkExtenderImpl) RunReservationScorePlugins(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, reservationInfos []*ReservationInfo, nodeName string) (ps PluginToReservationScores, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(PluginToReservationScores), nil
}

// TODO: Should support configure weight

// return error if score plugin returns invalid score.

func (ext *frameworkExtenderImpl) RunReservationPreAllocationScorePlugins(ctx context.Context, cycleState fwktype.CycleState, rInfo *ReservationInfo, pods []*corev1.Pod, nodeName string) (ps PluginToReservationScores, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return *new(PluginToReservationScores), nil
}

// each ReservationScore corresponds to a pod

// TODO: Should support configure weight (same as RunReservationScorePlugins)

// return error if score plugin returns invalid score.

func (ext *frameworkExtenderImpl) RegisterForgetPodHandler(handler ForgetPodHandler) {
	_ = "STUB: not implemented"
	return
}

func (ext *frameworkExtenderImpl) ForgetPod(logger klog.Logger, pod *corev1.Pod) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Always call plugin handlers even if ForgetPod fails.
// It is tolerated for multi-scheduler scenarios where the pod is not in assumed
// cache (because it was assumed bound but then failed), but we still need to clean
// up plugin-specific accounting.

// Return the original error after calling handlers

func (ext *frameworkExtenderImpl) RunNUMATopologyManagerAdmit(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, node *corev1.Node, numaNodes []int, policyType apiext.NUMATopologyPolicy, exclusivePolicy apiext.NumaTopologyExclusive, allNUMANodeStatus []apiext.NumaNodeStatus) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (ext *frameworkExtenderImpl) GetNUMATopologyHintProvider() []topologymanager.NUMATopologyHintProvider {
	_ = "STUB: not implemented"
	return nil
}

func (ext *frameworkExtenderImpl) RunReservePluginsReserve(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// FIXME: keep consistent behavior with the framework assuming

func (ext *frameworkExtenderImpl) RunResizePod(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}
