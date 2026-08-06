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

const (
	sha256Digest = "54ee30e752c6e751b7bfde666e638ec4e25f3974207258072299caa15e15820f"
	sha512Digest = "f758ae635ec9b464b46d89e05274b791b72ac4245f8eb18ccad836ef59238706fa201cf1616169813a8d0743f3e0275eb69bad17bcf1032217830d0aacbbc843"
)

func TestNewTLSARecord(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		target           string
		wantUsage        uint8
		wantSelector     uint8
		wantMatchingType uint8
		wantCertificate  string
	}{
		{
			name:             "DANE-EE SPKI SHA-256",
			target:           "3 1 1 " + sha256Digest,
			wantUsage:        3,
			wantSelector:     1,
			wantMatchingType: 1,
			wantCertificate:  sha256Digest,
		},
		{
			name:             "DANE-TA FullCert SHA-512",
			target:           "2 0 2 " + sha512Digest,
			wantUsage:        2,
			wantSelector:     0,
			wantMatchingType: 2,
			wantCertificate:  sha512Digest,
		},
		{
			name:             "PKIX-TA",
			target:           "0 0 1 " + sha256Digest,
			wantUsage:        0,
			wantSelector:     0,
			wantMatchingType: 1,
			wantCertificate:  sha256Digest,
		},
		{
			name:             "uppercase hex is normalised",
			target:           "3 1 1 " + strings.ToUpper(sha256Digest),
			wantUsage:        3,
			wantSelector:     1,
			wantMatchingType: 1,
			wantCertificate:  sha256Digest,
		},
		{
			name:             "colon separators are stripped",
			target:           "3 1 1 54:ee:30:e7:52:c6:e7:51:b7:bf:de:66:6e:63:8e:c4:e2:5f:39:74:20:72:58:07:22:99:ca:a1:5e:15:82:0f",
			wantUsage:        3,
			wantSelector:     1,
			wantMatchingType: 1,
			wantCertificate:  sha256Digest,
		},
		{
			name:             "extra whitespace between fields",
			target:           "3   1  1   " + sha256Digest,
			wantUsage:        3,
			wantSelector:     1,
			wantMatchingType: 1,
			wantCertificate:  sha256Digest,
		},
		{
			name:             "matching type 0 accepts arbitrary length",
			target:           "3 1 0 deadbeef",
			wantUsage:        3,
			wantSelector:     1,
			wantMatchingType: 0,
			wantCertificate:  "deadbeef",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewTLSARecord(tc.target)
			require.NoError(t, err)
			assert.Equal(t, tc.wantUsage, got.GetUsage())
			assert.Equal(t, tc.wantSelector, got.GetSelector())
			assert.Equal(t, tc.wantMatchingType, got.GetMatchingType())
			assert.Equal(t, tc.wantCertificate, got.GetCertificate())
		})
	}
}

func TestNewTLSARecordErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		target string
	}{
		{"empty", ""},
		{"too few fields", "3 1 1"},
		{"too many fields", "3 1 1 " + sha256Digest + " extra"},
		{"non-numeric usage", "x 1 1 " + sha256Digest},
		{"usage out of range", "4 1 1 " + sha256Digest},
		{"selector out of range", "3 2 1 " + sha256Digest},
		{"matching type out of range", "3 1 3 " + sha256Digest},
		{"negative usage", "-1 1 1 " + sha256Digest},
		{"non-hex certificate", "3 1 1 " + strings.Repeat("z", 64)},
		{"sha256 digest wrong length", "3 1 1 deadbeef"},
		{"sha512 digest wrong length", "3 1 2 " + sha256Digest},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewTLSARecord(tc.target)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}

// A record written to a provider and read back must render identically,
// otherwise every reconciliation would see a difference and issue an update.
func TestTLSARecordRoundTripIsStable(t *testing.T) {
	t.Parallel()

	canonical := "3 1 1 " + sha256Digest

	for _, variant := range []string{
		canonical,
		"3 1 1 " + strings.ToUpper(sha256Digest),
		"3   1 1  " + sha256Digest,
	} {
		first, err := NewTLSARecord(variant)
		require.NoError(t, err)
		assert.Equal(t, canonical, first.String())

		second, err := NewTLSARecord(first.String())
		require.NoError(t, err)
		assert.Equal(t, canonical, second.String(), "re-parsing a rendered target must be idempotent")
	}
}

func TestTLSARecordTypeIsKnown(t *testing.T) {
	t.Parallel()
	assert.Contains(t, KnownRecordTypes, RecordTypeTLSA)
}
