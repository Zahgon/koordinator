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

package fieldindex

import (
	"sync"

	"sigs.k8s.io/controller-runtime/pkg/cache"
)

const (
	IndexPodByNodeName        = "pod.spec.nodeName"
	IndexPodByOwnerRefUID     = "pod.ownerRefUID"
	IndexJobByPodUID          = "job.pod.uid"
	IndexJobPodNamespacedName = "job.pod.namespacedName"
	IndexJobByPodNamespace    = "job.pod.namespace"
)

var (
	registerOnce sync.Once
)

func RegisterFieldIndexes(c cache.Cache) error { _ = "STUB: not implemented"; return nil }
