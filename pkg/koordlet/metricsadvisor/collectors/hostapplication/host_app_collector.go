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

package hostapplication

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
	"go.uber.org/atomic"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	CollectorName = "HostApplicationCollector"
)

var (
	timeNow = time.Now
)

type hostAppCollector struct {
	collectInterval time.Duration
	started         *atomic.Bool
	appendableDB    metriccache.Appendable
	statesInformer  statesinformer.StatesInformer
	cgroupReader    resourceexecutor.CgroupReader
	lastAppCPUStat  *gocache.Cache
	sharedState     *framework.SharedState
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

var _ framework.Collector = &hostAppCollector{}

func (h *hostAppCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (h *hostAppCollector) Setup(c *framework.Context) { _ = "STUB: not implemented"; return }

func (h *hostAppCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Koordlet exit because of statesInformer sync failed.

func (h *hostAppCollector) Started() bool { _ = "STUB: not implemented"; return false }

var (
	defaultMemoryCollectPolicy slov1alpha1.NodeMemoryCollectPolicy = slov1alpha1.UsageWithoutPageCache
)

func (h *hostAppCollector) collectHostAppResUsed() { _ = "STUB: not implemented"; return }

// sum memory usage according to NodeMemoryCollectPolicy
