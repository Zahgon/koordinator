package core

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	ReasonPodDeleted = "PodDeleted"
	ReasonPodBound   = "PodBound"

	ReasonGangGroupEnterIntoScheduling = "GangGroupEnterIntoScheduling"
	ReasonGangBasicCheckUnsatisfied    = "GangBasicCheckUnsatisfied"
)

type GangGroupInfo struct {
	lock sync.RWMutex

	Initialized bool
	GangGroupId string
	GangGroup   []string

	// OnceResourceSatisfied indicates whether the gang has ever reached the ResourceSatisfied state，which means the
	// children number has reached the minNum in the early step,
	// once this variable is set true, it is irreversible.
	OnceResourceSatisfied bool

	/*
		WaitingGangIDs
		- is recorded when permit return wait
		- is cleared when
		  - gang is reject on postFilter
		  - all pods are deleted
		  - all gangs are deleted
		  - all pods are unreserved
		  - all pod are bound
	*/

	WaitingGangIDs sets.Set[string]

	/*
		RepresentativePodKey
		- is recorded when the first pod of gangGroup pass the PreEnqueue
		- is deleted when
		  - the pod is bound
		  - the pod is deleted
		- is replaced when some memberPod of gangGroup is firstly failed during some gangGroup schedulingContext
	*/
	RepresentativePodKey string

	// BindingMemberPods is the waiting pods when gang enters binding phase.
	// This value is locked when AllowGangGroup is called and used by all pods during PreBind.
	// It will be cleared when all gangs in the group complete binding.
	BindingMemberPods sets.Set[string]
}

func NewGangGroupInfo(gangGroupId string, gangGroup []string) *GangGroupInfo {
	_ = "STUB: not implemented"
	return nil
}

func (gg *GangGroupInfo) SetInitialized() { _ = "STUB: not implemented"; return }

func (gg *GangGroupInfo) IsInitialized() bool { _ = "STUB: not implemented"; return false }

func (gg *GangGroupInfo) isGangOnceResourceSatisfied() bool {
	_ = "STUB: not implemented"
	return false
}

func (gg *GangGroupInfo) setResourceSatisfied() { _ = "STUB: not implemented"; return }

func (gg *GangGroupInfo) AddWaitingGang() { _ = "STUB: not implemented"; return }

func (gg *GangGroupInfo) RemoveWaitingGang(gangID string) { _ = "STUB: not implemented"; return }

func (gg *GangGroupInfo) ClearWaitingGang() { _ = "STUB: not implemented"; return }

func (gg *GangGroupInfo) RecordIfNoRepresentatives(pod *corev1.Pod) string {
	_ = "STUB: not implemented"
	return ""
}

func (gg *GangGroupInfo) DeleteIfRepresentative(pod *corev1.Pod, reason string) {
	_ = "STUB: not implemented"
	return
}

func (gg *GangGroupInfo) IsRepresentative(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (gg *GangGroupInfo) ClearCurrentRepresentative(reason string) {
	_ = "STUB: not implemented"
	return
}

func (gg *GangGroupInfo) SetBindingMembers(pods sets.Set[string]) {
	_ = "STUB: not implemented"
	return
}

func (gg *GangGroupInfo) GetBindingMembers() sets.Set[string] {
	_ = "STUB: not implemented"
	return nil
}
