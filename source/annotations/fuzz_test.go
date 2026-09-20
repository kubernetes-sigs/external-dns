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

package annotations

import (
	"strings"
	"testing"
)

func FuzzParseTTL(f *testing.F) {
	f.Add("600")
	f.Add("10m")
	f.Add("1h30m")
	f.Add("1.5s")
	f.Add("")
	f.Add("abc")
	f.Add("0")
	f.Add("-1")
	f.Add("9999999999999")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 256 {
			t.Skip()
		}
		_, _ = parseTTL(input)
	})
}

func FuzzSplitHostnameAnnotation(f *testing.F) {
	f.Add("a.com,b.com")
	f.Add("")
	f.Add(" a.com , b.com ")
	f.Add(",")
	f.Add(",,,,")
	f.Add("single.host.com")
	f.Add("  ")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 4096 {
			t.Skip()
		}
		result := SplitHostnameAnnotation(input)
		for _, entry := range result {
			if strings.Contains(entry, " ") {
				t.Errorf("SplitHostnameAnnotation output should not contain spaces, got %q", entry)
			}
		}
	})
}

func FuzzParseFilter(f *testing.F) {
	f.Add("app=nginx")
	f.Add("env in (prod,staging)")
	f.Add("")
	f.Add("invalid(")
	f.Add("key!=value")
	f.Add("!key")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 1024 {
			t.Skip()
		}
		_, _ = ParseFilter(input)
	})
}
