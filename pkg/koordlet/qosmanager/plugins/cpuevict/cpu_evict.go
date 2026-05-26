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

package cpuevict

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/component-base/featuregate"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/framework"
	qosmanagerUtil "github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/plugins/util"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	CPUEvictName = "CPUEvict"

	beCPUSatisfactionLowPercentMax   = 60
	beCPUSatisfactionUpperPercentMax = 100
	beCPUUsageThresholdPercent       = 90

	defaultMinAllocatableBatchMilliCPU = 1
	cpuReleaseBufferPercent            = 2
)

var _ framework.QOSStrategy = &cpuEvictor{}

type cpuEvictor struct {
	evictInterval         time.Duration
	evictCoolingInterval  time.Duration
	metricCollectInterval time.Duration
	statesInformer        statesinformer.StatesInformer
	metricCache           metriccache.MetricCache
	lastEvictTime         time.Time
	evictExecutor         qosmanagerUtil.EvictionExecutor
}

func New(opt *framework.Options) framework.QOSStrategy {
	_ = "STUB: not implemented"
	return *new(framework.QOSStrategy)
}

func (c *cpuEvictor) Enabled() bool { _ = "STUB: not implemented"; return false }

func (c *cpuEvictor) Setup(ctx *framework.Context) { _ = "STUB: not implemented"; return }

func (c *cpuEvictor) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

type podEvictCPUInfo struct {
	milliRequest   int64
	milliUsedCores int64
	cpuUsage       float64 // cpuUsage = milliUsedCores / milliRequest
	pod            *corev1.Pod
}

// cpu evict triggered by configured mechanism:
func (c *cpuEvictor) cpuEvict() { _ = "STUB: not implemented"; return }

// runtime check

// build cpu evict tasks

// When both features.BECPUEvict and features.CPUEvict are enabled:
// - All eviction mechanisms will be activated.
// - BECPUEvict runs first, followed by CPUAllocatableEvict, then CPUEvict.
// - Resource release effects are cumulatively considered; the total reclaimed resources
//   from both phases are accounted for in scheduling and capacity planning.

// report and renew time

// calculate overall resource task need to release and functions to calculate resource from pod
func (c *cpuEvictor) calculateMilliReleaseByBESatisfaction(thresholdConfig *slov1alpha1.ResourceThresholdStrategy, node *corev1.Node, pods []*statesinformer.PodMeta) (overall corev1.ResourceList,
	calculateFunc func(podInfo *qosmanagerUtil.PodEvictInfo) corev1.ResourceList) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

// Sidecar containers run alongside regular containers and should be summed.

// Step1: Calculate release resource by BECPUResourceMetric in window

// BECPUUsage

// BECPURequest

// BECPULimit

// CPU Satisfaction considers the allocatable when policy=evictByAllocatable.

// get min count

// Step2: Calculate release resource current

// BECPUUsage

// BECPURequest

// BECPULimit

// CPU Satisfaction considers the allocatable when policy=evictByAllocatable.

// Requests and limits do not change frequently.
// If the current request and limit are equal to the average request and limit within the window period, there is no need to recalculate.

// Step3：release = min(releaseByAvg,releaseByCurrent)

func (c *cpuEvictor) calculateMilliReleaseByUsedThresholdPercent(thresholdConfig *slov1alpha1.ResourceThresholdStrategy, node *corev1.Node, pods []*statesinformer.PodMeta) (overall corev1.ResourceList,
	calculateFunc func(podInfo *qosmanagerUtil.PodEvictInfo) corev1.ResourceList) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

// cpu peak shaving

func (c *cpuEvictor) calculateMilliReleaseByAllocatableThresholdPercent(thresholdConfig *slov1alpha1.ResourceThresholdStrategy, node *corev1.Node, pods []*statesinformer.PodMeta) (overall corev1.ResourceList,
	calculateFunc func(podInfo *qosmanagerUtil.PodEvictInfo) corev1.ResourceList) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

// no need to release

// mid/batch are milli format on node: to confirm

// currently only support koord-batch/koord-mid

func isAvgQueryResultValid(windowSeconds, collectIntervalSeconds, count int64) bool {
	_ = "STUB: not implemented"
	return false
}

func isBECPUUsageHighEnough(beCPUMilliUsage, beCPUMilliRealLimit float64, thresholdPercent *int64) bool {
	_ = "STUB: not implemented"
	return false
}

func calculateResourceMilliToReleaseBySatisfaction(beCPUMilliRequest, beCPUMilliRealLimit float64, thresholdConfig *slov1alpha1.ResourceThresholdStrategy) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *cpuEvictor) buildEvictTask(feature featuregate.Feature, nodeSLO *slov1alpha1.NodeSLO, node *corev1.Node) (*qosmanagerUtil.EvictTaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// calculate overall resource task need to release and functions to calculate resource from pod

func (c *cpuEvictor) getPodEvictInfoAndSortByAllocatable(evictionPolicy string, thresholdConfig *slov1alpha1.ResourceThresholdStrategy, pods []*statesinformer.PodMeta) []*qosmanagerUtil.PodEvictInfo {
	_ = "STUB: not implemented"
	return nil
}

func (c *cpuEvictor) getPodEvictInfoAndSortByUsed(evictionPolicy string, thresholdConfig *slov1alpha1.ResourceThresholdStrategy, pods []*statesinformer.PodMeta) []*qosmanagerUtil.PodEvictInfo {
	_ = "STUB: not implemented"
	return nil
}

func (c *cpuEvictor) getPodEvictInfoAndSortByPriority(evictionPolicy string, priorityThreshold int32, pods []*statesinformer.PodMeta, subSortFun func(a, b *qosmanagerUtil.PodEvictInfo) bool) []*qosmanagerUtil.PodEvictInfo {
	_ = "STUB: not implemented"
	return nil
}

// higher priority pods are not allowed to evict
// 1. filter inactive / priority / eviction policy

// 2. filter: exclude: koordinator.sh/eviction-enabled

// 3. sort :  koordinator.sh/priority

// 4. filter no metrics

func (c *cpuEvictor) getBEPodEvictInfoAndSort(evictionPolicy string, thresholdConfig *slov1alpha1.ResourceThresholdStrategy, pods []*statesinformer.PodMeta) []*qosmanagerUtil.PodEvictInfo {
	_ = "STUB: not implemented"
	return nil
}

// Sidecar containers run alongside regular containers and should be summed.

func (c *cpuEvictor) getBEMilliAllocatable() float64 { _ = "STUB: not implemented"; return 0 }

// The batch allocatable value can be set to zero when high-priority util is high, where we still need to calculate
// the satisfaction rate. Here we use a small allocatable for the BE utilization check.

func isSatisfactionConfigValid(thresholdConfig *slov1alpha1.ResourceThresholdStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

func isUsedThresholdConfigValid(thresholdConfig *slov1alpha1.ResourceThresholdStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

func isAllocatableThresholdConfigValid(thresholdConfig *slov1alpha1.ResourceThresholdStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

func getBECPUMetric(resouceAllocation metriccache.MetricPropertyValue, querier metriccache.Querier, aggregateType metriccache.AggregationType) (float64, int64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func minInt64(num ...int64) int64 { _ = "STUB: not implemented"; return 0 }
