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

package rdma

import (
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
)

func GetNetDevice() (metriccache.Devices, error) {
	_ = "STUB: not implemented"
	return *new(metriccache.Devices), nil
}

// pf is loaded first, so the ibdev name is smaller

const (
	classIDBaseInt = 16
	classIDBitSize = 64
	netDevClassID  = 0x02
)

func isNetDevice(devClassID string) bool { _ = "STUB: not implemented"; return false }

// parseDeviceClassID returns device ID parsed from the string as 64bit integer
func parseDeviceClassID(deviceID string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
