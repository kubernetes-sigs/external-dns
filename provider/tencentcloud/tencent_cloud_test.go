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

package tencentcloud

import (
	"context"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/tencentcloud/cloudapi"

	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
)

func NewMockTencentCloudProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, zoneType string) *TencentCloudProvider {
	cfg := tencentCloudConfig{
		// SecretId:  "",
		// SecretKey: "",
		RegionId: "ap-shanghai",
		VPCId:    "vpc-abcdefg",
	}

	zoneId1 := common.StringPtr(cloudapi.RandStringRunes(8))

	privateZones := []*privatedns.PrivateZone{
		{
			ZoneId: zoneId1,
			Domain: common.StringPtr("external-dns-test.com"),
			VpcSet: []*privatedns.VpcInfo{
				{
					UniqVpcId: common.StringPtr("vpc-abcdefg"),
					Region:    common.StringPtr("ap-shanghai"),
				},
			},
		},
	}

	zoneRecordId1 := common.StringPtr(cloudapi.RandStringRunes(8))
	zoneRecordId2 := common.StringPtr(cloudapi.RandStringRunes(8))
	privateZoneRecords := map[string][]*privatedns.PrivateZoneRecord{
		*zoneId1: {
			{
				ZoneId:      zoneId1,
				RecordId:    zoneRecordId1,
				SubDomain:   common.StringPtr("nginx"),
				RecordType:  common.StringPtr("TXT"),
				RecordValue: common.StringPtr("heritage=external-dns,external-dns/owner=default"),
				TTL:         common.Int64Ptr(300),
			},
			{
				ZoneId:      zoneId1,
				RecordId:    zoneRecordId2,
				SubDomain:   common.StringPtr("nginx"),
				RecordType:  common.StringPtr("A"),
				RecordValue: common.StringPtr("10.10.10.10"),
				TTL:         common.Int64Ptr(300),
			},
		},
	}

	dnsDomainId1 := common.Uint64Ptr(cloudapi.RandUint64())
	dnsPodDomains := []*dnspod.DomainListItem{
		{
			DomainId: dnsDomainId1,
			Name:     common.StringPtr("external-dns-test.com"),
		},
	}
	dnsDomainRecordId1 := common.Uint64Ptr(cloudapi.RandUint64())
	dnsDomainRecordId2 := common.Uint64Ptr(cloudapi.RandUint64())
	dnspodRecords := map[string][]*dnspod.RecordListItem{
		"external-dns-test.com": {
			{
				RecordId: dnsDomainRecordId1,
				Value:    common.StringPtr("heritage=external-dns,external-dns/owner=default"),
				Name:     common.StringPtr("nginx"),
				Type:     common.StringPtr("TXT"),
				TTL:      common.Uint64Ptr(300),
			},
			{
				RecordId: dnsDomainRecordId2,
				Name:     common.StringPtr("nginx"),
				Type:     common.StringPtr("A"),
				Value:    common.StringPtr("10.10.10.10"),
				TTL:      common.Uint64Ptr(300),
			},
		},
	}

	var apiService cloudapi.TencentAPIService = cloudapi.NewMockService(privateZones, privateZoneRecords, dnsPodDomains, dnspodRecords)

	tencentCloudProvider := &TencentCloudProvider{
		domainFilter: domainFilter,
		zoneIDFilter: zoneIDFilter,
		apiService:   apiService,
		vpcID:        cfg.VPCId,
		privateZone:  zoneType == "private",
	}

	return tencentCloudProvider
}

func TestTencentPrivateProvider_Records(t *testing.T) {
	p := NewMockTencentCloudProvider(endpoint.NewDomainFilter([]string{"external-dns-test.com"}), provider.NewZoneIDFilter([]string{}), "private")
	endpoints, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Create、UpdateOld、UpdateNew、Delete
	// The base record will be created.
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "redis.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("4.3.2.1"),
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "nginx.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("10.10.10.10"),
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("1.2.3.4", "5.6.7.8"),
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "nginx.external-dns-test.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 3 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete one target
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 3 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete another target
	changes = &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "redis.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 3 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete another target
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("1.2.3.4"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}
}

func TestTencentPublicProvider_Records(t *testing.T) {
	p := NewMockTencentCloudProvider(endpoint.NewDomainFilter([]string{"external-dns-test.com"}), provider.NewZoneIDFilter([]string{}), "public")
	endpoints, err := p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Create、UpdateOld、UpdateNew、Delete
	changes := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "redis.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("4.3.2.1"),
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName:    "nginx.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("10.10.10.10"),
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("1.2.3.4", "5.6.7.8"),
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "nginx.external-dns-test.com",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete one target
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete another target
	changes = &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "redis.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("5.6.7.8"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 2 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}

	// Test for Delete another target
	changes = &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "tencent.external-dns-test.com",
				RecordType: "A",
				RecordTTL:  600,
				Targets:    endpoint.NewTargets("1.2.3.4"),
			},
		},
	}
	if err := p.ApplyChanges(context.Background(), changes); err != nil {
		t.Errorf("Failed to get records: %v", err)
	}
	endpoints, err = p.Records(context.Background())
	if err != nil {
		t.Errorf("Failed to get records: %v", err)
	} else {
		if len(endpoints) != 1 {
			t.Errorf("Incorrect number of records: %d", len(endpoints))
		}
		for _, endpoint := range endpoints {
			t.Logf("Endpoint for %+v", *endpoint)
		}
	}
}
