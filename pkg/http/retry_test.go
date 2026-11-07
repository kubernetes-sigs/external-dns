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

package http

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/wait"
)

var testBackoff = wait.Backoff{Duration: time.Millisecond, Factor: 1.5, Steps: 5}

func TestRequestWithRetry_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	}))
	defer server.Close()

	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := RequestWithRetry(client, req, testBackoff)
	require.NoError(t, err)
	require.NotNil(t, resp)
	DrainAndClose(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRequestWithRetry_NonRetriableStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := RequestWithRetry(client, req, testBackoff)
	require.EqualError(t, err, "unexpected status code 400")
	require.Nil(t, resp)
}

func TestRequestWithRetry_ServerErrorRetried(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	}))
	defer server.Close()

	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := RequestWithRetry(client, req, testBackoff)
	require.NoError(t, err)
	require.NotNil(t, resp)
	DrainAndClose(resp.Body)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 3, attempts)
}

func TestRequestWithRetry_RetriesExhausted(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts++
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, server.URL, nil)
	require.NoError(t, err)

	resp, err := RequestWithRetry(client, req, testBackoff)
	require.EqualError(t, err, "server error: status code 503")
	require.Nil(t, resp)
	require.Equal(t, testBackoff.Steps, attempts)
}

type retryRoundTripper func(*http.Request) (*http.Response, error)

func (f retryRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestRequestWithRetry_CancelledAfterServerError(t *testing.T) {
	ctx, cancel := context.WithCancelCause(t.Context())
	cause := errors.New("shutting down")
	client := &http.Client{Transport: retryRoundTripper(func(r *http.Request) (*http.Response, error) {
		cancel(cause)
		return &http.Response{StatusCode: http.StatusServiceUnavailable, Body: http.NoBody, Request: r}, nil
	})}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://webhook", nil)
	require.NoError(t, err)

	resp, err := RequestWithRetry(client, req, wait.Backoff{Duration: time.Minute, Steps: 2})
	require.ErrorIs(t, err, cause)
	require.Nil(t, resp)
}

func TestRequestWithRetry_NoSteps(t *testing.T) {
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, "http://127.0.0.1", nil)
	require.NoError(t, err)

	resp, err := RequestWithRetry(http.DefaultClient, req, wait.Backoff{})
	require.Error(t, err)
	require.Nil(t, resp)
}
