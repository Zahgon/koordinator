/*
Copyright 2022 The Koordinator Authors.
Copyright 2019 The Kubernetes Authors.

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

// Package log will be removed after switching to use core framework log.
// Do not make further changes here!
package log

func nowStamp() string { _ = "STUB: not implemented"; return "" }

func log(level string, format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Logf logs the info.
func Logf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Failf logs the fail info.
func Failf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// FailfWithOffset calls "Fail" and logs the error at "offset" levels above its caller
// (for example, for call chain f -> g -> FailfWithOffset(1, ...) error would be logged for "f").
func FailfWithOffset(offset int, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}
