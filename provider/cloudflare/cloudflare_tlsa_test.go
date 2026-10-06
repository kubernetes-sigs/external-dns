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
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
)

const (
	tlsaDigest = "54ee30e752c6e751b7bfde666e638ec4e25f3974207258072299caa15e15820f"
	tlsaTarget = "3 1 1 " + tlsaDigest
	tlsaName   = "_443._tcp.example.com"
)

func tlsaRecordResponse() dns.RecordResponse {
	return dns.RecordResponse{
		Name:    tlsaName,
		TTL:     300,
		Type:    dns.RecordResponseTypeTLSA,
		Content: tlsaTarget,
	}
}

// The API rejects TLSA sent as content, so writes must send structured data.
func assertTLSADataOnly(t *testing.T, body any) {
	t.Helper()

	raw, err := json.Marshal(body)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(raw, &decoded))

	assert.NotContains(t, decoded, "content", "TLSA writes must not send a content field")

	data, ok := decoded["data"].(map[string]any)
	require.True(t, ok, "TLSA writes must send a data object, got: %s", raw)
	assert.InDelta(t, 3, data["usage"], 0)
	assert.InDelta(t, 1, data["selector"], 0)
	assert.InDelta(t, 1, data["matching_type"], 0)
	assert.Equal(t, tlsaDigest, data["certificate"])
	assert.Equal(t, "TLSA", decoded["type"])
	assert.Equal(t, tlsaName, decoded["name"])
}

func TestGetCreateDNSRecordParamTLSAUsesData(t *testing.T) {
	t.Parallel()

	params, err := getCreateDNSRecordParam("zone-1", &cloudFlareChange{
		Action:         cloudFlareCreate,
		ResourceRecord: tlsaRecordResponse(),
	})
	require.NoError(t, err)
	assert.Equal(t, "zone-1", params.ZoneID.Value)

	body, ok := params.Body.(dns.TLSARecordParam)
	require.True(t, ok, "expected a TLSARecordParam body, got %T", params.Body)
	assertTLSADataOnly(t, body)
}

func TestGetUpdateDNSRecordParamTLSAUsesData(t *testing.T) {
	t.Parallel()

	params, err := getUpdateDNSRecordParam("zone-1", cloudFlareChange{
		Action:         cloudFlareUpdate,
		ResourceRecord: tlsaRecordResponse(),
	})
	require.NoError(t, err)
	assert.Equal(t, "zone-1", params.ZoneID.Value)

	body, ok := params.Body.(dns.TLSARecordParam)
	require.True(t, ok, "expected a TLSARecordParam body, got %T", params.Body)
	assertTLSADataOnly(t, body)
}

func TestBuildBatchPostParamTLSAUsesData(t *testing.T) {
	t.Parallel()

	param, ok := buildBatchPostParam(tlsaRecordResponse())
	require.True(t, ok)

	body, isTLSA := param.(dns.TLSARecordParam)
	require.True(t, isTLSA, "expected a TLSARecordParam, got %T", param)
	assertTLSADataOnly(t, body)
}

func TestBuildBatchPutParamTLSAUsesData(t *testing.T) {
	t.Parallel()

	param, ok := buildBatchPutParam("record-id", tlsaRecordResponse())
	require.True(t, ok)

	put, isTLSA := param.(dns.BatchPutTLSARecordParam)
	require.True(t, isTLSA, "expected a BatchPutTLSARecordParam, got %T", param)
	assert.Equal(t, "record-id", put.ID.Value)
	assertTLSADataOnly(t, put.TLSARecordParam)
}

// Other record types must keep using content, unchanged.
func TestNonTLSARecordsStillUseContent(t *testing.T) {
	t.Parallel()

	params, err := getCreateDNSRecordParam("zone-1", &cloudFlareChange{
		Action: cloudFlareCreate,
		ResourceRecord: dns.RecordResponse{
			Name:    "example.com",
			TTL:     120,
			Type:    dns.RecordResponseTypeA,
			Content: "1.2.3.4",
		},
	})
	require.NoError(t, err)

	body, ok := params.Body.(dns.RecordNewParamsBody)
	require.True(t, ok, "expected the generic body for an A record, got %T", params.Body)
	assert.Equal(t, "1.2.3.4", body.Content.Value)
}

func TestMalformedTLSATargetIsRejected(t *testing.T) {
	t.Parallel()

	bad := dns.RecordResponse{
		Name:    tlsaName,
		TTL:     300,
		Type:    dns.RecordResponseTypeTLSA,
		Content: "not a tlsa record",
	}

	_, err := getCreateDNSRecordParam("zone-1", &cloudFlareChange{ResourceRecord: bad})
	require.Error(t, err)

	_, err = getUpdateDNSRecordParam("zone-1", cloudFlareChange{ResourceRecord: bad})
	require.Error(t, err)

	_, ok := buildBatchPostParam(bad)
	assert.False(t, ok, "a malformed target must not produce a batch post")

	_, ok = buildBatchPutParam("record-id", bad)
	assert.False(t, ok, "a malformed target must not produce a batch put")
}

// Without normalisation a record never compares equal to itself and every sync updates.
func TestEndpointTargetFromCloudflareRecordNormalisesTLSA(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		content string
		want    string
	}{
		{"canonical", tlsaTarget, tlsaTarget},
		{"uppercase digest", "3 1 1 " + strings.ToUpper(tlsaDigest), tlsaTarget},
		{"extra spacing", "3  1  1  " + tlsaDigest, tlsaTarget},
		{"unparseable is passed through", "garbage", "garbage"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := endpointTargetFromCloudflareRecord(dns.RecordResponse{
				Name:    tlsaName,
				Type:    dns.RecordResponseTypeTLSA,
				Content: tc.content,
			})
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEndpointTargetFromCloudflareRecordLeavesOtherTypesAlone(t *testing.T) {
	t.Parallel()

	got := endpointTargetFromCloudflareRecord(dns.RecordResponse{
		Name:    "example.com",
		Type:    dns.RecordResponseTypeTXT,
		Content: "Some Mixed Case Value",
	})
	assert.Equal(t, "Some Mixed Case Value", got)
}

func TestTLSAIsReadableAndNotProxied(t *testing.T) {
	t.Parallel()

	p := &CloudFlareProvider{}
	assert.True(t, p.SupportedAdditionalRecordTypes(endpoint.RecordTypeTLSA),
		"TLSA must be readable or the planner never sees existing records")

	// Cloudflare only proxies A/AAAA/CNAME; asking it to proxy a TLSA record is
	// rejected, so proxying must be forced off even when --cloudflare-proxied is set.
	assert.True(t, recordTypeProxyNotSupported.Has(endpoint.RecordTypeTLSA))
	ep := endpoint.NewEndpoint(tlsaName, endpoint.RecordTypeTLSA, tlsaTarget)
	assert.False(t, shouldBeProxied(ep, true))
}
