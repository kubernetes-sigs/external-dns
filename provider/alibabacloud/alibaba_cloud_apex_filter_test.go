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
	"regexp"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
)

// fakeDomainRecordsClient overrides DescribeDomainRecords only and returns a
// fixed, single page of records.
type fakeDomainRecordsClient struct {
	AlibabaCloudDNSAPI
	records []alidns.Record
}

func (f *fakeDomainRecordsClient) DescribeDomainRecords(_ *alidns.DescribeDomainRecordsRequest) (*alidns.DescribeDomainRecordsResponse, error) {
	response := alidns.CreateDescribeDomainRecordsResponse()
	response.DomainRecords.Record = f.records
	response.PageNumber = 1
	response.TotalCount = int64(len(f.records))
	return response, nil
}

// Apex records are returned by Alibaba Cloud with RR "@". The domain filter
// must see the real name ("example.com"), not "@.example.com".
func TestGetDomainRecordsApexDomainFilter(t *testing.T) {
	apex := alidns.Record{RR: nullHostAlibabaCloud, DomainName: "example.com", Type: endpoint.RecordTypeA, Value: "1.1.1.1"}
	sub := alidns.Record{RR: "www", DomainName: "example.com", Type: endpoint.RecordTypeA, Value: "2.2.2.2"}

	tests := []struct {
		name   string
		filter *endpoint.DomainFilter
		want   []string // expected RR values
	}{
		{
			name:   "plain filter keeps apex and subdomain",
			filter: endpoint.NewDomainFilter([]string{"example.com"}),
			want:   []string{nullHostAlibabaCloud, "www"},
		},
		{
			name:   "dot-prefixed filter keeps subdomains only",
			filter: endpoint.NewDomainFilter([]string{".example.com"}),
			want:   []string{"www"},
		},
		{
			name:   "regex filter matching apex only",
			filter: endpoint.NewRegexDomainFilter(regexp.MustCompile(`^example\.com$`), nil),
			want:   []string{nullHostAlibabaCloud},
		},
		{
			name:   "regex exclusion of apex",
			filter: endpoint.NewRegexDomainFilter(nil, regexp.MustCompile(`^example\.com$`)),
			want:   []string{"www"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AlibabaCloudProvider{
				domainFilter: tt.filter,
				dnsClient:    &fakeDomainRecordsClient{records: []alidns.Record{apex, sub}},
			}

			records, err := p.getDomainRecords("example.com")

			require.NoError(t, err)
			var got []string
			for _, r := range records {
				got = append(got, r.RR)
			}
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
