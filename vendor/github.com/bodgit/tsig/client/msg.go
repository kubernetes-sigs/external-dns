// DNS packet assembly, see RFC 1035. Converting from - Unpack() -
// and to - Pack() - wire format.
// All the packers and unpackers take a (msg []byte, off int)
// and return (off1 int, ok bool).  If they return ok==false, they
// also return off1==len(msg), so that the next unpacker will
// also fail.  This lets us avoid checks of ok until the end of a
// packing sequence.

package client

import (
	"strings"

	"github.com/miekg/dns"
)

const (
	maxCompressionOffset = 2 << 13 // We have 14 bits for the compression pointer
)

// Helpers for dealing with escaped bytes
func isDigit(b byte) bool { return b >= '0' && b <= '9' }

// unpackRRslice unpacks msg[off:] into an []RR.
// If we cannot unpack the whole array, then it will return nil
func unpackRRslice(l int, msg []byte, off int) (dst1 []dns.RR, off1 int, err error) {
	var r dns.RR
	// Don't pre-allocate, l may be under attacker control
	var dst []dns.RR
	for i := 0; i < l; i++ {
		off1 := off
		r, off, err = dns.UnpackRR(msg, off)
		if err != nil {
			off = len(msg)
			break
		}
		// If offset does not increase anymore, l is a lie
		if off1 == off {
			break
		}
		dst = append(dst, r)
	}
	if err != nil && off == len(msg) {
		dst = nil
	}
	return dst, off, err
}

func domainNameLen(s string, off int, compression map[string]struct{}, compress bool) int {
	if s == "" || s == "." {
		return 1
	}

	escaped := strings.Contains(s, "\\")

	if compression != nil && (compress || off < maxCompressionOffset) {
		// compressionLenSearch will insert the entry into the compression
		// map if it doesn't contain it.
		if l, ok := compressionLenSearch(compression, s, off); ok && compress {
			if escaped {
				return escapedNameLen(s[:l]) + 2
			}

			return l + 2
		}
	}

	if escaped {
		return escapedNameLen(s) + 1
	}

	return len(s) + 1
}

func escapedNameLen(s string) int {
	nameLen := len(s)
	for i := 0; i < len(s); i++ {
		if s[i] != '\\' {
			continue
		}

		if i+3 < len(s) && isDigit(s[i+1]) && isDigit(s[i+2]) && isDigit(s[i+3]) {
			nameLen -= 3
			i += 3
		} else {
			nameLen--
			i++
		}
	}

	return nameLen
}

func compressionLenSearch(c map[string]struct{}, s string, msgOff int) (int, bool) {
	for off, end := 0, false; !end; off, end = dns.NextLabel(s, off) {
		if _, ok := c[s[off:]]; ok {
			return off, true
		}

		if msgOff+off < maxCompressionOffset {
			c[s[off:]] = struct{}{}
		}
	}

	return 0, false
}

func unpackQuestion(msg []byte, off int) (dns.Question, int, error) {
	var (
		q   dns.Question
		err error
	)
	q.Name, off, err = dns.UnpackDomainName(msg, off)
	if err != nil {
		return q, off, err
	}
	if off == len(msg) {
		return q, off, nil
	}
	q.Qtype, off, err = unpackUint16(msg, off)
	if err != nil {
		return q, off, err
	}
	if off == len(msg) {
		return q, off, nil
	}
	q.Qclass, off, err = unpackUint16(msg, off)
	if off == len(msg) {
		return q, off, nil
	}
	return q, off, err
}

func unpackMsgHdr(msg []byte, off int) (dns.Header, int, error) {
	var (
		dh  dns.Header
		err error
	)
	dh.Id, off, err = unpackUint16(msg, off)
	if err != nil {
		return dh, off, err
	}
	dh.Bits, off, err = unpackUint16(msg, off)
	if err != nil {
		return dh, off, err
	}
	dh.Qdcount, off, err = unpackUint16(msg, off)
	if err != nil {
		return dh, off, err
	}
	dh.Ancount, off, err = unpackUint16(msg, off)
	if err != nil {
		return dh, off, err
	}
	dh.Nscount, off, err = unpackUint16(msg, off)
	if err != nil {
		return dh, off, err
	}
	dh.Arcount, off, err = unpackUint16(msg, off)
	if err != nil {
		return dh, off, err
	}
	return dh, off, nil
}
