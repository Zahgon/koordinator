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

package deviceshare

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"
	schedconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"

	schedulerconfig "github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

func (p *Plugin) PreScore(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Score(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *Plugin) ScoreExtensions() fwktype.ScoreExtensions {
	_ = "STUB: not implemented"
	return *new(fwktype.ScoreExtensions)
}

func (p *Plugin) NormalizeScore(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, scores fwktype.NodeScoreList) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) ScoreReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, reservationInfo *frameworkext.ReservationInfo, nodeName string) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *Plugin) ReservationScoreExtensions() frameworkext.ReservationScoreExtensions {
	_ = "STUB: not implemented"
	return *new(frameworkext.ReservationScoreExtensions)
}

func (p *Plugin) NormalizeReservationScore(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, scores frameworkext.ReservationScoreList) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// deviceResourceStrategyTypeMap maps strategy to scorer implementation
var deviceResourceStrategyTypeMap = map[schedulerconfig.ScoringStrategyType]scorer{
	schedulerconfig.LeastAllocated: func(args *schedulerconfig.DeviceShareArgs) *resourceAllocationScorer {
		resToWeightMap := resourcesToWeightMap(args.ScoringStrategy.Resources)
		return &resourceAllocationScorer{
			Name:                string(schedconfig.LeastAllocated),
			scorer:              leastResourceScorer(resToWeightMap),
			resourceToWeightMap: resToWeightMap,
		}
	},
	schedulerconfig.MostAllocated: func(args *schedulerconfig.DeviceShareArgs) *resourceAllocationScorer {
		resToWeightMap := resourcesToWeightMap(args.ScoringStrategy.Resources)
		return &resourceAllocationScorer{
			Name:                string(schedconfig.MostAllocated),
			scorer:              mostResourceScorer(resToWeightMap),
			resourceToWeightMap: resToWeightMap,
		}
	},
}

// resourceToWeightMap contains resource name and weight.
type resourceToWeightMap map[corev1.ResourceName]int64

// scorer is decorator for resourceAllocationScorer
type scorer func(args *schedulerconfig.DeviceShareArgs) *resourceAllocationScorer

// resourceAllocationScorer contains information to calculate resource allocation score.
type resourceAllocationScorer struct {
	Name                string
	scorer              func(requested, allocatable resourceToValueMap) int64
	resourceToWeightMap resourceToWeightMap
}

// resourceToValueMap is keyed with resource name and valued with quantity.
type resourceToValueMap map[corev1.ResourceName]int64

// scoreDevice will use `scorer` function to calculate the score per device.
func (r *resourceAllocationScorer) scoreDevice(podRequest corev1.ResourceList, total, free corev1.ResourceList) int64 {
	_ = "STUB: not implemented"
	return 0
}

func (r *resourceAllocationScorer) scoreNode(podRequest corev1.ResourceList, totalDeviceResources, freeDeviceResources deviceResources) int64 {
	_ = "STUB: not implemented"
	return 0
}

// resourcesToWeightMap make weightmap from resources spec
func resourcesToWeightMap(resourceSpecs []schedconfig.ResourceSpec) resourceToWeightMap {
	_ = "STUB: not implemented"
	return *new(resourceToWeightMap)
}

func leastResourceScorer(resToWeightMap resourceToWeightMap) func(resourceToValueMap, resourceToValueMap) int64 {
	_ = "STUB: not implemented"
	return nil
}

func leastRequestedScore(requested, capacity int64) int64 { _ = "STUB: not implemented"; return 0 }

func mostResourceScorer(resToWeightMap resourceToWeightMap) func(requested, allocable resourceToValueMap) int64 {
	_ = "STUB: not implemented"
	return nil
}

func mostRequestedScore(requested, capacity int64) int64 { _ = "STUB: not implemented"; return 0 }

// `requested` might be greater than `capacity` because pods with no
// requests get minimum values.
