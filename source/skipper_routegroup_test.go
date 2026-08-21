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

package source

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"sigs.k8s.io/external-dns/internal/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
	"sigs.k8s.io/external-dns/source/types"
)

func TestRouteGroupClientUpdateToken(t *testing.T) {
	t.Parallel()

	t.Run("no token file path leaves token unchanged", func(t *testing.T) {
		t.Parallel()

		client := &routeGroupClient{token: "initial-token"}
		client.updateToken()

		assert.Equal(t, "initial-token", client.getToken())
	})

	t.Run("token is read from file", func(t *testing.T) {
		t.Parallel()

		tokenFile := filepath.Join(t.TempDir(), "token")
		require.NoError(t, os.WriteFile(tokenFile, []byte("updated-token"), 0o600))
		client := &routeGroupClient{token: "initial-token", tokenFile: tokenFile}

		client.updateToken()

		assert.Equal(t, "updated-token", client.getToken())
	})

	t.Run("unreadable token file leaves token unchanged", func(t *testing.T) {
		t.Parallel()

		client := &routeGroupClient{
			token:     "initial-token",
			tokenFile: filepath.Join(t.TempDir(), "missing-token"),
		}

		client.updateToken()

		assert.Equal(t, "initial-token", client.getToken())
	})
}

func TestRouteGroupClientGetRouteGroupList(t *testing.T) {
	t.Parallel()

	t.Run("successful response", func(t *testing.T) {
		t.Parallel()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(`{"kind":"RouteGroupList","apiVersion":"zalando.org/v1","items":[{"metadata":{"name":"rg1"}}]}`))
			assert.NoError(t, err)
		}))
		defer server.Close()

		client := &routeGroupClient{client: server.Client(), token: "test-token"}
		got, err := client.getRouteGroupList(server.URL)

		require.NoError(t, err)
		require.Len(t, got.Items, 1)
		assert.Equal(t, "RouteGroupList", got.Kind)
		assert.Equal(t, "zalando.org/v1", got.APIVersion)
		assert.Equal(t, "rg1", got.Items[0].Name)
	})

	t.Run("non-success response", func(t *testing.T) {
		t.Parallel()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "unavailable", http.StatusServiceUnavailable)
		}))
		defer server.Close()

		client := &routeGroupClient{client: server.Client()}
		got, err := client.getRouteGroupList(server.URL)

		assert.Nil(t, got)
		assert.ErrorContains(t, err, "503 Service Unavailable")
	})

	t.Run("invalid JSON response", func(t *testing.T) {
		t.Parallel()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, err := w.Write([]byte(`{"items":`))
			assert.NoError(t, err)
		}))
		defer server.Close()

		client := &routeGroupClient{client: server.Client()}
		got, err := client.getRouteGroupList(server.URL)

		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("request failure", func(t *testing.T) {
		t.Parallel()

		server := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
		client := &routeGroupClient{client: server.Client()}
		server.Close()

		got, err := client.getRouteGroupList(server.URL)

		assert.Nil(t, got)
		assert.Error(t, err)
	})
}

func TestRouteGroupClientGetAndDo(t *testing.T) {
	t.Parallel()

	t.Run("invalid URL", func(t *testing.T) {
		t.Parallel()

		client := &routeGroupClient{}
		resp, err := client.get("://invalid")

		assert.Nil(t, resp)
		assert.Error(t, err)
	})

	t.Run("existing authorization header is preserved", func(t *testing.T) {
		t.Parallel()

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "Bearer request-token", r.Header.Get("Authorization"))
			w.WriteHeader(http.StatusNoContent)
		}))
		defer server.Close()

		client := &routeGroupClient{client: server.Client(), token: "client-token"}
		req, err := http.NewRequest(http.MethodGet, server.URL, nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer request-token")

		resp, err := client.do(req)

		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}

func TestNewRouteGroupSource(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name              string
		apiServer         string
		namespace         string
		routeGroupVersion string
		wantServer        string
		wantAPI           string
		wantErr           bool
	}{
		{
			name:       "default version across all namespaces",
			apiServer:  "http://kubernetes.default.svc",
			wantServer: "http://kubernetes.default.svc",
			wantAPI:    "http://kubernetes.default.svc/apis/zalando.org/v1/routegroups",
		},
		{
			name:              "custom version in one namespace",
			apiServer:         "https://kubernetes.default.svc:8443",
			namespace:         "test-namespace",
			routeGroupVersion: "zalando.org/v1alpha1",
			wantServer:        "https://kubernetes.default.svc:8443",
			wantAPI:           "https://kubernetes.default.svc:8443/apis/zalando.org/v1alpha1/namespaces/test-namespace/routegroups",
		},
		{
			name:       "standard HTTPS port is removed",
			apiServer:  "https://kubernetes.default.svc:443",
			wantServer: "https://kubernetes.default.svc",
			wantAPI:    "https://kubernetes.default.svc/apis/zalando.org/v1/routegroups",
		},
		{
			name:       "standard HTTPS port is removed from IPv6 address",
			apiServer:  "https://[2001:db8::1]:443",
			wantServer: "https://[2001:db8::1]",
			wantAPI:    "https://[2001:db8::1]/apis/zalando.org/v1/routegroups",
		},
		{
			name:      "unparsable API server URL returns an error",
			apiServer: "://invalid",
			wantErr:   true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			config := Config{
				Namespace:                tt.namespace,
				SkipperRouteGroupVersion: tt.routeGroupVersion,
				KubeAPIRequestTimeout:    time.Second,
			}
			got, err := NewRouteGroupSource(&config, "test-token", "", tt.apiServer)
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)

			source, ok := got.(*routeGroupSource)
			require.True(t, ok)
			client, ok := source.cli.(*routeGroupClient)
			require.True(t, ok)
			t.Cleanup(func() { close(client.quit) })

			assert.Equal(t, tt.wantServer, source.apiServer)
			assert.Equal(t, tt.namespace, source.namespace)
			assert.Equal(t, tt.wantAPI, source.apiEndpoint)
			assert.Equal(t, "test-token", client.getToken())
		})
	}
}

func TestRouteGroupDeepCopyObject(t *testing.T) {
	t.Parallel()

	original := createTestRouteGroup(
		"namespace1",
		"rg1",
		map[string]string{"key": "value"},
		[]string{"rg1.example.org"},
		nil,
	)
	cloned, ok := original.DeepCopyObject().(*routeGroup)
	require.True(t, ok)

	assert.Equal(t, original, cloned)
	assert.NotSame(t, original, cloned)

	// Top-level fields are independent.
	cloned.Name = "rg2"
	assert.Equal(t, "rg1", original.Name)

	// Nested slices and maps are not.
	cloned.Spec.Hosts[0] = "rg2.example.org"
	cloned.Annotations["key"] = "changed"
	assert.Equal(t, "rg2.example.org", original.Spec.Hosts[0])
	assert.Equal(t, "changed", original.Annotations["key"])
}

func TestRouteGroupClientToken(t *testing.T) {
	t.Parallel()

	t.Run("trims direct token", func(t *testing.T) {
		client := newRouteGroupClient(" direct-token\r\n", "", time.Second)
		t.Cleanup(func() { close(client.quit) })

		assert.Equal(t, "direct-token", client.getToken())
	})

	t.Run("loads custom token file", func(t *testing.T) {
		tokenFile := filepath.Join(t.TempDir(), "token")
		require.NoError(t, os.WriteFile(tokenFile, []byte("custom-token\n"), 0o600))

		client := newRouteGroupClient("initial-token", tokenFile, time.Second)
		t.Cleanup(func() { close(client.quit) })

		assert.Equal(t, tokenFile, client.tokenFile)
		assert.Equal(t, "custom-token", client.getToken())
	})

	t.Run("reloads custom token file", func(t *testing.T) {
		tokenFile := filepath.Join(t.TempDir(), "token")
		require.NoError(t, os.WriteFile(tokenFile, []byte("initial-file-token"), 0o600))

		client := newRouteGroupClient("initial-token", tokenFile, time.Second)
		t.Cleanup(func() { close(client.quit) })

		require.NoError(t, os.WriteFile(tokenFile, []byte("rotated-file-token\r\n"), 0o600))
		client.updateToken()

		assert.Equal(t, "rotated-file-token", client.getToken())
	})
}

func createTestRouteGroup(ns, name string, annotations map[string]string, hosts []string, destinations []routeGroupLoadBalancer) *routeGroup {
	return &routeGroup{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   ns,
			Name:        name,
			Annotations: annotations,
		},
		Spec: routeGroupSpec{
			Hosts: hosts,
		},
		Status: routeGroupStatus{
			LoadBalancer: routeGroupLoadBalancerStatus{
				RouteGroup: destinations,
			},
		},
	}
}

func TestEndpointsFromRouteGroups(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		source *routeGroupSource
		rg     *routeGroup
		want   []*endpoint.Endpoint
	}{
		{
			name:   "Empty routegroup should return empty endpoints",
			source: &routeGroupSource{},
			rg:     &routeGroup{},
			want:   []*endpoint.Endpoint{},
		},
		{
			name:   "Routegroup without hosts and destinations create no endpoints",
			source: &routeGroupSource{},
			rg:     createTestRouteGroup("namespace1", "rg1", nil, nil, nil),
			want:   []*endpoint.Endpoint{},
		},
		{
			name:   "Routegroup without hosts create no endpoints",
			source: &routeGroupSource{},
			rg: createTestRouteGroup("namespace1", "rg1", nil, nil, []routeGroupLoadBalancer{
				{
					Hostname: "lb.example.org",
				},
			}),
			want: []*endpoint.Endpoint{},
		},
		{
			name:   "Routegroup without destinations create no endpoints",
			source: &routeGroupSource{},
			rg:     createTestRouteGroup("namespace1", "rg1", nil, []string{"rg1.k8s.example"}, nil),
			want:   []*endpoint.Endpoint{},
		},
		{
			name:   "Routegroup with hosts and destinations creates an endpoint",
			source: &routeGroupSource{},
			rg: createTestRouteGroup("namespace1", "rg1", nil, []string{"rg1.k8s.example"}, []routeGroupLoadBalancer{
				{
					Hostname: "lb.example.org",
				},
			}),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:   "Routegroup with hostname annotation, creates endpoints from the annotation ",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				map[string]string{
					annotations.HostnameKey: "my.example",
				},
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "my.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:   "Routegroup with hosts and destinations and ignoreHostnameAnnotation creates endpoints but ignores annotation",
			source: &routeGroupSource{ignoreHostnameAnnotation: true},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				map[string]string{
					annotations.HostnameKey: "my.example",
				},
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:   "Routegroup with hosts and destinations and ttl creates an endpoint with ttl",
			source: &routeGroupSource{ignoreHostnameAnnotation: true},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				map[string]string{
					annotations.TtlKey: "2189",
				},
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
					RecordTTL:  endpoint.TTL(2189),
				},
			},
		},
		{
			name:   "Routegroup with hosts and destination IP creates an endpoint",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				nil,
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						IP: "1.5.1.4",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets([]string{"1.5.1.4"}),
				},
			},
		},
		{
			name:   "Routegroup with hosts and destination IPv6 creates an endpoint",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				nil,
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						IP: "2001:DB8::1",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeAAAA,
					Targets:    endpoint.Targets([]string{"2001:DB8::1"}),
				},
			},
		},
		{
			name:   "Routegroup with hosts and mixed destinations creates endpoints",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				nil,
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
						IP:       "1.5.1.4",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets([]string{"1.5.1.4"}),
				},
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:   "Routegroup with hosts and mixed destinations (IPv6) creates endpoints",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				nil,
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
						IP:       "2001:DB8::1",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeAAAA,
					Targets:    endpoint.Targets([]string{"2001:DB8::1"}),
				},
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:   "Routegroup with provider-specific annotation creates endpoint with provider-specific property",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				map[string]string{
					annotations.AWSPrefix + "weight": "10",
				},
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/weight", Value: "10"},
					},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.source.endpointsFromRouteGroup(tt.rg)

			testutils.ValidateEndpoints(t, got, tt.want)
		})
	}
}

type fakeRouteGroupClient struct {
	returnErr bool
	rg        *routeGroupList
}

func (f *fakeRouteGroupClient) getRouteGroupList(string) (*routeGroupList, error) {
	if f.returnErr {
		return nil, errors.New("Fake route group list error")
	}
	return f.rg, nil
}

func TestRouteGroupsEndpoints(t *testing.T) {
	for _, tt := range []struct {
		name        string
		source      *routeGroupSource
		templates   string
		combineFQDN bool
		want        []*endpoint.Endpoint
		wantErr     bool
	}{
		{
			name: "routegroup client error should be returned",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{returnErr: true},
			},
			wantErr: true,
		},
		{
			name: "Empty routegroup should return empty endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{},
				},
			},
			want:    []*endpoint.Endpoint{},
			wantErr: false,
		},
		{
			name: "Single routegroup should return endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							{
								ObjectMeta: metav1.ObjectMeta{
									Namespace: "namespace1",
									Name:      "rg1",
									UID:       "skipper-rg-uid-1234",
								},
								Spec: routeGroupSpec{
									Hosts: []string{"rg1.k8s.example"},
								},
								Status: routeGroupStatus{
									LoadBalancer: routeGroupLoadBalancerStatus{
										RouteGroup: []routeGroupLoadBalancer{
											{
												Hostname: "lb.example.org",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				(&endpoint.Endpoint{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				}).WithRefObject(testutils.RefSource(string(types.SkipperRouteGroup))),
			},
		},
		{
			name:        "Single routegroup with combineFQDNAnnotation with fqdn template should return endpoints from fqdnTemplate and routegroup",
			templates:   "{{.Metadata.Name}}.{{.Metadata.Namespace}}.example",
			combineFQDN: true,
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg1.namespace1.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:      "Single routegroup without, with fqdn template should return endpoints from fqdnTemplate",
			templates: "{{.Metadata.Name}}.{{.Metadata.Namespace}}.example",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								nil,
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.namespace1.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:      "fqdn template execution error should be returned",
			templates: "{{index . 0}}",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								nil,
								[]routeGroupLoadBalancer{{Hostname: "lb.example.org"}},
							),
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name:      "Single routegroup without combineFQDNAnnotation with fqdn template should return endpoints not from fqdnTemplate",
			templates: "{{.Metadata.Name}}.{{.Metadata.Namespace}}.example",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "Single routegroup with TTL should return endpoint with TTL",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								map[string]string{
									annotations.TtlKey: "2189",
								},
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
					RecordTTL:  endpoint.TTL(2189),
				},
			},
		},
		{
			name: "Routegroup with hosts and mixed destinations creates endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
										IP:       "1.5.1.4",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets([]string{"1.5.1.4"}),
				},
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups should return endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace1",
								"rg2",
								nil,
								[]string{"rg2.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace2",
								"rg3",
								nil,
								[]string{"rg3.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace3",
								"rg",
								nil,
								[]string{"rg.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb2.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg2.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg3.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb2.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups with filter annotations should return only filtered endpoints",
			source: &routeGroupSource{
				annotationFilter: parseAnnotationFilterOrNil("kubernetes.io/ingress.class=skipper"),
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								map[string]string{
									"kubernetes.io/ingress.class": "skipper",
								},
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace1",
								"rg2",
								map[string]string{
									"kubernetes.io/ingress.class": "nginx",
								},
								[]string{"rg2.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace2",
								"rg3",
								map[string]string{
									"kubernetes.io/ingress.class": "",
								},
								[]string{"rg3.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace3",
								"rg",
								nil,
								[]string{"rg.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb2.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups with set operation annotation filter should return only filtered endpoints",
			source: &routeGroupSource{
				annotationFilter: parseAnnotationFilterOrNil("kubernetes.io/ingress.class in (nginx, skipper)"),
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								map[string]string{
									"kubernetes.io/ingress.class": "skipper",
								},
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace1",
								"rg2",
								map[string]string{
									"kubernetes.io/ingress.class": "nginx",
								},
								[]string{"rg2.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace2",
								"rg3",
								map[string]string{
									"kubernetes.io/ingress.class": "",
								},
								[]string{"rg3.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace3",
								"rg",
								nil,
								[]string{"rg.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb2.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg2.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups with matching label filter returns only labeled endpoints",
			source: &routeGroupSource{
				labelSelector: labels.SelectorFromSet(labels.Set{"app": "test"}),
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							{
								ObjectMeta: metav1.ObjectMeta{
									Namespace: "namespace1",
									Name:      "rg-match",
									Labels:    map[string]string{"app": "test"},
								},
								Spec: routeGroupSpec{Hosts: []string{"match.example.org"}},
								Status: routeGroupStatus{
									LoadBalancer: routeGroupLoadBalancerStatus{
										RouteGroup: []routeGroupLoadBalancer{{Hostname: "lb.example.org"}},
									},
								},
							},
							{
								ObjectMeta: metav1.ObjectMeta{
									Namespace: "namespace1",
									Name:      "rg-no-match",
									Labels:    map[string]string{"app": "other"},
								},
								Spec: routeGroupSpec{Hosts: []string{"no-match.example.org"}},
								Status: routeGroupStatus{
									LoadBalancer: routeGroupLoadBalancerStatus{
										RouteGroup: []routeGroupLoadBalancer{{Hostname: "lb.example.org"}},
									},
								},
							},
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "match.example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups with non-matching label filter returns no endpoints",
			source: &routeGroupSource{
				labelSelector: labels.SelectorFromSet(labels.Set{"app": "test"}),
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							{
								ObjectMeta: metav1.ObjectMeta{
									Namespace: "namespace1",
									Name:      "rg-no-match",
									Labels:    map[string]string{"app": "other"},
								},
								Spec: routeGroupSpec{Hosts: []string{"no-match.example.org"}},
								Status: routeGroupStatus{
									LoadBalancer: routeGroupLoadBalancerStatus{
										RouteGroup: []routeGroupLoadBalancer{{Hostname: "lb.example.org"}},
									},
								},
							},
						},
					},
				},
			},
			want: []*endpoint.Endpoint{},
		},
		{
			name: "multiple routegroups with controller annotation filter should not return filtered endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								map[string]string{
									annotations.ControllerKey: annotations.ControllerValue,
								},
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace1",
								"rg2",
								map[string]string{
									annotations.ControllerKey: "dns",
								},
								[]string{"rg2.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace2",
								"rg3",
								nil,
								[]string{"rg3.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg3.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if tt.templates != "" {
				if tt.combineFQDN {
					tt.source.templateEngine = templatetest.MustEngine(t, tt.templates, "", "", true)
				} else {
					tt.source.templateEngine = templatetest.MustEngine(t, tt.templates, "", "", false)
				}
			}

			got, err := tt.source.Endpoints(t.Context())
			if err != nil && !tt.wantErr {
				t.Errorf("Got error, but does not want to get an error: %v", err)
			}
			if tt.wantErr && err == nil {
				t.Fatal("Got no error, but we want to get an error")
			}

			testutils.ValidateEndpoints(t, got, tt.want)
		})
	}
}

func TestResourceLabelIsSet(t *testing.T) {
	source := &routeGroupSource{
		cli: &fakeRouteGroupClient{
			rg: &routeGroupList{
				Items: []*routeGroup{
					createTestRouteGroup(
						"namespace1",
						"rg1",
						nil,
						[]string{"rg1.k8s.example"},
						[]routeGroupLoadBalancer{
							{
								Hostname: "lb.example.org",
							},
						},
					),
				},
			},
		},
	}

	got, _ := source.Endpoints(t.Context())
	for _, ep := range got {
		if _, ok := ep.Labels[endpoint.ResourceLabelKey]; !ok {
			t.Errorf("Failed to set resource label on ep %v", ep)
		}
	}
}
