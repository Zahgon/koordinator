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

package resctrl

import (
	"time"

	"go.uber.org/atomic"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	CollectorName = "resctrlCollector"
)

type resctrlCollector struct {
	collectInterval      time.Duration
	started              *atomic.Bool
	metricCache          metriccache.MetricCache
	statesInformer       statesinformer.StatesInformer
	resctrlReader        resourceexecutor.ResctrlReader
	resctrlCollectorGate bool
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

// 1. config enable resctrl collector
// 2. cmdline, os, cpuid enable resctrl collector
// 3. check CPU vendor(Intel&AMD)
// 4. check resctrl collector feature gate
func (r *resctrlCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (r *resctrlCollector) Setup(c *framework.Context) { _ = "STUB: not implemented"; return }

func (r *resctrlCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (r *resctrlCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (r *resctrlCollector) collectQoSResctrlStat() { _ = "STUB: not implemented"; return }

// use resctrl/schemata to check mount state TODO: use another method to check

// save QoS resctrl data to tsdb

func (r *resctrlCollector) saveMetric(samples []metriccache.MetricSample) error {
	_ = "STUB: not implemented"
	return nil
}
