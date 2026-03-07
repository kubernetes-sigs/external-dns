/*
Copyright 2025 The Kubernetes Authors.

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

type testPiholeClientV6 struct {
	endpoints []*endpoint.Endpoint
	requests  *requestTrackerV6
	trigger   string
}

func (t *testPiholeClientV6) listRecords(_ context.Context, rtype string) ([]*endpoint.Endpoint, error) {
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

func (t *testPiholeClientV6) createRecord(_ context.Context, ep *endpoint.Endpoint) error {
	t.endpoints = append(t.endpoints, ep)
	t.requests.createRequests = append(t.requests.createRequests, ep)
	return nil
}

func (t *testPiholeClientV6) deleteRecord(_ context.Context, ep *endpoint.Endpoint) error {
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

type requestTrackerV6 struct {
	createRequests []*endpoint.Endpoint
	deleteRequests []*endpoint.Endpoint
}

func (r *requestTrackerV6) clear() {
	r.createRequests = nil
	r.deleteRequests = nil
}

func TestErrorHandling(t *testing.T) {
	requests := requestTrackerV6{}
	p := &PiholeProvider{
		api:        &testPiholeClientV6{endpoints: make([]*endpoint.Endpoint, 0), requests: &requests},
		apiVersion: "6",
	}

	p.api.(*testPiholeClientV6).trigger = "AERROR"
	_, err := p.Records(context.Background())
	if err.Error() != "AERROR" {
		t.Fatal(err)
	}

	p.api.(*testPiholeClientV6).trigger = "AAAAERROR"
	_, err = p.Records(context.Background())
	if err.Error() != "AAAAERROR" {
		t.Fatal(err)
	}

	p.api.(*testPiholeClientV6).trigger = "CNAMEERROR"
	_, err = p.Records(context.Background())
	if err.Error() != "CNAMEERROR" {
		t.Fatal(err)
	}

}

func TestNewPiholeProviderV6(t *testing.T) {
	// Test invalid configuration
	_, err := NewPiholeProvider(PiholeConfig{APIVersion: "7"})
	if err == nil {
		t.Error("Expected error from invalid configuration")
	}
	// Test valid configuration
	_, err = NewPiholeProvider(PiholeConfig{Server: "test.example.com", APIVersion: "6"})
	if err != nil {
		t.Error("Expected no error from valid configuration, got:", err)
	}
}

func TestProviderV6(t *testing.T) {
	requests := requestTrackerV6{}
	p := &PiholeProvider{
		api:        &testPiholeClientV6{endpoints: make([]*endpoint.Endpoint, 0), requests: &requests},
		apiVersion: "6",
	}

	t.Run("Initial Records", func(t *testing.T) {
		records, err := p.Records(context.Background())
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
		if err := p.ApplyChanges(context.Background(), &plan.Changes{Create: records}); err != nil {
			t.Fatal(err)
		}

		newRecords, err := p.Records(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if !cmp.Equal(newRecords, records, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort) {
			t.Error("Records are not equal:", cmp.Diff(newRecords, records, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort))
		}
		if !cmp.Equal(requests.createRequests, records, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort) {
			t.Error("Create requests are not equal:", cmp.Diff(requests.createRequests, records, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort))
		}
		if len(requests.deleteRequests) != 0 {
			t.Fatal("Expected no delete requests, got:", requests.deleteRequests)
		}
		requests.clear()
	})

	t.Run("Delete Records", func(t *testing.T) {
		recordToDeleteA := &endpoint.Endpoint{DNSName: "test3.example.com", Targets: []string{"192.168.1.3"}, RecordType: endpoint.RecordTypeA}
		if err := p.ApplyChanges(context.Background(), &plan.Changes{Delete: []*endpoint.Endpoint{recordToDeleteA}}); err != nil {
			t.Fatal(err)
		}
		recordToDeleteAAAA := &endpoint.Endpoint{DNSName: "test3.example.com", Targets: []string{"fc00::1:192:168:1:3"}, RecordType: endpoint.RecordTypeAAAA}
		if err := p.ApplyChanges(context.Background(), &plan.Changes{Delete: []*endpoint.Endpoint{recordToDeleteAAAA}}); err != nil {
			t.Fatal(err)
		}

		expectedRecords := []*endpoint.Endpoint{
			{DNSName: "test1.example.com", Targets: []string{"192.168.1.1"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test2.example.com", Targets: []string{"192.168.1.2"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test1.example.com", Targets: []string{"fc00::1:192:168:1:1"}, RecordType: endpoint.RecordTypeAAAA},
			{DNSName: "test2.example.com", Targets: []string{"fc00::1:192:168:1:2"}, RecordType: endpoint.RecordTypeAAAA},
		}
		newRecords, err := p.Records(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if !cmp.Equal(newRecords, expectedRecords, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort) {
			t.Error("Records are not equal:", cmp.Diff(newRecords, expectedRecords, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort))
		}
		if len(requests.createRequests) != 0 {
			t.Fatal("Expected no create requests, got:", requests.createRequests)
		}
		expectedDeletes := []*endpoint.Endpoint{recordToDeleteA, recordToDeleteAAAA}
		if !cmp.Equal(requests.deleteRequests, expectedDeletes, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort) {
			t.Error("Delete requests are not equal:", cmp.Diff(requests.deleteRequests, expectedDeletes, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort))
		}
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
		if err := p.ApplyChanges(context.Background(), &plan.Changes{UpdateOld: updateOld, UpdateNew: updateNew}); err != nil {
			t.Fatal(err)
		}

		expectedRecords := []*endpoint.Endpoint{
			{DNSName: "test1.example.com", Targets: []string{"192.168.1.1"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test2.example.com", Targets: []string{"10.0.0.1"}, RecordType: endpoint.RecordTypeA},
			{DNSName: "test1.example.com", Targets: []string{"fc00::1:192:168:1:1"}, RecordType: endpoint.RecordTypeAAAA},
			{DNSName: "test2.example.com", Targets: []string{"fc00::1:10:0:0:1"}, RecordType: endpoint.RecordTypeAAAA},
		}
		newRecords, err := p.Records(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if !cmp.Equal(newRecords, expectedRecords, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort) {
			t.Error("Records are not equal:", cmp.Diff(newRecords, expectedRecords, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort))
		}
		if !cmp.Equal(requests.createRequests, updateNew, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort) {
			t.Error("Create requests are not equal:", cmp.Diff(requests.createRequests, updateNew, cmpopts.IgnoreUnexported(endpoint.Endpoint{}), endpointSort))
		}
		if !cmp.Equal(requests.deleteRequests, updateOld, cmpopts.IgnoreUnexported(endpoint.Endpoint{})) {
			t.Error("Delete requests are not equal:", cmp.Diff(requests.deleteRequests, updateOld, cmpopts.IgnoreUnexported(endpoint.Endpoint{})))
		}
		requests.clear()
	})
}

func TestProviderV6MultipleTargets(t *testing.T) {
	requests := requestTrackerV6{}
	p := &PiholeProvider{
		api:        &testPiholeClientV6{endpoints: make([]*endpoint.Endpoint, 0), requests: &requests},
		apiVersion: "6",
	}

	t.Run("Update with multiple targets - merge and deduplicate", func(t *testing.T) {
		// Create initial records with multiple targets
		initialRecords := []*endpoint.Endpoint{
			{DNSName: "multi.example.com", Targets: []string{"192.168.1.1", "192.168.1.2"}, RecordType: endpoint.RecordTypeA},
		}
		if err := p.ApplyChanges(context.Background(), &plan.Changes{Create: initialRecords}); err != nil {
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
		if err := p.ApplyChanges(context.Background(), &plan.Changes{UpdateOld: updateOld, UpdateNew: updateNew}); err != nil {
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
		if err := p.ApplyChanges(context.Background(), &plan.Changes{UpdateOld: updateOld, UpdateNew: updateNew}); err != nil {
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
