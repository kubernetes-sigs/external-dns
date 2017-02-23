package plan

import (
	"fmt"
	"testing"
)

// TestCalculate tests that a plan can calculate actions to move a list of
// current records to a list of desired records.
func TestCalculate(t *testing.T) {
	// empty list of records
	empty := []DNSRecord{}
	// a simple entry
	fooV1 := []DNSRecord{{DNSName: "foo", Target: "v1"}}
	// the same entry but with different target
	fooV2 := []DNSRecord{{DNSName: "foo", Target: "v2"}}
	// another simple entry
	bar := []DNSRecord{{DNSName: "bar", Target: "v1"}}

	for _, tc := range []struct {
		current, desired, create, updateOld, updateNew, delete []DNSRecord
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
	foo := DNSRecord{DNSName: "foo", Target: "v1"}
	barV1 := DNSRecord{DNSName: "bar", Target: "v1"}
	barV2 := DNSRecord{DNSName: "bar", Target: "v2"}
	baz := DNSRecord{DNSName: "baz", Target: "v1"}

	plan := &Plan{
		Current: []DNSRecord{foo, barV1},
		Desired: []DNSRecord{barV2, baz},
	}

	for i := 0; i < b.N; i++ {
		plan.Calculate()
	}
}

// ExamplePlan shows how plan can be used.
func ExamplePlan() {
	foo := DNSRecord{DNSName: "foo.example.com", Target: "1.2.3.4"}
	barV1 := DNSRecord{DNSName: "bar.example.com", Target: "8.8.8.8"}
	barV2 := DNSRecord{DNSName: "bar.example.com", Target: "8.8.4.4"}
	baz := DNSRecord{DNSName: "baz.example.com", Target: "6.6.6.6"}

	// Plan where
	// * foo should be deleted
	// * bar should be updated from v1 to v2
	// * baz should be created
	plan := &Plan{
		Current: []DNSRecord{foo, barV1},
		Desired: []DNSRecord{barV2, baz},
	}

	// calculate actions
	plan = plan.Calculate()

	// print actions
	fmt.Println("Create:", plan.Changes.Create)
	fmt.Println("UpdateOld:", plan.Changes.UpdateOld)
	fmt.Println("UpdateNew:", plan.Changes.UpdateNew)
	fmt.Println("Delete:", plan.Changes.Delete)
	// Output:
	// Create: [{baz.example.com 6.6.6.6}]
	// UpdateOld: [{bar.example.com 8.8.8.8}]
	// UpdateNew: [{bar.example.com 8.8.4.4}]
	// Delete: [{foo.example.com 1.2.3.4}]
}

// validateEntries validates that the list of entries matches expected.
func validateEntries(t *testing.T, entries, expected []DNSRecord) {
	if len(entries) != len(expected) {
		t.Fatalf("expected %q to match %q", entries, expected)
	}

	for i := range entries {
		if entries[i] != expected[i] {
			t.Fatalf("expected %q to match %q", entries, expected)
		}
	}
}
