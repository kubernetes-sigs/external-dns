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

package scaleway

import (
	"context"
	"os"
	"reflect"
	"testing"

	domain "github.com/scaleway/scaleway-sdk-go/api/domain/v2beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockScalewayDomain struct {
	*domain.API
}

func (m *mockScalewayDomain) ListDNSZones(req *domain.ListDNSZonesRequest, opts ...scw.RequestOption) (*domain.ListDNSZonesResponse, error) {
	return &domain.ListDNSZonesResponse{
		DNSZones: []*domain.DNSZone{
			{
				Domain:    "example.com",
				Subdomain: "",
			},
			{
				Domain:    "example.com",
				Subdomain: "test",
			},
			{
				Domain:    "dummy.me",
				Subdomain: "",
			},
			{
				Domain:    "dummy.me",
				Subdomain: "test",
			},
		},
	}, nil
}

func (m *mockScalewayDomain) ListDNSZoneRecords(req *domain.ListDNSZoneRecordsRequest, opts ...scw.RequestOption) (*domain.ListDNSZoneRecordsResponse, error) {
	records := []*domain.Record{}
	if req.DNSZone == "example.com" {
		records = []*domain.Record{
			{
				Data:     "1.1.1.1",
				Name:     "one",
				TTL:      300,
				Priority: 0,
				Type:     domain.RecordTypeA,
			},
			{
				Data:     "1.1.1.2",
				Name:     "two",
				TTL:      300,
				Priority: 0,
				Type:     domain.RecordTypeA,
			},
			{
				Data:     "1.1.1.3",
				Name:     "two",
				TTL:      300,
				Priority: 0,
				Type:     domain.RecordTypeA,
			},
		}
	} else if req.DNSZone == "test.example.com" {
		records = []*domain.Record{
			{
				Data:     "1.1.1.1",
				Name:     "",
				TTL:      300,
				Priority: 0,
				Type:     domain.RecordTypeA,
			},
			{
				Data:     "test.example.com.",
				Name:     "two",
				TTL:      600,
				Priority: 30,
				Type:     domain.RecordTypeCNAME,
			},
		}
	}
	return &domain.ListDNSZoneRecordsResponse{
		Records: records,
	}, nil
}

func (m *mockScalewayDomain) UpdateDNSZoneRecords(req *domain.UpdateDNSZoneRecordsRequest, opts ...scw.RequestOption) (*domain.UpdateDNSZoneRecordsResponse, error) {
	return nil, nil
}

func TestScalewayProvider_NewScalewayProvider(t *testing.T) {
	profile := `profiles:
  foo:
    access_key: SCWXXXXXXXXXXXXXXXXX
    secret_key: 11111111-1111-1111-1111-111111111111
`
	tmpDir := t.TempDir()
	err := os.WriteFile(tmpDir+"/config.yaml", []byte(profile), 0600)
	if err != nil {
		t.Errorf("failed : %s", err)
	}
	_ = os.Setenv(scw.ScwActiveProfileEnv, "foo")
	_ = os.Setenv(scw.ScwConfigPathEnv, tmpDir+"/config.yaml")
	_, err = NewScalewayProvider(context.TODO(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}

	_ = os.Setenv(scw.ScwAccessKeyEnv, "SCWXXXXXXXXXXXXXXXXX")
	_ = os.Setenv(scw.ScwSecretKeyEnv, "11111111-1111-1111-1111-111111111111")
	_, err = NewScalewayProvider(context.TODO(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}

	_ = os.Unsetenv(scw.ScwSecretKeyEnv)
	_, err = NewScalewayProvider(context.TODO(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}

	_ = os.Setenv(scw.ScwSecretKeyEnv, "dummy")
	_, err = NewScalewayProvider(context.TODO(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}

	_ = os.Unsetenv(scw.ScwAccessKeyEnv)
	_ = os.Setenv(scw.ScwSecretKeyEnv, "11111111-1111-1111-1111-111111111111")
	_, err = NewScalewayProvider(context.TODO(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}

	_ = os.Setenv(scw.ScwAccessKeyEnv, "dummy")
	_, err = NewScalewayProvider(context.TODO(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestScalewayProvider_OptionnalConfigFile(t *testing.T) {
	_ = os.Setenv(scw.ScwAccessKeyEnv, "SCWXXXXXXXXXXXXXXXXX")
	_ = os.Setenv(scw.ScwSecretKeyEnv, "11111111-1111-1111-1111-111111111111")

	_, err := NewScalewayProvider(context.TODO(), endpoint.NewDomainFilter([]string{"example.com"}), true)
	assert.NoError(t, err)
}

func TestScalewayProvider_AdjustEndpoints(t *testing.T) {
	provider := &ScalewayProvider{}

	before := []*endpoint.Endpoint{
		{
			DNSName:    "one.example.com",
			RecordTTL:  300,
			RecordType: "A",
			Targets:    []string{"1.1.1.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
		{
			DNSName:    "two.example.com",
			RecordTTL:  0,
			RecordType: "A",
			Targets:    []string{"1.1.1.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "10",
				},
			},
		},
		{
			DNSName:          "three.example.com",
			RecordTTL:        600,
			RecordType:       "A",
			Targets:          []string{"1.1.1.1"},
			ProviderSpecific: endpoint.ProviderSpecific{},
		},
	}

	expected := []*endpoint.Endpoint{
		{
			DNSName:    "one.example.com",
			RecordTTL:  300,
			RecordType: "A",
			Targets:    []string{"1.1.1.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
		{
			DNSName:    "two.example.com",
			RecordTTL:  300,
			RecordType: "A",
			Targets:    []string{"1.1.1.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "10",
				},
			},
		},
		{
			DNSName:    "three.example.com",
			RecordTTL:  600,
			RecordType: "A",
			Targets:    []string{"1.1.1.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
	}

	after, err := provider.AdjustEndpoints(before)
	assert.NoError(t, err)
	for i := range after {
		if !checkRecordEquality(after[i], expected[i]) {
			t.Errorf("got record %s instead of %s", after[i], expected[i])
		}
	}
}

func TestScalewayProvider_Zones(t *testing.T) {
	mocked := mockScalewayDomain{nil}
	provider := &ScalewayProvider{
		domainAPI:    &mocked,
		domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	expected := []*domain.DNSZone{
		{
			Domain:    "example.com",
			Subdomain: "",
		},
		{
			Domain:    "example.com",
			Subdomain: "test",
		},
	}
	zones, err := provider.Zones(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, zones, len(expected))
	for i, zone := range zones {
		assert.Equal(t, expected[i], zone)
	}
}

func TestScalewayProvider_Records(t *testing.T) {
	mocked := mockScalewayDomain{nil}
	provider := &ScalewayProvider{
		domainAPI:    &mocked,
		domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	expected := []*endpoint.Endpoint{
		{
			DNSName:    "one.example.com",
			RecordTTL:  300,
			RecordType: "A",
			Targets:    []string{"1.1.1.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
		{
			DNSName:    "two.example.com",
			RecordTTL:  300,
			RecordType: "A",
			Targets:    []string{"1.1.1.2", "1.1.1.3"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
		{
			DNSName:    "test.example.com",
			RecordTTL:  300,
			RecordType: "A",
			Targets:    []string{"1.1.1.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
		{
			DNSName:    "two.test.example.com",
			RecordTTL:  600,
			RecordType: "CNAME",
			Targets:    []string{"test.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "30",
				},
			},
		},
	}

	records, err := provider.Records(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	require.Len(t, records, len(expected))
	for _, record := range records {
		found := false
		for _, expectedRecord := range expected {
			if checkRecordEquality(record, expectedRecord) {
				found = true
			}
		}
		assert.Equal(t, true, found)
	}
}

// this test is really ugly since we are working on maps, so array are randomly sorted
// feel free to modify if you have a better idea
func TestScalewayProvider_generateApplyRequests(t *testing.T) {
	mocked := mockScalewayDomain{nil}
	provider := &ScalewayProvider{
		domainAPI:    &mocked,
		domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	expected := []*domain.UpdateDNSZoneRecordsRequest{
		{
			DNSZone: "example.com",
			Changes: []*domain.RecordChange{
				{
					Add: &domain.RecordChangeAdd{
						Records: []*domain.Record{
							{
								Data:     "1.1.1.1",
								Name:     "",
								TTL:      300,
								Type:     domain.RecordTypeA,
								Priority: 0,
							},
							{
								Data:     "1.1.1.2",
								Name:     "",
								TTL:      300,
								Type:     domain.RecordTypeA,
								Priority: 0,
							},
							{
								Data:     "2.2.2.2",
								Name:     "me",
								TTL:      600,
								Type:     domain.RecordTypeA,
								Priority: 30,
							},
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: scw.StringPtr("3.3.3.3"),
							Name: "me",
							Type: domain.RecordTypeA,
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: scw.StringPtr("1.1.1.1"),
							Name: "here",
							Type: domain.RecordTypeA,
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: scw.StringPtr("1.1.1.2"),
							Name: "here",
							Type: domain.RecordTypeA,
						},
					},
				},
			},
		},
		{
			DNSZone: "test.example.com",
			Changes: []*domain.RecordChange{
				{
					Add: &domain.RecordChangeAdd{
						Records: []*domain.Record{
							{
								Data:     "example.com.",
								Name:     "",
								TTL:      600,
								Type:     domain.RecordTypeCNAME,
								Priority: 20,
							},
							{
								Data:     "1.2.3.4",
								Name:     "my",
								TTL:      300,
								Type:     domain.RecordTypeA,
								Priority: 0,
							},
							{
								Data:     "5.6.7.8",
								Name:     "my",
								TTL:      300,
								Type:     domain.RecordTypeA,
								Priority: 0,
							},
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: scw.StringPtr("1.1.1.1"),
							Name: "here.is.my",
							Type: domain.RecordTypeA,
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: scw.StringPtr("4.4.4.4"),
							Name: "my",
							Type: domain.RecordTypeA,
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: scw.StringPtr("5.5.5.5"),
							Name: "my",
							Type: domain.RecordTypeA,
						},
					},
				},
			},
		},
	}

	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "example.com",
				RecordType: "A",
				Targets:    []string{"1.1.1.1", "1.1.1.2"},
			},
			{
				DNSName:    "test.example.com",
				RecordType: "CNAME",
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  scalewayPriorityKey,
						Value: "20",
					},
				},
				RecordTTL: 600,
				Targets:   []string{"example.com"},
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "here.example.com",
				RecordType: "A",
				Targets:    []string{"1.1.1.1", "1.1.1.2"},
			},
			{
				DNSName:    "here.is.my.test.example.com",
				RecordType: "A",
				Targets:    []string{"1.1.1.1"},
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName: "me.example.com",
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  scalewayPriorityKey,
						Value: "30",
					},
				},
				RecordType: "A",
				RecordTTL:  600,
				Targets:    []string{"2.2.2.2"},
			},
			{
				DNSName:    "my.test.example.com",
				RecordType: "A",
				Targets:    []string{"1.2.3.4", "5.6.7.8"},
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName: "me.example.com",
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  scalewayPriorityKey,
						Value: "1234",
					},
				},
				RecordType: "A",
				Targets:    []string{"3.3.3.3"},
			},
			{
				DNSName:    "my.test.example.com",
				RecordType: "A",
				Targets:    []string{"4.4.4.4", "5.5.5.5"},
			},
		},
	}

	requests, err := provider.generateApplyRequests(context.TODO(), changes)
	if err != nil {
		t.Fatal(err)
	}

	require.Len(t, requests, len(expected))
	total := int(len(expected))
	for _, req := range requests {
		for _, exp := range expected {
			if checkScalewayReqChanges(req, exp) {
				total--
			}
		}
	}
	assert.Equal(t, 0, total)
}

func checkRecordEquality(record1, record2 *endpoint.Endpoint) bool {
	return record1.Targets.Same(record2.Targets) &&
		record1.DNSName == record2.DNSName &&
		record1.RecordTTL == record2.RecordTTL &&
		record1.RecordType == record2.RecordType &&
		reflect.DeepEqual(record1.ProviderSpecific, record2.ProviderSpecific)
}

func checkScalewayReqChanges(r1, r2 *domain.UpdateDNSZoneRecordsRequest) bool {
	if r1.DNSZone != r2.DNSZone {
		return false
	}
	if len(r1.Changes) != len(r2.Changes) {
		return false
	}
	total := int(len(r1.Changes))
	for _, c1 := range r1.Changes {
		for _, c2 := range r2.Changes {
			// we only have 1 add per request
			if c1.Add != nil && c2.Add != nil && checkScalewayRecords(c1.Add.Records, c2.Add.Records) {
				total--
			} else if c1.Delete != nil && c2.Delete != nil {
				if *c1.Delete.IDFields.Data == *c2.Delete.IDFields.Data && c1.Delete.IDFields.Name == c2.Delete.IDFields.Name && c1.Delete.IDFields.Type == c2.Delete.IDFields.Type {
					total--
				}
			}
		}
	}
	return total == 0
}

func checkScalewayRecords(rs1, rs2 []*domain.Record) bool {
	if len(rs1) != len(rs2) {
		return false
	}
	total := int(len(rs1))
	for _, r1 := range rs1 {
		for _, r2 := range rs2 {
			if r1.Data == r2.Data &&
				r1.Name == r2.Name &&
				r1.Priority == r2.Priority &&
				r1.TTL == r2.TTL &&
				r1.Type == r2.Type {
				total--
			}
		}
	}
	return total == 0
}
