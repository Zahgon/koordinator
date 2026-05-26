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

package sysresource

import (
	"time"

	"go.uber.org/atomic"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
)

const (
	CollectorName = "SystemResourceCollector"
)

var (
	timeNow = time.Now
)

type systemResourceCollector struct {
	collectInterval  time.Duration
	outdatedInterval time.Duration
	started          *atomic.Bool
	appendableDB     metriccache.Appendable
	sharedState      *framework.SharedState
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

func (s *systemResourceCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (s *systemResourceCollector) Setup(c *framework.Context) { _ = "STUB: not implemented"; return }

func (s *systemResourceCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (s *systemResourceCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (s *systemResourceCollector) collectSysResUsed() { _ = "STUB: not implemented"; return }

// get node resource usage

// get all pod resource usage

// get all host application resource usage

// calculate system resource usage

// commit metric sample

func (s *systemResourceCollector) getAllPodsResourceUsage() (cpuCore float64, memory float64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
