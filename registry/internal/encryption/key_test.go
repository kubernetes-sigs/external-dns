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

package encryption

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseAESKey(t *testing.T) {
	rawKey := []byte("passphrasewhichneedstobe32bytes!")
	encodedKey := []byte(base64.StdEncoding.EncodeToString(rawKey))
	encoded32ByteKey := []byte("12345678901234567890123456789012")

	tests := []struct {
		name    string
		input   []byte
		want    []byte
		wantErr bool
	}{
		{name: "nil key"},
		{name: "empty key", input: []byte{}},
		{name: "raw key", input: rawKey, want: rawKey},
		{name: "base64 key", input: encodedKey, want: rawKey},
		{name: "32-byte key takes precedence over base64 decoding", input: encoded32ByteKey, want: encoded32ByteKey},
		{name: "invalid base64", input: []byte("not-base64!"), wantErr: true},
		{name: "decoded key has invalid length", input: []byte(base64.StdEncoding.EncodeToString([]byte("short"))), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAESKey(tt.input)
			if tt.wantErr {
				require.EqualError(t, err, "the AES Encryption key must be 32 bytes long, in either plain text or base64-encoded format")
				assert.Nil(t, got)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
