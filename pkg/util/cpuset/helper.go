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

package cpuset

func int32sToInts(values []int32) []int { _ = "STUB: not implemented"; return nil }

func intsToInt32s(values []int) []int32 { _ = "STUB: not implemented"; return nil }

// MergeCPUSet merges the old cpuset with the new one, and also deduplicate and keeps a desc order by processor ids
// e.g. [1,0], [3,2,2,1] => [3,2,1,0]
func MergeCPUSet(old, new []int32) []int32 { _ = "STUB: not implemented"; return nil }

// ParseCPUSetStr parses cpuset string into a slice
// eg. "0-5,34,46-48" => [0,1,2,3,4,5,34,46,47,48]
func ParseCPUSetStr(cpusetStr string) ([]int32, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseCPUSet parses CPUSet object into an int32 slice
// eg. { elem: { 0:{}, 1:{}, 5:{}, 6:{}, 7:{} } } => [0,1,5,6,7]
func ParseCPUSet(cpus *CPUSet) []int32 { _ = "STUB: not implemented"; return nil }

// GenerateCPUSetStr generates the cpuset string from the cpuset slice
// eg. [3,2,1,0] => "3,2,1,0"
func GenerateCPUSetStr(cpuset []int32) string { _ = "STUB: not implemented"; return "" }
