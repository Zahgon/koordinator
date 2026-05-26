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

package elasticquota

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	policy "k8s.io/api/policy/v1"
	policylisters "k8s.io/client-go/listers/policy/v1"
	extenderv1 "k8s.io/kube-scheduler/extender/v1"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework/preemption"
)

func (g *Plugin) GetOffsetAndNumCandidates(nodes int32) (int32, int32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (g *Plugin) CandidatesToVictimsMap(candidates []preemption.Candidate) map[string]*extenderv1.Victims {
	_ = "STUB: not implemented"
	return nil
}

// PodEligibleToPreemptOthers determines whether this pod should be considered
// for preempting other pods or not. If this pod has already preempted other
// pods and those are in their graceful termination period, it shouldn't be
// considered for preemption.
// We look at the node that is nominated for this pod and as long as there are
// terminating pods on the node, we don't consider this for preempting more pods.
func (g *Plugin) PodEligibleToPreemptOthers(ctx context.Context, pod *corev1.Pod, nominatedNodeStatus *fwktype.Status) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// If the pod's nominated node is considered as UnschedulableAndUnresolvable by the filters,
// then the pod should be considered for preempting again.

// There is a terminating pod on the nominated node with the same quotaName.

// SelectVictimsOnNode finds minimum set of pods on the given node that should
// be preempted in order to make enough room for "pod" to be scheduled. The
// minimum set selected is subject to the constraint that a higher-priority pod
// is never preempted when a lower-priority pod could be (higher/lower relative
// to one another, not relative to the preemptor "pod").
// The algorithm first checks if the pod can be scheduled on the node when all the
// lower priority pods are gone. If so, it sorts all the lower priority pods by
// their priority and then puts them into two groups of those whose PodDisruptionBudget
// will be violated if preempted and other non-violating pods. Both groups are
// sorted by priority. It first tries to reprieve as many PDB violating pods as
// possible and then does them same for non-PDB-violating pods while checking
// that the "pod" can still fit on the node.
// NOTE: This function assumes that it is never called if "pod" cannot be scheduled
// due to pod affinity, node affinity, or node anti-affinity reasons. None of
// these predicates can be satisfied by removing more pods from the node.
func (g *Plugin) SelectVictimsOnNode(
	ctx context.Context,
	state fwktype.CycleState,
	pod *corev1.Pod,
	nodeInfo fwktype.NodeInfo,
	pdbs []*policy.PodDisruptionBudget,
) ([]*corev1.Pod, int, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// As the first step, remove all the lower priority pods from the node and
// check if the given pod can be scheduled.

// TODO only allow same quotaGroup preemption.

// No potential victims are found, and so we don't need to evaluate the node again since its state didn't change.

// If the new pod does not fit after removing all the lower priority pods,
// we are almost done and this node is not suitable for preemption. The only
// condition that we could check is if the "pod" is failing to schedule due to
// inter-pod affinity to one or more victims, but we have decided not to
// support this case for performance reasons. Having affinity to lower
// priority pods is not a recommended configuration anyway.

// Try to reprieve as many pods as possible. We first try to reprieve the PDB
// violating victims and then other non-violating ones. In both cases, we start
// from the highest priority victims.

// Now we try to reprieve non-violating victims.

// filterPodsWithPDBViolation groups the given "pods" into two groups of "violatingPods"
// and "nonViolatingPods" based on whether their PDBs will be violated if they are
// preempted.
// This function is stable and does not change the order of received pods. So, if it
// receives a sorted list, grouping will preserve the order of the input list.
func filterPodsWithPDBViolation(podInfos []fwktype.PodInfo, pdbs []*policy.PodDisruptionBudget) (violatingPodInfos, nonViolatingPodInfos []fwktype.PodInfo) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A pod with no labels will not match any PDB. So, no need to check.

// A PDB with a nil or empty selector matches nothing.

// Existing in DisruptedPods means it has been processed in API server,
// we don't treat it as a violating case.

// Only decrement the matched pdb when it's not in its <DisruptedPods>;
// otherwise we may over-decrement the budget number.

// We have found a matching PDB.

// TODO if the kubernetes version is before 1.20, will return nil.
func getPDBLister(handle fwktype.Handle) policylisters.PodDisruptionBudgetLister {
	_ = "STUB: not implemented"
	return *new(policylisters.PodDisruptionBudgetLister)
}

func (g *Plugin) OrderedScoreFuncs(ctx context.Context, nodesToVictims map[string]*extenderv1.Victims) []func(node string) int64 {
	_ = "STUB: not implemented"
	return nil
}

func (g *Plugin) canPreempt(pod, victim *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// The quota of pods that do not actively set through annotation will be recognized as the default quota.
// In some scenarios, many important pods are not scheduled by Koordinator, and their quotas are
// also recognized as default quota, this will cause the eviction operation to identify these important
// pods as victim pods, which is risky. When DisableDefaultQuotaPreemption is set to true, these pods can
// be avoided from being evicted.
