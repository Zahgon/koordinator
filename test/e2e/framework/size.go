/*
Copyright 2015 The Kubernetes Authors.

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

package framework

// ResizeGroup resizes an instance group
func ResizeGroup(group string, size int32) error { _ = "STUB: not implemented"; return nil }

// GetGroupNodes returns a node name for the specified node group
func GetGroupNodes(group string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// GroupSize returns the size of an instance group
func GroupSize(group string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WaitForGroupSize waits for node instance group reached the desired size
func WaitForGroupSize(group string, size int32) error { _ = "STUB: not implemented"; return nil }
