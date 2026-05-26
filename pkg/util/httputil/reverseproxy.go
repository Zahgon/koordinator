// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// HTTP reverse proxy handler

package httputil

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// ReverseProxy is an HTTP Handler that takes an incoming request and
// sends it to another server, proxying the response back to the
// client.
//
// ReverseProxy by default sets the client IP as the value of the
// X-Forwarded-For header.
//
// If an X-Forwarded-For header already exists, the client IP is
// appended to the existing values. As a special case, if the header
// exists in the Request.Header map but has a nil value (such as when
// set by the Director func), the X-Forwarded-For header is
// not modified.
//
// To prevent IP spoofing, be sure to delete any pre-existing
// X-Forwarded-For header coming from the client or
// an untrusted proxy.
type ReverseProxy struct {
	// Director must be a function which modifies
	// the request into a new request to be sent
	// using Transport. Its response is then copied
	// back to the original client unmodified.
	// Director must not access the provided Request
	// after returning.
	Director func(*http.Request)

	// The transport used to perform proxy requests.
	// If nil, http.DefaultTransport is used.
	Transport http.RoundTripper

	// FlushInterval specifies the flush interval
	// to flush to the client while copying the
	// response body.
	// If zero, no periodic flushing is done.
	// A negative value means to flush immediately
	// after each write to the client.
	// The FlushInterval is ignored when ReverseProxy
	// recognizes a response as a streaming response, or
	// if its ContentLength is -1; for such responses, writes
	// are flushed to the client immediately.
	FlushInterval time.Duration

	// ErrorLog specifies an optional logger for errors
	// that occur when attempting to proxy the request.
	// If nil, logging is done via the log package's standard logger.
	ErrorLog *log.Logger

	// BufferPool optionally specifies a buffer pool to
	// get byte slices for use by io.CopyBuffer when
	// copying HTTP response bodies.
	BufferPool BufferPool

	// ModifyResponse is an optional function that modifies the
	// Response from the backend. It is called if the backend
	// returns a response at all, with any HTTP status code.
	// If the backend is unreachable, the optional ErrorHandler is
	// called without any call to ModifyResponse.
	//
	// If ModifyResponse returns an error, ErrorHandler is called
	// with its error value. If ErrorHandler is nil, its default
	// implementation is used.
	ModifyResponse func(*http.Response) error

	// ErrorHandler is an optional function that handles errors
	// reaching the backend or errors from ModifyResponse.
	//
	// If nil, the default is to log the provided error and return
	// a 502 Status Bad Gateway response.
	ErrorHandler func(http.ResponseWriter, *http.Request, error)
}

// A BufferPool is an interface for getting and returning temporary
// byte slices for use by io.CopyBuffer.
type BufferPool interface {
	Get() []byte
	Put([]byte)
}

func singleJoiningSlash(a, b string) string { _ = "STUB: not implemented"; return "" }

func joinURLPath(a, b *url.URL) (path, rawpath string) { _ = "STUB: not implemented"; return "", "" }

// Same as singleJoiningSlash, but uses EscapedPath to determine
// whether a slash should be added

// NewSingleHostReverseProxy returns a new ReverseProxy that routes
// URLs to the scheme, host, and base path provided in target. If the
// target's path is "/base" and the incoming request was for "/dir",
// the target request will be for /base/dir.
// NewSingleHostReverseProxy does not rewrite the Host header.
// To rewrite Host headers, use ReverseProxy directly with a custom
// Director policy.
func NewSingleHostReverseProxy(target *url.URL) *ReverseProxy {
	_ = "STUB: not implemented"
	return nil
}

// explicitly disable User-Agent so it's not set to default value

func copyHeader(dst, src http.Header) { _ = "STUB: not implemented"; return }

// Hop-by-hop headers. These are removed when sent to the backend.
// As of RFC 7230, hop-by-hop headers are required to appear in the
// Connection header field. These are the headers defined by the
// obsoleted RFC 2616 (section 13.5.1) and are used for backward
// compatibility.
var hopHeaders = []string{
	"Connection",
	"Proxy-Connection", // non-standard but still sent by libcurl and rejected by e.g. google
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",      // canonicalized version of "TE"
	"Trailer", // not Trailers per URL above; https://www.rfc-editor.org/errata_search.php?eid=4522
	"Transfer-Encoding",
	"Upgrade",
}

func (p *ReverseProxy) defaultErrorHandler(rw http.ResponseWriter, req *http.Request, err error) {
	_ = "STUB: not implemented"
	return
}

func (p *ReverseProxy) getErrorHandler() func(http.ResponseWriter, *http.Request, error) {
	_ = "STUB: not implemented"
	return nil
}

// modifyResponse conditionally runs the optional ModifyResponse hook
// and reports whether the request should proceed.
func (p *ReverseProxy) modifyResponse(rw http.ResponseWriter, res *http.Response, req *http.Request) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Issue 16036: nil Body for http.Transport retries

// Reading from the request body after returning from a handler is not
// allowed, and the RoundTrip goroutine that reads the Body can outlive
// this handler. This can lead to a crash if the handler panics (see
// Issue 46866). Although calling Close doesn't guarantee there isn't
// any Read in flight after the handle returns, in practice it's safe to
// read after closing it.

// Issue 33142: historical behavior was to always allocate

// Remove hop-by-hop headers to the backend. Especially
// important is "Connection" because we want a persistent
// connection, regardless of what the client sent to us.

// Issue 21096: tell backend applications that care about trailer support
// that we support trailers. (We do, but we don't go out of our way to
// advertise that unless the incoming client request thought it was worth
// mentioning.) Note that we look at req.Header, not outreq.Header, since
// the latter has passed through removeConnectionHeaders.

// After stripping all the hop-by-hop connection headers above, add back any
// necessary for protocol upgrades, such as for websockets.

// If we aren't the first proxy retain prior
// X-Forwarded-For information as a comma+space
// separated list and fold multiple headers into one.

// Issue 38079: nil now means don't populate the header

// Deal with 101 Switching Protocols responses: (WebSocket, h2c, etc)

// The "Trailer" header isn't included in the Transport's response,
// at least for *http.Transport. Build it up from Trailer.

// Since we're streaming the response, if we run into an error all we can do
// is abort the request. Issue 23643: ReverseProxy should use ErrAbortHandler
// on read error while copying body.

// close now, instead of defer, to populate res.Trailer

// Force chunking if we saw a response trailer.
// This prevents net/http from calculating the length for short
// bodies and adding a Content-Length.

var inOurTests bool // whether we're in our own tests

// shouldPanicOnCopyError reports whether the reverse proxy should
// panic with http.ErrAbortHandler. This is the right thing to do by
// default, but Go 1.10 and earlier did not, so existing unit tests
// weren't expecting panics. Only panic in our own tests, or when
// running under the HTTP server.
func shouldPanicOnCopyError(req *http.Request) bool {
	_ = "STUB: not implemented"

	// Our tests know to handle this panic.
	return false
}

// We seem to be running under an HTTP server, so
// it'll recover the panic.

// Otherwise act like Go 1.10 and earlier to not break
// existing tests.

// removeConnectionHeaders removes hop-by-hop headers listed in the "Connection" header of h.
// See RFC 7230, section 6.1
func removeConnectionHeaders(h http.Header) { _ = "STUB: not implemented"; return }

// flushInterval returns the p.FlushInterval value, conditionally
// overriding its value for a specific request/response.
func (p *ReverseProxy) flushInterval(res *http.Response) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// For Server-Sent Events responses, flush immediately.
// The MIME type is defined in https://www.w3.org/TR/eventsource/#text-event-stream

// negative means immediately

// We might have the case of streaming for which Content-Length might be unset.

func (p *ReverseProxy) copyResponse(dst io.Writer, src io.Reader, flushInterval time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// set up initial timer so headers get flushed even if body writes are delayed

// copyBuffer returns any write errors or non-EOF read errors, and the amount
// of bytes written.
func (p *ReverseProxy) copyBuffer(dst io.Writer, src io.Reader, buf []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *ReverseProxy) logf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

type writeFlusher interface {
	io.Writer
	http.Flusher
}

type maxLatencyWriter struct {
	dst     writeFlusher
	latency time.Duration // non-zero; negative means to flush immediately

	mu           sync.Mutex // protects t, flushPending, and dst.Flush
	t            *time.Timer
	flushPending bool
}

func (m *maxLatencyWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *maxLatencyWriter) delayedFlush() { _ = "STUB: not implemented"; return }

// if stop was called but AfterFunc already started this goroutine

func (m *maxLatencyWriter) stop() { _ = "STUB: not implemented"; return }

func upgradeType(h http.Header) string { _ = "STUB: not implemented"; return "" }

func (p *ReverseProxy) handleUpgradeResponse(rw http.ResponseWriter, req *http.Request, res *http.Response) {
	_ = "STUB: not implemented"
	return
}

// Ensure that the cancelation of a request closes the backend.
// See issue https://golang.org/issue/35559.

// so res.Write only writes the headers; we have res.Body in backConn above

// CloseWriter implement CloseWrite method
type CloseWriter interface {
	CloseWrite() error
}

// switchProtocolCopier exists so goroutines proxying data back and
// forth have nice names in stacks.
type switchProtocolCopier struct {
	user, backend io.ReadWriter
}

func (c switchProtocolCopier) copyFromBackend(errc chan<- error) { _ = "STUB: not implemented"; return }

func (c switchProtocolCopier) copyToBackend() error { _ = "STUB: not implemented"; return nil }
