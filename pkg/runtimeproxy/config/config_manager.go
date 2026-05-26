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

package config

import (
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	defaultConfigFileNums = 2
)

type ManagerInterface interface {
	GetAllHook() []*RuntimeHookConfig
	Run() error
}

type Manager struct {
	sync.Mutex
	configs map[string]*RuntimeHookConfigItem
	watcher *fsnotify.Watcher
}

type RuntimeHookConfigItem struct {
	filePath   string
	fileIno    uint64
	updateTime time.Time
	*RuntimeHookConfig
}

func (m *Manager) GetAllHook() []*RuntimeHookConfig { _ = "STUB: not implemented"; return nil }

func (m *Manager) getAllRegisteredFiles() []string { _ = "STUB: not implemented"; return nil }

func NewConfigManager() *Manager { _ = "STUB: not implemented"; return nil }

func (m *Manager) registerFileToWatchIfNeed(file string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) removeFileToWatch(filepath string) { _ = "STUB: not implemented"; return }

func (m *Manager) needRefreshConfig(filepath string) bool { _ = "STUB: not implemented"; return false }

// updateHookConfig loads config file, and register file to fsnotify watcher to watch
// config file content changed
// the filepath should be absolute path
func (m *Manager) updateHookConfig(filepath string) error { _ = "STUB: not implemented"; return nil }

func (m *Manager) Run() error { _ = "STUB: not implemented"; return nil }

// watch the newly generated config file

// collect the existing config

func (m *Manager) collectAllConfigs() error { _ = "STUB: not implemented"; return nil }

func (m *Manager) syncLoop() error { _ = "STUB: not implemented"; return nil }

// only reload config when write/rename/remove events

// should add the config file to watcher if event.Op is fsnotify.Create

func (m *Manager) removeUnusedConfigs() { _ = "STUB: not implemented"; return }

func (m *Manager) healthCheck() { _ = "STUB: not implemented"; return }
