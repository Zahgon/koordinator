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

package resctrl

import (
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	util "github.com/koordinator-sh/koordinator/pkg/koordlet/util/resctrl"
)

type UpdateFunc func(resource util.ResctrlUpdater) error

type DefaultResctrlProtocolUpdater struct {
	hooksProtocol protocol.HooksProtocol
	group         string
	schemata      string
	updateFunc    UpdateFunc
}

func (u DefaultResctrlProtocolUpdater) Name() string { _ = "STUB: not implemented"; return "" }

func (u DefaultResctrlProtocolUpdater) Key() string { _ = "STUB: not implemented"; return "" }

func (u DefaultResctrlProtocolUpdater) Value() string { _ = "STUB: not implemented"; return "" }

func (r *DefaultResctrlProtocolUpdater) SetKey(group string) { _ = "STUB: not implemented"; return }

func (r *DefaultResctrlProtocolUpdater) SetValue(schemata string) {
	_ = "STUB: not implemented"
	return
}

func (u *DefaultResctrlProtocolUpdater) Update() error { _ = "STUB: not implemented"; return nil }

type Updater func(u DefaultResctrlProtocolUpdater) error

func NewCreateResctrlProtocolUpdater(hooksProtocol protocol.HooksProtocol) util.ResctrlUpdater {
	_ = "STUB: not implemented"
	return *new(util.ResctrlUpdater)
}

func NewRemoveResctrlProtocolUpdater(hooksProtocol protocol.HooksProtocol) util.ResctrlUpdater {
	_ = "STUB: not implemented"
	return *new(util.ResctrlUpdater)
}

func NewRemoveResctrlUpdater(group string) util.ResctrlUpdater {
	_ = "STUB: not implemented"
	return *new(util.ResctrlUpdater)
}

func CreateResctrlProtocolUpdaterFunc(u util.ResctrlUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveResctrlProtocolUpdaterFunc(u util.ResctrlUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveResctrlUpdaterFunc(u util.ResctrlUpdater) error { _ = "STUB: not implemented"; return nil }
