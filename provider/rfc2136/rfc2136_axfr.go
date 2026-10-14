/*
Copyright 2026 The Kubernetes Authors.

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

package rfc2136

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/bodgit/tsig"
	"github.com/miekg/dns"
)

const maxUnsignedAXFRMessages = 99

func (r *rfc2136Provider) signedIncomeTransfer(m *dns.Msg, nameserver string) (chan *dns.Envelope, error) {
	client, err := makeClient(r, nameserver)
	if err != nil {
		return nil, fmt.Errorf("error setting up TLS: %w", err)
	}
	conn, err := client.Dial(nameserver)
	if err != nil {
		return nil, fmt.Errorf("failed to connect for transfer: %w", err)
	}

	wire, requestMAC, err := dns.TsigGenerate(m, r.tsigSecret, "", false)
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("failed to sign AXFR request: %w", err)
	}
	if _, err := conn.Write(wire); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("failed to send AXFR request: %w", err)
	}

	return readSignedAXFR(conn, m.Id, requestMAC, r.tsigSecret, tsig.HMAC{r.tsigKeyName: r.tsigSecret}), nil
}

func readSignedAXFR(conn *dns.Conn, id uint16, requestMAC, secret string, provider dns.TsigProvider) chan *dns.Envelope {
	envelopes := make(chan *dns.Envelope)
	go func() {
		defer close(envelopes)
		defer conn.Close()

		first := true
		priorMAC := requestMAC
		var unsignedMessages [][]byte
		for {
			firstMessage := first
			if err := conn.Conn.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
				envelopes <- &dns.Envelope{Error: err}
				return
			}
			wire, err := conn.ReadMsgHeader(nil)
			if err != nil {
				envelopes <- &dns.Envelope{Error: err}
				return
			}
			message := new(dns.Msg)
			if err := message.Unpack(wire); err != nil {
				envelopes <- &dns.Envelope{Error: err}
				return
			}
			if message.Id != id {
				envelopes <- &dns.Envelope{RR: message.Answer, Error: dns.ErrId}
				return
			}
			if message.Rcode != dns.RcodeSuccess {
				envelopes <- &dns.Envelope{RR: message.Answer, Error: fmt.Errorf("bad AXFR rcode: %d", message.Rcode)}
				return
			}
			if first && (len(message.Answer) == 0 || message.Answer[0].Header().Rrtype != dns.TypeSOA) {
				envelopes <- &dns.Envelope{RR: message.Answer, Error: dns.ErrSoa}
				return
			}
			if len(message.Answer) == 0 {
				envelopes <- &dns.Envelope{Error: dns.ErrSoa}
				return
			}

			ts := message.IsTsig()
			if first {
				if ts == nil {
					envelopes <- &dns.Envelope{Error: dns.ErrNoSig}
					return
				}
				err = dns.TsigVerify(wire, secret, requestMAC, false)
			} else if ts == nil {
				if isSOALast(message) {
					envelopes <- &dns.Envelope{Error: dns.ErrNoSig}
					return
				}
				unsignedMessages = append(unsignedMessages, wire)
				if len(unsignedMessages) > maxUnsignedAXFRMessages {
					envelopes <- &dns.Envelope{Error: dns.ErrNoSig}
					return
				}
			} else {
				current, stripErr := stripAXFRTsig(wire, ts)
				if stripErr != nil {
					envelopes <- &dns.Envelope{Error: stripErr}
					return
				}
				macInput, inputErr := axfrTsigMACInput(priorMAC, unsignedMessages, current, ts)
				if inputErr != nil {
					envelopes <- &dns.Envelope{Error: inputErr}
					return
				}
				chainedProvider := axfrTsigProvider{provider: provider, macInput: macInput}
				err = dns.TsigVerifyWithProvider(wire, chainedProvider, priorMAC, true)
			}
			if err != nil {
				envelopes <- &dns.Envelope{Error: err}
				return
			}
			if ts != nil {
				priorMAC = ts.MAC
				unsignedMessages = nil
			}
			first = false

			envelopes <- &dns.Envelope{RR: message.Answer}
			if isSOALast(message) && (!firstMessage || len(message.Answer) > 1) {
				return
			}
		}
	}()
	return envelopes
}

func isSOALast(message *dns.Msg) bool {
	return len(message.Answer) > 0 && message.Answer[len(message.Answer)-1].Header().Rrtype == dns.TypeSOA
}

type axfrTsigProvider struct {
	provider dns.TsigProvider
	macInput []byte
}

func (p axfrTsigProvider) Generate(_ []byte, ts *dns.TSIG) ([]byte, error) {
	return p.provider.Generate(p.macInput, ts)
}

func (p axfrTsigProvider) Verify(_ []byte, ts *dns.TSIG) error {
	return p.provider.Verify(p.macInput, ts)
}

func axfrTsigMACInput(priorMAC string, unsignedMessages [][]byte, current []byte, ts *dns.TSIG) ([]byte, error) {
	mac, err := hex.DecodeString(priorMAC)
	if err != nil {
		return nil, fmt.Errorf("invalid prior TSIG MAC: %w", err)
	}
	if len(mac) > 0xffff {
		return nil, fmt.Errorf("prior TSIG MAC is too long")
	}
	input := make([]byte, 2, 2+len(mac)+len(current)+8)
	binary.BigEndian.PutUint16(input, uint16(len(mac)))
	input = append(input, mac...)
	for _, message := range unsignedMessages {
		input = append(input, message...)
	}
	input = append(input, current...)
	var timers [8]byte
	binary.BigEndian.PutUint64(timers[:], ts.TimeSigned)
	input = append(input, timers[2:]...)
	input = binary.BigEndian.AppendUint16(input, ts.Fudge)
	return input, nil
}

func stripAXFRTsig(wire []byte, ts *dns.TSIG) ([]byte, error) {
	if len(wire) < 12 {
		return nil, fmt.Errorf("invalid DNS message: header is truncated")
	}
	additionalCount := binary.BigEndian.Uint16(wire[10:12])
	if additionalCount == 0 {
		return nil, dns.ErrNoSig
	}
	offset := 12
	questionCount := binary.BigEndian.Uint16(wire[4:6])
	answerCount := binary.BigEndian.Uint16(wire[6:8])
	authorityCount := binary.BigEndian.Uint16(wire[8:10])
	for range questionCount {
		var err error
		offset, err = skipAXFRName(wire, offset)
		if err != nil || offset+4 > len(wire) {
			return nil, fmt.Errorf("invalid DNS question")
		}
		offset += 4
	}
	for range int(answerCount) + int(authorityCount) + int(additionalCount) - 1 {
		var err error
		offset, _, err = skipAXFRRR(wire, offset)
		if err != nil {
			return nil, err
		}
	}
	start, rrType, err := skipAXFRRR(wire, offset)
	if err != nil {
		return nil, err
	}
	if rrType != dns.TypeTSIG || start != len(wire) {
		return nil, fmt.Errorf("TSIG record is not the final additional record")
	}
	stripped := append([]byte(nil), wire[:offset]...)
	binary.BigEndian.PutUint16(stripped[0:2], ts.OrigId)
	binary.BigEndian.PutUint16(stripped[10:12], additionalCount-1)
	return stripped, nil
}

func skipAXFRName(wire []byte, offset int) (int, error) {
	for {
		if offset >= len(wire) {
			return 0, fmt.Errorf("invalid DNS name")
		}
		length := wire[offset]
		switch length & 0xc0 {
		case 0xc0:
			if offset+1 >= len(wire) {
				return 0, fmt.Errorf("invalid DNS compression pointer")
			}
			return offset + 2, nil
		case 0:
			offset++
			if length == 0 {
				return offset, nil
			}
			offset += int(length)
		default:
			return 0, fmt.Errorf("invalid DNS label length")
		}
	}
}

func skipAXFRRR(wire []byte, offset int) (int, uint16, error) {
	offset, err := skipAXFRName(wire, offset)
	if err != nil || offset+10 > len(wire) {
		return 0, 0, fmt.Errorf("invalid DNS resource record")
	}
	rrType := binary.BigEndian.Uint16(wire[offset : offset+2])
	rrLength := int(binary.BigEndian.Uint16(wire[offset+8 : offset+10]))
	end := offset + 10 + rrLength
	if end > len(wire) {
		return 0, 0, fmt.Errorf("truncated DNS resource record")
	}
	return end, rrType, nil
}
