/*
Copyright 2023 The Kubernetes Authors.

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

package adguardhome

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
)

func TestAdGuardHomeClient_ListRecords(t *testing.T) {
	// Mock server that returns a list of records
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prepare the mock response
		response := []adGuardHomeEntry{
			{Domain: "domain.filtered.com", Answer: "192.168.0.1"},
			{Domain: "a.example.com", Answer: "172.0.0.1"},
			{Domain: "aaaa.example.com", Answer: "2001:db8::1"},
			{Domain: "cname.example.com", Answer: "target.example.net"},
		}

		// Write the mock response as JSON
		err := json.NewEncoder(w).Encode(response)
		assert.NoError(t, err)
	}))
	defer server.Close()

	// Create a new AdGuardHomeClient with the mock server URL and required config fields
	cfg := AdGuardHomeConfig{
		Username:     "testuser",
		Password:     "testpassword",
		Server:       server.URL,
		DomainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	client, err := newAdGuardHomeClient(cfg)
	assert.NoError(t, err)

	// Call the listRecords method
	endpoints, err := client.listRecords(context.Background())

	// Verify the returned endpoints and error
	assert.NoError(t, err)
	assert.Equal(t, 3, len(endpoints))

	// Verify the contents of the returned endpoints
	expectedEndpoints := []*endpoint.Endpoint{
		{
			DNSName:    "a.example.com",
			Targets:    []string{"172.0.0.1"},
			RecordType: endpoint.RecordTypeA,
		},
		{
			DNSName:    "aaaa.example.com",
			Targets:    []string{"2001:db8::1"},
			RecordType: endpoint.RecordTypeAAAA,
		},
		{
			DNSName:    "cname.example.com",
			Targets:    []string{"target.example.net"},
			RecordType: endpoint.RecordTypeCNAME,
		},
	}
	assert.ElementsMatch(t, expectedEndpoints, endpoints)
}

func TestAdGuardHomeClient_CreateRecord(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/control/rewrite/add", r.URL.Path)
		// Check the request method
		assert.Equal(t, http.MethodPost, r.Method)
		// Respond with success
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a new AdGuardHomeClient with the mock server URL
	client, err := newAdGuardHomeClient(AdGuardHomeConfig{
		Server:   server.URL,
		Username: "test-username",
		Password: "test-password",
	})
	assert.NoError(t, err)

	// Test unsupported record type
	err = client.createRecord(context.Background(), &endpoint.Endpoint{
		DNSName:    "example.com",
		Targets:    []string{"192.168.0.1"},
		RecordType: "MX",
	})
	assert.Error(t, err)
	assert.Equal(t, "unsupported record type: MX for example.com", err.Error())

	// Test unsupported record type
	err = client.createRecord(context.Background(), &endpoint.Endpoint{
		DNSName:    "example.com",
		Targets:    []string{"192.168.0.1"},
		RecordType: "TXT",
	})
	assert.Error(t, err)
	assert.Equal(t, "unsupported record type: TXT for example.com", err.Error())

	// Test supported record types
	err = client.createRecord(context.Background(), &endpoint.Endpoint{
		DNSName:    "example.com",
		Targets:    []string{"192.168.0.1"},
		RecordType: "A",
	})
	assert.NoError(t, err)

	err = client.createRecord(context.Background(), &endpoint.Endpoint{
		DNSName:    "example.com",
		Targets:    []string{"2001:db8::1"},
		RecordType: "AAAA",
	})
	assert.NoError(t, err)

	err = client.createRecord(context.Background(), &endpoint.Endpoint{
		DNSName:    "example.com",
		Targets:    []string{"target.example.net"},
		RecordType: "CNAME",
	})
	assert.NoError(t, err)
}

func TestAdGuardHomeClient_DeleteRecord(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/control/rewrite/delete", r.URL.Path)
		// Check the request method
		assert.Equal(t, http.MethodPost, r.Method)
		// Respond with success
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a new AdGuardHomeClient with the mock server URL
	client, err := newAdGuardHomeClient(AdGuardHomeConfig{
		Server:   server.URL,
		Username: "test-username",
		Password: "test-password",
	})
	assert.NoError(t, err)

	// Create a sample endpoint
	endpoint := &endpoint.Endpoint{
		DNSName:    "example.com",
		Targets:    []string{"192.168.0.1"},
		RecordType: endpoint.RecordTypeA,
	}

	// Call the deleteRecord method
	err = client.deleteRecord(context.Background(), endpoint)

	// Verify the result
	assert.NoError(t, err)
}

func TestAdGuardHomeClient_UpdateRecord(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL
		assert.Equal(t, "/control/rewrite/update", r.URL.Path)

		// Read and parse the request body
		var request adGuardHomeUpdateEntry
		err := json.NewDecoder(r.Body).Decode(&request)
		assert.NoError(t, err)

		// Verify the content of the request
		assert.Equal(t, "old.example.com", request.Target.Domain)
		assert.Equal(t, "192.168.0.1", request.Target.Answer)
		assert.Equal(t, "new.example.com", request.Update.Domain)
		assert.Equal(t, "10.0.0.1", request.Update.Answer)

		// Respond with a successful result
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Create a new AdGuardHomeClient with the mock server URL and required config fields
	client, err := newAdGuardHomeClient(AdGuardHomeConfig{
		Server:   server.URL,
		Username: "your-username",
		Password: "your-password",
	})
	assert.NoError(t, err)

	// Define the old and new endpoints for the update
	oldEndpoint := &endpoint.Endpoint{
		DNSName:    "old.example.com",
		Targets:    []string{"192.168.0.1"},
		RecordType: endpoint.RecordTypeA,
	}
	newEndpoint := &endpoint.Endpoint{
		DNSName:    "new.example.com",
		Targets:    []string{"10.0.0.1"},
		RecordType: endpoint.RecordTypeA,
	}

	// Call the updateRecord method
	err = client.updateRecord(context.Background(), oldEndpoint, newEndpoint)

	// Verify the result
	assert.NoError(t, err)
}

func TestAdGuardHomeClient_do_Non200StatusCode(t *testing.T) {
	// Create a mock server that returns a non-200 status code
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound) // Return a non-200 status code
	}))
	defer server.Close()

	// Parse the mock server URL into a *url.URL object
	parsedURL, err := url.Parse(server.URL)
	assert.NoError(t, err)

	// Create a new AdGuardHomeClient with the parsed URL and required config fields
	client := &adGuardHomeClient{
		BaseURL:    parsedURL,
		cfg:        AdGuardHomeConfig{},
		httpClient: http.DefaultClient,
	}

	// Create a mock HTTP request
	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	assert.NoError(t, err)

	// Call the do method
	_, err = client.do(req)

	// Verify that the error is returned for non-200 status code
	assert.Error(t, err)
}

func TestNewAdGuardHomeClient_InvalidConfig(t *testing.T) {
	// Test case: Empty username
	cfg := AdGuardHomeConfig{
		Server:   "http://example.com",
		Username: "",
		Password: "password",
	}
	client, err := newAdGuardHomeClient(cfg)
	assert.Nil(t, client)
	assert.EqualError(t, err, "no username supplied, this is required")

	// Test case: Empty password
	cfg = AdGuardHomeConfig{
		Server:   "http://example.com",
		Username: "username",
		Password: "",
	}
	client, err = newAdGuardHomeClient(cfg)
	assert.Nil(t, client)
	assert.EqualError(t, err, "no password supplied, this is required")

	// Test case: Invalid server URL
	cfg = AdGuardHomeConfig{
		Server:   ":invalid-url",
		Username: "username",
		Password: "password",
	}
	client, err = newAdGuardHomeClient(cfg)
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.IsType(t, &url.Error{}, err) // Ensure the error type is *url.Error
}
