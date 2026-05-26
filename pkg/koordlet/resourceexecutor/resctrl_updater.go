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

package resourceexecutor

import (
	"github.com/koordinator-sh/koordinator/pkg/koordlet/audit"
	sysutil "github.com/koordinator-sh/koordinator/pkg/koordlet/util/system"
)

var _ ResourceUpdater = &ResctrlSchemataResourceUpdater{}

type ResctrlSchemataResourceUpdater struct {
	DefaultResourceUpdater
	schemataRaw *sysutil.ResctrlSchemataRaw
}

func (r *ResctrlSchemataResourceUpdater) Name() string { _ = "STUB: not implemented"; return "" }

func (r *ResctrlSchemataResourceUpdater) Key() string { _ = "STUB: not implemented"; return "" }

func (r *ResctrlSchemataResourceUpdater) update() error { _ = "STUB: not implemented"; return nil }

func (r *ResctrlSchemataResourceUpdater) Clone() ResourceUpdater {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater)
}

func NewResctrlSchemataResource(group, schemata string, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

// The current assumption is that the cache ids obtained through
// resctrl schemata will not go wrong. TODO: Use the ability of node info
// to obtain cache ids to replace the current method.

func NewCatGroupResource(group string, e *audit.EventHelper) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater), nil
}

func NewResctrlL3SchemataResource(group, schemataDelta string, l3Num int) ResourceUpdater {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater)
}

// The current assumption is that the cache ids obtained through
// resctrl schemata will not go wrong. TODO: Use the ability of node info
// to obtain cache ids to replace the current method.

func NewResctrlMbSchemataResource(group, schemataDelta string, l3Num int) ResourceUpdater {
	_ = "STUB: not implemented"
	return *new(ResourceUpdater)
}

// The current assumption is that the cache ids obtained through
// resctrl schemata will not go wrong. TODO: Use the ability of node info
// to obtain cache ids to replace the current method.

func CalculateResctrlL3TasksResource(group string, taskIds []int32) (ResourceUpdater, error) {
	_ = "STUB: not implemented"
	// join ids into updater value and make the id updates one by one
	return *new(ResourceUpdater), nil
}

// use ordered slice

func InitCatGroupFunc(u ResourceUpdater) error { _ = "STUB: not implemented"; return nil }

func UpdateResctrlSchemataFunc(u ResourceUpdater) error { _ = "STUB: not implemented"; return nil }

// schemata unchanged, no need to update

// NOTE: currently, only l3 and mba schemata are to update, so do not read or compare before the write
// eg.
// $ cat /sys/fs/resctrl/schemata/BE/schemata
// L3:0=7ff;1=7ff
// MB:0=100;1=100
// $ echo "L3:0=3f;1=3f" > /sys/fs/resctrl/BE/schemata
// $ cat /sys/fs/resctrl/BE/schemata
// L3:0=03f;1=03f
// MB:0=100;1=100

func UpdateResctrlTasksFunc(resource ResourceUpdater) error {
	_ = "STUB: not implemented"
	// NOTE: resctrl/{...}/tasks file is required to appending write a task id once a time, and any duplicate would be
	//
	//	dropped automatically without an exception
	//
	// eg.
	// $ echo 123 > /sys/fs/resctrl/BE/tasks
	// $ echo 124 > /sys/fs/resctrl/BE/tasks
	// $ echo 123 > /sys/fs/resctrl/BE/tasks
	// $ echo 122 > /sys/fs/resctrl/BE/tasks
	// $ tail -n 3 /sys/fs/resctrl/BE/tasks
	// 122
	// 123
	// 124
	return nil
}

// any thread can exit before the writing
