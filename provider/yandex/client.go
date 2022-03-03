package yandex

import (
	"context"

	dnsProto "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	opProto "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"github.com/yandex-cloud/go-sdk/gen/dns"
	"google.golang.org/grpc"
)

type DnsZoneIteratorAdapter interface {
	Next() bool
	Error() error
	Value() *dnsProto.DnsZone
}

type DnsZoneRecordSetIteratorAdapter interface {
	Next() bool
	Error() error
	Value() *dnsProto.RecordSet
}

type DnsZoneClient interface {
	DnsZoneIterator(ctx context.Context,
		req *dnsProto.ListDnsZonesRequest,
		opts ...grpc.CallOption,
	) DnsZoneIteratorAdapter

	DnsZoneRecordSetsIterator(ctx context.Context,
		req *dnsProto.ListDnsZoneRecordSetsRequest,
		opts ...grpc.CallOption,
	) DnsZoneRecordSetIteratorAdapter

	UpsertRecordSets(ctx context.Context,
		in *dnsProto.UpsertRecordSetsRequest,
		opts ...grpc.CallOption,
	) (*opProto.Operation, error)
}

type DNSZoneClientAdapter struct {
	c *dns.DnsZoneServiceClient
}

func (a *DNSZoneClientAdapter) DnsZoneIterator(ctx context.Context,
	req *dnsProto.ListDnsZonesRequest,
	opts ...grpc.CallOption,
) DnsZoneIteratorAdapter {
	return a.c.DnsZoneIterator(ctx, req, opts...)
}

func (a *DNSZoneClientAdapter) DnsZoneRecordSetsIterator(ctx context.Context,
	req *dnsProto.ListDnsZoneRecordSetsRequest,
	opts ...grpc.CallOption,
) DnsZoneRecordSetIteratorAdapter {
	return a.c.DnsZoneRecordSetsIterator(ctx, req, opts...)
}

func (a *DNSZoneClientAdapter) UpsertRecordSets(ctx context.Context,
	in *dnsProto.UpsertRecordSetsRequest,
	opts ...grpc.CallOption,
) (*opProto.Operation, error) {
	return a.c.UpsertRecordSets(ctx, in, opts...)
}
