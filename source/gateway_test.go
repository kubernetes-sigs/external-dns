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
	"fmt"
	"strconv"
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

func TestGatewayRouteCurrentParentStatuses(t *testing.T) {
	meta := &metav1.ObjectMeta{Namespace: "default"}

	statusCondition := func(status metav1.ConditionStatus) []metav1.Condition {
		return []metav1.Condition{{Type: string(v1.RouteConditionAccepted), Status: status}}
	}
	statusesFor := func(refs ...v1.ParentReference) []v1.RouteParentStatus {
		statuses := make([]v1.RouteParentStatus, 0, len(refs))
		for _, ref := range refs {
			statuses = append(statuses, v1.RouteParentStatus{
				ParentRef:  ref,
				Conditions: statusCondition(metav1.ConditionTrue),
			})
		}
		return statuses
	}
	rejected := func(ref v1.ParentReference) v1.RouteParentStatus {
		return v1.RouteParentStatus{ParentRef: ref, Conditions: statusCondition(metav1.ConditionFalse)}
	}

	// render identifies a status by the parts of its ParentRef that select listeners.
	render := func(statuses []v1.RouteParentStatus) []string {
		out := make([]string, 0, len(statuses))
		for _, status := range statuses {
			port := "-"
			if status.ParentRef.Port != nil {
				port = strconv.Itoa(int(*status.ParentRef.Port))
			}
			out = append(out, fmt.Sprintf("%s/%s/%s/%s",
				strVal((*string)(status.ParentRef.Kind), gatewayKind),
				status.ParentRef.Name, sectionVal(status.ParentRef.SectionName, "-"), port))
		}
		return out
	}

	tests := []struct {
		desc     string
		refs     []v1.ParentReference
		statuses []v1.RouteParentStatus
		want     []string
	}{
		{
			desc: "drops the status of a listener the route left",
			refs: []v1.ParentReference{gwParentRef("default", "gw", withSectionName("bar"))},
			statuses: statusesFor(
				gwParentRef("default", "gw", withSectionName("foo")),
				gwParentRef("default", "gw", withSectionName("bar")),
			),
			want: []string{"Gateway/gw/bar/-"},
		},
		{
			desc:     "keeps a status the controller has not updated yet",
			refs:     []v1.ParentReference{gwParentRef("default", "gw", withSectionName("bar"))},
			statuses: statusesFor(gwParentRef("default", "gw", withSectionName("foo"))),
			want:     []string{"Gateway/gw/foo/-"},
		},
		{
			desc: "keeps every listener the route currently attaches to",
			refs: []v1.ParentReference{
				gwParentRef("default", "gw", withSectionName("foo")),
				gwParentRef("default", "gw", withSectionName("bar")),
			},
			statuses: statusesFor(
				gwParentRef("default", "gw", withSectionName("foo")),
				gwParentRef("default", "gw", withSectionName("bar")),
			),
			want: []string{"Gateway/gw/foo/-", "Gateway/gw/bar/-"},
		},
		{
			desc: "distinguishes references by port",
			refs: []v1.ParentReference{gwParentRef("default", "gw", withPortNumber(8080))},
			statuses: statusesFor(
				gwParentRef("default", "gw", withPortNumber(80)),
				gwParentRef("default", "gw", withPortNumber(8080)),
			),
			want: []string{"Gateway/gw/-/8080"},
		},
		{
			desc:     "a status without a port does not match a ref that names one",
			refs:     []v1.ParentReference{gwParentRef("default", "gw", withPortNumber(8080))},
			statuses: statusesFor(gwParentRef("default", "gw")),
			want:     []string{"Gateway/gw/-/-"},
		},
		{
			desc: "does not depend on the order of the status list",
			refs: []v1.ParentReference{gwParentRef("default", "gw", withSectionName("bar"))},
			statuses: statusesFor(
				gwParentRef("default", "gw", withSectionName("bar")),
				gwParentRef("default", "gw", withSectionName("foo")),
			),
			want: []string{"Gateway/gw/bar/-"},
		},
		{
			desc:     "keeps sectioned statuses when the route names no section",
			refs:     []v1.ParentReference{gwParentRef("default", "gw")},
			statuses: statusesFor(gwParentRef("default", "gw", withSectionName("foo"))),
			want:     []string{"Gateway/gw/foo/-"},
		},
		{
			desc: "treats each parent object independently",
			refs: []v1.ParentReference{
				gwParentRef("default", "gw-a", withSectionName("bar")),
				gwParentRef("default", "gw-b"),
			},
			statuses: statusesFor(
				gwParentRef("default", "gw-a", withSectionName("foo")),
				gwParentRef("default", "gw-a", withSectionName("bar")),
				gwParentRef("default", "gw-b", withSectionName("foo")),
			),
			want: []string{"Gateway/gw-a/bar/-", "Gateway/gw-b/foo/-"},
		},
		{
			// A ListenerSet parent goes through the same resolver, and a Gateway of
			// the same name is a different parent object.
			desc: "filters ListenerSet parents and keeps them apart from a Gateway",
			refs: []v1.ParentReference{
				lsParentRef("default", "shared", withSectionName("bar")),
			},
			statuses: statusesFor(
				lsParentRef("default", "shared", withSectionName("foo")),
				lsParentRef("default", "shared", withSectionName("bar")),
				gwParentRef("default", "shared", withSectionName("foo")),
			),
			want: []string{"ListenerSet/shared/bar/-", "Gateway/shared/foo/-"},
		},
		{
			// One reference is moving from "old" to "new" while "stable" stays put.
			// Dropping "old" here would take the hostname down before "new" arrives.
			desc: "keeps a stale entry while another current listener has no status yet",
			refs: []v1.ParentReference{
				gwParentRef("default", "gw", withSectionName("stable")),
				gwParentRef("default", "gw", withSectionName("new")),
			},
			statuses: statusesFor(
				gwParentRef("default", "gw", withSectionName("stable")),
				gwParentRef("default", "gw", withSectionName("old")),
			),
			want: []string{"Gateway/gw/stable/-", "Gateway/gw/old/-"},
		},
		{
			desc: "prunes once every current listener has an accepted status",
			refs: []v1.ParentReference{
				gwParentRef("default", "gw", withSectionName("stable")),
				gwParentRef("default", "gw", withSectionName("new")),
			},
			statuses: statusesFor(
				gwParentRef("default", "gw", withSectionName("stable")),
				gwParentRef("default", "gw", withSectionName("old")),
				gwParentRef("default", "gw", withSectionName("new")),
			),
			want: []string{"Gateway/gw/stable/-", "Gateway/gw/new/-"},
		},
		{
			// The resolver rejects a status that is not accepted, so it cannot stand
			// as proof that the older entry is safe to drop.
			desc: "keeps a stale entry while the current listener is not accepted",
			refs: []v1.ParentReference{gwParentRef("default", "gw", withSectionName("bar"))},
			statuses: append(
				statusesFor(gwParentRef("default", "gw", withSectionName("foo"))),
				rejected(gwParentRef("default", "gw", withSectionName("bar"))),
			),
			want: []string{"Gateway/gw/foo/-", "Gateway/gw/bar/-"},
		},
		{
			desc:     "handles a route with no status",
			refs:     []v1.ParentReference{gwParentRef("default", "gw")},
			statuses: nil,
			want:     []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := gwRouteCurrentParentStatuses(meta, tt.refs, tt.statuses)
			require.Equal(t, tt.want, render(got))
		})
	}
}

func TestGatewayUniqueTargets(t *testing.T) {
	tests := []struct {
		desc string
		in   endpoint.Targets
		want endpoint.Targets
	}{
		{
			desc: "no targets",
			in:   endpoint.Targets{},
			want: endpoint.Targets{},
		},
		{
			desc: "a single target is returned as is",
			in:   endpoint.Targets{"1.2.3.4"},
			want: endpoint.Targets{"1.2.3.4"},
		},
		{
			desc: "distinct targets are sorted",
			in:   endpoint.Targets{"5.6.7.8", "1.2.3.4"},
			want: endpoint.Targets{"1.2.3.4", "5.6.7.8"},
		},
		{
			desc: "a target listed by two listeners appears once",
			in:   endpoint.Targets{"1.2.3.4", "5.6.7.8", "1.2.3.4"},
			want: endpoint.Targets{"1.2.3.4", "5.6.7.8"},
		},
		{
			desc: "every listener reports the same target",
			in:   endpoint.Targets{"1.2.3.4", "1.2.3.4", "1.2.3.4"},
			want: endpoint.Targets{"1.2.3.4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			require.Equal(t, tt.want, uniqueTargets(tt.in))
		})
	}
}
