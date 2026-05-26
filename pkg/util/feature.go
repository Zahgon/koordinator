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

package util

import (
	"k8s.io/component-base/featuregate"
)

// RunFeature runs moduleFunc only if interval > 0 AND at least one feature dependency is enabled
func RunFeature(moduleFunc func(), featureDependency []featuregate.Feature, interval int, stopCh <-chan struct{}) bool {
	_ = "STUB: not implemented"
	return false
}

// RunFeatureWithInit runs moduleFunc only if interval > 0 , at least one feature dependency is enabled
// and moduleInit function returns nil
func RunFeatureWithInit(moduleInit func() error, moduleFunc func(), featureDependency []featuregate.Feature, interval int, stopCh <-chan struct{}) bool {
	_ = "STUB: not implemented"
	return false
}
