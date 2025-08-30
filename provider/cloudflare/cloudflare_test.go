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
	"time"

	cloudflarev0 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v5/dns"
	"github.com/cloudflare/cloudflare-go/v5/zones"
	"github.com/maxatome/go-testdeep/td"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/source/annotations"
)

type MockAction struct {
	Name             string
	ZoneId           string
	RecordId         string
	RecordData       dns.RecordResponse
	RegionalHostname regionalHostname
}

type mockCloudFlareClient struct {
	Zones             map[string]string
	Records           map[string]map[string]dns.RecordResponse
	Actions           []MockAction
	listZonesError    error // For v4 ListZones
	getZoneError      error // For v4 GetZone
	dnsRecordsError   error
	customHostnames   map[string][]cloudflarev0.CustomHostname
	regionalHostnames map[string][]regionalHostname
}

var ExampleDomain = []dns.RecordResponse{
	{
		ID:      "1234567890",
		Name:    "foobar.bar.com",
		Type:    endpoint.RecordTypeA,
		TTL:     120,
		Content: "1.2.3.4",
		Proxied: false,
		Comment: "valid comment",
	},
	{
		ID:      "2345678901",
		Name:    "foobar.bar.com",
		Type:    endpoint.RecordTypeA,
		TTL:     120,
		Content: "3.4.5.6",
		Proxied: false,
	},
	{
		ID:      "1231231233",
		Name:    "bar.foo.com",
		Type:    endpoint.RecordTypeA,
		TTL:     1,
		Content: "2.3.4.5",
		Proxied: false,
	},
}

func NewMockCloudFlareClient() *mockCloudFlareClient {
	return &mockCloudFlareClient{
		Zones: map[string]string{
			"001": "bar.com",
			"002": "foo.com",
		},
		Records: map[string]map[string]dns.RecordResponse{
			"001": {},
			"002": {},
		},
		customHostnames:   map[string][]cloudflarev0.CustomHostname{},
		regionalHostnames: map[string][]regionalHostname{},
	}
}

func NewMockCloudFlareClientWithRecords(records map[string][]dns.RecordResponse) *mockCloudFlareClient {
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

func generateDNSRecordID(rrtype string, name string, content string) string {
	return fmt.Sprintf("%s-%s-%s", name, rrtype, content)
}

func (m *mockCloudFlareClient) CreateDNSRecord(ctx context.Context, params dns.RecordNewParams) (*dns.RecordResponse, error) {
	body := params.Body.(dns.RecordNewParamsBody)

	record := dns.RecordResponse{
		ID:       generateDNSRecordID(body.Type.String(), body.Name.Value, body.Content.Value),
		Name:     body.Name.Value,
		TTL:      dns.TTL(body.TTL.Value),
		Proxied:  body.Proxied.Value,
		Type:     dns.RecordResponseType(body.Type.String()),
		Content:  body.Content.Value,
		Priority: body.Priority.Value,
	}

	m.Actions = append(m.Actions, MockAction{
		Name:       "Create",
		ZoneId:     params.ZoneID.Value,
		RecordId:   record.ID,
		RecordData: record,
	})
	if zone, ok := m.Records[params.ZoneID.Value]; ok {
		zone[record.ID] = record
	}

	if record.Name == "newerror.bar.com" {
		return nil, fmt.Errorf("failed to create record")
	}
	return &record, nil
}

func (m *mockCloudFlareClient) ListDNSRecords(ctx context.Context, rc *cloudflarev0.ResourceContainer, rp cloudflarev0.ListDNSRecordsParams) ([]dns.RecordResponse, *cloudflarev0.ResultInfo, error) {
	if m.dnsRecordsError != nil {
		return nil, &cloudflarev0.ResultInfo{}, m.dnsRecordsError
	}
	result := []dns.RecordResponse{}
	if zone, ok := m.Records[rc.Identifier]; ok {
		for _, record := range zone {
			if strings.HasPrefix(record.Name, "newerror-list-") {
				m.DeleteDNSRecord(ctx, rc, record.ID)
				return nil, &cloudflarev0.ResultInfo{}, errors.New("failed to list erroring DNS record")
			}
			result = append(result, record)
		}
	}

	if len(result) == 0 || rp.PerPage == 0 {
		return result, &cloudflarev0.ResultInfo{Page: 1, TotalPages: 1, Count: 0, Total: 0}, nil
	}

	// if not pagination options were passed in, return the result as is
	if rp.Page == 0 {
		return result, &cloudflarev0.ResultInfo{Page: 1, TotalPages: 1, Count: len(result), Total: len(result)}, nil
	}

	// otherwise, split the result into chunks of size rp.PerPage to simulate the pagination from the API
	chunks := [][]dns.RecordResponse{}

	// to ensure consistency in the multiple calls to this function, sort the result slice
	sort.Slice(result, func(i, j int) bool { return strings.Compare(result[i].ID, result[j].ID) > 0 })
	for rp.PerPage < len(result) {
		result, chunks = result[rp.PerPage:], append(chunks, result[0:rp.PerPage])
	}
	chunks = append(chunks, result)

	// return the requested page
	partialResult := chunks[rp.Page-1]
	return partialResult, &cloudflarev0.ResultInfo{
		PerPage:    rp.PerPage,
		Page:       rp.Page,
		TotalPages: len(chunks),
		Count:      len(partialResult),
		Total:      len(result),
	}, nil
}

func (m *mockCloudFlareClient) UpdateDNSRecord(ctx context.Context, recordID string, params dns.RecordUpdateParams) (res *dns.RecordResponse, err error) {
	zoneID := params.ZoneID.String()
	body := params.Body.(dns.RecordUpdateParamsBody)

	record := dns.RecordResponse{
		ID:       recordID,
		Name:     body.Name.Value,
		TTL:      dns.TTL(body.TTL.Value),
		Proxied:  body.Proxied.Value,
		Type:     dns.RecordResponseType(body.Type.String()),
		Content:  body.Content.Value,
		Priority: body.Priority.Value,
	}

	m.Actions = append(m.Actions, MockAction{
		Name:       "Update",
		ZoneId:     zoneID,
		RecordId:   recordID,
		RecordData: record,
	})
	if zone, ok := m.Records[zoneID]; ok {
		if _, ok := zone[recordID]; ok {
			if strings.HasPrefix(record.Name, "newerror-update-") {
				return nil, errors.New("failed to update erroring DNS record")
			}
			zone[recordID] = record
		}
	}
	return &record, nil
}

func (m *mockCloudFlareClient) DeleteDNSRecord(ctx context.Context, rc *cloudflarev0.ResourceContainer, recordID string) error {
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

func (m *mockCloudFlareClient) CustomHostnames(ctx context.Context, zoneID string, page int, filter cloudflarev0.CustomHostname) ([]cloudflarev0.CustomHostname, cloudflarev0.ResultInfo, error) {
	var err error = nil
	perPage := 50 // cloudflare-go v0 API hardcoded

	if strings.HasPrefix(zoneID, "newerror-") {
		return nil, cloudflarev0.ResultInfo{}, errors.New("failed to list custom hostnames")
	}
	if filter.Hostname != "" {
		err = errors.New("filters are not supported for custom hostnames mock test")
		return nil, cloudflarev0.ResultInfo{}, err
	}
	if page < 1 {
		err = errors.New("incorrect page value for custom hostnames list")
		return nil, cloudflarev0.ResultInfo{}, err
	}

	result := []cloudflarev0.CustomHostname{}
	if chs, ok := m.customHostnames[zoneID]; ok {
		for idx := (page - 1) * perPage; idx < min(len(chs), page*perPage); idx++ {
			ch := m.customHostnames[zoneID][idx]
			if strings.HasPrefix(ch.Hostname, "newerror-list-") {
				m.DeleteCustomHostname(ctx, zoneID, ch.ID)
				return nil, cloudflarev0.ResultInfo{}, errors.New("failed to list erroring custom hostname")
			}
			result = append(result, ch)
		}
		return result,
			cloudflarev0.ResultInfo{
				Page:       page,
				PerPage:    perPage,
				Count:      len(result),
				Total:      len(chs),
				TotalPages: len(chs)/page + 1,
			}, err
	} else {
		return result,
			cloudflarev0.ResultInfo{
				Page:       page,
				PerPage:    perPage,
				Count:      0,
				Total:      0,
				TotalPages: 0,
			}, err
	}
}

func (m *mockCloudFlareClient) CreateCustomHostname(ctx context.Context, zoneID string, ch cloudflarev0.CustomHostname) (*cloudflarev0.CustomHostnameResponse, error) {
	if ch.Hostname == "" || ch.CustomOriginServer == "" || ch.Hostname == "newerror-create.foo.fancybar.com" {
		return nil, fmt.Errorf("Invalid custom hostname or origin hostname")
	}
	if _, ok := m.customHostnames[zoneID]; !ok {
		m.customHostnames[zoneID] = []cloudflarev0.CustomHostname{}
	}
	var newCustomHostname cloudflarev0.CustomHostname = ch
	newCustomHostname.ID = fmt.Sprintf("ID-%s", ch.Hostname)
	m.customHostnames[zoneID] = append(m.customHostnames[zoneID], newCustomHostname)
	return &cloudflarev0.CustomHostnameResponse{}, nil
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
	// Simulate iterator error (line 144)
	if m.listZonesError != nil {
		return "", fmt.Errorf("failed to list zones from CloudFlare API: %w", m.listZonesError)
	}

	for id, name := range m.Zones {
		if name == zoneName {
			return id, nil
		}
	}

	// Use the improved error message (line 147)
	return "", fmt.Errorf("zone %q not found in CloudFlare account - verify the zone exists and API credentials have access to it", zoneName)
}

// V4 Zone methods
func (m *mockCloudFlareClient) ListZones(ctx context.Context, params zones.ZoneListParams) autoPager[zones.Zone] {
	if m.listZonesError != nil {
		return &mockAutoPager[zones.Zone]{
			err: m.listZonesError,
		}
	}

	var results []zones.Zone

	for id, zoneName := range m.Zones {
		results = append(results, zones.Zone{
			ID:   id,
			Name: zoneName,
			Plan: zones.ZonePlan{IsSubscribed: strings.HasSuffix(zoneName, "bar.com")}, //nolint:SA1019 // Plan.IsSubscribed is deprecated but no replacement available yet
		})
	}

	return &mockAutoPager[zones.Zone]{
		items: results,
	}
}

func (m *mockCloudFlareClient) GetZone(ctx context.Context, zoneID string) (*zones.Zone, error) {
	if m.getZoneError != nil {
		return nil, m.getZoneError
	}

	for id, zoneName := range m.Zones {
		if zoneID == id {
			return &zones.Zone{
				ID:   zoneID,
				Name: zoneName,
				Plan: zones.ZonePlan{IsSubscribed: strings.HasSuffix(zoneName, "bar.com")}, //nolint:SA1019 // Plan.IsSubscribed is deprecated but no replacement available yet
			}, nil
		}
	}

	return nil, errors.New("Unknown zoneID: " + zoneID)
}

func getCustomHostnameIdxByID(chs []cloudflarev0.CustomHostname, customHostnameID string) int {
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: false,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("A", "bar.com", "127.0.0.2"),
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.2"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.2",
				TTL:     1,
				Proxied: false,
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("CNAME", "cname.bar.com", "google.com"),
				Type:    "CNAME",
				Name:    "cname.bar.com",
				Content: "google.com",
				TTL:     1,
				Proxied: false,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("CNAME", "cname.bar.com", "facebook.com"),
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("CNAME", "cname.bar.com", "facebook.com"),
				Type:    "CNAME",
				Name:    "cname.bar.com",
				Content: "facebook.com",
				TTL:     1,
				Proxied: false,
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
			RecordData: dns.RecordResponse{
				ID:       generateDNSRecordID("MX", "mx.bar.com", "google.com"),
				Type:     "MX",
				Name:     "mx.bar.com",
				Content:  "google.com",
				Priority: 10,
				TTL:      1,
				Proxied:  false,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("MX", "mx.bar.com", "facebook.com"),
			RecordData: dns.RecordResponse{
				ID:       generateDNSRecordID("MX", "mx.bar.com", "facebook.com"),
				Type:     "MX",
				Name:     "mx.bar.com",
				Content:  "facebook.com",
				Priority: 20,
				TTL:      1,
				Proxied:  false,
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("TXT", "txt.bar.com", "v=spf1 include:_spf.google.com ~all"),
				Type:    "TXT",
				Name:    "txt.bar.com",
				Content: "v=spf1 include:_spf.google.com ~all",
				TTL:     1,
				Proxied: false,
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "ttl.bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "ttl.bar.com",
				Content: "127.0.0.1",
				TTL:     120,
				Proxied: false,
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: true,
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: true,
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: false,
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "bar.com", "127.0.0.1"),
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: true,
			},
		},
	},
		[]string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	)
}

func TestCloudflareSetProxied(t *testing.T) {
	testCases := []struct {
		recordType string
		domain     string
		proxiable  bool
	}{
		{"A", "bar.com", true},
		{"CNAME", "bar.com", true},
		{"TXT", "bar.com", false},
		{"MX", "bar.com", false},
		{"NS", "bar.com", false},
		{"SPF", "bar.com", false},
		{"SRV", "bar.com", false},
		{"A", "*.bar.com", true},
		{"CNAME", "*.docs.bar.com", true},
	}

	for _, testCase := range testCases {
		var targets endpoint.Targets
		var content string
		var priority float64

		if testCase.recordType == "MX" {
			targets = endpoint.Targets{"10 mx.example.com"}
			content = "mx.example.com"
			priority = 10
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
		recordData := dns.RecordResponse{
			ID:      expectedID,
			Type:    dns.RecordResponseType(testCase.recordType),
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
	client.getZoneError = errors.New("zone lookup failed")

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
	client.listZonesError = &cloudflarev0.Error{
		StatusCode: 429,
		ErrorCodes: []int{10000},
		Type:       cloudflarev0.ErrorTypeRateLimit,
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
	client.listZonesError = errors.New("exceeded available rate limit retries")
	p := &CloudFlareProvider{Client: client}

	// Call the Zones function
	_, err := p.Zones(context.Background())

	// Assert that a soft error was returned
	assert.ErrorIs(t, err, provider.SoftError, "expected a rate limit error")
}

func TestCloudflareListZoneInternalErrors(t *testing.T) {
	// Create a mock client that returns a internal server error
	client := NewMockCloudFlareClient()
	client.listZonesError = &cloudflarev0.Error{
		StatusCode: 500,
		ErrorCodes: []int{20000},
		Type:       cloudflarev0.ErrorTypeService,
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
	client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
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
	client.listZonesError = &cloudflarev0.Error{
		StatusCode: 429,
		ErrorCodes: []int{10000},
		Type:       cloudflarev0.ErrorTypeRateLimit,
	}
	_, err = p.Records(ctx)
	// Assert that a soft error was returned
	if !errors.Is(err, provider.SoftError) {
		t.Error("expected a rate limit error")
	}

	client.listZonesError = &cloudflarev0.Error{
		StatusCode: 500,
		ErrorCodes: []int{10000},
		Type:       cloudflarev0.ErrorTypeService,
	}
	_, err = p.Records(ctx)
	// Assert that a soft error was returned
	if !errors.Is(err, provider.SoftError) {
		t.Error("expected a internal server error")
	}

	client.listZonesError = errors.New("failed to list zones")
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("", "new.bar.com", "target"),
				Name:    "new.bar.com",
				Content: "target",
				TTL:     1,
				Proxied: false,
			},
		},
		{
			Name:     "Create",
			ZoneId:   "001",
			RecordId: generateDNSRecordID("", "foobar.bar.com", "target-new"),
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("", "foobar.bar.com", "target-new"),
				Name:    "foobar.bar.com",
				Content: "target-new",
				TTL:     1,
				Proxied: false,
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

	assert.Empty(t, p.getRecordID(recordsMap, dns.RecordResponse{
		Name:    "foo.com",
		Type:    endpoint.RecordTypeA,
		Content: "foobar",
	}))

	assert.Empty(t, p.getRecordID(recordsMap, dns.RecordResponse{
		Name:    "foo.com",
		Type:    endpoint.RecordTypeCNAME,
		Content: "fizfuz",
	}))

	assert.Equal(t, "1", p.getRecordID(recordsMap, dns.RecordResponse{
		Name:    "foo.com",
		Type:    endpoint.RecordTypeCNAME,
		Content: "foobar",
	}))
	assert.Empty(t, p.getRecordID(recordsMap, dns.RecordResponse{
		Name:    "bar.de",
		Type:    endpoint.RecordTypeA,
		Content: "2.3.4.5",
	}))
	assert.Equal(t, "2", p.getRecordID(recordsMap, dns.RecordResponse{
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
		Records           []dns.RecordResponse
		ExpectedEndpoints []*endpoint.Endpoint
	}{
		{
			Name:              "empty",
			Records:           []dns.RecordResponse{},
			ExpectedEndpoints: []*endpoint.Endpoint{},
		},
		{
			Name: "single record - single target",
			Records: []dns.RecordResponse{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: false,
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
			Records: []dns.RecordResponse{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: false,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: false,
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
			Records: []dns.RecordResponse{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: false,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: false,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: false,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: false,
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
			Records: []dns.RecordResponse{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: false,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: false,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: false,
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
			Records: []dns.RecordResponse{
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: false,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultTTL,
					Proxied: false,
				},
				{
					Name:    "bar.de",
					Type:    "NOT SUPPORTED",
					Content: "10.10.10.1",
					TTL:     defaultTTL,
					Proxied: false,
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
	client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
		"001": {
			{
				ID:       "mx-1",
				Name:     "mx.bar.com",
				Type:     endpoint.RecordTypeMX,
				TTL:      3600,
				Content:  "mail.bar.com",
				Priority: 10,
			},
			{
				ID:       "mx-2",
				Name:     "mx.bar.com",
				Type:     endpoint.RecordTypeMX,
				TTL:      3600,
				Content:  "mail2.bar.com",
				Priority: 20,
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
		RecordsAreProxied        bool
		ShouldBeUpdated          bool
	}{
		{
			Name:                     "ProxyDefault: false, ShouldBeProxied: false, ExpectUpdates: false",
			ProviderProxiedByDefault: false,
			RecordsAreProxied:        false,
			ShouldBeUpdated:          false,
		},
		{
			Name:                     "ProxyDefault: true, ShouldBeProxied: true, ExpectUpdates: false",
			ProviderProxiedByDefault: true,
			RecordsAreProxied:        true,
			ShouldBeUpdated:          false,
		},
		{
			Name:                     "ProxyDefault: true, ShouldBeProxied: false, ExpectUpdates: true",
			ProviderProxiedByDefault: true,
			RecordsAreProxied:        false,
			ShouldBeUpdated:          true,
		},
		{
			Name:                     "ProxyDefault: false, ShouldBeProxied: true, ExpectUpdates: true",
			ProviderProxiedByDefault: false,
			RecordsAreProxied:        true,
			ShouldBeUpdated:          true,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
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
	client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
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
			RecordData: dns.RecordResponse{
				ID:      generateDNSRecordID("A", "foobar.bar.com", "2.3.4.5"),
				Name:    "foobar.bar.com",
				Type:    "A",
				Content: "2.3.4.5",
				TTL:     1,
				Proxied: true,
			},
		},
		{
			Name:     "Update",
			ZoneId:   "001",
			RecordId: "1234567890",
			RecordData: dns.RecordResponse{
				ID:      "1234567890",
				Name:    "foobar.bar.com",
				Type:    "A",
				Content: "1.2.3.4",
				TTL:     1,
				Proxied: true,
			},
		},
	})
}

func TestCustomTTLWithEnabledProxyNotChanged(t *testing.T) {
	client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
		"001": {
			{
				ID:      "1234567890",
				Name:    "foobar.bar.com",
				Type:    endpoint.RecordTypeA,
				TTL:     1,
				Content: "1.2.3.4",
				Proxied: true,
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
	if change.RegionalHostname.regionKey != "us" {
		t.Errorf("expected region key to be 'us', but got '%s'", change.RegionalHostname.regionKey)
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
	client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
		"001": {
			{
				ID:      "1234567890",
				Name:    "my-domain-here.app",
				Type:    endpoint.RecordTypeCNAME,
				TTL:     1,
				Content: "my-tunnel-guid-here.cfargotunnel.com",
				Proxied: true,
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
			ResourceRecord: dns.RecordResponse{
				Name:    "my-domain-here.app",
				Type:    endpoint.RecordTypeCNAME,
				ID:      "1234567890",
				Content: "my-tunnel-guid-here.cfargotunnel.com",
			},
			RegionalHostname: regionalHostname{
				hostname: "my-domain-here.app",
			},
		},
		{
			Action: cloudFlareUpdate,
			ResourceRecord: dns.RecordResponse{
				Name:    "my-domain-here.app",
				Type:    endpoint.RecordTypeTXT,
				ID:      "9876543210",
				Content: "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/external-dns/my-domain-here-app",
			},
			RegionalHostname: regionalHostname{
				hostname:  "my-domain-here.app",
				regionKey: "",
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
	client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
		"001": {
			{
				ID:      "1234567890",
				Name:    "@", // APEX record
				Type:    endpoint.RecordTypeCNAME,
				TTL:     1,
				Content: "my-tunnel-guid-here.cfargotunnel.com",
				Proxied: true,
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
			ResourceRecord: dns.RecordResponse{
				Name:    "@", // APEX record
				Type:    endpoint.RecordTypeCNAME,
				ID:      "1234567890",
				Content: "my-tunnel-guid-here.cfargotunnel.com",
			},
			RegionalHostname: regionalHostname{
				hostname: "@", // APEX record
			},
		},
		{
			Action: cloudFlareUpdate,
			ResourceRecord: dns.RecordResponse{
				Name:    "@", // APEX record
				Type:    endpoint.RecordTypeTXT,
				ID:      "9876543210",
				Content: "heritage=external-dns,external-dns/owner=default,external-dns/resource=service/external-dns/my-domain-here-app",
			},
			RegionalHostname: regionalHostname{
				hostname:  "@", // APEX record
				regionKey: "",
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
		Records:         map[string]map[string]dns.RecordResponse{},
		customHostnames: map[string][]cloudflarev0.CustomHostname{},
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
	client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
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
				ch := cloudflarev0.CustomHostname{
					Hostname:           "corrupted-newerror-getCustomHostnameOrigin.foo.fancybar.com",
					CustomOriginServer: oldCh.CustomOriginServer,
					SSL:                oldCh.SSL,
				}
				client.customHostnames[zoneID][oldIdx] = ch
			}
		} else if tc.preApplyHook == "duplicate" { // manually inject duplicating custom hostname with the same name and origin
			ch := cloudflarev0.CustomHostname{
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

	client.getZoneError = errors.New("zone lookup failed")
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

func TestCloudflareZoneChanges(t *testing.T) {
	client := NewMockCloudFlareClient()
	cfProvider := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com", "bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}

	// Test zone listing and filtering
	zones, err := cfProvider.Zones(context.Background())
	assert.NoError(t, err)
	assert.Len(t, zones, 2)

	// Verify zone names
	zoneNames := make([]string, len(zones))
	for i, zone := range zones {
		zoneNames[i] = zone.Name
	}
	assert.Contains(t, zoneNames, "foo.com")
	assert.Contains(t, zoneNames, "bar.com")

	// Test zone filtering with specific zone ID
	providerWithZoneFilter := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com", "bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{"001"}),
	}

	filteredZones, err := providerWithZoneFilter.Zones(context.Background())
	assert.NoError(t, err)
	assert.Len(t, filteredZones, 1)
	assert.Equal(t, "bar.com", filteredZones[0].Name) // zone 001 is bar.com
	assert.Equal(t, "001", filteredZones[0].ID)

	// Test zone changes grouping
	changes := []*cloudFlareChange{
		{
			Action:         cloudFlareCreate,
			ResourceRecord: dns.RecordResponse{Name: "test1.foo.com", Type: "A", Content: "1.2.3.4"},
		},
		{
			Action:         cloudFlareCreate,
			ResourceRecord: dns.RecordResponse{Name: "test2.foo.com", Type: "A", Content: "1.2.3.5"},
		},
		{
			Action:         cloudFlareCreate,
			ResourceRecord: dns.RecordResponse{Name: "test1.bar.com", Type: "A", Content: "1.2.3.6"},
		},
	}

	changesByZone := cfProvider.changesByZone(zones, changes)
	assert.Len(t, changesByZone, 2)
	assert.Len(t, changesByZone["001"], 1) // bar.com zone (test1.bar.com)
	assert.Len(t, changesByZone["002"], 2) // foo.com zone (test1.foo.com, test2.foo.com)

	// Test paid plan detection
	assert.False(t, cfProvider.ZoneHasPaidPlan("subdomain.foo.com")) // free plan
	assert.True(t, cfProvider.ZoneHasPaidPlan("subdomain.bar.com"))  // paid plan
}

func TestCloudflareZoneErrors(t *testing.T) {
	client := NewMockCloudFlareClient()

	// Test list zones error
	client.listZonesError = errors.New("failed to list zones")
	cfProvider := &CloudFlareProvider{
		Client: client,
	}

	zones, err := cfProvider.Zones(context.Background())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list zones")
	assert.Nil(t, zones)

	// Test get zone error
	client.listZonesError = nil
	client.getZoneError = errors.New("failed to get zone")

	// This should still work for listing but fail when getting individual zones
	zones, err = cfProvider.Zones(context.Background())
	assert.NoError(t, err) // List works, individual gets may fail internally
	assert.NotNil(t, zones)
}

func TestCloudflareZoneFiltering(t *testing.T) {
	client := NewMockCloudFlareClient()

	// Test with domain filter only
	cfProvider := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}

	zones, err := cfProvider.Zones(context.Background())
	assert.NoError(t, err)
	assert.Len(t, zones, 1)
	assert.Equal(t, "foo.com", zones[0].Name)

	// Test with zone ID filter
	providerWithIDFilter := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{"002"}),
	}

	filteredZones, err := providerWithIDFilter.Zones(context.Background())
	assert.NoError(t, err)
	assert.Len(t, filteredZones, 1)
	assert.Equal(t, "foo.com", filteredZones[0].Name) // zone 002 is foo.com
	assert.Equal(t, "002", filteredZones[0].ID)
}

func TestCloudflareZonePlanDetection(t *testing.T) {
	client := NewMockCloudFlareClient()
	cfProvider := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com", "bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}

	// Test free plan detection (foo.com)
	assert.False(t, cfProvider.ZoneHasPaidPlan("foo.com"))
	assert.False(t, cfProvider.ZoneHasPaidPlan("subdomain.foo.com"))
	assert.False(t, cfProvider.ZoneHasPaidPlan("deep.subdomain.foo.com"))

	// Test paid plan detection (bar.com)
	assert.True(t, cfProvider.ZoneHasPaidPlan("bar.com"))
	assert.True(t, cfProvider.ZoneHasPaidPlan("subdomain.bar.com"))
	assert.True(t, cfProvider.ZoneHasPaidPlan("deep.subdomain.bar.com"))

	// Test invalid domain
	assert.False(t, cfProvider.ZoneHasPaidPlan("invalid.domain.com"))

	// Test with zone error
	client.getZoneError = errors.New("zone lookup failed")
	providerWithError := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com", "bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}
	assert.False(t, providerWithError.ZoneHasPaidPlan("subdomain.foo.com"))
}

func TestCloudflareChangesByZone(t *testing.T) {
	client := NewMockCloudFlareClient()
	cfProvider := &CloudFlareProvider{
		Client:       client,
		domainFilter: endpoint.NewDomainFilter([]string{"foo.com", "bar.com"}),
		zoneIDFilter: provider.NewZoneIDFilter([]string{""}),
	}

	zones, err := cfProvider.Zones(context.Background())
	assert.NoError(t, err)
	assert.Len(t, zones, 2)

	// Test empty changes
	emptyChanges := []*cloudFlareChange{}
	changesByZone := cfProvider.changesByZone(zones, emptyChanges)
	assert.Len(t, changesByZone, 2)       // Should return map with zones but empty slices
	assert.Empty(t, changesByZone["001"]) // bar.com zone should have no changes
	assert.Empty(t, changesByZone["002"]) // foo.com zone should have no changes

	// Test changes for different zones
	changes := []*cloudFlareChange{
		{
			Action:         cloudFlareCreate,
			ResourceRecord: dns.RecordResponse{Name: "api.foo.com", Type: "A", Content: "1.2.3.4"},
		},
		{
			Action:         cloudFlareUpdate,
			ResourceRecord: dns.RecordResponse{Name: "www.foo.com", Type: "CNAME", Content: "foo.com"},
		},
		{
			Action:         cloudFlareCreate,
			ResourceRecord: dns.RecordResponse{Name: "mail.bar.com", Type: "MX", Content: "10 mail.bar.com"},
		},
		{
			Action:         cloudFlareDelete,
			ResourceRecord: dns.RecordResponse{Name: "old.bar.com", Type: "A", Content: "5.6.7.8"},
		},
	}

	changesByZone = cfProvider.changesByZone(zones, changes)
	assert.Len(t, changesByZone, 2)

	// Verify bar.com zone changes (zone 001)
	barChanges := changesByZone["001"]
	assert.Len(t, barChanges, 2)
	assert.Equal(t, "mail.bar.com", barChanges[0].ResourceRecord.Name)
	assert.Equal(t, "old.bar.com", barChanges[1].ResourceRecord.Name)

	// Verify foo.com zone changes (zone 002)
	fooChanges := changesByZone["002"]
	assert.Len(t, fooChanges, 2)
	assert.Equal(t, "api.foo.com", fooChanges[0].ResourceRecord.Name)
	assert.Equal(t, "www.foo.com", fooChanges[1].ResourceRecord.Name)
}

func TestConvertCloudflareError(t *testing.T) {
	tests := []struct {
		name            string
		inputError      error
		expectSoftError bool
		description     string
	}{
		{
			name:            "Rate limit error via Error type",
			inputError:      &cloudflarev0.Error{StatusCode: 429, Type: cloudflarev0.ErrorTypeRateLimit},
			expectSoftError: true,
			description:     "CloudFlare API rate limit error should be converted to soft error",
		},
		{
			name:            "Rate limit error via ClientRateLimited",
			inputError:      &cloudflarev0.Error{StatusCode: 429, ErrorCodes: []int{10000}, Type: cloudflarev0.ErrorTypeRateLimit}, // Complete rate limit error
			expectSoftError: true,
			description:     "CloudFlare client rate limited error should be converted to soft error",
		},
		{
			name:            "Server error 500",
			inputError:      &cloudflarev0.Error{StatusCode: 500},
			expectSoftError: true,
			description:     "Server error (500+) should be converted to soft error",
		},
		{
			name:            "Server error 502",
			inputError:      &cloudflarev0.Error{StatusCode: 502},
			expectSoftError: true,
			description:     "Server error (502) should be converted to soft error",
		},
		{
			name:            "Server error 503",
			inputError:      &cloudflarev0.Error{StatusCode: 503},
			expectSoftError: true,
			description:     "Server error (503) should be converted to soft error",
		},
		{
			name:            "Rate limit string error",
			inputError:      errors.New("exceeded available rate limit retries"),
			expectSoftError: true,
			description:     "String error containing rate limit message should be converted to soft error",
		},
		{
			name:            "Rate limit string error mixed case",
			inputError:      errors.New("request failed: exceeded available rate limit retries for this operation"),
			expectSoftError: true,
			description:     "String error containing rate limit message should be converted to soft error regardless of context",
		},
		{
			name:            "Client error 400",
			inputError:      &cloudflarev0.Error{StatusCode: 400},
			expectSoftError: false,
			description:     "Client error (400) should not be converted to soft error",
		},
		{
			name:            "Client error 401",
			inputError:      &cloudflarev0.Error{StatusCode: 401},
			expectSoftError: false,
			description:     "Client error (401) should not be converted to soft error",
		},
		{
			name:            "Client error 404",
			inputError:      &cloudflarev0.Error{StatusCode: 404},
			expectSoftError: false,
			description:     "Client error (404) should not be converted to soft error",
		},
		{
			name:            "Generic error",
			inputError:      errors.New("some generic error"),
			expectSoftError: false,
			description:     "Generic error should not be converted to soft error",
		},
		{
			name:            "Network error",
			inputError:      errors.New("connection refused"),
			expectSoftError: false,
			description:     "Network error should not be converted to soft error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertCloudflareError(tt.inputError)

			if tt.expectSoftError {
				assert.ErrorIs(t, result, provider.SoftError,
					"Expected soft error for %s: %s", tt.name, tt.description)

				// Verify the original error message is preserved in the soft error
				assert.Contains(t, result.Error(), tt.inputError.Error(),
					"Original error message should be preserved")
			} else {
				assert.NotErrorIs(t, result, provider.SoftError,
					"Expected non-soft error for %s: %s", tt.name, tt.description)
				assert.Equal(t, tt.inputError, result,
					"Non-soft errors should be returned unchanged")
			}
		})
	}
}

func TestConvertCloudflareErrorInContext(t *testing.T) {
	tests := []struct {
		name            string
		setupMock       func(*mockCloudFlareClient)
		function        func(*CloudFlareProvider) error
		expectSoftError bool
		description     string
	}{
		{
			name: "Zones with GetZone rate limit error",
			setupMock: func(client *mockCloudFlareClient) {
				client.Zones = map[string]string{"zone1": "example.com"}
				client.getZoneError = &cloudflarev0.Error{StatusCode: 429, Type: cloudflarev0.ErrorTypeRateLimit}
			},
			function: func(p *CloudFlareProvider) error {
				p.zoneIDFilter.ZoneIDs = []string{"zone1"}
				_, err := p.Zones(context.Background())
				return err
			},
			expectSoftError: true,
			description:     "Zones function should convert GetZone rate limit errors to soft errors",
		},
		{
			name: "Zones with GetZone server error",
			setupMock: func(client *mockCloudFlareClient) {
				client.Zones = map[string]string{"zone1": "example.com"}
				client.getZoneError = &cloudflarev0.Error{StatusCode: 500}
			},
			function: func(p *CloudFlareProvider) error {
				p.zoneIDFilter.ZoneIDs = []string{"zone1"}
				_, err := p.Zones(context.Background())
				return err
			},
			expectSoftError: true,
			description:     "Zones function should convert GetZone server errors to soft errors",
		},
		{
			name: "Zones with GetZone client error",
			setupMock: func(client *mockCloudFlareClient) {
				client.Zones = map[string]string{"zone1": "example.com"}
				client.getZoneError = &cloudflarev0.Error{StatusCode: 404}
			},
			function: func(p *CloudFlareProvider) error {
				p.zoneIDFilter.ZoneIDs = []string{"zone1"}
				_, err := p.Zones(context.Background())
				return err
			},
			expectSoftError: false,
			description:     "Zones function should not convert GetZone client errors to soft errors",
		},
		{
			name: "Zones with ListZones rate limit error",
			setupMock: func(client *mockCloudFlareClient) {
				client.listZonesError = errors.New("exceeded available rate limit retries")
			},
			function: func(p *CloudFlareProvider) error {
				_, err := p.Zones(context.Background())
				return err
			},
			expectSoftError: true,
			description:     "Zones function should convert ListZones rate limit string errors to soft errors",
		},
		{
			name: "Zones with ListZones server error",
			setupMock: func(client *mockCloudFlareClient) {
				client.listZonesError = &cloudflarev0.Error{StatusCode: 503}
			},
			function: func(p *CloudFlareProvider) error {
				_, err := p.Zones(context.Background())
				return err
			},
			expectSoftError: true,
			description:     "Zones function should convert ListZones server errors to soft errors",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewMockCloudFlareClient()
			tt.setupMock(client)

			p := &CloudFlareProvider{
				Client:       client,
				zoneIDFilter: provider.ZoneIDFilter{},
			}

			err := tt.function(p)
			assert.Error(t, err, "Expected an error from %s", tt.name)

			if tt.expectSoftError {
				assert.ErrorIs(t, err, provider.SoftError,
					"Expected soft error for %s: %s", tt.name, tt.description)
			} else {
				assert.NotErrorIs(t, err, provider.SoftError,
					"Expected non-soft error for %s: %s", tt.name, tt.description)
			}
		})
	}
}

func TestCloudFlareZonesDomainFilter(t *testing.T) {
	// Set required environment variables for CloudFlare provider
	t.Setenv("CF_API_TOKEN", "test-token")

	client := NewMockCloudFlareClient()

	// Create a domain filter that only matches "bar.com"
	// This should filter out "foo.com" and trigger the debug log
	domainFilter := endpoint.NewDomainFilter([]string{"bar.com"})

	p, err := NewCloudFlareProvider(
		domainFilter,
		provider.NewZoneIDFilter([]string{""}), // empty zone ID filter so it uses ListZones path
		false,                                  // proxied
		false,                                  // dry run
		RegionalServicesConfig{},
		CustomHostnamesConfig{},
		DNSRecordsConfig{PerPage: 50},
	)
	require.NoError(t, err)

	// Replace the real client with our mock
	p.Client = client

	// Capture debug logs to verify the filter log message
	oldLevel := log.GetLevel()
	log.SetLevel(log.DebugLevel)
	defer log.SetLevel(oldLevel)

	// Use a custom formatter to capture log output
	var logOutput strings.Builder
	log.SetOutput(&logOutput)
	defer log.SetOutput(os.Stderr)

	// Call Zones() which should trigger the domain filter logic
	zones, err := p.Zones(context.Background())
	require.NoError(t, err)

	// Should only return the "bar.com" zone since "foo.com" is filtered out
	assert.Len(t, zones, 1)
	assert.Equal(t, "bar.com", zones[0].Name)
	assert.Equal(t, "001", zones[0].ID)

	// Verify that the debug log was written for the filtered zone
	logString := logOutput.String()
	assert.Contains(t, logString, `zone \"foo.com\" not in domain filter`)
	assert.Contains(t, logString, "no zoneIDFilter configured, looking at all zones")
}

func TestZoneIDByNameIteratorError(t *testing.T) {
	client := NewMockCloudFlareClient()

	// Set up an error that will be returned by the ListZones iterator (line 144)
	client.listZonesError = fmt.Errorf("CloudFlare API connection timeout")

	// Call ZoneIDByName which should hit line 144 (iterator error handling)
	zoneID, err := client.ZoneIDByName("example.com")

	// Should return empty zone ID and the wrapped iterator error
	assert.Empty(t, zoneID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list zones from CloudFlare API")
	assert.Contains(t, err.Error(), "CloudFlare API connection timeout")
}

func TestZoneIDByNameZoneNotFound(t *testing.T) {
	client := NewMockCloudFlareClient()

	// Set up mock to return different zones but not the one we're looking for
	client.Zones = map[string]string{
		"zone456": "different.com",
		"zone789": "another.com",
	}

	// Call ZoneIDByName for a zone that doesn't exist, should hit line 147 (zone not found)
	zoneID, err := client.ZoneIDByName("nonexistent.com")

	// Should return empty zone ID and the improved error message
	assert.Empty(t, zoneID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), `zone "nonexistent.com" not found in CloudFlare account`)
	assert.Contains(t, err.Error(), "verify the zone exists and API credentials have access to it")
}

func TestDnsRecordFromLegacyAPI(t *testing.T) {
	parseTime := func(s string) time.Time {
		parsed, err := time.Parse(time.RFC3339, s)
		if err != nil {
			t.Fatal("failed to parse time:", err)
		}
		return parsed
	}

	tests := []struct {
		name   string
		input  cloudflarev0.DNSRecord
		expect dns.RecordResponse
	}{
		{
			name: "All fields set",
			input: cloudflarev0.DNSRecord{
				CreatedOn:  parseTime("2024-06-01T12:00:00Z"),
				ModifiedOn: parseTime("2024-06-02T12:00:00Z"),
				Type:       "A",
				Name:       "example.com",
				Content:    "1.2.3.4",
				Meta:       map[string]any{"foo": "bar"},
				Data:       map[string]any{"baz": "qux"},
				ID:         "record-id",
				Priority:   testutils.ToPtr(uint16(10)),
				TTL:        120,
				Proxied:    testutils.ToPtr(true),
				Proxiable:  true,
				Comment:    "test comment",
				Tags:       []string{"tag1", "tag2"},
			},
			expect: dns.RecordResponse{
				CreatedOn:  parseTime("2024-06-01T12:00:00Z"),
				ModifiedOn: parseTime("2024-06-02T12:00:00Z"),
				Type:       "A",
				Name:       "example.com",
				Content:    "1.2.3.4",
				Meta:       map[string]any{"foo": "bar"},
				Data:       map[string]any{"baz": "qux"},
				ID:         "record-id",
				Priority:   10,
				TTL:        120,
				Proxied:    true,
				Proxiable:  true,
				Comment:    "test comment",
				Tags:       []string{"tag1", "tag2"},
			},
		},
		{
			name: "Nil priority and proxied",
			input: cloudflarev0.DNSRecord{
				Type:      "TXT",
				Name:      "txt.example.com",
				Content:   "some text",
				Priority:  nil,
				TTL:       300,
				Proxied:   nil,
				Proxiable: false,
				Tags:      []string(nil),
			},
			expect: dns.RecordResponse{
				Type:      "TXT",
				Name:      "txt.example.com",
				Content:   "some text",
				Priority:  0,
				TTL:       300,
				Proxied:   false,
				Proxiable: false,
				Tags:      []string(nil),
			},
		},
		{
			name: "Proxied false",
			input: cloudflarev0.DNSRecord{
				Type:     "CNAME",
				Name:     "cname.example.com",
				Content:  "target.example.com",
				Proxied:  testutils.ToPtr(false),
				TTL:      60,
				Priority: nil,
				Tags:     []string(nil),
			},
			expect: dns.RecordResponse{
				Type:     "CNAME",
				Name:     "cname.example.com",
				Content:  "target.example.com",
				Proxied:  false,
				TTL:      60,
				Priority: 0,
				Tags:     []string(nil),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dnsRecordResponseFromLegacyDNSRecord(tt.input)
			assert.Equal(t, tt.expect, got)
		})
	}
}
