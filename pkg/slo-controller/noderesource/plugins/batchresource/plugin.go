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

package batchresource

import (
	topologyv1alpha1 "github.com/k8stopologyawareschedwg/noderesourcetopology-api/pkg/apis/topology/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/clock"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	"github.com/koordinator-sh/koordinator/apis/extension"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/noderesource/framework"
)

const PluginName = "BatchResource"

var (
	ResourceNames = []corev1.ResourceName{extension.BatchCPU, extension.BatchMemory}
)

var (
	Clock          clock.WithTickerAndDelayedExecution = clock.RealClock{} // for testing
	client         ctrlclient.Client
	nrtHandler     *NRTHandler
	nrtSyncContext *framework.SyncContext
)

// Plugin does 2 things:
// 1. calculate and update the extended resources of batch-cpu and batch-memory on the Node.
// 2. calculate and update the zone resources of batch-cpu and batch-memory on the NodeResourceTopology.
type Plugin struct{}

func (p *Plugin) Name() string {
	_ = "STUB: not implemented"

	// +kubebuilder:rbac:groups=topology.node.k8s.io,resources=noderesourcetopologies,verbs=get;list;watch;create;update
	return ""
}

func (p *Plugin) Setup(opt *framework.Option) error { _ = "STUB: not implemented"; return nil }

func (p *Plugin) PreUpdate(strategy *configuration.ColocationStrategy, node *corev1.Node, nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	// prepare for zone resources on NRT objects
	return nil
}

func (p *Plugin) NeedSync(strategy *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) (bool, string) {
	_ = "STUB: not implemented"
	// batch resource diff is bigger than ResourceDiffThreshold
	return false, ""
}

func (p *Plugin) Prepare(_ *configuration.ColocationStrategy, node *corev1.Node, nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	// prepare for node extended resources
	return nil
}

// set origin batch allocatable

//batchMilliCPU := util.GetBatchMilliCPUFromResourceList(node.Status.Allocatable)

// batch resources are reset, no need to recalculate node status

// subtract third party allocated

func (p *Plugin) Reset(node *corev1.Node, message string) []framework.ResourceItem {
	_ = "STUB: not implemented"
	return nil
}

// Calculate calculates Batch resources using the formula below:
// Node.Total - Node.Reserved - System.Used - Pod(High-Priority).Used, System.Used = Node.Used - Pod(All).Used.
// As node and podList are the nearly latest state at time T1, the resourceMetrics are the node metric and pod
// metrics collected and snapshot at time T0 (T0 < T1). There can be gaps between the states of T0 and T1.
// We firstly calculate an infimum of the batch allocatable at time T0.
// `BatchAllocatable0 = NodeAllocatable * ratio - SystemUsed0 - Pod(HP and in Pods1).Used0` - Pod(not in Pods1).Used0.
// Then we minus the sum requests of the pods newly scheduled but have not been reported metrics to give a safe result.
func (p *Plugin) Calculate(strategy *configuration.ColocationStrategy, node *corev1.Node, podList *corev1.PodList,
	resourceMetrics *framework.ResourceMetrics) ([]framework.ResourceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the node metric is abnormal, do degraded calculation

func (p *Plugin) calculate(strategy *configuration.ColocationStrategy, node *corev1.Node, podList *corev1.PodList,
	resourceMetrics *framework.ResourceMetrics) ([]framework.ResourceItem, error) {
	_ = "STUB: not implemented"
	// calculate node-level batch resources
	return nil, nil
}

// calculate NUMA-level batch resources on NodeResourceTopology

// In order to support the colocation requirements of different enterprise environments, a configurable colocation strategy is provided.
// The resource view from the node perspective is as follows:
//
//	https://github.com/koordinator-sh/koordinator/blob/main/docs/images/node-resource-model.png
//
// Typical colocation scenario:
//  1. default policy, and the CPU and memory that can be collocated are automatically calculated based on the load level of the node.
//  2. default policy on CPU, and the Memory is configured not to be overcommitted. This can reduce the probability of batch pods
//     being killed due to high memory water levels (reduce the kill rate)
//
// In each scenario, users can also adjust the resource water level configuration according to your own needs and control the deployment
// density of batch pods.
func (p *Plugin) calculateOnNode(strategy *configuration.ColocationStrategy, node *corev1.Node, podList *corev1.PodList,
	resourceMetrics *framework.ResourceMetrics) (corev1.ResourceList, string, string) {
	_ = "STUB: not implemented"
	// compute the requests and usages according to the pods' priority classes.
	// HP means High-Priority (i.e. not Batch or Free) pods
	return *new(corev1.ResourceList), "", ""
}

// podsAllUsed is the sum usage of all pods reported in NodeMetric.
// podsKnownUsed is the sum usage of pods which are both reported in NodeMetric and shown in current pod list.

// check if the pod has metrics

// count the high-priority usage

// ignore LP pods

// NOTE: Currently qos=LSE pods does not reclaim CPU resource.

// For the pods reported metrics but not shown in current list, count them according to the metric priority.

// resource usage of host applications with prod priority will be count as host system usage since they consumes the
// node reserved resource.

// System.Reserved = Node.Anno.Reserved, Node.Kubelet.Reserved)

// FIXME: resource reservation taking max is rather confusing.

func (p *Plugin) calculateOnNUMALevel(strategy *configuration.ColocationStrategy, node *corev1.Node, podList *corev1.PodList,
	resourceMetrics *framework.ResourceMetrics) (map[string]resource.Quantity, map[string]resource.Quantity, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// separate zone resources
// assert the zone is mapped into NUMA levels
// FIXME: Since NUMA-level metrics are not reported, we use an approximation here:
//        node reservation, system usage and unknown pods usage are the same in each zones.

// resource usage of host applications with prod priority will be count as host system usage since they consumes the
// node reserved resource. bind host app on single numa node is not supported yet. divide the usage by numa node number.

// check if the pod has metrics

// count the high-priority usage

// NOTE: Currently qos=LSE pods does not reclaim CPU resource.

// For the pods reported metrics but not shown in current list, count them according to the metric priority.

func (p *Plugin) isDegradeNeeded(strategy *configuration.ColocationStrategy, nodeMetric *slov1alpha1.NodeMetric, node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *Plugin) degradeCalculate(node *corev1.Node, message string) []framework.ResourceItem {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) prepareForNodeResourceTopology(strategy *configuration.ColocationStrategy, node *corev1.Node,
	nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) updateNRTIfNeeded(strategy *configuration.ColocationStrategy, node *corev1.Node,
	nrt *topologyv1alpha1.NodeResourceTopology, nr *framework.NodeResource) bool {
	_ = "STUB: not implemented"
	return false
}

// update time gap is bigger than UpdateTimeThresholdSeconds
