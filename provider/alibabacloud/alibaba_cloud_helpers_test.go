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

package alibabacloud

import (
	"errors"
	"slices"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pvtz"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

// fakeDescribeDomainsClient overrides DescribeDomains only. Any other method
// call panics through the nil embedded interface, which makes unexpected
// calls visible in tests.
type fakeDescribeDomainsClient struct {
	AlibabaCloudDNSAPI
	pages      map[string][]string // page number -> domain names
	totalCount int64
	err        error
	calls      int
}

func (f *fakeDescribeDomainsClient) DescribeDomains(request *alidns.DescribeDomainsRequest) (*alidns.DescribeDomainsResponse, error) {
	f.calls++
	if f.err != nil {
		return nil, f.err
	}
	page := string(request.PageNumber)
	response := alidns.CreateDescribeDomainsResponse()
	for _, name := range f.pages[page] {
		response.Domains.Domain = append(response.Domains.Domain, alidns.DomainInDescribeDomains{DomainName: name})
	}
	pageNumber, err := request.PageNumber.GetValue64()
	if err != nil {
		return nil, err
	}
	response.PageNumber = pageNumber
	response.TotalCount = f.totalCount
	return response, nil
}

// fakeZoneInfoClient overrides DescribeZoneInfo only.
type fakeZoneInfoClient struct {
	AlibabaCloudPrivateZoneAPI
	vpcIDs []string
	err    error
}

func (f *fakeZoneInfoClient) DescribeZoneInfo(_ *pvtz.DescribeZoneInfoRequest) (*pvtz.DescribeZoneInfoResponse, error) {
	if f.err != nil {
		return nil, f.err
	}
	response := pvtz.CreateDescribeZoneInfoResponse()
	for _, id := range f.vpcIDs {
		response.BindVpcs.Vpc = append(response.BindVpcs.Vpc, pvtz.VpcInDescribeZoneInfo{VpcId: id})
	}
	return response, nil
}

func TestGetNextPageNumber(t *testing.T) {
	tests := []struct {
		name       string
		pageNumber int64
		totalCount int64
		want       int64
	}{
		{name: "no results", pageNumber: 1, totalCount: 0, want: 0},
		{name: "exactly one full page", pageNumber: 1, totalCount: defaultAlibabaCloudPageSize, want: 0},
		{name: "one more than a page", pageNumber: 1, totalCount: defaultAlibabaCloudPageSize + 1, want: 2},
		{name: "last of several pages", pageNumber: 3, totalCount: 3 * defaultAlibabaCloudPageSize, want: 0},
		{name: "middle page", pageNumber: 2, totalCount: 3 * defaultAlibabaCloudPageSize, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getNextPageNumber(tt.pageNumber, tt.totalCount))
		})
	}
}

func TestWrapAndUnwrapQuotes(t *testing.T) {
	tests := []struct {
		name       string
		recordType string
		raw        string // value as stored by Alibaba Cloud
		wrapped    string // value as expected by ExternalDNS
	}{
		{
			name:       "TXT ownership record is quoted",
			recordType: endpoint.RecordTypeTXT,
			raw:        "heritage=external-dns,external-dns/owner=default",
			wrapped:    `"heritage=external-dns,external-dns/owner=default"`,
		},
		{
			name:       "TXT record without heritage prefix is unchanged",
			recordType: endpoint.RecordTypeTXT,
			raw:        "v=spf1 -all",
			wrapped:    "v=spf1 -all",
		},
		{
			name:       "A record is unchanged",
			recordType: endpoint.RecordTypeA,
			raw:        "1.2.3.4",
			wrapped:    "1.2.3.4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wrapped, wrapWithQuotes(tt.recordType, tt.raw))
			assert.Equal(t, tt.raw, unwrapQuotes(tt.recordType, tt.wrapped))
		})
	}
}

func TestGetDNSName(t *testing.T) {
	p := &AlibabaCloudProvider{}
	assert.Equal(t, "example.com", p.getDNSName(nullHostAlibabaCloud, "example.com"))
	assert.Equal(t, "www.example.com", p.getDNSName("www", "example.com"))
}

func TestGetRecordKey(t *testing.T) {
	p := &AlibabaCloudProvider{}

	apex := alidns.Record{RR: nullHostAlibabaCloud, DomainName: "example.com", Type: endpoint.RecordTypeA}
	sub := alidns.Record{RR: "www", DomainName: "example.com", Type: endpoint.RecordTypeA}

	assert.Equal(t, "A:example.com", p.getRecordKey(apex))
	assert.Equal(t, "A:www.example.com", p.getRecordKey(sub))

	// The key built from an endpoint must match the key built from the record
	// for the same name and type, otherwise updates and deletes are skipped.
	assert.Equal(t, p.getRecordKey(sub), p.getRecordKeyByEndpoint(endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "1.2.3.4")))
	assert.Equal(t, p.getRecordKey(apex), p.getRecordKeyByEndpoint(endpoint.NewEndpoint("example.com", endpoint.RecordTypeA, "1.2.3.4")))
}

func TestGroupRecords(t *testing.T) {
	p := &AlibabaCloudProvider{}
	records := []alidns.Record{
		{RR: "www", DomainName: "example.com", Type: endpoint.RecordTypeA, Value: "1.1.1.1"},
		{RR: "www", DomainName: "example.com", Type: endpoint.RecordTypeA, Value: "2.2.2.2"},
		{RR: "www", DomainName: "example.com", Type: endpoint.RecordTypeTXT, Value: "text"},
	}

	groups := p.groupRecords(records)

	require.Len(t, groups, 2)
	assert.Len(t, groups["A:www.example.com"], 2)
	assert.Len(t, groups["TXT:www.example.com"], 1)
}

func TestEquals(t *testing.T) {
	p := &AlibabaCloudProvider{}
	tests := []struct {
		name      string
		recordTTL int64
		epTTL     endpoint.TTL
		want      bool
	}{
		{name: "default TTL on record matches unset TTL", recordTTL: defaultTTL, epTTL: 0, want: true},
		{name: "unset TTL on record matches default TTL", recordTTL: 0, epTTL: defaultTTL, want: true},
		{name: "same custom TTL", recordTTL: 300, epTTL: 300, want: true},
		{name: "different TTL", recordTTL: 300, epTTL: defaultTTL, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			record := alidns.Record{TTL: tt.recordTTL}
			ep := endpoint.NewEndpointWithTTL("www.example.com", endpoint.RecordTypeA, tt.epTTL, "1.2.3.4")
			assert.Equal(t, tt.want, p.equals(record, ep))
		})
	}
}

func TestEqualsPrivateZone(t *testing.T) {
	p := &AlibabaCloudProvider{}
	tests := []struct {
		name      string
		recordTTL int
		epTTL     endpoint.TTL
		want      bool
	}{
		{name: "default TTL on record matches unset TTL", recordTTL: defaultAlibabaCloudPrivateZoneRecordTTL, epTTL: 0, want: true},
		{name: "same custom TTL", recordTTL: 300, epTTL: 300, want: true},
		{name: "different TTL", recordTTL: 300, epTTL: defaultAlibabaCloudPrivateZoneRecordTTL, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			record := pvtz.Record{Ttl: tt.recordTTL}
			ep := endpoint.NewEndpointWithTTL("www.example.com", endpoint.RecordTypeA, tt.epTTL, "1.2.3.4")
			assert.Equal(t, tt.want, p.equalsPrivateZone(record, ep))
		})
	}
}

func TestKeys(t *testing.T) {
	got := keys(map[string]int{"b.com": 2, "a.com": 1})
	slices.Sort(got)
	assert.Equal(t, []string{"a.com", "b.com"}, got)
	assert.Empty(t, keys(map[string]int{}))
}

func TestMatchVPC(t *testing.T) {
	tests := []struct {
		name   string
		client *fakeZoneInfoClient
		want   bool
	}{
		{name: "zone bound to provider VPC", client: &fakeZoneInfoClient{vpcIDs: []string{"vpc-other", "vpc-mine"}}, want: true},
		{name: "zone bound to other VPCs only", client: &fakeZoneInfoClient{vpcIDs: []string{"vpc-other"}}, want: false},
		{name: "zone not bound to any VPC", client: &fakeZoneInfoClient{}, want: false},
		{name: "API error", client: &fakeZoneInfoClient{err: errors.New("describe zone info failed")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AlibabaCloudProvider{vpcID: "vpc-mine", pvtzClient: tt.client}
			assert.Equal(t, tt.want, p.matchVPC("zone-1"))
		})
	}
}

func TestGetDomainListPagination(t *testing.T) {
	client := &fakeDescribeDomainsClient{
		pages: map[string][]string{
			"1": {"a.com"},
			"2": {"b.com"},
		},
		totalCount: defaultAlibabaCloudPageSize + 1,
	}
	p := &AlibabaCloudProvider{dnsClient: client}

	domains, err := p.getDomainList()

	require.NoError(t, err)
	assert.Equal(t, []string{"a.com", "b.com"}, domains)
	assert.Equal(t, 2, client.calls)
}

func TestGetDomainListError(t *testing.T) {
	p := &AlibabaCloudProvider{dnsClient: &fakeDescribeDomainsClient{err: errors.New("describe domains failed")}}

	domains, err := p.getDomainList()

	require.Error(t, err)
	assert.Nil(t, domains)
}

func TestRecordsReturnsErrorWhenDomainListFails(t *testing.T) {
	apiErr := errors.New("describe domains failed")
	p := &AlibabaCloudProvider{dnsClient: &fakeDescribeDomainsClient{err: apiErr}}

	endpoints, err := p.Records(t.Context())

	require.ErrorIs(t, err, apiErr)
	assert.Nil(t, endpoints)
}

func TestApplyChangesForDNSReturnsErrorWhenDomainListFails(t *testing.T) {
	apiErr := errors.New("describe domains failed")
	p := &AlibabaCloudProvider{dnsClient: &fakeDescribeDomainsClient{err: apiErr}}
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "1.2.3.4")},
	}

	err := p.ApplyChanges(t.Context(), changes)

	require.ErrorIs(t, err, apiErr)
}

func TestApplyChangesNoOp(t *testing.T) {
	// No clients are configured: any API call would panic, so a nil error
	// proves that empty change sets return early.
	p := &AlibabaCloudProvider{}

	require.NoError(t, p.ApplyChanges(t.Context(), nil))
	require.NoError(t, p.ApplyChanges(t.Context(), &plan.Changes{}))
}

func TestCreateRecordErrors(t *testing.T) {
	// No DNS client is configured: both cases must fail before any API call.
	p := &AlibabaCloudProvider{}
	ep := endpoint.NewEndpoint("www.example.com", endpoint.RecordTypeA, "1.2.3.4")

	t.Run("no hosted zones", func(t *testing.T) {
		err := p.createRecord(ep, "1.2.3.4", nil)
		require.EqualError(t, err, "zone not found")
	})

	t.Run("no matching hosted zone", func(t *testing.T) {
		err := p.createRecord(ep, "1.2.3.4", []string{"other.org"})
		require.EqualError(t, err, "no corresponding DNS zone found for this domain")
	})
}
