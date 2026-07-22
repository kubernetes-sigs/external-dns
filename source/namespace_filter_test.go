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

package source

import (
	"errors"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type mockNamespaceLister struct {
	namespaces map[string]*corev1.Namespace
}

func (m *mockNamespaceLister) List(selector labels.Selector) ([]*corev1.Namespace, error) {
	var list []*corev1.Namespace
	for _, ns := range m.namespaces {
		if selector.Matches(labels.Set(ns.Labels)) {
			list = append(list, ns)
		}
	}
	return list, nil
}

func (m *mockNamespaceLister) Get(name string) (*corev1.Namespace, error) {
	ns, ok := m.namespaces[name]
	if !ok {
		return nil, errors.New("not found")
	}
	return ns, nil
}

func TestNamespaceFilter(t *testing.T) {
	lister := &mockNamespaceLister{
		namespaces: map[string]*corev1.Namespace{
			"ns1": {
				ObjectMeta: metav1.ObjectMeta{
					Name:   "ns1",
					Labels: map[string]string{"tenant": "tenant1"},
				},
			},
			"ns2": {
				ObjectMeta: metav1.ObjectMeta{
					Name:   "ns2",
					Labels: map[string]string{"tenant": "tenant2"},
				},
			},
			"ns3": {
				ObjectMeta: metav1.ObjectMeta{
					Name:   "ns3",
					Labels: map[string]string{"tenant": "tenant1", "env": "prod"},
				},
			},
		},
	}

	tests := []struct {
		name                 string
		namespaces           []string
		namespaceLabelFilter string
		testNS               string
		expected             bool
	}{
		{
			name:       "empty filters matches everything",
			namespaces: nil,
			testNS:     "ns1",
			expected:   true,
		},
		{
			name:       "namespaces list filter matches match",
			namespaces: []string{"ns1", "ns3"},
			testNS:     "ns1",
			expected:   true,
		},
		{
			name:       "namespaces list filter mismatches mismatch",
			namespaces: []string{"ns1", "ns3"},
			testNS:     "ns2",
			expected:   false,
		},
		{
			name:                 "namespace label selector match",
			namespaceLabelFilter: "tenant=tenant1",
			testNS:               "ns1",
			expected:             true,
		},
		{
			name:                 "namespace label selector mismatch",
			namespaceLabelFilter: "tenant=tenant1",
			testNS:               "ns2",
			expected:             false,
		},
		{
			name:                 "namespace label selector compound match",
			namespaceLabelFilter: "tenant=tenant1,env=prod",
			testNS:               "ns3",
			expected:             true,
		},
		{
			name:                 "namespace label selector compound mismatch",
			namespaceLabelFilter: "tenant=tenant1,env=prod",
			testNS:               "ns1",
			expected:             false,
		},
		{
			name:                 "both filters match",
			namespaces:           []string{"ns1", "ns3"},
			namespaceLabelFilter: "tenant=tenant1",
			testNS:               "ns3",
			expected:             true,
		},
		{
			name:                 "both filters - namespace match label mismatch",
			namespaces:           []string{"ns1", "ns2"},
			namespaceLabelFilter: "tenant=tenant1",
			testNS:               "ns2",
			expected:             false,
		},
		{
			name:                 "wildcard namespace always matches",
			namespaces:           []string{"ns1"},
			namespaceLabelFilter: "tenant=tenant1",
			testNS:               "",
			expected:             true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var selector labels.Selector
			if tt.namespaceLabelFilter != "" {
				selector, _ = labels.Parse(tt.namespaceLabelFilter)
			}
			filter := NewNamespaceFilter(tt.namespaces, selector, lister)
			got := filter.Matches(tt.testNS)
			if got != tt.expected {
				t.Errorf("expected Matches(%q) = %v, got %v", tt.testNS, tt.expected, got)
			}
		})
	}
}
