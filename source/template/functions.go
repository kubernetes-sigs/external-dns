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

package template

import (
	"encoding/json"
	"strings"
	"text/template"

	"sigs.k8s.io/external-dns/endpoint"
)

var (
	// baseTemplate is the shared base with all template functions pre-registered.
	// It is package-level so functions are registered once rather than on every parse call.
	// Each call to parseTemplate clones it before adding new content.
	baseTemplate = template.Must(
		template.New("endpoint").Funcs(template.FuncMap{
			"contains":   strings.Contains,
			"trimPrefix": strings.TrimPrefix,
			"trimSuffix": strings.TrimSuffix,
			"trim":       strings.TrimSpace,
			"toLower":    strings.ToLower,
			"replace":    replace,
			"isIPv6":     isIPv6,
			"isIPv4":     isIPv4,
			"hasKey":     hasKey,
			"fromJson":   fromJson,
		}).Parse(""),
	)
)

// replace all instances of oldValue with newValue in target string.
// adheres to syntax from https://masterminds.github.io/sprig/strings.html.
func replace(oldValue, newValue, target string) string {
	return strings.ReplaceAll(target, oldValue, newValue)
}

// isIPv6 reports whether the target string is an IPv6 address,
// including IPv4-mapped IPv6 addresses.
func isIPv6(target string) bool {
	return endpoint.SuitableType(target) == endpoint.RecordTypeAAAA
}

// isIPv4 reports whether the target string is an IPv4 address.
func isIPv4(target string) bool {
	return endpoint.SuitableType(target) == endpoint.RecordTypeA
}

// hasKey checks if a key exists in a map. This is needed because Go templates'
// `index` function returns the zero value ("") for missing keys, which is
// indistinguishable from keys with empty values. Kubernetes uses empty-value
// labels for markers (e.g., `service.kubernetes.io/headless: ""`), so we need
// explicit key existence checking.
func hasKey(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}

// fromJson decodes a JSON string into a Go value (map, slice, etc.).
// This enables templates to work with structured data stored as JSON strings
// in complex labels or annotations or Configmap data fields, e.g. ranging over a list of entries:
//
//	{{ range $entry := (index .Data "entries" | fromJson) }}{{ index $entry "dns" }},{{ end }}
//
// Returns nil if the input is not valid JSON.
func fromJson(v string) any {
	var output any
	_ = json.Unmarshal([]byte(v), &output)
	return output
}
