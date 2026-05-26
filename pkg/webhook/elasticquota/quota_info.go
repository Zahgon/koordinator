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
)

type QuotaInfo struct {
	IsParent          bool
	AllowLentResource bool
	AllowForceUpdate  bool
	Name              string
	ParentName        string
	TreeID            string
	IsTreeRoot        bool
	CalculateInfo     QuotaCalculateInfo
}

type QuotaCalculateInfo struct {
	// The semantics of "max" is the quota group's upper limit of resources.
	Max v1.ResourceList
	// The semantics of "min" is the quota group's guaranteed resources, if quota group's "request" less than or
	// equal to "min", the quota group can obtain equivalent resources to the "request"
	Min v1.ResourceList

	Guaranteed v1.ResourceList
	Allocated  v1.ResourceList
}

func NewQuotaInfo(isParent, allowLentResource bool, name, parentName string) *QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

func NewQuotaInfoFromQuota(quota *v1alpha1.ElasticQuota) *QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

func (qi *QuotaInfo) setMaxQuotaNoLock(res v1.ResourceList) { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) setMinQuotaNoLock(res v1.ResourceList) { _ = "STUB: not implemented"; return }

func (qi *QuotaInfo) GetQuotaSummary() *QuotaInfoSummary { _ = "STUB: not implemented"; return nil }

type QuotaInfoSummary struct {
	Name              string `json:"name"`
	ParentName        string `json:"parentName"`
	IsParent          bool   `json:"isParent"`
	AllowLentResource bool   `json:"allowLentResource"`

	Max v1.ResourceList `json:"max"`
	Min v1.ResourceList `json:"min"`
}

func NewQuotaInfoSummary() *QuotaInfoSummary { _ = "STUB: not implemented"; return nil }
