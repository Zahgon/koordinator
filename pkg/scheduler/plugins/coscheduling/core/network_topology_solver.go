package core

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/parallelize"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/networktopology"
)

const (
	MessageNoCandidateTopologyNodes = "no candidate topology nodes can accommodate job, desiredOfferSlot: %d, %s, %s"
)

const (
	OperationCalculateNodeOfferSlot      = "CalculateNodeOfferSlot"
	OperationCalculateNodeExistingPodNum = "CalculateNodeExistingPodsNum"
)

type NetworkTopologySolver interface {
	// PlacePods place pods on nodes according to network topology. Please assure nodes is already cloned.
	// TODO Currently, only one score is supported for each node. Subsequent algorithms may need to support calling different plugins to score nodes and support configuring plugin weights.
	PlacePods(
		ctx context.Context,
		cycleStates map[string]fwktype.CycleState,
		toSchedulePods []*corev1.Pod,
		nodes []fwktype.NodeInfo,
		addPod podFunc,
		jobNetworkRequirements *JobTopologyRequirements,
		clusterNetworkTopology *networktopology.TreeNode,
		nodeToScore map[string]int,
	) (podToNode map[string]string, status *fwktype.Status)
}

var (
	_ NetworkTopologySolver = &networkTopologySolverImpl{}
)

type networkTopologySolverImpl struct {
	handle frameworkext.ExtendedHandle
}

func (solver *networkTopologySolverImpl) PlacePods(
	ctx context.Context,
	cycleStates map[string]fwktype.CycleState,
	toSchedulePods []*corev1.Pod,
	nodes []fwktype.NodeInfo,
	addPod podFunc,
	jobNetworkRequirements *JobTopologyRequirements,
	clusterNetworkTopology *networktopology.TreeNode,
	nodeToScore map[string]int,
) (map[string]string, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Append PodCountMultiple constraint information if present

func (solver *networkTopologySolverImpl) calculateNodeOfferSlot(
	ctx context.Context,
	cycleStates map[string]fwktype.CycleState,
	toSchedulePods []*corev1.Pod,
	nodeInfos []fwktype.NodeInfo,
	addPod podFunc,
) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

// TODO consider pod assume on reservation

func calculateNodeExistingPodsNum(
	ctx context.Context,
	parallelizer parallelize.Parallelizer,
	selectorKey string,
	nodeInfos []fwktype.NodeInfo) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

func enumerateNodeTopologyNode(
	clusterNetworkTopology *networktopology.TreeNode,
	nodesNum int,
) map[string]*networktopology.TreeNode {
	_ = "STUB: not implemented"
	return nil
}

func evaluateTopologyNode(
	nodeToTopologyNodes map[string]*networktopology.TreeNode,
	nodeOfferSlot map[string]int,
	nodeToScore map[string]int,
	nodeExitingPodsNum map[string]int,
) {
	_ = "STUB: not implemented"
	return
}

// constrainOfferSlotByPodCountMultiple traverses the topology tree bottom-up
// and constrains each node's OfferSlot based on PodCountMultiple requirements.
// After this, OfferSlot at each node reflects the maximum achievable capacity
// considering PodCountMultiple constraints.
func constrainOfferSlotByPodCountMultiple(
	root *networktopology.TreeNode,
	layerPodCountMultiple map[schedulingv1alpha1.TopologyLayer]int,
) {
	_ = "STUB: not implemented"
	return
}

func doConstrainOfferSlot(
	node *networktopology.TreeNode,
	layerPodCountMultiple map[schedulingv1alpha1.TopologyLayer]int,
) {
	_ = "STUB: not implemented"
	return
}

func searchMustGatherSatisfiedNodes(
	jobNetworkRequirements *JobTopologyRequirements,
	clusterNetworkTopology *networktopology.TreeNode,
) []*networktopology.TreeNode {
	_ = "STUB: not implemented"
	return nil
}

func searchOfferSlotSatisfiedNodes(
	jobNetworkRequirements *JobTopologyRequirements,
	mustGatherSatisfiedNodes []*networktopology.TreeNode,
) []*networktopology.TreeNode {
	_ = "STUB: not implemented"
	return nil
}

var topologyNodeLessFunc = func(a, b *networktopology.TreeNode, lowerOfferSlot bool) bool {
	// Compare ExistingPodNum layer by layer from the current node
	for nodeA, nodeB := a, b; nodeA != nil && nodeB != nil; nodeA, nodeB = nodeA.Parent, nodeB.Parent {
		if nodeA.ExistingPodNum != nodeB.ExistingPodNum {
			return nodeA.ExistingPodNum > nodeB.ExistingPodNum
		}
	}
	// Compare OfferSlot layer by layer from the current node
	for nodeA, nodeB := a, b; nodeA != nil && nodeB != nil; nodeA, nodeB = nodeA.Parent, nodeB.Parent {
		if nodeA.OfferSlot != nodeB.OfferSlot {
			return (nodeA.OfferSlot < nodeB.OfferSlot) == lowerOfferSlot
		}
	}
	if a.Score != b.Score {
		return a.Score > b.Score
	}
	return a.Name < b.Name
}

func distributeOfferSlot(
	desiredOfferSlot int,
	topologyNode *networktopology.TreeNode,
	distribution map[string]int,
	layerPodCountMultiple map[schedulingv1alpha1.TopologyLayer]int,
) (topologyOrderedNodes []string, offerSlot int) {
	_ = "STUB: not implemented"
	// Calculate the maximum slot this topology node can provide
	return nil, 0
}

func distributePods(
	toSchedulePods []*corev1.Pod,
	topologyOrderedNodes []string,
	nodeToOfferSlot map[string]int,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func NewNetworkTopologySolver(handle fwktype.Handle) NetworkTopologySolver {
	_ = "STUB: not implemented"
	return *new(NetworkTopologySolver)
}
