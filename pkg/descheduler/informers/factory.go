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

package informers

import (
	"reflect"
	"sync"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/informers/admissionregistration"
	"k8s.io/client-go/informers/apiserverinternal"
	"k8s.io/client-go/informers/apps"
	"k8s.io/client-go/informers/autoscaling"
	"k8s.io/client-go/informers/batch"
	"k8s.io/client-go/informers/certificates"
	"k8s.io/client-go/informers/coordination"
	"k8s.io/client-go/informers/core"
	"k8s.io/client-go/informers/discovery"
	"k8s.io/client-go/informers/events"
	"k8s.io/client-go/informers/extensions"
	"k8s.io/client-go/informers/flowcontrol"
	"k8s.io/client-go/informers/internalinterfaces"
	"k8s.io/client-go/informers/networking"
	"k8s.io/client-go/informers/node"
	"k8s.io/client-go/informers/policy"
	"k8s.io/client-go/informers/rbac"
	"k8s.io/client-go/informers/resource"
	"k8s.io/client-go/informers/scheduling"
	"k8s.io/client-go/informers/storage"
	storagemigration "k8s.io/client-go/informers/storagemigration"
	"k8s.io/client-go/tools/cache"
	ctrl "sigs.k8s.io/controller-runtime"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*sharedInformerFactory) *sharedInformerFactory

type sharedInformerFactory struct {
	manager          ctrl.Manager
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	lock      sync.Mutex
	informers map[reflect.Type]cache.SharedIndexInformer
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[v1.Object]time.Duration) SharedInformerOption {
	_ = "STUB: not implemented"
	return *new(SharedInformerOption)
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	_ = "STUB: not implemented"
	return *new(SharedInformerOption)
}

// WithNamespace limits the SharedInformerFactory to the specified namespace.
func WithNamespace(namespace string) SharedInformerOption {
	_ = "STUB: not implemented"
	return *new(SharedInformerOption)
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory for all namespaces.
func NewSharedInformerFactory(mgr ctrl.Manager, defaultResync time.Duration) informers.SharedInformerFactory {
	_ = "STUB: not implemented"
	return *new(informers.SharedInformerFactory)
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(mgr ctrl.Manager, defaultResync time.Duration, options ...SharedInformerOption) informers.SharedInformerFactory {
	_ = "STUB: not implemented"
	return *new(informers.SharedInformerFactory)
}

// Apply all options

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	_ = "STUB: not implemented"

	// WaitForCacheSync waits for all started informers' cache were synced.
	return
}

func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	_ = "STUB: not implemented"
	return nil
}

// InformerFor returns the SharedIndexInformer for obj using an internal client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedIndexInformer)
}

func (f *sharedInformerFactory) Shutdown() { _ = "STUB: not implemented"; return }

func (f *sharedInformerFactory) Admissionregistration() admissionregistration.Interface {
	_ = "STUB: not implemented"
	return *new(admissionregistration.Interface)
}

func (f *sharedInformerFactory) Internal() apiserverinternal.Interface {
	_ = "STUB: not implemented"
	return *new(apiserverinternal.Interface)
}

func (f *sharedInformerFactory) Apps() apps.Interface {
	_ = "STUB: not implemented"
	return *new(apps.Interface)
}

func (f *sharedInformerFactory) Autoscaling() autoscaling.Interface {
	_ = "STUB: not implemented"
	return *new(autoscaling.Interface)
}

func (f *sharedInformerFactory) Batch() batch.Interface {
	_ = "STUB: not implemented"
	return *new(batch.Interface)
}

func (f *sharedInformerFactory) Certificates() certificates.Interface {
	_ = "STUB: not implemented"
	return *new(certificates.Interface)
}

func (f *sharedInformerFactory) Coordination() coordination.Interface {
	_ = "STUB: not implemented"
	return *new(coordination.Interface)
}

func (f *sharedInformerFactory) Core() core.Interface {
	_ = "STUB: not implemented"
	return *new(core.Interface)
}

func (f *sharedInformerFactory) Discovery() discovery.Interface {
	_ = "STUB: not implemented"
	return *new(discovery.Interface)
}

func (f *sharedInformerFactory) Events() events.Interface {
	_ = "STUB: not implemented"
	return *new(events.Interface)
}

func (f *sharedInformerFactory) Extensions() extensions.Interface {
	_ = "STUB: not implemented"
	return *new(extensions.Interface)
}

func (f *sharedInformerFactory) Flowcontrol() flowcontrol.Interface {
	_ = "STUB: not implemented"
	return *new(flowcontrol.Interface)
}

func (f *sharedInformerFactory) Networking() networking.Interface {
	_ = "STUB: not implemented"
	return *new(networking.Interface)
}

func (f *sharedInformerFactory) Node() node.Interface {
	_ = "STUB: not implemented"
	return *new(node.Interface)
}

func (f *sharedInformerFactory) Policy() policy.Interface {
	_ = "STUB: not implemented"
	return *new(policy.Interface)
}

func (f *sharedInformerFactory) Rbac() rbac.Interface {
	_ = "STUB: not implemented"
	return *new(rbac.Interface)
}

func (f *sharedInformerFactory) Resource() resource.Interface {
	_ = "STUB: not implemented"
	return *new(resource.Interface)
}

func (f *sharedInformerFactory) Scheduling() scheduling.Interface {
	_ = "STUB: not implemented"
	return *new(scheduling.Interface)
}

func (f *sharedInformerFactory) Storage() storage.Interface {
	_ = "STUB: not implemented"
	return *new(storage.Interface)
}

func (f *sharedInformerFactory) Storagemigration() storagemigration.Interface {
	_ = "STUB: not implemented"
	return *new(storagemigration.Interface)
}
