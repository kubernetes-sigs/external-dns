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
	"strings"
	"testing"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/plan"
)

const (
	tlsaDigest      = "0b9fa5a59eed715c26c1020c711b4f6ec42d58b0015e14337a39dad301c5afc3"
	tlsaCanonical   = "3 1 1 " + tlsaDigest
	tlsaRecordName  = "_25._tcp.mail.foo.com"
	tlsaRecordFqdn  = tlsaRecordName + "."
	tlsaWireUpper   = "3 1 1 0B9FA5A59EED715C26C1020C711B4F6EC42D58B0015E14337A39DAD301C5AFC3"
	tlsaWithColons  = "3 1 1 0B9FA5A5:9EED715C:26C1020C:711B4F6E:C42D58B0:015E1433:7A39DAD3:01C5AFC3"
	tlsaInvalidData = "3 1 1 not-hex"
)

func TestRfc2136GetRecordsTLSA(t *testing.T) {
	stub := newStub()
	err := stub.setOutput([]string{
		tlsaRecordName + " 3600 IN TLSA " + tlsaWireUpper,
	})
	require.NoError(t, err)

	p, err := createRfc2136StubProvider(stub, "foo.com")
	require.NoError(t, err)

	recs, err := p.Records(t.Context())
	require.NoError(t, err)

	require.Len(t, recs, 1)
	assert.Equal(t, tlsaRecordName, recs[0].DNSName)
	assert.Equal(t, endpoint.RecordTypeTLSA, recs[0].RecordType)
	assert.Equal(t, endpoint.TTL(3600), recs[0].RecordTTL)
	// The nameserver answered in uppercase; the target is normalised so it
	// compares equal to the one the source asked for.
	assert.Equal(t, []string{tlsaCanonical}, []string(recs[0].Targets))
}

func TestRfc2136GetRecordsTLSAMultipleTargets(t *testing.T) {
	otherDigest := strings.Repeat("ab", 32)
	stub := newStub()
	err := stub.setOutput([]string{
		tlsaRecordName + " 3600 IN TLSA " + tlsaWireUpper,
		tlsaRecordName + " 3600 IN TLSA 3 1 1 " + otherDigest,
	})
	require.NoError(t, err)

	p, err := createRfc2136StubProvider(stub, "foo.com")
	require.NoError(t, err)

	recs, err := p.Records(t.Context())
	require.NoError(t, err)

	require.Len(t, recs, 1, "both targets belong to a single endpoint")
	assert.ElementsMatch(t, []string{tlsaCanonical, "3 1 1 " + otherDigest}, []string(recs[0].Targets))
}

func TestRfc2136TLSACreation(t *testing.T) {
	stub := newStub()
	p, err := createRfc2136StubProvider(stub, "foo.com")
	require.NoError(t, err)

	err = p.ApplyChanges(t.Context(), &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    tlsaRecordName,
				RecordType: endpoint.RecordTypeTLSA,
				Targets:    endpoint.Targets{tlsaWithColons},
			},
		},
	})
	require.NoError(t, err)

	require.Len(t, stub.createMsgs, 1)
	createMsg := strings.Join(strings.Fields(getSortedChanges(stub.createMsgs)[0]), " ")
	assert.Contains(t, createMsg, tlsaRecordFqdn+" 300 IN TLSA "+tlsaCanonical)
}

func TestRfc2136TLSADeletion(t *testing.T) {
	stub := newStub()
	p, err := createRfc2136StubProvider(stub, "foo.com")
	require.NoError(t, err)

	err = p.ApplyChanges(t.Context(), &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    tlsaRecordName,
				RecordType: endpoint.RecordTypeTLSA,
				Targets:    endpoint.Targets{tlsaWireUpper},
			},
		},
	})
	require.NoError(t, err)

	require.Len(t, stub.updateMsgs, 1)
	deleteMsg := strings.Join(strings.Fields(getSortedChanges(stub.updateMsgs)[0]), " ")
	assert.Contains(t, deleteMsg, tlsaRecordFqdn+" 0 NONE TLSA "+tlsaCanonical)
}

func TestRfc2136AdjustEndpointsNormalisesTLSA(t *testing.T) {
	stub := newStub()
	p, err := createRfc2136StubProvider(stub, "foo.com")
	require.NoError(t, err)

	eps, err := p.AdjustEndpoints([]*endpoint.Endpoint{
		endpoint.NewEndpoint(tlsaRecordName, endpoint.RecordTypeTLSA, tlsaWithColons, tlsaWireUpper),
		endpoint.NewEndpoint("v1.foo.com", endpoint.RecordTypeA, "1.2.3.4"),
	})
	require.NoError(t, err)

	require.Len(t, eps, 2)
	assert.Equal(t, []string{tlsaCanonical, tlsaCanonical}, []string(eps[0].Targets))
	assert.Equal(t, []string{"1.2.3.4"}, []string(eps[1].Targets), "non-TLSA targets are left alone")
}

func TestRfc2136AdjustEndpointsLeavesMalformedTLSAUnchanged(t *testing.T) {
	stub := newStub()
	p, err := createRfc2136StubProvider(stub, "foo.com")
	require.NoError(t, err)

	hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

	eps, err := p.AdjustEndpoints([]*endpoint.Endpoint{
		endpoint.NewEndpoint(tlsaRecordName, endpoint.RecordTypeTLSA, tlsaInvalidData),
	})
	require.NoError(t, err)

	require.Len(t, eps, 1)
	assert.Equal(t, []string{tlsaInvalidData}, []string(eps[0].Targets))
	logtest.TestHelperLogContains("could not parse TLSA target", hook, t)
}

func TestRfc2136MalformedTLSATargetIsRejected(t *testing.T) {
	stub := newStub()
	p, err := createRfc2136StubProvider(stub, "foo.com")
	require.NoError(t, err)

	r, ok := p.(*rfc2136Provider)
	require.True(t, ok)

	ep := &endpoint.Endpoint{
		DNSName:    tlsaRecordName,
		RecordType: endpoint.RecordTypeTLSA,
		Targets:    endpoint.Targets{tlsaInvalidData},
	}

	m := new(dns.Msg)
	m.SetUpdate("foo.com.")

	err = r.AddRecord(m, ep)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not valid hex")
	assert.Empty(t, m.Ns, "a malformed target must not reach the update message")

	err = r.RemoveRecord(m, ep)
	require.Error(t, err)
	assert.ErrorContains(t, err, "not valid hex")
	assert.Empty(t, m.Ns)
}

func TestTlsaTargetFallsBackToVerbatim(t *testing.T) {
	hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

	// A nameserver is not expected to answer with a usage this large, but the
	// record must still be surfaced rather than dropped.
	rr := &dns.TLSA{
		Hdr:          dns.RR_Header{Name: tlsaRecordFqdn, Rrtype: dns.TypeTLSA, Class: dns.ClassINET},
		Usage:        9,
		Selector:     1,
		MatchingType: 1,
		Certificate:  strings.ToUpper(tlsaDigest),
	}

	assert.Equal(t, "9 1 1 "+strings.ToUpper(tlsaDigest), tlsaTarget(rr))
	logtest.TestHelperLogContains("could not parse TLSA record", hook, t)
}
