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

type shutdownListener struct {
	net.Listener
	closed chan struct{}
	once   sync.Once
}

func (l *shutdownListener) Close() error {
	err := l.Listener.Close()
	l.once.Do(func() { close(l.closed) })
	return err
}

type webhookAPIFixture struct {
	listener *shutdownListener
	client   *http.Client
	cancel   context.CancelFunc
	result   <-chan error
}

func startWebhookAPI(t *testing.T, opts ServerOptions) *webhookAPIFixture {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	l := &shutdownListener{Listener: listener, closed: make(chan struct{})}
	if opts.Provider == nil {
		opts.Provider = FakeWebhookProvider{domainFilter: endpoint.NewDomainFilter(nil)}
	}
	ctx, cancel := context.WithCancel(t.Context())
	result := make(chan error, 1)
	done := make(chan struct{})
	go func() {
		result <- serveHTTPAPI(ctx, opts, l)
		close(done)
	}()
	transport := &http.Transport{}
	f := &webhookAPIFixture{
		listener: l, client: &http.Client{Transport: transport, Timeout: 5 * time.Second},
		cancel: cancel, result: result,
	}
	t.Cleanup(func() {
		cancel()
		transport.CloseIdleConnections()
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			_ = l.Close()
			t.Error("webhook API did not stop")
		}
	})
	return f
}

func waitWebhookResult(t *testing.T, result <-chan error) error {
	t.Helper()
	select {
	case err := <-result:
		return err
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for the webhook API")
		return nil
	}
}

func TestHTTPAPIShutdownIdle(t *testing.T) {
	started := make(chan struct{}, 2)
	f := startWebhookAPI(t, ServerOptions{StartedChan: started})
	select {
	case <-started:
	case <-time.After(5 * time.Second):
		t.Fatal("server did not report readiness")
	}
	resp, err := f.client.Get("http://" + f.listener.Addr().String())
	require.NoError(t, err)
	_, err = io.Copy(io.Discard, resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())
	require.Equal(t, http.StatusOK, resp.StatusCode)

	f.cancel()
	require.NoError(t, waitWebhookResult(t, f.result))
	assert.Empty(t, started, "readiness must be signaled only once")
}

func TestHTTPAPIShutdownWithUnreadReadiness(t *testing.T) {
	f := startWebhookAPI(t, ServerOptions{StartedChan: make(chan struct{})})
	f.cancel()
	require.NoError(t, waitWebhookResult(t, f.result))
}

type drainingWebhookProvider struct {
	FakeWebhookProvider
	entered chan struct{}
	release chan struct{}
}

func (p *drainingWebhookProvider) Records(context.Context) ([]*endpoint.Endpoint, error) {
	close(p.entered)
	<-p.release
	return []*endpoint.Endpoint{}, nil
}

type webhookRequestResult struct {
	status int
	err    error
}

func startWebhookRecordsRequest(t *testing.T, f *webhookAPIFixture, p *drainingWebhookProvider) <-chan webhookRequestResult {
	t.Helper()
	result := make(chan webhookRequestResult, 1)
	go func() {
		resp, err := f.client.Get("http://" + f.listener.Addr().String() + UrlRecords)
		if err != nil {
			result <- webhookRequestResult{err: err}
			return
		}
		_, err = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
		result <- webhookRequestResult{status: resp.StatusCode, err: err}
	}()
	select {
	case <-p.entered:
	case <-time.After(5 * time.Second):
		t.Fatal("request did not reach the provider")
	}
	return result
}

func waitWebhookRequest(t *testing.T, result <-chan webhookRequestResult) webhookRequestResult {
	t.Helper()
	select {
	case res := <-result:
		return res
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for the HTTP request")
		return webhookRequestResult{}
	}
}

func TestHTTPAPIShutdownDrainsActiveRequest(t *testing.T) {
	for _, timeout := range []time.Duration{0, -time.Second, time.Second} {
		t.Run(timeout.String(), func(t *testing.T) {
			p := &drainingWebhookProvider{entered: make(chan struct{}), release: make(chan struct{})}
			f := startWebhookAPI(t, ServerOptions{Provider: p, ShutdownTimeout: timeout})
			release := sync.OnceFunc(func() { close(p.release) })
			t.Cleanup(release)
			request := startWebhookRecordsRequest(t, f, p)

			f.cancel()
			select {
			case <-f.listener.closed:
			case <-time.After(5 * time.Second):
				t.Fatal("shutdown did not close the listener")
			}
			select {
			case err := <-f.result:
				t.Fatalf("server returned before the active request completed: %v", err)
			default:
			}

			release()
			res := waitWebhookRequest(t, request)
			require.NoError(t, res.err)
			assert.Equal(t, http.StatusOK, res.status)
			require.NoError(t, waitWebhookResult(t, f.result))
		})
	}
}

func TestHTTPAPIShutdownTimeoutClosesConnection(t *testing.T) {
	p := &drainingWebhookProvider{entered: make(chan struct{}), release: make(chan struct{})}
	f := startWebhookAPI(t, ServerOptions{Provider: p, ShutdownTimeout: 20 * time.Millisecond})
	t.Cleanup(sync.OnceFunc(func() { close(p.release) }))
	request := startWebhookRecordsRequest(t, f, p)
	f.cancel()
	require.ErrorIs(t, waitWebhookResult(t, f.result), context.DeadlineExceeded)
	res := waitWebhookRequest(t, request)
	require.Error(t, res.err, "the timed-out connection must be closed")
}

func TestHTTPAPIServeErrorIsReturned(t *testing.T) {
	f := startWebhookAPI(t, ServerOptions{})
	require.NoError(t, f.listener.Close())
	require.ErrorIs(t, waitWebhookResult(t, f.result), net.ErrClosed)
}

func TestStartHTTPAPIWithContextListenError(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer listener.Close()
	err = StartHTTPApiWithContext(t.Context(), ServerOptions{ProviderPort: listener.Addr().String()})
	require.ErrorContains(t, err, "listen for webhook HTTP API")
}

func TestStartHTTPAPIWithCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	started := make(chan struct{}, 1)
	err := StartHTTPApiWithContext(ctx, ServerOptions{ProviderPort: "not a valid address", StartedChan: started})
	require.NoError(t, err)
	assert.Empty(t, started)
}
