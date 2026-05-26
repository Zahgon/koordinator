package core

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	clientset "k8s.io/client-go/kubernetes"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

const (
	ReasonAlreadyPreempted                 = "preemption already attempted by %s with message %s"
	ReasonNoPendingPods                    = "no pending pods"
	ReasonPreemptionPolicyNever            = "not eligible due to preemptionPolicy=Never."
	ReasonTerminatingVictimOnNominatedNode = "not eligible due to terminating pod on the nominated node."
	ReasonListNode                         = "list nodes from snapshot err"
	ReasonNoNodesAvailable                 = "no nodes available"
	ReasonNoPotentialVictims               = "no potential victims"
	ReasonPreemptionNotHelpful             = "preemption not helpful"
	ReasonSelectVictimsOnNodeError         = "select victims on node error"
	ReasonPrepareCandidatesError           = "prepare candidates error"
	ReasonTriggerPodPreemptSuccess         = "preempt success, alreadyWaitingForBound: %d/%d"
)

const (
	OperationRemovePossibleVictims = "RemovePossibleVictims"
	OperationFindFeasibleNodes     = "FindFeasibleNodes"
	OperationSelectVictimsOnNode   = "SelectVictimsOnNode"
	OperationSetNominatedNode      = "SetNominatedNode"
	OperationPreemptPod            = "PreemptPod"
	OperationClearNominatedNode    = "ClearNominatedNode"
)

// IsEligiblePodFunc is a function which may be assigned to the DefaultPreemption plugin.
// This may implement rules/filtering around preemption eligibility, which is in addition to
// the internal requirement that the victim pod have lower priority than the preemptor pod.
// Any customizations should always allow system services to preempt normal pods, to avoid
// problems if system pods are unable to find space.
type IsEligiblePodFunc func(nodeInfo fwktype.NodeInfo, victim fwktype.PodInfo, preemptor *corev1.Pod) bool

type PreemptionEvaluator interface {
	Preempt(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, m fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status)
}

type preemptionEvaluatorImpl struct {
	// IsEligiblePod returns whether a victim pod is allowed to be preempted by a preemptor pod.
	// This filtering is in addition to the internal requirement that the victim pod have lower
	// priority than the preemptor pod. Any customizations should always allow system services
	// to preempt normal pods, to avoid problems if system pods are unable to find space.
	IsEligiblePod IsEligiblePodFunc

	handle                frameworkext.ExtendedHandle
	gangCache             *GangCache
	gangContextHolder     *GangSchedulingContextHolder
	networkTopologySolver NetworkTopologySolver
}

func NewPreemptionEvaluator(handle fwktype.Handle, gangCache *GangCache, gangContextHolder *GangSchedulingContextHolder, networkTopologySolver NetworkTopologySolver) PreemptionEvaluator {
	_ = "STUB: not implemented"
	return *new(PreemptionEvaluator)
}

type JobPreemptionStateContextKey struct {
}

type JobPreemptionState struct {
	TriggerPodKey string `json:"TriggerPodKey,omitempty"`
	PreemptorKey  string `json:"preemptorKey,omitempty"`

	gangSchedulingContext *GangSchedulingContext

	allPendingPods []*corev1.Pod
	allWaitingPods []*corev1.Pod
	allPods        []*corev1.Pod

	Reason                          string            `json:"reason,omitempty"`
	Message                         string            `json:"message,omitempty"`
	TerminatingPodOnNominatedNode   map[string]string `json:"terminatingPodOnNominatedNode,omitempty"`
	DurationOfNodeInfoClone         metav1.Duration   `json:"durationOfNodeInfoClone,omitempty"`
	DurationOfCycleStateClone       metav1.Duration   `json:"durationOfCycleStateClone,omitempty"`
	possibleVictims                 map[string][]fwktype.PodInfo
	DurationOfRemovePossibleVictims metav1.Duration   `json:"durationOfRemovePossibleVictims,omitempty"`
	PodToNominatedNode              map[string]string `json:"podToNominatedNode,omitempty"`
	DurationOfPlaceToSchedulePods   metav1.Duration   `json:"durationOfPlaceToSchedulePods,omitempty"`
	statusMap                       map[string]*fwktype.Status
	unschedulablePods               []*corev1.Pod
	selectVictimError               error
	DurationOfSelectVictimsOnNode   metav1.Duration `json:"durationOfSelectVictimsOnNode,omitempty"`
	victims                         map[string][]*corev1.Pod
	DurationOfPrepareCandidates     metav1.Duration   `json:"durationOfPrepareCandidates,omitempty"`
	ClearNominatedNodeFailedMsg     map[string]string `json:"clearNominatedNodeFailedMsg,omitempty"`
	DurationOfMakeNomination        metav1.Duration   `json:"durationOfMakeNomination,omitempty"`
	DurationOfCancelNomination      metav1.Duration   `json:"durationOfCancelNomination,omitempty"`

	SchedulingMode  frameworkext.SchedulingMode
	NodeToOfferSlot map[string]int

	PossibleVictims         []v1alpha1.NodePossibleVictim `json:"possibleVictims,omitempty"`
	UnschedulablePodsNumber int                           `json:"unschedulablePodsNumber,omitempty"`
	SelectVictimError       string                        `json:"selectVictimError,omitempty"`
	Victims                 []v1alpha1.NodePossibleVictim `json:"victims,omitempty"`
}

func (s *JobPreemptionState) addMoreDetailForStateToMarshal() { _ = "STUB: not implemented"; return }

func preemptionStateFromContext(ctx context.Context) *JobPreemptionState {
	_ = "STUB: not implemented"
	return nil
}

func contextWithJobPreemptionState(ctx context.Context, preemptionState *JobPreemptionState) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Preempt returns a PostFilterResult carrying suggested nominatedNodeName, along with a Status.
// The semantics of returned <PostFilterResult, Status> varies on different scenarios:
//
//   - <nil, Error>. This denotes it's a transient/rare error that may be self-healed in future cycles.
//
//   - <nil, Unschedulable>. This status is mostly as expected like the preemptor is waiting for the
//     victims to be fully terminated.
//
//   - In both cases above, a nil PostFilterResult is returned to keep the pod's nominatedNodeName unchanged.
//
//   - <non-nil PostFilterResult, Unschedulable>. It indicates the pod cannot be scheduled even with preemption.
//     In this case, a non-nil PostFilterResult is returned and result.NominatingMode instructs how to deal with
//     the nominatedNodeName.
//
//   - <non-nil PostFilterResult, Success>. It's the regular happy path
//     and the non-empty nominatedNodeName will be applied to the preemptor pod.

func (ev *preemptionEvaluatorImpl) Preempt(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, m fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// preempt implements the core preemption logic to help a high-priority pod (preemptor)
// find space by evicting lower-priority pods (victims) when no node can schedule it.
//
// This method performs the following steps:
// 1. Checks eligibility for preemption, considering PreemptionPolicy and ongoing terminations.
// 2. For gang-scheduled workloads, retrieves all associated pending/waiting pods in the same job.
// 3. Filters nodes where preemption may help (nodes with Unschedulable status).
// 4. Simulates removal of potential victims on candidate nodes ("dry-run").
// 5. Evaluates if scheduling becomes feasible after victim removal.
// 6. Selects optimal victims per node based on priority, job grouping, and cost estimation.
// 7. Triggers eviction of selected victims and nominates nodes for every member of the preemptor.
//
// The behavior differs slightly between regular and gang scheduling modes:
// - In regular mode: only the given pod is considered.
// - In gang mode: all pending pods in the same gang are evaluated together to maintain co-scheduling guarantees.
//
// Side effects include:
// - Updating NominatedNodeName for the preemptor and related pods.
// - Clearing nominations for lower-priority pods that may no longer fit.
// - Sending reject signals to waiting pods via Permit plugins.
func (ev *preemptionEvaluatorImpl) preempt(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, m fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use UnschedulableAndUnresolvable as absentNodesStatus so that NodesForStatusCode
// correctly handles nodes not in the explicit map without panicking on nil.

// Return a FitError only when there are no candidates that fit the pod.

// jobEligibleToPreemptOthers returns one bool and one string. The bool
// indicates whether this pod should be considered for preempting other pods or
// not. The string includes the reason if this pod isn't eligible.
// There are several reasons:
//  1. The pod has a preemptionPolicy of Never.
//  2. The pod has already preempted other pods and the victims are in their graceful termination period.
//     Currently, we check the node that is nominated for this pod, and as long as there are
//     terminating pods on this node, we don't attempt to preempt more pods.
func (ev *preemptionEvaluatorImpl) jobEligibleToPreemptOthers(ctx context.Context, triggerPod *corev1.Pod, allPendingPods []*corev1.Pod, m fwktype.NodeToStatusReader) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// TODO all pods of the same job must have the same preemptionPolicy, add a webhook for it.

func (ev *preemptionEvaluatorImpl) podEligibleToPreemptOthers(ctx context.Context, pod *corev1.Pod, m fwktype.NodeToStatusReader) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// If the pod's nominated node is considered as UnschedulableAndUnresolvable by the filters,
// then the pod should be considered for preempting again.

// There is a terminating pod on the nominated node.

// isPreemptionAllowed returns whether the victim residing on nodeInfo can be preempted by the preemptor
func (ev *preemptionEvaluatorImpl) isPreemptionAllowed(nodeInfo fwktype.NodeInfo, victim fwktype.PodInfo, preemptor *corev1.Pod) bool {
	_ = "STUB: not implemented"
	// The victim must have lower priority than the preemptor, in addition to any filtering implemented by IsEligiblePod
	return false
}

// podTerminatingByPreemption returns true if the pod is in the termination state caused by scheduler preemption.
func podTerminatingByPreemption(p *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

// FindCandidates calculates a slice of preemption candidates.
// Each candidate is executable to make the given <pod> schedulable.
func (ev *preemptionEvaluatorImpl) findCandidates(
	ctx context.Context,
	state fwktype.CycleState,
	allNodes []fwktype.NodeInfo,
	pod *corev1.Pod,
	m fwktype.NodeToStatusReader,
) (map[string]string, map[string][]*corev1.Pod, map[string]*fwktype.Status, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// nodesWherePreemptionMightHelp returns a list of nodes with failed predicates
// that may be satisfied by removing pods from the node.
func nodesWherePreemptionMightHelp(nodes []fwktype.NodeInfo, m fwktype.NodeToStatusReader) ([]fwktype.NodeInfo, map[string]*fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We only attempt preemption on nodes with status 'Unschedulable'. For
// diagnostic purposes, we propagate UnschedulableAndUnresolvable if either
// implied by absence in map or explicitly set.

// clone nodeInfo to avoid modifying the nodeInfoSnapshot

type podFunc = func(state fwktype.CycleState, pod *corev1.Pod, podInfo fwktype.PodInfo, nodeInfo fwktype.NodeInfo) error

// TODO consider PDB violation

// dryRunPreemption simulates the preemption process on a set of potential nodes in parallel.
// It performs a "what-if" analysis by removing potential victim pods from node resources
// and checking if the preemptor pod (and its associated gang pods) can then be scheduled.
//
// This method executes the following steps:
//  1. Removes all possible victims from each candidate node's NodeInfo.
//  2. Estimates a preemption cost per node based on victim priorities and job grouping.
//  3. Attempts to schedule pending pods (e.g., gang members) on the modified nodes,
//     considering assumed pods for sequential scheduling simulation.
//  4. Identifies feasible nodes where all required pods can fit after preemption.
//  5. Selects the best victims per feasible node by re-adding pods one-by-one and testing feasibility.
//
// The simulation uses cloned CycleState and NodeInfo objects to avoid affecting real scheduling state.
func (ev *preemptionEvaluatorImpl) dryRunPreemption(
	ctx context.Context,
	triggerPod *corev1.Pod,
	cycleStates map[string]fwktype.CycleState,
	potentialNodes []fwktype.NodeInfo,
) (map[string]string, map[string][]*corev1.Pod, map[string]*fwktype.Status, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (ev *preemptionEvaluatorImpl) removePossibleVictims(
	triggerPod *corev1.Pod,
	cycleStates map[string]fwktype.CycleState,
	potentialNodes []fwktype.NodeInfo,
	removePod podFunc,
) (map[string][]fwktype.PodInfo, map[string]*fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func estimatePreemptionCost(possibleVictims map[string][]fwktype.PodInfo) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

// estimate preemption cost in the job dimension

type Placements struct {
	nodeName string
	nodeInfo fwktype.NodeInfo
	pods     []*corev1.Pod
}

func (ev *preemptionEvaluatorImpl) placeToSchedulePods(
	ctx context.Context,
	toSchedulePods []*corev1.Pod,
	cycleStates map[string]fwktype.CycleState,
	potentialNodes []fwktype.NodeInfo,
	preemptionCosts map[string]int,
	addPod podFunc,
	statusMap map[string]*fwktype.Status,
) (podToNominatedNode map[string]string, successPods map[string]*Placements, unschedulablePods []*corev1.Pod) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// The lifecycle of assumedCycleState is one job schedule.
// During job preemption, there are two job schedules:
// 1. one to determine whether preemption is effective after removing all victims, and
// 2. another to determine the victim on the candidate node.
// Here we choose clone to avoid the modification of cycleState affecting the subsequent determination of Victim

// TODO consider pod assume on reservation

func (ev *preemptionEvaluatorImpl) findFeasibleNodes(
	ctx context.Context,
	toSchedulePod *corev1.Pod,
	cycleStates map[string]fwktype.CycleState,
	potentialNodes []fwktype.NodeInfo,
	assumedCycleStates map[string]fwktype.CycleState,
	assumedNodeInfos map[string]fwktype.NodeInfo,
	statusMap map[string]*fwktype.Status,
) (feasibleNodes []fwktype.NodeInfo) {
	_ = "STUB: not implemented"
	return nil
}

func (ev *preemptionEvaluatorImpl) selectVictims(
	ctx context.Context,
	possibleVictims map[string][]fwktype.PodInfo,
	cycleStates map[string]fwktype.CycleState,
	successPods map[string]*Placements,
	addPod podFunc,
	removePod podFunc,
) (victims map[string][]*corev1.Pod, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sortVictims(victims []fwktype.PodInfo) { _ = "STUB: not implemented"; return }

func (ev *preemptionEvaluatorImpl) prepareCandidates(ctx context.Context, candidatesByNode map[string][]*corev1.Pod, triggerPod *corev1.Pod) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

// If the victim Pod is already being deleted, we don't have to make another deletion api call.

// Lower priority pods nominated to run on this node, may no longer fit on
// this node. So, we should remove their nomination. Removing their
// nomination updates these pods and moves them to the active queue. It
// lets scheduler find another place for them.

// We do not return as this error is not critical.

func (ev *preemptionEvaluatorImpl) preemptPod(ctx context.Context, preemptor, victim *corev1.Pod, pluginName string) error {
	_ = "STUB: not implemented"
	return nil
}

// If the victim is a WaitingPod, send a reject message to the PermitPlugin.
// Otherwise, we should delete the victim.

// getLowerPriorityNominatedPods returns pods whose priority is smaller than the
// priority of the given "pod" and are nominated to run on the given node.
// Note: We could possibly check if the nominated lower priority pods still fit
// and return those that no longer fit, but that would require lots of
// manipulation of NodeInfo and PreFilter state per nominated pod. It may not be
// worth the complexity, especially because we generally expect to have a very
// small number of nominated pods per node.
func (ev *preemptionEvaluatorImpl) getLowerPriorityNominatedPods(triggerPod *corev1.Pod, candidatesByNode map[string][]*corev1.Pod) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// clearNominatedNodeName internally submit a patch request to API server
// to set each pods[*].Status.NominatedNodeName> to "".
func (ev *preemptionEvaluatorImpl) clearNominatedNodeName(ctx context.Context, cs clientset.Interface, pods ...*corev1.Pod) utilerrors.Aggregate {
	_ = "STUB: not implemented"
	return *new(utilerrors.Aggregate)
}

func (ev *preemptionEvaluatorImpl) makeNomination(ctx context.Context, podToNominatedNode map[string]string) {
	_ = "STUB: not implemented"
	return
}

func makeWaitingPodToNominatedNode(allWaitingPods []*corev1.Pod) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (ev *preemptionEvaluatorImpl) cancelNomination(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (ev *preemptionEvaluatorImpl) rejectAllWaitingPod(_ context.Context, allWaitingPods []*corev1.Pod, pluginName, msg string) {
	_ = "STUB: not implemented"
	return
}

func (ev *preemptionEvaluatorImpl) setAllNominatedNode(ctx context.Context, allPods []*corev1.Pod, podNominatedNodes, waitingPodToNominatedNode map[string]string) {
	_ = "STUB: not implemented"
	return
}

// TODO nominated reservationRelated
