/*
Copyright 2023 The Kubernetes Authors.

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
	"fmt"
	"testing"
)

func TestCidrToInAddr(t *testing.T) {
	var tests = []struct {
		in     string
		out    string
		errmsg string
	}{

		{"174.136.107.0/24", "107.136.174.in-addr.arpa", ""},
		{"174.136.107.1/24", "107.136.174.in-addr.arpa", "CIDR 174.136.107.1/24 has 1 bits beyond the mask"},

		{"174.136.0.0/16", "136.174.in-addr.arpa", ""},
		{"174.136.43.0/16", "136.174.in-addr.arpa", "CIDR 174.136.43.0/16 has 1 bits beyond the mask"},

		{"174.0.0.0/8", "174.in-addr.arpa", ""},
		{"174.136.43.0/8", "174.in-addr.arpa", "CIDR 174.136.43.0/8 has 1 bits beyond the mask"},
		{"174.136.0.44/8", "174.in-addr.arpa", "CIDR 174.136.0.44/8 has 1 bits beyond the mask"},
		{"174.136.45.45/8", "174.in-addr.arpa", "CIDR 174.136.45.45/8 has 1 bits beyond the mask"},

		{"2001::/16", "1.0.0.2.ip6.arpa", ""},
		{"2001:0db8:0123:4567:89ab:cdef:1234:5670/124", "7.6.5.4.3.2.1.f.e.d.c.b.a.9.8.7.6.5.4.3.2.1.0.8.b.d.0.1.0.0.2.ip6.arpa", ""},

		{"174.136.107.14/32", "14.107.136.174.in-addr.arpa", ""},
		{"2001:0db8:0123:4567:89ab:cdef:1234:5678/128", "8.7.6.5.4.3.2.1.f.e.d.c.b.a.9.8.7.6.5.4.3.2.1.0.8.b.d.0.1.0.0.2.ip6.arpa", ""},

		// IPv4 "Classless in-addr.arpa delegation" RFC2317.
		// From examples in the RFC:
		{"192.0.2.0/25", "0/25.2.0.192.in-addr.arpa", ""},
		{"192.0.2.128/26", "128/26.2.0.192.in-addr.arpa", ""},
		{"192.0.2.192/26", "192/26.2.0.192.in-addr.arpa", ""},
		// All the base cases:
		{"174.1.0.0/25", "0/25.0.1.174.in-addr.arpa", ""},
		{"174.1.0.0/26", "0/26.0.1.174.in-addr.arpa", ""},
		{"174.1.0.0/27", "0/27.0.1.174.in-addr.arpa", ""},
		{"174.1.0.0/28", "0/28.0.1.174.in-addr.arpa", ""},
		{"174.1.0.0/29", "0/29.0.1.174.in-addr.arpa", ""},
		{"174.1.0.0/30", "0/30.0.1.174.in-addr.arpa", ""},
		{"174.1.0.0/31", "0/31.0.1.174.in-addr.arpa", ""},
		// /25 (all cases)
		{"174.1.0.0/25", "0/25.0.1.174.in-addr.arpa", ""},
		{"174.1.0.128/25", "128/25.0.1.174.in-addr.arpa", ""},
		// /26 (all cases)
		{"174.1.0.0/26", "0/26.0.1.174.in-addr.arpa", ""},
		{"174.1.0.64/26", "64/26.0.1.174.in-addr.arpa", ""},
		{"174.1.0.128/26", "128/26.0.1.174.in-addr.arpa", ""},
		{"174.1.0.192/26", "192/26.0.1.174.in-addr.arpa", ""},
		// /27 (all cases)
		{"174.1.0.0/27", "0/27.0.1.174.in-addr.arpa", ""},
		{"174.1.0.32/27", "32/27.0.1.174.in-addr.arpa", ""},
		{"174.1.0.64/27", "64/27.0.1.174.in-addr.arpa", ""},
		{"174.1.0.96/27", "96/27.0.1.174.in-addr.arpa", ""},
		{"174.1.0.128/27", "128/27.0.1.174.in-addr.arpa", ""},
		{"174.1.0.160/27", "160/27.0.1.174.in-addr.arpa", ""},
		{"174.1.0.192/27", "192/27.0.1.174.in-addr.arpa", ""},
		{"174.1.0.224/27", "224/27.0.1.174.in-addr.arpa", ""},
		// /28 (first 2, last 2)
		{"174.1.0.0/28", "0/28.0.1.174.in-addr.arpa", ""},
		{"174.1.0.16/28", "16/28.0.1.174.in-addr.arpa", ""},
		{"174.1.0.224/28", "224/28.0.1.174.in-addr.arpa", ""},
		{"174.1.0.240/28", "240/28.0.1.174.in-addr.arpa", ""},
		// /29 (first 2 cases)
		{"174.1.0.0/29", "0/29.0.1.174.in-addr.arpa", ""},
		{"174.1.0.8/29", "8/29.0.1.174.in-addr.arpa", ""},
		// /30 (first 2 cases)
		{"174.1.0.0/30", "0/30.0.1.174.in-addr.arpa", ""},
		{"174.1.0.4/30", "4/30.0.1.174.in-addr.arpa", ""},
		// /31 (first 2 cases)
		{"174.1.0.0/31", "0/31.0.1.174.in-addr.arpa", ""},
		{"174.1.0.2/31", "2/31.0.1.174.in-addr.arpa", ""},

		// IPv4-mapped IPv6 addresses:
		{"::ffff:174.136.107.15", "15.107.136.174.in-addr.arpa", ""},

		// Error Cases:
		{"0.0.0.0/0", "", "cannot use /0 in reverse CIDR"},
		{"2001::/0", "", "CIDR 2001::/0 has 1 bits beyond the mask"},
		{"4.5/16", "", "invalid CIDR address: 4.5/16"},
		{"foo.com", "", "invalid CIDR address: foo.com"},
	}
	for i, tst := range tests {
		t.Run(fmt.Sprintf("%d--%s", i, tst.in), func(t *testing.T) {
			d, err := CidrToInAddr(tst.in)

			if tst.errmsg == "" {
				// We DO NOT expect an error.
				if err != nil {
					// ...but we got one.
					t.Errorf("Expected '%s' but got ERROR('%s')", tst.out, err)
				} else if (tst.errmsg == "") && d != tst.out {
					// but the expected output was wrong
					t.Errorf("Expected '%s' but got '%s'", tst.out, d)
				}
			} else {
				// We DO expect an error.
				if err == nil {
					// ...but we didn't get one.
					t.Errorf("Expected ERROR('%s') but got result '%s'", tst.errmsg, d)
				} else if err.Error() != tst.errmsg {
					// ...but not the right error.
					t.Errorf("Expected ERROR('%s') but got ERROR('%s')", tst.errmsg, err)
				}

			}
		})
	}
}
