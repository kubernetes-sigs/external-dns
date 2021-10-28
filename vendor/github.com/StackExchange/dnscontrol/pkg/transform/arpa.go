package transform

import (
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"
)

// ReverseDomainName turns a CIDR block into a reversed (in-addr) name.
func ReverseDomainName(cidr string) (string, error) {
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
		return "", errors.Errorf("CIDR %v has 1 bits beyond the mask", cidr)
	}

	bits, total := c.Mask.Size()
	var toTrim int
	if bits == 0 {
		return "", errors.Errorf("Cannot use /0 in reverse cidr")
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
			return "", errors.Errorf("IPv4 mask must be multiple of 8 bits")
		}
		toTrim = (total - bits) / 8
	} else if total == 128 {
		if bits%4 != 0 {
			return "", errors.Errorf("IPv6 mask must be multiple of 4 bits")
		}
		toTrim = (total - bits) / 4
	} else {
		return "", errors.Errorf("Address is not IPv4 or IPv6: %v", cidr)
	}

	parts := strings.SplitN(base, ".", toTrim+1)
	return parts[len(parts)-1], nil
}

// copied from go source.
// https://github.com/golang/go/blob/bfc164c64d33edfaf774b5c29b9bf5648a6447fb/src/net/dnsclient.go#L15

// reverseaddr returns the in-addr.arpa. or ip6.arpa. hostname of the IP
// address addr suitable for rDNS (PTR) record lookup or an error if it fails
// to parse the IP address.
func reverseaddr(addr string) (arpa string, err error) {
	ip := net.ParseIP(addr)
	if ip == nil {
		return "", &net.DNSError{Err: "unrecognized address", Name: addr}
	}
	if ip.To4() != nil {
		return uitoa(uint(ip[15])) + "." + uitoa(uint(ip[14])) + "." + uitoa(uint(ip[13])) + "." + uitoa(uint(ip[12])) + ".in-addr.arpa.", nil
	}
	// Must be IPv6
	buf := make([]byte, 0, len(ip)*4+len("ip6.arpa."))
	// Add it, in reverse, to the buffer
	for i := len(ip) - 1; i >= 0; i-- {
		v := ip[i]
		buf = append(buf, hexDigit[v&0xF])
		buf = append(buf, '.')
		buf = append(buf, hexDigit[v>>4])
		buf = append(buf, '.')
	}
	// Append "ip6.arpa." and return (buf already has the final .)
	buf = append(buf, "ip6.arpa."...)
	return string(buf), nil
}

// Convert unsigned integer to decimal string.
func uitoa(val uint) string {
	if val == 0 { // avoid string allocation
		return "0"
	}
	var buf [20]byte // big enough for 64bit value base 10
	i := len(buf) - 1
	for val >= 10 {
		q := val / 10
		buf[i] = byte('0' + val - q*10)
		i--
		val = q
	}
	// val < 10
	buf[i] = byte('0' + val)
	return string(buf[i:])
}

const hexDigit = "0123456789abcdef"
