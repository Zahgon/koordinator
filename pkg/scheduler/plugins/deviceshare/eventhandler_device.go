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

package deviceshare

import (
	koordinatorinformers "github.com/koordinator-sh/koordinator/pkg/client/informers/externalversions"
)

func registerDeviceEventHandler(deviceCache *nodeDeviceCache, koordSharedInformerFactory koordinatorinformers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

// make sure Device resources are loaded before Pods

func (n *nodeDeviceCache) onDeviceAdd(obj interface{}) { _ = "STUB: not implemented"; return }

func (n *nodeDeviceCache) onDeviceUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (n *nodeDeviceCache) onDeviceDelete(obj interface{}) { _ = "STUB: not implemented"; return }

//
// The user may accidentally delete the Device CRD object,
// and then the Device CRD object will be recreated by koordlet/DevicePlugin.
// During this period, the internal state can only be marked as invalid,
// otherwise the GPU may be repeatedly allocated to different Pods.
//
