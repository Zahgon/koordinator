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

package options

import (
	"flag"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type Options struct {
	ControllerAddFuncs  map[string]func(manager.Manager) error
	Controllers         []string
	ControllerInitFlags map[string]func(*flag.FlagSet)
}

func NewOptions() *Options { _ = "STUB: not implemented"; return nil }

func (o *Options) InitFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }

func (o *Options) ApplyTo(m manager.Manager) error { _ = "STUB: not implemented"; return nil }

func isControllerEnabled(controllerName string, controllers []string) bool {
	_ = "STUB: not implemented"
	return false
}
