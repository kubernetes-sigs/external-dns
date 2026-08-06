/*
Copyright 2026 The Kubernetes Authors.

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

package endpoint_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestAppendIfNotNil(t *testing.T) {
	first := endpoint.NewEndpoint("first.example.com", endpoint.RecordTypeA, "192.168.1.1")
	second := endpoint.NewEndpoint("second.example.com", endpoint.RecordTypeA, "192.168.1.2")
	require.NotNil(t, first)
	require.NotNil(t, second)

	t.Run("appends an endpoint to a nil slice", func(t *testing.T) {
		assert.Equal(t, []*endpoint.Endpoint{first}, endpoint.AppendIfNotNil(nil, first))
	})

	t.Run("appends an endpoint to a populated slice", func(t *testing.T) {
		got := endpoint.AppendIfNotNil([]*endpoint.Endpoint{first}, second)
		assert.Equal(t, []*endpoint.Endpoint{first, second}, got)
	})

	t.Run("drops a nil endpoint and leaves the slice untouched", func(t *testing.T) {
		got := endpoint.AppendIfNotNil([]*endpoint.Endpoint{first}, nil)
		assert.Equal(t, []*endpoint.Endpoint{first}, got)
	})

	t.Run("a nil endpoint leaves a nil slice nil rather than empty", func(t *testing.T) {
		var endpoints []*endpoint.Endpoint

		got := endpoint.AppendIfNotNil(endpoints, nil)

		// Nil rather than merely empty: EndpointsForHostname hands its slice
		// straight back to callers that compare it against []*Endpoint(nil),
		// so the helper has to return the input rather than allocate.
		assert.Nil(t, got)
	})

	t.Run("drops the endpoint a rejected DNS name produces", func(t *testing.T) {
		rejected := endpoint.NewEndpoint(strings.Repeat("a", 64)+".example.com", endpoint.RecordTypeA, "192.168.1.1")
		require.Nil(t, rejected, "an over-long label must be rejected for this case to prove anything")
		assert.Empty(t, endpoint.AppendIfNotNil(nil, rejected))
	})
}

func TestEndpointsForHostsAndTargets(t *testing.T) {
	tests := []struct {
		name      string
		hostnames []string
		targets   []string
		expected  []*endpoint.Endpoint
	}{
		{
			name:      "nil hostnames returns nil",
			hostnames: nil,
			targets:   []string{"192.168.1.1"},
			expected:  nil,
		},
		{
			name:      "nil targets returns nil",
			hostnames: []string{"example.com"},
			targets:   nil,
			expected:  nil,
		},
		{
			name:      "empty hostnames returns nil",
			hostnames: []string{},
			targets:   []string{"192.168.1.1"},
			expected:  nil,
		},
		{
			name:      "empty targets returns nil",
			hostnames: []string{"example.com"},
			targets:   []string{},
			expected:  nil,
		},
		{
			name:      "single hostname single IPv4 target",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1"),
			},
		},
		{
			name:      "single hostname single IPv6 target",
			hostnames: []string{"example.com"},
			targets:   []string{"2001:db8::1"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001:db8::1"),
			},
		},
		{
			name:      "single hostname CNAME target",
			hostnames: []string{"example.com"},
			targets:   []string{"lb.example.com"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			name:      "duplicate hostnames are deduplicated",
			hostnames: []string{"example.com", "example.com"},
			targets:   []string{"192.168.1.1"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1"),
			},
		},
		{
			name:      "duplicate targets of same type are deduplicated",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1", "192.168.1.1", "192.168.1.2"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1", "192.168.1.2"),
			},
		},
		{
			name:      "hostname with trailing dot is trimmed",
			hostnames: []string{"example.com."},
			targets:   []string{"192.168.1.1"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1"),
			},
		},
		{
			name:      "duplicate hostname with IPv4 and IPv6 targets",
			hostnames: []string{"example.com", "example.com"},
			targets:   []string{"192.168.1.1", "192.168.1.1", "2001:db8::1"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001:db8::1"),
			},
		},
		{
			name:      "multiple hostnames with single target",
			hostnames: []string{"example.com", "www.example.com"},
			targets:   []string{"192.168.1.1"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1"),
				endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "192.168.1.1"),
			},
		},
		{
			name:      "all three record types for same hostname",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1", "2001:db8::1", "lb.example.com"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001:db8::1"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "lb.example.com"),
			},
		},
		{
			name:      "multiple of each type grouped and sorted",
			hostnames: []string{"example.com"},
			targets:   []string{"192.168.1.1", "192.168.1.2", "2001:db8::1", "2001:db8::2", "a.example.com", "b.example.com"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "192.168.1.1", "192.168.1.2"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeAAAA, "2001:db8::1", "2001:db8::2"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.example.com", "b.example.com"),
			},
		},
		{
			name:      "output hostnames are sorted",
			hostnames: []string{"z.example.com", "a.example.com"},
			targets:   []string{"192.168.1.1"},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("a.example.com", endpoint.RecordTypeA, "192.168.1.1"),
				endpoint.NewEndpoint("z.example.com", endpoint.RecordTypeA, "192.168.1.1"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := endpoint.EndpointsForHostsAndTargets(tc.hostnames, tc.targets)
			if tc.expected == nil {
				assert.Nil(t, result)
				return
			}
			testutils.ValidateEndpoints(t, result, tc.expected)
		})
	}
}

// TestEndpointsForHostsAndTargetsSkipsRejectedHostnames covers a hostname the
// constructor refuses to build an endpoint for. It must be dropped rather than
// collected as a nil element: callers dereference every element of the result,
// so a leaked nil panics in MergeEndpoints and in the label loop of the
// unstructured source rather than surfacing as a rejected hostname.
func TestEndpointsForHostsAndTargetsSkipsRejectedHostnames(t *testing.T) {
	// The two names differ only in the length of the first label, 64 characters
	// against the 63 that RFC 1035 section 2.3.4 allows.
	rejected := strings.Repeat("a", 64) + ".example.com"
	accepted := strings.Repeat("a", 63) + ".example.com"

	t.Run("hostname is dropped rather than collected as nil", func(t *testing.T) {
		result := endpoint.EndpointsForHostsAndTargets([]string{rejected}, []string{"192.168.1.1"})
		assert.Empty(t, result)
		assert.NotContains(t, result, (*endpoint.Endpoint)(nil))
	})

	t.Run("remaining hostnames are still returned", func(t *testing.T) {
		result := endpoint.EndpointsForHostsAndTargets([]string{rejected, accepted}, []string{"192.168.1.1"})
		testutils.ValidateEndpoints(t, result, []*endpoint.Endpoint{
			endpoint.NewEndpoint(accepted, endpoint.RecordTypeA, "192.168.1.1"),
		})
	})

	t.Run("result is safe to merge", func(t *testing.T) {
		result := endpoint.EndpointsForHostsAndTargets([]string{rejected}, []string{"192.168.1.1"})
		assert.NotPanics(t, func() {
			endpoint.MergeEndpoints(result)
		})
	})
}
