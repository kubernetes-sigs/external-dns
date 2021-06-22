package tsig

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"time"

	c "github.com/bodgit/tsig/client"
	"github.com/miekg/dns"
)

const (
	// GSS is the RFC 3645 defined algorithm name
	GSS = "gss-tsig."
)

const (
	_ uint16 = iota // Reserved, RFC 2930, section 2.5
	// TkeyModeServer is used for server assigned keying
	TkeyModeServer
	// TkeyModeDH is used for Diffie-Hellman exchanged keying
	TkeyModeDH
	// TkeyModeGSS is used for GSS-API establishment
	TkeyModeGSS
	// TkeyModeResolver is used for resolver assigned keying
	TkeyModeResolver
	// TkeyModeDelete is used for key deletion
	TkeyModeDelete
)

// Exchanger is the interface a DNS client is expected to implement.
type Exchanger interface {
	Exchange(*dns.Msg, string) (*dns.Msg, time.Duration, error)
}

func calculateTimes(mode uint16, lifetime uint32) (uint32, uint32, error) {

	switch mode {
	case TkeyModeDH:
		fallthrough
	case TkeyModeGSS:
		now := time.Now().Unix()
		return uint32(now), uint32(now) + lifetime, nil
	case TkeyModeDelete:
		return 0, 0, nil
	default:
		return 0, 0, fmt.Errorf("Unsupported TKEY mode %d", mode)
	}
}

// SplitHostPort attempts to split a "hostname:port" string and return them
// as separate strings. If the host cannot be split then it is returned with
// the default DNS port "53".
func SplitHostPort(host string) (string, string) {

	hostname, port, err := net.SplitHostPort(host)
	if err != nil {
		return host, "53"
	}

	return hostname, port
}

func exchangeTKEY(client Exchanger, host, keyname, algorithm string, mode uint16, lifetime uint32, input []byte, extra []dns.RR, tsigname, tsigalgo, tsigmac *string) (*dns.TKEY, []dns.RR, error) {

	hostname, port := SplitHostPort(host)

	msg := &dns.Msg{
		MsgHdr: dns.MsgHdr{
			RecursionDesired: false,
		},
		Question: make([]dns.Question, 1),
		Extra:    make([]dns.RR, 1),
	}

	msg.Question[0] = dns.Question{
		Name:   keyname,
		Qtype:  dns.TypeTKEY,
		Qclass: dns.ClassANY,
	}

	msg.Id = dns.Id()

	inception, expiration, err := calculateTimes(mode, lifetime)
	if err != nil {
		return nil, nil, err
	}

	msg.Extra[0] = &dns.TKEY{
		Hdr: dns.RR_Header{
			Name:   keyname,
			Rrtype: dns.TypeTKEY,
			Class:  dns.ClassANY,
			Ttl:    0,
		},
		Algorithm:  algorithm,
		Mode:       mode,
		Inception:  inception,
		Expiration: expiration,
		KeySize:    uint16(len(input)),
		Key:        hex.EncodeToString(input),
	}

	msg.Extra = append(msg.Extra, extra...)

	if strings.ToLower(algorithm) != GSS && tsigname != nil && tsigalgo != nil && tsigmac != nil {
		msg.SetTsig(*tsigname, *tsigalgo, 300, time.Now().Unix())
	}

	rr, _, err := client.Exchange(msg, net.JoinHostPort(hostname, port))
	if err != nil {
		return nil, nil, err
	}

	if rr.Rcode != dns.RcodeSuccess {
		return nil, nil, fmt.Errorf("DNS error: %s (%d)", dns.RcodeToString[rr.Rcode], rr.Rcode)
	}

	additional := []dns.RR{}

	var tkey *dns.TKEY

	for _, ans := range rr.Answer {
		switch t := ans.(type) {
		case *dns.TKEY:
			// There mustn't be more than one TKEY answer RR
			if tkey != nil {
				return nil, nil, fmt.Errorf("Multiple TKEY responses")
			}
			tkey = t
		default:
			additional = append(additional, ans)
		}
	}

	// There should always be at least a TKEY answer RR
	if tkey == nil {
		return nil, nil, fmt.Errorf("Received no TKEY response")
	}

	if tkey.Error != 0 {
		return nil, nil, fmt.Errorf("TKEY error: %s (%d)", dns.RcodeToString[int(tkey.Error)], tkey.Error)
	}

	return tkey, additional, nil
}

// ExchangeTKEY exchanges TKEY records with the given host using the given
// key name, algorithm, mode, and lifetime with the provided input payload.
// Any additional DNS records are also sent and the exchange can be secured
// with TSIG if a key name, algorithm and MAC are provided.
// The TKEY record is returned along with any other DNS records in the
// response along with any error that occurred.
func ExchangeTKEY(host, keyname, algorithm string, mode uint16, lifetime uint32, input []byte, extra []dns.RR, tsigname, tsigalgo, tsigmac *string) (*dns.TKEY, []dns.RR, error) {

	client := c.Client{}

	// Use TCP regardless; TKEY queries can be in the range of ~ 1800 bytes
	client.Net = "tcp"

	// nsupdate(1) intentionally ignores the TSIG on the TKEY response for GSS
	if strings.ToLower(algorithm) == GSS {
		client.TsigAlgorithm = map[string]*c.TsigAlgorithm{GSS: {Generate: nil, Verify: nil}}
		client.TsigSecret = map[string]string{keyname: ""}
	} else if tsigname != nil && tsigmac != nil {
		client.TsigSecret = map[string]string{*tsigname: *tsigmac}
	}

	return exchangeTKEY(&client, host, keyname, algorithm, mode, lifetime, input, extra, tsigname, tsigalgo, tsigmac)
}
