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

package helper

import (
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/informers/core"
	"k8s.io/client-go/informers/internalinterfaces"
	"k8s.io/client-go/informers/storage"
	"k8s.io/client-go/tools/cache"
)

// forceSyncSharedInformerFactory wraps SharedInformerFactory so that every informer
// obtained via InformerFor is wrapped with forceSyncsharedIndexInformer, which
// intercepts AddEventHandler / AddEventHandlerWithResyncPeriod calls and collects
// the returned ResourceEventHandlerRegistration for WaitForHandlersSync.

var _ informers.SharedInformerFactory = &forceSyncSharedInformerFactory{}

type forceSyncSharedInformerFactory struct {
	informers.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	defaultResync    time.Duration
}

func NewForceSyncSharedInformerFactory(factory informers.SharedInformerFactory) informers.SharedInformerFactory {
	_ = "STUB: not implemented"
	return *new(informers.SharedInformerFactory)
}

func (f *forceSyncSharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedIndexInformer)
}

func (f *forceSyncSharedInformerFactory) Core() core.Interface {
	_ = "STUB: not implemented"
	return *new(core.Interface)
}

func (f *forceSyncSharedInformerFactory) Storage() storage.Interface {
	_ = "STUB: not implemented"
	return *new(storage.Interface)
}
