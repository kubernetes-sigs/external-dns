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
	"fmt"
	"time"

	"github.com/miekg/dns"
)

// RFC 8945 section 5.3.1 requires a signed message at least every 100 messages.
const maxUnsignedAXFRMessages = 99

func signAndSendAXFRRequest(conn *dns.Conn, message *dns.Msg, provider dns.TsigProvider) (string, error) {
	wire, requestMAC, err := dns.TsigGenerateWithProvider(message, provider, "", false)
	if err != nil {
		_ = conn.Close()
		return "", fmt.Errorf("failed to sign AXFR request: %w", err)
	}
	if _, err := conn.Write(wire); err != nil {
		_ = conn.Close()
		return "", fmt.Errorf("failed to send AXFR request: %w", err)
	}
	return requestMAC, nil
}

func readSignedAXFR(conn *dns.Conn, id uint16, requestMAC string, provider dns.TsigProvider) chan *dns.Envelope {
	envelopes := make(chan *dns.Envelope)
	go func() {
		defer close(envelopes)
		defer conn.Close()

		first := true
		priorMAC := requestMAC
		var unsignedMessages [][]byte
		var unsignedRecords []dns.RR
		for {
			firstMessage := first
			if err := conn.SetReadDeadline(time.Now().Add(2 * time.Second)); err != nil {
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
				envelopes <- &dns.Envelope{Error: dns.ErrId}
				return
			}
			if message.Rcode != dns.RcodeSuccess {
				envelopes <- &dns.Envelope{Error: fmt.Errorf("bad AXFR rcode: %d", message.Rcode)}
				return
			}
			if first && (len(message.Answer) == 0 || message.Answer[0].Header().Rrtype != dns.TypeSOA) {
				envelopes <- &dns.Envelope{Error: dns.ErrSoa}
				return
			}

			ts := message.IsTsig()
			switch {
			case ts == nil:
				if first || isSOALast(message) {
					envelopes <- &dns.Envelope{Error: dns.ErrNoSig}
					return
				}
				unsignedMessages = append(unsignedMessages, wire)
				unsignedRecords = append(unsignedRecords, message.Answer...)
				if len(unsignedMessages) > maxUnsignedAXFRMessages {
					envelopes <- &dns.Envelope{Error: dns.ErrNoSig}
					return
				}
				first = false
				continue
			case first:
				err = dns.TsigVerifyWithProvider(wire, provider, requestMAC, false)
			default:
				chainedProvider := axfrTsigProvider{
					provider:         provider,
					priorMACLen:      len(priorMAC) / 2,
					unsignedMessages: unsignedMessages,
				}
				err = dns.TsigVerifyWithProvider(wire, chainedProvider, priorMAC, true)
			}
			if err != nil {
				envelopes <- &dns.Envelope{Error: err}
				return
			}
			priorMAC = ts.MAC
			unsignedMessages = nil
			first = false

			// Hold records until the next signed message authenticates intervening packets.
			unsignedRecords = append(unsignedRecords, message.Answer...)
			envelopes <- &dns.Envelope{RR: unsignedRecords}
			unsignedRecords = nil
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

// miekg omits unsigned AXFR packets from the MAC chain, so insert them after the prior MAC.
type axfrTsigProvider struct {
	provider         dns.TsigProvider
	priorMACLen      int
	unsignedMessages [][]byte
}

func (p axfrTsigProvider) Generate(message []byte, ts *dns.TSIG) ([]byte, error) {
	return p.provider.Generate(message, ts)
}

func (p axfrTsigProvider) Verify(message []byte, ts *dns.TSIG) error {
	prefix := 2 + p.priorMACLen
	macInputSize := len(message)
	for _, unsignedMessage := range p.unsignedMessages {
		macInputSize += len(unsignedMessage)
	}
	macInput := make([]byte, 0, macInputSize)
	macInput = append(macInput, message[:prefix]...)
	for _, unsignedMessage := range p.unsignedMessages {
		macInput = append(macInput, unsignedMessage...)
	}
	macInput = append(macInput, message[prefix:]...)
	return p.provider.Verify(macInput, ts)
}
