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

package metriccache

import (
	"flag"
	"time"
)

type Config struct {
	MetricGCIntervalSeconds int
	MetricExpireSeconds     int

	TSDBPath              string
	TSDBRetentionDuration time.Duration
	TSDBEnablePromMetrics bool
	TSDBStripeSize        int
	TSDBMaxBytes          int64

	// not necessary now since it is in-memory empty dir now
	TSDBWALSegmentSize            int
	TSDBMaxBlockChunkSegmentSize  int64
	TSDBMinBlockDuration          time.Duration
	TSDBMaxBlockDuration          time.Duration
	TSDBHeadChunksWriteBufferSize int
}

func NewDefaultConfig() *Config { _ = "STUB: not implemented"; return nil }

// 100 MB

// 1 MB
// 5 MB
// 10 minutes
// 10 minutes
// 1 MB

func (c *Config) InitFlags(fs *flag.FlagSet) { _ = "STUB: not implemented"; return }
