/*
Copyright 2019 The Kubernetes Authors.

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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/fake"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestNodeSource(t *testing.T) {
	t.Parallel()

	t.Run("NewNodeSource", testNodeSourceNewNodeSource)
	t.Run("Endpoints", testNodeSourceEndpoints)
}

// testNodeSourceNewNodeSource tests that NewNodeService doesn't return an error.
func testNodeSourceNewNodeSource(t *testing.T) {
	t.Parallel()

	for _, ti := range []struct {
		title            string
		annotationFilter string
		fqdnTemplate     string
		expectError      bool
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
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			_, err := NewNodeSource(
				context.TODO(),
				fake.NewSimpleClientset(),
				ti.annotationFilter,
				ti.fqdnTemplate,
				labels.Everything(),
			)

			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// testNodeSourceEndpoints tests that various node generate the correct endpoints.
func testNodeSourceEndpoints(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		title            string
		annotationFilter string
		labelSelector    string
		fqdnTemplate     string
		nodeName         string
		nodeAddresses    []v1.NodeAddress
		labels           map[string]string
		annotations      map[string]string
		expected         []*endpoint.Endpoint
		expectError      bool
	}{
		{
			title:         "node with short hostname returns one endpoint",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:         "node with fqdn returns one endpoint",
			nodeName:      "node1.example.org",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:         "ipv6 node with fqdn returns one endpoint",
			nodeName:      "node1.example.org",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: "2001:DB8::8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "AAAA", DNSName: "node1.example.org", Targets: endpoint.Targets{"2001:DB8::8"}},
			},
		},
		{
			title:         "node with fqdn template returns endpoint with expanded hostname",
			fqdnTemplate:  "{{.Name}}.example.org",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:         "node with fqdn and fqdn template returns one endpoint",
			fqdnTemplate:  "{{.Name}}.example.org",
			nodeName:      "node1.example.org",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:         "node with fqdn template returns two endpoints with multiple IP addresses and expanded hostname",
			fqdnTemplate:  "{{.Name}}.example.org",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}, {Type: v1.NodeExternalIP, Address: "5.6.7.8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"1.2.3.4", "5.6.7.8"}},
			},
		},
		{
			title:         "node with fqdn template returns two endpoints with dual-stack IP addresses and expanded hostname",
			fqdnTemplate:  "{{.Name}}.example.org",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}, {Type: v1.NodeInternalIP, Address: "2001:DB8::8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
				{RecordType: "AAAA", DNSName: "node1.example.org", Targets: endpoint.Targets{"2001:DB8::8"}},
			},
		},
		{
			title:         "node with both external and internal IP returns an endpoint with external IP",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}, {Type: v1.NodeInternalIP, Address: "2.3.4.5"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:         "node with both external, internal, and IPv6 IP returns endpoints with external IPs",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}, {Type: v1.NodeInternalIP, Address: "2.3.4.5"}, {Type: v1.NodeInternalIP, Address: "2001:DB8::8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
				{RecordType: "AAAA", DNSName: "node1", Targets: endpoint.Targets{"2001:DB8::8"}},
			},
		},
		{
			title:         "node with only internal IP returns an endpoint with internal IP",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: "2.3.4.5"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"2.3.4.5"}},
			},
		},
		{
			title:         "node with only internal IPs returns endpoints with internal IPs",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: "2.3.4.5"}, {Type: v1.NodeInternalIP, Address: "2001:DB8::8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"2.3.4.5"}},
				{RecordType: "AAAA", DNSName: "node1", Targets: endpoint.Targets{"2001:DB8::8"}},
			},
		},
		{
			title:         "node with neither external nor internal IP returns no endpoints",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{},
			expectError:   true,
		},
		{
			title:         "node with target annotation",
			nodeName:      "node1.example.org",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/target": "203.2.45.7",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"203.2.45.7"}},
			},
		},
		{
			title:         "annotated node without annotation filter returns endpoint",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "OnlyLocal",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:            "annotated node with matching annotation filter returns endpoint",
			annotationFilter: "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			nodeName:         "node1",
			nodeAddresses:    []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "OnlyLocal",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:            "annotated node with non-matching annotation filter returns nothing",
			annotationFilter: "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			nodeName:         "node1",
			nodeAddresses:    []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "SomethingElse",
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:         "labeled node with matching label selector returns endpoint",
			labelSelector: "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			labels: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "OnlyLocal",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:         "labeled node with non-matching label selector returns nothing",
			labelSelector: "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			labels: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "SomethingElse",
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:         "our controller type is dns-controller",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				controllerAnnotationKey: controllerAnnotationValue,
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:         "different controller types are ignored",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				controllerAnnotationKey: "not-dns-controller",
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:         "ttl not annotated should have RecordTTL.IsConfigured set to false",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:         "ttl annotated but invalid should have RecordTTL.IsConfigured set to false",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				ttlAnnotationKey: "foo",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:         "ttl annotated and is valid should set Record.TTL",
			nodeName:      "node1",
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				ttlAnnotationKey: "10",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(10)},
			},
		},
	} {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			t.Parallel()

			labelSelector := labels.Everything()
			if tc.labelSelector != "" {
				var err error
				labelSelector, err = labels.Parse(tc.labelSelector)
				require.NoError(t, err)
			}

			// Create a Kubernetes testing client
			kubernetes := fake.NewSimpleClientset()

			node := &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name:        tc.nodeName,
					Labels:      tc.labels,
					Annotations: tc.annotations,
				},
				Status: v1.NodeStatus{
					Addresses: tc.nodeAddresses,
				},
			}

			_, err := kubernetes.CoreV1().Nodes().Create(context.Background(), node, metav1.CreateOptions{})
			require.NoError(t, err)

			// Create our object under test and get the endpoints.
			client, err := NewNodeSource(
				context.TODO(),
				kubernetes,
				tc.annotationFilter,
				tc.fqdnTemplate,
				labelSelector,
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
