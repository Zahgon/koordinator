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

package custompriority

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"

	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
	podutil "github.com/koordinator-sh/koordinator/pkg/descheduler/pod"
)

const (
	PluginCustomPriorityName = "CustomPriority"
)

var _ framework.BalancePlugin = &CustomPriority{}

// CustomPriority evicts pods from high priority (expensive) resources to low priority (cheap) resources
// based on user-defined priority order and resource availability.
type CustomPriority struct {
	handle    framework.Handle
	podFilter framework.FilterFunc
	args      *deschedulerconfig.CustomPriorityArgs
}

// NewCustomPriority builds plugin from its arguments while passing a handle
func NewCustomPriority(_ context.Context, args runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

// Name retrieves the plugin name
func (pl *CustomPriority) Name() string { _ = "STUB: not implemented"; return "" }

// Balance extension point implementation for the plugin
func (pl *CustomPriority) Balance(ctx context.Context, nodes []*corev1.Node) *framework.Status {
	_ = "STUB: not implemented"
	return nil
}

// Filter nodes based on NodeSelector

// Classify nodes by priority

// Process eviction from high priority to low priority

// filterNodes filters nodes based on NodeSelector
// todo: make as a common util
func (pl *CustomPriority) filterNodes(nodes []*corev1.Node) ([]*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// classifyNodesByPriority groups nodes by their priority level
func (pl *CustomPriority) classifyNodesByPriority(nodes []*corev1.Node) map[string][]*corev1.Node {
	_ = "STUB: not implemented"
	return nil
}

// processEvictions processes evictions from high priority to low priority
func (pl *CustomPriority) processEvictions(ctx context.Context, priorityNodes map[string][]*corev1.Node) *framework.Status {
	_ = "STUB: not implemented"
	// Sort priorities by eviction order (high priority first)
	return nil
}

// Process evictions respecting the configured mode

// BestEffort (default)

// todo：代码应该复用

// evictFromPriorityToTargets evicts pods from source priority nodes to target priority nodes
func (pl *CustomPriority) evictFromPriorityToTargets(ctx context.Context, sourcePriority string, sourceNodes []*corev1.Node, targetNodes []*corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Get pods on the source node

// Sort pods to improve fit rate: smaller requests first
// todo：看一下这个有没有可以复用的地方

// Check if pods can be evicted to target nodes

// Evict pods

// evictByDrainingNodes drains whole source nodes when all candidate pods can be placed onto target nodes.
func (pl *CustomPriority) evictByDrainingNodes(ctx context.Context, sourcePriority string, sourceNodes []*corev1.Node, targetNodes []*corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// Precompute virtual remaining resources for each target node
// CPU (milli), Memory (bytes), Pods(count)
// todo: 这个其实需要做的再高明一些，有很多其他限制条件会导致Pod无法驱逐

// Sort to improve fit rate

// Try to map all pods to target nodes with virtual capacity reservation

// make a working copy of virtual remaining for this source node attempt

// Non-resource checks and actual-remaining checks

// Virtual remaining checks to avoid overbooking in this batch

// Optionally cordon the source node
// todo: cordon的时机不对，但凡是前面没做检查、后面做检查的了的，都有可能先cordon节点后驱逐Pod

// Evict all assigned pods

// Commit virtual reservations into baseline for subsequent nodes

func (pl *CustomPriority) cordonNode(ctx context.Context, node *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (pl *CustomPriority) uncordonNode(ctx context.Context, node *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

func (pl *CustomPriority) nodeRemainingRequests(node *corev1.Node) (map[corev1.ResourceName]*resource.Quantity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pl *CustomPriority) requestsFit(reqs map[corev1.ResourceName]*resource.Quantity, remaining map[corev1.ResourceName]*resource.Quantity) bool {
	_ = "STUB: not implemented"
	return false
}

func (pl *CustomPriority) subRequests(remaining map[corev1.ResourceName]*resource.Quantity, reqs map[corev1.ResourceName]*resource.Quantity) {
	_ = "STUB: not implemented"
	return
}

// findEvictablePods finds pods that can be evicted to target nodes
func (pl *CustomPriority) findEvictablePods(nodeIndexer podutil.GetPodsAssignedToNodeFunc, pods []*corev1.Pod, targetNodes []*corev1.Node) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// simple early-stop to avoid long scans when nothing fits

// Check if pod can fit on any target node

// getPodRequests gets the resource requests of the pod
func (pl *CustomPriority) getPodRequests(pod *corev1.Pod) map[corev1.ResourceName]*resource.Quantity {
	_ = "STUB: not implemented"
	return nil
}

// DeepCopy() returns a value, we need a pointer

// sortPodsByRequestsAscending sorts pods by CPU (milli) then Memory (bytes) requests ascending
func (pl *CustomPriority) sortPodsByRequestsAscending(pods []*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// deterministic fallback

// validateCustomPriorityArgs validates the plugin arguments
func validateCustomPriorityArgs(args *deschedulerconfig.CustomPriorityArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate that each priority has a unique name

// filterPods creates a filter function for pods based on selectors
func filterPods(podSelectors []deschedulerconfig.CustomPriorityPodSelector) (framework.FilterFunc, error) {
	_ = "STUB: not implemented"
	return *new(framework.FilterFunc), nil
}
