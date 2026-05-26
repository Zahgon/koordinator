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

package sloconfig

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/koordinator-sh/koordinator/apis/configuration"
	"github.com/koordinator-sh/koordinator/apis/slo/v1alpha1"
)

const (
	// InvalidPercentageValueMsg is an error message for value must in percentage range.
	InvalidPercentageValueMsg string = `must be percentage value just in 0-100'`
)

// NetDirection represents the stored type of IngressOrEgress
type NetDirection int64

const (
	Ingress NetDirection = iota // The IngressOrEgress holds an ingress.
	Egress                      // The IngressOrEgress holds an egress.
)

var _ ConfigChecker = &ResourceQOSChecker{}

type ResourceQOSChecker struct {
	cfg    *configuration.ResourceQOSCfg
	sysCfg *configuration.SystemCfg
	CommonChecker
}

func NewResourceQOSChecker(oldConfig, newConfig *corev1.ConfigMap, needUnmarshal bool) *ResourceQOSChecker {
	_ = "STUB: not implemented"
	return nil
}

func (c *ResourceQOSChecker) ConfigParamValid() error { _ = "STUB: not implemented"; return nil }

func CheckNetQosCfg(cfg *configuration.ResourceQOSCfg, sysCfg *configuration.SystemCfg, fldPath *field.Path) error {
	_ = "STUB: not implemented"
	return nil
}

func CheckEachElementIsValid(cfg *configuration.ResourceQOSCfg, fldPath *field.Path) error {
	_ = "STUB: not implemented"
	return nil
}

// check netqos config is valid for cluster strategy

// check netqos config is valid for each node strategy

func CheckTotalPercentageIsValid(cfg *configuration.ResourceQOSCfg) error {
	_ = "STUB: not implemented"
	// check netqos config is valid for cluster strategy
	return nil
}

// check netqos config is valid for each node strategy

func doSubstration(strategy *v1alpha1.ResourceQOSStrategy, totalNetBandWidth int64) (leftIngress, leftEgress int64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func CheckTotalQuantityIsValid(cfg *configuration.ResourceQOSCfg, sysCfg *configuration.SystemCfg) error {
	_ = "STUB: not implemented"
	// check netqos config is valid for cluster strategy
	return nil
}

// check netqos config is valid for each node strategy

func getTotalNetBandWidth(labelSelector *metav1.LabelSelector, sysCfg *configuration.SystemCfg) *resource.Quantity {
	_ = "STUB: not implemented"
	return nil
}

func CheckOnlyOneFormat(cfg *configuration.ResourceQOSCfg) error {
	_ = "STUB: not implemented"
	return nil
}

// check netqos config is valid for cluster strategy

// check netqos config is valid for each node strategy

func CheckSubNetQos(fldPath *field.Path, qos *v1alpha1.NetworkQOSCfg) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidatePercentageOrQuantity tests if a given value is a valid percentage or
// quantity (defines in: https://github.com/kubernetes/apimachinery/blob/master/pkg/api/resource/quantity.go#L100-L111).
func ValidatePercentageOrQuantity(intOrPercent *intstr.IntOrString, fldPath *field.Path) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// ValidateQuantityField validates that given value is quantity format, like 50M.
func ValidateQuantityField(quantityStr string, fldPath *field.Path) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// ValidatePercentageField validates that given value is a percentage format just in 0-100.
func ValidatePercentageField(value int, fldPath *field.Path) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

func (c *ResourceQOSChecker) initConfig() error { _ = "STUB: not implemented"; return nil }

func (c *ResourceQOSChecker) getConfigProfiles() []configuration.NodeCfgProfile {
	_ = "STUB: not implemented"
	return nil
}
