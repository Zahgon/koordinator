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

const (
	ProcStatName    = "stat"
	ProcMemInfoName = "meminfo"
	ProcCPUInfoName = "cpuinfo"
)

func GetProcFilePath(procRelativePath string) string { _ = "STUB: not implemented"; return "" }

func GetProcRootDir() string { _ = "STUB: not implemented"; return "" }

// ProcStat is the content of /proc/<pid>/stat.
// https://manpages.ubuntu.com/manpages/xenial/en/man5/proc.5.html
type ProcStat struct {
	Pid   uint32
	Comm  string
	State byte
	Ppid  uint32
	Pgrp  uint32
	// TODO: add more fields if needed
}

func GetProcPIDStatPath(pid uint32) string { _ = "STUB: not implemented"; return "" }

func ParseProcPIDStat(content string) (*ProcStat, error) {
	_ = "STUB: not implemented"
	// pattern: `12345 (stress) S 12340 12344 12340 12300 12345 123450 151 0 0 0 0 0 ...`
	// splitAfterComm -> "12345 (stress", " S 12340 12344 12340 12300 12345 123450 151 0 0 0 0 0 ..."
	return nil, nil
}

// comm
// splitBeforeComm -> "12345 ", "stress"

// pid

// fieldsAfterComm -> "S", "12340", "12344", "12340", "12300", "12345", "123450", "151", "0", ...

// remaining fields are ignored

// state

// ppid

// pgrp/pgid

func GetPGIDForPID(pid uint32) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// GetPGIDsForPIDs gets the PGIDs for a cgroup's PIDs.
// It will consider the PID as PGID if its PGID does not exist anymore.
func GetPGIDsForPIDs(pids []uint32) ([]uint32, error) { _ = "STUB: not implemented"; return nil, nil }

// get PGID (pgrp) via /proc/$pid/stat

// verify if PGID lives in the pid list
// if not, consider the PID as PGID

// in ascending order

func GetContainerPGIDs(containerParentDir string) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
