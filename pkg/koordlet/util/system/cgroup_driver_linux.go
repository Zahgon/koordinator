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

import (
	"sync"
)

const (
	kubeletConfigCgroupDriverKey = "cgroupDriver"
)

var (
	isUnifiedOnce sync.Once
	isUnified     bool
)

func GetCgroupDriverFromCgroupName() CgroupDriverType {
	_ = "STUB: not implemented"
	return *new(CgroupDriverType)
}

// GetCgroupDriverFromKubeletPort get Kubelet's cgroup driver from kubelet port.
//  1. use KubeletPortToPid to get kubelet pid.
//  2. If '--cgroup-driver' in args, that's it.
//     else if '--config' not in args, is default driver('cgroupfs').
//     else go to step-3.
//  3. If kubelet config is relative path, join with /proc/${pidof kubelet}/cwd.
//     search 'cgroupDriver:' in kubelet config file, that's it.
func GetCgroupDriverFromKubeletPort(port int) (CgroupDriverType, error) {
	_ = "STUB: not implemented"
	return *new(CgroupDriverType), nil
}

// kubelet command-line args will override configuration from config file

// parse kubelet config file

// kubelet config file is in host path

// remove trailing ':' from key

// IsUsingCgroupsV2 checks once if the CGroup V2 is in use.
// modify base: github.com/opencontainers/runc/libcontainer/cgroups/utils.go IsCgroup2UnifiedMode
func IsUsingCgroupsV2() bool { _ = "STUB: not implemented"; return false }

// ignore the "not found" error if running in userns
