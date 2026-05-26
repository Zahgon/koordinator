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

package elasticquota

// webhook works with multiple copies at the same time, so it needs to watch other copies' writes to task effect locally.
// OnQuotaAdd/Update/DeleteHandlers no need to check, and if its parent doesn't exist, just create.

func (qt *quotaTopology) OnQuotaAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (qt *quotaTopology) OnQuotaUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

// parentQuotaName change

func (qt *quotaTopology) OnQuotaDelete(obj interface{}) { _ = "STUB: not implemented"; return }
