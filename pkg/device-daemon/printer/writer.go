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

package mamanger

import (
	"io"

	resourceconifg "github.com/koordinator-sh/koordinator/cmd/koord-device-daemon/config/v1"
	koordletuti "github.com/koordinator-sh/koordinator/pkg/koordlet/util"
)

// Writer defines a mechanism to output labels and device infos.
type Writer interface {
	OutputPrints(koordletuti.XPUDevices) error
}

// toFile writes to the specified file.
type toFile string

func (path *toFile) OutputPrints(devices koordletuti.XPUDevices) error {
	_ = "STUB: not implemented"
	return nil
}

type toWriter struct {
	io.Writer
}

func (output *toWriter) OutputPrints(devices koordletuti.XPUDevices) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPrintsWriter(config *resourceconifg.Config) (Writer, error) {
	_ = "STUB: not implemented"
	return *new(Writer), nil
}
