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
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	istionetworking "istio.io/api/networking/v1beta1"
	networkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	istiofake "istio.io/client-go/pkg/clientset/versioned/fake"
	v1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"

	"sigs.k8s.io/external-dns/endpoint"
)

// This is a compile-time validation that gatewaySource is a Source.
var _ Source = &gatewaySource{}

type GatewaySuite struct {
	suite.Suite
	source     Source
	lbServices []*v1.Service
	ingresses  []*networkv1.Ingress
}

func (suite *GatewaySuite) SetupTest() {
	fakeKubernetesClient := fake.NewClientset()
	fakeIstioClient := istiofake.NewSimpleClientset()
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
			namespace: "istio-other",
			name:      "istio-gateway2",
		}).Service(),
	}

	for _, service := range suite.lbServices {
		_, err = fakeKubernetesClient.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
		suite.NoError(err, "should succeed")
	}

	suite.ingresses = []*networkv1.Ingress{
		(fakeIngress{
			ips:       []string{"2.2.2.2"},
			hostnames: []string{"v2"},
			namespace: "istio-system",
			name:      "istio-ingress",
		}).Ingress(),
		(fakeIngress{
			ips:       []string{"3.3.3.3"},
			hostnames: []string{"v62"},
			namespace: "istio-system",
			name:      "istio-ingress2",
		}).Ingress(),
	}

	for _, ingress := range suite.ingresses {
		_, err = fakeKubernetesClient.NetworkingV1().Ingresses(ingress.Namespace).Create(context.Background(), ingress, metav1.CreateOptions{})
		suite.NoError(err, "should succeed")
	}

	suite.source, err = NewIstioGatewaySource(
		context.TODO(),
		fakeKubernetesClient,
		fakeIstioClient,
		"",
		"",
		"{{.Name}}",
		false,
		false,
	)
	suite.NoError(err, "should initialize gateway source")
	suite.NoError(err, "should succeed")
}

func (suite *GatewaySuite) TestResourceLabelIsSet() {
	endpoints, _ := suite.source.Endpoints(context.Background())
	for _, ep := range endpoints {
		suite.Equal("gateway/default/foo-gateway-with-targets", ep.Labels[endpoint.ResourceLabelKey], "should set correct resource label")
	}
}

func TestGateway(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(GatewaySuite))
	t.Run("endpointsFromGatewayConfig", testEndpointsFromGatewayConfig)
	t.Run("Endpoints", testGatewayEndpoints)
}

func TestNewIstioGatewaySource(t *testing.T) {
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
			annotationFilter: "kubernetes.io/gateway.class=nginx",
		},
	} {

		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			_, err := NewIstioGatewaySource(
				context.TODO(),
				fake.NewClientset(),
				istiofake.NewSimpleClientset(),
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

func testEndpointsFromGatewayConfig(t *testing.T) {
	t.Parallel()

	for _, ti := range []struct {
		title      string
		lbServices []fakeIngressGatewayService
		ingresses  []fakeIngress
		config     fakeGatewayConfig
		expected   []*endpoint.Endpoint
	}{
		{
			title: "one rule.host one lb.hostname",
			lbServices: []fakeIngressGatewayService{
				{
					hostnames: []string{"lb.com"}, // Kubernetes omits the trailing dot
				},
			},
			config: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"}, // Kubernetes requires removal of trailing dot
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
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
			config: fakeGatewayConfig{
				dnsnames: [][]string{
					{"my-namespace/foo.bar"}, // Kubernetes requires removal of trailing dot
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
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
			config: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
			},
		},
		{
			title: "one rule.host one ingress.IP",
			ingresses: []fakeIngress{
				{
					name: "ingress1",
					ips:  []string{"8.8.8.8"},
				},
			},
			config: fakeGatewayConfig{
				annotations: map[string]string{
					IstioGatewayIngressSource: "ingress1",
				},
				dnsnames: [][]string{
					{"foo.bar"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
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
			config: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8", "127.0.0.1"},
				},
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"elb.com", "alb.com"},
				},
			},
		},
		{
			title: "one rule.host two ingress.IP and two ingress.Hostname",
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			config: fakeGatewayConfig{
				annotations: map[string]string{
					IstioGatewayIngressSource: "ingress1",
				},
				dnsnames: [][]string{
					{"foo.bar"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8", "127.0.0.1"},
				},
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"elb.com", "alb.com"},
				},
			},
		},
		{
			title: "no rule.host",
			lbServices: []fakeIngressGatewayService{
				{
					ips:         []string{"8.8.8.8", "127.0.0.1"},
					hostnames:   []string{"elb.com", "alb.com"},
					externalIPs: []string{"1.1.1.1", "2.2.2.2"},
				},
			},
			config: fakeGatewayConfig{
				dnsnames: [][]string{},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "one empty rule.host",
			lbServices: []fakeIngressGatewayService{
				{
					ips:         []string{"8.8.8.8", "127.0.0.1"},
					hostnames:   []string{"elb.com", "alb.com"},
					externalIPs: []string{"1.1.1.1", "2.2.2.2"},
				},
			},
			config: fakeGatewayConfig{
				dnsnames: [][]string{
					{""},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "one empty rule.host with gateway ingress annotation",
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			config: fakeGatewayConfig{
				annotations: map[string]string{
					IstioGatewayIngressSource: "ingress1",
				},
				dnsnames: [][]string{
					{""},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:      "no targets",
			lbServices: []fakeIngressGatewayService{{}},
			config: fakeGatewayConfig{
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
			config: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"}, // Kubernetes requires removal of trailing dot
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com", "lb2.com"},
				},
			},
		},
		{
			title: "one gateway, ingress in separate namespace",
			ingresses: []fakeIngress{
				{
					hostnames: []string{"lb.com"},
					namespace: "istio-other2",
					name:      "ingress1",
				},
				{
					hostnames: []string{"lb2.com"},
					namespace: "istio-other",
					name:      "ingress2",
				},
			},
			config: fakeGatewayConfig{
				annotations: map[string]string{
					IstioGatewayIngressSource: "istio-other2/ingress1",
				},
				dnsnames: [][]string{
					{"foo.bar"}, // Kubernetes requires removal of trailing dot
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title: "one rule.host one lb.externalIP",
			lbServices: []fakeIngressGatewayService{
				{
					externalIPs: []string{"8.8.8.8"},
				},
			},
			config: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
			},
		},
		{
			title: "one rule.host two lb.IP, two lb.Hostname and two lb.externalIP",
			lbServices: []fakeIngressGatewayService{
				{
					ips:         []string{"8.8.8.8", "127.0.0.1"},
					hostnames:   []string{"elb.com", "alb.com"},
					externalIPs: []string{"1.1.1.1", "2.2.2.2"},
				},
			},
			config: fakeGatewayConfig{
				dnsnames: [][]string{
					{"foo.bar"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.1.1.1", "2.2.2.2"},
				},
			},
		},
	} {

		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			gatewayCfg := ti.config.Config()
			source, err := newTestGatewaySource(ti.lbServices, ti.ingresses)
			require.NoError(t, err)
			hostnames := source.hostNamesFromGateway(gatewayCfg)
			endpoints, err := source.endpointsFromGateway(hostnames, gatewayCfg)
			require.NoError(t, err)
			validateEndpoints(t, endpoints, ti.expected)
		})
	}
}

func testGatewayEndpoints(t *testing.T) {
	t.Parallel()

	for _, ti := range []struct {
		title                    string
		targetNamespace          string
		annotationFilter         string
		lbServices               []fakeIngressGatewayService
		ingresses                []fakeIngress
		configItems              []fakeGatewayConfig
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					dnsnames:  [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: "",
					dnsnames:  [][]string{{"new.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
				},
				{
					DNSName:    "new.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "new.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					dnsnames:  [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: "",
					dnsnames:  [][]string{{"new.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
				},
				{
					DNSName:    "new.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "new.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "testing1",
					dnsnames:  [][]string{{"example.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
				},
			},
		},
		{
			title:           "one simple gateways on different namespace and a target namespace, one ingress service",
			targetNamespace: "testing1",
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
					namespace: "testing2",
				},
			},
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "testing1",
					dnsnames:  [][]string{{"example.org"}},
					annotations: map[string]string{
						IstioGatewayIngressSource: "testing2/ingress1",
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						"kubernetes.io/gateway.class": "nginx",
					},
					dnsnames: [][]string{{"example.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						"kubernetes.io/gateway.class": "nginx",
					},
					dnsnames: [][]string{{"example.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.ControllerKey: annotations.ControllerValue,
					},
					dnsnames: [][]string{{"example.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.ControllerKey: "some-other-tool",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.ControllerKey: annotations.ControllerValue,
					},
					dnsnames: [][]string{},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "fake1.ext-dns.test.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "fake1.ext-dns.test.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"elb.com"},
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.ControllerKey: "other-controller",
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
			configItems: []fakeGatewayConfig{
				{
					name:        "fake1",
					namespace:   "",
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
			configItems: []fakeGatewayConfig{
				{
					name:        "fake1",
					namespace:   "",
					annotations: map[string]string{},
					dnsnames:    [][]string{},
				},
				{
					name:      "fake2",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "gateway-target.com",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "gateway-target.com",
					},
					dnsnames: [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "gateway-target.com",
					},
					dnsnames: [][]string{{"example2.org"}},
				},
				{
					name:      "fake3",
					namespace: "",
					annotations: map[string]string{
						IstioGatewayIngressSource: "not-real/ingress1",
						annotations.TargetKey:     "1.2.3.4",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.HostnameKey: "dns-through-hostname.com",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.HostnameKey: "dns-through-hostname.com, another-dns-through-hostname.com",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.HostnameKey: "dns-through-hostname.com",
						annotations.TargetKey:   "gateway-target.com",
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
			title:           "gateway rules with hostname, target and ingress annotation",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{},
				},
			},
			ingresses: []fakeIngress{
				{
					name: "ingress1",
					ips:  []string{},
				},
			},
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						IstioGatewayIngressSource: "ingress1",
						annotations.HostnameKey:   "dns-through-hostname.com",
						annotations.TargetKey:     "gateway-target.com",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "gateway-target.com",
						annotations.TtlKey:    "6",
					},
					dnsnames: [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "gateway-target.com",
						annotations.TtlKey:    "1",
					},
					dnsnames: [][]string{{"example2.org"}},
				},
				{
					name:      "fake3",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "gateway-target.com",
						annotations.TtlKey:    "10s",
					},
					dnsnames: [][]string{{"example3.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordTTL:  endpoint.TTL(6),
				},
				{
					DNSName:    "example2.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordTTL:  endpoint.TTL(1),
				},
				{
					DNSName:    "example3.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordTTL:  endpoint.TTL(10),
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "gateway-target.com",
					},
					dnsnames: [][]string{},
				},
				{
					name:      "fake2",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "gateway-target.com",
					},
					dnsnames: [][]string{},
				},
				{
					name:      "fake3",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "1.2.3.4",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.TargetKey: "",
					},
					dnsnames: [][]string{},
				},
			},
			expected:     []*endpoint.Endpoint{},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title:           "Gateway with empty ingress annotation",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{},
					hostnames: []string{},
				},
			},
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					ips:       []string{},
					hostnames: []string{},
				},
			},
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						IstioGatewayIngressSource: "",
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
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.HostnameKey: "ignore.me",
					},
					dnsnames: [][]string{{"example.org"}},
				},
				{
					name:      "fake2",
					namespace: "",
					annotations: map[string]string{
						annotations.HostnameKey: "ignore.me.too",
					},
					dnsnames: [][]string{{"new.org"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
				},
				{
					DNSName:    "new.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "new.org",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"lb.com"},
				},
			},
			ignoreHostnameAnnotation: true,
		},
		{
			title:           "gateways with wildcard host",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"1.2.3.4"},
				},
			},
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					dnsnames:  [][]string{{"*"}},
				},
				{
					name:      "fake2",
					namespace: "",
					dnsnames:  [][]string{{"some-namespace/*"}},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "gateways with wildcard host and hostname annotation",
			targetNamespace: "",
			lbServices: []fakeIngressGatewayService{
				{
					ips: []string{"1.2.3.4"},
				},
			},
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						annotations.HostnameKey: "fake1.dns-through-hostname.com",
					},
					dnsnames: [][]string{{"*"}},
				},
				{
					name:      "fake2",
					namespace: "",
					annotations: map[string]string{
						annotations.HostnameKey: "fake2.dns-through-hostname.com",
					},
					dnsnames: [][]string{{"some-namespace/*"}},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "fake1.dns-through-hostname.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.2.3.4"},
				},
				{
					DNSName:    "fake2.dns-through-hostname.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"1.2.3.4"},
				},
			},
		},
		{
			title:           "gateways with ingress annotation; ingress not found",
			targetNamespace: "",
			ingresses: []fakeIngress{
				{
					name: "ingress1",
					ips:  []string{"8.8.8.8"},
				},
			},
			configItems: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "",
					annotations: map[string]string{
						IstioGatewayIngressSource: "ingress2",
					},
					dnsnames: [][]string{{"new.org"}},
				},
			},
			expected:    []*endpoint.Endpoint{},
			expectError: true,
		},
	} {

		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			fakeKubernetesClient := fake.NewClientset()
			targetNamespace := ti.targetNamespace

			for _, lb := range ti.lbServices {
				service := lb.Service()
				_, err := fakeKubernetesClient.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			for _, ing := range ti.ingresses {
				ingress := ing.Ingress()
				if ingress.Namespace != targetNamespace {
					targetNamespace = v1.NamespaceAll
				}
				_, err := fakeKubernetesClient.NetworkingV1().Ingresses(ingress.Namespace).Create(context.Background(), ingress, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			fakeIstioClient := istiofake.NewSimpleClientset()
			for _, config := range ti.configItems {
				gatewayCfg := config.Config()
				_, err := fakeIstioClient.NetworkingV1beta1().Gateways(ti.targetNamespace).Create(context.Background(), gatewayCfg, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			gatewaySource, err := NewIstioGatewaySource(
				context.TODO(),
				fakeKubernetesClient,
				fakeIstioClient,
				targetNamespace,
				ti.annotationFilter,
				ti.fqdnTemplate,
				ti.combineFQDNAndAnnotation,
				ti.ignoreHostnameAnnotation,
			)
			require.NoError(t, err)

			res, err := gatewaySource.Endpoints(context.Background())
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			validateEndpoints(t, res, ti.expected)
		})
	}
}

func TestGatewaySource_GWSelectorMatchServiceSelector(t *testing.T) {
	tests := []struct {
		name      string
		selectors map[string]string
		expected  []*endpoint.Endpoint
	}{
		{
			name: "gw single selector match with single service selector",
			selectors: map[string]string{
				"version": "v1",
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.org", endpoint.RecordTypeA, "10.10.10.255").WithLabel("resource", "gateway/default/fake-gateway"),
			},
		},
		{
			name: "gw selector match all service selectors",
			selectors: map[string]string{
				"app":     "demo",
				"env":     "prod",
				"team":    "devops",
				"version": "v1",
				"release": "stable",
				"track":   "daily",
				"tier":    "backend",
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.org", endpoint.RecordTypeA, "10.10.10.255").WithLabel("resource", "gateway/default/fake-gateway"),
			},
		},
		{
			name: "gw selector has subset of service selectors",
			selectors: map[string]string{
				"version": "v1",
				"release": "stable",
				"tier":    "backend",
				"app":     "demo",
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpoint("example.org", endpoint.RecordTypeA, "10.10.10.255").WithLabel("resource", "gateway/default/fake-gateway"),
			},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeKubeClient := fake.NewClientset()
			fakeIstioClient := istiofake.NewSimpleClientset()

			svc := &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "fake-service",
					Namespace: "default",
					UID:       types.UID(fmt.Sprintf("fake-service-uid-%d", i)),
				},
				Spec: v1.ServiceSpec{
					Selector: map[string]string{
						"app":     "demo",
						"env":     "prod",
						"team":    "devops",
						"version": "v1",
						"release": "stable",
						"track":   "daily",
						"tier":    "backend",
					},
					ExternalIPs: []string{"10.10.10.255"},
				},
			}
			_, err := fakeKubeClient.CoreV1().Services(svc.Namespace).Create(t.Context(), svc, metav1.CreateOptions{})
			require.NoError(t, err)

			gw := &networkingv1beta1.Gateway{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "fake-gateway",
					Namespace: "default",
				},
				Spec: istionetworking.Gateway{
					Servers: []*istionetworking.Server{
						{
							Hosts: []string{"example.org"},
						},
					},
					Selector: tt.selectors,
				},
			}

			_, err = fakeIstioClient.NetworkingV1beta1().Gateways(gw.Namespace).Create(context.Background(), gw, metav1.CreateOptions{})
			require.NoError(t, err)

			src, err := NewIstioGatewaySource(
				t.Context(),
				fakeKubeClient,
				fakeIstioClient,
				"",
				"",
				"",
				false,
				false,
			)
			require.NoError(t, err)
			require.NotNil(t, src)

			res, err := src.Endpoints(t.Context())
			require.NoError(t, err)

			validateEndpoints(t, res, tt.expected)
		})
	}
}

func TestTransformerInIstioGatewaySource(t *testing.T) {
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-service",
			Namespace: "default",
			Labels: map[string]string{
				"label1": "value1",
				"label2": "value2",
				"label3": "value3",
			},
			Annotations: map[string]string{
				"user-annotation": "value",
				"external-dns.alpha.kubernetes.io/hostname": "test-hostname",
				"external-dns.alpha.kubernetes.io/random":   "value",
				"other/annotation":                          "value",
			},
			UID: "someuid",
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"selector":  "one",
				"selector2": "two",
				"selector3": "three",
			},
			ExternalIPs: []string{"1.2.3.4"},
			Ports: []v1.ServicePort{
				{
					Name:       "http",
					Port:       80,
					TargetPort: intstr.FromInt32(8080),
					Protocol:   v1.ProtocolTCP,
				},
				{
					Name:       "https",
					Port:       443,
					TargetPort: intstr.FromInt32(8443),
					Protocol:   v1.ProtocolTCP,
				},
			},
			Type: v1.ServiceTypeLoadBalancer,
		},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{
				Ingress: []v1.LoadBalancerIngress{
					{IP: "5.6.7.8", Hostname: "lb.example.com"},
				},
			},
			Conditions: []metav1.Condition{
				{
					Type:               "Available",
					Status:             metav1.ConditionTrue,
					Reason:             "MinimumReplicasAvailable",
					Message:            "Service is available",
					LastTransitionTime: metav1.Now(),
				},
			},
		},
	}

	fakeClient := fake.NewClientset()

	_, err := fakeClient.CoreV1().Services(svc.Namespace).Create(context.Background(), svc, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewIstioGatewaySource(
		t.Context(),
		fakeClient,
		istiofake.NewSimpleClientset(),
		"",
		"",
		"",
		false,
		false)
	require.NoError(t, err)
	gwSource, ok := src.(*gatewaySource)
	require.True(t, ok)

	rService, err := gwSource.serviceInformer.Lister().Services(svc.Namespace).Get(svc.Name)
	require.NoError(t, err)

	assert.Equal(t, "fake-service", rService.Name)
	assert.Empty(t, rService.Labels)
	assert.Empty(t, rService.Annotations)
	assert.Empty(t, rService.UID)
	assert.NotEmpty(t, rService.Status.LoadBalancer)
	assert.Empty(t, rService.Status.Conditions)
	assert.Equal(t, map[string]string{
		"selector":  "one",
		"selector2": "two",
		"selector3": "three",
	}, rService.Spec.Selector)
}

func TestSingleGatewayMultipleServicesPointingToSameLoadBalancer(t *testing.T) {
	fakeKubeClient := fake.NewClientset()
	fakeIstioClient := istiofake.NewSimpleClientset()

	gw := &networkingv1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "argocd",
			Namespace: "argocd",
		},
		Spec: istionetworking.Gateway{
			Servers: []*istionetworking.Server{
				{
					Hosts: []string{"example.org"},
					Tls: &istionetworking.ServerTLSSettings{
						HttpsRedirect: true,
					},
				},
				{
					Hosts: []string{"example.org"},
					Tls: &istionetworking.ServerTLSSettings{
						ServerCertificate: IstioGatewayIngressSource,
						Mode:              istionetworking.ServerTLSSettings_SIMPLE,
					},
				},
			},
			Selector: map[string]string{
				"istio": "ingressgateway",
			},
		},
	}

	services := []*v1.Service{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "istio-ingressgateway",
				Namespace: "default",
				Labels: map[string]string{
					"app":   "istio-ingressgateway",
					"istio": "ingressgateway",
				},
			},
			Spec: v1.ServiceSpec{
				Type:                  v1.ServiceTypeLoadBalancer,
				ClusterIP:             "10.118.223.3",
				ClusterIPs:            []string{"10.118.223.3"},
				ExternalTrafficPolicy: v1.ServiceExternalTrafficPolicyCluster,
				IPFamilies:            []v1.IPFamily{v1.IPv4Protocol},
				IPFamilyPolicy:        testutils.ToPtr(v1.IPFamilyPolicySingleStack),
				Ports: []v1.ServicePort{
					{
						Name:       "http2",
						Port:       80,
						Protocol:   v1.ProtocolTCP,
						TargetPort: intstr.FromInt32(8080),
						NodePort:   30127,
					},
				},
				Selector: map[string]string{
					"app":   "istio-ingressgateway",
					"istio": "ingressgateway",
				},
				SessionAffinity: v1.ServiceAffinityNone,
			},
			Status: v1.ServiceStatus{
				LoadBalancer: v1.LoadBalancerStatus{
					Ingress: []v1.LoadBalancerIngress{
						{
							IP:     "34.66.66.77",
							IPMode: testutils.ToPtr(v1.LoadBalancerIPModeVIP),
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "istio-ingressgatewayudp",
				Namespace: "default",
				Labels: map[string]string{
					"app":   "istio-ingressgatewayudp",
					"istio": "ingressgateway",
				},
			},
			Spec: v1.ServiceSpec{
				Type:                  v1.ServiceTypeLoadBalancer,
				ClusterIP:             "10.118.220.130",
				ClusterIPs:            []string{"10.118.220.130"},
				ExternalTrafficPolicy: v1.ServiceExternalTrafficPolicyCluster,
				IPFamilies:            []v1.IPFamily{v1.IPv4Protocol},
				IPFamilyPolicy:        testutils.ToPtr(v1.IPFamilyPolicySingleStack),
				Ports: []v1.ServicePort{
					{
						Name:       "upd-dns",
						Port:       53,
						Protocol:   v1.ProtocolUDP,
						TargetPort: intstr.FromInt32(5353),
						NodePort:   30873,
					},
				},
				Selector: map[string]string{
					"app":   "istio-ingressgatewayudp",
					"istio": "ingressgateway",
				},
				SessionAffinity: v1.ServiceAffinityNone,
			},
			Status: v1.ServiceStatus{
				LoadBalancer: v1.LoadBalancerStatus{
					Ingress: []v1.LoadBalancerIngress{
						{
							IP:     "34.66.66.77",
							IPMode: testutils.ToPtr(v1.LoadBalancerIPModeVIP),
						},
					},
				},
			},
		},
	}

	assert.NotNil(t, services)

	for _, svc := range services {
		_, err := fakeKubeClient.CoreV1().Services(svc.Namespace).Create(t.Context(), svc, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	_, err := fakeIstioClient.NetworkingV1beta1().Gateways(gw.Namespace).Create(t.Context(), gw, metav1.CreateOptions{})
	require.NoError(t, err)

	src, err := NewIstioGatewaySource(
		t.Context(),
		fakeKubeClient,
		fakeIstioClient,
		"",
		"",
		"",
		false,
		false,
	)
	require.NoError(t, err)
	require.NotNil(t, src)

	got, err := src.Endpoints(t.Context())
	require.NoError(t, err)

	validateEndpoints(t, got, []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.org", endpoint.RecordTypeA, "34.66.66.77").WithLabel(endpoint.ResourceLabelKey, "gateway/argocd/argocd"),
	})
}

// gateway specific helper functions
func newTestGatewaySource(loadBalancerList []fakeIngressGatewayService, ingressList []fakeIngress) (*gatewaySource, error) {
	fakeKubernetesClient := fake.NewClientset()
	fakeIstioClient := istiofake.NewSimpleClientset()

	for _, lb := range loadBalancerList {
		service := lb.Service()
		_, err := fakeKubernetesClient.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	for _, ing := range ingressList {
		ingress := ing.Ingress()
		_, err := fakeKubernetesClient.NetworkingV1().Ingresses(ingress.Namespace).Create(context.Background(), ingress, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}

	src, err := NewIstioGatewaySource(
		context.TODO(),
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

	gwsrc, ok := src.(*gatewaySource)
	if !ok {
		return nil, errors.New("underlying source type was not gateway")
	}

	return gwsrc, nil
}

type fakeIngressGatewayService struct {
	ips         []string
	hostnames   []string
	namespace   string
	name        string
	selector    map[string]string
	externalIPs []string
}

func (ig fakeIngressGatewayService) Service() *v1.Service {
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
		Spec: v1.ServiceSpec{
			Selector:    ig.selector,
			ExternalIPs: ig.externalIPs,
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

type fakeGatewayConfig struct {
	namespace   string
	name        string
	annotations map[string]string
	dnsnames    [][]string
	selector    map[string]string
}

func (c fakeGatewayConfig) Config() *networkingv1beta1.Gateway {
	gw := &networkingv1beta1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:        c.name,
			Namespace:   c.namespace,
			Annotations: c.annotations,
		},
		Spec: istionetworking.Gateway{
			Servers:  nil,
			Selector: c.selector,
		},
	}

	var servers []*istionetworking.Server
	for _, dnsnames := range c.dnsnames {
		servers = append(servers, &istionetworking.Server{
			Hosts: dnsnames,
		})
	}

	gw.Spec.Servers = servers

	return gw
}
