//go:build linux
// +build linux

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
	"regexp"
	"sync/atomic"
	"time"

	"github.com/cakturk/go-netstat/netstat"
	"github.com/vishvananda/netlink"
)

var (
	kubeletPortAndPid atomic.Value
	expireTime        = time.Minute
)

type portAndPid struct {
	port           int
	pid            int
	lastUpdateTime time.Time
}

func TCPSocks(fn netstat.AcceptFn) ([]netstat.SockTabEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KubeletPortToPid Query pid by tcp port number with the help of go-netstat
// note: Due to the low efficiency of full traversal, we cache the result and verify each time
func KubeletPortToPid(port int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// CmdLine returns the command line args of a process.
func ProcCmdLine(procRoot string, pid int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PidOf finds process(es) with a specified name (regexp match)
// and return their pid(s).
var PidOf = pidOfFn

// From k8s.io/kubernetes/pkg/util/procfs/procfs_linux.go
// caller should specify proc root dir
func pidOfFn(procRoot string, name string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPids(procRoot string, re *regexp.Regexp) []int { _ = "STUB: not implemented"; return nil }

// Read a small number at a time in case there are many entries, we don't want to
// allocate a lot here.

// If the directory is not a number (i.e. not a PID), skip it

// The bytes we read have '\0' as a separator for the command line

// Split the command line itself we are interested in just the first part

// Check if the name of the executable is what we are looking for

// Grab the PID from the directory path

// If running in container, exec command by 'nsenter --mount=/proc/1/ns/mnt ${cmds}'.
// return stdout, exitcode, error
var ExecCmdOnHost = execCmdOnHostFn

func execCmdOnHostFn(cmds []string) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func GetLinkInfoByDefaultRoute() (netlink.Link, error) {
	_ = "STUB: not implemented"
	return *new(netlink.Link), nil
}
