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

import (
	"k8s.io/apimachinery/pkg/util/sets"
)

/*
https://lkml.org/lkml/2020/9/7/597

Generic Initiators are a new ACPI concept that allows for the
description of proximity domains that contain a device which
performs memory access (such as a network card) but neither
host CPU nor Memory.

This file provides some GI acquisition and parsing logic.
*/

const (
	SysHasGenericInitiator = "devices/system/node/has_generic_initiator"
)

func GetHasGenericInitiatorPath() string { _ = "STUB: not implemented"; return "" }

func GetNUMANodesHasGI() sets.Set[int32] { _ = "STUB: not implemented"; return nil }

func parseIDs(s string) (sets.Set[int32], error) {
	_ = "STUB: not implemented"
	return nil,

		// Handle empty string.
		nil
}

// Split CPU list string:
// "0-5,34,46-48 => ["0-5", "34", "46-48"]

// Handle ranges that consist of only one element like "34".
// assert cpu id is in range of int32

// Handle multi-element ranges like "0-5".
// assert cpu id is in range of int32

// Add all elements to the result.
// e.g. "0-5", "46-48" => [0, 1, 2, 3, 4, 5, 46, 47, 48].
