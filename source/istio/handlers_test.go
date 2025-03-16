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

package istio

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestGatewayAddFuncEventHandler(t *testing.T) {
	tests := []struct {
		name           string
		obj            interface{}
		expectedLogMsg string
		logLevel       log.Level
	}{
		{
			name: "valid service object",
			obj: &networkingv1alpha3.Gateway{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "istio-system",
					Name:      "istio-gw",
				},
			},
			expectedLogMsg: "event handler added for 'gateway.networking.istio.io/v1alpha3' in 'namespace:istio-system'",
			logLevel:       log.DebugLevel,
		},
		{
			name:           "invalid object type",
			obj:            "invalid object",
			expectedLogMsg: "event handler not added. want 'gateway.networking.istio.io/v1alpha3'",
			logLevel:       log.DebugLevel,
		},
		{
			name:           "on info level no logging",
			obj:            "invalid object",
			expectedLogMsg: "",
			logLevel:       log.InfoLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := testutils.LogsToBuffer(tt.logLevel, t)
			GatewayAddFuncEventHander(tt.obj)
			if tt.expectedLogMsg != "" {
				assert.Contains(t, string(buf.Bytes()), tt.expectedLogMsg)
			} else {
				assert.Empty(t, string(buf.Bytes()))
			}
		})
	}
}

func TestVirtualServiceAddFuncEventHandler(t *testing.T) {
	tests := []struct {
		name           string
		obj            interface{}
		expectedLogMsg string
		logLevel       log.Level
	}{
		{
			name: "valid service object",
			obj: &networkingv1alpha3.VirtualService{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "istio-system",
					Name:      "istio-gw",
				},
			},
			expectedLogMsg: "event handler added for 'virtualservice.networking.istio.io/v1alpha3' in 'namespace:istio-system'",
			logLevel:       log.DebugLevel,
		},
		{
			name:           "invalid object type",
			obj:            "invalid object",
			expectedLogMsg: "vent handler not added. want 'virtualservice.networking.istio.io/v1alpha3'",
			logLevel:       log.DebugLevel,
		},
		{
			name:           "on info level no logging",
			obj:            "invalid object",
			expectedLogMsg: "",
			logLevel:       log.InfoLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := testutils.LogsToBuffer(tt.logLevel, t)
			VirtualServiceAddedEvent(tt.obj)
			if tt.expectedLogMsg != "" {
				assert.Contains(t, string(buf.Bytes()), tt.expectedLogMsg)
			} else {
				assert.Empty(t, string(buf.Bytes()))
			}
		})
	}
}
