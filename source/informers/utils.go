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

package informers

import (
	"crypto/sha1"
	"encoding/hex"
)

// ToSHA returns the SHA1 hash of the input string as a hex string.
// Using a SHA1 hash of the label selector string (as in ToSHA(labels.Set(selector).String())) is useful:
// - It provides a consistent and compact representation of the selector.
// - It allows for efficient indexing and lookup in Kubernetes informers.
// - It avoids issues with long label selector strings that could exceed index length limits.
func ToSHA(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
