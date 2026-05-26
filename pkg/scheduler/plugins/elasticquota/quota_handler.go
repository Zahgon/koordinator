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
	corev1 "k8s.io/api/core/v1"

	schedulerv1alpha1 "github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/plugins/elasticquota/core"
)

func (g *Plugin) OnQuotaAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (g *Plugin) OnQuotaUpdate(oldObj, newObj interface{}) { _ = "STUB: not implemented"; return }

// forbidden change quota tree.

// quota spec not change. return

// OnQuotaDelete if a quotaGroup is deleted, the pods should migrate to defaultQuotaGroup.
func (g *Plugin) OnQuotaDelete(obj interface{}) { _ = "STUB: not implemented"; return }

func (g *Plugin) ReplaceQuotas(objs []interface{}) error { _ = "STUB: not implemented"; return nil }

func (g *Plugin) GetQuotaSummary(quotaName string, includePods bool) (*core.QuotaInfoSummary, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (g *Plugin) GetQuotaSummaries(tree string, includePods bool) map[string]*core.QuotaInfoSummary {
	_ = "STUB: not implemented"
	return nil
}

func (g *Plugin) GetOrCreateGroupQuotaManagerForTree(treeID string) *core.GroupQuotaManager {
	_ = "STUB: not implemented"
	return nil
}

// return the default manager

// read lock

// write lock

func (g *Plugin) GetGroupQuotaManagerForTree(treeID string) *core.GroupQuotaManager {
	_ = "STUB: not implemented"
	return nil
}

func (g *Plugin) GetGroupQuotaManagerForQuota(quotaName string) *core.GroupQuotaManager {
	_ = "STUB: not implemented"
	return nil
}

func (g *Plugin) ListGroupQuotaManagersForQuotaTree() []*core.GroupQuotaManager {
	_ = "STUB: not implemented"
	return nil
}

func (g *Plugin) updateQuotaToTreeMap(quota, tree string) { _ = "STUB: not implemented"; return }

func (g *Plugin) deleteQuotaToTreeMap(quota string) { _ = "STUB: not implemented"; return }

// handlerQuotaForRoot will update quota tree total resource when the quota is root quota and enable MultiQuotaTree
func (g *Plugin) handlerQuotaWhenRoot(quota *schedulerv1alpha1.ElasticQuota, mgr *core.GroupQuotaManager, isDelete bool) {
	_ = "STUB: not implemented"
	return
}

// decrease the default GroupQuotaManager resource

func getTotalResource(quota *schedulerv1alpha1.ElasticQuota) (corev1.ResourceList, bool) {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceList), false
}
