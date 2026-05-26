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

func (qt *quotaTopology) validateQuotaSelfItem(quota *v1alpha1.ElasticQuota) error {
	_ = "STUB: not implemented"
	// min and max's each dimension should not have negative value
	return nil
}

// 1.check if sharewight is equal to

// 1. check if all key in min are included in max
// 2. check if all quantities in min <= that in max

// 1. check if all key in AnnotationMaxStrictCheckResourceKeys in max >= that in used

// check should be skipped if current used never counts

// validateQuotaTopology checks the quotaInfo's topology with its parent and its children.
// oldQuotaInfo is null when validate a new create request, and is the current quotaInfo when validate a update request.
func (qt *quotaTopology) validateQuotaTopology(oldQuotaInfo, newQuotaInfo *QuotaInfo, oldNamespaces []string) error {
	_ = "STUB: not implemented"
	return nil
}

// if the quotaInfo's parent is root and its IsParent is false, the following checks will be true, just return nil.

func (qt *quotaTopology) checkTreeID(oldQuotaInfo, quotaInfo *QuotaInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// check the parent tree id

// checkParentQuotaInfo has check parent exist

// check the children tree id

func (qt *quotaTopology) checkIsParentChange(oldQuotaInfo, quotaInfo *QuotaInfo, oldNamespaces []string) error {
	_ = "STUB: not implemented"
	// means create quota, no need check
	return nil
}

// checkParentQuotaInfo check parent exist
func (qt *quotaTopology) checkParentQuotaInfo(quotaName, parentName string) error {
	_ = "STUB: not implemented"
	return nil
}

// checkSubAndParentGroupQuotaKey check the quotaInfo's quota with its parent and its children
//
//	while enableResourceTypeUpdate=false, the quotaInfo's max quota key must be same as its children and parent's quota key
//	while enableResourceTypeUpdate=true, the quotaInfo's max quota key only need be included in its parent's quota key
//
// the quotaInfo's min quota key only need be included in its parent's quota key no matter when
func (qt *quotaTopology) checkSubAndParentGroupQuotaKey(quotaInfo *QuotaInfo, enableUpdateResourceKey bool) error {
	_ = "STUB: not implemented"
	return nil
}

// checkMinQuotaValidate will do two checks:
//  1. the sum of brothers' minquota should less than or equal to parentMinQuota.
//  2. the sum of children's minquota should less than or equal to newQuotaMin.
func (qt *quotaTopology) checkMinQuotaValidate(newQuotaInfo *QuotaInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// If the quota is tree root, we don't check it's min

// check children's minquota sum

func (qt *quotaTopology) getChildMinQuotaSumExceptSpecificChild(parentName, skipQuota string) (allChildQuotaSum v1.ResourceList, err error) {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList), nil
}

func toElasticQuota(obj interface{}) *v1alpha1.ElasticQuota { _ = "STUB: not implemented"; return nil }

func quotaFieldsCopy(q *v1alpha1.ElasticQuota) v1alpha1.ElasticQuota {
	_ = "STUB: not implemented"
	return *new(v1alpha1.ElasticQuota)
}

func checkQuotaKeySame(parent, child v1.ResourceList) bool { _ = "STUB: not implemented"; return false }

// checkQuotaKeyIncluded will check whether the parent quota includes all keys of child quota.
func checkQuotaKeyIncluded(parent, child v1.ResourceList) bool {
	_ = "STUB: not implemented"
	return false
}

func (qt *quotaTopology) checkGuaranteedForMin(quotaInfo *QuotaInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// If the quota is tree root, allow it change the min.

// if the new min less than guaranteed, which means that no more resource guarantee is needed, so it is allowed directly.

// Guaranteed resources are searched starting from the parent node of the current node until the root node of the tree.
// We need to meet guaranteed resources at least at a certain level to guarantee the resources of the current node.
// During the search process, there may be situations where min is not used up at the intermediate nodes.
// Therefore, we need to recursively accumulate the guaranteed resources that need to be satisfied until the root node is reached.
func (qt *quotaTopology) checkParentGuaranteed(newGuarantee v1.ResourceList, self, parentName string) error {
	_ = "STUB: not implemented"
	return nil
}
