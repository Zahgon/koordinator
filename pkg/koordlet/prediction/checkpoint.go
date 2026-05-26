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

package prediction

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/koordinator-sh/koordinator/pkg/util/histogram"
)

const (
	TmpFileSuffix = ".tmp"
)

// ModelCheckpoint represents a checkpoint for a model.
type ModelCheckpoint struct {
	UID         UIDType
	CPU         *histogram.HistogramCheckpoint
	Memory      *histogram.HistogramCheckpoint
	LastUpdated metav1.Time

	Error error `json:"-"`
}

// Checkpointer is an interface for saving and restoring model checkpoints.
type Checkpointer interface {
	Save(checkpoint ModelCheckpoint) error
	Remove(UID UIDType) error
	Restore() ([]*ModelCheckpoint, error)
}

// NewFileCheckpointer creates a new file-based checkpointer with the specified directory.
func NewFileCheckpointer(path string) *fileCheckpointer { _ = "STUB: not implemented"; return nil }

// fileCheckpointer is an implementation of the Checkpointer interface using files.
type fileCheckpointer struct {
	path string // The directory path where checkpoints are stored.
}

// Save saves the given model as a checkpoint with the specified UID.
func (f *fileCheckpointer) Save(checkpoint ModelCheckpoint) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove removes the given model the specified UID.
func (f *fileCheckpointer) Remove(UID UIDType) error { _ = "STUB: not implemented"; return nil }

// Restore returns a slice of ModelCheckpoint instances by scanning and decoding checkpoint files from the specified path.
func (f *fileCheckpointer) Restore() ([]*ModelCheckpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reset UID to the file name
