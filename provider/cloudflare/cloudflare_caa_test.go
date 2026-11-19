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
	"sigs.k8s.io/external-dns/provider"
)

const (
	caaName   = "example.com"
	caaTarget = `0 issue "letsencrypt.org; validationmethods=dns-01"`
)

func caaRecordResponse() dns.RecordResponse {
	return dns.RecordResponse{
		Name:    caaName,
		TTL:     300,
		Type:    dns.RecordResponseTypeCAA,
		Content: caaTarget,
	}
}

// The API rejects CAA sent as content ("flags is a required data field"), so writes must send structured data.
func assertCAADataOnly(t *testing.T, body any) {
	t.Helper()

	raw, err := json.Marshal(body)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(raw, &decoded))

	assert.NotContains(t, decoded, "content", "CAA writes must not send a content field")

	data, ok := decoded["data"].(map[string]any)
	require.True(t, ok, "CAA writes must send a data object, got: %s", raw)
	require.Contains(t, data, "flags", "flags 0 must still be sent")
	assert.InDelta(t, 0, data["flags"], 0)
	assert.Equal(t, "issue", data["tag"])
	assert.Equal(t, "letsencrypt.org; validationmethods=dns-01", data["value"])
	assert.Equal(t, "CAA", decoded["type"])
	assert.Equal(t, caaName, decoded["name"])
}

func TestGetCreateDNSRecordParamCAAUsesData(t *testing.T) {
	t.Parallel()

	params, err := getCreateDNSRecordParam("zone-1", &cloudFlareChange{
		Action:         cloudFlareCreate,
		ResourceRecord: caaRecordResponse(),
	})
	require.NoError(t, err)

	body, ok := params.Body.(dns.CAARecordParam)
	require.True(t, ok, "expected a CAARecordParam body, got %T", params.Body)
	assertCAADataOnly(t, body)
}

func TestGetUpdateDNSRecordParamCAAUsesData(t *testing.T) {
	t.Parallel()

	params, err := getUpdateDNSRecordParam("zone-1", cloudFlareChange{
		Action:         cloudFlareUpdate,
		ResourceRecord: caaRecordResponse(),
	})
	require.NoError(t, err)

	body, ok := params.Body.(dns.CAARecordParam)
	require.True(t, ok, "expected a CAARecordParam body, got %T", params.Body)
	assertCAADataOnly(t, body)
}

func TestBuildBatchPostParamCAAUsesData(t *testing.T) {
	t.Parallel()

	param, ok := buildBatchPostParam(caaRecordResponse())
	require.True(t, ok)

	body, isCAA := param.(dns.CAARecordParam)
	require.True(t, isCAA, "expected a CAARecordParam, got %T", param)
	assertCAADataOnly(t, body)
}

func TestBuildBatchPutParamCAAUsesData(t *testing.T) {
	t.Parallel()

	param, ok := buildBatchPutParam("record-id", caaRecordResponse())
	require.True(t, ok)

	put, isCAA := param.(dns.BatchPutCAARecordParam)
	require.True(t, isCAA, "expected a BatchPutCAARecordParam, got %T", param)
	assert.Equal(t, "record-id", put.ID.Value)
	assertCAADataOnly(t, put.CAARecordParam)
}

func TestMalformedCAATargetIsRejected(t *testing.T) {
	t.Parallel()

	bad := dns.RecordResponse{
		Name:    caaName,
		TTL:     300,
		Type:    dns.RecordResponseTypeCAA,
		Content: "0 issue",
	}

	_, err := getCreateDNSRecordParam("zone-1", &cloudFlareChange{ResourceRecord: bad})
	require.ErrorContains(t, err, "invalid CAA target")

	_, err = getUpdateDNSRecordParam("zone-1", cloudFlareChange{ResourceRecord: bad})
	require.Error(t, err)

	_, ok := buildBatchPostParam(bad)
	assert.False(t, ok, "a malformed target must not produce a batch post")

	_, ok = buildBatchPutParam("record-id", bad)
	assert.False(t, ok, "a malformed target must not produce a batch put")
}

func TestMalformedCAAEndpointReturnsSoftError(t *testing.T) {
	client := NewMockCloudFlareClient()
	p := &CloudFlareProvider{Client: client}

	err := p.ApplyChanges(t.Context(), &plan.Changes{Create: []*endpoint.Endpoint{{
		DNSName:    "bar.com",
		RecordType: endpoint.RecordTypeCAA,
		Targets:    endpoint.Targets{"0 issue"},
	}}})
	require.ErrorIs(t, err, provider.SoftError)
}

func TestEndpointTargetFromCloudflareRecordNormalisesCAA(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		record dns.RecordResponse
		want   string
	}{
		{
			name: "structured data",
			record: dns.RecordResponse{
				Type:    dns.RecordResponseTypeCAA,
				Content: "0 iodef mailto:security@example.com",
				Data:    dns.CAARecordData{Flags: 128, Tag: "iodef", Value: "mailto:security@example.com"},
			},
			want: `128 iodef "mailto:security@example.com"`,
		},
		{
			name:   "content only, unquoted",
			record: dns.RecordResponse{Type: dns.RecordResponseTypeCAA, Content: "0 issue letsencrypt.org"},
			want:   `0 issue "letsencrypt.org"`,
		},
		{
			name:   "unparseable content is kept verbatim",
			record: dns.RecordResponse{Type: dns.RecordResponseTypeCAA, Content: "garbage"},
			want:   "garbage",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, endpointTargetFromCloudflareRecord(tt.record))
		})
	}
}

func TestAdjustEndpointsCanonicalizesCAA(t *testing.T) {
	t.Parallel()

	p := &CloudFlareProvider{}
	adjusted, err := p.AdjustEndpoints([]*endpoint.Endpoint{{
		DNSName:    "bar.com",
		RecordType: endpoint.RecordTypeCAA,
		Targets:    endpoint.Targets{"0  issue  letsencrypt.org", "0 issue"},
	}})
	require.NoError(t, err)
	require.Len(t, adjusted, 1)
	assert.Equal(t, `0 issue "letsencrypt.org"`, adjusted[0].Targets[0])
	assert.Equal(t, "0 issue", adjusted[0].Targets[1], "a malformed CAA target must be left unchanged so it surfaces on apply")
}

func TestCloudflareCAARoundTrip(t *testing.T) {
	client := NewMockCloudFlareClient()
	p := &CloudFlareProvider{Client: client}
	ctx := t.Context()

	findCAA := func() *endpoint.Endpoint {
		records, err := p.Records(ctx)
		require.NoError(t, err)
		for _, r := range records {
			if r.RecordType == endpoint.RecordTypeCAA && r.DNSName == "bar.com" {
				return r
			}
		}
		return nil
	}

	ep := &endpoint.Endpoint{
		DNSName:    "bar.com",
		RecordType: endpoint.RecordTypeCAA,
		Targets:    endpoint.Targets{`0 issue "letsencrypt.org"`},
		RecordTTL:  endpoint.TTL(defaultTTL),
		Labels:     endpoint.Labels{},
	}
	require.NoError(t, p.ApplyChanges(ctx, &plan.Changes{Create: []*endpoint.Endpoint{ep}}))

	got := findCAA()
	require.NotNil(t, got, "created CAA record must be read back")
	assert.Equal(t, endpoint.Targets{`0 issue "letsencrypt.org"`}, got.Targets)

	updated := got.DeepCopy()
	updated.RecordTTL = 600
	require.NoError(t, p.ApplyChanges(ctx, &plan.Changes{UpdateOld: []*endpoint.Endpoint{got}, UpdateNew: []*endpoint.Endpoint{updated}}))
	got = findCAA()
	require.NotNil(t, got)
	assert.Equal(t, endpoint.TTL(600), got.RecordTTL)

	require.NoError(t, p.ApplyChanges(ctx, &plan.Changes{Delete: []*endpoint.Endpoint{got}}))
	assert.Nil(t, findCAA(), "deleted CAA record must not be read back")
}

func TestCloudflareCAAConverge(t *testing.T) {
	cases := []struct {
		name   string
		target string
	}{
		{"quoted", `0 issue "letsencrypt.org"`},
		{"unquoted", "0 issue letsencrypt.org"},
		{"multi-space", `0  iodef  "mailto:security@bar.com"`},
		{"tab", "0\tissue\t\"letsencrypt.org\""},
		{"parameters", `0 issue "letsencrypt.org; validationmethods=dns-01"`},
		{"critical flag", `128 issuewild ";"`},
		{"trailing dot", "0 issue letsencrypt.org."},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			planned := replanAfterApply(t, endpoint.RecordTypeCAA, "bar.com", c.target)
			assert.Empty(t, planned.Changes.Create, "should converge: unexpected create on second reconcile")
			assert.Empty(t, planned.Changes.UpdateNew, "should converge: unexpected update on second reconcile")
			assert.Empty(t, planned.Changes.Delete, "should converge: unexpected delete on second reconcile")
		})
	}
}

func TestCAAIsReadableAndNotProxied(t *testing.T) {
	t.Parallel()

	p := &CloudFlareProvider{}
	assert.True(t, p.SupportedAdditionalRecordTypes(endpoint.RecordTypeCAA),
		"CAA must be readable or the planner never sees existing records")

	assert.True(t, recordTypeProxyNotSupported.Has(endpoint.RecordTypeCAA))
	ep := endpoint.NewEndpoint(caaName, endpoint.RecordTypeCAA, caaTarget)
	assert.False(t, shouldBeProxied(ep, true))
}
