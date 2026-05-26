//go:build linux
// +build linux

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

package performance

import (
	"time"

	"go.uber.org/atomic"
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/util/perf"
)

type performanceCollector struct {
	cpiCollectInterval        time.Duration
	psiCollectInterval        time.Duration
	collectTimeWindowDuration time.Duration

	started        *atomic.Bool
	statesInformer statesinformer.StatesInformer
	metricCache    metriccache.MetricCache
	cgroupReader   resourceexecutor.CgroupReader
}

func New(opt *framework.Options) framework.Collector {
	_ = "STUB: not implemented"
	return *new(framework.Collector)
}

func (p *performanceCollector) Enabled() bool {
	_ = "STUB: not implemented"
	// TODO: add tma analyze feature gate
	return false
}

func (p *performanceCollector) EnabledPerf() bool { _ = "STUB: not implemented"; return false }

func (p *performanceCollector) Setup(s *framework.Context) { _ = "STUB: not implemented"; return }

func (p *performanceCollector) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

// Koordlet exit because of statesInformer sync failed.

// init perf buffer pool for different perf group collectors

func (p *performanceCollector) getLibpfm4EventMap() map[int]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (p *performanceCollector) Started() bool { _ = "STUB: not implemented"; return false }

func (p *performanceCollector) collectContainerCPI() { _ = "STUB: not implemented"; return }

// get container CPI collectors for each container

// collect container cpi

// save container CPI metric to tsdb

func (p *performanceCollector) getAndStartCollectorOnSingleContainer(podParentCgroupDir string, containerStatus *corev1.ContainerStatus, number int32, events []string) (perf.Collector, error) {
	_ = "STUB: not implemented"
	return *new(perf.Collector), nil
}

func (p *performanceCollector) profileCPIOnSingleContainer(status *corev1.ContainerStatus, collectorOnSingleContainer perf.Collector, pod *corev1.Pod) []metriccache.MetricSample {
	_ = "STUB: not implemented"
	return nil
}

func (p *performanceCollector) collectContainerPSI() { _ = "STUB: not implemented"; return }

// save container's psi metrics to tsdb

func (p *performanceCollector) collectSingleContainerPSI(podParentCgroupDir string, containerStatus *corev1.ContainerStatus, pod *corev1.Pod) []metriccache.MetricSample {
	_ = "STUB: not implemented"
	return nil
}

func (p *performanceCollector) collectPodPSI() { _ = "STUB: not implemented"; return }

// save pod psi metrics to tsdb

func (p *performanceCollector) collectSinglePodPSI(pod *corev1.Pod, podCgroupDir string) []metriccache.MetricSample {
	_ = "STUB: not implemented"
	return nil
}

func (p *performanceCollector) collectPSI(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"
	// CgroupV1 psi collector support only on anolis os currently
	return
}

//skip collect psi when system not support

func (p *performanceCollector) saveMetric(samples []metriccache.MetricSample) error {
	_ = "STUB: not implemented"
	return nil
}
