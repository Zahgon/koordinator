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

package metriccache

import (
	"time"
)

type InterferenceMetricName string

type QueryParam struct {
	Aggregate AggregationType
	Start     *time.Time
	End       *time.Time
}

type AggregateInfo struct {
	// TODO only support node resource metric now
	MetricStart *time.Time
	MetricEnd   *time.Time

	MetricsCount int64
}

func (a *AggregateInfo) TimeRangeDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type QueryResult struct {
	AggregateInfo *AggregateInfo
	Error         error
}

func (q *QueryParam) FillDefaultValue() {
	_ = "STUB: not implemented"
	// todo, set start time as unix-zero if nil, set end as now if nil
	return
}

type MetricCache interface {
	Run(stopCh <-chan struct{}) error
	TSDBStorage
	KVStorage
}

type metricCache struct {
	config *Config
	TSDBStorage
	KVStorage
}

func NewMetricCache(cfg *Config) (MetricCache, error) {
	_ = "STUB: not implemented"
	return *new(MetricCache), nil
}

func (m *metricCache) Run(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }
