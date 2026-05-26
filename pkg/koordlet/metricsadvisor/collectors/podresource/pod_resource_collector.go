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

package podresource

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
	"go.uber.org/atomic"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	CollectorName = "PodResourceCollector"
)

type podResourceCollector struct {
	collectInterval      time.Duration
	started              *atomic.Bool
	appendableDB         metriccache.Appendable
	statesInformer       statesinformer.StatesInformer
	cgroupReader         resourceexecutor.CgroupReader
	podFilter            framework.PodFilter
	lastPodCPUStat       *gocache.Cache
	lastContainerCPUStat *gocache.Cache

	deviceCollectors map[string]framework.DeviceCollector
	sharedState      *framework.SharedState
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

var _ framework.PodCollector = &podResourceCollector{}

func (p *podResourceCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (p *podResourceCollector) Setup(c *framework.Context) { _ = "STUB: not implemented"; return }

func (p *podResourceCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Koordlet exit because of statesInformer sync failed.

func (p *podResourceCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (p *podResourceCollector) FilterPod(meta *statesinformer.PodMeta) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (p *podResourceCollector) collectPodResUsed() { _ = "STUB: not implemented"; return }

// types.UID

// higher verbosity for probably non-running pods

// do subtraction and division first to avoid overflow

// update collect time

func (p *podResourceCollector) collectContainerResUsed(meta *statesinformer.PodMeta) []metriccache.MetricSample {
	_ = "STUB: not implemented"
	return nil
}

// higher verbosity for probably non-running pods

// do subtraction and division first to avoid overflow
