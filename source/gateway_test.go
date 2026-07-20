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
	"context"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"
)

var _ Source = &gatewaySource{}

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
			for range 2 {
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
			for range 2 {
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

func gatewayStatusHostname(hostnames ...string) v1.GatewayStatus {
	typ := v1.HostnameAddressType
	addrs := make([]v1.GatewayStatusAddress, len(hostnames))
	for i, h := range hostnames {
		addrs[i] = v1.GatewayStatusAddress{Type: &typ, Value: h}
	}
	return v1.GatewayStatus{Addresses: addrs}
}

func makeGateway(namespace, name string, annots map[string]string, status v1.GatewayStatus, labels map[string]string) *v1.Gateway {
	return &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   namespace,
			Annotations: annots,
			Labels:      labels,
		},
		Status: status,
	}
}

func TestGatewaySourceEndpoints(t *testing.T) {
	t.Parallel()

	hostnameAnnotation := func(hostnames ...string) map[string]string {
		return map[string]string{annotations.HostnameKey: strings.Join(hostnames, ",")}
	}

	tests := []struct {
		title     string
		config    *Config
		gateways  []*v1.Gateway
		endpoints []*endpoint.Endpoint
	}{
		{
			title:  "IP address in status produces A record",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "hostname address in status produces CNAME record",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatusHostname("lb.example.com"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpointWithTTL("foo.example.com", endpoint.RecordTypeCNAME, 0, "lb.example.com"),
			},
		},
		{
			title:  "IPv6 address in status produces AAAA record",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatus("2001:db8::1"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpointWithTTL("foo.example.com", endpoint.RecordTypeAAAA, 0, "2001:db8::1"),
			},
		},
		{
			title:  "multiple hostnames in annotation produce multiple endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("a.example.com", "b.example.com"), gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
				newTestEndpoint("b.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "multiple gateways produce multiple endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", hostnameAnnotation("a.example.com"), gatewayStatus("1.2.3.4"), nil),
				makeGateway("default", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
				newTestEndpoint("b.example.com", "5.6.7.8"),
			},
		},
		{
			title:  "GatewayName filter skips non-matching gateways",
			config: &Config{GatewayName: "gw1"},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", hostnameAnnotation("a.example.com"), gatewayStatus("1.2.3.4"), nil),
				makeGateway("default", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "GatewayNamespace filter limits to configured namespace",
			config: &Config{GatewayNamespace: "ns1"},
			gateways: []*v1.Gateway{
				makeGateway("ns1", "gw1", hostnameAnnotation("a.example.com"), gatewayStatus("1.2.3.4"), nil),
				makeGateway("ns2", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "Namespace filter limits to configured namespace",
			config: &Config{Namespace: "ns1"},
			gateways: []*v1.Gateway{
				makeGateway("ns1", "gw1", hostnameAnnotation("a.example.com"), gatewayStatus("1.2.3.4"), nil),
				makeGateway("ns2", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "GatewayLabelFilter skips gateways without matching label",
			config: &Config{GatewayLabelFilter: "env=prod"},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", hostnameAnnotation("a.example.com"), gatewayStatus("1.2.3.4"), map[string]string{"env": "prod"}),
				makeGateway("default", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), map[string]string{"env": "staging"}),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "AnnotationFilter skips gateways without matching annotation",
			config: &Config{AnnotationFilter: parseAnnotationFilterOrNil("custom=yes")},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", map[string]string{
					annotations.HostnameKey: "a.example.com",
					"custom":                "yes",
				}, gatewayStatus("1.2.3.4"), nil),
				makeGateway("default", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "IgnoreHostnameAnnotation produces no endpoints",
			config: &Config{IgnoreHostnameAnnotation: true},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: nil,
		},
		{
			title:  "no hostname annotation produces no endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", nil, gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: nil,
		},
		{
			title:  "no status addresses produces no endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), v1.GatewayStatus{}, nil),
			},
			endpoints: nil,
		},
		{
			title:  "target override annotation replaces status addresses",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", map[string]string{
					annotations.HostnameKey: "foo.example.com",
					annotations.TargetKey:   "override.example.com",
				}, gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpointWithTTL("foo.example.com", endpoint.RecordTypeCNAME, 0, "override.example.com"),
			},
		},
		{
			title:  "TTL annotation is applied to endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", map[string]string{
					annotations.HostnameKey: "foo.example.com",
					annotations.TtlKey:      "300",
				}, gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpointWithTTL("foo.example.com", endpoint.RecordTypeA, 300, "1.2.3.4"),
			},
		},
		{
			title:  "controller annotation mismatch skips gateway",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", map[string]string{
					annotations.HostnameKey:   "foo.example.com",
					annotations.ControllerKey: "other-controller",
				}, gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: nil,
		},
		{
			title:  "same hostname from two gateways merges targets",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", hostnameAnnotation("shared.example.com"), gatewayStatus("1.2.3.4"), nil),
				makeGateway("default", "gw2", hostnameAnnotation("shared.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("shared.example.com", "1.2.3.4", "5.6.7.8"),
			},
		},
		{
			title:  "multiple IP addresses produce multi-value A record",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatus("1.2.3.4", "5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.com", "1.2.3.4", "5.6.7.8"),
			},
		},
		{
			title:  "multiple hostname addresses produce multi-value CNAME record",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatusHostname("lb1.example.com", "lb2.example.com"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpointWithTTL("foo.example.com", endpoint.RecordTypeCNAME, 0, "lb1.example.com", "lb2.example.com"),
			},
		},
		{
			title:  "mixed IP and hostname addresses uses only IPs",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), func() v1.GatewayStatus {
					ipTyp := v1.IPAddressType
					hnTyp := v1.HostnameAddressType
					return v1.GatewayStatus{Addresses: []v1.GatewayStatusAddress{
						{Type: &ipTyp, Value: "1.2.3.4"},
						{Type: &hnTyp, Value: "lb.example.com"},
					}}
				}(), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "nil address type treated as IPAddress",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), v1.GatewayStatus{
					Addresses: []v1.GatewayStatusAddress{
						{Type: nil, Value: "1.2.3.4"},
					},
				}, nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "NamedAddress type is skipped",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), func() v1.GatewayStatus {
					namedTyp := v1.NamedAddressType
					return v1.GatewayStatus{Addresses: []v1.GatewayStatusAddress{
						{Type: &namedTyp, Value: "my-named-address"},
					}}
				}(), nil),
			},
			endpoints: nil,
		},
		{
			title:  "empty address value is skipped",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), func() v1.GatewayStatus {
					ipTyp := v1.IPAddressType
					return v1.GatewayStatus{Addresses: []v1.GatewayStatusAddress{
						{Type: &ipTyp, Value: ""},
					}}
				}(), nil),
			},
			endpoints: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
			defer cancel()

			gwClient := gatewayfake.NewSimpleClientset()
			for _, gw := range tt.gateways {
				_, err := gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
				require.NoError(t, err, "failed to create Gateway")
			}

			clients := new(testutils.MockClientGenerator)
			clients.On("GatewayClient").Return(gwClient, nil)

			src, err := NewGatewaySource(ctx, clients, tt.config)
			require.NoError(t, err, "failed to create Gateway source")

			endpoints, err := src.Endpoints(ctx)
			require.NoError(t, err, "failed to get endpoints")

			testutils.ValidateEndpoints(t, endpoints, tt.endpoints)
		})
	}
}
