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

package nodenumaresource

import (
	"reflect"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/apis/extension"
	schedulingconfig "github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
)

func GetDefaultNUMAAllocateStrategy(pluginArgs *schedulingconfig.NodeNUMAResourceArgs) schedulingconfig.NUMAAllocateStrategy {
	_ = "STUB: not implemented"
	return *new(schedulingconfig.NUMAAllocateStrategy)
}

func GetNUMAAllocateStrategy(node *corev1.Node, defaultNUMAtAllocateStrategy schedulingconfig.NUMAAllocateStrategy) schedulingconfig.NUMAAllocateStrategy {
	_ = "STUB: not implemented"
	return *new(schedulingconfig.NUMAAllocateStrategy)
}

func AllowUseCPUSet(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func mergeTopologyPolicy(nodePolicy, podPolicy extension.NUMATopologyPolicy) (extension.NUMATopologyPolicy, error) {
	_ = "STUB: not implemented"
	return *new(extension.NUMATopologyPolicy), nil
}

func getNUMATopologyPolicy(nodeLabels map[string]string, kubeletTopologyManagerPolicy extension.NUMATopologyPolicy) extension.NUMATopologyPolicy {
	_ = "STUB: not implemented"
	return *new(extension.NUMATopologyPolicy)
}

// amplifyNUMANodeResources amplifies the resources per NUMA Node.
// NOTE(joseph): After the NodeResource controller supports amplifying by ratios, should remove the function.
func amplifyNUMANodeResources(node *corev1.Node, topologyOptions *TopologyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func getCPUBindPolicy(topologyOptions *TopologyOptions, node *corev1.Node, requiredCPUBindPolicy, preferredCPUBindPolicy schedulingconfig.CPUBindPolicy) (schedulingconfig.CPUBindPolicy, bool, error) {
	_ = "STUB: not implemented"
	return *new(schedulingconfig.CPUBindPolicy), false, nil
}

func requestCPUBind(state *preFilterState, nodeCPUBindPolicy extension.NodeCPUBindPolicy) (bool, *fwktype.Status) {
	_ = "STUB: not implemented"
	return false, nil
}

func logStruct(v reflect.Value, key string, verbosity klog.Level) {
	_ = "STUB: not implemented"
	return
}

// logValue is a recursive function that prints the contents of any value.
// For pointers to structs, it recursively unwraps until it reaches the underlying struct.
func logValue(v reflect.Value, depth int, builder *strings.Builder) {
	_ = "STUB: not implemented"
	// Indent for pretty printing
	return
}

// For pointers, obtain the value being pointed to

// Appends "nil" representation for nil pointers

// Recursively log the pointer element

// For structs, iterate through all fields

// Appends the field name
// Recursively log each struct field

// For slices or arrays, iterate through each element

// Appends the index of the element
// Recursively log each element

// For other types, print the value directly
// Appends the value
