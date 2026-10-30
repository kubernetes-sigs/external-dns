/*
Copyright 2017 The Kubernetes Authors.

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

package godaddy

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
)

// Tests that
func TestClient_DoWhenQuotaExceeded(t *testing.T) {
	assert := assert.New(t)

	// Mock server to return 429 with a JSON payload
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		_, err := w.Write([]byte(`{"code": "QUOTA_EXCEEDED", "message": "rate limit exceeded"}`))
		if err != nil {
			t.Fatalf("Failed to write response: %v", err)
		}
	}))
	defer mockServer.Close()

	client := Client{
		APIKey:      "",
		APISecret:   "",
		APIEndPoint: mockServer.URL,
		Client:      &http.Client{},
		// Add one token every second
		Ratelimiter: rate.NewLimiter(rate.Every(time.Second), 60),
		Timeout:     DefaultTimeout,
	}

	req, err := client.NewRequest("GET", "/v1/domains/example.net/records", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	require.NoError(t, err, "A CODE_EXCEEDED response should not return an error")
	assert.Equal(http.StatusTooManyRequests, resp.StatusCode, "Expected a 429 response")

	respContents := GDErrorResponse{}
	err = client.UnmarshalResponse(resp, &respContents)
	require.Error(t, err)
	var apiErr *APIError
	require.ErrorAs(t, err, &apiErr)
	assert.Equal("QUOTA_EXCEEDED", apiErr.Code)
	assert.Equal("rate limit exceeded", apiErr.Message)
}

func TestClient_DoRetryRateLimitReset(t *testing.T) {
	var attempts atomic.Int32
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if attempts.Add(1) == 1 {
			w.Header().Set("RateLimit-Reset", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			_, err := w.Write([]byte(`{"code": "QUOTA_EXCEEDED", "message": "rate limit exceeded"}`))
			if err != nil {
				t.Fatalf("Failed to write response: %v", err)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`[]`))
		if err != nil {
			t.Fatalf("Failed to write response: %v", err)
		}
	}))
	defer mockServer.Close()

	client := Client{
		APIKey:      "",
		APISecret:   "",
		APIEndPoint: mockServer.URL,
		Client:      &http.Client{},
		Ratelimiter: rate.NewLimiter(rate.Every(time.Second), 60),
		Timeout:     DefaultTimeout,
	}

	req, err := client.NewRequest(http.MethodGet, "/v1/domains/example.net/records", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(2), attempts.Load())
}

func TestClient_DoRetryRetryAfter(t *testing.T) {
	var attempts atomic.Int32
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if attempts.Add(1) == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			_, err := w.Write([]byte(`{"code": "QUOTA_EXCEEDED", "message": "rate limit exceeded"}`))
			if err != nil {
				t.Fatalf("Failed to write response: %v", err)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`[]`))
		if err != nil {
			t.Fatalf("Failed to write response: %v", err)
		}
	}))
	defer mockServer.Close()

	client := Client{
		APIKey:      "",
		APISecret:   "",
		APIEndPoint: mockServer.URL,
		Client:      &http.Client{},
		Ratelimiter: rate.NewLimiter(rate.Every(time.Second), 60),
		Timeout:     DefaultTimeout,
	}

	req, err := client.NewRequest(http.MethodGet, "/v1/domains/example.net/records", nil)
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(2), attempts.Load())
}
