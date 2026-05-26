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

package util

import (
	"os"

	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/util/perf"
	perfgroup "github.com/koordinator-sh/koordinator/pkg/koordlet/util/perf_group"
)

func readTotalCPUStat(statPath string) (uint64, error) {
	_ = "STUB: not implemented"
	// stat usage: $user + $nice + $system + $irq + $softirq
	return 0, nil
}

// format: cpu $user $nice $system $idle $iowait $irq $softirq

// GetCPUStatUsageTicks returns the node's CPU usage ticks
func GetCPUStatUsageTicks() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func GetContainerPerfGroupCollector(podCgroupDir string, c *corev1.ContainerStatus, number int32, events []string) (*perfgroup.PerfGroupCollector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get file descriptor for cgroup mode perf_event_open

func GetContainerPerfCollector(podCgroupDir string, c *corev1.ContainerStatus, number int32) (*perf.PerfCollector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get file descriptor for cgroup mode perf_event_open

func GetContainerCyclesAndInstructions(collector perf.Collector) (float64, float64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func getContainerCgroupFile(podCgroupDir string, c *corev1.ContainerStatus) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
