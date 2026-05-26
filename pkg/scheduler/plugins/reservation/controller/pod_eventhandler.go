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

package controller

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
)

func (c *Controller) onPodAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (c *Controller) onPodUpdate(oldObj, newObj interface{}) { _ = "STUB: not implemented"; return }

// become unassigned, a special case for multi-scheduler

func (c *Controller) onPodDelete(obj interface{}) { _ = "STUB: not implemented"; return }

// enqueueIfPodBoundReservation will enqueue the reservation and  return the reservation allocated
// if the pod is bound to the reservation
func (c *Controller) enqueueIfPodBoundReservation(pod *corev1.Pod) *apiext.ReservationAllocated {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) updatePod(pod *corev1.Pod, rAllocated *apiext.ReservationAllocated) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) deletePod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (c *Controller) getPodsOnNode(nodeName string) map[types.UID]*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) getPodsOnReservation(rUID types.UID) map[types.UID]*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}
