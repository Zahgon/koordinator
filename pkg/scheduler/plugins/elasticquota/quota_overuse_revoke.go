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
	"context"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/plugins/elasticquota/core"
)

const (
	QuotaOverUsedRevokeControllerName = "QuotaOverUsedRevokeController"
)

type QuotaOverUsedGroupMonitor struct {
	groupQuotaManger             *core.GroupQuotaManager
	quotaName                    string
	lastUnderUsedTime            time.Time
	overUsedTriggerEvictDuration time.Duration
}

func NewQuotaOverUsedGroupMonitor(quotaName string, manager *core.GroupQuotaManager, overUsedTriggerEvictDuration time.Duration) *QuotaOverUsedGroupMonitor {
	_ = "STUB: not implemented"
	return nil
}

func (monitor *QuotaOverUsedGroupMonitor) monitor() bool { _ = "STUB: not implemented"; return false }

func (monitor *QuotaOverUsedGroupMonitor) getToRevokePodList(quotaName string) []*v1.Pod {
	_ = "STUB: not implemented"
	return nil
}

// order pod from low priority -> high priority

// first try revoke all until used <= runtime

// means should evict all

//try assign back from high->low

type QuotaOverUsedRevokeController struct {
	monitorsLock                 sync.RWMutex
	monitors                     map[string]*QuotaOverUsedGroupMonitor
	overUsedTriggerEvictDuration time.Duration
	revokePodCycle               time.Duration
	monitorAllQuotas             bool
	enableRuntimeQuota           bool
	plugin                       *Plugin
}

func NewQuotaOverUsedRevokeController(plugin *Plugin) *QuotaOverUsedRevokeController {
	_ = "STUB: not implemented"
	return nil
}

func (controller *QuotaOverUsedRevokeController) Name() string {
	_ = "STUB: not implemented"
	return ""
}

func (controller *QuotaOverUsedRevokeController) Start() { _ = "STUB: not implemented"; return }

func (controller *QuotaOverUsedRevokeController) revokePodDueToQuotaOverUsed() {
	_ = "STUB: not implemented"
	return
}

func (controller *QuotaOverUsedRevokeController) monitorAll() []*v1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (controller *QuotaOverUsedRevokeController) syncQuota() { _ = "STUB: not implemented"; return }

func (controller *QuotaOverUsedRevokeController) addQuota(quotaName string, mgr *core.GroupQuotaManager) {
	_ = "STUB: not implemented"
	return
}

func (controller *QuotaOverUsedRevokeController) deleteQuota(quotaName string) {
	_ = "STUB: not implemented"
	return
}

func (controller *QuotaOverUsedRevokeController) getToMonitorQuotas() map[string]*QuotaOverUsedGroupMonitor {
	_ = "STUB: not implemented"
	return nil
}

func EvictPod(ctx context.Context, client clientset.Interface, pod *v1.Pod, deleteOptions *metav1.DeleteOptions) error {
	_ = "STUB: not implemented"
	// In Kubernetes 1.25+, the v1beta1 eviction API was removed.
	// Always use policy/v1 Eviction API for k8s 1.25 and above.
	return nil
}
