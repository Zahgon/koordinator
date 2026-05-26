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

package proxyserver

import (
	"net"

	"google.golang.org/grpc"
	"k8s.io/client-go/tools/record"

	runtimeapi "github.com/koordinator-sh/koordinator/apis/runtime/v1alpha1"
	"github.com/koordinator-sh/koordinator/pkg/koordlet/resourceexecutor"
	"github.com/koordinator-sh/koordinator/pkg/runtimeproxy/config"
)

type Options struct {
	Network       string
	Address       string
	HostEndpoint  string
	FailurePolicy config.FailurePolicyType
	// support stop running other hooks once someone failed
	PluginFailurePolicy config.FailurePolicyType
	ConfigFilePath      string
	DisableStages       map[string]struct{}
	Executor            resourceexecutor.ResourceUpdateExecutor
	EventRecorder       record.EventRecorder
}

type Server interface {
	Setup() error
	Start() error
	Stop()
	Register() error
}

type server struct {
	listener net.Listener // socket our gRPC server listens on
	server   *grpc.Server // our gRPC server
	options  Options      // server options
	runtimeapi.UnimplementedRuntimeHookServiceServer
}

func (s *server) Setup() error { _ = "STUB: not implemented"; return nil }

func (s *server) Start() error { _ = "STUB: not implemented"; return nil }

func (s *server) Stop() { _ = "STUB: not implemented"; return }

func (s *server) Register() error { _ = "STUB: not implemented"; return nil }

func (s *server) createRPCServer() error { _ = "STUB: not implemented"; return nil }

func NewServer(opt Options) (Server, error) { _ = "STUB: not implemented"; return *new(Server), nil }
