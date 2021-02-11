/*
Copyright 2018 The Kubernetes Authors.

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

package source

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"sigs.k8s.io/external-dns/endpoint"
)

type HTTPClientSuite struct {
	suite.Suite
}

func (suite *HTTPClientSuite) SetupTest() {
}

func startTestHttpServer(t *testing.T, addr string, endpoints []*endpoint.Endpoint) *http.Server {
	t.Helper()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(200)
		enc := json.NewEncoder(w)
		enc.Encode(endpoints)
	})

	srv := &http.Server{Addr: addr, Handler: handler}

	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("HTTP server listening on %s", addr)

	go srv.Serve(ln)
	return srv
}

func TestHttpClientSource(t *testing.T) {
	suite.Run(t, new(HTTPClientSuite))
	t.Run("Interface", testHttpClientSourceImplementsSource)
	t.Run("Endpoints", testHttpClientSourceEndpoints)
}

// testHttpClientSourceImplementsSource tests that httpClientSource is a valid Source.
func testHttpClientSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*Source)(nil), new(httpClientSource))
}

// testHttpClientSourceEndpoints tests that NewHttpClientSource doesn't return an error.
func testHttpClientSourceEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title               string
		serverListenAddress string
		url                 string
		expected            []*endpoint.Endpoint
		expectError         bool
	}{
		{
			title:               "unavailable URL",
			serverListenAddress: "",
			url:                 "http://localhost:8091",
			expectError:         true,
		},
		{
			title:               "valid URL with no endpoints",
			serverListenAddress: "127.0.0.1:8080",
			url:                 "http://127.0.0.1:8080",
			expectError:         false,
		},
		{
			title:               "valid URL",
			serverListenAddress: "127.0.0.1:8081",
			url:                 "http://127.0.0.1:8081",
			expected: []*endpoint.Endpoint{
				{DNSName: "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectError: false,
		},
		{
			title:               "valid URL with multiple endpoints",
			serverListenAddress: "127.0.0.1:8082",
			url:                 "http://127.0.0.1:8082",
			expected: []*endpoint.Endpoint{
				{DNSName: "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
				{DNSName: "xyz.example.org",
					Targets:    endpoint.Targets{"abc.example.org"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  180,
				},
			},
			expectError: false,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			var srv *http.Server
			if ti.serverListenAddress != "" {
				srv = startTestHttpServer(t, ti.serverListenAddress, ti.expected)
			}
			hcs, _ := NewHTTPClientSource(ti.url)

			endpoints, err := hcs.Endpoints(context.Background())
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Validate returned endpoints against expected endpoints.
			validateEndpoints(t, endpoints, ti.expected)

			if srv != nil {
				srv.Shutdown(context.Background())
			}
		})
	}
}
