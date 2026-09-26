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

package informers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestNormalizeNamespaces(t *testing.T) {
	tests := []struct {
		name       string
		namespaces []string
		want       []string
	}{
		{
			name:       "nil is cluster wide",
			namespaces: nil,
			want:       []string{corev1.NamespaceAll},
		},
		{
			name:       "empty is cluster wide",
			namespaces: []string{},
			want:       []string{corev1.NamespaceAll},
		},
		{
			name:       "single namespace is kept",
			namespaces: []string{"team-a"},
			want:       []string{"team-a"},
		},
		{
			name:       "order is preserved",
			namespaces: []string{"team-b", "team-a", "team-c"},
			want:       []string{"team-b", "team-a", "team-c"},
		},
		{
			name:       "duplicates are removed keeping first occurrence",
			namespaces: []string{"team-a", "team-b", "team-a"},
			want:       []string{"team-a", "team-b"},
		},
		{
			name:       "empty entry subsumes every other namespace",
			namespaces: []string{"team-a", "", "team-b"},
			want:       []string{corev1.NamespaceAll},
		},
		{
			name:       "empty entry first subsumes every other namespace",
			namespaces: []string{"", "team-a"},
			want:       []string{corev1.NamespaceAll},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NormalizeNamespaces(tt.namespaces))
		})
	}
}
