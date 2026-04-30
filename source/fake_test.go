/*
Copyright 2017 The Kubernetes Authors.

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

package source

import (
	"net"
	"net/netip"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	templatetest "sigs.k8s.io/external-dns/source/template/testutil"
)

// Validate that fakeSource implements Source.
var _ Source = &fakeSource{}

func TestFakeSourceEndpoints(t *testing.T) {
	sc, err := NewFakeSource(&Config{})
	require.NoError(t, err)

	endpoints, err := sc.Endpoints(t.Context())
	require.NoError(t, err)

	// One endpoint per known record type.
	assert.Len(t, endpoints, len(endpoint.KnownRecordTypes))

	byType := make(map[string]*endpoint.Endpoint, len(endpoints))
	for _, ep := range endpoints {
		byType[ep.RecordType] = ep
	}

	for _, rt := range endpoint.KnownRecordTypes {
		assert.Contains(t, byType, rt, "missing endpoint for record type %s", rt)
	}
}

func TestFakeSource_RecordTypes(t *testing.T) {
	tests := []struct {
		recordType string
		check      func(*testing.T, *endpoint.Endpoint)
	}{
		{
			recordType: endpoint.RecordTypeA,
			check: func(t *testing.T, ep *endpoint.Endpoint) {
				t.Helper()
				require.Len(t, ep.Targets, 1)
				ip := net.ParseIP(ep.Targets[0])
				assert.NotNil(t, ip, "A record target %q is not a valid IP", ep.Targets[0])
				assert.NotNil(t, ip.To4(), "A record target %q must be IPv4", ep.Targets[0])
				require.NotNil(t, ep.RefObject())
				assert.Equal(t, "Pod", ep.RefObject().Kind)
			},
		},
		{
			recordType: endpoint.RecordTypeAAAA,
			check: func(t *testing.T, ep *endpoint.Endpoint) {
				t.Helper()
				require.Len(t, ep.Targets, 1)
				addr, err := netip.ParseAddr(ep.Targets[0])
				require.NoError(t, err, "AAAA record target %q is not a valid IP address", ep.Targets[0])
				assert.True(t, addr.Is6() && !addr.Is4In6(), "AAAA record target %q must be native IPv6", ep.Targets[0])
			},
		},
		{
			recordType: endpoint.RecordTypeCNAME,
			check: func(t *testing.T, ep *endpoint.Endpoint) {
				t.Helper()
				require.Len(t, ep.Targets, 1)
				assert.True(t, strings.HasSuffix(ep.Targets[0], "."+defaultFQDNTemplate),
					"CNAME target %q should be under %s", ep.Targets[0], defaultFQDNTemplate)
			},
		},
		{
			recordType: endpoint.RecordTypeTXT,
			check: func(t *testing.T, ep *endpoint.Endpoint) {
				t.Helper()
				require.NotEmpty(t, ep.Targets)
			},
		},
		{
			recordType: endpoint.RecordTypeSRV,
			check: func(t *testing.T, ep *endpoint.Endpoint) {
				t.Helper()
				assert.True(t, strings.HasPrefix(ep.DNSName, "_sip._udp."), "SRV DNSName %q should start with _sip._udp.", ep.DNSName)
				require.Len(t, ep.Targets, 1)
				assert.True(t, ep.Targets.ValidateSRVRecord(), "SRV target %q is invalid", ep.Targets[0])
			},
		},
		{
			recordType: endpoint.RecordTypeNS,
			check: func(t *testing.T, ep *endpoint.Endpoint) {
				t.Helper()
				assert.Equal(t, defaultFQDNTemplate, ep.DNSName)
				require.Len(t, ep.Targets, 1)
				assert.True(t, strings.HasSuffix(ep.Targets[0], "."+defaultFQDNTemplate),
					"NS target %q should be under %s", ep.Targets[0], defaultFQDNTemplate)
			},
		},
		{
			recordType: endpoint.RecordTypePTR,
			check: func(t *testing.T, ep *endpoint.Endpoint) {
				t.Helper()
				assert.True(t, ep.ValidatePTRRecord(), "PTR record is invalid: %v", ep)
			},
		},
		{
			recordType: endpoint.RecordTypeMX,
			check: func(t *testing.T, ep *endpoint.Endpoint) {
				t.Helper()
				assert.Equal(t, defaultFQDNTemplate, ep.DNSName)
				require.Len(t, ep.Targets, 1)
				_, err := endpoint.NewMXRecord(ep.Targets[0])
				assert.NoError(t, err, "MX target %q is invalid", ep.Targets[0])
			},
		},
		{
			recordType: endpoint.RecordTypeNAPTR,
			check: func(t *testing.T, ep *endpoint.Endpoint) {
				t.Helper()
				assert.True(t, strings.HasPrefix(ep.DNSName, "_sip._udp."), "NAPTR DNSName %q should start with _sip._udp.", ep.DNSName)
				require.NotEmpty(t, ep.Targets)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.recordType, func(t *testing.T) {
			ep := mustGenerateEndpointForType(t, tt.recordType)
			tt.check(t, ep)
		})
	}
}

func TestFakeSource_FQDNTemplate(t *testing.T) {
	tests := []struct {
		name       string
		template   string
		wantDomain string
	}{
		{
			name:       "template expression",
			template:   "{{.Name}}.my-company.com",
			wantDomain: "fake.my-company.com",
		},
		{
			name:       "plain domain",
			template:   "my-company.com",
			wantDomain: "my-company.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc, err := NewFakeSource(&Config{
				TemplateEngine: templatetest.MustEngine(t, tt.template, "", "", false),
			})
			require.NoError(t, err)

			endpoints, err := sc.Endpoints(t.Context())
			require.NoError(t, err)
			require.NotEmpty(t, endpoints)

			for _, ep := range endpoints {
				if ep.RecordType == endpoint.RecordTypePTR {
					continue // PTR names are reverse-DNS, not under the configured domain
				}
				assert.True(t, strings.HasSuffix(ep.DNSName, "."+tt.wantDomain) || ep.DNSName == tt.wantDomain,
					"endpoint DNSName %q should be under %s", ep.DNSName, tt.wantDomain)
			}
		})
	}
}

func TestFakeSource_FQDNTemplate_MultiDomain(t *testing.T) {
	domains := []string{"one.example.com", "two.example.com", "three.example.com"}
	sc, err := NewFakeSource(&Config{
		TemplateEngine: templatetest.MustEngine(t, strings.Join(domains, ","), "", "", false),
	})
	require.NoError(t, err)

	endpoints, err := sc.Endpoints(t.Context())
	require.NoError(t, err)

	assert.Len(t, endpoints, len(domains)*len(endpoint.KnownRecordTypes))

	for _, ep := range endpoints {
		switch ep.RecordType {
		case endpoint.RecordTypeA:
			assert.NotEmpty(t, ep.Targets, "A record %s should have at least one target", ep.DNSName)
		case endpoint.RecordTypeAAAA:
			assert.NotEmpty(t, ep.Targets, "AAAA record %s should have at least one target", ep.DNSName)
		}
	}
}

func TestFakeSource_AddEventHandler(t *testing.T) {
	sc, err := NewFakeSource(&Config{})
	require.NoError(t, err)
	sc.AddEventHandler(t.Context(), func() {})
}

func TestFakeSource_NewFakeSource_TemplateError(t *testing.T) {
	// Verify that a template which parses successfully but fails at execution is caught during NewFakeSource.
	_, err := NewFakeSource(&Config{
		TemplateEngine: templatetest.MustEngine(t, "{{call .Name}}", "", "", false),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "rendering fqdn template")
}

func TestFakeSource_generateEndpointForType_UnknownType(t *testing.T) {
	sc, err := NewFakeSource(&Config{})
	require.NoError(t, err)
	fs := sc.(*fakeSource)
	_, err = fs.generateEndpointForType("UNKNOWN", defaultFQDNTemplate)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported record type")
}

func TestFakePodName_Truncation(t *testing.T) {
	long := strings.Repeat("a", 249) // "fake-" (5) + 249 = 254 > 253
	assert.Len(t, fakePodName(long), 253)
}

func mustGenerateEndpointForType(t *testing.T, recordType string) *endpoint.Endpoint {
	t.Helper()
	sc, err := NewFakeSource(&Config{})
	require.NoError(t, err)
	fs := sc.(*fakeSource)
	ep, err := fs.generateEndpointForType(recordType, defaultFQDNTemplate)
	require.NoError(t, err)
	require.NotNil(t, ep, "endpoint for type %s should not be nil", recordType)
	return ep
}
