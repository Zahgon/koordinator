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

package cgreconcile

import (
	"time"

	corev1 "k8s.io/api/core/v1"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

const (
	CgroupReconcileName = "CgroupReconcile"
)

var _ framework.QOSStrategy = &cgroupResourcesReconcile{}

type cgroupResourcesReconcile struct {
	reconcileInterval time.Duration
	statesInformer    statesinformer.StatesInformer
	executor          resourceexecutor.ResourceUpdateExecutor
}

// cgroupResourceSummary summarizes values of cgroup resources to update; nil value means not to update
type cgroupResourceSummary struct {
	// Memory
	memoryMin              *int64
	memoryLow              *int64
	memoryHigh             *int64
	memoryWmarkRatio       *int64
	memoryWmarkScaleFactor *int64
	memoryWmarkMinAdj      *int64
	memoryUsePriorityOom   *int64
	memoryPriority         *int64
	memoryOomKillGroup     *int64
}

type cgroupResourceUpdaterMeta struct {
	resourceType system.ResourceType
	value        *int64
	isMergeable  bool
}

func New(opt *framework.Options) framework.QOSStrategy {
	_ = "STUB: not implemented"
	return *new(framework.QOSStrategy)
}

func (m *cgroupResourcesReconcile) Enabled() bool { _ = "STUB: not implemented"; return false }

func (m *cgroupResourcesReconcile) Setup(context *framework.Context) {
	_ = "STUB: not implemented"
	return
}

func (m *cgroupResourcesReconcile) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (m *cgroupResourcesReconcile) init(stopCh <-chan struct{}) {
	m.executor.Run(stopCh)
}

func (m *cgroupResourcesReconcile) reconcile() { _ = "STUB: not implemented"; return }

// do nothing if nodeSLO == nil || nodeSLO.Spec.ResourceQOSStrategy == nil

// apply CgroupReconcile: calculate resources to update, and then update them by a leveled order to avoid dynamic
// resource overcommitment/leak

func (m *cgroupResourcesReconcile) calculateAndUpdateResources(nodeSLO *slov1alpha1.NodeSLO) {
	_ = "STUB: not implemented"
	// 1. sort cgroup resources by the owner level (qos, pod, container).
	//    e.g. for hierarchical resources of memoryMin, when qos-level memoryMin increases, they should be updated from
	//         the top to bottom; while resources should be updated from the bottom to top when qos-level memoryMin
	//         decreases to avoid higher-level's over-commit.
	// 2. update resources in level order
	return
}

// calculate qos-level, pod-level and container-level resources

// to make sure the hierarchical cgroup resources are correctly updated, we simply update the resources by
// cgroup-level order.
// e.g. /kubepods.slice/memory.min, /kubepods.slice-podxxx/memory.min, /kubepods.slice-podxxx/docker-yyy/memory.min

// calculateResources calculates qos-level, pod-level and container-level resources with nodeCfg and podMetas
func (m *cgroupResourcesReconcile) calculateResources(nodeCfg *slov1alpha1.ResourceQOSStrategy, node *corev1.Node,
	podMetas []*statesinformer.PodMeta) (qosLevelResources, podLevelResources, containerLevelResources []resourceexecutor.ResourceUpdater) {
	_ = "STUB: not implemented"
	// TODO: check anolis os version
	return nil, nil, nil
}

// ignore non-running pods

// retrieve pod-level config
// assert kubeQoS belongs to {Guaranteed, Burstable, Besteffort}

// update summary for qos resources

// calculate pod-level and container-level resources and make resourceUpdaters

// summarize qos-level resources

// calculate qos-level resources with the qos summary
// NOTE: first visit Guaranteed since it actually has a higher level cgroup than others'

// make qos resourceUpdaters

func (m *cgroupResourcesReconcile) calculateQoSResources(summary *cgroupResourceSummary, qos corev1.PodQOSClass,
	qosCfg *slov1alpha1.ResourceQOS) []resourceexecutor.ResourceUpdater {
	_ = "STUB: not implemented"
	// double-check qosCfg is not nil
	return nil
}

// Mem QoS

func (m *cgroupResourcesReconcile) calculatePodAndContainerResources(podMeta *statesinformer.PodMeta, node *corev1.Node,
	podCfg *slov1alpha1.ResourceQOS) (podResources, containerResources []resourceexecutor.ResourceUpdater) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *cgroupResourcesReconcile) calculatePodResources(pod *corev1.Pod, parentDir string, podCfg *slov1alpha1.ResourceQOS) []resourceexecutor.ResourceUpdater {
	_ = "STUB: not implemented"
	// double-check qos config is not nil
	return nil
}

// Mem QoS
// resources statically use configured values

// resources calculated with pod spec

// memory.min, memory.low: just sum all containers' memory requests; regard as no memory protection when any
// of containers does not set request

// assert no overflow for request < 1PiB

// values improved: memory.low is no less than memory.min

func (m *cgroupResourcesReconcile) calculateContainerResources(container *corev1.Container, pod *corev1.Pod,
	node *corev1.Node, parentDir string, podCfg *slov1alpha1.ResourceQOS) []resourceexecutor.ResourceUpdater {
	_ = "STUB: not implemented"
	// double-check qos config is not nil
	return nil
}

// Mem QoS
// resources statically use configured values

// resources calculated with container spec

// when container request not set, memory request is counted as zero but not unlimited(-1)

// memory.min, memory.low: if container's memory request is not set, just consider it as zero

// memory.high: if container's memory throttling factor is set as zero, disable memory.high by set to maximal;
// else if factor is set while container's limit not set, set memory.high with node memory allocatable

// reset to system default if set 0
// writing MaxInt64 is equal to write "max"

// values improved: memory.low is no less than memory.min

// values improved: memory.high is no less than memory.min

// getMergedPodResourceQoS returns a merged ResourceQOS for the pod (i.e. a pod-level qos config).
// 1. merge pod-level cfg with node-level cfg if pod annotation of advanced qos config exists;
// 2. calculates and finally returns the pod-level cfg with each feature cfg (e.g. pod-level memory qos config).
func (m *cgroupResourcesReconcile) getMergedPodResourceQoS(pod *corev1.Pod, cfg *slov1alpha1.ResourceQOS) (*slov1alpha1.ResourceQOS, error) {
	_ = "STUB: not implemented"
	// deep-copy node config into pod config; assert cfg == NoneResourceQOS when node disables
	return nil, nil
}

// update with memory qos config

// mergePodResourceQoSForMemoryQoS merges pod-level memory qos config with node-level resource qos config
// config overwrite: pod-level config > pod policy template > node-level config
func (m *cgroupResourcesReconcile) mergePodResourceQoSForMemoryQoS(pod *corev1.Pod, cfg *slov1alpha1.ResourceQOS) {
	_ = "STUB: not implemented"
	// get the pod-level config and determine if the pod is allowed
	return
}

// get pod-level config

// ignore pod-level memory qos config when parse error

// policy="" is equal to policy="default"

// if policy is not default, replace memory qos config with the policy template
// fully disable memory qos for policy=None

// qos=None would be set with kubeQoS for policy=Auto

// no need to merge config if pod-level config is nil

// otherwise detailed pod-level config is specified, merge with node-level config for the pod
// node config has been deep-copied

// not change memory qos config if merge error

// updateCgroupSummaryForQoS updates qos cgroup summary by pod to summarize qos-level cgroup according to belonging pods
func updateCgroupSummaryForQoS(summary *cgroupResourceSummary, pod *corev1.Pod, podCfg *slov1alpha1.ResourceQOS) {
	_ = "STUB: not implemented"
	// Memory QoS
	// `memory.min` for qos := sum(requests of pod with the qos * minLimitPercent); if factor is nil, set kernel default
	// `memory.low` for qos := sum(requests of pod with the qos * lowLimitPercent); if factor is nil, set kernel default
	return
}

// if any container's memory request is not set, just consider it as zero

// assert no overflow for req < 1PiB

// completeCgroupSummaryForQoS completes qos cgroup summary considering Guaranteed qos is higher than the others
func completeCgroupSummaryForQoS(qosSummary map[corev1.PodQOSClass]*cgroupResourceSummary) {
	_ = "STUB: not implemented"
	// memory qos
	// Guaranteed cgroup is the ancestor node of Burstable and Besteffort, so the `min` and `low` derive from the sum
	return
}

func makeCgroupResources(parentDir string, summary *cgroupResourceSummary) []resourceexecutor.ResourceUpdater {
	_ = "STUB: not implemented"
	return nil
}

//Memory
// mergeable resources: memory.min, memory.low, memory.high

// TBD: handle memory priority and oom group
