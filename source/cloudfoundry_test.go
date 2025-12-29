/*
Copyright 2019 The Kubernetes Authors.

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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"sigs.k8s.io/external-dns/endpoint"
)

type RouteSuite struct {
	suite.Suite
}

func TestRouteSource(t *testing.T) {
	suite.Run(t, new(RouteSuite))
	t.Run("Interface", testRouteSourceImplementsSource)
}

// testRouteSourceImplementsSource tests that cloudfoundrySource is a valid Source.
func testRouteSourceImplementsSource(t *testing.T) {
	require.Implements(t, (*Source)(nil), new(cloudfoundrySource))
}

func TestCloudFoundrySourceEndpoints(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v2/private_domains":
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(map[string]any{
				"total_results": 1,
				"total_pages":   1,
				"next_url":      "",
				"resources": []map[string]any{
					{
						"metadata": map[string]any{
							"guid":       "domain-guid",
							"created_at": "now",
							"updated_at": "now",
						},
						"entity": map[string]any{
							"name": "example.com",
						},
					},
				},
			}); err != nil {
				t.Errorf("encode domains response: %v", err)
			}
		case "/v2/routes":
			if got := r.URL.Query().Get("q"); got != "domain_guid:domain-guid" {
				t.Errorf("expected domain query to be %q, got %q", "domain_guid:domain-guid", got)
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(map[string]any{
				"total_results": 2,
				"total_pages":   1,
				"next_url":      "",
				"resources": []map[string]any{
					{
						"metadata": map[string]any{
							"guid":       "route-guid",
							"created_at": "now",
							"updated_at": "now",
						},
						"entity": map[string]any{
							"host":        "app",
							"domain_guid": "domain-guid",
						},
					},
					{
						"metadata": map[string]any{
							"guid":       "route-guid-2",
							"created_at": "now",
							"updated_at": "now",
						},
						"entity": map[string]any{
							"host":        "api",
							"domain_guid": "domain-guid",
						},
					},
				},
			}); err != nil {
				t.Errorf("encode routes response: %v", err)
			}
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)

	client := &cfclient.Client{
		Config: cfclient.Config{
			ApiAddress: server.URL,
			HttpClient: server.Client(),
			UserAgent:  "cloudfoundry-test",
		},
	}

	source, err := NewCloudFoundrySource(client)
	require.NoError(t, err)

	source.AddEventHandler(context.Background(), func() {})
	endpoints, err := source.Endpoints(context.Background())
	require.NoError(t, err)

	expected := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("app.example.com", endpoint.RecordTypeCNAME, 300, "example.com"),
		endpoint.NewEndpointWithTTL("api.example.com", endpoint.RecordTypeCNAME, 300, "example.com"),
	}
	for _, ep := range expected {
		ep.Targets[0] = server.Listener.Addr().String()
	}
	require.ElementsMatch(t, expected, endpoints)
}

func TestCloudFoundrySourceEndpointsPanicsOnBadURL(t *testing.T) {
	client := &cfclient.Client{
		Config: cfclient.Config{
			ApiAddress: "http://[::1]:namedport",
			HttpClient: http.DefaultClient,
			UserAgent:  "cloudfoundry-test",
		},
	}

	source, err := NewCloudFoundrySource(client)
	require.NoError(t, err)

	require.Panics(t, func() {
		_, _ = source.Endpoints(context.Background())
	})
}
