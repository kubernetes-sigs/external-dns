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

package wrappers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/source"
)

var _ source.Source = &acmeDelegationSource{}

func TestNewACMEDelegationSource(t *testing.T) {
	t.Run("invalid template returns error", func(t *testing.T) {
		src, err := NewACMEDelegationSource(testutils.NewMockSource(), "{{ .Hostname", nil, 0)
		require.Error(t, err)
		assert.ErrorContains(t, err, "--acme-cname-delegation-target-template")
		assert.Nil(t, src)
	})

	t.Run("valid template", func(t *testing.T) {
		src, err := NewACMEDelegationSource(testutils.NewMockSource(), "{{ .Hostname }}.acme.example.net", nil, 0)
		require.NoError(t, err)
		assert.NotNil(t, src)
	})
}

func TestACMEDelegationSource(t *testing.T) {
	defaultTemplate := "{{ .HostnameWithoutWildcard }}.acme.example.net"
	tests := []struct {
		name         string
		template     string
		domainFilter []string
		ttl          time.Duration
		endpoints    []*endpoint.Endpoint
		expected     []*endpoint.Endpoint
	}{
		{
			name: "A record produces delegation CNAME",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "_acme-challenge.web.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.example.com.acme.example.net"}},
			},
		},
		{
			name: "AAAA record produces delegation CNAME",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "v6.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "v6.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
				{DNSName: "_acme-challenge.v6.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"v6.example.com.acme.example.net"}},
			},
		},
		{
			name: "CNAME record produces delegation CNAME",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "alias.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.example.com"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "alias.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.example.com"}},
				{DNSName: "_acme-challenge.alias.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"alias.example.com.acme.example.net"}},
			},
		},
		{
			name: "A and AAAA for same hostname produce a single delegation CNAME",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeAAAA, Targets: endpoint.Targets{"2001:db8::1"}},
				{DNSName: "_acme-challenge.web.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.example.com.acme.example.net"}},
			},
		},
		{
			name: "wildcard hostname maps to challenge name without the wildcard label",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "*.app.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "*.app.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "_acme-challenge.app.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"app.example.com.acme.example.net"}},
			},
		},
		{
			name:     "Hostname template field renders the raw wildcard hostname",
			template: "{{ .Hostname }}.acme.example.net",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "*.app.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "*.app.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "_acme-challenge.app.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"*.app.example.com.acme.example.net"}},
			},
		},
		{
			name: "wildcard and apex hostname share a single delegation CNAME",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "*.app.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "app.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "*.app.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "app.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "_acme-challenge.app.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"app.example.com.acme.example.net"}},
			},
		},
		{
			name: "unsupported record types pass through unchanged",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "txt.example.com", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{"\"value\""}},
				{DNSName: "mx.example.com", RecordType: endpoint.RecordTypeMX, Targets: endpoint.Targets{"10 mail.example.com"}},
				{DNSName: "srv.example.com", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"0 5 5060 sip.example.com"}},
				{DNSName: "ns.example.com", RecordType: endpoint.RecordTypeNS, Targets: endpoint.Targets{"ns1.example.com"}},
				{DNSName: "2.49.168.192.in-addr.arpa", RecordType: endpoint.RecordTypePTR, Targets: endpoint.Targets{"web.example.com"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "txt.example.com", RecordType: endpoint.RecordTypeTXT, Targets: endpoint.Targets{"\"value\""}},
				{DNSName: "mx.example.com", RecordType: endpoint.RecordTypeMX, Targets: endpoint.Targets{"10 mail.example.com"}},
				{DNSName: "srv.example.com", RecordType: endpoint.RecordTypeSRV, Targets: endpoint.Targets{"0 5 5060 sip.example.com"}},
				{DNSName: "ns.example.com", RecordType: endpoint.RecordTypeNS, Targets: endpoint.Targets{"ns1.example.com"}},
				{DNSName: "2.49.168.192.in-addr.arpa", RecordType: endpoint.RecordTypePTR, Targets: endpoint.Targets{"web.example.com"}},
			},
		},
		{
			name:         "domain filter limits generation to matching hostnames",
			domainFilter: []string{"example.com"},
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "web.other.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.3"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "web.other.org", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.3"}},
				{DNSName: "_acme-challenge.web.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.example.com.acme.example.net"}},
			},
		},
		{
			name: "hostname with challenge prefix is not prefixed again",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "_acme-challenge.web.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.custom-acme.example.net"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "_acme-challenge.web.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.custom-acme.example.net"}},
			},
		},
		{
			name: "explicitly defined challenge endpoint suppresses generation and is preserved",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "_acme-challenge.web.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.custom-acme.example.net"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "_acme-challenge.web.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.custom-acme.example.net"}},
			},
		},
		{
			name: "configured TTL applies to generated CNAMEs",
			ttl:  300 * time.Second,
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("web.example.com", endpoint.RecordTypeA, 60, "192.168.49.2"),
			},
			expected: []*endpoint.Endpoint{
				endpoint.NewEndpointWithTTL("web.example.com", endpoint.RecordTypeA, 60, "192.168.49.2"),
				endpoint.NewEndpointWithTTL("_acme-challenge.web.example.com", endpoint.RecordTypeCNAME, 300, "web.example.com.acme.example.net"),
			},
		},
		{
			name:     "shared template functions are available",
			template: `{{ trimPrefix .HostnameWithoutWildcard "www." }}.acme.example.net`,
			endpoints: []*endpoint.Endpoint{
				{DNSName: "www.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "www.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "_acme-challenge.www.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"example.com.acme.example.net"}},
			},
		},
		{
			name:     "template execution failure skips generation without error",
			template: "{{ .DoesNotExist }}.acme.example.net",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
		},
		{
			name:     "empty template output skips generation without error",
			template: "{{ if false }}unreachable{{ end }}",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
		},
		{
			name:     "trailing dot in template output is trimmed",
			template: "{{ .HostnameWithoutWildcard }}.acme.example.net.",
			endpoints: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
			},
			expected: []*endpoint.Endpoint{
				{DNSName: "web.example.com", RecordType: endpoint.RecordTypeA, Targets: endpoint.Targets{"192.168.49.2"}},
				{DNSName: "_acme-challenge.web.example.com", RecordType: endpoint.RecordTypeCNAME, Targets: endpoint.Targets{"web.example.com.acme.example.net"}},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockSource := new(testutils.MockSource)
			mockSource.On("Endpoints").Return(tc.endpoints, nil)

			tmpl := tc.template
			if tmpl == "" {
				tmpl = defaultTemplate
			}
			src, err := NewACMEDelegationSource(mockSource, tmpl, tc.domainFilter, tc.ttl)
			require.NoError(t, err)

			result, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			require.Len(t, result, len(tc.expected))
			for i, ep := range result {
				assert.Equal(t, tc.expected[i].DNSName, ep.DNSName)
				assert.Equal(t, tc.expected[i].RecordType, ep.RecordType)
				assert.Equal(t, tc.expected[i].RecordTTL, ep.RecordTTL)
				assert.ElementsMatch(t, tc.expected[i].Targets, ep.Targets)
			}
		})
	}
}

func TestACMEDelegationSource_LabelAndRefCarryOver(t *testing.T) {
	ep := endpoint.NewEndpoint("web.example.com", endpoint.RecordTypeA, "192.168.49.2").
		WithLabel(endpoint.ResourceLabelKey, "ingress/default/web").
		WithRefObject(testutils.RefSource("ingress"))

	mockSource := testutils.NewMockSource(ep)
	src, err := NewACMEDelegationSource(mockSource, "{{ .HostnameWithoutWildcard }}.acme.example.net", nil, 0)
	require.NoError(t, err)

	result, err := src.Endpoints(t.Context())
	require.NoError(t, err)
	require.Len(t, result, 2)

	cname := result[1]
	assert.Equal(t, "_acme-challenge.web.example.com", cname.DNSName)
	assert.Equal(t, "ingress/default/web", cname.Labels[endpoint.ResourceLabelKey])
	require.Len(t, cname.RefObjects(), 1)
	assert.Equal(t, "ingress", cname.RefObjects()[0].Source())
}

func TestACMEDelegationSource_ErrorPropagation(t *testing.T) {
	mockSource := new(testutils.MockSource)
	mockSource.On("Endpoints").Return([]*endpoint.Endpoint(nil), assert.AnError)

	src, err := NewACMEDelegationSource(mockSource, "{{ .HostnameWithoutWildcard }}.acme.example.net", nil, 0)
	require.NoError(t, err)

	result, err := src.Endpoints(t.Context())
	require.ErrorIs(t, err, assert.AnError)
	assert.Nil(t, result)
}

func TestACMEDelegationSource_AddEventHandler(t *testing.T) {
	mockSource := testutils.NewMockSource()

	src, err := NewACMEDelegationSource(mockSource, "{{ .HostnameWithoutWildcard }}.acme.example.net", nil, 0)
	require.NoError(t, err)
	src.AddEventHandler(t.Context(), func() {})

	mockSource.AssertNumberOfCalls(t, "AddEventHandler", 1)
}
