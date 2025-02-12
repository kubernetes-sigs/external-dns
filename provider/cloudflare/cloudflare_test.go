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
	"sort"
	"strings"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/stretchr/testify/assert"

	"github.com/maxatome/go-testdeep/td"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type MockAction struct {
	Name       string
	ZoneId     string
	RecordId   string
	RecordData cloudflare.DNSRecord
}

type mockCloudFlareClient struct {
	User                  cloudflare.User
	Zones                 map[string]string
	Records               map[string]map[string]cloudflare.DNSRecord
	Actions               []MockAction
	listZonesError        error
	listZonesContextError error
	dnsRecordsError       error
	customHostnames       map[string]map[string]cloudflare.CustomHostname
}

var ExampleDomain = []cloudflare.DNSRecord{
	{
		ID:      "1234567890",
		Name:    "foobar.bar.com",
		Type:    endpoint.RecordTypeA,
		TTL:     120,
		Content: "1.2.3.4",
		Proxied: proxyDisabled,
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
		User: cloudflare.User{ID: "xxxxxxxxxxxxxxxxxxx"},
		Zones: map[string]string{
			"001": "bar.com",
			"002": "foo.com",
		},
		Records: map[string]map[string]cloudflare.DNSRecord{
			"001": {},
			"002": {},
		},
		customHostnames: map[string]map[string]cloudflare.CustomHostname{},
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
		return cloudflare.DNSRecord{
			Name:    params.Name,
			TTL:     params.TTL,
			Proxied: params.Proxied,
			Type:    params.Type,
			Content: params.Content,
		}
	case cloudflare.UpdateDNSRecordParams:
		return cloudflare.DNSRecord{
			Name:    params.Name,
			TTL:     params.TTL,
			Proxied: params.Proxied,
			Type:    params.Type,
			Content: params.Content,
		}
	default:
		return cloudflare.DNSRecord{}
	}
}

func (m *mockCloudFlareClient) CreateDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.CreateDNSRecordParams) (cloudflare.DNSRecord, error) {
	recordData := getDNSRecordFromRecordParams(rp)
	m.Actions = append(m.Actions, MockAction{
		Name:       "Create",
		ZoneId:     rc.Identifier,
		RecordId:   rp.ID,
		RecordData: recordData,
	})
	if zone, ok := m.Records[rc.Identifier]; ok {
		zone[rp.ID] = recordData
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
			zone[rp.ID] = recordData
		}
	}
	return nil
}

func (m *mockCloudFlareClient) UpdateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.UpdateDataLocalizationRegionalHostnameParams) error {
	m.Actions = append(m.Actions, MockAction{
		Name:     "UpdateDataLocalizationRegionalHostname",
		ZoneId:   rc.Identifier,
		RecordId: "",
		RecordData: cloudflare.DNSRecord{
			Name: rp.Hostname,
		},
	})
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
			delete(zone, recordID)
			return nil
		}
	}
	return nil
}

func (m *mockCloudFlareClient) UserDetails(ctx context.Context) (cloudflare.User, error) {
	return m.User, nil
}

func (m *mockCloudFlareClient) CustomHostnames(ctx context.Context, zoneID string, page int, filter cloudflare.CustomHostname) ([]cloudflare.CustomHostname, cloudflare.ResultInfo, error) {
	var err error = nil

	if page != 1 || filter.Hostname != "" {
		err = errors.New("pages and filters are not supported for custom hostnames mock test")
	}

	result := []cloudflare.CustomHostname{}
	if zone, ok := m.customHostnames[zoneID]; ok {
		for _, ch := range zone {
			result = append(result, ch)
		}
	}

	return result,
		cloudflare.ResultInfo{
			Page:       1,
			PerPage:    100,
			Count:      len(result),
			Total:      len(result),
			TotalPages: 1,
		}, err
}

func (m *mockCloudFlareClient) CreateCustomHostname(ctx context.Context, zoneID string, ch cloudflare.CustomHostname) (*cloudflare.CustomHostnameResponse, error) {
	if _, ok := m.customHostnames[zoneID]; !ok {
		m.customHostnames[zoneID] = map[string]cloudflare.CustomHostname{}
	}
	m.customHostnames[zoneID][ch.ID] = ch
	return &cloudflare.CustomHostnameResponse{}, nil
}

func (m *mockCloudFlareClient) DeleteCustomHostname(ctx context.Context, zoneID string, customHostnameID string) error {
	if zone, ok := m.customHostnames[zoneID]; ok {
		if _, ok := zone[customHostnameID]; ok {
			delete(zone, customHostnameID)
			return nil
		}
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

func (m *mockCloudFlareClient) ListZones(ctx context.Context, zoneID ...string) ([]cloudflare.Zone, error) {
	if m.listZonesError != nil {
		return nil, m.listZonesError
	}

	result := []cloudflare.Zone{}

	for zoneID, zoneName := range m.Zones {
		result = append(result, cloudflare.Zone{
			ID:   zoneID,
			Name: zoneName,
		})
	}

	return result, nil
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
	for id, zoneName := range m.Zones {
		if zoneID == id {
			return cloudflare.Zone{
				ID:   zoneID,
				Name: zoneName,
			}, nil
		}
	}

	return cloudflare.Zone{}, errors.New("Unknown zoneID: " + zoneID)
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
		DomainFilter:   endpoint.MatchAllDomainFilters{&domainFilter},
		ManagedRecords: managedRecords,
	}

	changes := plan.Calculate().Changes

	// Records other than A, CNAME and NS are not supported by planner, just create them
	for _, endpoint := range endpoints {
		if endpoint.RecordType != "A" && endpoint.RecordType != "CNAME" && endpoint.RecordType != "NS" {
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
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
				Type:    "A",
				Name:    "bar.com",
				Content: "127.0.0.1",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
		{
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
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
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
				Type:    "CNAME",
				Name:    "cname.bar.com",
				Content: "google.com",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
		{
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
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
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
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
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
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
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
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
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
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
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
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
		endpoints := []*endpoint.Endpoint{
			{
				RecordType: testCase.recordType,
				DNSName:    testCase.domain,
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
				Name:   "Create",
				ZoneId: "001",
				RecordData: cloudflare.DNSRecord{
					Type:    testCase.recordType,
					Name:    testCase.domain,
					Content: "127.0.0.1",
					TTL:     1,
					Proxied: testCase.proxiable,
				},
			},
		}, []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME, endpoint.RecordTypeNS}, testCase.recordType+" record on "+testCase.domain)
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

	assert.Equal(t, 1, len(zones))
	assert.Equal(t, "bar.com", zones[0].Name)
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
	assert.Equal(t, 1, len(zones))
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
		Client:            client,
		DNSRecordsPerPage: 1,
	}
	ctx := context.Background()

	records, err := p.Records(ctx)
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}
	assert.Equal(t, 2, len(records))
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
	_ = os.Setenv("CF_API_TOKEN", "abc123def")
	_, err := NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"bar.com"}),
		provider.NewZoneIDFilter([]string{""}),
		false,
		true,
		5000,
		"")
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	_ = os.Unsetenv("CF_API_TOKEN")
	tokenFile := "/tmp/cf_api_token"
	if err := os.WriteFile(tokenFile, []byte("abc123def"), 0o644); err != nil {
		t.Errorf("failed to write token file, %s", err)
	}
	_ = os.Setenv("CF_API_TOKEN", tokenFile)
	_, err = NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"bar.com"}),
		provider.NewZoneIDFilter([]string{""}),
		false,
		true,
		5000,
		"")
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	_ = os.Unsetenv("CF_API_TOKEN")
	_ = os.Setenv("CF_API_KEY", "xxxxxxxxxxxxxxxxx")
	_ = os.Setenv("CF_API_EMAIL", "test@test.com")
	_, err = NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"bar.com"}),
		provider.NewZoneIDFilter([]string{""}),
		false,
		true,
		5000,
		"")
	if err != nil {
		t.Errorf("should not fail, %s", err)
	}

	_ = os.Unsetenv("CF_API_KEY")
	_ = os.Unsetenv("CF_API_EMAIL")
	_, err = NewCloudFlareProvider(
		endpoint.NewDomainFilter([]string{"bar.com"}),
		provider.NewZoneIDFilter([]string{""}),
		false,
		true,
		5000,
		"")
	if err == nil {
		t.Errorf("expected to fail")
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
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
				Name:    "new.bar.com",
				Content: "target",
				TTL:     1,
				Proxied: proxyDisabled,
			},
		},
		{
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
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
	records := []cloudflare.DNSRecord{
		{
			Name:    "foo.com",
			Type:    endpoint.RecordTypeCNAME,
			Content: "foobar",
			ID:      "1",
		},
		{
			Name: "bar.de",
			Type: endpoint.RecordTypeA,
			ID:   "2",
		},
		{
			Name:    "bar.de",
			Type:    endpoint.RecordTypeA,
			Content: "1.2.3.4",
			ID:      "2",
		},
	}

	assert.Equal(t, "", p.getRecordID(records, cloudflare.DNSRecord{
		Name:    "foo.com",
		Type:    endpoint.RecordTypeA,
		Content: "foobar",
	}))

	assert.Equal(t, "", p.getRecordID(records, cloudflare.DNSRecord{
		Name:    "foo.com",
		Type:    endpoint.RecordTypeCNAME,
		Content: "fizfuz",
	}))

	assert.Equal(t, "1", p.getRecordID(records, cloudflare.DNSRecord{
		Name:    "foo.com",
		Type:    endpoint.RecordTypeCNAME,
		Content: "foobar",
	}))
	assert.Equal(t, "", p.getRecordID(records, cloudflare.DNSRecord{
		Name:    "bar.de",
		Type:    endpoint.RecordTypeA,
		Content: "2.3.4.5",
	}))
	assert.Equal(t, "2", p.getRecordID(records, cloudflare.DNSRecord{
		Name:    "bar.de",
		Type:    endpoint.RecordTypeA,
		Content: "1.2.3.4",
	}))
}

func TestCloudflareGroupByNameAndType(t *testing.T) {
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
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
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
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
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
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
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
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
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
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "bar.de",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
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
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
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
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "foo.com",
					Type:    endpoint.RecordTypeA,
					Content: "10.10.10.2",
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
				{
					Name:    "bar.de",
					Type:    "NOT SUPPORTED",
					Content: "10.10.10.1",
					TTL:     defaultCloudFlareRecordTTL,
					Proxied: proxyDisabled,
				},
			},
			ExpectedEndpoints: []*endpoint.Endpoint{
				{
					DNSName:    "foo.com",
					Targets:    endpoint.Targets{"10.10.10.1", "10.10.10.2"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
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
		assert.ElementsMatch(t, groupByNameAndType(tc.Records), tc.ExpectedEndpoints)
	}
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
			assert.Equal(t, 1, len(current))

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
			assert.Equal(t, 0, len(plan.Changes.Create), "should not have creates")
			assert.Equal(t, 0, len(plan.Changes.Delete), "should not have deletes")

			if test.ShouldBeUpdated {
				assert.Equal(t, 1, len(plan.Changes.UpdateNew), "should not have new updates")
				assert.Equal(t, 1, len(plan.Changes.UpdateOld), "should not have old updates")
			} else {
				assert.Equal(t, 0, len(plan.Changes.UpdateNew), "should not have new updates")
				assert.Equal(t, 0, len(plan.Changes.UpdateOld), "should not have old updates")
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
			RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
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
		DomainFilter:   endpoint.MatchAllDomainFilters{&domainFilter},
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
			Name:   "Create",
			ZoneId: "001",
			RecordData: cloudflare.DNSRecord{
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
		DomainFilter:   endpoint.MatchAllDomainFilters{&domainFilter},
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
	}

	planned := plan.Calculate()

	assert.Equal(t, 0, len(planned.Changes.Create), "no new changes should be here")
	assert.Equal(t, 0, len(planned.Changes.UpdateNew), "no new changes should be here")
	assert.Equal(t, 0, len(planned.Changes.UpdateOld), "no new changes should be here")
	assert.Equal(t, 0, len(planned.Changes.Delete), "no new changes should be here")
}

func TestCloudFlareProvider_Region(t *testing.T) {
	_ = os.Setenv("CF_API_TOKEN", "abc123def")
	_ = os.Setenv("CF_API_EMAIL", "test@test.com")
	provider, err := NewCloudFlareProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.ZoneIDFilter{}, true, false, 50, "us")
	if err != nil {
		t.Fatal(err)
	}

	if provider.RegionKey != "us" {
		t.Errorf("expected region key to be 'us', but got '%s'", provider.RegionKey)
	}
}

func TestCloudFlareProvider_updateDataLocalizationRegionalHostnameParams(t *testing.T) {
	change := &cloudFlareChange{
		RegionalHostname: cloudflare.RegionalHostname{
			Hostname:  "example.com",
			RegionKey: "us",
		},
	}

	params := updateDataLocalizationRegionalHostnameParams(*change)
	if params.Hostname != "example.com" {
		t.Errorf("expected hostname to be 'example.com', but got '%s'", params.Hostname)
	}

	if params.RegionKey != "us" {
		t.Errorf("expected region key to be 'us', but got '%s'", params.RegionKey)
	}
}

func TestCloudFlareProvider_newCloudFlareChange(t *testing.T) {
	_ = os.Setenv("CF_API_KEY", "xxxxxxxxxxxxxxxxx")
	_ = os.Setenv("CF_API_EMAIL", "test@test.com")
	provider, err := NewCloudFlareProvider(endpoint.NewDomainFilter([]string{"example.com"}), provider.ZoneIDFilter{}, true, false, 50, "us")
	if err != nil {
		t.Fatal(err)
	}

	endpoint := &endpoint.Endpoint{
		DNSName: "example.com",
		Targets: []string{"192.0.2.1"},
	}

	change := provider.newCloudFlareChange(cloudFlareCreate, endpoint, endpoint.Targets[0])
	if change.RegionalHostname.RegionKey != "us" {
		t.Errorf("expected region key to be 'us', but got '%s'", change.RegionalHostname.RegionKey)
	}
}

func TestCloudflareCreateCustomHostname(t *testing.T) {
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

	testCases := []struct {
		Name                    string
		Endpoints               []*endpoint.Endpoint
		ExpectedCustomHostnames map[string]string
	}{
		{
			Name: "add custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "a.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4", "2.3.4.5"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "a.foo.fancybar.com",
						},
					},
				},
				{
					DNSName:    "txt.foo.bar.com",
					Targets:    endpoint.Targets{"value"},
					RecordType: endpoint.RecordTypeTXT,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "txt.foo.fancybar.com",
						},
					},
				},
			},
			ExpectedCustomHostnames: map[string]string{
				"a.foo.bar.com":   "a.foo.fancybar.com",
				"txt.foo.bar.com": "",
			},
		},
		{
			Name: "update custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "a.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4", "2.3.4.5"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
					ProviderSpecific: endpoint.ProviderSpecific{
						{
							Name:  "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname",
							Value: "a2.foo.fancybar.com",
						},
					},
				},
			},
			ExpectedCustomHostnames: map[string]string{"a.foo.bar.com": "a2.foo.fancybar.com"},
		},
		{
			Name: "delete custom hostname",
			Endpoints: []*endpoint.Endpoint{
				{
					DNSName:    "a.foo.bar.com",
					Targets:    endpoint.Targets{"1.2.3.4", "2.3.4.5"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  endpoint.TTL(defaultCloudFlareRecordTTL),
					Labels:     endpoint.Labels{},
				},
			},
			ExpectedCustomHostnames: map[string]string{"a.foo.bar.com": ""},
		},
	}

	for _, tc := range testCases {
		endpoints, err := provider.AdjustEndpoints(tc.Endpoints)

		assert.NoError(t, err)
		plan := &plan.Plan{
			Current:        records,
			Desired:        endpoints,
			DomainFilter:   endpoint.MatchAllDomainFilters{&domainFilter},
			ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeCNAME},
		}

		planned := plan.Calculate()

		err = provider.ApplyChanges(context.Background(), planned.Changes)
		if err != nil {
			t.Errorf("should not fail, %s", err)
		}

		chs, chErr := provider.listCustomHostnamesWithPagination(ctx, "001")
		if chErr != nil {
			t.Errorf("should not fail, %s", chErr)
		}

		for k, v := range tc.ExpectedCustomHostnames {
			_, ch := provider.getCustomHostnameIDbyOrigin(chs, k)
			assert.Equal(t, v, ch)
		}
	}
}
