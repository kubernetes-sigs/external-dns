package client

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"hash"
	"time"

	"github.com/miekg/dns"
)

type tsigAlgorithmGenerate func(msg []byte, algorithm, name, secret string) ([]byte, error)
type tsigAlgorithmVerify func(msg []byte, tsig *dns.TSIG, name, secret string) error

// TsigAlgorithm holds the two callbacks used to generate and verify the
// transaction signature of a message.
type TsigAlgorithm struct {
	Generate tsigAlgorithmGenerate
	Verify   tsigAlgorithmVerify
}

// TSIG is the RR the holds the transaction signature of a message.
// See RFC 2845 and RFC 4635.
type TSIG struct {
	dns.TSIG
}

// The following values must be put in wireformat, so that the MAC can be calculated.
// RFC 2845, section 3.4.2. TSIG Variables.
type tsigWireFmt struct {
	// From RR_Header
	Name  string `dns:"domain-name"`
	Class uint16
	Ttl   uint32
	// Rdata of the TSIG
	Algorithm  string `dns:"domain-name"`
	TimeSigned uint64 `dns:"uint48"`
	Fudge      uint16
	// MACSize, MAC and OrigId excluded
	Error     uint16
	OtherLen  uint16
	OtherData string `dns:"size-hex:OtherLen"`
}

// If we have the MAC use this type to convert it to wiredata. Section 3.4.3. Request MAC
type macWireFmt struct {
	MACSize uint16
	MAC     string `dns:"size-hex:MACSize"`
}

// 3.3. Time values used in TSIG calculations
type timerWireFmt struct {
	TimeSigned uint64 `dns:"uint48"`
	Fudge      uint16
}

func tsigGenerateHmac(msg []byte, algorithm, name, secret string) ([]byte, error) {
	// If we barf here, the caller is to blame
	rawsecret, err := fromBase64([]byte(secret))
	if err != nil {
		return nil, err
	}

	var h hash.Hash
	switch dns.CanonicalName(algorithm) {
	case dns.HmacMD5:
		h = hmac.New(md5.New, rawsecret)
	case dns.HmacSHA1:
		h = hmac.New(sha1.New, rawsecret)
	case dns.HmacSHA256:
		h = hmac.New(sha256.New, rawsecret)
	case dns.HmacSHA512:
		h = hmac.New(sha512.New, rawsecret)
	default:
		return nil, dns.ErrKeyAlg
	}
	h.Write(msg)

	return h.Sum(nil), nil
}

// TsigGenerate fills out the TSIG record attached to the message.
// The message should contain
// a "stub" TSIG RR with the algorithm, key name (owner name of the RR),
// time fudge (defaults to 300 seconds) and the current time
// The TSIG MAC is saved in that Tsig RR.
// When TsigGenerate is called for the first time requestMAC is set to the empty string and
// timersOnly is false.
// If something goes wrong an error is returned, otherwise it is nil.
func TsigGenerate(m *dns.Msg, secret, requestMAC string, timersOnly bool) ([]byte, string, error) {
	return TsigGenerateByAlgorithm(m, tsigGenerateHmac, "", secret, requestMAC, timersOnly)
}

// TsigGenerateByAlgorithm fills out the TSIG record attached to the message
// using a callback to implement the algorithm-specific generation.
// The message should contain
// a "stub" TSIG RR with the algorithm, key name (owner name of the RR),
// time fudge (defaults to 300 seconds) and the current time
// The TSIG MAC is saved in that Tsig RR.
// When TsigGenerate is called for the first time requestMAC is set to the empty string and
// timersOnly is false.
// If something goes wrong an error is returned, otherwise it is nil.
func TsigGenerateByAlgorithm(m *dns.Msg, cb tsigAlgorithmGenerate, name, secret, requestMAC string, timersOnly bool) ([]byte, string, error) {
	if m.IsTsig() == nil {
		panic("dns: TSIG not last RR in additional")
	}

	rr := m.Extra[len(m.Extra)-1].(*dns.TSIG)
	m.Extra = m.Extra[0 : len(m.Extra)-1] // kill the TSIG from the msg
	mbuf, err := m.Pack()
	if err != nil {
		return nil, "", err
	}
	buf, err := tsigBuffer(mbuf, rr, requestMAC, timersOnly)
	if err != nil {
		return nil, "", err
	}

	t := new(TSIG)

	h, err := cb(buf, rr.Algorithm, name, secret)
	if err != nil {
		return nil, "", err
	}

	// Copy all TSIG fields except MAC and its size, which are filled using the computed digest.
	t.TSIG = *rr
	t.MAC = hex.EncodeToString(h)
	t.MACSize = uint16(len(t.MAC) / 2) // Size is half!

	tbuf := make([]byte, dns.Len(t))
	off, err := dns.PackRR(t, tbuf, 0, nil, false)
	if err != nil {
		return nil, "", err
	}
	mbuf = append(mbuf, tbuf[:off]...)
	// Update the ArCount directly in the buffer.
	binary.BigEndian.PutUint16(mbuf[10:], uint16(len(m.Extra)+1))

	return mbuf, t.MAC, nil
}

func tsigVerifyHmac(msg []byte, tsig *dns.TSIG, name, secret string) error {
	rawsecret, err := fromBase64([]byte(secret))
	if err != nil {
		return err
	}

	msgMAC, err := hex.DecodeString(tsig.MAC)
	if err != nil {
		return err
	}

	var h hash.Hash
	switch dns.CanonicalName(tsig.Algorithm) {
	case dns.HmacMD5:
		h = hmac.New(md5.New, rawsecret)
	case dns.HmacSHA1:
		h = hmac.New(sha1.New, rawsecret)
	case dns.HmacSHA256:
		h = hmac.New(sha256.New, rawsecret)
	case dns.HmacSHA512:
		h = hmac.New(sha512.New, rawsecret)
	default:
		return dns.ErrKeyAlg
	}
	h.Write(msg)
	if !hmac.Equal(h.Sum(nil), msgMAC) {
		return dns.ErrSig
	}
	return nil
}

// TsigVerify verifies the TSIG on a message.
// If the signature does not validate err contains the
// error, otherwise it is nil.
func TsigVerify(msg []byte, secret, requestMAC string, timersOnly bool) error {
	return TsigVerifyByAlgorithm(msg, tsigVerifyHmac, "", secret, requestMAC, timersOnly)
}

// TsigVerifyByAlgorithm verifies the TSIG on a message using a callback to
// implement the algorithm-specific verification.
// If the signature does not validate err contains the
// error, otherwise it is nil.
func TsigVerifyByAlgorithm(msg []byte, cb tsigAlgorithmVerify, name, secret, requestMAC string, timersOnly bool) error {
	// Strip the TSIG from the incoming msg
	stripped, tsig, err := stripTsig(msg)
	if err != nil {
		return err
	}

	buf, err := tsigBuffer(stripped, tsig, requestMAC, timersOnly)
	if err != nil {
		return err
	}

	if err := cb(buf, tsig, name, secret); err != nil {
		return err
	}

	// Fudge factor works both ways. A message can arrive before it was signed because
	// of clock skew.
	// We check this after verifying the signature, following draft-ietf-dnsop-rfc2845bis
	// instead of RFC2845, in order to prevent a security vulnerability as reported in CVE-2017-3142/3143.
	now := uint64(time.Now().Unix())
	ti := now - tsig.TimeSigned
	if now < tsig.TimeSigned {
		ti = tsig.TimeSigned - now
	}
	if uint64(tsig.Fudge) < ti {
		return dns.ErrTime
	}

	return nil
}

// Create a wiredata buffer for the MAC calculation.
func tsigBuffer(msgbuf []byte, rr *dns.TSIG, requestMAC string, timersOnly bool) ([]byte, error) {
	var buf []byte
	if rr.TimeSigned == 0 {
		rr.TimeSigned = uint64(time.Now().Unix())
	}
	if rr.Fudge == 0 {
		rr.Fudge = 300 // Standard (RFC) default.
	}

	// Replace message ID in header with original ID from TSIG
	binary.BigEndian.PutUint16(msgbuf[0:2], rr.OrigId)

	if requestMAC != "" {
		m := new(macWireFmt)
		m.MACSize = uint16(len(requestMAC) / 2)
		m.MAC = requestMAC
		buf = make([]byte, len(requestMAC)) // long enough
		n, err := packMacWire(m, buf)
		if err != nil {
			return nil, err
		}
		buf = buf[:n]
	}

	tsigvar := make([]byte, dns.DefaultMsgSize)
	if timersOnly {
		tsig := new(timerWireFmt)
		tsig.TimeSigned = rr.TimeSigned
		tsig.Fudge = rr.Fudge
		n, err := packTimerWire(tsig, tsigvar)
		if err != nil {
			return nil, err
		}
		tsigvar = tsigvar[:n]
	} else {
		tsig := new(tsigWireFmt)
		tsig.Name = dns.CanonicalName(rr.Hdr.Name)
		tsig.Class = dns.ClassANY
		tsig.Ttl = rr.Hdr.Ttl
		tsig.Algorithm = dns.CanonicalName(rr.Algorithm)
		tsig.TimeSigned = rr.TimeSigned
		tsig.Fudge = rr.Fudge
		tsig.Error = rr.Error
		tsig.OtherLen = rr.OtherLen
		tsig.OtherData = rr.OtherData
		n, err := packTsigWire(tsig, tsigvar)
		if err != nil {
			return nil, err
		}
		tsigvar = tsigvar[:n]
	}

	if requestMAC != "" {
		x := append(buf, msgbuf...)
		buf = append(x, tsigvar...)
	} else {
		buf = append(msgbuf, tsigvar...)
	}
	return buf, nil
}

// Strip the TSIG from the raw message.
func stripTsig(msg []byte) ([]byte, *dns.TSIG, error) {
	// Copied from msg.go's Unpack() Header, but modified.
	var (
		dh  dns.Header
		err error
	)
	off, tsigoff := 0, 0

	if dh, off, err = unpackMsgHdr(msg, off); err != nil {
		return nil, nil, err
	}
	if dh.Arcount == 0 {
		return nil, nil, dns.ErrNoSig
	}

	// Rcode, see msg.go Unpack()
	if int(dh.Bits&0xF) == dns.RcodeNotAuth {
		return nil, nil, dns.ErrAuth
	}

	for i := 0; i < int(dh.Qdcount); i++ {
		_, off, err = unpackQuestion(msg, off)
		if err != nil {
			return nil, nil, err
		}
	}

	_, off, err = unpackRRslice(int(dh.Ancount), msg, off)
	if err != nil {
		return nil, nil, err
	}
	_, off, err = unpackRRslice(int(dh.Nscount), msg, off)
	if err != nil {
		return nil, nil, err
	}

	rr := new(dns.TSIG)
	var extra dns.RR
	for i := 0; i < int(dh.Arcount); i++ {
		tsigoff = off
		extra, off, err = dns.UnpackRR(msg, off)
		if err != nil {
			return nil, nil, err
		}
		if extra.Header().Rrtype == dns.TypeTSIG {
			rr = extra.(*dns.TSIG)
			// Adjust Arcount.
			arcount := binary.BigEndian.Uint16(msg[10:])
			binary.BigEndian.PutUint16(msg[10:], arcount-1)
			break
		}
	}
	if rr == nil {
		return nil, nil, dns.ErrNoSig
	}
	return msg[:tsigoff], rr, nil
}

func packTsigWire(tw *tsigWireFmt, msg []byte) (int, error) {
	// copied from zmsg.go TSIG packing
	// RR_Header
	off, err := dns.PackDomainName(tw.Name, msg, 0, nil, false)
	if err != nil {
		return off, err
	}
	off, err = packUint16(tw.Class, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint32(tw.Ttl, msg, off)
	if err != nil {
		return off, err
	}

	off, err = dns.PackDomainName(tw.Algorithm, msg, off, nil, false)
	if err != nil {
		return off, err
	}
	off, err = packUint48(tw.TimeSigned, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint16(tw.Fudge, msg, off)
	if err != nil {
		return off, err
	}

	off, err = packUint16(tw.Error, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packUint16(tw.OtherLen, msg, off)
	if err != nil {
		return off, err
	}
	off, err = packStringHex(tw.OtherData, msg, off)
	if err != nil {
		return off, err
	}
	return off, nil
}

func packMacWire(mw *macWireFmt, msg []byte) (int, error) {
	off, err := packUint16(mw.MACSize, msg, 0)
	if err != nil {
		return off, err
	}
	off, err = packStringHex(mw.MAC, msg, off)
	if err != nil {
		return off, err
	}
	return off, nil
}

func packTimerWire(tw *timerWireFmt, msg []byte) (int, error) {
	off, err := packUint48(tw.TimeSigned, msg, 0)
	if err != nil {
		return off, err
	}
	off, err = packUint16(tw.Fudge, msg, off)
	if err != nil {
		return off, err
	}
	return off, nil
}
