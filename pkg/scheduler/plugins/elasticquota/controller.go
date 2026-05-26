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

package elasticquota

import (
	v1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/plugins/elasticquota/core"
)

const (
	ControllerName = "ElasticQuotaController"
)

// Controller is a controller that update elastic quota crd
type Controller struct {
	plugin *Plugin
}

func NewElasticQuotaController(plugin *Plugin) *Controller { _ = "STUB: not implemented"; return nil }

func (ctrl *Controller) Name() string { _ = "STUB: not implemented"; return "" }

func (ctrl *Controller) Start() { _ = "STUB: not implemented"; return }

func (ctrl *Controller) syncElasticQuotaRuntimeWorker() { _ = "STUB: not implemented"; return }

func (ctrl *Controller) syncElasticQuotaStatusWorker() { _ = "STUB: not implemented"; return }

func (ctrl *Controller) syncElasticQuotaStatus(eq *v1alpha1.ElasticQuota) {
	_ = "STUB: not implemented"
	return
}

// Update the status of the elastic quota by hook plugins

// createMergePatch return patch generated from original and new interfaces
func createMergePatch(original, new interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var resourceDecorators []func(quota *v1alpha1.ElasticQuota, resource v1.ResourceList)

func decorateResource(quota *v1alpha1.ElasticQuota, resource v1.ResourceList) {
	_ = "STUB: not implemented"
	return
}

type traceChange struct {
	key      string
	original v1.ResourceList
	current  v1.ResourceList
}

func updateElasticQuotaStatusIfChanged(eq *v1alpha1.ElasticQuota, summary *core.QuotaInfoSummary, logChanges bool) (*v1alpha1.ElasticQuota, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isElasticQuotaAnnotationDiff(eq *v1alpha1.ElasticQuota, key string, resourceList v1.ResourceList) (bool, v1.ResourceList, error) {
	_ = "STUB: not implemented"
	return false, *new(v1.ResourceList), nil
}

func updateElasticQuotaAnnotation(eq *v1alpha1.ElasticQuota, key string, resourceList v1.ResourceList) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctrl *Controller) syncElasticQuotaStatusMetricsWorker() { _ = "STUB: not implemented"; return }

func syncElasticQuotaMetrics(eq *v1alpha1.ElasticQuota, summary *core.QuotaInfoSummary) {
	_ = "STUB: not implemented"
	return
}

func deleteElasticQuotaMetrics(eq *v1alpha1.ElasticQuota, summary *core.QuotaInfoSummary) {
	_ = "STUB: not implemented"
	return
}

func getResourceFieldAndQuotaLabels(eq *v1alpha1.ElasticQuota, summary *core.QuotaInfoSummary) (map[string]v1.ResourceList, map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// record the unschedulable resource
