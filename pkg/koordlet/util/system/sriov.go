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

package system

const (
	configuredVfFile = "sriov_numvfs"
)

// SriovConfigured returns true if sriov_numvfs reads > 0 else false
func SriovConfigured(addr string) bool { _ = "STUB: not implemented"; return false }

func extractNumber(pfDir string, s string) int { _ = "STUB: not implemented"; return 0 }

// GetVFList returns a List containing PCI addr for all VF discovered in a given PF
func GetVFList(pf string) (vfList []string, err error) { _ = "STUB: not implemented"; return nil, nil }

//TODO 排序

// Read all VF directory and get add VF PCI addr to the vfList

// GetVConfigured returns number of VF configured for a PF
func GetVConfigured(pf string) int { _ = "STUB: not implemented"; return 0 }

// IsSriovVF check if a pci device has link to a PF
func IsSriovVF(pciAddr string) bool { _ = "STUB: not implemented"; return false }
