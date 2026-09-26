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
	v1 "k8s.io/api/core/v1"

	"sigs.k8s.io/external-dns/internal/sets"
)

// NormalizeNamespaces dedups the namespaces to watch, keeping their order.
// NamespaceAll subsumes the others, since watching all and a subset duplicates every object.
func NormalizeNamespaces(namespaces []string) []string {
	seen := sets.New[string]()
	result := make([]string, 0, len(namespaces))
	for _, ns := range namespaces {
		if ns == v1.NamespaceAll {
			return []string{v1.NamespaceAll}
		}
		if seen.Has(ns) {
			continue
		}
		seen.Insert(ns)
		result = append(result, ns)
	}
	if len(result) == 0 {
		return []string{v1.NamespaceAll}
	}
	return result
}

// SingleNamespace returns the only namespace to watch, NamespaceAll when there is none.
// Used by sources watching a single namespace, which ValidateConfig guarantees.
func SingleNamespace(namespaces []string) string {
	return NormalizeNamespaces(namespaces)[0]
}
