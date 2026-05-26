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
	"flag"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	disableNoDeepCopy bool
)

func init() {
	flag.BoolVar(&disableNoDeepCopy, "disable-no-deepcopy", false, "If you are going to disable NoDeepCopy List in some controllers and webhooks.")
}

var _ client.NewClientFunc = NewClient

// NewClient creates the default caching client with disable deepcopy list from cache.
func NewClient(config *rest.Config, options client.Options) (client.Client, error) {
	_ = "STUB: not implemented"
	return *new(client.Client), nil
}

type delegatingClient struct {
	client.Reader
	client.Writer
	client.StatusClient
	client.SubResourceClientConstructor

	originClient client.Client
	scheme       *runtime.Scheme
	mapper       meta.RESTMapper
}

// Scheme returns the scheme this client is using.
func (d *delegatingClient) Scheme() *runtime.Scheme {
	_ = "STUB: not implemented"

	// RESTMapper returns the rest mapper this client is using.
	return nil
}

func (d *delegatingClient) RESTMapper() meta.RESTMapper {
	_ = "STUB: not implemented"

	// GroupVersionKindFor returns the GroupVersionKind for the given object.
	return *new(meta.RESTMapper)
}

func (d *delegatingClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionKind), nil
}

// IsObjectNamespaced returns true if the object is namespaced.
func (d *delegatingClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

var _ client.Reader = &delegatingReader{}

// delegatingReader forms a Reader that will cause Get and List requests for
// unstructured types to use the ClientReader while requests for any other type
// of object with use the CacheReader.  This avoids accidentally caching the
// entire cluster in the common case of loading arbitrary unstructured objects
// (e.g. from OwnerReferences).
type delegatingReader struct {
	CacheReader  client.Reader
	ClientReader client.Reader

	noDeepCopyLister *noDeepCopyLister

	uncachedGVKs      map[schema.GroupVersionKind]struct{}
	scheme            *runtime.Scheme
	cacheUnstructured bool
}

func (d *delegatingReader) shouldBypassCache(obj runtime.Object) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// TODO: this is producing unsafe guesses that don't actually work,
// but it matches ~99% of the cases out there.

// Get retrieves an obj for a given object key from the Kubernetes Cluster.
func (d *delegatingReader) Get(ctx context.Context, key client.ObjectKey, obj client.Object, option ...client.GetOption) error {
	_ = "STUB: not implemented"
	return nil
}

// List retrieves list of objects for a given namespace and list options.
func (d *delegatingReader) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

var DisableDeepCopy = disableDeepCopy{}

type disableDeepCopy struct{}

func (disableDeepCopy) ApplyToList(_ *client.ListOptions) { _ = "STUB: not implemented"; return }

func isDisableDeepCopy(opts []client.ListOption) bool { _ = "STUB: not implemented"; return false }
