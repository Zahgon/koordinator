/*
Copyright 2023 The Koordinator Authors.
Copyright 2017 The Kubernetes Authors.

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

package metricsserver

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	metricsapi "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type metricsClientTestCase struct {
	snapshotTimestamp    time.Time
	snapshotWindow       time.Duration
	namespace            *corev1.Namespace
	pod1Snaps, pod2Snaps []*ContainerMetricsSnapshot
}

type containerID struct {
	namespace     string
	podname       string
	containerName string
}

func newMetricsClientTestCase() *metricsClientTestCase { _ = "STUB: not implemented"; return nil }

func newEmptyMetricsClientTestCase() *metricsClientTestCase { _ = "STUB: not implemented"; return nil }

func (tc *metricsClientTestCase) newContainerMetricsSnapshot(id containerID, cpuUsage int64, memUsage int64) *ContainerMetricsSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (tc *metricsClientTestCase) createFakeMetricsClient() MetricsClient {
	_ = "STUB: not implemented"
	return *new(MetricsClient)
}

func (tc *metricsClientTestCase) getFakePodMetricsList() *metricsapi.PodMetricsList {
	_ = "STUB: not implemented"
	return nil
}

func (tc *metricsClientTestCase) getFakePodMetric() *metricsapi.PodMetrics {
	_ = "STUB: not implemented"
	return nil
}

func makePodMetrics(snaps []*ContainerMetricsSnapshot) metricsapi.PodMetrics {
	_ = "STUB: not implemented"
	return *new(metricsapi.PodMetrics)
}

func (tc *metricsClientTestCase) getAllSnaps() []*ContainerMetricsSnapshot {
	_ = "STUB: not implemented"
	return nil
}
