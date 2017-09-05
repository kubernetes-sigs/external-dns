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
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
	"github.com/kubernetes-incubator/external-dns/plan"

	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	_ Provider = &InMemoryProvider{}
)

func TestInMemoryProvider(t *testing.T) {
	t.Run("findByType", testInMemoryFindByType)
	t.Run("Records", testInMemoryRecords)
	t.Run("validateChangeBatch", testInMemoryValidateChangeBatch)
	t.Run("ApplyChanges", testInMemoryApplyChanges)
	t.Run("NewInMemoryProvider", testNewInMemoryProvider)
	t.Run("CreateZone", testInMemoryCreateZone)
}

func testInMemoryFindByType(t *testing.T) {
	for _, ti := range []struct {
		title         string
		findEndpoint  *endpoint.Endpoint
		records       []*endpoint.Endpoint
		expected      *endpoint.Endpoint
		expectedEmpty bool
	}{
		{
			title:         "no records, empty type",
			findEndpoint:  &endpoint.Endpoint{},
			records:       nil,
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title: "no records, non-empty type",
			findEndpoint: &endpoint.Endpoint{
				RecordType: endpoint.RecordTypeA,
			},
			records:       nil,
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title:        "one record, empty type",
			findEndpoint: &endpoint.Endpoint{},
			records: []*endpoint.Endpoint{
				{
					RecordType: endpoint.RecordTypeA,
				},
			},
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title: "one record, wrong type",
			findEndpoint: &endpoint.Endpoint{
				RecordType: endpoint.RecordTypeCNAME,
			},
			records: []*endpoint.Endpoint{
				{
					RecordType: endpoint.RecordTypeA,
				},
			},
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title: "one record, right type",
			findEndpoint: &endpoint.Endpoint{
				RecordType: endpoint.RecordTypeA,
			},
			records: []*endpoint.Endpoint{
				{
					RecordType: endpoint.RecordTypeA,
				},
			},
			expected: &endpoint.Endpoint{
				RecordType: endpoint.RecordTypeA,
			},
		},
		{
			title: "multiple records, right type",
			findEndpoint: &endpoint.Endpoint{
				RecordType: endpoint.RecordTypeA,
			},
			records: []*endpoint.Endpoint{
				{
					RecordType: endpoint.RecordTypeA,
				},
				{
					RecordType: endpoint.RecordTypeTXT,
				},
			},
			expected: &endpoint.Endpoint{
				RecordType: endpoint.RecordTypeA,
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			c := newInMemoryClient()
			record := c.findRecord(ti.findEndpoint, ti.records)
			if ti.expectedEmpty {
				assert.Nil(t, record)
			} else {
				require.NotNil(t, record)
				assert.Equal(t, *ti.expected, *record)
			}
		})
	}
}

func testInMemoryRecords(t *testing.T) {
	for _, ti := range []struct {
		title       string
		zone        InMemoryZone
		expectError bool
		init        []InMemoryZone
		expected    []*endpoint.Endpoint
	}{
		{
			title:       "no records, no zone",
			zone:        InMemoryZone{},
			init:        []InMemoryZone{},
			expectError: false,
		},
		{
			title: "records, wrong zone",
			zone:  InMemoryZone{ZoneID: "net"},
			init: []InMemoryZone{
				{
					ZoneID: "org",
				},
				{
					ZoneID: "com",
				},
			},
			expectError: false,
		},
		{
			title: "records, zone with records",
			zone:  InMemoryZone{ZoneID: "org"},
			init: []InMemoryZone{
				{
					ZoneID: "org",
					Endpoints: []*endpoint.Endpoint{
						{
							DNSName:    "example.org",
							Target:     "8.8.8.8",
							RecordType: endpoint.RecordTypeA,
						},
						{
							DNSName:    "example.org",
							RecordType: endpoint.RecordTypeTXT,
						},
						{
							DNSName:    "foo.org",
							Target:     "4.4.4.4",
							RecordType: endpoint.RecordTypeCNAME,
						},
					},
				},
				{
					ZoneID: "com",
					Endpoints: []*endpoint.Endpoint{
						{
							DNSName:    "example.com",
							Target:     "4.4.4.4",
							RecordType: endpoint.RecordTypeCNAME,
						},
					},
				},
			},
			expectError: false,
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Target:     "8.8.8.8",
					RecordType: endpoint.RecordTypeA,
				},
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeTXT,
				},
				{
					DNSName:    "foo.org",
					Target:     "4.4.4.4",
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			c := newInMemoryClient()
			c.zones = ti.init
			im := NewInMemoryProvider()
			im.client = c
			f := filter{domain: ti.zone.ZoneID}
			im.filter = &f
			records, err := im.Records()
			if ti.expectError {
				assert.Nil(t, records)
				assert.EqualError(t, err, ErrZoneNotFound.Error())
			} else {
				require.NoError(t, err)
				assert.True(t, testutils.SameEndpoints(ti.expected, records))
			}
		})
	}
}

func testInMemoryValidateChangeBatch(t *testing.T) {
	init := []InMemoryZone{
		{
			ZoneID: "org",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Target:     "8.8.8.8",
					RecordType: endpoint.RecordTypeA,
				},
				{
					Target: "example.org",
				},
				{
					DNSName:    "foo.org",
					Target:     "bar.org",
					RecordType: endpoint.RecordTypeCNAME,
				},
				{
					DNSName:    "foo.bar.org",
					Target:     "5.5.5.5",
					RecordType: endpoint.RecordTypeA,
				},
			},
		},
		{
			ZoneID: "com",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "example.com",
					Target:     "another-example.com",
					RecordType: endpoint.RecordTypeCNAME,
				},
			},
		},
	}
	for _, ti := range []struct {
		title       string
		expectError bool
		errorType   error
		init        []InMemoryZone
		changes     *plan.Changes
		zone        string
	}{
		{
			title:       "no zones, no update",
			expectError: true,
			zone:        "",
			init:        []InMemoryZone{},
			changes: &plan.Changes{
				Create:    []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrZoneNotFound,
		},
		{
			title:       "zones, no update",
			expectError: true,
			zone:        "",
			init:        init,
			changes: &plan.Changes{
				Create:    []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrZoneNotFound,
		},
		{
			title:       "zones, update, wrong zone",
			expectError: true,
			zone:        "test",
			init:        init,
			changes: &plan.Changes{
				Create:    []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrZoneNotFound,
		},
		{
			title:       "zones, update, right zone, invalid batch - already exists",
			expectError: true,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrRecordAlreadyExists,
		},
		{
			title:       "zones, update, right zone, invalid batch - record not found for update",
			expectError: true,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "foo.org",
						Target:     "4.4.4.4",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "foo.org",
						Target:     "4.4.4.4",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrRecordNotFound,
		},
		{
			title:       "zones, update, right zone, invalid batch - record not found for update",
			expectError: true,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "foo.org",
						Target:     "4.4.4.4",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "foo.org",
						Target:     "4.4.4.4",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrRecordNotFound,
		},
		{
			title:       "zones, update, right zone, invalid batch - duplicated create",
			expectError: true,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "foo.org",
						Target:     "4.4.4.4",
						RecordType: endpoint.RecordTypeA,
					},
					{
						DNSName:    "foo.org",
						Target:     "4.4.4.4",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrDuplicateRecordFound,
		},
		{
			title:       "zones, update, right zone, invalid batch - duplicated update/delete",
			expectError: true,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
				},
			},
			errorType: ErrDuplicateRecordFound,
		},
		{
			title:       "zones, update, right zone, invalid batch - duplicated update",
			expectError: true,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
					{
						DNSName:    "example.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrDuplicateRecordFound,
		},
		{
			title:       "zones, update, right zone, invalid batch - wrong update old",
			expectError: true,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create:    []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "new.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
				},
				Delete: []*endpoint.Endpoint{},
			},
			errorType: ErrRecordNotFound,
		},
		{
			title:       "zones, update, right zone, invalid batch - wrong delete",
			expectError: true,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create:    []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "new.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
				},
			},
			errorType: ErrRecordNotFound,
		},
		{
			title:       "zones, update, right zone, valid batch - delete",
			expectError: false,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create:    []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Target:     "5.5.5.5",
						RecordType: endpoint.RecordTypeA,
					},
				},
			},
		},
		{
			title:       "zones, update, right zone, valid batch - update and create",
			expectError: false,
			zone:        "org",
			init:        init,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.new.org",
						Target:     "4.8.8.9",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Target:     "4.8.8.4",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Target:     "5.5.5.5",
						RecordType: endpoint.RecordTypeA,
					},
				},
				Delete: []*endpoint.Endpoint{},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			c := &inMemoryClient{}
			c.zones = ti.init
			ichanges := &inMemoryChange{
				Create:    ti.changes.Create,
				UpdateNew: ti.changes.UpdateNew,
				UpdateOld: ti.changes.UpdateOld,
				Delete:    ti.changes.Delete,
			}
			err := c.validateChangeBatch(ti.zone, ichanges)
			if ti.expectError {
				assert.EqualError(t, err, ti.errorType.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func testInMemoryApplyChanges(t *testing.T) {
	for _, ti := range []struct {
		title              string
		expectError        bool
		init               []InMemoryZone
		changes            *plan.Changes
		zone               string
		expectedZonesState []InMemoryZone
	}{
		{
			title:       "expect error",
			expectError: true,
			zone:        "org",
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
				},
			},
		},
		{
			title:       "zones, update, right zone, valid batch - delete",
			expectError: false,
			zone:        "org",
			changes: &plan.Changes{
				Create:    []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Target:     "5.5.5.5",
						RecordType: endpoint.RecordTypeA,
					},
				},
			},
			expectedZonesState: []InMemoryZone{
				{
					ZoneID: "org",
					Endpoints: []*endpoint.Endpoint{
						{

							DNSName:    "example.org",
							Target:     "8.8.8.8",
							RecordType: endpoint.RecordTypeA,
						},
						{

							DNSName:    "example.org",
							RecordType: endpoint.RecordTypeTXT,
						},
						{

							DNSName:    "foo.org",
							Target:     "4.4.4.4",
							RecordType: endpoint.RecordTypeCNAME,
						},
					},
				},
				{
					ZoneID: "com",
					Endpoints: []*endpoint.Endpoint{
						{
							DNSName:    "example.com",
							Target:     "4.4.4.4",
							RecordType: endpoint.RecordTypeCNAME,
						},
					},
				},
			},
		},
		{
			title:       "zones, update, right zone, valid batch - update, create, delete",
			expectError: false,
			zone:        "org",
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.new.org",
						Target:     "4.8.8.9",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Target:     "4.8.8.4",
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Target:     "5.5.5.5",
						RecordType: endpoint.RecordTypeA,
					},
				},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Target:     "8.8.8.8",
						RecordType: endpoint.RecordTypeA,
					},
				},
			},
			expectedZonesState: []InMemoryZone{
				{
					ZoneID: "org",
					Endpoints: []*endpoint.Endpoint{
						{
							DNSName:    "example.org",
							RecordType: endpoint.RecordTypeTXT,
						},
						{
							DNSName:    "foo.org",
							Target:     "4.4.4.4",
							RecordType: endpoint.RecordTypeCNAME,
						},
						{
							DNSName:    "foo.bar.org",
							Target:     "4.8.8.4",
							RecordType: endpoint.RecordTypeA,
						},
						{
							DNSName:    "foo.bar.new.org",
							Target:     "4.8.8.9",
							RecordType: endpoint.RecordTypeA,
						},
					},
				},
				{
					ZoneID: "com",
					Endpoints: []*endpoint.Endpoint{
						{
							DNSName:    "example.com",
							Target:     "4.4.4.4",
							RecordType: endpoint.RecordTypeCNAME,
						},
					},
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			init := []InMemoryZone{
				{
					ZoneID: "org",
					Endpoints: []*endpoint.Endpoint{
						{
							DNSName:    "example.org",
							Target:     "8.8.8.8",
							RecordType: endpoint.RecordTypeA,
						},
						{
							DNSName:    "example.org",
							RecordType: endpoint.RecordTypeTXT,
						},
						{
							DNSName:    "foo.org",
							Target:     "4.4.4.4",
							RecordType: endpoint.RecordTypeCNAME,
						},
						{
							DNSName:    "foo.bar.org",
							Target:     "5.5.5.5",
							RecordType: endpoint.RecordTypeA,
						},
					},
				},
				{
					ZoneID: "com",
					Endpoints: []*endpoint.Endpoint{
						{
							DNSName:    "example.com",
							Target:     "4.4.4.4",
							RecordType: endpoint.RecordTypeCNAME,
						},
					},
				},
			}
			im := NewInMemoryProvider()
			c := &inMemoryClient{}
			c.zones = init
			im.client = c

			err := im.ApplyChanges(ti.changes)
			if ti.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NoError(t, checkSameZones(ti.expectedZonesState, c.zones))
			}
		})
	}
}

func testNewInMemoryProvider(t *testing.T) {
	cfg := NewInMemoryProvider()
	assert.NotNil(t, cfg.client)
}

func testInMemoryCreateZone(t *testing.T) {
	im := NewInMemoryProvider()
	err := im.CreateZone("zone")
	assert.NoError(t, err)
	err = im.CreateZone("zone")
	assert.EqualError(t, err, ErrZoneAlreadyExists.Error())
}

func checkSameZones(expected, current []InMemoryZone) error {
	for _, expectedZone := range expected {
		exists := false
		for _, currentZone := range current {
			if expectedZone.ZoneID == currentZone.ZoneID {
				exists = true
				if !testutils.SameEndpoints(expectedZone.Endpoints, currentZone.Endpoints) {
					return fmt.Errorf("Endpoints for zone do not match. Expected: Zone %s %s, "+
						"Actual: Zone %s %s do not match", expectedZone.ZoneID,
						expectedZone.Endpoints, currentZone.ZoneID, currentZone.Endpoints)
				}
			}
		}
		if !exists {
			return fmt.Errorf("Does not have zone %s", expectedZone.ZoneID)
		}
	}
	return nil
}
