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

package util

import (
	"regexp"
)

type LocalStorageInfo struct {
	// mapper of disk and disk number, such as "/dev/vda":"253:0"
	DiskNumberMap map[string]string
	// mapper of disk number and disk, such as "253:0":"/dev/vda"
	NumberDiskMap map[string]string
	// mapper of partition and its disk, such as "/dev/vdb3":"/dev/vdb"
	PartitionDiskMap map[string]string
	// mapper of volumegroup and its disk, such as "yoda-pool0":"/dev/vdb"
	VGDiskMap map[string]string
	// mapper of logicalvolume and its volumegroup, such as "/dev/mapper/yoda--pool0-yoda--2c52d97f--eab6--4ac5--ba8b--242f399470e1":"yoda-pool0"
	LVMapperVGMap map[string]string
	// mapper of mountpoint and its disk, such as "/var/lib/kubelet/pods/d806ee8d-fe28-4995-a836-d2356d44ec5f/volumes/kubernetes.io~csi/yoda-2c52d97f-eab6-4ac5-ba8b-242f399470e1/mount":"/dev/mapper/yoda--pool0-yoda--2c52d97f--eab6--4ac5--ba8b--242f399470e1"
	MPDiskMap map[string]string
}

var (
	deviceName    = "/dev/%s"
	lvmMapperName = "/dev/mapper/%s-%s"

	lsblkRE   = regexp.MustCompile(`([A-Z:]+)=(?:"(.*?)")`)
	vgsRE     = regexp.MustCompile(`[^\s]+`)
	lvsRE     = regexp.MustCompile(`[^\s]+`)
	findmntRE = regexp.MustCompile(`([A-Z:]+)=(?:"(.*?)")`)

	lsblkColumns = []string{
		"NAME",
		"TYPE",
		"MAJ:MIN",
	}
	vgsColumns = []string{
		"vg_name",
		"pv_count",
		"pv_name",
	}
	lvsColumns = []string{
		"lv_name",
		"vg_name",
	}
	findmntColumns = []string{
		"TARGET",
		"SOURCE",
	}
)

func (s *LocalStorageInfo) scanDevices() error {
	_ = "STUB: not implemented"
	// A Cmd cannot be reused after calling its Run, Output or CombinedOutput methods.
	// Or it will report error: Stdout already set
	return nil
}

// output fields as key=value pairs

func (s *LocalStorageInfo) scanVolumeGroups() error { _ = "STUB: not implemented"; return nil }

func (s *LocalStorageInfo) scanLogicalVolumes() error { _ = "STUB: not implemented"; return nil }

func (s *LocalStorageInfo) scanMountPoints() error { _ = "STUB: not implemented"; return nil }

func (s *LocalStorageInfo) scanDevicesOutput(output []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// text: [[NAME="vdd" NAME vdd] [TYPE="disk" TYPE disk] [MAJ:MIN="253:48" MAJ:MIN 253:48]]

func (s *LocalStorageInfo) scanVolumeGroupsOutput(output []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// text: [yoda-pool0 1 /dev/vdc3]

func (s *LocalStorageInfo) scanLogicalVolumesOutput(output []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// text: [yoda-15199981-7229-45e9-b3a0-b5b30a6a162b yoda-pool0]

func (s *LocalStorageInfo) isDeviceDisk(device string) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *LocalStorageInfo) scanMountPointOutput(output []byte) error {
	_ = "STUB: not implemented"
	// text: [[TARGET="/var/lib/kubelet/pods/0ef5bd7a-aa83-4242-8597-c7ab4afaf356/volumes/kubernetes.io~csi/yoda-15199981-7229-45e9-b3a0-b5b30a6a162b/mount" TARGET /var/lib/kubelet/pods/0ef5bd7a-aa83-4242-8597-c7ab4afaf356/volumes/kubernetes.io~csi/yoda-15199981-7229-45e9-b3a0-b5b30a6a162b/mount] [SOURCE="/dev/mapper/yoda--pool0-yoda--15199981--7229--45e9--b3a0--b5b30a6a162b" SOURCE /dev/mapper/yoda--pool0-yoda--15199981--7229--45e9--b3a0--b5b30a6a162b]]
	return nil
}

func GetLocalStorageInfo() (*LocalStorageInfo, error) { _ = "STUB: not implemented"; return nil, nil }

func getDiskNumberFromDeviceNumber(number string) string { _ = "STUB: not implemented"; return "" }

func getDeviceName(device string) string { _ = "STUB: not implemented"; return "" }

func getLVMMapperName(vgName, lvName string) string { _ = "STUB: not implemented"; return "" }
