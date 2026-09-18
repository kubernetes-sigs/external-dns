/*
Copyright 2017 The Kubernetes Authors.

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

package plan

import (
	"reflect"
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
)

// TestApply tests that applying a policy results in the correct set of changes.
func TestApply(t *testing.T) {
	// empty list of records
	empty := []*endpoint.Endpoint{}
	// a simple entry
	fooV1 := []*endpoint.Endpoint{{DNSName: "foo", Targets: endpoint.Targets{"v1"}}}
	// the same entry but with different target
	fooV2 := []*endpoint.Endpoint{{DNSName: "foo", Targets: endpoint.Targets{"v2"}}}
	// another two simple entries
	bar := []*endpoint.Endpoint{{DNSName: "bar", Targets: endpoint.Targets{"v1"}}}
	baz := []*endpoint.Endpoint{{DNSName: "baz", Targets: endpoint.Targets{"v1"}}}

	for _, tc := range []struct {
		policy   Policy
		changes  *Changes
		expected *Changes
	}{
		{
			// SyncPolicy doesn't modify the set of changes.
			&SyncPolicy{},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: bar},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: bar},
		},
		{
			// UpsertOnlyPolicy clears the list of deletions.
			&UpsertOnlyPolicy{},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: bar},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: empty},
		},
		{
			// CreateOnlyPolicy clears the list of updates and deletions.
			&CreateOnlyPolicy{},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: bar},
			&Changes{Create: baz, UpdateOld: empty, UpdateNew: empty, Delete: empty},
		},
	} {
		// apply policy
		changes := tc.policy.Apply(tc.changes)

		// validate changes after applying policy
		validateEntries(t, changes.Create, tc.expected.Create)
		validateEntries(t, changes.UpdateOld, tc.expected.UpdateOld)
		validateEntries(t, changes.UpdateNew, tc.expected.UpdateNew)
		validateEntries(t, changes.Delete, tc.expected.Delete)
	}
}

// TestPolicies tests that policies are correctly registered.
func TestPolicies(t *testing.T) {
	validatePolicy(t, Policies["sync"], &SyncPolicy{})
	validatePolicy(t, Policies["upsert-only"], &UpsertOnlyPolicy{})
	validatePolicy(t, Policies["create-only"], &CreateOnlyPolicy{})
}

// validatePolicy validates that a given policy is of the given type.
func validatePolicy(t *testing.T, policy, expected Policy) {
	policyType := reflect.TypeOf(policy).String()
	expectedType := reflect.TypeOf(expected).String()

	if policyType != expectedType {
		t.Errorf("expected %q to match %q", policyType, expectedType)
	}
}
