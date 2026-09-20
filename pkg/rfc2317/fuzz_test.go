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

package rfc2317

import (
	"testing"
)

func FuzzCidrToInAddr(f *testing.F) {
	f.Add("10.20.30.0/24")
	f.Add("2001::/16")
	f.Add("1.2.3.4")
	f.Add("10.20.30.0/25")
	f.Add("10.20.30.128/25")
	f.Add("0.0.0.0/0")
	f.Add("10.20.30.1/24")
	f.Add("::/0")
	f.Add("2001:db8::/32")
	f.Add("::ffff:192.0.2.1")
	f.Add("")
	f.Add("not-a-cidr")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 256 {
			t.Skip()
		}
		_, _ = CidrToInAddr(input)
	})
}
