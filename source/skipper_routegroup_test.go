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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	rgv1 "github.com/szuecs/routegroup-client/apis/zalando.org/v1"
	rgfake "github.com/szuecs/routegroup-client/client/clientset/versioned/fake"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
	"sigs.k8s.io/external-dns/source/types"
)

func createTestRouteGroup(ns, name string, anns map[string]string, hosts []string, destinations []rgv1.RouteGroupLoadBalancer) *rgv1.RouteGroup {
	return &rgv1.RouteGroup{
		Namespace:   ns,
		Name:        name,
		Annotations: anns,
		Spec: rgv1.RouteGroupSpec{
			Hosts: hosts,
		},
		Status: rgv1.RouteGroupStatus{
			LoadBalancer: rgv1.RouteGroupLoadBalancerStatus{
				RouteGroup: destinations,
			},
		},
	}
}

func newTestRouteGroupSource(t *testing.T, cfg *Config, rgs ...*rgv1.RouteGroup) Source {
	t.Helper()
	objects := make([]runtime.Object, len(rgs))
	for i, rg := range rgs {
		objects[i] = rg
	}
	fakeClient := rgfake.NewSimpleClientset(objects...)
	src, err := NewRouteGroupSource(t.Context(), fakeClient, cfg)
	require.NoError(t, err)
	return src
}

func TestNewRouteGroupSource(t *testing.T) {
	t.Parallel()

	t.Run("creates source successfully", func(t *testing.T) {
		t.Parallel()
		fakeClient := rgfake.NewSimpleClientset()
		src, err := NewRouteGroupSource(t.Context(), fakeClient, &Config{})
		require.NoError(t, err)
		_, ok := src.(*routeGroupSource)
		require.True(t, ok)
	})

	t.Run("respects namespace", func(t *testing.T) {
		t.Parallel()
		fakeClient := rgfake.NewSimpleClientset()
		src, err := NewRouteGroupSource(t.Context(), fakeClient, &Config{Namespace: "test-ns"})
		require.NoError(t, err)
		require.NotNil(t, src)
	})

	t.Run("logs warning for non-default version", func(t *testing.T) {
		t.Parallel()
		fakeClient := rgfake.NewSimpleClientset()
		src, err := NewRouteGroupSource(t.Context(), fakeClient, &Config{SkipperRouteGroupVersion: "zalando.org/v1alpha1"})
		require.NoError(t, err)
		require.NotNil(t, src)
	})
}

func TestEndpointsFromRouteGroups(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		source *routeGroupSource
		rg     *rgv1.RouteGroup
		want   []*endpoint.Endpoint
	}{
		{
			name:   "Empty routegroup should return empty endpoints",
			source: &routeGroupSource{},
			rg:     &rgv1.RouteGroup{},
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
			rg: createTestRouteGroup("namespace1", "rg1", nil, nil, []rgv1.RouteGroupLoadBalancer{
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
			rg: createTestRouteGroup("namespace1", "rg1", nil, []string{"rg1.k8s.example"}, []rgv1.RouteGroupLoadBalancer{
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
					annotations.HostnameKey: "my.example",
				},
				[]string{"rg1.k8s.example"},
				[]rgv1.RouteGroupLoadBalancer{
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
					annotations.HostnameKey: "my.example",
				},
				[]string{"rg1.k8s.example"},
				[]rgv1.RouteGroupLoadBalancer{
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
					annotations.TtlKey: "2189",
				},
				[]string{"rg1.k8s.example"},
				[]rgv1.RouteGroupLoadBalancer{
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
				[]rgv1.RouteGroupLoadBalancer{
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
			name:   "Routegroup with hosts and destination IPv6 creates an endpoint",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				nil,
				[]string{"rg1.k8s.example"},
				[]rgv1.RouteGroupLoadBalancer{
					{
						IP: "2001:DB8::1",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeAAAA,
					Targets:    endpoint.Targets([]string{"2001:DB8::1"}),
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
				[]rgv1.RouteGroupLoadBalancer{
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
		},
		{
			name:   "Routegroup with hosts and mixed destinations (IPv6) creates endpoints",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				nil,
				[]string{"rg1.k8s.example"},
				[]rgv1.RouteGroupLoadBalancer{
					{
						Hostname: "lb.example.org",
						IP:       "2001:DB8::1",
					},
				},
			),
			want: []*endpoint.Endpoint{
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeAAAA,
					Targets:    endpoint.Targets([]string{"2001:DB8::1"}),
				},
				{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name:   "Routegroup with provider-specific annotation creates endpoint with provider-specific property",
			source: &routeGroupSource{},
			rg: createTestRouteGroup(
				"namespace1",
				"rg1",
				map[string]string{
					annotations.AWSPrefix + "weight": "10",
				},
				[]string{"rg1.k8s.example"},
				[]rgv1.RouteGroupLoadBalancer{
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
					ProviderSpecific: endpoint.ProviderSpecific{
						{Name: "aws/weight", Value: "10"},
					},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.source.endpointsFromRouteGroup(tt.rg)

			testutils.ValidateEndpoints(t, got, tt.want)
		})
	}
}

func TestRouteGroupsEndpoints(t *testing.T) {
	for _, tt := range []struct {
		name        string
		rgs         []*rgv1.RouteGroup
		cfg         *Config
		templates   string
		combineFQDN bool
		want        []*endpoint.Endpoint
		wantErr     bool
	}{
		{
			name: "Empty routegroup should return empty endpoints",
			rgs:  nil,
			want: []*endpoint.Endpoint{},
		},
		{
			name: "Single routegroup should return endpoints",
			rgs: []*rgv1.RouteGroup{
				{
					Namespace: "namespace1",
					Name:      "rg1",
					UID:       "skipper-rg-uid-1234",
					Spec: rgv1.RouteGroupSpec{
						Hosts: []string{"rg1.k8s.example"},
					},
					Status: rgv1.RouteGroupStatus{
						LoadBalancer: rgv1.RouteGroupLoadBalancerStatus{
							RouteGroup: []rgv1.RouteGroupLoadBalancer{
								{
									Hostname: "lb.example.org",
								},
							},
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				(&endpoint.Endpoint{
					DNSName:    "rg1.k8s.example",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				}).WithRefObject(testutils.RefSource(string(types.SkipperRouteGroup))),
			},
		},
		{
			name:        "Single routegroup with combineFQDNAnnotation with fqdn template should return endpoints from fqdnTemplate and routegroup",
			templates:   "{{.Name}}.{{.Namespace}}.example",
			combineFQDN: true,
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					nil,
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
						},
					},
				),
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
			name:        "Single routegroup with combineFQDNAnnotation with fqdn template prefixed with .Metadata should return endpoints from fqdnTemplate and routegroup",
			templates:   "{{.Metadata.Name}}.{{.Metadata.Namespace}}.example",
			combineFQDN: true,
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					nil,
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
						},
					},
				),
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
			name:      "Single routegroup without hosts, with fqdn template should return endpoints from fqdnTemplate",
			templates: "{{.Name}}.{{.Namespace}}.example",
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					nil,
					nil,
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
						},
					},
				),
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
			name:      "Single routegroup without hosts, with fqdn template using .Metadata should return endpoints from fqdnTemplate",
			templates: "{{.Metadata.Name}}.{{.Metadata.Namespace}}.example",
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					nil,
					nil,
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
						},
					},
				),
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
			name:      "fqdn template execution error should be returned",
			templates: "{{index . 0}}",
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					nil,
					nil,
					[]rgv1.RouteGroupLoadBalancer{{Hostname: "lb.example.org"}},
				),
			},
			wantErr: true,
		},
		{
			name:      "Single routegroup without combineFQDNAnnotation with fqdn template should return endpoints not from fqdnTemplate",
			templates: "{{.Name}}.{{.Namespace}}.example",
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					nil,
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
						},
					},
				),
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
			name:      "Single routegroup without combineFQDNAnnotation with fqdn template using .Metadata should return endpoints not from fqdnTemplate",
			templates: "{{.Metadata.Name}}.{{.Metadata.Namespace}}.example",
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					nil,
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
						},
					},
				),
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
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					map[string]string{
						annotations.TtlKey: "2189",
					},
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
						},
					},
				),
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
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					nil,
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
							IP:       "1.5.1.4",
						},
					},
				),
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
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					nil,
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb2.example.org",
						},
					},
				),
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
			cfg:  &Config{AnnotationFilter: parseAnnotationFilterOrNil("kubernetes.io/ingress.class=skipper")},
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					map[string]string{
						"kubernetes.io/ingress.class": "skipper",
					},
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb2.example.org",
						},
					},
				),
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
			cfg:  &Config{AnnotationFilter: parseAnnotationFilterOrNil("kubernetes.io/ingress.class in (nginx, skipper)")},
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					map[string]string{
						"kubernetes.io/ingress.class": "skipper",
					},
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb2.example.org",
						},
					},
				),
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
			name: "multiple routegroups with matching label filter returns only labeled endpoints",
			cfg:  &Config{LabelFilter: labels.SelectorFromSet(labels.Set{"app": "test"})},
			rgs: []*rgv1.RouteGroup{
				{
					Namespace: "namespace1",
					Name:      "rg-match",
					Labels:    map[string]string{"app": "test"},
					Spec:      rgv1.RouteGroupSpec{Hosts: []string{"match.example.org"}},
					Status: rgv1.RouteGroupStatus{
						LoadBalancer: rgv1.RouteGroupLoadBalancerStatus{
							RouteGroup: []rgv1.RouteGroupLoadBalancer{{Hostname: "lb.example.org"}},
						},
					},
				},
				{
					Namespace: "namespace1",
					Name:      "rg-no-match",
					Labels:    map[string]string{"app": "other"},
					Spec:      rgv1.RouteGroupSpec{Hosts: []string{"no-match.example.org"}},
					Status: rgv1.RouteGroupStatus{
						LoadBalancer: rgv1.RouteGroupLoadBalancerStatus{
							RouteGroup: []rgv1.RouteGroupLoadBalancer{{Hostname: "lb.example.org"}},
						},
					},
				},
			},
			want: []*endpoint.Endpoint{
				{
					DNSName:    "match.example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets([]string{"lb.example.org"}),
				},
			},
		},
		{
			name: "multiple routegroups with non-matching label filter returns no endpoints",
			cfg:  &Config{LabelFilter: labels.SelectorFromSet(labels.Set{"app": "test"})},
			rgs: []*rgv1.RouteGroup{
				{
					Namespace: "namespace1",
					Name:      "rg-no-match",
					Labels:    map[string]string{"app": "other"},
					Spec:      rgv1.RouteGroupSpec{Hosts: []string{"no-match.example.org"}},
					Status: rgv1.RouteGroupStatus{
						LoadBalancer: rgv1.RouteGroupLoadBalancerStatus{
							RouteGroup: []rgv1.RouteGroupLoadBalancer{{Hostname: "lb.example.org"}},
						},
					},
				},
			},
			want: []*endpoint.Endpoint{},
		},
		{
			name: "multiple routegroups with controller annotation filter should not return filtered endpoints",
			rgs: []*rgv1.RouteGroup{
				createTestRouteGroup(
					"namespace1",
					"rg1",
					map[string]string{
						annotations.ControllerKey: annotations.ControllerValue,
					},
					[]string{"rg1.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
						},
					},
				),
				createTestRouteGroup(
					"namespace1",
					"rg2",
					map[string]string{
						annotations.ControllerKey: "dns",
					},
					[]string{"rg2.k8s.example"},
					[]rgv1.RouteGroupLoadBalancer{
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
					[]rgv1.RouteGroupLoadBalancer{
						{
							Hostname: "lb.example.org",
						},
					},
				),
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
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.cfg
			if cfg == nil {
				cfg = &Config{}
			}
			if tt.templates != "" {
				cfg.TemplateEngine = templatetest.MustEngine(t, tt.templates, "", "", tt.combineFQDN)
			}

			src := newTestRouteGroupSource(t, cfg, tt.rgs...)

			got, err := src.Endpoints(t.Context())
			if err != nil && !tt.wantErr {
				t.Errorf("Got error, but does not want to get an error: %v", err)
			}
			if tt.wantErr && err == nil {
				t.Fatal("Got no error, but we want to get an error")
			}

			testutils.ValidateEndpoints(t, got, tt.want)
		})
	}
}

func TestResourceLabelIsSet(t *testing.T) {
	src := newTestRouteGroupSource(t, &Config{},
		createTestRouteGroup(
			"namespace1",
			"rg1",
			nil,
			[]string{"rg1.k8s.example"},
			[]rgv1.RouteGroupLoadBalancer{
				{
					Hostname: "lb.example.org",
				},
			},
		),
	)

	got, _ := src.Endpoints(t.Context())
	for _, ep := range got {
		if _, ok := ep.Labels[endpoint.ResourceLabelKey]; !ok {
			t.Errorf("Failed to set resource label on ep %v", ep)
		}
	}
}

func TestRouteGroupAddEventHandler(t *testing.T) {
	t.Parallel()

	src := newTestRouteGroupSource(t, &Config{})
	called := false
	src.(*routeGroupSource).AddEventHandler(t.Context(), func() { called = true })
	assert.False(t, called, "handler should not be called immediately")
}
