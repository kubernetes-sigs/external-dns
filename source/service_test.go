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
	"net"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/pkg/api/v1"

	"github.com/kubernetes-incubator/external-dns/endpoint"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceSource(t *testing.T) {
	t.Run("Interface", testServiceSourceImplementsSource)
	t.Run("NewServiceSource", testServiceSourceNewServiceSource)
	t.Run("Endpoints", testServiceSourceEndpoints)
}

// testServiceSourceImplementsSource tests that serviceSource is a valid Source.
func testServiceSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*Source)(nil), new(serviceSource))
}

// testServiceSourceNewServiceSource tests that NewServiceSource doesn't return an error.
func testServiceSourceNewServiceSource(t *testing.T) {
	for _, ti := range []struct {
		title        string
		fqdnTemplate string
		expectError  bool
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
	} {
		t.Run(ti.title, func(t *testing.T) {
			_, err := NewServiceSource(
				fake.NewSimpleClientset(),
				"",
				ti.fqdnTemplate,
				"",
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
	for _, tc := range []struct {
		title           string
		targetNamespace string
		svcNamespace    string
		svcName         string
		compatibility   string
		fqdnTemplate    string
		labels          map[string]string
		annotations     map[string]string
		lbs             []string
		expected        []*endpoint.Endpoint
		expectError     bool
	}{
		{
			"no annotated services return no endpoints",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{},
			false,
		},
		{
			"annotated services return an endpoint with target IP",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
			},
			false,
		},
		{
			"annotated services with multiple hostnames return an endpoint with target IP",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org., bar.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
				{DNSName: "bar.example.org", Target: "1.2.3.4"},
			},
			false,
		},
		{
			"annotated services with multiple hostnames and without trailing period return an endpoint with target IP",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org, bar.example.org",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
				{DNSName: "bar.example.org", Target: "1.2.3.4"},
			},
			false,
		},
		{
			"annotated services return an endpoint with target hostname",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
			},
			[]string{"lb.example.com"}, // Kubernetes omits the trailing dot
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "lb.example.com"},
			},
			false,
		},
		{
			"annotated services can omit trailing dot",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org", // Trailing dot is omitted
			},
			[]string{"1.2.3.4", "lb.example.com"}, // Kubernetes omits the trailing dot
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
				{DNSName: "foo.example.org", Target: "lb.example.com"},
			},
			false,
		},
		{
			"our controller type is dns-controller",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				controllerAnnotationKey: controllerAnnotationValue,
				hostnameAnnotationKey:   "foo.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
			},
			false,
		},
		{
			"different controller types are ignored even (with template specified)",
			"",
			"testing",
			"foo",
			"",
			"{{.Name}}.ext-dns.test.com",
			map[string]string{},
			map[string]string{
				controllerAnnotationKey: "some-other-tool",
				hostnameAnnotationKey:   "foo.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{},
			false,
		},
		{
			"services are found in target namespace",
			"testing",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
			},
			false,
		},
		{
			"services that are not in target namespace are ignored",
			"testing",
			"other-testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{},
			false,
		},
		{
			"services are found in all namespaces",
			"",
			"other-testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
			},
			false,
		},
		{
			"no external entrypoints return no endpoints",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
			},
			[]string{},
			[]*endpoint.Endpoint{},
			false,
		},
		{
			"multiple external entrypoints return multiple endpoints",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
			},
			[]string{"1.2.3.4", "8.8.8.8"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
				{DNSName: "foo.example.org", Target: "8.8.8.8"},
			},
			false,
		},
		{
			"services annotated with legacy mate annotations are ignored in default mode",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				"zalando.org/dnsname": "foo.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{},
			false,
		},
		{
			"services annotated with legacy mate annotations return an endpoint in compatibility mode",
			"",
			"testing",
			"foo",
			"mate",
			"",
			map[string]string{},
			map[string]string{
				"zalando.org/dnsname": "foo.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
			},
			false,
		},
		{
			"services annotated with legacy molecule annotations return an endpoint in compatibility mode",
			"",
			"testing",
			"foo",
			"molecule",
			"",
			map[string]string{
				"dns": "route53",
			},
			map[string]string{
				"domainName": "foo.example.org., bar.example.org",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
				{DNSName: "bar.example.org", Target: "1.2.3.4"},
			},
			false,
		},
		{
			"not annotated services with set fqdnTemplate return an endpoint with target IP",
			"",
			"testing",
			"foo",
			"",
			"{{.Name}}.bar.example.com",
			map[string]string{},
			map[string]string{},
			[]string{"1.2.3.4", "elb.com"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.bar.example.com", Target: "1.2.3.4"},
				{DNSName: "foo.bar.example.com", Target: "elb.com"},
			},
			false,
		},
		{
			"annotated services with set fqdnTemplate annotation takes precedence",
			"",
			"testing",
			"foo",
			"",
			"{{.Name}}.bar.example.com",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
			},
			[]string{"1.2.3.4", "elb.com"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4"},
				{DNSName: "foo.example.org", Target: "elb.com"},
			},
			false,
		},
		{
			"compatibility annotated services with tmpl. compatibility takes precedence",
			"",
			"testing",
			"foo",
			"mate",
			"{{.Name}}.bar.example.com",
			map[string]string{},
			map[string]string{
				"zalando.org/dnsname": "mate.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "mate.example.org", Target: "1.2.3.4"},
			},
			false,
		},
		{
			"not annotated services with unknown tmpl field should not return anything",
			"",
			"testing",
			"foo",
			"",
			"{{.Calibre}}.bar.example.com",
			map[string]string{},
			map[string]string{},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{},
			true,
		},
		{
			"ttl not annotated should have RecordTTL.IsConfigured set to false",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4", RecordTTL: endpoint.TTL{IsConfigured: false}},
			},
			false,
		},
		{
			"ttl annotated but invalid should have RecordTTL.IsConfigured set to false",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
				ttlAnnotationKey:      "foo",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4", RecordTTL: endpoint.TTL{IsConfigured: false}},
			},
			false,
		},
		{
			"ttl annotated and is valid should set Record.TTL",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
				ttlAnnotationKey:      "10",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4", RecordTTL: endpoint.TTL{IsConfigured: true, Value: 10}},
			},
			false,
		},
		{
			"Negative ttl is not valid",
			"",
			"testing",
			"foo",
			"",
			"",
			map[string]string{},
			map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
				ttlAnnotationKey:      "-10",
			},
			[]string{"1.2.3.4"},
			[]*endpoint.Endpoint{
				{DNSName: "foo.example.org", Target: "1.2.3.4", RecordTTL: endpoint.TTL{IsConfigured: false}},
			},
			false,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			// Create a Kubernetes testing client
			kubernetes := fake.NewSimpleClientset()

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

			_, err := kubernetes.CoreV1().Services(service.Namespace).Create(service)
			require.NoError(t, err)

			// Create our object under test and get the endpoints.
			client, _ := NewServiceSource(
				kubernetes,
				tc.targetNamespace,
				tc.fqdnTemplate,
				tc.compatibility,
			)
			require.NoError(t, err)

			endpoints, err := client.Endpoints()
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

func BenchmarkServiceEndpoints(b *testing.B) {
	kubernetes := fake.NewSimpleClientset()

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "testing",
			Name:      "foo",
			Annotations: map[string]string{
				hostnameAnnotationKey: "foo.example.org.",
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

	_, err := kubernetes.CoreV1().Services(service.Namespace).Create(service)
	require.NoError(b, err)

	client, err := NewServiceSource(kubernetes, v1.NamespaceAll, "", "")
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, err := client.Endpoints()
		require.NoError(b, err)
	}
}
