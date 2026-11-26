/*
Copyright 2026 The Kubernetes Authors.

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

package api

import (
	"context"
	"io"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
)

// The tests below exercise the shutdown of the webhook HTTP API:
//
//  1. startTestServer runs serveHTTPAPI on a random local port.
//  2. A test can hold a GET /records request inside the provider with
//     blockingRecordsProvider, so that the request is still active when the
//     server is stopped.
//  3. testServer.Stop cancels the server context, which starts the shutdown.
//  4. testServer.Wait returns what serveHTTPAPI returned.

// testServer is a webhook HTTP API running on a random local port.
type testServer struct {
	url      string
	listener *trackingListener
	client   *http.Client
	cancel   context.CancelFunc
	result   chan error    // receives the error returned by serveHTTPAPI
	done     chan struct{} // closed when serveHTTPAPI has returned
}

type getResult struct {
	status int
	err    error
}

// trackingListener reports when the server closes it, which is the first step
// of http.Server.Shutdown: from then on, no new connections are accepted.
type trackingListener struct {
	net.Listener
	closed chan struct{}
	once   sync.Once
}

// blockingRecordsProvider blocks in Records until UnblockRecords is called,
// to keep a GET /records request active during the shutdown.
type blockingRecordsProvider struct {
	FakeWebhookProvider
	recordsCalled chan struct{}
	unblock       chan struct{}
	unblockOnce   sync.Once
}

// testWait bounds every wait in these tests so that a bug fails the test
// instead of hanging it.
const testWait = 5 * time.Second

// startTestServer starts serveHTTPAPI with opts and stops it at the end of the
// test. It uses FakeWebhookProvider when opts.Provider is nil.
func startTestServer(t *testing.T, opts ServerOptions) *testServer {
	t.Helper()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	listener := &trackingListener{Listener: l, closed: make(chan struct{})}
	if opts.Provider == nil {
		opts.Provider = FakeWebhookProvider{domainFilter: endpoint.NewDomainFilter(nil)}
	}

	ctx, cancel := context.WithCancel(t.Context())
	transport := &http.Transport{}
	s := &testServer{
		url:      "http://" + l.Addr().String(),
		listener: listener,
		client:   &http.Client{Transport: transport, Timeout: testWait},
		cancel:   cancel,
		result:   make(chan error, 1),
		done:     make(chan struct{}),
	}
	go func() {
		s.result <- serveHTTPAPI(ctx, opts, listener)
		close(s.done)
	}()

	t.Cleanup(func() {
		cancel()
		transport.CloseIdleConnections()
		select {
		case <-s.done:
		case <-time.After(testWait):
			t.Error("webhook API did not stop")
		}
	})
	return s
}

// Stop cancels the server context, which starts the graceful shutdown.
func (s *testServer) Stop() {
	s.cancel()
}

// Wait returns the error returned by serveHTTPAPI.
func (s *testServer) Wait(t *testing.T) error {
	t.Helper()
	select {
	case err := <-s.result:
		return err
	case <-time.After(testWait):
		t.Fatal("webhook API did not return")
		return nil
	}
}

// Get sends a GET request in the background and returns a channel that
// receives the result once the response body has been read.
func (s *testServer) Get(path string) <-chan getResult {
	result := make(chan getResult, 1)
	go func() {
		resp, err := s.client.Get(s.url + path)
		if err != nil {
			result <- getResult{err: err}
			return
		}
		defer resp.Body.Close()
		_, err = io.Copy(io.Discard, resp.Body)
		result <- getResult{status: resp.StatusCode, err: err}
	}()
	return result
}

func waitGet(t *testing.T, result <-chan getResult) getResult {
	t.Helper()
	select {
	case res := <-result:
		return res
	case <-time.After(testWait):
		t.Fatal("HTTP request did not complete")
		return getResult{}
	}
}

func (l *trackingListener) Close() error {
	err := l.Listener.Close()
	l.once.Do(func() { close(l.closed) })
	return err
}

func (l *trackingListener) WaitClosed(t *testing.T) {
	t.Helper()
	select {
	case <-l.closed:
	case <-time.After(testWait):
		t.Fatal("server did not stop accepting connections")
	}
}

// newBlockingRecordsProvider returns a provider whose Records call is released
// at the latest at the end of the test, so that the server can always stop.
func newBlockingRecordsProvider(t *testing.T) *blockingRecordsProvider {
	p := &blockingRecordsProvider{
		recordsCalled: make(chan struct{}),
		unblock:       make(chan struct{}),
	}
	t.Cleanup(p.UnblockRecords)
	return p
}

func (p *blockingRecordsProvider) Records(context.Context) ([]*endpoint.Endpoint, error) {
	close(p.recordsCalled)
	<-p.unblock
	return []*endpoint.Endpoint{}, nil
}

// WaitUntilRecordsCalled waits until a request is blocked in Records.
func (p *blockingRecordsProvider) WaitUntilRecordsCalled(t *testing.T) {
	t.Helper()
	select {
	case <-p.recordsCalled:
	case <-time.After(testWait):
		t.Fatal("request did not reach the provider")
	}
}

// UnblockRecords lets the blocked Records call return.
func (p *blockingRecordsProvider) UnblockRecords() {
	p.unblockOnce.Do(func() { close(p.unblock) })
}

func TestHTTPAPIShutdownWhenIdle(t *testing.T) {
	started := make(chan struct{}, 2)
	s := startTestServer(t, ServerOptions{StartedChan: started})
	select {
	case <-started:
	case <-time.After(testWait):
		t.Fatal("server did not signal that it started")
	}
	res := waitGet(t, s.Get("/"))
	require.NoError(t, res.err)
	require.Equal(t, http.StatusOK, res.status)

	s.Stop()

	require.NoError(t, s.Wait(t))
	assert.Empty(t, started, "the server must signal that it started only once")
}

func TestHTTPAPIShutdownWithoutStartedChanReceiver(t *testing.T) {
	// Nobody receives from StartedChan: the server must not block on it.
	s := startTestServer(t, ServerOptions{StartedChan: make(chan struct{})})

	s.Stop()

	require.NoError(t, s.Wait(t))
}

func TestHTTPAPIShutdownWaitsForActiveRequest(t *testing.T) {
	for name, shutdownTimeout := range map[string]time.Duration{
		"default timeout":  0,
		"negative timeout": -time.Second,
		"custom timeout":   time.Second,
	} {
		t.Run(name, func(t *testing.T) {
			p := newBlockingRecordsProvider(t)
			s := startTestServer(t, ServerOptions{Provider: p, ShutdownTimeout: shutdownTimeout})
			request := s.Get(UrlRecords)
			p.WaitUntilRecordsCalled(t)

			s.Stop()

			// The server stops accepting connections but keeps running while
			// the request is still being served.
			s.listener.WaitClosed(t)
			select {
			case err := <-s.result:
				t.Fatalf("server returned before the active request completed: %v", err)
			default:
			}

			p.UnblockRecords()

			// The request completes, then the server returns without error.
			res := waitGet(t, request)
			require.NoError(t, res.err)
			assert.Equal(t, http.StatusOK, res.status)
			require.NoError(t, s.Wait(t))
		})
	}
}

func TestHTTPAPIShutdownTimeoutClosesActiveRequest(t *testing.T) {
	// Records never returns before the shutdown timeout.
	p := newBlockingRecordsProvider(t)
	s := startTestServer(t, ServerOptions{Provider: p, ShutdownTimeout: 20 * time.Millisecond})
	request := s.Get(UrlRecords)
	p.WaitUntilRecordsCalled(t)

	s.Stop()

	require.ErrorIs(t, s.Wait(t), context.DeadlineExceeded)
	res := waitGet(t, request)
	require.Error(t, res.err, "the connection of the unfinished request must be closed")
}

func TestHTTPAPIServeErrorIsReturned(t *testing.T) {
	s := startTestServer(t, ServerOptions{})

	require.NoError(t, s.listener.Close())

	require.ErrorIs(t, s.Wait(t), net.ErrClosed)
}

func TestStartHTTPAPIWithContextListenError(t *testing.T) {
	inUse, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer inUse.Close()

	err = StartHTTPApiWithContext(t.Context(), ServerOptions{ProviderPort: inUse.Addr().String()})

	require.ErrorContains(t, err, "listen for webhook HTTP API")
}

func TestStartHTTPAPIWithContextAlreadyCanceled(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	err := StartHTTPApiWithContext(ctx, ServerOptions{
		Provider:     FakeWebhookProvider{},
		ProviderPort: "127.0.0.1:0",
	})

	require.NoError(t, err)
}
