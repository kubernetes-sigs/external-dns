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

	"github.com/kubernetes-sigs/external-dns/endpoint"

	//"strconv"
	//"sync"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	v1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

// This is a compile-time validation that gatewaySource is a Source.
var _ Source = &istioVirtualServiceSource{}

var virtualServiceType = istiomodel.VirtualService.Type

type VirtualServiceSuite struct {
	suite.Suite
	source     Source
	lbServices []*v1.Service
	gwconfig   istiomodel.Config
	vsconfig   istiomodel.Config
}

func (suite *VirtualServiceSuite) SetupTest() {
	fakeKubernetesClient := fake.NewSimpleClientset()
	fakeIstioClient := NewFakeConfigStore()
	var err error

	suite.lbServices = []*v1.Service{
		(fakeIngressGatewayService{
			ips:       []string{"8.8.8.8"},
			hostnames: []string{"v1"},
			namespace: "istio-system",
			name:      "istio-gateway1",
		}).Service(),
		(fakeIngressGatewayService{
			ips:       []string{"1.1.1.1"},
			hostnames: []string{"v42"},
			namespace: "istio-system",
			name:      "istio-gateway2",
		}).Service(),
	}

	for _, service := range suite.lbServices {
		_, err = fakeKubernetesClient.CoreV1().Services(service.Namespace).Create(service)
		suite.NoError(err, "should succeed")
	}

	suite.source, err = NewIstioVirtualServiceSource(
		fakeKubernetesClient,
		fakeIstioClient,
		"",
		"",
		"{{.Name}}",
		false,
		false,
	)
	suite.NoError(err, "should initialize virtualservice source")

	suite.gwconfig = (fakeGatewayConfig{
		name:      "foo-gateway-with-targets",
		namespace: "istio-system",
		dnsnames:  [][]string{{}},
	}).Config()
	_, err = fakeIstioClient.Create(suite.gwconfig)
	suite.NoError(err, "should succeed")

	suite.vsconfig = (fakeVirtualServiceConfig{
		name:      "foo-virtualservice",
		namespace: "istio-other",
		gateways:  []string{"istio-system/foo-gateway-with-targets"},
		dnsnames:  []string{"foo"},
	}).Config()
	_, err = fakeIstioClient.Create(suite.vsconfig)
	suite.NoError(err, "should succeed")
}

func (suite *VirtualServiceSuite) TestResourceLabelIsSet() {
	endpoints, err := suite.source.Endpoints()
	suite.NoError(err, "should succeed")
	suite.Equal(len(endpoints), 2, "should return the correct number of endpoints")
	for _, ep := range endpoints {
		suite.Equal("virtualservice/istio-other/foo-virtualservice", ep.Labels[endpoint.ResourceLabelKey], "should set correct resource label")
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
			annotationFilter: "kubernetes.io/gateway.class=nginx",
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			_, err := NewIstioVirtualServiceSource(
				fake.NewSimpleClientset(),
				NewFakeConfigStore(),
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
		lbServices []fakeIngressGatewayService
		gwconfig   fakeGatewayConfig
		vsconfig   fakeVirtualServiceConfig
		expected   []*endpoint.Endpoint
	}{
		{
			title: "one rule.host one lb.hostname",
			lbServices: []fakeIngressGatewayService{
				{
					hostnames: []string{"lb.com"}, // Kubernetes omits the trailing dot
				},
			},
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"}, // Kubernetes requires removal of trailing dot
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
			title: "one namespaced rule.host one lb.hostname",
			lbServices: []fakeIngressGatewayService{
				{
					hostnames: []string{"lb.com"}, // Kubernetes omits the trailing dot
				},
			},
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{
					{"my-namespace/foo.bar"}, // Kubernetes requires removal of trailing dot
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
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"},
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
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"},
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
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "one empty rule.host",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{
					{""},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:      "no targets",
			lbServices: []fakeIngressGatewayService{{}},
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{
					{""},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "one gateway, two ingressgateway loadbalancer hostnames",
			lbServices: []fakeIngressGatewayService{
				{
					hostnames: []string{"lb.com"},
					namespace: "istio-other",
					name:      "gateway1",
				},
				{
					hostnames: []string{"lb2.com"},
					namespace: "istio-other",
					name:      "gateway2",
				},
			},
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"}, // Kubernetes requires removal of trailing dot
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
			if source, err := newTestVirtualServiceSource(ti.lbServices, []fakeGatewayConfig{ti.gwconfig}); err != nil {
				require.NoError(t, err)
			} else if endpoints, err := source.endpointsFromVirtualServiceConfig(ti.vsconfig.Config()); err != nil {
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
		lbServices               []fakeIngressGatewayService
		gwConfigs                []fakeGatewayConfig
		vsConfigs                []fakeVirtualServiceConfig
		expected                 []*endpoint.Endpoint
		expectError              bool
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		ignoreHostnameAnnotation bool
	}{
		{
			title:           "no gateway",
			targetNamespace: "",
		},
		{
			title:           "two simple gateways, one ingressgateway loadbalancer service",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: namespace,
					dnsnames:  [][]string{{"new.org"}},
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
			title:           "two simple gateways on different namespaces, one ingressgateway loadbalancer service",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "testing1",
					dnsnames:  [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: "testing2",
					dnsnames:  [][]string{{"new.org"}},
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
			title:           "two simple gateways on different namespaces and a target namespace, one ingressgateway loadbalancer service",
			targetNamespace: "testing1",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
					namespace: "testing1",
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "testing1",
					dnsnames:  [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: "testing2",
					dnsnames:  [][]string{{"new.org"}},
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
			annotationFilter: "kubernetes.io/gateway.class in (alb, nginx)",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/gateway.class": "nginx",
					},
					dnsnames: [][]string{{"example.org"}},
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
			annotationFilter: "kubernetes.io/gateway.class in (alb, nginx)",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/gateway.class": "tectonic",
					},
					dnsnames: [][]string{{"example.org"}},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:            "invalid annotation filter expression",
			targetNamespace:  "",
			annotationFilter: "kubernetes.io/gateway.name in (a b)",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/gateway.class": "alb",
					},
					dnsnames: [][]string{{"example.org"}},
				},
			},
			expected:    []*endpoint.Endpoint{},
			expectError: true,
		},
		{
			title:            "valid matching annotation filter label",
			targetNamespace:  "",
			annotationFilter: "kubernetes.io/gateway.class=nginx",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/gateway.class": "nginx",
					},
					dnsnames: [][]string{{"example.org"}},
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
			annotationFilter: "kubernetes.io/gateway.class=nginx",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/gateway.class": "alb",
					},
					dnsnames: [][]string{{"example.org"}},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "our controller type is dns-controller",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: controllerAnnotationValue,
					},
					dnsnames: [][]string{{"example.org"}},
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
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: "some-other-tool",
					},
					dnsnames: [][]string{{"example.org"}},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "template for gateway if host is missing",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"elb.com"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: controllerAnnotationValue,
					},
					dnsnames: [][]string{},
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
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: "other-controller",
					},
					dnsnames: [][]string{},
				},
			},
			expected:     []*endpoint.Endpoint{},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title:           "multiple FQDN template hostnames",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:        "fake1",
					namespace:   namespace,
					annotations: map[string]string{},
					dnsnames:    [][]string{},
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
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:        "fake1",
					namespace:   namespace,
					annotations: map[string]string{},
					dnsnames:    [][]string{},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
					},
					dnsnames: [][]string{{"example.org"}},
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
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dns.test.com",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dna.test.com",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
			fqdnTemplate:             "{{.Name}}.ext-dns.test.com, {{.Name}}.ext-dna.test.com",
			combineFQDNAndAnnotation: true,
		},
		{
			title:           "gateway rules with annotation",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
					},
					dnsnames: [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
					},
					dnsnames: [][]string{{"example2.org"}},
				},
				{
					name:      "fake3",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "1.2.3.4",
					},
					dnsnames: [][]string{{"example3.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "example2.org",
					Targets:    endpoint.Targets{"gateway-target.com"},
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
			title:           "gateway rules with hostname annotation",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"1.2.3.4"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com",
					},
					dnsnames: [][]string{{"example.org"}},
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
			title:           "gateway rules with hostname annotation having multiple hostnames",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"1.2.3.4"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com, another-dns-through-hostname.com",
					},
					dnsnames: [][]string{{"example.org"}},
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
			title:           "gateway rules with hostname and target annotation",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "dns-through-hostname.com",
						targetAnnotationKey:   "gateway-target.com",
					},
					dnsnames: [][]string{{"example.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "dns-through-hostname.com",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
		},
		{
			title:           "gateway rules with annotation and custom TTL",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
						ttlAnnotationKey:    "6",
					},
					dnsnames: [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
						ttlAnnotationKey:    "1",
					},
					dnsnames: [][]string{{"example2.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:   "example.org",
					Targets:   endpoint.Targets{"gateway-target.com"},
					RecordTTL: endpoint.TTL(6),
				},
				{
					DNSName:   "example2.org",
					Targets:   endpoint.Targets{"gateway-target.com"},
					RecordTTL: endpoint.TTL(1),
				},
			},
		},
		{
			title:           "template for gateway with annotation",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{},
					hostnames: []string{},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
					},
					dnsnames: [][]string{},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
					},
					dnsnames: [][]string{},
				},
				{
					name:      "fake3",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "1.2.3.4",
					},
					dnsnames: [][]string{},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "fake1.ext-dns.test.com",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "fake2.ext-dns.test.com",
					Targets:    endpoint.Targets{"gateway-target.com"},
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
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{},
					hostnames: []string{},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "",
					},
					dnsnames: [][]string{},
				},
			},
			expected:     []*endpoint.Endpoint{},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title:           "ignore hostname annotations",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "ignore.me",
					},
					dnsnames: [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						hostnameAnnotationKey: "ignore.me.too",
					},
					dnsnames: [][]string{{"new.org"}},
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
			for _, item := range ti.gwConfigs {
				configs = append(configs, item.Config())
			}
			for _, item := range ti.vsConfigs {
				configs = append(configs, item.Config())
			}

			fakeKubernetesClient := fake.NewSimpleClientset()

			for _, lb := range ti.lbServices {
				service := lb.Service()
				_, err := fakeKubernetesClient.CoreV1().Services(service.Namespace).Create(service)
				require.NoError(t, err)
			}

			fakeIstioClient := NewFakeConfigStore()
			for _, config := range configs {
				_, err := fakeIstioClient.Create(config)
				require.NoError(t, err)
			}

			virtualServiceSource, err := NewIstioVirtualServiceSource(
				fakeKubernetesClient,
				fakeIstioClient,
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

func newTestVirtualServiceSource(loadBalancerList []fakeIngressGatewayService, gwList []fakeGatewayConfig) (*istioVirtualServiceSource, error) {
	fakeKubernetesClient := fake.NewSimpleClientset()
	fakeIstioClient := NewFakeConfigStore()

	for _, gw := range gwList {
		_, err := fakeIstioClient.Create(gw.Config())
		if err != nil {
			return nil, err
		}
	}

	for _, lb := range loadBalancerList {
		service := lb.Service()
		_, err := fakeKubernetesClient.CoreV1().Services(service.Namespace).Create(service)
		if err != nil {
			return nil, err
		}
	}

	src, err := NewIstioVirtualServiceSource(
		fakeKubernetesClient,
		fakeIstioClient,
		"",
		"",
		"{{.Name}}",
		false,
		false,
	)
	if err != nil {
		return nil, err
	}

	vssrc, ok := src.(*istioVirtualServiceSource)
	if !ok {
		return nil, errors.New("underlying source type was not virtualservice")
	}

	return vssrc, nil
}

type fakeVirtualServiceConfig struct {
	namespace   string
	name        string
	gateways    []string
	annotations map[string]string
	dnsnames    []string
}

func (c fakeVirtualServiceConfig) Config() istiomodel.Config {
	vs := &istionetworking.VirtualService{
		Gateways: c.gateways,
		Hosts:    c.dnsnames,
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