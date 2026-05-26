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
	"testing"
)

const (
	msgResourceSupportedForTesting = "resource is set supported for testing"
)

var (
	defaultAnolisOSResourcesForTesting = []Resource{
		CPUBurst,
		CPUBVTWarpNs,
		MemoryWmarkRatio,
		MemoryWmarkScaleFactor,
		MemoryWmarkMinAdj,
		MemoryMin,
		MemoryLow,
		MemoryHigh,
		MemoryPriority,
		MemoryUsePriorityOom,
		MemoryOomGroup,
		BlkioIOQoS,
		BlkioIOWeight,
		BlkioReadBps,
		BlkioReadIops,
		BlkioWriteBps,
		BlkioWriteIops,
		NetClsClassId,
	}
)

type FileTestUtil struct {
	// Temporary directory to store mock cgroup filesystem.
	TempDir string
	// whether to validate when writing cgroups resources
	ValidateResource bool
	// additional cleanup function for Config to be invoked in Cleanup()
	CleanupFn func(config *Config)

	t testing.TB
}

type MockMonData struct {
	CacheItems map[int]MockCacheItem
}

type MockCacheItem map[string]uint64

// create mock ctrl group mon_data directory and files
func TestingPrepareResctrlMondata(t *testing.T, sysFsRootPath, ctrlGrp string, mmd MockMonData) {
	_ = "STUB: not implemented"
	return
}

// NewFileTestUtil creates a new test util for the specified subsystem.
// NOTE: this function should be called only for testing purposes.
func NewFileTestUtil(t testing.TB) *FileTestUtil {
	_ = "STUB: not implemented"
	// NOTE: When $TMPDIR is not set, `t.TempDir()` can use different base directory on Mac OS X and Linux, which may
	// generates too long paths to test unix socket.
	return nil
}

func (c *FileTestUtil) Cleanup() { _ = "STUB: not implemented"; return }

func (c *FileTestUtil) SetResourcesSupported(supported bool, resources ...Resource) {
	_ = "STUB: not implemented"
	return
}

func (c *FileTestUtil) SetAnolisOSResourcesSupported(supported bool) {
	_ = "STUB: not implemented"
	return
}

func (c *FileTestUtil) SetCgroupsV2(useCgroupsV2 bool) { _ = "STUB: not implemented"; return }

func (c *FileTestUtil) SetValidateResource(enabled bool) { _ = "STUB: not implemented"; return }

func (c *FileTestUtil) SetConf(setFn, cleanupFn func(conf *Config)) {
	_ = "STUB: not implemented"
	return
}

// if dir contain TempDir, mkdir direct, else join with TempDir and mkdir
func (c *FileTestUtil) MkDirAll(testDir string) { _ = "STUB: not implemented"; return }

// if filePath contain TempDir, createFile direct, else join with TempDir and create
func (c *FileTestUtil) CreateFile(testFilePath string) { _ = "STUB: not implemented"; return }

// if filePath contain TempDir, write direct, else join with TempDir and write
func (c *FileTestUtil) WriteFileContents(testFilePath, contents string) {
	_ = "STUB: not implemented"
	return
}

// if filePath contain TempDir, read direct, else join with TempDir and read
func (c *FileTestUtil) ReadFileContents(testFilePath string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *FileTestUtil) CreateProcSubFile(fileRelativePath string) {
	_ = "STUB: not implemented"
	return
}

func (c *FileTestUtil) WriteProcSubFileContents(relativeFilePath string, contents string) {
	_ = "STUB: not implemented"
	return
}

func (c *FileTestUtil) ReadProcSubFileContents(relativeFilePath string) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *FileTestUtil) CreateCgroupFile(taskDir string, r Resource) {
	_ = "STUB: not implemented"
	return
}

// WriteCgroupFileContents is only intended for test functions. For specific read/write functionalities, please refer
// to the executor package.
func (c *FileTestUtil) WriteCgroupFileContents(taskDir string, r Resource, contents string) {
	_ = "STUB: not implemented"
	return
}

func (c *FileTestUtil) ReadCgroupFileContentsInt(taskDir string, r Resource) *int64 {
	_ = "STUB: not implemented"
	return nil
}

func (c *FileTestUtil) ReadCgroupFileContents(taskDir string, r Resource) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *FileTestUtil) stripPrefix(path string) string { _ = "STUB: not implemented"; return "" }
