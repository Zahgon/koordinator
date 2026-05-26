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

package blkio

import (
	"fmt"
	"time"

	"k8s.io/client-go/tools/cache"

	slov1alpha1 "github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/metriccache"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/qosmanager/framework"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

const (
	BlkIOReconcileName = "BlkioReconcile"

	DefaultReadIOPS           = 0
	DefaultWriteIOPS          = 0
	DefaultReadBPS            = 0
	DefaultWriteBPS           = 0
	DefaultIOWeightPercentage = 100
	DefaultIOLatency          = 3000
	DefaultLatencyPercent     = 95
)

var _ framework.QOSStrategy = &blkIOReconcile{}

type blkIOReconcile struct {
	reconcileInterval time.Duration
	statesInformer    statesinformer.StatesInformer
	metricCache       metriccache.MetricCache
	executor          resourceexecutor.ResourceUpdateExecutor
	storageInfo       *metriccache.NodeLocalStorageInfo
}

func (b *blkIOReconcile) Enabled() bool { _ = "STUB: not implemented"; return false }

func (b *blkIOReconcile) Setup(context *framework.Context) { _ = "STUB: not implemented"; return }

func (b *blkIOReconcile) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

type (
	GetUpdaterFunc      func(block *slov1alpha1.BlockCfg, diskNumber string, dynamicPath string) (resources []resourceexecutor.ResourceUpdater)
	GetRemoverFunc      func(diskNumber string, dynamicPath string) (resources []resourceexecutor.ResourceUpdater)
	GetDiskRecorderFunc func(absolutePath string) (map[string]bool, error)
)

func New(opt *framework.Options) framework.QOSStrategy {
	_ = "STUB: not implemented"
	return *new(framework.QOSStrategy)
}

func (b *blkIOReconcile) init(stopCh <-chan struct{}) error {
	b.executor.Run(stopCh)
	if !cache.WaitForCacheSync(stopCh, b.statesInformer.HasSynced) {
		return fmt.Errorf("%s: timed out waiting for pvc caches to sync", BlkIOReconcileName)
	}
	return nil
}

func (b *blkIOReconcile) reconcile() { _ = "STUB: not implemented"; return }

// get node local storage info

// get nodeslo

// update node blk qos by strategy defined in nodeslo

// lsr

// ls

// be

// root

// pods

// ignore unknown qos pods

type blkioUpdater struct {
	dynamicPath  string
	absolutePath string

	getDiskRecorder GetDiskRecorderFunc
	getUpdaterFunc  GetUpdaterFunc
	getRemoverFunc  GetRemoverFunc
}

// update blkio cgroup files
// podMeta == nil when BlockType is BlockTypeDevice or BlockTypeVolumeGroup
// podMeta != nil when BlockType is BlockTypePodVolume
func (b *blkIOReconcile) updateBlkIOConfig(blocks []*slov1alpha1.BlockCfg, podMeta *statesinformer.PodMeta, blkioUpdater blkioUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

// deviceName: /dev/sdb
// diskNumber: 253:16
func (b *blkIOReconcile) getDiskNumberFromDevice(deviceName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// vgName: yoda-pool
// diskNumber: 253:16
func (b *blkIOReconcile) getDiskNumberFromVolumeGroup(vgName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// volumeName is volume name of pod
// diskNumber: 253:16
func (b *blkIOReconcile) getDiskNumberFromPodVolume(podMeta *statesinformer.PodMeta, volumeName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// dynamicPath for be: kubepods.slice/kubepods-burstable.slice/
// dynamicPath for pod: kubepods.slice/kubepods-burstable.slice/kubepods-pod7712555c_ce62_454a_9e18_9ff0217b8941.slice/
func getBlkIOUpdaterFromBlockCfg(block *slov1alpha1.BlockCfg, diskNumber string, dynamicPath string) (resources []resourceexecutor.ResourceUpdater) {
	_ = "STUB: not implemented"
	return nil
}

// iops

// bps

// io weight

func (b *blkIOReconcile) getDiskNumberFromBlockCfg(block *slov1alpha1.BlockCfg, podMeta *statesinformer.PodMeta) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// check if kind of volume is pvc or csi ephemeral volume

// /var/lib/kubelet/pods/[pod uuid]/volumes/kubernetes.io~csi/[pv name]/mount

// /var/lib/kubelet/pods/[pod uuid]/volumes/kubernetes.io~csi/[pod ephemeral volume name]/mount

// configure cgroup root
// dynamicPath for root: ""
func getDiskConfigUpdaterFromBlockCfg(block *slov1alpha1.BlockCfg, diskNumber string, dynamicPath string) (resources []resourceexecutor.ResourceUpdater) {
	_ = "STUB: not implemented"
	return nil
}

// disk io weight latency

// disk io latency percent

// user cost model configuration

func parseBlkIOResult(blkioResult string) (*slov1alpha1.BlkIOQOS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// key of recorder is disk number
// value of recorder means whether to remove cgroup config of this disk
func getDiskRecorder(parentDir string, fileNames []string) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDiskNumbersFromCgroupFile(filePath string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBlkIORecorder(path string) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDiskConfigRecorder(path string) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getBlkIORemoverFromDiskNumber(diskNumber string, dynamicPath string) (resources []resourceexecutor.ResourceUpdater) {
	_ = "STUB: not implemented"
	return nil
}

func getDiskConfigRemoverFromDiskNumber(diskNumber string, dynamicPath string) (resources []resourceexecutor.ResourceUpdater) {
	_ = "STUB: not implemented"
	return nil
}

func getDiskNumber(s *metriccache.NodeLocalStorageInfo, disk string) string {
	_ = "STUB: not implemented"
	return ""
}

func getDiskByDevice(s *metriccache.NodeLocalStorageInfo, device string) string {
	_ = "STUB: not implemented"
	return ""
}

func getDiskByVG(s *metriccache.NodeLocalStorageInfo, vgName string) string {
	_ = "STUB: not implemented"
	return ""
}

func getDiskByMountPoint(s *metriccache.NodeLocalStorageInfo, mountpoint string) string {
	_ = "STUB: not implemented"
	return ""
}

// check if device is disk

// check if device is part

// check if device is lv

func isDeviceDisk(s *metriccache.NodeLocalStorageInfo, device string) bool {
	_ = "STUB: not implemented"
	return false
}
