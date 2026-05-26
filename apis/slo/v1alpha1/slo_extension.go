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

package v1alpha1

// ExtensionsMap if used for third-party plugins
// +k8s:deepcopy-gen=false
type ExtensionsMap struct {
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Object map[string]interface{} `json:",inline"`
}

func (in *ExtensionsMap) DeepCopyInto(out *ExtensionsMap) { _ = "STUB: not implemented"; return }

func (in *ExtensionsMap) DeepCopy() *ExtensionsMap { _ = "STUB: not implemented"; return nil }
