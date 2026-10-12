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

package cloudflare

import (
	"encoding/json"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const (
	naptrName   = "sip.example.com"
	naptrTarget = `10 20 "S" "SIP+D2U" "" _sip._udp.example.com.`
)

func naptrRecordResponse() dns.RecordResponse {
	return dns.RecordResponse{
		Name:    naptrName,
		TTL:     300,
		Type:    dns.RecordResponseTypeNAPTR,
		Content: naptrTarget,
	}
}

// The API rejects NAPTR sent as content, so writes must send structured data.
func assertNAPTRDataOnly(t *testing.T, body any) {
	t.Helper()

	raw, err := json.Marshal(body)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(raw, &decoded))

	assert.NotContains(t, decoded, "content", "NAPTR writes must not send a content field")

	data, ok := decoded["data"].(map[string]any)
	require.True(t, ok, "NAPTR writes must send a data object, got: %s", raw)
	assert.InDelta(t, 10, data["order"], 0)
	assert.InDelta(t, 20, data["preference"], 0)
	assert.Equal(t, "S", data["flags"])
	assert.Equal(t, "SIP+D2U", data["service"])
	assert.Empty(t, data["regex"])
	assert.Equal(t, "_sip._udp.example.com", data["replacement"], "replacement must be sent without the trailing dot")
	assert.Equal(t, "NAPTR", decoded["type"])
	assert.Equal(t, naptrName, decoded["name"])
}

func TestGetCreateDNSRecordParamNAPTRUsesData(t *testing.T) {
	t.Parallel()

	params, err := getCreateDNSRecordParam("zone-1", &cloudFlareChange{
		Action:         cloudFlareCreate,
		ResourceRecord: naptrRecordResponse(),
	})
	require.NoError(t, err)
	assert.Equal(t, "zone-1", params.ZoneID.Value)

	body, ok := params.Body.(dns.NAPTRRecordParam)
	require.True(t, ok, "expected a NAPTRRecordParam body, got %T", params.Body)
	assertNAPTRDataOnly(t, body)
}

func TestGetUpdateDNSRecordParamNAPTRUsesData(t *testing.T) {
	t.Parallel()

	params, err := getUpdateDNSRecordParam("zone-1", cloudFlareChange{
		Action:         cloudFlareUpdate,
		ResourceRecord: naptrRecordResponse(),
	})
	require.NoError(t, err)
	assert.Equal(t, "zone-1", params.ZoneID.Value)

	body, ok := params.Body.(dns.NAPTRRecordParam)
	require.True(t, ok, "expected a NAPTRRecordParam body, got %T", params.Body)
	assertNAPTRDataOnly(t, body)
}

func TestBuildBatchPostParamNAPTRUsesData(t *testing.T) {
	t.Parallel()

	param, ok := buildBatchPostParam(naptrRecordResponse())
	require.True(t, ok)

	body, isNAPTR := param.(dns.NAPTRRecordParam)
	require.True(t, isNAPTR, "expected a NAPTRRecordParam, got %T", param)
	assertNAPTRDataOnly(t, body)
}

func TestBuildBatchPutParamNAPTRUsesData(t *testing.T) {
	t.Parallel()

	param, ok := buildBatchPutParam("record-id", naptrRecordResponse())
	require.True(t, ok)

	put, isNAPTR := param.(dns.BatchPutNAPTRRecordParam)
	require.True(t, isNAPTR, "expected a BatchPutNAPTRRecordParam, got %T", param)
	assert.Equal(t, "record-id", put.ID.Value)
	assertNAPTRDataOnly(t, put.NAPTRRecordParam)
}

func TestMalformedNAPTRTargetIsRejected(t *testing.T) {
	t.Parallel()

	bad := dns.RecordResponse{
		Name:    naptrName,
		TTL:     300,
		Type:    dns.RecordResponseTypeNAPTR,
		Content: "10 20",
	}

	_, err := getCreateDNSRecordParam("zone-1", &cloudFlareChange{ResourceRecord: bad})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid NAPTR target")

	_, err = getUpdateDNSRecordParam("zone-1", cloudFlareChange{ResourceRecord: bad})
	require.Error(t, err)

	_, ok := buildBatchPostParam(bad)
	assert.False(t, ok, "a malformed target must not produce a batch post")

	_, ok = buildBatchPutParam("record-id", bad)
	assert.False(t, ok, "a malformed target must not produce a batch put")
}

func TestBuildNAPTRRecordParamReplacement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		target          string
		wantReplacement string
	}{
		{
			name:            "trailing dot is stripped for the API",
			target:          naptrTarget,
			wantReplacement: "_sip._udp.example.com",
		},
		{
			name:            "root replacement of a terminal rule is preserved",
			target:          `100 50 "U" "E2U+sip" "!^.*$!sip:info@example.com!" .`,
			wantReplacement: ".",
		},
		{
			name:            "replacement without trailing dot is accepted",
			target:          `10 20 "S" "SIP+D2U" "" _sip._udp.example.com`,
			wantReplacement: "_sip._udp.example.com",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := naptrRecordResponse()
			r.Content = tc.target
			body, err := buildNAPTRRecordParam(r)
			require.NoError(t, err)
			assert.Equal(t, tc.wantReplacement, body.Data.Value.Replacement.Value)
		})
	}
}

// Without normalisation a record never compares equal to itself and every sync updates.
func TestEndpointTargetFromCloudflareRecordNormalisesNAPTR(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		record dns.RecordResponse
		want   string
	}{
		{
			name: "structured data is rendered canonical with quoting and trailing dot",
			record: dns.RecordResponse{
				Name: naptrName,
				Type: dns.RecordResponseTypeNAPTR,
				Data: dns.NAPTRRecordData{
					Order:       10,
					Preference:  20,
					Flags:       "S",
					Service:     "SIP+D2U",
					Regex:       "",
					Replacement: "_sip._udp.example.com",
				},
				Content: `10 20 S SIP+D2U  _sip._udp.example.com`,
			},
			want: naptrTarget,
		},
		{
			name: "terminal rule with regexp and root replacement",
			record: dns.RecordResponse{
				Name: naptrName,
				Type: dns.RecordResponseTypeNAPTR,
				Data: dns.NAPTRRecordData{
					Order:       100,
					Preference:  50,
					Flags:       "U",
					Service:     "E2U+sip",
					Regex:       "!^.*$!sip:info@example.com!",
					Replacement: ".",
				},
			},
			want: `100 50 "U" "E2U+sip" "!^.*$!sip:info@example.com!" .`,
		},
		{
			name: "missing structured data falls back to content verbatim",
			record: dns.RecordResponse{
				Name:    naptrName,
				Type:    dns.RecordResponseTypeNAPTR,
				Content: "some raw content",
			},
			want: "some raw content",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, endpointTargetFromCloudflareRecord(tc.record))
		})
	}
}

func TestAdjustEndpointsCanonicalizesNAPTR(t *testing.T) {
	t.Parallel()

	p := &CloudFlareProvider{}
	adjusted, err := p.AdjustEndpoints([]*endpoint.Endpoint{
		{
			DNSName:    "sip.bar.com",
			RecordType: endpoint.RecordTypeNAPTR,
			Targets: endpoint.Targets{
				`10  20  "S"  "SIP+D2U"  ""  _sip._udp.bar.com`,
				"10 20",
			},
		},
		{
			DNSName:    "bar.com",
			RecordType: endpoint.RecordTypeA,
			Targets:    endpoint.Targets{"1.2.3.4"},
		},
	})
	require.NoError(t, err)
	require.Len(t, adjusted, 2)
	assert.Equal(t, `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`, adjusted[0].Targets[0],
		"valid NAPTR targets must be canonicalized to the read-back form")
	assert.Equal(t, "10 20", adjusted[0].Targets[1],
		"a malformed NAPTR target must be left unchanged so it surfaces on apply")
	assert.Equal(t, "1.2.3.4", adjusted[1].Targets[0])
}

func TestCloudflareNAPTRRoundTrip(t *testing.T) {
	client := NewMockCloudFlareClient()
	p := &CloudFlareProvider{Client: client}
	ctx := t.Context()

	ep := &endpoint.Endpoint{
		DNSName:    "bar.com",
		RecordType: endpoint.RecordTypeNAPTR,
		Targets:    endpoint.Targets{`10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`},
		RecordTTL:  endpoint.TTL(defaultTTL),
		Labels:     endpoint.Labels{},
	}

	require.NoError(t, p.ApplyChanges(ctx, &plan.Changes{Create: []*endpoint.Endpoint{ep}}))

	records, err := p.Records(ctx)
	require.NoError(t, err)
	var got *endpoint.Endpoint
	for _, r := range records {
		if r.RecordType == endpoint.RecordTypeNAPTR && r.DNSName == "bar.com" {
			got = r
			break
		}
	}
	require.NotNil(t, got, "created NAPTR record must be read back")
	require.Len(t, got.Targets, 1)
	assert.Equal(t, `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`, got.Targets[0])

	require.NoError(t, p.ApplyChanges(ctx, &plan.Changes{Delete: []*endpoint.Endpoint{got}}))
	records, err = p.Records(ctx)
	require.NoError(t, err)
	for _, r := range records {
		assert.NotEqual(t, endpoint.RecordTypeNAPTR, r.RecordType, "deleted NAPTR record must not be read back")
	}
}

// replanAfterApply applies one NAPTR endpoint, reads it back and re-plans; an empty plan proves the target converged.
func replanAfterApply(t *testing.T, recordType, dnsName, target string) *plan.Plan {
	t.Helper()
	client := NewMockCloudFlareClient()
	p := &CloudFlareProvider{Client: client}
	ctx := t.Context()

	desired := func() []*endpoint.Endpoint {
		return []*endpoint.Endpoint{{
			DNSName:    dnsName,
			RecordType: recordType,
			Targets:    endpoint.Targets{target},
			RecordTTL:  endpoint.TTL(defaultTTL),
			Labels:     endpoint.Labels{},
		}}
	}

	create, err := p.AdjustEndpoints(desired())
	require.NoError(t, err)
	require.NoError(t, p.ApplyChanges(ctx, &plan.Changes{Create: create}))

	records, err := p.Records(ctx)
	require.NoError(t, err)

	adjusted, err := p.AdjustEndpoints(desired())
	require.NoError(t, err)

	pl := &plan.Plan{
		Current:        records,
		Desired:        adjusted,
		DomainFilter:   endpoint.MatchAllDomainFilters{endpoint.NewDomainFilter([]string{"bar.com"})},
		ManagedRecords: []string{endpoint.RecordTypeNAPTR},
	}
	return pl.Calculate()
}

func TestCloudflareNAPTRConverge(t *testing.T) {
	cases := []struct {
		name   string
		target string
	}{
		{"canonical", `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com.`},
		{"multi-space", `10  20  "S"  "SIP+D2U"  ""  _sip._udp.bar.com.`},
		{"tab", "10\t20\t\"S\"\t\"SIP+D2U\"\t\"\"\t_sip._udp.bar.com."},
		{"no trailing dot", `10 20 "S" "SIP+D2U" "" _sip._udp.bar.com`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			planned := replanAfterApply(t, endpoint.RecordTypeNAPTR, "bar.com", c.target)
			assert.Empty(t, planned.Changes.Create, "should converge: unexpected create on second reconcile")
			assert.Empty(t, planned.Changes.UpdateNew, "should converge: unexpected update on second reconcile")
			assert.Empty(t, planned.Changes.Delete, "should converge: unexpected delete on second reconcile")
		})
	}
}

func TestNAPTRIsReadableAndNotProxied(t *testing.T) {
	t.Parallel()

	p := &CloudFlareProvider{}
	assert.True(t, p.SupportedAdditionalRecordTypes(endpoint.RecordTypeNAPTR),
		"NAPTR must be readable or the planner never sees existing records")

	// Cloudflare rejects proxying for NAPTR, so proxying must stay off even with --cloudflare-proxied.
	assert.True(t, recordTypeProxyNotSupported.Has(endpoint.RecordTypeNAPTR))
	ep := endpoint.NewEndpoint(naptrName, endpoint.RecordTypeNAPTR, naptrTarget)
	assert.False(t, shouldBeProxied(ep, true))
}
