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

package migration

import (
	"sync"

	"k8s.io/apimachinery/pkg/types"

	sev1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
)

type assumedCache struct {
	lock  sync.Mutex
	items map[types.UID]*sev1alpha1.PodMigrationJob
}

func newAssumedCache() *assumedCache { _ = "STUB: not implemented"; return nil }

func (c *assumedCache) assume(job *sev1alpha1.PodMigrationJob) { _ = "STUB: not implemented"; return }

func (c *assumedCache) delete(job *sev1alpha1.PodMigrationJob) { _ = "STUB: not implemented"; return }

func (c *assumedCache) isNewOrSameObj(job *sev1alpha1.PodMigrationJob) bool {
	_ = "STUB: not implemented"
	return false
}

func getObjVersion(name string, obj interface{}) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
