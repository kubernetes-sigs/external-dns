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

package provider

import (
	"net"
	"strings"
	"testing"
)

func FuzzEnsureTrailingDot(f *testing.F) {
	f.Add("example.com")
	f.Add("example.com.")
	f.Add("1.2.3.4")
	f.Add("::1")
	f.Add("2001:db8::1")
	f.Add("")
	f.Add(".")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 1024 {
			t.Skip()
		}
		result := EnsureTrailingDot(input)
		// Idempotency: applying twice should give the same result
		result2 := EnsureTrailingDot(result)
		if result != result2 {
			t.Errorf("EnsureTrailingDot is not idempotent: %q -> %q -> %q", input, result, result2)
		}
		// Non-IP results should end with "."
		if net.ParseIP(input) == nil && !strings.HasSuffix(result, ".") {
			t.Errorf("EnsureTrailingDot(%q) = %q should end with dot", input, result)
		}
	})
}
