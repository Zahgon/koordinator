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

package impl

import (
	"github.com/koordinator-sh/koordinator/pkg/koordlet/statesinformer"
)

type updateCallback struct {
	name        string
	description string
	fn          statesinformer.UpdateCbFn
}

type UpdateCbCtx struct{}

type callbackRunner struct {
	callbackChans        map[statesinformer.RegisterType]chan UpdateCbCtx
	stateUpdateCallbacks map[statesinformer.RegisterType][]updateCallback
	statesInformer       statesinformer.StatesInformer
}

func NewCallbackRunner() *callbackRunner { _ = "STUB: not implemented"; return nil }

func (s *callbackRunner) Setup(i statesinformer.StatesInformer) { _ = "STUB: not implemented"; return }

func (s *callbackRunner) RegisterCallbacks(rType statesinformer.RegisterType, name, description string, callbackFn statesinformer.UpdateCbFn) {
	_ = "STUB: not implemented"
	return
}

func (s *callbackRunner) SendCallback(objType statesinformer.RegisterType) {
	_ = "STUB: not implemented"
	return
}

func (s *callbackRunner) runCallbacks(objType statesinformer.RegisterType, obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *callbackRunner) Start(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (s *callbackRunner) getObjByType(objType statesinformer.RegisterType, cbCtx UpdateCbCtx) interface{} {
	_ = "STUB: not implemented"
	return nil
}
