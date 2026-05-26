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

package defaultevictor

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/descheduler/pkg/framework/plugins/defaultevictor"
	k8sdeschedulerframework "sigs.k8s.io/descheduler/pkg/framework/types"

	"github.com/koordinator-sh/koordinator/pkg/descheduler/evictions"
	"github.com/koordinator-sh/koordinator/pkg/descheduler/framework"
)

const (
	PluginName = defaultevictor.PluginName
)

type DefaultEvictorArgs = defaultevictor.DefaultEvictorArgs

type DefaultEvictor struct {
	handle        framework.Handle
	evictorFilter k8sdeschedulerframework.EvictorPlugin
	evictor       *evictions.PodEvictor
}

var _ framework.EvictPlugin = &DefaultEvictor{}
var _ framework.FilterPlugin = &DefaultEvictor{}
var _ k8sdeschedulerframework.EvictorPlugin = &DefaultEvictor{}

func New(ctx context.Context, args runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	_ = "STUB: not implemented"
	return *new(framework.Plugin), nil
}

func (d *DefaultEvictor) Name() string { _ = "STUB: not implemented"; return "" }

func (d *DefaultEvictor) Filter(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (d *DefaultEvictor) PreEvictionFilter(pod *corev1.Pod) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DefaultEvictor) Evict(ctx context.Context, pod *corev1.Pod, evictOptions framework.EvictOptions) bool {
	_ = "STUB: not implemented"
	return false
}
