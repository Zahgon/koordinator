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
	"time"

	gocache "github.com/patrickmn/go-cache"
)

const (
	Remove                = "Remove"
	Add                   = "Add"
	ExpirationTime  int64 = 10
	CleanupInterval       = 600 * time.Second
)

type Updater interface {
	Update(string) error
}

type SchemataUpdater interface {
	Update(id, schemata string) error
}

type ControlGroup struct {
	AppId       string
	GroupId     string
	Schemata    string
	Status      string
	CreatedTime int64
}

type ControlGroupManager struct {
	rdtcgs            *gocache.Cache
	reconcileInterval int64
	sync.Mutex
}

func NewControlGroupManager() ControlGroupManager {
	_ = "STUB: not implemented"
	return *new(ControlGroupManager)
}

func (c *ControlGroupManager) Init() { _ = "STUB: not implemented"; return }

// get resctrl filesystem root

// rebuild c.rdtcgs when restart

func (c *ControlGroupManager) AddPod(podid string, schemata string, fromNRI bool, createUpdater ResctrlUpdater, schemataUpdater ResctrlUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

// Update Schemata

func (c *ControlGroupManager) RemovePod(podid string, fromNRI bool, removeUpdater ResctrlUpdater) bool {
	_ = "STUB: not implemented"
	return false
}
