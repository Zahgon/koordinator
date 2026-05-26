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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/klog/v2"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

var AssignedPodDelete = fwktype.ClusterEvent{Resource: fwktype.Pod, ActionType: fwktype.Delete, CustomLabel: "AssignedPodDelete"}

var podPool = &sync.Pool{
	New: func() interface{} {
		uid := uuid.NewUUID()
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(uid),
				Namespace: "default",
				UID:       uid,
			},
		}
		return pod
	},
}

// Scheduler exports scheduler internal cache and queue interface for testability.
type Scheduler interface {
	GetCache() SchedulerCache
	GetSchedulingQueue() SchedulingQueue
}

type SchedulerCache interface {
	AddPod(logger klog.Logger, pod *corev1.Pod) error
	UpdatePod(logger klog.Logger, oldPod, newPod *corev1.Pod) error
	RemovePod(logger klog.Logger, pod *corev1.Pod) error
	AssumePod(logger klog.Logger, pod *corev1.Pod) error
	IsAssumedPod(pod *corev1.Pod) (bool, error)
	GetPod(pod *corev1.Pod) (*corev1.Pod, error)
	ForgetPod(logger klog.Logger, pod *corev1.Pod) error
	InvalidNodeInfo(logger klog.Logger, nodeName string) error
}

type PreEnqueueCheck func(pod *corev1.Pod) bool

type SchedulingQueue interface {
	Add(logger klog.Logger, pod *corev1.Pod)
	Update(logger klog.Logger, oldPod, newPod *corev1.Pod)
	Delete(pod *corev1.Pod)
	AddUnschedulableIfNotPresent(logger klog.Logger, pod *framework.QueuedPodInfo, podSchedulingCycle int64) error
	SchedulingCycle() int64
	AssignedPodAdded(logger klog.Logger, pod *corev1.Pod)
	AssignedPodUpdated(logger klog.Logger, oldPod, newPod *corev1.Pod, event fwktype.ClusterEvent)
	MoveAllToActiveOrBackoffQueue(logger klog.Logger, event fwktype.ClusterEvent, oldObj, newObj interface{}, preCheck PreEnqueueCheck)
	Activate(logger klog.Logger, pods map[string]*corev1.Pod)
	Done(types.UID)
}

var _ Scheduler = &SchedulerAdapter{}

type SchedulerAdapter struct {
	Scheduler *scheduler.Scheduler
}

func (s *SchedulerAdapter) GetCache() SchedulerCache {
	_ = "STUB: not implemented"
	return *new(SchedulerCache)
}

func (s *SchedulerAdapter) GetSchedulingQueue() SchedulingQueue {
	_ = "STUB: not implemented"
	return *new(SchedulingQueue)
}

func (s *SchedulerAdapter) MoveAllToActiveOrBackoffQueue(logger klog.Logger, event fwktype.ClusterEvent, oldObj, newObj interface{}, preCheck PreEnqueueCheck) {
	_ = "STUB: not implemented"
	return
}

var _ SchedulerCache = &cacheAdapter{}

type cacheAdapter struct {
	scheduler *scheduler.Scheduler
}

func (c *cacheAdapter) AddPod(logger klog.Logger, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cacheAdapter) UpdatePod(logger klog.Logger, oldPod, newPod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cacheAdapter) RemovePod(logger klog.Logger, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cacheAdapter) AssumePod(logger klog.Logger, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cacheAdapter) IsAssumedPod(pod *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *cacheAdapter) GetPod(pod *corev1.Pod) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *cacheAdapter) ForgetPod(logger klog.Logger, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *cacheAdapter) InvalidNodeInfo(logger klog.Logger, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

var _ SchedulingQueue = &queueAdapter{}

type queueAdapter struct {
	scheduler *scheduler.Scheduler
}

func (q *queueAdapter) Add(logger klog.Logger, pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (q *queueAdapter) Update(logger klog.Logger, oldPod, newPod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (q *queueAdapter) Delete(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (q *queueAdapter) AddUnschedulableIfNotPresent(logger klog.Logger, pInfo *framework.QueuedPodInfo, podSchedulingCycle int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (q *queueAdapter) SchedulingCycle() int64 { _ = "STUB: not implemented"; return 0 }

func (q *queueAdapter) AssignedPodAdded(logger klog.Logger, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (q *queueAdapter) AssignedPodUpdated(logger klog.Logger, oldPod, newPod *corev1.Pod, event fwktype.ClusterEvent) {
	_ = "STUB: not implemented"
	return
}

func (q *queueAdapter) MoveAllToActiveOrBackoffQueue(logger klog.Logger, event fwktype.ClusterEvent, oldObj, newObj interface{}, preCheck PreEnqueueCheck) {
	_ = "STUB: not implemented"
	return
}

func (q *queueAdapter) Activate(logger klog.Logger, pods map[string]*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (q *queueAdapter) Done(pod types.UID) { _ = "STUB: not implemented"; return }

var _ Scheduler = &FakeScheduler{}
var _ SchedulingQueue = &FakeQueue{}

type FakeScheduler struct {
	Pods       map[string]*corev1.Pod
	AssumedPod map[string]*corev1.Pod
	Queue      *FakeQueue
	lock       sync.Mutex
	NodeInfos  map[string]*framework.NodeInfo
}

func NewFakeScheduler() *FakeScheduler { _ = "STUB: not implemented"; return nil }

type FakeQueue struct {
	Pods                map[string]*corev1.Pod
	UnschedulablePods   map[string]*corev1.Pod
	AssignedPods        map[string]*corev1.Pod
	AssignedUpdatedPods map[string]*corev1.Pod
}

func (f *FakeScheduler) GetCache() SchedulerCache {
	_ = "STUB: not implemented"
	return *new(SchedulerCache)
}

func (f *FakeScheduler) GetSchedulingQueue() SchedulingQueue {
	_ = "STUB: not implemented"
	return *new(SchedulingQueue)
}

func (f *FakeScheduler) AddPod(logger klog.Logger, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeScheduler) UpdatePod(logger klog.Logger, oldPod, newPod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeScheduler) RemovePod(logger klog.Logger, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeScheduler) AssumePod(logger klog.Logger, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeScheduler) IsAssumedPod(pod *corev1.Pod) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *FakeScheduler) GetPod(pod *corev1.Pod) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeScheduler) ForgetPod(logger klog.Logger, pod *corev1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeScheduler) InvalidNodeInfo(logger klog.Logger, nodeName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeQueue) Add(logger klog.Logger, pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (f *FakeQueue) Update(logger klog.Logger, oldPod, newPod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeQueue) Delete(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (f *FakeQueue) AddUnschedulableIfNotPresent(logger klog.Logger, pod *framework.QueuedPodInfo, podSchedulingCycle int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeQueue) SchedulingCycle() int64 { _ = "STUB: not implemented"; return 0 }

func (f *FakeQueue) AssignedPodAdded(logger klog.Logger, pod *corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeQueue) AssignedPodUpdated(logger klog.Logger, oldPod, newPod *corev1.Pod, event fwktype.ClusterEvent) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeQueue) MoveAllToActiveOrBackoffQueue(logger klog.Logger, event fwktype.ClusterEvent, oldObj, newObj interface{}, preCheck PreEnqueueCheck) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeQueue) Activate(logger klog.Logger, pods map[string]*corev1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeQueue) Done(pod types.UID) { _ = "STUB: not implemented"; return }
