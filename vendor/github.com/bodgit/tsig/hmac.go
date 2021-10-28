package tsig

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"

	"github.com/miekg/dns"
)

// HMAC implements the standard HMAC TSIG methods using the dns.TsigProvider
// interface. It holds a map of TSIG key names to base64-encoded secrets. The
// key names should be in canonical form, see dns.CanonicalName.
type HMAC map[string]string

func fromBase64(s []byte) (buf []byte, err error) {
	buflen := base64.StdEncoding.DecodedLen(len(s))
	buf = make([]byte, buflen)
	n, err := base64.StdEncoding.Decode(buf, s)
	buf = buf[:n]
	return
}

// Generate generates the TSIG MAC using the HMAC algorithm indicated by
// t.Algorithm using h[t.Hdr.Name] as the key.
// It returns the bytes for the TSIG MAC and any error that occurred.
func (h HMAC) Generate(msg []byte, t *dns.TSIG) ([]byte, error) {
	var f func() hash.Hash
	switch dns.CanonicalName(t.Algorithm) {
	case dns.HmacMD5:
		f = md5.New
	case dns.HmacSHA1:
		f = sha1.New
	case dns.HmacSHA224:
		f = sha256.New224
	case dns.HmacSHA256:
		f = sha256.New
	case dns.HmacSHA384:
		f = sha512.New384
	case dns.HmacSHA512:
		f = sha512.New
	default:
		return nil, dns.ErrKeyAlg
	}
	secret, ok := h[t.Hdr.Name]
	if !ok {
		return nil, dns.ErrSecret
	}
	rawsecret, err := fromBase64([]byte(secret))
	if err != nil {
		return nil, err
	}
	m := hmac.New(f, rawsecret)
	m.Write(msg)
	return m.Sum(nil), nil
}

// Verify verifies the TSIG MAC using the HMAC algorithm indicated by
// t.Algorithm using h[t.Hdr.Name] as the key.
// It returns any error that occurred.
func (h HMAC) Verify(msg []byte, t *dns.TSIG) error {
	b, err := h.Generate(msg, t)
	if err != nil {
		return err
	}
	mac, err := hex.DecodeString(t.MAC)
	if err != nil {
		return err
	}
	if !hmac.Equal(b, mac) {
		return dns.ErrSig
	}
	return nil
}
