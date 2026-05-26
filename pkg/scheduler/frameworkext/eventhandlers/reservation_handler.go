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

package eventhandlers

import (
	"regexp"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/profile"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	koordclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	koordinatorinformers "github.com/koordinator-sh/koordinator/pkg/client/informers/externalversions"
	schedulingv1alpha1lister "github.com/koordinator-sh/koordinator/pkg/client/listers/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

// Register schedulingv1alpha1 scheme to report event
var _ = schedulingv1alpha1.AddToScheme(scheme.Scheme)

var (
	// for fit reservation error item
	fitErrPrefix = regexp.MustCompile("^0/[0-9]+ nodes are available: ")
	// for reservation total item
	reserveTotalRe = regexp.MustCompile("^([0-9]+) Reservation\\(s\\) matched owner total$")
	// for reservation name matched item
	reserveNameTotalRe = regexp.MustCompile("^([0-9]+) Reservation\\(s\\) exactly matches the requested reservation name$")
	// for node related item
	reserveNodeDetailRe = regexp.MustCompile("^([0-9]+ Reservation\\(s\\)) (for node reason that .*)$")
	// for reservation detail item
	reserveDetailRe = regexp.MustCompile("^([0-9]+) Reservation\\(s\\) .*$")
	// for reservation level failure message format
	reserveLevelMsgFmt = "0/%d reservations are available: %s."
)

func MakeReservationErrorHandler(
	sched *scheduler.Scheduler,
	schedAdapter frameworkext.Scheduler,
	koordClientSet koordclientset.Interface,
	koordSharedInformerFactory koordinatorinformers.SharedInformerFactory,
) frameworkext.PreErrorHandlerFilter {
	_ = "STUB: not implemented"
	return *new(frameworkext.PreErrorHandlerFilter)
}

// if the pod is not a reserve pod, use the default error handler
// If the Pod failed to schedule or no post-filter plugins, should remove exist NominatedReservation of the Pod.

// If the pod preempting successfully, we should keep the nomination of the pod to reservation, and other pods can not allocate
// the preempted reserved resources in the next cycle.

// export event on reservation level asynchronously if a normal pod specifies the reservation affinity

// for pod specified reservation affinity, export new event on reservation level

// user reason=FailedScheduling-Reservation to avoid event being auto-merged

// handle failure for the reserve pod

// not reservation CR, not pod with reservation affinity

func addNominatedReservation(f framework.Framework, podInfo *framework.QueuedPodInfo, nominatingInfo *fwktype.NominatingInfo) {
	_ = "STUB: not implemented"
	return
}

// input:
// "0/1 nodes are available: 3 Reservation(s) didn't match affinity rules, 1 Reservation(s) is unshedulable, 1 Reservation(s) is unavailable,
// 2 Reservation(s) Insufficient cpu, 1 Reservation(s) Insufficient memory, 1 Insufficient cpu, 1 Insufficient memory.
// 8 Reservation(s) matched owner total, Gang "default/demo-job-podgroup" gets rejected due to pod is unschedulable."
// output:
// "0/8 reservations are available: 3 Reservation(s) didn't match affinity rules, 1 Reservation(s) is unschedulable, 1 Reservation(s) is unavailable,
// 2 Reservation(s) Insufficient cpu, 1 Reservation(s) Insufficient memory."
func generatePodEventOnReservationLevel(errorMsg string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// expect: ["", "3 Reservation(s) ..."]

// "3 Reservations ..., 1 Reservation xxx. 1 Reservation ..."

// "3 Reservations ..., 1 Reservation xxx, 1 Reservation ..."

// ["3 Reservation(s) ...", " 1 Reservation(s) ...", ..., " 8 Reservation(s) matched owner total.", " Gang rejected..."]

// matched total item "8 Reservation(s) matched owner total"

// node related item, e.g. "2 Reservation(s) for node reason that node(s) didn't match pod affinity rules"

// expect: ["2 Reservation(s)", "didn't match pod affinity rules"]

// reservation itself item, append to details, e.g. " 1 Reservation(s) ..."

// put the reservation name at the front, and put the node-related details at the end

func handleReservationSchedulingFailure(sched *scheduler.Scheduler,
	schedAdapter frameworkext.Scheduler,
	reservationLister schedulingv1alpha1lister.ReservationLister,
	koordClientSet koordclientset.Interface) scheduler.FailureHandlerFn {
	_ = "STUB: not implemented"
	// Here we follow the procedure of the normal pod handling in the framework, except using the reservation object.
	return *new(scheduler.FailureHandlerFn)
}

// Basically, AddUnschedulableIfNotPresent calls DonePod internally.
// But, AddUnschedulableIfNotPresent isn't called in some corner cases.
// Here, we call DonePod explicitly to avoid leaking the pod.

// NOTE: If the pod is a reserve pod, we simply check the corresponding reservation status if the reserve pod
// need requeue for the next scheduling cycle.

// Inject UnschedulablePlugins to PodInfo, which will be used later for moving Pods between queues efficiently.

// Check if the corresponding reservation exists in informer cache.

// We need to call DonePod here because we don't call AddUnschedulableIfNotPresent in this case.

// The scheduler name of a reservation can change in-flight, so we need to double-check if the scheduler
// is not matched anymore. If unmatched, we should abort the failure handling to avoid applying a
// failure state with another scheduler concurrently.
// TODO: Add same check for pods

// In the case of extender, the pod may have been bound successfully, but timed out returning its response to the scheduler.
// It could result in the live version to carry .spec.nodeName, and that's inconsistent with the internal-queued version.

// We need to call DonePod here because we don't call AddUnschedulableIfNotPresent in this case.

// nominate for the reserve pod if it is
// FIXME: We expect use the default nominator for a nominated reserve pod, since it makes no benefit to
//   maintain another nominator. However, the default nominator relies the podLister to fetch the real pod
//   from the informer cache.

// TODO: use apiDispatcher

func updateReservationStatus(client koordclientset.Interface, reservationLister schedulingv1alpha1lister.ReservationLister, rName string, schedulingErr error) {
	_ = "STUB: not implemented"
	return
}

func truncateMessage(message string) string { _ = "STUB: not implemented"; return "" }

func reservationEventHandlers(sched *scheduler.Scheduler, schedAdapter frameworkext.Scheduler) cache.ResourceEventHandler {
	_ = "STUB: not implemented"
	// TODO: add metrics for handler latency
	return *new(cache.ResourceEventHandler)
}

// addReservation handles an Add event for a Reservation object.
// 1. If adding an unassigned and responsible Reservation, add it to the scheduling queue.
// 2. If adding an available (assigned and not terminated) Reservation, add it to the scheduler cache.
func addReservation(sched *scheduler.Scheduler, schedAdapter frameworkext.Scheduler, r *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	// To avoid scheduler cache corrupted, only add valid reservation into cache.
	return
}

// updateReservation handles an Update event for a Reservation object.
// Base states: unassigned (active, nodeName is empty), available (assigned, nodeName not empty), terminated (Failed/Succeeded).
// Supported transition states:
// 0. terminated -> terminated: keep terminal state
// 1. unassigned -> available: reservation gets scheduled
// 2. available -> terminated: reservation expires or completes
// 3. unassigned -> terminated: reservation scheduling fails permanently
// 4. available -> unassigned (extended): assumed binding failure, rollback for retry
// 5. available -> available with different nodeName (extended): node migration in multi-scheduler scenarios
func updateReservation(sched *scheduler.Scheduler, schedAdapter frameworkext.Scheduler, oldR, newR *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	// To avoid scheduler cache corrupted, only update valid reservation into cache.
	return
}

// Case 0: Keep terminated.

// Case 1: Keep available

// Update available reservation in cache (handles nodeName change internally)

// Case 2: From unassigned to available (scheduling succeeded)

// Remove from scheduling queue since it's now scheduled

// Case 3: From available to terminated

// Remove from cache, no need to requeue for terminal states

// Ensure it's removed from scheduling queue

// Case 4: From available to unassigned (assumed binding failure, rollback)
// This is an extended case for multi-scheduler scenarios. We cleanup
// cache before requeue the request to avoid cache leak.

// Remove from cache since it's no longer scheduled

// Requeue for retry

// Case 5: Keep unassigned (active)

// Handle scheduler name changes

// Case 6: From unassigned to terminated

// Unexpected state transitions

// deleteReservation handles a Delete event for a Reservation object.
// 1. Delete the Reservation from the scheduler cache if it exists.
// 2. If deleting an unassigned and responsible Reservation, delete it from the scheduling queue.
func deleteReservation(sched *scheduler.Scheduler, schedAdapter frameworkext.Scheduler, r *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	return
}

func toReservation(obj interface{}) *schedulingv1alpha1.Reservation {
	_ = "STUB: not implemented"
	return nil
}

func addReservationToSchedulerCache(sched frameworkext.Scheduler, r *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	return
}

// update pod cache and trigger pod assigned event for scheduling queue

func updateReservationInSchedulerCache(sched frameworkext.Scheduler, oldR, newR *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	// A fast lookup for unscheduled reservation.
	return
}

// A delete event followed by an immediate add event may be merged into a update event.
// In extended multi-scheduler scenarios the nodeName may change as a merged event.
// In this case, we should invalidate the old object, and then add the new object,
// while it does not guarantee the resources between the delete and add process.

func deleteReservationFromSchedulerCache(sched frameworkext.Scheduler, r *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	return
}

// FIXME: Due to the nodeInfo ports cleanup, here we need the ReservationInfo to get deleted by the global event handler.
//   Without this dependency, the plugin holds the reservation cache can safely delete the ReservationInfo, and no need to do it here.
// https://github.com/koordinator-sh/koordinator/issues/1247

// The ports allocated from the reservation by some owner pods should not be removed from the NodeInfo.
// https://github.com/koordinator-sh/koordinator/issues/1247

// The Pod status in the Cache must be refreshed once to ensure that subsequent deletions are valid.

func addReservationToSchedulingQueue(sched frameworkext.Scheduler, r *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	return
}

func updateReservationInSchedulingQueue(sched frameworkext.Scheduler, oldR, newR *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	// Bypass update event that carries identical objects to avoid duplicate scheduling.
	// https://github.com/kubernetes/kubernetes/pull/96071
	return
}

// Normal update flow: check if assumed

// SAFETY: Don't update assumed pods in the queue
// They are in the binding process and should not be disturbed

// Update in scheduling queue for other cases (e.g., spec changes)

func deleteReservationFromSchedulingQueue(sched *scheduler.Scheduler, schedAdapter frameworkext.Scheduler, r *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	return
}

// Try to delete from scheduling queue
// For scenario 2 (scheduling success), the pod may not be in the queue anymore (already assumed)

func isResponsibleForReservation(profiles profile.Map, r *schedulingv1alpha1.Reservation) bool {
	_ = "STUB: not implemented"
	return false
}

// isReservationActive checks if the reservation is unassigned (not scheduled to a node yet) and not terminated.
// A reservation is considered active if:
// 1. It has no nodeName assigned (unscheduled)
// 2. It is not in a terminal state (Failed or Succeeded)
func isReservationActive(r *schedulingv1alpha1.Reservation) bool {
	_ = "STUB: not implemented"
	return false
}
