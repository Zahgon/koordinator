//go:build !linux
// +build !linux

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

package perf

import "os"

type PerfCollector struct{}

func NewPerfCollector(cgroupFile *os.File, cpus []int) (*PerfCollector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetContainerCyclesAndInstructions(collector *PerfCollector) (float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func GetAndStartPerfCollectorOnContainer(cgroupFile *os.File, cpus []int) (*PerfCollector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *PerfCollector) stopAndClose() (err error) { _ = "STUB: not implemented"; return nil }

func (c *PerfCollector) closeOnSingleCPU(cpu int) error { _ = "STUB: not implemented"; return nil }

func (c *PerfCollector) stopOnSingleCPU(cpu int) error { _ = "STUB: not implemented"; return nil }

type HardwareProfile struct{}

func (c *PerfCollector) hwProfileOnSingleCPU(cpu int) (HardwareProfile, error) {
	_ = "STUB: not implemented"
	return *new(HardwareProfile), nil
}

type collectResult struct{}

func (c *PerfCollector) collect() (result collectResult, err error) {
	_ = "STUB: not implemented"
	return *new(collectResult), nil
}

func (c *PerfCollector) CleanUp() error { _ = "STUB: not implemented"; return nil }

type Collector interface{}
