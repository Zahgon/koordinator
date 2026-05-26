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
	"k8s.io/apimachinery/pkg/conversion"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
)

func Convert_v1alpha2_LowNodeLoadArgs_To_config_LowNodeLoadArgs(in *LowNodeLoadArgs, out *config.LowNodeLoadArgs, s conversion.Scope) error {
	_ = "STUB: not implemented"
	return nil
}

func Convert_v1alpha2_CustomPriorityArgs_To_config_CustomPriorityArgs(in *CustomPriorityArgs, out *config.CustomPriorityArgs, s conversion.Scope) error {
	_ = "STUB: not implemented"
	return nil
}
