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
)

const (
	DefaultExpireTime = 5 * time.Minute
	DefaultGCInterval = 1 * time.Minute
)

var defaultMetricGC = NewMetricGC(DefaultExpireTime, DefaultGCInterval)

type GCGaugeVec struct {
	name         string
	vec          *prometheus.GaugeVec
	expireStatus MetricGC
}

func NewGCGaugeVec(name string, vec *prometheus.GaugeVec) *GCGaugeVec {
	_ = "STUB: not implemented"
	return nil
}

func newGCGaugeVec(name string, vec *prometheus.GaugeVec, metricGC MetricGC) *GCGaugeVec {
	_ = "STUB: not implemented"
	return nil
}

func (g *GCGaugeVec) GetGaugeVec() *prometheus.GaugeVec { _ = "STUB: not implemented"; return nil }

func (g *GCGaugeVec) WithSet(labels prometheus.Labels, value float64) {
	_ = "STUB: not implemented"
	return
}

func (g *GCGaugeVec) Delete(labels prometheus.Labels) { _ = "STUB: not implemented"; return }

type GCCounterVec struct {
	name         string
	vec          *prometheus.CounterVec
	expireStatus MetricGC
}

func NewGCCounterVec(name string, vec *prometheus.CounterVec) *GCCounterVec {
	_ = "STUB: not implemented"
	return nil
}

func newGCCounterVec(name string, vec *prometheus.CounterVec, metricGC MetricGC) *GCCounterVec {
	_ = "STUB: not implemented"
	return nil
}

func (g *GCCounterVec) GetCounterVec() *prometheus.CounterVec {
	_ = "STUB: not implemented"
	return nil
}

func (g *GCCounterVec) WithInc(labels prometheus.Labels) { _ = "STUB: not implemented"; return }

func (g *GCCounterVec) Delete(labels prometheus.Labels) { _ = "STUB: not implemented"; return }

// GCHistogramVec wraps a prometheus.HistogramVec and integrates with MetricGC for expiration handling.
type GCHistogramVec struct {
	name         string
	vec          *prometheus.HistogramVec
	expireStatus MetricGC
}

func NewGCHistogramVec(name string, vec *prometheus.HistogramVec) *GCHistogramVec {
	_ = "STUB: not implemented"
	return nil
}

func newGCHistogramVec(name string, vec *prometheus.HistogramVec, metricGC MetricGC) *GCHistogramVec {
	_ = "STUB: not implemented"
	return nil
}

func (g *GCHistogramVec) GetHistogramVec() *prometheus.HistogramVec {
	_ = "STUB: not implemented"

	// WithObserve records a value in the histogram and updates the expiration status.
	return nil
}

func (g *GCHistogramVec) WithObserve(labels prometheus.Labels, value float64) {
	_ = "STUB: not implemented"
	return
}

// Delete removes the metric with the given labels and updates the expiration status.
func (g *GCHistogramVec) Delete(labels prometheus.Labels) { _ = "STUB: not implemented"; return }

type MetricVecGC interface {
	// Len returns the length of the alive metric statuses.
	Len() int
	// UpdateStatus updates the metric status with the given label values and timestamp (Unix seconds).
	UpdateStatus(updateTime int64, labels prometheus.Labels)
	// RemoveStatus removes the metric status with the given label values.
	RemoveStatus(labels prometheus.Labels)
	// ExpireMetrics expires all metric statuses which are updated before the expired time (Unix seconds).
	ExpireMetrics(expireTime int64) int
}

// record metric last updateTime and then can expire by updateTime
type metricVecGC struct {
	lock      sync.Mutex
	name      string
	metricVec *prometheus.MetricVec
	statuses  map[string]metricStatus // label values -> timestamp
}

type metricStatus struct {
	Labels          prometheus.Labels
	lastUpdatedUnix int64
}

func NewMetricVecGC(name string, metricVec *prometheus.MetricVec) MetricVecGC {
	_ = "STUB: not implemented"
	return *new(MetricVecGC)
}

func (v *metricVecGC) Len() int { _ = "STUB: not implemented"; return 0 }

func (v *metricVecGC) UpdateStatus(updateTime int64, labels prometheus.Labels) {
	_ = "STUB: not implemented"
	return
}

func (v *metricVecGC) updateStatus(statusKey string, status *metricStatus) {
	_ = "STUB: not implemented"
	return
}

func (v *metricVecGC) RemoveStatus(labels prometheus.Labels) { _ = "STUB: not implemented"; return }

func (v *metricVecGC) removeStatus(statusKey string) { _ = "STUB: not implemented"; return }

func (v *metricVecGC) ExpireMetrics(expireTime int64) int { _ = "STUB: not implemented"; return 0 }

type MetricGC interface {
	Run()
	Stop()
	AddMetric(name string, metric *prometheus.MetricVec)
	UpdateStatus(name string, labels prometheus.Labels)
	RemoveStatus(name string, labels prometheus.Labels)
	CountStatus(name string) int
}

type metricGC struct {
	globalLock sync.RWMutex
	metrics    map[string]MetricVecGC // metric name -> metricVecGC

	// expire time for metrics
	expireTime time.Duration
	// gc interval
	interval time.Duration

	stopCh chan struct{}
}

func NewMetricGC(expireTime time.Duration, interval time.Duration) MetricGC {
	_ = "STUB: not implemented"
	return *new(MetricGC)
}

func (e *metricGC) AddMetric(metricName string, metric *prometheus.MetricVec) {
	_ = "STUB: not implemented"
	return
}

func (e *metricGC) addMetric(metricName string, metric *prometheus.MetricVec) {
	_ = "STUB: not implemented"
	return
}

func (e *metricGC) Run() { _ = "STUB: not implemented"; return }

func (e *metricGC) run() { _ = "STUB: not implemented"; return }

func (e *metricGC) Stop() { _ = "STUB: not implemented"; return }

func (e *metricGC) UpdateStatus(metricName string, labels prometheus.Labels) {
	_ = "STUB: not implemented"
	return

	// different metric vectors can update simultaneously
}

func (e *metricGC) updateStatus(updateTime int64, metricName string, labels prometheus.Labels) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricGC) RemoveStatus(metricName string, labels prometheus.Labels) {
	_ = "STUB: not implemented"
	return
}

func (e *metricGC) removeStatus(metricName string, labels prometheus.Labels) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricGC) CountStatus(metricName string) int { _ = "STUB: not implemented"; return 0 }

func (e *metricGC) statusLen() int { _ = "STUB: not implemented"; return 0 }

func (e *metricGC) expire() error { _ = "STUB: not implemented"; return nil }

// labelsToKey generate a key for a metric with the given metric label pairs.
// NOTE: It assumes that the label keys of a metric vector are fixed.
// pattern: ${name}:${key1},${key2},...
func labelsToKey(labels prometheus.Labels) string { _ = "STUB: not implemented"; return "" }
