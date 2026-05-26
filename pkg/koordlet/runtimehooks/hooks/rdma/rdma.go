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
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/hooks"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/runtimehooks/protocol"
)

const (
	IBDevDir  = "/dev/infiniband"
	RdmaCmDir = "/dev/infiniband/rdma_cm"
	SysBusPci = "/sys/bus/pci/devices"
)

type rdmaPlugin struct{}

func (p *rdmaPlugin) Register(op hooks.Options) { _ = "STUB: not implemented"; return }

var singleton *rdmaPlugin

func Object() *rdmaPlugin { _ = "STUB: not implemented"; return nil }

func (p *rdmaPlugin) InjectDevice(proto protocol.HooksProtocol) error {
	_ = "STUB: not implemented"
	return nil
}

// Both VF and PF of the same device are not allowed

func getUVerbsViaPciAdd(pciAddress string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
