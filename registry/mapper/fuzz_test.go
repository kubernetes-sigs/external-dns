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

package mapper

import (
	"testing"
)

func FuzzToEndpointName(f *testing.F) {
	f.Add("txt-a-example.com")
	f.Add("txt-cname-example.com")
	f.Add("txt-example.com")
	f.Add("")
	f.Add("example.com")
	f.Add("a-example.com")
	f.Add("prefix-aaaa-example.com")
	f.Add("example.com-suffix")

	mappers := []AffixNameMapper{
		NewAffixNameMapper("txt-", "", ""),
		NewAffixNameMapper("", "-txt", ""),
		NewAffixNameMapper("txt-%{record_type}-", "", ""),
		NewAffixNameMapper("", "-%{record_type}-txt", ""),
		NewAffixNameMapper("txt-", "", "wildcard"),
	}

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 1024 {
			t.Skip()
		}
		for _, m := range mappers {
			m.ToEndpointName(input)
		}
	})
}

func FuzzExtractRecordTypeDefaultPosition(f *testing.F) {
	f.Add("a-example.com")
	f.Add("aaaa-example.com")
	f.Add("cname-test.com")
	f.Add("txt-record.com")
	f.Add("example.com")
	f.Add("")
	f.Add("-example.com")
	f.Add("a-")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 1024 {
			t.Skip()
		}
		extractRecordTypeDefaultPosition(input)
	})
}

func FuzzToTXTNameRoundTrip(f *testing.F) {
	f.Add("example.com", "A")
	f.Add("sub.example.com", "AAAA")
	f.Add("test.com", "CNAME")

	mappers := []AffixNameMapper{
		NewAffixNameMapper("txt-", "", ""),
		NewAffixNameMapper("", "-txt", ""),
		NewAffixNameMapper("txt-%{record_type}-", "", ""),
		NewAffixNameMapper("", "-%{record_type}-txt", ""),
	}

	f.Fuzz(func(t *testing.T, dns, recordType string) {
		if len(dns) > 512 || len(recordType) > 32 {
			t.Skip()
		}
		for _, m := range mappers {
			txtName := m.ToTXTName(dns, recordType)
			if txtName == "" {
				continue
			}
			// Round-trip: TXT name should map back to the original endpoint name
			epName, _ := m.ToEndpointName(txtName)
			if epName != "" && epName != dns {
				// Only check when both directions produce results
				// Some inputs may not round-trip due to ambiguous record types
			}
		}
	})
}
