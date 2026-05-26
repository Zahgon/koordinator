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

package metricsadvisor

import (
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metricsadvisor/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type MetricAdvisor interface {
	Run(stopCh <-chan struct{}) error
	HasSynced() bool
}

type metricAdvisor struct {
	options *framework.Options
	context *framework.Context
}

func NewMetricAdvisor(cfg *framework.Config, statesInformer statesinformer.StatesInformer, metricCache metriccache.MetricCache) MetricAdvisor {
	_ = "STUB: not implemented"
	return *new(MetricAdvisor)
}

func (m *metricAdvisor) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (m *metricAdvisor) Run(stopCh <-chan struct{}) error { _ = "STUB: not implemented"; return nil }

func (m *metricAdvisor) setup() { _ = "STUB: not implemented"; return }

func (m *metricAdvisor) shutdown() { _ = "STUB: not implemented"; return }
