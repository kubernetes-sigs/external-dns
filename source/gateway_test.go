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

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "sigs.k8s.io/gateway-api/apis/v1"

	"sigs.k8s.io/external-dns/endpoint"
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

func TestGatewayRouteStatusIsCurrent(t *testing.T) {
	ref := gwParentRef("default", "gw", withSectionName("bar"))

	condition := func(status metav1.ConditionStatus, observed int64) v1.RouteParentStatus {
		return v1.RouteParentStatus{
			ParentRef: ref,
			Conditions: []metav1.Condition{{
				Type:               string(v1.RouteConditionAccepted),
				Status:             status,
				ObservedGeneration: observed,
			}},
		}
	}

	tests := []struct {
		desc       string
		status     v1.RouteParentStatus
		generation int64
		want       bool
	}{
		{"accepted for the generation in hand", condition(metav1.ConditionTrue, 3), 3, true},
		{"accepted, but for an older generation", condition(metav1.ConditionTrue, 1), 3, false},
		{"accepted, but for a generation newer than the object", condition(metav1.ConditionTrue, 4), 3, false},
		{"accepted, but with no generation recorded", condition(metav1.ConditionTrue, 0), 3, false},
		{"not accepted", condition(metav1.ConditionFalse, 3), 3, false},
		{"acceptance still Unknown", condition(metav1.ConditionUnknown, 3), 3, false},
		{"no accepted condition at all", v1.RouteParentStatus{ParentRef: ref}, 3, false},
		{
			"looks past the conditions a controller writes alongside Accepted",
			v1.RouteParentStatus{
				ParentRef: ref,
				Conditions: []metav1.Condition{
					{Type: string(v1.RouteConditionResolvedRefs), Status: metav1.ConditionTrue, ObservedGeneration: 1},
					{Type: string(v1.RouteConditionAccepted), Status: metav1.ConditionTrue, ObservedGeneration: 3},
				},
			},
			3, true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			require.Equal(t, tt.want, gwRouteStatusIsCurrent(tt.status, tt.generation))
		})
	}
}

func TestGatewayRouteHasParentRef(t *testing.T) {
	meta := &metav1.ObjectMeta{Namespace: "default"}
	gatewayAPIGroup := v1.Group(gatewayGroup)

	omitted := v1.ParentReference{Name: "gw"}
	explicitGateway := v1.ParentReference{Name: "gw", Group: &gatewayAPIGroup}

	tests := []struct {
		desc string
		spec v1.ParentReference
		got  v1.ParentReference
		want bool
	}{
		{"an omitted group matches an omitted group", omitted, omitted, true},
		{"an omitted group matches the spelled out Gateway API group", omitted, explicitGateway, true},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			require.Equal(t, tt.want, gwRouteHasParentRef([]v1.ParentReference{tt.spec}, tt.got, meta))
		})
	}
}

func TestGatewayMatchRouteToParent(t *testing.T) {
	c := &gatewayRouteResolver{src: &gatewayRouteSource{rtKind: "HTTPRoute"}}
	rt := &gatewayHTTPRoute{route: v1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Namespace: "default", Name: "route"},
	}}
	rtHosts := []string{"example.net"}

	// A listener with no hostname matches any of the route's hostnames.
	parent := func(gw *v1.Gateway, overrides endpoint.Targets) *resolvedParent {
		return &resolvedParent{
			kind:      "Gateway",
			namespace: "default",
			ref:       gwParentRef("default", "gw"),
			obj: &listenerObject{
				gateway:        gw,
				ownerNamespace: "default",
				overrides:      overrides,
			},
			listeners: []listenerSection{{
				listener: v1.Listener{Name: "web", Protocol: v1.HTTPProtocolType},
				attached: true,
			}},
		}
	}

	withAddress := &v1.Gateway{Status: v1.GatewayStatus{
		Addresses: []v1.GatewayStatusAddress{{Value: "203.0.113.1"}},
	}}

	tests := []struct {
		desc    string
		parent  *resolvedParent
		want    bool
		targets endpoint.Targets
	}{
		{"a matched listener with no target does not attach", parent(&v1.Gateway{}, nil), false, nil},
		{"a Gateway status address is a target", parent(withAddress, nil), true, endpoint.Targets{"203.0.113.1"}},
		{"a target annotation override is a target", parent(&v1.Gateway{}, endpoint.Targets{"lb.example.net"}), true, endpoint.Targets{"lb.example.net"}},
		{"a blank target annotation override is not a target", parent(&v1.Gateway{}, endpoint.Targets{"", " "}), false, nil},
		{"a blank override falls back to the Gateway address", parent(withAddress, endpoint.Targets{"", " "}), true, endpoint.Targets{"203.0.113.1"}},
		{"a usable override keeps the Gateway address out", parent(withAddress, endpoint.Targets{"lb.example.net", " "}), true, endpoint.Targets{"lb.example.net"}},
	}

	// Targets are asserted too: filtering blank overrides decides whether the Gateway
	// address is used instead.
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			hostTargets := map[string]endpoint.Targets{}
			require.Equal(t, tt.want, c.matchRouteToParent(rt, rtHosts, tt.parent, hostTargets))
			require.Equal(t, tt.targets, hostTargets["example.net"])
		})
	}
}
