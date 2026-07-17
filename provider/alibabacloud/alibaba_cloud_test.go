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

package alibabacloud

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/pvtz"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
)

type MockAlibabaCloudDNSAPI struct {
	counter int64
	domains []*alidns.Domain
	records map[string][]*alidns.Record
}

func (m *MockAlibabaCloudDNSAPI) AddDomainRecord(request *alidns.AddDomainRecordRequest) (*alidns.AddDomainRecordResponse, error) {
	if _, exist := m.records[request.DomainName]; !exist {
		return alidns.CreateAddDomainRecordResponse(), fmt.Errorf("Domain not found")
	}

	// Check if the record already exists.
	value := unwrapQuotes(request.Type, request.Value)
	records, _ := m.records[request.DomainName]
	for _, record := range records {
		if record.Type == request.Type && record.RR == request.RR && record.Value == value {
			return alidns.CreateAddDomainRecordResponse(), fmt.Errorf("The DNS record already exists")
		}
	}

	ttl, _ := request.TTL.GetValue64()
	recordId := m.nextID()
	m.records[request.DomainName] = append(m.records[request.DomainName], &alidns.Record{
		RecordId:   recordId,
		DomainName: request.DomainName,
		Type:       request.Type,
		RR:         request.RR,
		TTL:        ttl,
		Value:      value,
	})

	response := alidns.CreateAddDomainRecordResponse()
	response.RecordId = recordId
	return response, nil
}

func (m *MockAlibabaCloudDNSAPI) DeleteDomainRecord(request *alidns.DeleteDomainRecordRequest) (*alidns.DeleteDomainRecordResponse, error) {
	if request == nil || request.RecordId == "" {
		return alidns.CreateDeleteDomainRecordResponse(), fmt.Errorf("Invalid request")
	}

	for domain, records := range m.records {
		m.records[domain] = slices.DeleteFunc(records, func(item *alidns.Record) bool {
			return item.RecordId == request.RecordId
		})
	}
	response := alidns.CreateDeleteDomainRecordResponse()
	response.RecordId = request.RecordId
	return response, nil
}

func (m *MockAlibabaCloudDNSAPI) UpdateDomainRecord(request *alidns.UpdateDomainRecordRequest) (*alidns.UpdateDomainRecordResponse, error) {
	if request == nil || request.RecordId == "" {
		return alidns.CreateUpdateDomainRecordResponse(), fmt.Errorf("Invalid request")
	}

	ttl, _ := request.TTL.GetValue64()
	for _, records := range m.records {
		idx := slices.IndexFunc(records, func(item *alidns.Record) bool {
			return item.RecordId == request.RecordId
		})
		if idx != -1 {
			record := records[idx]
			if request.Value != "" {
				value := unwrapQuotes(record.Type, request.Value)
				// Check if the modified value already exists.
				if exist := slices.ContainsFunc(records, func(item *alidns.Record) bool {
					return item.RecordId != request.RecordId &&
						item.Type == record.Type &&
						item.RR == record.RR &&
						item.Value == value
				}); exist {
					return alidns.CreateUpdateDomainRecordResponse(), fmt.Errorf("The DNS record value is duplicated.")
				}
				records[idx].Value = value
			}
			if ttl > 0 {
				records[idx].TTL = ttl
			}
			break
		}
	}
	response := alidns.CreateUpdateDomainRecordResponse()
	response.RecordId = request.RecordId
	return response, nil
}

func (m *MockAlibabaCloudDNSAPI) DescribeDomains(_ *alidns.DescribeDomainsRequest) (*alidns.DescribeDomainsResponse, error) {
	var result alidns.DomainsInDescribeDomains
	for _, domain := range m.domains {
		result.Domain = append(result.Domain, alidns.DomainInDescribeDomains{
			DomainId:   domain.DomainId,
			DomainName: domain.DomainName,
		})
	}
	response := alidns.CreateDescribeDomainsResponse()
	response.Domains = result
	return response, nil
}

func (m *MockAlibabaCloudDNSAPI) DescribeDomainRecords(request *alidns.DescribeDomainRecordsRequest) (*alidns.DescribeDomainRecordsResponse, error) {
	if request == nil || request.DomainName == "" {
		return alidns.CreateDescribeDomainRecordsResponse(), fmt.Errorf("Invalid request")
	}

	var result []alidns.Record
	if records, exist := m.records[request.DomainName]; exist {
		for _, record := range records {
			result = append(result, *record)
		}
	}
	response := alidns.CreateDescribeDomainRecordsResponse()
	response.DomainRecords.Record = result
	return response, nil
}

func (m *MockAlibabaCloudDNSAPI) nextID() string {
	id := atomic.AddInt64(&m.counter, 1)
	return strconv.FormatInt(id, 10)
}

func (m *MockAlibabaCloudDNSAPI) newRecord(ep *endpoint.Endpoint, target, domain string) *alidns.Record {
	subname := getSubname(domain, ep)
	value := unwrapQuotes(ep.RecordType, target)

	return &alidns.Record{
		DomainName: domain,
		RecordId:   m.nextID(),
		Type:       ep.RecordType,
		TTL:        int64(ep.RecordTTL),
		RR:         subname,
		Value:      value,
	}
}

type MockAlibabaCloudPrivateZoneAPI struct {
	counter int64
	zones   []*pvtz.Zone
	records map[string][]*pvtz.Record
}

func (m *MockAlibabaCloudPrivateZoneAPI) AddZoneRecord(request *pvtz.AddZoneRecordRequest) (*pvtz.AddZoneRecordResponse, error) {
	if _, exist := m.records[request.ZoneId]; !exist {
		return pvtz.CreateAddZoneRecordResponse(), fmt.Errorf("Zone not found")
	}

	// Check if the record already exists.
	value := unwrapQuotes(request.Type, request.Value)
	records, _ := m.records[request.ZoneId]
	for _, record := range records {
		if record.Type == request.Type && record.Rr == request.Rr && record.Value == value {
			return pvtz.CreateAddZoneRecordResponse(), fmt.Errorf("The DNS record already exists")
		}
	}

	ttl, _ := request.Ttl.GetValue()
	recordId := m.nextID()
	m.records[request.ZoneId] = append(m.records[request.ZoneId], &pvtz.Record{
		RecordId: recordId,
		ZoneId:   request.ZoneId,
		Type:     request.Type,
		Rr:       request.Rr,
		Ttl:      ttl,
		Value:    value,
	})

	response := pvtz.CreateAddZoneRecordResponse()
	response.RecordId = recordId
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) DeleteZoneRecord(request *pvtz.DeleteZoneRecordRequest) (*pvtz.DeleteZoneRecordResponse, error) {
	if request == nil || request.RecordId == "" {
		return pvtz.CreateDeleteZoneRecordResponse(), fmt.Errorf("Invalid request")
	}

	recordID, _ := request.RecordId.GetValue64()
	for domain, records := range m.records {
		m.records[domain] = slices.DeleteFunc(records, func(item *pvtz.Record) bool {
			return item.RecordId == recordID
		})
	}
	response := pvtz.CreateDeleteZoneRecordResponse()
	response.RecordId = recordID
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) UpdateZoneRecord(request *pvtz.UpdateZoneRecordRequest) (*pvtz.UpdateZoneRecordResponse, error) {
	if request == nil || request.RecordId == "" {
		return pvtz.CreateUpdateZoneRecordResponse(), fmt.Errorf("Invalid request")
	}

	ttl, _ := request.Ttl.GetValue()
	recordID, _ := request.RecordId.GetValue64()
	for _, records := range m.records {
		idx := slices.IndexFunc(records, func(item *pvtz.Record) bool {
			return item.RecordId == recordID
		})
		if idx != -1 {
			record := records[idx]
			if request.Value != "" {
				value := unwrapQuotes(record.Type, request.Value)
				// Check if the modified value already exists.
				if exist := slices.ContainsFunc(records, func(item *pvtz.Record) bool {
					return item.RecordId != recordID &&
						item.Type == record.Type &&
						item.Rr == record.Rr &&
						item.Value == value
				}); exist {
					return pvtz.CreateUpdateZoneRecordResponse(), fmt.Errorf("The DNS record value is duplicated.")
				}
				records[idx].Value = value
			}
			if ttl > 0 {
				records[idx].Ttl = ttl
			}
			break
		}
	}
	response := pvtz.CreateUpdateZoneRecordResponse()
	response.RecordId = recordID
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) DescribeZoneRecords(request *pvtz.DescribeZoneRecordsRequest) (*pvtz.DescribeZoneRecordsResponse, error) {
	if request == nil || request.ZoneId == "" {
		return pvtz.CreateDescribeZoneRecordsResponse(), fmt.Errorf("Invalid request")
	}

	var result []pvtz.Record
	if records, exist := m.records[request.ZoneId]; exist {
		for _, record := range records {
			result = append(result, *record)
		}
	}
	response := pvtz.CreateDescribeZoneRecordsResponse()
	response.Records.Record = result
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) DescribeZones(_ *pvtz.DescribeZonesRequest) (*pvtz.DescribeZonesResponse, error) {
	var result pvtz.ZonesInDescribeZones
	for _, zone := range m.zones {
		result.Zone = append(result.Zone, pvtz.Zone{
			ZoneId:   zone.ZoneId,
			ZoneName: zone.ZoneName,
		})
	}
	response := pvtz.CreateDescribeZonesResponse()
	response.Zones = result
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) DescribeZoneInfo(request *pvtz.DescribeZoneInfoRequest) (*pvtz.DescribeZoneInfoResponse, error) {
	if request == nil || request.ZoneId == "" {
		return pvtz.CreateDescribeZoneInfoResponse(), fmt.Errorf("Invalid request")
	}

	response := pvtz.CreateDescribeZoneInfoResponse()
	for _, zone := range m.zones {
		if zone.ZoneId == request.ZoneId {
			response.ZoneId = zone.ZoneId
			response.ZoneName = zone.ZoneName
			response.BindVpcs = pvtz.BindVpcsInDescribeZoneInfo{Vpc: make([]pvtz.VpcInDescribeZoneInfo, len(zone.Vpcs.Vpc))}
			for i, vpc := range zone.Vpcs.Vpc {
				response.BindVpcs.Vpc[i] = pvtz.VpcInDescribeZoneInfo{VpcName: vpc.VpcName, VpcId: vpc.VpcId, VpcType: vpc.VpcType, RegionName: vpc.RegionName, RegionId: vpc.RegionId}
			}
		}
	}
	return response, nil
}

func (m *MockAlibabaCloudPrivateZoneAPI) nextID() int64 {
	id := atomic.AddInt64(&m.counter, 1)
	return id
}

func (m *MockAlibabaCloudPrivateZoneAPI) newRecord(ep *endpoint.Endpoint, target, domain, zoneId string) *pvtz.Record {
	subname := getSubname(domain, ep)
	value := unwrapQuotes(ep.RecordType, target)

	return &pvtz.Record{
		ZoneId:   zoneId,
		RecordId: m.nextID(),
		Type:     ep.RecordType,
		Ttl:      int(ep.RecordTTL),
		Rr:       subname,
		Value:    value,
	}
}

func newTestAlibabaCloudProvider(private bool) *AlibabaCloudProvider {
	domain := "container-service.top"
	return newTestAlibabaCloudProviderWithConfig(
		endpoint.NewDomainFilter([]string{domain}),
		private, map[string][]*endpoint.Endpoint{
			domain: createDefaultEndpoints(domain),
		},
	)
}

func newTestAlibabaCloudProviderWithConfig(domainFilter *endpoint.DomainFilter, private bool, endpointsMap map[string][]*endpoint.Endpoint) *AlibabaCloudProvider {
	cfg := alibabaCloudConfig{
		RegionID: "cn-beijing",
		VPCID:    "vpc-xxxxxx",
	}

	dnsClient := &MockAlibabaCloudDNSAPI{
		counter: 0,
		domains: make([]*alidns.Domain, 0),
		records: make(map[string][]*alidns.Record, 0),
	}
	pvtzClient := &MockAlibabaCloudPrivateZoneAPI{
		counter: 0,
		zones:   make([]*pvtz.Zone, 0),
		records: make(map[string][]*pvtz.Record, 0),
	}

	if private {
		for domain, endpoints := range endpointsMap {
			zoneId := uuid.New().String()
			pvtzClient.zones = append(pvtzClient.zones, &pvtz.Zone{
				ZoneId:   zoneId,
				ZoneName: domain,
				Vpcs: pvtz.Vpcs{
					Vpc: []pvtz.Vpc{{
						RegionId: cfg.RegionID,
						VpcId:    cfg.VPCID,
					}},
				},
			})
			pvtzClient.records[zoneId] = make([]*pvtz.Record, 0)
			for _, ep := range endpoints {
				for _, target := range ep.Targets {
					pvtzClient.records[zoneId] = append(
						pvtzClient.records[zoneId],
						pvtzClient.newRecord(ep, target, domain, zoneId),
					)
				}
			}
		}
	} else {
		for domain, endpoints := range endpointsMap {
			dnsClient.domains = append(dnsClient.domains, &alidns.Domain{
				DomainId:   uuid.New().String(),
				DomainName: domain,
			})
			dnsClient.records[domain] = make([]*alidns.Record, 0)
			for _, ep := range endpoints {
				for _, target := range ep.Targets {
					dnsClient.records[domain] = append(
						dnsClient.records[domain],
						dnsClient.newRecord(ep, target, domain),
					)
				}
			}
		}
	}

	return &AlibabaCloudProvider{
		domainFilter: domainFilter,
		vpcID:        cfg.VPCID,
		dryRun:       false,
		dnsClient:    dnsClient,
		pvtzClient:   pvtzClient,
		privateZone:  private,
	}
}

func getSubname(domain string, ep *endpoint.Endpoint) string {
	name := strings.TrimSuffix(ep.DNSName, ".")
	name = strings.TrimSuffix(name, strings.TrimSuffix(domain, "."))
	name = strings.TrimSuffix(name, ".")

	if name == "" {
		return "@"
	}
	return name
}

func createDefaultEndpoints(domain string) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("abc."+domain, "A", 300, "1.2.3.4"),
		endpoint.NewEndpointWithTTL("abc."+domain, "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
		endpoint.NewEndpointWithTTL("a-abc."+domain, "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
	}

	return endpoints
}

func TestAlibabaCloudProvider_Records(t *testing.T) {
	domain := "container-service.top"
	defaultEndpoints := createDefaultEndpoints(domain)
	provider := newTestAlibabaCloudProviderWithConfig(
		endpoint.NewDomainFilter([]string{domain}), false, map[string][]*endpoint.Endpoint{
			domain: defaultEndpoints,
		},
	)
	endpoints, err := provider.Records(t.Context())

	require.NoError(t, err, "Failed to get records: %v", err)
	require.Len(t, endpoints, len(defaultEndpoints), "Incorrect number of records: %d", len(endpoints))
	assert.True(t, testutils.SameEndpoints(defaultEndpoints, endpoints), "expected and actual endpoints don't match. %s:%s", defaultEndpoints, endpoints)
}

func TestAlibabaCloudProvider_ApplyChanges(t *testing.T) {
	provider := newTestAlibabaCloudProvider(false)
	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("xyz.container-service.top", "A", 300, "4.3.2.1"),
			endpoint.NewEndpointWithTTL("ttl.container-service.top", "A", defaultTTL, "4.3.2.1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("abc.container-service.top", "A", 500, "1.2.3.4", "5.6.7.8"),
			endpoint.NewEndpointWithTTL("a-abc.container-service.top", "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
		},
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("abc.container-service.top", "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
		},
	}

	ctx := t.Context()
	err := provider.ApplyChanges(ctx, &changes)
	require.NoError(t, err, "Failed to apply changes: %v", err)

	endpoints, err := provider.Records(ctx)
	require.NoError(t, err, "Failed to get records: %v", err)

	changedEndpoints := append([]*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("abc.container-service.top", "A", 500, "1.2.3.4", "5.6.7.8"),
		endpoint.NewEndpointWithTTL("a-abc.container-service.top", "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
	}, changes.Create...)

	require.Len(t, endpoints, len(changedEndpoints), "Incorrect number of records: %d", len(endpoints))
	assert.True(t, testutils.SameEndpoints(changedEndpoints, endpoints), "expected and actual endpoints don't match. %s:%s", changedEndpoints, endpoints)
}

func TestAlibabaCloudProvider_ApplyChanges_UndefinedZoneDomain(t *testing.T) {
	provider := newTestAlibabaCloudProvider(false)
	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			// no found this zone by API: DescribeDomains
			endpoint.NewEndpointWithTTL("www.example.com", "A", 300, "9.9.9.9"),
			// can create this domain record
			endpoint.NewEndpointWithTTL("ttl.container-service.top", "A", defaultTTL, "4.3.2.1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("abc.container-service.top", "A", 500, "1.2.3.4", "5.6.7.8"),
		},
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("abc.container-service.top", "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
		},
	}

	ctx := t.Context()
	err := provider.ApplyChanges(ctx, &changes)
	require.NoError(t, err, "Failed to apply changes: %v", err)

	endpoints, err := provider.Records(ctx)
	require.NoError(t, err, "Failed to get records: %v", err)

	changedEndpoints := append([]*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("abc.container-service.top", "A", 500, "1.2.3.4", "5.6.7.8"),
		endpoint.NewEndpointWithTTL("a-abc.container-service.top", "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
	}, changes.Create[1]) // Only one new one was created

	require.Len(t, endpoints, len(changedEndpoints), "Incorrect number of records: %d", len(endpoints))
	assert.True(t, testutils.SameEndpoints(changedEndpoints, endpoints), "expected and actual endpoints don't match. %s:%s", changedEndpoints, endpoints)
}

func TestAlibabaCloudProvider_PrivateZone_Records(t *testing.T) {
	domain := "container-service.top"
	defaultEndpoints := createDefaultEndpoints(domain)
	provider := newTestAlibabaCloudProviderWithConfig(
		endpoint.NewDomainFilter([]string{domain}), true, map[string][]*endpoint.Endpoint{
			domain: defaultEndpoints,
		},
	)
	endpoints, err := provider.Records(t.Context())

	require.NoError(t, err, "Failed to get records: %v", err)
	require.Len(t, endpoints, len(defaultEndpoints), "Incorrect number of records: %d", len(endpoints))
	assert.True(t, testutils.SameEndpoints(defaultEndpoints, endpoints), "expected and actual endpoints don't match. %s:%s", defaultEndpoints, endpoints)
}

func TestAlibabaCloudProvider_PrivateZone_ApplyChanges(t *testing.T) {
	provider := newTestAlibabaCloudProvider(true)
	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("xyz.container-service.top", "A", 300, "4.3.2.1"),
		},
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("abc.container-service.top", "A", 500, "1.2.3.4", "5.6.7.8"),
			endpoint.NewEndpointWithTTL("a-abc.container-service.top", "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
		},
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpointWithTTL("abc.container-service.top", "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
		},
	}

	ctx := t.Context()
	err := provider.ApplyChanges(ctx, &changes)
	require.NoError(t, err, "Failed to apply changes: %v", err)

	endpoints, err := provider.Records(ctx)
	require.NoError(t, err, "Failed to get records: %v", err)

	changedEndpoints := append([]*endpoint.Endpoint{
		endpoint.NewEndpointWithTTL("abc.container-service.top", "A", 500, "1.2.3.4", "5.6.7.8"),
		endpoint.NewEndpointWithTTL("a-abc.container-service.top", "TXT", 300, "\"heritage=external-dns,external-dns/owner=default\""),
	}, changes.Create...)

	require.Len(t, endpoints, len(changedEndpoints), "Incorrect number of records: %d", len(endpoints))
	assert.True(t, testutils.SameEndpoints(changedEndpoints, endpoints), "expected and actual endpoints don't match. %s:%s", changedEndpoints, endpoints)
}

func TestAlibabaCloudProvider_splitDNSName(t *testing.T) {
	p := newTestAlibabaCloudProvider(false)
	endpoint := &endpoint.Endpoint{}
	hostedZoneDomains := []string{"container-service.top", "example.org"}

	var emptyZoneDomains []string

	endpoint.DNSName = "www.example.org"
	rr, domain := p.splitDNSName(endpoint.DNSName, hostedZoneDomains)
	if rr != "www" || domain != "example.org" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = ".example.org"
	rr, domain = p.splitDNSName(endpoint.DNSName, hostedZoneDomains)
	if rr != "@" || domain != "example.org" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "www"
	rr, domain = p.splitDNSName(endpoint.DNSName, hostedZoneDomains)
	if rr != "@" || domain != "" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = ""
	rr, domain = p.splitDNSName(endpoint.DNSName, hostedZoneDomains)
	if rr != "@" || domain != "" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "_30000._tcp.container-service.top"
	rr, domain = p.splitDNSName(endpoint.DNSName, hostedZoneDomains)
	if rr != "_30000._tcp" || domain != "container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "container-service.top"
	rr, domain = p.splitDNSName(endpoint.DNSName, hostedZoneDomains)
	if rr != "@" || domain != "container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "a.b.container-service.top"
	rr, domain = p.splitDNSName(endpoint.DNSName, hostedZoneDomains)
	if rr != "a.b" || domain != "container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "a.b.c.container-service.top"
	rr, domain = p.splitDNSName(endpoint.DNSName, hostedZoneDomains)
	if rr != "a.b.c" || domain != "container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	endpoint.DNSName = "a.b.c.container-service.top"
	rr, domain = p.splitDNSName(endpoint.DNSName, []string{"c.container-service.top"})
	if rr != "a.b" || domain != "c.container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}

	endpoint.DNSName = "a.b.c.container-service.top"
	rr, domain = p.splitDNSName(endpoint.DNSName, []string{"container-service.top", "c.container-service.top"})
	if rr != "a.b" || domain != "c.container-service.top" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	rr, domain = p.splitDNSName(endpoint.DNSName, emptyZoneDomains)
	if rr != "@" || domain != "" {
		t.Errorf("Failed to splitDNSName with emptyZoneDomains for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
	rr, domain = p.splitDNSName(endpoint.DNSName, []string{"example.com"})
	if rr != "@" || domain != "" {
		t.Errorf("Failed to splitDNSName for %s: rr=%s, domain=%s", endpoint.DNSName, rr, domain)
	}
}

func TestAlibabaCloudProvider_TXTEndpoint(t *testing.T) {
	const recordValue = "heritage=external-dns,external-dns/owner=default"
	const endpointTarget = "\"heritage=external-dns,external-dns/owner=default\""

	newTarget := wrapWithQuotes("TXT", recordValue)
	assert.Equal(t, endpointTarget, newTarget, "Failed to wrapWithQuotes: %s", newTarget)

	newValue := unwrapQuotes("TXT", endpointTarget)
	assert.Equal(t, recordValue, newValue, "Failed to unwrapQuotes: %s", newValue)
}
