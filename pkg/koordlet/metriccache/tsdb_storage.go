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
	"github.com/prometheus/prometheus/tsdb"
)

// TSDBStorage defines time-series type DB, providing insert and query interface
type TSDBStorage interface {
	Appendable
	Queryable

	// Close closes the storage and all its underlying resources.
	Close() error
}

// Appendable allows creating appenders.
type Appendable interface {
	// Appender returns a new appender for the storage.
	Appender() Appender
}

// Appender provides batch appends of MetricSample against a storage
type Appender interface {
	// Append adds a list of MetricSample for the given series.
	Append(s []MetricSample) error
	// Commit submits the MetricSample and purges the batch. If Commit
	// returns a non-nil error, it also rolls back all modifications made in
	// the appender so far, as Rollback would do. In any case, an Appender
	// must not be used anymore after Commit has been called.
	Commit() error
}

// Queryable handles queries against a storage.
type Queryable interface {
	// Querier returns a new querier over the data partition for the given time range.
	Querier(startTime, endTime time.Time) (Querier, error)
}

// Querier provides querying access over time series data of a fixed time range.
type Querier interface {
	// Query add series to MetricResult that matches the given meta.
	// It allows passing hints that can help in optimizing select,
	// but it's up to implementation of MetricResult how to use it.
	Query(meta MetricMeta, hints *QueryHints, result MetricResult) error

	// QueryAndClose is same with Query, but close the querier after return the result
	QueryAndClose(meta MetricMeta, hints *QueryHints, result MetricResult) error

	// Close closes the querier and releases its resources.
	Close()
}

// SelectHints specifies hints passed for query.
// It is only an option for implementation of MetricResult to use, e.g. GroupedResult
type QueryHints struct {
	// GroupBy []string
}

var _ TSDBStorage = &tsdbStorage{}

// tsdbStorage implements TSDBStorage
type tsdbStorage struct {
	db *tsdb.DB
}

func (t *tsdbStorage) Appender() Appender { _ = "STUB: not implemented"; return *new(Appender) }

func (t *tsdbStorage) Querier(startTime, endTime time.Time) (Querier, error) {
	_ = "STUB: not implemented"
	return *new(Querier), nil
}

func (t *tsdbStorage) Close() error { _ = "STUB: not implemented"; return nil }

func NewTSDBStorage(conf *Config) (TSDBStorage, error) {
	_ = "STUB: not implemented"
	return *new(TSDBStorage), nil
}

// avoid conflicts using prometheus.tsdb v0.39 or higher
// prometheus.tsdb(0.37) requires all sample must following the time order
// new sample >= TSDB.MaxTime, out of order sample could not be appended until v0.39 with outOfOrderTimeWindow
// option enabled.
// oooTimeWindow must follow the grain of metric series

var _ Appender = &tsdbAppender{}

// tsdbAppender implements Appender
type tsdbAppender struct {
	appender promstorage.Appender
}

func (t *tsdbAppender) Append(samples []MetricSample) error { _ = "STUB: not implemented"; return nil }

// TODO cache the seriesRef to accelerate calls

func (t *tsdbAppender) Commit() error { _ = "STUB: not implemented"; return nil }

var _ Querier = &tsdbQuerier{}

// tsdbQuerier implements Querier
type tsdbQuerier struct {
	querier promstorage.Querier
}

func (t *tsdbQuerier) Query(meta MetricMeta, hints *QueryHints, result MetricResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tsdbQuerier) QueryAndClose(meta MetricMeta, hints *QueryHints, result MetricResult) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tsdbQuerier) Close() { _ = "STUB: not implemented"; return }
