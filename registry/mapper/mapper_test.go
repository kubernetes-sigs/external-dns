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
			name:             "prefix with record type in affix",
			mapper:           NewAffixNameMapper("%{record_type}-", "", ""),
			input:            "a-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   "A",
		},
		{
			name:             "suffix with record type in affix",
			mapper:           NewAffixNameMapper("", "-%{record_type}", ""),
			input:            "foo.example.com-a",
			wantEndpointName: "foo.example.com",
			wantRecordType:   "A",
		},
		{
			name:             "prefix without record type in affix",
			mapper:           NewAffixNameMapper("txt-", "", ""),
			input:            "a-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   "A",
		},
		{
			name:             "suffix without record type in affix",
			mapper:           NewAffixNameMapper("", "-txt", ""),
			input:            "foo.example.com-a",
			wantEndpointName: "foo.example.com",
			wantRecordType:   "A",
		},
		{
			name:             "no affix",
			mapper:           NewAffixNameMapper("", "", ""),
			input:            "a-foo.example.com",
			wantEndpointName: "foo.example.com",
			wantRecordType:   "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotType := tt.mapper.ToEndpointName(tt.input)
			assert.Equal(t, tt.wantEndpointName, gotName)
			assert.Equal(t, tt.wantRecordType, gotType)
		})
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
			name:        "prefix with record type in affix",
			mapper:      NewAffixNameMapper("%{record_type}-", "", ""),
			dns:         "foo.example.com",
			recordType:  "A",
			wantTXTName: "a-foo.example.com",
		},
		{
			name:        "suffix with record type in affix",
			mapper:      NewAffixNameMapper("", "-%{record_type}", ""),
			dns:         "foo.example.com",
			recordType:  "A",
			wantTXTName: "foo-a.example.com",
		},
		{
			name:        "wildcard replacement",
			mapper:      NewAffixNameMapper("txt-", "", "wild"),
			dns:         "*.example.com",
			recordType:  "A",
			wantTXTName: "txt-a-wild.example.com",
		},
		{
			name:        "no affix",
			mapper:      NewAffixNameMapper("", "", ""),
			dns:         "foo.example.com",
			recordType:  "A",
			wantTXTName: "a-foo.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mapper.ToTXTName(tt.dns, tt.recordType)
			assert.Equal(t, tt.wantTXTName, got)
		})
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
