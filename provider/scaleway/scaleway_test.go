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
	"io"
	"os"
	"reflect"
	"testing"

	domain "github.com/scaleway/scaleway-sdk-go/api/domain/v2beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockScalewayDomain struct {
	*domain.API
}

func (m *mockScalewayDomain) ListDNSZones(_ *domain.ListDNSZonesRequest, _ ...scw.RequestOption) (*domain.ListDNSZonesResponse, error) {
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

func (m *mockScalewayDomain) ListDNSZoneRecords(req *domain.ListDNSZoneRecordsRequest, _ ...scw.RequestOption) (*domain.ListDNSZoneRecordsResponse, error) {
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
			{
				Data:     "one.example.com.",
				Name:     "",
				TTL:      300,
				Priority: 0,
				Type:     domain.RecordTypeALIAS,
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
			{
				Data:     "foo.example.com.",
				Name:     "www",
				TTL:      300,
				Priority: 0,
				Type:     domain.RecordTypeALIAS,
			},
		}
	}
	return &domain.ListDNSZoneRecordsResponse{
		Records: records,
	}, nil
}

func (m *mockScalewayDomain) UpdateDNSZoneRecords(_ *domain.UpdateDNSZoneRecordsRequest, _ ...scw.RequestOption) (*domain.UpdateDNSZoneRecordsResponse, error) {
	return &domain.UpdateDNSZoneRecordsResponse{}, nil
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
	t.Setenv(scw.ScwActiveProfileEnv, "foo")
	t.Setenv(scw.ScwConfigPathEnv, tmpDir+"/config.yaml")
	_, err = newProvider(endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}

	t.Setenv(scw.ScwAccessKeyEnv, "SCWXXXXXXXXXXXXXXXXX")
	t.Setenv(scw.ScwSecretKeyEnv, "11111111-1111-1111-1111-111111111111")
	_, err = newProvider(endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}

	_ = os.Unsetenv(scw.ScwSecretKeyEnv)
	_, err = newProvider(endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}

	t.Setenv(scw.ScwSecretKeyEnv, "dummy")
	_, err = newProvider(endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}

	_ = os.Unsetenv(scw.ScwAccessKeyEnv)
	t.Setenv(scw.ScwSecretKeyEnv, "11111111-1111-1111-1111-111111111111")
	_, err = newProvider(endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}

	t.Setenv(scw.ScwAccessKeyEnv, "dummy")
	_, err = newProvider(endpoint.NewDomainFilter([]string{"example.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestScalewayProvider_OptionnalConfigFile(t *testing.T) {
	log.SetOutput(io.Discard)
	t.Setenv(scw.ScwAccessKeyEnv, "SCWXXXXXXXXXXXXXXXXX")
	t.Setenv(scw.ScwSecretKeyEnv, "11111111-1111-1111-1111-111111111111")

	_, err := newProvider(endpoint.NewDomainFilter([]string{"example.com"}), true)
	assert.NoError(t, err)
}

func TestScalewayProvider_AdjustEndpoints(t *testing.T) {
	provider := &ScalewayProvider{
		zoneNames: map[string]struct{}{
			"example.com": {},
		},
	}

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
		{
			// CNAME at the zone apex gets the alias property added
			DNSName:    "example.com",
			RecordTTL:  300,
			RecordType: "CNAME",
			Targets:    []string{"foo.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
		{
			// CNAME with the alias annotation keeps the alias property
			DNSName:    "four.example.com",
			RecordTTL:  300,
			RecordType: "CNAME",
			Targets:    []string{"foo.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  endpoint.ProviderSpecificAlias,
					Value: "true",
				},
			},
		},
		{
			// CNAME with the alias property set to false gets it removed
			DNSName:    "five.example.com",
			RecordTTL:  300,
			RecordType: "CNAME",
			Targets:    []string{"foo.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  endpoint.ProviderSpecificAlias,
					Value: "false",
				},
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
		{
			// non-CNAME with the alias property gets it removed
			DNSName:    "six.example.com",
			RecordTTL:  300,
			RecordType: "A",
			Targets:    []string{"1.1.1.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  endpoint.ProviderSpecificAlias,
					Value: "true",
				},
			},
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
		{
			DNSName:    "example.com",
			RecordTTL:  300,
			RecordType: "CNAME",
			Targets:    []string{"foo.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
				{
					Name:  endpoint.ProviderSpecificAlias,
					Value: "true",
				},
			},
		},
		{
			DNSName:    "four.example.com",
			RecordTTL:  300,
			RecordType: "CNAME",
			Targets:    []string{"foo.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  endpoint.ProviderSpecificAlias,
					Value: "true",
				},
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
		{
			DNSName:    "five.example.com",
			RecordTTL:  300,
			RecordType: "CNAME",
			Targets:    []string{"foo.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
			},
		},
		{
			DNSName:    "six.example.com",
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
	}

	after, err := provider.AdjustEndpoints(before)
	require.NoError(t, err)
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
	zones, err := provider.Zones(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	require.Len(t, zones, len(expected))
	for i, zone := range zones {
		assert.Equal(t, expected[i], zone)
	}

	// Zones caches the zone names for apex detection in AdjustEndpoints
	assert.Contains(t, provider.zoneNames, "example.com")
	assert.Contains(t, provider.zoneNames, "test.example.com")
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
		{
			// ALIAS records at the zone apex are read back as CNAME endpoints
			DNSName:    "example.com",
			RecordTTL:  300,
			RecordType: "CNAME",
			Targets:    []string{"one.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
				{
					Name:  endpoint.ProviderSpecificAlias,
					Value: "true",
				},
			},
		},
		{
			// ALIAS records below the zone apex are read back as CNAME endpoints too
			DNSName:    "www.test.example.com",
			RecordTTL:  300,
			RecordType: "CNAME",
			Targets:    []string{"foo.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  scalewayPriorityKey,
					Value: "0",
				},
				{
					Name:  endpoint.ProviderSpecificAlias,
					Value: "true",
				},
			},
		},
	}

	records, err := provider.Records(t.Context())
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
		assert.True(t, found)
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
							Data: new("3.3.3.3"),
							Name: "me",
							Type: domain.RecordTypeA,
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: new("1.1.1.1"),
							Name: "here",
							Type: domain.RecordTypeA,
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: new("1.1.1.2"),
							Name: "here",
							Type: domain.RecordTypeA,
						},
					},
				},
				{
					// apex CNAME records are deleted as ALIAS records
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: new("foo.example.com."),
							Name: "",
							Type: domain.RecordTypeALIAS,
						},
					},
				},
				{
					// CNAME records with the alias property are deleted as ALIAS records
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: new("bar.example.com."),
							Name: "old",
							Type: domain.RecordTypeALIAS,
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
								// apex CNAME records are created as ALIAS records
								Data:     "example.com.",
								Name:     "",
								TTL:      600,
								Type:     domain.RecordTypeALIAS,
								Priority: 20,
							},
							{
								// CNAME records with the alias property are created as ALIAS records
								Data:     "foo.example.com.",
								Name:     "www",
								TTL:      600,
								Type:     domain.RecordTypeALIAS,
								Priority: 0,
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
							Data: new("1.1.1.1"),
							Name: "here.is.my",
							Type: domain.RecordTypeA,
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: new("4.4.4.4"),
							Name: "my",
							Type: domain.RecordTypeA,
						},
					},
				},
				{
					Delete: &domain.RecordChangeDelete{
						IDFields: &domain.RecordIdentifier{
							Data: new("5.5.5.5"),
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
			{
				DNSName:    "www.test.example.com",
				RecordType: "CNAME",
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  endpoint.ProviderSpecificAlias,
						Value: "true",
					},
				},
				RecordTTL: 600,
				Targets:   []string{"foo.example.com"},
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
			{
				DNSName:    "example.com",
				RecordType: "CNAME",
				Targets:    []string{"foo.example.com"},
			},
			{
				DNSName:    "old.example.com",
				RecordType: "CNAME",
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  endpoint.ProviderSpecificAlias,
						Value: "true",
					},
				},
				Targets: []string{"bar.example.com"},
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

	requests, err := provider.generateApplyRequests(t.Context(), changes)
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

// TestScalewayProvider_AliasRoundTrip ensures that desired CNAME endpoints
// stored as ALIAS records, once adjusted, exactly match the endpoints read
// back from those ALIAS records. If they diverge, the plan detects a
// difference and rewrites the records on every reconciliation cycle.
func TestScalewayProvider_AliasRoundTrip(t *testing.T) {
	mocked := mockScalewayDomain{nil}
	provider := &ScalewayProvider{
		domainAPI:    &mocked,
		domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
	}

	// Records also warms the zone name cache used for apex detection
	current, err := provider.Records(t.Context())
	require.NoError(t, err)

	desired := []*endpoint.Endpoint{
		{
			// apex CNAME without any annotation
			DNSName:    "example.com",
			RecordType: "CNAME",
			RecordTTL:  300,
			Targets:    []string{"one.example.com"},
		},
		{
			// CNAME below the apex opted in via the alias annotation
			DNSName:    "www.test.example.com",
			RecordType: "CNAME",
			RecordTTL:  300,
			Targets:    []string{"foo.example.com"},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  endpoint.ProviderSpecificAlias,
					Value: "true",
				},
			},
		},
	}
	adjusted, err := provider.AdjustEndpoints(desired)
	require.NoError(t, err)

	for _, d := range adjusted {
		found := false
		for _, c := range current {
			if c.DNSName == d.DNSName && c.RecordType == d.RecordType {
				found = true
				assert.True(t, c.Targets.Same(d.Targets), "targets mismatch for %s", d.DNSName)
				assert.Equal(t, c.RecordTTL, d.RecordTTL, "TTL mismatch for %s", d.DNSName)
				assert.ElementsMatch(t, c.ProviderSpecific, d.ProviderSpecific,
					"provider-specific mismatch for %s would cause an update on every cycle", d.DNSName)
			}
		}
		assert.True(t, found, "no record found for %s", d.DNSName)
	}
}

func TestScalewayProvider_scalewayRecordType(t *testing.T) {
	tests := []struct {
		name               string
		ep                 *endpoint.Endpoint
		relativeRecordName string
		expected           domain.RecordType
	}{
		{
			name:               "CNAME at zone apex becomes ALIAS",
			ep:                 &endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME},
			relativeRecordName: "",
			expected:           domain.RecordTypeALIAS,
		},
		{
			name:               "CNAME below zone apex stays CNAME",
			ep:                 &endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME},
			relativeRecordName: "www",
			expected:           domain.RecordTypeCNAME,
		},
		{
			name: "CNAME with alias property becomes ALIAS",
			ep: (&endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME}).
				WithProviderSpecific(endpoint.ProviderSpecificAlias, "true"),
			relativeRecordName: "www",
			expected:           domain.RecordTypeALIAS,
		},
		{
			name: "CNAME with alias property set to false stays CNAME",
			ep: (&endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME}).
				WithProviderSpecific(endpoint.ProviderSpecificAlias, "false"),
			relativeRecordName: "www",
			expected:           domain.RecordTypeCNAME,
		},
		{
			// "A" and "AAAA" are AWS-specific dual-stack alias values, they do
			// not opt in to Scaleway ALIAS records
			name: "CNAME with alias property set to A stays CNAME",
			ep: (&endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME}).
				WithProviderSpecific(endpoint.ProviderSpecificAlias, "A"),
			relativeRecordName: "www",
			expected:           domain.RecordTypeCNAME,
		},
		{
			name: "CNAME with alias property set to AAAA stays CNAME",
			ep: (&endpoint.Endpoint{RecordType: endpoint.RecordTypeCNAME}).
				WithProviderSpecific(endpoint.ProviderSpecificAlias, "AAAA"),
			relativeRecordName: "www",
			expected:           domain.RecordTypeCNAME,
		},
		{
			name:               "A at zone apex stays A",
			ep:                 &endpoint.Endpoint{RecordType: endpoint.RecordTypeA},
			relativeRecordName: "",
			expected:           domain.RecordTypeA,
		},
		{
			name: "A with alias property stays A",
			ep: (&endpoint.Endpoint{RecordType: endpoint.RecordTypeA}).
				WithProviderSpecific(endpoint.ProviderSpecificAlias, "true"),
			relativeRecordName: "www",
			expected:           domain.RecordTypeA,
		},
		{
			name:               "TXT below zone apex stays TXT",
			ep:                 &endpoint.Endpoint{RecordType: endpoint.RecordTypeTXT},
			relativeRecordName: "www",
			expected:           domain.RecordTypeTXT,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, scalewayRecordType(tt.ep, tt.relativeRecordName))
		})
	}
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
