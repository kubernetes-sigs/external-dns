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
	"fmt"
	"io"
	"net/http"
	"path"
	"runtime"
	"strconv"
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
	assert.Equal(t, "subscription-override", cfg.SubscriptionID)
	assert.Equal(t, "rg-override", cfg.ResourceGroup)
	assert.Equal(t, "aad-endpoint-override", cfg.ActiveDirectoryAuthorityHost)
}

// Test for custom header policy
type transportFunc func(*http.Request) (*http.Response, error)

func (f transportFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestCustomHeaderPolicyWithRetries(t *testing.T) {
	// Set up test environment
	defaultRetries := 3
	flagValue := "-6"
	isSet := true

	retries, err := parseMaxRetries(flagValue, defaultRetries)
	if err != nil {
		t.Fatalf("Failed to parse retries: %v", err)
	}
	maxRetries := int32(retries)
	if !isSet || (isSet && flagValue == "0") {
		// Use default if flag not provided OR if flag is "0"
		maxRetries = int32(defaultRetries)
		t.Logf("Using default value: %d (flag provided: %v, value: %q)",
			defaultRetries, isSet, flagValue)
	} else {
		// Flag was provided with non-zero value
		retries, err := parseMaxRetries(flagValue, defaultRetries)
		if err != nil {
			t.Fatalf("Failed to parse retries: %v", err)
		}
		maxRetries = int32(retries)
		t.Logf("Using provided flag value: %d", retries)
	}

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
		if maxRetries < 0 || attempt <= maxRetries {
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
	var expectedAttempts int32
	if maxRetries < 0 {
		expectedAttempts = 1 // For negative retries, only one attempt should be made
	} else {
		expectedAttempts = maxRetries + 1 // For zero or positive retries, attempts = retries + 1
	}

	if attempt != expectedAttempts {
		t.Errorf("Wrong number of attempts: got %d, want %d", attempt, expectedAttempts)
	}

	t.Logf("Test completed with %d attempts, all with request ID: %s", attempt, firstRequestID)
}

func TestMaxRetriesCount(t *testing.T) {
	defaultRetries := 3

	tests := []struct {
		name        string
		input       string
		isSet       bool // indicates if flag was provided
		expected    int
		shouldError bool
		description string
	}{
		{
			name:        "FlagNotProvided",
			input:       "",
			isSet:       false,
			expected:    defaultRetries,
			shouldError: false,
			description: "When flag is not provided, should use default value",
		},
		{
			name:        "FlagProvidedEmpty",
			input:       "",
			isSet:       true,
			expected:    0,
			shouldError: true,
			description: "When flag is provided but empty, should error",
		},
		{
			name:        "ValidPositive",
			input:       "5",
			isSet:       true,
			expected:    5,
			shouldError: false,
			description: "Valid positive number should be accepted",
		},
		{
			name:        "ZeroRetries",
			input:       "0",
			isSet:       true,
			expected:    0,
			shouldError: false,
			description: "Zero should be accepted and handled by SDK",
		},
		{
			name:        "NegativeRetries",
			input:       "-2",
			isSet:       true,
			expected:    -2,
			shouldError: false,
			description: "Negative values should be accepted and  handled by SDK",
		},
		{
			name:        "InvalidString",
			input:       "abc",
			isSet:       true,
			expected:    0,
			shouldError: true,
			description: "Non-numeric string should error",
		},
		{
			name:        "Whitespace",
			input:       "   ",
			isSet:       true,
			expected:    0,
			shouldError: true,
			description: "Whitespace should error",
		},
		{
			name:        "SpecialChars",
			input:       "@#$%",
			isSet:       true,
			expected:    0,
			shouldError: true,
			description: "Special characters should error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("=== Test Case: %s ===", tt.name)
			t.Logf("Description: %s", tt.description)
			t.Logf("Input: %q (flag provided: %v)", tt.input, tt.isSet)

			// Handle flag not provided case
			if !tt.isSet {
				t.Logf("Using default value: %d", defaultRetries)
				return
			}

			retries, err := parseMaxRetries(tt.input, defaultRetries)

			// Check error condition
			if tt.shouldError {
				if err == nil {
					t.Errorf("Expected error for input %q but got none", tt.input)
				} else {
					t.Logf("Got expected error: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if retries != tt.expected {
					t.Errorf("Got %d retries, want %d", retries, tt.expected)
				} else {
					t.Logf("Got expected value: %d", retries)
				}
			}
		})
	}
}

// Helper function to parse max retries value
func parseMaxRetries(value string, defaultValue int) (int, error) {
	// Trim whitespace
	value = strings.TrimSpace(value)

	// Empty string or whitespace should error
	if value == "" {
		return 0, fmt.Errorf("retry count must be provided when flag is set")
	}

	retries, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid retry count %q: %w", value, err)
	}

	return retries, nil
}
