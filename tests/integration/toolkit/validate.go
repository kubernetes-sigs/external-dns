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

package toolkit

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
)

// ValidateScenarioEndpoints asserts that the actual endpoints match the expected
// scenario endpoints. Standard fields (DNSName, Targets, RecordType, TTL, etc.)
// are validated via the shared testutils helper. When an ExpectedEndpoint carries
// a non-empty RefObjects list, the actual endpoint is additionally checked to have
// exactly that many ref objects, with each one satisfying the partial match
// (only non-empty fields in ExpectedRefObject are compared).
func ValidateScenarioEndpoints(t *testing.T, got []*endpoint.Endpoint, expected []*ExpectedEndpoint) {
	t.Helper()

	plain := make([]*endpoint.Endpoint, len(expected))
	for i, e := range expected {
		plain[i] = e.ToEndpoint()
	}
	testutils.ValidateEndpoints(t, got, plain)

	// After ValidateEndpoints has sorted both slices in-place, pair them up and
	// check refObjects where the expected entry specifies them.
	// We re-sort expected by the same key so the pairing is stable.
	sortExpected(expected)

	for i := range min(len(got), len(expected)) {
		exp := expected[i]
		if len(exp.RefObjects) == 0 {
			continue
		}
		act := got[i]
		actualRefs := act.RefObjects()
		if len(actualRefs) != len(exp.RefObjects) {
			t.Errorf("endpoint %q: expected %d refObjects, got %d",
				exp.DNSName, len(exp.RefObjects), len(actualRefs))
			continue
		}
		errs := matchRefObjects(exp.DNSName, actualRefs, exp.RefObjects)
		for _, e := range errs {
			t.Error(e)
		}
	}
}

// matchRefObjects checks that every ExpectedRefObject.Key is satisfied by some
// actual ref's Key() (order-independent).
func matchRefObjects(dnsName string, actual []*endpoint.ObjectRef, expected []ExpectedRefObject) []string {
	used := make([]bool, len(actual))
	var errs []string

	for _, exp := range expected {
		matched := false
		for j, ref := range actual {
			if used[j] {
				continue
			}
			if ref.Key() == exp.Key {
				used[j] = true
				matched = true
				break
			}
		}
		if !matched {
			errs = append(errs, fmt.Sprintf(
				"endpoint %q: no actual refObject has key %q",
				dnsName, exp.Key,
			))
		}
	}
	return errs
}

// sortExpected sorts expected endpoints by the same key as testutils.ValidateEndpoints
// so that the pairing with got[] is stable.
func sortExpected(expected []*ExpectedEndpoint) {
	slices.SortFunc(expected, compareExpected)
}

func compareExpected(a, b *ExpectedEndpoint) int {
	if n := strings.Compare(a.DNSName, b.DNSName); n != 0 {
		return n
	}
	if n := strings.Compare(a.RecordType, b.RecordType); n != 0 {
		return n
	}
	return strings.Compare(a.Targets.String(), b.Targets.String())
}
