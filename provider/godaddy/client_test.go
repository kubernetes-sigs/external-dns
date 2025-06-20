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
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

// Tests that
func TestClient_DoWhenQuotaExceeded(t *testing.T) {
	assert := assert.New(t)

	// Mock server to return 429 with a JSON payload
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	assert.NoError(err, "A CODE_EXCEEDED response should not return an error")
	assert.Equal(http.StatusTooManyRequests, resp.StatusCode, "Expected a 429 response")

	respContents := GDErrorResponse{}
	err = client.UnmarshalResponse(resp, &respContents)
	if assert.Error(err) {
		var apiErr *APIError
		errors.As(err, &apiErr)
		assert.Equal("QUOTA_EXCEEDED", apiErr.Code)
		assert.Equal("rate limit exceeded", apiErr.Message)
	}
}
