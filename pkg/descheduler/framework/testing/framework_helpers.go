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

package testing

import (
	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config/scheme"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework/runtime"
)

var configDecoder = scheme.Codecs.UniversalDecoder()

// NewFramework creates a Framework from the register functions and options.
func NewFramework(fns []RegisterPluginFunc, profileName string, opts ...runtime.Option) (framework.Handle, error) {
	_ = "STUB: not implemented"
	return *new(framework.Handle), nil
}

type RegisterPluginFunc func(reg *runtime.Registry, profile *deschedulerconfig.DeschedulerProfile)

func RegisterDeschedulePlugin(pluginName string, pluginNewFunc runtime.PluginFactory) RegisterPluginFunc {
	_ = "STUB: not implemented"
	return *new(RegisterPluginFunc)
}

func RegisterBalancePlugin(pluginName string, pluginNewFunc runtime.PluginFactory) RegisterPluginFunc {
	_ = "STUB: not implemented"
	return *new(RegisterPluginFunc)
}

func RegisterEvictPlugin(pluginName string, pluginNewFunc runtime.PluginFactory) RegisterPluginFunc {
	_ = "STUB: not implemented"
	return *new(RegisterPluginFunc)
}

func RegisterFilterPlugin(pluginName string, pluginNewFunc runtime.PluginFactory) RegisterPluginFunc {
	_ = "STUB: not implemented"
	return *new(RegisterPluginFunc)
}

// RegisterPluginAsExtensions returns a function to register a Plugin as given extensionPoints to a given registry.
func RegisterPluginAsExtensions(pluginName string, pluginNewFunc runtime.PluginFactory, extensions ...string) RegisterPluginFunc {
	_ = "STUB: not implemented"
	return *new(RegisterPluginFunc)
}

// RegisterPluginAsExtensionsWithWeight returns a function to register a Plugin as given extensionPoints with weight to a given registry.
func RegisterPluginAsExtensionsWithWeight(pluginName string, weight int32, pluginNewFunc runtime.PluginFactory, extensions ...string) RegisterPluginFunc {
	_ = "STUB: not implemented"
	return *new(RegisterPluginFunc)
}

// Use defaults from latest config API version.

func getPluginSetByExtension(plugins *deschedulerconfig.Plugins, extension string) *deschedulerconfig.PluginSet {
	_ = "STUB: not implemented"
	return nil
}
