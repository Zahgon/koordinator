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

var _ ConfigChecker = &CPUBurstChecker{}

type CPUBurstChecker struct {
	cfg *configuration.CPUBurstCfg
	CommonChecker
}

func NewCPUBurstChecker(oldConfig, newConfig *corev1.ConfigMap, needUnmarshal bool) *CPUBurstChecker {
	_ = "STUB: not implemented"
	return nil
}

func (c *CPUBurstChecker) ConfigParamValid() error { _ = "STUB: not implemented"; return nil }

func (c *CPUBurstChecker) initConfig() error { _ = "STUB: not implemented"; return nil }

func (c *CPUBurstChecker) getConfigProfiles() []configuration.NodeCfgProfile {
	_ = "STUB: not implemented"
	return nil
}
