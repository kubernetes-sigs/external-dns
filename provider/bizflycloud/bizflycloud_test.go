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

package bizflycloud

import (
	"context"
	"errors"
	"os"
	"testing"

	bizflycloud "github.com/bizflycloud/gobizfly"
	"github.com/maxatome/go-testdeep/td"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type MockAction struct {
	Name       string
	ZoneId     string
	RecordData bizflycloud.Record
}

type mockBizflyCloudClient struct {
	Zones   map[string]string
	Records map[string]bizflycloud.Record
	Actions []MockAction
}

var ExampleRecrods = []bizflycloud.Record{
	{
		ID:     "R001",
		ZoneID: "Z001",
		Name:   "foobar",
		Type:   endpoint.RecordTypeA,
		TTL:    120,
		Data:   makeRecordData([]string{"1.2.3.4", "3.4.5.6"}),
	},
	{
		ID:     "R002",
		ZoneID: "Z001",
		Name:   "foo",
		Type:   endpoint.RecordTypeA,
		TTL:    120,
		Data:   makeRecordData([]string{"3.4.5.6"}),
	},
	{
		ID:     "R003",
		ZoneID: "Z002",
		Name:   "bar",
		Type:   endpoint.RecordTypeA,
		TTL:    1,
		Data:   makeRecordData([]string{"2.3.4.5"}),
	},
}

func makeRecordData(listData []string) []interface{} {
	recordData := make([]interface{}, 0)
	for _, data := range listData {
		recordData = append(recordData, data)
	}
	return recordData
}

func NewMockBizflyCloudClient() *mockBizflyCloudClient {
	return &mockBizflyCloudClient{
		Zones: map[string]string{
			"Z001": "bar.com",
			"Z002": "foo.com",
		},
		Records: map[string]bizflycloud.Record{
			"R001": {},
			"R002": {},
		},
	}
}

func NewMockBizflyCloudClientWithRecords(records []bizflycloud.Record) *mockBizflyCloudClient {
	m := NewMockBizflyCloudClient()

	for _, record := range records {
		m.Records[record.ID] = record
	}

	return m
}

func getDNSRecordFromRecordParams(crpl interface{}, zoneID string, recordID string) bizflycloud.Record {
	switch params := crpl.(type) {
	case bizflycloud.CreateNormalRecordPayload:
		return bizflycloud.Record{
			Name:   params.Name,
			TTL:    params.TTL,
			Type:   params.Type,
			ZoneID: zoneID,
			Data:   makeRecordData(params.Data),
		}
	case bizflycloud.UpdateNormalRecordPayload:
		return bizflycloud.Record{
			ID:     recordID,
			Name:   params.Name,
			TTL:    params.TTL,
			Type:   params.Type,
			ZoneID: zoneID,
			Data:   makeRecordData(params.Data),
		}
	default:
		return bizflycloud.Record{}
	}
}

func (m *mockBizflyCloudClient) CreateRecord(ctx context.Context, zoneID string, crpl interface{}) (*bizflycloud.Record, error) {
	recordData := getDNSRecordFromRecordParams(crpl, zoneID, "")
	m.Actions = append(m.Actions, MockAction{
		Name:       "Create",
		ZoneId:     zoneID,
		RecordData: recordData,
	})
	m.Records["R003"] = recordData
	return &bizflycloud.Record{}, nil
}

func (m *mockBizflyCloudClient) UpdateRecord(ctx context.Context, recordID string, urpl interface{}) (*bizflycloud.Record, error) {
	if record, ok := m.Records[recordID]; ok {
		zoneID := record.ZoneID
		recordData := getDNSRecordFromRecordParams(urpl, zoneID, recordID)
		m.Actions = append(m.Actions, MockAction{
			Name:       "Update",
			ZoneId:     zoneID,
			RecordData: recordData,
		})
		return &bizflycloud.Record{}, nil
	}
	return nil, errors.New("Unknown zoneID: " + recordID)
}

func (m *mockBizflyCloudClient) DeleteRecord(ctx context.Context, recordID string) error {
	if record, ok := m.Records[recordID]; ok {
		zoneID := record.ZoneID
		m.Actions = append(m.Actions, MockAction{
			Name:   "Delete",
			ZoneId: zoneID,
			RecordData: bizflycloud.Record{
				ID: record.ID,
			},
		})
		delete(m.Records, recordID)
	}
	return nil
}

func (m *mockBizflyCloudClient) ListZones(ctx context.Context, opts *bizflycloud.ListOptions) (*bizflycloud.ListZoneResp, error) {
	result := bizflycloud.ListZoneResp{}

	for zoneID, zoneName := range m.Zones {
		result.Zones = append(result.Zones, bizflycloud.Zone{
			ID:   zoneID,
			Name: zoneName,
		})
	}

	return &result, nil
}

func (m *mockBizflyCloudClient) GetZone(ctx context.Context, zoneID string) (*bizflycloud.ExtendedZone, error) {
	recordSet := []bizflycloud.Record{}
	for _, record := range m.Records {
		if record.ZoneID == zoneID {
			recordSet = append(recordSet, record)
		}
	}
	for id, zoneName := range m.Zones {
		if zoneID == id {
			return &bizflycloud.ExtendedZone{
				Zone: bizflycloud.Zone{
					ID:   zoneID,
					Name: zoneName,
				},
				RecordsSet: recordSet,
			}, nil
		}
	}

	return &bizflycloud.ExtendedZone{}, errors.New("Unknown zoneID: " + zoneID)
}

func AssertActions(t *testing.T, provider *BizflyCloudProvider, endpoints []*endpoint.Endpoint, actions []MockAction, managedRecords []string, args ...interface{}) {
	t.Helper()

	var client *mockBizflyCloudClient

	if provider.Client == nil {
		client = NewMockBizflyCloudClient()
		provider.Client = client
	} else {
		client = provider.Client.(*mockBizflyCloudClient)
	}

	ctx := context.Background()

	records, err := provider.Records(ctx)
	if err != nil {
		t.Fatalf("cannot fetch records, %s", err)
	}

	plan := &plan.Plan{
		Current:        records,
		Desired:        endpoints,
		DomainFilter:   endpoint.NewDomainFilter([]string{"bar.com"}),
		ManagedRecords: managedRecords,
	}

	changes := plan.Calculate().Changes

	// Records other than A, CNAME and NS are not supported by planner, just create them
	for _, endpoint := range endpoints {
		if endpoint.RecordType != "A" && endpoint.RecordType != "CNAME" && endpoint.RecordType != "NS" {
			changes.Create = append(changes.Create, endpoint)
		}
	}

	err = provider.ApplyChanges(context.Background(), changes)

	if err != nil {
		t.Fatalf("cannot apply changes, %s", err)
	}

	td.Cmp(t, client.Actions, actions, args...)
}

func TestBizflycloudA(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "A",
			DNSName:    "bar.com",
			Targets:    endpoint.Targets{"127.0.0.1", "127.0.0.2"},
		},
	}

	AssertActions(t, &BizflyCloudProvider{}, endpoints, []MockAction{
		{
			Name:   "Create",
			ZoneId: "Z001",
			RecordData: bizflycloud.Record{
				Type:   "A",
				Name:   "bar.com",
				TTL:    60,
				ZoneID: "Z001",
				Data:   makeRecordData([]string{"127.0.0.1", "127.0.0.2"}),
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestBizflycloudCname(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "CNAME",
			DNSName:    "cname.bar.com",
			Targets:    endpoint.Targets{"google.com"},
		},
	}

	AssertActions(t, &BizflyCloudProvider{}, endpoints, []MockAction{
		{
			Name:   "Create",
			ZoneId: "Z001",
			RecordData: bizflycloud.Record{
				Type:   "CNAME",
				Name:   "cname.bar.com",
				TTL:    60,
				ZoneID: "Z001",
				Data:   makeRecordData([]string{"google.com"}),
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestBizflycloudCustomTTL(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "A",
			DNSName:    "ttl.bar.com",
			Targets:    endpoint.Targets{"127.0.0.1"},
			RecordTTL:  120,
		},
	}

	AssertActions(t, &BizflyCloudProvider{}, endpoints, []MockAction{
		{
			Name:   "Create",
			ZoneId: "Z001",
			RecordData: bizflycloud.Record{
				Type:   "A",
				Name:   "ttl.bar.com",
				TTL:    120,
				ZoneID: "Z001",
				Data:   makeRecordData([]string{"127.0.0.1"}),
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestBizflycloudZones(t *testing.T) {
	provider := &BizflyCloudProvider{
		Client:       NewMockBizflyCloudClient(),
		domainFilter: endpoint.NewDomainFilter([]string{"bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}

	zones, err := provider.Zones(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(zones))
	assert.Equal(t, "bar.com", zones[0].Name)
}

func TestBizflyCloudZonesWithIDFilter(t *testing.T) {
	client := NewMockBizflyCloudClient()
	provider := &BizflyCloudProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"bar.com", "foo.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{"Z001"}),
	}

	zones, err := provider.Zones(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// foo.com should *not* be returned as it doesn't match ZoneID filter
	assert.Equal(t, 1, len(zones))
	assert.Equal(t, "bar.com", zones[0].Name)
}

func TestBizflycloudRecords(t *testing.T) {
	client := NewMockBizflyCloudClientWithRecords(ExampleRecrods)

	// Set DNSRecordsPerPage to 1 test the pagination behaviour
	provider := &BizflyCloudProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"bar.com"}),
	}
	ctx := context.Background()

	records, err := provider.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	assert.Equal(t, 2, len(records))
}

func TestBizflycloudProvider(t *testing.T) {
	_ = os.Setenv("BFC_APP_CREDENTIAL_ID", "34d3c157d2de499eb6a39ab85945e7bd")
	_ = os.Setenv("BFC_APP_CREDENTIAL_SECRET", "DmfebiHV6b5aFJ0AD9g9z3m9rGqgFMz95ClAH-wb5WNZH3QYZnm3NzTwa_TtST0QqY4H-AywMHECzsmg6mfISA")
	ctx := context.Background()
	_, err := NewBizflyCloudProvider(
		ctx,
		endpoint.NewDomainFilter([]string{"bar.com"}),
		provider.NewZoneIDFilter([]string{""}),
		false,
		"HN",
		1000)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	_ = os.Unsetenv("BFC_APP_CREDENTIAL_ID")
	_ = os.Unsetenv("BFC_APP_CREDENTIAL_SECRET")
	_, err = NewBizflyCloudProvider(
		ctx,
		endpoint.NewDomainFilter([]string{"bar.com"}),
		provider.NewZoneIDFilter([]string{""}),
		false,
		"HN",
		1000)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestBizflycloudApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	client := NewMockBizflyCloudClientWithRecords(ExampleRecrods)

	provider := &BizflyCloudProvider{
		Client: client,
	}

	changes.Create = []*endpoint.Endpoint{{
		DNSName:    "new.bar.com",
		RecordTTL:  60,
		RecordType: "A",
		Targets:    endpoint.Targets{"target"},
	}, {
		DNSName:    "new.ext-dns-test.unrelated.to",
		RecordTTL:  60,
		RecordType: "A",
		Targets:    endpoint.Targets{"target"},
	}}
	changes.Delete = []*endpoint.Endpoint{{
		DNSName:    "foobar.bar.com",
		RecordTTL:  60,
		RecordType: "A",
		Targets:    endpoint.Targets{"target"},
	}}
	changes.UpdateOld = []*endpoint.Endpoint{{
		DNSName:    "foobar.bar.com",
		RecordTTL:  60,
		RecordType: "A",
		Targets:    endpoint.Targets{"target-old"},
	}}
	changes.UpdateNew = []*endpoint.Endpoint{{
		DNSName:    "foobar.bar.com",
		RecordTTL:  60,
		RecordType: "A",
		Targets:    endpoint.Targets{"target-new"},
	}}
	err := provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	td.Cmp(t, client.Actions, []MockAction{
		{
			Name:   "Create",
			ZoneId: "Z001",
			RecordData: bizflycloud.Record{
				Name:   "new.bar.com",
				ZoneID: "Z001",
				Type:   "A",
				TTL:    60,
				Data:   makeRecordData([]string{"target"}),
			},
		},
		{
			Name:   "Update",
			ZoneId: "Z001",
			RecordData: bizflycloud.Record{
				Name:   "foobar.bar.com",
				ZoneID: "Z001",
				Type:   "A",
				TTL:    60,
				ID:     "R001",
				Data:   makeRecordData([]string{"target-new"}),
			},
		},
		{
			Name:   "Delete",
			ZoneId: "Z001",
			RecordData: bizflycloud.Record{
				ID: "R001",
			},
		},
	})

	// empty changes
	changes.Create = []*endpoint.Endpoint{}
	changes.Delete = []*endpoint.Endpoint{}
	changes.UpdateOld = []*endpoint.Endpoint{}
	changes.UpdateNew = []*endpoint.Endpoint{}

	err = provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestBizflycloudGetRecordID(t *testing.T) {
	p := &BizflyCloudProvider{}
	records := []bizflycloud.Record{
		{
			ID:     "1",
			Name:   "cname",
			Type:   endpoint.RecordTypeCNAME,
			ZoneID: "Z001",
			Data:   makeRecordData([]string{"foo.bar.com"}),
		},
		{
			ID:     "2",
			Name:   "@",
			Type:   endpoint.RecordTypeA,
			ZoneID: "Z001",
			Data:   makeRecordData([]string{"1.2.3.4"}),
		},
		{
			ID:     "3",
			Name:   "foo",
			Type:   endpoint.RecordTypeA,
			ZoneID: "Z001",
			Data:   makeRecordData([]string{"1.2.3.4"}),
		},
	}
	zone := bizflycloud.ExtendedZone{
		Zone: bizflycloud.Zone{
			Name: "bar.com",
		},
		RecordsSet: records,
	}

	assert.Equal(t, "", p.getRecordID(&zone, NormalRecord{
		Name: "bar.com",
		Type: endpoint.RecordTypeCNAME,
	}))

	assert.Equal(t, "", p.getRecordID(&zone, NormalRecord{
		Name: "cname",
		Type: endpoint.RecordTypeA,
	}))

	assert.Equal(t, "1", p.getRecordID(&zone, NormalRecord{
		Name: "cname.bar.com",
		Type: endpoint.RecordTypeCNAME,
	}))
	assert.Equal(t, "2", p.getRecordID(&zone, NormalRecord{
		Name: "bar.com",
		Type: endpoint.RecordTypeA,
	}))
	assert.Equal(t, "3", p.getRecordID(&zone, NormalRecord{
		Name: "foo.bar.com",
		Type: endpoint.RecordTypeA,
	}))
}

func TestBizflycloudComplexUpdate(t *testing.T) {
	client := NewMockBizflyCloudClientWithRecords(ExampleRecrods)

	provider := &BizflyCloudProvider{
		Client: client,
	}
	ctx := context.Background()

	records, err := provider.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	plan := &plan.Plan{
		Current: records,
		Desired: []*endpoint.Endpoint{
			{
				DNSName:    "foobar.bar.com",
				Targets:    endpoint.Targets{"1.2.3.4", "2.3.4.5"},
				RecordType: endpoint.RecordTypeA,
				RecordTTL:  endpoint.TTL(defaultBizflyCloudRecordTTL),
				Labels:     endpoint.Labels{},
			},
		},
		DomainFilter:   endpoint.NewDomainFilter([]string{"bar.com"}),
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	}

	planned := plan.Calculate()

	err = provider.ApplyChanges(context.Background(), planned.Changes)

	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	td.CmpDeeply(t, client.Actions, []MockAction{
		{
			Name:   "Update",
			ZoneId: "Z001",
			RecordData: bizflycloud.Record{
				ID:     "R001",
				Name:   "foobar.bar.com",
				Type:   "A",
				ZoneID: "Z001",
				TTL:    60,
				Data:   makeRecordData([]string{"1.2.3.4", "2.3.4.5"}),
			},
		},
		{
			Name:   "Delete",
			ZoneId: "Z001",
			RecordData: bizflycloud.Record{
				ID: "R002",
			},
		},
	})
}
