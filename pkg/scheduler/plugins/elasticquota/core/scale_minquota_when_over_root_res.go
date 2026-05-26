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
)

// ScaleMinQuotaManager The child nodes under each node will be divided into two categories, one allows
// scaling min quota, and the other does not allow. When our total resources are insufficient, the min quota of child
// nodes will be proportionally reduced. In order to calculate the scaling of the min quota, we will count the sum of
// the min quota of the child nodes of each node in advance. The quota of the child nodes that allow scaling is stored
// in enableScaleSubsSumMinQuotaMap, and the quota of child nodes that do not allow scaling is stored in
// disableScaleSubsSumMinQuotaMap
type ScaleMinQuotaManager struct {
	lock sync.RWMutex
	// enableScaleSubsSumMinQuotaMap key: quotaName, val: sum of its enableScale children's minQuota
	enableScaleSubsSumMinQuotaMap map[string]v1.ResourceList
	// disableScaleSubsSumMinQuotaMap key: quotaName, val: sum of its disableScale children's minQuota
	disableScaleSubsSumMinQuotaMap map[string]v1.ResourceList
	// originalMinQuotaMap stores the original minQuota, when children's sum minQuota is smaller than the
	// totalRes, just return the originalMinQuota.
	originalMinQuotaMap         map[string]v1.ResourceList
	quotaEnableMinQuotaScaleMap map[string]bool
}

func NewScaleMinQuotaManager() *ScaleMinQuotaManager { _ = "STUB: not implemented"; return nil }

func (s *ScaleMinQuotaManager) update(parQuotaName, subQuotaName string, subMinQuota v1.ResourceList, enableScaleMinQuota bool) {
	_ = "STUB: not implemented"
	return
}

// step1: delete the oldMinQuota if present

// step2: add the newMinQuota

// step3: record the newMinQuota

func (s *ScaleMinQuotaManager) remove(parQuotaName, subQuotaName string) {
	_ = "STUB: not implemented"
	return
}

func (s *ScaleMinQuotaManager) getScaledMinQuota(newTotalRes v1.ResourceList, parQuotaName, subQuotaName string) (bool, v1.ResourceList) {
	_ = "STUB: not implemented"
	return false, *new(v1.ResourceList)
}

// get the dimensions where children's minQuota sum is larger than newTotalRes

//  children's minQuota sum is smaller than totalRes in all dimensions

// ensure the disableScale children's minQuota first

// if still has minQuota left, enableScaleMinQuota children partition it according to their minQuotaValue.
