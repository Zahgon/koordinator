/*
Copyright 2021 The Kubernetes Authors.

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

package pod

import (
	"context"
	"net"
	"net/http"
	"regexp"
	"time"

	"k8s.io/apimachinery/pkg/util/httpstream"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// NewTransport creates a transport which uses the port forward dialer.
// URLs must use <namespace>.<pod>:<port> as host.
func NewTransport(client kubernetes.Interface, restConfig *rest.Config) *http.Transport {
	_ = "STUB: not implemented"
	return nil
}

// NewDialer creates a dialer that supports connecting to container ports.
func NewDialer(client kubernetes.Interface, restConfig *rest.Config) *Dialer {
	_ = "STUB: not implemented"
	return nil
}

// Dialer holds the relevant parameters that are independent of a particular connection.
type Dialer struct {
	client     kubernetes.Interface
	restConfig *rest.Config
}

// DialContainerPort connects to a certain container port in a pod.
func (d *Dialer) DialContainerPort(ctx context.Context, addr Addr) (conn net.Conn, finalErr error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// The setup code around the actual portforward is from
// https://github.com/kubernetes/kubernetes/blob/c652ffbe4a29143623a1aaec39f745575f7e43ad/staging/src/k8s.io/kubectl/pkg/cmd/portforward/portforward.go

// create error stream

// We're not writing to this stream, just reading an error message from it.
// This happens asynchronously.

// create data stream

// Addr contains all relevant parameters for a certain port in a pod.
// The container should be running before connections are attempted,
// otherwise the connection will fail.
type Addr struct {
	Namespace, PodName string
	Port               int
}

var _ net.Addr = Addr{}

func (a Addr) Network() string { _ = "STUB: not implemented"; return "" }

func (a Addr) String() string { _ = "STUB: not implemented"; return "" }

// ParseAddr expects a <namespace>.<pod>:<port number> as produced
// by Addr.String.
func ParseAddr(addr string) (*Addr, error) { _ = "STUB: not implemented"; return nil, nil }

var addrRegex = regexp.MustCompile(`^([^\.]+)\.([^:]+):(\d+)$`)

type stream struct {
	addr Addr
	httpstream.Stream
	streamConn httpstream.Connection
}

var _ net.Conn = &stream{}

func (s *stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *stream) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (s *stream) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (s *stream) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *stream) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (s *stream) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

type LocalAddr struct{}

var _ net.Addr = LocalAddr{}

func (l LocalAddr) Network() string { _ = "STUB: not implemented"; return "" }
func (l LocalAddr) String() string  { _ = "STUB: not implemented"; return "" }
