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
	"net"
	"strconv"
	"strings"
)

// CidrToInAddr converts a CIDR block into its reverse lookup (in-addr) name.
// Given "2001::/16" returns "1.0.0.2.ip6.arpa"
// Given "10.20.30.0/24" returns "30.20.10.in-addr.arpa"
// Given "10.20.30.0/25" returns "0/25.30.20.10.in-addr.arpa" (RFC2317)
func CidrToInAddr(cidr string) (string, error) {
	// If the user sent an IP instead of a CIDR (i.e. no "/"), turn it
	// into a CIDR by adding /32 or /128 as appropriate.
	ip := net.ParseIP(cidr)
	if ip != nil {
		if ip.To4() != nil {
			cidr = ip.String() + "/32"
			// Older code used `cidr + "/32"` but that didn't work with
			// "IPv4 mapped IPv6 address". ip.String() returns the IPv4
			// address for all IPv4 addresses no matter how they are
			// expressed internally.
		} else {
			cidr = cidr + "/128"
		}
	}

	a, c, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}
	base, err := reverseaddr(a.String())
	if err != nil {
		return "", err
	}
	base = strings.TrimRight(base, ".")
	if !a.Equal(c.IP) {
		return "", fmt.Errorf("CIDR %v has 1 bits beyond the mask", cidr)
	}

	bits, total := c.Mask.Size()
	var toTrim int
	if bits == 0 {
		return "", fmt.Errorf("cannot use /0 in reverse CIDR")
	}

	// Handle IPv4 "Classless in-addr.arpa delegation" RFC2317:
	if total == 32 && bits >= 25 && bits < 32 {
		// first address / netmask . Class-b-arpa.
		fparts := strings.Split(c.IP.String(), ".")
		first := fparts[3]
		bparts := strings.SplitN(base, ".", 2)
		return fmt.Sprintf("%s/%d.%s", first, bits, bparts[1]), nil
	}

	// Handle IPv4 Class-full and IPv6:
	if total == 32 {
		if bits%8 != 0 {
			return "", fmt.Errorf("IPv4 mask must be multiple of 8 bits")
		}
		toTrim = (total - bits) / 8
	} else if total == 128 {
		if bits%4 != 0 {
			return "", fmt.Errorf("IPv6 mask must be multiple of 4 bits")
		}
		toTrim = (total - bits) / 4
	} else {
		return "", fmt.Errorf("invalid address (not IPv4 or IPv6): %v", cidr)
	}

	parts := strings.SplitN(base, ".", toTrim+1)
	return parts[len(parts)-1], nil
}

// copied from go source.
// https://github.com/golang/go/blob/38b2c06e144c6ea7087c575c76c66e41265ae0b7/src/net/dnsclient.go#L26C1-L51C1
// The go source does not export this function so we copy it here.

// reverseaddr returns the in-addr.arpa. or ip6.arpa. hostname of the IP
// address addr suitable for rDNS (PTR) record lookup or an error if it fails
// to parse the IP address.
func reverseaddr(addr string) (arpa string, err error) {
	ip := net.ParseIP(addr)
	if ip == nil {
		return "", &net.DNSError{Err: "unrecognized address", Name: addr}
	}
	if ip.To4() != nil {
		return Uitoa(uint(ip[15])) + "." + Uitoa(uint(ip[14])) + "." + Uitoa(uint(ip[13])) + "." + Uitoa(uint(ip[12])) + ".in-addr.arpa.", nil
	}
	// Must be IPv6
	buf := make([]byte, 0, len(ip)*4+len("ip6.arpa."))
	// Add it, in reverse, to the buffer
	for i := len(ip) - 1; i >= 0; i-- {
		v := ip[i]
		buf = append(buf, hexDigit[v&0xF],
			'.',
			hexDigit[v>>4],
			'.')
	}
	// Append "ip6.arpa." and return (buf already has the final .)
	buf = append(buf, "ip6.arpa."...)
	return string(buf), nil
}

const hexDigit = "0123456789abcdef"

func Uitoa(val uint) string {
	return strconv.FormatInt(int64(val), 10)
}
