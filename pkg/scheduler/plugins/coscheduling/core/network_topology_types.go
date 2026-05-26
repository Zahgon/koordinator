package core

import (
	"context"

	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/extension"
	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/networktopology"
)

type TopologyState struct {
	JobTopologyRequirements  *JobTopologyRequirements
	NodeOfferSlot            map[string]int
	NodeToStatusMap          map[string]*fwktype.Status
	MustGatheredTopologyNode []*networktopology.TreeNode
}

type ContextKey struct {
}

func TopologyStateFromContext(ctx context.Context) *TopologyState {
	_ = "STUB: not implemented"
	return nil
}

func ContextWithTopologyState(ctx context.Context, topologyState *TopologyState) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type JobTopologyRequirements struct {
	TopologyLayerMustGather schedulingv1alpha1.TopologyLayer
	DesiredOfferSlot        int
	// LayerPodCountMultiple specifies the pod count multiple constraint for each topology layer.
	// The number of Pods placed in a topology node of the specified layer must be
	// a multiple of the corresponding value.
	LayerPodCountMultiple map[schedulingv1alpha1.TopologyLayer]int
}

func GetMustGatherLayer(spec *extension.NetworkTopologySpec, isLayerAncestorFunc networktopology.IsLayerAncestorFunc) schedulingv1alpha1.TopologyLayer {
	_ = "STUB: not implemented"
	return *new(schedulingv1alpha1.TopologyLayer)
}

func GetLayerPodCountMultiple(spec *extension.NetworkTopologySpec) map[schedulingv1alpha1.TopologyLayer]int {
	_ = "STUB: not implemented"
	return nil
}
