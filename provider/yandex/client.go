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

	dnsProto "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	opProto "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"github.com/yandex-cloud/go-sdk/gen/dns"
	"google.golang.org/grpc"
)

// ZoneIteratorAdapter is an interface of dns.DnsZoneIterator that can be stubbed for testing.
type ZoneIteratorAdapter interface {
	Next() bool
	Error() error
	Value() *dnsProto.DnsZone
}

// RecordSetIteratorAdapter is an interface of dns.DnsZoneRecordSetsIterator that can be stubbed for testing.
type RecordSetIteratorAdapter interface {
	Next() bool
	Error() error
	Value() *dnsProto.RecordSet
}

// DNSClient is an interface of dns.DnsZoneServiceClient that can be stubbed for testing.
type DNSClient interface {
	ZoneIterator(ctx context.Context,
		req *dnsProto.ListDnsZonesRequest,
		opts ...grpc.CallOption,
	) ZoneIteratorAdapter

	RecordSetIterator(ctx context.Context,
		req *dnsProto.ListDnsZoneRecordSetsRequest,
		opts ...grpc.CallOption,
	) RecordSetIteratorAdapter

	UpsertRecordSets(ctx context.Context,
		in *dnsProto.UpsertRecordSetsRequest,
		opts ...grpc.CallOption,
	) (*opProto.Operation, error)
}

// DNSZoneClientAdapter is an implementation of DNSClient interface for *dns.DnsZoneServiceClient.
type DNSZoneClientAdapter struct {
	c *dns.DnsZoneServiceClient
}

// ZoneIterator returns an ZoneIteratorAdapter that iterates over DNS zones (dns.DnsZoneIterator).
func (a *DNSZoneClientAdapter) ZoneIterator(ctx context.Context,
	req *dnsProto.ListDnsZonesRequest,
	opts ...grpc.CallOption,
) ZoneIteratorAdapter {
	return a.c.DnsZoneIterator(ctx, req, opts...)
}

// RecordSetIterator returns an RecordSetIteratorAdapter that iterates over DNS record sets (dns.DnsZoneRecordSetsIterator).
func (a *DNSZoneClientAdapter) RecordSetIterator(ctx context.Context,
	req *dnsProto.ListDnsZoneRecordSetsRequest,
	opts ...grpc.CallOption,
) RecordSetIteratorAdapter {
	return a.c.DnsZoneRecordSetsIterator(ctx, req, opts...)
}

// UpsertRecordSets executes dns.UpsertRecordSets operation.
func (a *DNSZoneClientAdapter) UpsertRecordSets(ctx context.Context,
	in *dnsProto.UpsertRecordSetsRequest,
	opts ...grpc.CallOption,
) (*opProto.Operation, error) {
	return a.c.UpsertRecordSets(ctx, in, opts...)
}
