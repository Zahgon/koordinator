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

package nodeslo

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
)

func getResourceThresholdSpec(node *corev1.Node, cfg *configuration.ResourceThresholdCfg) (*slov1alpha1.ResourceThresholdStrategy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getResourceQOSSpec(node *corev1.Node, cfg *configuration.ResourceQOSCfg) (*slov1alpha1.ResourceQOSStrategy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCPUBurstConfigSpec(node *corev1.Node, cfg *configuration.CPUBurstCfg) (*slov1alpha1.CPUBurstStrategy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSystemConfigSpec(node *corev1.Node, cfg *configuration.SystemCfg) (*slov1alpha1.SystemStrategy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find strategy that matches current node.

// If there is no matched node strategy, use cluster strategy.

// Check whether node bandwidth is specified on current node, which takes HIGHER priority
// than ones configured in cluster strategy or node strategy.
// Error is returned if failed to parse node total bandwidth from annotation, which is not
// supposed to happen because we will check the validity of the annotation value in node
// plugins of validating webhook.

func getHostApplicationConfig(node *corev1.Node, cfg *configuration.HostApplicationCfg) ([]slov1alpha1.HostApplicationSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func calculateResourceThresholdCfgMerged(oldCfg configuration.ResourceThresholdCfg, configMap *corev1.ConfigMap) (configuration.ResourceThresholdCfg, error) {
	_ = "STUB: not implemented"
	return *new(configuration.ResourceThresholdCfg), nil
}

// merge ClusterStrategy

// merge with clusterStrategy

func calculateResourceQOSCfgMerged(oldCfg configuration.ResourceQOSCfg, configMap *corev1.ConfigMap) (configuration.ResourceQOSCfg, error) {
	_ = "STUB: not implemented"
	return *new(configuration.ResourceQOSCfg), nil
}

// merge ClusterStrategy

// merge with clusterStrategy

func calculateCPUBurstCfgMerged(oldCfg configuration.CPUBurstCfg, configMap *corev1.ConfigMap) (configuration.CPUBurstCfg, error) {
	_ = "STUB: not implemented"
	return *new(configuration.CPUBurstCfg), nil
}

// merge ClusterStrategy

// merge with clusterStrategy

func calculateSystemConfigMerged(oldCfg configuration.SystemCfg, configMap *corev1.ConfigMap) (configuration.SystemCfg, error) {
	_ = "STUB: not implemented"
	return *new(configuration.SystemCfg), nil
}

// merge ClusterStrategy

// merge with clusterStrategy

func calculateHostAppConfigMerged(oldCfg configuration.HostApplicationCfg, configMap *corev1.ConfigMap) (configuration.HostApplicationCfg, error) {
	_ = "STUB: not implemented"
	return *new(configuration.HostApplicationCfg), nil
}
