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

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/maxatome/go-testdeep/td"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/source/annotations"
)

// proxyEnabled and proxyDisabled are pointers to bool values used to set if a record should be proxied through Cloudflare.
var (
	proxyEnabled  *bool = testutils.ToPtr(true)
	proxyDisabled *bool = testutils.ToPtr(false)
)

type MockAction struct {
	Name             string
	ZoneId           string
	RecordId         string
	RecordData       cloudflare.DNSRecord
	RegionalHostname cloudflare.RegionalHostname
}

type mockCloudFlareClient struct {
	Zones                 map[string]string
	Records               map[string]map[string]cloudflare.DNSRecord
	Actions               []MockAction
	listZonesError        error
	zoneDetailsError      error
	listZonesContextError error
	dnsRecordsError       error
	customHostnames       map[string][]cloudflare.CustomHostname
	regionalHostnames     map[string][]cloudflare.RegionalHostname
}

var ExampleDomain = []cloudflare.DNSRecord{
	{
		ID:      "1234567890",
		Name:    "foobar.bar.com",
		Type:    endpoint.RecordTypeA,
		TTL:     120,
		Content: "1.2.3.4",
		Proxied: proxyDisabled,
		Comment: "valid comment",
	},
	{
		ID:      "2345678901",
		Name:    "foobar.bar.com",
		Type:    endpoint.RecordTypeA,
		TTL:     120,
		Content: "3.4.5.6",
		Proxied: proxyDisabled,
	},
	{
		ID:      "1231231233",
		Name:    "bar.foo.com",
		Type:    endpoint.RecordTypeA,
		TTL:     1,
		Content: "2.3.4.5",
		Proxied: proxyDisabled,
	},
}

func NewMockCloudFlareClient() *mockCloudFlareClient {
	return &mockCloudFlareClient{
		Zones: map[string]string{
			"001": "bar.com",
			"002": "foo.com",
		},
		Records: map[string]map[string]cloudflare.DNSRecord{
			"001": {},
			"002": {},
		},
		customHostnames:   map[string][]cloudflare.CustomHostname{},
		regionalHostnames: map[string][]cloudflare.RegionalHostname{},
	}
}

func NewMockCloudFlareClientWithRecords(records map[string][]cloudflare.DNSRecord) *mockCloudFlareClient {
	m := NewMockCloudFlareClient()

	for zoneID, zoneRecords := range records {
		if zone, ok := m.Records[zoneID]; ok {
			for _, record := range zoneRecords {
				zone[record.ID] = record
			}
		}
	}

	return m
}

func getDNSRecordFromRecordParams(rp any) cloudflare.DNSRecord {
	switch params := rp.(type) {
	case cloudflare.CreateDNSRecordParams:
		record := cloudflare.DNSRecord{
			ID:      params.ID,
			Name:    params.Name,
			TTL:     params.TTL,
			Proxied: params.Proxied,
			Type:    params.Type,
			Content: params.Content,
		}
		if params.Type == "MX" {
			record.Priority = params.Priority
		}
		return record
	case cloudflare.UpdateDNSRecordParams:
		record := cloudflare.DNSRecord{
			ID:      params.ID,
			Name:    params.Name,
			TTL:     params.TTL,
			Proxied: params.Proxied,
			Type:    params.Type,
			Content: params.Content,
		}
		if params.Type == "MX" {
			record.Priority = params.Priority
		}
		return record
	default:
		return cloudflare.DNSRecord{}
	}
}

func generateDNSRecordID(rrtype string, name string, content string) string {
	return fmt.Sprintf("%s-%s-%s", name, rrtype, content)
}

func (m *mockCloudFlareClient) CreateDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.CreateDNSRecordParams) (cloudflare.DNSRecord, error) {
	recordData := getDNSRecordFromRecordParams(rp)
	if recordData.ID == "" {
		recordData.ID = generateDNSRecordID(recordData.Type, recordData.Name, recordData.Content)
	}
	m.Actions = append(m.Actions, MockAction{
		Name:       "Create",
		ZoneId:     rc.Identifier,
		RecordId:   recordData.ID,
		RecordData: recordData,
	})
	if zone, ok := m.Records[rc.Identifier]; ok {
		zone[recordData.ID] = recordData
	}

	if recordData.Name == "newerror.bar.com" {
		return cloudflare.DNSRecord{}, fmt.Errorf("failed to create record")
	}
	return cloudflare.DNSRecord{}, nil
}

func (m *mockCloudFlareClient) ListDNSRecords(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.ListDNSRecordsParams) ([]cloudflare.DNSRecord, *cloudflare.ResultInfo, error) {
	if m.dnsRecordsError != nil {
		return nil, &cloudflare.ResultInfo{}, m.dnsRecordsError
	}
	result := []cloudflare.DNSRecord{}
	if zone, ok := m.Records[rc.Identifier]; ok {
		for _, record := range zone {
			if strings.HasPrefix(record.Name, "newerror-list-") {
				m.DeleteDNSRecord(ctx, rc, record.ID)
				return nil, &cloudflare.ResultInfo{}, errors.New("failed to list erroring DNS record")
			}
			result = append(result, record)
		}
	}

	if len(result) == 0 || rp.PerPage == 0 {
		return result, &cloudflare.ResultInfo{Page: 1, TotalPages: 1, Count: 0, Total: 0}, nil
	}

	// if not pagination options were passed in, return the result as is
	if rp.Page == 0 {
		return result, &cloudflare.ResultInfo{Page: 1, TotalPages: 1, Count: len(result), Total: len(result)}, nil
	}

	// otherwise, split the result into chunks of size rp.PerPage to simulate the pagination from the API
	chunks := [][]cloudflare.DNSRecord{}

	// to ensure consistency in the multiple calls to this function, sort the result slice
	sort.Slice(result, func(i, j int) bool { return strings.Compare(result[i].ID, result[j].ID) > 0 })
	for rp.PerPage < len(result) {
		result, chunks = result[rp.PerPage:], append(chunks, result[0:rp.PerPage])
	}
	chunks = append(chunks, result)

	// return the requested page
	partialResult := chunks[rp.Page-1]
	return partialResult, &cloudflare.ResultInfo{
		PerPage:    rp.PerPage,
		Page:       rp.Page,
		TotalPages: len(chunks),
		Count:      len(partialResult),
		Total:      len(result),
	}, nil
}

func (m *mockCloudFlareClient) UpdateDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.UpdateDNSRecordParams) error {
	recordData := getDNSRecordFromRecordParams(rp)
	m.Actions = append(m.Actions, MockAction{
		Name:       "Update",
		ZoneId:     rc.Identifier,
		RecordId:   rp.ID,
		RecordData: recordData,
	})
	if zone, ok := m.Records[rc.Identifier]; ok {
		if _, ok := zone[rp.ID]; ok {
			if strings.HasPrefix(recordData.Name, "newerror-update-") {
				return errors.New("failed to update erroring DNS record")
			}
			zone[rp.ID] = recordData
		}
	}
	return nil
}

func (m *mockCloudFlareClient) DeleteDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, recordID string) error {
	m.Actions = append(m.Actions, MockAction{
		Name:     "Delete",
		ZoneId:   rc.Identifier,
		RecordId: recordID,
	})
	if zone, ok := m.Records[rc.Identifier]; ok {
		if _, ok := zone[recordID]; ok {
			name := zone[recordID].Name
			delete(zone, recordID)
			if strings.HasPrefix(name, "newerror-delete-") {
				return errors.New("failed to delete erroring DNS record")
			}
			return nil
		}
	}
	return nil
}

func (m *mockCloudFlareClient) CustomHostnames(ctx context.Context, zoneID string, page int, filter cloudflare.CustomHostname) ([]cloudflare.CustomHostname, cloudflare.ResultInfo, error) {
	var err error = nil
	perPage := 50 // cloudflare-go v0 API hardcoded

	if strings.HasPrefix(zoneID, "newerror-") {
		return nil, cloudflare.ResultInfo{}, errors.New("failed to list custom hostnames")
	}
	if filter.Hostname != "" {
		err = errors.New("filters are not supported for custom hostnames mock test")
		return nil, cloudflare.ResultInfo{}, err
	}
	if page < 1 {
		err = errors.New("incorrect page value for custom hostnames list")
		return nil, cloudflare.ResultInfo{}, err
	}

	result := []cloudflare.CustomHostname{}
	if chs, ok := m.customHostnames[zoneID]; ok {
		for idx := (page - 1) * perPage; idx < min(len(chs), page*perPage); idx++ {
			ch := m.customHostnames[zoneID][idx]
			if strings.HasPrefix(ch.Hostname, "newerror-list-") {
				m.DeleteCustomHostname(ctx, zoneID, ch.ID)
				return nil, cloudflare.ResultInfo{}, errors.New("failed to list erroring custom hostname")
			}
			result = append(result, ch)
		}
		return result,
			cloudflare.ResultInfo{
				Page:       page,
				PerPage:    perPage,
				Count:      len(result),
				Total:      len(chs),
				TotalPages: len(chs)/page + 1,
			}, err
	} else {
		return result,
			cloudflare.ResultInfo{
				Page:       page,
				PerPage:    perPage,
				Count:      0,
				Total:      0,
				TotalPages: 0,
			}, err
	}
}

func (m *mockCloudFlareClient) CreateCustomHostname(ctx context.Context, zoneID string, ch cloudflare.CustomHostname) (*cloudflare.CustomHostnameResponse, error) {
	if ch.Hostname == "" || ch.CustomOriginServer == "" || ch.Hostname == "newerror-create.foo.fancybar.com" {
		return nil, fmt.Errorf("Invalid custom hostname or origin hostname")
	}
	if _, ok := m.customHostnames[zoneID]; !ok {
		m.customHostnames[zoneID] = []cloudflare.CustomHostname{}
	}
	var newCustomHostname cloudflare.CustomHostname = ch
	newCustomHostname.ID = fmt.Sprintf("ID-%s", ch.Hostname)
	m.customHostnames[zoneID] = append(m.customHostnames[zoneID], newCustomHostname)
	return &cloudflare.CustomHostnameResponse{}, nil
}

func (m *mockCloudFlareClient) DeleteCustomHostname(ctx context.Context, zoneID string, customHostnameID string) error {
	idx := 0
	if idx = getCustomHostnameIdxByID(m.customHostnames[zoneID], customHostnameID); idx < 0 {
		return fmt.Errorf("Invalid custom hostname ID to delete")
	}

	m.customHostnames[zoneID] = append(m.customHostnames[zoneID][:idx], m.customHostnames[zoneID][idx+1:]...)

	if customHostnameID == "ID-newerror-delete.foo.fancybar.com" {
		return fmt.Errorf("Invalid custom hostname to delete")
	}
	return nil
}

func (m *mockCloudFlareClient) ZoneIDByName(zoneName string) (string, error) {
	for id, name := range m.Zones {
		if name == zoneName {
			return id, nil
		}
	}

	return "", errors.New("Unknown zone: " + zoneName)
}

func (m *mockCloudFlareClient) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	if m.listZonesContextError != nil {
		return cloudflare.ZonesResponse{}, m.listZonesContextError
	}

	result := []cloudflare.Zone{}

	for zoneId, zoneName := range m.Zones {
		result = append(result, cloudflare.Zone{
			ID:   zoneId,
			Name: zoneName,
		})
	}

	return cloudflare.ZonesResponse{
		Result: result,
		ResultInfo: cloudflare.ResultInfo{
			Page:       1,
			TotalPages: 1,
		},
	}, nil
}

func (m *mockCloudFlareClient) ZoneDetails(ctx context.Context, zoneID string) (cloudflare.Zone, error) {
	if m.zoneDetailsError != nil {
		return cloudflare.Zone{}, m.zoneDetailsError
	}

	for id, zoneName := range m.Zones {
		if zoneID == id {
			return cloudflare.Zone{
				ID:   zoneID,
				Name: zoneName,
				Plan: cloudflare.ZonePlan{IsSubscribed: strings.HasSuffix(zoneName, "bar.com")},
			}, nil
		}
	}

	return cloudflare.Zone{}, errors.New("Unknown zoneID: " + zoneID)
}

func getCustomHostnameIdxByID(chs []cloudflare.CustomHostname, customHostnameID string) int {
	for idx, ch := range chs {
		if ch.ID == customHostnameID {
			return idx
		}
	}
	return -1
}

func AssertActions(t *testing.T, provider *CloudFlareProvider, endpoints []*endpoint.Endpoint, actions []MockAction, managedRecords []string, args ...interface{}) {
	t.Helper()

	var client *mockCloudFlareClient

	if provider.Client == nil {
		client = NewMockCloudFlareClient()
		provider.Client = client
	} else {
		client = provider.Client.(*mockCloudFlareClient)
	}

	ctx := context.Background()

	records, err := provider.Records(ctx)
	if err != nil {
		t.Fatalf("cannot fetch records, %s", err)
	}

	endpoints, err = provider.AdjustEndpoints(endpoints)
	assert.NoError(t, err)
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})
	plan := &plan.Plan{
		Current:        records,
		Desired:        endpoints,
		DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
		ManagedRecords: managedRecords,
	}

	changes := plan.Calculate().Changes

	// Records other than A, CNAME and NS are not supported by planner, just create them
	for _, endpoint := range endpoints {
		if !slices.Contains(managedRecords, endpoint.RecordType) {
			changes.Create = append(changes.Create, endpoint)
		}
	}

	err = provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Fatalf("cannot apply changes, %s", err)
	}

	td.Cmp(t, client.Actions, actions, args...)
}

func TestCloudflareA(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "A",
			DNSName:    "bar.com",
			Targets:    endpoint.Targets{"127.0.0.1", "127.0.0.2"},
		},
	}

	AssertActions(t, &CloudFlareProvider{}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.1"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.2"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.2"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.2",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestCloudflareCname(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "CNAME",
			DNSName:    "cname.bar.com",
			Targets:    endpoint.Targets{"google.com", "facebook.com"},
		},
	}

	AssertActions(t, &CloudFlareProvider{}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("CNAME", "cname.bar.com", "google.com"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("CNAME", "cname.bar.com", "google.com"),
				Type:    "CNAME",
				Name:    "cname.bar.com",
				Content: "google.com",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("CNAME", "cname.bar.com", "facebook.com"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("CNAME", "cname.bar.com", "facebook.com"),
				Type:    "CNAME",
				Name:    "cname.bar.com",
				Content: "facebook.com",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestCloudflareMx(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "MX",
			DNSName:    "mx.bar.com",
			Targets:    endpoint.Targets{"10 google.com", "20 facebook.com"},
		},
	}

	AssertActions(t, &CloudFlareProvider{}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("MX", "mx.bar.com", "google.com"),
			RecordData: cloudflare.DNSRecord{
				ID:       generateDNSRecordID("MX", "mx.bar.com", "google.com"),
				Type:     "MX",
				Name:     "mx.bar.com",
				Content:  "google.com",
				Priority: cloudflare.Uint16Ptr(10),
				TTL:      1,
				Proxied:  proxyDisabled,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("MX", "mx.bar.com", "facebook.com"),
			RecordData: cloudflare.DNSRecord{
				ID:       generateDNSRecordID("MX", "mx.bar.com", "facebook.com"),
				Type:     "MX",
				Name:     "mx.bar.com",
				Content:  "facebook.com",
				Priority: cloudflare.Uint16Ptr(20),
				TTL:      1,
				Proxied:  proxyDisabled,
			},
		},
	},
		[]string{endpoint.RecordTypeMX},
	)
}

func TestCloudflareTxt(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "TXT",
			DNSName:    "txt.bar.com",
			Targets:    endpoint.Targets{"v=spf1 include:_spf.google.com ~all"},
		},
	}

	AssertActions(t, &CloudFlareProvider{}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("TXT", "txt.bar.com", "v=spf1 include:_spf.google.com ~all"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("TXT", "txt.bar.com", "v=spf1 include:_spf.google.com ~all"),
				Type:    "TXT",
				Name:    "txt.bar.com",
				Content: "v=spf1 include:_spf.google.com ~all",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
	},
		[]string{endpoint.RecordTypeTXT},
	)
}

func TestCloudflareCustomTTL(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "A",
			DNSName:    "ttl.bar.com",
			Targets:    endpoint.Targets{"127.0.0.1"},
			RecordTTL:  120,
		},
	}

	AssertActions(t, &CloudFlareProvider{}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "ttl.bar.com", "127.0.0.1"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "ttl.bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "ttl.bar.com",
				Content: "127.0.0.1",
				TTL:     120,
				Proxied: proxyDisabled,
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestCloudflareProxiedDefault(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "A",
			DNSName:    "bar.com",
			Targets:    endpoint.Targets{"127.0.0.1"},
		},
	}

	AssertActions(t, &CloudFlareProvider{proxiedByDefault: true}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.1"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: proxyEnabled,
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestCloudflareProxiedOverrideTrue(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "A",
			DNSName:    "bar.com",
			Targets:    endpoint.Targets{"127.0.0.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
					Value: "true",
				},
			},
		},
	}

	AssertActions(t, &CloudFlareProvider{}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.1"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: proxyEnabled,
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestCloudflareProxiedOverrideFalse(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "A",
			DNSName:    "bar.com",
			Targets:    endpoint.Targets{"127.0.0.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
					Value: "false",
				},
			},
		},
	}

	AssertActions(t, &CloudFlareProvider{proxiedByDefault: true}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.1"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestCloudflareProxiedOverrideIllegal(t *testing.T) {
	endpoints := []*endpoint.Endpoint{
		{
			RecordType: "A",
			DNSName:    "bar.com",
			Targets:    endpoint.Targets{"127.0.0.1"},
			ProviderSpecific: endpoint.ProviderSpecific{
				endpoint.ProviderSpecificProperty{
					Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
					Value: "asfasdfa",
				},
			},
		},
	}

	AssertActions(t, &CloudFlareProvider{proxiedByDefault: true}, endpoints, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.1"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: proxyEnabled,
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestCloudflareSetProxied(t *testing.T) {
	var proxied *bool = proxyEnabled
	var notProxied *bool = proxyDisabled
	testCases := []struct {
		recordType string
		domain     string
		proxiable  *bool
	}{
		{"A", "bar.com", proxied},
		{"CNAME", "bar.com", proxied},
		{"TXT", "bar.com", notProxied},
		{"MX", "bar.com", notProxied},
		{"NS", "bar.com", notProxied},
		{"SPF", "bar.com", notProxied},
		{"SRV", "bar.com", notProxied},
		{"A", "*.bar.com", proxied},
		{"CNAME", "*.docs.bar.com", proxied},
	}

	for _, testCase := range testCases {
		var targets endpoint.Targets
		var content string
		var priority *uint16

		if testCase.recordType == "MX" {
			targets = endpoint.Targets{"10 mx.example.com"}
			content = "mx.example.com"
			priority = cloudflare.Uint16Ptr(10)
		} else {
			targets = endpoint.Targets{"127.0.0.1"}
			content = "127.0.0.1"
		}

		endpoints := []*endpoint.Endpoint{
			{
				RecordType: testCase.recordType,
				DNSName:    testCase.domain,
				Targets:    endpoint.Targets{targets[0]},
				ProviderSpecific: endpoint.ProviderSpecific{
					endpoint.ProviderSpecificProperty{
						Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
						Value: "true",
					},
				},
			},
		}
		expectedID := fmt.Sprintf("%s-%s-%s", testCase.domain, testCase.recordType, content)
		recordData := cloudflare.DNSRecord{
			ID:      expectedID,
			Type:    testCase.recordType,
			Name:    testCase.domain,
			Content: content,
			TTL:     1,
			Proxied: testCase.proxiable,
		}
		if testCase.recordType == "MX" {
			recordData.Priority = priority
		}
		AssertActions(t, &CloudFlareProvider{}, endpoints, []MockAction{
			{
				Name:       "Create",
				ZoneId:     "001",
				RecordId:   expectedID,
				RecordData: recordData,
			},
		}, []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME, endpoint.RecordTypeNS, endpoint.RecordTypeMX}, testCase.recordType+" record on "+testCase.domain)
	}
}

func TestCloudflareZones(t *testing.T) {
	provider := &CloudFlareProvider{
		Client:       NewMockCloudFlareClient(),
		domainFilter: endpoint.NewDomainFilter([]string{"bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}

	zones, err := provider.Zones(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, zones, 1)
	assert.Equal(t, "bar.com", zones[0].Name)
}

// test failures on zone lookup
func TestCloudflareZonesFailed(t *testing.T) {

	client := NewMockCloudFlareClient()
	client.zoneDetailsError = errors.New("zone lookup failed")

	provider := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{"001"}),
	}

	_, err := provider.Zones(context.Background())
	if err == nil {
		t.Errorf("should fail, %s", err)
	}
}

func TestCloudFlareZonesWithIDFilter(t *testing.T) {
	client := NewMockCloudFlareClient()
	client.listZonesError = errors.New("shouldn't need to list zones when ZoneIDFilter in use")
	provider := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"bar.com", "foo.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{"001"}),
	}

	zones, err := provider.Zones(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// foo.com should *not* be returned as it doesn't match ZoneID filter
	assert.Len(t, zones, 1)
	assert.Equal(t, "bar.com", zones[0].Name)
}

func TestCloudflareListZonesRateLimited(t *testing.T) {
	// Create a mock client that returns a rate limit error
	client := NewMockCloudFlareClient()
	client.listZonesContextError = &cloudflare.Error{
		StatusCode: 429,
		ErrorCodes: []int{10000},
		Type:       cloudflare.ErrorTypeRateLimit,
	}
	p := &CloudFlareProvider{Client: client}

	// Call the Zones function
	_, err := p.Zones(context.Background())

	// Assert that a soft error was returned
	if !errors.Is(err, provider.SoftError) {
		t.Error("expected a rate limit error")
	}
}

func TestCloudflareListZonesRateLimitedStringError(t *testing.T) {
	// Create a mock client that returns a rate limit error
	client := NewMockCloudFlareClient()
	client.listZonesContextError = errors.New("exceeded available rate limit retries")
	p := &CloudFlareProvider{Client: client}

	// Call the Zones function
	_, err := p.Zones(context.Background())

	// Assert that a soft error was returned
	assert.ErrorIs(t, err, provider.SoftError, "expected a rate limit error")
}

func TestCloudflareListZoneInternalErrors(t *testing.T) {
	// Create a mock client that returns a internal server error
	client := NewMockCloudFlareClient()
	client.listZonesContextError = &cloudflare.Error{
		StatusCode: 500,
		ErrorCodes: []int{20000},
		Type:       cloudflare.ErrorTypeService,
	}
	p := &CloudFlareProvider{Client: client}

	// Call the Zones function
	_, err := p.Zones(context.Background())

	// Assert that a soft error was returned
	t.Log(err)
	if !errors.Is(err, provider.SoftError) {
		t.Errorf("expected a internal error")
	}
}

func TestCloudflareRecords(t *testing.T) {
	client := NewMockCloudFlareClientWithRecords(map[string][]cloudflare.DNSRecord{
		"001": ExampleDomain,
	})

	// Set DNSRecordsPerPage to 1 test the pagination behaviour
	p := &CloudFlareProvider{
		Client:           client,
		DNSRecordsConfig: DNSRecordsConfig{PerPage: 1},
	}
	ctx := context.Background()

	records, err := p.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	assert.Len(t, records, 2)
	client.dnsRecordsError = errors.New("failed to list dns records")
	_, err = p.Records(ctx)
	if err == nil {
		t.Errorf("expected to fail")
	}
	client.dnsRecordsError = nil
	client.listZonesContextError = &cloudflare.Error{
		StatusCode: 429,
		ErrorCodes: []int{10000},
		Type:       cloudflare.ErrorTypeRateLimit,
	}
	_, err = p.Records(ctx)
	// Assert that a soft error was returned
	if !errors.Is(err, provider.SoftError) {
		t.Error("expected a rate limit error")
	}

	client.listZonesContextError = &cloudflare.Error{
		StatusCode: 500,
		ErrorCodes: []int{10000},
		Type:       cloudflare.ErrorTypeService,
	}
	_, err = p.Records(ctx)
	// Assert that a soft error was returned
	if !errors.Is(err, provider.SoftError) {
		t.Error("expected a internal server error")
	}

	client.listZonesContextError = errors.New("failed to list zones")
	_, err = p.Records(ctx)
	if err == nil {
		t.Errorf("expected to fail")
	}
}

func TestCloudflareProvider(t *testing.T) {
	var err error

	type EnvVar struct {
		Key   string
		Value string
	}

	tokenFile := "/tmp/cf_api_token"
	if err := os.WriteFile(tokenFile, []byte("abc123def"), 0o644); err != nil {
		t.Errorf("failed to write token file, %s", err)
	}

	testCases := []struct {
		Name        string
		Environment []EnvVar
		ShouldFail  bool
	}{
		{
			Name: "use_api_token",
			Environment: []EnvVar{
				{Key: "CF_API_TOKEN", Value: "abc123def"},
			},
			ShouldFail: false,
		},
		{
			Name: "use_api_token_file_contents",
			Environment: []EnvVar{
				{Key: "CF_API_TOKEN", Value: tokenFile},
			},
			ShouldFail: false,
		},
		{
			Name: "use_email_and_key",
			Environment: []EnvVar{
				{Key: "CF_API_KEY", Value: "xxxxxxxxxxxxxxxxx"},
				{Key: "CF_API_EMAIL", Value: "test@test.com"},
			},
			ShouldFail: false,
		},
		{
			Name:        "no_use_email_and_key",
			Environment: []EnvVar{},
			ShouldFail:  true,
		},
		{
			Name: "use_credentials_in_missing_file",
			Environment: []EnvVar{
				{Key: "CF_API_TOKEN", Value: "file://abc"},
			},
			ShouldFail: true,
		},
		{
			Name: "use_credentials_in_missing_file",
			Environment: []EnvVar{
				{Key: "CF_API_TOKEN", Value: "file:/tmp/cf_api_token"},
			},
			ShouldFail: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			for _, env := range tc.Environment {
				t.Setenv(env.Key, env.Value)
			}

			_, err = NewCloudFlareProvider(
				endpoint.NewDomainFilter([]string{"bar.com"}),
				provider.NewZoneIDFilter([]string{""}),
				false,
				true,
				RegionalServicesConfig{Enabled: false},
				CustomHostnamesConfig{Enabled: false},
				DNSRecordsConfig{PerPage: 5000, Comment: ""},
			)
			if err != nil && !tc.ShouldFail {
				t.Errorf("should not fail, %s", err)
			}
			if err == nil && tc.ShouldFail {
				t.Errorf("should fail, %s", err)
			}
		})

	}
}

func TestCloudflareApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client: client,
	}
	changes.Create = []*endpoint.Endpoint{{
		DNSName: "new.bar.com",
		Targets: endpoint.Targets{"target"},
	}, {
		DNSName: "new.ext-dns-test.unrelated.to",
		Targets: endpoint.Targets{"target"},
	}}
	changes.Delete = []*endpoint.Endpoint{{
		DNSName: "foobar.bar.com",
		Targets: endpoint.Targets{"target"},
	}}
	changes.UpdateOld = []*endpoint.Endpoint{{
		DNSName: "foobar.bar.com",
		Targets: endpoint.Targets{"target-old"},
	}}
	changes.UpdateNew = []*endpoint.Endpoint{{
		DNSName: "foobar.bar.com",
		Targets: endpoint.Targets{"target-new"},
	}}
	err := provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	td.Cmp(t, client.Actions, []MockAction{
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("", "new.bar.com", "target"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("", "new.bar.com", "target"),
				Name:    "new.bar.com",
				Content: "target",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("", "foobar.bar.com", "target-new"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("", "foobar.bar.com", "target-new"),
				Name:    "foobar.bar.com",
				Content: "target-new",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
	})

	// empty changes
	changes.Create = []*endpoint.Endpoint{}
	changes.Delete = []*endpoint.Endpoint{}
	changes.UpdateOld = []*endpoint.Endpoint{}
	changes.UpdateNew = []*endpoint.Endpoint{}

	err = provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestCloudflareDryRunApplyChanges(t *testing.T) {
	changes := &plan.Changes{}
	client := NewMockCloudFlareClient()

	provider := &CloudFlareProvider{
		Client: client,
		DryRun: true,
	}
	changes.Create = []*endpoint.Endpoint{{
		DNSName: "new.bar.com",
		Targets: endpoint.Targets{"target"},
	}}
	err := provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	ctx := context.Background()
	records, err := provider.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	assert.Empty(t, records, "should not have any records")
}

func TestCloudflareApplyChangesError(t *testing.T) {
	changes := &plan.Changes{}
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client: client,
	}
	changes.Create = []*endpoint.Endpoint{{
		DNSName: "newerror.bar.com",
		Targets: endpoint.Targets{"target"},
	}}
	err := provider.ApplyChanges(context.Background(), changes)
	if err == nil {
		t.Errorf("should fail, %s", err)
	}
}

func TestCloudflareGetRecordID(t *testing.T) {
	p := &CloudFlareProvider{}
	recordsMap := DNSRecordsMap{
		{Name: "foo.com", Type: endpoint.RecordTypeCNAME, Content: "foobar"}: {
			Name:    "foo.com",
			Type:    endpoint.RecordTypeCNAME,
			Content: "foobar",
			ID:      "1",
		},
		{Name: "bar.de", Type: endpoint.RecordTypeA}: {
			Name: "bar.de",
			Type: endpoint.RecordTypeA,
			ID:   "2",
		},
		{Name: "bar.de", Type: endpoint.RecordTypeA, Content: "1.2.3.4"}: {
			Name:    "bar.de",
			Type:    endpoint.RecordTypeA,
			Content: "1.2.3.4",
			ID:      "2",
		},
	}

	assert.Empty(t, p.getRecordID(recordsMap, cloudflare.DNSRecord{
		Name:    "foo.com",
		Type:    endpoint.RecordTypeA,
		Content: "foobar",
	}))

	assert.Empty(t, p.getRecordID(recordsMap, cloudflare.DNSRecord{
		Name:    "foo.com",
		Type:    endpoint.RecordTypeCNAME,
		Content: "fizfuz",
	}))

	assert.Equal(t, "1", p.getRecordID(recordsMap, cloudflare.DNSRecord{
		Name:    "foo.com",
		Type:    endpoint.RecordTypeCNAME,
		Content: "foobar",
	}))
	assert.Empty(t, p.getRecordID(recordsMap, cloudflare.DNSRecord{
		Name:    "bar.de",
		Type:    endpoint.RecordTypeA,
		Content: "2.3.4.5",
	}))
	assert.Equal(t, "2", p.getRecordID(recordsMap, cloudflare.DNSRecord{
		Name:    "bar.de",
		Type:    endpoint.RecordTypeA,
		Content: "1.2.3.4",
	}))
}

func TestCloudflareGroupByNameAndType(t *testing.T) {
	provider := &CloudFlareProvider{
		Client:       NewMockCloudFlareClient(),
		domainFilter: endpoint.NewDomainFilter([]string{"bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}
	testCases := []struct {
		Name              string
		Records           []cloudflare.DNSRecord
		ExpectedEndpoints []*endpoint.Endpoint
	}{
		{
			Name:              "empty",
			Records:           []cloudflare.DNSRecord{},
			ExpectedEndpoints: []*endpoint.Endpoint{},
		},
		{
			Name: "single record - single target",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
		{
			Name: "single record - multiple targets",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
		{
			Name: "multiple record - multiple targets",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
				{
					DNSName:    "bar.de",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
		{
			Name: "multiple record - mixed single/multiple targets",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
				{
					DNSName:    "bar.de",
					Targets:    endpoint.Targets{"10.10.10.1"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
		{
			Name: "unsupported record type",
			Records: []cloudflare.DNSRecord{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "bar.de",
					Type:    "NOT SUPPORTED",
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
							Value: "false",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		records := make(DNSRecordsMap)
		for _, r := range tc.Records {
			records[newDNSRecordIndex(r)] = r
		}
		endpoints := provider.groupByNameAndTypeWithCustomHostnames(records, CustomHostnamesMap{})
		// Targets order could be random with underlying map
		for _, ep := range endpoints {
			slices.Sort(ep.Targets)
		}
		for _, ep := range tc.ExpectedEndpoints {
			slices.Sort(ep.Targets)
		}
		assert.ElementsMatch(t, endpoints, tc.ExpectedEndpoints)
	}
}

func TestGroupByNameAndTypeWithCustomHostnames_MX(t *testing.T) {
	client := NewMockCloudFlareClientWithRecords(map[string][]cloudflare.DNSRecord{
		"001": {
			{
				ID:       "mx-1",
				Name:     "mx.bar.com",
				Type:     endpoint.RecordTypeMX,
				TTL:      3600,
				Content:  "mail.bar.com",
				Priority: cloudflare.Uint16Ptr(10),
			},
			{
				ID:       "mx-2",
				Name:     "mx.bar.com",
				Type:     endpoint.RecordTypeMX,
				TTL:      3600,
				Content:  "mail2.bar.com",
				Priority: cloudflare.Uint16Ptr(20),
			},
		},
	})
	provider := &CloudFlareProvider{
		Client: client,
	}
	ctx := context.Background()
	chs := CustomHostnamesMap{}
	records, err := provider.listDNSRecordsWithAutoPagination(ctx, "001")
	assert.NoError(t, err)

	endpoints := provider.groupByNameAndTypeWithCustomHostnames(records, chs)
	assert.Len(t, endpoints, 1)
	mxEndpoint := endpoints[0]
	assert.Equal(t, "mx.bar.com", mxEndpoint.DNSName)
	assert.Equal(t, endpoint.RecordTypeMX, mxEndpoint.RecordType)
	assert.ElementsMatch(t, []string{"10 mail.bar.com", "20 mail2.bar.com"}, mxEndpoint.Targets)
	assert.Equal(t, endpoint.TTL(3600), mxEndpoint.RecordTTL)
}

func TestProviderPropertiesIdempotency(t *testing.T) {
	testCases := []struct {
		Name                     string
		ProviderProxiedByDefault bool
		RecordsAreProxied        *bool
		ShouldBeUpdated          bool
	}{
		{
			Name:                     "ProxyDefault: false, ShouldBeProxied: false, ExpectUpdates: false",
			ProviderProxiedByDefault: false,
			RecordsAreProxied:        proxyDisabled,
			ShouldBeUpdated:          false,
		},
		{
			Name:                     "ProxyDefault: true, ShouldBeProxied: true, ExpectUpdates: false",
			ProviderProxiedByDefault: true,
			RecordsAreProxied:        proxyEnabled,
			ShouldBeUpdated:          false,
		},
		{
			Name:                     "ProxyDefault: true, ShouldBeProxied: false, ExpectUpdates: true",
			ProviderProxiedByDefault: true,
			RecordsAreProxied:        proxyDisabled,
			ShouldBeUpdated:          true,
		},
		{
			Name:                     "ProxyDefault: false, ShouldBeProxied: true, ExpectUpdates: true",
			ProviderProxiedByDefault: false,
			RecordsAreProxied:        proxyEnabled,
			ShouldBeUpdated:          true,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			client := NewMockCloudFlareClientWithRecords(map[string][]cloudflare.DNSRecord{
				"001": {
					{
						ID:      "1234567890",
						Name:    "foobar.bar.com",
						Type:    endpoint.RecordTypeA,
						TTL:     120,
						Content: "1.2.3.4",
						Proxied: test.RecordsAreProxied,
					},
				},
			})

			provider := &CloudFlareProvider{
				Client:           client,
				proxiedByDefault: test.ProviderProxiedByDefault,
			}
			ctx := context.Background()

			current, err := provider.Records(ctx)
			if err != nil {
				t.Errorf("should not fail, %s", err)
			}
			assert.Len(t, current, 1)

			desired := []*endpoint.Endpoint{}
			for _, c := range current {
				// Copy all except ProviderSpecific fields
				desired = append(desired, &endpoint.Endpoint{
					DNSName:       c.DNSName,
					Targets:       c.Targets,
					RecordType:    c.RecordType,
					SetIdentifier: c.SetIdentifier,
					RecordTTL:     c.RecordTTL,
					Labels:        c.Labels,
				})
			}

			desired, err = provider.AdjustEndpoints(desired)
			assert.NoError(t, err)

			plan := plan.Plan{
				Current:        current,
				Desired:        desired,
				ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
			}

			plan = *plan.Calculate()
			assert.NotNil(t, plan.Changes, "should have plan")
			if plan.Changes == nil {
				return
			}
			assert.Empty(t, plan.Changes.Create, "should not have creates")
			assert.Empty(t, plan.Changes.Delete, "should not have deletes")

			if test.ShouldBeUpdated {
				assert.Len(t, plan.Changes.UpdateNew, 1, "should not have new updates")
				assert.Len(t, plan.Changes.UpdateOld, 1, "should not have old updates")
			} else {
				assert.Empty(t, plan.Changes.UpdateNew, "should not have new updates")
				assert.Empty(t, plan.Changes.UpdateOld, "should not have old updates")
			}
		})
	}
}

func TestCloudflareComplexUpdate(t *testing.T) {
	client := NewMockCloudFlareClientWithRecords(map[string][]cloudflare.DNSRecord{
		"001": ExampleDomain,
	})

	provider := &CloudFlareProvider{
		Client: client,
	}
	ctx := context.Background()

	records, err := provider.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})
	endpoints, err := provider.AdjustEndpoints([]*endpoint.Endpoint{
		{
			DNSName:    "foobar.bar.com",
			Targets:    endpoint.Targets{"1.2.3.4", "2.3.4.5"},
			RecordType: endpoint.RecordTypeA,
			RecordTTL:  endpoint.TTL(defaultTTL),
			Labels:     endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
					Value: "true",
				},
			},
		},
	})
	assert.NoError(t, err)
	plan := &plan.Plan{
		Current:        records,
		Desired:        endpoints,
		DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	}

	planned := plan.Calculate()

	err = provider.ApplyChanges(context.Background(), planned.Changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	td.CmpDeeply(t, client.Actions, []MockAction{
		{
			Name:     "Delete",
			ZoneId:   "001",
			RecordId: "2345678901",
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "foobar.bar.com", "2.3.4.5"),
			RecordData: cloudflare.DNSRecord{
				ID:      generateDNSRecordID("A", "foobar.bar.com", "2.3.4.5"),
				Name:    "foobar.bar.com",
				Type:    "A",
				Content: "2.3.4.5",
				TTL:     1,
				Proxied: proxyEnabled,
			},
		},
		{
			Name:     "Update",
			ZoneId:   "001",
			RecordId: "1234567890",
			RecordData: cloudflare.DNSRecord{
				ID:      "1234567890",
				Name:    "foobar.bar.com",
				Type:    "A",
				Content: "1.2.3.4",
				TTL:     1,
				Proxied: proxyEnabled,
			},
		},
	})
}

func TestCustomTTLWithEnabledProxyNotChanged(t *testing.T) {
	client := NewMockCloudFlareClientWithRecords(map[string][]cloudflare.DNSRecord{
		"001": {
			{
				ID:      "1234567890",
				Name:    "foobar.bar.com",
				Type:    endpoint.RecordTypeA,
				TTL:     1,
				Content: "1.2.3.4",
				Proxied: proxyEnabled,
			},
		},
	})

	provider := &CloudFlareProvider{
		Client: client,
	}

	records, err := provider.Records(context.Background())
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	endpoints := []*endpoint.Endpoint{
		{
			DNSName:    "foobar.bar.com",
			Targets:    endpoint.Targets{"1.2.3.4"},
			RecordType: endpoint.RecordTypeA,
			RecordTTL:  300,
			Labels:     endpoint.Labels{},
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  "external-dns.alpha.kubernetes.io/cloudflare-proxied",
					Value: "true",
				},
			},
		},
	}

	provider.AdjustEndpoints(endpoints)

	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})
	plan := &plan.Plan{
		Current:        records,
		Desired:        endpoints,
		DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	}

	planned := plan.Calculate()

	assert.Empty(t, planned.Changes.Create, "no new changes should be here")
	assert.Empty(t, planned.Changes.UpdateNew, "no new changes should be here")
	assert.Empty(t, planned.Changes.UpdateOld, "no new changes should be here")
	assert.Empty(t, planned.Changes.Delete, "no new changes should be here")
}

func TestCloudFlareProvider_Region(t *testing.T) {
	t.Setenv("CF_API_TOKEN", "abc123def")
	t.Setenv("CF_API_EMAIL", "test@test.com")
	provider, err := NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"example.com"}),
		provider.ZoneIDFilter{},
		true,
		false,
		RegionalServicesConfig{Enabled: false, RegionKey: "us"},
		CustomHostnamesConfig{Enabled: false},
		DNSRecordsConfig{PerPage: 50, Comment: ""},
	)
	assert.NoError(t, err, "should not fail to create provider")
	assert.True(t, provider.RegionalServicesConfig.Enabled, "expect regional services to be enabled")
	assert.Equal(t, "us", provider.RegionalServicesConfig.RegionKey, "expected region key to be 'us'")
}

func TestCloudFlareProvider_newCloudFlareChange(t *testing.T) {
	_ = os.Setenv("CF_API_KEY", "xxxxxxxxxxxxxxxxx")
	_ = os.Setenv("CF_API_EMAIL", "test@test.com")

	p, err := NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"example.com"}),
		provider.ZoneIDFilter{},
		true,
		false,
		RegionalServicesConfig{Enabled: true, RegionKey: "us"},
		CustomHostnamesConfig{Enabled: false},
		DNSRecordsConfig{PerPage: 50},
	)
	if err != nil {
		t.Fatal(err)
	}

	ep := &endpoint.Endpoint{
		DNSName:    "example.com",
		RecordType: "A",
		Targets:    []string{"192.0.2.1"},
	}

	change, _ := p.newCloudFlareChange(cloudFlareCreate, ep, ep.Targets[0], nil)
	if change.RegionalHostname.RegionKey != "us" {
		t.Errorf("expected region key to be 'us', but got '%s'", change.RegionalHostname.RegionKey)
	}

	var freeValidCommentBuilder strings.Builder
	for range freeZoneMaxCommentLength {
		freeValidCommentBuilder.WriteString("x")
	}

	var freeInvalidCommentBuilder strings.Builder
	for range freeZoneMaxCommentLength + 1 {
		freeInvalidCommentBuilder.WriteString("x")
	}

	var paidValidCommentBuilder strings.Builder
	for range paidZoneMaxCommentLength {
		paidValidCommentBuilder.WriteString("x")
	}
	var paidInvalidCommentBuilder strings.Builder
	for range paidZoneMaxCommentLength + 1 {
		paidInvalidCommentBuilder.WriteString("x")
	}

	paidProvider, err := NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"bar.com"}),
		provider.ZoneIDFilter{},
		true,
		false,
		RegionalServicesConfig{Enabled: true, RegionKey: "us"},
		CustomHostnamesConfig{Enabled: false},
		DNSRecordsConfig{PerPage: 50, Comment: paidValidCommentBuilder.String()},
	)
	if err != nil {
		t.Fatal(err)
	}

	paidProvider.Client = NewMockCloudFlareClient()
	commentTestCases := []struct {
		name     string
		provider *CloudFlareProvider
		endpoint *endpoint.Endpoint
		expected int
	}{
		{
			name:     "For free Zone respecting comment length, expect no trimming",
			provider: p,
			endpoint: &endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "A",
				Targets:    []string{"192.0.2.1"},
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  annotations.CloudflareRecordCommentKey,
						Value: freeValidCommentBuilder.String(),
					},
				},
			},
			expected: len(freeValidCommentBuilder.String()),
		},
		{
			name:     "For free Zones not respecting comment length, expect trimmed comments",
			provider: p,
			endpoint: &endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "A",
				Targets:    []string{"192.0.2.1"},
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  annotations.CloudflareRecordCommentKey,
						Value: freeInvalidCommentBuilder.String(),
					},
				},
			},
			expected: freeZoneMaxCommentLength,
		},
		{
			name:     "For paid Zones respecting comment length, expect no trimming",
			provider: paidProvider,
			endpoint: &endpoint.Endpoint{
				DNSName:    "bar.com",
				RecordType: "A",
				Targets:    []string{"192.0.2.1"},
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  annotations.CloudflareRecordCommentKey,
						Value: paidValidCommentBuilder.String(),
					},
				},
			},
			expected: len(paidValidCommentBuilder.String()),
		},
		{
			name:     "For paid Zones not respecting comment length, expect trimmed comments",
			provider: paidProvider,
			endpoint: &endpoint.Endpoint{
				DNSName:    "bar.com",
				RecordType: "A",
				Targets:    []string{"192.0.2.1"},
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  annotations.CloudflareRecordCommentKey,
						Value: paidInvalidCommentBuilder.String(),
					},
				},
			},
			expected: paidZoneMaxCommentLength,
		},
	}

	for _, test := range commentTestCases {
		t.Run(test.name, func(t *testing.T) {
			change, err := test.provider.newCloudFlareChange(cloudFlareCreate, test.endpoint, test.endpoint.Targets[0], nil)
			assert.NoError(t, err)
			if len(change.ResourceRecord.Comment) != test.expected {
				t.Errorf("expected comment to be %d characters long, but got %d", test.expected, len(change.ResourceRecord.Comment))
			}
		})
	}
}

func TestCloudFlareProvider_submitChangesCNAME(t *testing.T) {
	client := NewMockCloudFlareClientWithRecords(map[string][]cloudflare.DNSRecord{
		"001": {
			{
				ID:      "1234567890",
				Name:    "my-domain-here.app",
				Type:    endpoint.RecordTypeCNAME,
				TTL:     1,
				Content: "my-tunnel-guid-here.cfargotunnel.com",
				Proxied: proxyEnabled,
			},
			{
				ID:      "9876543210",
				Name:    "my-domain-here.app",
				Type:    endpoint.RecordTypeTXT,
				TTL:     1,
				Content: "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/external-dns/my-domain-here-app",
			},
		},
	})
	// zoneIdFilter := provider.NewZoneIDFilter([]string{"001"})
	provider := &CloudFlareProvider{
		Client: client,
	}

	changes := []*cloudFlareChange{
		{
			Action: cloudFlareUpdate,
			ResourceRecord: cloudflare.DNSRecord{
				Name:    "my-domain-here.app",
				Type:    endpoint.RecordTypeCNAME,
				ID:      "1234567890",
				Content: "my-tunnel-guid-here.cfargotunnel.com",
			},
			RegionalHostname: cloudflare.RegionalHostname{
				Hostname: "my-domain-here.app",
			},
		},
		{
			Action: cloudFlareUpdate,
			ResourceRecord: cloudflare.DNSRecord{
				Name:    "my-domain-here.app",
				Type:    endpoint.RecordTypeTXT,
				ID:      "9876543210",
				Content: "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/external-dns/my-domain-here-app",
			},
			RegionalHostname: cloudflare.RegionalHostname{
				Hostname:  "my-domain-here.app",
				RegionKey: "",
			},
		},
	}

	// Should not return an error
	err := provider.submitChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestCloudFlareProvider_submitChangesApex(t *testing.T) {
	// Create a mock CloudFlare client with APEX records
	client := NewMockCloudFlareClientWithRecords(map[string][]cloudflare.DNSRecord{
		"001": {
			{
				ID:      "1234567890",
				Name:    "@", // APEX record
				Type:    endpoint.RecordTypeCNAME,
				TTL:     1,
				Content: "my-tunnel-guid-here.cfargotunnel.com",
				Proxied: proxyEnabled,
			},
			{
				ID:      "9876543210",
				Name:    "@", // APEX record
				Type:    endpoint.RecordTypeTXT,
				TTL:     1,
				Content: "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/external-dns/my-domain-here-app",
			},
		},
	})

	// Create a CloudFlare provider instance
	provider := &CloudFlareProvider{
		Client: client,
	}

	// Define changes to submit
	changes := []*cloudFlareChange{
		{
			Action: cloudFlareUpdate,
			ResourceRecord: cloudflare.DNSRecord{
				Name:    "@", // APEX record
				Type:    endpoint.RecordTypeCNAME,
				ID:      "1234567890",
				Content: "my-tunnel-guid-here.cfargotunnel.com",
			},
			RegionalHostname: cloudflare.RegionalHostname{
				Hostname: "@", // APEX record
			},
		},
		{
			Action: cloudFlareUpdate,
			ResourceRecord: cloudflare.DNSRecord{
				Name:    "@", // APEX record
				Type:    endpoint.RecordTypeTXT,
				ID:      "9876543210",
				Content: "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/external-dns/my-domain-here-app",
			},
			RegionalHostname: cloudflare.RegionalHostname{
				Hostname:  "@", // APEX record
				RegionKey: "",
			},
		},
	}

	// Submit changes and verify no error is returned
	err := provider.submitChanges(context.Background(), changes)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
}

func TestCloudflareZoneRecordsFail(t *testing.T) {
	client := &mockCloudFlareClient{
		Zones: map[string]string{
			"newerror-001": "bar.com",
		},
		Records:         map[string]map[string]cloudflare.DNSRecord{},
		customHostnames: map[string][]cloudflare.CustomHostname{},
	}
	failingProvider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: true},
	}
	ctx := context.Background()

	_, err := failingProvider.Records(ctx)
	if err == nil {
		t.Errorf("should fail - invalid zone id, %s", err)
	}
}

// TestCloudflareLongRecordsErrorLog checks if the error is logged when a record name exceeds 63 characters
// it's not likely to happen in practice, as the Cloudflare API should reject having it
func TestCloudflareLongRecordsErrorLog(t *testing.T) {
	client := NewMockCloudFlareClientWithRecords(map[string][]cloudflare.DNSRecord{
		"001": {
			{
				ID:      "1234567890",
				Name:    "very-very-very-very-very-very-very-long-name-more-than-63-bytes-long.bar.com",
				Type:    endpoint.RecordTypeTXT,
				TTL:     120,
				Content: "some-content",
			},
		},
	})
	hook := testutils.LogsUnderTestWithLogLevel(log.InfoLevel, t)
	p := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: true},
	}
	ctx := context.Background()
	_, err := p.Records(ctx)
	if err != nil {
		t.Errorf("should not fail - too long record, %s", err)
	}
	testutils.TestHelperLogContains("s longer than 63 characters. Cannot create endpoint", hook, t)
}

// check if the error is expected
func checkFailed(name string, err error, shouldFail bool) error {
	if errors.Is(err, nil) && shouldFail {
		return fmt.Errorf("should fail - %q", name)
	}
	if !errors.Is(err, nil) && !shouldFail {
		return fmt.Errorf("should not fail - %q, %w", name, err)
	}
	return nil
}

func TestCloudflareDNSRecordsOperationsFail(t *testing.T) {
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: true},
	}
	ctx := context.Background()
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	testFailCases := []struct {
		Name                    string
		Endpoints               []*endpoint.Endpoint
		ExpectedCustomHostnames map[string]string
		shouldFail              bool
	}{
		{
			Name: "failing to create dns record",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "newerror.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
				},
			},
			shouldFail: true,
		},
		{
			Name: "adding failing to list DNS record",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "newerror-list-1.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
				},
			},
			shouldFail: false,
		},
		{
			Name:       "causing to list failing to list DNS record",
			Endpoints:  []*endpoint.Endpoint{},
			shouldFail: true,
		},
		{
			Name: "create failing to update DNS record",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "newerror-update-1.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
				},
			},
			shouldFail: false,
		},
		{
			Name: "failing to update DNS record",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "newerror-update-1.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  1234,
					Labels:     endpoint.Labels{},
				},
			},
			shouldFail: true,
		},
		{
			Name: "create failing to delete DNS record",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "newerror-delete-1.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  1234,
					Labels:     endpoint.Labels{},
				},
			},
			shouldFail: false,
		},
		{
			Name:       "failing to delete erroring DNS record",
			Endpoints:  []*endpoint.Endpoint{},
			shouldFail: true,
		},
	}

	for _, tc := range testFailCases {
		var err error
		var records, endpoints []*endpoint.Endpoint

		records, err = provider.Records(ctx)
		if errors.Is(err, nil) {
			endpoints, err = provider.AdjustEndpoints(tc.Endpoints)
		}
		if errors.Is(err, nil) {
			plan := &plan.Plan{
				Current:        records,
				Desired:        endpoints,
				DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
				ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
			}
			planned := plan.Calculate()
			err = provider.ApplyChanges(context.Background(), planned.Changes)
		}
		if e := checkFailed(tc.Name, err, tc.shouldFail); !errors.Is(e, nil) {
			t.Error(e)
		}
	}
}

func TestCloudflareCustomHostnameOperations(t *testing.T) {
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: true},
	}
	ctx := context.Background()
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	testFailCases := []struct {
		Name                    string
		Endpoints               []*endpoint.Endpoint
		ExpectedCustomHostnames map[string]string
	}{}

	for _, tc := range testFailCases {
		records, err := provider.Records(ctx)
		if err != nil {
			t.Errorf("should not fail, %v", err)
		}

		endpoints, err := provider.AdjustEndpoints(tc.Endpoints)

		assert.NoError(t, err)
		plan := &plan.Plan{
			Current:        records,
			Desired:        endpoints,
			DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
			ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
		}

		planned := plan.Calculate()

		err = provider.ApplyChanges(context.Background(), planned.Changes)
		if e := checkFailed(tc.Name, err, false); !errors.Is(e, nil) {
			t.Error(e)
		}

		chs, chErr := provider.listCustomHostnamesWithPagination(ctx, "001")
		if e := checkFailed(tc.Name, chErr, false); !errors.Is(e, nil) {
			t.Error(e)
		}

		actualCustomHostnames := map[string]string{}
		for _, ch := range chs {
			actualCustomHostnames[ch.Hostname] = ch.CustomOriginServer
		}
		if len(actualCustomHostnames) == 0 {
			actualCustomHostnames = nil
		}
		assert.Equal(t, tc.ExpectedCustomHostnames, actualCustomHostnames, "custom hostnames should be the same")
	}
}

func TestCloudflareDisabledCustomHostnameOperations(t *testing.T) {
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: false},
	}
	ctx := context.Background()
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	testCases := []struct {
		Name        string
		Endpoints   []*endpoint.Endpoint
		testChanges bool
	}{
		{
			Name: "add custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "a.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.11"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "a.foo.fancybar.com",
						},
					},
				},
				{
					DNSName:    "b.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.12"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
				},
				{
					DNSName:    "c.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.13"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "c1.foo.fancybar.com",
						},
					},
				},
			},
			testChanges: false,
		},
		{
			Name: "add custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "a.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.11"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
				},
				{
					DNSName:    "b.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.12"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "b.foo.fancybar.com",
						},
					},
				},
				{
					DNSName:    "c.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.13"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "c2.foo.fancybar.com",
						},
					},
				},
			},
			testChanges: true,
		},
	}

	for _, tc := range testCases {
		records, err := provider.Records(ctx)
		if err != nil {
			t.Errorf("should not fail, %v", err)
		}

		endpoints, err := provider.AdjustEndpoints(tc.Endpoints)

		assert.NoError(t, err)
		plan := &plan.Plan{
			Current:        records,
			Desired:        endpoints,
			DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
			ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
		}
		planned := plan.Calculate()
		err = provider.ApplyChanges(context.Background(), planned.Changes)
		if e := checkFailed(tc.Name, err, false); !errors.Is(e, nil) {
			t.Error(e)
		}
		if tc.testChanges {
			assert.False(t, planned.Changes.HasChanges(), "no new changes should be here")
		}
	}
}

func TestCloudflareCustomHostnameNotFoundOnRecordDeletion(t *testing.T) {
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: true},
	}
	ctx := context.Background()
	zoneID := "001"
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	testCases := []struct {
		Name                    string
		Endpoints               []*endpoint.Endpoint
		ExpectedCustomHostnames map[string]string
		preApplyHook            string
		logOutput               string
	}{
		{
			Name: "create DNS record with custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "create.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "newerror-getCustomHostnameOrigin.foo.fancybar.com",
						},
					},
				},
			},
			preApplyHook: "",
			logOutput:    "",
		},
		{
			Name:         "remove DNS record with unexpectedly missing custom hostname",
			Endpoints:    []*endpoint.Endpoint{},
			preApplyHook: "corrupt",
			logOutput:    "failed to delete custom hostname \"newerror-getCustomHostnameOrigin.foo.fancybar.com\": failed to get custom hostname: \"newerror-getCustomHostnameOrigin.foo.fancybar.com\" not found",
		},
		{
			Name:         "duplicate custom hostname",
			Endpoints:    []*endpoint.Endpoint{},
			preApplyHook: "duplicate",
			logOutput:    "",
		},
		{
			Name: "create DNS record with custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "a.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "a.foo.fancybar.com",
						},
					},
				},
			},
			preApplyHook: "",
			logOutput:    "custom hostname \"a.foo.fancybar.com\" already exists with the same origin \"a.foo.bar.com\", continue",
		},
	}

	for _, tc := range testCases {
		hook := testutils.LogsUnderTestWithLogLevel(log.InfoLevel, t)

		records, err := provider.Records(ctx)
		if err != nil {
			t.Errorf("should not fail, %v", err)
		}

		endpoints, err := provider.AdjustEndpoints(tc.Endpoints)

		assert.NoError(t, err)
		plan := &plan.Plan{
			Current:        records,
			Desired:        endpoints,
			DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
			ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
		}

		planned := plan.Calculate()

		// manually corrupt custom hostname before the deletion step
		// the purpose is to cause getCustomHostnameOrigin() to fail on change.Action == cloudFlareDelete
		chs, chErr := provider.listCustomHostnamesWithPagination(ctx, zoneID)
		if e := checkFailed(tc.Name, chErr, false); !errors.Is(e, nil) {
			t.Error(e)
		}
		if tc.preApplyHook == "corrupt" {
			if ch, err := getCustomHostname(chs, "newerror-getCustomHostnameOrigin.foo.fancybar.com"); errors.Is(err, nil) {
				chID := ch.ID
				t.Logf("corrupting custom hostname %q", chID)
				oldIdx := getCustomHostnameIdxByID(client.customHostnames[zoneID], chID)
				oldCh := client.customHostnames[zoneID][oldIdx]
				ch := cloudflare.CustomHostname{
					Hostname:           "corrupted-newerror-getCustomHostnameOrigin.foo.fancybar.com",
					CustomOriginServer: oldCh.CustomOriginServer,
					SSL:                oldCh.SSL,
				}
				client.customHostnames[zoneID][oldIdx] = ch
			}
		} else if tc.preApplyHook == "duplicate" { // manually inject duplicating custom hostname with the same name and origin
			ch := cloudflare.CustomHostname{
				ID:                 "ID-random-123",
				Hostname:           "a.foo.fancybar.com",
				CustomOriginServer: "a.foo.bar.com",
			}
			client.customHostnames[zoneID] = append(client.customHostnames[zoneID], ch)
		}
		err = provider.ApplyChanges(context.Background(), planned.Changes)
		if e := checkFailed(tc.Name, err, false); !errors.Is(e, nil) {
			t.Error(e)
		}

		testutils.TestHelperLogContains(tc.logOutput, hook, t)
	}
}

func TestCloudflareListCustomHostnamesWithPagionation(t *testing.T) {
	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client:                client,
		CustomHostnamesConfig: CustomHostnamesConfig{Enabled: true},
	}
	ctx := context.Background()
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	const CustomHostnamesNumber = 342
	var generatedEndpoints []*endpoint.Endpoint
	for i := 0; i < CustomHostnamesNumber; i++ {
		ep := []*endpoint.Endpoint{
			{
				DNSName:    fmt.Sprintf("host-%d.foo.bar.com", i),
				Targets:    endpoint.Targets{fmt.Sprintf("cname-%d.foo.bar.com", i)},
				RecordType: endpoint.RecordTypeCNAME,
				RecordTTL:  endpoint.TTL(defaultTTL),
				Labels:     endpoint.Labels{},
				ProviderSpecific: endpoint.ProviderSpecific{
					{
						Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
						Value: fmt.Sprintf("host-%d.foo.fancybar.com", i),
					},
				},
			},
		}
		generatedEndpoints = append(generatedEndpoints, ep...)
	}

	records, err := provider.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %v", err)
	}

	endpoints, err := provider.AdjustEndpoints(generatedEndpoints)

	assert.NoError(t, err)
	plan := &plan.Plan{
		Current:        records,
		Desired:        endpoints,
		DomainFilter:   endpoint.MatchAllDomainFilters{domainFilter},
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	}

	planned := plan.Calculate()

	err = provider.ApplyChanges(context.Background(), planned.Changes)
	if err != nil {
		t.Errorf("should not fail - %v", err)
	}

	chs, chErr := provider.listCustomHostnamesWithPagination(ctx, "001")
	if chErr != nil {
		t.Errorf("should not fail - %v", chErr)
	}
	assert.Len(t, chs, CustomHostnamesNumber)
}

func TestZoneHasPaidPlan(t *testing.T) {
	client := NewMockCloudFlareClient()
	cfprovider := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com", "bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}

	assert.False(t, cfprovider.ZoneHasPaidPlan("subdomain.foo.com"))
	assert.True(t, cfprovider.ZoneHasPaidPlan("subdomain.bar.com"))
	assert.False(t, cfprovider.ZoneHasPaidPlan("invaliddomain"))

	client.zoneDetailsError = errors.New("zone lookup failed")
	cfproviderWithZoneError := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com", "bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}
	assert.False(t, cfproviderWithZoneError.ZoneHasPaidPlan("subdomain.foo.com"))
}
func TestCloudflareApplyChanges_AllErrorLogPaths(t *testing.T) {
	hook := testutils.LogsUnderTestWithLogLevel(log.ErrorLevel, t)

	client := NewMockCloudFlareClient()
	provider := &CloudFlareProvider{
		Client: client,
	}

	cases := []struct {
		name                   string
		changes                *plan.Changes
		customHostnamesEnabled bool
		errorLogCount          int
	}{
		{
			name: "Create error (custom hostnames enabled)",
			changes: &plan.Changes{
				Create: []*endpoint.Endpoint{{
					DNSName:    "bad-create.bar.com",
					RecordType: "MX",
					Targets:    endpoint.Targets{"not-a-valid-mx"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "bad-create-custom.bar.com",
						},
					},
				}},
			},
			customHostnamesEnabled: true,
			errorLogCount:          1,
		},
		{
			name: "Delete error (custom hostnames enabled)",
			changes: &plan.Changes{
				Delete: []*endpoint.Endpoint{{
					DNSName:    "bad-delete.bar.com",
					RecordType: "MX",
					Targets:    endpoint.Targets{"not-a-valid-mx"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "bad-delete-custom.bar.com",
						},
					},
				}},
			},
			customHostnamesEnabled: true,
			errorLogCount:          1,
		},
		{
			name: "Update add/remove error (custom hostnames enabled)",
			changes: &plan.Changes{
				UpdateNew: []*endpoint.Endpoint{{
					DNSName:    "bad-update-add.bar.com",
					RecordType: "MX",
					Targets:    endpoint.Targets{"not-a-valid-mx"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "bad-update-add-custom.bar.com",
						},
					},
				}},
				UpdateOld: []*endpoint.Endpoint{{
					DNSName:    "old-bad-update-add.bar.com",
					RecordType: "MX",
					Targets:    endpoint.Targets{"not-a-valid-mx-but-still-updated"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "bad-update-add-custom.bar.com",
						},
					},
				}},
			},
			customHostnamesEnabled: true,
			errorLogCount:          2,
		},
		{
			name: "Update leave error (custom hostnames enabled)",
			changes: &plan.Changes{
				UpdateOld: []*endpoint.Endpoint{{
					DNSName:    "bad-update-leave.bar.com",
					RecordType: "MX",
					Targets:    endpoint.Targets{"not-a-valid-mx"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "bad-update-leave-custom.bar.com",
						},
					},
				}},
				UpdateNew: []*endpoint.Endpoint{{
					DNSName:    "bad-update-leave.bar.com",
					RecordType: "MX",
					Targets:    endpoint.Targets{"not-a-valid-mx"},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "bad-update-leave-custom.bar.com",
						},
					},
				}},
			},
			customHostnamesEnabled: true,
			errorLogCount:          1,
		},
		{
			name: "Delete error (custom hostnames disabled)",
			changes: &plan.Changes{
				Delete: []*endpoint.Endpoint{{
					DNSName:    "bad-delete2.bar.com",
					RecordType: "MX",
					Targets:    endpoint.Targets{"not-a-valid-mx"},
				}},
			},
			customHostnamesEnabled: false,
			errorLogCount:          1,
		},
	}

	// Test with custom hostnames enabled and disabled
	for _, tc := range cases {
		if tc.customHostnamesEnabled {
			provider.CustomHostnamesConfig = CustomHostnamesConfig{Enabled: true}
		} else {
			provider.CustomHostnamesConfig = CustomHostnamesConfig{Enabled: false}
		}
		hook.Reset()
		err := provider.ApplyChanges(context.Background(), tc.changes)
		assert.NoError(t, err, "ApplyChanges should not return error for newCloudFlareChange error (it should log and continue)")
		errorLogCount := 0
		for _, entry := range hook.Entries {
			if entry.Level == log.ErrorLevel &&
				strings.Contains(entry.Message, "failed to create cloudflare change") {
				errorLogCount++
			}
		}
		assert.Equal(t, tc.errorLogCount, errorLogCount, "expected error log count for %s", tc.name)
	}
}

func TestCloudFlareProvider_SupportedAdditionalRecordTypes(t *testing.T) {
	provider := &CloudFlareProvider{}

	tests := []struct {
		recordType string
		expected   bool
	}{
		{endpoint.RecordTypeMX, true},
		{endpoint.RecordTypeA, true},
		{endpoint.RecordTypeCNAME, true},
		{endpoint.RecordTypeTXT, true},
		{endpoint.RecordTypeNS, true},
		{"SRV", true},
		{"SPF", false},
		{"LOC", false},
		{"UNKNOWN", false},
	}

	for _, tt := range tests {
		t.Run(tt.recordType, func(t *testing.T) {
			result := provider.SupportedAdditionalRecordTypes(tt.recordType)
			assert.Equal(t, tt.expected, result)
		})
	}
}
