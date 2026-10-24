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
	"errors"
	"net"
	"testing"
	"time"

	"github.com/bodgit/tsig"
	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	axfrTestKeyName = "key.example.com."
	axfrTestSecret  = "c2VjcmV0"
)

func signedAXFRStart(t *testing.T) (uint16, string, []byte) {
	t.Helper()

	query := new(dns.Msg)
	query.SetAxfr("example.com.")
	query.SetTsig(axfrTestKeyName, dns.HmacSHA256, clockSkew, time.Now().Unix())
	_, requestMAC, err := dns.TsigGenerate(query, axfrTestSecret, "", false)
	require.NoError(t, err)

	soa, err := dns.NewRR("example.com. 300 IN SOA ns.example.com. hostmaster.example.com. 1 7200 3600 1209600 300")
	require.NoError(t, err)
	response := new(dns.Msg)
	response.SetReply(query)
	response.Answer = []dns.RR{soa}
	response.SetTsig(axfrTestKeyName, dns.HmacSHA256, clockSkew, time.Now().Unix())
	wire, _, err := dns.TsigGenerate(response, axfrTestSecret, requestMAC, false)
	require.NoError(t, err)

	return query.Id, requestMAC, wire
}

func packAXFRMessage(t *testing.T, message *dns.Msg) []byte {
	t.Helper()
	wire, err := message.Pack()
	require.NoError(t, err)
	return wire
}

func axfrResponse(id uint16, answers ...dns.RR) *dns.Msg {
	return &dns.Msg{
		MsgHdr: dns.MsgHdr{Id: id, Response: true},
		Answer: answers,
	}
}

func newAXFRTestQuery() *dns.Msg {
	query := new(dns.Msg)
	query.SetAxfr("example.com.")
	query.SetTsig(axfrTestKeyName, dns.HmacSHA256, clockSkew, time.Now().Unix())
	return query
}

func readAXFRPackets(t *testing.T, id uint16, requestMAC string, packets ...[]byte) ([]dns.RR, error) {
	t.Helper()
	client, server := net.Pipe()
	t.Cleanup(func() {
		_ = server.Close()
	})
	go func() {
		for _, packet := range packets {
			frame := make([]byte, 2, len(packet)+2)
			binary.BigEndian.PutUint16(frame, uint16(len(packet)))
			frame = append(frame, packet...)
			if _, err := server.Write(frame); err != nil {
				return
			}
		}
	}()

	var records []dns.RR
	var transferErr error
	for envelope := range readSignedAXFR(&dns.Conn{Conn: client}, id, requestMAC, axfrTestSecret, tsig.HMAC{axfrTestKeyName: axfrTestSecret}) {
		if envelope.Error != nil {
			transferErr = envelope.Error
			continue
		}
		records = append(records, envelope.RR...)
	}
	return records, transferErr
}

func TestReadSignedAXFRRejectsInvalidMessages(t *testing.T) {
	id, requestMAC, first := signedAXFRStart(t)
	soa, err := dns.NewRR("example.com. 300 IN SOA ns.example.com. hostmaster.example.com. 1 7200 3600 1209600 300")
	require.NoError(t, err)
	aRecord, err := dns.NewRR("www.example.com. 300 IN A 192.0.2.1")
	require.NoError(t, err)

	wrongID := axfrResponse(id+1, soa)
	badRcode := axfrResponse(id, soa)
	badRcode.Rcode = dns.RcodeNameError
	nonSOA := axfrResponse(id, aRecord)
	empty := axfrResponse(id)
	unsignedFinal := axfrResponse(id, soa)
	badSignature := axfrResponse(id, soa)
	badSignature.SetTsig(axfrTestKeyName, dns.HmacSHA256, clockSkew, time.Now().Unix())
	badSignatureWire, _, err := dns.TsigGenerate(badSignature, "d3Jvbmc=", "", false)
	require.NoError(t, err)
	malformedDNS := packAXFRMessage(t, axfrResponse(id, aRecord))
	malformedDNS = malformedDNS[:len(malformedDNS)-1]

	tests := []struct {
		name       string
		packets    [][]byte
		wantErr    error
		wantErrMsg string
	}{
		{name: "malformed DNS message", packets: [][]byte{malformedDNS}},
		{name: "mismatched ID", packets: [][]byte{packAXFRMessage(t, wrongID)}, wantErr: dns.ErrId},
		{name: "bad response code", packets: [][]byte{packAXFRMessage(t, badRcode)}, wantErrMsg: "bad AXFR rcode"},
		{name: "first record is not SOA", packets: [][]byte{packAXFRMessage(t, nonSOA)}, wantErr: dns.ErrSoa},
		{name: "first response has no answer", packets: [][]byte{packAXFRMessage(t, empty)}, wantErr: dns.ErrSoa},
		{name: "first response is unsigned", packets: [][]byte{packAXFRMessage(t, axfrResponse(id, soa))}, wantErr: dns.ErrNoSig},
		{name: "later response has no answers", packets: [][]byte{first, packAXFRMessage(t, empty)}, wantErr: dns.ErrSoa},
		{name: "unsigned final SOA", packets: [][]byte{first, packAXFRMessage(t, unsignedFinal)}, wantErr: dns.ErrNoSig},
		{name: "invalid signature", packets: [][]byte{first, badSignatureWire}, wantErrMsg: "dns: bad signature"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, transferErr := readAXFRPackets(t, id, requestMAC, test.packets...)
			require.Error(t, transferErr)
			if test.wantErr != nil {
				assert.ErrorIs(t, transferErr, test.wantErr)
			}
			if test.wantErrMsg != "" {
				assert.ErrorContains(t, transferErr, test.wantErrMsg)
			}
		})
	}

	t.Run("unsigned message limit", func(t *testing.T) {
		packets := make([][]byte, 1, maxUnsignedAXFRMessages+2)
		packets[0] = first
		middle := packAXFRMessage(t, axfrResponse(id, aRecord))
		for range maxUnsignedAXFRMessages + 1 {
			packets = append(packets, middle)
		}
		_, transferErr := readAXFRPackets(t, id, requestMAC, packets...)
		assert.ErrorIs(t, transferErr, dns.ErrNoSig)
	})
}

type readDeadlineErrorConn struct {
	net.Conn
	err error
}

func (c readDeadlineErrorConn) SetReadDeadline(time.Time) error {
	return c.err
}

type readErrorConn struct {
	net.Conn
	err error
}

func (c readErrorConn) Read([]byte) (int, error) {
	return 0, c.err
}

func TestReadSignedAXFRReadErrors(t *testing.T) {
	t.Run("setting read deadline", func(t *testing.T) {
		wantErr := errors.New("deadline unavailable")
		client, server := net.Pipe()
		defer server.Close()
		conn := &dns.Conn{Conn: readDeadlineErrorConn{Conn: client, err: wantErr}}
		envelope := <-readSignedAXFR(conn, 1, "", axfrTestSecret, tsig.HMAC{})
		assert.ErrorIs(t, envelope.Error, wantErr)
	})

	t.Run("reading message", func(t *testing.T) {
		client, server := net.Pipe()
		defer server.Close()
		wantErr := errors.New("read failed")
		conn := &dns.Conn{Conn: readErrorConn{Conn: client, err: wantErr}}
		envelope := <-readSignedAXFR(conn, 1, "", axfrTestSecret, tsig.HMAC{})
		assert.ErrorIs(t, envelope.Error, wantErr)
	})
}

func TestAXFRTsigHelpersRejectMalformedInput(t *testing.T) {
	t.Run("invalid prior MAC", func(t *testing.T) {
		_, err := axfrTsigMACInput("not-hex", nil, nil, &dns.TSIG{})
		assert.ErrorContains(t, err, "invalid prior TSIG MAC")
	})

	t.Run("prior MAC too long", func(t *testing.T) {
		_, err := axfrTsigMACInput(hex.EncodeToString(make([]byte, 0x10000)), nil, nil, &dns.TSIG{})
		assert.ErrorContains(t, err, "prior TSIG MAC is too long")
	})

	t.Run("valid prior MAC", func(t *testing.T) {
		input, err := axfrTsigMACInput("0102", [][]byte{{3}}, []byte{4}, &dns.TSIG{TimeSigned: 5, Fudge: 6})
		require.NoError(t, err)
		assert.Equal(t, []byte{0, 2, 1, 2, 3, 4, 0, 0, 0, 0, 0, 5, 0, 6}, input)
	})

	t.Run("strip TSIG errors", func(t *testing.T) {
		wireWithAdditional := make([]byte, 12)
		binary.BigEndian.PutUint16(wireWithAdditional[10:], 1)
		wireWithQuestion := make([]byte, 12)
		binary.BigEndian.PutUint16(wireWithQuestion[4:], 1)
		binary.BigEndian.PutUint16(wireWithQuestion[10:], 1)
		wireWithAnswer := make([]byte, 12)
		binary.BigEndian.PutUint16(wireWithAnswer[6:], 1)
		binary.BigEndian.PutUint16(wireWithAnswer[10:], 1)
		wireWithTrailingTSIGByte := make([]byte, 12)
		binary.BigEndian.PutUint16(wireWithTrailingTSIGByte[10:], 1)
		wireWithTrailingTSIGByte = append(wireWithTrailingTSIGByte, minimalTSIGWire()...)
		wireWithTrailingTSIGByte = append(wireWithTrailingTSIGByte, 0)

		tests := []struct {
			name string
			wire []byte
		}{
			{name: "short header", wire: make([]byte, 11)},
			{name: "no additional record", wire: make([]byte, 12)},
			{name: "invalid question", wire: wireWithQuestion},
			{name: "invalid answer", wire: wireWithAnswer},
			{name: "invalid TSIG record", wire: wireWithAdditional},
			{name: "TSIG is not final", wire: wireWithTrailingTSIGByte},
			{name: "not TSIG", wire: minimalRRWire(dns.TypeA)},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				_, err := stripAXFRTsig(test.wire, &dns.TSIG{})
				assert.Error(t, err)
			})
		}
	})

	t.Run("skip names", func(t *testing.T) {
		tests := []struct {
			name string
			wire []byte
			err  bool
		}{
			{name: "empty", wire: nil, err: true},
			{name: "truncated pointer", wire: []byte{0xc0}, err: true},
			{name: "compressed name", wire: []byte{0xc0, 0}, err: false},
			{name: "reserved label type", wire: []byte{0x40}, err: true},
			{name: "truncated label", wire: []byte{2, 'a'}, err: true},
			{name: "root name", wire: []byte{0}, err: false},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				_, err := skipAXFRName(test.wire, 0)
				if test.err {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})

	t.Run("skip resource records", func(t *testing.T) {
		tests := []struct {
			name string
			wire []byte
		}{
			{name: "invalid name", wire: nil},
			{name: "short record header", wire: []byte{0}},
			{name: "truncated data", wire: minimalRRWire(dns.TypeA, 4)},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				_, _, err := skipAXFRRR(test.wire, 0)
				assert.Error(t, err)
			})
		}
	})
}

func minimalTSIGWire() []byte {
	return minimalRRWire(dns.TypeTSIG)
}

func minimalRRWire(rrType uint16, rdlength ...uint16) []byte {
	wire := []byte{0}
	wire = binary.BigEndian.AppendUint16(wire, rrType)
	wire = binary.BigEndian.AppendUint16(wire, dns.ClassANY)
	wire = binary.BigEndian.AppendUint32(wire, 0)
	var length uint16
	if len(rdlength) > 0 {
		length = rdlength[0]
	}
	wire = binary.BigEndian.AppendUint16(wire, length)
	if length > 0 {
		wire = append(wire, make([]byte, int(length)-1)...)
	}
	return wire
}

func TestSignedIncomeTransferSetupErrors(t *testing.T) {
	t.Run("TLS setup", func(t *testing.T) {
		r := &rfc2136Provider{tlsConfig: TLSConfig{UseTLS: true, CAFilePath: "/does/not/exist"}}
		_, err := r.signedIncomeTransfer(new(dns.Msg), "127.0.0.1:53")
		assert.ErrorContains(t, err, "error setting up TLS")
	})

	t.Run("dial", func(t *testing.T) {
		r := &rfc2136Provider{}
		_, err := r.signedIncomeTransfer(new(dns.Msg), "127.0.0.1:1")
		assert.ErrorContains(t, err, "failed to connect for transfer")
	})

	t.Run("signing", func(t *testing.T) {
		listener, err := net.Listen("tcp", "127.0.0.1:0")
		require.NoError(t, err)
		defer listener.Close()
		accepted := make(chan net.Conn, 1)
		go func() {
			conn, acceptErr := listener.Accept()
			if acceptErr == nil {
				accepted <- conn
			}
		}()

		r := &rfc2136Provider{tsigSecret: "invalid-base64"}
		_, err = r.signedIncomeTransfer(newAXFRTestQuery(), listener.Addr().String())
		assert.ErrorContains(t, err, "failed to sign AXFR request")
		select {
		case conn := <-accepted:
			_ = conn.Close()
		case <-time.After(time.Second):
			t.Fatal("server did not accept AXFR connection")
		}
	})
}

func TestSignAndSendAXFRRequestErrors(t *testing.T) {
	t.Run("signing", func(t *testing.T) {
		client, server := net.Pipe()
		defer server.Close()
		conn := &dns.Conn{Conn: client}
		_, err := signAndSendAXFRRequest(conn, newAXFRTestQuery(), "invalid-base64")
		assert.ErrorContains(t, err, "failed to sign AXFR request")
		_, writeErr := conn.Write([]byte{1})
		assert.Error(t, writeErr)
	})

	t.Run("writing", func(t *testing.T) {
		client, server := net.Pipe()
		_ = server.Close()
		conn := &dns.Conn{Conn: client}
		_, err := signAndSendAXFRRequest(conn, newAXFRTestQuery(), axfrTestSecret)
		assert.ErrorContains(t, err, "failed to send AXFR request")
	})
}

func TestSignedIncomeTransferReturnsSignedStream(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer listener.Close()
	serverErr := make(chan error, 1)
	go func() {
		rawConn, acceptErr := listener.Accept()
		if acceptErr != nil {
			serverErr <- acceptErr
			return
		}
		defer rawConn.Close()
		conn := &dns.Conn{Conn: rawConn}
		wire, readErr := conn.ReadMsgHeader(nil)
		if readErr != nil {
			serverErr <- readErr
			return
		}
		request := new(dns.Msg)
		if unpackErr := request.Unpack(wire); unpackErr != nil {
			serverErr <- unpackErr
			return
		}
		response := new(dns.Msg)
		response.SetReply(request)
		soa, rrErr := dns.NewRR("example.com. 300 IN SOA ns.example.com. hostmaster.example.com. 1 7200 3600 1209600 300")
		if rrErr != nil {
			serverErr <- rrErr
			return
		}
		response.Answer = []dns.RR{soa, soa}
		response.SetTsig(axfrTestKeyName, dns.HmacSHA256, clockSkew, time.Now().Unix())
		responseWire, _, signErr := dns.TsigGenerate(response, axfrTestSecret, request.IsTsig().MAC, false)
		if signErr != nil {
			serverErr <- signErr
			return
		}
		frame := make([]byte, 2, len(responseWire)+2)
		binary.BigEndian.PutUint16(frame, uint16(len(responseWire)))
		frame = append(frame, responseWire...)
		_, writeErr := rawConn.Write(frame)
		serverErr <- writeErr
	}()

	query := newAXFRTestQuery()
	r := &rfc2136Provider{tsigKeyName: axfrTestKeyName, tsigSecret: axfrTestSecret}
	envelopes, err := r.IncomeTransfer(query, listener.Addr().String())
	require.NoError(t, err)
	var records []dns.RR
	for envelope := range envelopes {
		require.NoError(t, envelope.Error)
		records = append(records, envelope.RR...)
	}
	require.NoError(t, <-serverErr)
	assert.Len(t, records, 2)
}

func TestIncomeTransferUnsignedTLSSetupError(t *testing.T) {
	r := &rfc2136Provider{
		axfrInsecure: true,
		tlsConfig:    TLSConfig{UseTLS: true, CAFilePath: "/does/not/exist"},
	}
	_, err := r.IncomeTransfer(newAXFRTestQuery(), "127.0.0.1:53")
	assert.ErrorContains(t, err, "error setting up TLS")
}

func TestVerifySignedAXFRMessageErrors(t *testing.T) {
	ts := &dns.TSIG{}
	assert.Error(t, verifySignedAXFRMessage(make([]byte, 12), ts, "", nil, tsig.HMAC{}))

	wire := make([]byte, 12)
	binary.BigEndian.PutUint16(wire[10:], 1)
	wire = append(wire, minimalTSIGWire()...)
	assert.ErrorContains(t, verifySignedAXFRMessage(wire, ts, "not-hex", nil, tsig.HMAC{}), "invalid prior TSIG MAC")
}
