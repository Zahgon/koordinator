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

	gocache "github.com/patrickmn/go-cache"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"

	koordslolisters "github.com/koordinator-sh/koordinator/pkg/client/listers/slo/v1alpha1"
	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

const (
	LowNodeLoadName = "LowNodeLoad"
)

var _ framework.BalancePlugin = &LowNodeLoad{}

// LowNodeLoad evicts pods from overutilized nodes to underutilized nodes.
// Note that the plugin refers to the actual usage of the node.
type LowNodeLoad struct {
	handle               framework.Handle
	podFilter            framework.FilterFunc
	nodeMetricLister     koordslolisters.NodeMetricLister
	args                 *deschedulerconfig.LowNodeLoadArgs
	nodeAnomalyDetectors *gocache.Cache
	prodAnomalyDetectors *gocache.Cache
}

// NewLowNodeLoad builds plugin from its arguments while passing a handle
func NewLowNodeLoad(ctx context.Context, args runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

// Name retrieves the plugin name
func (pl *LowNodeLoad) Name() string { _ = "STUB: not implemented"; return "" }

// TODO(joseph): Do we need to consider filtering out nodes of certain specifications?
//  Consider a cluster with nodes of various specifications. The large specification is 96C512GiB, while the small one may be 2C4GiB.
//  It is very likely that the nodes with small specifications will be frequently descheduled.
//  Even some nodes have high utilization, but in fact, the utilization of system components is high,
//  while the utilization of applications is low.
//  Similarly, because there are two Pod filtering mechanisms, EvictableNamespaces and PodSelectors,
//  it is possible that the utilization rate of the filtered Pods is higher than that of the candidate Pods to be descheduled.

// Balance extension point implementation for the plugin
func (pl *LowNodeLoad) Balance(ctx context.Context, nodes []*corev1.Node) *framework.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pl *LowNodeLoad) processOneNodePool(ctx context.Context, nodePool *deschedulerconfig.LowNodeLoadNodePool, nodes []*corev1.Node, processedNodes sets.String) *framework.Status {
	_ = "STUB: not implemented"
	return nil
}

func resetNodesAsNormal(lowNodes []NodeInfo, nodeAnomalyDetectors *gocache.Cache) {
	_ = "STUB: not implemented"
	return
}

func tryMarkNodesAsNormal(nodes []NodeInfo, nodeAnomalyDetectors *gocache.Cache) {
	_ = "STUB: not implemented"
	return
}

func filterRealAbnormalNodes(sourceNodes []NodeInfo, nodeAnomalyDetectors *gocache.Cache, anomalyCondition *deschedulerconfig.LoadAnomalyCondition) []NodeInfo {
	_ = "STUB: not implemented"
	return nil
}

func newThresholds(useDeviationThresholds bool, low, high, lowProd, highProd deschedulerconfig.ResourceThresholds) (thresholds, highThresholds, prodThreshold, highProdThreshold deschedulerconfig.ResourceThresholds) {
	_ = "STUB: not implemented"
	return *new(deschedulerconfig.ResourceThresholds), *new(deschedulerconfig.ResourceThresholds), *new(deschedulerconfig.ResourceThresholds), *new(deschedulerconfig.ResourceThresholds)
}

func lowThresholdFilter(usage *NodeUsage, threshold NodeThresholds) bool {
	_ = "STUB: not implemented"
	return false
}

func prodLowThresholdFilter(usage *NodeUsage, threshold NodeThresholds) bool {
	_ = "STUB: not implemented"
	return false
}

func highThresholdFilter(usage *NodeUsage, threshold NodeThresholds) bool {
	_ = "STUB: not implemented"
	return false
}

func prodHighThresholdFilter(usage *NodeUsage, threshold NodeThresholds) bool {
	_ = "STUB: not implemented"
	return false
}

func filterNodes(nodeSelector *metav1.LabelSelector, nodes []*corev1.Node, processedNodes sets.String) ([]*corev1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterPods(podSelectors []deschedulerconfig.LowNodeLoadPodSelector) (framework.FilterFunc, error) {
	_ = "STUB: not implemented"
	return *new(framework.FilterFunc), nil
}

func logUtilizationCriteria(nodePoolName, message string, lowThresholds, highThresholds, prodLowThresholds, prodHighThresholds deschedulerconfig.ResourceThresholds,
	totalLowNodesNumber, totalHighNodesNumber, prodLowNodesNumber, prodHighNodesNumber, bothLowNodesNumber, totalNumber int) {
	_ = "STUB: not implemented"
	return
}

func overUtilizedEvictionReason(highThresholds, prodHighThresholds deschedulerconfig.ResourceThresholds) evictionReasonGeneratorFn {
	_ = "STUB: not implemented"
	return *new(evictionReasonGeneratorFn)
}
