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

package frameworkext

import (
	fwktype "k8s.io/kube-scheduler/framework"
)

var (
	ControllerPlugins = []string{"*"}
)

func isControllerPluginEnabled(pluginName string) bool { _ = "STUB: not implemented"; return false }

func getPluginControllerNames(pluginControllers map[string]Controller) []string {
	_ = "STUB: not implemented"
	return nil
}

type ControllerProvider interface {
	NewControllers() ([]Controller, error)
}

type Controller interface {
	Start()
	Name() string
}

type ControllersMap struct {
	controllers map[string]map[string]Controller
}

func NewControllersMap() *ControllersMap { _ = "STUB: not implemented"; return nil }

func (cm *ControllersMap) RegisterControllers(plugin fwktype.Plugin, profileName string) {
	_ = "STUB: not implemented"
	return
}

func (cm *ControllersMap) Start() { _ = "STUB: not implemented"; return }
