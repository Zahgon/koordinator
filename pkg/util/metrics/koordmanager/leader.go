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

package koordmanager

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"k8s.io/client-go/tools/leaderelection"
)

var leaderMetric = promauto.NewGauge(
	prometheus.GaugeOpts{
		Name: "koordinator_manager_is_leader",
		Help: "Gauge the leader of koordinator-manager",
	},
)

func init() {
	InternalMustRegister(leaderMetric)
	leaderelection.SetProvider(newMetricProvider())
}

func newMetricProvider() leaderelection.MetricsProvider {
	_ = "STUB: not implemented"
	return *new(leaderelection.MetricsProvider)
}

type metricsProvider struct{}
type switchMetric struct{}

var (
	_ leaderelection.MetricsProvider = metricsProvider{}
	_ leaderelection.LeaderMetric    = switchMetric{}
)

func (metricsProvider) NewLeaderMetric() leaderelection.LeaderMetric {
	_ = "STUB: not implemented"
	return *new(leaderelection.LeaderMetric)
}

func (switchMetric) On(_ string) { _ = "STUB: not implemented"; return }

func (s switchMetric) Off(string) { _ = "STUB: not implemented"; return }

func (switchMetric) SlowpathExercised(_ string) { _ = "STUB: not implemented"; return }
