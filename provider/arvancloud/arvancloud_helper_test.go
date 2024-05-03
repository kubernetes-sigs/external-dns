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

package arvancloud

import (
	"context"
	"sigs.k8s.io/external-dns/provider/arvancloud/dto"
	"testing"
)

type mockClient struct {
	t                  *testing.T
	getDomainsOut      func() ([]dto.Zone, error)
	getDnsRecordsOut   func() ([]dto.DnsRecord, error)
	createDnsRecordOut func() (dto.DnsRecord, error)
	updateDnsRecordOut func() (dto.DnsRecord, error)
	deleteDnsRecordOut func() error
}

var _ ArvanAdapter = (*mockClient)(nil)

func (m *mockClient) GetDomains(ctx context.Context, perPage ...int) ([]dto.Zone, error) {
	m.t.Helper()

	return m.getDomainsOut()
}

func (m *mockClient) GetDnsRecords(ctx context.Context, zone string, perPage ...int) ([]dto.DnsRecord, error) {
	m.t.Helper()

	return m.getDnsRecordsOut()
}

func (m *mockClient) CreateDnsRecord(ctx context.Context, zone string, record dto.DnsRecord) (dto.DnsRecord, error) {
	m.t.Helper()

	return m.createDnsRecordOut()
}

func (m *mockClient) UpdateDnsRecord(ctx context.Context, zone string, record dto.DnsRecord) (dto.DnsRecord, error) {
	m.t.Helper()

	return m.updateDnsRecordOut()
}

func (m *mockClient) DeleteDnsRecord(ctx context.Context, zone, recordId string) error {
	m.t.Helper()

	return m.deleteDnsRecordOut()
}
