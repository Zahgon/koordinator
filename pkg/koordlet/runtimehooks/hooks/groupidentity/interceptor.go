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

package groupidentity

import (
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
)

func (b *bvtPlugin) SetPodBvtValue(p protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// pod annotatations take precedence to the default CPUQoS which retrieve from getPodBvtValue

// we can change group identity by pod annotations only when QoS=LS
// because LS=2 and BE=-1 by default
// we only allow LS change to BE
// and not allow BE change to LS
// and do not change LSR/LSE

// if disabled, set bvt to zero

func (b *bvtPlugin) SetKubeQOSBvtValue(p protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bvtPlugin) SetHostAppBvtValue(p protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bvtPlugin) prepare() *bvtRule { _ = "STUB: not implemented"; return nil }

// no need to update cgroups if both rule and sysctl disabled
