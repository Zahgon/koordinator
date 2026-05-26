/*
Copyright 2022 The Kubernetes Authors.

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

package image

import (
	"regexp"
)

// All of the image tags are of the format k8s.gcr.io/sig-storage/hostpathplugin:v1.7.3.
var imageRE = regexp.MustCompile(`^(.*)/([^/:]*):(.*)$`)

// appendCSIImageConfigs extracts image repo, name and version from
// the YAML files under test/e2e/testing-manifests/storage-csi and
// creates new config entries  for them.
func appendCSIImageConfigs(configs map[int]Config) { _ = "STUB: not implemented"; return }

// We add our images with index numbers that start after the highest existing number.

// Split at the "---" separator before working on
// individual item. Only works for .yaml.
//
// We need to split ourselves because we need access
// to each original chunk of data for
// runtime.DecodeInto. kubectl has its own
// infrastructure for this, but that is a lot of code
// with many dependencies.

// We don't care what the actual type is. We just
// unmarshal into generic maps and then look up images.

// Will be called for all image strings.

// These paths match plain Pods and more complex types
// like Deployments.

// findStrings recursively decends into an object along a certain path.  Path
// elements are the named fields. If a field references a list, each of the
// list elements will be followed.
//
// Conceptually this is similar to a JSON path.
func findStrings(object interface{}, visit func(value string), path ...string) {
	_ = "STUB: not implemented"
	return

	// Found it. May or may not be a string, though.
}

// If we are in a list, check each entry.

// Follow path if possible
