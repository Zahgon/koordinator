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
	"sync"

	apiext "github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
	sysutil "github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

const (
	// Max memory bandwidth for AMD CPU, Gb/s, since the extreme limit is hard to reach, we set a discount by 0.8
	// TODO The max memory bandwidth varies across SKU, so koordlet should be aware of the maximum automatically,
	// or support an configuration list.
	// Currently, the value is measured on "AMD EPYC(TM) MILAN"

	AMDCCDMaxMBGbps = 25 * 8 * 0.8

	// the AMD CPU use 2048 to express the unlimited memory bandwidth
	AMDCCDUnlimitedMB = 2048
)

type ResctrlUpdater interface {
	Name() string
	Key() string
	Value() string
	Update() error
	SetKey(key string)
	SetValue(key string)
}

const ClosdIdPrefix = "koordlet-"

type App struct {
	Resctrl    *sysutil.ResctrlSchemataRaw
	Closid     string
	Annotation string
}

type ResctrlEngine interface {
	Rebuild()
	RegisterApp(podid, annotation string, fromNRI bool, updater ResctrlUpdater) error
	UnRegisterApp(podid string, fromNRI bool, updater ResctrlUpdater) error
	GetApp(podid string) (App, bool)
	GetApps() map[string]App
}

func NewRDTEngine(vendor string) (ResctrlEngine, error) {
	_ = "STUB: not implemented"
	return *new(ResctrlEngine), nil
}

type RDTEngine struct {
	Apps       map[string]App
	Cgm        ControlGroupManager
	CtrlGroups map[string]apiext.Resctrl
	l          sync.RWMutex
	CBM        uint
	Vendor     string
}

func (R *RDTEngine) GetApps() map[string]App { _ = "STUB: not implemented"; return nil }

func (R *RDTEngine) Rebuild() { _ = "STUB: not implemented"; return }

func (R *RDTEngine) RegisterApp(podid, annotation string, fromNRI bool, updater ResctrlUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

func (R *RDTEngine) UnRegisterApp(podid string, fromNRI bool, updater ResctrlUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

func (R *RDTEngine) GetApp(id string) (App, bool) {
	_ = "STUB: not implemented"
	return *new(App), false
}

func (R *RDTEngine) calculateMba(mbaPercent int64) int64 { _ = "STUB: not implemented"; return 0 }

func calculateIntelMba(mbaPercent int64) int64 { _ = "STUB: not implemented"; return 0 }

func calculateAMDMba(mbaPercent int64) int64 { _ = "STUB: not implemented"; return 0 }

func (R *RDTEngine) ParseSchemata(config apiext.ResctrlConfig, cbm uint) *sysutil.ResctrlSchemataRaw {
	_ = "STUB: not implemented"
	return nil
}

// l3 mask MUST be a valid hex

func GetPodCgroupNewTaskIdsFromPodCtx(podMeta *protocol.PodContext, tasksMap map[int32]struct{}) []int32 {
	_ = "STUB: not implemented"
	return nil
}

func GetNewTaskIds(ids []int32, tasksMap map[int32]struct{}) ([]int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// only append the non-mapped ids
