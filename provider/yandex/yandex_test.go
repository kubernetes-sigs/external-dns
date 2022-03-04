/*
Copyright 2022 The Kubernetes Authors.

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

package yandex

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	dnsProto "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	opProto "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"google.golang.org/grpc"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/provider"
)

type mockDNSClient struct {
	zoneIterator       *mockZoneIterator
	recordSetIterators map[string]*mockRecordSetIterator
	dryRun             bool
	domainFilter       endpoint.DomainFilter
	zoneNameFilter     endpoint.DomainFilter
	zoneIDFilter       provider.ZoneIDFilter
}

type mockIterator struct {
	error  error
	index  int
	values []interface{}
}

type mockZoneIterator struct {
	*mockIterator
}

type mockRecordSetIterator struct {
	*mockIterator
}

func (i *mockIterator) SetValues(values ...interface{}) {
	for _, value := range values {
		i.values = append(i.values, value)
	}
	i.index = -1
}

func (i *mockIterator) SetError(err error) {
	i.error = err
	i.values = nil
	i.index = -1
}

func (i *mockIterator) Error() error {
	return i.error
}

func (i *mockIterator) Next() bool {
	i.index++
	return i.index < len(i.values)
}

func (i *mockZoneIterator) Value() *dnsProto.DnsZone {
	return i.values[i.index].(*dnsProto.DnsZone)
}

func (i *mockRecordSetIterator) Value() *dnsProto.RecordSet {
	return i.values[i.index].(*dnsProto.RecordSet)
}

func newMockDNSClient() *mockDNSClient {
	return &mockDNSClient{
		zoneIterator:       &mockZoneIterator{&mockIterator{values: make([]interface{}, 0)}},
		recordSetIterators: map[string]*mockRecordSetIterator{},
	}
}

func (c *mockDNSClient) ZoneIterator(_ context.Context, _ *dnsProto.ListDnsZonesRequest, _ ...grpc.CallOption) ZoneIteratorAdapter {
	return c.zoneIterator
}

func (c *mockDNSClient) RecordSetIterator(_ context.Context, req *dnsProto.ListDnsZoneRecordSetsRequest, _ ...grpc.CallOption) RecordSetIteratorAdapter {
	return c.recordSetIterators[req.DnsZoneId]
}

func (c *mockDNSClient) UpsertRecordSets(_ context.Context, _ *dnsProto.UpsertRecordSetsRequest, _ ...grpc.CallOption) (*opProto.Operation, error) {
	return nil, nil
}

func (c *mockDNSClient) WithZoneRecords(zone string, sets ...*dnsProto.RecordSet) *mockDNSClient {
	key := strconv.Itoa(len(c.zoneIterator.values) + 1)

	c.zoneIterator.SetValues(&dnsProto.DnsZone{
		Id:   key,
		Zone: zone,
	})

	for _, rs := range sets {
		it, ok := c.recordSetIterators[key]
		if !ok {
			it = &mockRecordSetIterator{&mockIterator{values: make([]interface{}, 0)}}
			c.recordSetIterators[key] = it
		}
		it.SetValues(rs)
	}

	return c
}

func (c *mockDNSClient) WithDryRun() *mockDNSClient {
	c.dryRun = true
	return c
}

func (c *mockDNSClient) WithDomainFilter(domains ...string) *mockDNSClient {
	c.domainFilter = endpoint.NewDomainFilter(domains)
	return c
}

func (c *mockDNSClient) WithZoneNameFilter(f endpoint.DomainFilter) *mockDNSClient {
	c.zoneNameFilter = f
	return c
}

func (c *mockDNSClient) WithZoneIDFilter(f provider.ZoneIDFilter) *mockDNSClient {
	c.zoneIDFilter = f
	return c
}

func (c *mockDNSClient) BuildProvider() *YandexProvider {
	return &YandexProvider{
		DomainFilter:   c.domainFilter,
		ZoneNameFilter: c.zoneNameFilter,
		ZoneIDFilter:   c.zoneIDFilter,
		DryRun:         c.dryRun,
		FolderID:       "folderId",
		client:         c,
	}
}

func assertEndpointsAreSame(t *testing.T, expected, actual []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(expected, actual), "expected and actual endpoints don't match. %s:%s", actual, expected)
}

func TestYandexRecord(t *testing.T) {
	p := newMockDNSClient().
		WithZoneRecords("yandex.io.",
			&dnsProto.RecordSet{Name: "test.yandex.io.", Type: "A", Ttl: 10, Data: []string{"1.2.3.4"}},
			&dnsProto.RecordSet{Name: "another.yandex.io.", Type: "CNAME", Ttl: 10, Data: []string{"test2.yandex.io"}},
		).
		BuildProvider()

	actual, err := p.Records(context.Background())

	if err != nil {
		t.Fatal(err)
	}
	assertEndpointsAreSame(t, []*endpoint.Endpoint{
		{DNSName: "test.yandex.io", Targets: []string{"1.2.3.4"}, RecordType: "A", RecordTTL: 10},
		{DNSName: "another.yandex.io", Targets: []string{"test2.yandex.io"}, RecordType: "CNAME", RecordTTL: 10},
	}, actual)
}

func TestYandexRecordWithDomainFilter(t *testing.T) {
	p := newMockDNSClient().
		WithZoneRecords("another.io.",
			&dnsProto.RecordSet{Name: "internal.another.io.", Type: "A", Ttl: 10, Data: []string{"1.2.3.4"}},
			&dnsProto.RecordSet{Name: "stub.another.io.", Type: "CNAME", Ttl: 10, Data: []string{"test2.yandex.io"}},
		).
		WithZoneRecords("yandex.io.",
			&dnsProto.RecordSet{Name: "test.yandex.io.", Type: "A", Ttl: 10, Data: []string{"1.2.3.4"}},
			&dnsProto.RecordSet{Name: "another.yandex.io.", Type: "CNAME", Ttl: 10, Data: []string{"test2.yandex.io"}},
		).
		WithDomainFilter("yandex.io").
		BuildProvider()

	actual, err := p.Records(context.Background())

	if err != nil {
		t.Fatal(err)
	}
	assertEndpointsAreSame(t, []*endpoint.Endpoint{
		{DNSName: "test.yandex.io", Targets: []string{"1.2.3.4"}, RecordType: "A", RecordTTL: 10},
		{DNSName: "another.yandex.io", Targets: []string{"test2.yandex.io"}, RecordType: "CNAME", RecordTTL: 10},
	}, actual)
}
