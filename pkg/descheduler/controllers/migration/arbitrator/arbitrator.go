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

package arbitrator

import (
	"context"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/events"
	"k8s.io/klog/v2"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

const (
	AnnotationPassedArbitration = "descheduler.koordinator.sh/passed-arbitration"
	AnnotationPodArbitrating    = "descheduler.koordinator.sh/pod-arbitrating"
)

var enqueueLog = klog.Background().WithName("eventHandler").WithName("arbitratorImpl")

type MigrationFilter interface {
	Filter(pod *corev1.Pod) bool
	PreEvictionFilter(pod *corev1.Pod) bool
}

type Arbitrator interface {
	MigrationFilter
	AddPodMigrationJob(job *v1alpha1.PodMigrationJob)
	DeletePodMigrationJob(job *v1alpha1.PodMigrationJob)
}

// SortFn stably sorts PodMigrationJobs slice based on a certain strategy. Users
// can implement different SortFn according to their needs.
type SortFn func(jobs []*v1alpha1.PodMigrationJob, podOfJob map[*v1alpha1.PodMigrationJob]*corev1.Pod) []*v1alpha1.PodMigrationJob

type arbitratorImpl struct {
	waitingCollection map[types.UID]*v1alpha1.PodMigrationJob
	interval          time.Duration

	sorts  []SortFn
	filter *filter

	client        client.Client
	eventRecorder events.EventRecorder
	mu            sync.Mutex
}

// New creates an arbitratorImpl based on parameters.
func New(args *config.MigrationControllerArgs, options Options) (Arbitrator, error) {
	_ = "STUB: not implemented"
	return *new(Arbitrator), nil
}

// AddPodMigrationJob adds a PodMigrationJob waiting to be arbitrated to Arbitrator.
// It is safe to be called concurrently by multiple goroutines.
func (a *arbitratorImpl) AddPodMigrationJob(job *v1alpha1.PodMigrationJob) {
	_ = "STUB: not implemented"
	return
}

// DeletePodMigrationJob removes a deleted PodMigrationJob from Arbitrator.
// It is safe to be called concurrently by multiple goroutines.
func (a *arbitratorImpl) DeletePodMigrationJob(job *v1alpha1.PodMigrationJob) {
	_ = "STUB: not implemented"
	return
}

// Start starts the goroutine to arbitrate jobs periodically.
func (a *arbitratorImpl) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Filter checks if a pod can be evicted
func (a *arbitratorImpl) Filter(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (a *arbitratorImpl) PreEvictionFilter(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

// sort stably sorts jobs, outputs the sorted results and corresponding ranking map.
func (a *arbitratorImpl) sort(jobs []*v1alpha1.PodMigrationJob, podOfJob map[*v1alpha1.PodMigrationJob]*corev1.Pod) []*v1alpha1.PodMigrationJob {
	_ = "STUB: not implemented"
	return nil
}

// filtering calls nonRetryablePodFilter and retryablePodFilter to filter one PodMigrationJob.
func (a *arbitratorImpl) filtering(pod *corev1.Pod) (isFailed, isPassed bool) {
	_ = "STUB: not implemented"
	return false, false
}

// updatePassedJob does something after PodMigrationJob passed the filter.
func (a *arbitratorImpl) updatePassedJob(job *v1alpha1.PodMigrationJob) {
	_ = "STUB: not implemented"
	// add annotation AnnotationPassedArbitration
	return
}

// remove job from the waitingCollection

// doOnceArbitrate performs an arbitrate operation on PodMigrationJobs in the waitingCollection.
func (a *arbitratorImpl) doOnceArbitrate() {
	_ = "STUB: not implemented"
	// copy jobs from waitingCollection
	return
}

// sort

// filter

// copyJobs copy jobs from waitingCollection
func (a *arbitratorImpl) copyJobs() []*v1alpha1.PodMigrationJob {
	_ = "STUB: not implemented"
	return nil
}

func (a *arbitratorImpl) updateFailedJob(job *v1alpha1.PodMigrationJob, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	// change phase to Failed
	return
}

// delete from waitingCollection

type Options struct {
	Client        client.Client
	EventRecorder events.EventRecorder
	Manager       controllerruntime.Manager
	Handle        framework.Handle
}

func getPodForJob(c client.Client, jobs []*v1alpha1.PodMigrationJob) map[*v1alpha1.PodMigrationJob]*corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func markPodArbitrating(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func checkPodArbitrating(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }
