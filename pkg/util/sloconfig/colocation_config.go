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

package sloconfig

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/configuration"
)

func NewDefaultColocationCfg() *configuration.ColocationCfg { _ = "STUB: not implemented"; return nil }

func DefaultColocationCfg() configuration.ColocationCfg {
	_ = "STUB: not implemented"
	return *new(configuration.ColocationCfg)
}

func DefaultColocationStrategy() configuration.ColocationStrategy {
	_ = "STUB: not implemented"
	return *new(configuration.ColocationStrategy)
}

func IsColocationStrategyValid(strategy *configuration.ColocationStrategy) bool {
	_ = "STUB: not implemented"
	return false
}

func IsNodeColocationCfgValid(nodeCfg *configuration.NodeColocationCfg) bool {
	_ = "STUB: not implemented"
	return false
}

// node colocation should not be empty

func GetNodeColocationStrategy(cfg *configuration.ColocationCfg, node *corev1.Node) *configuration.ColocationStrategy {
	_ = "STUB: not implemented"
	return nil
}

// update strategy according to node metadata

func UpdateColocationStrategyForNode(strategy *configuration.ColocationStrategy, node *corev1.Node) {
	_ = "STUB: not implemented"
	return
}

// GetColocationStrategyOnNode gets the colocation strategy in the node annotations.
func GetColocationStrategyOnNode(node *corev1.Node) (*configuration.ColocationStrategy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNodeReclaimPercent(node *corev1.Node, key string) *int64 {
	_ = "STUB: not implemented"
	return nil
}
