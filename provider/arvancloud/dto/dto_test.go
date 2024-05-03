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

package dto

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestArvanCloud_DnsRecord_UnmarshalJSON(t *testing.T) {
	type given struct {
		v    *DnsRecord
		data string
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		given    given
		wantsErr expectedErr
		name     string
		wants    DnsRecord
	}{
		{
			name: "should successfully parse empty json with default values",
			given: given{
				data: "{}",
				v:    &DnsRecord{},
			},
		},
		{
			name: "should successfully parse data without record type",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: %q, %q: %q, %q: %q, %q: %d, %q: %t, %q: %t}",
					"created_at", "2024-04-14T01:02:03+00:00",
					"updated_at", "2024-04-14T01:02:03+00:00",
					"ID", "00000000-0000-0000-0000-000000000000",
					"name", "record",
					"ttl", 120,
					"cloud", true,
					"is_protected", true,
				),
				v: &DnsRecord{},
			},
			wants: DnsRecord{
				CreatedAt: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
					return t
				}(),
				UpdateAt: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
					return t
				}(),
				ID:          "00000000-0000-0000-0000-000000000000",
				Name:        "record",
				TTL:         120,
				Cloud:       true,
				IsProtected: true,
			},
		},
		{
			name: "should error parse json when type record exist but not match with expected records",
			given: given{
				data: fmt.Sprintf("{%q: %q, %q: %q}", "type", "unknown", "value", ""),
				v:    &DnsRecord{},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &ProviderError{},
				action:   ParseRecordActErr,
			},
		},
		{
			name: "should error parse json when A record is object",
			given: given{
				data: fmt.Sprintf("{%q: %q, %q: {}}", "type", "a", "value"),
				v:    &DnsRecord{},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when A record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}, {%q: %q}]}",
					"type", "a",
					"value",
					"ip", "192.168.1.1",
					"ip", "192.168.1.2",
				),
				v: &DnsRecord{
					Type: AType,
				},
			},
			wants: DnsRecord{
				Type:     AType,
				Value:    []ARecord{{IP: "192.168.1.1"}, {IP: "192.168.1.2"}},
				Contents: []string{"192.168.1.1", "192.168.1.2"},
			},
		},
		{
			name: "should successfully parse json when A record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}, {%q: %q}]}",
					"type", "A",
					"value",
					"ip", "192.168.1.1",
					"ip", "192.168.1.2",
				),
				v: &DnsRecord{
					Type: AType,
				},
			},
			wants: DnsRecord{
				Type:     AType,
				Value:    []ARecord{{IP: "192.168.1.1"}, {IP: "192.168.1.2"}},
				Contents: []string{"192.168.1.1", "192.168.1.2"},
			},
		},
		{
			name: "should error is object parse json when AAAA record",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {}}",
					"type", "aaaa",
					"value",
				),
				v: &DnsRecord{
					Type: AAAAType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when AAAA record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}, {%q: %q}]}",
					"type", "aaaa",
					"value",
					"ip", "2001:db8:3333:4444:5555:6666:7777:8888",
					"ip", "2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF",
				),
				v: &DnsRecord{
					Type: AAAAType,
				},
			},
			wants: DnsRecord{
				Type: AAAAType,
				Value: []AAAARecord{
					{IP: "2001:db8:3333:4444:5555:6666:7777:8888"},
					{IP: "2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF"},
				},
				Contents: []string{"2001:db8:3333:4444:5555:6666:7777:8888", "2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF"},
			},
		},
		{
			name: "should successfully parse json when AAAA record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}, {%q: %q}]}",
					"type", "AAAA",
					"value",
					"ip", "2001:db8:3333:4444:5555:6666:7777:8888",
					"ip", "2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF",
				),
				v: &DnsRecord{
					Type: AAAAType,
				},
			},
			wants: DnsRecord{
				Type: AAAAType,
				Value: []AAAARecord{
					{IP: "2001:db8:3333:4444:5555:6666:7777:8888"},
					{IP: "2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF"},
				},
				Contents: []string{"2001:db8:3333:4444:5555:6666:7777:8888", "2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF"},
			},
		},
		{
			name: "should successfully parse json when AAAA record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}, {%q: %q}]}",
					"type", "AaAa",
					"value",
					"ip", "2001:db8:3333:4444:5555:6666:7777:8888",
					"ip", "2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF",
				),
				v: &DnsRecord{
					Type: AAAAType,
				},
			},
			wants: DnsRecord{
				Type: AAAAType,
				Value: []AAAARecord{
					{IP: "2001:db8:3333:4444:5555:6666:7777:8888"},
					{IP: "2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF"},
				},
				Contents: []string{"2001:db8:3333:4444:5555:6666:7777:8888", "2001:db8:3333:4444:CCCC:DDDD:EEEE:FFFF"},
			},
		},
		{
			name: "should error parse json when NS record in not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}]}",
					"type", "ns",
					"value",
					"host", "ns.example.com",
				),
				v: &DnsRecord{
					Type: NSType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when NS record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "ns",
					"value",
					"host", "ns.example.com",
				),
				v: &DnsRecord{
					Type: NSType,
				},
			},
			wants: DnsRecord{
				Type:     NSType,
				Value:    NSRecord{Host: "ns.example.com"},
				Contents: []string{"ns.example.com"},
			},
		},
		{
			name: "should successfully parse json when NS record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "NS",
					"value",
					"host", "ns.example.com",
				),
				v: &DnsRecord{
					Type: NSType,
				},
			},
			wants: DnsRecord{
				Type:     NSType,
				Value:    NSRecord{Host: "ns.example.com"},
				Contents: []string{"ns.example.com"},
			},
		},
		{
			name: "should successfully parse json when NS record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "nS",
					"value",
					"host", "ns.example.com",
				),
				v: &DnsRecord{
					Type: NSType,
				},
			},
			wants: DnsRecord{
				Type:     NSType,
				Value:    NSRecord{Host: "ns.example.com"},
				Contents: []string{"ns.example.com"},
			},
		},
		{
			name: "should error parse json when TXT record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}]}",
					"type", "TXT",
					"value",
					"text", "this-is-a-text-record",
				),
				v: &DnsRecord{
					Type: TXTType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when TXT record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "txt",
					"value",
					"text", "this-is-a-text-record",
				),
				v: &DnsRecord{
					Type: TXTType,
				},
			},
			wants: DnsRecord{
				Type:     TXTType,
				Value:    TXTRecord{Text: "this-is-a-text-record"},
				Contents: []string{"this-is-a-text-record"},
			},
		},
		{
			name: "should successfully parse json when TXT record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "TXT",
					"value",
					"text", "this-is-a-text-record",
				),
				v: &DnsRecord{
					Type: TXTType,
				},
			},
			wants: DnsRecord{
				Type:     TXTType,
				Value:    TXTRecord{Text: "this-is-a-text-record"},
				Contents: []string{"this-is-a-text-record"},
			},
		},
		{
			name: "should successfully parse json when TXT record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "TxT",
					"value",
					"text", "this-is-a-text-record",
				),
				v: &DnsRecord{
					Type: TXTType,
				},
			},
			wants: DnsRecord{
				Type:     TXTType,
				Value:    TXTRecord{Text: "this-is-a-text-record"},
				Contents: []string{"this-is-a-text-record"},
			},
		},
		{
			name: "should error parse json when CNAME record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}]}",
					"type", "cname",
					"value",
					"host", "is an alias of example.com",
				),
				v: &DnsRecord{
					Type: CNAMEType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when CNAME record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "cname",
					"value",
					"host", "is an alias of example.com",
				),
				v: &DnsRecord{
					Type: CNAMEType,
				},
			},
			wants: DnsRecord{
				Type:     CNAMEType,
				Value:    CNAMERecord{Host: "is an alias of example.com"},
				Contents: []string{"is an alias of example.com"},
			},
		},
		{
			name: "should successfully parse json when CNAME record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "CNAME",
					"value",
					"host", "is an alias of example.com",
				),
				v: &DnsRecord{
					Type: CNAMEType,
				},
			},
			wants: DnsRecord{
				Type:     CNAMEType,
				Value:    CNAMERecord{Host: "is an alias of example.com"},
				Contents: []string{"is an alias of example.com"},
			},
		},
		{
			name: "should successfully parse json when CNAME record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "cNaMe",
					"value",
					"host", "is an alias of example.com",
				),
				v: &DnsRecord{
					Type: CNAMEType,
				},
			},
			wants: DnsRecord{
				Type:     CNAMEType,
				Value:    CNAMERecord{Host: "is an alias of example.com"},
				Contents: []string{"is an alias of example.com"},
			},
		},
		{
			name: "should error parse json when ANAME record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}]}",
					"type", "aname",
					"value",
					"location", "example.com",
				),
				v: &DnsRecord{
					Type: ANAMEType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when ANAME record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "aname",
					"value",
					"location", "example.com",
				),
				v: &DnsRecord{
					Type: ANAMEType,
				},
			},
			wants: DnsRecord{
				Type:     ANAMEType,
				Value:    ANAMERecord{Location: "example.com"},
				Contents: []string{"example.com"},
			},
		},
		{
			name: "should successfully parse json when ANAME record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "ANAME",
					"value",
					"location", "example.com",
				),
				v: &DnsRecord{
					Type: ANAMEType,
				},
			},
			wants: DnsRecord{
				Type:     ANAMEType,
				Value:    ANAMERecord{Location: "example.com"},
				Contents: []string{"example.com"},
			},
		},
		{
			name: "should successfully parse json when ANAME record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "AnAmE",
					"value",
					"location", "example.com",
				),
				v: &DnsRecord{
					Type: ANAMEType,
				},
			},
			wants: DnsRecord{
				Type:     ANAMEType,
				Value:    ANAMERecord{Location: "example.com"},
				Contents: []string{"example.com"},
			},
		},
		{
			name: "should error parse json when MX record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q, %q: %d}]}",
					"type", "mx",
					"value",
					"host", "mx.example.com",
					"priority", 10,
				),
				v: &DnsRecord{
					Type: MXType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when MX record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %d}}",
					"type", "mx",
					"value",
					"host", "mx.example.com",
					"priority", 10,
				),
				v: &DnsRecord{
					Type: MXType,
				},
			},
			wants: DnsRecord{
				Type:     MXType,
				Value:    MXRecord{Host: "mx.example.com", Priority: 10},
				Contents: []string{"10 mx.example.com"},
			},
		},
		{
			name: "should successfully parse json when MX record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %d}}",
					"type", "MX",
					"value",
					"host", "mx.example.com",
					"priority", 10,
				),
				v: &DnsRecord{
					Type: MXType,
				},
			},
			wants: DnsRecord{
				Type:     MXType,
				Value:    MXRecord{Host: "mx.example.com", Priority: 10},
				Contents: []string{"10 mx.example.com"},
			},
		},
		{
			name: "should successfully parse json when MX record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %d}}",
					"type", "mX",
					"value",
					"host", "mx.example.com",
					"priority", 10,
				),
				v: &DnsRecord{
					Type: MXType,
				},
			},
			wants: DnsRecord{
				Type:     MXType,
				Value:    MXRecord{Host: "mx.example.com", Priority: 10},
				Contents: []string{"10 mx.example.com"},
			},
		},
		{
			name: "should error parse json when SRV record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q, %q: %d}]}",
					"type", "srv",
					"value",
					"target", "server.example.com",
					"port", 80,
				),
				v: &DnsRecord{
					Type: SRVType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when SRV record (lowercase error)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %d}}",
					"type", "srv",
					"value",
					"target", "server.example.com",
					"port", 80,
				),
				v: &DnsRecord{
					Type: SRVType,
				},
			},
			wants: DnsRecord{
				Type:     SRVType,
				Value:    SRVRecord{Target: "server.example.com", Port: 80},
				Contents: []string{"0 0 80 server.example.com"},
			},
		},
		{
			name: "should successfully parse json when SRV record (uppercase error)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %d}}",
					"type", "SRV",
					"value",
					"target", "server.example.com",
					"port", 80,
				),
				v: &DnsRecord{
					Type: SRVType,
				},
			},
			wants: DnsRecord{
				Type:     SRVType,
				Value:    SRVRecord{Target: "server.example.com", Port: 80},
				Contents: []string{"0 0 80 server.example.com"},
			},
		},
		{
			name: "should successfully parse json when SRV record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %d}}",
					"type", "SrV",
					"value",
					"target", "server.example.com",
					"port", 80,
				),
				v: &DnsRecord{
					Type: SRVType,
				},
			},
			wants: DnsRecord{
				Type:     SRVType,
				Value:    SRVRecord{Target: "server.example.com", Port: 80},
				Contents: []string{"0 0 80 server.example.com"},
			},
		},
		{
			name: "should error parse json when SPF record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}]}",
					"type", "spf",
					"value",
					"text", "v=spf1 include:_spf.example.com -all",
				),
				v: &DnsRecord{
					Type: SPFType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when SPF record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "spf",
					"value",
					"text", "v=spf1 include:_spf.example.com -all",
				),
				v: &DnsRecord{
					Type: SPFType,
				},
			},
			wants: DnsRecord{
				Type:     SPFType,
				Value:    SPFRecord{Text: "v=spf1 include:_spf.example.com -all"},
				Contents: []string{"v=spf1 include:_spf.example.com -all"},
			},
		},
		{
			name: "should successfully parse json when SPF record (uppercase error)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "spf",
					"value",
					"text", "v=spf1 include:_spf.example.com -all",
				),
				v: &DnsRecord{
					Type: SPFType,
				},
			},
			wants: DnsRecord{
				Type:     SPFType,
				Value:    SPFRecord{Text: "v=spf1 include:_spf.example.com -all"},
				Contents: []string{"v=spf1 include:_spf.example.com -all"},
			},
		},
		{
			name: "should successfully parse json when SPF record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "sPf",
					"value",
					"text", "v=spf1 include:_spf.example.com -all",
				),
				v: &DnsRecord{
					Type: SPFType,
				},
			},
			wants: DnsRecord{
				Type:     SPFType,
				Value:    SPFRecord{Text: "v=spf1 include:_spf.example.com -all"},
				Contents: []string{"v=spf1 include:_spf.example.com -all"},
			},
		},
		{
			name: "should error parse json when DKIM record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}]}",
					"type", "dkim",
					"value",
					"text", "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE",
				),
				v: &DnsRecord{
					Type: DKIMType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when DKIM record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "dkim",
					"value",
					"text", "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE",
				),
				v: &DnsRecord{
					Type: DKIMType,
				},
			},
			wants: DnsRecord{
				Type:     DKIMType,
				Value:    DKIMRecord{Text: "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE"},
				Contents: []string{"v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE"},
			},
		},
		{
			name: "should successfully parse json when DKIM record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "dkim",
					"value",
					"text", "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE",
				),
				v: &DnsRecord{
					Type: DKIMType,
				},
			},
			wants: DnsRecord{
				Type:     DKIMType,
				Value:    DKIMRecord{Text: "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE"},
				Contents: []string{"v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE"},
			},
		},
		{
			name: "should successfully parse json when DKIM record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "DkIm",
					"value",
					"text", "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE",
				),
				v: &DnsRecord{
					Type: DKIMType,
				},
			},
			wants: DnsRecord{
				Type:     DKIMType,
				Value:    DKIMRecord{Text: "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE"},
				Contents: []string{"v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE"},
			},
		},
		{
			name: "should error parse json when PTR record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q}]}",
					"type", "ptr",
					"value",
					"domain", "example.com",
				),
				v: &DnsRecord{
					Type: PTRType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when PTR record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "ptr",
					"value",
					"domain", "example.com",
				),
				v: &DnsRecord{
					Type: PTRType,
				},
			},
			wants: DnsRecord{
				Type:     PTRType,
				Value:    PTRRecord{Domain: "example.com"},
				Contents: []string{"example.com"},
			},
		},
		{
			name: "should successfully parse json when PTR record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "ptr",
					"value",
					"domain", "example.com",
				),
				v: &DnsRecord{
					Type: PTRType,
				},
			},
			wants: DnsRecord{
				Type:     PTRType,
				Value:    PTRRecord{Domain: "example.com"},
				Contents: []string{"example.com"},
			},
		},
		{
			name: "should successfully parse json when PTR record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q}}",
					"type", "pTr",
					"value",
					"domain", "example.com",
				),
				v: &DnsRecord{
					Type: PTRType,
				},
			},
			wants: DnsRecord{
				Type:     PTRType,
				Value:    PTRRecord{Domain: "example.com"},
				Contents: []string{"example.com"},
			},
		},
		{
			name: "should error parse json when TLSA record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q, %q: %q, %q: %q, %q: %q}]}",
					"type", "tlsa",
					"value",
					"usage", "3",
					"selector", "1",
					"matching_type", "1",
					"certificate", "0D6FCE13243AA7",
				),
				v: &DnsRecord{
					Type: TLSAType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when TLSA record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %q, %q: %q, %q: %q}}",
					"type", "tlsa",
					"value",
					"usage", "3",
					"selector", "1",
					"matching_type", "1",
					"certificate", "0D6FCE13243AA7",
				),
				v: &DnsRecord{
					Type: TLSAType,
				},
			},
			wants: DnsRecord{
				Type: TLSAType,
				Value: TLSARecord{
					Usage:        "3",
					Selector:     "1",
					MatchingType: "1",
					Certificate:  "0D6FCE13243AA7",
				},
				Contents: []string{"3 1 1 0D6FCE13243AA7"},
			},
		},
		{
			name: "should successfully parse json when TLSA record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %q, %q: %q, %q: %q}}",
					"type", "tlsa",
					"value",
					"usage", "3",
					"selector", "1",
					"matching_type", "1",
					"certificate", "0D6FCE13243AA7",
				),
				v: &DnsRecord{
					Type: TLSAType,
				},
			},
			wants: DnsRecord{
				Type: TLSAType,
				Value: TLSARecord{
					Usage:        "3",
					Selector:     "1",
					MatchingType: "1",
					Certificate:  "0D6FCE13243AA7",
				},
				Contents: []string{"3 1 1 0D6FCE13243AA7"},
			},
		},
		{
			name: "should successfully parse json when TLSA record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %q, %q: %q, %q: %q}}",
					"type", "tlsa",
					"value",
					"usage", "3",
					"selector", "1",
					"matching_type", "1",
					"certificate", "0D6FCE13243AA7",
				),
				v: &DnsRecord{
					Type: TLSAType,
				},
			},
			wants: DnsRecord{
				Type: TLSAType,
				Value: TLSARecord{
					Usage:        "3",
					Selector:     "1",
					MatchingType: "1",
					Certificate:  "0D6FCE13243AA7",
				},
				Contents: []string{"3 1 1 0D6FCE13243AA7"},
			},
		},
		{
			name: "should error parse json when CAA record is not object",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: [{%q: %q, %q: %q}]}",
					"type", "caa",
					"value",
					"value", strconv.Quote("letsencrypt.org"),
					"tag", "issue",
				),
				v: &DnsRecord{
					Type: CAAType,
				},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "should successfully parse json when CAA record (lowercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %q}}",
					"type", "caa",
					"value",
					"value", strconv.Quote("letsencrypt.org"),
					"tag", "issue",
				),
				v: &DnsRecord{
					Type: CAAType,
				},
			},
			wants: DnsRecord{
				Type:     CAAType,
				Value:    CAARecord{Value: strconv.Quote("letsencrypt.org"), Tag: "issue"},
				Contents: []string{`issue "letsencrypt.org"`},
			},
		},
		{
			name: "should successfully parse json when CAA record (uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %q}}",
					"type", "caa",
					"value",
					"value", strconv.Quote("letsencrypt.org"),
					"tag", "issue",
				),
				v: &DnsRecord{
					Type: CAAType,
				},
			},
			wants: DnsRecord{
				Type:     CAAType,
				Value:    CAARecord{Value: strconv.Quote("letsencrypt.org"), Tag: "issue"},
				Contents: []string{`issue "letsencrypt.org"`},
			},
		},
		{
			name: "should successfully parse json when CAA record (random lowercase and uppercase record)",
			given: given{
				data: fmt.Sprintf(
					"{%q: %q, %q: {%q: %q, %q: %q}}",
					"type", "caa",
					"value",
					"value", strconv.Quote("letsencrypt.org"),
					"tag", "issue",
				),
				v: &DnsRecord{
					Type: CAAType,
				},
			},
			wants: DnsRecord{
				Type:     CAAType,
				Value:    CAARecord{Value: strconv.Quote("letsencrypt.org"), Tag: "issue"},
				Contents: []string{`issue "letsencrypt.org"`},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(tt.given.data), tt.given.v)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.wants, *tt.given.v)
		})
	}
}

func TestArvanCloud_DnsRecord_UnmarshalJSON2(t *testing.T) {
	type given struct {
		data string
		v    DnsRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		wantsErr expectedErr
		name     string
		given    given
		wants    DnsRecord
	}{
		{
			name: "should error parse invalid json",
			given: given{
				data: "invalid-json",
				v:    DnsRecord{},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.given.v.UnmarshalJSON([]byte(tt.given.data))

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.wants, tt.given.v)
		})
	}
}

func TestArvanCloud_DnsRecord_MarshalJSON(t *testing.T) {
	type given struct {
		v DnsRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}

	tests := []struct {
		wantsErr expectedErr
		name     string
		wants    string
		given    given
	}{
		{
			name: "should successfully marshal dns record",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:null,%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "",
				"value",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal empty A record",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        AType,
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:null,%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "A",
				"value",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal A record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        AType,
					Value:       []ARecord{{IP: "192.168.1.1"}},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:[{%q:%q}],%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "A",
				"value",
				"ip", "192.168.1.1",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal A record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        AType,
					Contents:    []string{""},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:[{%q:%q}],%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "A",
				"value",
				"ip", "",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal AAAA record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        AAAAType,
					Value:       []AAAARecord{{IP: "2001:db8:3333:4444:5555:6666:7777:8888"}},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:[{%q:%q}],%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "AAAA",
				"value",
				"ip", "2001:db8:3333:4444:5555:6666:7777:8888",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal AAAA record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        AAAAType,
					Contents:    []string{"2001:db8:3333:4444:5555:6666:7777:8888"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:[{%q:%q}],%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "AAAA",
				"value",
				"ip", "2001:db8:3333:4444:5555:6666:7777:8888",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal NS record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        NSType,
					Value:       NSRecord{Host: "ns.example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "NS",
				"value",
				"host", "ns.example.com",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal NS record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        NSType,
					Contents:    []string{"ns.example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "NS",
				"value",
				"host", "ns.example.com",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal TXT record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        TXTType,
					Value:       TXTRecord{Text: "this-is-a-text-record"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "TXT",
				"value",
				"text", "this-is-a-text-record",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal TXT record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        TXTType,
					Contents:    []string{"this-is-a-text-record"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "TXT",
				"value",
				"text", "this-is-a-text-record",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal TXT record with is filled contents (value is start and end with double quote)",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        TXTType,
					Contents:    []string{strconv.Quote("this-is-a-text-record")},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "TXT",
				"value",
				"text", "this-is-a-text-record",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal CNAME record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        CNAMEType,
					Value:       CNAMERecord{Host: "is an alias of example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "CNAME",
				"value",
				"host", "is an alias of example.com",
				"host_header", "",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal CNAME record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        CNAMEType,
					Contents:    []string{"is an alias of example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "CNAME",
				"value",
				"host", "is an alias of example.com",
				"host_header", "",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal ANAME record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        ANAMEType,
					Value:       ANAMERecord{Location: "example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "ANAME",
				"value",
				"location", "example.com",
				"host_header", "",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal ANAME record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        ANAMEType,
					Contents:    []string{"example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "ANAME",
				"value",
				"location", "example.com",
				"host_header", "",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal MX record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        MXType,
					Value:       MXRecord{Host: "mx.example.com", Priority: 10},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%d},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "MX",
				"value",
				"host", "mx.example.com",
				"priority", 10,
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should error marshal MX record with invalid contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        MXType,
					Contents:    []string{""},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &json.MarshalerError{},
			},
		},
		{
			name: "should successfully marshal MX record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        MXType,
					Contents:    []string{"10 mx.example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%d},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "MX",
				"value",
				"host", "mx.example.com",
				"priority", 10,
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal SRV record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        SRVType,
					Value:       SRVRecord{Target: "server.example.com", Port: 80},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%d},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "SRV",
				"value",
				"target", "server.example.com",
				"port", 80,
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should error marshal SRV record with invalid contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        SRVType,
					Contents:    []string{""},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &json.MarshalerError{},
			},
		},
		{
			name: "should successfully marshal SRV record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        SRVType,
					Contents:    []string{"0 0 80 server.example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%d},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "SRV",
				"value",
				"target", "server.example.com",
				"port", 80,
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal SRV record with is filled contents (all field is filled)",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        SRVType,
					Contents:    []string{"10 20 80 server.example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%d,%q:%d,%q:%d},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "SRV",
				"value",
				"target", "server.example.com",
				"port", 80,
				"weight", 20,
				"priority", 10,
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal SPF record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        SPFType,
					Value:       SPFRecord{Text: "v=spf1 include:_spf.google.com ~all"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "SPF",
				"value",
				"text", "v=spf1 include:_spf.google.com ~all",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal SPF record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        SPFType,
					Contents:    []string{"v=spf1 include:_spf.google.com ~all"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "SPF",
				"value",
				"text", "v=spf1 include:_spf.google.com ~all",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal DKIM record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        DKIMType,
					Value:       DKIMRecord{Text: "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "DKIM",
				"value",
				"text", "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal DKIM record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        DKIMType,
					Contents:    []string{"v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "DKIM",
				"value",
				"text", "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal PTR record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        PTRType,
					Value:       PTRRecord{Domain: "example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "PTR",
				"value",
				"domain", "example.com",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal PTR record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        PTRType,
					Contents:    []string{"example.com"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "PTR",
				"value",
				"domain", "example.com",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal TLSA record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        TLSAType,
					Value:       TLSARecord{Usage: "3", Selector: "1", MatchingType: "1", Certificate: "0D6FCE13243AA7"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%q,%q:%q,%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "TLSA",
				"value",
				"usage", "3",
				"selector", "1",
				"matching_type", "1",
				"certificate", "0D6FCE13243AA7",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal TLSA record with invalid contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        TLSAType,
					Contents:    []string{""},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &json.MarshalerError{},
			},
		},
		{
			name: "should successfully marshal TLSA record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        TLSAType,
					Contents:    []string{"3 1 1 0D6FCE13243AA7"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%q,%q:%q,%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "TLSA",
				"value",
				"usage", "3",
				"selector", "1",
				"matching_type", "1",
				"certificate", "0D6FCE13243AA7",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal TLSA record with is filled contents (without certificate)",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        TLSAType,
					Contents:    []string{"3 1 1"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%q,%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "TLSA",
				"value",
				"usage", "3",
				"selector", "1",
				"matching_type", "1",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal CAA record with is filled value",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        CAAType,
					Value:       CAARecord{Value: strconv.Quote("letsencrypt.org"), Tag: "issue"},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "CAA",
				"value",
				"value", strconv.Quote("letsencrypt.org"),
				"tag", "issue",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
		{
			name: "should successfully marshal CAA record with invalid contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        CAAType,
					Contents:    []string{""},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &json.MarshalerError{},
			},
		},
		{
			name: "should successfully marshal CAA record with is filled contents",
			given: given{
				v: DnsRecord{
					CreatedAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					UpdateAt: func() time.Time {
						t, _ := time.Parse(time.RFC3339, "2024-04-14T01:02:03+00:00")
						return t
					}(),
					ID:          "00000000-0000-0000-0000-000000000000",
					Name:        "record",
					Type:        CAAType,
					Contents:    []string{"issue \"letsencrypt.org\""},
					TTL:         120,
					Cloud:       true,
					IsProtected: true,
				},
			},
			wants: fmt.Sprintf(
				"{%q:%q,%q:%q,%q:%q,%q:%q,%q:%q,%q:{%q:%q,%q:%q},%q:%d,%q:%t,%q:%t}",
				"created_at", "2024-04-14T01:02:03Z",
				"updated_at", "2024-04-14T01:02:03Z",
				"id", "00000000-0000-0000-0000-000000000000",
				"name", "record",
				"type", "CAA",
				"value",
				"value", strconv.Quote("letsencrypt.org"),
				"tag", "issue",
				"ttl", 120,
				"cloud", true,
				"is_protected", true,
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.given.v)

			if tt.wantsErr.happened {
				assert.Error(t, err)
				if tt.wantsErr.errType != nil {
					assert.IsType(t, err, tt.wantsErr.errType)
					assert.ErrorAs(t, err, &tt.wantsErr.errType)
					if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&ProviderError{}) {
						assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*ProviderError).GetAction())
					}
				}
				return
			}
			assert.NoError(t, err)

			assert.Equal(t, tt.wants, string(data))
		})
	}
}

func helperPtrJsonRawMessage(text string) *json.RawMessage {
	data := json.RawMessage(text)

	return &data
}

func TestArvanCloud_DnsRecord_ValueParse(t *testing.T) {
	type given struct {
		ns       *parseDnsRecord
		valuePtr interface{}
		v        *DnsRecord
	}
	type expectedErr struct {
		errType interface{}
		action  string
	}

	tests := []struct {
		name     string
		given    given
		wantsErr expectedErr
	}{
		{
			name: "Should error if valuePtr is not pointer",
			given: given{
				ns:       &parseDnsRecord{},
				valuePtr: ARecord{},
				v:        &DnsRecord{},
			},
			wantsErr: expectedErr{
				errType: &ProviderError{},
				action:  NonPointerActErr,
			},
		},
		{
			name: "Should error if raw json is invalid",
			given: given{
				ns:       &parseDnsRecord{Value: helperPtrJsonRawMessage("invalid-json")},
				valuePtr: &ARecord{},
				v:        &DnsRecord{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.given.v.valueParse(tt.given.ns, tt.given.valuePtr)

			assert.Error(t, err)
			if tt.wantsErr.errType != nil {
				assert.IsType(t, err, tt.wantsErr.errType)
				assert.ErrorAs(t, err, &tt.wantsErr.errType)
				if reflect.TypeOf(tt.wantsErr.errType) == reflect.TypeOf(&ProviderError{}) {
					assert.Equal(t, tt.wantsErr.action, tt.wantsErr.errType.(*ProviderError).GetAction())
				}
			}
		})
	}
}
