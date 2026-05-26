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

package core

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/informers"
	listerv1 "k8s.io/client-go/listers/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	pgclientset "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/clientset/versioned"
	pgformers "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/informers/externalversions"
	pglister "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/listers/scheduling/v1alpha1"
	koordinatorinformers "github.com/koordinator-sh/koordinator/pkg/client/informers/externalversions"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/workloadauditor"
)

type Status string

const (
	Name     = "Coscheduling"
	stateKey = Name

	// PodGroupNotSpecified denotes no PodGroup is specified in the Pod spec.
	PodGroupNotSpecified Status = "PodGroup not specified"
	// PodGroupNotFound denotes the specified PodGroup in the Pod spec is
	// not found in API server.
	PodGroupNotFound Status = "PodGroup not found"
	Success          Status = "Success"
	Wait             Status = "Wait"
)

// Manager defines the interfaces for PodGroup management.
type Manager interface {
	NextPod() *corev1.Pod
	FindOneNode(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, result *fwktype.PreFilterResult) (string, *fwktype.Status)
	SucceedGangScheduling()
	PreEnqueue(context.Context, *corev1.Pod) (err error)
	BeforePreFilter(context.Context, fwktype.CycleState, *corev1.Pod) (err error)
	PreFilter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) (*fwktype.PreFilterResult, *fwktype.Status)
	PreScore(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod, nodes []fwktype.NodeInfo) *fwktype.Status
	Score(ctx context.Context, state fwktype.CycleState, p *corev1.Pod, nodeInfo fwktype.NodeInfo) (int64, *fwktype.Status)
	Permit(context.Context, *corev1.Pod) (time.Duration, Status)
	PostBind(context.Context, *corev1.Pod, string)
	PostFilter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, m fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status)
	AfterPostFilter(context.Context, fwktype.CycleState, *corev1.Pod, fwktype.Handle, string, fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status)
	GetAllPodsFromGang(string) []*corev1.Pod
	AllowGangGroup(*corev1.Pod, fwktype.Handle, string)
	Unreserve(context.Context, fwktype.CycleState, *corev1.Pod, string, fwktype.Handle, string)

	GetGangSummary(gangId string) (*GangSummary, bool)
	GetGangSummaries() map[string]*GangSummary

	GetBoundPodNumber(gangId string) int32

	GetGangBindingInfo(pod *corev1.Pod) *GangBindingInfo
}

// PodGroupManager defines the scheduling operation called
type PodGroupManager struct {
	handle fwktype.Handle
	args   *config.CoschedulingArgs
	// pgClient is a podGroup client
	pgClient pgclientset.Interface
	// pgLister is podgroup lister
	pgLister pglister.PodGroupLister
	// podLister is pod lister
	podLister listerv1.PodLister
	// cache stores gang info
	cache                 *GangCache
	holder                GangSchedulingContextHolder
	preemptionEvaluator   PreemptionEvaluator
	networkTopologySolver NetworkTopologySolver

	workloadAuditor workloadauditor.WorkloadAuditor
}

// NewPodGroupManager creates a new operation object.
func NewPodGroupManager(
	handle fwktype.Handle,
	args *config.CoschedulingArgs,
	pgClient pgclientset.Interface,
	pgSharedInformerFactory pgformers.SharedInformerFactory,
	sharedInformerFactory informers.SharedInformerFactory,
	koordSharedInformerFactory koordinatorinformers.SharedInformerFactory,
) *PodGroupManager {
	_ = "STUB: not implemented"
	return nil
}

func (pgMgr *PodGroupManager) NextPod() *corev1.Pod { _ = "STUB: not implemented"; return nil }

// the podGroup is deleted

// iterate over each gangGroup, get all the pods

// correct podInfo.Time and podInfo.Attempts

func (pgMgr *PodGroupManager) SucceedGangScheduling() { _ = "STUB: not implemented"; return }

// PreEnqueue
// TODO Turning it on may result in no Pod scheduling events, and an external check should be done through the controller later.
func (pgMgr *PodGroupManager) PreEnqueue(ctx context.Context, pod *corev1.Pod) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// check if gang is initialized

// resourceSatisfied means pod will directly pass the PreFilter

// use IsPodRepresentative to avoid write lock collisions.

// It's possible that another Pod arrives and satisfies the Gang Basic Check.
// Here, we determine this by checking whether RepresentativePod is itself within a write lock.

// Subsequent Pods should be prevented from entering ActiveQ or BackoffQ to avoid the time-consuming deletion of them

// PreEnqueue
// i.Check whether children in Gang has met the requirements of minimum number under each Gang, and reject the pod if negative.
// ii.Check whether the Gang is inited, and reject the pod if positive.
// iii.Check whether the Gang is OnceResourceSatisfied
func (pgMgr *PodGroupManager) basicGangRequirementsCheck(gang *Gang, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (pgMgr *PodGroupManager) BeforePreFilter(ctx context.Context, cycleState fwktype.CycleState, pod *corev1.Pod) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// check if gang is initialized

// resourceSatisfied means pod will directly pass the PreFilter

// use IsPodRepresentative to avoid write lock collisions.

// It's possible that another Pod arrives and satisfies the Gang Basic Check.
// Here, we determine this by checking whether RepresentativePod is itself within a write lock.

// clear the current representative because representative is already enter into scheduling

// PostFilter invoked at the postFilter extension point.
func (pgMgr *PodGroupManager) PostFilter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, m fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AfterPostFilter
// i. If strict-mode, we will set scheduleCycleValid to false and release all assumed pods.
// ii. If non-strict mode, we will do nothing.
func (pgMgr *PodGroupManager) AfterPostFilter(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, handle fwktype.Handle, pluginName string, filteredNodeStatusMap fwktype.NodeToStatusReader) (*fwktype.PostFilterResult, *fwktype.Status) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pgMgr *PodGroupManager) summaryAndRecordFailedMessage(state fwktype.CycleState, triggerPod *corev1.Pod, filteredNodeStatusMap fwktype.NodeToStatusReader) string {
	_ = "STUB: not implemented"
	return ""
}

// Permit
// we will calculate all Gangs in GangGroup whether the current number of assumed-pods in each Gang meets the Gang's minimum requirement.
// and decide whether we should let the pod wait in Permit stage or let the whole gangGroup go binding
func (pgMgr *PodGroupManager) Permit(ctx context.Context, pod *corev1.Pod) (time.Duration, Status) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(Status)
}

// first add pod to the gang's WaitingPodsMap

// check each gang group

// Unreserve
// if gang is resourceSatisfied, we only delAssumedPod
// if gang is not resourceSatisfied and is in StrictMode, we release all the assumed pods
func (pgMgr *PodGroupManager) Unreserve(ctx context.Context, state fwktype.CycleState, pod *corev1.Pod, nodeName string, handle fwktype.Handle, pluginName string) {
	_ = "STUB: not implemented"
	return
}

// first delete the pod from gang's waitingFroBindChildren map

// TODO we should record failed message when current pod is the first failed pod of gang, now we just let it go, so quick fail is not supported

func (pgMgr *PodGroupManager) rejectGangGroupById(handle fwktype.Handle, pluginName, gangId, message string) {
	_ = "STUB: not implemented"
	return
}

// iterate over each gangGroup, get all the pods

func (pgMgr *PodGroupManager) rejectGangGroup(handle fwktype.Handle, gangSet sets.Set[string], message string) {
	_ = "STUB: not implemented"
	return
}

// PostBind updates a PodGroup's status.
func (pgMgr *PodGroupManager) PostBind(ctx context.Context, pod *corev1.Pod, nodeName string) {
	_ = "STUB: not implemented"
	return
}

// update gang in cache

func (pgMgr *PodGroupManager) AllowGangGroup(pod *corev1.Pod, handle fwktype.Handle, pluginName string) {
	_ = "STUB: not implemented"
	return
}

// record binding members if it is unset

func (pgMgr *PodGroupManager) GetGangByPod(pod *corev1.Pod) *Gang {
	_ = "STUB: not implemented"
	return nil
}

func (pgMgr *PodGroupManager) GetAllPodsFromGang(gangId string) []*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (pgMgr *PodGroupManager) GetGangSummary(gangId string) (*GangSummary, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pgMgr *PodGroupManager) GetGangSummaries() map[string]*GangSummary {
	_ = "STUB: not implemented"
	return nil
}

func (pgMgr *PodGroupManager) GetBoundPodNumber(gangId string) int32 {
	_ = "STUB: not implemented"
	return 0
}

type GangBindingInfo struct {
	GangGroupId string
	MemberCount int
}

func (pgMgr *PodGroupManager) GetGangBindingInfo(pod *corev1.Pod) *GangBindingInfo {
	_ = "STUB: not implemented"
	return nil
}

// Get the snapshot members from GangGroupInfo.
// This value was set in AllowGangGroup and persists through binding cycle.

// skip gang binding info when no member pods or single pod

// only record binding info for member pods
