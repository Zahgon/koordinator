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

	v1 "k8s.io/api/core/v1"
	listerv1 "k8s.io/client-go/listers/core/v1"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
	pgclientset "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/clientset/versioned"
	pglister "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/generated/listers/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/workloadauditor"
)

type GangCache struct {
	lock             *sync.RWMutex
	gangItems        map[string]*Gang
	gangGroupInfoMap map[string]*GangGroupInfo
	pluginArgs       *config.CoschedulingArgs
	podLister        listerv1.PodLister
	pgLister         pglister.PodGroupLister
	pgClient         pgclientset.Interface
	handle           fwktype.Handle

	workloadAuditor workloadauditor.WorkloadAuditor
}

func NewGangCache(args *config.CoschedulingArgs, podLister listerv1.PodLister, pgLister pglister.PodGroupLister, client pgclientset.Interface, handle fwktype.Handle) *GangCache {
	_ = "STUB: not implemented"
	return nil
}

func (gangCache *GangCache) getGangGroupInfo(gangGroupId string, gangGroup []string, createIfNotExist bool) (gangGroupInfo *GangGroupInfo, created bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (gangCache *GangCache) deleteGangGroupInfo(gangGroupId string) {
	_ = "STUB: not implemented"
	return
}

func (gangCache *GangCache) getGangFromCacheByGangId(gangId string, createIfNotExist bool) *Gang {
	_ = "STUB: not implemented"
	return nil
}

func (gangCache *GangCache) getAllGangsFromCache() map[string]*Gang {
	_ = "STUB: not implemented"
	return nil
}

func (gangCache *GangCache) deleteGangFromCacheByGangId(gangId string) {
	_ = "STUB: not implemented"
	return
}

func (gangCache *GangCache) onPodAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (gangCache *GangCache) onPodAddInternal(obj interface{}, action string) {
	_ = "STUB: not implemented"
	return
}

// the gang is created in Annotation way

// Detect initial gating state for the newly added pod

// only UT will go here

func (gangCache *GangCache) onPodUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// Detect gating transitions for gang pods

func (gangCache *GangCache) onPodDelete(obj interface{}) { _ = "STUB: not implemented"; return }

func (gangCache *GangCache) onPodGroupAdd(obj interface{}) { _ = "STUB: not implemented"; return }

// only UT will go here

func (gangCache *GangCache) onPodGroupUpdate(oldObj interface{}, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// When PodGroup transitions from a pending phase to a non-pending phase,
// delete the gang record from the workload auditor.

// only UT will go here

func (gangCache *GangCache) onPodGroupDelete(obj interface{}) { _ = "STUB: not implemented"; return }

func (gangCache *GangCache) getPendingPods(gangGroup []string) []*v1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (gangCache *GangCache) getPendingPodsNum(gangGroup []string) int {
	_ = "STUB: not implemented"
	return 0
}

func (gangCache *GangCache) getWaitingPods(gangGroup []string) []*v1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (gangCache *GangCache) getWaitingPodsNum(gangGroup []string) int {
	_ = "STUB: not implemented"
	return 0
}

// isPodGroupPendingPhase returns true if the PodGroup phase indicates
// it has not yet finished scheduling (empty, Pending, PreScheduling, Scheduling).
func isPodGroupPendingPhase(phase v1alpha1.PodGroupPhase) bool {
	_ = "STUB: not implemented"
	return false
}

// isResponsibleForPod returns true if the pod's scheduler name matches
// the profile name of this scheduler instance.
func (gangCache *GangCache) isResponsibleForPod(pod *v1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}
