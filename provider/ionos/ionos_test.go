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

package ionos

import (
	"context"
	"fmt"
	"os"
	"sort"
	"testing"

	sdk "github.com/ionos-developer/dns-sdk-go"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockIonosDnsService struct {
	TestErrorReturned bool
}

func TestNewIonosProvider(t *testing.T) {
	_ = os.Setenv("IONOS_API_KEY", "1")
	p, err := NewIonosProvider(endpoint.NewDomainFilter([]string{"a.de."}), true)

	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	require.Equal(t, true, p.DryRun)
	require.Equal(t, true, p.DomainFilter.IsConfigured())
	require.Equal(t, false, p.DomainFilter.Match("b.de."))

	p, err = NewIonosProvider(endpoint.DomainFilter{}, false)

	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	require.Equal(t, false, p.DryRun)
	require.Equal(t, false, p.DomainFilter.IsConfigured())
	require.Equal(t, true, p.DomainFilter.Match("a.de."))

	_ = os.Unsetenv("IONOS_API_KEY")
	_, err = NewIonosProvider(endpoint.DomainFilter{}, true)

	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestRecords(t *testing.T) {
	ctx := context.Background()

	provider := &IonosProvider{Client: mockIonosDnsService{TestErrorReturned: false}}
	endpoints, err := provider.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	require.Equal(t, 5, len(endpoints))

	provider = &IonosProvider{Client: mockIonosDnsService{TestErrorReturned: true}}
	_, err = provider.Records(ctx)

	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func TestApplyChanges(t *testing.T) {
	ctx := context.Background()

	provider := &IonosProvider{Client: mockIonosDnsService{TestErrorReturned: false}}
	err := provider.ApplyChanges(ctx, mockChanges())
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	// 3 records must be deleted
	require.Equal(t, deletedRecords["b"], []string{"6"})
	sort.Strings(deletedRecords["a"])
	require.Equal(t, deletedRecords["a"], []string{"1", "2"})
	// 3 records must be created
	if !isRecordCreated("a", "a.de", sdk.A, "3.3.3.3", 2000) {
		t.Errorf("Record a.de A 3.3.3.3 not created")
	}
	if !isRecordCreated("a", "a.de", sdk.A, "4.4.4.4", 2000) {
		t.Errorf("Record a.de A 4.4.4.4 not created")
	}
	if !isRecordCreated("a", "new.a.de", sdk.CNAME, "a.de", 0) {
		t.Errorf("Record new.a.de CNAME a.de not created")
	}

	provider = &IonosProvider{Client: mockIonosDnsService{TestErrorReturned: true}}
	err = provider.ApplyChanges(ctx, nil)

	if err == nil {
		t.Errorf("expected to fail, %s", err)
	}
}

func (m mockIonosDnsService) GetZones(ctx context.Context) ([]sdk.Zone, error) {
	if m.TestErrorReturned {
		return nil, fmt.Errorf("GetZones failed")
	}

	a := sdk.NewZone()
	a.SetId("a")
	a.SetName("a.de")

	b := sdk.NewZone()
	b.SetId("b")
	b.SetName("b.de")

	return []sdk.Zone{*a, *b}, nil
}

func (m mockIonosDnsService) GetZone(ctx context.Context, zoneId string) (*sdk.CustomerZone, error) {
	if m.TestErrorReturned {
		return nil, fmt.Errorf("GetZone failed")
	}

	zoneName := zoneIdToZoneName[zoneId]
	zone := sdk.NewCustomerZone()
	zone.Id = &zoneId
	zone.Name = &zoneName
	if zoneName == "a.de" {
		zone.Records = []sdk.RecordResponse{record(1, "a.de", sdk.A, "1.1.1.1", 1000),
			record(2, "a.de", sdk.A, "2.2.2.2", 1000),
			record(3, "cname.a.de", sdk.CNAME, "cname.de", 1000),
			record(4, "aaaa.a.de", sdk.AAAA, "1::", 1000),
			record(5, "aaaa.a.de", sdk.AAAA, "2::", 2000),
		}
	} else {
		zone.Records = []sdk.RecordResponse{record(6, "b.de", sdk.A, "5.5.5.5", 1000)}
	}

	return zone, nil
}

func (m mockIonosDnsService) CreateRecords(ctx context.Context, zoneId string, records []sdk.Record) error {
	createdRecords[zoneId] = append(createdRecords[zoneId], records...)
	return nil
}

func (m mockIonosDnsService) DeleteRecord(ctx context.Context, zoneId string, recordId string) error {
	deletedRecords[zoneId] = append(deletedRecords[zoneId], recordId)
	return nil
}

func record(id int, name string, recordType sdk.RecordTypes, content string, ttl int32) sdk.RecordResponse {
	r := sdk.NewRecordResponse()
	idStr := fmt.Sprint(id)
	r.Id = &idStr
	r.Name = &name
	r.Type = &recordType
	r.Content = &content
	r.Ttl = &ttl
	return *r
}

func mockChanges() *plan.Changes {
	changes := &plan.Changes{}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "new.a.de", Targets: endpoint.Targets{"a.de"}, RecordType: "CNAME"},
	}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "b.de", RecordType: "A", Targets: endpoint.Targets{"5.5.5.5"}}}
	changes.UpdateOld = []*endpoint.Endpoint{{DNSName: "a.de", RecordType: "A", Targets: endpoint.Targets{"1.1.1.1", "2.2.2.2"}, RecordTTL: 1000}}
	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "a.de", RecordType: "A", Targets: endpoint.Targets{"3.3.3.3", "4.4.4.4"}, RecordTTL: 2000}}

	return changes
}

var zoneIdToZoneName = map[string]string{
	"a": "a.de",
	"b": "b.de",
}

var createdRecords = map[string][]sdk.Record{"a": {}, "b": {}}
var deletedRecords = map[string][]string{"a": {}, "b": {}}

func isRecordCreated(zoneId string, name string, recordType sdk.RecordTypes, content string, ttl int32) bool {
	for _, record := range createdRecords[zoneId] {
		if *record.Name == name && *record.Type == recordType && *record.Content == content && (ttl == 0 || *record.Ttl == ttl) {
			return true
		}
	}

	return false
}
