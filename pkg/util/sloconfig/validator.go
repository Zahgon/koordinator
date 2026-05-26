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
	"sync"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var validatorInstance = &DefaultValidator{}

type DefaultValidator struct {
	once      sync.Once
	validator *validator.Validate
	trans     *ut.Translator
}

func (v *DefaultValidator) StructWithTrans(config interface{}) (validator.ValidationErrorsTranslations, error) {
	_ = "STUB: not implemented"
	return *new(validator.ValidationErrorsTranslations), nil
}

func GetValidatorInstance() *DefaultValidator { _ = "STUB: not implemented"; return nil }

func createValidator() (*validator.Validate, *ut.Translator) {
	_ = "STUB: not implemented"
	return nil, nil
}

func registerEnTranslator(instance *validator.Validate) *ut.Translator {
	_ = "STUB: not implemented"
	return nil
}
