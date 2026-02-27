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

package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

func TestReverseAddr(t *testing.T) {
	tests := []struct {
		name    string
		ip      string
		want    string
		wantErr bool
	}{
		{
			name: "IPv4 simple",
			ip:   "1.2.3.4",
			want: "4.3.2.1.in-addr.arpa.",
		},
		{
			name: "IPv4 8.8.8.8",
			ip:   "8.8.8.8",
			want: "8.8.8.8.in-addr.arpa.",
		},
		{
			name: "IPv4 192.168.1.1",
			ip:   "192.168.1.1",
			want: "1.1.168.192.in-addr.arpa.",
		},
		{
			name: "IPv6 loopback",
			ip:   "::1",
			want: "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa.",
		},
		{
			name: "IPv6 2001:db8::1",
			ip:   "2001:db8::1",
			want: "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa.",
		},
		{
			name:    "invalid IP",
			ip:      "not-an-ip",
			wantErr: true,
		},
		{
			name:    "empty string",
			ip:      "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReverseAddr(tt.ip)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGeneratePTREndpoints(t *testing.T) {
	tests := []struct {
		name      string
		endpoints []*endpoint.Endpoint
		wantCount int
		wantPTRs  map[string][]string // reverse-addr -> hostnames (without trailing dot, per endpoint normalization)
	}{
		{
			name: "A record generates PTR",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
			},
			wantCount: 1,
			wantPTRs: map[string][]string{
				"8.8.8.8.in-addr.arpa": {"example.com"},
			},
		},
		{
			name: "AAAA record generates PTR",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeAAAA, endpoint.TTL(600), "2001:db8::1"),
			},
			wantCount: 1,
			wantPTRs: map[string][]string{
				"1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa": {"example.com"},
			},
		},
		{
			name: "CNAME record does NOT generate PTR",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("cname.example.com", endpoint.RecordTypeCNAME, endpoint.TTL(300), "example.com"),
			},
			wantCount: 0,
		},
		{
			name: "TXT record does NOT generate PTR",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "some-text"),
			},
			wantCount: 0,
		},
		{
			name: "mixed records only generate PTR for A/AAAA",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "1.2.3.4"),
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeTXT, endpoint.TTL(300), "heritage=external-dns"),
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeCNAME, endpoint.TTL(300), "other.com"),
			},
			wantCount: 1,
			wantPTRs: map[string][]string{
				"4.3.2.1.in-addr.arpa": {"example.com"},
			},
		},
		{
			name: "A record with multiple different IPs generates multiple PTRs",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "1.2.3.4", "5.6.7.8"),
			},
			wantCount: 2,
			wantPTRs: map[string][]string{
				"4.3.2.1.in-addr.arpa": {"example.com"},
				"8.7.6.5.in-addr.arpa": {"example.com"},
			},
		},
		{
			name: "multiple A records with same IP are merged into one PTR",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("web.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "192.168.49.2"),
				endpoint.NewEndpointWithTTL("app.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "192.168.49.2"),
			},
			wantCount: 1,
			wantPTRs: map[string][]string{
				"2.49.168.192.in-addr.arpa": {"web.example.com", "app.example.com"},
			},
		},
		{
			name: "TTL is preserved",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(3600), "10.0.0.1"),
			},
			wantCount: 1,
		},
		{
			name:      "empty input",
			endpoints: []*endpoint.Endpoint{},
			wantCount: 0,
		},
		{
			name:      "nil input",
			endpoints: nil,
			wantCount: 0,
		},
		{
			name: "wildcard A record is skipped",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("*.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "1.2.3.4"),
			},
			wantCount: 0,
		},
		{
			name: "wildcard AAAA record is skipped",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("*.example.com", endpoint.RecordTypeAAAA, endpoint.TTL(300), "2001:db8::1"),
			},
			wantCount: 0,
		},
		{
			name: "wildcard A mixed with normal A — only normal generates PTR",
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("*.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "1.2.3.4"),
				endpoint.NewEndpointWithTTL("www.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "5.6.7.8"),
			},
			wantCount: 1,
			wantPTRs: map[string][]string{
				"8.7.6.5.in-addr.arpa": {"www.example.com"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GeneratePTREndpoints(tt.endpoints, nil)
			assert.Len(t, result, tt.wantCount)

			for _, ep := range result {
				assert.Equal(t, "PTR", ep.RecordType)
			}

			if tt.wantPTRs != nil {
				for _, ep := range result {
					expectedTargets, ok := tt.wantPTRs[ep.DNSName]
					assert.True(t, ok, "unexpected PTR endpoint: %s", ep.DNSName)
					assert.ElementsMatch(t, expectedTargets, ep.Targets)
				}
			}

			// Verify TTL is preserved
			if tt.name == "TTL is preserved" && len(result) > 0 {
				assert.Equal(t, endpoint.TTL(3600), result[0].RecordTTL)
			}
		})
	}
}

func TestGeneratePTREndpoints_DomainFilter(t *testing.T) {
	df := endpoint.NewDomainFilter([]string{"example.local"})

	eps := []*endpoint.Endpoint{
		// Matches domain filter — should produce PTR
		endpoint.NewEndpoint("web.example.local", "A", "192.168.49.2"),
		// Does NOT match domain filter — should be skipped
		endpoint.NewEndpoint("web.example.com", "A", "192.168.49.3"),
	}

	result := GeneratePTREndpoints(eps, df)
	assert.Len(t, result, 1)
	assert.Equal(t, "2.49.168.192.in-addr.arpa", result[0].DNSName)
	assert.Equal(t, "PTR", result[0].RecordType)
	assert.Contains(t, result[0].Targets, "web.example.local")
}

func TestGeneratePTREndpoints_NilDomainFilter(t *testing.T) {
	eps := []*endpoint.Endpoint{
		endpoint.NewEndpoint("web.example.local", "A", "192.168.49.2"),
		endpoint.NewEndpoint("web.example.com", "A", "192.168.49.3"),
	}

	// nil filter means no filtering — both produce PTRs
	result := GeneratePTREndpoints(eps, nil)
	assert.Len(t, result, 2)
}

// mockProvider is a test helper that records ApplyChanges calls
type mockProvider struct {
	BaseProvider
	appliedChanges []*plan.Changes
	recordsResult  []*endpoint.Endpoint
}

func (m *mockProvider) Records(_ context.Context) ([]*endpoint.Endpoint, error) {
	return m.recordsResult, nil
}

func (m *mockProvider) ApplyChanges(_ context.Context, changes *plan.Changes) error {
	m.appliedChanges = append(m.appliedChanges, changes)
	return nil
}

func TestPTRProviderAdjustEndpoints(t *testing.T) {
	t.Run("A record generates PTR in adjusted endpoints", func(t *testing.T) {
		mock := &mockProvider{}
		p := NewPTRProvider(mock)

		endpoints := []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
		}

		result, err := p.AdjustEndpoints(endpoints)
		require.NoError(t, err)
		// Should have original A + generated PTR
		assert.Len(t, result, 2)

		var ptrFound bool
		for _, ep := range result {
			if ep.RecordType == "PTR" {
				ptrFound = true
				assert.Equal(t, "8.8.8.8.in-addr.arpa", ep.DNSName)
				require.Len(t, ep.Targets, 1)
				assert.Equal(t, "example.com", ep.Targets[0])
			}
		}
		assert.True(t, ptrFound, "expected PTR endpoint in adjusted endpoints")
	})

	t.Run("AAAA record generates PTR", func(t *testing.T) {
		mock := &mockProvider{}
		p := NewPTRProvider(mock)

		endpoints := []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeAAAA, endpoint.TTL(300), "2001:db8::1"),
		}

		result, err := p.AdjustEndpoints(endpoints)
		require.NoError(t, err)
		assert.Len(t, result, 2)

		var ptrFound bool
		for _, ep := range result {
			if ep.RecordType == "PTR" {
				ptrFound = true
				assert.Equal(t, "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa", ep.DNSName)
			}
		}
		assert.True(t, ptrFound, "expected PTR endpoint for AAAA record")
	})

	t.Run("CNAME records are not affected", func(t *testing.T) {
		mock := &mockProvider{}
		p := NewPTRProvider(mock)

		endpoints := []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("cname.example.com", endpoint.RecordTypeCNAME, endpoint.TTL(300), "example.com"),
		}

		result, err := p.AdjustEndpoints(endpoints)
		require.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, endpoint.RecordTypeCNAME, result[0].RecordType)
	})

	t.Run("wildcard A records are skipped", func(t *testing.T) {
		mock := &mockProvider{}
		p := NewPTRProvider(mock)

		endpoints := []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("*.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "1.2.3.4"),
			endpoint.NewEndpointWithTTL("www.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "5.6.7.8"),
		}

		result, err := p.AdjustEndpoints(endpoints)
		require.NoError(t, err)
		// 2 original + 1 PTR (only for www, not wildcard)
		assert.Len(t, result, 3)

		var ptrCount int
		for _, ep := range result {
			if ep.RecordType == "PTR" {
				ptrCount++
				assert.Equal(t, "8.7.6.5.in-addr.arpa", ep.DNSName)
			}
		}
		assert.Equal(t, 1, ptrCount, "expected exactly 1 PTR (wildcard should be skipped)")
	})

	t.Run("multiple A records with same IP produce one merged PTR", func(t *testing.T) {
		mock := &mockProvider{}
		p := NewPTRProvider(mock)

		endpoints := []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("web.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "192.168.49.2"),
			endpoint.NewEndpointWithTTL("app.example.com", endpoint.RecordTypeA, endpoint.TTL(300), "192.168.49.2"),
		}

		result, err := p.AdjustEndpoints(endpoints)
		require.NoError(t, err)
		// 2 original A records + 1 merged PTR
		assert.Len(t, result, 3)

		var ptrCount int
		for _, ep := range result {
			if ep.RecordType == "PTR" {
				ptrCount++
				assert.Equal(t, "2.49.168.192.in-addr.arpa", ep.DNSName)
				assert.ElementsMatch(t, []string{"web.example.com", "app.example.com"}, ep.Targets)
			}
		}
		assert.Equal(t, 1, ptrCount, "expected exactly 1 merged PTR endpoint")
	})

	t.Run("empty endpoints pass through", func(t *testing.T) {
		mock := &mockProvider{}
		p := NewPTRProvider(mock)

		result, err := p.AdjustEndpoints([]*endpoint.Endpoint{})
		require.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("delegates to underlying provider AdjustEndpoints", func(t *testing.T) {
		mock := &mockProvider{}
		p := NewPTRProvider(mock)

		endpoints := []*endpoint.Endpoint{
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "8.8.8.8"),
		}

		// BaseProvider.AdjustEndpoints returns endpoints as-is
		result, err := p.AdjustEndpoints(endpoints)
		require.NoError(t, err)
		// Should include original + generated PTR
		assert.Len(t, result, 2)
	})

	t.Run("Records delegates to underlying provider", func(t *testing.T) {
		expected := []*endpoint.Endpoint{
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "8.8.8.8"),
		}
		mock := &mockProvider{recordsResult: expected}
		p := NewPTRProvider(mock)

		result, err := p.Records(context.Background())
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("ApplyChanges delegates to underlying provider unchanged", func(t *testing.T) {
		mock := &mockProvider{}
		p := NewPTRProvider(mock)

		changes := &plan.Changes{
			Create: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, endpoint.TTL(300), "8.8.8.8"),
			},
		}

		err := p.ApplyChanges(context.Background(), changes)
		require.NoError(t, err)
		require.Len(t, mock.appliedChanges, 1)
		// ApplyChanges should NOT augment changes — that's AdjustEndpoints' job
		assert.Len(t, mock.appliedChanges[0].Create, 1)
		assert.Equal(t, endpoint.RecordTypeA, mock.appliedChanges[0].Create[0].RecordType)
	})
}
