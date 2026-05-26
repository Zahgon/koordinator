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

package health

import (
	"net/http"
	"path"
	"sync"

	"github.com/fsnotify/fsnotify"

	webhookutil "github.com/koordinator-sh/koordinator/pkg/webhook/util"
)

var (
	caCertFilePath = path.Join(webhookutil.GetCertDir(), "ca-cert.pem")

	onceWatch sync.Once
	lock      sync.Mutex
	client    *http.Client
)

func loadHTTPClientWithCACert() error { _ = "STUB: not implemented"; return nil }

func watchCACert(watcher *fsnotify.Watcher) { _ = "STUB: not implemented"; return }

// Channel is closed.

// Only care about events which may modify the contents of the file.

// If the file was removed, re-add the watch.

// Channel is closed.

func isWrite(event fsnotify.Event) bool { _ = "STUB: not implemented"; return false }

func isCreate(event fsnotify.Event) bool { _ = "STUB: not implemented"; return false }

func isRemove(event fsnotify.Event) bool { _ = "STUB: not implemented"; return false }

func Checker(_ *http.Request) error { _ = "STUB: not implemented"; return nil }
