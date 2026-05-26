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

package migration

import (
	"context"
	"sync"
	"time"

	gocache "github.com/patrickmn/go-cache"
	"golang.org/x/time/rate"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/client-go/tools/events"
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/controllers/migration/arbitrator"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/controllers/migration/controllerfinder"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/controllers/migration/evictor"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/controllers/migration/reservation"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/controllers/names"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

const (
	Name                = names.MigrationController
	defaultRequeueAfter = 3 * time.Second
)

var (
	UUIDGenerateFn = uuid.NewUUID
)

var _ framework.EvictPlugin = &Reconciler{}
var _ framework.FilterPlugin = &Reconciler{}

type Reconciler struct {
	client.Client
	args                   *deschedulerconfig.MigrationControllerArgs
	eventRecorder          events.EventRecorder
	reservationInterpreter reservation.Interpreter
	evictorInterpreter     evictor.Interpreter
	controllerFinder       controllerfinder.Interface
	assumedCache           *assumedCache
	clock                  clock.Clock

	arbitrator arbitrator.Arbitrator

	limiterMap      map[deschedulerconfig.MigrationLimitObjectType]map[string]*rate.Limiter
	limiterCacheMap map[deschedulerconfig.MigrationLimitObjectType]*gocache.Cache
	limiterLock     sync.Mutex

	reconcilerUID types.UID
}

func New(ctx context.Context, args runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

// TODO(joseph): It's better that delete reservation asynchronously

func newReconciler(args *deschedulerconfig.MigrationControllerArgs, handle framework.Handle) (*Reconciler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reconciler) Name() string { _ = "STUB: not implemented"; return "" }

func (r *Reconciler) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *Reconciler) scavenger(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (r *Reconciler) doScavenge() { _ = "STUB: not implemented"; return }

// +kubebuilder:rbac:groups=scheduling.koordinator.sh,resources=podmigrationjobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=scheduling.koordinator.sh,resources=podmigrationjobs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=scheduling.koordinator.sh,resources=reservations,verbs=get;list;watch;create;update;patch;delete

// Reconcile reads that state of the cluster for a PodMigrationJob object and makes changes based on the state read
// and what is in the Spec
// Automatically generate RBAC rules to allow the Controller to read and write Deployments
func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// handle the case that the job is created by other reconciler
// NOTE:
// 1. only the PMJs created by MigrationController have annotation AnnotationJobCreatedBy, if a PMJ was created
//    manually by users, we should handle it anyway.
// 2. if a PMJ was indeed created by MigrationController but koord-descheduler has restarted(new UUID generated),
//    we just ignore it.

func (r *Reconciler) getPodByJob(ctx context.Context, job *sev1alpha1.PodMigrationJob) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reconciler) doMigrate(ctx context.Context, job *sev1alpha1.PodMigrationJob) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// sync reservation Unschedulable message to PodMigrationJob

func (r *Reconciler) preparePendingJob(ctx context.Context, job *sev1alpha1.PodMigrationJob) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (r *Reconciler) preparePodRef(ctx context.Context, job *sev1alpha1.PodMigrationJob) (bool, *corev1.Pod, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (r *Reconciler) checkPodExceedObjectLimiter(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Reconciler) exceeded(limiterKey string, limiterType deschedulerconfig.MigrationLimitObjectType) bool {
	_ = "STUB: not implemented"
	return false
}

func getLimiterKeyAndProcessScope(pod *corev1.Pod, limiterType deschedulerconfig.MigrationLimitObjectType) (limiterKey, processScope string) {
	_ = "STUB: not implemented"
	return "", ""
}

func getLogInfo(pod *corev1.Pod, limiterType deschedulerconfig.MigrationLimitObjectType, processScope string) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) requeueJobIfObjectLimiterFailed(ctx context.Context, job *sev1alpha1.PodMigrationJob) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Reconciler) abortJobIfTimeout(ctx context.Context, job *sev1alpha1.PodMigrationJob) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Reconciler) abortJobByInvalidPodRef(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) abortJobByMissingPod(ctx context.Context, job *sev1alpha1.PodMigrationJob, podNamespacedName types.NamespacedName) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) abortJobByMissingReservation(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) abortJobByReservationExpired(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) abortJobByReservationBound(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) abortJobIfReservationBoundByAnotherPod(ctx context.Context, job *sev1alpha1.PodMigrationJob, pod *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Reconciler) abortJobIfReserveOnSameNode(ctx context.Context, job *sev1alpha1.PodMigrationJob, reservationObj reservation.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Reconciler) abortJobByReservationUnschedulable(ctx context.Context, job *sev1alpha1.PodMigrationJob, reservationObj reservation.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) syncReservationScheduleFailed(ctx context.Context, job *sev1alpha1.PodMigrationJob, reservationObj reservation.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) waitForPodBindReservation(ctx context.Context, job *sev1alpha1.PodMigrationJob, reservationObj reservation.Object) (bool, reconcile.Result, error) {
	_ = "STUB: not implemented"
	return false, *new(reconcile.Result), nil
}

func (r *Reconciler) waitForPodReady(ctx context.Context, job *sev1alpha1.PodMigrationJob, podNamespacedName types.NamespacedName) (bool, reconcile.Result, error) {
	_ = "STUB: not implemented"
	return false, *new(reconcile.Result), nil
}

func (r *Reconciler) evictPodDirectly(ctx context.Context, job *sev1alpha1.PodMigrationJob) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (r *Reconciler) evictPod(ctx context.Context, job *sev1alpha1.PodMigrationJob) (bool, reconcile.Result, error) {
	_ = "STUB: not implemented"
	return false, *new(reconcile.Result), nil
}

func (r *Reconciler) prepareJobWithReservationScheduleSuccess(ctx context.Context, job *sev1alpha1.PodMigrationJob, reservationObj reservation.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) trackEvictedPod(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

// namespace limiter should only accept int value. If a percent value is provided, it will be parsed as 0, thus just return.

func (r *Reconciler) track(limit rate.Limit, limiterKey, processScope string, limiterType deschedulerconfig.MigrationLimitObjectType, burst int) {
	_ = "STUB: not implemented"
	return
}

func (r *Reconciler) deleteReservation(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) createReservation(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) setReservationOrder(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) handleReservationCreateSuccess(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) handleReservationBoundSuccess(ctx context.Context, job *sev1alpha1.PodMigrationJob, boundPod *corev1.ObjectReference) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) handleBoundPodReadySuccess(ctx context.Context, job *sev1alpha1.PodMigrationJob) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) waitForPendingPodScheduled(ctx context.Context, job *sev1alpha1.PodMigrationJob) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (r *Reconciler) updateCondition(ctx context.Context, job *sev1alpha1.PodMigrationJob, cond *sev1alpha1.PodMigrationJobCondition) error {
	_ = "STUB: not implemented"
	return nil
}

// Filter checks if a pod can be evicted
func (r *Reconciler) Filter(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (r *Reconciler) PreEvictionFilter(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Reconciler) initObjectLimiters() { _ = "STUB: not implemented"; return }
