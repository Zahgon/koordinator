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

package nodenumaresource

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"
	schedconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/framework"

	schedulingconfig "github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
)

// resourceStrategyTypeMap maps strategy to scorer implementation
var resourceStrategyTypeMap = map[schedulingconfig.ScoringStrategyType]scorer{
	schedulingconfig.LeastAllocated: func(args *schedulingconfig.NodeNUMAResourceArgs) *resourceAllocationScorer {
		resToWeightMap := resourcesToWeightMap(args.ScoringStrategy.Resources)
		return &resourceAllocationScorer{
			Name:                string(schedconfig.LeastAllocated),
			scorer:              leastResourceScorer(resToWeightMap),
			resourceToWeightMap: resToWeightMap,
		}
	},
	schedulingconfig.MostAllocated: func(args *schedulingconfig.NodeNUMAResourceArgs) *resourceAllocationScorer {
		resToWeightMap := resourcesToWeightMap(args.ScoringStrategy.Resources)
		return &resourceAllocationScorer{
			Name:                string(schedconfig.MostAllocated),
			scorer:              mostResourceScorer(resToWeightMap),
			resourceToWeightMap: resToWeightMap,
		}
	},
}

func (p *Plugin) PreScore(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Score(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

// we have check in filter, so we will not get error in reserve

// Score-layer contract: the pod has already passed Filter on this node, so returning a
// non-Success status here would be upgraded to framework.Error by the scheduling framework
// and surface as SchedulerError on the Pod condition. Degrade to score=0 with nil status
// so the scheduler can continue selecting other candidates naturally, and record a warning
// for later investigation (it implies a mismatch between Filter and Score strictness).

// Same reasoning as above: degrade to score=0 instead of propagating Unschedulable/Error.

func (p *Plugin) scoreWithAmplifiedCPUs(state *preFilterState, nodeInfo fwktype.NodeInfo, resourceOptions *ResourceOptions) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *Plugin) calculateAllocatableAndRequested(
	nodeName string,
	nodeInfo fwktype.NodeInfo,
	podAllocation *PodAllocation,
	resourceOptions *ResourceOptions,
) (allocatable, requested *framework.Resource) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) ScoreExtensions() fwktype.ScoreExtensions {
	_ = "STUB: not implemented"

	// resourceToWeightMap contains resource name and weight.
	return *new(fwktype.ScoreExtensions)
}

type resourceToWeightMap map[corev1.ResourceName]int64

// scorer is decorator for resourceAllocationScorer
type scorer func(args *schedulingconfig.NodeNUMAResourceArgs) *resourceAllocationScorer

// resourceAllocationScorer contains information to calculate resource allocation score.
type resourceAllocationScorer struct {
	Name                string
	scorer              func(requested, allocatable resourceToValueMap) int64
	resourceToWeightMap resourceToWeightMap
}

// resourceToValueMap is keyed with resource name and valued with quantity.
type resourceToValueMap map[corev1.ResourceName]int64

// score will use `scorer` function to calculate the score.
func (r *resourceAllocationScorer) score(totalRequested, totalAllocatable, podRequests *framework.Resource) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Only fill the extended resource entry when it's non-zero.

func calculateResourceAllocatableRequest(allocatable, requested, podRequests *framework.Resource, resourceName corev1.ResourceName) (int64, int64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// If it's an extended resource, and the pod doesn't request it. We return (0, 0)
// as an implication to bypass scoring on this resource.

func getResourceQuantity(m *framework.Resource, resourceName corev1.ResourceName) int64 {
	_ = "STUB: not implemented"
	return 0
}

// resourcesToWeightMap make weightmap from resources spec
func resourcesToWeightMap(resources []schedconfig.ResourceSpec) resourceToWeightMap {
	_ = "STUB: not implemented"
	return *new(resourceToWeightMap)
}
