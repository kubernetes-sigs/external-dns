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
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
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
		title             string
		findType          string
		findSetIdentifier string
		records           []*inMemoryRecord
		expected          *inMemoryRecord
		expectedEmpty     bool
	}{
		{
			title:         "no records, empty type",
			findType:      "",
			records:       nil,
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title:         "no records, non-empty type",
			findType:      endpoint.RecordTypeA,
			records:       nil,
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title:    "one record, empty type",
			findType: "",
			records: []*inMemoryRecord{
				{
					Type: endpoint.RecordTypeA,
				},
			},
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title:    "one record, wrong type",
			findType: endpoint.RecordTypeCNAME,
			records: []*inMemoryRecord{
				{
					Type: endpoint.RecordTypeA,
				},
			},
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title:    "one record, right type",
			findType: endpoint.RecordTypeA,
			records: []*inMemoryRecord{
				{
					Type: endpoint.RecordTypeA,
				},
			},
			expected: &inMemoryRecord{
				Type: endpoint.RecordTypeA,
			},
		},
		{
			title:    "multiple records, right type",
			findType: endpoint.RecordTypeA,
			records: []*inMemoryRecord{
				{
					Type: endpoint.RecordTypeA,
				},
				{
					Type: endpoint.RecordTypeTXT,
				},
			},
			expected: &inMemoryRecord{
				Type: endpoint.RecordTypeA,
			},
		},
		{
			title:             "multiple records, right type and set identifier",
			findType:          endpoint.RecordTypeA,
			findSetIdentifier: "test-set-1",
			records: []*inMemoryRecord{
				{
					Type:          endpoint.RecordTypeA,
					SetIdentifier: "test-set-1",
				},
				{
					Type:          endpoint.RecordTypeA,
					SetIdentifier: "test-set-2",
				},
				{
					Type: endpoint.RecordTypeTXT,
				},
			},
			expected: &inMemoryRecord{
				Type:          endpoint.RecordTypeA,
				SetIdentifier: "test-set-1",
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			c := newInMemoryClient()
			record := c.findByTypeAndSetIdentifier(ti.findType, ti.findSetIdentifier, ti.records)
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
		zone        string
		expectError bool
		init        map[string]zone
		expected    []*endpoint.Endpoint
	}{
		{
			title:       "no records, no zone",
			zone:        "",
			init:        map[string]zone{},
			expectError: false,
		},
		{
			title: "records, wrong zone",
			zone:  "net",
			init: map[string]zone{
				"org": {},
				"com": {},
			},
			expectError: false,
		},
		{
			title: "records, zone with records",
			zone:  "org",
			init: map[string]zone{
				"org": {
					"example.org": []*inMemoryRecord{
						{
							Name:   "example.org",
							Target: "8.8.8.8",
							Type:   endpoint.RecordTypeA,
						},
						{
							Name: "example.org",
							Type: endpoint.RecordTypeTXT,
						},
					},
					"foo.org": []*inMemoryRecord{
						{
							Name:   "foo.org",
							Target: "4.4.4.4",
							Type:   endpoint.RecordTypeCNAME,
						},
					},
				},
				"com": {
					"example.com": []*inMemoryRecord{
						{
							Name:   "example.com",
							Target: "4.4.4.4",
							Type:   endpoint.RecordTypeCNAME,
						},
					},
				},
			},
			expectError: false,
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "example.org",
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordType: endpoint.RecordTypeA,
				},
				{
					DNSName:    "example.org",
					RecordType: endpoint.RecordTypeTXT,
					Targets:    endpoint.Targets{""},
				},
				{
					DNSName:    "foo.org",
					Targets:    endpoint.Targets{"4.4.4.4"},
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
			f := filter{domain: ti.zone}
			im.filter = &f
			records, err := im.Records(context.Background())
			if ti.expectError {
				assert.Nil(t, records)
				assert.EqualError(t, err, ErrZoneNotFound.Error())
			} else {
				require.NoError(t, err)
				assert.True(t, testutils.SameEndpoints(ti.expected, records), "Endpoints not the same: Expected: %+v Records: %+v", ti.expected, records)
			}
		})
	}
}

func testInMemoryValidateChangeBatch(t *testing.T) {
	init := map[string]zone{
		"org": {
			"example.org": []*inMemoryRecord{
				{
					Name:   "example.org",
					Target: "8.8.8.8",
					Type:   endpoint.RecordTypeA,
				},
				{
					Name: "example.org",
				},
			},
			"foo.org": []*inMemoryRecord{
				{
					Name:   "foo.org",
					Target: "bar.org",
					Type:   endpoint.RecordTypeCNAME,
				},
			},
			"foo.bar.org": []*inMemoryRecord{
				{
					Name:   "foo.bar.org",
					Target: "5.5.5.5",
					Type:   endpoint.RecordTypeA,
				},
			},
		},
		"com": {
			"example.com": []*inMemoryRecord{
				{
					Name:   "example.com",
					Target: "another-example.com",
					Type:   endpoint.RecordTypeCNAME,
				},
			},
		},
	}
	for _, ti := range []struct {
		title       string
		expectError bool
		errorType   error
		init        map[string]zone
		changes     *plan.Changes
		zone        string
	}{
		{
			title:       "no zones, no update",
			expectError: true,
			zone:        "",
			init:        map[string]zone{},
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
						Targets:    endpoint.Targets{"8.8.8.8"},
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
						Targets:    endpoint.Targets{"4.4.4.4"},
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "foo.org",
						Targets:    endpoint.Targets{"4.4.4.4"},
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
						Targets:    endpoint.Targets{"4.4.4.4"},
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "foo.org",
						Targets:    endpoint.Targets{"4.4.4.4"},
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
						Targets:    endpoint.Targets{"4.4.4.4"},
						RecordType: endpoint.RecordTypeA,
					},
					{
						DNSName:    "foo.org",
						Targets:    endpoint.Targets{"4.4.4.4"},
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
						Targets:    endpoint.Targets{"8.8.8.8"},
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Targets:    endpoint.Targets{"8.8.8.8"},
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
						Targets:    endpoint.Targets{"8.8.8.8"},
						RecordType: endpoint.RecordTypeA,
					},
					{
						DNSName:    "example.org",
						Targets:    endpoint.Targets{"8.8.8.8"},
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
						Targets:    endpoint.Targets{"8.8.8.8"},
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
						Targets:    endpoint.Targets{"8.8.8.8"},
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
						Targets:    endpoint.Targets{"5.5.5.5"},
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
						Targets:    endpoint.Targets{"4.8.8.9"},
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Targets:    endpoint.Targets{"4.8.8.4"},
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Targets:    endpoint.Targets{"5.5.5.5"},
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
				Create:    convertToInMemoryRecord(ti.changes.Create),
				UpdateNew: convertToInMemoryRecord(ti.changes.UpdateNew),
				UpdateOld: convertToInMemoryRecord(ti.changes.UpdateOld),
				Delete:    convertToInMemoryRecord(ti.changes.Delete),
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

func getInitData() map[string]zone {
	return map[string]zone{
		"org": {
			"example.org": []*inMemoryRecord{
				{
					Name:   "example.org",
					Target: "8.8.8.8",
					Type:   endpoint.RecordTypeA,
				},
				{
					Name: "example.org",
					Type: endpoint.RecordTypeTXT,
				},
			},
			"foo.org": []*inMemoryRecord{
				{
					Name:   "foo.org",
					Target: "4.4.4.4",
					Type:   endpoint.RecordTypeCNAME,
				},
			},
			"foo.bar.org": []*inMemoryRecord{
				{
					Name:   "foo.bar.org",
					Target: "5.5.5.5",
					Type:   endpoint.RecordTypeA,
				},
			},
		},
		"com": {
			"example.com": []*inMemoryRecord{
				{
					Name:   "example.com",
					Target: "4.4.4.4",
					Type:   endpoint.RecordTypeCNAME,
				},
			},
		},
	}
}

func testInMemoryApplyChanges(t *testing.T) {
	for _, ti := range []struct {
		title              string
		expectError        bool
		init               map[string]zone
		changes            *plan.Changes
		expectedZonesState map[string]zone
	}{
		{
			title:       "unmatched zone, should be ignored in the apply step",
			expectError: false,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{{
					DNSName:    "example.de",
					Targets:    endpoint.Targets{"8.8.8.8"},
					RecordType: endpoint.RecordTypeA,
				}},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			expectedZonesState: getInitData(),
		},
		{
			title:       "expect error",
			expectError: true,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Targets:    endpoint.Targets{"8.8.8.8"},
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Targets:    endpoint.Targets{"8.8.8.8"},
						RecordType: endpoint.RecordTypeA,
					},
				},
			},
		},
		{
			title:       "zones, update, right zone, valid batch - delete",
			expectError: false,
			changes: &plan.Changes{
				Create:    []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Targets:    endpoint.Targets{"5.5.5.5"},
						RecordType: endpoint.RecordTypeA,
					},
				},
			},
			expectedZonesState: map[string]zone{
				"org": {
					"example.org": []*inMemoryRecord{
						{

							Name:   "example.org",
							Target: "8.8.8.8",
							Type:   endpoint.RecordTypeA,
						},
						{

							Name: "example.org",
							Type: endpoint.RecordTypeTXT,
						},
					},
					"foo.org": []*inMemoryRecord{
						{

							Name:   "foo.org",
							Target: "4.4.4.4",
							Type:   endpoint.RecordTypeCNAME,
						},
					},
					"foo.bar.org": []*inMemoryRecord{},
				},
				"com": {
					"example.com": []*inMemoryRecord{
						{
							Name:   "example.com",
							Target: "4.4.4.4",
							Type:   endpoint.RecordTypeCNAME,
						},
					},
				},
			},
		},
		{
			title:       "zones, update, right zone, valid batch - update, create, delete",
			expectError: false,
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.new.org",
						Targets:    endpoint.Targets{"4.8.8.9"},
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Targets:    endpoint.Targets{"4.8.8.4"},
						RecordType: endpoint.RecordTypeA,
					},
				},
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName:    "foo.bar.org",
						Targets:    endpoint.Targets{"5.5.5.5"},
						RecordType: endpoint.RecordTypeA,
					},
				},
				Delete: []*endpoint.Endpoint{
					{
						DNSName:    "example.org",
						Targets:    endpoint.Targets{"8.8.8.8"},
						RecordType: endpoint.RecordTypeA,
					},
				},
			},
			expectedZonesState: map[string]zone{
				"org": {
					"example.org": []*inMemoryRecord{
						{
							Name: "example.org",
							Type: endpoint.RecordTypeTXT,
						},
					},
					"foo.org": []*inMemoryRecord{
						{
							Name:   "foo.org",
							Target: "4.4.4.4",
							Type:   endpoint.RecordTypeCNAME,
						},
					},
					"foo.bar.org": []*inMemoryRecord{
						{
							Name:   "foo.bar.org",
							Target: "4.8.8.4",
							Type:   endpoint.RecordTypeA,
						},
					},
					"foo.bar.new.org": []*inMemoryRecord{
						{
							Name:   "foo.bar.new.org",
							Target: "4.8.8.9",
							Type:   endpoint.RecordTypeA,
						},
					},
				},
				"com": {
					"example.com": []*inMemoryRecord{
						{
							Name:   "example.com",
							Target: "4.4.4.4",
							Type:   endpoint.RecordTypeCNAME,
						},
					},
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {

			im := NewInMemoryProvider()
			c := &inMemoryClient{}
			c.zones = getInitData()
			im.client = c

			err := im.ApplyChanges(context.Background(), ti.changes)
			if ti.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, ti.expectedZonesState, c.zones)
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
