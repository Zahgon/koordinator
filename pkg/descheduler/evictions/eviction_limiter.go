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

package evictions

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
)

type EvictionLimiter struct {
	maxPodsToEvictPerNode      *uint
	maxPodsToEvictPerNamespace *uint
	maxPodsToEvictTotal        *uint
	lock                       sync.RWMutex
	totalCount                 uint
	nodePodCount               nodePodEvictedCount
	namespacePodCount          namespacePodEvictCount
}

func NewEvictionLimiter(
	maxPodsToEvictPerNode *uint,
	maxPodsToEvictPerNamespace *uint,
	maxPodsToEvictTotal *uint,
) *EvictionLimiter {
	_ = "STUB: not implemented"
	return nil
}

func (pe *EvictionLimiter) Reset() { _ = "STUB: not implemented"; return }

// NodeEvicted gives a number of pods evicted for node
func (pe *EvictionLimiter) NodeEvicted(nodeName string) uint { _ = "STUB: not implemented"; return 0 }

func (pe *EvictionLimiter) NamespaceEvicted(namespace string) uint {
	_ = "STUB: not implemented"
	return 0
}

// TotalEvicted gives a number of pods evicted through all nodes
func (pe *EvictionLimiter) TotalEvicted() uint { _ = "STUB: not implemented"; return 0 }

// NodeLimitExceeded checks if the number of evictions for a node was exceeded
func (pe *EvictionLimiter) NodeLimitExceeded(node *corev1.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (pe *EvictionLimiter) NamespaceLimitExceeded(namespace string) bool {
	_ = "STUB: not implemented"
	return false
}

func (pe *EvictionLimiter) AllowEvict(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (pe *EvictionLimiter) Done(pod *corev1.Pod) { _ = "STUB: not implemented"; return }
