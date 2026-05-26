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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	_ = "STUB: not implemented"

	// Lister returns the GenericLister.
	return *new(cache.SharedIndexInformer)
}

func (f *genericInformer) Lister() cache.GenericLister {
	_ = "STUB: not implemented"
	return *new(cache.GenericLister)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error) {
	_ = "STUB: not implemented"

	// Group=admissionregistration.k8s.io, Version=v1
	return *new(informers.GenericInformer), nil
}

// Group=admissionregistration.k8s.io, Version=v1alpha1

// Group=admissionregistration.k8s.io, Version=v1beta1

// Group=apps, Version=v1

// Group=apps, Version=v1beta1

// Group=apps, Version=v1beta2

// Group=autoscaling, Version=v1

// Group=autoscaling, Version=v2

// Group=batch, Version=v1

// Group=certificates.k8s.io, Version=v1

// Group=certificates.k8s.io, Version=v1alpha1

// Group=certificates.k8s.io, Version=v1beta1

// Group=coordination.k8s.io, Version=v1

// Group=coordination.k8s.io, Version=v1beta1

// Group=core, Version=v1

// Group=discovery.k8s.io, Version=v1

// Group=discovery.k8s.io, Version=v1beta1

// Group=events.k8s.io, Version=v1

// Group=events.k8s.io, Version=v1beta1

// Group=extensions, Version=v1beta1

// Group=flowcontrol.apiserver.k8s.io, Version=v1beta1

// Group=flowcontrol.apiserver.k8s.io, Version=v1beta2

// Group=flowcontrol.apiserver.k8s.io, Version=v1beta3

// Group=internal.apiserver.k8s.io, Version=v1alpha1

// Group=networking.k8s.io, Version=v1

// Group=networking.k8s.io, Version=v1beta1

// Group=node.k8s.io, Version=v1

// Group=node.k8s.io, Version=v1alpha1

// Group=node.k8s.io, Version=v1beta1

// Group=policy, Version=v1

// Group=rbac.authorization.k8s.io, Version=v1

// Group=rbac.authorization.k8s.io, Version=v1alpha1

// Group=rbac.authorization.k8s.io, Version=v1beta1

// Group=resource.k8s.io, Version=v1beta1

// Group=scheduling.k8s.io, Version=v1

// Group=scheduling.k8s.io, Version=v1alpha1

// Group=scheduling.k8s.io, Version=v1beta1

// Group=storage.k8s.io, Version=v1

// Group=storage.k8s.io, Version=v1alpha1

// Group=storagemigration.k8s.io, Version=v1beta1

// Group=storage.k8s.io, Version=v1beta1
