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

package pihole

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

var (
	endpointSort = cmpopts.SortSlices(func(x, y *endpoint.Endpoint) bool {
		if x.DNSName < y.DNSName {
			return true
		}
		if x.DNSName > y.DNSName {
			return false
		}
		if x.RecordType < y.RecordType {
			return true
		}
		if x.RecordType > y.RecordType {
			return false
		}
		return x.Targets.String() < y.Targets.String()
	})
)

func assertEndpointsEqual(t *testing.T, got, want []*endpoint.Endpoint) {
	t.Helper()
	if diff := cmp.Diff(want, got, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort); diff != "" {
		t.Error(diff)
	}
}

type testPiholeClient struct {
	endpoints []*endpoint.Endpoint
	requests  *requestTracker
	trigger   string
}

func (t *testPiholeClient) listRecords(_ context.Context, rtype string) ([]*endpoint.Endpoint, error) {
	out := make([]*endpoint.Endpoint, 0)
	if t.trigger == "AERROR" {
		return nil, errors.New("AERROR")
	}
	if t.trigger == "AAAAERROR" {
		return nil, errors.New("AAAAERROR")
	}
	if t.trigger == "CNAMEERROR" {
		return nil, errors.New("CNAMEERROR")
	}
	for _, ep := range t.endpoints {
		if ep.RecordType == rtype {
			out = append(out, ep)
		}
	}
	return out, nil
}

func (t *testPiholeClient) createRecord(_ context.Context, ep *endpoint.Endpoint) error {
	t.endpoints = append(t.endpoints, ep)
	t.requests.createRequests = append(t.requests.createRequests, ep)
	return nil
}

func (t *testPiholeClient) deleteRecord(_ context.Context, ep *endpoint.Endpoint) error {
	newEPs := make([]*endpoint.Endpoint, 0)
	for _, existing := range t.endpoints {
		if existing.DNSName != ep.DNSName || cmp.Diff(existing.Targets, ep.Targets) != "" || existing.RecordType != ep.RecordType {
			newEPs = append(newEPs, existing)
		}
	}
	t.endpoints = newEPs
	t.requests.deleteRequests = append(t.requests.deleteRequests, ep)
	return nil
}

type requestTracker struct {
	createRequests []*endpoint.Endpoint
	deleteRequests []*endpoint.Endpoint
}

func (r *requestTracker) clear() {
	r.createRequests = nil
	r.deleteRequests = nil
}

func TestErrorHandling(t *testing.T) {
	requests := requestTracker{}
	p := &PiholeProvider{
		api: &testPiholeClient{endpoints: make([]*endpoint.Endpoint, 0), requests: &requests},
	}

	p.api.(*testPiholeClient).trigger = "AERROR"
	_, err := p.Records(t.Context())
	if err.Error() != "AERROR" {
		t.Fatal(err)
	}

	p.api.(*testPiholeClient).trigger = "AAAAERROR"
	_, err = p.Records(t.Context())
	if err.Error() != "AAAAERROR" {
		t.Fatal(err)
	}

	p.api.(*testPiholeClient).trigger = "CNAMEERROR"
	_, err = p.Records(t.Context())
	if err.Error() != "CNAMEERROR" {
		t.Fatal(err)
	}

}

func TestNewPiholeProvider(t *testing.T) {
	// Test invalid configuration
	_, err := newProvider(PiholeConfig{})
	if err == nil {
		t.Error("Expected error from invalid configuration")
	}
	// Test valid configuration
	_, err = newProvider(PiholeConfig{Server: "test.example.com"})
	if err != nil {
		t.Error("Expected no error from valid configuration, got:", err)
	}
}

func TestProvider(t *testing.T) {
	requests := requestTracker{}
	p := &PiholeProvider{
		api: &testPiholeClient{endpoints: make([]*endpoint.Endpoint, 0), requests: &requests},
	}

	t.Run("Initial Records", func(t *testing.T) {
		records, err := p.Records(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		if len(records) != 0 {
			t.Fatal("Expected empty list of records, got:", records)
		}
	})

	t.Run("Create Records", func(t *testing.T) {
		records := []*endpoint.Endpoint{
			{DNSName: "test1.example.com", Targets: []string{"192.168.1.1"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test2.example.com", Targets: []string{"192.168.1.2"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test3.example.com", Targets: []string{"192.168.1.3"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test1.example.com", Targets: []string{"fc00::1:192:168:1:1"}, RecordType: endpoint.RecordTypeAAAA},
			{DNSName: "test2.example.com", Targets: []string{"fc00::1:192:168:1:2"}, RecordType: endpoint.RecordTypeAAAA},
			{DNSName: "test3.example.com", Targets: []string{"fc00::1:192:168:1:3"}, RecordType: endpoint.RecordTypeAAAA},
		}
		if err := p.ApplyChanges(t.Context(), &plan.Changes{Create: records}); err != nil {
			t.Fatal(err)
		}

		newRecords, err := p.Records(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		assertEndpointsEqual(t, newRecords, records)
		assertEndpointsEqual(t, requests.createRequests, records)
		if len(requests.deleteRequests) != 0 {
			t.Fatal("Expected no delete requests, got:", requests.deleteRequests)
		}
		requests.clear()
	})

	t.Run("Delete Records", func(t *testing.T) {
		recordToDeleteA := &endpoint.Endpoint{DNSName: "test3.example.com", Targets: []string{"192.168.1.3"}, RecordType: endpoint.RecordTypeA}
		if err := p.ApplyChanges(t.Context(), &plan.Changes{Delete: []*endpoint.Endpoint{recordToDeleteA}}); err != nil {
			t.Fatal(err)
		}
		recordToDeleteAAAA := &endpoint.Endpoint{DNSName: "test3.example.com", Targets: []string{"fc00::1:192:168:1:3"}, RecordType: endpoint.RecordTypeAAAA}
		if err := p.ApplyChanges(t.Context(), &plan.Changes{Delete: []*endpoint.Endpoint{recordToDeleteAAAA}}); err != nil {
			t.Fatal(err)
		}

		expectedRecords := []*endpoint.Endpoint{
			{DNSName: "test1.example.com", Targets: []string{"192.168.1.1"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test2.example.com", Targets: []string{"192.168.1.2"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test1.example.com", Targets: []string{"fc00::1:192:168:1:1"}, RecordType: endpoint.RecordTypeAAAA},
			{DNSName: "test2.example.com", Targets: []string{"fc00::1:192:168:1:2"}, RecordType: endpoint.RecordTypeAAAA},
		}
		newRecords, err := p.Records(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		assertEndpointsEqual(t, newRecords, expectedRecords)
		if len(requests.createRequests) != 0 {
			t.Fatal("Expected no create requests, got:", requests.createRequests)
		}
		assertEndpointsEqual(t, requests.deleteRequests, []*endpoint.Endpoint{recordToDeleteA, recordToDeleteAAAA})
		requests.clear()
	})

	t.Run("Update Records", func(t *testing.T) {
		updateOld := []*endpoint.Endpoint{
			{DNSName: "test2.example.com", Targets: []string{"192.168.1.2"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test2.example.com", Targets: []string{"fc00::1:192:168:1:2"}, RecordType: endpoint.RecordTypeAAAA},
		}
		updateNew := []*endpoint.Endpoint{
			{DNSName: "test2.example.com", Targets: []string{"10.0.0.1"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test2.example.com", Targets: []string{"fc00::1:10:0:0:1"}, RecordType: endpoint.RecordTypeAAAA},
		}
		if err := p.ApplyChanges(t.Context(), &plan.Changes{UpdateOld: updateOld, UpdateNew: updateNew}); err != nil {
			t.Fatal(err)
		}

		expectedRecords := []*endpoint.Endpoint{
			{DNSName: "test1.example.com", Targets: []string{"192.168.1.1"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test2.example.com", Targets: []string{"10.0.0.1"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test1.example.com", Targets: []string{"fc00::1:192:168:1:1"}, RecordType: endpoint.RecordTypeAAAA},
			{DNSName: "test2.example.com", Targets: []string{"fc00::1:10:0:0:1"}, RecordType: endpoint.RecordTypeAAAA},
		}
		newRecords, err := p.Records(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		assertEndpointsEqual(t, newRecords, expectedRecords)
		assertEndpointsEqual(t, requests.createRequests, updateNew)
		assertEndpointsEqual(t, requests.deleteRequests, updateOld)
		requests.clear()
	})
}

func TestProviderMultipleTargets(t *testing.T) {
	requests := requestTracker{}
	p := &PiholeProvider{
		api: &testPiholeClient{endpoints: make([]*endpoint.Endpoint, 0), requests: &requests},
	}

	t.Run("Update with multiple targets - merge and deduplicate", func(t *testing.T) {
		// Create initial records with multiple targets
		initialRecords := []*endpoint.Endpoint{
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.1", "192.168.1.2"}, RecordType: endpoint.RecordTypeA},
		}
		if err := p.ApplyChanges(t.Context(), &plan.Changes{Create: initialRecords}); err != nil {
			t.Fatal(err)
		}
		requests.clear()

		// Update with new targets that should be merged
		updateOld := []*endpoint.Endpoint{
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.1", "192.168.1.2"}, RecordType: endpoint.RecordTypeA},
		}
		updateNew := []*endpoint.Endpoint{
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.3"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.4"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.3"}, RecordType: endpoint.RecordTypeA}, // Duplicate to test deduplication
		}
		if err := p.ApplyChanges(t.Context(), &plan.Changes{UpdateOld: updateOld, UpdateNew: updateNew}); err != nil {
			t.Fatal(err)
		}

		// Verify that targets were merged and deduplicated
		expectedCreate := []*endpoint.Endpoint{
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.3", "192.168.1.4"}, RecordType: endpoint.RecordTypeA},
		}
		if len(requests.createRequests) != 1 {
			t.Fatalf("Expected 1 create request, got %d", len(requests.createRequests))
		}
		if !cmp.Equal(requests.createRequests[0].Targets, expectedCreate[0].Targets) {
			t.Error("Targets not merged correctly:", cmp.Diff(requests.createRequests[0].Targets, expectedCreate[0].Targets))
		}
		if len(requests.deleteRequests) != 1 {
			t.Fatalf("Expected 1 delete request, got %d", len(requests.deleteRequests))
		}
		requests.clear()
	})

	t.Run("Update with exact match - should skip delete", func(t *testing.T) {
		// Update where old and new have the same targets (exact match)
		updateOld := []*endpoint.Endpoint{
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.3", "192.168.1.4"}, RecordType: endpoint.RecordTypeA},
		}
		updateNew := []*endpoint.Endpoint{
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.3", "192.168.1.4"}, RecordType: endpoint.RecordTypeA},
		}
		if err := p.ApplyChanges(t.Context(), &plan.Changes{UpdateOld: updateOld, UpdateNew: updateNew}); err != nil {
			t.Fatal(err)
		}

		// Should not create or delete anything since targets are the same
		if len(requests.createRequests) != 0 {
			t.Fatalf("Expected no create requests for exact match, got %d", len(requests.createRequests))
		}
		if len(requests.deleteRequests) != 0 {
			t.Fatalf("Expected no delete requests for exact match, got %d", len(requests.deleteRequests))
		}
		requests.clear()
	})
	
	t.Run("Update with unordered match - should skip delete", func(t *testing.T) {
		// Update where old and new have the same targets (exact match)
		updateOld := []*endpoint.Endpoint{
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.4", "192.168.1.3"}, RecordType: endpoint.RecordTypeA},
		}
		updateNew := []*endpoint.Endpoint{
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.3", "192.168.1.4"}, RecordType: endpoint.RecordTypeA},
		}
		if err := p.ApplyChanges(t.Context(), &plan.Changes{UpdateOld: updateOld, UpdateNew: updateNew}); err != nil {
			t.Fatal(err)
		}

		// Should not create or delete anything since targets are the same
		if len(requests.createRequests) != 0 {
			t.Fatalf("Expected no create requests for exact match, got %d", len(requests.createRequests))
		}
		if len(requests.deleteRequests) != 0 {
			t.Fatalf("Expected no delete requests for exact match, got %d", len(requests.deleteRequests))
		}
		requests.clear()
	})
}
