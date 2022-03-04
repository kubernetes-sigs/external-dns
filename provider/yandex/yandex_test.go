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

	dnsInt "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	op "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"github.com/yandex-cloud/go-sdk/gen/dns"
	"google.golang.org/grpc"
)

type mockDNSClient struct {
}

type mockZoneIterator struct {
}

type mockRecordSetIterator struct {
}

func (c *mockDNSClient) ZoneIterator(ctx context.Context, req *dnsInt.ListDnsZonesRequest, _ ...grpc.CallOption) ZoneIteratorAdapter {
	return &dns.DnsZoneIterator{}
}

func (c *mockDNSClient) RecordSetsIterator(ctx context.Context, req *dnsInt.ListDnsZoneRecordSetsRequest, _ ...grpc.CallOption) RecordSetIteratorAdapter {
	return nil
}

func (c *mockDNSClient) UpsertRecordSets(ctx context.Context, in *dnsInt.UpsertRecordSetsRequest, _ ...grpc.CallOption) (*op.Operation, error) {
	return nil, nil
}
