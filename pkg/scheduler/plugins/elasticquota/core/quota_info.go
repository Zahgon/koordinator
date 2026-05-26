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

package core

import (
	"sync"

	v1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
)

type QuotaCalculateInfo struct {
	// The semantics of "max" is the quota group's upper limit of resources.
	Max v1.ResourceList
	// The semantics of "min" is the quota group's guaranteed resources, if quota group's "request" less than or
	// equal to "min", the quota group can obtain equivalent resources to the "request"
	Min v1.ResourceList
	// If Child's sumMin is larger than totalResource, the value of Min should be scaled in equal proportion
	// to ensure the correctness and fairness of min
	AutoScaleMin v1.ResourceList
	// All assigned pods used
	Used v1.ResourceList
	// All non-preemptible pods used
	NonPreemptibleUsed v1.ResourceList
	// All pods request
	Request v1.ResourceList
	// All non-preemptible pods request
	NonPreemptibleRequest v1.ResourceList
	// ChildRquest is the sum of child quota requests.
	// If the quota is leaf, it's the sum of pods requests
	ChildRequest              v1.ResourceList
	SelfRequest               v1.ResourceList
	SelfUsed                  v1.ResourceList
	SelfNonPreemptibleUsed    v1.ResourceList
	SelfNonPreemptibleRequest v1.ResourceList
	// SharedWeight determines the ability of quota groups to compete for shared resources
	SharedWeight v1.ResourceList
	// Runtime is the current actual resource that can be used by the quota group
	Runtime v1.ResourceList

	// If the allocated is greater than min, the guaranteed resource is the allocated,
	// else the guaranteed is the min.
	Guaranteed v1.ResourceList
	// Allocated is the allocated resource. It's the sum of children quota guarantee. If the quota is leaf, it's
	// the sum of scheduled pods
	Allocated v1.ResourceList
}

type QuotaInfo struct {
	// Name
	Name string
	// Quota's ParentName
	ParentName string
	// IsParent quota group
	IsParent bool
	// If runtimeVersion not equal to quotaTree runtimeVersion, means runtime has been updated.
	RuntimeVersion int64
	// Allow lent resource to other quota group
	AllowLentResource bool
	CalculateInfo     QuotaCalculateInfo
	PodCache          map[string]*PodInfo
	lock              sync.RWMutex
}

func NewQuotaInfo(isParent, allowLentResource bool, name, parentName string) *QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

func (qi *QuotaInfo) DeepCopy() *QuotaInfo { _ = "STUB: not implemented"; return nil }

func (qi *QuotaInfo) GetQuotaSummary(treeID string, includePods bool) *QuotaInfoSummary {
	_ = "STUB: not implemented"
	return nil
}

// updateQuotaInfoFromRemote the CRD(max/oriMin/sharedWeight/allowLentResource/isParent/ParentName) of the quota maybe changed,
// so need update localQuotaInfo's information from inputQuotaInfo.
func (qi *QuotaInfo) updateQuotaInfoFromRemote(quotaInfo *QuotaInfo) {
	_ = "STUB: not implemented"
	return
}

// getLimitRequestNoLock returns the min value of request and max, as max is the quotaGroup's upper limit of resources.
// As the multi-hierarchy quota Model described in the PR, when passing a request upwards, passing a request exceeding its
// max will result in a wrong/invalid runtime distribution. For example, parentQuotaGroup's Max is 20, childGroup's Max
// is 10, and the childGroup's request is 30. If the child passes 30 request upwards and get a 20 runtime back
// (limited by the parent's max is 20), the child can only use 10 (limited by its max).
func (qi *QuotaInfo) getLimitRequestNoLock() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

// req > max, limitRequest = max

func (qi *QuotaInfo) setMaxNoLock(max v1.ResourceList) { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) setMinNoLock(min v1.ResourceList) { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) addRequestNonNegativeNoLock(delta, deltaNonPreemptibleRequest v1.ResourceList, isSelfRequest bool) {
	_ = "STUB: not implemented"
	return
}

func (qi *QuotaInfo) addChildRequestNonNegativeNoLock(delta v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func (qi *QuotaInfo) GetGuaranteed() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetAllocated() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) addUsedNonNegativeNoLock(delta, deltaNonPreemptibleUsed v1.ResourceList, isSelfUsed bool) {
	_ = "STUB: not implemented"
	return
}

func (qi *QuotaInfo) addAllocatedQuotaNoLock(delta v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func (qi *QuotaInfo) setMaxQuotaNoLock(res v1.ResourceList) { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) setMinQuotaNoLock(res v1.ResourceList) { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) setAutoScaleMinQuotaNoLock(res v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func (qi *QuotaInfo) setSharedWeightNoLock(res v1.ResourceList) { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) GetRequest() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetChildRequest() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetUsed() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetNonPreemptibleUsed() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetNonPreemptibleRequest() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetSelfRequest() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetSelfUsed() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetSelfNonPreemptibleUsed() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetSelfNonPreemptibleRequest() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetRuntime() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetMax() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) GetMin() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func NewQuotaInfoFromQuota(quota *v1alpha1.ElasticQuota) *QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

func (qi *QuotaInfo) getMaskedRuntimeNoLock() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (qi *QuotaInfo) clearForResetNoLock() { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) IsQuotaMetaChange(quotaInfo *QuotaInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (qi *QuotaInfo) IsQuotaChange(quotaInfo *QuotaInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (qi *QuotaInfo) IsQuotaParentChange(quotaInfo *QuotaInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (qi *QuotaInfo) IsPodExist(pod *v1.Pod) bool { _ = "STUB: not implemented"; return false }

func (qi *QuotaInfo) addPodIfNotPresent(pod *v1.Pod) { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) removePodIfPresent(pod *v1.Pod) { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) UpdatePodIsAssigned(pod *v1.Pod, isAssigned bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (qi *QuotaInfo) GetPodCache() map[string]*v1.Pod { _ = "STUB: not implemented"; return nil }

func (qi *QuotaInfo) CheckPodIsAssigned(pod *v1.Pod) bool { _ = "STUB: not implemented"; return false }

func (qi *QuotaInfo) GetPodThatIsAssigned() []*v1.Pod { _ = "STUB: not implemented"; return nil }

func (qi *QuotaInfo) Lock() { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) UnLock() {
	_ = "STUB: not implemented"

	// QuotaTopoNode only contains the topology of the parent/child relationship,
	// helps to reconstruct quotaTree from the rootQuotaGroup to all the leafQuotaNode.
	return
}

type QuotaTopoNode struct {
	name                 string
	quotaInfo            *QuotaInfo
	parQuotaTopoNode     *QuotaTopoNode
	childGroupQuotaInfos map[string]*QuotaTopoNode
}

func NewQuotaTopoNode(name string, quotaInfo *QuotaInfo) *QuotaTopoNode {
	_ = "STUB: not implemented"
	return nil
}

// not deepCopy

func (qtn *QuotaTopoNode) addChildGroupQuotaInfo(childNode *QuotaTopoNode) {
	_ = "STUB: not implemented"
	return
}

func (qtn *QuotaTopoNode) getChildGroupQuotaInfos() map[string]*QuotaTopoNode {
	_ = "STUB: not implemented"
	return nil
}

type PodInfo struct {
	pod        *v1.Pod
	isAssigned bool
	resource   v1.ResourceList
}

func NewPodInfo(pod *v1.Pod) *PodInfo { _ = "STUB: not implemented"; return nil }

func (pInfo *PodInfo) DeepCopy() *PodInfo { _ = "STUB: not implemented"; return nil }

func (pInfo *PodInfo) GetPod() *v1.Pod { _ = "STUB: not implemented"; return nil }

func generatePodCacheKey(pod *v1.Pod) string { _ = "STUB: not implemented"; return "" }
