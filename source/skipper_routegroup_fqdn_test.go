/*
Copyright 2017 The Kubernetes Authors.

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

	rgv1 "github.com/szuecs/routegroup-client/apis/zalando.org/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

func TestRouteGroupFQDNTemplate(t *testing.T) {
	t.Parallel()

	lb := []rgv1.RouteGroupLoadBalancer{{Hostname: "lb.example.com"}}

	tests := []struct {
		title        string
		routeGroup   *rgv1.RouteGroup
		fqdnTemplate string
		combine      bool
		expected     []*endpoint.Endpoint
	}{
		{
			title:        "fqdn-template generates endpoint when no spec hosts",
			routeGroup:   createTestRouteGroup("ns", "my-rg", nil, nil, lb),
			fqdnTemplate: "{{.Name}}.example.com",
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "my-rg.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.example.com"},
					Labels:     endpoint.Labels{endpoint.ResourceLabelKey: "routegroup/ns/my-rg"},
				},
			},
		},
		{
			title:        "fqdn-template with combine=true adds template endpoint alongside spec-hosts endpoint",
			routeGroup:   createTestRouteGroup("ns", "my-rg", nil, []string{"spec.example.com"}, lb),
			fqdnTemplate: "{{.Name}}.example.com",
			combine:      true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "spec.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.example.com"},
					Labels:     endpoint.Labels{endpoint.ResourceLabelKey: "routegroup/ns/my-rg"},
				},
				{
					DNSName:    "my-rg.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.example.com"},
					Labels:     endpoint.Labels{endpoint.ResourceLabelKey: "routegroup/ns/my-rg"},
				},
			},
		},
		{
			title:        "fqdn-template without combine is ignored when spec-hosts exist",
			routeGroup:   createTestRouteGroup("ns", "my-rg", nil, []string{"spec.example.com"}, lb),
			fqdnTemplate: "{{.Name}}.example.com",
			combine:      false,
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "spec.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.example.com"},
					Labels:     endpoint.Labels{endpoint.ResourceLabelKey: "routegroup/ns/my-rg"},
				},
			},
		},
		{
			title: "fqdn-template can reference .Kind",
			routeGroup: &rgv1.RouteGroup{
				APIVersion: "zalando.org/v1", Kind: "RouteGroup",
				Name: "my-rg", Namespace: "ns",
				Status: rgv1.RouteGroupStatus{
					LoadBalancer: rgv1.RouteGroupLoadBalancerStatus{RouteGroup: lb},
				},
			},
			fqdnTemplate: "{{.Kind | toLower}}.{{.Name}}.example.com",
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "routegroup.my-rg.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.example.com"},
					Labels:     endpoint.Labels{endpoint.ResourceLabelKey: "routegroup/ns/my-rg"},
				},
			},
		},
		{
			title: "fqdn-template can reference .APIVersion",
			routeGroup: &rgv1.RouteGroup{
				APIVersion: "zalando.org/v1", Kind: "RouteGroup",
				Name: "my-rg", Namespace: "ns",
				Status: rgv1.RouteGroupStatus{
					LoadBalancer: rgv1.RouteGroupLoadBalancerStatus{RouteGroup: lb},
				},
			},
			fqdnTemplate: `{{.Name}}.{{replace "/" "." .APIVersion}}.example.com`,
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "my-rg.zalando.org.v1.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.example.com"},
					Labels:     endpoint.Labels{endpoint.ResourceLabelKey: "routegroup/ns/my-rg"},
				},
			},
		},
		{
			title:        "fqdn-template can reference .Namespace",
			routeGroup:   createTestRouteGroup("production", "my-rg", nil, nil, lb),
			fqdnTemplate: "{{.Name}}.{{.Namespace}}.example.com",
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "my-rg.production.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.example.com"},
					Labels:     endpoint.Labels{endpoint.ResourceLabelKey: "routegroup/production/my-rg"},
				},
			},
		},
		{
			title:        "backward-compat .Metadata.Name prefix in fqdn-template",
			routeGroup:   createTestRouteGroup("ns", "my-rg", nil, nil, lb),
			fqdnTemplate: "{{.Metadata.Name}}.example.com",
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "my-rg.example.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.example.com"},
					Labels:     endpoint.Labels{endpoint.ResourceLabelKey: "routegroup/ns/my-rg"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			src := newTestRouteGroupSource(t, &Config{
				TemplateEngine: templatetest.MustEngine(t, tt.fqdnTemplate, "", "", tt.combine),
			}, tt.routeGroup)

			got, err := src.Endpoints(t.Context())
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			testutils.ValidateEndpoints(t, got, tt.expected)
		})
	}
}
