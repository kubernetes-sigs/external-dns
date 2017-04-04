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

package provider

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"

	dns "google.golang.org/api/dns/v1"
	googleapi "google.golang.org/api/googleapi"
)

var (
	expectedZones      = []*dns.ManagedZone{{Name: "expected"}}
	expectedRecordSets = []*dns.ResourceRecordSet{
		{
			Type:    "A",
			Name:    "expected",
			Rrdatas: []string{"target"},
		},
		{
			Type:    "CNAME",
			Name:    "unexpected",
			Rrdatas: []string{"target"},
		},
	}
)

type mockManagedZonesCreateCall struct{}

func (m *mockManagedZonesCreateCall) Do(opts ...googleapi.CallOption) (*dns.ManagedZone, error) {
	return nil, nil
}

type mockErrManagedZonesCreateCall struct{}

func (m *mockErrManagedZonesCreateCall) Do(opts ...googleapi.CallOption) (*dns.ManagedZone, error) {
	return nil, fmt.Errorf("failed")
}

type mockManagedZonesDeleteCall struct{}

func (m *mockManagedZonesDeleteCall) Do(opts ...googleapi.CallOption) error {
	return nil
}

type mockErrManagedZonesDeleteCall struct{}

func (m *mockErrManagedZonesDeleteCall) Do(opts ...googleapi.CallOption) error {
	return fmt.Errorf("failed")
}

type mockManagedZonesListCall struct{}

func (m *mockManagedZonesListCall) Pages(ctx context.Context, f func(*dns.ManagedZonesListResponse) error) error {
	return f(&dns.ManagedZonesListResponse{ManagedZones: expectedZones})
}

type mockErrManagedZonesListCall struct{}

func (m *mockErrManagedZonesListCall) Pages(ctx context.Context, f func(*dns.ManagedZonesListResponse) error) error {
	return fmt.Errorf("failed")
}

type mockManagedZonesClient struct{}

func (m *mockManagedZonesClient) Create(project string, managedzone *dns.ManagedZone) managedZonesCreateCallInterface {
	return &mockManagedZonesCreateCall{}
}

func (m *mockManagedZonesClient) Delete(project string, managedZone string) managedZonesDeleteCallInterface {
	return &mockManagedZonesDeleteCall{}
}

func (m *mockManagedZonesClient) List(project string) managedZonesListCallInterface {
	return &mockManagedZonesListCall{}
}

type mockErrManagedZonesClient struct{}

func (m *mockErrManagedZonesClient) Create(project string, managedzone *dns.ManagedZone) managedZonesCreateCallInterface {
	return &mockErrManagedZonesCreateCall{}
}

func (m *mockErrManagedZonesClient) Delete(project string, managedZone string) managedZonesDeleteCallInterface {
	return &mockErrManagedZonesDeleteCall{}
}

func (m *mockErrManagedZonesClient) List(project string) managedZonesListCallInterface {
	return &mockErrManagedZonesListCall{}
}

type mockResourceRecordSetsListCall struct{}

func (m *mockResourceRecordSetsListCall) Pages(ctx context.Context, f func(*dns.ResourceRecordSetsListResponse) error) error {
	return f(&dns.ResourceRecordSetsListResponse{Rrsets: expectedRecordSets})
}

type mockErrResourceRecordSetsListCall struct{}

func (m *mockErrResourceRecordSetsListCall) Pages(ctx context.Context, f func(*dns.ResourceRecordSetsListResponse) error) error {
	return fmt.Errorf("failed")
}

type mockResourceRecordSetsClient struct{}

func (m *mockResourceRecordSetsClient) List(project string, managedZone string) resourceRecordSetsListCallInterface {
	return &mockResourceRecordSetsListCall{}
}

type mockErrResourceRecordSetsClient struct{}

func (m *mockErrResourceRecordSetsClient) List(project string, managedZone string) resourceRecordSetsListCallInterface {
	return &mockErrResourceRecordSetsListCall{}
}

type mockChangesCreateCall struct{}

func (m *mockChangesCreateCall) Do(opts ...googleapi.CallOption) (*dns.Change, error) {
	return nil, nil
}

type mockErrChangesCreateCall struct{}

func (m *mockErrChangesCreateCall) Do(opts ...googleapi.CallOption) (*dns.Change, error) {
	return nil, fmt.Errorf("failed")
}

type mockChangesClient struct{}

func (m *mockChangesClient) Create(project string, managedZone string, change *dns.Change) changesCreateCallInterface {
	return &mockChangesCreateCall{}
}

type mockErrChangesClient struct{}

func (m *mockErrChangesClient) Create(project string, managedZone string, change *dns.Change) changesCreateCallInterface {
	return &mockErrChangesCreateCall{}
}

func TestGoogleZones(t *testing.T) {
	provider := &googleProvider{
		project:            "project",
		managedZonesClient: &mockManagedZonesClient{},
	}

	zones, err := provider.Zones()
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}

	if len(zones) != len(expectedZones) {
		t.Errorf("expected %d zones, got %d", len(expectedZones), len(zones))
	}

	provider.managedZonesClient = &mockErrManagedZonesClient{}

	_, err = provider.Zones()
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestGoogleCreateZone(t *testing.T) {
	provider := &googleProvider{
		project:            "project",
		managedZonesClient: &mockManagedZonesClient{},
	}

	err := provider.CreateZone("name", "domain")
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}

	provider.managedZonesClient = &mockErrManagedZonesClient{}

	err = provider.CreateZone("name", "domain")
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestGoogleDeleteZone(t *testing.T) {
	provider := &googleProvider{
		project:            "project",
		managedZonesClient: &mockManagedZonesClient{},
	}

	err := provider.DeleteZone("name")
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}

	provider.managedZonesClient = &mockErrManagedZonesClient{}

	err = provider.DeleteZone("name")
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestGoogleRecords(t *testing.T) {
	provider := &googleProvider{
		project:                  "project",
		resourceRecordSetsClient: &mockResourceRecordSetsClient{},
	}

	endpoints, err := provider.Records("zone")
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}

	if len(endpoints) != len(expectedRecordSets)-1 {
		t.Errorf("expected %d endpoints, got %d", len(expectedRecordSets)-1, len(endpoints))
	}

	provider.resourceRecordSetsClient = &mockErrResourceRecordSetsClient{}

	_, err = provider.Records("zone")
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestGoogleCreateRecords(t *testing.T) {
	provider := &googleProvider{
		project:       "project",
		changesClient: &mockChangesClient{},
	}

	endpoints := []*endpoint.Endpoint{
		{
			DNSName: "dns-name",
			Target:  "target",
		},
	}

	err := provider.CreateRecords("zone", endpoints)
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}

	provider.changesClient = &mockErrChangesClient{}

	err = provider.CreateRecords("zone", endpoints)
	if err == nil {
		t.Errorf("expected error")
	}
}

func TestGoogleUpdateRecords(t *testing.T) {
	provider := &googleProvider{
		project:       "project",
		changesClient: &mockChangesClient{},
	}

	records := []*endpoint.Endpoint{
		{
			DNSName: "dns-name",
			Target:  "target",
		},
	}

	oldRecords := []*endpoint.Endpoint{
		{
			DNSName: "dns-name",
			Target:  "target",
		},
	}

	err := provider.UpdateRecords("zone", records, oldRecords)
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}

	err = provider.UpdateRecords("zone", nil, nil)
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}

	provider.dryRun = true

	err = provider.UpdateRecords("zone", records, oldRecords)
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}
}

func TestGoogleDeleteRecords(t *testing.T) {
	provider := &googleProvider{
		project:       "project",
		changesClient: &mockChangesClient{},
	}

	endpoints := []*endpoint.Endpoint{
		{
			DNSName: "dns-name",
			Target:  "target",
		},
	}

	err := provider.DeleteRecords("zone", endpoints)
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}
}

func TestGoogleApplyChanges(t *testing.T) {
	provider := &googleProvider{
		project:       "project",
		changesClient: &mockChangesClient{},
	}

	changes := &plan.Changes{}

	err := provider.ApplyChanges("zone", changes)
	if err != nil {
		t.Errorf("should not fail: %s", err)
	}
}
