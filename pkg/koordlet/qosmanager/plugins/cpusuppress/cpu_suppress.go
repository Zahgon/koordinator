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

package cpusuppress

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	koordletutil "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
	"github.com/koordinator-sh/koordinator/pkg/util/cpuset"
)

const (
	CPUSuppressName = "CPUSuppresss"

	// if destQuota - currentQuota < suppressMinQuotaDeltaRatio * totalCpu; then bypass;
	suppressBypassQuotaDeltaRatio = 0.01

	beMinCPUSetCores        = 2
	beMinQuota              = 2000
	beMaxIncreaseCPUPercent = 0.1 // scale up slow
	beUnsetQuota            = -1
)

type suppressPolicyStatus string

var (
	policyUsing     suppressPolicyStatus = "using"
	policyRecovered suppressPolicyStatus = "recovered"
)

var _ framework.QOSStrategy = &CPUSuppress{}

type CPUSuppress struct {
	interval               time.Duration
	metricCollectInterval  time.Duration
	statesInformer         statesinformer.StatesInformer
	metricCache            metriccache.MetricCache
	executor               resourceexecutor.ResourceUpdateExecutor
	cgroupReader           resourceexecutor.CgroupReader
	suppressPolicyStatuses map[string]suppressPolicyStatus
}

func New(opt *framework.Options) framework.QOSStrategy {
	_ = "STUB: not implemented"
	return *new(framework.QOSStrategy)
}

func (r *CPUSuppress) Enabled() bool { _ = "STUB: not implemented"; return false }

func (r *CPUSuppress) Setup(*framework.Context) { _ = "STUB: not implemented"; return }

func (r *CPUSuppress) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (r *CPUSuppress) init(stopCh <-chan struct{}) {
	r.executor.Run(stopCh)
}

// writeBECgroupsCPUSet writes the be cgroups cpuset by order
func (r *CPUSuppress) writeBECgroupsCPUSet(paths []string, cpusetStr string, isReversed bool) {
	_ = "STUB: not implemented"
	return
}

// calculateBESuppressCPU calculates the quantity of cpuset cpus for suppressing BE pods.
func (r *CPUSuppress) calculateBESuppressCPU(node *corev1.Node, nodeMetric float64, podMetrics map[string]float64,
	podMetas []*statesinformer.PodMeta, hostApps []slov1alpha1.HostApplicationSpec,
	hostAppMetrics map[string]float64, beCPUUsedThreshold int64, beCPUMinThreshold *int64) *resource.Quantity {
	_ = "STUB: not implemented"
	return nil
}

// calculate pod(non-BE).Used and system.Used

// suppress(BE) := node.Capacity * SLOPercent - pod(non-BE).Used - max(system.Used, node.anno.reserved, node.kubelet.reserved)
// NOTE: valid milli-cpu values should not larger than 2^20, so there is no overflow during the calculation

func (r *CPUSuppress) applyBESuppressCPUSet(beCPUSet []int32, oldCPUSet []int32) error {
	_ = "STUB: not implemented"
	return nil
}

// applyCPUSetWithNonePolicy applies the be suppress policy by writing best-effort cgroups
func (r *CPUSuppress) applyCPUSetWithNonePolicy(cpus []int32, oldCPUSet []int32) error {
	_ = "STUB: not implemented"
	// 1. get current be cgroups cpuset
	// 2. temporarily write with a union of old cpuset and new cpuset from upper to lower, to avoid cgroup conflicts
	// 3. write with the new cpuset from lower to upper to apply the real policy
	return nil
}

// write a loose cpuset for all be cgroups before applying the real policy

// apply the suppress policy from lower to upper

func (r *CPUSuppress) applyCPUSetWithStaticPolicy(cpus []int32) error {
	_ = "STUB: not implemented"
	return nil
}

// suppressBECPU adjusts the cpusets of BE pods to suppress BE cpu usage
func (r *CPUSuppress) suppressBECPU() {
	_ = "STUB: not implemented"
	//  1. calculate be suppress threshold and check if the suppress is needed
	//     1.1. retrieve latest node resource usage from the metricCache
	//     1.2  calculate the quantity of be suppress cpuset cpus
	//  2. calculate be suppress policy
	//     2.1. new policy should try to get cpuset cpus scattered by numa node, paired by ht core, no less than 2,
	//     less jitter as far as possible
	//  3. apply best-effort cgroups cpuset or cfsquota
	return
}

// Step 0.

// Step 1.

// Step 2.

func (r *CPUSuppress) adjustByCPUSet(cpusetQuantity *resource.Quantity, nodeCPUInfo *metriccache.NodeCPUInfo) {
	_ = "STUB: not implemented"
	return
}

// value: 0 -> lse, 1 -> lsr, not exists -> others

// system qos exclusive cpuset

// FIXME: be pods might be starved since lse pods can run out of all cpus

// set the number of cpuset cpus no less than 2

// the new be suppress always need to apply since:
// - for a reduce of BE cpuset, we should make effort to protecting LS no matter how huge the decrease is;
// - for a enlargement of BE cpuset, it is welcome and costless for BE processes.

// recover cpuset path as be share pool for the following dirs:
// - besteffort dir
// - besteffort/pod dir
// - besteffort/pod/container if pod does not specify resource status(cpuset/numa node)
func (r *CPUSuppress) recoverCPUSetForBECPUManager() { _ = "STUB: not implemented"; return }

// cpuset path under besteffort dir, include root, pod

// cpuset path under besteffort, only include container/sandbox dir

// cgroup path of pods which has specified resource status

// no resource status or not specified

// exclude cgroupPathOfPodWithSpecifiedCPUSet from cpusetPathOfAllBEContainer

func (r *CPUSuppress) recoverCPUSetIfNeed(maxDepth int) { _ = "STUB: not implemented"; return }

func (r *CPUSuppress) calcBECPUSet() (*cpuset.CPUSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LSE pod, reserved cpu and system qos exclusive

// System pod, exclude cpuset if exclusive

// exclude reserved cpuset

func (r *CPUSuppress) adjustByCfsQuota(cpuQuantity *resource.Quantity, node *corev1.Node) {
	_ = "STUB: not implemented"
	return
}

// read current offline quota

//  delta is large enough

// scale with steps after quota has set

func (r *CPUSuppress) recoverCFSQuotaIfNeed() { _ = "STUB: not implemented"; return }

// calculateBESuppressPolicy calculates the be cpu suppress policy with cpuset cpus number and node cpu info
func calculateBESuppressCPUSetPolicy(cpus int32, processorInfos []koordletutil.ProcessorInfo) []int32 {
	_ = "STUB: not implemented"
	return nil
}

// getNodeIndex is a function to calculate an index for every numa node or socket

// (nodeId, socketId) => nodeIndex

// change cpuBucket map to array

// select same core cpu id

// if we don't pick any cpu, we need break this cycle

// select single cpu id

// if we don't pick any cpu, we need break this cycle

func getSystemQOSExclusiveCPU(nodeTopoAnno map[string]string) (cpuset.CPUSet, error) {
	_ = "STUB: not implemented"
	return *new(cpuset.CPUSet), nil
}

// system qos cpuset exist and exclusive

// parse cpuset string to struct
