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

package options

import (
	"context"

	scheduleroptions "k8s.io/kubernetes/cmd/kube-scheduler/app/options"

	schedulerappconfig "github.com/koordinator-sh/koordinator/cmd/koord-scheduler/app/config"
)

// Options has all the params needed to run a Scheduler
type Options struct {
	*scheduleroptions.Options
	CombinedInsecureServing *CombinedInsecureServingOptions
}

// NewOptions returns default scheduler app options.
func NewOptions() *Options { _ = "STUB: not implemented"; return nil }

func (o *Options) Validate() []error { _ = "STUB: not implemented"; return nil }

// Config return a scheduler config object
func (o *Options) Config(ctx context.Context) (*schedulerappconfig.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE(joseph): When the K8s Scheduler Framework starts, the thread that constructs NodeInfo
// and the scheduling thread are not synchronized. In this way, when the Pod on a Node is not
// filled in the NodeInfo, the Node is scheduled for a new Pod. This behavior is not expected.
// The K8s community itself has also noticed this issue https://github.com/kubernetes/kubernetes/issues/116717,
// but it was only fixed in the K8s v1.28 version https://github.com/kubernetes/kubernetes/pull /116729.
// So we need to fix it ourselves.

// use json for CRD clients
