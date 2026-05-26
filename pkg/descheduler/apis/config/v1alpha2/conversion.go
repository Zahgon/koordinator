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

import (
	"sync"

	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
)

var (
	// pluginArgConversionScheme is a scheme with internal and v1alpha2 registered,
	// used for defaulting/converting typed PluginConfig Args.
	// Access via getPluginArgConversionScheme()
	pluginArgConversionScheme     *runtime.Scheme
	initPluginArgConversionScheme sync.Once
)

func GetPluginArgConversionScheme() *runtime.Scheme { _ = "STUB: not implemented"; return nil }

// set up the scheme used for plugin arg conversion

func Convert_v1alpha2_DeschedulerConfiguration_To_config_DeschedulerConfiguration(in *DeschedulerConfiguration, out *config.DeschedulerConfiguration, s conversion.Scope) error {
	_ = "STUB: not implemented"
	return nil
}

// convertToInternalPluginConfigArgs converts PluginConfig#Args into internal
// types using a scheme, after applying defaults.
func convertToInternalPluginConfigArgs(out *config.DeschedulerConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}

func Convert_config_DeschedulerConfiguration_To_v1alpha2_DeschedulerConfiguration(in *config.DeschedulerConfiguration, out *DeschedulerConfiguration, s conversion.Scope) error {
	_ = "STUB: not implemented"
	return nil
}

// convertToExternalPluginConfigArgs converts PluginConfig#Args into
// external (versioned) types using a scheme.
func convertToExternalPluginConfigArgs(out *DeschedulerConfiguration) error {
	_ = "STUB: not implemented"
	return nil
}
