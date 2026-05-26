/*
Copyright 2022 The Koordinator Authors.
Copyright 2016 The Kubernetes Authors.

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

package framework

import (
	"sync"
)

// CleanupActionHandle is an integer pointer type for handling cleanup action
type CleanupActionHandle *int
type cleanupFuncHandle struct {
	actionHandle CleanupActionHandle
	actionHook   func()
}

var cleanupActionsLock sync.Mutex
var cleanupHookList = []cleanupFuncHandle{}

// AddCleanupAction installs a function that will be called in the event of the
// whole test being terminated.  This allows arbitrary pieces of the overall
// test to hook into SynchronizedAfterSuite().
// The hooks are called in last-in-first-out order.
func AddCleanupAction(fn func()) CleanupActionHandle {
	_ = "STUB: not implemented"
	return *new(CleanupActionHandle)
}

// RemoveCleanupAction removes a function that was installed by
// AddCleanupAction.
func RemoveCleanupAction(p CleanupActionHandle) { _ = "STUB: not implemented"; return }

// RunCleanupActions runs all functions installed by AddCleanupAction.  It does
// not remove them (see RemoveCleanupAction) but it does run unlocked, so they
// may remove themselves.
func RunCleanupActions() { _ = "STUB: not implemented"; return }

// Run unlocked.
