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

package resourceamplification

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/noderesource/framework"
)

const PluginName = "ResourceAmplification"

var (
	cfgHandler *configHandler
)

// Plugin calculates and updates final node resource amplification ratios automatically
// based on user config and node cpu normalization ratio.
type Plugin struct{}

func (p *Plugin) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Plugin) NeedSyncMeta(_ *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
func (p *Plugin) Setup(opt *framework.Option) error { _ = "STUB: not implemented"; return nil }

func (p *Plugin) Prepare(_ *configuration.ColocationStrategy, node *corev1.Node, nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Reset(node *corev1.Node, message string) []framework.ResourceItem {
	_ = "STUB: not implemented"
	// Currently we have no user configurations so there's no need to reset.
	return nil
}

// Calculate calculates resource amplification ratio, resource final ratio should >= 1
// cpuAmplificationRatio = amplificationStrategy.resourceAmplificationRatio["cpu"] * CPUNormalizationRatio
// otherResourceAmplificationRatio = amplificationStrategy.resourceAmplificationRatio["other-resource"]
func (p *Plugin) Calculate(_ *configuration.ColocationStrategy, node *corev1.Node, _ *corev1.PodList, _ *framework.ResourceMetrics) ([]framework.ResourceItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func calculateCPUAmplificationRatio(node *corev1.Node, ratio extension.Ratio) (extension.Ratio, error) {
	_ = "STUB: not implemented"
	return *new(extension.Ratio), nil
}

func validateResourceAmplificationRatios(ratios map[corev1.ResourceName]extension.Ratio) error {
	_ = "STUB: not implemented"
	return nil
}

func getResourceAmplificationRatios(node *corev1.Node) (map[corev1.ResourceName]extension.Ratio, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
