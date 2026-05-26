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

package loadaware

import (
	"time"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/services"
)

var _ services.APIServiceProvider = &Plugin{}

type NodeAssignInfoData struct {
	Pods []PodAssignInfoData `json:"pods,omitempty"`

	ProdUsage         corev1.ResourceList `json:"prodUsage,omitempty"`
	NodeDelta         corev1.ResourceList `json:"nodeDelta,omitempty"`
	ProdDelta         corev1.ResourceList `json:"prodDelta,omitempty"`
	NodeEstimated     corev1.ResourceList `json:"nodeEstimated,omitempty"`
	NodeDeltaPods     []string            `json:"nodeDeltaPods,omitempty"`
	ProdDeltaPods     []string            `json:"prodDeltaPods,omitempty"`
	NodeEstimatedPods []string            `json:"nodeEstimatedPods,omitempty"`
}

type PodAssignInfoData struct {
	Timestamp time.Time   `json:"timestamp,omitempty"`
	Pod       *corev1.Pod `json:"pod,omitempty"`
}

func (p *Plugin) RegisterEndpoints(group *gin.RouterGroup) { _ = "STUB: not implemented"; return }
