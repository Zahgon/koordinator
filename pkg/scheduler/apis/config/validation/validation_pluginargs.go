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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"
	schedconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"

	"github.com/koordinator-sh/koordinator/apis/extension"
	"github.com/koordinator-sh/koordinator/pkg/scheduler/apis/config"
)

// ValidateLoadAwareSchedulingArgs validates that LoadAwareSchedulingArgs are correct.
func ValidateLoadAwareSchedulingArgs(args *config.LoadAwareSchedulingArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func validateAggregatedArgs(
	aggregated *config.LoadAwareSchedulingAggregatedArgs,
	fldPath *field.Path,
) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func validateAggregationType(aggType extension.AggregationType, fldPath *field.Path) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateResourceWeights(resources map[corev1.ResourceName]int64) error {
	_ = "STUB: not implemented"
	return nil
}

func validateResourceThresholds(thresholds map[corev1.ResourceName]int64) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEstimatedScalingFactors(scalingFactors map[corev1.ResourceName]int64) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateElasticQuotaArgs(elasticArgs *config.ElasticQuotaArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateCoschedulingArgs(coeSchedulingArgs *config.CoschedulingArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func validateResources(resources []schedconfig.ResourceSpec, p *field.Path) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func ValidateDeviceShareArgs(path *field.Path, args *config.DeviceShareArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateReservationArgs(path *field.Path, args *config.ReservationArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateNodeNUMAResourceArgs(path *field.Path, args *config.NodeNUMAResourceArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateSchedulingHintArgs validates that SchedulingHintArgs are correct.
func ValidateSchedulingHintArgs(path *field.Path, args *config.SchedulingHintArgs) error {
	_ = "STUB: not implemented"
	return nil
}
