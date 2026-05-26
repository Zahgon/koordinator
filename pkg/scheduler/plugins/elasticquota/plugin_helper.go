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
	fwktype "k8s.io/kube-scheduler/framework"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/plugins/elasticquota/core"
)

// getPodAssociateQuotaName If pod's don't have the "quota-name" label, we will use the namespace to associate pod with quota
// group. If the plugin can't find the matched quota group, it will force the pod to associate with the "default-group".
func (g *Plugin) getPodAssociateQuotaName(pod *v1.Pod) string { _ = "STUB: not implemented"; return "" }

// getPodAssociateQuotaNameAndTreeID will return the quota and tree related the pod
// If pod's don't have the "quota-name" label, we will return the default quota and tree
// If pod has a quota label which not exists, we will also return the default quota and tree
func (g *Plugin) getPodAssociateQuotaNameAndTreeID(pod *v1.Pod) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (g *Plugin) GetQuotaName(pod *v1.Pod) string { _ = "STUB: not implemented"; return "" }

// migrateDefaultQuotaGroupsPod traverse all the pods in DefaultQuotaGroup, if the pod's QuotaName is not DefaultQuotaName,
// then erase the pod from DefaultQuotaGroup, Request. If the pod is Running, update Used.
func (g *Plugin) migrateDefaultQuotaGroupsPod() { _ = "STUB: not implemented"; return }

// different tree.

// the same tree.

// migratePods if a quotaGroup is deleted, migrate its pods to defaultQuotaGroup
func (g *Plugin) migratePods(out, in string) { _ = "STUB: not implemented"; return }

// createDefaultQuotaIfNotPresent create DefaultQuotaGroup's CRD
func (g *Plugin) createDefaultQuotaIfNotPresent() { _ = "STUB: not implemented"; return }

// defaultQuotaInfo and systemQuotaInfo are created once the groupQuotaManager is created, but we also want to see
// the used/request of the two quotaGroups, so we create the two quota's CRD if not present.
func (g *Plugin) createSystemQuotaIfNotPresent() { _ = "STUB: not implemented"; return }

// createRootQuotaIfNotPresent create RootQuotaGroup's CRD
func (g *Plugin) createRootQuotaIfNotPresent() { _ = "STUB: not implemented"; return }

func (g *Plugin) snapshotPostFilterState(quotaInfo *core.QuotaInfo, state fwktype.CycleState) *PostFilterState {
	_ = "STUB: not implemented"
	return nil
}

func (g *Plugin) skipPostFilterState(state fwktype.CycleState) { _ = "STUB: not implemented"; return }

func getPostFilterState(cycleState fwktype.CycleState) (*PostFilterState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Plugin) checkQuotaRecursive(mgr *core.GroupQuotaManager, curQuotaName string, quotaNameTopo []string, podRequest v1.ResourceList) *fwktype.Status {
	_ = "STUB: not implemented"
	return nil
}

func printResourceList(rl v1.ResourceList) string { _ = "STUB: not implemented"; return "" }

func (g *Plugin) getQuotaInfoUsedLimit(quotaInfo *core.QuotaInfo) v1.ResourceList {
	_ = "STUB: not implemented"
	return *new(v1.ResourceList)
}
