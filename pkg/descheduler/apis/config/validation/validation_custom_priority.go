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

package validation

import (
	deschedulerconfig "github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
)

// ValidateCustomPriorityArgs validates the CustomPriorityArgs
func ValidateCustomPriorityArgs(args *deschedulerconfig.CustomPriorityArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate each priority level

// Check for duplicate priority names

// Validate pod selectors

// Validate mode

// empty handled by defaulting to BestEffort

// validateResourcePriority validates a single resource priority
func validateResourcePriority(priority deschedulerconfig.ResourcePriority, index int) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate thresholds if provided

// validatePodSelector validates a pod selector
func validatePodSelector(selector deschedulerconfig.CustomPriorityPodSelector) error {
	_ = "STUB: not implemented"
	return nil
}

// validateResourceThresholds validates resource thresholds
func validateResourceThresholds(thresholds deschedulerconfig.ResourceThresholds) error {
	_ = "STUB: not implemented"
	return nil
}
