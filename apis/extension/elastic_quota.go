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

package extension

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
)

// RootQuotaName means quotaTree's root\head.
const (
	SystemQuotaName                      = "koordinator-system-quota"
	RootQuotaName                        = "koordinator-root-quota"
	DefaultQuotaName                     = "koordinator-default-quota"
	QuotaKoordinatorPrefix               = "quota.scheduling.koordinator.sh"
	LabelQuotaIsParent                   = QuotaKoordinatorPrefix + "/is-parent"
	LabelQuotaParent                     = QuotaKoordinatorPrefix + "/parent"
	LabelAllowLentResource               = QuotaKoordinatorPrefix + "/allow-lent-resource"
	LabelQuotaName                       = QuotaKoordinatorPrefix + "/name"
	LabelQuotaProfile                    = QuotaKoordinatorPrefix + "/profile"
	LabelQuotaIsRoot                     = QuotaKoordinatorPrefix + "/is-root"
	LabelQuotaTreeID                     = QuotaKoordinatorPrefix + "/tree-id"
	LabelQuotaIgnoreDefaultTree          = QuotaKoordinatorPrefix + "/ignore-default-tree"
	LabelPreemptible                     = QuotaKoordinatorPrefix + "/preemptible"
	LabelAllowForceUpdate                = QuotaKoordinatorPrefix + "/allow-force-update"
	AnnotationSharedWeight               = QuotaKoordinatorPrefix + "/shared-weight"
	AnnotationRuntime                    = QuotaKoordinatorPrefix + "/runtime"
	AnnotationRequest                    = QuotaKoordinatorPrefix + "/request"
	AnnotationChildRequest               = QuotaKoordinatorPrefix + "/child-request"
	AnnotationResourceKeys               = QuotaKoordinatorPrefix + "/resource-keys"
	AnnotationTotalResource              = QuotaKoordinatorPrefix + "/total-resource"
	AnnotationUnschedulableResource      = QuotaKoordinatorPrefix + "/unschedulable-resource"
	AnnotationQuotaNamespaces            = QuotaKoordinatorPrefix + "/namespaces"
	AnnotationGuaranteed                 = QuotaKoordinatorPrefix + "/guaranteed"
	AnnotationAllocated                  = QuotaKoordinatorPrefix + "/allocated"
	AnnotationNonPreemptibleRequest      = QuotaKoordinatorPrefix + "/non-preemptible-request"
	AnnotationNonPreemptibleUsed         = QuotaKoordinatorPrefix + "/non-preemptible-used"
	AnnotationAdmission                  = QuotaKoordinatorPrefix + "/admission"
	AnnotationMaxStrictCheckResourceKeys = QuotaKoordinatorPrefix + "/max-strict-check-resource-keys"
)

func GetParentQuotaName(quota *v1alpha1.ElasticQuota) string { _ = "STUB: not implemented"; return "" }

//default return RootQuotaName

func IsParentQuota(quota *v1alpha1.ElasticQuota) bool { _ = "STUB: not implemented"; return false }

func IsAllowLentResource(quota *v1alpha1.ElasticQuota) bool {
	_ = "STUB: not implemented"
	return false
}

func IsAllowForceUpdate(quota *v1alpha1.ElasticQuota) bool { _ = "STUB: not implemented"; return false }

func IsTreeRootQuota(quota *v1alpha1.ElasticQuota) bool { _ = "STUB: not implemented"; return false }

func IsPodNonPreemptible(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func GetQuotaTreeID(quota *v1alpha1.ElasticQuota) string { _ = "STUB: not implemented"; return "" }

func GetSharedWeight(quota *v1alpha1.ElasticQuota) corev1.ResourceList {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList)
}

//default equals to max

func IsForbiddenModify(quota *v1alpha1.ElasticQuota) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// can't modify SystemQuotaGroup

func GetQuotaName(pod *corev1.Pod) string { _ = "STUB: not implemented"; return "" }

func GetAnnotationQuotaNamespaces(quota *v1alpha1.ElasticQuota) []string {
	_ = "STUB: not implemented"
	return nil
}

func GetNonPreemptibleRequest(quota *v1alpha1.ElasticQuota) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func GetNonPreemptibleUsed(quota *v1alpha1.ElasticQuota) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func GetGuaranteed(quota *v1alpha1.ElasticQuota) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func GetAllocated(quota *v1alpha1.ElasticQuota) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func GetRuntime(quota *v1alpha1.ElasticQuota) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func GetRequest(quota *v1alpha1.ElasticQuota) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func GetChildRequest(quota *v1alpha1.ElasticQuota) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func GetUnschedulableResource(quota *v1alpha1.ElasticQuota) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func GetAdmission(quota *v1alpha1.ElasticQuota) (corev1.ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), nil
}

func GetMaxStrictCheckResourceKeys(quota *v1alpha1.ElasticQuota) ([]corev1.ResourceName, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
