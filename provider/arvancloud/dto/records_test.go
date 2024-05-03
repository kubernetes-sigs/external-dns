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

	"github.com/stretchr/testify/assert"
)

func TestArvanCloud_Uint16(t *testing.T) {
	type obj struct {
		Value Uint16 `json:"value"`
	}
	type given struct {
		v    *obj
		data string
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		given    given
		name     string
		wantsErr expectedErr
		wants    obj
	}{
		{
			name: "Should successfully unmarshal uint16 when value is null",
			given: given{
				data: fmt.Sprintf("{%q:%s}", "value", "null"),
				v:    &obj{},
			},
			wants: obj{
				Value: 0,
			},
		},
		{
			name: "Should successfully unmarshal uint16 when value contains double quote (a number string)",
			given: given{
				data: fmt.Sprintf("{%q:%q}", "value", "1"),
				v:    &obj{},
			},
			wants: obj{
				Value: 1,
			},
		},
		{
			name: "Should successfully unmarshal uint16 when value is number",
			given: given{
				data: fmt.Sprintf("{%q:%d}", "value", 2),
				v:    &obj{},
			},
			wants: obj{
				Value: 2,
			},
		},
		{
			name: "Should error unmarshal uint16 when value is not number",
			given: given{
				data: fmt.Sprintf("{%q:%q}", "value", "invalid"),
				v:    &obj{},
			},
			wantsErr: expectedErr{
				happened: true,
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

func TestArvanCloud_ARecordMarshalRec(t *testing.T) {
	type given struct {
		v *ARecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal A record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal A record and return empty if record is empty",
			given: given{
				v: &ARecord{},
			},
			wants: "",
		},
		{
			name: "Should successfully marshal A record and return IP address",
			given: given{
				v: &ARecord{IP: "192.168.1.1"},
			},
			wants: "192.168.1.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_AAAARecordMarshalRec(t *testing.T) {
	type given struct {
		v *AAAARecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal AAAA record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal AAAA record and return empty if record is empty",
			given: given{
				v: &AAAARecord{},
			},
			wants: "",
		},
		{
			name: "Should successfully marshal AAAA record and return IP address",
			given: given{
				v: &AAAARecord{IP: "2001:db8:3333:4444:5555:6666:7777:8888"},
			},
			wants: "2001:db8:3333:4444:5555:6666:7777:8888",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_NSRecordMarshalRec(t *testing.T) {
	type given struct {
		v *NSRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal NS record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal NS record and return empty if record is empty",
			given: given{
				v: &NSRecord{},
			},
			wants: "",
		},
		{
			name: "Should successfully marshal NS record and return host",
			given: given{
				v: &NSRecord{Host: "ns.example.com"},
			},
			wants: "ns.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_TXTRecordMarshalRec(t *testing.T) {
	type given struct {
		v *TXTRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal TXT record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal TXT record and return empty if record is empty",
			given: given{
				v: &TXTRecord{},
			},
			wants: "",
		},
		{
			name: "Should successfully marshal TXT record",
			given: given{
				v: &TXTRecord{Text: "this-is-a-text-record"},
			},
			wants: "this-is-a-text-record",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_CNAMERecordMarshalRec(t *testing.T) {
	type given struct {
		v *CNAMERecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal CNAME record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal CNAME record and return empty if record is empty",
			given: given{
				v: &CNAMERecord{},
			},
			wants: "",
		},
		{
			name: "Should successfully marshal CNAME record and return host",
			given: given{
				v: &CNAMERecord{Host: "is an alias of example.com"},
			},
			wants: "is an alias of example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_ANAMERecordMarshalRec(t *testing.T) {
	type given struct {
		v *ANAMERecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal ANAME record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal ANAME record and return empty if record is empty",
			given: given{
				v: &ANAMERecord{},
			},
			wants: "",
		},
		{
			name: "Should successfully marshal ANAME record and return location",
			given: given{
				v: &ANAMERecord{Location: "example.com"},
			},
			wants: "example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_MXRecordMarshalRec(t *testing.T) {
	type given struct {
		v *MXRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal MX record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal MX record and return empty if record is empty",
			given: given{
				v: &MXRecord{},
			},
			wants: "0 ",
		},
		{
			name: "Should successfully marshal MX record and return host and priority",
			given: given{
				v: &MXRecord{Host: "mx.example.com", Priority: 10},
			},
			wants: "10 mx.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_MXRecordUnmarshalRec(t *testing.T) {
	type given struct {
		data []byte
		v    MXRecord
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
		wants    MXRecord
	}{
		{
			name: "Should error unmarshal MX record when data is invalid length",
			given: given{
				data: []byte{},
				v:    MXRecord{},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &ProviderError{},
				action:   ParseMXRecordActErr,
			},
		},
		{
			name: "Should error unmarshal MX record when parse priority",
			given: given{
				data: []byte("invalid mx.example.com"),
				v:    MXRecord{},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "Should successfully unmarshal MX record",
			given: given{
				data: []byte("10 mx.example.com"),
				v:    MXRecord{},
			},
			wants: MXRecord{
				Host:     "mx.example.com",
				Priority: 10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.given.v.UnmarshalRec(tt.given.data)

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

func TestArvanCloud_SRVRecordMarshalRec(t *testing.T) {
	type given struct {
		v *SRVRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal SRV record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal SRV record and return empty if record is empty",
			given: given{
				v: &SRVRecord{},
			},
			wants: "0 0 0 ",
		},
		{
			name: "Should successfully marshal SRV record and return optional priority and optional weight and port and target",
			given: given{
				v: &SRVRecord{Target: "server.example.com", Port: 80},
			},
			wants: "0 0 80 server.example.com",
		},
		{
			name: "Should successfully marshal SRV record and return priority and weight and port and target",
			given: given{
				v: &SRVRecord{Target: "server.example.com", Port: 80, Priority: 10, Weight: 20},
			},
			wants: "10 20 80 server.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_SRVRecordUnmarshalRec(t *testing.T) {
	type given struct {
		data []byte
		v    SRVRecord
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
		wants    SRVRecord
	}{
		{
			name: "Should error unmarshal SRV record when data is invalid length",
			given: given{
				data: []byte{},
				v:    SRVRecord{},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &ProviderError{},
				action:   ParseSRVRecordActErr,
			},
		},
		{
			name: "Should error unmarshal SRV record when parse priority",
			given: given{
				data: []byte("invalid 20 80 server.example.com"),
				v:    SRVRecord{},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "Should error unmarshal SRV record when parse weight",
			given: given{
				data: []byte("10 invalid 80 server.example.com"),
				v:    SRVRecord{},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "Should error unmarshal SRV record when parse port",
			given: given{
				data: []byte("10 20 invalid server.example.com"),
				v:    SRVRecord{},
			},
			wantsErr: expectedErr{
				happened: true,
			},
		},
		{
			name: "Should successfully unmarshal SRV record",
			given: given{
				data: []byte("10 20 80 server.example.com"),
				v:    SRVRecord{},
			},
			wants: SRVRecord{
				Target:   "server.example.com",
				Port:     80,
				Priority: 10,
				Weight:   20,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.given.v.UnmarshalRec(tt.given.data)

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

func TestArvanCloud_SPFRecordMarshalRec(t *testing.T) {
	type given struct {
		v *SPFRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal SPF record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal SPF record",
			given: given{
				v: &SPFRecord{Text: "v=spf1 include:_spf.google.com ~all"},
			},
			wants: "v=spf1 include:_spf.google.com ~all",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_DKIMRecordMarshalRec(t *testing.T) {
	type given struct {
		v *DKIMRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal DKIM record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal DKIM record and return priority and weight and port and target",
			given: given{
				v: &DKIMRecord{Text: "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE"},
			},
			wants: "v=DKIM1; p=76E629F05F70 9EF665853333 EEC3F5ADE69A 2362BECE4065 8267AB2FC3CB 6CBE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_PTRRecordMarshalRec(t *testing.T) {
	type given struct {
		v *PTRRecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal PTR record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal PTR record and return priority and weight and port and target",
			given: given{
				v: &PTRRecord{Domain: "example.com"},
			},
			wants: "example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_TLSARecordMarshalRec(t *testing.T) {
	type given struct {
		v *TLSARecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal TLSA record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal TLSA record and return empty if record is empty",
			given: given{
				v: &TLSARecord{},
			},
			wants: "   ",
		},
		{
			name: "Should successfully marshal TLSA record and return priority and weight and port and target",
			given: given{
				v: &TLSARecord{Usage: "3", Selector: "1", MatchingType: "1", Certificate: "0D6FCE13243AA7"},
			},
			wants: "3 1 1 0D6FCE13243AA7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_TLSARecordUnmarshalRec(t *testing.T) {
	type given struct {
		v    TLSARecord
		data []byte
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    TLSARecord
		wantsErr expectedErr
	}{
		{
			name: "Should error unmarshal TLSA record when data is invalid length",
			given: given{
				data: []byte{},
				v:    TLSARecord{},
			},
			wantsErr: expectedErr{
				happened: true,
				errType:  &ProviderError{},
				action:   ParseTLSARecordActErr,
			},
		},
		{
			name: "Should successfully unmarshal TLSA record when it contains certificate",
			given: given{
				data: []byte("3 1 1 0D6FCE13243AA7"),
				v:    TLSARecord{},
			},
			wants: TLSARecord{
				Usage:        "3",
				Selector:     "1",
				MatchingType: "1",
				Certificate:  "0D6FCE13243AA7",
			},
		},
		{
			name: "Should successfully unmarshal TLSA record when it doesn't contain certificate",
			given: given{
				data: []byte("3 1 1"),
				v:    TLSARecord{},
			},
			wants: TLSARecord{
				Usage:        "3",
				Selector:     "1",
				MatchingType: "1",
				Certificate:  "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.given.v.UnmarshalRec(tt.given.data)

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

func TestArvanCloud_CAARecordMarshalRec(t *testing.T) {
	type given struct {
		v *CAARecord
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    string
		wantsErr expectedErr
	}{
		{
			name: "Should successfully marshal CAA record and return empty if record is nil",
			given: given{
				v: nil,
			},
			wants: "",
		},
		{
			name: "Should successfully marshal CAA record and return empty if record is empty",
			given: given{
				v: &CAARecord{},
			},
			wants: " ",
		},
		{
			name: "Should successfully marshal CAA record and return priority and weight and port and target",
			given: given{
				v: &CAARecord{Value: strconv.Quote("letsencrypt.org"), Tag: "issue"},
			},
			wants: "issue \"letsencrypt.org\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.given.v.MarshalRec()

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

func TestArvanCloud_CAARecordUnmarshalRec(t *testing.T) {
	type given struct {
		v    CAARecord
		data []byte
	}
	type expectedErr struct {
		errType  interface{}
		action   string
		happened bool
	}
	tests := []struct {
		name     string
		given    given
		wants    CAARecord
		wantsErr expectedErr
	}{
		{
			name: "Should error unmarshalCAA record when data is invalid length",
			given: given{
				data: []byte{},
				v:    CAARecord{},
			},
			wantsErr: expectedErr{happened: true, errType: &ProviderError{}, action: ParseCAARecordActErr},
		},
		{
			name: "Should successfully unmarshal CAA record when tag and value are empty",
			given: given{
				data: []byte(" "),
				v:    CAARecord{},
			},
			wants: CAARecord{},
		},
		{
			name: "Should successfully unmarshal CAA record",
			given: given{
				data: []byte("issue \"letsencrypt.org\""),
				v:    CAARecord{},
			},
			wants: CAARecord{
				Value: strconv.Quote("letsencrypt.org"),
				Tag:   "issue",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.given.v.UnmarshalRec(tt.given.data)

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
