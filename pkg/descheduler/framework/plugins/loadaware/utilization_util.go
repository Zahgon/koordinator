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
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	slolisters "github.com/koordinator-sh/koordinator/pkg/client/listers/slo/v1alpha1"
	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
	podutil "github.com/koordinator-sh/koordinator/pkg/descheduler/pod"
)

type Percentage = deschedulerconfig.Percentage
type ResourceThresholds = deschedulerconfig.ResourceThresholds

type NodeUsage struct {
	node       *corev1.Node
	allPods    []*corev1.Pod
	prodPods   []*corev1.Pod
	usage      map[corev1.ResourceName]*resource.Quantity
	prodUsage  map[corev1.ResourceName]*resource.Quantity
	podMetrics map[types.NamespacedName]*slov1alpha1.ResourceMap
}

type NodeThresholds struct {
	lowResourceThreshold      map[corev1.ResourceName]*resource.Quantity
	highResourceThreshold     map[corev1.ResourceName]*resource.Quantity
	prodLowResourceThreshold  map[corev1.ResourceName]*resource.Quantity
	prodHighResourceThreshold map[corev1.ResourceName]*resource.Quantity
}

type NodeInfo struct {
	*NodeUsage
	thresholds NodeThresholds
}

type continueEvictionCond func(nodeInfo NodeInfo, totalAvailableUsages map[corev1.ResourceName]*resource.Quantity, prod bool) bool

type evictionReasonGeneratorFn func(nodeInfo NodeInfo, prod bool) string

const (
	MinResourcePercentage = 0
	MaxResourcePercentage = 100
)

func normalizePercentage(percent Percentage) Percentage {
	_ = "STUB: not implemented"
	return *new(Percentage)
}

func getNodeThresholds(
	nodeUsages map[string]*NodeUsage,
	lowThreshold, highThreshold, prodLowThreshold, prodHighThreshold ResourceThresholds,
	resourceNames []corev1.ResourceName,
	useDeviationThresholds bool,
) map[string]NodeThresholds {
	_ = "STUB: not implemented"
	return nil
}

func resourceThreshold(nodeCapacity corev1.ResourceList, resourceName corev1.ResourceName, threshold Percentage) *resource.Quantity {
	_ = "STUB: not implemented"
	return nil
}

// A threshold is in percentages but in <0;100> interval.
// Performing `threshold * 0.01` will convert <0;100> interval into <0;1>.
// Multiplying it with capacity will give fraction of the capacity corresponding to the given resource threshold in Quantity units.

func getNodeUsage(nodes []*corev1.Node, resourceNames []corev1.ResourceName, nodeMetricLister slolisters.NodeMetricLister, getPodsAssignedToNode podutil.GetPodsAssignedToNodeFunc, nodeMetricExpirationSeconds *int64) map[string]*NodeUsage {
	_ = "STUB: not implemented"
	return nil
}

// We should check if NodeMetric is expired.

func ResetResourceUsageIsZero(resourceName corev1.ResourceName, usageQuantity resource.Quantity) resource.Quantity {
	_ = "STUB: not implemented"
	return *new(resource.Quantity)
}

// classifyNodes classifies the nodes into low-utilization or high-utilization nodes.
// If a node lies between low and high thresholds, it is simply ignored.
func classifyNodes(
	nodeUsages map[string]*NodeUsage,
	nodeThresholds map[string]NodeThresholds,
	lowThresholdFilter, highThresholdFilter, prodLowThresholdFilter, prodHighThresholdFilter func(usage *NodeUsage, threshold NodeThresholds) bool,
) (lowNodes []NodeInfo, highNodes []NodeInfo, prodLowNodes []NodeInfo, prodHighNodes []NodeInfo, bothLowNodes []NodeInfo) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

func resourceUsagePercentages(nodeUsage *NodeUsage, prod bool) map[corev1.ResourceName]float64 {
	_ = "STUB: not implemented"
	return nil
}

func evictPodsFromSourceNodes(
	ctx context.Context,
	nodePoolName string,
	sourceNodes, destinationNodes,
	prodSourceNodes, prodDestinationNodes, bothDestinationNodes []NodeInfo,
	nodeUsages map[string]*NodeUsage,
	nodeThresholds map[string]NodeThresholds,
	dryRun bool,
	nodeFit bool,
	resourceWeights map[corev1.ResourceName]int64,
	podEvictor framework.Evictor,
	podFilter framework.FilterFunc,
	nodeIndexer podutil.GetPodsAssignedToNodeFunc,
	resourceNames []corev1.ResourceName,
	continueEviction continueEvictionCond,
	evictionReasonGenerator evictionReasonGeneratorFn,
) {
	_ = "STUB: not implemented"
	return
}

// bothLowNode will be used by nodeHigh and prodHigh nodes, needs sub resources used by pods on nodeHigh.

// A part of bothTotalAvailableUsage has been used,
// then the remaining part of nodeTotalAvailableUsage can be utilized.

// add min(prodBothAvailableUsage, bothTotalAvailableUsage) to prodTotalAvailableUsages

func newAvailableUsage(resourceNames []corev1.ResourceName) map[corev1.ResourceName]*resource.Quantity {
	_ = "STUB: not implemented"
	return nil
}

func balancePods(ctx context.Context,
	nodePoolName string,
	sourceNodes []NodeInfo,
	targetNodes []*corev1.Node,
	nodeUsages map[string]*NodeUsage,
	nodeThresholds map[string]NodeThresholds,
	totalAvailableUsages map[corev1.ResourceName]*resource.Quantity,
	dryRun bool,
	nodeFit, prod bool,
	resourceWeights map[corev1.ResourceName]int64,
	podEvictor framework.Evictor,
	podFilter framework.FilterFunc,
	nodeIndexer podutil.GetPodsAssignedToNodeFunc,
	continueEviction continueEvictionCond,
	evictionReasonGenerator evictionReasonGeneratorFn) {
	_ = "STUB: not implemented"
	return
}

func targetAvailableUsage(destinationNodes []NodeInfo, resourceNames []corev1.ResourceName, prod bool) (map[corev1.ResourceName]*resource.Quantity, []*corev1.Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

func evictPods(
	ctx context.Context,
	nodePoolName string,
	dryRun bool,
	prod bool,
	inputPods []*corev1.Pod,
	nodeInfo NodeInfo,
	totalAvailableUsages map[corev1.ResourceName]*resource.Quantity,
	podEvictor framework.Evictor,
	podFilter framework.FilterFunc,
	continueEviction continueEvictionCond,
	evictionReasonGenerator evictionReasonGeneratorFn,
) {
	_ = "STUB: not implemented"
	return
}

// sortNodesByUsage sorts nodes based on usage.
func sortNodesByUsage(nodes []NodeInfo, resourceToWeightMap map[corev1.ResourceName]int64, ascending, prod bool) {
	_ = "STUB: not implemented"
	return
}

func usageToResourceList(usage map[corev1.ResourceName]*resource.Quantity) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

func isNodeOverutilized(usage, thresholds map[corev1.ResourceName]*resource.Quantity) (corev1.ResourceList, bool) {
	_ = "STUB: not implemented"
	// At least one resource has to be above the threshold
	return *new(corev1.ResourceList), false
}

func isNodeUnderutilized(usage, thresholds map[corev1.ResourceName]*resource.Quantity) bool {
	_ = "STUB: not implemented"
	// All resources have to be below the low threshold
	return false
}

func isNodeMetricExpired(lastUpdateTime *metav1.Time, nodeMetricExpirationSeconds int64) bool {
	_ = "STUB: not implemented"
	return false
}

func getResourceNames(thresholds ResourceThresholds) []corev1.ResourceName {
	_ = "STUB: not implemented"
	return nil
}

func classifyPods(pods []*corev1.Pod, filter func(pod *corev1.Pod) bool) ([]*corev1.Pod, []*corev1.Pod) {
	_ = "STUB: not implemented"
	return nil, nil
}

func calcAverageResourceUsagePercent(nodeUsages map[string]*NodeUsage) (ResourceThresholds, ResourceThresholds) {
	_ = "STUB: not implemented"
	return *new(ResourceThresholds), *new(ResourceThresholds)
}

func sortPodsOnOneOverloadedNode(srcNode NodeInfo, removablePods []*corev1.Pod, resourceWeights map[corev1.ResourceName]int64, prod bool) {
	_ = "STUB: not implemented"
	return
}

// get the overused resource of this node, and the weights of appropriately using resources will be zero.

// podFitsAnyNodeWithThreshold checks if the given pod will fit any of the given nodes. It also checks if the node
// utilization will exceed the threshold after this pod was scheduled on it.
func podFitsAnyNodeWithThreshold(nodeIndexer podutil.GetPodsAssignedToNodeFunc, pod *corev1.Pod, nodes []*corev1.Node,
	nodeUsages map[string]*NodeUsage, nodeThresholds map[string]NodeThresholds, prod bool, podMetric *slov1alpha1.ResourceMap) bool {
	_ = "STUB: not implemented"
	return false
}

// check if node utilization exceeds threshold if pod scheduled

// revert the change

// GetNodeRawAllocatableFromNode gets the raw allocatable from node annotation.
// In the cpu-normalization or amplification scenario, node Allocatable will be amplified,
// so raw-allocatable needs to be obtained during descheduling to accurately calculate node usage percent.
// If raw-allocatable is not set or fails to parse, returns the amplified Allocatable as fallback.
func GetNodeRawAllocatableFromNode(node *corev1.Node) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}
