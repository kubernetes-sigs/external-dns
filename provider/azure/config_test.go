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

package azure

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/assert"
)

func TestGetCloudConfiguration(t *testing.T) {
	tests := map[string]struct {
		cloudName string
		expected  cloud.Configuration
	}{
		"AzureChinaCloud":   {"AzureChinaCloud", cloud.AzureChina},
		"AzurePublicCloud":  {"", cloud.AzurePublic},
		"AzureUSGovernment": {"AzureUSGovernmentCloud", cloud.AzureGovernment},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cloudCfg, err := getCloudConfiguration(test.cloudName)
			if err != nil {
				t.Errorf("got unexpected err %v", err)
			}
			if cloudCfg.ActiveDirectoryAuthorityHost != test.expected.ActiveDirectoryAuthorityHost {
				t.Errorf("got %v, want %v", cloudCfg, test.expected)
			}
		})
	}
}

func TestOverrideConfiguration(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	configFile := path.Join(path.Dir(filename), "fixtures/config_test.json")
	cfg, err := getConfig(configFile, "subscription-override", "rg-override", "", "aad-endpoint-override")
	if err != nil {
		t.Errorf("got unexpected err %v", err)
	}
	assert.Equal(t, cfg.SubscriptionID, "subscription-override")
	assert.Equal(t, cfg.ResourceGroup, "rg-override")
	assert.Equal(t, cfg.ActiveDirectoryAuthorityHost, "aad-endpoint-override")
}

func TestGetMaxRetries(t *testing.T) {
	defaultRetries := 3
	tests := []struct {
		name         string
		envValue     string
		expected     int
		shouldSetEnv bool
	}{
		{"UnsetEnvVar", "", defaultRetries, false},
		{"ValidPositive", "5", 5, true},
		{"ZeroRetries", "0", 0, true},
		{"NegativeRetries", "-2", -2, true},
		{"InvalidString", "abc", defaultRetries, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test: %s", tt.name)
			if tt.shouldSetEnv {
				os.Setenv("AZURE_SDK_MAX_RETRIES", tt.envValue)
				t.Logf("Set AZURE_SDK_MAX_RETRIES=%s", tt.envValue)
				t.Cleanup(func() { os.Unsetenv("AZURE_SDK_MAX_RETRIES") }) // Clean up after test
			} else {
				os.Unsetenv("AZURE_SDK_MAX_RETRIES")
			}
			got := GetMaxRetries()
			t.Logf("GetMaxRetries() returned: %d (expected: %d)", got, tt.expected)
			if got != tt.expected {
				t.Errorf("GetMaxRetries() = %d; want %d", got, tt.expected)
			}
		})
	}
}

// Test for custom header policy
type transportFunc func(*http.Request) (*http.Response, error)

func (f transportFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestCustomHeaderPolicyWithRetries(t *testing.T) {
	// Set up test environment
	os.Setenv("AZURE_SDK_MAX_RETRIES", "6")
	defer os.Unsetenv("AZURE_SDK_MAX_RETRIES")

	maxRetries := int32(GetMaxRetries())
	var attempt int32
	var firstRequestID string

	// Create mock transport that simulates 429 responses
	mockTransport := transportFunc(func(req *http.Request) (*http.Response, error) {
		attempt++

		// Get the request ID from header
		requestID := req.Header.Get("x-ms-client-request-id")
		if requestID == "" {
			t.Fatalf("Request ID missing on attempt %d", attempt)
		}

		// On first attempt, store the request ID
		if attempt == 1 {
			firstRequestID = requestID
			t.Logf("Initial request ID: %s", firstRequestID)
		} else {
			// On subsequent attempts, verify it matches the first request ID
			if requestID != firstRequestID {
				t.Fatalf("Request ID changed on retry %d: got %s, want %s",
					attempt, requestID, firstRequestID)
			} else {
				t.Logf("Request ID preserved on attempt %d: %s", attempt, requestID)
			}
		}

		// Verify the ID is also in the context
		if ctxID, ok := req.Context().Value(clientRequestIDKey).(string); !ok || ctxID != requestID {
			t.Errorf("Context ID mismatch on attempt %d: got %v, want %s",
				attempt, ctxID, requestID)
		}

		// Return 429 for all but the last attempt
		if attempt <= maxRetries {
			t.Logf("Attempt %d: THROTTLED (429) - Request ID: %s", attempt, requestID)
			return &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Body:       io.NopCloser(strings.NewReader("Too many requests")),
				Request:    req,
				Header: http.Header{
					"x-ms-client-request-id": []string{requestID},
					"Retry-After":            []string{"1"},
				},
			}, nil
		}

		// Return 200 on final attempt
		t.Logf("Attempt %d: SUCCESS (200) - Request ID: %s", attempt, requestID)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("Success")),
			Request:    req,
			Header: http.Header{
				"x-ms-client-request-id": []string{requestID},
			},
		}, nil
	})

	// Create pipeline with retry policy and custom header policy
	mockPipeline := azruntime.NewPipeline(
		"testmodule",
		"1.0",
		azruntime.PipelineOptions{
			PerCall: []policy.Policy{
				CustomHeaderPolicynew(),
			},
		},
		&policy.ClientOptions{
			Retry: policy.RetryOptions{
				MaxRetries: maxRetries,
			},
			Transport: mockTransport,
		},
	)
	// Create request and execute
	req, err := azruntime.NewRequest(context.Background(), http.MethodGet, "https://example.com")
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := mockPipeline.Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// Verify we got the expected number of attempts
	expectedAttempts := maxRetries + 1
	if attempt != expectedAttempts {
		t.Errorf("Wrong number of attempts: got %d, want %d", attempt, expectedAttempts)
	}
	t.Logf("Test completed with %d attempts, all with request ID: %s", attempt, firstRequestID)
	t.Logf("Test summary:")
	t.Logf("- Total attempts: %d", attempt)
	t.Logf("- Throttled responses: %d", maxRetries)
	t.Logf("- Final status: Success")
	t.Logf("- Consistent Request ID: %s", firstRequestID)
}
