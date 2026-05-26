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

package v1alpha2

// getDefaultPlugins returns the default set of plugins.
func getDefaultPlugins() *Plugins { _ = "STUB: not implemented"; return nil }

// NOTE: add default deschedule plugins here.

// mergePlugins merges the custom set into the given default one, handling disabled sets.
func mergePlugins(defaultPlugins, customPlugins *Plugins) *Plugins {
	_ = "STUB: not implemented"
	return nil
}

type pluginIndex struct {
	index  int
	plugin Plugin
}

func mergePluginSet(defaultPluginSet, customPluginSet PluginSet) PluginSet {
	_ = "STUB: not implemented"
	return *new(PluginSet)
}

// replacedPluginIndex is a set of index of plugins, which have replaced the default plugins.

// The default plugin is explicitly re-configured, update the default plugin accordingly.

// Update the default plugin in place to preserve order.

// Append all the custom plugins which haven't replaced any default plugins.
// Note: duplicated custom plugins will still be appended here.
// If so, the instantiation of descheduler framework will detect it and abort.
