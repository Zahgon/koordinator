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
	"k8s.io/client-go/tools/cache"
)

// ReservationToPodEventHandler can be used to handle reservation events with a pod event handler, which converts
// each reservation object into the corresponding reserve pod object.
//
//	e.g.
//	func registerReservationEventHandler(handle framework.Handle, podHandler podHandler) {
//	  extendedHandle, ok := handle.(frameworkext.ExtendedHandle)
//	  if !ok { // if not implement extendedHandle, ignore reservation events
//	    klog.V(3).Infof("registerReservationEventHandler aborted, cannot convert handle to frameworkext.ExtendedHandle, got %T", handle)
//	    return
//	  }
//	  extendedHandle.KoordinatorSharedInformerFactory().Scheduling().V1alpha1().Reservations().Informer().
//	 	 AddEventHandler(util.NewReservationToPodEventHandler(&podHandler, IsObjValidActiveReservation))
//	}
type ReservationToPodEventHandler struct {
	handler cache.ResourceEventHandler
}

var _ cache.ResourceEventHandler = &ReservationToPodEventHandler{}

func NewReservationToPodEventHandler(handler cache.ResourceEventHandler, filters ...func(obj interface{}) bool) cache.ResourceEventHandler {
	_ = "STUB: not implemented"
	return *new(cache.ResourceEventHandler)
}

func (r ReservationToPodEventHandler) OnAdd(obj interface{}, isInInitialList bool) {
	_ = "STUB: not implemented"
	return
}

// OnUpdate calls UpdateFunc if it's not nil.
func (r ReservationToPodEventHandler) OnUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// OnDelete calls DeleteFunc if it's not nil.
func (r ReservationToPodEventHandler) OnDelete(obj interface{}) { _ = "STUB: not implemented"; return }

func IsObjValidActiveReservation(obj interface{}) bool { _ = "STUB: not implemented"; return false }
