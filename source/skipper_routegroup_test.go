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
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
)

func createTestRouteGroup(ns, name string, annotations map[string]string, hosts []string, destinations []routeGroupLoadBalancer) *routeGroup {
	return &routeGroup{
		Metadata: itemMetadata{
			Namespace:   ns,
			Name:        name,
			Annotations: annotations,
		},
		Spec: routeGroupSpec{
			Hosts: hosts,
		},
		Status: routeGroupStatus{
			LoadBalancer: routeGroupLoadBalancerStatus{
				RouteGroup: destinations,
			},
		},
	}
}

func TestEndpointsFromRouteGroups(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		source *routeGroupSource
		rg     *routeGroup
		want   []*endpoint.Endpoint
	}{
		{
			name:   "Empty routegroup should return empty endpoints",
			source: &routeGroupSource{},
			rg:     &routeGroup{},
			want:   []*endpoint.Endpoint{},
		},
		{
			name:   "Routegroup without hosts and destinations create no endpoints",
			source: &routeGroupSource{},
			rg:     createTestRouteGroup("namespace1", "rg1", nil, nil, nil),
			want:   []*endpoint.Endpoint{},
		},
		{
			name:   "Routegroup without hosts create no endpoints",
			source: &routeGroupSource{},
			rg: createTestRouteGroup("namespace1", "rg1", nil, nil, []routeGroupLoadBalancer{
				{
					Hostname: "lb.example.org",
				},
			}),
			want: []*endpoint.Endpoint{},
		},
		{
			name:   "Routegroup without destinations create no endpoints",
			source: &routeGroupSource{},
			rg:     createTestRouteGroup("namespace1", "rg1", nil, []string{"rg1.k8s.example"}, nil),
			want:   []*endpoint.Endpoint{},
		},
		{
			name:   "Routegroup with hosts and destinations creates an endpoint",
			source: &routeGroupSource{},
			rg: createTestRouteGroup("namespace1", "rg1", nil, []string{"rg1.k8s.example"}, []routeGroupLoadBalancer{
				{
					Hostname: "lb.example.org",
				},
			}),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:   "Routegroup with hostname annotation, creates endpoints from the annotation ",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				map[string]string{
					hostnameAnnotationKey: "my.example",
				},
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "my.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:   "Routegroup with hosts and destinations and ignoreHostnameAnnotation creates endpoints but ignores annotation",
			source: &routeGroupSource{ignoreHostnameAnnotation: true},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				map[string]string{
					hostnameAnnotationKey: "my.example",
				},
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:   "Routegroup with hosts and destinations and ttl creates an endpoint with ttl",
			source: &routeGroupSource{ignoreHostnameAnnotation: true},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				map[string]string{
					ttlAnnotationKey: "2189",
				},
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
					RecordTTL:  endpoint.TTL(2189),
				},
			},
		},
		{
			name:   "Routegroup with hosts and destination IP creates an endpoint",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				nil,
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						IP: "1.5.1.4",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets([]string{"1.5.1.4"}),
				},
			},
		},
		{
			name:   "Routegroup with hosts and mixed destinations creates endpoints",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				nil,
				[]string{"rg1.k8s.example"},
				[]routeGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
						IP:       "1.5.1.4",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets([]string{"1.5.1.4"}),
				},
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		}} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.source.endpointsFromRouteGroup(tt.rg)

			validateEndpoints(t, got, tt.want)
		})
	}

}

type fakeRouteGroupClient struct {
	returnErr bool
	rg        *routeGroupList
}

func (f *fakeRouteGroupClient) getRouteGroupList(string) (*routeGroupList, error) {
	if f.returnErr {
		return nil, errors.New("Fake route group list error")
	}
	return f.rg, nil
}

func TestRouteGroupsEndpoints(t *testing.T) {
	for _, tt := range []struct {
		name         string
		source       *routeGroupSource
		fqdnTemplate string
		want         []*endpoint.Endpoint
		wantErr      bool
	}{
		{
			name: "Empty routegroup should return empty endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{},
				},
			},
			want:    []*endpoint.Endpoint{},
			wantErr: false,
		},
		{
			name: "Single routegroup should return endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:         "Single routegroup with combineFQDNAnnotation with fqdn template should return endpoints from fqdnTemplate and routegroup",
			fqdnTemplate: "{{.Metadata.Name}}.{{.Metadata.Namespace}}.example",
			source: &routeGroupSource{
				combineFQDNAnnotation: true,
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg1.namespace1.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:         "Single routegroup without, with fqdn template should return endpoints from fqdnTemplate",
			fqdnTemplate: "{{.Metadata.Name}}.{{.Metadata.Namespace}}.example",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								nil,
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.namespace1.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:         "Single routegroup without combineFQDNAnnotation with fqdn template should return endpoints not from fqdnTemplate",
			fqdnTemplate: "{{.Metadata.Name}}.{{.Metadata.Namespace}}.example",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "Single routegroup with TTL should return endpoint with TTL",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								map[string]string{
									ttlAnnotationKey: "2189",
								},
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
					RecordTTL:  endpoint.TTL(2189),
				},
			},
		},
		{
			name: "Routegroup with hosts and mixed destinations creates endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
										IP:       "1.5.1.4",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets([]string{"1.5.1.4"}),
				},
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups should return endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								nil,
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace1",
								"rg2",
								nil,
								[]string{"rg2.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace2",
								"rg3",
								nil,
								[]string{"rg3.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace3",
								"rg",
								nil,
								[]string{"rg.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb2.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg2.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg3.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb2.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups with filter annotations should return only filtered endpoints",
			source: &routeGroupSource{
				annotationFilter: "kubernetes.io/ingress.class=skipper",
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								map[string]string{
									"kubernetes.io/ingress.class": "skipper",
								},
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace1",
								"rg2",
								map[string]string{
									"kubernetes.io/ingress.class": "nginx",
								},
								[]string{"rg2.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace2",
								"rg3",
								map[string]string{
									"kubernetes.io/ingress.class": "",
								},
								[]string{"rg3.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace3",
								"rg",
								nil,
								[]string{"rg.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb2.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups with set operation annotation filter should return only filtered endpoints",
			source: &routeGroupSource{
				annotationFilter: "kubernetes.io/ingress.class in (nginx, skipper)",
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								map[string]string{
									"kubernetes.io/ingress.class": "skipper",
								},
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace1",
								"rg2",
								map[string]string{
									"kubernetes.io/ingress.class": "nginx",
								},
								[]string{"rg2.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace2",
								"rg3",
								map[string]string{
									"kubernetes.io/ingress.class": "",
								},
								[]string{"rg3.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace3",
								"rg",
								nil,
								[]string{"rg.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb2.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg2.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups with controller annotation filter should not return filtered endpoints",
			source: &routeGroupSource{
				cli: &fakeRouteGroupClient{
					rg: &routeGroupList{
						Items: []*routeGroup{
							createTestRouteGroup(
								"namespace1",
								"rg1",
								map[string]string{
									controllerAnnotationKey: controllerAnnotationValue,
								},
								[]string{"rg1.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace1",
								"rg2",
								map[string]string{
									controllerAnnotationKey: "dns",
								},
								[]string{"rg2.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
							createTestRouteGroup(
								"namespace2",
								"rg3",
								nil,
								[]string{"rg3.k8s.example"},
								[]routeGroupLoadBalancer{
									{
										Hostname: "lb.example.org",
									},
								},
							),
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
				{
					DNSName:    "rg3.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		}} {
		t.Run(tt.name, func(t *testing.T) {
			if tt.fqdnTemplate != "" {
				tmpl, err := parseTemplate(tt.fqdnTemplate)
				if err != nil {
					t.Fatalf("Failed to parse template: %v", err)
				}
				tt.source.fqdnTemplate = tmpl
			}

			got, err := tt.source.Endpoints(context.Background())
			if err != nil && !tt.wantErr {
				t.Errorf("Got error, but does not want to get an error: %v", err)
			}
			if tt.wantErr && err == nil {
				t.Fatal("Got no error, but we want to get an error")
			}

			validateEndpoints(t, got, tt.want)
		})
	}
}

func TestResourceLabelIsSet(t *testing.T) {
	source := &routeGroupSource{
		cli: &fakeRouteGroupClient{
			rg: &routeGroupList{
				Items: []*routeGroup{
					createTestRouteGroup(
						"namespace1",
						"rg1",
						nil,
						[]string{"rg1.k8s.example"},
						[]routeGroupLoadBalancer{
							{
								Hostname: "lb.example.org",
							},
						},
					),
				},
			},
		},
	}

	got, _ := source.Endpoints(context.Background())
	for _, ep := range got {
		if _, ok := ep.Labels[endpoint.ResourceLabelKey]; !ok {
			t.Errorf("Failed to set resource label on ep %v", ep)
		}
	}
}

func TestDualstackLabelIsSet(t *testing.T) {
	source := &routeGroupSource{
		cli: &fakeRouteGroupClient{
			rg: &routeGroupList{
				Items: []*routeGroup{
					createTestRouteGroup(
						"namespace1",
						"rg1",
						map[string]string{
							ALBDualstackAnnotationKey: ALBDualstackAnnotationValue,
						},
						[]string{"rg1.k8s.example"},
						[]routeGroupLoadBalancer{
							{
								Hostname: "lb.example.org",
							},
						},
					),
				},
			},
		},
	}

	got, _ := source.Endpoints(context.Background())
	for _, ep := range got {
		if v, ok := ep.Labels[endpoint.DualstackLabelKey]; !ok || v != "true" {
			t.Errorf("Failed to set resource label on ep %v", ep)
		}
	}
}

func TestParseTemplate(t *testing.T) {
	for _, tt := range []struct {
		name                     string
		annotationFilter         string
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		expectError              bool
	}{
		{
			name:         "invalid template",
			expectError:  true,
			fqdnTemplate: "{{.Name",
		},
		{
			name:        "valid empty template",
			expectError: false,
		},
		{
			name:         "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com",
		},
		{
			name:         "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
		},
		{
			name:                     "valid template",
			expectError:              false,
			fqdnTemplate:             "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
			combineFQDNAndAnnotation: true,
		},
		{
			name:             "non-empty annotation filter label",
			expectError:      false,
			annotationFilter: "kubernetes.io/ingress.class=nginx",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseTemplate(tt.fqdnTemplate)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
