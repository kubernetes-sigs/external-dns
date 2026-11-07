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

package annotations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const addressesTmpl = "{{ range .status.addresses }}{{ .value }},{{ end }}"

func TestParseTargetFrom(t *testing.T) {
	for _, tt := range []struct {
		value   string
		want    TargetFromReference
		wantErr bool
	}{
		{
			value: "gateways.gateway.networking.k8s.io/istio-ingress/central#" + addressesTmpl,
			want: TargetFromReference{
				Resource:  schema.GroupResource{Group: "gateway.networking.k8s.io", Resource: "gateways"},
				Namespace: "istio-ingress", Name: "central", Template: addressesTmpl,
			},
		},
		{
			value: "services/shared-lb#{{ .spec.clusterIP }}",
			want:  TargetFromReference{Resource: schema.GroupResource{Resource: "services"}, Name: "shared-lb", Template: "{{ .spec.clusterIP }}"},
		},
		{
			// Only the first '#' splits; the template keeps any later '#' and '/'.
			value: `services/a#{{ "x#y/z" }}`,
			want:  TargetFromReference{Resource: schema.GroupResource{Resource: "services"}, Name: "a", Template: `{{ "x#y/z" }}`},
		},
		{value: "services/a", wantErr: true},
		{value: "services/a#", wantErr: true},
		{value: "services/a#  ", wantErr: true},
		{value: "a#{{ . }}", wantErr: true},
		{value: "/a#{{ . }}", wantErr: true},
		{value: "services/#{{ . }}", wantErr: true},
		{value: "services//a#{{ . }}", wantErr: true},
		{value: "services/ns/a/b#{{ . }}", wantErr: true},
	} {
		t.Run(tt.value, func(t *testing.T) {
			got, err := ParseTargetFrom(tt.value)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
