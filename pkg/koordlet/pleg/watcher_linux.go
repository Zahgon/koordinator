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

package pleg

import (
	"sync"

	"k8s.io/utils/inotify"
)

// NewWatcher creates and returns a new inotify instance using inotify_init(2)
func NewWatcher() (Watcher, error) { _ = "STUB: not implemented"; return *new(Watcher), nil }

type inotifyWatcher struct {
	// FIXME the underlying mutex in inotify.watcher has a bug, so add a mutex here
	sync.Mutex

	watcher *inotify.Watcher
}

// Close closes an inotify watcher instance
// It sends a message to the reader goroutine to quit and removes all watches
// associated with the inotify instance
func (w *inotifyWatcher) Close() error { _ = "STUB: not implemented"; return nil }

// AddWatch adds path to the watched file set.
func (w *inotifyWatcher) AddWatch(path string) error { _ = "STUB: not implemented"; return nil }

// RemoveWatch removes path from the watched file set.
func (w *inotifyWatcher) RemoveWatch(path string) error { _ = "STUB: not implemented"; return nil }

// Event returns the undlying event channel
func (w *inotifyWatcher) Event() chan *inotify.Event { _ = "STUB: not implemented"; return nil }

// Event returns the undlying error channel
func (w *inotifyWatcher) Error() chan error { _ = "STUB: not implemented"; return nil }

// TypeOf tell the type of event
func TypeOf(event *inotify.Event) EventType { _ = "STUB: not implemented"; return *new(EventType) }
