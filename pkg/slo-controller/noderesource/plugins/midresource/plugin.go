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

package midresource

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/clock"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	"github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/noderesource/framework"
)

const PluginName = "MidResource"

// ResourceNames defines the Mid-tier extended resource names to update.
var ResourceNames = []corev1.ResourceName{extension.MidCPU, extension.MidMemory}

var clk clock.WithTickerAndDelayedExecution = clock.RealClock{} // for testing

type Plugin struct{}

func (p *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) NeedSync(strategy *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) (bool, string) {
	_ = "STUB: not implemented"
	// mid resource diff is bigger than ResourceDiffThreshold
	return false, ""
}

func (p *Plugin) Prepare(_ *configuration.ColocationStrategy, node *corev1.Node, nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Reset(node *corev1.Node, message string) []framework.ResourceItem {
	_ = "STUB: not implemented"
	return nil
}

// Calculate calculates Mid resources using the formula below:
// if midReclaimMode = "static":
// NodeReclaimable[Mid] = NodeCapacity * staticReservedRatio
// else:
// NodeReclaimable[Mid] = min(NodeMetricReclaimable[Mid], NodeUnused) + Unallocated[Mid] * midUnallocatedRatio
// Allocatable[Mid] = min(NodeReclaimable[Mid], NodeAllocatable * midThresholdRatio)
func (p *Plugin) Calculate(strategy *configuration.ColocationStrategy, node *corev1.Node, podList *corev1.PodList,
	metrics *framework.ResourceMetrics) ([]framework.ResourceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the node metric is abnormal, do degraded calculation

func (p *Plugin) isDegradeNeeded(strategy *configuration.ColocationStrategy, nodeMetric *slov1alpha1.NodeMetric, node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Plugin) degradeCalculate(node *corev1.Node, message string) []framework.ResourceItem {
	_ = "STUB: not implemented"
	return nil
}

// Unallocated[Mid] = max(NodeCapacity - NodeReserved - Allocated[Prod], 0)
func (p *Plugin) getUnallocated(nodeName string, podList *corev1.PodList, nodeCapacity, nodeReserved corev1.ResourceList) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

// If the pod is not marked as low priority, it is considered high priority

func (p *Plugin) calculate(strategy *configuration.ColocationStrategy, node *corev1.Node, podList *corev1.PodList,
	resourceMetrics *framework.ResourceMetrics) []framework.ResourceItem {
	_ = "STUB: not implemented"
	/*
		if midReclaimMode = "static":
		NodeReclaimable[Mid] = NodeCapacity * staticReservedRatio
		else:
		NodeReclaimable[Mid] = min(NodeMetricReclaimable[Mid], NodeUnused) + Unallocated[Mid] * midUnallocatedRatio
		Allocatable[Mid] = min(NodeReclaimable[Mid], NodeAllocatable * midThresholdRatio)
	*/return nil
}

// resource usage of host applications with prod priority will be count as host system usage since they consume the
// node reserved resource.

// System.Reserved = Node.Anno.Reserved, Node.Kubelet.Reserved)

// FIXME: resource reservation taking max is rather confusing.

// failed to get nodeUsage, so radically belief that there is no resource left
// to keep mid-resource calculations relatively strict

// in milli-cores

func getNodeUnused(node *corev1.Node, nodeMetrics *slov1alpha1.NodeMetric) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	// nodeCapacity - nodeUsed
	return *new(corev1.ResourceList), nil
}
