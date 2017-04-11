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
	// empty list of records
	empty := []*endpoint.Endpoint{}
	// a simple entry
	fooV1 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v1", "CNAME")}
	// the same entry but with different target
	fooV2 := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v2", "CNAME")}
	// another simple entry
	bar := []*endpoint.Endpoint{endpoint.NewEndpoint("bar", "v1", "CNAME")}

	// test case with labels
	noLabels := []*endpoint.Endpoint{endpoint.NewEndpoint("foo", "v2", "CNAME")}
	labeledV2 := []*endpoint.Endpoint{newEndpointWithOwner("foo", "v2", "123")}
	labeledV1 := []*endpoint.Endpoint{newEndpointWithOwner("foo", "v1", "123")}

	for _, tc := range []struct {
		current, desired, create, updateOld, updateNew, delete []*endpoint.Endpoint
	}{
		// Nothing exists and nothing desired doesn't change anything.
		{empty, empty, empty, empty, empty, empty},
		// More desired than current creates the desired.
		{empty, fooV1, fooV1, empty, empty, empty},
		// Desired equals current doesn't change anything.
		{fooV1, fooV1, empty, empty, empty, empty},
		// Nothing is desired deletes the current.
		{fooV1, empty, empty, empty, empty, fooV1},
		// Current and desired match but Target is different triggers an update.
		{fooV1, fooV2, empty, fooV1, fooV2, empty},
		// Both exist but are different creates desired and deletes current.
		{fooV1, bar, bar, empty, empty, fooV1},
		// Labels should be inherited
		{labeledV1, noLabels, empty, labeledV1, labeledV2, empty},
	} {
		// setup plan
		plan := &Plan{
			Current: tc.current,
			Desired: tc.desired,
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
		Current: []*endpoint.Endpoint{foo, barV1},
		Desired: []*endpoint.Endpoint{barV2, baz},
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
	e := endpoint.NewEndpoint(dnsName, target, "CNAME")
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	return e
}
