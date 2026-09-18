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

package source

import (
	"testing"
)

func FuzzParseIngress(f *testing.F) {
	f.Add("default/nginx")
	f.Add("nginx")
	f.Add("")
	f.Add("a/b/c")
	f.Add("/")
	f.Add("namespace/")
	f.Add("/name")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 1024 {
			t.Skip()
		}
		_, _, _ = ParseIngress(input)
	})
}
