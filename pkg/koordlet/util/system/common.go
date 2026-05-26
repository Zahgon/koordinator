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

package system

import (
	"syscall"
)

var (
	CommonRootDir = "" // for uni-test
	PageSize      = int64(syscall.Getpagesize())
)

func CommonFileRead(file string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func CommonFileWriteIfDifferent(file string, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func CommonFileWrite(file string, data string) error { _ = "STUB: not implemented"; return nil }

// ReadFileNoStat uses io.ReadAll to read contents of entire file.
// This is similar to io.ReadFile but without the call to os.Stat, because
// many files in /proc and /sys report incorrect file sizes (either 0 or 4096).
// Reads a max file size of 512kB.  For files larger than this, a scanner
// should be used.
func ReadFileNoStat(filename string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func FileExists(path string) bool { _ = "STUB: not implemented"; return false }

func PathExists(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ParseKVMap parses a file content into a KV map.
// e.g. `user 100\nsystem 20` -> `{"user": "100", "system": "20"}`
func ParseKVMap(content string) map[string]string { _ = "STUB: not implemented"; return nil }

// GoWithNewThread synchronously runs the function in a new goroutine bound to a new OS thread.
func GoWithNewThread(f func() interface{}) interface{} {
	_ = "STUB: not implemented"
	// Lock the thread of the caller goroutine to ensure the thread does not change outside the new goroutine.
	return nil
}

// When the calling goroutine exits without unlocking the thread, the thread will be terminated.
// It helps the function to lock with an individual thread so not to affect the caller goroutine.
