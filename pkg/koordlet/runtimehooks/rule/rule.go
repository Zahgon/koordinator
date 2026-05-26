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

package rule

import (
	"sync"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

func init() {
	globalHookRules = map[string]*Rule{}
}

type Rule struct {
	name            string
	description     string
	parseRuleType   statesinformer.RegisterType
	parseRuleFn     ParseRuleFn
	callbacks       []UpdateCbFn
	systemSupported bool
}

type ParseRuleFn func(interface{}) (bool, error)
type UpdateCbFn func(target *statesinformer.CallbackTarget) error
type SysSupportFn func() bool

var globalHookRules map[string]*Rule
var globalRWMutex sync.RWMutex

func Register(name, description string, injectOpts ...InjectOption) *Rule {
	_ = "STUB: not implemented"
	return nil
}

func (r *Rule) runUpdateCallbacks(target *statesinformer.CallbackTarget) {
	_ = "STUB: not implemented"
	return
}

func find(name string) (*Rule, bool) { _ = "STUB: not implemented"; return nil, false }

func UpdateRules(ruleType statesinformer.RegisterType, ruleObj interface{}, targets *statesinformer.CallbackTarget) {
	_ = "STUB: not implemented"
	return
}
