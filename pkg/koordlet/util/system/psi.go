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
	"io"
)

const psiLineFormat = "avg10=%f avg60=%f avg300=%f total=%d"

type PSIPath struct {
	CPU string
	Mem string
	IO  string
}

type PSIByResource struct {
	CPU PSIStats
	Mem PSIStats
	IO  PSIStats
}

type PSILine struct {
	Avg10  float64
	Avg60  float64
	Avg300 float64
	Total  uint64
}

type PSIStats struct {
	Some *PSILine
	Full *PSILine

	FullSupported bool
}

// parsePSIStats parses the specified file for pressure stall information.
func ParsePSIStats(r io.Reader) (PSIStats, error) {
	_ = "STUB: not implemented"
	return *new(PSIStats), nil
}

// full cpu pressure not supported in old kernel versions

func GetPSIByResource(paths PSIPath) (*PSIByResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readPSI(pressureFilePath string) (PSIStats, error) {
	_ = "STUB: not implemented"
	return *new(PSIStats), nil
}
