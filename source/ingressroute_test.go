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

	contour "github.com/heptio/contour/apis/contour/v1beta1"
	fakeContour "github.com/heptio/contour/apis/generated/clientset/versioned/fake"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fakeKube "k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
)

// This is a compile-time validation that ingressRouteSource is a Source.
var _ Source = &ingressRouteSource{}

type IngressRouteSuite struct {
	suite.Suite
	source       Source
	loadBalancer *v1.Service
	ingressRoute *contour.IngressRoute
}

func (suite *IngressRouteSuite) SetupTest() {
	fakeKubernetesClient := fakeKube.NewSimpleClientset()
	fakeContourClient := fakeContour.NewSimpleClientset()
	var err error

	suite.loadBalancer = (fakeLoadBalancerService{
		ips:       []string{"8.8.8.8"},
		hostnames: []string{"v1"},
		namespace: "heptio-contour/contour",
		name:      "contour",
	}).Service()

	_, err = fakeKubernetesClient.CoreV1().Services(suite.loadBalancer.Namespace).Create(suite.loadBalancer)
	suite.NoError(err, "should succeed")

	suite.source, err = NewContourIngressRouteSource(
		fakeKubernetesClient,
		fakeContourClient,
		"heptio-contour/contour",
		"default",
		"",
		"{{.Name}}",
		false,
		false,
	)
	suite.NoError(err, "should initialize ingressroute source")

	suite.ingressRoute = (fakeIngressRoute{
		name:      "foo-ingressroute-with-targets",
		namespace: "default",
		host:      "example.com",
	}).IngressRoute()
	_, err = fakeContourClient.ContourV1beta1().IngressRoutes(suite.ingressRoute.Namespace).Create(suite.ingressRoute)
	suite.NoError(err, "should succeed")
}

func (suite *IngressRouteSuite) TestResourceLabelIsSet() {
	endpoints, _ := suite.source.Endpoints()
	for _, ep := range endpoints {
		suite.Equal("ingressroute/default/foo-ingressroute-with-targets", ep.Labels[endpoint.ResourceLabelKey], "should set correct resource label")
	}
}

func TestIngressRoute(t *testing.T) {
	suite.Run(t, new(IngressRouteSuite))
	t.Run("endpointsFromIngressRoute", testEndpointsFromIngressRoute)
	t.Run("Endpoints", testIngressRouteEndpoints)
}

func TestNewContourIngressRouteSource(t *testing.T) {
	for _, ti := range []struct {
		title                    string
		annotationFilter         string
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		expectError              bool
	}{
		{
			title:        "invalid template",
			expectError:  true,
			fqdnTemplate: "{{.Name",
		},
		{
			title:       "valid empty template",
			expectError: false,
		},
		{
			title:        "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com",
		},
		{
			title:        "valid template",
			expectError:  false,
			fqdnTemplate: "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
		},
		{
			title:                    "valid template",
			expectError:              false,
			fqdnTemplate:             "{{.Name}}-{{.Namespace}}.ext-dns.test.com, {{.Name}}-{{.Namespace}}.ext-dna.test.com",
			combineFQDNAndAnnotation: true,
		},
		{
			title:            "non-empty annotation filter label",
			expectError:      false,
			annotationFilter: "contour.heptio.com/ingress.class=contour",
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			_, err := NewContourIngressRouteSource(
				fakeKube.NewSimpleClientset(),
				fakeContour.NewSimpleClientset(),
				"heptio-contour/contour",
				"",
				ti.annotationFilter,
				ti.fqdnTemplate,
				ti.combineFQDNAndAnnotation,
				false,
			)
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func testEndpointsFromIngressRoute(t *testing.T) {
	for _, ti := range []struct {
		title        string
		loadBalancer fakeLoadBalancerService
		ingressRoute fakeIngressRoute
		expected     []*endpoint.Endpoint
	}{
		{
			title: "one rule.host one lb.hostname",
			loadBalancer: fakeLoadBalancerService{
				hostnames: []string{"lb.com"}, // Kubernetes omits the trailing dot
			},
			ingressRoute: fakeIngressRoute{
				host: "foo.bar", // Kubernetes requires removal of trailing dot
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo.bar",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title: "one rule.host one lb.IP",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRoute: fakeIngressRoute{
				host: "foo.bar",
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo.bar",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
			},
		},
		{
			title: "one rule.host two lb.IP and two lb.Hostname",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8", "127.0.0.1"},
				hostnames: []string{"elb.com", "alb.com"},
			},
			ingressRoute: fakeIngressRoute{
				host: "foo.bar",
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo.bar",
					Targets: endpoint.Targets{"8.8.8.8", "127.0.0.1"},
				},
				{
					DNSName: "foo.bar",
					Targets: endpoint.Targets{"elb.com", "alb.com"},
				},
			},
		},
		{
			title: "no rule.host",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8", "127.0.0.1"},
				hostnames: []string{"elb.com", "alb.com"},
			},
			ingressRoute: fakeIngressRoute{},
			expected:     []*endpoint.Endpoint{},
		},
		{
			title: "one rule.host invalid ingressroute",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8", "127.0.0.1"},
				hostnames: []string{"elb.com", "alb.com"},
			},
			ingressRoute: fakeIngressRoute{
				host:    "foo.bar",
				invalid: true,
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:        "no targets",
			loadBalancer: fakeLoadBalancerService{},
			ingressRoute: fakeIngressRoute{},
			expected:     []*endpoint.Endpoint{},
		},
		{
			title: "delegate ingressroute",
			loadBalancer: fakeLoadBalancerService{
				hostnames: []string{"lb.com"},
			},
			ingressRoute: fakeIngressRoute{
				delegate: true,
			},
			expected: []*endpoint.Endpoint{},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			if source, err := newTestIngressRouteSource(ti.loadBalancer); err != nil {
				require.NoError(t, err)
			} else if endpoints, err := source.endpointsFromIngressRoute(ti.ingressRoute.IngressRoute()); err != nil {
				require.NoError(t, err)
			} else {
				validateEndpoints(t, endpoints, ti.expected)
			}
		})
	}
}

func testIngressRouteEndpoints(t *testing.T) {
	namespace := "testing"
	for _, ti := range []struct {
		title                    string
		targetNamespace          string
		annotationFilter         string
		loadBalancer             fakeLoadBalancerService
		ingressRouteItems        []fakeIngressRoute
		expected                 []*endpoint.Endpoint
		expectError              bool
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		ignoreHostnameAnnotation bool
	}{
		{
			title:           "no ingressroute",
			targetNamespace: "",
		},
		{
			title:           "two simple ingressroutes",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					host:      "example.org",
				},
				{
					name:      "fake2",
					namespace: namespace,
					host:      "new.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"lb.com"},
				},
				{
					DNSName: "new.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "new.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:           "two simple ingressroutes on different namespaces",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: "testing1",
					host:      "example.org",
				},
				{
					name:      "fake2",
					namespace: "testing2",
					host:      "new.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"lb.com"},
				},
				{
					DNSName: "new.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "new.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:           "two simple ingressroutes on different namespaces and a target namespace",
			targetNamespace: "testing1",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: "testing1",
					host:      "example.org",
				},
				{
					name:      "fake2",
					namespace: "testing2",
					host:      "new.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:            "valid matching annotation filter expression",
			targetNamespace:  "",
			annotationFilter: "contour.heptio.com/ingress.class in (alb, contour)",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"contour.heptio.com/ingress.class": "contour",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
			},
		},
		{
			title:            "valid non-matching annotation filter expression",
			targetNamespace:  "",
			annotationFilter: "contour.heptio.com/ingress.class in (alb, contour)",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"contour.heptio.com/ingress.class": "tectonic",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:            "invalid annotation filter expression",
			targetNamespace:  "",
			annotationFilter: "contour.heptio.com/ingress.name in (a b)",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"contour.heptio.com/ingress.class": "alb",
					},
					host: "example.org",
				},
			},
			expected:    []*endpoint.Endpoint{},
			expectError: true,
		},
		{
			title:            "valid matching annotation filter label",
			targetNamespace:  "",
			annotationFilter: "contour.heptio.com/ingress.class=contour",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"contour.heptio.com/ingress.class": "contour",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
			},
		},
		{
			title:            "valid non-matching annotation filter label",
			targetNamespace:  "",
			annotationFilter: "contour.heptio.com/ingress.class=contour",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"contour.heptio.com/ingress.class": "alb",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "our controller type is dns-controller",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: controllerAnnotationValue,
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
			},
		},
		{
			title:           "different controller types are ignored",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: "some-other-tool",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "template for ingressroute if host is missing",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"elb.com"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: controllerAnnotationValue,
					},
					host: "",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "fake1.ext-dns.test.com",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "fake1.ext-dns.test.com",
					Targets: endpoint.Targets{"elb.com"},
				},
			},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title:           "another controller annotation skipped even with template",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: "other-controller",
					},
					host: "",
				},
			},
			expected:     []*endpoint.Endpoint{},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title:           "multiple FQDN template hostnames",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:        "fake1",
					namespace:   namespace,
					annotations: map[string]string{},
					host:        "",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "fake1.ext-dns.test.com",
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordType: endpoint.RecordTypeA,
				},
				{
					DNSName:    "fake1.ext-dna.test.com",
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordType: endpoint.RecordTypeA,
				},
			},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com, {{.Name}}.ext-dna.test.com",
		},
		{
			title:           "multiple FQDN template hostnames",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:        "fake1",
					namespace:   namespace,
					annotations: map[string]string{},
					host:        "",
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "ingressroute-target.com",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "fake1.ext-dns.test.com",
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordType: endpoint.RecordTypeA,
				},
				{
					DNSName:    "fake1.ext-dna.test.com",
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordType: endpoint.RecordTypeA,
				},
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"ingressroute-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dns.test.com",
					Targets:    endpoint.Targets{"ingressroute-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dna.test.com",
					Targets:    endpoint.Targets{"ingressroute-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
			fqdnTemplate:             "{{.Name}}.ext-dns.test.com, {{.Name}}.ext-dna.test.com",
			combineFQDNAndAnnotation: true,
		},
		{
			title:           "ingressroute rules with annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "ingressroute-target.com",
					},
					host: "example.org",
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "ingressroute-target.com",
					},
					host: "example2.org",
				},
				{
					name:      "fake3",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "1.2.3.4",
					},
					host: "example3.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"ingressroute-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "example2.org",
					Targets:    endpoint.Targets{"ingressroute-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "example3.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
				},
			},
		},
		{
			title:           "ingressroute rules with hostname annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"1.2.3.4"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
				},
				{
					DNSName:    "dns-through-hostname.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
				},
			},
		},
		{
			title:           "ingressroute rules with hostname annotation having multiple hostnames",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"1.2.3.4"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com, another-dns-through-hostname.com",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
				},
				{
					DNSName:    "dns-through-hostname.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
				},
				{
					DNSName:    "another-dns-through-hostname.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
				},
			},
		},
		{
			title:           "ingressroute rules with hostname and target annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com",
						targetAnnotationKey:   "ingressroute-target.com",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"ingressroute-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "dns-through-hostname.com",
					Targets:    endpoint.Targets{"ingressroute-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
		},
		{
			title:           "ingressroute rules with annotation and custom TTL",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "ingressroute-target.com",
						ttlAnnotationKey:    "6",
					},
					host: "example.org",
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "ingressroute-target.com",
						ttlAnnotationKey:    "1",
					},
					host: "example2.org",
				},
				{
					name:      "fake3",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "ingressroute-target.com",
						ttlAnnotationKey:    "10s",
					},
					host: "example3.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:   "example.org",
					Targets:   endpoint.Targets{"ingressroute-target.com"},
					RecordTTL: endpoint.TTL(6),
				},
				{
					DNSName:   "example2.org",
					Targets:   endpoint.Targets{"ingressroute-target.com"},
					RecordTTL: endpoint.TTL(1),
				},
				{
					DNSName:   "example3.org",
					Targets:   endpoint.Targets{"ingressroute-target.com"},
					RecordTTL: endpoint.TTL(10),
				},
			},
		},
		{
			title:           "template for ingressroute with annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{},
				hostnames: []string{},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "ingressroute-target.com",
					},
					host: "",
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "ingressroute-target.com",
					},
					host: "",
				},
				{
					name:      "fake3",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "1.2.3.4",
					},
					host: "",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "fake1.ext-dns.test.com",
					Targets:    endpoint.Targets{"ingressroute-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dns.test.com",
					Targets:    endpoint.Targets{"ingressroute-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake3.ext-dns.test.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
				},
			},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title:           "ingressroute with empty annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{},
				hostnames: []string{},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "",
					},
					host: "",
				},
			},
			expected:     []*endpoint.Endpoint{},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title:           "ignore hostname annotations",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
			},
			ingressRouteItems: []fakeIngressRoute{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "ignore.me",
					},
					host: "example.org",
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "ignore.me.too",
					},
					host: "new.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "example.org",
					Targets: endpoint.Targets{"lb.com"},
				},
				{
					DNSName: "new.org",
					Targets: endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName: "new.org",
					Targets: endpoint.Targets{"lb.com"},
				},
			},
			ignoreHostnameAnnotation: true,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			ingressRoutes := make([]*contour.IngressRoute, 0)
			for _, item := range ti.ingressRouteItems {
				ingressRoutes = append(ingressRoutes, item.IngressRoute())
			}

			fakeKubernetesClient := fakeKube.NewSimpleClientset()

			lbService := ti.loadBalancer.Service()
			_, err := fakeKubernetesClient.CoreV1().Services(lbService.Namespace).Create(lbService)
			if err != nil {
				require.NoError(t, err)
			}

			fakeContourClient := fakeContour.NewSimpleClientset()
			for _, ingressRoute := range ingressRoutes {
				_, err := fakeContourClient.ContourV1beta1().IngressRoutes(ingressRoute.Namespace).Create(ingressRoute)
				require.NoError(t, err)
			}

			ingressRouteSource, err := NewContourIngressRouteSource(
				fakeKubernetesClient,
				fakeContourClient,
				lbService.Namespace+"/"+lbService.Name,
				ti.targetNamespace,
				ti.annotationFilter,
				ti.fqdnTemplate,
				ti.combineFQDNAndAnnotation,
				ti.ignoreHostnameAnnotation,
			)
			require.NoError(t, err)

			res, err := ingressRouteSource.Endpoints()
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			validateEndpoints(t, res, ti.expected)
		})
	}
}

// ingressroute specific helper functions
func newTestIngressRouteSource(loadBalancer fakeLoadBalancerService) (*ingressRouteSource, error) {
	fakeKubernetesClient := fakeKube.NewSimpleClientset()
	fakeContourClient := fakeContour.NewSimpleClientset()

	lbService := loadBalancer.Service()
	_, err := fakeKubernetesClient.CoreV1().Services(lbService.Namespace).Create(lbService)
	if err != nil {
		return nil, err
	}

	src, err := NewContourIngressRouteSource(
		fakeKubernetesClient,
		fakeContourClient,
		lbService.Namespace+"/"+lbService.Name,
		"default",
		"",
		"{{.Name}}",
		false,
		false,
	)
	if err != nil {
		return nil, err
	}

	irsrc, ok := src.(*ingressRouteSource)
	if !ok {
		return nil, errors.New("underlying source type was not ingressroute")
	}

	return irsrc, nil
}

type fakeLoadBalancerService struct {
	ips       []string
	hostnames []string
	namespace string
	name      string
}

func (ig fakeLoadBalancerService) Service() *v1.Service {
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: ig.namespace,
			Name:      ig.name,
		},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{
				Ingress: []v1.LoadBalancerIngress{},
			},
		},
	}

	for _, ip := range ig.ips {
		svc.Status.LoadBalancer.Ingress = append(svc.Status.LoadBalancer.Ingress, v1.LoadBalancerIngress{
			IP: ip,
		})
	}
	for _, hostname := range ig.hostnames {
		svc.Status.LoadBalancer.Ingress = append(svc.Status.LoadBalancer.Ingress, v1.LoadBalancerIngress{
			Hostname: hostname,
		})
	}

	return svc
}

type fakeIngressRoute struct {
	namespace   string
	name        string
	annotations map[string]string

	host     string
	invalid  bool
	delegate bool
}

func (ir fakeIngressRoute) IngressRoute() *contour.IngressRoute {
	var status string
	if ir.invalid {
		status = "invalid"
	} else {
		status = "valid"
	}

	var spec contour.IngressRouteSpec
	if ir.delegate {
		spec = contour.IngressRouteSpec{}
	} else {
		spec = contour.IngressRouteSpec{
			VirtualHost: &contour.VirtualHost{
				Fqdn: ir.host,
			},
		}
	}

	ingressRoute := &contour.IngressRoute{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   ir.namespace,
			Name:        ir.name,
			Annotations: ir.annotations,
		},
		Spec: spec,
		Status: contour.Status{
			CurrentStatus: status,
		},
	}

	return ingressRoute
}
