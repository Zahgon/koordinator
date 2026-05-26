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

package loadaware

import (
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/clock"

	"github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/plugins/loadaware/estimator"
)

// podAssignCache stores the NodeMetric and Pod information that has been successfully scheduled or is about to be bound.
//
// The cache must handle multiple goroutines' object delta events concurrently.
//  1. When NodeMetric or Pod add or update event on a new node is received, a new nodeInfo should be created and stored in the cache.
//  2. When NodeMetric or Pod event on an existing node in cache is received, the corresponding nodeInfo should be updated.
//  3. When NodeMetric and Pod on a node are all deleted, the corresponding nodeInfo will be empty and should be deleted.
//
// # Implementation
//
// We implement the cache with nodeInfo items indexed by nodeName and maintain them in following rules:
//  1. To delete empty nodeInfo from cache correctly, handlers should hold both the cache lock and this nodeInfo lock when deleting.
//  2. To minimize event handlers' lock occupation on the entire cache, especially write lock, we choose to lock the nodeInfo
//     first and the cache then for deletion. We don't need to hold read lock the cache once for update,
//     release it and write lock the cache again for deletion if events makes nodeInfo empty.
//  3. To prevent deadlock risk with Rule 2, we MUST not lock any nodeInfo when holding cache lock already,
//     which means we have to unlock the cache before locking the nodeInfo.
//  4. Goroutine might get a write locked nodeInfo from cache because of Rule 3, which another goroutine is adding
//     or deleting object on it. When it grabs the lock, this nodeInfo might be empty and removed from cache.
//     We add a deleted flag in nodeInfo, which indicates this case. For adding or updating object event, we MUST not use
//     the deleted nodeInfo, instead, we should retry to get or create a new one from cache to prevent object event missing.
//     For deleting object event or reading nodeInfo data, we can ignore this flag.
//  5. When add or update object, we will try to get a loaded nodeInfo from cache, if not exist, we will create and store a new one.
//     When new nodeInfo is stored, it should be locked already to avoid race conditions on reading brought by Rule 3.
//
// # Concurrent Scenarios
//
// Adding on existing nodeInfo
//   - Alice: <get nodeInfo> - (nodeInfo Lock) - (nodeInfo Update) - (nodeInfo Unlock)
//   - Bob: <get nodeInfo> - (nodeInfo Lock after Alice Unlock) - (nodeInfo Update) - (nodeInfo Unlock)
//
// Adding for not existing nodeInfo
//   - Alice: <create locked nodeInfo, store, return> - (nodeInfo Update) - (nodeInfo Unlock)
//   - Bob: <get nodeInfo after Alice Unlock cache> - (nodeInfo Lock after Alice Unlock) - (nodeInfo Update) - (nodeInfo Unlock)
//
// Deleting and adding nodeInfo
//   - Alice: <get nodeInfo> - (nodeInfo Lock) - (nodeInfo Update) - (nodeInfo tag deleted) - <delete nodeInfo> - (nodeInfo Unlock)
//   - Bob: <get nodeInfo when cache is not locked> - (nodeInfo Lock after Alice Unlock) - (nodeInfo find deleted) - (nodeInfo Unlock) - return failed and retry adding
//
// Reading and writing on existing nodeInfo
//   - Alice: <get nodeInfo> - (nodeInfo Lock) - (nodeInfo Update) - (nodeInfo Unlock)
//   - Bob: <get nodeInfo> - (nodeInfo RLock after Alice Unlock)- (nodeInfo Read) - (nodeInfo RUnlock)
//
// Reading and writing on not existing nodeInfo
//   - Alice: <create locked nodeInfo, store, return> - (nodeInfo Update) - (nodeInfo Unlock)
//   - Bob: <get nodeInfo after Alice Unlock cache> - (nodeInfo RLock after Alice Unlock)- (nodeInfo Read) - (nodeInfo RUnlock)
//
// <...> represents cache is locked before action and unlocked after action
type podAssignCache struct {
	items      sync.Map // items stores nodeInfo, using sync.Map because nodes are not frequently added or deleted.
	estimator  estimator.Estimator
	vectorizer ResourceVectorizer
	clock      clock.Clock
	args       *config.LoadAwareSchedulingArgs
}

// nodeInfo stores
//  1. assigned pods on this node
//  2. this node's metric collected by koordlet
//  3. cached calculation result for node usage and estimation.
type nodeInfo struct {
	sync.RWMutex
	// nodeInfo is marked as deleted when podInfos and nodeMetric is empty.
	// Deleted info should not be updated anymore.
	deleted bool
	// podAssignInfo is indexed using the Pod's types.UID
	podInfos map[types.UID]*podAssignInfo

	nodeMetric     *slov1alpha1.NodeMetric
	updateTime     time.Time
	reportInterval time.Duration

	podUsages     map[NamespacedName]ResourceVector
	prodPods      sets.Set[NamespacedName]
	nodeUsage     ResourceVector
	prodUsage     ResourceVector
	aggUsages     map[aggUsageKey]ResourceVector
	nodeDelta     ResourceVector // delta estimated resources of existing pods
	prodDelta     ResourceVector // delta estimated resources of existing prod pods
	nodeEstimated ResourceVector // sum of full estimated resources of all existing pods

	nodeDeltaPods     sets.Set[NamespacedName] // pods that is part of delta estimated, used in logging only
	prodDeltaPods     sets.Set[NamespacedName] // pods that is part of delta estimated for prod, used in logging only
	nodeEstimatedPods sets.Set[NamespacedName] // pods that is part of full estimated, used in logging only
}

type podAssignInfo struct {
	timestamp         time.Time
	pod               *corev1.Pod
	estimated         ResourceVector
	estimatedDeadline time.Time
}

// implements klog.KMetadata
type NamespacedName struct {
	Namespace string
	Name      string
}

func (n NamespacedName) GetName() string { _ = "STUB: not implemented"; return "" }

func (n NamespacedName) GetNamespace() string { _ = "STUB: not implemented"; return "" }

type aggUsageKey struct {
	Type extension.AggregationType
	// aggDuration == 0 means the non-empty maximum period recorded
	Duration time.Duration
}

func newPodAssignCache(estimator estimator.Estimator, vectorizer ResourceVectorizer, args *config.LoadAwareSchedulingArgs) *podAssignCache {
	_ = "STUB: not implemented"
	return nil
}

func (p *podAssignCache) GetNodeMetricAndEstimatedOfExisting(name string, prodPod bool,
	aggregatedDuration metav1.Duration, aggregationType extension.AggregationType, logEnabled bool) (
	nodeMetric *slov1alpha1.NodeMetric, estimated ResourceVector, estimatedPods []NamespacedName, _ error) {
	_ = "STUB: not implemented"
	return nil, *new(ResourceVector), nil, nil
}

func (n *nodeInfo) getTargetAggregatedUsage(aggregatedDuration metav1.Duration, aggregationType extension.AggregationType) ResourceVector {
	_ = "STUB: not implemented"
	// If no specific period is set, the non-empty maximum period recorded by NodeMetrics will be used by default.
	// This is a default policy.
	return *new(ResourceVector)
}

// All values in aggregatedDuration are empty, downgrade to use the values in NodeUsage

func (p *podAssignCache) getPodAssignInfo(nodeName string, pod *corev1.Pod) *podAssignInfo {
	_ = "STUB: not implemented"
	return nil
}

func (p *podAssignCache) getClonedNodeInfo(nodeName string) *nodeInfo {
	_ = "STUB: not implemented"
	return nil
}

func (p *podAssignCache) getNodeInfo(nodeName string) (*nodeInfo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// getOrCreateNodeInfo returns the nodeInfo for the given nodeName.
// If the nodeInfo does not exist, it will be created with a locked mutex
// which prevent the empty nodeInfo is read and used by plugin with RLock before objects' updating.
//
// NOTICE: it should only be called in objects' add or update methods.
func (p *podAssignCache) getOrCreateNodeInfo(nodeName string) (_ *nodeInfo, created bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// tryCleanup cleans up the nodeInfo from cache.items if it's empty.
//
// NOTICE: nodeInfo should be locked before calling this method.
func (p *podAssignCache) tryCleanup(name string, n *nodeInfo) { _ = "STUB: not implemented"; return }

// only delete action has the chance that goroutine holds two locks,
// and the order always will be nodeInfo lock first, then podAssignCache.items lock

// add or update pod with node name provided
func (p *podAssignCache) assign(nodeName string, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// try to use time from PodScheduled condition first

// if PodScheduled condition not found, fallback to use assign timestamp from scheduler internal, which cannot be zero.

// We only try 2 times here, to avoid any bug in AddOrUpdate that keeps returning false and retrying forever,
// which will block the whole informer event handling procedure.

// if nodeInfo is created in getOrCreate, it is locked already

func (p *podAssignCache) shouldEstimatePodDeadline(pod *corev1.Pod, timestamp time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// if EstimatedSecondsAfterPodScheduled is set and pod is initialized, ignore EstimatedSecondsAfterPodScheduled
// EstimatedSecondsAfterPodScheduled might be set to a long duration to wait for time consuming init containers in pod.

func (p *podAssignCache) unAssign(nodeName string, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (p *podAssignCache) OnAdd(obj interface{}, isInInitialList bool) {
	_ = "STUB: not implemented"
	return
}

func (p *podAssignCache) OnUpdate(oldObj, newObj interface{}) { _ = "STUB: not implemented"; return }

// pod was not cached

// pod has nodeName & pod become terminated

// pod spec or pod conditions changed, renew cached pod

func (p *podAssignCache) OnDelete(obj interface{}) { _ = "STUB: not implemented"; return }

// AddOrUpdatePod add or update pod to nodeInfo.
// It returns false is nodeInfo is already deleted and caller should get a new nodeInfo and retry.
// Unlock is called whether nodeInfo is locked before calling or not.
func (n *nodeInfo) AddOrUpdatePod(pod *podAssignInfo, locked bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *nodeInfo) DeletePod(name string, uid types.UID, p *podAssignCache) {
	_ = "STUB: not implemented"
	return
}

func (p *podAssignCache) NodeMetricHandler() cache.ResourceEventHandler {
	_ = "STUB: not implemented"
	return *new(cache.ResourceEventHandler)
}

func (p *podAssignCache) AddOrUpdateNodeMetric(metric *slov1alpha1.NodeMetric) {
	_ = "STUB: not implemented"
	// We only try 2 times here, to avoid any bug in AddOrUpdate that keeps returning false and retrying forever,
	// which will block the whole informer event handling procedure.
	return
}

// if nodeInfo is created in getOrCreate, it is locked already

func (p *podAssignCache) DeleteNodeMetric(name string) { _ = "STUB: not implemented"; return }

// AddOrUpdateNodeMetric add or update node metric to nodeInfo.
// It returns false is nodeInfo is already deleted and caller should get a new nodeInfo and retry.
// Unlock is called whether nodeInfo is locked before calling or not.
func (n *nodeInfo) AddOrUpdateNodeMetric(metric *slov1alpha1.NodeMetric, p *podAssignCache, locked bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *nodeInfo) DeleteNodeMetric(name string, p *podAssignCache) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeInfo) addPod(pod *podAssignInfo) { _ = "STUB: not implemented"; return }

// Only use prod pod's usage when both pod claims and node metric reports it as prod.
// 1. pod priority class are updated dynamically
// 2. prod / non prod pod is wrongly reported or terminated pod is leaked in node metrics status

// 1. when usage is not collected
// 2. when pod miss lastest metrics update
// 3. when pod metrics is still in the report interval
// 4. when pod is configured in estimation

func (n *nodeInfo) updatePod(oldPod, newPod *podAssignInfo) { _ = "STUB: not implemented"; return }

// reverse procedure of addPod
func (n *nodeInfo) deletePod(pod *podAssignInfo) { _ = "STUB: not implemented"; return }
