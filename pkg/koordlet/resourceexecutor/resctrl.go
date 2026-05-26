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
	"github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

const ErrResctrlDir = "resctrl path or file not exist"
const CacheIdIndex = 2

// NewResctrlReader: lazy resctrl reader, just check vendor to generate specific reader
func NewResctrlReader() ResctrlReader {
	_ = "STUB: not implemented"
	// Support two main platforms; other platforms need to add their implementation of the resctrl interface.
	return *new(ResctrlReader)
}

type CacheId int

// parent for resctrl is like: `BE`, `LS`
type ResctrlReader interface {
	ReadResctrlL3Stat(parent string) (map[CacheId]uint64, error)
	ReadResctrlMBStat(parent string) (map[CacheId]system.MBStatData, error)
}

type ResctrlBaseReader struct {
}

type ResctrlRDTReader struct {
	ResctrlBaseReader
}
type ResctrlAMDReader struct {
	ResctrlBaseReader
}

type fakeReader struct {
	ResctrlBaseReader
}

func (rr *fakeReader) ReadResctrlL3Stat(parent string) (map[CacheId]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rr *fakeReader) ReadResctrlMBStat(parent string) (map[CacheId]system.MBStatData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewResctrlRDTReader() ResctrlReader { _ = "STUB: not implemented"; return *new(ResctrlReader) }

func NewResctrlQoSReader() ResctrlReader { _ = "STUB: not implemented"; return *new(ResctrlReader) }

// ReadResctrlL3Stat: Reads the resctrl L3 cache statistics based on NUMA domain.
// For more information about x86 resctrl, refer to: https://docs.kernel.org/arch/x86/resctrl.html
func (rr *ResctrlBaseReader) ReadResctrlL3Stat(parent string) (map[CacheId]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read all l3-memory domains

// Convert the cache ID from the domain name string to an integer.

// Construct the path to the resctrl L3 cache occupancy file.

// Parse the L3 cache usage data from the file content.

// ReadResctrlMBStat: Reads the resctrl memory bandwidth statistics based on NUMA domain.
// For more information about x86 resctrl, refer to: https://docs.kernel.org/arch/x86/resctrl.html
func (rr *ResctrlBaseReader) ReadResctrlMBStat(parent string) (map[CacheId]system.MBStatData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read all l3-memory domains

// Parse the L3 cache usage data from the file content.

// Read the memory bandwidth statistics for the local and total memory bandwidth.
// The local memory bandwidth is the memory bandwidth consumed by the domain itself.
// The total memory bandwidth is the memory bandwidth consumed by the domain and accessed by other domains.
