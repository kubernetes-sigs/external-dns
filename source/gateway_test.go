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

package source

import (
	"strings"
	"testing"

	v1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestGatewayMatchingHost(t *testing.T) {
	tests := []struct {
		desc string
		a, b string
		host string
		ok   bool
	}{
		{
			desc: "ipv4-rejected",
			a:    "1.2.3.4",
			ok:   false,
		},
		{
			desc: "ipv6-rejected",
			a:    "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			ok:   false,
		},
		{
			desc: "empty-matches-empty",
			ok:   true,
		},
		{
			desc: "empty-matches-nonempty",
			a:    "example.net",
			host: "example.net",
			ok:   true,
		},
		{
			desc: "simple-match",
			a:    "example.net",
			b:    "example.net",
			host: "example.net",
			ok:   true,
		},
		{
			desc: "wildcard-matches-longer",
			a:    "*.example.net",
			b:    "test.example.net",
			host: "test.example.net",
			ok:   true,
		},
		{
			desc: "wildcard-matches-equal-length",
			a:    "*.example.net",
			b:    "a.example.net",
			host: "a.example.net",
			ok:   true,
		},
		{
			desc: "wildcard-matches-multiple-subdomains",
			a:    "*.example.net",
			b:    "foo.bar.test.example.net",
			host: "foo.bar.test.example.net",
			ok:   true,
		},
		{
			desc: "wildcard-doesnt-match-parent",
			a:    "*.example.net",
			b:    "example.net",
			ok:   false,
		},
		{
			desc: "wildcard-must-be-complete-label",
			a:    "*example.net",
			b:    "test.example.net",
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			for i := 0; i < 2; i++ {
				if host, ok := gwMatchingHost(tt.a, tt.b); host != tt.host || ok != tt.ok {
					t.Errorf(
						"gwMatchingHost(%q, %q); got: %q, %v; want: %q, %v",
						tt.a, tt.b, host, ok, tt.host, tt.ok,
					)
				}
				tt.a, tt.b = tt.b, tt.a
			}
		})

	}
}

func TestGatewayMatchingProtocol(t *testing.T) {
	tests := []struct {
		route, lis string
		desc       string
		ok         bool
	}{
		{
			desc:  "protocol-matches-lis-https-route-http",
			route: "HTTP",
			lis:   "HTTPS",
			ok:    true,
		},
		{
			desc:  "protocol-match-invalid-list-https-route-tcp",
			route: "TCP",
			lis:   "HTTPS",
			ok:    false,
		},
		{
			desc:  "protocol-match-valid-lis-tls-route-tls",
			route: "TLS",
			lis:   "TLS",
			ok:    true,
		},
		{
			desc:  "protocol-match-valid-lis-TLS-route-TCP",
			route: "TCP",
			lis:   "TLS",
			ok:    true,
		},
		{
			desc:  "protocol-match-valid-lis-TLS-route-TCP",
			route: "TLS",
			lis:   "TCP",
			ok:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			for i := 0; i < 2; i++ {
				if ok := gwProtocolMatches(v1.ProtocolType(tt.route), v1.ProtocolType(tt.lis)); ok != tt.ok {
					t.Errorf(
						"gwProtocolMatches(%q, %q); got: %v; want: %v",
						tt.route, tt.lis, ok, tt.ok,
					)
				}
				// tt.a, tt.b = tt.b, tt.a
			}
		})

	}
}

func TestIsDNS1123Domain(t *testing.T) {
	tests := []struct {
		desc string
		in   string
		ok   bool
	}{
		{
			desc: "empty",
			ok:   false,
		},
		{
			desc: "label-too-long",
			in:   strings.Repeat("x", 64) + ".example.net",
			ok:   false,
		},
		{
			desc: "domain-too-long",
			in:   strings.Repeat("testing.", 256/(len("testing."))) + "example.net",
			ok:   false,
		},
		{
			desc: "hostname",
			in:   "example",
			ok:   true,
		},
		{
			desc: "domain",
			in:   "example.net",
			ok:   true,
		},
		{
			desc: "subdomain",
			in:   "test.example.net",
			ok:   true,
		},
		{
			desc: "dashes",
			in:   "test-with-dash.example.net",
			ok:   true,
		},
		{
			desc: "dash-prefix",
			in:   "-dash-prefix.example.net",
			ok:   false,
		},
		{
			desc: "dash-suffix",
			in:   "dash-suffix-.example.net",
			ok:   false,
		},
		{
			desc: "underscore",
			in:   "under_score.example.net",
			ok:   false,
		},
		{
			desc: "plus",
			in:   "pl+us.example.net",
			ok:   false,
		},
		{
			desc: "brackets",
			in:   "bra[k]ets.example.net",
			ok:   false,
		},
		{
			desc: "parens",
			in:   "pa[re]ns.example.net",
			ok:   false,
		},
		{
			desc: "wild",
			in:   "*.example.net",
			ok:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if ok := isDNS1123Domain(tt.in); ok != tt.ok {
				t.Errorf("isDNS1123Domain(%q); got: %v; want: %v", tt.in, ok, tt.ok)
			}
		})
	}
}

func TestMergeAnnotations(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc     string
		gateway  map[string]string
		route    map[string]string
		expected map[string]string
	}{
		{
			desc:     "nil gateway annotations",
			gateway:  nil,
			route:    map[string]string{"key": "route-value"},
			expected: map[string]string{"key": "route-value"},
		},
		{
			desc:     "nil route annotations",
			gateway:  map[string]string{"key": "gateway-value"},
			route:    nil,
			expected: map[string]string{"key": "gateway-value"},
		},
		{
			desc:     "both nil",
			gateway:  nil,
			route:    nil,
			expected: map[string]string{},
		},
		{
			desc:     "empty gateway annotations",
			gateway:  map[string]string{},
			route:    map[string]string{"key": "route-value"},
			expected: map[string]string{"key": "route-value"},
		},
		{
			desc:     "empty route annotations",
			gateway:  map[string]string{"key": "gateway-value"},
			route:    map[string]string{},
			expected: map[string]string{"key": "gateway-value"},
		},
		{
			desc:     "both empty",
			gateway:  map[string]string{},
			route:    map[string]string{},
			expected: map[string]string{},
		},
		{
			desc:     "route overrides gateway",
			gateway:  map[string]string{"key": "gateway-value"},
			route:    map[string]string{"key": "route-value"},
			expected: map[string]string{"key": "route-value"},
		},
		{
			desc:     "merge different keys",
			gateway:  map[string]string{"gw-key": "gw-value"},
			route:    map[string]string{"rt-key": "rt-value"},
			expected: map[string]string{"gw-key": "gw-value", "rt-key": "rt-value"},
		},
		{
			desc: "partial override with real annotation keys",
			gateway: map[string]string{
				"external-dns.alpha.kubernetes.io/target": "172.16.6.6",
				"external-dns.alpha.kubernetes.io/ttl":    "300",
			},
			route: map[string]string{
				"external-dns.alpha.kubernetes.io/target": "1.2.3.4",
			},
			expected: map[string]string{
				"external-dns.alpha.kubernetes.io/target": "1.2.3.4",
				"external-dns.alpha.kubernetes.io/ttl":    "300",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			result := mergeAnnotations(tt.gateway, tt.route)
			if len(result) != len(tt.expected) {
				t.Errorf("mergeAnnotations(); got len %d; want len %d", len(result), len(tt.expected))
			}
			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("mergeAnnotations()[%q]; got %q; want %q", k, result[k], v)
				}
			}
		})
	}
}

func TestMergeAnnotationsDoesNotMutateInputs(t *testing.T) {
	t.Parallel()
	gateway := map[string]string{"gw-key": "gw-value", "shared": "gateway"}
	route := map[string]string{"rt-key": "rt-value", "shared": "route"}

	// Copy original values
	gwOriginal := make(map[string]string)
	for k, v := range gateway {
		gwOriginal[k] = v
	}
	rtOriginal := make(map[string]string)
	for k, v := range route {
		rtOriginal[k] = v
	}

	result := mergeAnnotations(gateway, route)

	// Verify result is correct
	if result["shared"] != "route" {
		t.Errorf("expected route to override gateway for 'shared' key")
	}

	// Verify inputs were not mutated
	for k, v := range gwOriginal {
		if gateway[k] != v {
			t.Errorf("gateway map was mutated: key %q changed from %q to %q", k, v, gateway[k])
		}
	}
	for k, v := range rtOriginal {
		if route[k] != v {
			t.Errorf("route map was mutated: key %q changed from %q to %q", k, v, route[k])
		}
	}

	// Verify modifying result doesn't affect inputs
	result["new-key"] = "new-value"
	if _, ok := gateway["new-key"]; ok {
		t.Error("modifying result affected gateway map")
	}
	if _, ok := route["new-key"]; ok {
		t.Error("modifying result affected route map")
	}
}
