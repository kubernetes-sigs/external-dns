/*
Copyright 2020 The Kubernetes Authors.

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

package vultr

import (
	"context"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vultr/govultr/v2"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockVultrDomain struct {
	client *govultr.Client
}

func (m mockVultrDomain) Create(ctx context.Context, domainReq *govultr.DomainReq) (*govultr.Domain, error) {
	return nil, nil
}

func (m mockVultrDomain) Get(ctx context.Context, domain string) (*govultr.Domain, error) {
	return nil, nil
}

func (m mockVultrDomain) Update(ctx context.Context, domain, dnsSec string) error {
	return nil
}

func (m mockVultrDomain) Delete(ctx context.Context, domain string) error {
	return nil
}

func (m mockVultrDomain) List(ctx context.Context, options *govultr.ListOptions) ([]govultr.Domain, *govultr.Meta, error) {
	return []govultr.Domain{{Domain: "test.com", DateCreated: "1234"}}, &govultr.Meta{
		Total: 1,
		Links: &govultr.Links{
			Next: "",
			Prev: "",
		},
	}, nil
}

func (m mockVultrDomain) GetSoa(ctx context.Context, domain string) (*govultr.Soa, error) {
	return nil, nil
}

func (m mockVultrDomain) UpdateSoa(ctx context.Context, domain string, soaReq *govultr.Soa) error {
	return nil
}

func (m mockVultrDomain) GetDNSSec(ctx context.Context, domain string) ([]string, error) {
	return nil, nil
}

type mockVultrRecord struct {
	client *govultr.Client
}

func (m mockVultrRecord) Create(ctx context.Context, domain string, domainRecordReq *govultr.DomainRecordReq) (*govultr.DomainRecord, error) {
	return nil, nil
}

func (m mockVultrRecord) Get(ctx context.Context, domain, recordID string) (*govultr.DomainRecord, error) {
	return nil, nil
}

func (m mockVultrRecord) Update(ctx context.Context, domain, recordID string, domainRecordReq *govultr.DomainRecordReq) error {
	return nil
}

func (m mockVultrRecord) Delete(ctx context.Context, domain, recordID string) error {
	return nil
}

func (m mockVultrRecord) List(ctx context.Context, domain string, options *govultr.ListOptions) ([]govultr.DomainRecord, *govultr.Meta, error) {
	return []govultr.DomainRecord{{ID: "123", Type: "A", Name: "test", Data: "192.168.1.1", TTL: 300}}, &govultr.Meta{
		Total: 1,
		Links: &govultr.Links{
			Next: "",
			Prev: "",
		},
	}, nil
}

func TestNewVultrProvider(t *testing.T) {
	_ = os.Setenv("VULTR_API_KEY", "")
	_, err := NewVultrProvider(context.Background(), endpoint.NewDomainFilter([]string{"test.vultr.com"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}

	_ = os.Unsetenv("VULTR_API_KEY")
	_, err = NewVultrProvider(context.Background(), endpoint.NewDomainFilter([]string{"test.vultr.com"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestVultrProvider_Zones(t *testing.T) {
	mocked := mockVultrDomain{nil}
	provider := &VultrProvider{
		client: govultr.Client{
			Domain: &mocked,
		},
	}

	expected, _, err := provider.client.Domain.List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	provider.Zones(context.Background())
	zones, err := provider.Zones(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, zones) {
		t.Fatal(err)
	}
}

func TestVultrProvider_Records(t *testing.T) {
	mocked := mockVultrRecord{nil}
	mockedDomain := mockVultrDomain{nil}

	provider := &VultrProvider{
		client: govultr.Client{
			DomainRecord: &mocked,
			Domain:       &mockedDomain,
		},
	}

	expected, _, _ := provider.client.DomainRecord.List(context.Background(), "test.com", nil)
	records, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range records {
		assert.Equal(t, strings.TrimSuffix(v.DNSName, ".test.com"), expected[0].Name)
		assert.Equal(t, v.RecordType, expected[0].Type)
		assert.Equal(t, int(v.RecordTTL), expected[0].TTL)
	}
}

func TestVultrProvider_ApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	mocked := mockVultrRecord{nil}
	mockedDomain := mockVultrDomain{nil}

	provider := &VultrProvider{
		client: govultr.Client{
			DomainRecord: &mocked,
			Domain:       &mockedDomain,
		},
	}

	changes.Create = []*endpoint.Endpoint{
		{DNSName: "test.com", Targets: endpoint.Targets{"target"}},
		{DNSName: "ttl.test.com", Targets: endpoint.Targets{"target"}, RecordTTL: 100},
	}

	changes.UpdateNew = []*endpoint.Endpoint{{DNSName: "test.test.com", Targets: endpoint.Targets{"target-new"}, RecordType: "A", RecordTTL: 100}}
	changes.Delete = []*endpoint.Endpoint{{DNSName: "test.test.com", Targets: endpoint.Targets{"target"}, RecordType: "A"}}
	err := provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestVultrProvider_getRecordID(t *testing.T) {
	mocked := mockVultrRecord{nil}
	mockedDomain := mockVultrDomain{nil}

	provider := &VultrProvider{
		client: govultr.Client{
			DomainRecord: &mocked,
			Domain:       &mockedDomain,
		},
	}

	record := &govultr.DomainRecordReq{
		Type: "A",
		Name: "test.test.com",
	}
	id, err := provider.getRecordID(context.Background(), "test.com", record)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, id, "123")
}
