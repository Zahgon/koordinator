/*
Copyright 2022 The Koordinator Authors.
Copyright 2020 The Kubernetes Authors.

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
	policy "k8s.io/api/policy/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	policylisters "k8s.io/client-go/listers/policy/v1"
	extenderv1 "k8s.io/kube-scheduler/extender/v1"
	fwktype "k8s.io/kube-scheduler/framework"
	schedulerconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/kubernetes/pkg/scheduler/framework/plugins/defaultpreemption"
	"k8s.io/kubernetes/pkg/scheduler/framework/preemption"

	listerschedulingv1alpha1 "github.com/koordinator-sh/koordinator/pkg/client/listers/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

var (
	_ fwktype.PostFilterPlugin = &PreemptionMgr{}
	_ preemption.Interface     = &PreemptionMgr{}
)

// PreemptionMgr is a wrapper of defaultpreemption.DefaultPreemption, supporting the following preemption behaviors:
// 1. a pod preempt a pod.
// 2. a reservation preempt a pod.
// TODO: support the reservation being preempted
type PreemptionMgr struct {
	*defaultpreemption.DefaultPreemption
	fh                frameworkext.ExtendedHandle
	podLister         corelisters.PodLister
	pdbLister         policylisters.PodDisruptionBudgetLister
	reservationLister listerschedulingv1alpha1.ReservationLister
}

func newPreemptionMgr(pluginArgs *config.ReservationArgs, extendedHandle frameworkext.ExtendedHandle,
	podLister corelisters.PodLister, rLister listerschedulingv1alpha1.ReservationLister) (*PreemptionMgr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pm *PreemptionMgr) Name() string { _ = "STUB: not implemented"; return "" }

func (pm *PreemptionMgr) PostFilter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, m fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SelectVictimsOnNode finds minimum set of pods on the given node that should be preempted in order to make enough room
// for "pod" to be scheduled.
// Note that both `state` and `nodeInfo` are deep-copied.
// We delegate the function to extend the preemption rules:
// If a pod is marked as non-preemptible, it will not be selected as the victim.
func (pm *PreemptionMgr) SelectVictimsOnNode(
	ctx context.Context,
	state fwktype.CycleState,
	pod *corev1.Pod,
	nodeInfo fwktype.NodeInfo,
	pdbs []*policy.PodDisruptionBudget) ([]*corev1.Pod, int, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// As the first step, remove all the lower priority pods from the node and
// check if the given pod can be scheduled.

// NOTE: Ignore the non-preemptible pod.

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

var _ corelisters.PodLister = &delegatingPodLister{}

// delegatingPodLister delegates the PodLister interface to get a Pod or a Reservation from the reserve pod.
type delegatingPodLister struct {
	corelisters.PodLister
	reservationLister listerschedulingv1alpha1.ReservationLister
	cachedPod         *corev1.Pod
}

func newDelegatingPodLister(podLister corelisters.PodLister,
	reservationLister listerschedulingv1alpha1.ReservationLister,
	pod *corev1.Pod) corelisters.PodLister {
	_ = "STUB: not implemented"
	return *new(corelisters.PodLister)
}

func (dp *delegatingPodLister) Pods(namespace string) corelisters.PodNamespaceLister {
	_ = "STUB: not implemented"
	// only delegate the default namespace since reserve pod is forced to the namespace
	return *new(corelisters.PodNamespaceLister)
}

var _ corelisters.PodNamespaceLister = &delegatingPodNamespaceLister{}

type delegatingPodNamespaceLister struct {
	corelisters.PodNamespaceLister
	reservationLister listerschedulingv1alpha1.ReservationLister
	cachedPod         *corev1.Pod
}

func (dpn *delegatingPodNamespaceLister) Get(name string) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the cached pod not found from the informer, try to get a corresponding reservation

// Is it necessary to regenerate the reserve pod?

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

// This object has an invalid selector, it does not match the pod

// A PDB with a nil or empty selector matches nothing.

// Existing in DisruptedPods means it has been processed in API server,
// we don't treat it as a violating case.

// Only decrement the matched pdb when it's not in its <DisruptedPods>;
// otherwise we may over-decrement the budget number.

// We have found a matching PDB.

// OrderedScoreFuncs returns nil as this plugin does not implement ordered score functions.
func (pm *PreemptionMgr) OrderedScoreFuncs(ctx context.Context, nodesToVictims map[string]*extenderv1.Victims) []func(node string) int64 {
	_ = "STUB: not implemented"
	return nil
}

func getPreemptionArgs(pluginArgs *config.ReservationArgs) (*schedulerconfig.DefaultPreemptionArgs, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
