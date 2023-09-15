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
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"istio.io/api/meta/v1alpha1"
	istionetworking "istio.io/api/networking/v1alpha3"
	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	istiofake "istio.io/client-go/pkg/clientset/versioned/fake"
	fakenetworking3 "istio.io/client-go/pkg/clientset/versioned/typed/networking/v1alpha3/fake"
	v1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	k8sclienttesting "k8s.io/client-go/testing"

	"sigs.k8s.io/external-dns/endpoint"
)

// This is a compile-time validation that istioVirtualServiceSource is a Source.
var _ Source = &virtualServiceSource{}

type VirtualServiceSuite struct {
	suite.Suite
	source     Source
	lbServices []*v1.Service
	ingresses  []*networkv1.Ingress
	gwconfig   *networkingv1alpha3.Gateway
	vsconfig   *networkingv1alpha3.VirtualService
}

func (suite *VirtualServiceSuite) SetupTest() {
	fakeKubernetesClient := fake.NewSimpleClientset()
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
			namespace: "istio-system",
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

	suite.gwconfig = (fakeGatewayConfig{
		name:      "foo-gateway-with-targets",
		namespace: "istio-system",
		dnsnames:  [][]string{{"*"}},
	}).Config()
	_, err = fakeIstioClient.NetworkingV1alpha3().Gateways(suite.gwconfig.Namespace).Create(context.Background(), suite.gwconfig, metav1.CreateOptions{})
	suite.NoError(err, "should succeed")

	suite.vsconfig = (fakeVirtualServiceConfig{
		name:      "foo-virtualservice",
		namespace: "istio-other",
		gateways:  []string{"istio-system/foo-gateway-with-targets"},
		dnsnames:  []string{"foo"},
	}).Config()
	_, err = fakeIstioClient.NetworkingV1alpha3().VirtualServices(suite.vsconfig.Namespace).Create(context.Background(), suite.vsconfig, metav1.CreateOptions{})
	suite.NoError(err, "should succeed")

	suite.source, err = NewIstioVirtualServiceSource(
		context.TODO(),
		fakeKubernetesClient,
		fakeIstioClient,
		"",
		"",
		"{{.Name}}",
		false,
		false,
	)
	suite.NoError(err, "should initialize virtualservice source")
}

func (suite *VirtualServiceSuite) TestResourceLabelIsSet() {
	endpoints, err := suite.source.Endpoints(context.Background())
	suite.NoError(err, "should succeed")
	suite.Equal(len(endpoints), 2, "should return the correct number of endpoints")
	for _, ep := range endpoints {
		suite.Equal("virtualservice/istio-other/foo-virtualservice", ep.Labels[endpoint.ResourceLabelKey], "should set correct resource label")
	}
}

func TestVirtualService(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(VirtualServiceSuite))
	t.Run("virtualServiceBindsToGateway", testVirtualServiceBindsToGateway)
	t.Run("endpointsFromVirtualServiceConfig", testEndpointsFromVirtualServiceConfig)
	t.Run("Endpoints", testVirtualServiceEndpoints)
	t.Run("gatewaySelectorMatchesService", testGatewaySelectorMatchesService)
}

func TestNewIstioVirtualServiceSource(t *testing.T) {
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
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			_, err := NewIstioVirtualServiceSource(
				context.TODO(),
				fake.NewSimpleClientset(),
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

func testVirtualServiceBindsToGateway(t *testing.T) {
	for _, ti := range []struct {
		title    string
		gwconfig fakeGatewayConfig
		vsconfig fakeVirtualServiceConfig
		vsHost   string
		expected bool
	}{
		{
			title: "matching host *",
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{},
			vsHost:   "foo.bar",
			expected: true,
		},
		{
			title: "matching host *.<domain>",
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{{"*.foo.bar"}},
			},
			vsconfig: fakeVirtualServiceConfig{},
			vsHost:   "baz.foo.bar",
			expected: true,
		},
		{
			title: "not matching host *.<domain>",
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{{"*.foo.bar"}},
			},
			vsconfig: fakeVirtualServiceConfig{},
			vsHost:   "foo.bar",
			expected: false,
		},
		{
			title: "not matching host *.<domain>",
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{{"*.foo.bar"}},
			},
			vsconfig: fakeVirtualServiceConfig{},
			vsHost:   "bazfoo.bar",
			expected: false,
		},
		{
			title: "not matching host *.<domain>",
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{{"*.foo.bar"}},
			},
			vsconfig: fakeVirtualServiceConfig{},
			vsHost:   "*foo.bar",
			expected: false,
		},
		{
			title: "matching host */*",
			gwconfig: fakeGatewayConfig{
				dnsnames: [][]string{{"*/*"}},
			},
			vsconfig: fakeVirtualServiceConfig{},
			vsHost:   "foo.bar",
			expected: true,
		},
		{
			title: "matching host <namespace>/*",
			gwconfig: fakeGatewayConfig{
				namespace: "istio-system",
				dnsnames:  [][]string{{"myns/*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				namespace: "myns",
			},
			vsHost:   "foo.bar",
			expected: true,
		},
		{
			title: "matching host ./*",
			gwconfig: fakeGatewayConfig{
				namespace: "istio-system",
				dnsnames:  [][]string{{"./*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				namespace: "istio-system",
			},
			vsHost:   "foo.bar",
			expected: true,
		},
		{
			title: "not matching host ./*",
			gwconfig: fakeGatewayConfig{
				namespace: "istio-system",
				dnsnames:  [][]string{{"./*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				namespace: "myns",
			},
			vsHost:   "foo.bar",
			expected: false,
		},
		{
			title: "not matching host <namespace>/*",
			gwconfig: fakeGatewayConfig{
				namespace: "istio-system",
				dnsnames:  [][]string{{"myns/*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				namespace: "otherns",
			},
			vsHost:   "foo.bar",
			expected: false,
		},
		{
			title: "not matching host <namespace>/*",
			gwconfig: fakeGatewayConfig{
				namespace: "istio-system",
				dnsnames:  [][]string{{"myns/*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				namespace: "otherns",
			},
			vsHost:   "foo.bar",
			expected: false,
		},
		{
			title: "matching exportTo *",
			gwconfig: fakeGatewayConfig{
				namespace: "istio-system",
				dnsnames:  [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				namespace: "otherns",
				exportTo:  "*",
			},
			vsHost:   "foo.bar",
			expected: true,
		},
		{
			title: "matching exportTo <namespace>",
			gwconfig: fakeGatewayConfig{
				namespace: "istio-system",
				dnsnames:  [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				namespace: "otherns",
				exportTo:  "istio-system",
			},
			vsHost:   "foo.bar",
			expected: true,
		},
		{
			title: "not matching exportTo <namespace>",
			gwconfig: fakeGatewayConfig{
				namespace: "istio-system",
				dnsnames:  [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				namespace: "otherns",
				exportTo:  "myns",
			},
			vsHost:   "foo.bar",
			expected: false,
		},
		{
			title: "not matching exportTo .",
			gwconfig: fakeGatewayConfig{
				namespace: "istio-system",
				dnsnames:  [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				namespace: "otherns",
				exportTo:  ".",
			},
			vsHost:   "foo.bar",
			expected: false,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			vsconfig := ti.vsconfig.Config()
			gwconfig := ti.gwconfig.Config()
			require.Equal(t, ti.expected, virtualServiceBindsToGateway(vsconfig, gwconfig, ti.vsHost))
		})
	}
}

func testEndpointsFromVirtualServiceConfig(t *testing.T) {
	t.Parallel()

	for _, ti := range []struct {
		title      string
		lbServices []fakeIngressGatewayService
		ingresses  []fakeIngress
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
				name:     "mygw",
				dnsnames: [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				gateways: []string{"mygw"},
				dnsnames: []string{"foo.bar"},
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
			gwconfig: fakeGatewayConfig{
				name:     "mygw",
				dnsnames: [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				gateways: []string{"mygw"},
				dnsnames: []string{"foo.bar"},
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
			gwconfig: fakeGatewayConfig{
				name:     "mygw",
				dnsnames: [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				gateways: []string{"mygw"},
				dnsnames: []string{"foo.bar"},
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
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			gwconfig: fakeGatewayConfig{
				name:     "mygw",
				dnsnames: [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				gateways: []string{"mygw"},
				dnsnames: []string{},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "no rule.gateway",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8", "127.0.0.1"},
					hostnames: []string{"elb.com", "alb.com"},
				},
			},
			gwconfig: fakeGatewayConfig{
				name:     "mygw",
				dnsnames: [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				gateways: []string{},
				dnsnames: []string{"foo.bar"},
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
				name:     "mygw",
				dnsnames: [][]string{{"*"}},
			},
			vsconfig: fakeVirtualServiceConfig{
				gateways: []string{},
				dnsnames: []string{"foo.bar"},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "matching selectors for service and gateway",
			lbServices: []fakeIngressGatewayService{
				{
					name: "service1",
					selector: map[string]string{
						"app": "myservice",
					},
					hostnames: []string{"elb.com", "alb.com"},
				},
				{
					name: "service2",
					selector: map[string]string{
						"app": "otherservice",
					},
					ips: []string{"8.8.8.8", "127.0.0.1"},
				},
			},
			gwconfig: fakeGatewayConfig{
				name:     "mygw",
				dnsnames: [][]string{{"*"}},
				selector: map[string]string{
					"app": "myservice",
				},
			},
			vsconfig: fakeVirtualServiceConfig{
				gateways: []string{"mygw"},
				dnsnames: []string{"foo.bar"},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"elb.com", "alb.com"},
				},
			},
		},
		{
			title: "ingress gateway annotation same namespace",
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					hostnames: []string{"alb.com", "elb.com"},
				},
				{
					name: "ingress2",
					ips:  []string{"127.0.0.1", "8.8.8.8"},
				},
			},
			gwconfig: fakeGatewayConfig{
				name:     "mygw",
				dnsnames: [][]string{{"*"}},
				annotations: map[string]string{
					IstioGatewayIngressSource: "ingress1",
				},
			},
			vsconfig: fakeVirtualServiceConfig{
				gateways: []string{"mygw"},
				dnsnames: []string{"foo.bar"},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"alb.com", "elb.com"},
				},
			},
		},
		{
			title: "ingress gateway annotation separate namespace",
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					namespace: "ingress",
					hostnames: []string{"alb.com", "elb.com"},
				},
				{
					name:      "ingress2",
					namespace: "ingress",
					ips:       []string{"127.0.0.1", "8.8.8.8"},
				},
			},
			gwconfig: fakeGatewayConfig{
				name:     "mygw",
				dnsnames: [][]string{{"*"}},
				annotations: map[string]string{
					IstioGatewayIngressSource: "ingress/ingress2",
				},
			},
			vsconfig: fakeVirtualServiceConfig{
				gateways: []string{"mygw"},
				dnsnames: []string{"foo.bar"},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"127.0.0.1", "8.8.8.8"},
				},
			},
		},
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			if source, err := newTestVirtualServiceSource(ti.lbServices, ti.ingresses, []fakeGatewayConfig{ti.gwconfig}); err != nil {
				require.NoError(t, err)
			} else if endpoints, err := source.endpointsFromVirtualService(context.Background(), ti.vsconfig.Config()); err != nil {
				require.NoError(t, err)
			} else {
				validateEndpoints(t, endpoints, ti.expected)
			}
		})
	}
}

func testVirtualServiceEndpoints(t *testing.T) {
	t.Parallel()

	namespace := "testing"
	for _, ti := range []struct {
		title                    string
		targetNamespace          string
		annotationFilter         string
		lbServices               []fakeIngressGatewayService
		ingresses                []fakeIngress
		gwConfigs                []fakeGatewayConfig
		vsConfigs                []fakeVirtualServiceConfig
		expected                 []*endpoint.Endpoint
		expectError              bool
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		ignoreHostnameAnnotation bool
	}{
		{
			title: "two simple virtualservices with one gateway each, one ingressgateway loadbalancer service",
			lbServices: []fakeIngressGatewayService{
				{
					namespace: namespace,
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
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: namespace,
					gateways:  []string{"fake2"},
					dnsnames:  []string{"new.org"},
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
			title: "two simple virtualservices with one gateway each, one ingress",
			lbServices: []fakeIngressGatewayService{
				{
					namespace: namespace,
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
				},
			},
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					namespace: namespace,
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"example.org"}},
					annotations: map[string]string{
						IstioGatewayIngressSource: "ingress1",
					},
				},
				{
					name:      "fake2",
					namespace: namespace,
					dnsnames:  [][]string{{"new.org"}},
					annotations: map[string]string{
						IstioGatewayIngressSource: "ingress1",
					},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: namespace,
					gateways:  []string{"fake2"},
					dnsnames:  []string{"new.org"},
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
			title: "one virtualservice with two gateways, one ingressgateway loadbalancer service",
			lbServices: []fakeIngressGatewayService{
				{
					namespace: namespace,
					ips:       []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "gw1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
				{
					name:      "gw2",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs",
					namespace: namespace,
					gateways:  []string{"gw1", "gw2"},
					dnsnames:  []string{"example.org"},
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
			title: "two simple virtualservices on different namespaces with the same target gateway, one ingressgateway loadbalancer service",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
					namespace: "istio-system",
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "istio-system",
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: "testing1",
					gateways:  []string{"istio-system/fake1"},
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: "testing2",
					gateways:  []string{"istio-system/fake1"},
					dnsnames:  []string{"new.org"},
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
			title:           "two simple virtualservices with one gateway on different namespaces and a target namespace, one ingressgateway loadbalancer service",
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
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: "testing1",
					gateways:  []string{"testing1/fake1"},
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: "testing2",
					gateways:  []string{"testing1/fake1"},
					dnsnames:  []string{"new.org"},
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
			title:           "two simple virtualservices with one gateway on different namespaces and a target namespace, one ingress",
			targetNamespace: "testing1",
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
					namespace: "testing1",
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "testing1",
					dnsnames:  [][]string{{"*"}},
					annotations: map[string]string{
						IstioGatewayIngressSource: "ingress1",
					},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: "testing1",
					gateways:  []string{"testing1/fake1"},
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: "testing2",
					gateways:  []string{"testing1/fake1"},
					dnsnames:  []string{"new.org"},
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
			annotationFilter: "kubernetes.io/virtualservice.class in (alb, nginx)",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/virtualservice.class": "nginx",
					},
					gateways: []string{"fake1"},
					dnsnames: []string{"example.org"},
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
			annotationFilter: "kubernetes.io/gateway.class in (alb, nginx)",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					annotations: map[string]string{
						"kubernetes.io/virtualservice.class": "tectonic",
					},
					gateways: []string{"fake1"},
					dnsnames: []string{"example.org"},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:            "invalid annotation filter expression",
			annotationFilter: "kubernetes.io/gateway.name in (a b)",
			expected:         []*endpoint.Endpoint{},
			expectError:      true,
		},
		{
			title: "gateway ingress annotation; ingress not found",
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					namespace: namespace,
					ips:       []string{"8.8.8.8"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
					annotations: map[string]string{
						IstioGatewayIngressSource: "ingress2",
					},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{"example.org"},
				},
			},
			expected:    []*endpoint.Endpoint{},
			expectError: true,
		},
		{
			title: "our controller type is dns-controller",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: controllerAnnotationValue,
					},
					gateways: []string{"fake1"},
					dnsnames: []string{"example.org"},
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
			title: "different controller types are ignored",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					annotations: map[string]string{
						controllerAnnotationKey: "some-other-tool",
					},
					gateways: []string{"fake1"},
					dnsnames: []string{"example.org"},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "template for virtualservice if host is missing",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"elb.com"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{""},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "vs1.ext-dns.test.com",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
				},
				{
					DNSName:    "vs1.ext-dns.test.com",
					RecordType: endpoint.RecordTypeCNAME,
					Targets:    endpoint.Targets{"elb.com"},
				},
			},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title: "multiple FQDN template hostnames",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{""},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "vs1.ext-dns.test.com",
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordType: endpoint.RecordTypeA,
				},
				{
					DNSName:    "vs1.ext-dna.test.com",
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordType: endpoint.RecordTypeA,
				},
			},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com, {{.Name}}.ext-dna.test.com",
		},
		{
			title: "multiple FQDN template hostnames with restricted gw.hosts",
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
					},
					dnsnames: [][]string{{"*.org", "*.ext-dns.test.com"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "vs1.ext-dns.test.com",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "vs2.ext-dns.test.com",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
			fqdnTemplate:             "{{.Name}}.ext-dns.test.com, {{.Name}}.ext-dna.test.com",
			combineFQDNAndAnnotation: true,
		},
		{
			title: "virtualservice with target annotation",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					namespace: namespace,
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					annotations: map[string]string{
						targetAnnotationKey: "virtualservice-target.com",
					},
					dnsnames: []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: namespace,
					gateways:  []string{"fake1"},
					annotations: map[string]string{
						targetAnnotationKey: "virtualservice-target.com",
					},
					dnsnames: []string{"example2.org"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"virtualservice-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "example2.org",
					Targets:    endpoint.Targets{"virtualservice-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
		},
		{
			title: "virtualservice; gateway with target annotation",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
					},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{"example2.org"},
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
			},
		},
		{
			title: "virtualservice; gateway with target and ingress annotation",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					namespace: namespace,
				},
			},
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					ips:       []string{"1.1.1.1"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
					annotations: map[string]string{
						targetAnnotationKey:       "gateway-target.com",
						IstioGatewayIngressSource: "ingress1",
					},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{"example2.org"},
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
			},
		},
		{
			title: "virtualservice with hostname annotation",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"1.2.3.4"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
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
			title: "virtualservice with hostname annotation having multiple hostnames, restricted by gw.hosts",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"1.2.3.4"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*.bar.com"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					annotations: map[string]string{
						hostnameAnnotationKey: "foo.bar.com, another-dns-through-hostname.com",
					},
					dnsnames: []string{"baz.bar.org"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
				},
			},
		},
		{
			title: "virtualservices with annotation and custom TTL",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					annotations: map[string]string{
						ttlAnnotationKey: "6",
					},
					dnsnames: []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: namespace,
					gateways:  []string{"fake1"},
					annotations: map[string]string{
						ttlAnnotationKey: "1",
					},
					dnsnames: []string{"example2.org"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordTTL:  endpoint.TTL(6),
				},
				{
					DNSName:    "example2.org",
					RecordType: endpoint.RecordTypeA,
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordTTL:  endpoint.TTL(1),
				},
			},
		},
		{
			title: "template for virtualservice; gateway with target annotation",
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
					},
					dnsnames: [][]string{{"*"}},
				},
				{
					name:      "fake2",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "gateway-target.com",
					},
					dnsnames: [][]string{{"*"}},
				},
				{
					name:      "fake3",
					namespace: namespace,
					annotations: map[string]string{
						targetAnnotationKey: "1.2.3.4",
					},
					dnsnames: [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					dnsnames:  []string{},
				},
				{
					name:      "vs2",
					namespace: namespace,
					gateways:  []string{"fake2"},
					dnsnames:  []string{},
				},
				{
					name:      "vs3",
					namespace: namespace,
					gateways:  []string{"fake3"},
					dnsnames:  []string{},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "vs1.ext-dns.test.com",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "vs2.ext-dns.test.com",
					Targets:    endpoint.Targets{"gateway-target.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "vs3.ext-dns.test.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
				},
			},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
		{
			title: "ignore hostname annotations",
			lbServices: []fakeIngressGatewayService{
				{
					ips:       []string{"8.8.8.8"},
					hostnames: []string{"lb.com"},
					namespace: namespace,
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: namespace,
					dnsnames:  [][]string{{"*"}},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: namespace,
					gateways:  []string{"fake1"},
					annotations: map[string]string{
						hostnameAnnotationKey: "ignore.me",
					},
					dnsnames: []string{"example.org"},
				},
				{
					name:      "vs2",
					namespace: namespace,
					gateways:  []string{"fake1"},
					annotations: map[string]string{
						hostnameAnnotationKey: "ignore.me.too",
					},
					dnsnames: []string{"new.org"},
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
			title: "complex setup with multiple gateways and multiple vs.hosts only matching some of the gateway",
			lbServices: []fakeIngressGatewayService{
				{
					name: "svc1",
					selector: map[string]string{
						"app": "igw1",
					},
					hostnames: []string{"target1.com"},
					namespace: "istio-system",
				},
				{
					name: "svc2",
					selector: map[string]string{
						"app": "igw2",
					},
					hostnames: []string{"target2.com"},
					namespace: "testing1",
				},
				{
					name: "svc3",
					selector: map[string]string{
						"app": "igw3",
					},
					hostnames: []string{"target3.com"},
					namespace: "testing2",
				},
			},
			ingresses: []fakeIngress{
				{
					name:      "ingress1",
					namespace: "testing1",
					hostnames: []string{"target4.com"},
				},
				{
					name:      "ingress3",
					namespace: "testing3",
					hostnames: []string{"target5.com"},
				},
			},
			gwConfigs: []fakeGatewayConfig{
				{
					name:      "fake1",
					namespace: "istio-system",
					dnsnames:  [][]string{{"*"}},
					selector: map[string]string{
						"app": "igw1",
					},
				},
				{
					name:      "fake2",
					namespace: "testing1",
					dnsnames:  [][]string{{"*.baz.com"}, {"*.bar.com"}},
					selector: map[string]string{
						"app": "igw2",
					},
				},
				{
					name:      "fake3",
					namespace: "testing2",
					dnsnames:  [][]string{{"*.bax.com", "*.bar.com"}},
					selector: map[string]string{
						"app": "igw3",
					},
				},
				{
					name:      "fake4",
					namespace: "testing3",
					dnsnames:  [][]string{{"*.bax.com", "*.bar.com"}},
					selector: map[string]string{
						"app": "igw4",
					},
					annotations: map[string]string{
						IstioGatewayIngressSource: "testing1/ingress1",
					},
				},
			},
			vsConfigs: []fakeVirtualServiceConfig{
				{
					name:      "vs1",
					namespace: "testing3",
					gateways:  []string{"istio-system/fake1", "testing1/fake2"},
					dnsnames:  []string{"somedomain.com", "foo.bar.com"},
				},
				{
					name:      "vs2",
					namespace: "testing2",
					gateways:  []string{"testing1/fake2", "fake3"},
					dnsnames:  []string{"hello.bar.com", "hello.bax.com", "hello.bak.com"},
				},
				{
					name:      "vs3",
					namespace: "testing1",
					gateways:  []string{"istio-system/fake1", "testing2/fake3"},
					dnsnames:  []string{"world.bax.com", "world.bak.com"},
				},
				{
					name:      "vs4",
					namespace: "testing1",
					gateways:  []string{"istio-system/fake1", "testing3/fake4"},
					dnsnames:  []string{"foo.bax.com", "foo.bak.com"},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "somedomain.com",
					Targets:    endpoint.Targets{"target1.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "foo.bar.com",
					Targets:    endpoint.Targets{"target1.com", "target2.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "hello.bar.com",
					Targets:    endpoint.Targets{"target2.com", "target3.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "hello.bax.com",
					Targets:    endpoint.Targets{"target3.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "world.bak.com",
					Targets:    endpoint.Targets{"target1.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "world.bax.com",
					Targets:    endpoint.Targets{"target1.com", "target3.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "foo.bak.com",
					Targets:    endpoint.Targets{"target1.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "foo.bax.com",
					Targets:    endpoint.Targets{"target1.com", "target4.com"},
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
		},
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			var gateways []*networkingv1alpha3.Gateway
			var virtualservices []*networkingv1alpha3.VirtualService

			for _, gwItem := range ti.gwConfigs {
				gateways = append(gateways, gwItem.Config())
			}
			for _, vsItem := range ti.vsConfigs {
				virtualservices = append(virtualservices, vsItem.Config())
			}

			fakeKubernetesClient := fake.NewSimpleClientset()

			for _, lb := range ti.lbServices {
				service := lb.Service()
				_, err := fakeKubernetesClient.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			for _, ing := range ti.ingresses {
				ingress := ing.Ingress()
				_, err := fakeKubernetesClient.NetworkingV1().Ingresses(ingress.Namespace).Create(context.Background(), ingress, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			fakeIstioClient := istiofake.NewSimpleClientset()

			for _, gateway := range gateways {
				_, err := fakeIstioClient.NetworkingV1alpha3().Gateways(gateway.Namespace).Create(context.Background(), gateway, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			for _, virtualservice := range virtualservices {
				_, err := fakeIstioClient.NetworkingV1alpha3().VirtualServices(virtualservice.Namespace).Create(context.Background(), virtualservice, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			virtualServiceSource, err := NewIstioVirtualServiceSource(
				context.TODO(),
				fakeKubernetesClient,
				fakeIstioClient,
				ti.targetNamespace,
				ti.annotationFilter,
				ti.fqdnTemplate,
				ti.combineFQDNAndAnnotation,
				ti.ignoreHostnameAnnotation,
			)
			require.NoError(t, err)

			res, err := virtualServiceSource.Endpoints(context.Background())
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			validateEndpoints(t, res, ti.expected)
		})
	}
}

func testGatewaySelectorMatchesService(t *testing.T) {
	for _, ti := range []struct {
		title      string
		gwSelector map[string]string
		lbSelector map[string]string
		expected   bool
	}{
		{
			title:      "gw selector matches lb selector",
			gwSelector: map[string]string{"istio": "ingressgateway"},
			lbSelector: map[string]string{"istio": "ingressgateway"},
			expected:   true,
		},
		{
			title:      "gw selector matches lb selector partially",
			gwSelector: map[string]string{"istio": "ingressgateway"},
			lbSelector: map[string]string{"release": "istio", "istio": "ingressgateway"},
			expected:   true,
		},
		{
			title:      "gw selector does not match lb selector",
			gwSelector: map[string]string{"app": "mytest"},
			lbSelector: map[string]string{"istio": "ingressgateway"},
			expected:   false,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			require.Equal(t, ti.expected, gatewaySelectorMatchesServiceSelector(ti.gwSelector, ti.lbSelector))
		})
	}
}

func newTestVirtualServiceSource(loadBalancerList []fakeIngressGatewayService, ingressList []fakeIngress, gwList []fakeGatewayConfig) (*virtualServiceSource, error) {
	fakeKubernetesClient := fake.NewSimpleClientset()
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

	for _, gw := range gwList {
		gwObj := gw.Config()
		// use create instead of add
		// https://github.com/kubernetes/client-go/blob/92512ee2b8cf6696e9909245624175b7f0c971d9/testing/fixture.go#LL336C3-L336C52
		_, err := fakeIstioClient.NetworkingV1alpha3().Gateways(gw.namespace).Create(context.Background(), gwObj, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}

	src, err := NewIstioVirtualServiceSource(
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

	vssrc, ok := src.(*virtualServiceSource)
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
	exportTo    string
}

func (c fakeVirtualServiceConfig) Config() *networkingv1alpha3.VirtualService {
	vs := istionetworking.VirtualService{
		Gateways: c.gateways,
		Hosts:    c.dnsnames,
	}
	if c.exportTo != "" {
		vs.ExportTo = []string{c.exportTo}
	}

	return &networkingv1alpha3.VirtualService{
		ObjectMeta: metav1.ObjectMeta{
			Name:        c.name,
			Namespace:   c.namespace,
			Annotations: c.annotations,
		},
		Spec: *vs.DeepCopy(),
	}
}

func TestVirtualServiceSourceGetGateway(t *testing.T) {
	type fields struct {
		virtualServiceSource *virtualServiceSource
	}
	type args struct {
		ctx            context.Context
		gatewayStr     string
		virtualService *networkingv1alpha3.VirtualService
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           *networkingv1alpha3.Gateway
		expectedErrStr string
	}{
		{name: "EmptyGateway", fields: fields{
			virtualServiceSource: func() *virtualServiceSource { vs, _ := newTestVirtualServiceSource(nil, nil, nil); return vs }(),
		}, args: args{
			ctx:            context.TODO(),
			gatewayStr:     "",
			virtualService: nil,
		}, want: nil, expectedErrStr: ""},
		{name: "MeshGateway", fields: fields{
			virtualServiceSource: func() *virtualServiceSource { vs, _ := newTestVirtualServiceSource(nil, nil, nil); return vs }(),
		}, args: args{
			ctx:            context.TODO(),
			gatewayStr:     IstioMeshGateway,
			virtualService: nil,
		}, want: nil, expectedErrStr: ""},
		{name: "MissingGateway", fields: fields{
			virtualServiceSource: func() *virtualServiceSource { vs, _ := newTestVirtualServiceSource(nil, nil, nil); return vs }(),
		}, args: args{
			ctx:        context.TODO(),
			gatewayStr: "doesnt/exist",
			virtualService: &networkingv1alpha3.VirtualService{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{Name: "exist", Namespace: "doesnt"},
				Spec:       istionetworking.VirtualService{},
				Status:     v1alpha1.IstioStatus{},
			},
		}, want: nil, expectedErrStr: ""},
		{name: "InvalidGatewayStr", fields: fields{
			virtualServiceSource: func() *virtualServiceSource { vs, _ := newTestVirtualServiceSource(nil, nil, nil); return vs }(),
		}, args: args{
			ctx:            context.TODO(),
			gatewayStr:     "1/2/3/",
			virtualService: &networkingv1alpha3.VirtualService{},
		}, want: nil, expectedErrStr: "invalid gateway name (name or namespace/name) found '1/2/3/'"},
		{name: "ExistingGateway", fields: fields{
			virtualServiceSource: func() *virtualServiceSource {
				vs, _ := newTestVirtualServiceSource(nil, nil, []fakeGatewayConfig{{
					namespace: "bar",
					name:      "foo",
				}})
				return vs
			}(),
		}, args: args{
			ctx:        context.TODO(),
			gatewayStr: "bar/foo",
			virtualService: &networkingv1alpha3.VirtualService{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec:       istionetworking.VirtualService{},
				Status:     v1alpha1.IstioStatus{},
			},
		}, want: &networkingv1alpha3.Gateway{
			TypeMeta:   metav1.TypeMeta{},
			ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "bar"},
			Spec:       istionetworking.Gateway{},
			Status:     v1alpha1.IstioStatus{},
		}, expectedErrStr: ""},
		{name: "ErrorGettingGateway", fields: fields{
			virtualServiceSource: func() *virtualServiceSource {
				istioFake := istiofake.NewSimpleClientset()
				istioFake.NetworkingV1alpha3().(*fakenetworking3.FakeNetworkingV1alpha3).PrependReactor("get", "gateways", func(action k8sclienttesting.Action) (handled bool, ret runtime.Object, err error) {
					return true, &networkingv1alpha3.Gateway{}, fmt.Errorf("error getting gateway")
				})
				vs, _ := NewIstioVirtualServiceSource(
					context.TODO(),
					fake.NewSimpleClientset(),
					istioFake,
					"",
					"",
					"{{.Name}}",
					false,
					false,
				)
				return vs.(*virtualServiceSource)
			}(),
		}, args: args{
			ctx:        context.TODO(),
			gatewayStr: "foo/bar",
			virtualService: &networkingv1alpha3.VirtualService{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{Name: "gateway", Namespace: "error"},
				Spec:       istionetworking.VirtualService{},
				Status:     v1alpha1.IstioStatus{},
			},
		}, want: nil, expectedErrStr: "error getting gateway"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.virtualServiceSource.getGateway(tt.args.ctx, tt.args.gatewayStr, tt.args.virtualService)
			if tt.expectedErrStr != "" {
				assert.EqualError(t, err, tt.expectedErrStr, fmt.Sprintf("getGateway(%v, %v, %v)", tt.args.ctx, tt.args.gatewayStr, tt.args.virtualService))
				return
			} else {
				require.NoError(t, err)
			}
			if tt.want != nil && got != nil {
				tt.want.Spec.ProtoReflect()
				tt.want.Status.ProtoReflect()
				assert.Equalf(t, tt.want, got, "getGateway(%v, %v, %v)", tt.args.ctx, tt.args.gatewayStr, tt.args.virtualService)
			} else {
				assert.Equalf(t, tt.want, got, "getGateway(%v, %v, %v)", tt.args.ctx, tt.args.gatewayStr, tt.args.virtualService)
			}
		})
	}
}
