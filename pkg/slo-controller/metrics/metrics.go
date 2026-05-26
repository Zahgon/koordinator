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

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	_ "k8s.io/component-base/metrics/prometheus/clientgo" // load restclient and workqueue metrics
)

const (
	SLOControllerSubsystem = "slo_controller"

	NodeKey = "node"

	StatusKey       = "status"
	StatusSucceeded = "succeeded"
	StatusFailed    = "failed"

	ReasonKey   = "reason"
	PluginKey   = "plugin"
	ResourceKey = "resource"

	UnitKey     = "unit"
	UnitCore    = "core"
	UnitByte    = "byte"
	UnitInteger = "integer"
)

func genNodeLabels(node *corev1.Node) prometheus.Labels {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels)
}

func withNodeStatusLabels(labels map[string]string, isSucceeded bool, reason string) prometheus.Labels {
	_ = "STUB: not implemented"
	return *new(prometheus.Labels)
}

func recordNodeCountMetric(vec *prometheus.CounterVec, isSucceeded bool, reason string) {
	_ = "STUB: not implemented"
	return
}
