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

package elasticquota

import (
	"context"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/client-go/listers/core/v1"
	policylisters "k8s.io/client-go/listers/policy/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/clientset/versioned"
	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/informers/externalversions"
	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/listers/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/plugins/elasticquota/core"
)

const (
	Name                              = "ElasticQuota"
	MigrateDefaultQuotaGroupsPodCycle = 1 * time.Second
	postFilterKey                     = "PostFilter" + Name
)

type PostFilterState struct {
	skip               bool
	quotaInfo          *core.QuotaInfo
	used               corev1.ResourceList
	nonPreemptibleUsed corev1.ResourceList
	usedLimit          corev1.ResourceList
}

func (p *PostFilterState) Clone() fwktype.StateData {
	_ = "STUB: not implemented"
	return *new(fwktype.StateData)
}

type Plugin struct {
	handle                    fwktype.Handle
	client                    versioned.Interface
	pluginArgs                *config.ElasticQuotaArgs
	scheSharedInformerFactory externalversions.SharedInformerFactory
	quotaLister               v1alpha1.ElasticQuotaLister
	quotaInformer             cache.SharedIndexInformer
	podLister                 v1.PodLister
	pdbLister                 policylisters.PodDisruptionBudgetLister
	nodeLister                v1.NodeLister
	groupQuotaManager         *core.GroupQuotaManager

	quotaManagerLock sync.RWMutex
	// groupQuotaManagersForQuotaTree store the GroupQuotaManager of all quota trees. The key is the quota tree id
	groupQuotaManagersForQuotaTree map[string]*core.GroupQuotaManager

	quotaToTreeMapLock sync.RWMutex
	// quotaToTreeMap store the relationship of quota and quota tree
	// the key is the quota name, the value is the tree id
	quotaToTreeMap map[string]string

	// quotaSnapshot stores the quota snapshot for each quota tree
	// The key is tree ID, the value is the snapshot
	// This snapshot is updated periodically in background and doesn't need to stay in sync with quotaInfoMap
	quotaSnapshotLock sync.RWMutex
	quotaSnapshot     map[string]*core.QuotaSnapshot

	// quotaToTreeMapSnapshot stores a snapshot of quotaToTreeMap
	// This snapshot is updated periodically in background and doesn't need to stay in sync with quotaToTreeMap
	quotaToTreeMapSnapshotLock sync.RWMutex
	quotaToTreeMapSnapshot     map[string]string
}

var (
	_ fwktype.EnqueueExtensions            = &Plugin{}
	_ fwktype.PreFilterPlugin              = &Plugin{}
	_ fwktype.PostFilterPlugin             = &Plugin{}
	_ fwktype.ReservePlugin                = &Plugin{}
	_ frameworkext.InformerFactoryProvider = &Plugin{}
)

func New(ctx context.Context, args runtime.Object, handle fwktype.Handle) (fwktype.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.Plugin), nil
}

func (g *Plugin) Start() { _ = "STUB: not implemented"; return }

// Start background goroutine to periodically update quota parent snapshot

func (g *Plugin) NewControllers() ([]frameworkext.Controller, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

func (g *Plugin) EventsToRegister(_ context.Context) ([]fwktype.ClusterEventWithHint, error) {
	_ = "STUB: not implemented"
	// To register a custom event, follow the naming convention at:
	// https://git.k8s.io/kubernetes/pkg/scheduler/eventhandlers.go#L403-L410
	return nil, nil
}

// Only set QueueingHintFn if EnableQueueHint is enabled

func (g *Plugin) PreFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) (*fwktype.PreFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Plugin) PreFilterExtensions() fwktype.PreFilterExtensions {
	_ = "STUB: not implemented"

	// getQuotaSnapshot gets quota snapshot for the given tree ID
	// Returns the snapshot and whether it exists
	return *new(fwktype.PreFilterExtensions)
}

func (g *Plugin) getQuotaSnapshot(treeID string) (*core.QuotaSnapshot, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// getQuotaToTreeMapCopy creates a copy of quotaToTreeMap
// This is thread-safe and returns a new map that can be safely used without locking
func (g *Plugin) getQuotaToTreeMapCopy() map[string]string { _ = "STUB: not implemented"; return nil }

// getPodAssociateQuotaNameAndTreeIDFromSnapshot gets quota name and tree ID using snapshot
// This avoids locking quotaToTreeMap
func (g *Plugin) getPodAssociateQuotaNameAndTreeIDFromSnapshot(pod *corev1.Pod) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// If not found in snapshot, fallback to default quota

// isSchedulableAfterQuotaChanged determines if a pod becomes schedulable after quota is updated.
// QueueAfterBackoff is default queueingHintFn behavior.
func (g *Plugin) isSchedulableAfterQuotaChanged(logger klog.Logger, pod *corev1.Pod, oldObj, newObj interface{}) (fwktype.QueueingHint, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.QueueingHint), nil
}

// Use snapshot to get pod quota name and tree ID without locking

// Check if modified quota is in the same tree as the pod

// Create quota info from original and modified quota

// Quota has changed, check if modified quota is in the pod's path to root

// Modified quota is the pod's quota, allow queueing

// Modified quota is not the pod's quota, check if it's in the path to root

// Get the path from pod's quota to root using the snapshot

// Modified quota is not in the pod's path to root, skip

// isSchedulableAfterPodDeletion determines if a pod becomes schedulable after another pod is deleted.
// QueueAfterBackoff is default queueingHintFn behavior.
func (g *Plugin) isSchedulableAfterPodDeletion(logger klog.Logger, pod *corev1.Pod, oldObj, newObj interface{}) (fwktype.QueueingHint, error) {
	_ = "STUB: not implemented"
	return *new(fwktype.QueueingHint), nil
}

// Use snapshot to get quota names and tree IDs without locking

// Check if deleted pod and unschedulable pod are in the same tree

// Get quota info from snapshot and check if pod is assigned

// Deleted pod is in the same quota as the unschedulable pod, allow queueing

// Check if deleted pod's quota is in the unschedulable pod's path to root

// Get the path from unschedulable pod's quota to root using the snapshot

// Deleted pod's quota is in the path to root, allow queueing

// Deleted pod's quota is not in the unschedulable pod's path to root, skip

// AddPod is called by the framework while trying to evaluate the impact
// of adding podToAdd to the node while scheduling podToSchedule.
func (g *Plugin) AddPod(ctx context.Context, state fwktype.CycleState, podToSchedule *corev1.Pod, podInfoToAdd fwktype.PodInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// RemovePod is called by the framework while trying to evaluate the impact
// of removing podToRemove from the node while scheduling podToSchedule.
func (g *Plugin) RemovePod(ctx context.Context, state fwktype.CycleState, podToSchedule *corev1.Pod, podInfoToRemove fwktype.PodInfo, nodeInfo fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// PostFilter modify the defaultPreemption, only allow pods in the same quota can preempt others.
func (g *Plugin) PostFilter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, filteredNodeStatusMap fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Plugin) Reserve(ctx context.Context, state fwktype.CycleState, p *corev1.Pod, nodeName string) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (g *Plugin) Unreserve(ctx context.Context, state fwktype.CycleState, p *corev1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

func (g *Plugin) GetQuotaInformer() cache.SharedIndexInformer {
	_ = "STUB: not implemented" // expose for extensions
	return *new(cache.SharedIndexInformer)
}

// GetInformerFactories returns the ElasticQuota informer factory for central startup management.
func (g *Plugin) GetInformerFactories() []frameworkext.SharedInformerFactory {
	_ = "STUB: not implemented"
	return nil
}

// updateQuotaSnapshot periodically updates quota snapshot for all quota trees
// This runs in background and doesn't block the main scheduling path
func (g *Plugin) updateQuotaSnapshot() {
	_ = "STUB: not implemented"
	// Copy quotaToTreeMap
	return
}

// Get managers and generate snapshots

// Update snapshots
