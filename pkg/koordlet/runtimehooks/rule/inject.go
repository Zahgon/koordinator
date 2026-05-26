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
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type InjectOption interface {
	Apply(interface{}) error
}

type funcInject struct {
	f func(interface{}) error
}

func (fi *funcInject) Apply(o interface{}) error { _ = "STUB: not implemented"; return nil }

func NewFuncInject(f func(interface{}) error) *funcInject { _ = "STUB: not implemented"; return nil }

func WithParseFunc(t statesinformer.RegisterType, parseFunc ParseRuleFn) InjectOption {
	_ = "STUB: not implemented"
	return *new(InjectOption)
}

func WithUpdateCallback(updateCb UpdateCbFn) InjectOption {
	_ = "STUB: not implemented"
	return *new(InjectOption)
}

func WithSystemSupported(sysSupportFn SysSupportFn) InjectOption {
	_ = "STUB: not implemented"
	return *new(InjectOption)
}
