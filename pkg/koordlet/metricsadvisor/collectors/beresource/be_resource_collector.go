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

package beresource

import (
	"time"

	"go.uber.org/atomic"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	CollectorName = "BEResourceCollector"
)

type beResourceCollector struct {
	collectInterval time.Duration
	started         *atomic.Bool
	metricCache     metriccache.MetricCache
	statesInformer  statesinformer.StatesInformer
	cgroupReader    resourceexecutor.CgroupReader

	lastBECPUStat *framework.CPUStat
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

func (b *beResourceCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (b *beResourceCollector) Setup(s *framework.Context) { _ = "STUB: not implemented"; return }

func (b *beResourceCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Koordlet exit because of statesInformer sync failed.

func (b *beResourceCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (b *beResourceCollector) collectBECPUResourceMetric() { _ = "STUB: not implemented"; return }

func (b *beResourceCollector) getBECPURealMilliLimit() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// -1 means not suppress by cfs_quota

func (b *beResourceCollector) getBECPURequestMilliCores() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (b *beResourceCollector) getBECPUUsageMilliCores() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// NOTE:
// 1. do subtraction and division first to avoid overflow.
// 2. To solve the problem of insufficient precision of nanosecond floating-point numbers, convert to Milliseconds

// 1.0 CPU = 1000 Milli-CPU
// cpuUsageCores := resource.NewMilliQuantity(int64(cpuUsageValue*1000), resource.DecimalSI)
