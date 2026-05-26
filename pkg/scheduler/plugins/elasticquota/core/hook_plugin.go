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

package core

import (
	"sync"

	v1 "k8s.io/api/core/v1"

	"github.com/koordinator-sh/koordinator/apis/thirdparty/scheduler-plugins/pkg/apis/scheduling/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
)

// QuotaUpdateState provides a shared state storage for hookPlugins during quota updates
type QuotaUpdateState struct {
	Storage sync.Map
}

// QuotaInfoReader provides read-only access to quota information
type QuotaInfoReader struct {
	GetQuotaInfo       func(quotaName string) *QuotaInfo
	GetChildQuotaInfos func(quotaName string) map[string]*QuotaInfo

	// NoLock methods are used for read-only access without acquiring locks
	GetQuotaInfoNoLock       func(quotaName string) *QuotaInfo
	GetChildQuotaInfosNoLock func(quotaName string) map[string]*QuotaInfo
}

// QuotaHookPlugin defines the interface for quota management hookPlugins
type QuotaHookPlugin interface {
	// GetKey returns a unique key for the hook plugin
	GetKey() string

	// IsQuotaUpdated determines if the quota has been updated based on the hook plugin's perspective.
	// This method is invoked before updating the quota (excluding create/delete operations) with the GQM lock held.
	IsQuotaUpdated(oldQuotaInfo, newQuotaInfo *QuotaInfo, newQuota *v1alpha1.ElasticQuota) (isUpdated bool)

	// PreQuotaUpdate is invoked before performing quota create, update, or delete operations with the GQM lock held.
	// Parameters:
	// - oldQuotaInfo: The original quota information before the update. It is nil for create operations.
	// - newQuotaInfo: The updated quota information after the update. It is nil for delete operations.
	// - quota: The quota being updated. For delete operations, it is oldQuota;
	//			for create/update operations, it is newQuota.
	// - state: A shared state storage for hookPlugins during quota updates.
	PreQuotaUpdate(oldQuotaInfo, newQuotaInfo *QuotaInfo, quota *v1alpha1.ElasticQuota, state *QuotaUpdateState)

	// PostQuotaUpdate is invoked after performing quota create, update, or delete operations with the GQM lock held.
	// Parameters are the same as PreQuotaUpdate.
	PostQuotaUpdate(oldQuotaInfo, newQuotaInfo *QuotaInfo, quota *v1alpha1.ElasticQuota, state *QuotaUpdateState)

	// OnPodUpdated is triggered whenever an assigned pod's state changes within the quota,
	// including add, delete, update, reserve, or unreserve operations with the GQM lock held.
	// This method enables hook plugins to handle or modify the updated pod as needed.
	OnPodUpdated(quotaName string, oldPod, newPod *v1.Pod)

	// UpdateQuotaStatus is invoked periodically to update the quota status if needed.
	// This method allows hook plugins to modify quota status.
	// Parameters:
	// - oldQuota is the origin quota before updating status, it won't be nil.
	// - newQuota is a newly cloned quota with updated status, could be nil if no changes.
	// Return a newly cloned quota with updated status if changes are detected, otherwise nil.
	UpdateQuotaStatus(oldQuota, newQuota *v1alpha1.ElasticQuota) *v1alpha1.ElasticQuota

	// CheckPod verifies if a pod can be scheduled within the specified quota,
	// Return nil if the pod can be scheduled, otherwise return an error.
	CheckPod(quotaName string, pod *v1.Pod) error
}

type HookPluginFactory func(qiProvider *QuotaInfoReader, key, args string) (QuotaHookPlugin, error)

// hookPluginFactories is a global map of registered hook-plugin factories, keyed by factory key.
var hookPluginFactories = map[string]HookPluginFactory{}

func RegisterHookPluginFactory(factoryKey string, hookPluginFactory HookPluginFactory) {
	_ = "STUB: not implemented"
	return
}

func GetHookPluginFactory(factoryKey string) (HookPluginFactory, error) {
	_ = "STUB: not implemented"
	return *new(HookPluginFactory), nil
}

func initHookPlugins(qiProvider *QuotaInfoReader, args *config.ElasticQuotaArgs) (
	plugins []QuotaHookPlugin, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MetricsWrapper is a wrapper for QuotaHookPlugin that records metrics
type MetricsWrapper struct {
	plugin QuotaHookPlugin
}

// NewMetricsWrapper creates a new MetricsWrapper for a QuotaHookPlugin
func NewMetricsWrapper(plugin QuotaHookPlugin) QuotaHookPlugin {
	_ = "STUB: not implemented"
	return *new(QuotaHookPlugin)
}

// WrapWithMetrics wraps hook plugins with metrics
func WrapWithMetrics(plugins []QuotaHookPlugin) []QuotaHookPlugin {
	_ = "STUB: not implemented"
	return nil
}

func (w *MetricsWrapper) GetKey() string { _ = "STUB: not implemented"; return "" }

func (w *MetricsWrapper) IsQuotaUpdated(oldQuotaInfo, newQuotaInfo *QuotaInfo, newQuota *v1alpha1.ElasticQuota) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *MetricsWrapper) PreQuotaUpdate(oldQuotaInfo, newQuotaInfo *QuotaInfo, quota *v1alpha1.ElasticQuota, state *QuotaUpdateState) {
	_ = "STUB: not implemented"
	return
}

func (w *MetricsWrapper) PostQuotaUpdate(oldQuotaInfo, newQuotaInfo *QuotaInfo, quota *v1alpha1.ElasticQuota, state *QuotaUpdateState) {
	_ = "STUB: not implemented"
	return
}

func (w *MetricsWrapper) OnPodUpdated(quotaName string, oldPod, newPod *v1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (w *MetricsWrapper) UpdateQuotaStatus(oldQuota, newQuota *v1alpha1.ElasticQuota) *v1alpha1.ElasticQuota {
	_ = "STUB: not implemented"
	return nil
}

func (w *MetricsWrapper) CheckPod(quotaName string, pod *v1.Pod) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *MetricsWrapper) GetPlugin() QuotaHookPlugin {
	_ = "STUB: not implemented"

	// The following methods in GroupQuotaManager handle interactions with hook plugins
	return *new(QuotaHookPlugin)
}

func (gqm *GroupQuotaManager) InitHookPlugins(pluginArgs *config.ElasticQuotaArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) GetHookPlugins() []QuotaHookPlugin {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) SetHookPlugins(hookPlugins []QuotaHookPlugin) {
	_ = "STUB: not implemented"
	// Wrap plugins with metrics
	return
}

func (gqm *GroupQuotaManager) GetQuotaInfoReader() *QuotaInfoReader {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) GetChildGroupQuotaInfos(quotaName string) map[string]*QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

func (gqm *GroupQuotaManager) getChildGroupQuotaInfosNoLock(quotaName string) map[string]*QuotaInfo {
	_ = "STUB: not implemented"
	return nil
}

// runPreQuotaUpdateHooks executes all pre-update hooks for quota changes
func (gqm *GroupQuotaManager) runPreQuotaUpdateHooks(oldQuotaInfo, newQuotaInfo *QuotaInfo,
	quota *v1alpha1.ElasticQuota) *QuotaUpdateState {
	_ = "STUB: not implemented"
	return nil
}

// runPostQuotaUpdateHooks executes all post-update hooks for quota changes
// state could be nil if quota is newly added
func (gqm *GroupQuotaManager) runPostQuotaUpdateHooks(oldQuotaInfo, newQuotaInfo *QuotaInfo,
	quota *v1alpha1.ElasticQuota, state *QuotaUpdateState) {
	_ = "STUB: not implemented"
	return
}

// runPodUpdateHooks executes all hooks after pod used resource changed
func (gqm *GroupQuotaManager) runPodUpdateHooks(quotaName string, oldPod, newPod *v1.Pod) {
	_ = "STUB: not implemented"
	return
}

func (gqm *GroupQuotaManager) isQuotaUpdatedNoLock(oldQuotaInfo, newQuotaInfo *QuotaInfo,
	newQuota *v1alpha1.ElasticQuota) bool {
	_ = "STUB: not implemented"
	return false
}

// IsQuotaUpdated checks if the quota has been updated from the hook plugins' perspective.
func (gqm *GroupQuotaManager) IsQuotaUpdated(oldQuotaInfo, newQuotaInfo *QuotaInfo,
	newQuota *v1alpha1.ElasticQuota) bool {
	_ = "STUB: not implemented"
	return false
}

// ResetQuotasForHookPlugins resets quotas with pre-update and post-update hooks
func (gqm *GroupQuotaManager) ResetQuotasForHookPlugins(quotas map[string]*v1alpha1.ElasticQuota) {
	_ = "STUB: not implemented"
	return
}
