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

package resourceexecutor

import (
	sysutil "github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

const (
	// CgroupMaxSymbolStr only appears in cgroups-v2 files, we consider the value as MaxInt64
	CgroupMaxSymbolStr string = "max"
	// CgroupMaxValueStr math.MaxInt64; writing `memory.high` with this do the same as set as "max"
	CgroupMaxValueStr string = "9223372036854775807"
)

const EmptyValueError string = "EmptyValueError"
const ErrCgroupDir = "cgroup path or file not exist"

// CgroupFileWriteIfDifferent writes the cgroup file if current value is different from the given value.
func cgroupFileWriteIfDifferent(cgroupTaskDir string, r sysutil.Resource, value string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// FIXME(saintube): Instead of handling cpuset resource in writing function, we should use a updater and do
//  MergeUpdate in resourceexecutor's LeveledUpdateBatch.

// compatible with cgroup valued "max"

// CgroupFileWrite writes the cgroup file with the given value.
func cgroupFileWrite(cgroupTaskDir string, r sysutil.Resource, value string) error {
	_ = "STUB: not implemented"
	return nil
}

// CgroupFileReadInt reads the cgroup file and returns an int64 value.
func cgroupFileReadInt(cgroupTaskDir string, r sysutil.Resource) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// compatible with cgroup valued "max"

// CgroupFileRead reads the cgroup file.
func cgroupFileRead(cgroupTaskDir string, r sysutil.Resource) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func readCgroupAndParseInt64(parentDir string, r sysutil.Resource) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// "max" means unlimited

// content: `%lld`

func readCgroupAndParseUint32(parentDir string, r sysutil.Resource) (uint32, error) {
	_ = "STUB: not implemented"
	// TODO: refactor with generics
	return 0, nil
}

// "max" means unlimited

func readCgroupAndParseUint64(parentDir string, r sysutil.Resource) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// "max" means unlimited

// content: `%llu`

// ReadCgroupAndParseInt32Slice reads the given cgroup content and parses it into an int32 slice.
// e.g. content: "1\n23\n0\n4\n56789" -> []int32{ 1, 23, 0, 4, 56789 }
// TODO: refactor via Generics.
func readCgroupAndParseInt32Slice(parentDir string, r sysutil.Resource) ([]int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// content: "%d\n%d\n%d\n..."

func IsCgroupPathExist(parentDir string, r sysutil.Resource) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func ResourceCgroupDirErr(msg string) error { _ = "STUB: not implemented"; return nil }

func IsCgroupDirErr(err error) bool { _ = "STUB: not implemented"; return false }
