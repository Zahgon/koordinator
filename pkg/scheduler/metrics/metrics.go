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

package metrics

import (
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/component-base/metrics"
	schedulermetrics "k8s.io/kubernetes/pkg/scheduler/metrics"

	utilmetrics "github.com/koordinator-sh/koordinator/pkg/util/metrics"
)

const (
	NodeNameKey = "node_name"
)

// All the histogram based metrics have 1ms as size for the smallest bucket.
var (
	SchedulingTimeout = metrics.NewCounterVec(
		&metrics.CounterOpts{
			Subsystem:      schedulermetrics.SchedulerSubsystem,
			Name:           "scheduling_timeout",
			Help:           "The currently scheduled Pod exceeds the maximum acceptable time interval",
			StabilityLevel: metrics.STABLE,
		}, []string{"profile"})
	ReservationStatusPhase = metrics.NewGaugeVec(
		&metrics.GaugeOpts{
			Subsystem: schedulermetrics.SchedulerSubsystem,
			Name:      "reservation_status_phase",
			Help:      `The current number of reservations in each status phase (e.g. Pending, Available, Succeeded, Failed)`,
		}, []string{"name", "phase"})
	ReservationResource = utilmetrics.NewGCGaugeVec(
		"reservation_resource",
		prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: schedulermetrics.SchedulerSubsystem,
				Name:      "reservation_resource",
				Help:      "Resource metrics for a reservation, including allocatable, allocated, and utilization with unit.",
			},
			[]string{"type", "name", "resource", "unit"},
		),
	)

	PodSchedulingEvaluatedNodes = metrics.NewHistogram(
		&metrics.HistogramOpts{
			Subsystem: schedulermetrics.SchedulerSubsystem,
			Name:      "pod_scheduling_evaluated_nodes",
			Help:      "The number of nodes the scheduler evaluated the pod against in the filtering phase and beyond when find the suggested node",
			Buckets:   metrics.ExponentialBuckets(1, 2, 24),
		})
	PodSchedulingFeasibleNodes = metrics.NewHistogram(
		&metrics.HistogramOpts{
			Subsystem: schedulermetrics.SchedulerSubsystem,
			Name:      "pod_scheduling_feasible_nodes",
			Help:      "The number of of nodes out of the evaluated ones that fit the pod when find the suggested node",
			Buckets:   metrics.ExponentialBuckets(1, 2, 24),
		})

	ElasticQuotaProcessLatency = metrics.NewHistogramVec(
		&metrics.HistogramOpts{
			Subsystem: schedulermetrics.SchedulerSubsystem,
			Name:      "elastic_quota_process_latency",
			Help:      "elastic quota process latency in second",
			Buckets:   metrics.ExponentialBuckets(0.00001, 2, 24),
		},
		[]string{"operation"},
	)
	SecondaryDeviceNotWellPlannedNodes = metrics.NewGaugeVec(
		&metrics.GaugeOpts{
			Subsystem: schedulermetrics.SchedulerSubsystem,
			Name:      "secondary_device_not_well_planned",
			Help:      "The number of secondary device not well planned",
		},
		[]string{NodeNameKey},
	)
	WaitingGangGroupNumber = metrics.NewGaugeVec(
		&metrics.GaugeOpts{
			Subsystem:      schedulermetrics.SchedulerSubsystem,
			Name:           "waiting_gang_group_number",
			Help:           "The number of GangGroups in Waiting",
			StabilityLevel: metrics.STABLE,
		}, nil)
	NextPodDeleteFromQueueLatency = metrics.NewHistogramVec(
		&metrics.HistogramOpts{
			Subsystem: schedulermetrics.SchedulerSubsystem,
			Name:      "next_pod_delete_from_queue_latency",
			Help:      "run next pod plugins, the latency of deleting pod from queue",
			Buckets:   metrics.ExponentialBuckets(0.00001, 2, 24),
		}, nil)
	ElasticQuotaHookPluginLatency = metrics.NewHistogramVec(
		&metrics.HistogramOpts{
			Subsystem:      schedulermetrics.SchedulerSubsystem,
			Name:           "elastic_quota_hook_plugin_latency",
			Help:           "elastic quota hook plugin latency in seconds",
			Buckets:        metrics.ExponentialBuckets(0.000001, 2, 24),
			StabilityLevel: metrics.ALPHA,
		},
		[]string{"plugin", "operation"},
	)
	JobPreemptionDuration = utilmetrics.NewGCHistogramVec("job_preemption_duration_seconds", prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: schedulermetrics.SchedulerSubsystem,
			Name:      "job_preemption_duration_seconds",
			Help:      "Latency for running Coscheduling plugin postFilter.",
			// Start with 0.1ms with the last bucket being [~200ms, Inf)
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 20),
		}, []string{"jobName", "result"}))

	GangScheduleCycleDuration = metrics.NewHistogramVec(
		&metrics.HistogramOpts{
			Subsystem: schedulermetrics.SchedulerSubsystem,
			Name:      "gang_schedule_cycle_duration_seconds",
			Help:      "Duration of a gang scheduling cycle, from the first pod passing PreFilter until the context is cleared, labeled by clear reason and job size bucket.",
			// Explicit buckets tuned for a 10s SLO with current p50 ~30s.
			// Dense around the 10s boundary (7.5/10/12.5) to get accurate
			// quantiles there, and 15/20/25/30 to observe the current range.
			Buckets: []float64{
				0.1, 0.5, 1, 2, 3, 5, 7.5, 10, 12.5, 15, 20, 25, 30, 45, 60, 120,
			},
			StabilityLevel: metrics.ALPHA,
		},
		[]string{"reason", "job_size"},
	)

	metricsList = []metrics.Registerable{
		SchedulingTimeout,
		ReservationStatusPhase,
		PodSchedulingEvaluatedNodes,
		PodSchedulingFeasibleNodes,
		ElasticQuotaProcessLatency,
		SecondaryDeviceNotWellPlannedNodes,
		WaitingGangGroupNumber,
		NextPodDeleteFromQueueLatency,
		ElasticQuotaHookPluginLatency,
		GangScheduleCycleDuration,
	}

	gcMetricsList = []prometheus.Collector{
		ReservationResource.GetGaugeVec(),
		JobPreemptionDuration.GetHistogramVec(),
	}
)

const (
	reservationNameKey         = "name"
	reservationPhaseKey        = "phase"
	reservationResourceKey     = "resource"
	reservationResourceTypeKey = "type"
	reservationResourceUnitKey = "unit"
)

const (
	UnitCore  = "core"
	UnitGiB   = "Gi"
	UnitRaw   = "raw"
	UnitRatio = "ratio"
)

const (
	TypeAllocatable = "allocatable"
	TypeAllocated   = "allocated"
	TypeUtilization = "utilization"
)

var registerMetrics sync.Once

// Register all metrics.
func Register() {
	_ = "STUB: not implemented"
	// Register the metrics.
	return
}

// RegisterMetrics registers a list of metrics.
// This function is exported because it is intended to be used by out-of-tree plugins to register their custom metrics.
func RegisterMetrics(extraMetrics ...metrics.Registerable) { _ = "STUB: not implemented"; return }

// RegisterGCMetrics registers garbage collection metrics.
func RegisterGCMetrics(gcMetrics ...prometheus.Collector) { _ = "STUB: not implemented"; return }

// RecordReservationPhase records the phase of a reservation as a metric.
// It uses the provided name, phase, and value to set the metric with specific labels.
func RecordReservationPhase(name string, phase string, value float64) {
	_ = "STUB: not implemented"
	return
}

func ResetReservationPhase() { _ = "STUB: not implemented"; return }

// RecordReservationResourceByTypeWithUnit records the resource record of a reservation as a metric.
func RecordReservationResourceByTypeWithUnit(name, resource, typ, unit string, value float64) {
	_ = "STUB: not implemented"
	return
}

func RecordElasticQuotaProcessLatency(operation string, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

func RecordSecondaryDeviceNotWellPlanned(nodeName string, notWellPlanned bool) {
	_ = "STUB: not implemented"
	return
}

// only for UT

func RecordNextPodPluginsDeletePodFromQueue(latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

func RecordElasticQuotaHookPluginLatency(plugin, operation string, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

func RecordJobPreemptionDuration(jobName string, result string, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

// GangJobSizeBucket maps a gang job size (number of attempted pods) to a
// bounded-cardinality label value, so it can be safely used as a Prometheus
// label without causing a cardinality explosion. Buckets are 100-wide up to
// 1000, with a single "1000+" catch-all above that.
func GangJobSizeBucket(n int) string { _ = "STUB: not implemented"; return "" }

// RecordGangScheduleCycleDuration records how long a gang scheduling cycle
// took, bucketed by clear reason and job size.
func RecordGangScheduleCycleDuration(reason, jobSize string, latency time.Duration) {
	_ = "STUB: not implemented"
	return
}
