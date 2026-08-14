/*
Copyright 2026 The Kubernetes Authors.

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

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsKnown(t *testing.T) {
	for _, tt := range []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "known type, exact case",
			input:    "ingress",
			expected: true,
		},
		{
			name:     "known type, different case",
			input:    "Ingress",
			expected: true,
		},
		{
			name:     "known hyphenated type",
			input:    "gateway-httproute",
			expected: true,
		},
		{
			name:     "unknown type",
			input:    "ingres",
			expected: false,
		},
		{
			name:     "empty string",
			input:    "",
			expected: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsKnown(tt.input))
		})
	}
}

func TestAllContainsEveryConstant(t *testing.T) {
	expected := []Type{
		Node, Service, Ingress, Pod,
		GatewayHttpRoute, GatewayGrpcRoute, GatewayTlsRoute, GatewayTcpRoute, GatewayUdpRoute,
		IstioGateway, IstioVirtualService,
		AmbassadorHost, ContourHTTPProxy, GlooProxy, TraefikProxy, OpenShiftRoute,
		Fake, Connector, CRD, SkipperRouteGroup, KongTCPIngress,
		F5VirtualServer, F5TransportServer, Unstructured,
	}
	assert.ElementsMatch(t, expected, All)
}
