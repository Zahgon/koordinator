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

package memoryevict

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
	MemoryEvictName = "memoryEvict"

	memoryReleaseBufferPercent = 2
)

var _ framework.QOSStrategy = &memoryEvictor{}

type memoryEvictor struct {
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

func (m *memoryEvictor) Enabled() bool { _ = "STUB: not implemented"; return false }

func (m *memoryEvictor) Setup(ctx *framework.Context) { _ = "STUB: not implemented"; return }

func (m *memoryEvictor) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (m *memoryEvictor) memoryEvict() { _ = "STUB: not implemented"; return }

// build memory evict tasks

// triggerFeatures defines the features that can trigger pod eviction based on memory pressure.
// - features.BEMemoryEvict and features.MemoryEvict can operate independently.
// - All evaluate eviction decisions using the same memory watermark (threshold).
// - When both are enabled:
//   * BEMemoryEvict runs first, followed by MemoryAllocatableEvict, then MemoryEvict.
//   * They target different sets of pods (no duplicate victims), but the overall resource
//     release effect is cumulative and considered together during scheduling and capacity planning.
// - Although safe to enable concurrently, it is generally NOT RECOMMENDED to run both simultaneously,
//   as this may lead to redundant eviction logic and increased system complexity without clear benefit.

// report and renew time

func (m *memoryEvictor) buildEvictTask(feature featuregate.Feature, nodeSLO *slov1alpha1.NodeSLO, node *corev1.Node) (*qosmanagerUtil.EvictTaskInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateConfigCheck(feature featuregate.Feature) func(*slov1alpha1.ResourceThresholdStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

func isAllocatableThresholdConfigValid(thresholdConfig *slov1alpha1.ResourceThresholdStrategy) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *memoryEvictor) calculateReleaseByUsedThresholdPercent(thresholdConfig *slov1alpha1.ResourceThresholdStrategy, node *corev1.Node, pods []*statesinformer.PodMeta) (overall corev1.ResourceList,
	calculateFunc func(podInfo *qosmanagerUtil.PodEvictInfo) corev1.ResourceList) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func (m *memoryEvictor) calculateReleaseByAllocatableThresholdPercent(thresholdConfig *slov1alpha1.ResourceThresholdStrategy, node *corev1.Node, pods []*statesinformer.PodMeta) (overall corev1.ResourceList,
	calculateFunc func(podInfo *qosmanagerUtil.PodEvictInfo) corev1.ResourceList) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

// no need to release

// currently only support koord-batch/koord-mid

func (m *memoryEvictor) getSortedBEPodInfos(evictionPolicy string, thresholdConfig *slov1alpha1.ResourceThresholdStrategy, pods []*statesinformer.PodMeta) []*qosmanagerUtil.PodEvictInfo {
	_ = "STUB: not implemented"
	return nil
}

// TODO: https://github.com/koordinator-sh/koordinator/pull/65#discussion_r849048467
// compare priority > podMetric > name

//return bePodInfos[i].podMetric.MemoryUsed.MemoryWithoutCache.Value() > bePodInfos[j].podMetric.MemoryUsed.MemoryWithoutCache.Value()

func (m *memoryEvictor) getPodEvictInfoAndSortByAllocatable(evictionPolicy string, thresholdConfig *slov1alpha1.ResourceThresholdStrategy, pods []*statesinformer.PodMeta) []*qosmanagerUtil.PodEvictInfo {
	_ = "STUB: not implemented"
	return nil
}

func (m *memoryEvictor) getPodEvictInfoAndSortByUsed(evictionPolicy string, thresholdConfig *slov1alpha1.ResourceThresholdStrategy, pods []*statesinformer.PodMeta) []*qosmanagerUtil.PodEvictInfo {
	_ = "STUB: not implemented"
	return nil
}

func (m *memoryEvictor) getPodEvictInfoAndSortByPriority(evictionPolicy string, priorityThreshold int32, pods []*statesinformer.PodMeta, subSortFun func(a, b *qosmanagerUtil.PodEvictInfo) bool) []*qosmanagerUtil.PodEvictInfo {
	_ = "STUB: not implemented"
	return nil
}

// higher priority pods are not allowed to evict
// 1. filter inactive / priority / eviction policy

// 2. filter: exclude: koordinator.sh/eviction-enabled

// 3. sort :  koordinator.sh/priority

// 4. filter no metrics
