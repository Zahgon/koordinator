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

package cpunormalization

import (
	corev1 "k8s.io/api/core/v1"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/slo-controller/noderesource/framework"
)

const PluginName = "CPUNormalization"

const (
	defaultRatioStr = "1.00"
	// in case of unexpected resource amplification
	// NOTE: Currently we do not support the scaling factor below 1.0.
	defaultMinRatio = 1.0
	defaultMaxRatio = 5.0
)

var (
	client     ctrlclient.Client
	cfgHandler *configHandler
)

type Plugin struct{}

func (p *Plugin) Name() string {
	_ = "STUB: not implemented"

	// +kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
	// +kubebuilder:rbac:groups=topology.node.k8s.io,resources=noderesourcetopologies,verbs=get;list;watch;create;update;patch;delete
	return ""
}

func (p *Plugin) Setup(opt *framework.Option) error { _ = "STUB: not implemented"; return nil }

// NeedSyncMeta checks if the node annotation of cpu normalization ratio to update is different from the current
func (p *Plugin) NeedSyncMeta(_ *configuration.ColocationStrategy, oldNode, newNode *corev1.Node) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// no annotation to set

// annotation to create

// annotation to remove

func (p *Plugin) Prepare(_ *configuration.ColocationStrategy, node *corev1.Node, nr *framework.NodeResource) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) Reset(node *corev1.Node, message string) []framework.ResourceItem {
	_ = "STUB: not implemented"
	// TBD: do we need to delete existing annotations/labels?
	return nil
}

// Calculate retrieves the CPUNormalizationStrategy from ConfigMap and the CPUBasicInfo from the NRT, get the cpu
// normalization ratio annotation
func (p *Plugin) Calculate(_ *configuration.ColocationStrategy, node *corev1.Node, _ *corev1.PodList, _ *framework.ResourceMetrics) ([]framework.ResourceItem, error) {
	_ = "STUB: not implemented"
	// If any of the necessary inputs like config, CPUBasicInfo is missing, abort to output and skip updating the
	// node annotation.
	return nil, nil
}

// node label take precedence to the strategy, <nodeEnabled, strategyEnabled>:
// - <true, strategyEnabled> -> true
// - <false, strategyEnabled> -> false
// - <nil, true> -> true
// - otherwise -> false

func isCPUBasicInfoChanged(infoOld, infoNew *extension.CPUBasicInfo) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func getCPUNormalizationRatio(info *extension.CPUBasicInfo, strategy *configuration.CPUNormalizationStrategy) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getCPUNormalizationRatioFromModel(info *extension.CPUBasicInfo, strategy *configuration.CPUNormalizationStrategy) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// HT=on, Turbo=on

// HT=on, Turbo=off

// HT=off, Turbo=on

// HT=off, Turbo=off

func isCPUNormalizationRatioValid(ratio float64) error { _ = "STUB: not implemented"; return nil }
