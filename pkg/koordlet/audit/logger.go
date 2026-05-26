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
	"bufio"
	"os"
	"time"
)

var (
	ChunkSize = 4096
)

type LogWriter interface {
	Append([]byte) error
	Size() int
	Flush() error
	Close() error
}

type LogReader interface {
	Read() ([]byte, error)
	Offset() int
	Close() error
}

func OpenLogWriter(name string) (LogWriter, error) {
	_ = "STUB: not implemented"
	return *new(LogWriter), nil
}

type logWriter struct {
	size          int
	lastFlushTime time.Time

	file   *os.File
	writer *bufio.Writer
}

func (w *logWriter) Append(log []byte) error { _ = "STUB: not implemented"; return nil }

func (w *logWriter) Size() int { _ = "STUB: not implemented"; return 0 }

func (w *logWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func (w *logWriter) Close() error { _ = "STUB: not implemented"; return nil }

func OpenlogReader(name string) (LogReader, error) {
	_ = "STUB: not implemented"
	return *new(LogReader), nil
}

type logReader struct {
	file       *os.File
	fileOffset int
	offset     int

	current *chunkFile
}

func (r *logReader) readChunks() (*chunkFile, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *logReader) Read() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *logReader) Offset() int { _ = "STUB: not implemented"; return 0 }

func (r *logReader) Close() error { _ = "STUB: not implemented"; return nil }

func readChunk(file *os.File, endOffset int) (*chunkFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// shift to first '\n'

type chunkFile struct {
	startOffset int
	endOffset   int
	lines       [][]byte
	offset      int
}

func (r *chunkFile) Read() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *chunkFile) Size() int { _ = "STUB: not implemented"; return 0 }

func (r *chunkFile) RemainLines() int { _ = "STUB: not implemented"; return 0 }
