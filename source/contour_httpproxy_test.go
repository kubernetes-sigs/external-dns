/*
Copyright 2020 The Kubernetes Authors.

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

	fakeDynamic "k8s.io/client-go/dynamic/fake"

	"github.com/pkg/errors"
	projectcontour "github.com/projectcontour/contour/apis/projectcontour/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/external-dns/endpoint"
)

// This is a compile-time validation that httpProxySource is a Source.
var _ Source = &httpProxySource{}

type HTTPProxySuite struct {
	suite.Suite
	source    Source
	httpProxy *projectcontour.HTTPProxy
}

func newDynamicKubernetesClient() (*fakeDynamic.FakeDynamicClient, *runtime.Scheme) {
	s := runtime.NewScheme()
	_ = projectcontour.AddToScheme(s)
	return fakeDynamic.NewSimpleDynamicClient(s), s
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

func (suite *HTTPProxySuite) SetupTest() {
	fakeDynamicClient, s := newDynamicKubernetesClient()
	var err error

	suite.source, err = NewContourHTTPProxySource(
		context.TODO(),
		fakeDynamicClient,
		"default",
		"",
		"{{.Name}}",
		false,
		false,
	)
	suite.NoError(err, "should initialize httpproxy source")

	suite.httpProxy = (fakeHTTPProxy{
		name:      "foo-httpproxy-with-targets",
		namespace: "default",
		host:      "example.com",
	}).HTTPProxy()

	// Convert to unstructured
	unstructuredHTTPProxy, err := convertHTTPProxyToUnstructured(suite.httpProxy, s)
	if err != nil {
		suite.Error(err)
	}

	_, err = fakeDynamicClient.Resource(projectcontour.HTTPProxyGVR).Namespace(suite.httpProxy.Namespace).Create(context.Background(), unstructuredHTTPProxy, metav1.CreateOptions{})
	suite.NoError(err, "should succeed")
}

func (suite *HTTPProxySuite) TestResourceLabelIsSet() {
	endpoints, _ := suite.source.Endpoints(context.Background())
	for _, ep := range endpoints {
		suite.Equal("httpproxy/default/foo-httpproxy-with-targets", ep.Labels[endpoint.ResourceLabelKey], "should set correct resource label")
	}
}

func convertHTTPProxyToUnstructured(hp *projectcontour.HTTPProxy, s *runtime.Scheme) (*unstructured.Unstructured, error) {
	unstructuredHTTPProxy := &unstructured.Unstructured{}
	if err := s.Convert(hp, unstructuredHTTPProxy, context.Background()); err != nil {
		return nil, err
	}
	return unstructuredHTTPProxy, nil
}

func TestHTTPProxy(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(HTTPProxySuite))
	t.Run("endpointsFromHTTPProxy", testEndpointsFromHTTPProxy)
	t.Run("Endpoints", testHTTPProxyEndpoints)
}

func TestNewContourHTTPProxySource(t *testing.T) {
	t.Parallel()

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
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			fakeDynamicClient, _ := newDynamicKubernetesClient()

			_, err := NewContourHTTPProxySource(
				context.TODO(),
				fakeDynamicClient,
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

func testEndpointsFromHTTPProxy(t *testing.T) {
	t.Parallel()

	for _, ti := range []struct {
		title     string
		httpProxy fakeHTTPProxy
		expected  []*endpoint.Endpoint
	}{
		{
			title: "one rule.host one lb.hostname",
			httpProxy: fakeHTTPProxy{
				host: "foo.bar", // Kubernetes requires removal of trailing dot
				loadBalancer: fakeLoadBalancerService{
					hostnames: []string{"lb.com"}, // Kubernetes omits the trailing dot
				},
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
			httpProxy: fakeHTTPProxy{
				host: "foo.bar",
				loadBalancer: fakeLoadBalancerService{
					ips: []string{"8.8.8.8"},
				},
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
			httpProxy: fakeHTTPProxy{
				host: "foo.bar",
				loadBalancer: fakeLoadBalancerService{
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
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
			title:     "no rule.host",
			httpProxy: fakeHTTPProxy{},
			expected:  []*endpoint.Endpoint{},
		},
		{
			title: "one rule.host invalid httpproxy",
			httpProxy: fakeHTTPProxy{
				host:    "foo.bar",
				invalid: true,
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:     "no targets",
			httpProxy: fakeHTTPProxy{},
			expected:  []*endpoint.Endpoint{},
		},
		{
			title: "delegate httpproxy",
			httpProxy: fakeHTTPProxy{
				delegate: true,
			},
			expected: []*endpoint.Endpoint{},
		},
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			if source, err := newTestHTTPProxySource(); err != nil {
				require.NoError(t, err)
			} else if endpoints, err := source.endpointsFromHTTPProxy(ti.httpProxy.HTTPProxy()); err != nil {
				require.NoError(t, err)
			} else {
				validateEndpoints(t, endpoints, ti.expected)
			}
		})
	}
}

func testHTTPProxyEndpoints(t *testing.T) {
	t.Parallel()

	namespace := "testing"
	for _, ti := range []struct {
		title                    string
		targetNamespace          string
		annotationFilter         string
		loadBalancer             fakeLoadBalancerService
		httpProxyItems           []fakeHTTPProxy
		expected                 []*endpoint.Endpoint
		expectError              bool
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		ignoreHostnameAnnotation bool
	}{
		{
			title:           "no httpproxy",
			targetNamespace: "",
		},
		{
			title:           "two simple httpproxys",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
			},
			httpProxyItems: []fakeHTTPProxy{
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
			title:           "two simple httpproxys on different namespaces",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
			},
			httpProxyItems: []fakeHTTPProxy{
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
			title:           "two simple httpproxys on different namespaces and a target namespace",
			targetNamespace: "testing1",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"lb.com"},
			},
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
			title:           "template for httpproxy if host is missing",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{"8.8.8.8"},
				hostnames: []string{"elb.com"},
			},
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
						targetAnnotationKey: "httpproxy-target.com",
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
					Targets:    endpoint.Targets{"httpproxy-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dns.test.com",
					Targets:    endpoint.Targets{"httpproxy-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dna.test.com",
					Targets:    endpoint.Targets{"httpproxy-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
			fqdnTemplate:             "{{.Name}}.ext-dns.test.com, {{.Name}}.ext-dna.test.com",
			combineFQDNAndAnnotation: true,
		},
		{
			title:           "httpproxy rules with annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			httpProxyItems: []fakeHTTPProxy{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "httpproxy-target.com",
					},
					host: "example.org",
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "httpproxy-target.com",
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
					Targets:    endpoint.Targets{"httpproxy-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "example2.org",
					Targets:    endpoint.Targets{"httpproxy-target.com"},
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
			title:           "httpproxy rules with hostname annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"1.2.3.4"},
			},
			httpProxyItems: []fakeHTTPProxy{
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
			title:           "httpproxy rules with hostname annotation having multiple hostnames",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"1.2.3.4"},
			},
			httpProxyItems: []fakeHTTPProxy{
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
			title:           "httpproxy rules with hostname and target annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{},
			},
			httpProxyItems: []fakeHTTPProxy{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com",
						targetAnnotationKey:   "httpproxy-target.com",
					},
					host: "example.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"httpproxy-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "dns-through-hostname.com",
					Targets:    endpoint.Targets{"httpproxy-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
		},
		{
			title:           "httpproxy rules with annotation and custom TTL",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips: []string{"8.8.8.8"},
			},
			httpProxyItems: []fakeHTTPProxy{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "httpproxy-target.com",
						ttlAnnotationKey:    "6",
					},
					host: "example.org",
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "httpproxy-target.com",
						ttlAnnotationKey:    "1",
					},
					host: "example2.org",
				},
				{
					name:      "fake3",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "httpproxy-target.com",
						ttlAnnotationKey:    "10s",
					},
					host: "example3.org",
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:   "example.org",
					Targets:   endpoint.Targets{"httpproxy-target.com"},
					RecordTTL: endpoint.TTL(6),
				},
				{
					DNSName:   "example2.org",
					Targets:   endpoint.Targets{"httpproxy-target.com"},
					RecordTTL: endpoint.TTL(1),
				},
				{
					DNSName:   "example3.org",
					Targets:   endpoint.Targets{"httpproxy-target.com"},
					RecordTTL: endpoint.TTL(10),
				},
			},
		},
		{
			title:           "template for httpproxy with annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{},
				hostnames: []string{},
			},
			httpProxyItems: []fakeHTTPProxy{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "httpproxy-target.com",
					},
					host: "",
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "httpproxy-target.com",
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
					Targets:    endpoint.Targets{"httpproxy-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dns.test.com",
					Targets:    endpoint.Targets{"httpproxy-target.com"},
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
			title:           "httpproxy with empty annotation",
			targetNamespace: "",
			loadBalancer: fakeLoadBalancerService{
				ips:       []string{},
				hostnames: []string{},
			},
			httpProxyItems: []fakeHTTPProxy{
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
			httpProxyItems: []fakeHTTPProxy{
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
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			httpProxies := make([]*projectcontour.HTTPProxy, 0)
			for _, item := range ti.httpProxyItems {
				item.loadBalancer = ti.loadBalancer
				httpProxies = append(httpProxies, item.HTTPProxy())
			}

			fakeDynamicClient, scheme := newDynamicKubernetesClient()
			for _, httpProxy := range httpProxies {
				converted, err := convertHTTPProxyToUnstructured(httpProxy, scheme)
				require.NoError(t, err)
				_, err = fakeDynamicClient.Resource(projectcontour.HTTPProxyGVR).Namespace(httpProxy.Namespace).Create(context.Background(), converted, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			httpProxySource, err := NewContourHTTPProxySource(
				context.TODO(),
				fakeDynamicClient,
				ti.targetNamespace,
				ti.annotationFilter,
				ti.fqdnTemplate,
				ti.combineFQDNAndAnnotation,
				ti.ignoreHostnameAnnotation,
			)
			require.NoError(t, err)

			res, err := httpProxySource.Endpoints(context.Background())
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			validateEndpoints(t, res, ti.expected)
		})
	}
}

// httpproxy specific helper functions
func newTestHTTPProxySource() (*httpProxySource, error) {
	fakeDynamicClient, _ := newDynamicKubernetesClient()

	src, err := NewContourHTTPProxySource(
		context.TODO(),
		fakeDynamicClient,
		"default",
		"",
		"{{.Name}}",
		false,
		false,
	)
	if err != nil {
		return nil, err
	}

	irsrc, ok := src.(*httpProxySource)
	if !ok {
		return nil, errors.New("underlying source type was not httpproxy")
	}

	return irsrc, nil
}

type fakeHTTPProxy struct {
	namespace   string
	name        string
	annotations map[string]string

	host         string
	invalid      bool
	delegate     bool
	loadBalancer fakeLoadBalancerService
}

func (ir fakeHTTPProxy) HTTPProxy() *projectcontour.HTTPProxy {
	var status string
	if ir.invalid {
		status = "invalid"
	} else {
		status = "valid"
	}

	var spec projectcontour.HTTPProxySpec
	if ir.delegate {
		spec = projectcontour.HTTPProxySpec{}
	} else {
		spec = projectcontour.HTTPProxySpec{
			VirtualHost: &projectcontour.VirtualHost{
				Fqdn: ir.host,
			},
		}
	}

	lb := v1.LoadBalancerStatus{
		Ingress: []v1.LoadBalancerIngress{},
	}

	for _, ip := range ir.loadBalancer.ips {
		lb.Ingress = append(lb.Ingress, v1.LoadBalancerIngress{
			IP: ip,
		})
	}
	for _, hostname := range ir.loadBalancer.hostnames {
		lb.Ingress = append(lb.Ingress, v1.LoadBalancerIngress{
			Hostname: hostname,
		})
	}

	httpProxy := &projectcontour.HTTPProxy{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   ir.namespace,
			Name:        ir.name,
			Annotations: ir.annotations,
		},
		Spec: spec,
		Status: projectcontour.HTTPProxyStatus{
			CurrentStatus: status,
			LoadBalancer:  lb,
		},
	}

	return httpProxy
}
