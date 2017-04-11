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

package controller

import (
	"errors"
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
	"github.com/kubernetes-incubator/external-dns/registry"
	"github.com/kubernetes-incubator/external-dns/source"
)

// mockProvider returns mock endpoints and validates changes.
type mockProvider struct {
	RecordsStore  []*endpoint.Endpoint
	ExpectZone    string
	ExpectChanges *plan.Changes
}

// Records returns the desired mock endpoints.
func (p *mockProvider) Records(zone string) ([]*endpoint.Endpoint, error) {
	return p.RecordsStore, nil
}

// ApplyChanges validates that the passed in changes satisfy the assumtions.
func (p *mockProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	if zone != p.ExpectZone {
		return errors.New("zone is incorrect")
	}

	if len(changes.Create) != len(p.ExpectChanges.Create) {
		return errors.New("number of created records is wrong")
	}

	for i := range changes.Create {
		if changes.Create[i].DNSName != p.ExpectChanges.Create[i].DNSName || changes.Create[i].Target != p.ExpectChanges.Create[i].Target {
			return errors.New("created record is wrong")
		}
	}

	for i := range changes.UpdateNew {
		if changes.UpdateNew[i].DNSName != p.ExpectChanges.UpdateNew[i].DNSName || changes.UpdateNew[i].Target != p.ExpectChanges.UpdateNew[i].Target {
			return errors.New("delete record is wrong")
		}
	}

	for i := range changes.UpdateOld {
		if changes.UpdateOld[i].DNSName != p.ExpectChanges.UpdateOld[i].DNSName || changes.UpdateOld[i].Target != p.ExpectChanges.UpdateOld[i].Target {
			return errors.New("delete record is wrong")
		}
	}

	for i := range changes.Delete {
		if changes.Delete[i].DNSName != p.ExpectChanges.Delete[i].DNSName || changes.Delete[i].Target != p.ExpectChanges.Delete[i].Target {
			return errors.New("delete record is wrong")
		}
	}

	return nil
}

// newMockProvider creates a new mockProvider returning the given endpoints and validating the desired changes.
func newMockProvider(endpoints []*endpoint.Endpoint, zone string, changes *plan.Changes) provider.Provider {
	dnsProvider := &mockProvider{
		RecordsStore:  endpoints,
		ExpectZone:    zone,
		ExpectChanges: changes,
	}

	return dnsProvider
}

// TestRunOnce tests that RunOnce correctly orchestrates the different components.
func TestRunOnce(t *testing.T) {
	// Fake some desired endpoints coming from our source.
	source := source.NewMockSource(
		[]*endpoint.Endpoint{
			{
				DNSName: "create-record",
				Target:  "1.2.3.4",
			},
			{
				DNSName: "update-record",
				Target:  "8.8.4.4",
			},
		},
	)

	// Fake some existing records in our DNS provider and validate some desired changes.
	provider := newMockProvider(
		[]*endpoint.Endpoint{
			{
				DNSName: "update-record",
				Target:  "8.8.8.8",
			},
			{
				DNSName: "delete-record",
				Target:  "4.3.2.1",
			},
		},
		"test-zone",
		&plan.Changes{
			Create: []*endpoint.Endpoint{
				{DNSName: "create-record", Target: "1.2.3.4"},
			},
			UpdateNew: []*endpoint.Endpoint{
				{DNSName: "update-record", Target: "8.8.4.4"},
			},
			UpdateOld: []*endpoint.Endpoint{
				{DNSName: "update-record", Target: "8.8.8.8"},
			},
			Delete: []*endpoint.Endpoint{
				{DNSName: "delete-record", Target: "4.3.2.1"},
			},
		},
	)

	r, _ := registry.NewNoopRegistry(provider)

	// Run our controller once to trigger the validation.
	ctrl := &Controller{
		Zone:     "test-zone",
		Source:   source,
		Registry: r,
		Policy:   &plan.SyncPolicy{},
	}

	err := ctrl.RunOnce()

	if err != nil {
		t.Fatal(err)
	}
}
