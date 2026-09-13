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

package endpoint

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/yaml"
)

// The chart copies are included because they are what users install.
var crdManifests = []string{
	"../config/crd/standard/dnsendpoints.externaldns.k8s.io.yaml",
	"../config/crd/standard/dnsrecords.externaldns.k8s.io.yaml",
	"../charts/external-dns/crds/dnsendpoints.externaldns.k8s.io.yaml",
	"../charts/external-dns/crds/dnsrecords.externaldns.k8s.io.yaml",
}

// validate-crd.yml cannot catch this drift: it only checks that `make crd` is a
// no-op, and a stale hand-written marker regenerates cleanly.
func TestRecordTypeEnumMatchesKnownRecordTypes(t *testing.T) {
	for _, manifest := range crdManifests {
		t.Run(filepath.Base(filepath.Dir(manifest))+"/"+filepath.Base(manifest), func(t *testing.T) {
			raw, err := os.ReadFile(manifest)
			require.NoError(t, err)

			var crd map[string]any
			require.NoError(t, yaml.Unmarshal(raw, &crd))

			enums := findRecordTypeEnums(crd)
			require.NotEmpty(t, enums, "no recordType enum found; the schema layout changed and this test no longer checks anything")

			for _, enum := range enums {
				assert.ElementsMatch(t, KnownRecordTypes, enum,
					"recordType enum differs from KnownRecordTypes; update the Enum marker on Endpoint.RecordType and run `make crd`")
			}
		})
	}
}

// Walks rather than indexing a fixed path: DNSEndpoint nests a list of endpoints,
// DNSRecord holds a single one.
func findRecordTypeEnums(node any) [][]string {
	var found [][]string

	switch typed := node.(type) {
	case map[string]any:
		if property, ok := typed["recordType"].(map[string]any); ok {
			if enum, ok := property["enum"].([]any); ok {
				values := make([]string, 0, len(enum))
				for _, value := range enum {
					if str, ok := value.(string); ok {
						values = append(values, str)
					}
				}
				found = append(found, values)
			}
		}
		for _, child := range typed {
			found = append(found, findRecordTypeEnums(child)...)
		}
	case []any:
		for _, child := range typed {
			found = append(found, findRecordTypeEnums(child)...)
		}
	}

	return found
}
