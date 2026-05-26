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

// ResourceValidator validates the resource value
type ResourceValidator interface {
	Validate(value string) (isValid bool, msg string)
}

type RangeValidator struct {
	max int64
	min int64
}

func (r *RangeValidator) Validate(value string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// compatible to cgroup-v2 file valued "max"

type CPUSetStrValidator struct{}

func (c *CPUSetStrValidator) Validate(value string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

type BlkIORangeValidator struct {
	resource string
	max      int64
	min      int64
}

func (r *BlkIORangeValidator) Validate(value string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// 253:16 2048
// 253:16 0

// 253:16 enable=1 ctrl=user rpct=95 rlat=3000 wpct=95 wlat=4000
// 253:16 enable=0

// 253:16 ctrl=user rbps=3324911720 rseqiops=168274 rrandiops=352545 wbps=2765819289 wseqiops=367565 wrandiops=339390
// 253:16 ctrl=auto

// compatible to cgroup-v2 file valued "max"

type NetClsRangeValidator struct {
	resource string
}

const (
	maxClassIdDecimal = 41231686041
	maxClassIdHex     = 99999999
)

func (r *NetClsRangeValidator) Validate(value string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

// You can write hexadecimal values to net_cls.classid; the format for these values is 0xAAAABBBB;
// AAAA is the major handle number and BBBB is the minor handle number. Reading net_cls.classid yields a decimal result.
// so, the max length of this value is 8.
