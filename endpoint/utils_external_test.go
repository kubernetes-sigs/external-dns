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
	"testing"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

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
