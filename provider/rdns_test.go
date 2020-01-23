/*
Copyright 2019 The Kubernetes Authors.

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

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type fakeEtcdv3Client struct {
	rs map[string]RDNSRecord
}

func (c fakeEtcdv3Client) Get(key string) ([]RDNSRecord, error) {
	rs := make([]RDNSRecord, 0)
	for k, v := range c.rs {
		if strings.Contains(k, key) {
			rs = append(rs, v)
		}
	}
	return rs, nil
}

func (c fakeEtcdv3Client) List(rootDomain string) ([]RDNSRecord, error) {
	var result []RDNSRecord
	for key, value := range c.rs {
		rootPath := rdnsPrefix + dnsNameToKey(rootDomain)
		if strings.HasPrefix(key, rootPath) {
			value.Key = key
			result = append(result, value)
		}
	}

	r := &clientv3.GetResponse{}

	for _, v := range result {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		k := &mvccpb.KeyValue{
			Key:   []byte(v.Key),
			Value: b,
		}

		r.Kvs = append(r.Kvs, k)
	}

	return c.aggregationRecords(r)
}

func (c fakeEtcdv3Client) Set(r RDNSRecord) error {
	c.rs[r.Key] = r
	return nil
}

func (c fakeEtcdv3Client) Delete(key string) error {
	ks := make([]string, 0)
	for k := range c.rs {
		if strings.Contains(k, key) {
			ks = append(ks, k)
		}
	}

	for _, v := range ks {
		delete(c.rs, v)
	}

	return nil
}

func TestARecordTranslation(t *testing.T) {
	expectedTarget1 := "1.2.3.4"
	expectedTarget2 := "2.3.4.5"
	expectedTargets := []string{expectedTarget1, expectedTarget2}
	expectedDNSName := "p1xaf1.lb.rancher.cloud"
	expectedRecordType := endpoint.RecordTypeA

	client := fakeEtcdv3Client{
		map[string]RDNSRecord{
			"/rdnsv3/cloud/rancher/lb/p1xaf1/1_2_3_4": {Host: expectedTarget1},
			"/rdnsv3/cloud/rancher/lb/p1xaf1/2_3_4_5": {Host: expectedTarget2},
		},
	}

	provider := RDNSProvider{
		client:     client,
		rootDomain: "lb.rancher.cloud",
	}

	endpoints, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(endpoints) != 1 {
		t.Fatalf("got unexpected number of endpoints: %d", len(endpoints))
	}

	ep := endpoints[0]
	if ep.DNSName != expectedDNSName {
		t.Errorf("got unexpected DNS name: %s != %s", ep.DNSName, expectedDNSName)
	}
	assert.Contains(t, expectedTargets, ep.Targets[0])
	assert.Contains(t, expectedTargets, ep.Targets[1])
	if ep.RecordType != expectedRecordType {
		t.Errorf("got unexpected DNS record type: %s != %s", ep.RecordType, expectedRecordType)
	}
}

func TestTXTRecordTranslation(t *testing.T) {
	expectedTarget := "string"
	expectedDNSName := "p1xaf1.lb.rancher.cloud"
	expectedRecordType := endpoint.RecordTypeTXT

	client := fakeEtcdv3Client{
		map[string]RDNSRecord{
			"/rdnsv3/cloud/rancher/lb/p1xaf1": {Text: expectedTarget},
		},
	}

	provider := RDNSProvider{
		client:     client,
		rootDomain: "lb.rancher.cloud",
	}

	endpoints, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(endpoints) != 1 {
		t.Fatalf("got unexpected number of endpoints: %d", len(endpoints))
	}
	if endpoints[0].DNSName != expectedDNSName {
		t.Errorf("got unexpected DNS name: %s != %s", endpoints[0].DNSName, expectedDNSName)
	}
	if endpoints[0].Targets[0] != expectedTarget {
		t.Errorf("got unexpected DNS target: %s != %s", endpoints[0].Targets[0], expectedTarget)
	}
	if endpoints[0].RecordType != expectedRecordType {
		t.Errorf("got unexpected DNS record type: %s != %s", endpoints[0].RecordType, expectedRecordType)
	}
}

func TestAWithTXTRecordTranslation(t *testing.T) {
	expectedTargets := map[string]string{
		endpoint.RecordTypeA:   "1.2.3.4",
		endpoint.RecordTypeTXT: "string",
	}
	expectedDNSName := "p1xaf1.lb.rancher.cloud"

	client := fakeEtcdv3Client{
		map[string]RDNSRecord{
			"/rdnsv3/cloud/rancher/lb/p1xaf1": {Host: "1.2.3.4", Text: "string"},
		},
	}

	provider := RDNSProvider{
		client:     client,
		rootDomain: "lb.rancher.cloud",
	}

	endpoints, err := provider.Records(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(endpoints) != len(expectedTargets) {
		t.Fatalf("got unexpected number of endpoints: %d", len(endpoints))
	}

	for _, ep := range endpoints {
		expectedTarget := expectedTargets[ep.RecordType]
		if expectedTarget == "" {
			t.Errorf("got unexpected DNS record type: %s", ep.RecordType)
			continue
		}
		delete(expectedTargets, ep.RecordType)

		if ep.DNSName != expectedDNSName {
			t.Errorf("got unexpected DNS name: %s != %s", ep.DNSName, expectedDNSName)
		}

		if ep.Targets[0] != expectedTarget {
			t.Errorf("got unexpected DNS target: %s != %s", ep.Targets[0], expectedTarget)
		}
	}
}

func TestRDNSApplyChanges(t *testing.T) {
	client := fakeEtcdv3Client{
		map[string]RDNSRecord{},
	}

	provider := RDNSProvider{
		client:     client,
		rootDomain: "lb.rancher.cloud",
	}

	changes1 := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("p1xaf1.lb.rancher.cloud", endpoint.RecordTypeA, "5.5.5.5", "6.6.6.6"),
			endpoint.NewEndpoint("p1xaf1.lb.rancher.cloud", endpoint.RecordTypeTXT, "string1"),
		},
	}

	if err := provider.ApplyChanges(context.Background(), changes1); err != nil {
		t.Error(err)
	}

	expectedRecords1 := map[string]RDNSRecord{
		"/rdnsv3/cloud/rancher/lb/p1xaf1/5_5_5_5": {Host: "5.5.5.5", Text: "string1"},
		"/rdnsv3/cloud/rancher/lb/p1xaf1/6_6_6_6": {Host: "6.6.6.6", Text: "string1"},
	}

	client.validateRecords(client.rs, expectedRecords1, t)

	changes2 := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("abx1v1.lb.rancher.cloud", endpoint.RecordTypeA, "7.7.7.7"),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpoint("p1xaf1.lb.rancher.cloud", endpoint.RecordTypeA, "8.8.8.8", "9.9.9.9"),
		},
	}

	records, _ := provider.Records(context.Background())
	for _, ep := range records {
		if ep.DNSName == "p1xaf1.lb.rancher.cloud" {
			changes2.UpdateOld = append(changes2.UpdateOld, ep)
		}
	}

	if err := provider.ApplyChanges(context.Background(), changes2); err != nil {
		t.Error(err)
	}

	expectedRecords2 := map[string]RDNSRecord{
		"/rdnsv3/cloud/rancher/lb/p1xaf1/8_8_8_8": {Host: "8.8.8.8"},
		"/rdnsv3/cloud/rancher/lb/p1xaf1/9_9_9_9": {Host: "9.9.9.9"},
		"/rdnsv3/cloud/rancher/lb/abx1v1/7_7_7_7": {Host: "7.7.7.7"},
	}

	client.validateRecords(client.rs, expectedRecords2, t)

	changes3 := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpoint("p1xaf1.lb.rancher.cloud", endpoint.RecordTypeA, "8.8.8.8", "9.9.9.9"),
		},
	}

	if err := provider.ApplyChanges(context.Background(), changes3); err != nil {
		t.Error(err)
	}

	expectedRecords3 := map[string]RDNSRecord{
		"/rdnsv3/cloud/rancher/lb/abx1v1/7_7_7_7": {Host: "7.7.7.7"},
	}

	client.validateRecords(client.rs, expectedRecords3, t)

}

func (c fakeEtcdv3Client) aggregationRecords(result *clientv3.GetResponse) ([]RDNSRecord, error) {
	var rs []RDNSRecord
	bx := make(map[RDNSRecordType]RDNSRecord)

	for _, n := range result.Kvs {
		r := new(RDNSRecord)
		if err := json.Unmarshal(n.Value, r); err != nil {
			return nil, fmt.Errorf("%s: %s", n.Key, err.Error())
		}

		r.Key = string(n.Key)

		if r.Host == "" && r.Text == "" {
			continue
		}

		if r.Host != "" {
			c := RDNSRecord{
				AggregationHosts: r.AggregationHosts,
				Host:             r.Host,
				Text:             r.Text,
				TTL:              r.TTL,
				Key:              r.Key,
			}
			n, isContinue := appendRecords(c, endpoint.RecordTypeA, bx, rs)
			if isContinue {
				continue
			}
			rs = n
		}

		if r.Text != "" && r.Host == "" {
			c := RDNSRecord{
				AggregationHosts: []string{},
				Host:             r.Host,
				Text:             r.Text,
				TTL:              r.TTL,
				Key:              r.Key,
			}
			n, isContinue := appendRecords(c, endpoint.RecordTypeTXT, bx, rs)
			if isContinue {
				continue
			}
			rs = n
		}
	}

	return rs, nil
}

func (c fakeEtcdv3Client) validateRecords(rs, expectedRs map[string]RDNSRecord, t *testing.T) {
	if len(rs) != len(expectedRs) {
		t.Errorf("wrong number of records: %d != %d", len(rs), len(expectedRs))
	}
	for key, value := range rs {
		if _, ok := expectedRs[key]; !ok {
			t.Errorf("unexpected record %s", key)
			continue
		}
		expected := expectedRs[key]
		delete(expectedRs, key)
		if value.Host != expected.Host {
			t.Errorf("wrong host for record %s: %s != %s", key, value.Host, expected.Host)
		}
		if value.Text != expected.Text {
			t.Errorf("wrong text for record %s: %s != %s", key, value.Text, expected.Text)
		}
	}
}
