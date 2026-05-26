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

	promstorage "github.com/prometheus/prometheus/storage"
)

// MetricResult contains s set of series, it can also produce final result like aggregation value
type MetricResult interface {
	MetricMeta
	// AddSeries receives series and saves in MetricResult, which can be used for generating the final value
	AddSeries(promstorage.Series) error
}

// AggregateResultFactory generates AggregateResult according to MetricMeta
type AggregateResultFactory interface {
	New(meta MetricMeta) AggregateResult
}

var DefaultAggregateResultFactory AggregateResultFactory = &aggregateResultFactory{}

// aggregateResultFactory implements AggregateResultFactory
type aggregateResultFactory struct{}

func (f *aggregateResultFactory) New(meta MetricMeta) AggregateResult {
	_ = "STUB: not implemented"
	return *new(AggregateResult)
}

// AggregateResult inherits MetricResult, which can also generate value according to the give AggregationType
type AggregateResult interface {
	MetricResult
	Count() int
	Value(t AggregationType) (float64, error)
	TimeRangeDuration() time.Duration
}

var _ AggregateResult = &aggregateResult{}

func newAggregateResult(meta MetricMeta) AggregateResult {
	_ = "STUB: not implemented"
	return *new(AggregateResult)
}

// aggregateResult implements AggregateResult
type aggregateResult struct {
	metricKind       string
	metricProperties map[string]string
	points           []*Point
	metricStart      time.Time
	metricsEnd       time.Time
}

type AggregationType string

const (
	AggregationTypeAVG   AggregationType = "avg"
	AggregationTypeP99   AggregationType = "p99"
	AggregationTypeP95   AggregationType = "P95"
	AggregationTypeP90   AggregationType = "P90"
	AggregationTypeP50   AggregationType = "p50"
	AggregationTypeLast  AggregationType = "last"
	AggregationTypeCount AggregationType = "count"
)

// AggregateParam defines the field name of value and time in series struct
type AggregateParam struct {
	ValueFieldName string
	TimeFieldName  string
}

// AggregationFunc receives a list of series and generate the final value according to AggregateParam
type AggregationFunc func(interface{}, AggregateParam) (float64, error)

func (r *aggregateResult) AddSeries(series promstorage.Series) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *aggregateResult) GetKind() string { _ = "STUB: not implemented"; return "" }

func (r *aggregateResult) GetProperties() map[string]string { _ = "STUB: not implemented"; return nil }

// Count return the size of metric series
func (r *aggregateResult) Count() int { _ = "STUB: not implemented"; return 0 }

// Value returns the result by AggregationType
func (r *aggregateResult) Value(t AggregationType) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TimeRangeDuration returns the time duration of metric series
func (r *aggregateResult) TimeRangeDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

var pointsDefaultAggregateParam = AggregateParam{
	ValueFieldName: "Value",
	TimeFieldName:  "Timestamp",
}

func getAggregateFunc(aggregationType AggregationType) AggregationFunc {
	_ = "STUB: not implemented"
	return *new(AggregationFunc)
}
