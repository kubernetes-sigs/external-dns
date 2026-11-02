/*
Copyright 2026 The Kubernetes Authors.

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

package endpoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNAPTRRecord(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		target          string
		wantOrder       uint16
		wantPreference  uint16
		wantFlags       string
		wantService     string
		wantRegexp      string
		wantReplacement string
	}{
		{
			name:            "terminal rule with regexp and root replacement",
			target:          `100 50 "U" "E2U+sip" "!^.*$!sip:info@bar.com!" .`,
			wantOrder:       100,
			wantPreference:  50,
			wantFlags:       "U",
			wantService:     "E2U+sip",
			wantRegexp:      "!^.*$!sip:info@bar.com!",
			wantReplacement: ".",
		},
		{
			name:            "replacement rule with empty regexp",
			target:          `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`,
			wantOrder:       10,
			wantPreference:  20,
			wantFlags:       "S",
			wantService:     "SIP+D2U",
			wantRegexp:      "",
			wantReplacement: "_sip._udp.bar.com.",
		},
		{
			name:            "replacement without trailing dot is normalized",
			target:          `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com`,
			wantOrder:       10,
			wantPreference:  20,
			wantFlags:       "S",
			wantService:     "SIP+D2U",
			wantRegexp:      "",
			wantReplacement: "_sip._udp.bar.com.",
		},
		{
			name:            "multi-space separators",
			target:          `10  20  "S"  "SIP+D2U"  ""  _sip._udp.bar.com.`,
			wantOrder:       10,
			wantPreference:  20,
			wantFlags:       "S",
			wantService:     "SIP+D2U",
			wantRegexp:      "",
			wantReplacement: "_sip._udp.bar.com.",
		},
		{
			name:            "tab separators",
			target:          "10\t20\t\"S\"\t\"SIP+D2U\"\t\"\"\t_sip._udp.bar.com.",
			wantOrder:       10,
			wantPreference:  20,
			wantFlags:       "S",
			wantService:     "SIP+D2U",
			wantRegexp:      "",
			wantReplacement: "_sip._udp.bar.com.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			naptr, err := NewNAPTRRecord(tt.target)
			require.NoError(t, err)
			assert.Equal(t, tt.wantOrder, naptr.GetOrder())
			assert.Equal(t, tt.wantPreference, naptr.GetPreference())
			assert.Equal(t, tt.wantFlags, naptr.GetFlags())
			assert.Equal(t, tt.wantService, naptr.GetService())
			assert.Equal(t, tt.wantRegexp, naptr.GetRegexp())
			assert.Equal(t, tt.wantReplacement, naptr.GetReplacement())
		})
	}
}

func TestNewNAPTRRecordInvalid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		target string
	}{
		{name: "missing field", target: `10 "S" "SIP+D2U" "" _sip._udp.bar.com.`},
		{name: "non-numeric order", target: `abc 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`},
		{name: "order above uint16", target: `99999 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`},
		{name: "preference above uint16", target: `10 99999 "S" "SIP+D2U" "" _sip._udp.bar.com.`},
		{name: "empty target", target: ""},
		{name: "replacement with double trailing dot", target: `10 20 "S" "SIP+D2U" "" host..`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := NewNAPTRRecord(tt.target)
			require.Error(t, err)
			assert.Contains(t, err.Error(), "invalid NAPTR target")
		})
	}
}

func TestNAPTRTargetString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		target string
		want   string
	}{
		{
			name:   "canonical input round-trips unchanged",
			target: `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`,
			want:   `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`,
		},
		{
			name:   "multi-space separators are collapsed",
			target: `10  20  "S"  "SIP+D2U"  ""  _sip._udp.bar.com.`,
			want:   `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`,
		},
		{
			name:   "tab separators are collapsed",
			target: "10\t20\t\"S\"\t\"SIP+D2U\"\t\"\"\t_sip._udp.bar.com.",
			want:   `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`,
		},
		{
			name:   "missing trailing dot is added",
			target: `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com`,
			want:   `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`,
		},
		{
			name:   "regexp with spaces stays quoted",
			target: `100 50 "U" "E2U+sip" "!^.*$!sip:info@bar.com!" .`,
			want:   `100 50 "U" "E2U+sip" "!^.*$!sip:info@bar.com!" .`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			naptr, err := NewNAPTRRecord(tt.target)
			require.NoError(t, err)
			assert.Equal(t, tt.want, naptr.String())
		})
	}
}

func TestValidateNAPTRRecord(t *testing.T) {
	t.Parallel()

	valid := Targets{
		`10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`,
		`100 50 "U" "E2U+sip" "!^.*$!sip:info@bar.com!" .`,
	}
	assert.True(t, valid.ValidateNAPTRRecord())

	invalid := Targets{`10 20`}
	assert.False(t, invalid.ValidateNAPTRRecord())

	mixed := Targets{
		`10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`,
		`10 20`,
	}
	assert.False(t, mixed.ValidateNAPTRRecord())
}

func TestCheckEndpointNAPTR(t *testing.T) {
	t.Parallel()

	valid := &Endpoint{
		DNSName:    "bar.com",
		RecordType: RecordTypeNAPTR,
		Targets:    Targets{`10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`},
	}
	assert.True(t, valid.CheckEndpoint())

	invalid := &Endpoint{
		DNSName:    "bar.com",
		RecordType: RecordTypeNAPTR,
		Targets:    Targets{`10 20`},
	}
	assert.False(t, invalid.CheckEndpoint())
}
