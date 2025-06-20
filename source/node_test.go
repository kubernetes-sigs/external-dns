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
	"fmt"
	"maps"
	"math/rand"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"k8s.io/client-go/kubernetes"

	"sigs.k8s.io/external-dns/internal/testutils"

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
	t.Run("EndpointsIPv6", testNodeEndpointsWithIPv6)
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
			title:        "complex template",
			expectError:  false,
			fqdnTemplate: "{{range .Status.Addresses}}{{if and (eq .Type \"ExternalIP\") (isIPv4 .Address)}}{{.Address | replace \".\" \"-\"}}{{break}}{{end}}{{end}}.ext-dns.test.com",
		},
		{
			title:            "non-empty annotation filter label",
			expectError:      false,
			annotationFilter: "kubernetes.io/ingress.class=nginx",
		},
	} {

		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			_, err := NewNodeSource(
				context.TODO(),
				fake.NewClientset(),
				ti.annotationFilter,
				ti.fqdnTemplate,
				labels.Everything(),
				true,
				true,
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

// testNodeSourceEndpoints tests that various node generate the correct endpoints.
func testNodeSourceEndpoints(t *testing.T) {
	for _, tc := range []struct {
		title                string
		annotationFilter     string
		labelSelector        string
		fqdnTemplate         string
		nodeName             string
		nodeAddresses        []v1.NodeAddress
		labels               map[string]string
		annotations          map[string]string
		excludeUnschedulable bool // default to false
		exposeInternalIPv6   bool // default to true for this version. Change later when the next minor version is released.
		unschedulable        bool // default to false
		expected             []*endpoint.Endpoint
		expectError          bool
		expectedLogs         []string
		expectedAbsentLogs   []string
	}{
		{
			title:              "node with short hostname returns one endpoint",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "node with fqdn returns one endpoint",
			nodeName:           "node1.example.org",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "ipv6 node with fqdn returns one endpoint",
			nodeName:           "node1.example.org",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: "2001:DB8::8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "AAAA", DNSName: "node1.example.org", Targets: endpoint.Targets{"2001:DB8::8"}},
			},
		},
		{
			title:              "node with fqdn template returns endpoint with expanded hostname",
			fqdnTemplate:       "{{.Name}}.example.org",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "node with fqdn and fqdn template returns one endpoint",
			fqdnTemplate:       "{{.Name}}.example.org",
			nodeName:           "node1.example.org",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "node with fqdn template returns two endpoints with multiple IP addresses and expanded hostname",
			fqdnTemplate:       "{{.Name}}.example.org",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}, {Type: v1.NodeExternalIP, Address: "5.6.7.8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"1.2.3.4", "5.6.7.8"}},
			},
		},
		{
			title:              "node with fqdn template returns two endpoints with dual-stack IP addresses and expanded hostname",
			fqdnTemplate:       "{{.Name}}.example.org",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}, {Type: v1.NodeInternalIP, Address: "2001:DB8::8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"1.2.3.4"}},
				{RecordType: "AAAA", DNSName: "node1.example.org", Targets: endpoint.Targets{"2001:DB8::8"}},
			},
		},
		{
			title:              "node with both external and internal IP returns an endpoint with external IP",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}, {Type: v1.NodeInternalIP, Address: "2.3.4.5"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "node with both external, internal, and IPv6 IP returns endpoints with external IPs",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}, {Type: v1.NodeInternalIP, Address: "2.3.4.5"}, {Type: v1.NodeInternalIP, Address: "2001:DB8::8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
				{RecordType: "AAAA", DNSName: "node1", Targets: endpoint.Targets{"2001:DB8::8"}},
			},
		},
		{
			title:              "node with only internal IP returns an endpoint with internal IP",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: "2.3.4.5"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"2.3.4.5"}},
			},
		},
		{
			title:              "node with only internal IPs returns endpoints with internal IPs",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: "2.3.4.5"}, {Type: v1.NodeInternalIP, Address: "2001:DB8::8"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"2.3.4.5"}},
				{RecordType: "AAAA", DNSName: "node1", Targets: endpoint.Targets{"2001:DB8::8"}},
			},
		},
		{
			title:              "node with only internal IPs with expose internal IP as false shouldn't return AAAA endpoints with internal IPs",
			nodeName:           "node1",
			exposeInternalIPv6: false,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: "2.3.4.5"}, {Type: v1.NodeInternalIP, Address: "2001:DB8::9"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"2.3.4.5"}},
				{RecordType: "AAAA", DNSName: "node1", Targets: endpoint.Targets{"2001:DB8::9"}},
			},
		},
		{
			title:              "node with neither external nor internal IP returns no endpoints",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{},
			expectError:        true,
		},
		{
			title:              "node with target annotation",
			nodeName:           "node1.example.org",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				"external-dns.alpha.kubernetes.io/target": "203.2.45.7",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1.example.org", Targets: endpoint.Targets{"203.2.45.7"}},
			},
		},
		{
			title:              "annotated node without annotation filter returns endpoint",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "OnlyLocal",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "annotated node with matching annotation filter returns endpoint",
			annotationFilter:   "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "OnlyLocal",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "annotated node with non-matching annotation filter returns nothing",
			annotationFilter:   "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "SomethingElse",
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:              "labeled node with matching label selector returns endpoint",
			labelSelector:      "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			labels: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "OnlyLocal",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "labeled node with non-matching label selector returns nothing",
			labelSelector:      "service.beta.kubernetes.io/external-traffic in (Global, OnlyLocal)",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			labels: map[string]string{
				"service.beta.kubernetes.io/external-traffic": "SomethingElse",
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:              "our controller type is dns-controller",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				controllerAnnotationKey: controllerAnnotationValue,
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
		},
		{
			title:              "different controller types are ignored",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				controllerAnnotationKey: "not-dns-controller",
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:              "ttl not annotated should have RecordTTL.IsConfigured set to false",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:              "ttl annotated but invalid should have RecordTTL.IsConfigured set to false",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				ttlAnnotationKey: "foo",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(0)},
			},
		},
		{
			title:              "ttl annotated and is valid should set Record.TTL",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			annotations: map[string]string{
				ttlAnnotationKey: "10",
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}, RecordTTL: endpoint.TTL(10)},
			},
		},
		{
			title:                "unschedulable node return nothing with excludeUnschedulable=true",
			nodeName:             "node1",
			exposeInternalIPv6:   true,
			nodeAddresses:        []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			unschedulable:        true,
			excludeUnschedulable: true,
			expected:             []*endpoint.Endpoint{},
			expectedLogs: []string{
				"Skipping node node1 because it is unschedulable",
			},
		},
		{
			title:                "unschedulable node returns node with excludeUnschedulable=false",
			nodeName:             "node1",
			nodeAddresses:        []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}},
			unschedulable:        true,
			excludeUnschedulable: false,
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
			},
			expectedAbsentLogs: []string{
				"Skipping node node1 because it is unschedulable",
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			hook := testutils.LogsUnderTestWithLogLevel(log.DebugLevel, t)

			labelSelector := labels.Everything()
			if tc.labelSelector != "" {
				var err error
				labelSelector, err = labels.Parse(tc.labelSelector)
				require.NoError(t, err)
			}

			// Create a Kubernetes testing client
			kubeClient := fake.NewClientset()

			node := &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name:        tc.nodeName,
					Labels:      tc.labels,
					Annotations: tc.annotations,
				},
				Spec: v1.NodeSpec{
					Unschedulable: tc.unschedulable,
				},
				Status: v1.NodeStatus{
					Addresses: tc.nodeAddresses,
				},
			}

			_, err := kubeClient.CoreV1().Nodes().Create(context.Background(), node, metav1.CreateOptions{})
			require.NoError(t, err)

			// Create our object under test and get the endpoints.
			client, err := NewNodeSource(
				context.TODO(),
				kubeClient,
				tc.annotationFilter,
				tc.fqdnTemplate,
				labelSelector,
				tc.exposeInternalIPv6,
				tc.excludeUnschedulable,
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

			for _, entry := range tc.expectedLogs {
				testutils.TestHelperLogContains(entry, hook, t)
			}
			for _, entry := range tc.expectedAbsentLogs {
				testutils.TestHelperLogNotContains(entry, hook, t)
			}
		})
	}
}

func testNodeEndpointsWithIPv6(t *testing.T) {
	for _, tc := range []struct {
		title                string
		annotationFilter     string
		labelSelector        string
		fqdnTemplate         string
		nodeName             string
		nodeAddresses        []v1.NodeAddress
		labels               map[string]string
		annotations          map[string]string
		excludeUnschedulable bool // defaults to false
		exposeInternalIPv6   bool // default to true for this version. Change later when the next minor version is released.
		unschedulable        bool // default to false
		expected             []*endpoint.Endpoint
		expectError          bool
	}{
		{
			title:              "node with only internal IPs should return internal IPvs irrespective of exposeInternalIPv6",
			nodeName:           "node1",
			exposeInternalIPv6: false,
			nodeAddresses:      []v1.NodeAddress{{Type: v1.NodeInternalIP, Address: "2.3.4.5"}, {Type: v1.NodeInternalIP, Address: "2001:DB8::9"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"2.3.4.5"}},
				{RecordType: "AAAA", DNSName: "node1", Targets: endpoint.Targets{"2001:DB8::9"}},
			},
		},
		{
			title:              "node with both external, internal, and IPv6 IP returns endpoints with external IPs",
			nodeName:           "node1",
			exposeInternalIPv6: false,
			nodeAddresses: []v1.NodeAddress{{Type: v1.NodeExternalIP, Address: "1.2.3.4"}, {
				Type:    v1.NodeExternalIP,
				Address: "2001:DB8::8",
			}, {Type: v1.NodeInternalIP, Address: "2001:DB8::9"}},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.4"}},
				{RecordType: "AAAA", DNSName: "node1", Targets: endpoint.Targets{"2001:DB8::8"}},
			},
		},
		{
			title:              "node with both external and internal IPs should return internal IPv6 if exposeInternalIPv6 is true",
			nodeName:           "node1",
			exposeInternalIPv6: true,
			nodeAddresses: []v1.NodeAddress{
				{Type: v1.NodeExternalIP, Address: "1.2.3.5"},
				{Type: v1.NodeInternalIP, Address: "2001:DB8::9"},
			},
			expected: []*endpoint.Endpoint{
				{RecordType: "A", DNSName: "node1", Targets: endpoint.Targets{"1.2.3.5"}},
				{RecordType: "AAAA", DNSName: "node1", Targets: endpoint.Targets{"2001:DB8::9"}},
			},
		},
	} {
		labelSelector := labels.Everything()
		if tc.labelSelector != "" {
			var err error
			labelSelector, err = labels.Parse(tc.labelSelector)
			require.NoError(t, err)
		}

		// Create a Kubernetes testing client
		kubeClient := fake.NewClientset()

		node := &v1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name:        tc.nodeName,
				Labels:      tc.labels,
				Annotations: tc.annotations,
			},
			Spec: v1.NodeSpec{
				Unschedulable: tc.unschedulable,
			},
			Status: v1.NodeStatus{
				Addresses: tc.nodeAddresses,
			},
		}

		_, err := kubeClient.CoreV1().Nodes().Create(t.Context(), node, metav1.CreateOptions{})
		require.NoError(t, err)

		var hook *test.Hook
		if tc.exposeInternalIPv6 {
			hook = testutils.LogsUnderTestWithLogLevel(log.WarnLevel, t)
		}

		// Create our object under test and get the endpoints.
		client, err := NewNodeSource(
			t.Context(),
			kubeClient,
			tc.annotationFilter,
			tc.fqdnTemplate,
			labelSelector,
			tc.exposeInternalIPv6,
			tc.excludeUnschedulable,
			false,
		)
		require.NoError(t, err)

		endpoints, err := client.Endpoints(t.Context())
		if tc.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)

			if tc.exposeInternalIPv6 && hook != nil {
				testutils.TestHelperLogContainsWithLogLevel(warningMsg, log.WarnLevel, hook, t)
			}
		}

		// Validate returned endpoints against desired endpoints.
		validateEndpoints(t, endpoints, tc.expected)

		// TODO; when all resources have the resource label, we could add this check to the validateEndpoints function.
		for _, ep := range endpoints {
			require.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
		}
	}
}

func TestResourceLabelIsSetForEachNodeEndpoint(t *testing.T) {
	kubeClient := fake.NewClientset()

	nodes := helperNodeBuilder().
		withNode(nil).
		withNode(nil).
		withNode(nil).
		withNode(nil).
		build()

	for _, node := range nodes.Items {
		_, err := kubeClient.CoreV1().Nodes().Create(t.Context(), &node, metav1.CreateOptions{})
		require.NoError(t, err, "Failed to create node %s", node.Name)
	}

	client, err := NewNodeSource(
		t.Context(),
		kubeClient,
		"",
		"",
		labels.Everything(),
		false,
		true,
		false,
	)
	require.NoError(t, err)

	got, err := client.Endpoints(t.Context())
	require.NoError(t, err)
	for _, ep := range got {
		assert.NotEmpty(t, ep.Labels, "Labels should not be empty for endpoint %s", ep.DNSName)
		assert.Contains(t, ep.Labels, endpoint.ResourceLabelKey)
	}
}

type nodeListBuilder struct {
	nodes []v1.Node
}

func helperNodeBuilder() *nodeListBuilder {
	return &nodeListBuilder{nodes: []v1.Node{}}
}

func (b *nodeListBuilder) withNode(labels map[string]string) *nodeListBuilder {
	idx := len(b.nodes) + 1
	nodeName := fmt.Sprintf("ip-10-1-176-%d.internal", idx)
	b.nodes = append(b.nodes, v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: nodeName,
			Labels: func() map[string]string {
				base := map[string]string{
					"test-label":                    "test-value",
					"name":                          nodeName,
					"topology.kubernetes.io/region": "eu-west-1",
					"node.kubernetes.io/lifecycle":  "spot",
				}
				maps.Copy(base, labels)
				return base
			}(),
			Annotations: map[string]string{
				"volumes.kubernetes.io/controller-managed-attach-detach": "true",
				"alpha.kubernetes.io/provided-node-ip":                   fmt.Sprintf("10.1.176.%d", idx),
				"external-dns.alpha.kubernetes.io/hostname":              fmt.Sprintf("node-%d.example.com", idx),
			},
		},
		Spec: v1.NodeSpec{
			Unschedulable: false,
		},
		Status: v1.NodeStatus{
			Addresses: []v1.NodeAddress{
				{Type: v1.NodeInternalIP, Address: fmt.Sprintf("10.1.176.%d", idx)},
				{Type: v1.NodeInternalIP, Address: fmt.Sprintf("fc00:f853:ccd:e793::%d", idx)},
			},
		},
	})

	return b
}

func (b *nodeListBuilder) build() v1.NodeList {
	if len(b.nodes) > 1 {
		// Shuffle the result to ensure randomness in the order.
		rand.New(rand.NewSource(time.Now().UnixNano()))
		rand.Shuffle(len(b.nodes), func(i, j int) {
			b.nodes[i], b.nodes[j] = b.nodes[j], b.nodes[i]
		})
	}
	return v1.NodeList{Items: b.nodes}
}

func (b *nodeListBuilder) apply(t *testing.T, kubeClient kubernetes.Interface) v1.NodeList {
	for _, node := range b.nodes {
		_, err := kubeClient.CoreV1().Nodes().Create(t.Context(), &node, metav1.CreateOptions{})
		require.NoError(t, err, "Failed to create node %s", node.Name)
	}
	return v1.NodeList{Items: b.nodes}
}
