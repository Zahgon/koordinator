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

package generator

import (
	"time"
)

// ValidCACert think cert and key are valid if they meet the following requirements:
// - key and cert are valid pair
// - caCert is the root ca of cert
// - cert is for dnsName
// - cert won't expire before time
func ValidCACert(key, cert, caCert []byte, dnsName string, time time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// Verify key and cert are valid pair

// Verify cert is valid for at least 1 year.
