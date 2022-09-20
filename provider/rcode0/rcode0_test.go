/*
Copyright 2019 The Kubernetes Authors.

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

package rcode0

import (
	"context"
	"fmt"
	"os"
	"testing"

	rc0 "github.com/nic-at/rc0go"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const (
	testZoneOne = "testzone1.at"
	testZoneTwo = "testzone2.at"

	rrsetChangesUnsupportedChangeType = 0
)

type mockRcodeZeroClient rc0.Client

type mockZoneManagementService struct {
	TestNilZonesReturned bool
	TestErrorReturned    bool
}

type mockRRSetService struct {
	TestErrorReturned bool
}

func (m *mockZoneManagementService) resetTestConditions() {
	m.TestNilZonesReturned = false
	m.TestErrorReturned = false
}

func TestRcodeZeroProvider_Records(t *testing.T) {
	mockRRSetService := &mockRRSetService{}
	mockZoneManagementService := &mockZoneManagementService{}

	provider := &RcodeZeroProvider{
		Client: (*rc0.Client)(&mockRcodeZeroClient{
			Zones: mockZoneManagementService,
			RRSet: mockRRSetService,
		}),
	}

	ctx := context.Background()

	endpoints, err := provider.Records(ctx) // should return 6 rrs
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	require.Equal(t, 10, len(endpoints))

	mockRRSetService.TestErrorReturned = true

	_, err = provider.Records(ctx)
	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func TestRcodeZeroProvider_ApplyChanges(t *testing.T) {
	mockRRSetService := &mockRRSetService{}
	mockZoneManagementService := &mockZoneManagementService{}

	provider := &RcodeZeroProvider{
		Client: (*rc0.Client)(&mockRcodeZeroClient{
			Zones: mockZoneManagementService,
			RRSet: mockRRSetService,
		}),
		DomainFilter: endpoint.NewDomainFilter([]string{testZoneOne}),
	}

	changes := mockChanges()

	err := provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestRcodeZeroProvider_NewRcodezeroChanges(t *testing.T) {
	provider := &RcodeZeroProvider{}

	changes := mockChanges()

	createChanges := provider.NewRcodezeroChanges(testZoneOne, changes.Create)
	require.Equal(t, 4, len(createChanges))

	deleteChanges := provider.NewRcodezeroChanges(testZoneOne, changes.Delete)
	require.Equal(t, 1, len(deleteChanges))

	updateOldChanges := provider.NewRcodezeroChanges(testZoneOne, changes.UpdateOld)
	require.Equal(t, 1, len(updateOldChanges))

	updateNewChanges := provider.NewRcodezeroChanges(testZoneOne, changes.UpdateNew)
	require.Equal(t, 1, len(updateNewChanges))
}

func TestRcodeZeroProvider_NewRcodezeroChange(t *testing.T) {
	_endpoint := &endpoint.Endpoint{
		RecordType: "A",
		DNSName:    "app." + testZoneOne,
		RecordTTL:  300,
		Targets:    endpoint.Targets{"target"},
	}

	provider := &RcodeZeroProvider{}

	rrsetChange := provider.NewRcodezeroChange(testZoneOne, _endpoint)

	require.Equal(t, _endpoint.RecordType, rrsetChange.Type)
	require.Equal(t, _endpoint.DNSName, rrsetChange.Name)
	require.Equal(t, _endpoint.Targets[0], rrsetChange.Records[0].Content)
	// require.Equal(t, endpoint.RecordTTL, rrsetChange.TTL)
}

func Test_submitChanges(t *testing.T) {
	mockRRSetService := &mockRRSetService{}
	mockZoneManagementService := &mockZoneManagementService{}

	provider := &RcodeZeroProvider{
		Client: (*rc0.Client)(&mockRcodeZeroClient{
			Zones: mockZoneManagementService,
			RRSet: mockRRSetService,
		}),
		DomainFilter: endpoint.NewDomainFilter([]string{testZoneOne}),
	}

	changes := mockRRSetChanges(rrsetChangesUnsupportedChangeType)

	err := provider.submitChanges(changes)

	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func mockRRSetChanges(condition int) []*rc0.RRSetChange {
	switch condition {
	case rrsetChangesUnsupportedChangeType:
		return []*rc0.RRSetChange{
			{
				Name:       testZoneOne,
				Type:       "A",
				ChangeType: "UNSUPPORTED",
				Records:    []*rc0.Record{{Content: "fail"}},
			},
		}
	default:
		return nil
	}
}

func mockChanges() *plan.Changes {
	changes := &plan.Changes{}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "new.ext-dns-test." + testZoneOne, Targets: endpoint.Targets{"target"}, RecordType: "A"},
		{DNSName: "new.ext-dns-test-with-ttl." + testZoneOne, Targets: endpoint.Targets{"target"}, RecordType: "A", RecordTTL: 100},
		{DNSName: "new.ext-dns-test.unexpected.com", Targets: endpoint.Targets{"target"}, RecordType: "AAAA"},
		{DNSName: testZoneOne, Targets: endpoint.Targets{"target"}, RecordType: "CNAME"},
	}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test." + testZoneOne, Targets: endpoint.Targets{"target"}}}
	changes.UpdateOld = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test." + testZoneOne, Targets: endpoint.Targets{"target-old"}}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "foobar.ext-dns-test." + testZoneOne, Targets: endpoint.Targets{"target-new"}, RecordType: "CNAME", RecordTTL: 100}}

	return changes
}

func TestRcodeZeroProvider_Zones(t *testing.T) {
	mockRRSetService := &mockRRSetService{}
	mockZoneManagementService := &mockZoneManagementService{}

	provider := &RcodeZeroProvider{
		Client: (*rc0.Client)(&mockRcodeZeroClient{
			Zones: mockZoneManagementService,
			RRSet: mockRRSetService,
		}),
	}

	mockZoneManagementService.TestNilZonesReturned = true

	zones, err := provider.Zones()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, 0, len(zones))
	mockZoneManagementService.resetTestConditions()

	mockZoneManagementService.TestErrorReturned = true

	_, err = provider.Zones()
	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func TestNewRcodeZeroProvider(t *testing.T) {
	_ = os.Setenv("RC0_API_KEY", "123")
	p, err := NewRcodeZeroProvider(endpoint.NewDomainFilter([]string{"ext-dns-test." + testZoneOne + "."}), true, true)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	require.Equal(t, true, p.DryRun)
	require.Equal(t, true, p.TXTEncrypt)
	require.Equal(t, true, p.DomainFilter.IsConfigured())
	require.Equal(t, false, p.DomainFilter.Match("ext-dns-test."+testZoneTwo+".")) // filter is set, so it should match only provided domains

	p, err = NewRcodeZeroProvider(endpoint.DomainFilter{}, false, false)

	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	require.Equal(t, false, p.DryRun)
	require.Equal(t, false, p.DomainFilter.IsConfigured())
	require.Equal(t, true, p.DomainFilter.Match("ext-dns-test."+testZoneOne+".")) // filter is not set, so it should match any

	_ = os.Unsetenv("RC0_API_KEY")
	_, err = NewRcodeZeroProvider(endpoint.DomainFilter{}, false, false)

	if err == nil {
		t.Errorf("expected to fail")
	}
}

/* mocking mockRRSetServiceInterface */

func (m *mockRRSetService) List(zone string, options *rc0.ListOptions) ([]*rc0.RRType, *rc0.Page, error) {
	if m.TestErrorReturned {
		return nil, nil, fmt.Errorf("operation RRSet.List failed")
	}

	return mockRRSet(zone), nil, nil
}

func mockRRSet(zone string) []*rc0.RRType {
	return []*rc0.RRType{
		{
			Name: "app." + zone + ".",
			Type: "TXT",
			TTL:  300,
			Records: []*rc0.Record{
				{
					Content:  "\"heritage=external-dns,external-dns/owner=default,external-dns/resource=ingress/default/app\"",
					Disabled: false,
				},
			},
		},
		{
			Name: "app." + zone + ".",
			Type: "A",
			TTL:  300,
			Records: []*rc0.Record{
				{
					Content:  "127.0.0.1",
					Disabled: false,
				},
			},
		},
		{
			Name: "www." + zone + ".",
			Type: "A",
			TTL:  300,
			Records: []*rc0.Record{
				{
					Content:  "127.0.0.1",
					Disabled: false,
				},
			},
		},
		{
			Name: zone + ".",
			Type: "SOA",
			TTL:  3600,
			Records: []*rc0.Record{
				{
					Content:  "sec1.rcode0.net. rcodezero-soa.ipcom.at. 2019011616 10800 3600 604800 3600",
					Disabled: false,
				},
			},
		},
		{
			Name: zone + ".",
			Type: "NS",
			TTL:  3600,
			Records: []*rc0.Record{
				{
					Content:  "sec2.rcode0.net.",
					Disabled: false,
				},
				{
					Content:  "sec1.rcode0.net.",
					Disabled: false,
				},
			},
		},
	}
}

func (m *mockRRSetService) Create(zone string, rrsetCreate []*rc0.RRSetChange) (*rc0.StatusResponse, error) {
	return &rc0.StatusResponse{Status: "ok", Message: "pass"}, nil
}

func (m *mockRRSetService) Edit(zone string, rrsetEdit []*rc0.RRSetChange) (*rc0.StatusResponse, error) {
	return &rc0.StatusResponse{Status: "ok", Message: "pass"}, nil
}

func (m *mockRRSetService) Delete(zone string, rrsetDelete []*rc0.RRSetChange) (*rc0.StatusResponse, error) {
	return &rc0.StatusResponse{Status: "ok", Message: "pass"}, nil
}

func (m *mockRRSetService) SubmitChangeSet(zone string, changeSet []*rc0.RRSetChange) (*rc0.StatusResponse, error) {
	return &rc0.StatusResponse{Status: "ok", Message: "pass"}, nil
}

func (m *mockRRSetService) EncryptTXT(key []byte, rrType *rc0.RRSetChange) {}

func (m *mockRRSetService) DecryptTXT(key []byte, rrType *rc0.RRType) {}

/* mocking ZoneManagementServiceInterface */

func (m *mockZoneManagementService) List(options *rc0.ListOptions) ([]*rc0.Zone, *rc0.Page, error) {
	if m.TestNilZonesReturned {
		return nil, nil, nil
	}

	if m.TestErrorReturned {
		return nil, nil, fmt.Errorf("operation Zone.List failed")
	}

	zones := []*rc0.Zone{
		{
			Domain: testZoneOne,
			Type:   "SLAVE",
			// "dnssec": "yes",                   @todo: add this
			// "created": "2018-04-09T09:27:31Z", @todo: add this
			LastCheck: "",
			Serial:    20180411,
			Masters: []string{
				"193.0.2.2",
				"2001:db8::2",
			},
		},
		{
			Domain: testZoneTwo,
			Type:   "MASTER",
			// "dnssec": "no",                    @todo: add this
			// "created": "2019-01-15T13:20:10Z", @todo: add this
			LastCheck: "",
			Serial:    2019011616,
			Masters: []string{
				"",
			},
		},
	}

	return zones, nil, nil
}

func (m *mockZoneManagementService) Get(zone string) (*rc0.Zone, error) { return nil, nil }
func (m *mockZoneManagementService) Create(zoneCreate *rc0.ZoneCreate) (*rc0.StatusResponse, error) {
	return nil, nil
}

func (m *mockZoneManagementService) Edit(zone string, zoneEdit *rc0.ZoneEdit) (*rc0.StatusResponse, error) {
	return nil, nil
}
func (m *mockZoneManagementService) Delete(zone string) (*rc0.StatusResponse, error) { return nil, nil }
func (m *mockZoneManagementService) Transfer(zone string) (*rc0.StatusResponse, error) {
	return nil, nil
}
