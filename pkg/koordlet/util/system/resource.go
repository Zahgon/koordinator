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

package system

const ErrResourceUnsupportedPrefix = "resource is unsupported"

type ResourceType string

type Resource interface {
	// ResourceType is the type of system resource. e.g. "cpu.cfs_quota_us", "cpu.cfs_period_us", "schemata"
	ResourceType() ResourceType
	// Path is the generated system file path according to the given parent directory.
	// e.g. "/host-cgroup/kubepods/kubepods-podxxx/cpu.shares"
	Path(dynamicPath string) string
	// IsSupported checks whether the system resource is supported in current platform
	IsSupported(dynamicPath string) (bool, string)
	// IsValid checks whether the given value is valid for the system resource's content
	IsValid(v string) (bool, string)
	// WithValidator sets the ResourceValidator for the resource
	WithValidator(validator ResourceValidator) Resource
	// WithSupported sets the Supported status of the resource when it is initialized.
	WithSupported(supported bool, msg string) Resource
	// WithCheckSupported sets the check function for the Supported status of given resource and parent directory.
	WithCheckSupported(checkSupportedFn func(r Resource, dynamicPath string) (isSupported bool, msg string)) Resource
	// WithCheckOnce sets the check function only checking once and then use the result as the Supported status.
	WithCheckOnce(isCheckOnce bool) Resource
}

func GetDefaultResourceType(subfs string, filename string) ResourceType {
	_ = "STUB: not implemented"
	return *new(ResourceType)
}

func ValidateResourceValue(value *int64, dynamicPath string, r Resource) bool {
	_ = "STUB: not implemented"
	return false
}

func ResourceUnsupportedErr(msg string) error { _ = "STUB: not implemented"; return nil }

func IsResourceUnsupportedErr(err error) bool { _ = "STUB: not implemented"; return false }

func SupportedIfFileExistsInKubepods(r Resource, _ string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func SupportedIfFileExists(r Resource, dynamicPath string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func SupportedIfFileExistsInRootCgroup(filename string, subfs string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func CheckIfAllSupported(checkSupportedFns ...func() (bool, string)) func() (bool, string) {
	_ = "STUB: not implemented"
	return nil
}
