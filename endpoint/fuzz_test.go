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
	"encoding/base64"
	"encoding/json"
	"testing"
)

func FuzzNewLabelsFromStringPlain(f *testing.F) {
	f.Add("heritage=external-dns,external-dns/owner=foo")
	f.Add("heritage=external-dns,external-dns/owner=foo,external-dns/resource=bar")
	f.Add("heritage=other-tool")
	f.Add("")
	f.Add("no-equals-sign")
	f.Add(",,,")
	f.Add("heritage=external-dns")
	f.Add(`"heritage=external-dns,external-dns/owner=test"`)

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 4096 {
			t.Skip()
		}
		_, _ = NewLabelsFromStringPlain(input)
	})
}

func FuzzNewMXRecord(f *testing.F) {
	f.Add("10 mail.example.com")
	f.Add("0 .")
	f.Add("65535 host")
	f.Add("")
	f.Add("notanumber host")
	f.Add("10")
	f.Add("10 a b")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 1024 {
			t.Skip()
		}
		mx, err := NewMXRecord(input)
		if err == nil {
			p := mx.GetPriority()
			if p == nil {
				t.Error("priority should not be nil on success")
			}
			h := mx.GetHost()
			if h == nil || *h == "" {
				t.Error("host should not be empty on success")
			}
		}
	})
}

func FuzzValidateSRVRecord(f *testing.F) {
	f.Add("10 5 5060 example.com.")
	f.Add("0 0 0 .")
	f.Add("")
	f.Add("10 5 5060 example.com")
	f.Add("notanum 5 5060 example.com.")
	f.Add("10 5 99999 example.com.")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 1024 {
			t.Skip()
		}
		Targets{input}.ValidateSRVRecord()
	})
}

func FuzzSuitableType(f *testing.F) {
	f.Add("1.2.3.4")
	f.Add("::1")
	f.Add("2001:db8::1")
	f.Add("example.com")
	f.Add("")
	f.Add("999.999.999.999")
	f.Add("::ffff:192.0.2.1")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 1024 {
			t.Skip()
		}
		result := SuitableType(input)
		switch result {
		case RecordTypeA, RecordTypeAAAA, RecordTypeCNAME:
			// valid
		default:
			t.Errorf("SuitableType returned unexpected type: %q", result)
		}
	})
}

func FuzzDomainFilterMatch(f *testing.F) {
	f.Add("example.com")
	f.Add("sub.example.com")
	f.Add("")
	f.Add(".")
	f.Add("example.com.")
	f.Add("münchen.de")
	f.Add("xn--mnchen-3ya.de")

	filter := NewDomainFilterWithExclusions([]string{"example.com"}, []string{"excluded.example.com"})

	f.Fuzz(func(t *testing.T, domain string) {
		if len(domain) > 1024 {
			t.Skip()
		}
		filter.Match(domain)
		filter.MatchParent(domain)
	})
}

func FuzzDomainFilterUnmarshalJSON(f *testing.F) {
	f.Add([]byte(`{"Include":["example.com"]}`))
	f.Add([]byte(`{}`))
	f.Add([]byte(`{"Include":["example.com"],"Exclude":["sub.example.com"]}`))
	f.Add([]byte(`{"regexInclude":"^.*\\.example\\.com$"}`))
	f.Add([]byte(`{"regexExclude":"^internal\\."}`))
	f.Add([]byte(`{"Include":["example.com"],"regexInclude":"x"}`))
	f.Add([]byte(`""`))
	f.Add([]byte(`null`))

	f.Fuzz(func(t *testing.T, input []byte) {
		if len(input) > 4096 {
			t.Skip()
		}
		var df DomainFilter
		_ = json.Unmarshal(input, &df)
	})
}

func FuzzDecryptText(f *testing.F) {
	f.Add("randomgarbage")
	f.Add("")
	f.Add(base64.StdEncoding.EncodeToString([]byte("short")))
	f.Add(base64.StdEncoding.EncodeToString(make([]byte, 20)))

	fixedKey := []byte("0123456789abcdef0123456789abcdef") // 32 bytes

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 4096 {
			t.Skip()
		}
		_, _, _ = DecryptText(input, fixedKey)
	})
}

func FuzzDecryptTextRoundTrip(f *testing.F) {
	f.Add("heritage=external-dns,external-dns/owner=foo")
	f.Add("hello world")
	f.Add("")
	f.Add("a")

	fixedKey := []byte("0123456789abcdef0123456789abcdef")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > 4096 {
			t.Skip()
		}
		nonce, err := GenerateNonce()
		if err != nil {
			t.Skip()
		}
		encrypted, err := EncryptText(input, fixedKey, nonce)
		if err != nil {
			return
		}
		decrypted, _, err := DecryptText(encrypted, fixedKey)
		if err != nil {
			t.Fatalf("failed to decrypt round-trip: %v", err)
		}
		if decrypted != input {
			t.Fatalf("round-trip mismatch: got %q, want %q", decrypted, input)
		}
	})
}
