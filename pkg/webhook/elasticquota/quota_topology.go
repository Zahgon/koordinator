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
	"sync"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
)

type quotaTopology struct {
	lock sync.Mutex
	// quotaInfoMap stores all quota information
	quotaInfoMap map[string]*QuotaInfo
	// namespaceMap key: annotationNamespace, val: quotaName
	namespaceToQuotaMap map[string]string
	// quotaHierarchyInfo stores the quota's all children
	quotaHierarchyInfo map[string]map[string]struct{}

	client client.Client
}

func NewQuotaTopology(client client.Client) *quotaTopology { _ = "STUB: not implemented"; return nil }

func (qt *quotaTopology) ValidAddQuota(quota *v1alpha1.ElasticQuota) error {
	_ = "STUB: not implemented"
	return nil
}

func (qt *quotaTopology) ValidUpdateQuota(oldQuota, newQuota *v1alpha1.ElasticQuota) error {
	_ = "STUB: not implemented"
	return nil
}

func (qt *quotaTopology) ValidDeleteQuota(quota *v1alpha1.ElasticQuota) error {
	_ = "STUB: not implemented"
	return nil
}

// check has child quota.

// fillQuotaDefaultInformation fills quota with default information if not be configured
func (qt *quotaTopology) fillQuotaDefaultInformation(quota *v1alpha1.ElasticQuota) error {
	_ = "STUB: not implemented"
	return nil
}

// add tree id, if the parent has tree id

type QuotaTopologySummary struct {
	QuotaInfoMap       map[string]*QuotaInfoSummary `json:"quotaInfoMap"`
	QuotaHierarchyInfo map[string][]string          `json:"quotaHierarchyInfo"`
}

func NewQuotaTopologySummary() *QuotaTopologySummary { _ = "STUB: not implemented"; return nil }

func (qt *quotaTopology) getQuotaTopologyInfo() *QuotaTopologySummary {
	_ = "STUB: not implemented"
	return nil
}

func (qt *quotaTopology) getQuotaInfo(name, namespace string) *QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

// fixedSharedWeight keep keys in sharedWeight and maxQuota same
// if key in maxQuota not included in sharedWeight, add key/value in sharedWeight
// if key in sharedWeight not included in maxQuota, delete key/value in sharedWeight
// if fixed, return true
func fixedSharedWeight(sharedWeight, maxQuota corev1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}
