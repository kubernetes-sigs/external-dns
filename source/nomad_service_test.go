/*
Copyright 2025 The Kubernetes Authors.

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
	"encoding/json"
	"maps"
	"net/http"
	"net/http/httptest"
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/nomad/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"sigs.k8s.io/external-dns/endpoint"
)

type NomadServiceSuite struct {
	suite.Suite
	helloService *api.ServiceRegistration
	ts           *httptest.Server
	sc           Source
}

func (suite *NomadServiceSuite) SetupTest() {
	suite.helloService = &api.ServiceRegistration{
		ID:          "_nomad-task-71a63a80-a98a-93ee-4fd7-73b808577c20-group-foo-foo-http",
		ServiceName: "foo",
		Namespace:   "default",
		NodeID:      "6d7f412e-e7ff-2e66-d47b-867b0e9d8726",
		Datacenter:  "dc1",
		JobID:       "echo",
		AllocID:     "71a63a80-a98a-93ee-4fd7-73b808577c20",
		Tags:        []string{},
		Address:     "127.0.0.1",
		Port:        20627,
		CreateIndex: 18,
		ModifyIndex: 18,
	}

	fakeClient, ts, err := fakeNomadClient([]*api.ServiceRegistration{suite.helloService})
	suite.ts = ts

	suite.sc, err = NewNomadServiceSource(
		context.TODO(),
		fakeClient,
		"",
		"{{.Name}}",
		false,
		false,
	)
	suite.NoError(err, "should initialize nomad-service source")
}

func (suite *NomadServiceSuite) TearDownTest() {
	suite.ts.Close()
}

func (suite *NomadServiceSuite) TestResourceLabelIsSet() {
	endpoints, _ := suite.sc.Endpoints(context.Background())
	for _, ep := range endpoints {
		suite.Equal("service/default/foo", ep.Labels[endpoint.ResourceLabelKey], "should set correct resource label")
	}
}

func TestNomadServiceSource(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(NomadServiceSuite))
	t.Run("Interface", testNomadServiceSourceImplementsSource)
	t.Run("NewServiceSource", testNomadServiceSourceNewServiceSource)
	t.Run("Endpoints", testNomadServiceSourceEndpoints)
	t.Run("MultipleServices", testMultipleNomadServicesEndpoints)
}

// testNomadServiceSourceImplementsSource tests that serviceSource is a valid Source.
func testNomadServiceSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*Source)(nil), new(nomadServiceSource))
}

// testNomadServiceSourceNewServiceSource tests that NewServiceSource doesn't return an error.
func testNomadServiceSourceNewServiceSource(t *testing.T) {
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
		// {
		// 	title:            "non-empty annotation filter label",
		// 	expectError:      false,
		// 	annotationFilter: "kubernetes.io/ingress.class=nginx",
		// },
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			fakeClient, ts, err := fakeNomadClient([]*api.ServiceRegistration{})
			t.Cleanup(ts.Close)
			require.NoError(t, err)

			_, err = NewNomadServiceSource(
				context.TODO(),
				fakeClient,
				"",
				ti.fqdnTemplate,
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

// testNomadServiceSourceEndpoints tests that various services generate the correct endpoints.
func testNomadServiceSourceEndpoints(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		svcNamespace             string
		svcName                  string
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		ignoreHostnameAnnotation bool
		tags                     []string
		address                  string
		expected                 []*endpoint.Endpoint
		expectError              bool
	}{
		{
			title:        "no annotated services return no endpoints",
			svcNamespace: "testing",
			svcName:      "foo",
			tags:         []string{},
			address:      "1.2.3.4",
			expected:     []*endpoint.Endpoint{},
		},
		{
			title:                    "no annotated services return no endpoints when ignoring annotations",
			svcNamespace:             "testing",
			svcName:                  "foo",
			ignoreHostnameAnnotation: true,
			tags:                     []string{},
			address:                  "1.2.3.4",
			expected:                 []*endpoint.Endpoint{},
		},
		{
			title:        "annotated services return an endpoint with target IP",
			svcNamespace: "testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:                    "hostname annotation on services is ignored",
			svcNamespace:             "testing",
			svcName:                  "foo",
			ignoreHostnameAnnotation: true,
			tags: []string{
				"external-dns.hostname=foo.example.org.",
			},
			address:  "1.2.3.4",
			expected: []*endpoint.Endpoint{},
		},
		{
			title:        "FQDN template with multiple hostnames return an endpoint with target IP",
			svcNamespace: "testing",
			svcName:      "foo",
			fqdnTemplate: "{{.Name}}.fqdn.org,{{.Name}}.fqdn.com",
			address:      "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.fqdn.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.fqdn.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:                    "FQDN template with multiple hostnames return an endpoint with target IP when ignoring annotations",
			svcNamespace:             "testing",
			svcName:                  "foo",
			fqdnTemplate:             "{{.Name}}.fqdn.org,{{.Name}}.fqdn.com",
			ignoreHostnameAnnotation: true,
			address:                  "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.fqdn.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.fqdn.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:                    "FQDN template and annotation both with multiple hostnames return an endpoint with target IP",
			svcNamespace:             "testing",
			svcName:                  "foo",
			fqdnTemplate:             "{{.Name}}.fqdn.org,{{.Name}}.fqdn.com",
			combineFQDNAndAnnotation: true,
			tags: []string{
				"external-dns.hostname=foo.example.org., bar.example.org.",
			},
			address: "1.2.3.4",
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
			fqdnTemplate:             "{{.Name}}.fqdn.org,{{.Name}}.fqdn.com",
			combineFQDNAndAnnotation: true,
			ignoreHostnameAnnotation: true,
			tags: []string{
				"external-dns.hostname=foo.example.org., bar.example.org.",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.fqdn.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "foo.fqdn.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "annotated services with multiple hostnames return an endpoint with target IP",
			svcNamespace: "testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org., bar.example.org.",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "annotated services with multiple hostnames and without trailing period return an endpoint with target IP",
			svcNamespace: "testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org, bar.example.org",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "our controller type is kops dns controller",
			svcNamespace: "testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.controller=" + controllerAnnotationValue,
				"external-dns.hostname=foo.example.org.",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "different controller types are ignored even (with template specified)",
			svcNamespace: "testing",
			svcName:      "foo",
			fqdnTemplate: "{{.Name}}.ext-dns.test.com",
			tags: []string{
				"external-dns.controller=some-other-tool",
				"external-dns.hostname=foo.example.org.",
			},
			address:  "1.2.3.4",
			expected: []*endpoint.Endpoint{},
		},
		{
			title:           "services are found in target namespace",
			targetNamespace: "testing",
			svcNamespace:    "testing",
			svcName:         "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:           "services that are not in target namespace are ignored",
			targetNamespace: "testing",
			svcNamespace:    "other-testing",
			svcName:         "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
			},
			address:  "1.2.3.4",
			expected: []*endpoint.Endpoint{},
		},
		{
			title:        "services are found in all namespaces",
			svcNamespace: "other-testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "not annotated services with set fqdnTemplate return an endpoint with target IP",
			svcNamespace: "testing",
			svcName:      "foo",
			fqdnTemplate: "{{.Name}}.bar.example.com",
			tags:         []string{},
			address:      "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.bar.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "annotated services with set fqdnTemplate annotation takes precedence",
			svcNamespace: "testing",
			svcName:      "foo",
			fqdnTemplate: "{{.Name}}.bar.example.com",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:        "not annotated services with unknown tmpl field should not return anything",
			svcNamespace: "testing",
			svcName:      "foo",
			fqdnTemplate: "{{.Calibre}}.bar.example.com",
			tags:         []string{},
			address:      "1.2.3.4",
			expected:     []*endpoint.Endpoint{},
			expectError:  true,
		},
		{
			title:        "ttl not annotated should have RecordTTL.IsConfigured set to false",
			svcNamespace: "testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:        "ttl annotated but invalid should have RecordTTL.IsConfigured set to false",
			svcNamespace: "testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
				"external-dns.ttl=foo",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:        "ttl annotated and is valid should set Record.TTL",
			svcNamespace: "testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
				"external-dns.ttl=10",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(10)},
			},
		},
		{
			title:        "ttl annotated (in duration format) and is valid should set Record.TTL",
			svcNamespace: "testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
				"external-dns.ttl=1m",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(60)},
			},
		},
		{
			title:        "Negative ttl is not valid",
			svcNamespace: "testing",
			svcName:      "foo",
			tags: []string{
				"external-dns.hostname=foo.example.org.",
				"external-dns.ttl=-10",
			},
			address: "1.2.3.4",
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:        "IPv6 service gets IPv6 endpoint",
			svcNamespace: "testing",
			svcName:      "foobar-v6",
			tags: []string{
				"external-dns.hostname=foobar-v6.example.org",
			},
			address: "2001:db8::2",
			expected: []*endpoint.Endpoint{
				{DNSName: "foobar-v6.example.org", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::2"}},
			},
		},
	} {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			var service = &api.ServiceRegistration{
				ID:          "_nomad-task-71a63a80-a98a-93ee-4fd7-73b808577c20-group-foo-foo-http",
				ServiceName: tc.svcName,
				Namespace:   tc.svcNamespace,
				NodeID:      "6d7f412e-e7ff-2e66-d47b-867b0e9d8726",
				Datacenter:  "dc1",
				JobID:       "echo",
				AllocID:     "71a63a80-a98a-93ee-4fd7-73b808577c20",
				Tags:        tc.tags,
				Address:     tc.address,
				Port:        20627,
				CreateIndex: 18,
				ModifyIndex: 18,
			}

			fakeClient, ts, err := fakeNomadClient([]*api.ServiceRegistration{service})
			t.Cleanup(ts.Close)
			require.NoError(t, err)

			// Create our object under test and get the endpoints.
			client, err := NewNomadServiceSource(
				context.TODO(),
				fakeClient,
				tc.targetNamespace,
				tc.fqdnTemplate,
				tc.combineFQDNAndAnnotation,
				tc.ignoreHostnameAnnotation,
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

// testMultipleNomadServicesEndpoints tests that multiple services generate correct merged endpoints
func testMultipleNomadServicesEndpoints(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title                    string
		targetNamespace          string
		svcNamespace             string
		svcName                  string
		fqdnTemplate             string
		combineFQDNAndAnnotation bool
		ignoreHostnameAnnotation bool
		tags                     []string
		services                 map[string][]string
		expected                 []*endpoint.Endpoint
		expectError              bool
	}{
		{
			title:                    "test service returns a correct end point",
			targetNamespace:          "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			services: map[string][]string{
				"1.2.3.4": {"external-dns.hostname=foo.example.org"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
			},
			expectError: false,
		},
		{
			title:                    "multiple services that share same DNS should be merged into one endpoint",
			targetNamespace:          "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			services: map[string][]string{
				"1.2.3.4": {"external-dns.hostname=foo.example.org"},
				"1.2.3.5": {"external-dns.hostname=foo.example.org"},
				"1.2.3.6": {"external-dns.hostname=foo.example.org"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5", "1.2.3.6"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
			},
			expectError: false,
		},
		{
			title:                    "test that services with different hostnames do not get merged together",
			targetNamespace:          "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			services: map[string][]string{
				"1.2.3.5":  {"external-dns.hostname=foo.example.org"},
				"10.1.1.3": {"external-dns.hostname=bar.example.org"},
				"10.1.1.1": {"external-dns.hostname=bar.example.org"},
				"1.2.3.4":  {"external-dns.hostname=foo.example.org"},
				"10.1.1.2": {"external-dns.hostname=bar.example.org"},
				"20.1.1.1": {"external-dns.hostname=foobar.example.org"},
				"1.2.3.6":  {"external-dns.hostname=foo.example.org"},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.4", "1.2.3.5", "1.2.3.6"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
				{DNSName: "bar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.1", "10.1.1.2", "10.1.1.3"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo10.1.1.1"}},
				{DNSName: "foobar.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"20.1.1.1"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo20.1.1.1"}},
			},
			expectError: false,
		},
		{
			title:                    "test that services with different set-identifier do not get merged together",
			targetNamespace:          "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			services: map[string][]string{
				"1.2.3.5": {
					"external-dns.hostname=foo.example.org",
					"external-dns.set-identifier=a",
				},
				"10.1.1.3": {
					"external-dns.hostname=foo.example.org",
					"external-dns.set-identifier=b",
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"1.2.3.5"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.5"}, SetIdentifier: "a"},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"10.1.1.3"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo10.1.1.3"}, SetIdentifier: "b"},
			},
			expectError: false,
		},
		{
			title:                    "test that services with CNAME types do not get merged together",
			targetNamespace:          "",
			svcNamespace:             "testing",
			svcName:                  "foo",
			fqdnTemplate:             "",
			combineFQDNAndAnnotation: false,
			ignoreHostnameAnnotation: false,
			services: map[string][]string{
				"1.2.3.4": []string{
					"external-dns.hostname=foo.example.org",
					"external-dns.target=a.elb.com",
				},
				"1.2.3.5": []string{
					"external-dns.hostname=foo.example.org",
					"external-dns.target=b.elb.com",
				},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"a.elb.com"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.4"}},
				{DNSName: "foo.example.org", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"b.elb.com"}, Labels: map[string]string{endpoint.ResourceLabelKey: "service/testing/foo1.2.3.5"}},
			},
			expectError: false,
		},
	} {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			serviceRegistrations := make([]*api.ServiceRegistration, 0, len(tc.services))

			// Create services to test against
			for address, tags := range tc.services {
				serviceRegistrations = append(serviceRegistrations, &api.ServiceRegistration{
					ID:          "_nomad-task-71a63a80-a98a-93ee-4fd7-73b808577c20-group-foo-foo-http",
					ServiceName: tc.svcName + address,
					Namespace:   tc.svcNamespace,
					NodeID:      "6d7f412e-e7ff-2e66-d47b-867b0e9d8726",
					Datacenter:  "dc1",
					JobID:       "echo",
					AllocID:     "71a63a80-a98a-93ee-4fd7-73b808577c20",
					Tags:        tags,
					Address:     address,
					Port:        20627,
					CreateIndex: 18,
					ModifyIndex: 18,
				})
			}

			fakeClient, ts, err := fakeNomadClient(serviceRegistrations)
			t.Cleanup(ts.Close)
			require.NoError(t, err)

			// Create our object under test and get the endpoints.
			client, err := NewNomadServiceSource(
				context.TODO(),
				fakeClient,
				tc.targetNamespace,
				tc.fqdnTemplate,
				tc.combineFQDNAndAnnotation,
				tc.ignoreHostnameAnnotation,
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

func BenchmarkNomadServiceEndpoints(b *testing.B) {
	serviceRegistration := &api.ServiceRegistration{
		ID:          "_nomad-task-71a63a80-a98a-93ee-4fd7-73b808577c20-group-foo-foo-http",
		ServiceName: "foo",
		Namespace:   "testing",
		NodeID:      "6d7f412e-e7ff-2e66-d47b-867b0e9d8726",
		Datacenter:  "dc1",
		JobID:       "echo",
		AllocID:     "71a63a80-a98a-93ee-4fd7-73b808577c20",
		Tags:        []string{"external-dns.hostname=foo.example.org."},
		Address:     "1.2.3.4",
		Port:        20627,
		CreateIndex: 18,
		ModifyIndex: 18,
	}

	fakeClient, ts, err := fakeNomadClient([]*api.ServiceRegistration{serviceRegistration})
	b.Cleanup(ts.Close)
	require.NoError(b, err)

	client, err := NewNomadServiceSource(
		context.TODO(),
		fakeClient,
		"",
		"",
		false,
		false,
	)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, err := client.Endpoints(context.Background())
		require.NoError(b, err)
	}
}

func fakeNomadClient(services []*api.ServiceRegistration) (*api.Client, *httptest.Server, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/services", func(w http.ResponseWriter, r *http.Request) {
		namespaceFilter := r.FormValue("namespace")
		if namespaceFilter == "" {
			namespaceFilter = "default"
		}

		// Service stubs grouped by namespace and service name
		uniqueServices := make(map[string]map[string]*api.ServiceRegistrationStub)
		for _, service := range services {
			if _, ok := uniqueServices[service.Namespace]; !ok {
				uniqueServices[service.Namespace] = make(map[string]*api.ServiceRegistrationStub)
			}

			if _, ok := uniqueServices[service.Namespace][service.ServiceName]; !ok {
				uniqueServices[service.Namespace][service.ServiceName] = &api.ServiceRegistrationStub{
					ServiceName: service.ServiceName,
					Tags:        service.Tags,
				}
			}
		}

		resp := make([]*api.ServiceRegistrationListStub, 0, len(uniqueServices))
		for svcNamespace, services := range uniqueServices {
			if svcNamespace == namespaceFilter || namespaceFilter == "*" {
				resp = append(resp, &api.ServiceRegistrationListStub{
					Namespace: svcNamespace,
					Services:  slices.Collect(maps.Values(services)),
				})
			}
		}
		respBytes, _ := json.Marshal(resp)
		_, _ = w.Write(respBytes)
	})

	mux.HandleFunc("/v1/service/{serviceName}", func(w http.ResponseWriter, r *http.Request) {
		namespaceFilter := r.FormValue("namespace")
		serviceName := r.PathValue("serviceName")

		resp := make([]*api.ServiceRegistration, 0)
		for _, service := range services {
			if service.Namespace == namespaceFilter && service.ServiceName == serviceName {
				resp = append(resp, service)
			}
		}
		respBytes, _ := json.Marshal(resp)
		_, _ = w.Write(respBytes)
	})

	ts := httptest.NewServer(mux)

	fakeClient, err := NewNomadClient(ts.URL, "", "", 0)

	if err != nil {
		return nil, ts, err
	}

	return fakeClient, ts, nil
}
