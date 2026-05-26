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

package system

// MountResctrlSubsystem mounts resctrl fs under the sysFSRoot to enable the kernel feature on supported environment
// NOTE: Linux kernel (>= 4.10), Intel cpu and bare-mental host are required; Also, Intel RDT
// features should be enabled in kernel configurations and kernel commandline.
// For more info, please see https://github.com/intel/intel-cmt-cat/wiki/resctrl
func MountResctrlSubsystem() (bool, error) {
	_ = "STUB: not implemented"
	// use schemata path to check since the subsystem root dir could keep exist when unmounted
	return false, nil
}
