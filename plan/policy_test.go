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
		policy                   Policy
		changes                  *Changes
		expected                 *Changes
		expectedSuppressed       []*endpoint.Endpoint
		expectedSuppressedUpdate []*endpoint.Endpoint
	}{
		{
			// SyncPolicy doesn't modify the set of changes.
			// SuppressedDelete is nil: sync allows deletions, nothing is suppressed.
			&SyncPolicy{},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: bar},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: bar},
			nil,
			nil,
		},
		{
			// UpsertOnlyPolicy clears the list of deletions.
			&UpsertOnlyPolicy{},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: bar},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: empty},
			bar,
			nil,
		},
		{
			// CreateOnlyPolicy clears the list of updates and deletions.
			// SuppressedUpdateOld carries UpdateOld for downstream logging.
			&CreateOnlyPolicy{},
			&Changes{Create: baz, UpdateOld: fooV1, UpdateNew: fooV2, Delete: bar},
			&Changes{Create: baz, UpdateOld: empty, UpdateNew: empty, Delete: empty},
			bar,
			fooV1,
		},
	} {
		// apply policy
		changes := tc.policy.Apply(tc.changes)

		// validate changes after applying policy
		validateEntries(t, changes.Create, tc.expected.Create)
		validateEntries(t, changes.UpdateOld, tc.expected.UpdateOld)
		validateEntries(t, changes.UpdateNew, tc.expected.UpdateNew)
		validateEntries(t, changes.Delete, tc.expected.Delete)
		validateEntries(t, changes.SuppressedDelete, tc.expectedSuppressed)
		validateEntries(t, changes.SuppressedUpdateOld, tc.expectedSuppressedUpdate)
	}
}

// TestPolicies tests that policies are correctly registered.
func TestPolicies(t *testing.T) {
	validatePolicy(t, Policies["sync"], &SyncPolicy{})
	validatePolicy(t, Policies["upsert-only"], &UpsertOnlyPolicy{})
	validatePolicy(t, Policies["create-only"], &CreateOnlyPolicy{})
}

// TestPolicyNameMatchesRegistryKey guards against drift between the
// registry key and the policy's own Name()/policyName() reporting: an
// alert built on a metric label or log field whose value doesn't match
// the configured `--policy=` CLI flag would be a silent observability bug.
func TestPolicyNameMatchesRegistryKey(t *testing.T) {
	for key, pol := range Policies {
		if got := policyName(pol); got != key {
			t.Errorf("policyName(%T) = %q, registered under key %q", pol, got, key)
		}
	}
}

// TestPolicyNameUnknownForThirdParty guards the optional-method fallback:
// a Policy that does not implement Name() must not break observability,
// and the rendered name must carry the concrete Go type so operators can
// identify the implementation from a debug log alone.
func TestPolicyNameUnknownForThirdParty(t *testing.T) {
	got := policyName(thirdPartyPolicy{})
	want := "unknown(plan.thirdPartyPolicy)"
	if got != want {
		t.Errorf("policyName(thirdPartyPolicy) = %q, want %q", got, want)
	}
}

type thirdPartyPolicy struct{}

func (thirdPartyPolicy) Apply(c *Changes) *Changes { return c }

// TestConcatSuppressed pins the nil-preservation contract of
// concatSuppressed: all-empty inputs must return nil, not an empty
// non-nil slice. The wantNil branch below asserts that directly via
// `got != nil`, so a future refactor that dropped the explicit guard
// and returned []*Endpoint{} for the empty case would fail this test
// instead of silently shifting the contract.
func TestConcatSuppressed(t *testing.T) {
	x := &endpoint.Endpoint{DNSName: "x"}
	y := &endpoint.Endpoint{DNSName: "y"}

	for _, tc := range []struct {
		name     string
		prev     []*endpoint.Endpoint
		deletes  []*endpoint.Endpoint
		want     []*endpoint.Endpoint
		wantNil  bool
	}{
		{"both nil returns nil", nil, nil, nil, true},
		{"both empty returns nil", []*endpoint.Endpoint{}, []*endpoint.Endpoint{}, nil, true},
		{"nil prev + non-empty deletes", nil, []*endpoint.Endpoint{x}, []*endpoint.Endpoint{x}, false},
		{"non-empty prev + nil deletes", []*endpoint.Endpoint{x}, nil, []*endpoint.Endpoint{x}, false},
		{"non-empty prev + non-empty deletes", []*endpoint.Endpoint{x}, []*endpoint.Endpoint{y}, []*endpoint.Endpoint{x, y}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := concatSuppressed(tc.prev, tc.deletes)
			if tc.wantNil {
				if got != nil {
					t.Errorf("concatSuppressed(%v, %v) = %v (len=%d), want nil", tc.prev, tc.deletes, got, len(got))
				}
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("concatSuppressed(%v, %v) = %v, want %v", tc.prev, tc.deletes, got, tc.want)
			}
		})
	}
}

// validatePolicy validates that a given policy is of the given type.
func validatePolicy(t *testing.T, policy, expected Policy) {
	policyType := reflect.TypeOf(policy).String()
	expectedType := reflect.TypeOf(expected).String()

	if policyType != expectedType {
		t.Errorf("expected %q to match %q", policyType, expectedType)
	}
}
