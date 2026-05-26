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

package frameworkext

import (
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	schedulingv1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

type ReservationCache interface {
	DeleteReservation(r *schedulingv1alpha1.Reservation) *ReservationInfo
	GetReservationInfoByPod(pod *corev1.Pod, nodeName string) *ReservationInfo
}

// TODO(joseph): Considering the amount of changed code,
// temporarily use global variable to store ReservationCache instance,
// and then refactor to separate ReservationCache later.
var reservationCacheMap = &sync.Map{}

func GetAllReservationCaches() map[string]ReservationCache { _ = "STUB: not implemented"; return nil }

func SetReservationCache(cache ReservationCache, profileName string) {
	_ = "STUB: not implemented"
	return
}

func ClearReservationCache() { _ = "STUB: not implemented"; return }

var _ ReservationCache = &FakeReservationCache{}

type FakeReservationCache struct {
	RInfo *ReservationInfo
}

func NewFakeReservationCache() *FakeReservationCache { _ = "STUB: not implemented"; return nil }

func (f *FakeReservationCache) DeleteReservation(r *schedulingv1alpha1.Reservation) *ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeReservationCache) GetReservationInfoByPod(pod *corev1.Pod, nodeName string) *ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeReservationCache) GetReservationInfo(uid types.UID) *ReservationInfo {
	_ = "STUB: not implemented"
	return nil
}
