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

package resctrl

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"

	"github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	ResctrlReconcileName = "ResctrlReconcile"

	// LSRResctrlGroup is the name of LSR resctrl group
	LSRResctrlGroup = "LSR"
	// LSResctrlGroup is the name of LS resctrl group
	LSResctrlGroup = "LS"
	// BEResctrlGroup is the name of BE resctrl group
	BEResctrlGroup = "BE"
	// UnknownResctrlGroup is the resctrl group which is unknown to reconcile
	UnknownResctrlGroup = "Unknown"

	// Max memory bandwidth for AMD CPU, Gb/s, since the extreme limit is hard to reach, we set a discount by 0.8
	// TODO The max memory bandwidth varies across SKU, so koordlet should be aware of the maximum automatically,
	// or support an configuration list.
	// Currently, the value is measured on "AMD EPYC(TM) MILAN"

	AMDCCDMaxMBGbps = 25 * 8 * 0.8

	// the AMD CPU use 2048 to express the unlimited memory bandwidth
	AMDCCDUnlimitedMB = "2048"
)

var (
	// resctrlGroupList is the list of resctrl groups to be reconcile
	resctrlGroupList = []string{LSRResctrlGroup, LSResctrlGroup, BEResctrlGroup}
)

var _ framework.QOSStrategy = &resctrlReconcile{}

type resctrlReconcile struct {
	reconcileInterval time.Duration
	executor          resourceexecutor.ResourceUpdateExecutor
	statesInformer    statesinformer.StatesInformer
	metricCache       metriccache.MetricCache
	cgroupReader      resourceexecutor.CgroupReader
	eventRecorder     record.EventRecorder
}

func New(opt *framework.Options) framework.QOSStrategy {
	_ = "STUB: not implemented"
	return *new(framework.QOSStrategy)
}

func (r *resctrlReconcile) Enabled() bool { _ = "STUB: not implemented"; return false }

func (r *resctrlReconcile) Setup(context *framework.Context) { _ = "STUB: not implemented"; return }

func (r *resctrlReconcile) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (r *resctrlReconcile) init(stopCh <-chan struct{}) {
	r.executor.Run(stopCh)
}

func getPodResctrlGroup(pod *corev1.Pod) string { _ = "STUB: not implemented"; return "" }

func getResourceQOSForResctrlGroup(strategy *slov1alpha1.ResourceQOSStrategy, group string) *slov1alpha1.ResourceQOS {
	_ = "STUB: not implemented"
	return nil
}

func initCatResctrl() error {
	_ = "STUB: not implemented"
	// check if the resctrl root and l3_cat feature are enabled correctly
	return nil
}

func initCatGroupIfNotExist(group string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func calculateMbaPercentForGroup(group string, mbaPercentConfig *int64, cpuBasicInfo extension.CPUBasicInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func calculateIntel(mbaPercent int64) string { _ = "STUB: not implemented"; return "" }

func calculateAMDMba(mbaPercent int64) string { _ = "STUB: not implemented"; return "" }

func (r *resctrlReconcile) getContainerCgroupNewTaskIds(containerParentDir string, tasksMap map[int32]struct{}) ([]int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only append the non-mapped ids

func (r *resctrlReconcile) getPodCgroupNewTaskIds(podMeta *statesinformer.PodMeta, tasksMap map[int32]struct{}) []int32 {
	_ = "STUB: not implemented"
	return nil
}

// reconcile containers

// try retrieve task IDs from the sandbox container, especially for VM-based container runtime

func (r *resctrlReconcile) calculateAndApplyRDTL3PolicyForGroup(group string, cbm uint, l3Num int,
	resourceQoS *slov1alpha1.ResourceQOS) error {
	_ = "STUB: not implemented"
	return nil
}

// calculate policy

// calculate updating resource

// write policy into resctrl files if need update

func (r *resctrlReconcile) calculateAndApplyRDTMbPolicyForGroup(group string, l3Num int, cpuBasicInfo extension.CPUBasicInfo, resourceQoS *slov1alpha1.ResourceQOS) error {
	_ = "STUB: not implemented"
	return nil
}

// calculate updating resource

// write policy into resctrl files if need update

func (r *resctrlReconcile) calculateAndApplyRDTL3GroupTasks(group string, taskIds []int32) error {
	_ = "STUB: not implemented"
	return nil
}

// write policy into resctrl files
// NOTE: the operation should not be cacheable, since old tid has chance to be reused by a new task and here the
// tasks ids are the realtime diff between cgroup and resctrl

func (r *resctrlReconcile) reconcileRDTResctrlPolicy(qosStrategy *slov1alpha1.ResourceQOSStrategy) {
	_ = "STUB: not implemented"
	// 1. retrieve rdt configs from nodeSLOSpec
	// 2.1 get cbm and l3 numbers, which are general for all resctrl groups
	// 2.2 calculate applying resctrl policies, like cat policy and so on, with each rdt config
	// 3. apply the policies onto resctrl groups
	return
}

// read cat l3 cbm

// get the number of l3 caches; it is larger than 0

// calculate and apply l3 cat policy for each group

func (r *resctrlReconcile) reconcileResctrlGroups(qosStrategy *slov1alpha1.ResourceQOSStrategy) {
	_ = "STUB: not implemented"
	// 1. retrieve task ids for each slo by reading cgroup task file of every pod container
	// 2. add the related task ids in resctrl groups
	return
}

// NOTE: pid_max can be found in `/proc/sys/kernel/pid_max` on linux.
// the maximum pid on 32-bit/64-bit platforms is always less than 4194304, so the int type is bigger enough.
// here we only append the task ids which only appear in cgroup but not in resctrl to reduce resctrl writes

// only QoS class level pod are considered

// only Running and Pending pods are considered

// only extension-QoS-specified pod are considered

// TODO https://github.com/koordinator-sh/koordinator/pull/94#discussion_r858779795

// write Cat L3 tasks for each resctrl group

func (r *resctrlReconcile) reconcile() {
	_ = "STUB: not implemented"
	// Step 0. create and init them if resctrl groups do not exist
	// Step 1. reconcile rdt policies against `schemata` file
	// Step 2. reconcile resctrl groups against `tasks` file
	return
}

// Step 0.

// do nothing if nodeSLO == nil || nodeSLO.spec.ResourceStrategy == nil

// skip if host not support resctrl
