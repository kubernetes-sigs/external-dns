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

	dnsProto "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	opProto "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"google.golang.org/grpc"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider"
)

type fixture struct {
	dryRun         bool
	domainFilter   endpoint.DomainFilter
	zoneNameFilter endpoint.DomainFilter
	zoneIDFilter   provider.ZoneIDFilter
	provider       *YandexProvider
	client         *mockDNSClient
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

type mockDNSClient struct {
	zoneIterator       *mockZoneIterator
	recordSetIterators map[string]*mockRecordSetIterator
	upsertRequest      *upsertRecordSetsRequests
}

type upsertRecordSetsRequests struct {
	Deletions    []*endpoint.Endpoint
	Merges       []*endpoint.Endpoint
	Replacements []*endpoint.Endpoint
}

func (c *mockDNSClient) ZoneIterator(_ context.Context, _ *dnsProto.ListDnsZonesRequest, _ ...grpc.CallOption) ZoneIteratorAdapter {
	return c.zoneIterator
}

func (c *mockDNSClient) RecordSetIterator(_ context.Context, req *dnsProto.ListDnsZoneRecordSetsRequest, _ ...grpc.CallOption) RecordSetIteratorAdapter {
	return c.recordSetIterators[req.DnsZoneId]
}

func (c *mockDNSClient) UpsertRecordSets(_ context.Context, r *dnsProto.UpsertRecordSetsRequest, _ ...grpc.CallOption) (*opProto.Operation, error) {
	c.upsertRequest = &upsertRecordSetsRequests{
		Deletions:    []*endpoint.Endpoint{},
		Merges:       []*endpoint.Endpoint{},
		Replacements: []*endpoint.Endpoint{},
	}

	for _, r := range r.Deletions {
		c.upsertRequest.Deletions = append(c.upsertRequest.Deletions, toEndpoint(r))
	}
	for _, r := range r.Merges {
		c.upsertRequest.Merges = append(c.upsertRequest.Merges, toEndpoint(r))
	}
	for _, r := range r.Replacements {
		c.upsertRequest.Replacements = append(c.upsertRequest.Replacements, toEndpoint(r))
	}

	return nil, nil
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

func newFixture() *fixture {
	return &fixture{
		dryRun:         false,
		domainFilter:   endpoint.DomainFilter{},
		zoneNameFilter: endpoint.DomainFilter{},
		zoneIDFilter:   provider.ZoneIDFilter{},
		client: &mockDNSClient{
			zoneIterator:       &mockZoneIterator{&mockIterator{values: make([]interface{}, 0)}},
			recordSetIterators: map[string]*mockRecordSetIterator{},
		},
	}
}

func (f *fixture) WithZoneRecords(zone string, sets ...*dnsProto.RecordSet) *fixture {
	key := strconv.Itoa(len(f.client.zoneIterator.values) + 1)

	f.client.zoneIterator.SetValues(&dnsProto.DnsZone{
		Id:   key,
		Zone: zone,
	})

	for _, rs := range sets {
		it, ok := f.client.recordSetIterators[key]
		if !ok {
			it = &mockRecordSetIterator{&mockIterator{values: make([]interface{}, 0)}}
			f.client.recordSetIterators[key] = it
		}
		it.SetValues(rs)
	}

	return f
}

func (f *fixture) WithDryRun() *fixture {
	f.dryRun = true
	return f
}

func (f *fixture) WithDomainFilter(domains ...string) *fixture {
	f.domainFilter = endpoint.NewDomainFilter(domains)
	return f
}

func (f *fixture) WithZoneNameFilter(filter endpoint.DomainFilter) *fixture {
	f.zoneNameFilter = filter
	return f
}

func (f *fixture) WithZoneIDFilter(filter provider.ZoneIDFilter) *fixture {
	f.zoneIDFilter = filter
	return f
}

func (f *fixture) Client() *mockDNSClient {
	return f.client
}

func (f *fixture) Provider() *YandexProvider {
	if f.provider == nil {
		f.provider = &YandexProvider{
			DomainFilter:   f.domainFilter,
			ZoneNameFilter: f.zoneNameFilter,
			ZoneIDFilter:   f.zoneIDFilter,
			DryRun:         f.dryRun,
			FolderID:       "folderId",
			client:         f.client,
		}
	}

	return f.provider
}
