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
	"sync"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const ratioDiffEpsilon = 0.01

type Rule struct {
	lock     sync.RWMutex
	enable   bool
	curRatio float64
}

func newRule() *Rule { _ = "STUB: not implemented"; return nil }

func (r *Rule) IsEnabled() bool { _ = "STUB: not implemented"; return false }

func (r *Rule) GetCPUNormalizationRatio() float64 { _ = "STUB: not implemented"; return 0 }

func (r *Rule) UpdateRule(ratio float64) bool { _ = "STUB: not implemented"; return false }

func (p *Plugin) parseRule(nodeIf interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *Plugin) ruleUpdateCb(target *statesinformer.CallbackTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: if the ratio becomes bigger, scale top down, otherwise, scale bottom up

// pod-level

// container-level

// ignore sandbox containers

// NOTE: Update cgroups by the level since some resources like cfs quota requires the upper level value is no less
//       than the lower.
