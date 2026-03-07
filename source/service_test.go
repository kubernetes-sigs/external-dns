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
	"fmt"
	"maps"
	"math/rand"
	"net"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/intstr"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers/fakes"
)

type ServiceSuite struct {
	suite.Suite
	sc             Source
	fooWithTargets *v1.Service
}

func (suite *ServiceSuite) SetupTest() {
	fakeClient := fake.NewClientset()

	suite.fooWithTargets = &v1.Service{
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   "default",
			Name:        "foo-with-targets",
			Annotations: map[string]string{},
		},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{
				Ingress: []v1.LoadBalancerIngress{
					{IP: "8.8.8.8"},
					{Hostname: "foo"},
				},
			},
		},
	}
	_, err := fakeClient.CoreV1().Services(suite.fooWithTargets.Namespace).Create(context.Background(), suite.fooWithTargets, metav1.CreateOptions{})
	suite.NoError(err, "should successfully create service")

	suite.sc, err = NewServiceSource(
		context.TODO(),
		fakeClient,
		"",
		"",
		"{{.Name}}",
		false,
		"",
		false,
		false,
		false,
		[]string{},
		false,
		labels.Everything(),
		false,
		false,
		false,
		false,
	)
	suite.NoError(err, "should initialize service source")
}

func (suite *ServiceSuite) TestResourceLabelIsSet() {
	endpoints, _ := suite.sc.Endpoints(context.Background())
	for _, ep := range endpoints {
		suite.Equal("service/default/foo-with-targets", ep.Labels[endpoint.ResourceLabelKey], "should set correct resource label")
	}
}

func TestServiceSource(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(ServiceSuite))
	t.Run("Interface", testServiceSourceImplementsSource)
	t.Run("NewServiceSource", testServiceSourceNewServiceSource)
	t.Run("Endpoints", testServiceSourceEndpoints)
	t.Run("MultipleServices", testMultipleServicesEndpoints)
}

// testServiceSourceImplementsSource tests that serviceSource is a valid Source.
func testServiceSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*Source)(nil), new(serviceSource))
}

// testServiceSourceNewServiceSource tests that NewServiceSource doesn't return an error.
func testServiceSourceNewServiceSource(t *testing.T) {
	t.Parallel()

	for _, ti := range []struct {
		title              string
		annotationFilter   string
		fqdnTemplate       string
		serviceTypesFilter []string
		expectError        bool
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
			title:            "non-empty annotation filter label",
			expectError:      false,
			annotationFilter: "kubernetes.io/ingress.class=nginx",
		},
		{
			title:              "non-empty service types filter",
			expectError:        false,
			serviceTypesFilter: []string{string(v1.ServiceTypeClusterIP)},
		},
	} {

		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			_, err := NewServiceSource(
				context.TODO(),
				fake.NewClientset(),
				"",
				ti.annotationFilter,
				ti.fqdnTemplate,
				false,
				"",
				false,
				false,
				false,
				ti.serviceTypesFilter,
				false,
				labels.Everything(),
				false,
				false,
				false,
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

// testServiceSourceEndpoints tests that various services generate the correct endpoints.
func testServiceSourceEndpoints(t *testing.T) {
	exampleDotComIP4, err := net.DefaultResolver.LookupNetIP(context.Background(), "ip4", "example.com")
	assert.NoError(t, err)
	exampleDotComIP6, err := net.DefaultResolver.LookupNetIP(context.Background(), "ip6", "example.com")
	assert.NoError(t, err)

	t.Parallel()

	for _, tc := range []struct {
		title                       string
		targetNamespace             string
		annotationFilter            string
		svcNamespace                string
		svcName                     string
		svcType                     v1.ServiceType
		compatibility               string
		fqdnTemplate                string
		combineFQDNAndAnnotation    bool
		ignoreHostnameAnnotation    bool
		labels                      map[string]string
		annotations                 map[string]string
		clusterIP                   string
		externalIPs                 []string
		lbs                         []string
		serviceTypesFilter          []string
		expected                    []*endpoint.Endpoint
		expectError                 bool
		serviceLabelSelector        string
		resolveLoadBalancerHostname bool
	}{
		{
			title:              "no annotated services return no endpoints",
			svcNamespace:       "testing",
			svcName:            "foo",
			svcType:            v1.ServiceTypeLoadBalancer,
			labels:             map[string]string{},
			annotations:        map[string]string{},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:                    "no annotated services return no endpoints when ignoring annotations",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			ignoreHostnameAnnotation: true,
			labels:                   map[string]string{},
			annotations:              map[string]string{},
			externalIPs:              []string{},
			lbs:                      []string{"1.2.3.4"},
			serviceTypesFilter:       []string{},
			expected:                 []*endpoint.Endpoint{},
		},
		{
			title:        "annotated services return an endpoint with target IP",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{string(v1.ServiceTypeLoadBalancer)},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:                    "hostname annotation on services is ignored",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			ignoreHostnameAnnotation: true,
			labels:                   map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:        "annotated ClusterIp aren't processed without explicit authorization",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			clusterIP:          "1.2.3.4",
			externalIPs:        []string{},
			lbs:                []string{},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:              "FQDN template with multiple hostnames return an endpoint with target IP",
			svcNamespace:       "testing",
			svcName:            "foo",
			svcType:            v1.ServiceTypeLoadBalancer,
			fqdnTemplate:       "{{.Name}}.fqdn.org,{{.Name}}.fqdn.com",
			labels:             map[string]string{},
			annotations:        map[string]string{},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{string(v1.ServiceTypeLoadBalancer), string(v1.ServiceTypeNodePort)},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.fqdn.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.fqdn.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "with excluded service type should not generate endpoints",
			svcNamespace:       "testing",
			svcName:            "foo",
			svcType:            v1.ServiceTypeLoadBalancer,
			fqdnTemplate:       "{{.Name}}.fqdn.org,{{.Name}}.fqdn.com",
			labels:             map[string]string{},
			annotations:        map[string]string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{string(v1.ServiceTypeNodePort)},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:                    "FQDN template with multiple hostnames return an endpoint with target IP when ignoring annotations",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			fqdnTemplate:             "{{.Name}}.fqdn.org,{{.Name}}.fqdn.com",
			ignoreHostnameAnnotation: true,
			labels:                   map[string]string{},
			annotations:              map[string]string{},
			externalIPs:              []string{},
			lbs:                      []string{"1.2.3.4"},
			serviceTypesFilter:       []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.fqdn.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.fqdn.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:                    "FQDN template and annotation both with multiple hostnames return an endpoint with target IP",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			fqdnTemplate:             "{{.Name}}.fqdn.org,{{.Name}}.fqdn.com",
			combineFQDNAndAnnotation: true,
			labels:                   map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org., bar.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.fqdn.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.fqdn.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:                    "FQDN template and annotation both with multiple hostnames while ignoring annotations will only return FQDN endpoints",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			fqdnTemplate:             "{{.Name}}.fqdn.org,{{.Name}}.fqdn.com",
			combineFQDNAndAnnotation: true,
			ignoreHostnameAnnotation: true,
			labels:                   map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org., bar.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.fqdn.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.fqdn.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "annotated services with multiple hostnames return an endpoint with target IP",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org., bar.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "annotated services with multiple hostnames and without trailing period return an endpoint with target IP",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org, bar.example.org",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "annotated services return an endpoint with target hostname",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"lb.example.com"}, // Kubernetes omits the trailing dot
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.com"}},
			},
		},
		{
			title:        "annotated services return an endpoint with hostname then resolve hostname",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:                 []string{},
			lbs:                         []string{"example.com"}, // Use a resolvable hostname for testing.
			serviceTypesFilter:          []string{},
			resolveLoadBalancerHostname: true,
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: testutils.NewTargetsFromAddr(exampleDotComIP4)},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: testutils.NewTargetsFromAddr(exampleDotComIP6)},
			},
		},
		{
			title:        "annotated services can omit trailing dot",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org", // Trailing dot is omitted
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4", "lb.example.com"}, // Kubernetes omits the trailing dot
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.com"}},
			},
		},
		{
			title:        "our controller type is kops dns controller",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.ControllerKey: annotations.ControllerValue,
				annotations.HostnameKey:   "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{string(v1.ServiceTypeLoadBalancer), string(v1.ServiceTypeNodePort)},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "different controller types are ignored even (with template specified)",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.ControllerKey: "some-other-tool",
				annotations.HostnameKey:   "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:           "services are found in target namespace",
			targetNamespace: "testing",
			svcNamespace:    "testing",
			svcName:         "foo",
			svcType:         v1.ServiceTypeLoadBalancer,
			labels:          map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:           "services that are not in target namespace are ignored",
			targetNamespace: "testing",
			svcNamespace:    "other-testing",
			svcName:         "foo",
			svcType:         v1.ServiceTypeLoadBalancer,
			labels:          map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:        "services are found in all namespaces",
			svcNamespace: "other-testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:            "valid matching annotation filter expression",
			annotationFilter: "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeLoadBalancer,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey:                       "foo.example.org.",
				"service.beta.kubernetes.io/external-traffic": "OnlyLocal",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:            "valid non-matching annotation filter expression",
			annotationFilter: "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeLoadBalancer,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey:                       "foo.example.org.",
				"service.beta.kubernetes.io/external-traffic": "SomethingElse",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:            "invalid annotation filter expression",
			annotationFilter: "service.beta.kubernetes.io/external-traffic in (Global OnlyLocal)",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeLoadBalancer,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey:                       "foo.example.org.",
				"service.beta.kubernetes.io/external-traffic": "OnlyLocal",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
			expectError:        true,
		},
		{
			title:            "valid matching annotation filter label",
			annotationFilter: "service.beta.kubernetes.io/external-traffic=Global",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeLoadBalancer,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey:                       "foo.example.org.",
				"service.beta.kubernetes.io/external-traffic": "Global",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:            "valid non-matching annotation filter label",
			annotationFilter: "service.beta.kubernetes.io/external-traffic=Global",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeLoadBalancer,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey:                       "foo.example.org.",
				"service.beta.kubernetes.io/external-traffic": "OnlyLocal",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:        "no external entrypoints return no endpoints",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:        "annotated service with externalIPs returns a single endpoint with multiple targets",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{"10.2.3.4", "11.2.3.4"},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.2.3.4", "11.2.3.4"}},
			},
		},
		{
			title:        "multiple external entrypoints return a single endpoint with multiple targets",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4", "8.8.8.8"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "8.8.8.8"}},
			},
		},
		{
			title:        "services annotated with legacy mate annotations are ignored in default mode",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				"zalando.org/dnsname": "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:         "services annotated with legacy mate annotations return an endpoint in compatibility mode",
			svcNamespace:  "testing",
			svcName:       "foo",
			svcType:       v1.ServiceTypeLoadBalancer,
			compatibility: "mate",
			labels:        map[string]string{},
			annotations: map[string]string{
				"zalando.org/dnsname": "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:         "services annotated with legacy molecule annotations return an endpoint in compatibility mode",
			svcNamespace:  "testing",
			svcName:       "foo",
			svcType:       v1.ServiceTypeLoadBalancer,
			compatibility: "molecule",
			labels: map[string]string{
				"dns": "route53",
			},
			annotations: map[string]string{
				"domainName": "foo.example.org., bar.example.org",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:         "load balancer services annotated with DNS Controller annotations return an endpoint with A and CNAME targets in compatibility mode",
			svcNamespace:  "testing",
			svcName:       "foo",
			svcType:       v1.ServiceTypeLoadBalancer,
			compatibility: "kops-dns-controller",
			labels:        map[string]string{},
			annotations: map[string]string{
				kopsDNSControllerInternalHostnameAnnotationKey: "internal.foo.example.org",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4", "lb.example.com"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "internal.foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "internal.foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"lb.example.com"}},
			},
		},
		{
			title:         "load balancer services annotated with DNS Controller annotations return an endpoint with both annotations in compatibility mode",
			svcNamespace:  "testing",
			svcName:       "foo",
			svcType:       v1.ServiceTypeLoadBalancer,
			compatibility: "kops-dns-controller",
			labels:        map[string]string{},
			annotations: map[string]string{
				kopsDNSControllerInternalHostnameAnnotationKey: "internal.foo.example.org., internal.bar.example.org",
				kopsDNSControllerHostnameAnnotationKey:         "foo.example.org., bar.example.org",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "internal.foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "internal.bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "not annotated services with set fqdnTemplate return an endpoint with target IP",
			svcNamespace:       "testing",
			svcName:            "foo",
			svcType:            v1.ServiceTypeLoadBalancer,
			fqdnTemplate:       "{{.Name}}.bar.example.com",
			labels:             map[string]string{},
			annotations:        map[string]string{},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4", "elb.com"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.bar.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.bar.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"elb.com"}},
			},
		},
		{
			title:        "annotated services with set fqdnTemplate annotation takes precedence",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			fqdnTemplate: "{{.Name}}.bar.example.com",
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4", "elb.com"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"elb.com"}},
			},
		},
		{
			title:         "compatibility annotated services with tmpl. compatibility takes precedence",
			svcNamespace:  "testing",
			svcName:       "foo",
			svcType:       v1.ServiceTypeLoadBalancer,
			compatibility: "mate",
			fqdnTemplate:  "{{.Name}}.bar.example.com",
			labels:        map[string]string{},
			annotations: map[string]string{
				"zalando.org/dnsname": "mate.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "mate.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "not annotated services with unknown tmpl field should not return anything",
			svcNamespace:       "testing",
			svcName:            "foo",
			svcType:            v1.ServiceTypeLoadBalancer,
			fqdnTemplate:       "{{.Calibre}}.bar.example.com",
			labels:             map[string]string{},
			annotations:        map[string]string{},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected:           []*endpoint.Endpoint{},
			expectError:        true,
		},
		{
			title:        "ttl not annotated should have RecordTTL.IsConfigured set to false",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:        "ttl annotated but invalid should have RecordTTL.IsConfigured set to false",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TtlKey:      "foo",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:        "ttl annotated and is valid should set Record.TTL",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TtlKey:      "10",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(10)},
			},
		},
		{
			title:        "ttl annotated (in duration format) and is valid should set Record.TTL",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TtlKey:      "1m",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(60)},
			},
		},
		{
			title:        "Negative ttl is not valid",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TtlKey:      "-10",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:        "filter on service types should include matching services",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{string(v1.ServiceTypeLoadBalancer)},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "filter on service types should exclude non-matching services",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeNodePort,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{string(v1.ServiceTypeLoadBalancer)},
			expected:           []*endpoint.Endpoint{},
		},
		{
			title:        "internal-host annotated and host annotated clusterip services return an endpoint with Cluster IP",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey:         "foo.example.org.",
				annotations.InternalHostnameKey: "foo.internal.example.org.",
			},
			clusterIP:          "1.1.1.1",
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.internal.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
			},
		},
		{
			title:        "internal-host annotated loadbalancer services return an endpoint with Cluster IP",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.InternalHostnameKey: "foo.internal.example.org.",
			},
			clusterIP:          "1.1.1.1",
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.internal.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
			},
		},
		{
			title:        "internal-host annotated and host annotated loadbalancer services return an endpoint with Cluster IP and an endpoint with lb IP",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels:       map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey:         "foo.example.org.",
				annotations.InternalHostnameKey: "foo.internal.example.org.",
			},
			clusterIP:          "1.1.1.1",
			externalIPs:        []string{},
			lbs:                []string{"1.2.3.4"},
			serviceTypesFilter: []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.internal.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "service with matching labels and fqdn filter should be included",
			svcNamespace: "testing",
			svcName:      "fqdn",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels: map[string]string{
				"app": "web-external",
			},
			clusterIP:            "1.1.1.1",
			externalIPs:          []string{},
			lbs:                  []string{"1.2.3.4"},
			serviceTypesFilter:   []string{},
			serviceLabelSelector: "app=web-external",
			fqdnTemplate:         "{{.Name}}.bar.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "fqdn.bar.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "service with matching labels and hostname annotation should be included",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels: map[string]string{
				"app": "web-external",
			},
			clusterIP:            "1.1.1.1",
			externalIPs:          []string{},
			lbs:                  []string{"1.2.3.4"},
			serviceTypesFilter:   []string{},
			serviceLabelSelector: "app=web-external",
			annotations:          map[string]string{annotations.HostnameKey: "annotation.bar.example.com"},
			expected: []*endpoint.Endpoint{
				{DNSName: "annotation.bar.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "service without matching labels and fqdn filter should be excluded",
			svcNamespace: "testing",
			svcName:      "fqdn",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels: map[string]string{
				"app": "web-internal",
			},
			clusterIP:            "1.1.1.1",
			externalIPs:          []string{},
			lbs:                  []string{"1.2.3.4"},
			serviceTypesFilter:   []string{},
			serviceLabelSelector: "app=web-external",
			fqdnTemplate:         "{{.Name}}.bar.example.com",
			expected:             []*endpoint.Endpoint{},
		},
		{
			title:        "service without matching labels and hostname annotation should be excluded",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeLoadBalancer,
			labels: map[string]string{
				"app": "web-internal",
			},
			clusterIP:            "1.1.1.1",
			externalIPs:          []string{},
			lbs:                  []string{"1.2.3.4"},
			serviceTypesFilter:   []string{},
			serviceLabelSelector: "app=web-external",
			annotations:          map[string]string{annotations.HostnameKey: "annotation.bar.example.com"},
			expected:             []*endpoint.Endpoint{},
		},
		{
			title:              "dual-stack load-balancer service gets both addresses",
			svcNamespace:       "testing",
			svcName:            "foobar",
			svcType:            v1.ServiceTypeLoadBalancer,
			labels:             map[string]string{},
			clusterIP:          "1.1.1.2,2001:db8::2",
			externalIPs:        []string{},
			lbs:                []string{"1.1.1.1", "2001:db8::1"},
			serviceTypesFilter: []string{},
			annotations:        map[string]string{annotations.HostnameKey: "foobar.example.org"},
			expected: []*endpoint.Endpoint{
				{DNSName: "foobar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "foobar.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
			},
		},
		{
			title:              "IPv6-only load-balancer service gets IPv6 endpoint",
			svcNamespace:       "testing",
			svcName:            "foobar-v6",
			svcType:            v1.ServiceTypeLoadBalancer,
			labels:             map[string]string{},
			clusterIP:          "2001:db8::1",
			externalIPs:        []string{},
			lbs:                []string{"2001:db8::2"},
			serviceTypesFilter: []string{},
			annotations:        map[string]string{annotations.HostnameKey: "foobar-v6.example.org"},
			expected: []*endpoint.Endpoint{
				{DNSName: "foobar-v6.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::2"}},
			},
		},
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			// Create a Kubernetes testing client
			kubernetes := fake.NewClientset()

			// Create a service to test against
			ingresses := []v1.LoadBalancerIngress{}
			for _, lb := range tc.lbs {
				if net.ParseIP(lb) != nil {
					ingresses = append(ingresses, v1.LoadBalancerIngress{IP: lb})
				} else {
					ingresses = append(ingresses, v1.LoadBalancerIngress{Hostname: lb})
				}
			}

			service := &v1.Service{
				Spec: v1.ServiceSpec{
					Type:        tc.svcType,
					ClusterIP:   tc.clusterIP,
					ExternalIPs: tc.externalIPs,
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   tc.svcNamespace,
					Name:        tc.svcName,
					Labels:      tc.labels,
					Annotations: tc.annotations,
				},
				Status: v1.ServiceStatus{
					LoadBalancer: v1.LoadBalancerStatus{
						Ingress: ingresses,
					},
				},
			}

			_, err := kubernetes.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
			require.NoError(t, err)

			var sourceLabel labels.Selector
			if tc.serviceLabelSelector != "" {
				sourceLabel, err = labels.Parse(tc.serviceLabelSelector)
				require.NoError(t, err)
			} else {
				sourceLabel = labels.Everything()
			}

			// Create our object under test and get the endpoints.
			client, err := NewServiceSource(
				context.TODO(),
				kubernetes,
				tc.targetNamespace,
				tc.annotationFilter,
				tc.fqdnTemplate,
				tc.combineFQDNAndAnnotation,
				tc.compatibility,
				false,
				false,
				false,
				tc.serviceTypesFilter,
				tc.ignoreHostnameAnnotation,
				sourceLabel,
				tc.resolveLoadBalancerHostname,
				false,
				false,
				false,
			)

			require.NoError(t, err)

			res, err := client.Endpoints(context.Background())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, res, tc.expected)
		})
	}
}

// testMultipleServicesEndpoints tests that multiple services generate correct merged endpoints
func testMultipleServicesEndpoints(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		annotationFilter         string
		svcNamespace             string
		svcName                  string
		svcType                  v1.ServiceType
		compatibility            string
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		ignoreHostnameAnnotation bool
		labels                   map[string]string
		clusterIP                string
		services                 map[string]map[string]string
		serviceTypesFilter       []string
		expected                 []*endpoint.Endpoint
		expectError              bool
	}{
		{
			"test service returns a correct end point",
			"",
			"",
			"testing",
			"foo",
			v1.ServiceTypeLoadBalancer,
			"",
			"",
			false,
			false,
			map[string]string{},
			"",
			map[string]map[string]string{
				"1.2.3.4": {annotations.HostnameKey: "foo.example.org"},
			},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
			},
			false,
		},
		{
			"multiple services that share same DNS should be merged into one endpoint",
			"",
			"",
			"testing",
			"foo",
			v1.ServiceTypeLoadBalancer,
			"",
			"",
			false,
			false,
			map[string]string{},
			"",
			map[string]map[string]string{
				"1.2.3.4": {annotations.HostnameKey: "foo.example.org"},
				"1.2.3.5": {annotations.HostnameKey: "foo.example.org"},
				"1.2.3.6": {annotations.HostnameKey: "foo.example.org"},
			},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5", "1.2.3.6"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
			},
			false,
		},
		{
			"test that services with different hostnames do not get merged together",
			"",
			"",
			"testing",
			"foo",
			v1.ServiceTypeLoadBalancer,
			"",
			"",
			false,
			false,
			map[string]string{},
			"",
			map[string]map[string]string{
				"1.2.3.5":  {annotations.HostnameKey: "foo.example.org"},
				"10.1.1.3": {annotations.HostnameKey: "bar.example.org"},
				"10.1.1.1": {annotations.HostnameKey: "bar.example.org"},
				"1.2.3.4":  {annotations.HostnameKey: "foo.example.org"},
				"10.1.1.2": {annotations.HostnameKey: "bar.example.org"},
				"20.1.1.1": {annotations.HostnameKey: "foobar.example.org"},
				"1.2.3.6":  {annotations.HostnameKey: "foo.example.org"},
			},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5", "1.2.3.6"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.1", "10.1.1.2", "10.1.1.3"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo10.1.1.1"}},
				{DNSName: "foobar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"20.1.1.1"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo20.1.1.1"}},
			},
			false,
		},
		{
			"test that services with different set-identifier do not get merged together",
			"",
			"",
			"testing",
			"foo",
			v1.ServiceTypeLoadBalancer,
			"",
			"",
			false,
			false,
			map[string]string{},
			"",
			map[string]map[string]string{
				"1.2.3.5":  {annotations.HostnameKey: "foo.example.org", annotations.SetIdentifierKey: "a"},
				"10.1.1.3": {annotations.HostnameKey: "foo.example.org", annotations.SetIdentifierKey: "b"},
			},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.5"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.5"}, SetIdentifier: "a"},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.3"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo10.1.1.3"}, SetIdentifier: "b"},
			},
			false,
		},
		{
			"test that services with CNAME types do not get merged together",
			"",
			"",
			"testing",
			"foo",
			v1.ServiceTypeLoadBalancer,
			"",
			"",
			false,
			false,
			map[string]string{},
			"",
			map[string]map[string]string{
				"a.elb.com": {annotations.HostnameKey: "foo.example.org"},
				"b.elb.com": {annotations.HostnameKey: "foo.example.org"},
			},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"a.elb.com"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/fooa.elb.com"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"b.elb.com"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foob.elb.com"}},
			},
			false,
		},
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			// Create a Kubernetes testing client
			kubernetes := fake.NewClientset()

			// Create services to test against
			for lb, ants := range tc.services {
				ingresses := []v1.LoadBalancerIngress{}
				ingresses = append(ingresses, v1.LoadBalancerIngress{IP: lb})

				service := &v1.Service{
					Spec: v1.ServiceSpec{
						Type:      tc.svcType,
						ClusterIP: tc.clusterIP,
					},
					ObjectMeta: metav1.ObjectMeta{
						Namespace:   tc.svcNamespace,
						Name:        tc.svcName + lb,
						Labels:      tc.labels,
						Annotations: ants,
					},
					Status: v1.ServiceStatus{
						LoadBalancer: v1.LoadBalancerStatus{
							Ingress: ingresses,
						},
					},
				}

				_, err := kubernetes.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			// Create our object under test and get the endpoints.
			client, err := NewServiceSource(
				context.TODO(),
				kubernetes,
				tc.targetNamespace,
				tc.annotationFilter,
				tc.fqdnTemplate,
				tc.combineFQDNAndAnnotation,
				tc.compatibility,
				false,
				false,
				false,
				tc.serviceTypesFilter,
				tc.ignoreHostnameAnnotation,
				labels.Everything(),
				false,
				false,
				false,
				false,
			)
			require.NoError(t, err)

			res, err := client.Endpoints(context.Background())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, res, tc.expected)
			// Test that endpoint resourceLabelKey matches desired endpoint
			sort.SliceStable(res, func(i, j int) bool {
				return strings.Compare(res[i].DNSName, res[j].DNSName) < 0
			})
			sort.SliceStable(tc.expected, func(i, j int) bool {
				return strings.Compare(tc.expected[i].DNSName, tc.expected[j].DNSName) < 0
			})

			for i := range res {
				if res[i].Labels[endpoint.ResourceLabelKey] != tc.expected[i].Labels[endpoint.ResourceLabelKey] {
					t.Errorf("expected %s, got %s", tc.expected[i].Labels[endpoint.ResourceLabelKey], res[i].Labels[endpoint.ResourceLabelKey])
				}
			}
		})
	}
}

// testServiceSourceEndpoints tests that various services generate the correct endpoints.
func TestClusterIpServices(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		annotationFilter         string
		svcNamespace             string
		svcName                  string
		svcType                  v1.ServiceType
		compatibility            string
		fqdnTemplate             string
		ignoreHostnameAnnotation bool
		labels                   map[string]string
		annotations              map[string]string
		clusterIP                string
		expected                 []*endpoint.Endpoint
		expectError              bool
		labelSelector            string
	}{
		{
			title:        "hostname annotated ClusterIp services return an endpoint with Cluster IP",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			clusterIP: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "target annotated ClusterIp services return an endpoint with the specified A",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TargetKey:   "4.3.2.1",
			},
			clusterIP: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"4.3.2.1"}},
			},
		},
		{
			title:        "target annotated ClusterIp services return an endpoint with the specified CNAME",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TargetKey:   "bar.example.org.",
			},
			clusterIP: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"bar.example.org"}},
			},
		},
		{
			title:        "target annotated ClusterIp services return an endpoint with the specified AAAA",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TargetKey:   "2001:DB8::1",
			},
			clusterIP: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1"}},
			},
		},
		{
			title:        "multiple target annotated ClusterIp services return an endpoint with the specified CNAMES",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TargetKey:   "bar.example.org.,baz.example.org.",
			},
			clusterIP: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"bar.example.org", "baz.example.org"}},
			},
		},
		{
			title:        "multiple target annotated ClusterIp services return two endpoints with the specified CNAMES and AAAA",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TargetKey:   "bar.example.org.,baz.example.org.,2001:DB8::1",
			},
			clusterIP: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"bar.example.org", "baz.example.org"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1"}},
			},
		},
		{
			title:                    "hostname annotated ClusterIp services are ignored",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeClusterIP,
			ignoreHostnameAnnotation: true,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			clusterIP: "1.2.3.4",
			expected:  []*endpoint.Endpoint{},
		},
		{
			title:                    "hostname and target annotated ClusterIp services are ignored",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeClusterIP,
			ignoreHostnameAnnotation: true,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TargetKey:   "bar.example.org.",
			},
			clusterIP: "1.2.3.4",
			expected:  []*endpoint.Endpoint{},
		},
		{
			title:        "hostname and target annotated ClusterIp services return an endpoint with the specified CNAME",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TargetKey:   "bar.example.org.",
			},
			clusterIP: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"bar.example.org"}},
			},
		},
		{
			title:        "non-annotated ClusterIp services with set fqdnTemplate return an endpoint with target IP",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			fqdnTemplate: "{{.Name}}.bar.example.com",
			clusterIP:    "4.5.6.7",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.bar.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"4.5.6.7"}},
			},
		},
		{
			title:        "Headless services do not generate endpoints",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			clusterIP:    v1.ClusterIPNone,
			expected:     []*endpoint.Endpoint{},
		},
		{
			title:        "Headless services generate endpoints when target is specified",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.TargetKey:   "bar.example.org.",
			},
			clusterIP: v1.ClusterIPNone,
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"bar.example.org"}},
			},
		},
		{
			title:        "ClusterIP service with matching label generates an endpoint",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			fqdnTemplate: "{{.Name}}.bar.example.com",
			labels:       map[string]string{"app": "web-internal"},
			clusterIP:    "4.5.6.7",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.bar.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"4.5.6.7"}},
			},
			labelSelector: "app=web-internal",
		},
		{
			title:        "ClusterIP service with matching label and target generates a CNAME endpoint",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			fqdnTemplate: "{{.Name}}.bar.example.com",
			labels:       map[string]string{"app": "web-internal"},
			annotations:  map[string]string{annotations.TargetKey: "bar.example.com."},
			clusterIP:    "4.5.6.7",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.bar.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"bar.example.com"}},
			},
			labelSelector: "app=web-internal",
		},
		{
			title:         "ClusterIP service without matching label generates an endpoint",
			svcNamespace:  "testing",
			svcName:       "foo",
			svcType:       v1.ServiceTypeClusterIP,
			fqdnTemplate:  "{{.Name}}.bar.example.com",
			labels:        map[string]string{"app": "web-internal"},
			clusterIP:     "4.5.6.7",
			expected:      []*endpoint.Endpoint{},
			labelSelector: "app=web-external",
		},
		{
			title:        "invalid hostname does not generate endpoints",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				annotations.HostnameKey: "this-is-an-exceedingly-long-label-that-external-dns-should-reject.example.org.",
			},
			clusterIP: "1.2.3.4",
			expected:  []*endpoint.Endpoint{},
		},
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			// Create a Kubernetes testing client
			kubernetes := fake.NewClientset()

			// Create a service to test against
			service := &v1.Service{
				Spec: v1.ServiceSpec{
					Type:      tc.svcType,
					ClusterIP: tc.clusterIP,
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   tc.svcNamespace,
					Name:        tc.svcName,
					Labels:      tc.labels,
					Annotations: tc.annotations,
				},
			}

			_, err := kubernetes.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
			require.NoError(t, err)

			var labelSelector labels.Selector
			if tc.labelSelector != "" {
				labelSelector, err = labels.Parse(tc.labelSelector)
				require.NoError(t, err)
			} else {
				labelSelector = labels.Everything()
			}
			// Create our object under test and get the endpoints.
			client, _ := NewServiceSource(
				context.TODO(),
				kubernetes,
				tc.targetNamespace,
				tc.annotationFilter,
				tc.fqdnTemplate,
				false,
				tc.compatibility,
				true,
				false,
				false,
				[]string{},
				tc.ignoreHostnameAnnotation,
				labelSelector,
				false,
				false,
				false,
				false,
			)
			require.NoError(t, err)

			endpoints, err := client.Endpoints(context.Background())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)
		})
	}
}

// testNodePortServices tests that various services generate the correct endpoints.
func TestServiceSourceNodePortServices(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		annotationFilter         string
		svcNamespace             string
		svcName                  string
		svcType                  v1.ServiceType
		svcTrafficPolicy         v1.ServiceExternalTrafficPolicyType
		compatibility            string
		fqdnTemplate             string
		ignoreHostnameAnnotation bool
		exposeInternalIPv6       bool
		ignoreUnscheduledNodes   bool
		labels                   map[string]string
		annotations              map[string]string
		lbs                      []string
		expected                 []*endpoint.Endpoint
		expectError              bool
		nodes                    []*v1.Node
		podNames                 []string
		nodeIndex                []int
		phases                   []v1.PodPhase
		conditions               []v1.PodCondition
		labelSelector            labels.Selector
		deletionTimestamp        []*metav1.Time
	}{
		{
			title:            "annotated NodePort services return an endpoint with IP addresses of the cluster's nodes",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeCluster,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"54.10.11.1", "54.10.11.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::3"}, RecordType: endpoint.RecordTypeAAAA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::3"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::4"},
					},
				},
			}},
		},
		{
			title:                    "hostname annotated NodePort services are ignored",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeNodePort,
			svcTrafficPolicy:         v1.ServiceExternalTrafficPolicyTypeCluster,
			ignoreHostnameAnnotation: true,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::3"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::4"},
					},
				},
			}},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:            "non-annotated NodePort services with set fqdnTemplate return an endpoint with target IP",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeCluster,
			fqdnTemplate:     "{{.Name}}.bar.example.com",
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.bar.example.com", Targets: endpoint.Targets{"0 50 30192 foo.bar.example.com."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.bar.example.com", Targets: endpoint.Targets{"54.10.11.1", "54.10.11.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.bar.example.com", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::3"}, RecordType: endpoint.RecordTypeAAAA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::3"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::4"},
					},
				},
			}},
		},
		{
			title:            "annotated NodePort services return an endpoint with IP addresses of the private cluster's nodes",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeCluster,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"10.0.1.1", "10.0.1.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}, RecordType: endpoint.RecordTypeAAAA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::1"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}},
		},
		{
			title:            "annotated NodePort services with ExternalTrafficPolicy=Local return an endpoint with IP addresses of the cluster's nodes where pods is running only",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeLocal,
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"54.10.11.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"2001:DB8::3"}, RecordType: endpoint.RecordTypeAAAA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::3"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::4"},
					},
				},
			}},
			podNames:          []string{"pod-0"},
			nodeIndex:         []int{1},
			phases:            []v1.PodPhase{v1.PodRunning},
			conditions:        []v1.PodCondition{{Type: v1.PodReady, Status: v1.ConditionFalse}},
			deletionTimestamp: []*metav1.Time{{}},
		},
		{
			title:            "annotated NodePort services with ExternalTrafficPolicy=Local and multiple pods on a single node return an endpoint with unique IP addresses of the cluster's nodes where pods is running only",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeLocal,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"54.10.11.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"2001:DB8::3"}, RecordType: endpoint.RecordTypeAAAA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::3"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::4"},
					},
				},
			}},
			podNames:  []string{"pod-0", "pod-1"},
			nodeIndex: []int{1, 1},
			phases:    []v1.PodPhase{v1.PodRunning, v1.PodRunning},
			conditions: []v1.PodCondition{
				{Type: v1.PodReady, Status: v1.ConditionFalse},
				{Type: v1.PodReady, Status: v1.ConditionFalse},
			},
			deletionTimestamp: []*metav1.Time{{}, {}},
		},
		{
			title:            "annotated NodePort services with ExternalTrafficPolicy=Local return pods in Ready & Running state",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeLocal,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"54.10.11.1"}, RecordType: endpoint.RecordTypeA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
					},
				},
			}},
			podNames:  []string{"pod-0", "pod-1"},
			nodeIndex: []int{0, 1},
			phases:    []v1.PodPhase{v1.PodRunning, v1.PodRunning},
			conditions: []v1.PodCondition{
				{Type: v1.PodReady, Status: v1.ConditionTrue},
				{Type: v1.PodReady, Status: v1.ConditionFalse},
			},
			deletionTimestamp: []*metav1.Time{{}, {}},
		},
		{
			title:            "annotated NodePort services with ExternalTrafficPolicy=Local return pods in Ready & Running state & not in Terminating",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeLocal,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"54.10.11.1"}, RecordType: endpoint.RecordTypeA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node3",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.3"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.3"},
					},
				},
			}},
			podNames:  []string{"pod-0", "pod-1", "pod-2"},
			nodeIndex: []int{0, 1, 2},
			phases:    []v1.PodPhase{v1.PodRunning, v1.PodRunning, v1.PodRunning},
			conditions: []v1.PodCondition{
				{Type: v1.PodReady, Status: v1.ConditionTrue},
				{Type: v1.PodReady, Status: v1.ConditionFalse},
				{Type: v1.PodReady, Status: v1.ConditionTrue},
			},
			deletionTimestamp: []*metav1.Time{nil, nil, {}},
		},
		{
			title:            "access=private annotation NodePort services return an endpoint with private IP addresses of the cluster's nodes",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeCluster,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.AccessKey:   "private",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"10.0.1.1", "10.0.1.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}, RecordType: endpoint.RecordTypeAAAA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::1"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}},
		},
		{
			title:            "access=public annotation NodePort services return an endpoint with external IP addresses of the cluster's nodes if exposeInternalIPv6 is unset",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeCluster,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.AccessKey:   "public",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"54.10.11.1", "54.10.11.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::3"}, RecordType: endpoint.RecordTypeAAAA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::3"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::4"},
					},
				},
			}},
		},
		{
			title:            "access=public annotation NodePort services return an endpoint with public IP addresses of the cluster's nodes if exposeInternalIPv6 is set to true",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeCluster,
			labels:           map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.AccessKey:   "public",
			},
			exposeInternalIPv6: true,
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"54.10.11.1", "54.10.11.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2", "2001:DB8::3", "2001:DB8::4"}, RecordType: endpoint.RecordTypeAAAA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::3"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::4"},
					},
				},
			}},
		},
		{
			title:            "node port services annotated DNS Controller annotations return an endpoint where all targets has the node role",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeCluster,
			compatibility:    "kops-dns-controller",
			labels:           map[string]string{},
			annotations: map[string]string{
				kopsDNSControllerInternalHostnameAnnotationKey: "internal.foo.example.org., internal.bar.example.org",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "internal.foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.1.1"}},
				{DNSName: "internal.foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1"}},
				{DNSName: "internal.bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.1.1"}},
				{DNSName: "internal.bar.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1"}},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
					Labels: map[string]string{
						"node-role.kubernetes.io/node": "",
					},
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::1"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
					Labels: map[string]string{
						"node-role.kubernetes.io/control-plane": "",
					},
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}},
		},
		{
			title:            "node port services annotated with internal DNS Controller annotations return an endpoint in compatibility mode",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeCluster,
			compatibility:    "kops-dns-controller",
			annotations: map[string]string{
				kopsDNSControllerInternalHostnameAnnotationKey: "internal.foo.example.org., internal.bar.example.org",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "internal.foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.1.1", "10.0.1.2"}},
				{DNSName: "internal.foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}},
				{DNSName: "internal.bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.0.1.1", "10.0.1.2"}},
				{DNSName: "internal.bar.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
					Labels: map[string]string{
						"node-role.kubernetes.io/node": "",
					},
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::1"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
					Labels: map[string]string{
						"node-role.kubernetes.io/node": "",
					},
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}},
		},
		{
			title:              "node port services annotated with external DNS Controller annotations return an endpoint in compatibility mode with exposeInternalIPv6 flag set",
			svcNamespace:       "testing",
			svcName:            "foo",
			svcType:            v1.ServiceTypeNodePort,
			svcTrafficPolicy:   v1.ServiceExternalTrafficPolicyTypeCluster,
			compatibility:      "kops-dns-controller",
			exposeInternalIPv6: true,
			annotations: map[string]string{
				kopsDNSControllerHostnameAnnotationKey: "foo.example.org., bar.example.org",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"54.10.11.1", "54.10.11.2"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"54.10.11.1", "54.10.11.2"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:DB8::1", "2001:DB8::2"}},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
					Labels: map[string]string{
						"node-role.kubernetes.io/node": "",
					},
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::1"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
					Labels: map[string]string{
						"node-role.kubernetes.io/node": "",
					},
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}},
		},
		{
			title:            "node port services annotated with both kops dns controller annotations return an empty set of addons",
			svcNamespace:     "testing",
			svcName:          "foo",
			svcType:          v1.ServiceTypeNodePort,
			svcTrafficPolicy: v1.ServiceExternalTrafficPolicyTypeCluster,
			compatibility:    "kops-dns-controller",
			labels:           map[string]string{},
			annotations: map[string]string{
				kopsDNSControllerInternalHostnameAnnotationKey: "internal.foo.example.org., internal.bar.example.org",
				kopsDNSControllerHostnameAnnotationKey:         "foo.example.org., bar.example.org",
			},
			expected: []*endpoint.Endpoint{},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
					Labels: map[string]string{
						"node-role.kubernetes.io/node": "",
					},
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::1"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
					Labels: map[string]string{
						"node-role.kubernetes.io/node": "",
					},
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}},
		},
		{
			title:                  "NodePort services ignore unschedulable node",
			ignoreUnscheduledNodes: true,
			svcNamespace:           "testing",
			svcName:                "foo",
			svcType:                v1.ServiceTypeNodePort,
			svcTrafficPolicy:       v1.ServiceExternalTrafficPolicyTypeCluster,
			labels:                 map[string]string{},
			annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
				annotations.AccessKey:   "public",
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_foo._tcp.foo.example.org", Targets: endpoint.Targets{"0 50 30192 foo.example.org."}, RecordType: endpoint.RecordTypeSRV},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"54.10.11.2"}, RecordType: endpoint.RecordTypeA},
				{DNSName: "foo.example.org", Targets: endpoint.Targets{"2001:DB8::3"}, RecordType: endpoint.RecordTypeAAAA},
			},
			nodes: []*v1.Node{{
				ObjectMeta: metav1.ObjectMeta{
					Name: "node1",
				},
				Spec: v1.NodeSpec{
					Unschedulable: true,
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.1"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.1"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::1"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::2"},
					},
				},
			}, {
				ObjectMeta: metav1.ObjectMeta{
					Name: "node2",
				},
				Spec: v1.NodeSpec{
					Unschedulable: false,
				},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{Type: v1.NodeExternalIP, Address: "54.10.11.2"},
						{Type: v1.NodeInternalIP, Address: "10.0.1.2"},
						{Type: v1.NodeExternalIP, Address: "2001:DB8::3"},
						{Type: v1.NodeInternalIP, Address: "2001:DB8::4"},
					},
				},
			}},
		},
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			// Create a Kubernetes testing client
			kubernetes := fake.NewClientset()

			// Create the nodes
			for _, node := range tc.nodes {
				if _, err := kubernetes.CoreV1().Nodes().Create(context.Background(), node, metav1.CreateOptions{}); err != nil {
					t.Fatal(err)
				}
			}

			// Create  pods
			for i, podname := range tc.podNames {
				pod := &v1.Pod{
					Spec: v1.PodSpec{
						Containers: []v1.Container{},
						Hostname:   podname,
						NodeName:   tc.nodes[tc.nodeIndex[i]].Name,
					},
					ObjectMeta: metav1.ObjectMeta{
						Namespace:         tc.svcNamespace,
						Name:              podname,
						Labels:            tc.labels,
						Annotations:       tc.annotations,
						DeletionTimestamp: tc.deletionTimestamp[i],
					},
					Status: v1.PodStatus{
						Phase:      tc.phases[i],
						Conditions: []v1.PodCondition{tc.conditions[i]},
					},
				}

				_, err := kubernetes.CoreV1().Pods(tc.svcNamespace).Create(context.Background(), pod, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			// Create a service to test against
			service := &v1.Service{
				Spec: v1.ServiceSpec{
					Type:                  tc.svcType,
					ExternalTrafficPolicy: tc.svcTrafficPolicy,
					Ports: []v1.ServicePort{
						{
							NodePort: 30192,
						},
					},
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   tc.svcNamespace,
					Name:        tc.svcName,
					Labels:      tc.labels,
					Annotations: tc.annotations,
				},
			}

			_, err := kubernetes.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
			require.NoError(t, err)

			// Create our object under test and get the endpoints.
			client, _ := NewServiceSource(
				context.TODO(),
				kubernetes,
				tc.targetNamespace,
				tc.annotationFilter,
				tc.fqdnTemplate,
				false,
				tc.compatibility,
				true,
				false,
				false,
				[]string{},
				tc.ignoreHostnameAnnotation,
				labels.Everything(),
				false,
				false,
				tc.exposeInternalIPv6,
				tc.ignoreUnscheduledNodes,
			)
			require.NoError(t, err)

			endpoints, err := client.Endpoints(context.Background())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)
		})
	}
}

// TestHeadlessServices tests that headless services generate the correct endpoints.
func TestHeadlessServices(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		svcNamespace             string
		svcName                  string
		svcType                  v1.ServiceType
		compatibility            string
		fqdnTemplate             string
		ignoreHostnameAnnotation bool
		exposeInternalIPv6       bool
		labels                   map[string]string
		svcAnnotations           map[string]string
		podAnnotations           map[string]string
		clusterIP                string
		podIPs                   []string
		hostIPs                  []string
		selector                 map[string]string
		lbs                      []string
		podnames                 []string
		hostnames                []string
		podsReady                []bool
		publishNotReadyAddresses bool
		nodes                    []v1.Node
		serviceTypesFilter       []string
		expected                 []*endpoint.Endpoint
		expectError              bool
	}{
		{
			"annotated Headless services return IPv4 endpoints for each selected Pod",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			[]string{"", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.2"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "1.1.1.2"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv6 endpoints for each selected Pod",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"2001:db8::1", "2001:db8::2"},
			[]string{"", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			false,
			[]v1.Node{},
			[]string{string(v1.ServiceTypeClusterIP), string(v1.ServiceTypeLoadBalancer)},
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::2"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1", "2001:db8::2"}},
			},
			false,
		},
		{
			"hostname annotated Headless services are ignored",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			true,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			[]string{"", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{},
			false,
		},
		{
			"annotated Headless services return IPv4 endpoints with TTL for each selected Pod",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
				annotations.TtlKey:      "1",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			[]string{"", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}, RecordTTL: endpoint.TTL(1)},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.2"}, RecordTTL: endpoint.TTL(1)},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "1.1.1.2"}, RecordTTL: endpoint.TTL(1)},
			},
			false,
		},
		{
			"annotated Headless services return IPv6 endpoints with TTL for each selected Pod",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
				annotations.TtlKey:      "1",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"2001:db8::1", "2001:db8::2"},
			[]string{"", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}, RecordTTL: endpoint.TTL(1)},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::2"}, RecordTTL: endpoint.TTL(1)},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1", "2001:db8::2"}, RecordTTL: endpoint.TTL(1)},
			},
			false,
		},
		{
			"annotated Headless services return endpoints for each selected Pod, which are in running state",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			[]string{"", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, false},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
			},
			false,
		},
		{
			"annotated Headless services return endpoints for all Pod if publishNotReadyAddresses is set",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			[]string{"", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, false},
			true,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.2"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "1.1.1.2"}},
			},
			false,
		},
		{
			"annotated Headless services return endpoints for pods missing hostname",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			[]string{"", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"", ""},
			[]bool{true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "1.1.1.2"}},
			},
			false,
		},
		{
			"annotated Headless services return only a unique set of IPv4 targets",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.1", "1.1.1.2"},
			[]string{"", "", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1", "foo-3"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "1.1.1.2"}},
			},
			false,
		},
		{
			"annotated Headless services return only a unique set of IPv6 targets",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"2001:db8::1", "2001:db8::1", "2001:db8::2"},
			[]string{"", "", ""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1", "foo-3"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1", "2001:db8::2"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv4 targets from pod annotation",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{
				annotations.TargetKey: "1.2.3.4",
			},
			v1.ClusterIPNone,
			[]string{"1.1.1.1"},
			[]string{""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{},
			[]string{string(v1.ServiceTypeClusterIP)},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv6 targets from pod annotation",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			map[string]string{
				annotations.TargetKey: "2001:db8::4",
			},
			v1.ClusterIPNone,
			[]string{"2001:db8::1"},
			[]string{""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::4"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv4 targets from node external IP if endpoints-type annotation is set",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey:      "service.example.org",
				annotations.EndpointsTypeKey: EndpointsTypeNodeExternalIP,
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1"},
			[]string{""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{
				{
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{
								Type:    v1.NodeExternalIP,
								Address: "1.2.3.4",
							},
						},
					},
				},
			},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			false,
		},
		{
			"annotated Headless services return only external IPv6 targets from node IP if endpoints-type annotation is set and exposeInternalIPv6 flag is unset",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey:      "service.example.org",
				annotations.EndpointsTypeKey: EndpointsTypeNodeExternalIP,
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"2001:db8::1"},
			[]string{""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{
				{
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{
								Type:    v1.NodeInternalIP,
								Address: "2001:db8::4",
							},
							{
								Type:    v1.NodeExternalIP,
								Address: "2001:db8::5",
							},
						},
					},
				},
			},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::5"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv6 targets from node external IP if endpoints-type annotation is set and exposeInternalIPv6 flag set",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			true,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey:      "service.example.org",
				annotations.EndpointsTypeKey: EndpointsTypeNodeExternalIP,
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"2001:db8::1"},
			[]string{""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{
				{
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{
								Type:    v1.NodeInternalIP,
								Address: "2001:db8::4",
							},
						},
					},
				},
			},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::4"}},
			},
			false,
		},
		{
			"annotated Headless services return dual-stack targets from node external IP if endpoints-type annotation is set and exposeInternalIPv6 flag set",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			true,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey:      "service.example.org",
				annotations.EndpointsTypeKey: EndpointsTypeNodeExternalIP,
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1"},
			[]string{""},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{
				{
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{
								Type:    v1.NodeExternalIP,
								Address: "1.2.3.4",
							},
							{
								Type:    v1.NodeInternalIP,
								Address: "2001:db8::4",
							},
						},
					},
				},
			},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::4"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv4 targets from hostIP if endpoints-type annotation is set",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey:      "service.example.org",
				annotations.EndpointsTypeKey: EndpointsTypeHostIP,
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"1.1.1.1"},
			[]string{"1.2.3.4"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv6 targets from hostIP if endpoints-type annotation is set",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey:      "service.example.org",
				annotations.EndpointsTypeKey: EndpointsTypeHostIP,
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"2001:db8::1"},
			[]string{"2001:db8::4"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::4"}},
			},
			false,
		},
		{
			"headless service with endpoints-type annotation is outside of serviceTypeFilter scope",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey:      "service.example.org",
				annotations.EndpointsTypeKey: EndpointsTypeNodeExternalIP,
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"2001:db8::1"},
			[]string{"2001:db8::4"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{
				{
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{
								Type:    v1.NodeExternalIP,
								Address: "1.2.3.4",
							},
							{
								Type:    v1.NodeInternalIP,
								Address: "10.0.10.12",
							},
						},
					},
				},
			},
			[]string{string(v1.ServiceTypeClusterIP)},
			[]*endpoint.Endpoint{},
			false,
		},
		{
			"headless service with endpoints-type annotation is in the scope of serviceTypeFilter",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey:      "service.example.org",
				annotations.EndpointsTypeKey: EndpointsTypeNodeExternalIP,
			},
			map[string]string{},
			v1.ClusterIPNone,
			[]string{"2001:db8::1"},
			[]string{"1.2.3.4"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo"},
			[]string{"", "", ""},
			[]bool{true, true, true},
			false,
			[]v1.Node{
				{
					Status: v1.NodeStatus{
						Addresses: []v1.NodeAddress{
							{
								Type:    v1.NodeExternalIP,
								Address: "1.2.3.4",
							},
							{
								Type:    v1.NodeInternalIP,
								Address: "10.0.10.12",
							},
						},
					},
				},
			},
			[]string{string(v1.ServiceTypeClusterIP), string(v1.ServiceTypeNodePort)},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
			false,
		},
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			// Create a Kubernetes testing client
			kubernetes := fake.NewClientset()

			service := &v1.Service{
				Spec: v1.ServiceSpec{
					Type:                     tc.svcType,
					ClusterIP:                tc.clusterIP,
					Selector:                 tc.selector,
					PublishNotReadyAddresses: tc.publishNotReadyAddresses,
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   tc.svcNamespace,
					Name:        tc.svcName,
					Labels:      tc.labels,
					Annotations: tc.svcAnnotations,
				},
				Status: v1.ServiceStatus{},
			}
			_, err := kubernetes.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
			require.NoError(t, err)

			var endpointSliceEndpoints []discoveryv1.Endpoint
			for i, podName := range tc.podnames {
				pod := &v1.Pod{
					Spec: v1.PodSpec{
						Containers: []v1.Container{},
						Hostname:   tc.hostnames[i],
					},
					ObjectMeta: metav1.ObjectMeta{
						Namespace:   tc.svcNamespace,
						Name:        podName,
						Labels:      tc.labels,
						Annotations: tc.podAnnotations,
					},
					Status: v1.PodStatus{
						PodIP:  tc.podIPs[i],
						HostIP: tc.hostIPs[i],
					},
				}

				_, err = kubernetes.CoreV1().Pods(tc.svcNamespace).Create(context.Background(), pod, metav1.CreateOptions{})
				require.NoError(t, err)

				ep := discoveryv1.Endpoint{
					Addresses: []string{tc.podIPs[i]},
					TargetRef: &v1.ObjectReference{
						APIVersion: "",
						Kind:       "Pod",
						Name:       podName,
					},
					Conditions: discoveryv1.EndpointConditions{
						Ready: &tc.podsReady[i],
					},
				}
				endpointSliceEndpoints = append(endpointSliceEndpoints, ep)
			}
			endpointSliceLabels := maps.Clone(tc.labels)
			endpointSliceLabels[discoveryv1.LabelServiceName] = tc.svcName
			endpointSlice := &discoveryv1.EndpointSlice{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: tc.svcNamespace,
					Name:      tc.svcName,
					Labels:    endpointSliceLabels,
				},
				AddressType: discoveryv1.AddressTypeIPv4,
				Endpoints:   endpointSliceEndpoints,
			}
			_, err = kubernetes.DiscoveryV1().EndpointSlices(tc.svcNamespace).Create(context.Background(), endpointSlice, metav1.CreateOptions{})
			require.NoError(t, err)
			for _, node := range tc.nodes {
				_, err = kubernetes.CoreV1().Nodes().Create(context.Background(), &node, metav1.CreateOptions{})
				require.NoError(t, err)
			}

			// Create our object under test and get the endpoints.
			client, _ := NewServiceSource(
				context.TODO(),
				kubernetes,
				tc.targetNamespace,
				"",
				tc.fqdnTemplate,
				false,
				tc.compatibility,
				true,
				false,
				false,
				tc.serviceTypesFilter,
				tc.ignoreHostnameAnnotation,
				labels.Everything(),
				false,
				false,
				tc.exposeInternalIPv6,
				false,
			)
			require.NoError(t, err)

			endpoints, err := client.Endpoints(context.Background())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)
		})
	}
}

func TestMultipleServicesPointingToSameLoadBalancer(t *testing.T) {
	kubernetes := fake.NewClientset()

	services := []*v1.Service{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "istio-ingressgateway",
				Namespace: "default",
				Labels: map[string]string{
					"app":   "istio-ingressgateway",
					"istio": "ingressgateway",
				},
				Annotations: map[string]string{
					"external-dns.alpha.kubernetes.io/hostname": "example.org",
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
				Annotations: map[string]string{
					"external-dns.alpha.kubernetes.io/hostname": "example.org",
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
		_, err := kubernetes.CoreV1().Services(svc.Namespace).Create(t.Context(), svc, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	src, err := NewServiceSource(
		t.Context(),
		kubernetes,
		v1.NamespaceAll,
		"",
		"",
		false,
		"",
		false,
		false,
		false,
		[]string{},
		false,
		labels.Everything(),
		false,
		false,
		false,
		true,
	)
	require.NoError(t, err)
	assert.NotNil(t, src)

	got, err := src.Endpoints(t.Context())
	require.NoError(t, err)

	validateEndpoints(t, got, []*endpoint.Endpoint{
		endpoint.NewEndpoint("example.org", endpoint.RecordTypeA, "34.66.66.77").WithLabel(endpoint.ResourceLabelKey, "service/default/istio-ingressgateway"),
	})
}

func TestMultipleHeadlessServicesPointingToPodsOnTheSameNode(t *testing.T) {
	kubernetes := fake.NewClientset()

	headless := []*v1.Service{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "kafka",
				Namespace: "default",
				Labels: map[string]string{
					"app": "kafka",
				},
				Annotations: map[string]string{
					annotations.HostnameKey: "example.org",
				},
			},
			Spec: v1.ServiceSpec{
				Type:                  v1.ServiceTypeClusterIP,
				ClusterIP:             v1.ClusterIPNone,
				ClusterIPs:            []string{v1.ClusterIPNone},
				InternalTrafficPolicy: testutils.ToPtr(v1.ServiceInternalTrafficPolicyCluster),
				IPFamilies:            []v1.IPFamily{v1.IPv4Protocol},
				IPFamilyPolicy:        testutils.ToPtr(v1.IPFamilyPolicySingleStack),
				Ports: []v1.ServicePort{
					{
						Name:       "web",
						Port:       80,
						Protocol:   v1.ProtocolTCP,
						TargetPort: intstr.FromInt32(80),
					},
				},
				Selector: map[string]string{
					"app": "kafka",
				},
				SessionAffinity: v1.ServiceAffinityNone,
			},
			Status: v1.ServiceStatus{
				LoadBalancer: v1.LoadBalancerStatus{},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "kafka-2",
				Namespace: "default",
				Labels: map[string]string{
					"app": "kafka",
				},
				Annotations: map[string]string{
					annotations.HostnameKey: "example.org",
				},
			},
			Spec: v1.ServiceSpec{
				Type:                  v1.ServiceTypeClusterIP,
				ClusterIP:             v1.ClusterIPNone,
				ClusterIPs:            []string{v1.ClusterIPNone},
				InternalTrafficPolicy: testutils.ToPtr(v1.ServiceInternalTrafficPolicyCluster),
				IPFamilies:            []v1.IPFamily{v1.IPv4Protocol},
				IPFamilyPolicy:        testutils.ToPtr(v1.IPFamilyPolicySingleStack),
				Ports: []v1.ServicePort{
					{
						Name:       "web",
						Port:       80,
						Protocol:   v1.ProtocolTCP,
						TargetPort: intstr.FromInt32(80),
					},
				},
				Selector: map[string]string{
					"app": "kafka",
				},
				SessionAffinity: v1.ServiceAffinityNone,
			},
			Status: v1.ServiceStatus{
				LoadBalancer: v1.LoadBalancerStatus{},
			},
		},
	}

	assert.NotNil(t, headless)

	pods := []*v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "kafka-0",
				Namespace: "default",
				Labels: map[string]string{
					"app":                                 "kafka",
					appsv1.PodIndexLabel:                  "0",
					appsv1.ControllerRevisionHashLabelKey: "kafka-b8d79cdb6",
					appsv1.StatefulSetPodNameLabel:        "kafka-0",
				},
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: "apps/v1",
						Kind:       "StatefulSet",
						Name:       "kafka",
					},
				},
			},
			Spec: v1.PodSpec{
				Hostname:  "kafka-0",
				Subdomain: "kafka",
				NodeName:  "local-dev-worker",
				Containers: []v1.Container{
					{
						Name: "nginx",
						Ports: []v1.ContainerPort{
							{Name: "web", ContainerPort: 80, Protocol: v1.ProtocolTCP},
						},
					},
				},
			},
			Status: v1.PodStatus{
				Phase:   v1.PodRunning,
				PodIP:   "10.244.1.2",
				PodIPs:  []v1.PodIP{{IP: "10.244.1.2"}},
				HostIP:  "172.18.0.2",
				HostIPs: []v1.HostIP{{IP: "172.18.0.2"}},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "kafka-1",
				Namespace: "default",
				Labels: map[string]string{
					"app":                                 "kafka",
					appsv1.PodIndexLabel:                  "1",
					appsv1.ControllerRevisionHashLabelKey: "kafka-b8d79cdb6",
					appsv1.StatefulSetPodNameLabel:        "kafka-1",
				},
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: "apps/v1",
						Kind:       "StatefulSet",
						Name:       "kafka",
					},
				},
			},
			Spec: v1.PodSpec{
				Hostname:  "kafka-1",
				Subdomain: "kafka",
				NodeName:  "local-dev-worker",
				Containers: []v1.Container{
					{
						Name: "nginx",
						Ports: []v1.ContainerPort{
							{Name: "web", ContainerPort: 80, Protocol: v1.ProtocolTCP},
						},
					},
				},
			},
			Status: v1.PodStatus{
				Phase:   v1.PodRunning,
				PodIP:   "10.244.1.3",
				PodIPs:  []v1.PodIP{{IP: "10.244.1.3"}},
				HostIP:  "172.18.0.2",
				HostIPs: []v1.HostIP{{IP: "172.18.0.2"}},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "kafka-2",
				Namespace: "default",
				Labels: map[string]string{
					"app":                                 "kafka",
					appsv1.PodIndexLabel:                  "2",
					appsv1.ControllerRevisionHashLabelKey: "kafka-b8d79cdb6",
					appsv1.StatefulSetPodNameLabel:        "kafka-2",
				},
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: "apps/v1",
						Kind:       "StatefulSet",
						Name:       "kafka",
					},
				},
			},
			Spec: v1.PodSpec{
				Hostname:  "kafka-2",
				Subdomain: "kafka",
				NodeName:  "local-dev-worker",
				Containers: []v1.Container{
					{
						Name: "nginx",
						Ports: []v1.ContainerPort{
							{Name: "web", ContainerPort: 80, Protocol: v1.ProtocolTCP},
						},
					},
				},
			},
			Status: v1.PodStatus{
				Phase:   v1.PodRunning,
				PodIP:   "10.244.1.4",
				PodIPs:  []v1.PodIP{{IP: "10.244.1.4"}},
				HostIP:  "172.18.0.2",
				HostIPs: []v1.HostIP{{IP: "172.18.0.2"}},
			},
		},
	}
	assert.Len(t, pods, 3)

	endpoints := []*discoveryv1.EndpointSlice{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "kafka-xhrc9",
				Namespace: "default",
				Labels: map[string]string{
					"app":                        "kafka",
					discoveryv1.LabelServiceName: "kafka",
					discoveryv1.LabelManagedBy:   "endpointslice-controller.k8s.io",
					v1.IsHeadlessService:         "",
				},
			},
			AddressType: discoveryv1.AddressTypeIPv4,
			Endpoints: []discoveryv1.Endpoint{
				{
					Addresses: []string{"10.244.1.2"},
					Hostname:  testutils.ToPtr("kafka-0"),
					NodeName:  testutils.ToPtr("local-dev-worker"),
					TargetRef: &v1.ObjectReference{
						Kind:      "Pod",
						Name:      "kafka-0",
						Namespace: "default",
					},
					Conditions: discoveryv1.EndpointConditions{
						Ready:       testutils.ToPtr(true),
						Serving:     testutils.ToPtr(true),
						Terminating: testutils.ToPtr(false),
					},
				},
				{
					Addresses: []string{"10.244.1.3"},
					Hostname:  testutils.ToPtr("kafka-1"),
					NodeName:  testutils.ToPtr("local-dev-worker"),
					TargetRef: &v1.ObjectReference{
						Kind:      "Pod",
						Name:      "kafka-1",
						Namespace: "default",
					},
					Conditions: discoveryv1.EndpointConditions{
						Ready:       testutils.ToPtr(true),
						Serving:     testutils.ToPtr(true),
						Terminating: testutils.ToPtr(false),
					},
				},
				{
					Addresses: []string{"10.244.1.4"},
					Hostname:  testutils.ToPtr("kafka-2"),
					NodeName:  testutils.ToPtr("local-dev-worker"),
					TargetRef: &v1.ObjectReference{
						Kind:      "Pod",
						Name:      "kafka-2",
						Namespace: "default",
					},
					Conditions: discoveryv1.EndpointConditions{
						Ready:       testutils.ToPtr(true),
						Serving:     testutils.ToPtr(true),
						Terminating: testutils.ToPtr(false),
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "kafka-2-svwsg",
				Namespace: "default",
				Labels: map[string]string{
					"app":                        "kafka",
					discoveryv1.LabelServiceName: "kafka-2",
					discoveryv1.LabelManagedBy:   "endpointslice-controller.k8s.io",
					v1.IsHeadlessService:         "",
				},
			},
			AddressType: discoveryv1.AddressTypeIPv4,
			Endpoints: []discoveryv1.Endpoint{
				{
					Addresses: []string{"10.244.1.2"},
					Hostname:  testutils.ToPtr("kafka-0"),
					NodeName:  testutils.ToPtr("local-dev-worker"),
					TargetRef: &v1.ObjectReference{
						Kind:      "Pod",
						Name:      "kafka-0",
						Namespace: "default",
					},
					Conditions: discoveryv1.EndpointConditions{
						Ready:       testutils.ToPtr(true),
						Serving:     testutils.ToPtr(true),
						Terminating: testutils.ToPtr(false),
					},
				},
				{
					Addresses: []string{"10.244.1.3"},
					Hostname:  testutils.ToPtr("kafka-1"),
					NodeName:  testutils.ToPtr("local-dev-worker"),
					TargetRef: &v1.ObjectReference{
						Kind:      "Pod",
						Name:      "kafka-1",
						Namespace: "default",
					},
					Conditions: discoveryv1.EndpointConditions{
						Ready:       testutils.ToPtr(true),
						Serving:     testutils.ToPtr(true),
						Terminating: testutils.ToPtr(false),
					},
				},
				{
					Addresses: []string{"10.244.1.4"},
					Hostname:  testutils.ToPtr("kafka-2"),
					NodeName:  testutils.ToPtr("local-dev-worker"),
					TargetRef: &v1.ObjectReference{
						Kind:      "Pod",
						Name:      "kafka-2",
						Namespace: "default",
					},
					Conditions: discoveryv1.EndpointConditions{
						Ready:       testutils.ToPtr(true),
						Serving:     testutils.ToPtr(true),
						Terminating: testutils.ToPtr(false),
					},
				},
			},
		},
	}

	for _, svc := range headless {
		_, err := kubernetes.CoreV1().Services(svc.Namespace).Create(context.Background(), svc, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	for _, pod := range pods {
		_, err := kubernetes.CoreV1().Pods(pod.Namespace).Create(context.Background(), pod, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	for _, ep := range endpoints {
		_, err := kubernetes.DiscoveryV1().EndpointSlices(ep.Namespace).Create(context.Background(), ep, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	src, err := NewServiceSource(
		t.Context(),
		kubernetes,
		v1.NamespaceAll,
		"",
		"",
		false,
		"",
		false,
		false,
		false,
		[]string{},
		false,
		labels.Everything(),
		false,
		false,
		false,
		true,
	)
	require.NoError(t, err)
	assert.NotNil(t, src)

	got, err := src.Endpoints(context.Background())
	require.NoError(t, err)

	want := []*endpoint.Endpoint{
		// TODO: root domain records should not be created. Address them in a follow-up PR.
		{DNSName: "example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.244.1.2", "10.244.1.3", "10.244.1.4"}},
		{DNSName: "kafka-0.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.244.1.2"}},
		{DNSName: "kafka-1.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.244.1.3"}},
		{DNSName: "kafka-2.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.244.1.4"}},
	}

	validateEndpoints(t, got, want)
}

// TestHeadlessServices tests that headless services generate the correct endpoints.
func TestHeadlessServicesHostIP(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		svcNamespace             string
		svcName                  string
		svcType                  v1.ServiceType
		compatibility            string
		fqdnTemplate             string
		ignoreHostnameAnnotation bool
		labels                   map[string]string
		annotations              map[string]string
		clusterIP                string
		hostIPs                  []string
		selector                 map[string]string
		lbs                      []string
		podnames                 []string
		hostnames                []string
		podsReady                []bool
		targetRefs               []*v1.ObjectReference
		publishNotReadyAddresses bool
		expected                 []*endpoint.Endpoint
		expectError              bool
	}{
		{
			"annotated Headless services return IPv4 endpoints for each selected Pod",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			[]*v1.ObjectReference{
				{APIVersion: "", Kind: "Pod", Name: "foo-0"},
				{APIVersion: "", Kind: "Pod", Name: "foo-1"},
			},
			false,
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.2"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "1.1.1.2"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv6 endpoints for each selected Pod",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			v1.ClusterIPNone,
			[]string{"2001:db8::1", "2001:db8::2"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			[]*v1.ObjectReference{
				{APIVersion: "", Kind: "Pod", Name: "foo-0"},
				{APIVersion: "", Kind: "Pod", Name: "foo-1"},
			},
			false,
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::2"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1", "2001:db8::2"}},
			},
			false,
		},
		{
			"hostname annotated Headless services are ignored",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			true,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			[]*v1.ObjectReference{
				{APIVersion: "", Kind: "Pod", Name: "foo-0"},
				{APIVersion: "", Kind: "Pod", Name: "foo-1"},
			},
			false,
			[]*endpoint.Endpoint{},
			false,
		},
		{
			"annotated Headless services return IPv4 endpoints with TTL for each selected Pod",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
				annotations.TtlKey:      "1",
			},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			[]*v1.ObjectReference{
				{APIVersion: "", Kind: "Pod", Name: "foo-0"},
				{APIVersion: "", Kind: "Pod", Name: "foo-1"},
			},
			false,
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}, RecordTTL: endpoint.TTL(1)},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.2"}, RecordTTL: endpoint.TTL(1)},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "1.1.1.2"}, RecordTTL: endpoint.TTL(1)},
			},
			false,
		},
		{
			"annotated Headless services return IPv6 endpoints with TTL for each selected Pod",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
				annotations.TtlKey:      "1",
			},
			v1.ClusterIPNone,
			[]string{"2001:db8::1", "2001:db8::2"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, true},
			[]*v1.ObjectReference{
				{APIVersion: "", Kind: "Pod", Name: "foo-0"},
				{APIVersion: "", Kind: "Pod", Name: "foo-1"},
			},
			false,
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}, RecordTTL: endpoint.TTL(1)},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::2"}, RecordTTL: endpoint.TTL(1)},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1", "2001:db8::2"}, RecordTTL: endpoint.TTL(1)},
			},
			false,
		},
		{
			"annotated Headless services return endpoints for each selected Pod, which are in running state",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, false},
			[]*v1.ObjectReference{
				{APIVersion: "", Kind: "Pod", Name: "foo-0"},
				{APIVersion: "", Kind: "Pod", Name: "foo-1"},
			},
			false,
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
			},
			false,
		},
		{
			"annotated Headless services return endpoints for all Pod if publishNotReadyAddresses is set",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"foo-0", "foo-1"},
			[]bool{true, false},
			[]*v1.ObjectReference{
				{APIVersion: "", Kind: "Pod", Name: "foo-0"},
				{APIVersion: "", Kind: "Pod", Name: "foo-1"},
			},
			true,
			[]*endpoint.Endpoint{
				{DNSName: "foo-0.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1"}},
				{DNSName: "foo-1.service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.2"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "1.1.1.2"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv4 endpoints for pods missing hostname",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			v1.ClusterIPNone,
			[]string{"1.1.1.1", "1.1.1.2"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"", ""},
			[]bool{true, true},
			[]*v1.ObjectReference{
				{APIVersion: "", Kind: "Pod", Name: "foo-0"},
				{APIVersion: "", Kind: "Pod", Name: "foo-1"},
			},
			false,
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.1.1.1", "1.1.1.2"}},
			},
			false,
		},
		{
			"annotated Headless services return IPv6 endpoints for pods missing hostname",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			v1.ClusterIPNone,
			[]string{"2001:db8::1", "2001:db8::2"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0", "foo-1"},
			[]string{"", ""},
			[]bool{true, true},
			[]*v1.ObjectReference{
				{APIVersion: "", Kind: "Pod", Name: "foo-0"},
				{APIVersion: "", Kind: "Pod", Name: "foo-1"},
			},
			false,
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1", "2001:db8::2"}},
			},
			false,
		},
		{
			"annotated Headless services without a targetRef has no endpoints",
			"",
			"testing",
			"foo",
			v1.ServiceTypeClusterIP,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			v1.ClusterIPNone,
			[]string{"1.1.1.1"},
			map[string]string{
				"component": "foo",
			},
			[]string{},
			[]string{"foo-0"},
			[]string{"foo-0"},
			[]bool{true, true},
			[]*v1.ObjectReference{nil},
			false,
			[]*endpoint.Endpoint{},
			false,
		},
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			kubernetes := fake.NewClientset()

			service := &v1.Service{
				Spec: v1.ServiceSpec{
					Type:                     tc.svcType,
					ClusterIP:                tc.clusterIP,
					Selector:                 tc.selector,
					PublishNotReadyAddresses: tc.publishNotReadyAddresses,
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   tc.svcNamespace,
					Name:        tc.svcName,
					Labels:      tc.labels,
					Annotations: tc.annotations,
				},
				Status: v1.ServiceStatus{},
			}
			_, err := kubernetes.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
			require.NoError(t, err)

			var endpointsSlicesEndpoints []discoveryv1.Endpoint
			for i, podname := range tc.podnames {
				pod := &v1.Pod{
					Spec: v1.PodSpec{
						Containers: []v1.Container{},
						Hostname:   tc.hostnames[i],
					},
					ObjectMeta: metav1.ObjectMeta{
						Namespace:   tc.svcNamespace,
						Name:        podname,
						Labels:      tc.labels,
						Annotations: tc.annotations,
					},
					Status: v1.PodStatus{
						HostIP: tc.hostIPs[i],
					},
				}

				_, err = kubernetes.CoreV1().Pods(tc.svcNamespace).Create(context.Background(), pod, metav1.CreateOptions{})
				require.NoError(t, err)

				ep := discoveryv1.Endpoint{
					Addresses: []string{"4.3.2.1"},
					TargetRef: tc.targetRefs[i],
					Conditions: discoveryv1.EndpointConditions{
						Ready: &tc.podsReady[i],
					},
				}
				endpointsSlicesEndpoints = append(endpointsSlicesEndpoints, ep)
			}
			endpointSliceLabels := maps.Clone(tc.labels)
			endpointSliceLabels[discoveryv1.LabelServiceName] = tc.svcName
			endpointSlice := &discoveryv1.EndpointSlice{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: tc.svcNamespace,
					Name:      tc.svcName,
					Labels:    endpointSliceLabels,
				},
				AddressType: discoveryv1.AddressTypeIPv4,
				Endpoints:   endpointsSlicesEndpoints,
			}
			_, err = kubernetes.DiscoveryV1().EndpointSlices(tc.svcNamespace).Create(context.Background(), endpointSlice, metav1.CreateOptions{})
			require.NoError(t, err)

			// Create our object under test and get the endpoints.
			client, _ := NewServiceSource(
				context.TODO(),
				kubernetes,
				tc.targetNamespace,
				"",
				tc.fqdnTemplate,
				false,
				tc.compatibility,
				true,
				true,
				false,
				[]string{},
				tc.ignoreHostnameAnnotation,
				labels.Everything(),
				false,
				false,
				false,
				true,
			)
			require.NoError(t, err)

			endpoints, err := client.Endpoints(context.Background())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)

			// TODO; when all resources have the resource label, we could add this check to the validateEndpoints function.
			for _, ep := range endpoints {
				require.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
			}
		})
	}
}

// TestExternalServices tests that external services generate the correct endpoints.
func TestExternalServices(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		svcNamespace             string
		svcName                  string
		svcType                  v1.ServiceType
		compatibility            string
		fqdnTemplate             string
		ignoreHostnameAnnotation bool
		labels                   map[string]string
		annotations              map[string]string
		externalName             string
		externalIPs              []string
		serviceTypeFilter        []string
		expected                 []*endpoint.Endpoint
		expectError              bool
	}{
		{
			"external services return an A endpoint for the external name that is an IPv4 address",
			"",
			"testing",
			"foo",
			v1.ServiceTypeExternalName,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			"111.111.111.111",
			[]string{},
			[]string{string(v1.ServiceTypeNodePort), string(v1.ServiceTypeExternalName)},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", Targets: endpoint.Targets{"111.111.111.111"}, RecordType: endpoint.RecordTypeA},
			},
			false,
		},
		{
			"external services return an AAAA endpoint for the external name that is an IPv6 address",
			"",
			"testing",
			"foo",
			v1.ServiceTypeExternalName,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			"2001:db8::111",
			[]string{},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", Targets: endpoint.Targets{"2001:db8::111"}, RecordType: endpoint.RecordTypeAAAA},
			},
			false,
		},
		{
			"external services return a CNAME endpoint for the external name that is a domain",
			"",
			"testing",
			"foo",
			v1.ServiceTypeExternalName,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			"remote.example.com",
			[]string{},
			[]string{string(v1.ServiceTypeExternalName)},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", Targets: endpoint.Targets{"remote.example.com"}, RecordType: endpoint.RecordTypeCNAME},
			},
			false,
		},
		{
			"annotated ExternalName service with externalIPs returns a single endpoint with multiple targets",
			"",
			"testing",
			"foo",
			v1.ServiceTypeExternalName,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			"service.example.org",
			[]string{"10.2.3.4", "11.2.3.4"},
			[]string{},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.2.3.4", "11.2.3.4"}},
			},
			false,
		},
		{
			"annotated ExternalName service with externalIPs of dualstack addresses returns 2 endpoints with multiple targets",
			"",
			"testing",
			"foo",
			v1.ServiceTypeExternalName,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			"service.example.org",
			[]string{"10.2.3.4", "11.2.3.4", "2001:db8::1", "2001:db8::2"},
			[]string{string(v1.ServiceTypeNodePort), string(v1.ServiceTypeExternalName)},
			[]*endpoint.Endpoint{
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.2.3.4", "11.2.3.4"}},
				{DNSName: "service.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1", "2001:db8::2"}},
			},
			false,
		},
		{
			"annotated ExternalName service with externalIPs of dualstack and excluded in serviceTypeFilter",
			"",
			"testing",
			"foo",
			v1.ServiceTypeExternalName,
			"",
			"",
			false,
			map[string]string{"component": "foo"},
			map[string]string{
				annotations.HostnameKey: "service.example.org",
			},
			"service.example.org",
			[]string{"10.2.3.4", "11.2.3.4", "2001:db8::1", "2001:db8::2"},
			[]string{string(v1.ServiceTypeNodePort), string(v1.ServiceTypeClusterIP)},
			[]*endpoint.Endpoint{},
			false,
		},
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			// Create a Kubernetes testing client
			kubernetes := fake.NewClientset()

			service := &v1.Service{
				Spec: v1.ServiceSpec{
					Type:         tc.svcType,
					ExternalName: tc.externalName,
					ExternalIPs:  tc.externalIPs,
				},
				ObjectMeta: metav1.ObjectMeta{
					Namespace:   tc.svcNamespace,
					Name:        tc.svcName,
					Labels:      tc.labels,
					Annotations: tc.annotations,
				},
				Status: v1.ServiceStatus{},
			}
			_, err := kubernetes.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
			require.NoError(t, err)

			// Create our object under test and get the endpoints.
			client, _ := NewServiceSource(
				context.TODO(),
				kubernetes,
				tc.targetNamespace,
				"",
				tc.fqdnTemplate,
				false,
				tc.compatibility,
				true,
				false,
				false,
				tc.serviceTypeFilter,
				tc.ignoreHostnameAnnotation,
				labels.Everything(),
				false,
				false,
				false,
				true,
			)
			require.NoError(t, err)

			endpoints, err := client.Endpoints(context.Background())
			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// Validate returned endpoints against desired endpoints.
			validateEndpoints(t, endpoints, tc.expected)

			// TODO; when all resources have the resource label, we could add this check to the validateEndpoints function.
			for _, ep := range endpoints {
				require.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
			}
		})
	}
}

func BenchmarkServiceEndpoints(b *testing.B) {
	kubernetes := fake.NewClientset()

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "testing",
			Name:      "foo",
			Annotations: map[string]string{
				annotations.HostnameKey: "foo.example.org.",
			},
		},
		Status: v1.ServiceStatus{
			LoadBalancer: v1.LoadBalancerStatus{
				Ingress: []v1.LoadBalancerIngress{
					{IP: "1.2.3.4"},
					{IP: "8.8.8.8"},
				},
			},
		},
	}

	_, err := kubernetes.CoreV1().Services(service.Namespace).Create(context.Background(), service, metav1.CreateOptions{})
	require.NoError(b, err)

	client, err := NewServiceSource(
		context.TODO(),
		kubernetes,
		v1.NamespaceAll,
		"",
		"",
		false,
		"",
		false,
		false,
		false,
		[]string{},
		false,
		labels.Everything(),
		false,
		false,
		false,
		true,
	)
	require.NoError(b, err)

	for b.Loop() {
		_, err := client.Endpoints(context.Background())
		require.NoError(b, err)
	}
}

func TestNewServiceSourceInformersEnabled(t *testing.T) {
	tests := []struct {
		name      string
		asserts   func(svc *serviceSource)
		svcFilter []string
	}{
		{
			name: "serviceTypeFilter is set to empty",
			asserts: func(svc *serviceSource) {
				assert.NotNil(t, svc)
				assert.NotNil(t, svc.serviceTypeFilter)
				assert.False(t, svc.serviceTypeFilter.enabled)
				assert.NotNil(t, svc.nodeInformer)
				assert.NotNil(t, svc.serviceInformer)
				assert.NotNil(t, svc.endpointSlicesInformer)
			},
		},
		{
			name:      "serviceTypeFilter contains NodePort",
			svcFilter: []string{string(v1.ServiceTypeClusterIP)},
			asserts: func(svc *serviceSource) {
				assert.NotNil(t, svc)
				assert.NotNil(t, svc.serviceTypeFilter)
				assert.True(t, svc.serviceTypeFilter.enabled)
				assert.NotNil(t, svc.serviceInformer)
				assert.Nil(t, svc.nodeInformer)
				assert.NotNil(t, svc.endpointSlicesInformer)
				assert.NotNil(t, svc.podInformer)
			},
		},
		{
			name:      "serviceTypeFilter contains NodePort and ExternalName",
			svcFilter: []string{string(v1.ServiceTypeNodePort), string(v1.ServiceTypeExternalName)},
			asserts: func(svc *serviceSource) {
				assert.NotNil(t, svc)
				assert.NotNil(t, svc.serviceTypeFilter)
				assert.True(t, svc.serviceTypeFilter.enabled)
				assert.NotNil(t, svc.serviceInformer)
				assert.NotNil(t, svc.nodeInformer)
				assert.NotNil(t, svc.endpointSlicesInformer)
				assert.NotNil(t, svc.podInformer)
			},
		},
		{
			name:      "serviceTypeFilter contains ExternalName",
			svcFilter: []string{string(v1.ServiceTypeExternalName)},
			asserts: func(svc *serviceSource) {
				assert.NotNil(t, svc)
				assert.NotNil(t, svc.serviceTypeFilter)
				assert.True(t, svc.serviceTypeFilter.enabled)
				assert.NotNil(t, svc.serviceInformer)
				assert.Nil(t, svc.nodeInformer)
				assert.Nil(t, svc.endpointSlicesInformer)
				assert.Nil(t, svc.podInformer)
			},
		},
		{
			name:      "serviceTypeFilter contains LoadBalancer",
			svcFilter: []string{string(v1.ServiceTypeLoadBalancer)},
			asserts: func(svc *serviceSource) {
				assert.NotNil(t, svc)
				assert.NotNil(t, svc.serviceTypeFilter)
				assert.True(t, svc.serviceTypeFilter.enabled)
				assert.NotNil(t, svc.serviceInformer)
				assert.Nil(t, svc.nodeInformer)
				assert.Nil(t, svc.endpointSlicesInformer)
				assert.Nil(t, svc.podInformer)
			},
		},
	}

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			svc, err := NewServiceSource(
				t.Context(),
				fake.NewClientset(),
				"default",
				"",
				"",
				false,
				"",
				true,
				false,
				false,
				ts.svcFilter,
				false,
				labels.Everything(),
				false,
				false,
				false,
				false,
			)
			require.NoError(t, err)
			svcSrc, ok := svc.(*serviceSource)
			if !ok {
				require.Fail(t, "expected serviceSource")
			}
			ts.asserts(svcSrc)
		})
	}
}

func TestNewServiceSourceWithServiceTypeFilters_Unsupported(t *testing.T) {
	serviceTypeFilter := []string{"ClusterIP", "ServiceTypeNotExist"}

	svc, err := NewServiceSource(
		context.TODO(),
		fake.NewClientset(),
		"default",
		"",
		"",
		false,
		"",
		false,
		false,
		false,
		serviceTypeFilter,
		false,
		labels.Everything(),
		false,
		false,
		false,
		false,
	)
	require.Errorf(t, err, "unsupported service type filter: \"UnknownType\". Supported types are: [\"ClusterIP\" \"NodePort\" \"LoadBalancer\" \"ExternalName\"]")
	require.Nil(t, svc, "ServiceSource should be nil when an unsupported service type is provided")
}

func TestNewServiceTypes(t *testing.T) {
	tests := []struct {
		name        string
		filter      []string
		wantEnabled bool
		wantTypes   map[v1.ServiceType]bool
		wantErr     bool
	}{
		{
			name:        "empty filter disables serviceTypes",
			filter:      []string{},
			wantEnabled: false,
			wantTypes:   nil,
			wantErr:     false,
		},
		{
			name:        "filter with empty string disables serviceTypes",
			filter:      []string{""},
			wantEnabled: false,
			wantTypes:   nil,
			wantErr:     false,
		},
		{
			name:        "valid filter enables serviceTypes",
			filter:      []string{string(v1.ServiceTypeClusterIP), string(v1.ServiceTypeNodePort)},
			wantEnabled: true,
			wantTypes: map[v1.ServiceType]bool{
				v1.ServiceTypeClusterIP: true,
				v1.ServiceTypeNodePort:  true,
			},
			wantErr: false,
		},
		{
			name:        "filter with unknown type returns error",
			filter:      []string{"UnknownType"},
			wantEnabled: false,
			wantTypes:   nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st, err := newServiceTypesFilter(tt.filter)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, st)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantEnabled, st.enabled)
				if tt.wantTypes != nil {
					assert.Equal(t, tt.wantTypes, st.types)
				}
			}
		})
	}
}

func TestFilterByServiceType_WithFixture(t *testing.T) {
	tests := []struct {
		name            string
		filter          *serviceTypes
		currentServices []*v1.Service
		expected        int
	}{
		{
			name: "all types of services with filter enabled for ServiceTypeNodePort and ServiceTypeClusterIP",
			currentServices: createTestServicesByType("kube-system", map[v1.ServiceType]int{
				v1.ServiceTypeLoadBalancer: 3,
				v1.ServiceTypeNodePort:     4,
				v1.ServiceTypeClusterIP:    5,
				v1.ServiceTypeExternalName: 2,
			}),
			filter: &serviceTypes{
				enabled: true,
				types: map[v1.ServiceType]bool{
					v1.ServiceTypeNodePort:  true,
					v1.ServiceTypeClusterIP: true,
				},
			},
			expected: 4 + 5,
		},
		{
			name: "all types of services with filter enabled for ServiceTypeLoadBalancer",
			currentServices: createTestServicesByType("default", map[v1.ServiceType]int{
				v1.ServiceTypeLoadBalancer: 3,
				v1.ServiceTypeNodePort:     4,
				v1.ServiceTypeClusterIP:    5,
				v1.ServiceTypeExternalName: 2,
			}),
			filter: &serviceTypes{
				enabled: true,
				types: map[v1.ServiceType]bool{
					v1.ServiceTypeLoadBalancer: true,
				},
			},
			expected: 3,
		},
		{
			name: "enabled for ServiceTypeLoadBalancer when not all types are present",
			currentServices: createTestServicesByType("default", map[v1.ServiceType]int{
				v1.ServiceTypeNodePort:     4,
				v1.ServiceTypeClusterIP:    5,
				v1.ServiceTypeExternalName: 2,
			}),
			filter: &serviceTypes{
				enabled: true,
				types: map[v1.ServiceType]bool{
					v1.ServiceTypeLoadBalancer: true,
				},
			},
			expected: 0,
		},
		{
			name: "filter disabled returns all services",
			currentServices: createTestServicesByType("default", map[v1.ServiceType]int{
				v1.ServiceTypeLoadBalancer: 3,
				v1.ServiceTypeNodePort:     4,
				v1.ServiceTypeClusterIP:    5,
				v1.ServiceTypeExternalName: 2,
			}),
			filter: &serviceTypes{
				enabled: false,
				types:   map[v1.ServiceType]bool{},
			},
			expected: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &serviceSource{serviceTypeFilter: tt.filter}
			assert.NotNil(t, sc)
			got := sc.filterByServiceType(tt.currentServices)
			assert.Len(t, got, tt.expected)
		})
	}
}

func TestPodTransformerInServiceSource(t *testing.T) {
	ctx := t.Context()
	fakeClient := fake.NewClientset()

	es := &discoveryv1.EndpointSlice{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-ns",
			Name:      "test-slice",
			Labels:    map[string]string{discoveryv1.LabelServiceName: "test-service"},
			Annotations: map[string]string{
				"some-annotation": "value",
			},
			UID: "esuid",
			ManagedFields: []metav1.ManagedFieldsEntry{
				{Manager: "endpointslice-controller", Operation: metav1.ManagedFieldsOperationUpdate},
			},
		},
		AddressType: discoveryv1.AddressTypeIPv4,
		Endpoints:   []discoveryv1.Endpoint{{Addresses: []string{"10.0.0.1"}}},
	}

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{{
				Name: "test",
			}},
			Hostname: "test-hostname",
			NodeName: "test-node",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test-ns",
			Name:      "test-name",
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
		Status: v1.PodStatus{
			PodIP:  "127.0.0.1",
			HostIP: "127.0.0.2",
			Conditions: []v1.PodCondition{{
				Type:   v1.PodReady,
				Status: v1.ConditionTrue,
			}, {
				Type:   v1.ContainersReady,
				Status: v1.ConditionFalse,
			}},
		},
	}

	_, err := fakeClient.CoreV1().Pods(pod.Namespace).Create(context.Background(), pod, metav1.CreateOptions{})
	require.NoError(t, err)
	_, err = fakeClient.DiscoveryV1().EndpointSlices(es.Namespace).Create(context.Background(), es, metav1.CreateOptions{})
	require.NoError(t, err)

	// Should not error when creating the source
	src, err := NewServiceSource(
		ctx,
		fakeClient,
		"",
		"",
		"{{.Name}}",
		false,
		"",
		false,
		false,
		false,
		[]string{},
		false,
		labels.Everything(),
		false,
		false,
		false,
		false,
	)
	require.NoError(t, err)
	ss, ok := src.(*serviceSource)
	require.True(t, ok)

	retrieved, err := ss.podInformer.Lister().Pods("test-ns").Get("test-name")
	require.NoError(t, err)

	// Metadata
	assert.Equal(t, "test-name", retrieved.Name)
	assert.Equal(t, "test-ns", retrieved.Namespace)
	assert.Equal(t, pod.UID, retrieved.UID)
	assert.Equal(t, pod.Labels, retrieved.Labels)
	// Annotations filtered to external-dns prefix only
	assert.Equal(t, map[string]string{
		"external-dns.alpha.kubernetes.io/hostname": "test-hostname",
		"external-dns.alpha.kubernetes.io/random":   "value",
	}, retrieved.Annotations)

	// Spec — fully preserved
	assert.NotEmpty(t, retrieved.Spec.Containers)
	assert.Equal(t, "test-hostname", retrieved.Spec.Hostname)
	assert.Equal(t, "test-node", retrieved.Spec.NodeName)

	// Status — fully preserved
	assert.Equal(t, "127.0.0.2", retrieved.Status.HostIP)
	assert.Equal(t, "127.0.0.1", retrieved.Status.PodIP)
	assert.ElementsMatch(t, []v1.PodCondition{{
		Type:   v1.PodReady,
		Status: v1.ConditionTrue,
	}, {
		Type:   v1.ContainersReady,
		Status: v1.ConditionFalse,
	}}, retrieved.Status.Conditions)

	// EndpointSlice — managedFields stripped, everything else preserved
	retrievedES, err := ss.endpointSlicesInformer.Lister().EndpointSlices(es.Namespace).Get(es.Name)
	require.NoError(t, err)
	assert.Empty(t, retrievedES.ManagedFields)
	assert.Equal(t, es.Labels, retrievedES.Labels)
	assert.Equal(t, es.Annotations, retrievedES.Annotations)
	assert.Equal(t, es.Endpoints, retrievedES.Endpoints)
}

func TestNodeTransformerInServiceSource(t *testing.T) {
	ctx := t.Context()

	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-node",
			Labels: map[string]string{
				"label1": "value1",
			},
			Annotations: map[string]string{
				"user-annotation":              "value",
				v1.LastAppliedConfigAnnotation: `{"apiVersion":"v1"}`,
			},
			UID: "someuid",
			ManagedFields: []metav1.ManagedFieldsEntry{
				{Manager: "kubectl", Operation: metav1.ManagedFieldsOperationApply},
			},
		},
		Status: v1.NodeStatus{
			Addresses: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "1.2.3.4"},
			},
			Conditions: []v1.NodeCondition{
				{Type: v1.NodeReady, Status: v1.ConditionTrue},
			},
		},
	}

	fakeClient := fake.NewClientset(node)

	src, err := NewServiceSource(
		ctx, fakeClient, "", "", "", false, "", false, false, false,
		[]string{string(v1.ServiceTypeNodePort)},
		false, labels.Everything(), false, false, false, false,
	)
	require.NoError(t, err)
	ss, ok := src.(*serviceSource)
	require.True(t, ok)

	retrieved, err := ss.nodeInformer.Lister().Get(node.Name)
	require.NoError(t, err)

	assert.Equal(t, node.Name, retrieved.Name)
	assert.Equal(t, node.Labels, retrieved.Labels)
	assert.Equal(t, node.UID, retrieved.UID)
	assert.Empty(t, retrieved.ManagedFields)
	assert.NotContains(t, retrieved.Annotations, v1.LastAppliedConfigAnnotation)
	assert.Contains(t, retrieved.Annotations, "user-annotation")
	// Status.Addresses preserved — used for endpoint generation
	assert.Equal(t, node.Status.Addresses, retrieved.Status.Addresses)
	// Status.Conditions stripped
	assert.Empty(t, retrieved.Status.Conditions)
}

// createTestServicesByType creates the requested number of services per type in the given namespace.
func createTestServicesByType(ns string, typeCounts map[v1.ServiceType]int) []*v1.Service {
	var services []*v1.Service
	idx := 0
	for svcType, count := range typeCounts {
		for range count {
			svc := &v1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("svc-%s-%d", svcType, idx),
					Namespace: ns,
				},
				Spec: v1.ServiceSpec{
					Type: svcType,
				},
			}
			if svcType == v1.ServiceTypeExternalName {
				svc.Spec.ExternalName = fmt.Sprintf("external-%d.example.com", idx)
			}
			services = append(services, svc)
			idx++
		}
	}
	// Shuffle the resulting services to ensure randomness in the order.
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(services), func(i, j int) {
		services[i], services[j] = services[j], services[i]
	})
	return services
}

func TestServiceTypes_isNodeInformerRequired(t *testing.T) {
	tests := []struct {
		name     string
		filter   []string
		required []v1.ServiceType
		want     bool
	}{
		{
			name:     "NodePort required and filter is empty",
			filter:   []string{},
			required: []v1.ServiceType{v1.ServiceTypeNodePort},
			want:     true,
		},
		{
			name:     "NodePort type present",
			filter:   []string{string(v1.ServiceTypeNodePort)},
			required: []v1.ServiceType{v1.ServiceTypeNodePort},
			want:     true,
		},
		{
			name:     "NodePort type absent, filter enabled",
			filter:   []string{string(v1.ServiceTypeLoadBalancer)},
			required: []v1.ServiceType{v1.ServiceTypeNodePort},
			want:     false,
		},
		{
			name:     "NodePort and other filters present",
			filter:   []string{string(v1.ServiceTypeLoadBalancer), string(v1.ServiceTypeNodePort)},
			required: []v1.ServiceType{v1.ServiceTypeNodePort},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter, _ := newServiceTypesFilter(tt.filter)
			got := filter.isRequired(tt.required...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestServiceSource_AddEventHandler(t *testing.T) {
	var fakeServiceInformer *fakes.FakeServiceInformer
	var fakeEdpInformer *fakes.FakeEndpointSliceInformer
	var fakeNodeInformer *fakes.FakeNodeInformer
	tests := []struct {
		name    string
		filter  []string
		times   int
		asserts func(t *testing.T)
	}{
		{
			name:   "AddEventHandler should trigger all event handlers when empty filter is provided",
			filter: []string{},
			times:  2,
			asserts: func(t *testing.T) {
				fakeServiceInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeEdpInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeNodeInformer.AssertNumberOfCalls(t, "Informer", 0)
			},
		},
		{
			name:   "AddEventHandler should trigger only service event handler",
			filter: []string{string(v1.ServiceTypeExternalName), string(v1.ServiceTypeLoadBalancer)},
			times:  1,
			asserts: func(t *testing.T) {
				fakeServiceInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeEdpInformer.AssertNumberOfCalls(t, "Informer", 0)
				fakeNodeInformer.AssertNumberOfCalls(t, "Informer", 0)
			},
		},
		{
			name:   "AddEventHandler should configure only service event handler",
			filter: []string{string(v1.ServiceTypeExternalName), string(v1.ServiceTypeLoadBalancer), string(v1.ServiceTypeClusterIP)},
			times:  2,
			asserts: func(t *testing.T) {
				fakeServiceInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeEdpInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeNodeInformer.AssertNumberOfCalls(t, "Informer", 0)
			},
		},
		{
			name:   "AddEventHandler should configure all service event handlers",
			filter: []string{string(v1.ServiceTypeNodePort)},
			times:  2,
			asserts: func(t *testing.T) {
				fakeServiceInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeEdpInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeNodeInformer.AssertNumberOfCalls(t, "Informer", 0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeServiceInformer = new(fakes.FakeServiceInformer)
			infSvc := testInformer{}
			fakeServiceInformer.On("Informer").Return(&infSvc)

			fakeEdpInformer = new(fakes.FakeEndpointSliceInformer)
			infEdp := testInformer{}
			fakeEdpInformer.On("Informer").Return(&infEdp)

			fakeNodeInformer = new(fakes.FakeNodeInformer)
			infNode := testInformer{}
			fakeNodeInformer.On("Informer").Return(&infNode)

			filter, _ := newServiceTypesFilter(tt.filter)

			svcSource := &serviceSource{
				endpointSlicesInformer: fakeEdpInformer,
				serviceInformer:        fakeServiceInformer,
				nodeInformer:           fakeNodeInformer,
				serviceTypeFilter:      filter,
				listenEndpointEvents:   true,
			}

			svcSource.AddEventHandler(t.Context(), func() {})

			assert.Equal(t, tt.times, infSvc.times+infEdp.times+infNode.times)

			tt.asserts(t)
		})
	}
}

// Test helper functions created during extractHeadlessEndpoints refactoring
func TestConvertToEndpointSlices(t *testing.T) {
	t.Run("converts valid EndpointSlices", func(t *testing.T) {
		validSlice := &discoveryv1.EndpointSlice{
			ObjectMeta:  metav1.ObjectMeta{Name: "valid-slice"},
			AddressType: discoveryv1.AddressTypeIPv4,
		}

		rawObjects := []any{validSlice}
		result := convertToEndpointSlices(rawObjects)

		assert.Len(t, result, 1)
		assert.Equal(t, "valid-slice", result[0].Name)
	})

	t.Run("skips invalid objects", func(t *testing.T) {
		invalidObject := "not-an-endpoint-slice"
		validSlice := &discoveryv1.EndpointSlice{
			ObjectMeta:  metav1.ObjectMeta{Name: "valid-slice"},
			AddressType: discoveryv1.AddressTypeIPv4,
		}

		rawObjects := []any{invalidObject, validSlice}
		result := convertToEndpointSlices(rawObjects)

		assert.Len(t, result, 1)
		assert.Equal(t, "valid-slice", result[0].Name)
	})

	t.Run("handles empty input", func(t *testing.T) {
		result := convertToEndpointSlices([]any{})
		assert.Empty(t, result)
	})

	t.Run("handles all invalid objects", func(t *testing.T) {
		rawObjects := []any{"invalid1", 123, map[string]string{"key": "value"}}
		result := convertToEndpointSlices(rawObjects)
		assert.Empty(t, result)
	})
}

// Test for processEndpointSlice: publishPodIPs true and pod == nil

func TestProcessEndpointSlices_PublishPodIPsPodNil(t *testing.T) {
	sc := &serviceSource{}

	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "slice1", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeIPv4,
		Endpoints: []discoveryv1.Endpoint{
			{
				TargetRef:  &v1.ObjectReference{Kind: "Pod", Name: "missing-pod"},
				Conditions: discoveryv1.EndpointConditions{Ready: testutils.ToPtr(true)},
			},
		},
	}
	pods := []*v1.Pod{} // No pods, so pod == nil
	hostname := "test.example.com"
	endpointsType := "IPv4"
	publishPodIPs := true
	publishNotReadyAddresses := false

	result := sc.processHeadlessEndpointsFromSlices(
		pods, []*discoveryv1.EndpointSlice{endpointSlice},
		hostname, endpointsType, publishPodIPs, publishNotReadyAddresses)
	assert.Empty(t, result, "No targets should be added when pod is nil and publishPodIPs is true")
}

// Test for processEndpointSlice: publishPodIPs true and unsupported address type triggers log.Debugf skip
func TestProcessEndpointSlices_PublishPodIPsUnsupportedAddressType(t *testing.T) {
	sc := &serviceSource{}

	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "slice2", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeFQDN, // unsupported type
		Endpoints: []discoveryv1.Endpoint{
			{
				TargetRef:  &v1.ObjectReference{Kind: "Pod", Name: "some-pod"},
				Conditions: discoveryv1.EndpointConditions{Ready: testutils.ToPtr(true)},
			},
		},
	}
	pods := []*v1.Pod{{ObjectMeta: metav1.ObjectMeta{Name: "some-pod"}}}
	hostname := "test.example.com"
	endpointsType := "FQDN"
	publishPodIPs := true
	publishNotReadyAddresses := false

	result := sc.processHeadlessEndpointsFromSlices(
		pods, []*discoveryv1.EndpointSlice{endpointSlice},
		hostname, endpointsType, publishPodIPs, publishNotReadyAddresses)
	assert.Empty(t, result, "No targets should be added for unsupported address type when publishPodIPs is true")
}

// Test for missing coverage: publishPodIPs false scenario
func TestProcessEndpointSlices_PublishPodIPsFalse(t *testing.T) {
	sc := &serviceSource{}

	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "slice1", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeIPv4,
		Endpoints: []discoveryv1.Endpoint{
			{
				TargetRef:  &v1.ObjectReference{Kind: "Pod", Name: "test-pod"},
				Conditions: discoveryv1.EndpointConditions{Ready: testutils.ToPtr(true)},
				Addresses:  []string{"10.0.0.1"},
			},
		},
	}
	pods := []*v1.Pod{{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
		Status:     v1.PodStatus{PodIP: "10.0.0.1"},
	}}
	hostname := "test.example.com"
	endpointsType := "IPv4"
	publishPodIPs := false // This should allow processing
	publishNotReadyAddresses := false

	result := sc.processHeadlessEndpointsFromSlices(
		pods, []*discoveryv1.EndpointSlice{endpointSlice},
		hostname, endpointsType, publishPodIPs, publishNotReadyAddresses)
	assert.NotEmpty(t, result, "Targets should be added when publishPodIPs is false")
}

// Test for missing coverage: not ready endpoints with publishNotReadyAddresses true
func TestProcessEndpointSlices_NotReadyWithPublishNotReady(t *testing.T) {
	sc := &serviceSource{}

	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "slice1", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeIPv4,
		Endpoints: []discoveryv1.Endpoint{
			{
				TargetRef:  &v1.ObjectReference{Kind: "Pod", Name: "test-pod"},
				Conditions: discoveryv1.EndpointConditions{Ready: testutils.ToPtr(false)}, // Not ready
				Addresses:  []string{"10.0.0.1"},
			},
		},
	}
	pods := []*v1.Pod{{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
		Status:     v1.PodStatus{PodIP: "10.0.0.1"},
	}}
	hostname := "test.example.com"
	endpointsType := "IPv4"
	publishPodIPs := false
	publishNotReadyAddresses := true // This should allow not-ready endpoints

	result := sc.processHeadlessEndpointsFromSlices(
		pods, []*discoveryv1.EndpointSlice{endpointSlice},
		hostname, endpointsType, publishPodIPs, publishNotReadyAddresses)
	assert.NotEmpty(t, result, "Not ready endpoints should be processed when publishNotReadyAddresses is true")
}

// Test getTargetsForDomain with empty ep.Addresses
func TestGetTargetsForDomain_EmptyAddresses(t *testing.T) {
	sc := &serviceSource{}
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
		Status:     v1.PodStatus{PodIP: "10.0.0.1"},
	}
	ep := discoveryv1.Endpoint{
		Addresses: []string{}, // Empty addresses
	}
	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "test-slice", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeIPv4,
	}
	endpointsType := "IPv4"
	headlessDomain := "test.example.com"

	targets := sc.getTargetsForDomain(pod, ep, endpointSlice, endpointsType, headlessDomain)
	assert.Empty(t, targets, "Should return empty targets when ep.Addresses is empty")
}

// Test getTargetsForDomain with EndpointsTypeHostIP
func TestGetTargetsForDomain_HostIP(t *testing.T) {
	sc := &serviceSource{publishHostIP: false}
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
		Status:     v1.PodStatus{HostIP: "192.168.1.100", PodIP: "10.0.0.1"},
	}
	ep := discoveryv1.Endpoint{
		Addresses: []string{"10.0.0.1"},
	}
	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "test-slice", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeIPv4,
	}
	endpointsType := EndpointsTypeHostIP
	headlessDomain := "test.example.com"

	targets := sc.getTargetsForDomain(pod, ep, endpointSlice, endpointsType, headlessDomain)
	assert.Contains(t, targets, "192.168.1.100", "Should return HostIP when endpointsType is HostIP")
}

// Test getTargetsForDomain with NodeExternalIP and nodeInformer
func TestGetTargetsForDomain_NodeExternalIP(t *testing.T) {
	// Create a fake node informer with a node
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{Name: "test-node"},
		Status: v1.NodeStatus{
			Addresses: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "203.0.113.10"},
				{Type: v1.NodeInternalIP, Address: "10.0.0.10"},
			},
		},
	}

	client := fake.NewSimpleClientset(node)
	kubeInformers := kubeinformers.NewSharedInformerFactory(client, 0)
	nodeInformer := kubeInformers.Core().V1().Nodes()

	// Add the node to the informer
	nodeInformer.Informer().GetStore().Add(node)

	sc := &serviceSource{
		nodeInformer: nodeInformer,
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
		Spec:       v1.PodSpec{NodeName: "test-node"},
		Status:     v1.PodStatus{PodIP: "10.0.0.1"},
	}
	ep := discoveryv1.Endpoint{
		Addresses: []string{"10.0.0.1"},
	}
	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "test-slice", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeIPv4,
	}
	endpointsType := EndpointsTypeNodeExternalIP
	headlessDomain := "test.example.com"

	targets := sc.getTargetsForDomain(pod, ep, endpointSlice, endpointsType, headlessDomain)
	assert.Contains(t, targets, "203.0.113.10", "Should return NodeExternalIP")
}

func TestFindPodForEndpoint(t *testing.T) {
	pods := []*v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{Name: "pod1"},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Name: "pod2"},
		},
	}

	t.Run("finds matching pod", func(t *testing.T) {
		endpoint := discoveryv1.Endpoint{
			TargetRef: &v1.ObjectReference{
				Kind: "Pod",
				Name: "pod1",
			},
		}

		result := findPodForEndpoint(endpoint, pods)
		assert.NotNil(t, result)
		assert.Equal(t, "pod1", result.Name)
	})

	t.Run("returns nil for nil TargetRef", func(t *testing.T) {
		endpoint := discoveryv1.Endpoint{
			TargetRef: nil,
		}

		result := findPodForEndpoint(endpoint, pods)
		assert.Nil(t, result)
	})

	t.Run("returns nil for non-Pod kind", func(t *testing.T) {
		endpoint := discoveryv1.Endpoint{
			TargetRef: &v1.ObjectReference{
				Kind: "Service",
				Name: "pod1",
			},
		}

		result := findPodForEndpoint(endpoint, pods)
		assert.Nil(t, result)
	})

	t.Run("returns nil for non-empty APIVersion", func(t *testing.T) {
		endpoint := discoveryv1.Endpoint{
			TargetRef: &v1.ObjectReference{
				Kind:       "Pod",
				Name:       "pod1",
				APIVersion: "v1",
			},
		}

		result := findPodForEndpoint(endpoint, pods)
		assert.Nil(t, result)
	})

	t.Run("returns nil for non-existent pod", func(t *testing.T) {
		endpoint := discoveryv1.Endpoint{
			TargetRef: &v1.ObjectReference{
				Kind: "Pod",
				Name: "non-existent-pod",
			},
		}

		result := findPodForEndpoint(endpoint, pods)
		assert.Nil(t, result)
	})
}

func TestBuildHeadlessEndpoints(t *testing.T) {
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-service",
			Namespace: "default",
		},
	}

	t.Run("builds endpoints from targets", func(t *testing.T) {
		targetsByHeadlessDomainAndType := map[endpoint.EndpointKey]endpoint.Targets{
			{DNSName: "test.example.com", RecordType: endpoint.RecordTypeA}:    {"1.2.3.4", "5.6.7.8"},
			{DNSName: "test.example.com", RecordType: endpoint.RecordTypeAAAA}: {"2001:db8::1"},
		}

		result := buildHeadlessEndpoints(svc, targetsByHeadlessDomainAndType, endpoint.TTL(0))

		assert.Len(t, result, 2)

		// Check A record
		aRecord := findEndpointByType(result, endpoint.RecordTypeA)
		assert.NotNil(t, aRecord)
		assert.Equal(t, "test.example.com", aRecord.DNSName)
		assert.Contains(t, aRecord.Targets, "1.2.3.4")
		assert.Contains(t, aRecord.Targets, "5.6.7.8")
		assert.Equal(t, "service/default/test-service", aRecord.Labels[endpoint.ResourceLabelKey])

		// Check AAAA record
		aaaaRecord := findEndpointByType(result, endpoint.RecordTypeAAAA)
		assert.NotNil(t, aaaaRecord)
		assert.Equal(t, "test.example.com", aaaaRecord.DNSName)
		assert.Contains(t, aaaaRecord.Targets, "2001:db8::1")
	})

	t.Run("deduplicates targets", func(t *testing.T) {
		targetsByHeadlessDomainAndType := map[endpoint.EndpointKey]endpoint.Targets{
			{DNSName: "test.example.com", RecordType: endpoint.RecordTypeA}: {"1.2.3.4", "1.2.3.4", "5.6.7.8"},
		}

		result := buildHeadlessEndpoints(svc, targetsByHeadlessDomainAndType, endpoint.TTL(0))

		assert.Len(t, result, 1)
		assert.Len(t, result[0].Targets, 2)
		assert.Contains(t, result[0].Targets, "1.2.3.4")
		assert.Contains(t, result[0].Targets, "5.6.7.8")
	})

	t.Run("handles TTL configuration", func(t *testing.T) {
		targetsByHeadlessDomainAndType := map[endpoint.EndpointKey]endpoint.Targets{
			{DNSName: "test.example.com", RecordType: endpoint.RecordTypeA}: {"1.2.3.4"},
		}

		result := buildHeadlessEndpoints(svc, targetsByHeadlessDomainAndType, endpoint.TTL(300))

		assert.Len(t, result, 1)
		assert.Equal(t, endpoint.TTL(300), result[0].RecordTTL)
	})

	t.Run("sorts endpoints deterministically", func(t *testing.T) {
		targetsByHeadlessDomainAndType := map[endpoint.EndpointKey]endpoint.Targets{
			{DNSName: "z.example.com", RecordType: endpoint.RecordTypeA}:    {"1.2.3.4"},
			{DNSName: "a.example.com", RecordType: endpoint.RecordTypeA}:    {"5.6.7.8"},
			{DNSName: "a.example.com", RecordType: endpoint.RecordTypeAAAA}: {"2001:db8::1"},
		}

		result := buildHeadlessEndpoints(svc, targetsByHeadlessDomainAndType, endpoint.TTL(0))

		assert.Len(t, result, 3)
		// Should be sorted by DNSName first, then by RecordType
		assert.Equal(t, "a.example.com", result[0].DNSName)
		assert.Equal(t, endpoint.RecordTypeA, result[0].RecordType)
		assert.Equal(t, "a.example.com", result[1].DNSName)
		assert.Equal(t, endpoint.RecordTypeAAAA, result[1].RecordType)
		assert.Equal(t, "z.example.com", result[2].DNSName)
	})

	t.Run("handles empty targets", func(t *testing.T) {
		result := buildHeadlessEndpoints(svc, map[endpoint.EndpointKey]endpoint.Targets{}, endpoint.TTL(0))
		assert.Empty(t, result)
	})
}

// Test for missing coverage: pod with hostname creates additional headless domains
func TestProcessEndpointSlices_PodWithHostname(t *testing.T) {
	sc := &serviceSource{}

	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "slice1", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeIPv4,
		Endpoints: []discoveryv1.Endpoint{
			{
				TargetRef:  &v1.ObjectReference{Kind: "Pod", Name: "test-pod"},
				Conditions: discoveryv1.EndpointConditions{Ready: testutils.ToPtr(true)},
				Addresses:  []string{"10.0.0.1"},
			},
		},
	}
	pods := []*v1.Pod{{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod"},
		Spec:       v1.PodSpec{Hostname: "my-pod"}, // Non-empty hostname
		Status:     v1.PodStatus{PodIP: "10.0.0.1"},
	}}
	hostname := "test.example.com"
	endpointsType := "IPv4"
	publishPodIPs := false
	publishNotReadyAddresses := false

	result := sc.processHeadlessEndpointsFromSlices(
		pods, []*discoveryv1.EndpointSlice{endpointSlice},
		hostname, endpointsType, publishPodIPs, publishNotReadyAddresses)

	assert.NotEmpty(t, result, "Should create targets for pod with hostname")

	// Check that both the base hostname and pod-specific hostname are created
	var foundBaseHostname, foundPodHostname bool
	for key := range result {
		if key.DNSName == "test.example.com" {
			foundBaseHostname = true
		}
		if key.DNSName == "my-pod.test.example.com" {
			foundPodHostname = true
		}
	}

	assert.True(t, foundBaseHostname, "Should create endpoint for base hostname")
	assert.True(t, foundPodHostname, "Should create endpoint for pod-specific hostname when pod.Spec.Hostname is set")
}

// Helper function to find endpoint by record type
func findEndpointByType(endpoints []*endpoint.Endpoint, recordType string) *endpoint.Endpoint {
	for _, ep := range endpoints {
		if ep.RecordType == recordType {
			return ep
		}
	}
	return nil
}
