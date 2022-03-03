package yandex

import (
	"context"

	dnsInt "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	op "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"github.com/yandex-cloud/go-sdk/gen/dns"
	"google.golang.org/grpc"
)

type mockDNSZoneClient struct {
}

type mockDnsZoneIterator struct {
}

type mockDnsZoneRecordSetIterator struct {
}

func (c *mockDNSZoneClient) DnsZoneIterator(ctx context.Context, req *dnsInt.ListDnsZonesRequest, _ ...grpc.CallOption) DnsZoneIteratorAdapter {
	return &dns.DnsZoneIterator{}
}

func (c *mockDNSZoneClient) DnsZoneRecordSetsIterator(ctx context.Context, req *dnsInt.ListDnsZoneRecordSetsRequest, _ ...grpc.CallOption) DnsZoneRecordSetIteratorAdapter {
	return nil
}

func (c *mockDNSZoneClient) UpsertRecordSets(ctx context.Context, in *dnsInt.UpsertRecordSetsRequest, _ ...grpc.CallOption) (*op.Operation, error) {
	return nil, nil
}
