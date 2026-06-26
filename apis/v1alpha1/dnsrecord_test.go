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

package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestMergeProviderLabels(t *testing.T) {
	tests := []struct {
		name           string
		current        endpoint.Labels
		providerLabels endpoint.Labels
		ownerID        string
		want           endpoint.Labels
	}{
		{
			name:           "provider labels replace current, owner is stamped",
			current:        endpoint.Labels{endpoint.OwnerLabelKey: "old"},
			providerLabels: endpoint.Labels{"foo": "bar"},
			ownerID:        "me",
			want:           endpoint.Labels{"foo": "bar", endpoint.OwnerLabelKey: "me"},
		},
		{
			name:           "resource label is preserved across the merge",
			current:        endpoint.Labels{endpoint.ResourceLabelKey: "ingress/default/web", endpoint.OwnerLabelKey: "me"},
			providerLabels: endpoint.Labels{"foo": "bar"},
			ownerID:        "me",
			want:           endpoint.Labels{"foo": "bar", endpoint.OwnerLabelKey: "me", endpoint.ResourceLabelKey: "ingress/default/web"},
		},
		{
			name:           "nil provider labels still yields an owner-labeled map",
			current:        endpoint.Labels{endpoint.OwnerLabelKey: "me"},
			providerLabels: nil,
			ownerID:        "me",
			want:           endpoint.Labels{endpoint.OwnerLabelKey: "me"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DNSRecord{
				Spec: DNSRecordSpec{
					Endpoint: endpoint.Endpoint{Labels: tt.current},
				},
			}
			r.MergeProviderLabels(tt.providerLabels, tt.ownerID)
			assert.Equal(t, tt.want, r.Spec.Endpoint.Labels)
		})
	}
}
