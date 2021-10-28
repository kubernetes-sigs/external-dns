package transform

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// PtrNameMagic implements the PTR magic.
func PtrNameMagic(name, domain string) (string, error) {
	// Implement the PTR name magic.  If the name is a properly formed
	// IPv4 or IPv6 address, we replace it with the right string (i.e
	// reverse it and truncate it).

	// If the name is already in-addr.arpa or ipv6.arpa,
	// make sure the domain matches.
	if strings.HasSuffix(name, ".in-addr.arpa.") || strings.HasSuffix(name, ".ip6.arpa.") {
		if strings.HasSuffix(name, "."+domain+".") {
			return strings.TrimSuffix(name, "."+domain+"."), nil
		}
		return name, errors.Errorf("PTR record %v in wrong domain (%v)", name, domain)
	}

	// If the domain is .arpa, we do magic.
	if strings.HasSuffix(domain, ".in-addr.arpa") {
		return ipv4magic(name, domain)
	} else if strings.HasSuffix(domain, ".ip6.arpa") {
		return ipv6magic(name, domain)
	} else {
		return name, nil
	}
}

func ipv4magic(name, domain string) (string, error) {
	// Not a valid IPv4 address. Leave it alone.
	ip := net.ParseIP(name)
	if ip == nil || ip.To4() == nil || !strings.Contains(name, ".") {
		return name, nil
	}

	// Reverse it.
	rev, err := ReverseDomainName(ip.String() + "/32")
	if err != nil {
		return name, err
	}
	result := strings.TrimSuffix(rev, "."+domain)

	// Are we in the right domain?
	if strings.HasSuffix(rev, "."+domain) {
		return result, nil
	}
	if ipMatchesClasslessDomain(ip, domain) {
		return strings.SplitN(rev, ".", 2)[0], nil
	}

	return "", errors.Errorf("PTR record %v in wrong IPv4 domain (%v)", name, domain)
}

var isRfc2317Format1 = regexp.MustCompile(`(\d{1,3})/(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.in-addr\.arpa$`)

// ipMatchesClasslessDomain returns true if ip is appropriate for domain.
// domain is a reverse DNS lookup zone (in-addr.arpa) as described in RFC2317.
func ipMatchesClasslessDomain(ip net.IP, domain string) bool {

	// The unofficial but preferred format in RFC2317:
	m := isRfc2317Format1.FindStringSubmatch(domain)
	if m != nil {
		// IP:          Domain:
		// 172.20.18.27 128/27.18.20.172.in-addr.arpa
		// A   B  C  D  F   M  X  Y  Z
		// The following should be true:
		//   A==Z, B==Y, C==X.
		//   If you mask ip by M, the last octet should be F.
		ii := ip.To4()
		a, b, c, _ := ii[0], ii[1], ii[2], ii[3]
		f, m, x, y, z := atob(m[1]), atob(m[2]), atob(m[3]), atob(m[4]), atob(m[5])
		masked := ip.Mask(net.CIDRMask(int(m), 32))
		if a == z && b == y && c == x && masked.Equal(net.IPv4(a, b, c, f)) {
			return true
		}
	}

	// To extend this to include other formats, add them here.

	return false
}

// atob converts a to a byte value or panics.
func atob(s string) byte {
	if i, err := strconv.Atoi(s); err == nil {
		if i < 256 {
			return byte(i)
		}
	}
	panic(fmt.Sprintf("(%v) matched \\d{1,3} but is not a byte", s))
}

func ipv6magic(name, domain string) (string, error) {
	// Not a valid IPv6 address. Leave it alone.
	ip := net.ParseIP(name)
	if ip == nil || len(ip) != 16 || !strings.Contains(name, ":") {
		return name, nil
	}

	// Reverse it.
	rev, err := ReverseDomainName(ip.String() + "/128")
	if err != nil {
		return name, err
	}
	if !strings.HasSuffix(rev, "."+domain) {
		err = errors.Errorf("PTR record %v in wrong IPv6 domain (%v)", name, domain)
	}
	return strings.TrimSuffix(rev, "."+domain), err
}
