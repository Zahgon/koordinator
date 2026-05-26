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
	"k8s.io/apimachinery/pkg/labels"

	"github.com/koordinator-sh/koordinator/apis/configuration"
)

const InitSuccess = "Success"
const NotInit = "NotInit"

type ConfigChecker interface {
	IsCfgNotEmptyAndChanged() bool
	InitStatus() string
	ConfigParamValid() error
	NodeConfigProfileChecker
}

type NodeConfigProfileChecker interface {
	HasMultiNodeConfigs() bool                 //nodeConfigs nums > 1
	ProfileParamValid() error                  //name must not conflict
	NodeSelectorOverlap() error                //check nodeselector if overlap
	ExistNodeConflict(node *corev1.Node) error //check a node if exist conflict
}

type CommonChecker struct {
	configKey    string
	OldConfigMap *corev1.ConfigMap
	NewConfigMap *corev1.ConfigMap

	initStatus string
	NodeConfigProfileChecker
}

func (c *CommonChecker) IsCfgNotEmptyAndChanged() bool { _ = "STUB: not implemented"; return false }

func (c *CommonChecker) InitStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CommonChecker) CheckByValidator(config interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type nodeConfigProfileChecker struct {
	cfgName     string
	nodeConfigs []profileCheckInfo
}

type profileCheckInfo struct {
	profile   configuration.NodeCfgProfile
	selectors labels.Selector
}

func CreateNodeConfigProfileChecker(configName string, profiles func() []configuration.NodeCfgProfile) (NodeConfigProfileChecker, error) {
	_ = "STUB: not implemented"
	return *new(NodeConfigProfileChecker), nil
}

//checkNodeSelector empty

func (n *nodeConfigProfileChecker) HasMultiNodeConfigs() bool {
	_ = "STUB: not implemented"
	return false
}

func (n *nodeConfigProfileChecker) ProfileParamValid() error { _ = "STUB: not implemented"; return nil }

func (n *nodeConfigProfileChecker) checkName() error { _ = "STUB: not implemented"; return nil }

//checkName

/*
NodeSelectorOverlap detect overlap by nodeSelector before checkConflict By node.
example: config1.nodeSelector{aa=true}, config2.nodeSelector{aa=true,bb=true} ,config1 contains config2
*/
func (n *nodeConfigProfileChecker) NodeSelectorOverlap() error {
	_ = "STUB: not implemented"
	return nil
}

/*
ExistNodeConflict checkConflict By node.
example: config1.nodeSelector{aa=true}, config2.nodeSelector{aa=true,bb=true} conflict with node have label{aa=true,bb=true}
*/
func (n *nodeConfigProfileChecker) ExistNodeConflict(node *corev1.Node) error {
	_ = "STUB: not implemented"
	return nil
}
