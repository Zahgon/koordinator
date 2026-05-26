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

package resourceexecutor

import (
	"sync"
	"time"

	"github.com/koordinator-sh/koordinator/pkg/koordlet/audit"
	sysutil "github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

var DefaultCgroupUpdaterFactory = NewCgroupUpdaterFactory()

func init() {
	// register the update logic for system resources
	// NOTE: should exclude the read-only resources, e.g. `cpu.stat`.
	// common
	DefaultCgroupUpdaterFactory.Register(NewCgroupUpdaterWithUpdateFunc(CgroupUpdateWithUnlimitedFunc),
		sysutil.CPUCFSPeriodName,
		sysutil.MemoryLimitName,
	)
	DefaultCgroupUpdaterFactory.Register(NewCommonCgroupUpdater,
		sysutil.CPUBurstName,
		sysutil.CPUBVTWarpNsName,
		sysutil.CPUIdleName,
		sysutil.CPUTasksName,
		sysutil.CPUProcsName,
		sysutil.MemoryWmarkRatioName,
		sysutil.MemoryWmarkScaleFactorName,
		sysutil.MemoryWmarkMinAdjName,
		sysutil.MemoryPriorityName,
		sysutil.MemoryUsePriorityOomName,
		sysutil.MemoryOomGroupName,
		sysutil.NetClsClassIdName,
	)
	// special cases
	DefaultCgroupUpdaterFactory.Register(NewCgroupUpdaterWithUpdateFunc(CgroupUpdateCPUSharesFunc), sysutil.CPUSharesName)
	DefaultCgroupUpdaterFactory.Register(NewMergeableCgroupUpdaterWithConditionFunc(CgroupUpdateWithUnlimitedFunc, MergeConditionIfCFSQuotaIsLarger),
		sysutil.CPUCFSQuotaName,
	)
	DefaultCgroupUpdaterFactory.Register(NewMergeableCgroupUpdaterIfValueLarger,
		sysutil.MemoryMinName,
		sysutil.MemoryLowName,
		sysutil.MemoryHighName,
	)
	DefaultCgroupUpdaterFactory.Register(NewMergeableCgroupUpdaterWithConditionFunc(CommonCgroupUpdateFunc, MergeConditionIfCPUSetIsLooser),
		sysutil.CPUSetCPUSName,
	)
	DefaultCgroupUpdaterFactory.Register(NewBlkIOResourceUpdater,
		sysutil.BlkioTRIopsName,
		sysutil.BlkioTRBpsName,
		sysutil.BlkioTWIopsName,
		sysutil.BlkioTWBpsName,
		sysutil.BlkioIOQoSName,
		sysutil.BlkioIOModelName,
		sysutil.BlkioIOWeightName,
	)
}

type UpdateFunc func(resource ResourceUpdater) error

type MergeUpdateFunc func(resource ResourceUpdater) (ResourceUpdater, error)

type ResourceUpdater interface {
	// Name returns the name of the resource updater
	Name() string
	ResourceType() sysutil.ResourceType
	Key() string
	Path() string
	Value() string
	MergeUpdate() (ResourceUpdater, error)
	Clone() ResourceUpdater
	GetLastUpdateTimestamp() time.Time
	UpdateLastUpdateTimestamp(time time.Time)
	update() error
}

type CgroupResourceUpdater struct {
	file      sysutil.Resource
	parentDir string
	value     string

	lastUpdateTimestamp time.Time
	updateFunc          UpdateFunc
	// MergeableResourceUpdater implementation (used by LeveledCacheExecutor):
	// For cgroup interfaces like `cpuset.cpus` and `memory.min`, reconciliation from top to bottom should keep the
	// upper value larger/broader than the lower. Thus a Leveled updater is implemented as follows:
	// 1. update batch of cgroup resources group by cgroup interface, i.e. cgroup filename.
	// 2. update each cgroup resource by the order of layers: firstly update resources from upper to lower by merging
	//    the new value with old value; then update resources from lower to upper with the new value.
	mergeUpdateFunc MergeUpdateFunc
	eventHelper     *audit.EventHelper
}

func (u *CgroupResourceUpdater) Name() string { _ = "STUB: not implemented"; return "" }

func (u *CgroupResourceUpdater) ResourceType() sysutil.ResourceType {
	_ = "STUB: not implemented"
	return *new(sysutil.ResourceType)
}

func (u *CgroupResourceUpdater) Key() string { _ = "STUB: not implemented"; return "" }

func (u *CgroupResourceUpdater) Path() string { _ = "STUB: not implemented"; return "" }

func (u *CgroupResourceUpdater) Value() string { _ = "STUB: not implemented"; return "" }

func (u *CgroupResourceUpdater) update() error { _ = "STUB: not implemented"; return nil }

func (u *CgroupResourceUpdater) GetEventHelper() *audit.EventHelper {
	_ = "STUB: not implemented"
	return nil
}

func (u *CgroupResourceUpdater) SetEventHelper(a *audit.EventHelper) {
	_ = "STUB: not implemented"
	return
}

func (u *CgroupResourceUpdater) MergeUpdate() (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

func (u *CgroupResourceUpdater) Clone() ResourceUpdater {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater)
}

func (u *CgroupResourceUpdater) GetLastUpdateTimestamp() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (u *CgroupResourceUpdater) UpdateLastUpdateTimestamp(time time.Time) {
	_ = "STUB: not implemented"
	return
}

func (u *CgroupResourceUpdater) WithUpdateFunc(updateFunc UpdateFunc) *CgroupResourceUpdater {
	_ = "STUB: not implemented"
	return nil
}

func (u *CgroupResourceUpdater) WithMergeUpdateFunc(mergeUpdateFunc MergeUpdateFunc) *CgroupResourceUpdater {
	_ = "STUB: not implemented"
	return nil
}

type DefaultResourceUpdater struct {
	key                 string // the cache key to identify the updater (can be the filepath or other custom key)
	value               string
	file                string // the real filepath
	lastUpdateTimestamp time.Time
	updateFunc          UpdateFunc
	eventHelper         *audit.EventHelper
}

func (u *DefaultResourceUpdater) Name() string { _ = "STUB: not implemented"; return "" }

func (u *DefaultResourceUpdater) ResourceType() sysutil.ResourceType {
	_ = "STUB: not implemented"
	return *new(sysutil.ResourceType)
}

func (u *DefaultResourceUpdater) Key() string { _ = "STUB: not implemented"; return "" }

func (u *DefaultResourceUpdater) Path() string {
	_ = "STUB: not implemented"
	// no additional parent dir here
	return ""
}

func (u *DefaultResourceUpdater) Value() string { _ = "STUB: not implemented"; return "" }

func (u *DefaultResourceUpdater) update() error { _ = "STUB: not implemented"; return nil }

func (u *DefaultResourceUpdater) MergeUpdate() (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

func (u *DefaultResourceUpdater) Clone() ResourceUpdater {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater)
}

func (u *DefaultResourceUpdater) GetLastUpdateTimestamp() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (u *DefaultResourceUpdater) UpdateLastUpdateTimestamp(time time.Time) {
	_ = "STUB: not implemented"
	return
}

// NewCommonDefaultUpdater returns a DefaultResourceUpdater for update general files.
func NewCommonDefaultUpdater(key string, file string, value string, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

// NewCommonDefaultUpdaterWithUpdateFunc returns a DefaultResourceUpdater for update general files with the given update function.
func NewCommonDefaultUpdaterWithUpdateFunc(key string, file string, value string, updateFunc UpdateFunc, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

type NewResourceUpdaterFunc func(resourceType sysutil.ResourceType, parentDir string, value string, e *audit.EventHelper) (ResourceUpdater, error)

type ResourceUpdaterFactory interface {
	Register(g NewResourceUpdaterFunc, resourceTypes ...sysutil.ResourceType)
	New(resourceType sysutil.ResourceType, parentDir string, value string, e *audit.EventHelper) (ResourceUpdater, error)
}

func NewCgroupUpdater(resourceType sysutil.ResourceType, parentDir string, value string, updateFunc UpdateFunc, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

func NewCgroupUpdaterWithUpdateFunc(updateFn UpdateFunc) func(resourceType sysutil.ResourceType, parentDir string, value string, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return nil
}

// NewCommonCgroupUpdater returns a CgroupResourceUpdater for updating known cgroup resources.
func NewCommonCgroupUpdater(resourceType sysutil.ResourceType, parentDir string, value string, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

func NewMergeableCgroupUpdaterWithCondition(resourceType sysutil.ResourceType, parentDir string, value string, updateFunc UpdateFunc, mergeCondition MergeConditionFunc, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

func NewMergeableCgroupUpdaterWithConditionFunc(updateFn UpdateFunc, mergeCondition MergeConditionFunc) func(resourceType sysutil.ResourceType, parentDir string, value string, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return nil
}

func NewMergeableCgroupUpdaterIfValueLarger(resourceType sysutil.ResourceType, parentDir string, value string, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

// NewDetailCgroupUpdater returns a new *CgroupResourceUpdater according to the given Resource, which is generally used
// for backwards compatibility. It is not guaranteed for updating successfully since it does not retrieve from the
// known cgroup resources.
func NewDetailCgroupUpdater(resource sysutil.Resource, parentDir string, value string, updateFunc UpdateFunc, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

type CgroupUpdaterFactoryImpl struct {
	lock     sync.RWMutex
	registry map[sysutil.ResourceType]NewResourceUpdaterFunc
}

func NewCgroupUpdaterFactory() ResourceUpdaterFactory {
	_ = "STUB: not implemented"
	return *new(ResourceUpdaterFactory)
}

func (f *CgroupUpdaterFactoryImpl) Register(g NewResourceUpdaterFunc, resourceTypes ...sysutil.ResourceType) {
	_ = "STUB: not implemented"
	return
}

func (f *CgroupUpdaterFactoryImpl) New(resourceType sysutil.ResourceType, parentDir string, value string, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

func CommonCgroupUpdateFunc(resource ResourceUpdater) error { _ = "STUB: not implemented"; return nil }

func CommonDefaultUpdateFunc(resource ResourceUpdater) error { _ = "STUB: not implemented"; return nil }

func CgroupUpdateWithUnlimitedFunc(resource ResourceUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: convert "-1" to "max", since some cgroups-v2 files only accept "max" to unlimit resource instead of "-1".
//       DO NOT use it on the cgroups which has a valid value of "-1".

func CgroupUpdateCPUSharesFunc(resource ResourceUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

// convert values in `cpu.shares` (v1) into values in `cpu.weight` (v2)

type MergeConditionFunc func(oldValue, newValue string) (mergedValue string, needMerge bool, err error)

func MergeFuncUpdateCgroup(resource ResourceUpdater, mergeCondition MergeConditionFunc) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

// skip the write when merge condition is not meet

// otherwise, do write for the current value

// suppose current value is different

// MergeConditionIfValueIsLarger returns a merge condition where only do update when the new value is larger.
func MergeConditionIfValueIsLarger(oldValue, newValue string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// compatible with cgroup valued "max"

func MergeConditionIfCFSQuotaIsLarger(oldValue, newValue string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// cgroup-v2 content: "max 100000", "100000 100000"

// cgroup-v1 content: "-1", "100000"
// compatible with cgroup valued "max"

// MergeConditionIfCPUSetIsLooser returns a merge condition where only do update when the new cpuset value is looser.
func MergeConditionIfCPUSetIsLooser(oldValue, newValue string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// no need to merge if new cpuset is equal to old

// no need to merge if new cpuset is a subset of the old

// need to update with the merged of old and new cpuset values

func cgroupWriteIfDifferentWithLog(c *CgroupResourceUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

func commonWriteIfDifferentWithLog(c *DefaultResourceUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

func BlkIOUpdateFunc(resource ResourceUpdater) error { _ = "STUB: not implemented"; return nil }

func NewBlkIOResourceUpdater(resourceType sysutil.ResourceType, parentDir string, value string, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

func cgroupBlkIOFileWriteIfDifferent(cgroupTaskDir string, file sysutil.Resource, value string) error {
	_ = "STUB: not implemented"
	return nil
}

// https://www.alibabacloud.com/help/en/elastic-compute-service/latest/configure-the-weight-based-throttling-feature-of-blk-iocost
func CheckIfBlkRootConfigNeedUpdate(oldValue string, newValue string) bool {
	_ = "STUB: not implemented"
	return false
}

// blkio.cost.weight: configure iocost weight
// blkio.throttle.read_bps_device: configure read bps
// blkio.throttle.read_iops_device: configure read iops
// blkio.throttle.write_bps_device: configure write bps
// blkio.throttle.write_iops_device: configure write iops
func CheckIfBlkQOSNeedUpdate(oldValue string, newValue string) bool {
	_ = "STUB: not implemented"
	return false
}

// newValue: "253:0 30"
// out: [["253:0 0" "253:0" "0"]]

// If majminWithZero is not empty, it means to assign a zero value to a device

// If majminWithZero is empty, it means to update the blkio cgroup file

// If currentValue does not completely contain newValue, a write operation is required
