/*
Copyright 2025 The Kubernetes Authors.

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

package wrappers

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestWithProviderLabel(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectProvider string
		isConfigured   bool
	}{
		{
			name:           "valid provider",
			input:          "aws",
			expectProvider: "aws",
			isConfigured:   true,
		},
		{
			name:         "empty string",
			input:        "",
			isConfigured: false,
		},
		{
			name:         "whitespace only",
			input:        "   ",
			isConfigured: false,
		},
		{
			name:           "provider with surrounding whitespace",
			input:          "  aws  ",
			expectProvider: "aws",
			isConfigured:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &PostProcessorConfig{}
			opt := WithProviderLabel(tt.input)
			opt(cfg)

			require.Equal(t, tt.isConfigured, cfg.isConfigured, "isConfigured mismatch")
			require.Equal(t, tt.expectProvider, cfg.provider, "provider mismatch")
		})
	}
}

func TestWithTTL(t *testing.T) {
	tests := []struct {
		name         string
		ttlStr       string
		expectErr    bool
		expectTTL    int64
		isConfigured bool
	}{
		{
			name:         "valid 10m6s",
			ttlStr:       "10m6s",
			expectErr:    false,
			expectTTL:    606,
			isConfigured: true,
		},
		{
			name:         "valid 5m",
			ttlStr:       "5m",
			expectTTL:    300,
			isConfigured: true,
		},
		{
			name:      "zero duration",
			ttlStr:    "0s",
			expectTTL: 0,
		},
		{
			name:      "empty duration",
			ttlStr:    "0s",
			expectTTL: 0,
		},
		{
			name:      "invalid duration",
			ttlStr:    "notaduration",
			expectErr: true,
			expectTTL: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &PostProcessorConfig{}
			ttl, err := time.ParseDuration(tt.ttlStr)
			if tt.expectErr {
				require.Error(t, err, "should fail to parse duration string")
				return
			}
			require.NoError(t, err, "should parse duration string")

			opt := WithTTL(ttl)
			opt(cfg)

			require.Equal(t, tt.isConfigured, cfg.isConfigured, "isConfigured mismatch")
			require.Equal(t, tt.expectTTL, cfg.ttl, "ttl mismatch")
		})
	}
}

func TestPostProcessorEndpointsWithTTL(t *testing.T) {
	tests := []struct {
		title     string
		ttl       string
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
		expectErr bool
	}{
		{
			title: "process endpoints with TTL set",
			ttl:   "6s",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo-1", "A", "1.2.3.4"),
				endpoint.NewEndpointWithTTL("foo-2", "A", 60, "1.2.3.5"),
				endpoint.NewEndpointWithTTL("foo-3", "A", 0, "1.2.3.6"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("foo-1", "A", 6, "1.2.3.4"),
				endpoint.NewEndpointWithTTL("foo-2", "A", 60, "1.2.3.5"),
				endpoint.NewEndpointWithTTL("foo-3", "A", 6, "1.2.3.6"),
			},
		},
		{
			title: "skip endpoints processing with TTL set to 0",
			ttl:   "0s",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo-1", "A", "1.2.3.4"),
				endpoint.NewEndpointWithTTL("foo-2", "A", 60, "1.2.3.5"),
				endpoint.NewEndpointWithTTL("foo-3", "A", 0, "1.2.3.6"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo-1", "A", "1.2.3.4"),
				endpoint.NewEndpointWithTTL("foo-2", "A", 60, "1.2.3.5"),
				endpoint.NewEndpointWithTTL("foo-3", "A", 0, "1.2.3.6"),
			},
		},
		{
			title: "skip endpoints processing for nill endpoint",
			ttl:   "0s",
			endpoints: []*endpoint.Endpoint{
				nil,
				endpoint.NewEndpointWithTTL("foo-2", "A", 60, "1.2.3.5"),
			},
			expected: []*endpoint.Endpoint{
				nil,
				endpoint.NewEndpointWithTTL("foo-2", "A", 60, "1.2.3.5"),
			},
		},
		{
			title: "endpoint foo-2 with TTL configured while foo-1 without TTL configured",
			ttl:   "1s",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "foo-1", Targets: endpoint.Targets{"1.2.3.5"}},
				{DNSName: "foo-2", Targets: endpoint.Targets{"1.2.3.6"}, RecordTTL: endpoint.TTL(0)},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo-1", Targets: endpoint.Targets{"1.2.3.5"}, RecordTTL: endpoint.TTL(1)},
				{DNSName: "foo-2", Targets: endpoint.Targets{"1.2.3.6"}, RecordTTL: endpoint.TTL(1)},
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.title, func(t *testing.T) {

			ms := new(testutils.MockSource)
			ms.On("Endpoints").Return(tt.endpoints, nil)
			ttl, _ := time.ParseDuration(tt.ttl)
			src := NewPostProcessor(ms, WithTTL(ttl))

			endpoints, err := src.Endpoints(context.Background())
			require.NoError(t, err)
			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestPostProcessorEndpointsWithProviderFilter(t *testing.T) {
	tests := []struct {
		title     string
		provider  string
		endpoints []*endpoint.Endpoint
		expected  []*endpoint.Endpoint
	}{
		{
			title:    "no provider configured, properties untouched",
			provider: "",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/evaluate-target-health", Value: "true"},
						{Name: "coredns/group", Value: "my-group"},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/evaluate-target-health", Value: "true"},
						{Name: "coredns/group", Value: "my-group"},
					},
				},
			},
		},
		{
			title:    "provider configured, all properties match",
			provider: "aws",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/evaluate-target-health", Value: "true"},
						{Name: "aws/weight", Value: "10"},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/evaluate-target-health", Value: "true"},
						{Name: "aws/weight", Value: "10"},
					},
				},
			},
		},
		{
			title:    "provider configured, mixed properties, only provider retained",
			provider: "aws",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/evaluate-target-health", Value: "true"},
						{Name: "coredns/group", Value: "my-group"},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/evaluate-target-health", Value: "true"},
					},
				},
			},
		},
		{
			title:    "provider configured, no matching properties, empty result",
			provider: "aws",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "coredns/group", Value: "my-group"},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
				},
			},
		},
		{
			title:    "provider agnostic properties without prefix are retained",
			provider: "aws",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "alias", Value: "true"},
						{Name: "aws/evaluate-target-health", Value: "true"},
						{Name: "coredns/group", Value: "my-group"},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "alias", Value: "true"},
						{Name: "aws/evaluate-target-health", Value: "true"},
					},
				},
			},
		},
		{
			title:    "cloudflare retains all properties regardless of prefix",
			provider: "cloudflare",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "external-dns.alpha.kubernetes.io/cloudflare-tags", Value: "tag1"},
						{Name: "aws/evaluate-target-health", Value: "true"},
						{Name: "alias", Value: "false"},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "alias", Value: "false"},
						{Name: "aws/evaluate-target-health", Value: "true"},
						{Name: "external-dns.alpha.kubernetes.io/cloudflare-tags", Value: "tag1"},
					},
				},
			},
		},
		{
			title:    "cloudflare properties are sorted",
			provider: "cloudflare",
			endpoints: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "external-dns.alpha.kubernetes.io/cloudflare-tags", Value: "tag1"},
						{Name: "external-dns.alpha.kubernetes.io/cloudflare-proxied", Value: "true"},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo-1",
					Targets: endpoint.Targets{"1.2.3.4"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "external-dns.alpha.kubernetes.io/cloudflare-proxied", Value: "true"},
						{Name: "external-dns.alpha.kubernetes.io/cloudflare-tags", Value: "tag1"},
					},
				},
			},
		},
		{
			title:    "nil endpoint is skipped",
			provider: "aws",
			endpoints: []*endpoint.Endpoint{
				nil,
				{
					DNSName: "foo-2",
					Targets: endpoint.Targets{"1.2.3.5"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/weight", Value: "10"},
						{Name: "coredns/group", Value: "my-group"},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				nil,
				{
					DNSName: "foo-2",
					Targets: endpoint.Targets{"1.2.3.5"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/weight", Value: "10"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			ms := new(testutils.MockSource)
			ms.On("Endpoints").Return(tt.endpoints, nil)
			src := NewPostProcessor(ms, WithProviderLabel(tt.provider))

			endpoints, err := src.Endpoints(context.Background())
			require.NoError(t, err)
			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}

func TestPostProcessor_AddEventHandler(t *testing.T) {
	tests := []struct {
		title string
		input []string
		times int
	}{
		{
			title: "should add event handler",
			times: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			mockSource := testutils.NewMockSource()

			src := NewPostProcessor(mockSource)
			src.AddEventHandler(t.Context(), func() {})

			mockSource.AssertNumberOfCalls(t, "AddEventHandler", tt.times)
		})
	}
}

func TestPostProcessorEndpointsWithPreferAlias(t *testing.T) {
	tests := []struct {
		title       string
		preferAlias bool
		endpoints   []*endpoint.Endpoint
		expected    []*endpoint.Endpoint
	}{
		{
			title:       "CNAME records get alias annotation when preferAlias is enabled",
			preferAlias: true,
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeCNAME, "target.example.com"),
				endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeCNAME, "target.example.com").WithProviderSpecific("alias", "true"),
				endpoint.NewEndpoint("bar.example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			title:       "CNAME records remain unchanged when preferAlias is disabled",
			preferAlias: false,
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeCNAME, "target.example.com"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("foo.example.com", endpoint.RecordTypeCNAME, "target.example.com"),
			},
		},
		{
			title:       "only CNAME records are affected, A records are unchanged",
			preferAlias: true,
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("aaaa.example.com", endpoint.RecordTypeAAAA, "::1"),
				endpoint.NewEndpoint("cname.example.com", endpoint.RecordTypeCNAME, "target.example.com"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("aaaa.example.com", endpoint.RecordTypeAAAA, "::1"),
				endpoint.NewEndpoint("cname.example.com", endpoint.RecordTypeCNAME, "target.example.com").WithProviderSpecific("alias", "true"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			ms := new(testutils.MockSource)
			ms.On("Endpoints").Return(tt.endpoints, nil)
			src := NewPostProcessor(ms, WithPostProcessorPreferAlias(tt.preferAlias))

			endpoints, err := src.Endpoints(context.Background())
			require.NoError(t, err)
			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}
