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

package reservation

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

const (
	mostPreferredScore = 1000
)

func (pl *Plugin) PreScore(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfos []fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) preScoreForNormalPod(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodes []*corev1.Node) *fwktype.Status {
	_ = "STUB: not implemented"
	// if the pod is reservation-ignored, it does not want a nominated reservation
	return nil
}

func (pl *Plugin) preScoreForPreAllocation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodes []*corev1.Node) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// For multiple pre-allocatable pods mode, skip the nomination in PreScore phase.
// The selectedPreAllocatablePods are already determined during Filter phase,
// and Score phase will use them directly without needing a single nominated pod.

func (pl *Plugin) Score(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pl *Plugin) scoreForNormalPod(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pl *Plugin) scoreForPreAllocation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodeName string) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If there is no pre-allocatable pod nominated on this node, give it the highest priority.

// The more pre-allocatable pods are selected on this node, the lower priority it will get.

func (pl *Plugin) ScoreExtensions() fwktype.ScoreExtensions {
	_ = "STUB: not implemented"
	return *new(fwktype.ScoreExtensions)
}

func (pl *Plugin) NormalizeScore(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, scores fwktype.NodeScoreList) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pl *Plugin) ScoreReservation(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, rInfo *frameworkext.ReservationInfo, nodeName string) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Here we use MostAllocated (simply set all weights as 1.0)

func (pl *Plugin) ReservationScoreExtensions() frameworkext.ReservationScoreExtensions {
	_ = "STUB: not implemented"
	return *new(frameworkext.ReservationScoreExtensions)
}

func findMostPreferredReservationByOrder(rOnNode []*frameworkext.ReservationInfo) (*frameworkext.ReservationInfo, int64) {
	_ = "STUB: not implemented"
	return nil, 0
}

// The smaller the order value is, the reservation will be selected first
