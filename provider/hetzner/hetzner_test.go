/*
Copyright 2020 The Kubernetes Authors.
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

package hetzner

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"

	hclouddns "git.blindage.org/21h/hcloud-dns"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockHCloudClientAdapter interface {
	GetZone(ID string) (hclouddns.HCloudAnswerGetZone, error)
	GetZones(params hclouddns.HCloudGetZonesParams) (hclouddns.HCloudAnswerGetZones, error)
	UpdateZone(zone hclouddns.HCloudZone) (hclouddns.HCloudAnswerGetZone, error)
	DeleteZone(ID string) (hclouddns.HCloudAnswerDeleteZone, error)
	CreateZone(zone hclouddns.HCloudZone) (hclouddns.HCloudAnswerGetZone, error)
	ImportZoneString(zoneID string, zonePlainText string) (hclouddns.HCloudAnswerGetZone, error)
	ExportZoneToString(zoneID string) (hclouddns.HCloudAnswerGetZonePlainText, error)
	ValidateZoneString(zonePlainText string) (hclouddns.HCloudAnswerZoneValidate, error)
	GetRecord(ID string) (hclouddns.HCloudAnswerGetRecord, error)
	GetRecords(params hclouddns.HCloudGetRecordsParams) (hclouddns.HCloudAnswerGetRecords, error)
	UpdateRecord(record hclouddns.HCloudRecord) (hclouddns.HCloudAnswerGetRecord, error)
	DeleteRecord(ID string) (hclouddns.HCloudAnswerDeleteRecord, error)
	CreateRecord(record hclouddns.HCloudRecord) (hclouddns.HCloudAnswerGetRecord, error)
	CreateRecordBulk(record []hclouddns.HCloudRecord) (hclouddns.HCloudAnswerCreateRecords, error)
	UpdateRecordBulk(record []hclouddns.HCloudRecord) (hclouddns.HCloudAnswerUpdateRecords, error)
}

type mockHCloudClient struct {
	Token string `yaml:"token"`
}

// New instance
func mockHCloudNew(t string) mockHCloudClientAdapter {
	return &mockHCloudClient{
		Token: t,
	}
}

// Mock all methods

func (m *mockHCloudClient) GetZone(ID string) (hclouddns.HCloudAnswerGetZone, error) {
	return hclouddns.HCloudAnswerGetZone{}, nil
}

func (m *mockHCloudClient) GetZones(params hclouddns.HCloudGetZonesParams) (hclouddns.HCloudAnswerGetZones, error) {
	return hclouddns.HCloudAnswerGetZones{
		Zones: []hclouddns.HCloudZone{
			{
				ID:           "HetznerZoneID",
				Name:         "blindage.org",
				TTL:          666,
				RecordsCount: 1,
			},
		},
	}, nil
}

// zones
func (m *mockHCloudClient) UpdateZone(zone hclouddns.HCloudZone) (hclouddns.HCloudAnswerGetZone, error) {
	return hclouddns.HCloudAnswerGetZone{}, nil
}
func (m *mockHCloudClient) DeleteZone(ID string) (hclouddns.HCloudAnswerDeleteZone, error) {
	return hclouddns.HCloudAnswerDeleteZone{}, nil
}
func (m *mockHCloudClient) CreateZone(zone hclouddns.HCloudZone) (hclouddns.HCloudAnswerGetZone, error) {
	return hclouddns.HCloudAnswerGetZone{}, nil
}
func (m *mockHCloudClient) ImportZoneString(zoneID string, zonePlainText string) (hclouddns.HCloudAnswerGetZone, error) {
	return hclouddns.HCloudAnswerGetZone{}, nil
}
func (m *mockHCloudClient) ExportZoneToString(zoneID string) (hclouddns.HCloudAnswerGetZonePlainText, error) {
	return hclouddns.HCloudAnswerGetZonePlainText{}, nil
}
func (m *mockHCloudClient) ValidateZoneString(zonePlainText string) (hclouddns.HCloudAnswerZoneValidate, error) {
	return hclouddns.HCloudAnswerZoneValidate{}, nil
}

// records

func (m *mockHCloudClient) GetRecord(ID string) (hclouddns.HCloudAnswerGetRecord, error) {
	return hclouddns.HCloudAnswerGetRecord{}, nil
}

func (m *mockHCloudClient) GetRecords(params hclouddns.HCloudGetRecordsParams) (hclouddns.HCloudAnswerGetRecords, error) {
	return hclouddns.HCloudAnswerGetRecords{
		Records: []hclouddns.HCloudRecord{
			{
				RecordType: hclouddns.RecordType("A"),
				ID:         "ATypeRecordID",
				ZoneID:     "HetznerZoneID",
				Name:       "@",
				Value:      "127.0.0.1",
				TTL:        666,
			},
		},
	}, nil
}
func (m *mockHCloudClient) UpdateRecord(record hclouddns.HCloudRecord) (hclouddns.HCloudAnswerGetRecord, error) {
	return hclouddns.HCloudAnswerGetRecord{}, nil
}
func (m *mockHCloudClient) DeleteRecord(ID string) (hclouddns.HCloudAnswerDeleteRecord, error) {
	return hclouddns.HCloudAnswerDeleteRecord{}, nil
}
func (m *mockHCloudClient) CreateRecord(record hclouddns.HCloudRecord) (hclouddns.HCloudAnswerGetRecord, error) {
	return hclouddns.HCloudAnswerGetRecord{}, nil
}
func (m *mockHCloudClient) CreateRecordBulk(record []hclouddns.HCloudRecord) (hclouddns.HCloudAnswerCreateRecords, error) {
	return hclouddns.HCloudAnswerCreateRecords{}, nil
}
func (m *mockHCloudClient) UpdateRecordBulk(record []hclouddns.HCloudRecord) (hclouddns.HCloudAnswerUpdateRecords, error) {
	return hclouddns.HCloudAnswerUpdateRecords{}, nil
}

func TestNewHetznerProvider(t *testing.T) {
	_ = os.Setenv("HETZNER_TOKEN", "myHetznerToken")
	_, err := NewHetznerProvider(context.Background(), endpoint.NewDomainFilter([]string{"blindage.org"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}

	_ = os.Unsetenv("HETZNER_TOKEN")
	_, err = NewHetznerProvider(context.Background(), endpoint.NewDomainFilter([]string{"blindage.org"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestHetznerProvider_TestData(t *testing.T) {

	mockedClient := mockHCloudNew("myHetznerToken")

	// Check test zone data is ok
	expectedZonesAnswer := hclouddns.HCloudAnswerGetZones{
		Zones: []hclouddns.HCloudZone{
			{
				ID:           "HetznerZoneID",
				Name:         "blindage.org",
				TTL:          666,
				RecordsCount: 1,
			},
		},
	}

	testingZonesAnswer, err := mockedClient.GetZones(hclouddns.HCloudGetZonesParams{})
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	if !reflect.DeepEqual(expectedZonesAnswer, testingZonesAnswer) {
		t.Errorf("should be equal, %s", err)
	}

	// Check test record data is ok
	expectedRecordsAnswer := hclouddns.HCloudAnswerGetRecords{
		Records: []hclouddns.HCloudRecord{
			{
				RecordType: hclouddns.RecordType("A"),
				ID:         "ATypeRecordID",
				ZoneID:     "HetznerZoneID",
				Name:       "@",
				Value:      "127.0.0.1",
				TTL:        666,
			},
		},
	}

	testingRecordsAnswer, err := mockedClient.GetRecords(hclouddns.HCloudGetRecordsParams{})
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	if !reflect.DeepEqual(expectedRecordsAnswer, testingRecordsAnswer) {
		t.Errorf("should be equal, %s", err)
	}

}

func TestHetznerProvider_Records(t *testing.T) {

	mockedClient := mockHCloudNew("myHetznerToken")

	mockedProvider := &HetznerProvider{
		Client: mockedClient,
	}

	// Now check Records function of provider, if ZoneID equal "blindage.org" must be returned
	endpoints, err := mockedProvider.Records(context.Background())
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	fmt.Printf("%+v\n", endpoints[0].DNSName)
	assert.Equal(t, "blindage.org", endpoints[0].DNSName)
}

func TestHetznerProvider_ApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	mockedClient := mockHCloudNew("myHetznerToken")
	mockedProvider := &HetznerProvider{
		Client: mockedClient,
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "test.org", Targets: endpoint.Targets{"target"}},
		{DNSName: "test.test.org", Targets: endpoint.Targets{"target"}, RecordTTL: 666},
	}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test.test.org", Targets: endpoint.Targets{"target-new"}, RecordType: "A", RecordTTL: 777}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "test.test.org", Targets: endpoint.Targets{"target"}, RecordType: "A"}}

	err := mockedProvider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}
