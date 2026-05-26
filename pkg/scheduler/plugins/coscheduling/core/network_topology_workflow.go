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

package core

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/networktopology"
)

const (
	ErrNoClusterNetworkTopology = "no cluster network topology"
	ErrorNoPlannedNodes         = "no planned pods"
	ErrorNoPendingPods          = "no pending pods"
	ErrorInvalidPlan            = "plan become invalid, wait for the next round"
)

// FIXME Currently, our workflow and solver only supports scenarios where
// 1. there are no Bound member Pods
// 2. the total number of Job member Pods is equal to the minimum number.
// 3. pods have no other topological requirements

func (pgMgr *PodGroupManager) PreFilter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) (*fwktype.PreFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this shouldn't happen. cause BeforePreFilter will return fwktype.UnschedulableAndUnresolvable if failedMessage != ""

// this shouldn't happen. If happens, return err to exposing problems

// first pod, the FindOneNode will take care of it

func (pgMgr *PodGroupManager) FindOneNode(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, preRes *fwktype.PreFilterResult) (string, *fwktype.Status) {
	_ = "STUB: not implemented"
	return "", nil
}

// this shouldn't happen. If happens, return err to exposing problems

// not first pod, the PreFilter will take care of it

// TODO: fill clusterNetworkTopology

func (ev *preemptionEvaluatorImpl) PlanNodes(
	ctx context.Context,
	networkTopologySpec *extension.NetworkTopologySpec,
	allPendingPods []*corev1.Pod,
	nodes []fwktype.NodeInfo,
	cycleStates map[string]fwktype.CycleState,
	addPod podFunc,
	preemptionCosts map[string]int,
) (podToNominatedNode map[string]string, successPods map[string]*Placements, statusMap map[string]*fwktype.Status, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// TODO: fill clusterNetworkTopology

const (
	preScoreStateKey = Name + "/pre-score-state"
)

type PreScoreState struct {
	nodesIndex map[string]int
}

func (p *PreScoreState) Clone() fwktype.StateData {
	_ = "STUB: not implemented"
	return *new(fwktype.StateData)
}

func (pgMgr *PodGroupManager) PreScore(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func (pgMgr *PodGroupManager) sortNodesByTopology(
	ctx context.Context,
	clusterNetworkTopology *networktopology.TreeSnapshot,
	podSelector string,
	candidateNodes []fwktype.NodeInfo,
	nodeInfos []fwktype.NodeInfo,
) []fwktype.NodeInfo {
	_ = "STUB: not implemented"
	// FIXME here we assert that every node only accommodates one pod
	return nil
}

// Compare ExistingPodNum layer by layer from the current node

// Compare OfferSlot layer by layer from the current node

func (pgMgr *PodGroupManager) Score(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodeInfo fwktype.NodeInfo) (int64, *fwktype.Status) {
	_ = "STUB: not implemented"
	return 0, nil
}
