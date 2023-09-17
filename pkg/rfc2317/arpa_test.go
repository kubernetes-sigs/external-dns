package rfc2317

import (
	"fmt"
	"testing"
)

func TestCidrToInAddr(t *testing.T) {
	var tests = []struct {
		in      string
		isError bool
		out     string
	}{
		{"174.136.107.0/24", false, "107.136.174.in-addr.arpa"},
		{"174.136.107.1/24", true, "107.136.174.in-addr.arpa"},

		{"174.136.0.0/16", false, "136.174.in-addr.arpa"},
		{"174.136.43.0/16", true, "136.174.in-addr.arpa"},

		{"174.0.0.0/8", false, "174.in-addr.arpa"},
		{"174.136.43.0/8", true, "174.in-addr.arpa"},
		{"174.136.0.44/8", true, "174.in-addr.arpa"},
		{"174.136.45.45/8", true, "174.in-addr.arpa"},

		{"2001::/16", false, "1.0.0.2.ip6.arpa"},
		{"2001:0db8:0123:4567:89ab:cdef:1234:5670/124", false, "7.6.5.4.3.2.1.f.e.d.c.b.a.9.8.7.6.5.4.3.2.1.0.8.b.d.0.1.0.0.2.ip6.arpa"},

		{"174.136.107.14/32", false, "14.107.136.174.in-addr.arpa"},
		{"2001:0db8:0123:4567:89ab:cdef:1234:5678/128", false, "8.7.6.5.4.3.2.1.f.e.d.c.b.a.9.8.7.6.5.4.3.2.1.0.8.b.d.0.1.0.0.2.ip6.arpa"},

		// IPv4 "Classless in-addr.arpa delegation" RFC2317.
		// From examples in the RFC:
		{"192.0.2.0/25", false, "0/25.2.0.192.in-addr.arpa"},
		{"192.0.2.128/26", false, "128/26.2.0.192.in-addr.arpa"},
		{"192.0.2.192/26", false, "192/26.2.0.192.in-addr.arpa"},
		// All the base cases:
		{"174.1.0.0/25", false, "0/25.0.1.174.in-addr.arpa"},
		{"174.1.0.0/26", false, "0/26.0.1.174.in-addr.arpa"},
		{"174.1.0.0/27", false, "0/27.0.1.174.in-addr.arpa"},
		{"174.1.0.0/28", false, "0/28.0.1.174.in-addr.arpa"},
		{"174.1.0.0/29", false, "0/29.0.1.174.in-addr.arpa"},
		{"174.1.0.0/30", false, "0/30.0.1.174.in-addr.arpa"},
		{"174.1.0.0/31", false, "0/31.0.1.174.in-addr.arpa"},
		// /25 (all cases)
		{"174.1.0.0/25", false, "0/25.0.1.174.in-addr.arpa"},
		{"174.1.0.128/25", false, "128/25.0.1.174.in-addr.arpa"},
		// /26 (all cases)
		{"174.1.0.0/26", false, "0/26.0.1.174.in-addr.arpa"},
		{"174.1.0.64/26", false, "64/26.0.1.174.in-addr.arpa"},
		{"174.1.0.128/26", false, "128/26.0.1.174.in-addr.arpa"},
		{"174.1.0.192/26", false, "192/26.0.1.174.in-addr.arpa"},
		// /27 (all cases)
		{"174.1.0.0/27", false, "0/27.0.1.174.in-addr.arpa"},
		{"174.1.0.32/27", false, "32/27.0.1.174.in-addr.arpa"},
		{"174.1.0.64/27", false, "64/27.0.1.174.in-addr.arpa"},
		{"174.1.0.96/27", false, "96/27.0.1.174.in-addr.arpa"},
		{"174.1.0.128/27", false, "128/27.0.1.174.in-addr.arpa"},
		{"174.1.0.160/27", false, "160/27.0.1.174.in-addr.arpa"},
		{"174.1.0.192/27", false, "192/27.0.1.174.in-addr.arpa"},
		{"174.1.0.224/27", false, "224/27.0.1.174.in-addr.arpa"},
		// /28 (first 2, last 2)
		{"174.1.0.0/28", false, "0/28.0.1.174.in-addr.arpa"},
		{"174.1.0.16/28", false, "16/28.0.1.174.in-addr.arpa"},
		{"174.1.0.224/28", false, "224/28.0.1.174.in-addr.arpa"},
		{"174.1.0.240/28", false, "240/28.0.1.174.in-addr.arpa"},
		// /29 (first 2 cases)
		{"174.1.0.0/29", false, "0/29.0.1.174.in-addr.arpa"},
		{"174.1.0.8/29", false, "8/29.0.1.174.in-addr.arpa"},
		// /30 (first 2 cases)
		{"174.1.0.0/30", false, "0/30.0.1.174.in-addr.arpa"},
		{"174.1.0.4/30", false, "4/30.0.1.174.in-addr.arpa"},
		// /31 (first 2 cases)
		{"174.1.0.0/31", false, "0/31.0.1.174.in-addr.arpa"},
		{"174.1.0.2/31", false, "2/31.0.1.174.in-addr.arpa"},

		// Error Cases:
		{"0.0.0.0/0", true, ""},
		{"2001::/0", true, ""},
		{"4.5/16", true, ""},
		{"foo.com", true, ""},
	}
	for i, tst := range tests {
		t.Run(fmt.Sprintf("%d--%s", i, tst.in), func(t *testing.T) {
			d, err := CidrToInAddr(tst.in)
			if err != nil && !tst.isError {
				t.Error("Should not have errored ", err)
			} else if tst.isError && err == nil {
				t.Errorf("Should have errored, but didn't. Got %s", d)
			} else if (!tst.isError) && d != tst.out {
				t.Errorf("Expected '%s' but got '%s'", tst.out, d)
			}
		})
	}
}
