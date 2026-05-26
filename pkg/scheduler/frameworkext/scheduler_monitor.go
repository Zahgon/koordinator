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
	"time"

	"github.com/spf13/pflag"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	fwktype "k8s.io/kube-scheduler/framework"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

var (
	schedulerMonitorPeriod         = 10 * time.Second
	schedulingTimeout              = 30 * time.Second
	schedulingDropUnhandledTimeout = 5 * time.Second

	logWarningF = klog.Warningf
)

func init() {
	pflag.DurationVar(&schedulerMonitorPeriod, "scheduler-monitor-period", schedulerMonitorPeriod, "Execution period of scheduler monitor")
	pflag.DurationVar(&schedulingTimeout, "scheduling-timeout", schedulingTimeout, "The maximum acceptable scheduling time interval. After timeout, the metric will be updated and the log will be printed.")
	pflag.DurationVar(&schedulingDropUnhandledTimeout, "scheduling-drop-unhandled-timeout", schedulingDropUnhandledTimeout, "The maximum acceptable scheduling time interval to drop the invalid pod context when the pod is dequeued but has not handled to schedule.")
}

var (
	StartMonitor       = defaultStartMonitor       // start a schedule attempt
	CompleteMonitor    = defaultCompleteMonitor    // complete a schedule attempt
	RecordQueuePodInfo = defaultRecordQueuePodInfo // dequeue a schedule attempt
	GCMonitor          = defaultGCMonitor          // garbage collect an unhandled schedule attempt
)

type SchedulerMonitor struct {
	timeout          time.Duration
	unhandledTimeout time.Duration
	lock             sync.Mutex
	schedulingPods   map[types.UID]podScheduleState
}

type podScheduleState struct {
	namespace     string
	name          string
	schedulerName string
	// scheduling info
	start time.Time
	// queue info
	dequeued        time.Time
	lastEnqueued    time.Time
	attempts        int
	initialEnqueued *time.Time
	// for extensions
	extensionInfo interface{}
}

func NewSchedulerMonitor(period time.Duration, timeout time.Duration) *SchedulerMonitor {
	_ = "STUB: not implemented"
	return nil
}

func (m *SchedulerMonitor) monitor() { _ = "STUB: not implemented"; return }

func (m *SchedulerMonitor) RecordNextPod(podInfo *framework.QueuedPodInfo) {
	_ = "STUB: not implemented"
	return
}

// clean up from the cache when the pod is terminating

func (m *SchedulerMonitor) StartMonitoring(pod *corev1.Pod) { _ = "STUB: not implemented"; return }

func (m *SchedulerMonitor) Complete(pod *corev1.Pod, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return
}

func isPodUnhandledExceedingTimeout(state *podScheduleState, now time.Time, timeout time.Duration) (skipped bool, toDelete bool) {
	_ = "STUB: not implemented"
	return false,
		// pod is handled
		false
}

// unhandled exceeding timeout

// unhandled in timeout

func defaultStartMonitor(pod *corev1.Pod, state *podScheduleState) {
	_ = "STUB: not implemented"
	return
}

func defaultCompleteMonitor(pod *corev1.Pod, state *podScheduleState, end time.Time, timeout time.Duration, status *fwktype.Status) {
	_ = "STUB: not implemented"
	return
}

func defaultRecordQueuePodInfo(podInfo *framework.QueuedPodInfo, state *podScheduleState) {
	_ = "STUB: not implemented"
	return
}

func defaultGCMonitor(uid types.UID, state *podScheduleState, end time.Time) {
	_ = "STUB: not implemented"
	return
}

func recordIfSchedulingTimeout(uid types.UID, state *podScheduleState, now time.Time, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}
