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
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
)

type NodeResource struct {
	Resources     map[corev1.ResourceName]*resource.Quantity `json:"resources,omitempty"`
	ZoneResources map[string]corev1.ResourceList             `json:"zoneResources,omitempty"`
	Labels        map[string]string                          `json:"labels,omitempty"`
	Annotations   map[string]string                          `json:"annotations,omitempty"`
	Messages      map[corev1.ResourceName]string             `json:"messages,omitempty"`
	Resets        map[corev1.ResourceName]bool               `json:"resets,omitempty"`
}

func NewNodeResource(items ...ResourceItem) *NodeResource { _ = "STUB: not implemented"; return nil }

func (nr *NodeResource) Set(items ...ResourceItem) { _ = "STUB: not implemented"; return }

// omit empty message

func (nr *NodeResource) SetResourceList(rl corev1.ResourceList, message string) {
	_ = "STUB: not implemented"
	return
}

func (nr *NodeResource) Delete(items ...ResourceItem) { _ = "STUB: not implemented"; return }

func (nr *NodeResource) Get(name corev1.ResourceName) *resource.Quantity {
	_ = "STUB: not implemented"
	return nil
}

type ResourceItem struct {
	Name         corev1.ResourceName          `json:"name,omitempty"`
	Quantity     *resource.Quantity           `json:"quantity,omitempty"`
	ZoneQuantity map[string]resource.Quantity `json:"zoneQuantity,omitempty"`
	Labels       map[string]string            `json:"labels,omitempty"`
	Annotations  map[string]string            `json:"annotations,omitempty"`
	Message      string                       `json:"message,omitempty"` // the message about the resource calculation
	Reset        bool                         `json:"reset,omitempty"`   // whether to reset the resource or not
}

type ResourceMetrics struct {
	NodeMetric *slov1alpha1.NodeMetric `json:"nodeMetric,omitempty"`
	// extended metrics
	Extensions *slov1alpha1.ExtensionsMap `json:"extensions,omitempty"`
}

type SyncContext struct {
	lock       sync.RWMutex
	contextMap map[string]time.Time
}

func NewSyncContext() *SyncContext { _ = "STUB: not implemented"; return nil }

func (s *SyncContext) WithContext(m map[string]time.Time) *SyncContext {
	_ = "STUB: not implemented"
	return nil
}

func (s *SyncContext) Load(key string) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (s *SyncContext) Store(key string, value time.Time) { _ = "STUB: not implemented"; return }

func (s *SyncContext) Delete(key string) { _ = "STUB: not implemented"; return }
