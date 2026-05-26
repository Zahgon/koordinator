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

package podthrottled

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
	CollectorName = "PodThrottledCollector"
)

// TODO more ut is needed for this plugin
type podThrottledCollector struct {
	collectInterval time.Duration
	started         *atomic.Bool
	appendableDB    metriccache.Appendable
	statesInformer  statesinformer.StatesInformer
	cgroupReader    resourceexecutor.CgroupReader
	podFilter       framework.PodFilter

	lastPodCPUThrottled       *gocache.Cache
	lastContainerCPUThrottled *gocache.Cache
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

var _ framework.PodCollector = &podThrottledCollector{}

func (c *podThrottledCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

func (c *podThrottledCollector) Setup(ctx *framework.Context) { _ = "STUB: not implemented"; return }

func (c *podThrottledCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Koordlet exit because of statesInformer sync failed.

func (c *podThrottledCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (c *podThrottledCollector) FilterPod(meta *statesinformer.PodMeta) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func (c *podThrottledCollector) collectPodThrottledInfo() { _ = "STUB: not implemented"; return }

// types.UID

// print running pod collection error

// collect container-level metrics

// end for podMeta

func (c *podThrottledCollector) collectContainerThrottledInfo(podMeta *statesinformer.PodMeta) []metriccache.MetricSample {
	_ = "STUB: not implemented"
	return nil
}

// higher verbosity for probably non-running pods

// end for container status
