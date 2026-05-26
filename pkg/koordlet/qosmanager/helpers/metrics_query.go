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

package helpers

import (
	"time"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

var (
	timeNow = time.Now
)

func CollectorNodeMetricLast(metricCache metriccache.MetricCache, queryMeta metriccache.MetricMeta, metricCollectInterval time.Duration) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func CollectNodeMetrics(metricCache metriccache.MetricCache, start, end time.Time, queryMeta metriccache.MetricMeta) (metriccache.AggregateResult, error) {
	_ = "STUB: not implemented"
	return *new(metriccache.AggregateResult), nil
}

func CollectAllHostAppMetricsLast(hostApps []slov1alpha1.HostApplicationSpec, metricCache metriccache.MetricCache,
	metricResource metriccache.MetricResource, metricCollectInterval time.Duration) map[string]float64 {
	_ = "STUB: not implemented"
	return nil
}

func CollectAllHostAppMetrics(hostApps []slov1alpha1.HostApplicationSpec, metricCache metriccache.MetricCache,
	queryParam metriccache.QueryParam, metricResource metriccache.MetricResource) map[string]float64 {
	_ = "STUB: not implemented"
	return nil
}

func CollectAllPodMetricsLast(statesInformer statesinformer.StatesInformer, metricCache metriccache.MetricCache,
	metricResource metriccache.MetricResource, metricCollectInterval time.Duration) map[string]float64 {
	_ = "STUB: not implemented"
	return nil
}

func CollectAllPodMetrics(statesInformer statesinformer.StatesInformer, metricCache metriccache.MetricCache,
	queryParam metriccache.QueryParam, metricResource metriccache.MetricResource) map[string]float64 {
	_ = "STUB: not implemented"
	return nil
}

func CollectPodMetricLast(metricCache metriccache.MetricCache, queryMeta metriccache.MetricMeta,
	metricCollectInterval time.Duration) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func CollectPodMetric(metricCache metriccache.MetricCache, queryMeta metriccache.MetricMeta, start, end time.Time) (metriccache.AggregateResult, error) {
	_ = "STUB: not implemented"
	return *new(metriccache.AggregateResult), nil
}

func CollectContainerResMetricLast(metricCache metriccache.MetricCache, queryMeta metriccache.MetricMeta,
	metricCollectInterval time.Duration) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func CollectContainerThrottledMetric(metricCache metriccache.MetricCache, containerID *string,
	metricCollectInterval time.Duration) (metriccache.AggregateResult, error) {
	_ = "STUB: not implemented"
	return *new(metriccache.AggregateResult), nil
}

func GenerateQueryParamsAvg(windowDuration time.Duration) *metriccache.QueryParam {
	_ = "STUB: not implemented"
	return nil
}

func GenerateQueryParamsLast(windowDuration time.Duration) *metriccache.QueryParam {
	_ = "STUB: not implemented"
	return nil
}

func Query(querier metriccache.Querier, resource metriccache.MetricResource, properties map[metriccache.MetricProperty]string) (metriccache.AggregateResult, error) {
	_ = "STUB: not implemented"
	return *new(metriccache.AggregateResult), nil
}
