/*
Copyright 2022 The Koordinator Authors.
Copyright 2019 The Kubernetes Authors.

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

package node

import (
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	clientset "k8s.io/client-go/kubernetes"
)

const (
	// poll is how often to Poll pods, nodes and claims.
	poll = 2 * time.Second

	// singleCallTimeout is how long to try single API calls (like 'get' or 'list'). Used to prevent
	// transient failures from failing tests.
	singleCallTimeout = 5 * time.Minute

	// ssh port
	sshPort = "22"
)

var (
	// unreachableTaintTemplate is the taint for when a node becomes unreachable.
	// Copied from pkg/controller/nodelifecycle to avoid pulling extra dependencies
	unreachableTaintTemplate = &v1.Taint{
		Key:    v1.TaintNodeUnreachable,
		Effect: v1.TaintEffectNoExecute,
	}

	// notReadyTaintTemplate is the taint for when a node is not ready for executing pods.
	// Copied from pkg/controller/nodelifecycle to avoid pulling extra dependencies
	notReadyTaintTemplate = &v1.Taint{
		Key:    v1.TaintNodeNotReady,
		Effect: v1.TaintEffectNoExecute,
	}

	// updateTaintBackOff contains the maximum retries and the wait interval between two retries.
	updateTaintBackOff = wait.Backoff{
		Steps:    5,
		Duration: 100 * time.Millisecond,
		Jitter:   1.0,
	}
)

// PodNode is a pod-node pair indicating which node a given pod is running on
type PodNode struct {
	// Pod represents pod name
	Pod string
	// Node represents node name
	Node string
}

// FirstAddress returns the first address of the given type of each node.
func FirstAddress(nodelist *v1.NodeList, addrType v1.NodeAddressType) string {
	_ = "STUB: not implemented"
	return ""
}

func isNodeConditionSetAsExpected(node *v1.Node, conditionType v1.NodeConditionType, wantTrue, silent bool) bool {
	_ = "STUB: not implemented"
	// Check the node readiness condition (logging all).
	return false
}

// Ensure that the condition type and the status matches as desired.

// For NodeReady condition we need to check Taints as well

// For NodeReady we need to check if Taints are gone as well

// TODO: check if the Node is tainted once we enable NC notReady/unreachable taints by default

// IsConditionSetAsExpected returns a wantTrue value if the node has a match to the conditionType, otherwise returns an opposite value of the wantTrue with detailed logging.
func IsConditionSetAsExpected(node *v1.Node, conditionType v1.NodeConditionType, wantTrue bool) bool {
	_ = "STUB: not implemented"
	return false
}

// IsConditionSetAsExpectedSilent returns a wantTrue value if the node has a match to the conditionType, otherwise returns an opposite value of the wantTrue.
func IsConditionSetAsExpectedSilent(node *v1.Node, conditionType v1.NodeConditionType, wantTrue bool) bool {
	_ = "STUB: not implemented"
	return false
}

// isConditionUnset returns true if conditions of the given node do not have a match to the given conditionType, otherwise false.
func isConditionUnset(node *v1.Node, conditionType v1.NodeConditionType) bool {
	_ = "STUB: not implemented"
	return false
}

// Filter filters nodes in NodeList in place, removing nodes that do not
// satisfy the given condition
func Filter(nodeList *v1.NodeList, fn func(node v1.Node) bool) { _ = "STUB: not implemented"; return }

// TotalRegistered returns number of schedulable Nodes.
func TotalRegistered(c clientset.Interface) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// TotalReady returns number of ready schedulable Nodes.
func TotalReady(c clientset.Interface) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Filter out not-ready nodes.

// GetExternalIP returns node external IP concatenated with port 22 for ssh
// e.g. 1.2.3.4:22
func GetExternalIP(node *v1.Node) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetInternalIP returns node internal IP
func GetInternalIP(node *v1.Node) (string, error) { _ = "STUB: not implemented"; return "", nil }

// FirstAddressByTypeAndFamily returns the first address that matches the given type and family of the list of nodes
func FirstAddressByTypeAndFamily(nodelist *v1.NodeList, addrType v1.NodeAddressType, family v1.IPFamily) string {
	_ = "STUB: not implemented"
	return ""
}

// GetAddressesByTypeAndFamily returns a list of addresses of the given addressType for the given node
// and filtered by IPFamily
func GetAddressesByTypeAndFamily(node *v1.Node, addressType v1.NodeAddressType, family v1.IPFamily) (ips []string) {
	_ = "STUB: not implemented"
	return nil
}

// GetAddresses returns a list of addresses of the given addressType for the given node
func GetAddresses(node *v1.Node, addressType v1.NodeAddressType) (ips []string) {
	_ = "STUB: not implemented"
	return nil
}

// CollectAddresses returns a list of addresses of the given addressType for the given list of nodes
func CollectAddresses(nodes *v1.NodeList, addressType v1.NodeAddressType) []string {
	_ = "STUB: not implemented"
	return nil
}

// PickIP picks one public node IP
func PickIP(c clientset.Interface) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetPublicIps returns a public IP list of nodes.
func GetPublicIps(c clientset.Interface) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If ExternalIP isn't set, assume the test programs can reach the InternalIP

// GetReadySchedulableNodes addresses the common use case of getting nodes you can do work on.
// 1) Needs to be schedulable.
// 2) Needs to be ready.
// If EITHER 1 or 2 is not true, most tests will want to ignore the node entirely.
// If there are no nodes that are both ready and schedulable, this will return an error.
func GetReadySchedulableNodes(c clientset.Interface) (nodes *v1.NodeList, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBoundedReadySchedulableNodes is like GetReadySchedulableNodes except that it returns
// at most maxNodes nodes. Use this to keep your test case from blowing up when run on a
// large cluster.
func GetBoundedReadySchedulableNodes(c clientset.Interface, maxNodes int) (nodes *v1.NodeList, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRandomReadySchedulableNode gets a single randomly-selected node which is available for
// running pods on. If there are no available nodes it will return an error.
func GetRandomReadySchedulableNode(c clientset.Interface) (*v1.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetReadyNodesIncludingTainted returns all ready nodes, even those which are tainted.
// There are cases when we care about tainted nodes
// E.g. in tests related to nodes with gpu we care about nodes despite
// presence of nvidia.com/gpu=present:NoSchedule taint
func GetReadyNodesIncludingTainted(c clientset.Interface) (nodes *v1.NodeList, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isNodeUntainted tests whether a fake pod can be scheduled on "node", given its current taints.
// TODO: need to discuss wether to return bool and error type
func isNodeUntainted(node *v1.Node) bool { _ = "STUB: not implemented"; return false }

// isNodeUntaintedWithNonblocking tests whether a fake pod can be scheduled on "node"
// but allows for taints in the list of non-blocking taints.
func isNodeUntaintedWithNonblocking(node *v1.Node, nonblockingTaints string) bool {
	_ = "STUB: not implemented"
	return false
}

// Simple lookup for nonblocking taints based on comma-delimited list.

func toleratesTaintsWithNoScheduleNoExecuteEffects(taints []v1.Taint, tolerations []v1.Toleration) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNodeSchedulable returns true if:
// 1) doesn't have "unschedulable" field set
// 2) it also returns true from IsNodeReady
func IsNodeSchedulable(node *v1.Node) bool { _ = "STUB: not implemented"; return false }

// IsNodeReady returns true if:
// 1) it's Ready condition is set to true
// 2) doesn't have NetworkUnavailable condition set to true
func IsNodeReady(node *v1.Node) bool { _ = "STUB: not implemented"; return false }

// isNodeSchedulableWithoutTaints returns true if:
// 1) doesn't have "unschedulable" field set
// 2) it also returns true from IsNodeReady
// 3) it also returns true from isNodeUntainted
func isNodeSchedulableWithoutTaints(node *v1.Node) bool { _ = "STUB: not implemented"; return false }

// hasNonblockingTaint returns true if the node contains at least
// one taint with a key matching the regexp.
func hasNonblockingTaint(node *v1.Node, nonblockingTaints string) bool {
	_ = "STUB: not implemented"
	return false
}

// Simple lookup for nonblocking taints based on comma-delimited list.

// PodNodePairs return podNode pairs for all pods in a namespace
func PodNodePairs(c clientset.Interface, ns string) ([]PodNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetClusterZones returns the values of zone label collected from all nodes.
func GetClusterZones(c clientset.Interface) (sets.String, error) {
	_ = "STUB: not implemented"
	return *new(sets.String), nil
}

// collect values of zone label from all nodes

// GetSchedulableClusterZones returns the values of zone label collected from all nodes which are schedulable.
func GetSchedulableClusterZones(c clientset.Interface) (sets.String, error) {
	_ = "STUB: not implemented"
	return *new(sets.String), nil
}

// collect values of zone label from all nodes

// We should have at least 1 node in the zone which is schedulable.

// CreatePodsPerNodeForSimpleApp creates pods w/ labels.  Useful for tests which make a bunch of pods w/o any networking.
func CreatePodsPerNodeForSimpleApp(c clientset.Interface, namespace, appName string, podSpec func(n v1.Node) v1.PodSpec, maxCount int) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// TODO use wrapper methods in expect.go after removing core e2e dependency on node

// TODO use wrapper methods in expect.go after removing core e2e dependency on node

// RemoveTaintsOffNode removes a list of taints from the given node
// It is simply a helper wrapper for RemoveTaintOffNode
func RemoveTaintsOffNode(c clientset.Interface, nodeName string, taints []v1.Taint) {
	_ = "STUB: not implemented"
	return
}

// RemoveTaintOffNode removes the given taint from the given node.
func RemoveTaintOffNode(c clientset.Interface, nodeName string, taint v1.Taint) {
	_ = "STUB: not implemented"
	return
}

// TODO use wrapper methods in expect.go after removing core e2e dependency on node

// AddOrUpdateTaintOnNode adds the given taint to the given node or updates taint.
func AddOrUpdateTaintOnNode(c clientset.Interface, nodeName string, taint v1.Taint) {
	_ = "STUB: not implemented"
	// TODO use wrapper methods in expect.go after removing the dependency on this
	// package from the core e2e framework.
	return
}

// addOrUpdateTaintOnNode add taints to the node. If taint was added into node, it'll issue API calls
// to update nodes; otherwise, no API calls. Return error if any.
// copied from pkg/controller/controller_utils.go AddOrUpdateTaintOnNode()
func addOrUpdateTaintOnNode(c clientset.Interface, nodeName string, taints ...*v1.Taint) error {
	_ = "STUB: not implemented"
	return nil
}

// First we try getting node from the API server cache, as it's cheaper. If it fails
// we get it from etcd to be sure to have fresh data.

// addOrUpdateTaint tries to add a taint to annotations list. Returns a new copy of updated Node and true if something was updated
// false otherwise.
// copied from pkg/util/taints/taints.go AddOrUpdateTaint()
func addOrUpdateTaint(node *v1.Node, taint *v1.Taint) (*v1.Node, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// semantic can do semantic deep equality checks for core objects.
// Example: apiequality.Semantic.DeepEqual(aPod, aPodWithNonNilButEmptyMaps) == true
// copied from pkg/apis/core/helper/helpers.go Semantic
var semantic = conversion.EqualitiesOrDie(
	func(a, b resource.Quantity) bool {
		// Ignore formatting, only care that numeric value stayed the same.
		// TODO: if we decide it's important, it should be safe to start comparing the format.
		//
		// Uninitialized quantities are equivalent to 0 quantities.
		return a.Cmp(b) == 0
	},
	func(a, b metav1.MicroTime) bool {
		return a.UTC() == b.UTC()
	},
	func(a, b metav1.Time) bool {
		return a.UTC() == b.UTC()
	},
	func(a, b labels.Selector) bool {
		return a.String() == b.String()
	},
	func(a, b fields.Selector) bool {
		return a.String() == b.String()
	},
)

// removeNodeTaint is for cleaning up taints temporarily added to node,
// won't fail if target taint doesn't exist or has been removed.
// If passed a node it'll check if there's anything to be done, if taint is not present it won't issue
// any API calls.
func removeNodeTaint(c clientset.Interface, nodeName string, node *v1.Node, taints ...*v1.Taint) error {
	_ = "STUB: not implemented"
	return nil
}

// Short circuit for limiting amount of API calls.

// First we try getting node from the API server cache, as it's cheaper. If it fails
// we get it from etcd to be sure to have fresh data.

// patchNodeTaints patches node's taints.
func patchNodeTaints(c clientset.Interface, nodeName string, oldNode *v1.Node, newNode *v1.Node) error {
	_ = "STUB: not implemented"
	return nil
}

// removeTaint tries to remove a taint from annotations list. Returns a new copy of updated Node and true if something was updated
// false otherwise.
func removeTaint(node *v1.Node, taint *v1.Taint) (*v1.Node, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// deleteTaint removes all the taints that have the same key and effect to given taintToDelete.
func deleteTaint(taints []v1.Taint, taintToDelete *v1.Taint) ([]v1.Taint, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func verifyThatTaintIsGone(c clientset.Interface, nodeName string, taint *v1.Taint) {
	_ = "STUB: not implemented"
	return
}

// TODO use wrapper methods in expect.go after removing core e2e dependency on node

// taintExists checks if the given taint exists in list of taints. Returns true if exists false otherwise.
func taintExists(taints []v1.Taint, taintToFind *v1.Taint) bool {
	_ = "STUB: not implemented"
	return false
}
