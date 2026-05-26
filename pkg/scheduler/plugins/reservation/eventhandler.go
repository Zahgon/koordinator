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

package reservation

import (
	koordinatorinformers "github.com/koordinator-sh/koordinator/pkg/client/informers/externalversions"
)

type reservationEventHandler struct {
	cache       *reservationCache
	rrNominator *nominator
}

func registerReservationEventHandler(cache *reservationCache, koordinatorInformerFactory koordinatorinformers.SharedInformerFactory,
	rrNominator *nominator) {
	_ = "STUB: not implemented"
	return
}

func (h *reservationEventHandler) OnAdd(obj interface{}, isInInitialList bool) {
	_ = "STUB: not implemented"
	return
}

func (h *reservationEventHandler) OnUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// Here it is only marked that ReservationInfo is unavailable,
// and the real deletion operation is executed in deleteReservationFromCache(pkg/scheduler/frameworkext/eventhandlers/reservation_handler.go).
// This ensures that the Reserve Pod and the resources it holds are deleted correctly.
// NOTE: For the update event from available to terminated triggers the deleteReservationFromCache.

func (h *reservationEventHandler) OnDelete(obj interface{}) { _ = "STUB: not implemented"; return }

// Here it is only marked that ReservationInfo is unavailable,
// and the real deletion operation is executed in deleteReservationFromCache(pkg/scheduler/frameworkext/eventhandlers/reservation_handler.go).
// This ensures that the Reserve Pod and the resources it holds are deleted correctly.
