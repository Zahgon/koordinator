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

package perf_group

import (
	"io"
	"os"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -lpfm
#include <perfmon/pfmlib.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

const (
	CYCLES       = "cycles"
	INSTRUCTIONS = "instructions"
)

var (
	initlibpfm    sync.Once
	closelibpfm   sync.Once
	perfValuePool sync.Pool
	BufPools      map[int]*sync.Pool
	EventsMap     = map[string][]string{
		"CPICollector": {"cycles", "instructions"},
	}
	attrMap = make(map[string]*unix.PerfEventAttr)
)

func InitBufferPool(eventsNums map[int]struct{}) { _ = "STUB: not implemented"; return }

// https://man7.org/linux/man-pages/man2/perf_event_open.2.html#Reading%20results
// 24 means the size of nr, time_enabled, time_running
// 16 means the size of value, id
// struct read_format {
//  u64 nr;            /* The number of events */
//  u64 time_enabled;  /* if PERF_FORMAT_TOTAL_TIME_ENABLED */
//  u64 time_running;  /* if PERF_FORMAT_TOTAL_TIME_RUNNING */
//  struct {
//      u64 value;     /* The value of the event */
//      u64 id;        /* if PERF_FORMAT_ID */
//      u64 lost;      /* if PERF_FORMAT_LOST */
//  } values[nr];
// };

func LibInit() { _ = "STUB: not implemented"; return }

func LibFinalize() { _ = "STUB: not implemented"; return }

type PerfGroupCollector struct {
	cgroupFile     *os.File
	cpus           []int
	perfCollectors map[int]*perfCollector
	idEventMap     map[uint64]string
	resultMap      map[string]float64
	valueCh        chan perfValue
	syscall6       func(trap uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (r1 uintptr, r2 uintptr, err syscall.Errno)
	closeCh        chan struct{}
}

type perfCollector struct {
	cpu      int
	syscall6 func(trap uintptr, a1 uintptr, a2 uintptr, a3 uintptr, a4 uintptr, a5 uintptr, a6 uintptr) (r1 uintptr, r2 uintptr, err syscall.Errno)
	leaderFd io.ReadCloser
	fds      []io.ReadCloser
}

type perfValue struct {
	Value float64
	ID    uint64
}

type value struct {
	Value uint64
	ID    uint64
}

type perfValueHeader struct {
	Nr          uint64 // number of events
	TimeEnabled uint64 // time event active
	TimeRunning uint64 // time event on CPU
}

// first event is group leader
func NewPerfGroupCollector(cgroupFile *os.File, cpus []int, events []string, syscallFunc func(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)) (collector *PerfGroupCollector, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create perf group

// enable perf group

// collect and statistic perf result

func GetAndStartPerfGroupCollectorOnContainer(cgroupFile *os.File, cpus []int, events []string) (*PerfGroupCollector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetContainerPerfResult(collector *PerfGroupCollector) (map[string]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetContainerCyclesAndInstructionsGroup(collector *PerfGroupCollector) (float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (c *PerfGroupCollector) cleanUp() error { _ = "STUB: not implemented"; return nil }

// caller must free the memory
func createPerfConfig(event string) (*unix.PerfEventAttr, error) {
	_ = "STUB: not implemented"
	// https://pkg.go.dev/cmd/cgo OOM instread of check malloc error
	return nil, nil
}

// pfmPerfEncodeArgT represents structure that is used to parse perf event name
// into perf_event_attr using libpfm4.
type pfmPerfEncodeArgT struct {
	attr unsafe.Pointer
	fstr unsafe.Pointer
	size C.size_t
	_    C.int // idx
	_    C.int // cpu
	_    C.int // flags
}

// https://man7.org/linux/man-pages/man3/pfm_get_os_event_encoding.3.html
func pfmGetOsEventEncoding(event string, perfEventAttrPtr unsafe.Pointer) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *perfCollector) collect(ch chan perfValue) error { _ = "STUB: not implemented"; return nil }

// stop stops perf group counter
func (p *perfCollector) stop() error { _ = "STUB: not implemented"; return nil }

// close closes all perf fds
func (p *perfCollector) close() error { _ = "STUB: not implemented"; return nil }
