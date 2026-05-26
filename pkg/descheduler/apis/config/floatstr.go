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

package config

type Float64OrString struct {
	Type     Type
	FloatVal float64
	StrVal   string
}

// Type represents the stored type of Float64OrString.
type Type int64

const (
	Float  Type = iota // The Float64OrString holds a Float.
	String             // The Float64OrString holds a string.
)

// UnmarshalJSON implements the json.Unmarshaller interface.
func (floatstr *Float64OrString) UnmarshalJSON(value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// String returns the string value, or the Itoa of the float value.
func (floatstr *Float64OrString) String() string { _ = "STUB: not implemented"; return "" }

// FloatValue returns the FloatVal if type Float, or if
// it is a String, will attempt a conversion to float64,
// returning 0 if a parsing error occurs.
func (floatstr *Float64OrString) FloatValue() float64 { _ = "STUB: not implemented"; return 0 }

// MarshalJSON implements the json.Marshaller interface.
func (floatstr *Float64OrString) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
