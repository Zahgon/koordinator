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
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/networktopology"
)

const (
	ReasonGangIsNil                        = "gang is nil, may be already deleted"
	ReasonAllPendingPodsIsAlreadyAttempted = "all pending pods is already been attempted"
	ReasonGangIsSucceed                    = "gang is already succeed"

	ReasonFirstPodPassPreFilter = "the first pod of gang pass preFilter"
)

/*
GangSchedulingContextHolder is used to hold the gang scheduling context
- It will be set
  - 1. when the first Pod of the Gang pass PreFilter

- It will be cleared when
  - 1. gang is deleted or
  - 2. the gang has no pending pods that have not been retried or
  - 3. gang is already succeed
*/
type GangSchedulingContextHolder struct {
	sync.RWMutex
	gangSchedulingContext *GangSchedulingContext
}

func (h *GangSchedulingContextHolder) getCurrentGangSchedulingContext() *GangSchedulingContext {
	_ = "STUB: not implemented"
	return nil
}

func (h *GangSchedulingContextHolder) clearGangSchedulingContext(reason string) {
	_ = "STUB: not implemented"
	return
}

func (h *GangSchedulingContextHolder) setGangSchedulingContext(gangSchedulingContext *GangSchedulingContext, reason string) {
	_ = "STUB: not implemented"
	return
}

type GangSchedulingContext struct {
	startTime   time.Time
	gangGroup   sets.Set[string]
	gangGroupID string
	firstPod    *corev1.Pod

	// secure alreadyAttemptedPods to avoid concurrent map read and write
	sync.RWMutex
	alreadyAttemptedPods sets.Set[string]

	triggerPod        *corev1.Pod
	failedMessage     string
	preemptionMessage string
	suggestion        *frameworkext.ScheduleSuggestion

	networkTopologySpec         *extension.NetworkTopologySpec
	networkTopologySnapshot     *networktopology.TreeSnapshot
	networkTopologyPlannedNodes map[string]string // podKey -> nodeName
}
