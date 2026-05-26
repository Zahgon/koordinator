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

package asynclog

import (
	"bytes"
	"io"
	"sync"

	"github.com/spf13/pflag"
)

var (
	enableAsync = pflag.Bool("async-log", false, "Enable asynchronous logging to improve performance. When enabling async-log, should enable logtostderr and disable alsologtostderr at the same time. By default, klog outputs logs synchronously to stderr, which will affect performance when there are too many logs.")
	queueLength = pflag.Int("async-log-queue-length", 10000, "Control the log queue length to tune performance.")

	globalOutput *output
	once         sync.Once

	dataPool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		},
	}
)

func EnableAsyncIfNeed() bool { _ = "STUB: not implemented"; return false }

func FlushAndExit() { _ = "STUB: not implemented"; return }

type output struct {
	w            io.Writer
	logCh        chan *bytes.Buffer
	quit         chan bool
	shuttingDown int32
}

func newOutput(queueLength int) *output { _ = "STUB: not implemented"; return nil }

func (o *output) logger() { _ = "STUB: not implemented"; return }

func (o *output) FlushAndExit() { _ = "STUB: not implemented"; return }

func (o *output) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// The data must be copied to a temporary buffer because the data may be reused
