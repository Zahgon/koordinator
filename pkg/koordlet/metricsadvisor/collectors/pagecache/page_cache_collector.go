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

package pagecache

import (
	"time"

	"go.uber.org/atomic"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	CollectorName = "PageCacheCollector"
)

var (
	timeNow = time.Now
)

type pageCacheCollector struct {
	collectInterval       time.Duration
	started               *atomic.Bool
	appendableDB          metriccache.Appendable
	metricDB              metriccache.MetricCache
	statesInformer        statesinformer.StatesInformer
	cgroupReader          resourceexecutor.CgroupReader
	podFilter             framework.PodFilter
	coldPageCollectorGate bool
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

func (p *pageCacheCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (p *pageCacheCollector) Setup(c *framework.Context) { _ = "STUB: not implemented"; return }

func (p *pageCacheCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (p *pageCacheCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (p *pageCacheCollector) FilterPod(meta *statesinformer.PodMeta) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (p *pageCacheCollector) collectPageCache() { _ = "STUB: not implemented"; return }

func (p *pageCacheCollector) collectNodePageCache() { _ = "STUB: not implemented"; return }

// NOTE: The collected memory usage is in kilobytes not bytes.

func (p *pageCacheCollector) collectPodPageCache() { _ = "STUB: not implemented"; return }

// types.UID

// higher verbosity for probably non-running pods

func (p *pageCacheCollector) collectContainerPageCache(meta *statesinformer.PodMeta) []metriccache.MetricSample {
	_ = "STUB: not implemented"
	return nil
}

// higher verbosity for probably non-running pods
