package wrappers

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
)

type mockSource struct {
	endpoints []*endpoint.Endpoint
	err       error
}

func (m *mockSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	for _, ep := range m.endpoints {
		if ep == nil {
			return m.endpoints, fmt.Errorf("skipped nil endpoint")
		}
	}
	return m.endpoints, m.err
}
func (m *mockSource) AddEventHandler(_ context.Context, _ func()) {}

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
			title: "skip endpoints processing as nill endpoint detected",
			ttl:   "0s",
			endpoints: []*endpoint.Endpoint{
				nil,
				endpoint.NewEndpointWithTTL("foo-2", "A", 60, "1.2.3.5"),
			},
			expected: []*endpoint.Endpoint{
				nil,
				endpoint.NewEndpointWithTTL("foo-2", "A", 60, "1.2.3.5"),
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.title, func(t *testing.T) {

			ms := mockSource{endpoints: tt.endpoints}
			ttl, _ := time.ParseDuration(tt.ttl)
			src := NewPostProcessor(&ms, WithTTL(ttl))

			endpoints, err := src.Endpoints(context.Background())
			if tt.expectErr {
				require.Error(t, err, "expected error for test case: %s", tt.title)
				return
			}
			validateEndpoints(t, endpoints, tt.expected)
		})
	}
}
