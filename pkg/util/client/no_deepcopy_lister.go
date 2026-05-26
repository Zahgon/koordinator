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

package client

import (
	"context"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	toolscache "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type noDeepCopyLister struct {
	cache  cache.Cache
	scheme *runtime.Scheme
}

func (r *noDeepCopyLister) List(ctx context.Context, out client.ObjectList, opts ...client.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

// list all objects by the field selector.  If this is namespaced and we have one, ask for the
// namespaced index key.  Otherwise, ask for the non-namespaced variant by using the fake "all namespaces"
// namespace.

// if the Limit option is set and the number of items
// listed exceeds this limit, then stop reading.

func (r *noDeepCopyLister) getIndexerByGVK(ctx context.Context, gvk schema.GroupVersionKind) (toolscache.Indexer, error) {
	_ = "STUB: not implemented"
	return *new(toolscache.Indexer), nil
}

// objectTypeForListObject tries to find the runtime.Object and associated GVK
// for a single object corresponding to the passed-in list type. We need them
// because they are used as cache map key.
func (r *noDeepCopyLister) objectTypeForListObject(list client.ObjectList) (*schema.GroupVersionKind, runtime.Object, error) {
	_ = "STUB: not implemented"
	return nil, *new(runtime.Object), nil
}

// we need the non-list GVK, so chop off the "List" from the end of the kind

// http://knowyourmeme.com/memes/this-is-fine

// requiresExactMatch checks if the given field selector is of the form `k=v` or `k==v`.
func requiresExactMatch(sel fields.Selector) (field, val string, required bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// FieldIndexName constructs the name of the index over the given field,
// for use with an indexer.
func FieldIndexName(field string) string { _ = "STUB: not implemented"; return "" }

// noNamespaceNamespace is used as the "namespace" when we want to list across all namespaces.
const allNamespacesNamespace = "__all_namespaces"

// KeyToNamespacedKey prefixes the given index key with a namespace
// for use in field selector indexes.
func KeyToNamespacedKey(ns string, baseKey string) string { _ = "STUB: not implemented"; return "" }
