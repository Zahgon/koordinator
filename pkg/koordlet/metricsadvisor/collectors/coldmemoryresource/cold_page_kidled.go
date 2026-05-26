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

package coldmemoryresource

import (
	"time"

	"go.uber.org/atomic"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

var (
	timeNow = time.Now
)

type kidledcoldPageCollector struct {
	collectInterval time.Duration
	started         *atomic.Bool
	cgroupReader    resourceexecutor.CgroupReader
	statesInformer  statesinformer.StatesInformer
	podFilter       framework.PodFilter
	appendableDB    metriccache.Appendable
	metricDB        metriccache.MetricCache // TODO remove redundant var
	coldBoundary    int
}

func (k *kidledcoldPageCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (k *kidledcoldPageCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (k *kidledcoldPageCollector) Enabled() bool { _ = "STUB: not implemented"; return false }

// start kidled

func (k *kidledcoldPageCollector) Setup(c1 *framework.Context) { _ = "STUB: not implemented"; return }

func (k *kidledcoldPageCollector) collectColdPageInfo() { _ = "STUB: not implemented"; return }

func (k *kidledcoldPageCollector) collectNodeColdPageInfo() ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kidledcoldPageCollector) collectPodsColdPageInfo() ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// types.UID

func (k *kidledcoldPageCollector) collectContainersColdPageInfo(meta *statesinformer.PodMeta) ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kidledcoldPageCollector) collectHostAppsColdPageInfo() ([]metriccache.MetricSample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *kidledcoldPageCollector) FilterPod(meta *statesinformer.PodMeta) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}
