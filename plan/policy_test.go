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
	// more simple entries
	barV1 := []*endpoint.Endpoint{{DNSName: "bar1", Targets: endpoint.Targets{"v1"}}}
	barV2 := []*endpoint.Endpoint{{DNSName: "bar2", Targets: endpoint.Targets{"v2"}}}
	bazV1 := []*endpoint.Endpoint{{DNSName: "baz1", Targets: endpoint.Targets{"v1"}}}
	bazV2 := []*endpoint.Endpoint{{DNSName: "baz2", Targets: endpoint.Targets{"v2"}}}
	bazV3 := []*endpoint.Endpoint{{DNSName: "baz3", Targets: endpoint.Targets{"v3"}}}
	// Compose input
	barAndBaz := []*endpoint.Endpoint{barV1[0], barV2[0], bazV1[0], bazV2[0], bazV3[0]}
	fooAndBar := []*endpoint.Endpoint{barV1[0], barV2[0], fooV1[0], fooV2[0]}
	// Expected halves
	bar := []*endpoint.Endpoint{barV1[0], barV2[0]}
	baz := []*endpoint.Endpoint{bazV1[0], bazV2[0], bazV3[0]}
	foo := []*endpoint.Endpoint{fooV1[0], fooV2[0]}

	for _, tc := range []struct {
		policy   Policy
		changes  *Changes
		expected *Changes
	}{
		{
			// SyncPolicy doesn't modify the set of changes.
			&SyncPolicy{},
			&Changes{Create: bazV1, UpdateOld: fooV1, UpdateNew: fooV2, Delete: barV1},
			&Changes{Create: bazV1, UpdateOld: fooV1, UpdateNew: fooV2, Delete: barV1},
		},
		{
			// UpsertOnlyPolicy clears the list of deletions.
			&UpsertOnlyPolicy{},
			&Changes{Create: bazV1, UpdateOld: fooV1, UpdateNew: fooV2, Delete: barV1},
			&Changes{Create: bazV1, UpdateOld: fooV1, UpdateNew: fooV2, Delete: empty},
		},
		{
			// CreateOnlyPolicy clears the list of updates and deletions.
			&CreateOnlyPolicy{},
			&Changes{Create: bazV1, UpdateOld: fooV1, UpdateNew: fooV2, Delete: barV1},
			&Changes{Create: bazV1, UpdateOld: empty, UpdateNew: empty, Delete: empty},
		},
		{
			// FirstHalfChangesPolicy limits list to Create's first half
			&FirstHalfChangesPolicy{},
			&Changes{Create: barAndBaz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: barV1},
			&Changes{Create: bar, UpdateOld: empty, UpdateNew: empty, Delete: empty},
		},
		{
			// FirstHalfChangesPolicy limits list to Update's first half
			&FirstHalfChangesPolicy{},
			&Changes{Create: empty, UpdateOld: barAndBaz, UpdateNew: fooAndBar, Delete: barV1},
			&Changes{Create: empty, UpdateOld: bar, UpdateNew: bar, Delete: empty},
		},
		{
			// FirstHalfChangesPolicy limits list to Delete's first half
			&FirstHalfChangesPolicy{},
			&Changes{Create: empty, UpdateOld: empty, UpdateNew: empty, Delete: barAndBaz},
			&Changes{Create: empty, UpdateOld: empty, UpdateNew: empty, Delete: bar},
		},
		{
			// LastHalfChangesPolicy limits list to Create's last half
			&LastHalfChangesPolicy{},
			&Changes{Create: barAndBaz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: barV1},
			&Changes{Create: baz, UpdateOld: empty, UpdateNew: empty, Delete: empty},
		},
		{
			// LastHalfChangesPolicy limits list to Update's last half
			&LastHalfChangesPolicy{},
			&Changes{Create: empty, UpdateOld: barAndBaz, UpdateNew: fooAndBar, Delete: barV1},
			&Changes{Create: empty, UpdateOld: baz, UpdateNew: foo, Delete: empty},
		},
		{
			// LastHalfChangesPolicy limits list to Delete's last half
			&LastHalfChangesPolicy{},
			&Changes{Create: empty, UpdateOld: empty, UpdateNew: empty, Delete: barAndBaz},
			&Changes{Create: empty, UpdateOld: empty, UpdateNew: empty, Delete: baz},
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
