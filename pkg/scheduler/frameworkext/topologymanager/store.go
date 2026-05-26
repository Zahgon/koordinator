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

package topologymanager

import (
	"sync"

	fwktype "k8s.io/kube-scheduler/framework"
)

const (
	affinityStateKey = "koordinator.sh/topology-affinity-store"
)

type Store struct {
	affinityMap sync.Map
}

func InitStore(cycleState fwktype.CycleState) { _ = "STUB: not implemented"; return }

func GetStore(cycleState fwktype.CycleState) *Store { _ = "STUB: not implemented"; return nil }

func (s *Store) Clone() fwktype.StateData {
	_ = "STUB: not implemented"
	return *new(fwktype.StateData)
}

func (s *Store) SetAffinity(nodeName string, affinity NUMATopologyHint) {
	_ = "STUB: not implemented"
	return
}

func (s *Store) GetAffinity(nodeName string) (NUMATopologyHint, bool) {
	_ = "STUB: not implemented"
	return *new(NUMATopologyHint), false
}
