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

package cpuburst

import (
	"math/rand"
	"time"

	corev1 "k8s.io/api/core/v1"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	CPUBurstName = "CPUBurst"

	cfsIncreaseStep = 1.2
	cfsDecreaseStep = 0.8

	sharePoolCoolingThresholdRatio = 0.9

	cpuThresholdPercentForLimiterConsumeTokens = 100
	cpuThresholdPercentForLimiterSavingTokens  = 60
)

// cfsOperation is used for CFSQuotaBurst strategy
type cfsOperation int64

const (
	cfsScaleUp cfsOperation = iota
	cfsScaleDown
	cfsRemain
	cfsReset
)

func (o cfsOperation) String() string { _ = "STUB: not implemented"; return "" }

// nodeStateForBurst depends on cpu-share-pool usage, used for CFSBurstStrategy
type nodeStateForBurst int64

const (
	// cpu-share-pool usage >= threshold
	nodeBurstOverload nodeStateForBurst = iota
	// threshold * 0.9 <= cpu-share-pool usage < threshold
	nodeBurstCooling nodeStateForBurst = iota
	// cpu-share-pool usage < threshold * 0.9
	nodeBurstIdle nodeStateForBurst = iota
	// cpu-share-pool is unknown
	nodeBurstUnknown nodeStateForBurst = iota
)

func (s nodeStateForBurst) String() string { _ = "STUB: not implemented"; return "" }

// burstLimiter is a token bucket limiter for CFSQuotaBurst strategy, limit container continuously overused
// bucket capacity = burstCfg.CFSQuotaBurstPeriodSeconds * burstCfg.CFSQuotaBurstPercent
// bucket accumulate/consume = (currentUsageScalePercent - 100) * int64(timePastSec)
type burstLimiter struct {
	bucketCapacity int64
	currentToken   int64
	lastUpdateTime time.Time
	expireDuration time.Duration
}

func newBurstLimiter(burstPeriodSec, maxScalePercent int64) *burstLimiter {
	_ = "STUB: not implemented"
	return nil
}

func (l *burstLimiter) init(burstPeriodSec, maxScalePercent int64) {
	capacity := burstPeriodSec * (maxScalePercent - 100)
	// init currentToken with capacity * randomInitRatio, which in range [0-0.5)
	randomInitRatio := rand.Float64() / 2
	initSize := float64(capacity) * randomInitRatio
	l.bucketCapacity = capacity
	l.currentToken = int64(initSize)
	l.lastUpdateTime = time.Now()
	l.expireDuration = time.Duration(2*burstPeriodSec) * time.Second
}

func (l *burstLimiter) Allow(now time.Time, currentUsageScalePercent int64) (bool, int64) {
	_ = "STUB: not implemented"
	return false, 0
}

func (l *burstLimiter) UpdateIfChanged(burstPeriodSec, maxScalePercent int64) {
	_ = "STUB: not implemented"
	// update if config changed
	return
}

func (l *burstLimiter) Expire() bool { _ = "STUB: not implemented"; return false }

var _ framework.QOSStrategy = &cpuBurst{}

type cpuBurst struct {
	reconcileInterval     time.Duration
	metricCollectInterval time.Duration
	statesInformer        statesinformer.StatesInformer
	metricCache           metriccache.MetricCache
	executor              resourceexecutor.ResourceUpdateExecutor
	cgroupReader          resourceexecutor.CgroupReader
	nodeCPUBurstStrategy  *slov1alpha1.CPUBurstStrategy
	containerLimiter      map[string]*burstLimiter
}

func New(opt *framework.Options) framework.QOSStrategy {
	_ = "STUB: not implemented"
	return *new(framework.QOSStrategy)
}

func (b *cpuBurst) Enabled() bool { _ = "STUB: not implemented"; return false }

func (b *cpuBurst) Setup(ctx *framework.Context) { _ = "STUB: not implemented"; return }

func (b *cpuBurst) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (b *cpuBurst) init(stopCh <-chan struct{}) {
	b.executor.Run(stopCh)
}

func (b *cpuBurst) start() { _ = "STUB: not implemented"; return }

// at the beginning of appling cpu burst strategy, we should reset all metrics belongs to pods and containers

// sync config from node slo

// get node state by node share pool usage

// ignore non-burstable pod, e.g. LSR, BE pods

// ignore pods that status.phase is not pending or running

// merge burst config from pod and node

// set cpu.cfs_burst_us for pod and containers

// scale cpu.cfs_quota_us for pod and containers

// getNodeStateForBurst checks whether node share pool cpu usage beyonds the threshold
// return isOverload, share pool usage ratio and message detail
func (b *cpuBurst) getNodeStateForBurst(sharePoolThresholdPercent int64,
	podsMeta []*statesinformer.PodMeta) nodeStateForBurst {
	_ = "STUB: not implemented"
	return *new(nodeStateForBurst)
}

// calculate cpu share pool info; for conservative reason, include system usage in share pool

// exclude LSE/LSR pod cpu from cpu share pool

// exclude LSE/LSR/BE pod cpu usage from cpu share pool

// end for podsMeta

// calculate cpu share pool usage ratio

// generate node burst state by cpu share pool usage

// sharePoolUsageRatio < sharePoolCoolingRatio

// scale cpu.cfs_quota_us for pod/containers by container throttled state and node state
func (b *cpuBurst) applyCFSQuotaBurst(burstCfg *slov1alpha1.CPUBurstConfig, podMeta *statesinformer.PodMeta,
	nodeState nodeStateForBurst) {
	_ = "STUB: not implemented"
	return
}

// end for containers

// check if cfs burst for container is allowed by limiter config, return true if allowed
func (b *cpuBurst) cfsBurstAllowedByLimiter(burstCfg *slov1alpha1.CPUBurstConfig, container *corev1.Container,
	containerID *string) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *cpuBurst) genOperationByContainer(burstCfg *slov1alpha1.CPUBurstConfig, pod *corev1.Pod,
	container *corev1.Container, containerStat *corev1.ContainerStatus) cfsOperation {
	_ = "STUB: not implemented"
	return *new(cfsOperation)
}

func (b *cpuBurst) applyContainerCFSQuota(podMeta *statesinformer.PodMeta, containerStat *corev1.ContainerStatus,
	curContaienrCFS, deltaContainerCFS int64) error {
	_ = "STUB: not implemented"
	return nil
}

// no need to adjust pod cpu.cfs_quota if it is already -1

// cfs scale down, order: container->pod

// cfs scale up, order: pod->container

// set cpu.cfs_burst_us for containers
func (b *cpuBurst) applyCPUBurst(burstCfg *slov1alpha1.CPUBurstConfig, podMeta *statesinformer.PodMeta) {
	_ = "STUB: not implemented"
	return
}

// normally cpu burst resource not supported on current system

// end for containers

// normally cpu burst resource not supported on current system

func (b *cpuBurst) Recycle() { _ = "STUB: not implemented"; return }

// container cpu.cfs_burst_us = container.limit * burstCfg.CPUBurstPercent * cfs_period_us
func calcStaticCPUBurstVal(container *corev1.Container, burstCfg *slov1alpha1.CPUBurstConfig) int64 {
	_ = "STUB: not implemented"
	return 0
}

// use node config by default, overlap if pod specify config
func genPodBurstConfig(pod *corev1.Pod, nodeCfg *slov1alpha1.CPUBurstConfig) *slov1alpha1.CPUBurstConfig {
	_ = "STUB: not implemented"
	return nil
}

func cpuBurstEnabled(burstPolicy slov1alpha1.CPUBurstPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func cfsQuotaBurstEnabled(burstPolicy slov1alpha1.CPUBurstPolicy) bool {
	_ = "STUB: not implemented"
	return false
}

func changeOperationByNode(nodeState nodeStateForBurst, originOperation cfsOperation) (bool, cfsOperation) {
	_ = "STUB: not implemented"
	return false, *new(cfsOperation)
}
