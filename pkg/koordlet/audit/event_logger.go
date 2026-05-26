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

package audit

import (
	"io/fs"
	"sync"
)

const (
	MaxFileSize                   = 1 << 21 // 2MB
	DefaultLogDirMode fs.FileMode = 0755
)

// NewEventLogger create an EventWriter, it won't open the underly file until do a log call.
// verbose=0 means no restrictions on verbose
func NewEventLogger(dir string, sizeMB int, verbose int) EventWriter {
	_ = "STUB: not implemented"
	return *new(EventWriter)
}

// we should ensure the log directory exist and have write to permission

// NewFluentEventLogger create an EventFluentWriter to simplify the audit.
func NewFluentEventLogger(dir string, sizeMB int, verbose int) EventFluentWriter {
	_ = "STUB: not implemented"
	return *new(EventFluentWriter)
}

func NewEventReader(dir string) EventReader { _ = "STUB: not implemented"; return *new(EventReader) }

type eventWriter struct {
	mutex       sync.Mutex
	dir         string
	verbose     int
	maxFileSize int
	maxFileNum  int

	logWriter LogWriter
}

// Log write an event to the underly storage
func (e *eventWriter) Log(verbose int, event *Event) error { _ = "STUB: not implemented"; return nil }

// reset err to rotate status

func (e *eventWriter) openNewLogWriter() error { _ = "STUB: not implemented"; return nil }

// rename audit.log to audit-2021-04-05-12-05-01.log

// if rename failed, return error

// remove the old files

// Flush flush events to the underly storage
func (e *eventWriter) Flush() error { _ = "STUB: not implemented"; return nil }

// Close close the writer
func (e *eventWriter) Close() error { _ = "STUB: not implemented"; return nil }

type eventFluentWriter struct {
	writer EventWriter
}

// V create an eventHelper with Level verbose
func (e *eventFluentWriter) V(verbose int) *EventHelper { _ = "STUB: not implemented"; return nil }

// Flush flush events to the underly storage
func (e *eventFluentWriter) Flush() error { _ = "STUB: not implemented"; return nil }

// Close close the underly writer
func (e *eventFluentWriter) Close() error { _ = "STUB: not implemented"; return nil }

type eventReader struct {
	dir string
}

func (e *eventReader) NewReverseInterator() EventIterator {
	_ = "STUB: not implemented"
	return *new(EventIterator)
}

func newReverseEventIterator(dir string) *reverseEventIterator {
	_ = "STUB: not implemented"
	return nil
}

func listAuditfiles(dir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// reverse files

type reverseEventIterator struct {
	dir        string
	curentFile string
	reader     LogReader
}

func (e *reverseEventIterator) openNextLogReader() error { _ = "STUB: not implemented"; return nil }

func (e *reverseEventIterator) Next() (*Event, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *reverseEventIterator) Close() error { _ = "STUB: not implemented"; return nil }
