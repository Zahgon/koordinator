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

package nri

import (
	"sync"
	"time"

	"github.com/containerd/nri/pkg/stub"
)

const (
	events = "RunPodSandbox,RemovePodSandbox,CreateContainer,UpdateContainer"
)

var (
	_ Server = &NriServer{}
)

type NriServer struct {
	cfg      nriConfig
	mask     stub.EventMask
	options  Options       // server options
	stubOpts []stub.Option // nri stub options

	mutex            sync.Mutex
	stopped          bool // if false, the stub will try to reconnect when stub.OnClose is invoked
	stub             StubInterface
	connID           int64
	connClosedSignal chan struct{}
}

func NewNriServer(opt Options) (Server, error) { _ = "STUB: not implemented"; return *new(Server), nil }

func (p *NriServer) Start() (err error) { _ = "STUB: not implemented"; return nil }

// connect at first when start nri server

// then keep connection alive forever until stopped

func (p *NriServer) Stop() { _ = "STUB: not implemented"; return }

func (p *NriServer) keepAlive() {
	_ = "STUB: not implemented"
	// this action is important, we should not catch panic
	// so that we can find the problem
	return
}

// close channel when no consumer

// still not connected ,retry it in next turn

func (p *NriServer) connected() bool { _ = "STUB: not implemented"; return false }

func (p *NriServer) connect() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (p *NriServer) tryDial(connID int64, stubIns stub.Stub, atLeastLiveTime time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// make sure the consumer exists, otherwise the goroutine will be blocked
// and never exit, which will cause memory leak

func (p *NriServer) reconnect(connID int64) func() { _ = "STUB: not implemented"; return nil }

func (p *NriServer) cleanConnection(connID int64) { _ = "STUB: not implemented"; return }

// wait for the plugin to stop
