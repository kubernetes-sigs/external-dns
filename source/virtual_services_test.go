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

	istionetworking "istio.io/api/networking/v1alpha3"
	istiomodel "istio.io/istio/pilot/pkg/model"

	"strconv"
	"sync"

	"github.com/kubernetes-incubator/external-dns/endpoint"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

// This is a compile-time validation that virtualServiceSource is a Source.
var _ Source = &virtualServiceSource{}

var virtualServiceType = istiomodel.VirtualService.Type

type VirtualServiceSuite struct {
	suite.Suite
	source     Source
	lbServices []*v1.Service
	config     istiomodel.Config
}

func (suite *VirtualServiceSuite) SetupTest() {
	fakeKubernetesClient := fake.NewSimpleClientset()
	fakeIstioClient := NewFakeVsConfigStore()
	var err error

	suite.lbServices = []*v1.Service{
		(fakeIngressVirtualServiceService{
			ips:       []string{"8.8.8.8"},
			hostnames: []string{"v1"},
			namespace: "istio-system",
			name:      "istio-virtualService1",
		}).Service(),
		(fakeIngressVirtualServiceService{
			ips:       []string{"1.1.1.1"},
			hostnames: []string{"v42"},
			namespace: "istio-other",
			name:      "istio-virtualService2",
		}).Service(),
	}

	for _, loadBalancer := range suite.lbServices {
		_, err = fakeKubernetesClient.CoreV1().Services(loadBalancer.Namespace).Create(loadBalancer)
		suite.NoError(err, "should succeed")
	}

	suite.source, err = NewIstioVirtualServiceSource(
		fakeKubernetesClient,
		fakeIstioClient,
		[]string{"istio-system/istio-ingressvirtualService"},
		"default",
		"",
		"{{.Name}}",
		false,
		false,
	)
	suite.NoError(err, "should initialize virtualService source")

	suite.config = (fakeVirtualServiceConfig{
		name:      "foo-virtualService-with-targets",
		namespace: "default",
		dnsnames:  []string{"foo"},
	}).Config()
	_, err = fakeIstioClient.Create(suite.config)
	suite.NoError(err, "should succeed")
}

func (suite *VirtualServiceSuite) TestResourceLabelIsSet() {
	endpoints, _ := suite.source.Endpoints()
	for _, ep := range endpoints {
		suite.Equal("virtualService/default/foo-virtualService-with-targets", ep.Labels[endpoint.ResourceLabelKey], "should set correct resource label")
	}
}

func TestVirtualService(t *testing.T) {
	suite.Run(t, new(VirtualServiceSuite))
	t.Run("endpointsFromVirtualServiceConfig", testEndpointsFromVirtualServiceConfig)
	t.Run("Endpoints", testVirtualServiceEndpoints)
}

func TestNewIstioVirtualServiceSource(t *testing.T) {
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
			fqdnTemplate: "{{.Name}",
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
			annotationFilter: "kubernetes.io/virtualService.class=nginx",
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			_, err := NewIstioVirtualServiceSource(
				fake.NewSimpleClientset(),
				NewFakeVsConfigStore(),
				[]string{"istio-system/istio-ingressvirtualService"},
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

func testEndpointsFromVirtualServiceConfig(t *testing.T) {
	for _, ti := range []struct {
		title      string
		lbServices []fakeIngressVirtualServiceService
		config     fakeVirtualServiceConfig
		expected   []*endpoint.Endpoint
	}{
		{
			title: "one rule.host one lb.hostname",
			lbServices: []fakeIngressVirtualServiceService{
				{
					hostnames: []string{"lb.com"}, // Kubernetes omits the trailing dot
				},
			},
			config: fakeVirtualServiceConfig{
				dnsnames: []string{
					"foo.bar", // Kubernetes requires removal of trailing dot
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
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			config: fakeVirtualServiceConfig{
				dnsnames: []string{
					"foo.bar",
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
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			config: fakeVirtualServiceConfig{
				dnsnames: []string{
					"foo.bar",
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
			title: "no rule.host",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			config: fakeVirtualServiceConfig{
				dnsnames: []string{},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "one empty rule.host",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			config: fakeVirtualServiceConfig{
				dnsnames: []string{
					"",
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:      "no targets",
			lbServices: []fakeIngressVirtualServiceService{{}},
			config: fakeVirtualServiceConfig{
				dnsnames: []string{
					"",
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "one virtualService, two ingressvirtualService loadbalancer hostnames",
			lbServices: []fakeIngressVirtualServiceService{
				{
					hostnames: []string{"lb.com"},
					namespace: "istio-other",
					name:      "virtualService1",
				},
				{
					hostnames: []string{"lb2.com"},
					namespace: "istio-other",
					name:      "virtualService2",
				},
			},
			config: fakeVirtualServiceConfig{
				dnsnames: []string{
					"foo.bar", // Kubernetes requires removal of trailing dot
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "foo.bar",
					Targets: endpoint.Targets{"lb.com", "lb2.com"},
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			if source, err := newTestVirtualServiceSource(ti.lbServices); err != nil {
				require.NoError(t, err)
			} else if endpoints, err := source.endpointsFromVirtualServiceConfig(ti.config.Config()); err != nil {
				require.NoError(t, err)
			} else {
				validateEndpoints(t, endpoints, ti.expected)
			}
		})
	}
}

func testVirtualServiceEndpoints(t *testing.T) {
	namespace := "testing"
	for _, ti := range []struct {
		title                    string
		targetNamespace          string
		annotationFilter         string
		lbServices               []fakeIngressVirtualServiceService
		configItems              []fakeVirtualServiceConfig
		expected                 []*endpoint.Endpoint
		expectError              bool
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		ignoreHostnameAnnotation bool
	}{
		{
			title:           "no virtualService",
			targetNamespace: "",
		},
		{
			title:           "two simple virtualServices, one ingressvirtualService loadbalancer service",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "fake2",
					namespace: namespace,
					dnsnames:  []string{"new.org"},
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
			title:           "two simple virtualServices on different namespaces, one ingressvirtualService loadbalancer service",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: "testing1",
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "fake2",
					namespace: "testing2",
					dnsnames:  []string{"new.org"},
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
			title:           "two simple virtualServices on different namespaces and a target namespace, one ingressvirtualService loadbalancer service",
			targetNamespace: "testing1",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: "testing1",
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "fake2",
					namespace: "testing2",
					dnsnames:  []string{"new.org"},
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
			annotationFilter: "kubernetes.io/virtualService.class in (alb, nginx)",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/virtualService.class": "nginx",
					},
					dnsnames: []string{"example.org"},
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
			annotationFilter: "kubernetes.io/virtualService.class in (alb, nginx)",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/virtualService.class": "tectonic",
					},
					dnsnames: []string{"example.org"},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:            "invalid annotation filter expression",
			targetNamespace:  "",
			annotationFilter: "kubernetes.io/virtualService.name in (a b)",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/virtualService.class": "alb",
					},
					dnsnames: []string{"example.org"},
				},
			},
			expected:    []*endpoint.Endpoint{},
			expectError: true,
		},
		{
			title:            "valid matching annotation filter label",
			targetNamespace:  "",
			annotationFilter: "kubernetes.io/virtualService.class=nginx",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/virtualService.class": "nginx",
					},
					dnsnames: []string{"example.org"},
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
			annotationFilter: "kubernetes.io/virtualService.class=nginx",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/virtualService.class": "alb",
					},
					dnsnames: []string{"example.org"},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "our controller type is dns-controller",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: controllerAnnotationValue,
					},
					dnsnames: []string{"example.org"},
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
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: "some-other-tool",
					},
					dnsnames: []string{"example.org"},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "template for virtualService if host is missing",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"elb.com"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: controllerAnnotationValue,
					},
					dnsnames: []string{},
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
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: "other-controller",
					},
					dnsnames: []string{},
				},
			},
			expected:     []*endpoint.Endpoint{},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title:           "multiple FQDN template hostnames",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:        "fake1",
					namespace:   namespace,
					annotations: map[string]string{},
					dnsnames:    []string{},
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
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:        "fake1",
					namespace:   namespace,
					annotations: map[string]string{},
					dnsnames:    []string{},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "virtualService-target.com",
					},
					dnsnames: []string{"example.org"},
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
					Targets:    endpoint.Targets{"virtualService-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dns.test.com",
					Targets:    endpoint.Targets{"virtualService-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dna.test.com",
					Targets:    endpoint.Targets{"virtualService-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
			fqdnTemplate:             "{{.Name}}.ext-dns.test.com, {{.Name}}.ext-dna.test.com",
			combineFQDNAndAnnotation: true,
		},
		{
			title:           "virtualService rules with annotation",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "virtualService-target.com",
					},
					dnsnames: []string{"example.org"},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "virtualService-target.com",
					},
					dnsnames: []string{"example2.org"},
				},
				{
					name:      "fake3",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "1.2.3.4",
					},
					dnsnames: []string{"example3.org"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"virtualService-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "example2.org",
					Targets:    endpoint.Targets{"virtualService-target.com"},
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
			title:           "virtualService rules with hostname annotation",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"1.2.3.4"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com",
					},
					dnsnames: []string{"example.org"},
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
			title:           "virtualService rules with hostname annotation having multiple hostnames",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"1.2.3.4"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com, another-dns-through-hostname.com",
					},
					dnsnames: []string{"example.org"},
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
			title:           "virtualService rules with hostname and target annotation",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com",
						targetAnnotationKey:   "virtualService-target.com",
					},
					dnsnames: []string{"example.org"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"virtualService-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "dns-through-hostname.com",
					Targets:    endpoint.Targets{"virtualService-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
		},
		{
			title:           "virtualService rules with annotation and custom TTL",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "virtualService-target.com",
						ttlAnnotationKey:    "6",
					},
					dnsnames: []string{"example.org"},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "virtualService-target.com",
						ttlAnnotationKey:    "1",
					},
					dnsnames: []string{"example2.org"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:   "example.org",
					Targets:   endpoint.Targets{"virtualService-target.com"},
					RecordTTL: endpoint.TTL(6),
				},
				{
					DNSName:   "example2.org",
					Targets:   endpoint.Targets{"virtualService-target.com"},
					RecordTTL: endpoint.TTL(1),
				},
			},
		},
		{
			title:           "template for virtualService with annotation",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{},
					hostnames: []string{},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "virtualService-target.com",
					},
					dnsnames: []string{},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "virtualService-target.com",
					},
					dnsnames: []string{},
				},
				{
					name:      "fake3",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "1.2.3.4",
					},
					dnsnames: []string{},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "fake1.ext-dns.test.com",
					Targets:    endpoint.Targets{"virtualService-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dns.test.com",
					Targets:    endpoint.Targets{"virtualService-target.com"},
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
			title:           "Ingress with empty annotation",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{},
					hostnames: []string{},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "",
					},
					dnsnames: []string{},
				},
			},
			expected:     []*endpoint.Endpoint{},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title:           "ignore hostname annotations",
			targetNamespace: "",
			lbServices: []fakeIngressVirtualServiceService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
				},
			},
			configItems: []fakeVirtualServiceConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "ignore.me",
					},
					dnsnames: []string{"example.org"},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "ignore.me.too",
					},
					dnsnames: []string{"new.org"},
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
			configs := make([]istiomodel.Config, 0)
			for _, item := range ti.configItems {
				configs = append(configs, item.Config())
			}

			fakeKubernetesClient := fake.NewSimpleClientset()

			var fakeLoadBalancerList []string
			for _, lb := range ti.lbServices {
				lbService := lb.Service()
				_, err := fakeKubernetesClient.CoreV1().Services(lbService.Namespace).Create(lbService)
				if err != nil {
					require.NoError(t, err)
				}
				fakeLoadBalancerList = append(fakeLoadBalancerList, lbService.Namespace+"/"+lbService.Name)
			}

			fakeIstioClient := NewFakeVsConfigStore()
			for _, config := range configs {
				_, err := fakeIstioClient.Create(config)
				require.NoError(t, err)
			}

			virtualServiceSource, err := NewIstioVirtualServiceSource(
				fakeKubernetesClient,
				fakeIstioClient,
				fakeLoadBalancerList,
				ti.targetNamespace,
				ti.annotationFilter,
				ti.fqdnTemplate,
				ti.combineFQDNAndAnnotation,
				ti.ignoreHostnameAnnotation,
			)
			require.NoError(t, err)

			res, err := virtualServiceSource.Endpoints()
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			validateEndpoints(t, res, ti.expected)
		})
	}
}

// virtualService specific helper functions
func newTestVirtualServiceSource(loadBalancerList []fakeIngressVirtualServiceService) (*virtualServiceSource, error) {
	fakeKubernetesClient := fake.NewSimpleClientset()
	fakeIstioClient := NewFakeVsConfigStore()

	var lbList []string
	for _, lb := range loadBalancerList {
		lbService := lb.Service()
		_, err := fakeKubernetesClient.CoreV1().Services(lbService.Namespace).Create(lbService)
		if err != nil {
			return nil, err
		}
		lbList = append(lbList, lbService.Namespace+"/"+lbService.Name)
	}

	src, err := NewIstioVirtualServiceSource(
		fakeKubernetesClient,
		fakeIstioClient,
		lbList,
		"default",
		"",
		"{{.Name}}",
		false,
		false,
	)
	if err != nil {
		return nil, err
	}

	vssrc, ok := src.(*virtualServiceSource)
	if !ok {
		return nil, errors.New("underlying source type was not virtualService")
	}

	return vssrc, nil
}

type fakeIngressVirtualServiceService struct {
	ips       []string
	hostnames []string
	namespace string
	name      string
}

func (ig fakeIngressVirtualServiceService) Service() *v1.Service {
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

type fakeVirtualServiceConfig struct {
	namespace   string
	name        string
	annotations map[string]string
	dnsnames    []string
}

func (c fakeVirtualServiceConfig) Config() istiomodel.Config {
	vs := &istionetworking.VirtualService{
		Hosts: c.dnsnames,
	}

	config := istiomodel.Config{
		ConfigMeta: istiomodel.ConfigMeta{
			Namespace:   c.namespace,
			Name:        c.name,
			Type:        virtualServiceType,
			Annotations: c.annotations,
		},
		Spec: vs,
	}

	return config
}

type fakeVsConfigStore struct {
	descriptor istiomodel.ConfigDescriptor
	configs    []*istiomodel.Config
	sync.RWMutex
}

func NewFakeVsConfigStore() istiomodel.ConfigStore {
	return &fakeVsConfigStore{
		descriptor: istiomodel.ConfigDescriptor{
			istiomodel.VirtualService,
		},
		configs: make([]*istiomodel.Config, 0),
	}
}

func (f *fakeVsConfigStore) ConfigDescriptor() istiomodel.ConfigDescriptor {
	return f.descriptor
}

func (f *fakeVsConfigStore) Get(typ, name, namespace string) (config *istiomodel.Config) {
	f.RLock()
	defer f.RUnlock()

	if cfg, _ := f.get(typ, name, namespace); cfg != nil {
		config = cfg
	}

	return
}

func (f *fakeVsConfigStore) get(typ, name, namespace string) (*istiomodel.Config, int) {
	for idx, cfg := range f.configs {
		if cfg.Type == typ && cfg.Name == name && cfg.Namespace == namespace {
			return cfg, idx
		}
	}

	return nil, -1
}

func (f *fakeVsConfigStore) List(typ, namespace string) (configs []istiomodel.Config, err error) {
	f.RLock()
	defer f.RUnlock()

	if namespace == "" {
		for _, cfg := range f.configs {
			configs = append(configs, *cfg)
		}
	} else {
		for _, cfg := range f.configs {
			if cfg.Type == typ && cfg.Namespace == namespace {
				configs = append(configs, *cfg)
			}
		}
	}

	return
}

func (f *fakeVsConfigStore) Create(config istiomodel.Config) (revision string, err error) {
	f.Lock()
	defer f.Unlock()

	if cfg, _ := f.get(config.Type, config.Name, config.Namespace); cfg != nil {
		err = errors.New("config already exists")
	} else {
		revision = "0"
		cfg := &config
		cfg.ResourceVersion = revision
		f.configs = append(f.configs, cfg)
	}

	return
}

func (f *fakeVsConfigStore) Update(config istiomodel.Config) (newRevision string, err error) {
	f.Lock()
	defer f.Unlock()

	if oldCfg, idx := f.get(config.Type, config.Name, config.Namespace); oldCfg == nil {
		err = errors.New("config does not exist")
	} else if oldRevision, e := strconv.Atoi(oldCfg.ResourceVersion); e != nil {
		err = e
	} else {
		newRevision = strconv.Itoa(oldRevision + 1)
		cfg := &config
		cfg.ResourceVersion = newRevision
		f.configs[idx] = cfg
	}

	return
}

func (f *fakeVsConfigStore) Delete(typ, name, namespace string) error {
	f.Lock()
	defer f.Unlock()

	_, idx := f.get(typ, name, namespace)
	if idx < 0 {
		return errors.New("config does not exist")
	}

	copy(f.configs[idx:], f.configs[idx+1:])
	f.configs[len(f.configs)-1] = nil
	f.configs = f.configs[:len(f.configs)-1]

	return nil
}
