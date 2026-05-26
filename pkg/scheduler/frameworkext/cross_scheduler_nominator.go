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

package frameworkext

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	fwktype "k8s.io/kube-scheduler/framework"
)

// CrossSchedulerPodNominator tracks nominated pods from other schedulers
// to enable cross-scheduler resource accounting during filter phase.
type CrossSchedulerPodNominator struct {
	mu sync.RWMutex
	// nominatedPods indexes nominated pods by node name
	nominatedPods map[string][]fwktype.PodInfo
	// nominatedPodToNode maps pod UID to nominated node name for quick lookup
	nominatedPodToNode map[types.UID]string
	// localProfileNames contains the scheduler profile names of the current scheduler instance.
	// Pods from these profiles are excluded (they're handled by the native PodNominator).
	localProfileNames sets.Set[string]
}

// NewCrossSchedulerPodNominator creates a new CrossSchedulerPodNominator with an empty local profile set.
// Profile names are registered lazily via AddLocalProfileName after each framework profile is built.
func NewCrossSchedulerPodNominator() *CrossSchedulerPodNominator {
	_ = "STUB: not implemented"
	return nil
}

// AddLocalProfileName adds a profile name to the local profiles set.
// Pods from local profiles are excluded from cross-scheduler nomination tracking
// because they are handled by the native PodNominator.
// This method is called when a new framework profile is created.
func (n *CrossSchedulerPodNominator) AddLocalProfileName(profileName string) {
	_ = "STUB: not implemented"
	return
}

// NominatedPodsForNode returns the cross-scheduler nominated pods for the given node.
func (n *CrossSchedulerPodNominator) NominatedPodsForNode(nodeName string) []fwktype.PodInfo {
	_ = "STUB: not implemented"
	return nil
}

// Make a copy of the nominated Pods so the caller can mutate safely.

// addNominatedPod adds or updates the nominated pod record.
// It must be called with the write lock held (or in a context where no lock is needed).
func (n *CrossSchedulerPodNominator) addNominatedPod(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

// Remove the old entry first to handle updates (node name may have changed).

// deleteNominatedPod removes the nominated pod record. Must be called with write lock held.
func (n *CrossSchedulerPodNominator) deleteNominatedPod(pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (n *CrossSchedulerPodNominator) ShouldHandle(obj interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

// Handle tombstone objects from the informer cache.

func (n *CrossSchedulerPodNominator) OnAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (n *CrossSchedulerPodNominator) OnUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// We update irrespective of the nominatedNodeName changed or not, to ensure that pod pointer is updated.

func (n *CrossSchedulerPodNominator) OnDelete(obj interface{}) { _ = "STUB: not implemented"; return }

// Handle tombstone objects from the informer cache.
