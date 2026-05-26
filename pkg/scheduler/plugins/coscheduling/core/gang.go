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
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
)

const (
	ErrPodHasNotBeenAttempted         = "gangGroup %s is scheduling and this pod has not been attempted"
	ErrRepresentativePodAlreadyExists = "representative pod %s of gangGroupID %s already exists"
	ErrPodIsNotExistsInGangCache      = "pod %s is not exists in gangCache"
)

var (
	timeNowFn = time.Now
)

const (
	GangFromPodGroupCrd   string = "GangFromPodGroupCrd"
	GangFromPodAnnotation string = "GangFromPodAnnotation"
)

// Gang  basic podGroup info recorded in gangCache:
type Gang struct {
	Name       string
	WaitTime   time.Duration
	CreateTime time.Time

	// strict-mode or non-strict-mode
	Mode                string
	MinRequiredNumber   int
	TotalChildrenNum    int
	GangGroupId         string
	GangGroup           []string
	NetworkTopologySpec *extension.NetworkTopologySpec

	GangGroupInfo *GangGroupInfo

	Children        map[string]*v1.Pod
	PendingChildren map[string]*v1.Pod
	// pods that have already assumed(waiting in Permit stage)
	WaitingForBindChildren map[string]*v1.Pod
	// pods that have already bound
	BoundChildren map[string]*v1.Pod

	// only-waiting, only consider waiting pods
	// waiting-and-running, consider waiting and running pods
	// once-satisfied, once gang is satisfied, no need to consider any status pods
	GangMatchPolicy string

	GangFrom    string
	HasGangInit bool

	lock sync.RWMutex
}

func NewGang(gangName string) *Gang { _ = "STUB: not implemented"; return nil }

func (gang *Gang) tryInitByPodConfig(pod *v1.Pod, args *config.CoschedulingArgs) bool {
	_ = "STUB: not implemented"
	return false
}

// here we assume that Coscheduling's CreateTime equal with the pod's CreateTime

func (gang *Gang) tryInitByPodGroup(pg *v1alpha1.PodGroup, args *config.CoschedulingArgs) {
	_ = "STUB: not implemented"
	return
}

// here we assume that Coscheduling's CreateTime equal with the podGroup CRD CreateTime

func (gang *Gang) SetGangGroupInfo(gangGroupInfo *GangGroupInfo) { _ = "STUB: not implemented"; return }

func (gang *Gang) deletePod(pod *v1.Pod) bool { _ = "STUB: not implemented"; return false }

func (gang *Gang) getGangWaitTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (gang *Gang) getChildrenNum() int { _ = "STUB: not implemented"; return 0 }

func (gang *Gang) getPendingChildrenNum() int { _ = "STUB: not implemented"; return 0 }

func (gang *Gang) getGangMinNum() int { _ = "STUB: not implemented"; return 0 }

func (gang *Gang) getGangTotalNum() int { _ = "STUB: not implemented"; return 0 }

func (gang *Gang) getBoundPodNum() int32 { _ = "STUB: not implemented"; return 0 }

func (gang *Gang) getGangMode() string { _ = "STUB: not implemented"; return "" }

func (gang *Gang) getGangMatchPolicy() string { _ = "STUB: not implemented"; return "" }

func (gang *Gang) getGangAssumedPods() int { _ = "STUB: not implemented"; return 0 }

func (gang *Gang) getGangWaitingPods() int { _ = "STUB: not implemented"; return 0 }

func (gang *Gang) getCreateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (gang *Gang) getGangGroup() []string { _ = "STUB: not implemented"; return nil }

func (gang *Gang) isGangOnceResourceSatisfied() bool { _ = "STUB: not implemented"; return false }

func (gang *Gang) setChild(pod *v1.Pod) { _ = "STUB: not implemented"; return }

func (gang *Gang) addAssumedPod(pod *v1.Pod) { _ = "STUB: not implemented"; return }

func (gang *Gang) delAssumedPod(pod *v1.Pod) { _ = "STUB: not implemented"; return }

func (gang *Gang) getChildrenFromGang() (children []*v1.Pod) { _ = "STUB: not implemented"; return nil }

func (gang *Gang) getPendingChildrenFromGang() (children []*v1.Pod) {
	_ = "STUB: not implemented"
	return nil
}

func (gang *Gang) getWaitingChildrenFromGang() (children []*v1.Pod) {
	_ = "STUB: not implemented"
	return nil
}

func (gang *Gang) isGangFromAnnotation() bool { _ = "STUB: not implemented"; return false }

func (gang *Gang) setResourceSatisfied() { _ = "STUB: not implemented"; return }

func (gang *Gang) addBoundPod(pod *v1.Pod) { _ = "STUB: not implemented"; return }

func (gang *Gang) addWaitingGang() { _ = "STUB: not implemented"; return }

func (gang *Gang) clearWaitingGang() { _ = "STUB: not implemented"; return }

func (gang *Gang) removeWaitingGang() { _ = "STUB: not implemented"; return }

func (gang *Gang) setBindingMembers(pods sets.Set[string]) { _ = "STUB: not implemented"; return }

func (gang *Gang) getBindingMembers() sets.Set[string] { _ = "STUB: not implemented"; return nil }

func (gang *Gang) isGangWorthRequeue() bool { _ = "STUB: not implemented"; return false }

func (gang *Gang) pickSomeChildren() *v1.Pod { _ = "STUB: not implemented"; return nil }

func (gang *Gang) isGangValidForPermit() bool { _ = "STUB: not implemented"; return false }

func (gang *Gang) RecordIfNoRepresentatives(pod *v1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// avoid pod is not exists in gang cache, resulting representativePodKey leak

func (gang *Gang) IsPodRepresentative(pod *v1.Pod) bool { _ = "STUB: not implemented"; return false }

func (gang *Gang) DeleteIfRepresentative(pod *v1.Pod, reason string) {
	_ = "STUB: not implemented"
	return
}

func (gang *Gang) ClearCurrentRepresentative(reason string) { _ = "STUB: not implemented"; return }
