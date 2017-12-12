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
	"fmt"
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
)

// TestCalculate tests that a plan can calculate actions to move a list of
// current records to a list of desired records.
func TestCalculate(t *testing.T) {
	// we need different TTLs to create differing Endpoints with the same name and target
	ttl := endpoint.TTL(300)
	ttl2 := endpoint.TTL(50)

	// empty list of records
	empty := []*endpoint.Endpoint{}
	// a simple entry
	fooV1 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v1", endpoint.RecordTypeCNAME)}
	// the same entry but with different target
	fooV2 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v2", endpoint.RecordTypeCNAME)}
	// the same entry as before but with varying TTLs
	fooV2ttl1 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v2", endpoint.RecordTypeCNAME, ttl)}
	fooV2ttl2 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v2", endpoint.RecordTypeCNAME, ttl2)}
	// another simple entry
	bar := []*endpoint.Endpoint{endpoint.NewEndpoint("bar", "v1", endpoint.RecordTypeCNAME)}

	// test case with labels
	unlabeledTTL2 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v2", endpoint.RecordTypeCNAME, ttl2)}
	labeledTTL1 := []*endpoint.Endpoint{newEndpointWithOwnerAndTTL("foo", "v2", "123", ttl)}
	labeledTTL2 := []*endpoint.Endpoint{newEndpointWithOwnerAndTTL("foo", "v2", "123", ttl2)}

	// test case with type inheritance
	untypedTTL2 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v2", "", ttl2)}
	typedTTL1 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v2", endpoint.RecordTypeA, ttl)}
	typedTTL2 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v2", endpoint.RecordTypeA, ttl2)}

	// explicit TTL test cases
	ttlV1 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v1", endpoint.RecordTypeCNAME, ttl)}
	ttlV2 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v1", endpoint.RecordTypeCNAME)}
	ttlV3 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v1", endpoint.RecordTypeCNAME, ttl)}
	ttlV4 := []*endpoint.Endpoint{endpoint.NewEndpointWithTTL("foo", "v1", endpoint.RecordTypeCNAME, ttl2)}

	for _, tc := range []struct {
		policies                             []Policy
		current, desired                     []*endpoint.Endpoint
		create, updateOld, updateNew, delete []*endpoint.Endpoint
	}{
		// Nothing exists and nothing desired doesn't change anything.
		{[]Policy{&SyncPolicy{}}, empty, empty, empty, empty, empty, empty},
		// More desired than current creates the desired.
		{[]Policy{&SyncPolicy{}}, empty, fooV1, fooV1, empty, empty, empty},
		// Desired equals current doesn't change anything.
		{[]Policy{&SyncPolicy{}}, fooV1, fooV1, empty, empty, empty, empty},
		// Nothing is desired deletes the current.
		{[]Policy{&SyncPolicy{}}, fooV1, empty, empty, empty, empty, fooV1},
		// Current and desired match but TTL is different triggers an update.
		{[]Policy{&SyncPolicy{}}, fooV2ttl1, fooV2ttl2, empty, fooV2ttl1, fooV2ttl2, empty},
		// Both exist but are different creates desired and deletes current.
		{[]Policy{&SyncPolicy{}}, fooV1, bar, bar, empty, empty, fooV1},
		// Same thing with current and desired only having different targets
		{[]Policy{&SyncPolicy{}}, fooV1, fooV2, fooV2, empty, empty, fooV1},
		// Nothing is desired but policy doesn't allow deletions.
		{[]Policy{&UpsertOnlyPolicy{}}, fooV1, empty, empty, empty, empty, empty},
		// Labels should be inherited
		{[]Policy{&SyncPolicy{}}, labeledTTL1, unlabeledTTL2, empty, labeledTTL1, labeledTTL2, empty},
		// RecordType should be inherited
		{[]Policy{&SyncPolicy{}}, typedTTL1, untypedTTL2, empty, typedTTL1, typedTTL2, empty},
		// If desired TTL is not configured, do not update
		{[]Policy{&SyncPolicy{}}, ttlV1, ttlV2, empty, empty, empty, empty},
		// If desired TTL is configured but is the same as current TTL, do not update
		{[]Policy{&SyncPolicy{}}, ttlV1, ttlV3, empty, empty, empty, empty},
		// If desired TTL is configured and is not the same as current TTL, need to update
		{[]Policy{&SyncPolicy{}}, ttlV1, ttlV4, empty, ttlV1, ttlV4, empty},
	} {
		// setup plan
		plan := &Plan{
			Policies: tc.policies,
			Current:  tc.current,
			Desired:  tc.desired,
		}
		// calculate actions
		plan = plan.Calculate()

		// validate actions
		validateEntries(t, plan.Changes.Create, tc.create)
		validateEntries(t, plan.Changes.UpdateOld, tc.updateOld)
		validateEntries(t, plan.Changes.UpdateNew, tc.updateNew)
		validateEntries(t, plan.Changes.Delete, tc.delete)
	}
}

// BenchmarkCalculate benchmarks the Calculate method.
func BenchmarkCalculate(b *testing.B) {
	foo := endpoint.NewEndpoint("foo", "v1", "")
	barV1 := endpoint.NewEndpoint("bar", "v1", "")
	barV2 := endpoint.NewEndpoint("bar", "v2", "")
	baz := endpoint.NewEndpoint("baz", "v1", "")

	plan := &Plan{
		Current: []*endpoint.Endpoint{foo, barV1},
		Desired: []*endpoint.Endpoint{barV2, baz},
	}

	for i := 0; i < b.N; i++ {
		plan.Calculate()
	}
}

// ExamplePlan shows how plan can be used.
func ExamplePlan() {
	foo := endpoint.NewEndpoint("foo.example.com", "1.2.3.4", "")
	barV1 := endpoint.NewEndpoint("bar.example.com", "8.8.8.8", "")
	barV2 := endpoint.NewEndpoint("bar.example.com", "8.8.4.4", "")
	baz := endpoint.NewEndpoint("baz.example.com", "6.6.6.6", "")

	// Plan where
	// * foo should be deleted
	// * bar should be updated from v1 to v2
	// * baz should be created
	plan := &Plan{
		Policies: []Policy{&SyncPolicy{}},
		Current:  []*endpoint.Endpoint{foo, barV1},
		Desired:  []*endpoint.Endpoint{barV2, baz},
	}

	// calculate actions
	plan = plan.Calculate()

	// print actions
	fmt.Println("Create:")
	for _, ep := range plan.Changes.Create {
		fmt.Println(ep)
	}
	fmt.Println("UpdateOld:")
	for _, ep := range plan.Changes.UpdateOld {
		fmt.Println(ep)
	}
	fmt.Println("UpdateNew:")
	for _, ep := range plan.Changes.UpdateNew {
		fmt.Println(ep)
	}
	fmt.Println("Delete:")
	for _, ep := range plan.Changes.Delete {
		fmt.Println(ep)
	}
	// Create:
	// &{baz.example.com 6.6.6.6 map[] }
	// UpdateOld:
	// &{bar.example.com 8.8.8.8 map[] }
	// UpdateNew:
	// &{bar.example.com 8.8.4.4 map[] }
	// Delete:
	// &{foo.example.com 1.2.3.4 map[] }
}

// validateEntries validates that the list of entries matches expected.
func validateEntries(t *testing.T, entries, expected []*endpoint.Endpoint) {
	if len(entries) != len(expected) {
		t.Fatalf("expected %q to match %q", entries, expected)
	}

	for i := range entries {
		if !testutils.SameEndpoint(entries[i], expected[i]) {
			t.Fatalf("expected %q to match %q", entries, expected)
		}
	}
}

func newEndpointWithOwner(dnsName, target, ownerID string) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, target, endpoint.RecordTypeCNAME)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	return e
}

func newEndpointWithOwnerAndTTL(dnsName, target, ownerID string, ttl endpoint.TTL) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, target, endpoint.RecordTypeCNAME)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	e.RecordTTL = ttl
	return e
}
