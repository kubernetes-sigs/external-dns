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
	"net"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/fake"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/informers"
)

// pointerToBool returns a pointer to the given bool value.
func pointerToBool(b bool) *bool {
	return &b
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org., bar.example.org.",
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
				hostnameAnnotationKey: "foo.example.org., bar.example.org.",
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
				hostnameAnnotationKey: "foo.example.org., bar.example.org.",
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
				hostnameAnnotationKey: "foo.example.org, bar.example.org",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org", // Trailing dot is omitted
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
				controllerAnnotationKey: controllerAnnotationValue,
				hostnameAnnotationKey:   "foo.example.org.",
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
				controllerAnnotationKey: "some-other-tool",
				hostnameAnnotationKey:   "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey:                         "foo.example.org.",
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
				hostnameAnnotationKey:                         "foo.example.org.",
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
				hostnameAnnotationKey:                         "foo.example.org.",
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
				hostnameAnnotationKey:                         "foo.example.org.",
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
				hostnameAnnotationKey:                         "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
				ttlAnnotationKey:      "foo",
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
				hostnameAnnotationKey: "foo.example.org.",
				ttlAnnotationKey:      "10",
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
				hostnameAnnotationKey: "foo.example.org.",
				ttlAnnotationKey:      "1m",
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
				hostnameAnnotationKey: "foo.example.org.",
				ttlAnnotationKey:      "-10",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey:         "foo.example.org.",
				internalHostnameAnnotationKey: "foo.internal.example.org.",
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
				internalHostnameAnnotationKey: "foo.internal.example.org.",
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
				hostnameAnnotationKey:         "foo.example.org.",
				internalHostnameAnnotationKey: "foo.internal.example.org.",
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
			annotations:          map[string]string{hostnameAnnotationKey: "annotation.bar.example.com"},
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
			annotations:          map[string]string{hostnameAnnotationKey: "annotation.bar.example.com"},
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
			annotations:        map[string]string{hostnameAnnotationKey: "foobar.example.org"},
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
			annotations:        map[string]string{hostnameAnnotationKey: "foobar-v6.example.org"},
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
		labelSelector               string
		resolveLoadBalancerHostname bool
	}{
		{
			title:                    "multiple services that share same DNS should be merged into one endpoint",
			targetNamespace:          "",
			annotationFilter:         "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			compatibility:            "",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			labels:                   map[string]string{},
			annotations:              map[string]string{},
			clusterIP:                "",
			externalIPs:              []string{},
			lbs:                      []string{},
			serviceTypesFilter:       []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
			},
			expectError:                 false,
			labelSelector:               "",
			resolveLoadBalancerHostname: false,
		},
		{
			title:                    "multiple services that share same DNS should be merged into one endpoint",
			targetNamespace:          "",
			annotationFilter:         "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			compatibility:            "",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			labels:                   map[string]string{},
			annotations:              map[string]string{},
			clusterIP:                "",
			externalIPs:              []string{},
			lbs:                      []string{"1.2.3.4", "1.2.3.5", "1.2.3.6"},
			serviceTypesFilter:       []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5", "1.2.3.6"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
			},
			expectError:                 false,
			labelSelector:               "",
			resolveLoadBalancerHostname: false,
		},
		{
			title:                    "test that services with different hostnames do not get merged together",
			targetNamespace:          "",
			annotationFilter:         "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			compatibility:            "",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			labels:                   map[string]string{},
			annotations:              map[string]string{},
			clusterIP:                "",
			externalIPs:              []string{},
			lbs:                      []string{"1.2.3.5", "10.1.1.3", "10.1.1.1", "1.2.3.4", "10.1.1.2", "20.1.1.1", "1.2.3.6"},
			serviceTypesFilter:       []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5", "1.2.3.6"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.1", "10.1.1.2", "10.1.1.3"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo10.1.1.1"}},
				{DNSName: "foobar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"20.1.1.1"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo20.1.1.1"}},
			},
			expectError:                 false,
			labelSelector:               "",
			resolveLoadBalancerHostname: false,
		},
		{
			title:                    "test that services with different hostnames do not get merged together",
			targetNamespace:          "",
			annotationFilter:         "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			compatibility:            "",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			labels:                   map[string]string{},
			annotations:              map[string]string{},
			clusterIP:                "",
			externalIPs:              []string{},
			lbs:                      []string{"1.2.3.5", "10.1.1.3", "10.1.1.1", "1.2.3.4", "10.1.1.2", "20.1.1.1", "1.2.3.6"},
			serviceTypesFilter:       []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5", "1.2.3.6"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.1", "10.1.1.2", "10.1.1.3"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo10.1.1.1"}},
				{DNSName: "foobar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"20.1.1.1"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo20.1.1.1"}},
			},
			expectError:                 false,
			labelSelector:               "",
			resolveLoadBalancerHostname: false,
		},
		{
			title:                    "test that services with CNAME types do not get merged together",
			targetNamespace:          "",
			annotationFilter:         "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			svcType:                  v1.ServiceTypeLoadBalancer,
			compatibility:            "",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			labels:                   map[string]string{},
			annotations:              map[string]string{},
			clusterIP:                "",
			externalIPs:              []string{},
			lbs:                      []string{"a.elb.com", "b.elb.com"},
			serviceTypesFilter:       []string{},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"a.elb.com"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/fooa.elb.com"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"b.elb.com"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foob.elb.com"}},
			},
			expectError:                 false,
			labelSelector:               "",
			resolveLoadBalancerHostname: false,
		},
	} {

		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			// Create a Kubernetes testing client
			kubernetes := fake.NewClientset()

			// Create services to test against
			for _, lb := range tc.lbs {
				ingresses := []v1.LoadBalancerIngress{{IP: lb}}
				service := &v1.Service{
					Spec: v1.ServiceSpec{
						Type:      tc.svcType,
						ClusterIP: tc.clusterIP,
					},
					ObjectMeta: metav1.ObjectMeta{
						Namespace:   tc.svcNamespace,
						Name:        tc.svcName + lb,
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
		labelSelector               string
		resolveLoadBalancerHostname bool
	}{
		{
			title:        "ClusterIp services return an endpoint with the specified A",
			svcNamespace: "testing",
			svcName:      "foo",
			svcType:      v1.ServiceTypeClusterIP,
			annotations: map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
				targetAnnotationKey:   "4.3.2.1",
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
				hostnameAnnotationKey: "foo.example.org.",
				targetAnnotationKey:   "bar.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
				targetAnnotationKey:   "2001:DB8::1",
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
				hostnameAnnotationKey: "foo.example.org.",
				targetAnnotationKey:   "bar.example.org.,baz.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
				targetAnnotationKey:   "bar.example.org.,baz.example.org.,2001:DB8::1",
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
				hostnameAnnotationKey: "foo.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
				targetAnnotationKey:   "bar.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
				targetAnnotationKey:   "bar.example.org.",
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
				hostnameAnnotationKey: "foo.example.org.",
				targetAnnotationKey:   "bar.example.org.",
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
			annotations:  map[string]string{targetAnnotationKey: "bar.example.com."},
			clusterIP:    "4.5.6.7",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.bar.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"bar.example.com"}},
			},
			labelSelector: "app=web-internal",
		},
		{
			title:         "ClusterIP service without matching label does not generate an endpoint",
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
				hostnameAnnotationKey: "this-is-an-exceedingly-long-label-that-external-dns-should-reject.example.org.",
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

// Test for processHeadlessEndpointSlices: not ready and publishNotReadyAddresses false
func TestProcessHeadlessEndpointSlices_NotReadyNotPublished(t *testing.T) {
	sc := &serviceSource{}
	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "slice3", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeIPv4,
		Endpoints: []discoveryv1.Endpoint{
			{
				TargetRef:  &v1.ObjectReference{Kind: "Pod", Name: "some-pod"},
				Conditions: discoveryv1.EndpointConditions{Ready: pointerToBool(false)},
			},
		},
	}
	pods := []*v1.Pod{{ObjectMeta: metav1.ObjectMeta{Name: "some-pod"}}}
	params := HeadlessEndpointParams{
		sc:                       sc,
		endpointSlices:           []*discoveryv1.EndpointSlice{endpointSlice},
		pods:                     pods,
		hostname:                 "test.example.com",
		endpointsType:            "IPv4",
		publishPodIPs:            true,
		publishNotReadyAddresses: false,
	}
	result := processHeadlessEndpointSlices(params)
	assert.Empty(t, result, "No targets should be added when endpoint is not ready and publishNotReadyAddresses is false")
}

// Test for processHeadlessEndpointSlices: not ready but publishNotReadyAddresses true
func TestProcessHeadlessEndpointSlices_NotReadyPublished(t *testing.T) {
	sc := &serviceSource{}
	endpointSlice := &discoveryv1.EndpointSlice{
		ObjectMeta:  metav1.ObjectMeta{Name: "slice4", Namespace: "default"},
		AddressType: discoveryv1.AddressTypeIPv4,
		Endpoints: []discoveryv1.Endpoint{
			{
				TargetRef:  &v1.ObjectReference{Kind: "Pod", Name: "some-pod"},
				Conditions: discoveryv1.EndpointConditions{Ready: pointerToBool(false)},
				Addresses:  []string{"1.2.3.4"},
			},
		},
	}
	pods := []*v1.Pod{{ObjectMeta: metav1.ObjectMeta{Name: "some-pod"}}}
	params := HeadlessEndpointParams{
		sc:                       sc,
		endpointSlices:           []*discoveryv1.EndpointSlice{endpointSlice},
		pods:                     pods,
		hostname:                 "test.example.com",
		endpointsType:            "IPv4",
		publishPodIPs:            true,
		publishNotReadyAddresses: true,
	}
	result := processHeadlessEndpointSlices(params)
	assert.NotEmpty(t, result, "Targets should be added when endpoint is not ready but publishNotReadyAddresses is true")
}

func TestServiceTypes_isNodeInformerRequired(t *testing.T) {
	tests := []struct {
		name     string
		filter   []v1.ServiceType
		required []v1.ServiceType
		want     bool
	}{
		{
			name:     "NodePort required and filter is empty",
			filter:   []v1.ServiceType{},
			required: []v1.ServiceType{v1.ServiceTypeNodePort},
			want:     true,
		},
		{
			name:     "NodePort type present",
			filter:   []v1.ServiceType{v1.ServiceTypeNodePort},
			required: []v1.ServiceType{v1.ServiceTypeNodePort},
			want:     true,
		},
		{
			name:     "NodePort type absent, filter enabled",
			filter:   []v1.ServiceType{v1.ServiceTypeLoadBalancer},
			required: []v1.ServiceType{v1.ServiceTypeNodePort},
			want:     false,
		},
		{
			name:     "NodePort and other filters present",
			filter:   []v1.ServiceType{v1.ServiceTypeLoadBalancer, v1.ServiceTypeNodePort},
			required: []v1.ServiceType{v1.ServiceTypeNodePort},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert []v1.ServiceType to []string for filter
			filterStr := make([]string, len(tt.filter))
			for i, v := range tt.filter {
				filterStr[i] = string(v)
			}
			requiredTypes := make([]v1.ServiceType, len(tt.required))
			for i, v := range tt.required {
				requiredTypes[i] = v
			}
			filter, _ := newServiceTypesFilter(filterStr)
			got := filter.isRequired(requiredTypes...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestServiceSource_AddEventHandler(t *testing.T) {
	var fakeServiceInformer *informers.FakeServiceInformer
	var fakeEdpInformer *informers.FakeEndpointSliceInformer
	var fakeNodeInformer *informers.FakeNodeInformer
	tests := []struct {
		name    string
		filter  []string
		times   int
		asserts func(t *testing.T, s *serviceSource)
	}{
		{
			name:   "AddEventHandler should trigger all event handlers when empty filter is provided",
			filter: []string{},
			times:  3,
			asserts: func(t *testing.T, s *serviceSource) {
				fakeServiceInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeEdpInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeNodeInformer.AssertNumberOfCalls(t, "Informer", 1)
			},
		},
		{
			name:   "AddEventHandler should trigger only service event handler",
			filter: []string{string(v1.ServiceTypeExternalName), string(v1.ServiceTypeLoadBalancer)},
			times:  1,
			asserts: func(t *testing.T, s *serviceSource) {
				fakeServiceInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeEdpInformer.AssertNumberOfCalls(t, "Informer", 0)
				fakeNodeInformer.AssertNumberOfCalls(t, "Informer", 0)
			},
		},
		{
			name:   "AddEventHandler should configure only service event handler",
			filter: []string{string(v1.ServiceTypeExternalName), string(v1.ServiceTypeLoadBalancer), string(v1.ServiceTypeClusterIP)},
			times:  2,
			asserts: func(t *testing.T, s *serviceSource) {
				fakeServiceInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeEdpInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeNodeInformer.AssertNumberOfCalls(t, "Informer", 0)
			},
		},
		{
			name:   "AddEventHandler should configure all service event handlers",
			filter: []string{string(v1.ServiceTypeNodePort)},
			times:  3,
			asserts: func(t *testing.T, s *serviceSource) {
				fakeServiceInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeEdpInformer.AssertNumberOfCalls(t, "Informer", 1)
				fakeNodeInformer.AssertNumberOfCalls(t, "Informer", 1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeServiceInformer = new(informers.FakeServiceInformer)
			infSvc := testInformer{}
			fakeServiceInformer.On("Informer").Return(&infSvc)

			fakeEdpInformer = new(informers.FakeEndpointSliceInformer)
			infEdp := testInformer{}
			fakeEdpInformer.On("Informer").Return(&infEdp)

			fakeNodeInformer = new(informers.FakeNodeInformer)
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

			tt.asserts(t, svcSource)
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

		rawObjects := []interface{}{validSlice}
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

		rawObjects := []interface{}{invalidObject, validSlice}
		result := convertToEndpointSlices(rawObjects)

		assert.Len(t, result, 1)
		assert.Equal(t, "valid-slice", result[0].Name)
	})

	t.Run("handles empty input", func(t *testing.T) {
		result := convertToEndpointSlices([]interface{}{})
		assert.Empty(t, result)
	})

	t.Run("handles all invalid objects", func(t *testing.T) {
		rawObjects := []interface{}{"invalid1", 123, map[string]string{"key": "value"}}
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
				Conditions: discoveryv1.EndpointConditions{Ready: pointerToBool(true)},
			},
		},
	}
	pods := []*v1.Pod{} // No pods, so pod == nil
	hostname := "test.example.com"
	endpointsType := "IPv4"
	publishPodIPs := true
	publishNotReadyAddresses := false

	params := HeadlessEndpointParams{
		sc:                       sc,
		endpointSlices:           []*discoveryv1.EndpointSlice{endpointSlice},
		pods:                     pods,
		hostname:                 hostname,
		endpointsType:            endpointsType,
		publishPodIPs:            publishPodIPs,
		publishNotReadyAddresses: publishNotReadyAddresses,
	}
	result := processHeadlessEndpointSlices(params)
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
				Conditions: discoveryv1.EndpointConditions{Ready: pointerToBool(true)},
			},
		},
	}
	pods := []*v1.Pod{{ObjectMeta: metav1.ObjectMeta{Name: "some-pod"}}}
	hostname := "test.example.com"
	endpointsType := "FQDN"
	publishPodIPs := true
	publishNotReadyAddresses := false

	params := HeadlessEndpointParams{
		sc:                       sc,
		endpointSlices:           []*discoveryv1.EndpointSlice{endpointSlice},
		pods:                     pods,
		hostname:                 hostname,
		endpointsType:            endpointsType,
		publishPodIPs:            publishPodIPs,
		publishNotReadyAddresses: publishNotReadyAddresses,
	}
	result := processHeadlessEndpointSlices(params)
	assert.Empty(t, result, "No targets should be added for unsupported address type when publishPodIPs is true")
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

		result := buildHeadlessEndpoints(targetsByHeadlessDomainAndType, endpoint.TTL(0), svc)

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

		result := buildHeadlessEndpoints(targetsByHeadlessDomainAndType, endpoint.TTL(0), svc)

		assert.Len(t, result, 1)
		assert.Len(t, result[0].Targets, 2)
		assert.Contains(t, result[0].Targets, "1.2.3.4")
		assert.Contains(t, result[0].Targets, "5.6.7.8")
	})

	t.Run("handles TTL configuration", func(t *testing.T) {
		targetsByHeadlessDomainAndType := map[endpoint.EndpointKey]endpoint.Targets{
			{DNSName: "test.example.com", RecordType: endpoint.RecordTypeA}: {"1.2.3.4"},
		}

		result := buildHeadlessEndpoints(targetsByHeadlessDomainAndType, endpoint.TTL(300), svc)

		assert.Len(t, result, 1)
		assert.Equal(t, endpoint.TTL(300), result[0].RecordTTL)
	})

	t.Run("sorts endpoints deterministically", func(t *testing.T) {
		targetsByHeadlessDomainAndType := map[endpoint.EndpointKey]endpoint.Targets{
			{DNSName: "z.example.com", RecordType: endpoint.RecordTypeA}:    {"1.2.3.4"},
			{DNSName: "a.example.com", RecordType: endpoint.RecordTypeA}:    {"5.6.7.8"},
			{DNSName: "a.example.com", RecordType: endpoint.RecordTypeAAAA}: {"2001:db8::1"},
		}

		result := buildHeadlessEndpoints(targetsByHeadlessDomainAndType, endpoint.TTL(0), svc)

		assert.Len(t, result, 3)
		// Should be sorted by DNSName first, then by RecordType
		assert.Equal(t, "a.example.com", result[0].DNSName)
		assert.Equal(t, endpoint.RecordTypeA, result[0].RecordType)
		assert.Equal(t, "a.example.com", result[1].DNSName)
		assert.Equal(t, endpoint.RecordTypeAAAA, result[1].RecordType)
		assert.Equal(t, "z.example.com", result[2].DNSName)
	})

	t.Run("handles empty targets", func(t *testing.T) {
		result := buildHeadlessEndpoints(map[endpoint.EndpointKey]endpoint.Targets{}, endpoint.TTL(0), svc)
		assert.Empty(t, result)
	})
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
