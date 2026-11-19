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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCAARecord(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		wantFlags uint8
		wantTag   string
		wantValue string
		wantStr   string
	}{
		{"quoted", `0 issue "letsencrypt.org"`, 0, "issue", "letsencrypt.org", `0 issue "letsencrypt.org"`},
		{"unquoted", "0 issue letsencrypt.org", 0, "issue", "letsencrypt.org", `0 issue "letsencrypt.org"`},
		{"iodef mailto", `0 iodef "mailto:security@example.com"`, 0, "iodef", "mailto:security@example.com", `0 iodef "mailto:security@example.com"`},
		{"critical flag", `128 issue "ca.example.net"`, 128, "issue", "ca.example.net", `128 issue "ca.example.net"`},
		{"parameters with spaces", `0 issue "letsencrypt.org; validationmethods=dns-01"`, 0, "issue", "letsencrypt.org; validationmethods=dns-01", `0 issue "letsencrypt.org; validationmethods=dns-01"`},
		{"unquoted parameters keep the whole rest", "0 issue letsencrypt.org; validationmethods=dns-01", 0, "issue", "letsencrypt.org; validationmethods=dns-01", `0 issue "letsencrypt.org; validationmethods=dns-01"`},
		{"empty value forbids all issuers", `0 issue ""`, 0, "issue", "", `0 issue ""`},
		{"semicolon value", `0 issuewild ";"`, 0, "issuewild", ";", `0 issuewild ";"`},
		{"extra whitespace", "  0 \t issue \t \"letsencrypt.org\"  ", 0, "issue", "letsencrypt.org", `0 issue "letsencrypt.org"`},
		{"trailing dot kept", "0 issue letsencrypt.org.", 0, "issue", "letsencrypt.org.", `0 issue "letsencrypt.org."`},
		{"longest tag", `0 ` + strings.Repeat("a", 255) + ` "ca.example.net"`, 0, strings.Repeat("a", 255), "ca.example.net", `0 ` + strings.Repeat("a", 255) + ` "ca.example.net"`},
		{"value over 255 octets", `0 issue "` + strings.Repeat("a", 300) + `"`, 0, "issue", strings.Repeat("a", 300), `0 issue "` + strings.Repeat("a", 300) + `"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			caa, err := NewCAARecord(tt.target)
			require.NoError(t, err)
			assert.Equal(t, tt.wantFlags, caa.GetFlags())
			assert.Equal(t, tt.wantTag, caa.GetTag())
			assert.Equal(t, tt.wantValue, caa.GetValue())
			assert.Equal(t, tt.wantStr, caa.String())

			again, err := NewCAARecord(caa.String())
			require.NoError(t, err)
			assert.Equal(t, caa, again, "String must round-trip through NewCAARecord")
		})
	}
}

func TestNewCAARecordInvalid(t *testing.T) {
	tests := []struct {
		name    string
		target  string
		wantErr string
	}{
		{"empty", "", "expected <flags> <tag> <value>"},
		{"missing value", "0 issue", "expected <flags> <tag> <value>"},
		{"non-numeric flags", `x issue "ca.example.net"`, "flags"},
		{"flags out of range", `256 issue "ca.example.net"`, "flags"},
		{"tag with dash", `0 bad-tag "ca.example.net"`, "tag"},
		{"tag too long", `0 ` + strings.Repeat("a", 256) + ` "ca.example.net"`, "tag"},
		{"unterminated quote", `0 issue "ca.example.net`, "unterminated quote"},
		{"lone quote", `0 issue "`, "unterminated quote"},
		{"embedded quote", `0 issue "ca"example.net"`, `must not contain`},
		{"backslash", `0 issue "ca\046example.net"`, `must not contain`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCAARecord(tt.target)
			require.ErrorContains(t, err, tt.wantErr)
		})
	}
}

func TestValidateCAAEndpoint(t *testing.T) {
	require.NoError(t, NewEndpoint("example.com", RecordTypeCAA, `0 issue "letsencrypt.org"`).Validate())
	require.Error(t, NewEndpoint("example.com", RecordTypeCAA, "0 issue").Validate())
}

func TestNewEndpointKeepsCAATrailingDot(t *testing.T) {
	ep := NewEndpoint("example.com", RecordTypeCAA, "0 issue letsencrypt.org.")
	assert.Equal(t, Targets{"0 issue letsencrypt.org."}, ep.Targets)
}
