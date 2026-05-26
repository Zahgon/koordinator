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
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	corelister "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/util/workqueue"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	koordclientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	koordinatorinformers "github.com/koordinator-sh/koordinator/pkg/client/informers/externalversions"
	schedulinglister "github.com/koordinator-sh/koordinator/pkg/client/listers/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext"
)

const (
	Name = "reservationController"

	minRetryAfterTime = 3 * time.Second
	maxRetryAfterTime = 15 * time.Second
)

var _ frameworkext.Controller = &Controller{}

type Controller struct {
	sharedInformerFactory      informers.SharedInformerFactory
	koordSharedInformerFactory koordinatorinformers.SharedInformerFactory
	nodeLister                 corelister.NodeLister
	podLister                  corelister.PodLister
	reservationLister          schedulinglister.ReservationLister
	client                     clientset.Interface
	koordClientSet             koordclientset.Interface
	queue                      workqueue.RateLimitingInterface
	numWorker                  int
	isGCDisabled               bool
	gcDuration                 time.Duration
	gcInterval                 time.Duration
	resyncInterval             time.Duration

	lock   sync.RWMutex
	pods   map[string]map[types.UID]*corev1.Pod    // nodeName -> podUID -> pod
	podToR map[types.UID]types.UID                 // podUID to reservationUID
	rToPod map[types.UID]map[types.UID]*corev1.Pod // reservationUID -> podUID -> pod
}

func New(
	sharedInformerFactory informers.SharedInformerFactory,
	koordSharedInformerFactory koordinatorinformers.SharedInformerFactory,
	client clientset.Interface,
	koordClientSet koordclientset.Interface,
	args *config.ReservationArgs,
) *Controller {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Controller) Start() { _ = "STUB: not implemented"; return }

func (c *Controller) worker() { _ = "STUB: not implemented"; return }

func (c *Controller) processNextWorkItem() bool { _ = "STUB: not implemented"; return false }

type result struct {
	requeue      bool
	requeueAfter time.Duration
}

func (c *Controller) sync(key string) (result, error) {
	_ = "STUB: not implemented"
	return *new(result), nil
}

// Clean the reservation-allocated annotation for owner pods when a reservation is deleted.

func (c *Controller) syncPodsForTerminatedReservation(rName string, rUID types.UID) error {
	_ = "STUB: not implemented"
	// If the reservation is deleted, remove the reservationAllocation of the owner pods.
	return nil
}

func (c *Controller) expireReservation(reservation *schedulingv1alpha1.Reservation) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Controller) syncAssignedReservation(reservation *schedulingv1alpha1.Reservation) error {
	_ = "STUB: not implemented"
	return nil
}

// use a pods snapshot to avoid the inconsistency between pods and reservation status

func (c *Controller) syncStatus(reservation *schedulingv1alpha1.Reservation, pods map[types.UID]*corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

// must be called after actualAllocated

func (c *Controller) updateReservationStatus(reservation *schedulingv1alpha1.Reservation) error {
	_ = "STUB: not implemented"
	return nil
}

func isReservationNeedExpiration(r *schedulingv1alpha1.Reservation) bool {
	_ = "STUB: not implemented"
	// 1. failed or succeeded reservations does not need to expire
	return false
}

// 2. disable expiration if TTL is set as 0

// 3. if both TTL and Expires are set, firstly check Expires

func nextSyncTime(r *schedulingv1alpha1.Reservation) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

const (
	defaultResyncInterval = 60 * time.Second
)

func (c *Controller) resyncReservations() { _ = "STUB: not implemented"; return }

// record metrics

// RecordReservationPhases records all possible phases of a reservation as metrics.
// For each phase, it sets the value to 1.0 if it matches the current phase of the reservation,
// otherwise, it sets the value to 0.0.
func RecordReservationPhases(reservation *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	return
}

// If the reservation doesn't have a phase yet,
// consider the pending phase to be the current phase.

// Record the phase with a value of 1.0 if it's the current phase, otherwise 0.0.

func RecordReservationResource(reservation *schedulingv1alpha1.Reservation) {
	_ = "STUB: not implemented"
	return
}

// mCPU -> Core

// bytes -> GiB

// allocatable

// allocated

// utilization
