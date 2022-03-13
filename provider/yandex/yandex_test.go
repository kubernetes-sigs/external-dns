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
	"testing"

	"github.com/stretchr/testify/assert"
	dnsProto "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/provider"
)

func assertEndpointsAreSame(t *testing.T, expected, actual []*endpoint.Endpoint) {
	assert.True(t, testutils.SameEndpoints(expected, actual), "expected and actual endpoints don't match. %s:%s", actual, expected)
}

func TestYandexRecord(t *testing.T) {
	f := newFixture().
		WithZoneRecords("yandex.io.",
			&dnsProto.RecordSet{Name: "test.yandex.io.", Type: "A", Ttl: 10, Data: []string{"1.2.3.4"}},
			&dnsProto.RecordSet{Name: "another.yandex.io.", Type: "CNAME", Ttl: 10, Data: []string{"test2.yandex.io"}},
		)

	actual, err := f.Provider().Records(context.Background())

	if err != nil {
		t.Fatal(err)
	}
	assertEndpointsAreSame(t, []*endpoint.Endpoint{
		{DNSName: "test.yandex.io", Targets: []string{"1.2.3.4"}, RecordType: "A", RecordTTL: 10},
		{DNSName: "another.yandex.io", Targets: []string{"test2.yandex.io"}, RecordType: "CNAME", RecordTTL: 10},
	}, actual)
}

func TestYandexRecordWithDomainFilter(t *testing.T) {
	f := newFixture().
		WithZoneRecords("another.io.",
			&dnsProto.RecordSet{Name: "internal.another.io.", Type: "A", Ttl: 10, Data: []string{"1.2.3.4"}},
			&dnsProto.RecordSet{Name: "stub.another.io.", Type: "CNAME", Ttl: 10, Data: []string{"test2.yandex.io"}},
		).
		WithZoneRecords("yandex.io.",
			&dnsProto.RecordSet{Name: "test.yandex.io.", Type: "A", Ttl: 10, Data: []string{"1.2.3.4"}},
			&dnsProto.RecordSet{Name: "another.yandex.io.", Type: "CNAME", Ttl: 10, Data: []string{"test2.yandex.io"}},
		).
		WithDomainFilter("yandex.io")

	actual, err := f.Provider().Records(context.Background())

	if err != nil {
		t.Fatal(err)
	}
	assertEndpointsAreSame(t, []*endpoint.Endpoint{
		{DNSName: "test.yandex.io", Targets: []string{"1.2.3.4"}, RecordType: "A", RecordTTL: 10},
		{DNSName: "another.yandex.io", Targets: []string{"test2.yandex.io"}, RecordType: "CNAME", RecordTTL: 10},
	}, actual)
}
