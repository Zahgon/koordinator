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

package framework

import (
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type PodFilter interface {
	// FilterPod returns true if the collector should skip collecting the pod.
	FilterPod(podMeta *statesinformer.PodMeta) (shouldSkip bool, msg string)
}

// DefaultPodFilter filters nothing and allows collection for all pods.
var DefaultPodFilter PodFilter = &NothingPodFilter{}

type NothingPodFilter struct{}

func (n *NothingPodFilter) FilterPod(podMeta *statesinformer.PodMeta) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

type TerminatedPodFilter struct{}

func (t *TerminatedPodFilter) FilterPod(podMeta *statesinformer.PodMeta) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

type AndPodFilter struct {
	Filters []PodFilter
}

func NewAndPodFilter(filters ...PodFilter) PodFilter {
	_ = "STUB: not implemented"
	return *new(PodFilter)
}

func (a *AndPodFilter) FilterPod(podMeta *statesinformer.PodMeta) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}
