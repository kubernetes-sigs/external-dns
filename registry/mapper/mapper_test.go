/*
Copyright 2025 The Kubernetes Authors.

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

package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
)

var (
	_ NameMapper = AffixNameMapper{}
)

func TestAffixNameMapper_ToEndpointName(t *testing.T) {
	tests := []struct {
		name             string
		mapper           AffixNameMapper
		input            string
		wantEndpointName string
		wantRecordType   string
	}{
		{
			name:             "prefix with A record type in affix",
			mapper:           NewAffixNameMapper("%{record_type}-", "", ""),
			input:            "a-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeA,
		},
		{
			name:             "prefix with AAAA record type in affix",
			mapper:           NewAffixNameMapper("%{record_type}-", "", ""),
			input:            "aaaa-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeAAAA,
		},
		{
			name:             "prefix with CNAME record type in affix",
			mapper:           NewAffixNameMapper("%{record_type}-", "", ""),
			input:            "cname-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeCNAME,
		},
		{
			name:             "prefix with NS record type in affix",
			mapper:           NewAffixNameMapper("%{record_type}-", "", ""),
			input:            "ns-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeNS,
		},
		{
			name:             "prefix with MX record type in affix",
			mapper:           NewAffixNameMapper("%{record_type}-", "", ""),
			input:            "mx-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeMX,
		},
		{
			name:             "prefix with SRV record type in affix",
			mapper:           NewAffixNameMapper("%{record_type}-", "", ""),
			input:            "srv-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeSRV,
		},
		{
			name:             "prefix with NAPTR record type in affix",
			mapper:           NewAffixNameMapper("%{record_type}-", "", ""),
			input:            "naptr-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeNAPTR,
		},
		{
			name:             "suffix with A record type in affix",
			mapper:           NewAffixNameMapper("", "-%{record_type}", ""),
			input:            "foo-a.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeA,
		},
		{
			name:             "suffix with CNAME record type in affix",
			mapper:           NewAffixNameMapper("", "-%{record_type}", ""),
			input:            "foo-cname.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeCNAME,
		},
		{
			name:             "no affix with A record",
			mapper:           NewAffixNameMapper("", "", ""),
			input:            "a-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeA,
		},
		{
			name:             "no affix with AAAA record",
			mapper:           NewAffixNameMapper("", "", ""),
			input:            "aaaa-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeAAAA,
		},
		{
			name:             "no affix with CNAME record",
			mapper:           NewAffixNameMapper("", "", ""),
			input:            "cname-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeCNAME,
		},
		{
			name:             "no affix with NS record",
			mapper:           NewAffixNameMapper("", "", ""),
			input:            "ns-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeNS,
		},
		{
			name:             "no affix with MX record",
			mapper:           NewAffixNameMapper("", "", ""),
			input:            "mx-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeMX,
		},
		{
			name:             "no affix with SRV record",
			mapper:           NewAffixNameMapper("", "", ""),
			input:            "srv-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeSRV,
		},
		{
			name:             "no affix with NAPTR record",
			mapper:           NewAffixNameMapper("", "", ""),
			input:            "naptr-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   endpoint.RecordTypeNAPTR,
		},
		{
			name:             "suffix with txt record",
			mapper:           NewAffixNameMapper("", "", ""),
			input:            "txt-foo.example.com",
			wantEndpointName: "txt-foo.example.com",
			wantRecordType:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotType := tt.mapper.ToEndpointName(tt.input)
			assert.Equal(t, tt.wantEndpointName, gotName)
			assert.Equal(t, tt.wantRecordType, gotType)
		})
	}

	// Verify all supported records are tested
	testedRecords := make(map[string]bool)
	for _, tt := range tests {
		if tt.wantRecordType != "" {
			testedRecords[tt.wantRecordType] = true
		}
	}

	for _, recordType := range supportedRecords {
		assert.True(t, testedRecords[recordType], "Record type %s is in supportedRecords but not tested in TestAffixNameMapper_ToEndpointName", recordType)
	}
}

func TestAffixNameMapper_ToTXTName(t *testing.T) {
	tests := []struct {
		name        string
		mapper      AffixNameMapper
		dns         string
		recordType  string
		wantTXTName string
	}{
		{
			name:        "prefix with A record type in affix",
			mapper:      NewAffixNameMapper("%{record_type}-", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeA,
			wantTXTName: "a-foo.example.com",
		},
		{
			name:        "prefix with AAAA record type in affix",
			mapper:      NewAffixNameMapper("%{record_type}-", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeAAAA,
			wantTXTName: "aaaa-foo.example.com",
		},
		{
			name:        "prefix with CNAME record type in affix",
			mapper:      NewAffixNameMapper("%{record_type}-", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeCNAME,
			wantTXTName: "cname-foo.example.com",
		},
		{
			name:        "prefix with NS record type in affix",
			mapper:      NewAffixNameMapper("%{record_type}-", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeNS,
			wantTXTName: "ns-foo.example.com",
		},
		{
			name:        "prefix with MX record type in affix",
			mapper:      NewAffixNameMapper("%{record_type}-", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeMX,
			wantTXTName: "mx-foo.example.com",
		},
		{
			name:        "prefix with SRV record type in affix",
			mapper:      NewAffixNameMapper("%{record_type}-", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeSRV,
			wantTXTName: "srv-foo.example.com",
		},
		{
			name:        "prefix with NAPTR record type in affix",
			mapper:      NewAffixNameMapper("%{record_type}-", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeNAPTR,
			wantTXTName: "naptr-foo.example.com",
		},
		{
			name:        "suffix with A record type in affix",
			mapper:      NewAffixNameMapper("", "-%{record_type}", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeA,
			wantTXTName: "foo-a.example.com",
		},
		{
			name:        "suffix with CNAME record type in affix",
			mapper:      NewAffixNameMapper("", "-%{record_type}", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeCNAME,
			wantTXTName: "foo-cname.example.com",
		},
		{
			name:        "wildcard replacement with A record",
			mapper:      NewAffixNameMapper("txt-", "", "wild"),
			dns:         "*.example.com",
			recordType:  endpoint.RecordTypeA,
			wantTXTName: "txt-a-wild.example.com",
		},
		{
			name:        "wildcard replacement with MX record",
			mapper:      NewAffixNameMapper("txt-", "", "wild"),
			dns:         "*.example.com",
			recordType:  endpoint.RecordTypeMX,
			wantTXTName: "txt-mx-wild.example.com",
		},
		{
			name:        "no affix with A record",
			mapper:      NewAffixNameMapper("", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeA,
			wantTXTName: "a-foo.example.com",
		},
		{
			name:        "no affix with AAAA record",
			mapper:      NewAffixNameMapper("", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeAAAA,
			wantTXTName: "aaaa-foo.example.com",
		},
		{
			name:        "no affix with CNAME record",
			mapper:      NewAffixNameMapper("", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeCNAME,
			wantTXTName: "cname-foo.example.com",
		},
		{
			name:        "no affix with NS record",
			mapper:      NewAffixNameMapper("", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeNS,
			wantTXTName: "ns-foo.example.com",
		},
		{
			name:        "no affix with MX record",
			mapper:      NewAffixNameMapper("", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeMX,
			wantTXTName: "mx-foo.example.com",
		},
		{
			name:        "no affix with SRV record",
			mapper:      NewAffixNameMapper("", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeSRV,
			wantTXTName: "srv-foo.example.com",
		},
		{
			name:        "no affix with NAPTR record",
			mapper:      NewAffixNameMapper("", "", ""),
			dns:         "foo.example.com",
			recordType:  endpoint.RecordTypeNAPTR,
			wantTXTName: "naptr-foo.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mapper.ToTXTName(tt.dns, tt.recordType)
			assert.Equal(t, tt.wantTXTName, got)
		})
	}

	// Verify all supported records are tested
	testedRecords := make(map[string]bool)
	for _, tt := range tests {
		testedRecords[tt.recordType] = true
	}

	for _, recordType := range supportedRecords {
		assert.True(t, testedRecords[recordType], "Record type %s is in supportedRecords but not tested in TestAffixNameMapper_ToTXTName", recordType)
	}
}

func TestAffixNameMapper_RecordTypeInAffix(t *testing.T) {
	tests := []struct {
		name   string
		mapper AffixNameMapper
		want   bool
	}{
		{
			name:   "prefix contains record type",
			mapper: NewAffixNameMapper("%{record_type}-", "", ""),
			want:   true,
		},
		{
			name:   "suffix contains record type",
			mapper: NewAffixNameMapper("", "-%{record_type}", ""),
			want:   true,
		},
		{
			name:   "no record type in affix",
			mapper: NewAffixNameMapper("txt-", "-txt", ""),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mapper.RecordTypeInAffix()
			assert.Equal(t, tt.want, got)
		})
	}
}
