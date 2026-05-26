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

package impl

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/utils/ptr"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	clientset "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned"
	clientsetv1alpha1 "github.com/koordinator-sh/koordinator/pkg/client/clientset/versioned/typed/slo/v1alpha1"
	listerv1alpha1 "github.com/koordinator-sh/koordinator/pkg/client/listers/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/prediction"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	koordletutil "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

const (
	nodeMetricInformerName PluginName = "nodeMetricInformer"

	// defaultAggregateDurationSeconds is the default metric aggregate duration by seconds
	minAggregateDurationSeconds     = 60
	defaultAggregateDurationSeconds = 300

	defaultReportIntervalSeconds = 60
	minReportIntervalSeconds     = 30

	// metric is valid only if its (lastSample.Time - firstSample.Time) > 0.5 * targetTimeRange
	// used during checking node aggregate usage for cold start
	validateTimeRangeRatio = 0.5
)

var (
	scheme                                                         = runtime.NewScheme()
	defaultMemoryCollectPolicy slov1alpha1.NodeMemoryCollectPolicy = slov1alpha1.UsageWithoutPageCache
	defaultNodeMetricSpec                                          = slov1alpha1.NodeMetricSpec{
		CollectPolicy: &slov1alpha1.NodeMetricCollectPolicy{
			AggregateDurationSeconds: ptr.To[int64](defaultAggregateDurationSeconds),
			ReportIntervalSeconds:    ptr.To[int64](defaultReportIntervalSeconds),
			NodeAggregatePolicy: &slov1alpha1.AggregatePolicy{
				Durations: []metav1.Duration{
					{Duration: 5 * time.Minute},
					{Duration: 10 * time.Minute},
					{Duration: 30 * time.Minute},
				},
			},
			NodeMemoryCollectPolicy: &defaultMemoryCollectPolicy,
		},
	}
	timeNow = time.Now
)

type nodeMetricInformer struct {
	reportEnabled      bool
	nodeName           string
	nodeMetricInformer cache.SharedIndexInformer
	nodeMetricLister   listerv1alpha1.NodeMetricLister
	eventRecorder      record.EventRecorder
	statusUpdater      *statusUpdater

	podsInformer     *podsInformer
	nodeInformer     *nodeInformer
	nodeSLOInformer  *nodeSLOInformer
	metricCache      metriccache.MetricCache
	predictorFactory prediction.PredictorFactory

	rwMutex    sync.RWMutex
	nodeMetric *slov1alpha1.NodeMetric
}

func NewNodeMetricInformer() *nodeMetricInformer { _ = "STUB: not implemented"; return nil }

func (r *nodeMetricInformer) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (r *nodeMetricInformer) Setup(ctx *PluginOption, state *PluginState) {
	_ = "STUB: not implemented"
	return
}

func (r *nodeMetricInformer) ReportEvent(object runtime.Object, eventType, reason, message string) {
	_ = "STUB: not implemented"
	return
}

func (r *nodeMetricInformer) Start(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (r *nodeMetricInformer) syncNodeMetricWorker(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	return
}

func (r *nodeMetricInformer) getNodeMetricReportInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *nodeMetricInformer) getNodeMetricAggregateDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *nodeMetricInformer) getNodeMetricSpec() *slov1alpha1.NodeMetricSpec {
	_ = "STUB: not implemented"
	return nil
}

func (r *nodeMetricInformer) sync() { _ = "STUB: not implemented"; return }

func newNodeMetricInformer(client clientset.Interface, nodeName string) cache.SharedIndexInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedIndexInformer)
}

func (r *nodeMetricInformer) isNodeMetricInited() bool { _ = "STUB: not implemented"; return false }

func (r *nodeMetricInformer) updateMetricSpec(newNodeMetric *slov1alpha1.NodeMetric) {
	_ = "STUB: not implemented"
	return
}

// generateQueryDuration generate query params. It assumes the nodeMetric is initialized
func (r *nodeMetricInformer) generateQueryDuration() (start time.Time, end time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time)
}

func (r *nodeMetricInformer) collectMetric() (*slov1alpha1.NodeMetricInfo, []*slov1alpha1.PodMetricInfo,
	[]*slov1alpha1.HostApplicationMetricInfo, *slov1alpha1.ReclaimableMetric) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// predict pods which have valid metrics; ignore prediction failures

// Todo: expose status in nodeMetrics

func (r *nodeMetricInformer) queryNodeMetric(start time.Time, end time.Time, aggregateType metriccache.AggregationType,
	coldStartFilter bool) slov1alpha1.ResourceMap {
	_ = "STUB: not implemented"
	return *new(slov1alpha1.ResourceMap)
}

func metricsInColdStart(queryStart, queryEnd time.Time, duration time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *nodeMetricInformer) collectNodeMetric(queryparam metriccache.QueryParam) (corev1.ResourceList, time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), *new(time.Duration), nil
}

// report usageMemoryWithoutPageCache

// report usageMemoryWithHotPageCache

// report usageMemoryWithPageCache

// degrade and apply default memory reporting policy: usageWithoutPageCache

func (r *nodeMetricInformer) collectNodeGPUMetric(queryparam metriccache.QueryParam, gpus koordletutil.GPUDevices) ([]schedulingv1alpha1.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: how to check the health status of GPU

func (r *nodeMetricInformer) collectNodeAggregateMetric(endTime time.Time, aggregatePolicy *slov1alpha1.AggregatePolicy) []slov1alpha1.AggregatedUsage {
	_ = "STUB: not implemented"
	return nil
}

func (r *nodeMetricInformer) collectSystemMetric(queryparam metriccache.QueryParam) (corev1.ResourceList, time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), *new(time.Duration), nil
}

func (r *nodeMetricInformer) querySystemMetric(start time.Time, end time.Time, aggregateType metriccache.AggregationType,
	coldStartFilter bool) slov1alpha1.ResourceMap {
	_ = "STUB: not implemented"
	return *new(slov1alpha1.ResourceMap)
}

func (r *nodeMetricInformer) collectSystemAggregateMetric(endTime time.Time, aggregatePolicy *slov1alpha1.AggregatePolicy) []slov1alpha1.AggregatedUsage {
	_ = "STUB: not implemented"
	return nil
}

func (r *nodeMetricInformer) collectPodMetric(podMeta *statesinformer.PodMeta, queryParam metriccache.QueryParam) (*slov1alpha1.PodMetricInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// slov1alpha1.UsageWithoutPageCache

func (r *nodeMetricInformer) collectHostAppMetric(hostApp *slov1alpha1.HostApplicationSpec, queryParam metriccache.QueryParam) (*slov1alpha1.HostApplicationMetricInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *nodeMetricInformer) collectPodGPUMetric(queryparam metriccache.QueryParam, uid string, gpus koordletutil.GPUDevices) ([]schedulingv1alpha1.DeviceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: how to check the health status of GPU

func (r *nodeMetricInformer) fillGPUMetrics(queryparam metriccache.QueryParam, info *slov1alpha1.PodMetricInfo, uid string, gpus koordletutil.GPUDevices) {
	_ = "STUB: not implemented"
	return
}

const (
	statusUpdateQPS   = 0.1
	statusUpdateBurst = 2
)

type statusUpdater struct {
	nodeMetricClient  clientsetv1alpha1.NodeMetricInterface
	previousTimestamp time.Time
	rateLimiter       *rate.Limiter
}

func newStatusUpdater(nodeMetricClient clientsetv1alpha1.NodeMetricInterface) *statusUpdater {
	_ = "STUB: not implemented"
	return nil
}

func (su *statusUpdater) updateStatus(nodeMetric *slov1alpha1.NodeMetric, newStatus *slov1alpha1.NodeMetricStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func doQuery(querier metriccache.Querier, resource metriccache.MetricResource, properties map[metriccache.MetricProperty]string) (metriccache.AggregateResult, error) {
	_ = "STUB: not implemented"
	return *new(metriccache.AggregateResult), nil
}
