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

package source

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
)

func TestSuitableType(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		expected string
	}{
		{
			name:     "valid IPv4 address",
			target:   "192.168.1.1",
			expected: endpoint.RecordTypeA,
		},
		{
			name:     "valid IPv6 address",
			target:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expected: endpoint.RecordTypeAAAA,
		},
		{
			name:     "invalid IP address, should return CNAME",
			target:   "example.com",
			expected: endpoint.RecordTypeCNAME,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := suitableType(tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseIngress(t *testing.T) {
	tests := []struct {
		name      string
		ingress   string
		wantNS    string
		wantName  string
		wantError bool
	}{
		{
			name:      "valid namespace and name",
			ingress:   "default/test-ingress",
			wantNS:    "default",
			wantName:  "test-ingress",
			wantError: false,
		},
		{
			name:      "only name provided",
			ingress:   "test-ingress",
			wantNS:    "",
			wantName:  "test-ingress",
			wantError: false,
		},
		{
			name:      "invalid format",
			ingress:   "default/test/ingress",
			wantNS:    "",
			wantName:  "",
			wantError: true,
		},
		{
			name:      "empty string",
			ingress:   "",
			wantNS:    "",
			wantName:  "",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNS, gotName, err := ParseIngress(tt.ingress)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantNS, gotNS)
			assert.Equal(t, tt.wantName, gotName)
		})
	}
}

func TestSelectorMatchesService(t *testing.T) {
	tests := []struct {
		name        string
		selector    map[string]string
		svcSelector map[string]string
		expected    bool
	}{
		{
			name:        "all key-value pairs match",
			selector:    map[string]string{"app": "nginx", "env": "prod"},
			svcSelector: map[string]string{"app": "nginx", "env": "prod"},
			expected:    true,
		},
		{
			name:        "one key-value pair does not match",
			selector:    map[string]string{"app": "nginx", "env": "prod"},
			svcSelector: map[string]string{"app": "nginx", "env": "dev"},
			expected:    false,
		},
		{
			name:        "key not present in svcSelector",
			selector:    map[string]string{"app": "nginx", "env": "prod"},
			svcSelector: map[string]string{"app": "nginx"},
			expected:    false,
		},
		{
			name:        "empty selector",
			selector:    map[string]string{},
			svcSelector: map[string]string{"app": "nginx", "env": "prod"},
			expected:    true,
		},
		{
			name:        "empty svcSelector",
			selector:    map[string]string{"app": "nginx", "env": "prod"},
			svcSelector: map[string]string{},
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MatchesServiceSelector(tt.selector, tt.svcSelector)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMergeEndpoints(t *testing.T) {
	tests := []struct {
		name     string
		input    []*endpoint.Endpoint
		expected []*endpoint.Endpoint
	}{
		{
			name:     "nil input returns nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty input returns empty",
			input:    []*endpoint.Endpoint{},
			expected: []*endpoint.Endpoint{},
		},
		{
			name: "single endpoint unchanged",
			input: []*endpoint.Endpoint{
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			name: "different keys not merged",
			input: []*endpoint.Endpoint{
				{DNSName: "a.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "b.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"5.6.7.8"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "a.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "b.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"5.6.7.8"}},
			},
		},
		{
			name: "same DNSName different RecordType not merged",
			input: []*endpoint.Endpoint{
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
			},
		},
		{
			name: "same key merged with sorted targets",
			input: []*endpoint.Endpoint{
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"5.6.7.8"}},
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "5.6.7.8"}},
			},
		},
		{
			name: "multiple endpoints same key merged",
			input: []*endpoint.Endpoint{
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"3.3.3.3"}},
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"2.2.2.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "2.2.2.2", "3.3.3.3"}},
			},
		},
		{
			name: "mixed merge and no merge",
			input: []*endpoint.Endpoint{
				{DNSName: "a.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "b.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"2.2.2.2"}},
				{DNSName: "a.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"3.3.3.3"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "a.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "3.3.3.3"}},
				{DNSName: "b.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"2.2.2.2"}},
			},
		},
		{
			name: "duplicate targets deduplicated",
			input: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4", "1.2.3.4", "5.6.7.8"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4", "5.6.7.8"),
			},
		},
		{
			name: "duplicate targets across merged endpoints deduplicated",
			input: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4", "5.6.7.8"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4", "5.6.7.8"),
			},
		},
		{
			name: "CNAME endpoints not merged",
			input: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "b.elb.com"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "b.elb.com"),
			},
		},
		{
			name: "CNAME with no targets is skipped",
			input: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4"),
			},
		},
		{
			name: "identical CNAME endpoints deduplicated",
			input: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
			},
		},
		{
			name: "same key with different TTL not merged",
			input: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, 300, "1.2.3.4"),
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, 600, "5.6.7.8"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, 300, "1.2.3.4"),
				endpoint.NewEndpointWithTTL("example.com", endpoint.RecordTypeA, 600, "5.6.7.8"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeEndpoints(tt.input)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func TestMergeEndpointsLogging(t *testing.T) {
	t.Run("warns on CNAME conflict", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

		MergeEndpoints([]*endpoint.Endpoint{
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "b.elb.com"),
		})

		logtest.TestHelperLogContainsWithLogLevel("Only one CNAME per name", log.WarnLevel, hook, t)
		logtest.TestHelperLogContains("example.com CNAME a.elb.com", hook, t)
		logtest.TestHelperLogContains("example.com CNAME b.elb.com", hook, t)
	})

	t.Run("no warning for identical CNAMEs", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

		MergeEndpoints([]*endpoint.Endpoint{
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
		})

		logtest.TestHelperLogNotContains("Only one CNAME per name", hook, t)
	})

	t.Run("no warning for same DNSName with different SetIdentifier", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

		MergeEndpoints([]*endpoint.Endpoint{
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "a.elb.com").WithSetIdentifier("weight-1"),
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME, "b.elb.com").WithSetIdentifier("weight-2"),
		})

		logtest.TestHelperLogNotContains("Only one CNAME per name", hook, t)
	})

	t.Run("debug log for CNAME with no targets", func(t *testing.T) {
		hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

		MergeEndpoints([]*endpoint.Endpoint{
			endpoint.NewEndpoint("example.com", endpoint.RecordTypeCNAME),
		})

		logtest.TestHelperLogContainsWithLogLevel("Skipping CNAME endpoint", log.DebugLevel, hook, t)
		logtest.TestHelperLogContains("example.com", hook, t)
	})
}
