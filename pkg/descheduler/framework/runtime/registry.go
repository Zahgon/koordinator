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

package runtime

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

type PluginFactory func(ctx context.Context, args runtime.Object, handle framework.Handle) (framework.Plugin, error)

// DecodeInto decodes configuration whose type is *runtime.Unknown to the interface into.
func DecodeInto(obj runtime.Object, into interface{}) error { _ = "STUB: not implemented"; return nil }

// If ContentType is empty, it means ContentTypeJSON by default.

type Registry map[string]PluginFactory

// Register adds a new plugin to the registry. If a plugin with the same name
// exists, it returns an error.
func (r Registry) Register(name string, factory PluginFactory) error {
	_ = "STUB: not implemented"
	return nil
}

// Unregister removes an existing plugin from the registry. If no plugin with
// the provided name exists, it returns an error.
func (r Registry) Unregister(name string) error { _ = "STUB: not implemented"; return nil }

func (r Registry) Merge(in Registry) error { _ = "STUB: not implemented"; return nil }
