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
	"slices"
	"strconv"
	"strings"
)

// CidrToInAddr converts a CIDR block into its reverse lookup (in-addr) name.
// Given "2001::/16" returns "1.0.0.2.ip6.arpa"
// Given "10.20.30.0/24" returns "30.20.10.in-addr.arpa"
// Given "10.20.30.0/25" returns "0/25.30.20.10.in-addr.arpa" (RFC2317)
func CidrToInAddr(cidr string) (string, error) {
	cidr = normalizeCIDR(cidr)

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
	if bits == 0 {
		return "", fmt.Errorf("cannot use /0 in reverse CIDR")
	}

	// Handle IPv4 "Classless in-addr.arpa delegation" RFC2317:
	if name, ok := classlessIPv4Name(c, base, total, bits); ok {
		return name, nil
	}

	toTrim, err := reverseNameTrimCount(total, bits, cidr)
	if err != nil {
		return "", err
	}

	parts := strings.SplitN(base, ".", toTrim+1)
	return parts[len(parts)-1], nil
}

// normalizeCIDR turns a bare IP address into a host-sized CIDR. Using
// ip.String() for IPv4 also normalizes IPv4-mapped IPv6 addresses.
func normalizeCIDR(cidr string) string {
	ip := net.ParseIP(cidr)
	if ip == nil {
		return cidr
	}
	if ip.To4() != nil {
		return ip.String() + "/32"
	}
	return cidr + "/128"
}

func classlessIPv4Name(network *net.IPNet, base string, total, bits int) (string, bool) {
	if total != 32 || bits < 25 || bits >= 32 {
		return "", false
	}

	// first address / netmask . Class-b-arpa.
	fparts := strings.Split(network.IP.String(), ".")
	first := fparts[3]
	bparts := strings.SplitN(base, ".", 2)
	return fmt.Sprintf("%s/%d.%s", first, bits, bparts[1]), true
}

func reverseNameTrimCount(total, bits int, cidr string) (int, error) {
	switch total {
	case 32:
		if bits%8 != 0 {
			return 0, fmt.Errorf("IPv4 mask must be multiple of 8 bits")
		}
		return (total - bits) / 8, nil
	case 128:
		if bits%4 != 0 {
			return 0, fmt.Errorf("IPv6 mask must be multiple of 4 bits")
		}
		return (total - bits) / 4, nil
	default:
		return 0, fmt.Errorf("invalid address (not IPv4 or IPv6): %v", cidr)
	}
}

// copied from go source.
// https://github.com/golang/go/blob/38b2c06e144c6ea7087c575c76c66e41265ae0b7/src/net/dnsclient.go#L26C1-L51C1
// The go source does not export this function so we copy it here.

// reverseaddr returns the in-addr.arpa. or ip6.arpa. hostname of the IP
// address addr suitable for rDNS (PTR) record lookup or an error if it fails
// to parse the IP address.
func reverseaddr(addr string) (string, error) {
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
	for _, v := range slices.Backward(ip) {
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
