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
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayfake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source/annotations"
)

func gatewayStatusHostname(hostnames ...string) v1.GatewayStatus {
	typ := v1.HostnameAddressType
	addrs := make([]v1.GatewayStatusAddress, len(hostnames))
	for i, h := range hostnames {
		addrs[i] = v1.GatewayStatusAddress{Type: &typ, Value: h}
	}
	return v1.GatewayStatus{Addresses: addrs}
}

func makeGateway(namespace, name string, annots map[string]string, status v1.GatewayStatus, labels map[string]string) *v1.Gateway {
	return &v1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Namespace:   namespace,
			Annotations: annots,
			Labels:      labels,
		},
		Status: status,
	}
}

func TestGatewayResourceSourceEndpoints(t *testing.T) {
	t.Parallel()

	hostnameAnnotation := func(hostnames ...string) map[string]string {
		return map[string]string{annotations.HostnameKey: strings.Join(hostnames, ",")}
	}

	tests := []struct {
		title     string
		config    *Config
		gateways  []*v1.Gateway
		endpoints []*endpoint.Endpoint
	}{
		{
			title:  "IP address in status produces A record",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("foo.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "hostname address in status produces CNAME record",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatusHostname("lb.example.com"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpointWithTTL("foo.example.com", endpoint.RecordTypeCNAME, 0, "lb.example.com"),
			},
		},
		{
			title:  "IPv6 address in status produces AAAA record",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatus("2001:db8::1"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpointWithTTL("foo.example.com", endpoint.RecordTypeAAAA, 0, "2001:db8::1"),
			},
		},
		{
			title:  "multiple hostnames in annotation produce multiple endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("a.example.com", "b.example.com"), gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
				newTestEndpoint("b.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "multiple gateways produce multiple endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", hostnameAnnotation("a.example.com"), gatewayStatus("1.2.3.4"), nil),
				makeGateway("default", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
				newTestEndpoint("b.example.com", "5.6.7.8"),
			},
		},
		{
			title:  "GatewayName filter skips non-matching gateways",
			config: &Config{GatewayName: "gw1"},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", hostnameAnnotation("a.example.com"), gatewayStatus("1.2.3.4"), nil),
				makeGateway("default", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "GatewayNamespace filter limits to configured namespace",
			config: &Config{GatewayNamespace: "ns1"},
			gateways: []*v1.Gateway{
				makeGateway("ns1", "gw1", hostnameAnnotation("a.example.com"), gatewayStatus("1.2.3.4"), nil),
				makeGateway("ns2", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "GatewayLabelFilter skips gateways without matching label",
			config: &Config{GatewayLabelFilter: "env=prod"},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", hostnameAnnotation("a.example.com"), gatewayStatus("1.2.3.4"), map[string]string{"env": "prod"}),
				makeGateway("default", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), map[string]string{"env": "staging"}),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "AnnotationFilter skips gateways without matching annotation",
			config: &Config{AnnotationFilter: "custom=yes"},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", map[string]string{
					annotations.HostnameKey: "a.example.com",
					"custom":                "yes",
				}, gatewayStatus("1.2.3.4"), nil),
				makeGateway("default", "gw2", hostnameAnnotation("b.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("a.example.com", "1.2.3.4"),
			},
		},
		{
			title:  "IgnoreHostnameAnnotation produces no endpoints",
			config: &Config{IgnoreHostnameAnnotation: true},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: nil,
		},
		{
			title:  "no hostname annotation produces no endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", nil, gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: nil,
		},
		{
			title:  "no status addresses produces no endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", hostnameAnnotation("foo.example.com"), v1.GatewayStatus{}, nil),
			},
			endpoints: nil,
		},
		{
			title:  "target override annotation replaces status addresses",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", map[string]string{
					annotations.HostnameKey: "foo.example.com",
					annotations.TargetKey:   "override.example.com",
				}, gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpointWithTTL("foo.example.com", endpoint.RecordTypeCNAME, 0, "override.example.com"),
			},
		},
		{
			title:  "TTL annotation is applied to endpoints",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", map[string]string{
					annotations.HostnameKey: "foo.example.com",
					annotations.TtlKey:      "300",
				}, gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpointWithTTL("foo.example.com", endpoint.RecordTypeA, 300, "1.2.3.4"),
			},
		},
		{
			title:  "controller annotation mismatch skips gateway",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw", map[string]string{
					annotations.HostnameKey:   "foo.example.com",
					annotations.ControllerKey: "other-controller",
				}, gatewayStatus("1.2.3.4"), nil),
			},
			endpoints: nil,
		},
		{
			title:  "same hostname from two gateways merges targets",
			config: &Config{},
			gateways: []*v1.Gateway{
				makeGateway("default", "gw1", hostnameAnnotation("shared.example.com"), gatewayStatus("1.2.3.4"), nil),
				makeGateway("default", "gw2", hostnameAnnotation("shared.example.com"), gatewayStatus("5.6.7.8"), nil),
			},
			endpoints: []*endpoint.Endpoint{
				newTestEndpoint("shared.example.com", "1.2.3.4", "5.6.7.8"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
			defer cancel()

			gwClient := gatewayfake.NewSimpleClientset()
			for _, gw := range tt.gateways {
				_, err := gwClient.GatewayV1().Gateways(gw.Namespace).Create(ctx, gw, metav1.CreateOptions{})
				require.NoError(t, err, "failed to create Gateway")
			}

			clients := new(testutils.MockClientGenerator)
			clients.On("GatewayClient").Return(gwClient, nil)

			src, err := NewGatewaySource(ctx, clients, tt.config)
			require.NoError(t, err, "failed to create Gateway source")

			endpoints, err := src.Endpoints(ctx)
			require.NoError(t, err, "failed to get endpoints")

			testutils.ValidateEndpoints(t, endpoints, tt.endpoints)
		})
	}
}
