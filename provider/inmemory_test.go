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
	"reflect"
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/internal/testutils"
	"github.com/kubernetes-incubator/external-dns/plan"
)

var _ Provider = &InMemoryProvider{}

func TestInMemoryProvider(t *testing.T) {
	t.Run("Records", testInMemoryRecords)
	t.Run("endpoints", testInMemoryEndpoints)
	t.Run("findByType", testInMemoryFindByType)
	t.Run("validateChangeBatch", testInMemoryValidateChangeBatch)
	t.Run("ApplyChanges", testInMemoryApplyChanges)
	t.Run("NewInMemoryProvider", testNewInMemoryProvider)
	t.Run("CreateZone", testInMemoryCreateZone)
}

func testInMemoryFindByType(t *testing.T) {
	for _, ti := range []struct {
		title         string
		findType      string
		records       []*InMemoryRecord
		expected      *InMemoryRecord
		expectedEmpty bool
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
			findType:      "A",
			records:       nil,
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title:    "one record, empty type",
			findType: "",
			records: []*InMemoryRecord{
				{
					Type: "A",
				},
			},
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title:    "one record, wrong type",
			findType: "CNAME",
			records: []*InMemoryRecord{
				{
					Type: "A",
				},
			},
			expected:      nil,
			expectedEmpty: true,
		},
		{
			title:    "one record, right type",
			findType: "A",
			records: []*InMemoryRecord{
				{
					Type: "A",
				},
			},
			expected: &InMemoryRecord{
				Type: "A",
			},
		},
		{
			title:    "multiple records, right type",
			findType: "A",
			records: []*InMemoryRecord{
				{
					Type: "A",
				},
				{
					Type: "TXT",
				},
			},
			expected: &InMemoryRecord{
				Type: "A",
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			im := NewInMemoryProvider()
			record := im.findByType(ti.findType, ti.records)
			if ti.expectedEmpty && record != nil {
				t.Errorf("should return nil")
			}
			if !ti.expectedEmpty && record == nil {
				t.Errorf("should not return nil")
			}
			if !ti.expectedEmpty && record != nil && !reflect.DeepEqual(*record, *ti.expected) {
				t.Errorf("wrong record found")
			}
		})
	}
}

func testInMemoryEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title    string
		zone     string
		init     map[string]zone
		expected []*endpoint.Endpoint
	}{
		{
			title:    "no records, no zone",
			zone:     "",
			init:     map[string]zone{},
			expected: []*endpoint.Endpoint{},
		},
		{
			title:    "no records, zone",
			zone:     "central",
			init:     map[string]zone{},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "records, no zone",
			zone:  "",
			init: map[string]zone{
				"org": {
					"example.org": []*InMemoryRecord{
						{},
					},
					"foo.org": []*InMemoryRecord{
						{},
					},
				},
				"com": {
					"example.com": []*InMemoryRecord{
						{},
					},
					"foo.com": []*InMemoryRecord{
						{},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "records, zone with no records",
			zone:  "",
			init: map[string]zone{
				"org": {
					"example.org": []*InMemoryRecord{
						{},
					},
					"foo.org": []*InMemoryRecord{
						{},
					},
				},
				"com": {
					"example.com": []*InMemoryRecord{
						{},
					},
					"foo.com": []*InMemoryRecord{
						{},
					},
				},
			},
			expected: []*endpoint.Endpoint{},
		},
		{
			title: "records, zone with records",
			zone:  "org",
			init: map[string]zone{
				"org": {
					"example.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.org",
								Target:  "8.8.8.8",
							},
							Type: defaultType,
						},
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.org",
							},
							Type: "TXT",
						},
					},
					"foo.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "foo.org",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
						},
					},
				},
				"com": {
					"example.com": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.com",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
						},
					},
				},
			},
			expected: []*endpoint.Endpoint{
				{
					DNSName: "example.org",
					Target:  "8.8.8.8",
				},
				{
					DNSName: "example.org",
				},
				{
					DNSName: "foo.org",
					Target:  "4.4.4.4",
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			im := &InMemoryProvider{zones: ti.init}
			if !testutils.SameEndpoints(im.endpoints(ti.zone), ti.expected) {
				t.Errorf("endpoints returned wrong set")
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
	}{
		{
			title:       "no records, no zone",
			zone:        "",
			init:        map[string]zone{},
			expectError: true,
		},
		{
			title: "records, wrong zone",
			zone:  "net",
			init: map[string]zone{
				"org": {},
				"com": {},
			},
			expectError: true,
		},
		{
			title: "records, zone with records",
			zone:  "org",
			init: map[string]zone{
				"org": {
					"example.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.org",
								Target:  "8.8.8.8",
							},
							Type: defaultType,
						},
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.org",
							},
							Type: "TXT",
						},
					},
					"foo.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "foo.org",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
						},
					},
				},
				"com": {
					"example.com": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.com",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
						},
					},
				},
			},
			expectError: false,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			im := &InMemoryProvider{zones: ti.init}
			records, err := im.Records(ti.zone)
			if ti.expectError && records != nil {
				t.Errorf("wrong zone should not return records")
			}
			if ti.expectError && err != ErrZoneNotFound {
				t.Errorf("expected error")
			}
			if !ti.expectError && err != nil {
				t.Errorf("unexpected error")
			}
			if !ti.expectError && !testutils.SameEndpoints(im.endpoints(ti.zone), records) {
				t.Errorf("endpoints returned wrong set")
			}
		})
	}
}

func testInMemoryValidateChangeBatch(t *testing.T) {
	init := map[string]zone{
		"org": {
			"example.org": []*InMemoryRecord{
				{
					Endpoint: &endpoint.Endpoint{
						DNSName: "example.org",
						Target:  "8.8.8.8",
					},
					Type: defaultType,
				},
				{
					Endpoint: &endpoint.Endpoint{
						DNSName: "example.org",
					},
					Type: "TXT",
				},
			},
			"foo.org": []*InMemoryRecord{
				{
					Endpoint: &endpoint.Endpoint{
						DNSName: "foo.org",
						Target:  "4.4.4.4",
					},
					Type: "CNAME",
				},
			},
			"foo.bar.org": []*InMemoryRecord{
				{
					Endpoint: &endpoint.Endpoint{
						DNSName: "foo.bar.org",
						Target:  "5.5.5.5",
					},
					Type: defaultType,
				},
			},
		},
		"com": {
			"example.com": []*InMemoryRecord{
				{
					Endpoint: &endpoint.Endpoint{
						DNSName: "example.com",
						Target:  "4.4.4.4",
					},
					Type: "CNAME",
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
						DNSName: "example.org",
						Target:  "8.8.8.8",
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
						DNSName: "foo.org",
						Target:  "4.4.4.4",
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName: "foo.org",
						Target:  "4.4.4.4",
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
						DNSName: "foo.org",
						Target:  "4.4.4.4",
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName: "foo.org",
						Target:  "4.4.4.4",
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
						DNSName: "foo.org",
						Target:  "4.4.4.4",
					},
					{
						DNSName: "foo.org",
						Target:  "4.4.4.4",
					},
				},
				UpdateNew: []*endpoint.Endpoint{},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrInvalidBatchRequest,
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
						DNSName: "example.org",
						Target:  "8.8.8.8",
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName: "example.org",
						Target:  "8.8.8.8",
					},
				},
			},
			errorType: ErrInvalidBatchRequest,
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
						DNSName: "example.org",
						Target:  "8.8.8.8",
					},
					{
						DNSName: "example.org",
						Target:  "8.8.8.8",
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete:    []*endpoint.Endpoint{},
			},
			errorType: ErrInvalidBatchRequest,
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
						DNSName: "new.org",
						Target:  "8.8.8.8",
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
						DNSName: "new.org",
						Target:  "8.8.8.8",
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
						DNSName: "foo.bar.org",
						Target:  "5.5.5.5",
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
						DNSName: "foo.bar.new.org",
						Target:  "4.8.8.9",
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName: "foo.bar.org",
						Target:  "4.8.8.4",
					},
				},
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName: "foo.bar.org",
						Target:  "5.5.5.5",
					},
				},
				Delete: []*endpoint.Endpoint{},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			im := &InMemoryProvider{
				zones: ti.init,
			}
			err := im.validateChangeBatch(ti.zone, ti.changes)
			if ti.expectError && err != ti.errorType {
				t.Errorf("returns wrong type of error: %v, expected: %v", err, ti.errorType)
			}
			if !ti.expectError && err != nil {
				t.Error(err)
			}
		})
	}
}

func testInMemoryApplyChanges(t *testing.T) {
	for _, ti := range []struct {
		title              string
		expectError        bool
		init               map[string]zone
		changes            *plan.Changes
		zone               string
		expectedZonesState map[string]zone
	}{
		{
			title:       "expect error",
			expectError: true,
			zone:        "org",
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName: "example.org",
						Target:  "8.8.8.8",
					},
				},
				UpdateOld: []*endpoint.Endpoint{},
				Delete: []*endpoint.Endpoint{
					{
						DNSName: "example.org",
						Target:  "8.8.8.8",
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
						DNSName: "foo.bar.org",
						Target:  "5.5.5.5",
					},
				},
			},
			expectedZonesState: map[string]zone{
				"org": {
					"example.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.org",
								Target:  "8.8.8.8",
							},
							Type: defaultType,
						},
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.org",
							},
							Type: "TXT",
						},
					},
					"foo.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "foo.org",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
						},
					},
					"foo.bar.org": []*InMemoryRecord{},
				},
				"com": {
					"example.com": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.com",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
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
						DNSName: "foo.bar.new.org",
						Target:  "4.8.8.9",
					},
				},
				UpdateNew: []*endpoint.Endpoint{
					{
						DNSName: "foo.bar.org",
						Target:  "4.8.8.4",
					},
				},
				UpdateOld: []*endpoint.Endpoint{
					{
						DNSName: "foo.bar.org",
						Target:  "5.5.5.5",
					},
				},
				Delete: []*endpoint.Endpoint{
					{
						DNSName: "example.org",
						Target:  "8.8.8.8",
					},
				},
			},
			expectedZonesState: map[string]zone{
				"org": {
					"example.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.org",
							},
							Type: "TXT",
						},
					},
					"foo.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "foo.org",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
						},
					},
					"foo.bar.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "foo.bar.org",
								Target:  "4.8.8.4",
							},
							Type: defaultType,
						},
					},
					"foo.bar.new.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "foo.bar.new.org",
								Target:  "4.8.8.9",
							},
							Type: defaultType,
						},
					},
				},
				"com": {
					"example.com": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.com",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
						},
					},
				},
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			init := map[string]zone{
				"org": {
					"example.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.org",
								Target:  "8.8.8.8",
							},
							Type: defaultType,
						},
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.org",
							},
							Type: "TXT",
						},
					},
					"foo.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "foo.org",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
						},
					},
					"foo.bar.org": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "foo.bar.org",
								Target:  "5.5.5.5",
							},
							Type: defaultType,
						},
					},
				},
				"com": {
					"example.com": []*InMemoryRecord{
						{
							Endpoint: &endpoint.Endpoint{
								DNSName: "example.com",
								Target:  "4.4.4.4",
							},
							Type: "CNAME",
						},
					},
				},
			}
			im := &InMemoryProvider{
				zones: init,
			}
			err := im.ApplyChanges(ti.zone, ti.changes)
			if ti.expectError && err == nil {
				t.Errorf("should return an error")
			}
			if !ti.expectError && err != nil {
				t.Error(err)
			}
			if !ti.expectError {
				if !reflect.DeepEqual(im.zones, ti.expectedZonesState) {
					t.Errorf("invalid update")
				}
			}
		})
	}
}

func testNewInMemoryProvider(t *testing.T) {
	cfg := NewInMemoryProvider()
	if cfg.zones == nil {
		t.Error("nil map")
	}
}

func testInMemoryCreateZone(t *testing.T) {
	im := NewInMemoryProvider()
	if err := im.CreateZone("zone"); err != nil {
		t.Error(err)
	}
	if err := im.CreateZone("zone"); err != ErrZoneAlreadyExists {
		t.Errorf("should fail with zone already exists")
	}
}
