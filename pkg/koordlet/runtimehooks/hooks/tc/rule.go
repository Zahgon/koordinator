//go:build linux
// +build linux

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

package tc

import (
	"k8s.io/apimachinery/pkg/types"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type tcRule struct {
	enable      bool
	netCfg      *NetQosGlobalConfig
	speed       uint64
	uidToHandle map[types.UID]uint32
	handleToUid map[uint32]types.UID
}

func newRule() *tcRule { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) getRule() *tcRule { _ = "STUB: not implemented"; return nil }

func (p *tcPlugin) updateRule(newRule *tcRule) bool { _ = "STUB: not implemented"; return false }

func (p *tcPlugin) parseRuleForNodeSLO(mergedNodeSLOIf interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// default policy enables

func (p *tcPlugin) parseForAllPods(e interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (p *tcPlugin) ruleUpdateCbForNodeSlo(target *statesinformer.CallbackTarget) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *tcPlugin) ruleUpdateCbForPod(target *statesinformer.CallbackTarget) error {
	_ = "STUB: not implemented"
	return nil
}

// cache classId from net_cls cgroup when first sync.

// mark the class that has been used

// netClsId is a decimal number.
