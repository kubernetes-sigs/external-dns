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
	"errors"
)

var errInvalidAESKey = errors.New("the AES Encryption key must be 32 bytes long, in either plain text or base64-encoded format")

// ParseAESKey accepts an unset key, a raw 32-byte key, or a base64-encoded 32-byte key.
func ParseAESKey(key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, nil
	}
	if len(key) == 32 {
		return key, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(string(key))
	if err != nil || len(decoded) != 32 {
		return nil, errInvalidAESKey
	}

	return decoded, nil
}
