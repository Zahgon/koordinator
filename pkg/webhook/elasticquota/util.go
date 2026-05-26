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

type PodWrapper struct{ *v1.Pod }

func MakePod(namespace, name string) *PodWrapper { _ = "STUB: not implemented"; return nil }

func (p *PodWrapper) Label(string1, string2 string) *PodWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodWrapper) Container(request v1.ResourceList) *PodWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (p *PodWrapper) Obj() *v1.Pod { _ = "STUB: not implemented"; return nil }

type QuotaWrapper struct {
	*v1alpha1.ElasticQuota
}

func MakeQuota(name string) *QuotaWrapper { _ = "STUB: not implemented"; return nil }

func (q *QuotaWrapper) Namespace(ns string) *QuotaWrapper { _ = "STUB: not implemented"; return nil }

func (q *QuotaWrapper) Min(min v1.ResourceList) *QuotaWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (q *QuotaWrapper) Max(max v1.ResourceList) *QuotaWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (q *QuotaWrapper) Used(used v1.ResourceList) *QuotaWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (q *QuotaWrapper) ChildRequest(request v1.ResourceList) *QuotaWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (q *QuotaWrapper) Admission(request v1.ResourceList) *QuotaWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (q *QuotaWrapper) TreeID(tree string) *QuotaWrapper { _ = "STUB: not implemented"; return nil }

func (q *QuotaWrapper) Guaranteed(guaranteed v1.ResourceList) *QuotaWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (q *QuotaWrapper) IsRoot(isRoot bool) *QuotaWrapper { _ = "STUB: not implemented"; return nil }

func (q *QuotaWrapper) sharedWeight(sharedWeight v1.ResourceList) *QuotaWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (q *QuotaWrapper) IsParent(isParent bool) *QuotaWrapper { _ = "STUB: not implemented"; return nil }

func (q *QuotaWrapper) ParentName(parentName string) *QuotaWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (q *QuotaWrapper) Annotations(annotations map[string]string) *QuotaWrapper {
	_ = "STUB: not implemented"
	return nil
}

func (q *QuotaWrapper) Obj() *v1alpha1.ElasticQuota { _ = "STUB: not implemented"; return nil }

type resourceWrapper struct{ v1.ResourceList }

func MakeResourceList() *resourceWrapper { _ = "STUB: not implemented"; return nil }

func (r *resourceWrapper) CPU(val int64) *resourceWrapper { _ = "STUB: not implemented"; return nil }

func (r *resourceWrapper) Mem(val int64) *resourceWrapper { _ = "STUB: not implemented"; return nil }

func (r *resourceWrapper) GPU(val int64) *resourceWrapper { _ = "STUB: not implemented"; return nil }

func (r *resourceWrapper) Obj() v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}
