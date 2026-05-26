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
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/apis/config"
)

func ValidateDeschedulerConfiguration(cc *config.DeschedulerConfiguration) utilerrors.Aggregate {
	_ = "STUB: not implemented"
	return *new(utilerrors.Aggregate)
}

// isValidSocketAddr checks whether addr is a valid host:port or ip:port socket address.
func isValidSocketAddr(addr string) []string { _ = "STUB: not implemented"; return nil }

// try to resolve as hostname

func validateDeschedulerProfile(path *field.Path, profile *config.DeschedulerProfile) []error {
	_ = "STUB: not implemented"
	return nil
}

func validatePluginConfig(path *field.Path, profile *config.DeschedulerProfile) []error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: you can add the in-tree plugins configuration validation function

// type mismatch, no need to validate the `args`.

// It's possible that validation function return a Aggregate, just append here and it will be flattened at the end of CC validation.
