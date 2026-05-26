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

type GroupQuotaManager struct {
	// hierarchyUpdateLock used for resourceKeys/quotaInfoMap/quotaTreeWrapper change
	hierarchyUpdateLock sync.RWMutex
	// totalResource without systemQuotaGroup and DefaultQuotaGroup's used Quota
	totalResourceExceptSystemAndDefaultUsed v1.ResourceList
	// totalResource with systemQuotaGroup and DefaultQuotaGroup's used Quota
	totalResource v1.ResourceList
	// resourceKeys helps to store runtimeQuotaCalculators' resourceKey
	resourceKeys map[v1.ResourceName]struct{}
	// quotaInfoMap stores all the nodes, it can help get all parents conveniently
	quotaInfoMap map[string]*QuotaInfo
	// runtimeQuotaCalculatorMap helps calculate the subGroups' runtimeQuota in one quotaGroup
	runtimeQuotaCalculatorMap map[string]*RuntimeQuotaCalculator
	// quotaTopoNodeMap only stores the topology of the quota
	quotaTopoNodeMap     map[string]*QuotaTopoNode
	scaleMinQuotaEnabled bool
	// scaleMinQuotaManager is used when overRootResource
	scaleMinQuotaManager *ScaleMinQuotaManager
	once                 sync.Once

	// nodeResourceMapLock  used to lock the nodeResourceMapLock.
	nodeResourceMapLock sync.Mutex
	// nodeResourceMap store the nodes belong to the manager.
	nodeResourceMap map[string]struct{}

	// treeID is the quota tree id
	treeID string

	// hookPlugins contains all registered hookPlugins
	hookPlugins []QuotaHookPlugin
}

func NewGroupQuotaManager(treeID string, enableMinQuotaScale bool, systemGroupMax, defaultGroupMax v1.ResourceList) *GroupQuotaManager {
	_ = "STUB: not implemented"
	return nil
}

// only default GroupQuotaManager need system quota and deault quota.

func (gqm *GroupQuotaManager) setScaleMinQuotaEnabled(flag bool) { _ = "STUB: not implemented"; return }

func (gqm *GroupQuotaManager) UpdateClusterTotalResource(deltaRes v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

// updateClusterTotalResourceNoLock no need to lock gqm.hierarchyUpdateLock and system/defaultQuotaGroup's lock
func (gqm *GroupQuotaManager) updateClusterTotalResourceNoLock(deltaRes v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) GetClusterTotalResource() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (gqm *GroupQuotaManager) SetTotalResourceForTree(total v1.ResourceList) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

// updateGroupDeltaRequestNoLock no need lock gqm.lock
func (gqm *GroupQuotaManager) updateGroupDeltaRequestNoLock(quotaName string, deltaReq, deltaNonPreemptibleRequest v1.ResourceList, selfQuotaIndex int) {
	_ = "STUB: not implemented"
	return
}

// recursiveUpdateGroupTreeWithDeltaRequest update the quota of a node, also need update all parentNode, the lock operation
// of all quotaInfo is done by gqm. scopedLockForQuotaInfo, so just get treeWrappers' lock when calling treeWrappers' function
func (gqm *GroupQuotaManager) recursiveUpdateGroupTreeWithDeltaRequest(deltaReq, deltaNonPreemptibleRequest v1.ResourceList, curToAllParInfos []*QuotaInfo, selfQuotaIndex int) {
	_ = "STUB: not implemented"
	return
}

// If the quota not allow to lent resource. we should request for min

// updateGroupDeltaUsedNoLock updates the usedQuota of a node, it also updates all parent nodes
// no need to lock gqm.hierarchyUpdateLock
func (gqm *GroupQuotaManager) updateGroupDeltaUsedNoLock(quotaName string, delta, deltaNonPreemptibleUsed v1.ResourceList, selfQuotaIndex int) {
	_ = "STUB: not implemented"
	return
}

// if systemQuotaGroup or DefaultQuotaGroup's used change, update cluster total resource.

func (gqm *GroupQuotaManager) RefreshRuntime(quotaName string) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

func (gqm *GroupQuotaManager) refreshRuntimeNoLock(quotaName string) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}

// 1. execute scaleMin logic with totalRes and update scaledMin if needed

// 2. update parent's runtimeQuota

// 3. update subGroup's cluster resource  when i >= 1 (still has children)

// 4. update totalRes

// updateOneGroupAutoScaleMinQuotaNoLock no need to lock gqm.lock
func (gqm *GroupQuotaManager) updateOneGroupAutoScaleMinQuotaNoLock(quotaInfo *QuotaInfo, newMinRes v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) getCurToAllParentGroupQuotaInfoNoLock(quotaName string) []*QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) GetQuotaInfoByName(quotaName string) *QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) getQuotaInfoByNameNoLock(quotaName string) *QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) getRuntimeQuotaCalculatorByNameNoLock(quotaName string) *RuntimeQuotaCalculator {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) scopedLockForQuotaInfo(quotaList []*QuotaInfo) func() {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) UpdateQuota(quota *v1alpha1.ElasticQuota) error {
	_ = "STUB: not implemented"
	return nil
}

// update the local quotaInfo's crd

// if the quotaMeta doesn't change, only runtime/used/request/min/max/sharedWeight change causes update,
// no need to call updateQuotaGroupConfigNoLock.

// update quota internal with pre/post hookPlugins

// update quota internal with pre/post hookPlugins

func (gqm *GroupQuotaManager) DeleteQuota(quota *v1alpha1.ElasticQuota) error {
	_ = "STUB: not implemented"
	return nil
}

// run pre-quota-update hookPlugins

// run post-quota-update hookPlugins

func (gqm *GroupQuotaManager) UpdateQuotaInfo(quota *v1alpha1.ElasticQuota) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) ResetQuota() { _ = "STUB: not implemented"; return }

func (gqm *GroupQuotaManager) resetQuotaNoLock() { _ = "STUB: not implemented"; return }

// rebuild gqm.quotaTopoNodeMap

// reset gqm.runtimeQuotaCalculator

// buildSubParGroupTopoNoLock rebuild a nodeTree from root, no need to lock gqm.lock
func (gqm *GroupQuotaManager) rebuildQuotaTopoNodeMapNoLock() {
	_ = "STUB: not implemented"
	// rebuild QuotaTopoNodeMap
	return
}

// add node according to the quotaInfoMap

// build tree according to the parGroupName

// incase load child before its parent

// rebuildAllGroupQuotaNoLock will reset quota info and runtimeQuotaCalculator
func (gqm *GroupQuotaManager) rebuildAllGroupQuotaNoLock() { _ = "STUB: not implemented"; return }

// clear old runtimeQuotaCalculator

// reset runtimeQuotaCalculator

// subGroup's topo relation may change; refresh the request/used from bottom to top

// ResetAllGroupQuotaRecursiveNoLock no need to lock gqm.lock
func (gqm *GroupQuotaManager) resetAllGroupQuotaRecursiveNoLock(rootNode *QuotaTopoNode) {
	_ = "STUB: not implemented"
	return
}

// updateOneGroupMaxQuotaNoLock no need to lock gqm.lock
func (gqm *GroupQuotaManager) updateOneGroupMaxQuotaNoLock(quotaInfo *QuotaInfo) {
	_ = "STUB: not implemented"
	return
}

// updateMinQuotaNoLock no need to lock gqm.lock
func (gqm *GroupQuotaManager) updateMinQuotaNoLock(quotaInfo *QuotaInfo) {
	_ = "STUB: not implemented"
	return
}

// updateOneGroupOriginalMinQuotaNoLock no need to lock gqm.lock
func (gqm *GroupQuotaManager) updateOneGroupOriginalMinQuotaNoLock(quotaInfo *QuotaInfo) {
	_ = "STUB: not implemented"
	return
}

// updateOneGroupSharedWeightNoLock no need to lock gqm.lock
func (gqm *GroupQuotaManager) updateOneGroupSharedWeightNoLock(quotaInfo *QuotaInfo) {
	_ = "STUB: not implemented"
	return
}

// updateResourceKeyNoLock based on quotaInfo.CalculateInfo.Max of self
// Note: RootQuotaName need to be updated as allResourceKeys
func (gqm *GroupQuotaManager) updateResourceKeyNoLock() {
	_ = "STUB: not implemented"
	// collect all dimensions
	return
}

// update right now

// keep special ones same

func (gqm *GroupQuotaManager) GetAllQuotaNames() map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// QuotaSnapshot is a snapshot of quotaInfoMap
// It contains a copy of all QuotaInfo objects and doesn't need to be kept in sync with quotaInfoMap
type QuotaSnapshot struct {
	lock         sync.RWMutex
	quotaInfoMap map[string]*QuotaInfo
}

// GetQuotaSnapshot returns a snapshot of quotaInfoMap
// This snapshot is taken at a point in time and doesn't need to stay in sync with quotaInfoMap
// It's safe to use this snapshot without holding locks after it's created
func (gqm *GroupQuotaManager) GetQuotaSnapshot() *QuotaSnapshot {
	_ = "STUB: not implemented"
	return nil
}

// Deep copy QuotaInfo for snapshot

// GetQuotaInfoByName returns a QuotaInfo from the snapshot by name
func (s *QuotaSnapshot) GetQuotaInfoByName(quotaName string) *QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

// GetQuotaPathToRoot returns the path from the given quota to root using the snapshot
// This is thread-safe as it only uses the snapshot data without accessing quotaInfoMap
func (s *QuotaSnapshot) GetQuotaPathToRoot(quotaName string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Add current quota to path

// Check if we've reached root

// Get parent from snapshot

func (gqm *GroupQuotaManager) updatePodRequestNoLock(quotaName string, oldPod, newPod *v1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) updatePodUsedNoLock(quotaName string, oldPod, newPod *v1.Pod) {
	_ = "STUB: not implemented"
	return
}

// run pod-update hooks regardless of whether used resources have changed.

func (gqm *GroupQuotaManager) updatePodCacheNoLock(quotaName string, pod *v1.Pod, isAdd bool) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) UpdatePodIsAssigned(quotaName string, pod *v1.Pod, isAssigned bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) updatePodIsAssignedNoLock(quotaName string, pod *v1.Pod, isAssigned bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) getPodIsAssignedNoLock(quotaName string, pod *v1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (gqm *GroupQuotaManager) MigratePod(pod *v1.Pod, out, in string) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) GetQuotaSummary(quotaName string, includePods bool) (*QuotaInfoSummary, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (gqm *GroupQuotaManager) GetQuotaSummaries(includePods bool) map[string]*QuotaInfoSummary {
	_ = "STUB: not implemented"
	return nil
}

// Skip koordinator-root-quota since it's an abstract entity

func (gqm *GroupQuotaManager) OnPodAdd(quotaName string, pod *v1.Pod) {
	_ = "STUB: not implemented"
	return
}

// if the quotaInfo is nil or include the pod, skip it.

// in case failOver, update pod isAssigned explicitly according to its phase and NodeName.

func (gqm *GroupQuotaManager) OnPodUpdate(newQuotaName, oldQuotaName string, newPod, oldPod *v1.Pod) {
	_ = "STUB: not implemented"
	return
}

// it's means the pod creation is before quota creation.

// reserve phase will assign the pod. Just update it.
// upgrade will change the resource.

// assign it

// remove the old resource.

func (gqm *GroupQuotaManager) OnPodDelete(quotaName string, pod *v1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) ReservePod(quotaName string, p *v1.Pod) {
	_ = "STUB: not implemented"
	return
}

// write lock to avoid concurrent pod events:
//  t1: start reserve pod
//  t2: delete pod
//  t3：finish reserve pod

func (gqm *GroupQuotaManager) UnreservePod(quotaName string, p *v1.Pod) {
	_ = "STUB: not implemented"
	return
}

func getPodName(oldPod, newPod *v1.Pod) string { _ = "STUB: not implemented"; return "" }

func (gqm *GroupQuotaManager) OnNodeAdd(node *v1.Node) { _ = "STUB: not implemented"; return }

func (gqm *GroupQuotaManager) OnNodeUpdate(oldNode, newNode *v1.Node) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) OnNodeDelete(node *v1.Node) { _ = "STUB: not implemented"; return }

func (gqm *GroupQuotaManager) GetTreeID() string { _ = "STUB: not implemented"; return "" }

func (gqm *GroupQuotaManager) resetRootQuotaUsedAndRequest() { _ = "STUB: not implemented"; return }

func (gqm *GroupQuotaManager) recursiveUpdateGroupTreeWithDeltaAllocated(deltaAllocated v1.ResourceList, curToAllParInfos []*QuotaInfo) {
	_ = "STUB: not implemented"
	return
}

// update the guarantee.

func (gqm *GroupQuotaManager) updateQuotaInternalNoLock(newQuotaInfo, oldQuotaInfo *QuotaInfo) {
	_ = "STUB: not implemented"
	// update topogy node map
	return
}

// update quota info map

// max changed

// update resource keys

// min changed

// sharedweight changed

func (gqm *GroupQuotaManager) updateQuotaNoLockWhenParentChange(newQuota *v1alpha1.ElasticQuota) {
	_ = "STUB: not implemented"
	// 1. save old quotaInfo and quotaCalculator and quotaTopoNode
	return
}

// 2. delete old quota
// run pre-quota-update hookPlugins

// run post-quota-update hookPlugins

// 3. add new quota info

// clean the min/max/shared weight

// copy pod cache

// run pre-quota-update hookPlugins

// reuse runtimeQuotaCalculator

// 4. update max/min/shared weight

// 5. add requests and used

// run post-quota-update hookPlugins

func (gqm *GroupQuotaManager) updateQuotaTopoNodeNoLock(newQuotaInfo, oldQuotaInfo *QuotaInfo) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) doUpdateOneGroupMaxQuotaNoLock(quotaName string, newMax v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) doUpdateOneGroupMinQuotaNoLock(quotaName string, newMin v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

// Update quota info

// Update request. If the quota not allow to lent resource, the new min will effect the request.

// update parent runtime calculator for min changed

// update the guarantee.

func (gqm *GroupQuotaManager) doUpdateOneGroupSharedWeightNoLock(quotaName string, newSharedWeight v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

func shouldBeIgnored(pod *v1.Pod) bool { _ = "STUB: not implemented"; return false }

func (gqm *GroupQuotaManager) deleteQuotaNoLock(quota *v1alpha1.ElasticQuota) error {
	_ = "STUB: not implemented"
	return nil
}

// handle runtimeQuotaCalculator.

// remove scale min.

// remove topology node.

// update resource keys

// update request

// update used
