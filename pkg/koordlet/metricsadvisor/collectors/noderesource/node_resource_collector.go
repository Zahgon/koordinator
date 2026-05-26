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

package noderesource

import (
	"time"

	"go.uber.org/atomic"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
)

const (
	CollectorName = "NodeResourceCollector"
)

var (
	timeNow = time.Now
)

// TODO more ut is needed for this plugin
type nodeResourceCollector struct {
	collectInterval time.Duration
	started         *atomic.Bool
	appendableDB    metriccache.Appendable
	metricDB        metriccache.MetricCache

	lastNodeCPUStat *framework.CPUStat

	sharedState      *framework.SharedState
	deviceCollectors map[string]framework.DeviceCollector
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

func (n *nodeResourceCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (n *nodeResourceCollector) Setup(c *framework.Context) { _ = "STUB: not implemented"; return }

func (n *nodeResourceCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Koordlet exit because of statesInformer sync failed.

func (n *nodeResourceCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (n *nodeResourceCollector) collectNodeResUsed() { _ = "STUB: not implemented"; return }

// get the accumulated cpu ticks

// NOTE: The collected memory usage is in kilobytes not bytes.

// 1 jiffy can be 10ms by default.
// NOTE: do subtraction and division first to avoid overflow

// update collect time

// in cpu cores
