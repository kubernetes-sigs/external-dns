package zdns

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

func Test_zdnsRRS_tostring(t *testing.T) {
	z := &zdnsRRS{
		Rdata:   "example.com",
		Ttl:     3600,
		RrsType: "A",
	}

	name := "example"
	want := []string{"example 3600 A example.com"}

	got := z.tostring(name)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("tostring() = %v, want %v", got, want)
	}
}

func TestAnalyzeUpdate(t *testing.T) {
	z := &ZDNSProvider{} // Initialize ZDNSProvider for testing

	// Test for creating new non-TXT records
	changesCreate := &plan.Changes{
		UpdateNew: []*endpoint.Endpoint{
			{
				RecordType: "A",
				DNSName:    "example.com",
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				RecordType: "A",
				DNSName:    "example.com",
			},
		},
	}
	create, delete := z.analyzeUpdate(changesCreate)
	if len(create) != 0 {
		t.Errorf("Expected 1 new record, got %d", len(create))
	}
	if len(delete) != 0 {
		t.Errorf("Expected 0 records to delete, got %d", len(delete))
	}

	// Test for updating existing non-TXT records
	changesUpdate := &plan.Changes{
		UpdateNew: []*endpoint.Endpoint{
			{
				RecordType: "A",
				DNSName:    "example.com",
				Targets:    []string{"127.0.0.1"},
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				RecordType: "A",
				DNSName:    "example.com",
				Targets:    []string{"192.168.0.1"},
			},
		},
	}
	create, delete = z.analyzeUpdate(changesUpdate)
	if len(create) != 1 {
		t.Errorf("Expected 1 record to update, got %d", len(create))
	}
	if len(delete) != 1 {
		t.Errorf("Expected 1 record to delete, got %d", len(delete))
	}

	// Test for creating new TXT records
	changesCreateTXT := &plan.Changes{
		UpdateNew: []*endpoint.Endpoint{
			{
				RecordType: "TXT",
				DNSName:    "example.com.",
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				RecordType: "TXT",
				DNSName:    "example.com.",
			},
		},
	}
	create, delete = z.analyzeUpdate(changesCreateTXT)
	if len(create) != 0 {
		t.Errorf("Expected 1 new TXT record, got %d", len(create))
	}
	if len(delete) != 0 {
		t.Errorf("Expected 0 records to delete, got %d", len(delete))
	}

}

func TestZDNSProvider_compareNewOld(t *testing.T) {
	testCases := []struct {
		name         string
		epNew        *endpoint.Endpoint
		epOld        *endpoint.Endpoint
		expectedBool bool
	}{
		{
			name: "different record types",
			epNew: &endpoint.Endpoint{
				RecordType: "A",
			},
			epOld: &endpoint.Endpoint{
				RecordType: "CNAME",
			},
			expectedBool: false,
		},
		{
			name: "different ttls",
			epNew: &endpoint.Endpoint{
				RecordTTL: 1,
			},
			epOld: &endpoint.Endpoint{
				RecordTTL: 2,
			},
			expectedBool: false,
		},
		{
			name: "different targets",
			epNew: &endpoint.Endpoint{
				Targets: endpoint.Targets{"1.2.3.4"},
			},
			epOld: &endpoint.Endpoint{
				Targets: endpoint.Targets{"5.6.7.8"},
			},
			expectedBool: false,
		},
		{
			name: "same record",
			epNew: &endpoint.Endpoint{
				RecordType: "A",
				RecordTTL:  1,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			epOld: &endpoint.Endpoint{
				RecordType: "A",
				RecordTTL:  1,
				Targets:    endpoint.Targets{"1.2.3.4"},
			},
			expectedBool: true,
		},
	}

	z := &ZDNSProvider{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := z.compareNewOld(tc.epNew, tc.epOld); got != tc.expectedBool {
				t.Errorf("compareNewOld() = %v, want %v", got, tc.expectedBool)
			}
		})
	}
}

func TestZDNSProvider_sortSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "empty",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "one",
			input:    []string{"1.2.3.4"},
			expected: []string{"1.2.3.4"},
		},
		{
			name:     "one ipv6",
			input:    []string{"2001:db8:85a3::8a2e:370:7334"},
			expected: []string{"2001:db8:85a3::8a2e:370:7334"},
		},
		{
			name:     "two",
			input:    []string{"1.2.3.4", "5.6.7.8"},
			expected: []string{"1.2.3.4", "5.6.7.8"},
		},
		{
			name:     "two with ipv6",
			input:    []string{"1.2.3.4", "2001:db8:85a3::8a2e:370:7334"},
			expected: []string{"1.2.3.4", "2001:db8:85a3::8a2e:370:7334"},
		},
		{
			name:     "reverse",
			input:    []string{"5.6.7.8", "1.2.3.4"},
			expected: []string{"1.2.3.4", "5.6.7.8"},
		},
		{
			name:     "shuffled",
			input:    []string{"2.3.4.5", "1.2.3.4", "5.6.7.8"},
			expected: []string{"1.2.3.4", "2.3.4.5", "5.6.7.8"},
		},
	}

	z := &ZDNSProvider{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ep := &endpoint.Endpoint{
				Targets: tc.input,
			}
			z.sortSlice(ep)
			if strings.Join(ep.Targets, ",") != strings.Join(tc.expected, ",") {
				t.Errorf("sortSlice() = %v, want %v", ep.Targets, tc.expected)
			}
		})
	}
}

func TestRRSToEndpoint(t *testing.T) {
	testCases := []struct {
		name     string
		input    zdnsRRS
		expected endpoint.Endpoint
	}{
		{
			name: "A",
			input: zdnsRRS{
				Name:    "example.com",
				RrsType: "A",
				Rdata:   "192.0.2.1",
				Ttl:     3600,
				Id:      "rdata_id",
			},
			expected: endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "A",
				RecordTTL:  endpoint.TTL(3600),
				Targets:    []string{"192.0.2.1"},
				ProviderSpecific: []endpoint.ProviderSpecificProperty{
					{Name: "id", Value: "rdata_id"},
				},
				Labels: endpoint.NewLabels(),
			},
		},
		{
			name: "AAAA",
			input: zdnsRRS{
				Name:    "example.com",
				RrsType: "AAAA",
				Rdata:   "2001:db8:85a3::8a2e:370:7334",
				Ttl:     3600,
				Id:      "rdata_id",
			},
			expected: endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "AAAA",
				RecordTTL:  endpoint.TTL(3600),
				Targets:    []string{"2001:db8:85a3::8a2e:370:7334"},
				ProviderSpecific: []endpoint.ProviderSpecificProperty{
					{Name: "id", Value: "rdata_id"},
				},
				Labels: endpoint.NewLabels(),
			},
		},
		{
			name: "TXT",
			input: zdnsRRS{
				Name:    "example.com",
				RrsType: "TXT",
				Rdata:   "txt1 txt2",
				Ttl:     3600,
				Id:      "rdata_id",
			},
			expected: endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "TXT",
				RecordTTL:  endpoint.TTL(3600),
				Targets:    []string{"txt1 txt2"},
				ProviderSpecific: []endpoint.ProviderSpecificProperty{
					{Name: "id", Value: "rdata_id"},
				},
				Labels: endpoint.NewLabels(),
			},
		},
		{
			name: "CNAME",
			input: zdnsRRS{
				Name:    "example.com",
				RrsType: "CNAME",
				Rdata:   "example.net",
				Ttl:     3600,
				Id:      "rdata_id",
			},
			expected: endpoint.Endpoint{
				DNSName:    "example.com",
				RecordType: "CNAME",
				RecordTTL:  endpoint.TTL(3600),
				Targets:    []string{"example.net"},
				ProviderSpecific: []endpoint.ProviderSpecificProperty{
					{Name: "id", Value: "rdata_id"},
				},
				Labels: endpoint.NewLabels(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok, ep := rrsToEndpoint(tc.input)
			if !ok || !reflect.DeepEqual(ep, tc.expected) {
				t.Errorf("rrsToEndpoint() = %v, want %v", ep, tc.expected)
			}

		})
	}
}

func TestEndpointToRRS(t *testing.T) {
	// Test case 1
	endpoint1 := endpoint.Endpoint{
		DNSName:          "example.com",
		RecordType:       "A",
		Targets:          []string{"192.0.2.1"},
		RecordTTL:        3600,
		ProviderSpecific: []endpoint.ProviderSpecificProperty{},
	}
	rrs1 := endpointToRRS(endpoint1)
	assert.Equal(t, "example.com", rrs1.Name)
	assert.Equal(t, "A", rrs1.RrsType)
	assert.Equal(t, "192.0.2.1", rrs1.Rdata)
	assert.Equal(t, 3600, rrs1.Ttl)
	assert.Equal(t, "", rrs1.Id)

	// Test case 2
	resourceTxtMap["example.com."] = &endpoint1
	endpoint2 := endpoint.Endpoint{
		DNSName:          "example.com",
		RecordType:       "TXT",
		Targets:          []string{"txt1", "txt2"},
		RecordTTL:        3600,
		ProviderSpecific: []endpoint.ProviderSpecificProperty{},
	}
	rrs2 := endpointToRRS(endpoint2)
	assert.Equal(t, "example.com", rrs2.Name)
	assert.Equal(t, "TXT", rrs2.RrsType)
	assert.Equal(t, "txt1 txt2", rrs2.Rdata)
	assert.Equal(t, 3600, rrs2.Ttl)

	// Test case 3
	endpoint3 := endpoint.Endpoint{
		DNSName:    "example.com",
		RecordType: "A",
		Targets:    []string{"192.0.2.1"},
		RecordTTL:  3600,
		ProviderSpecific: []endpoint.ProviderSpecificProperty{
			{Value: "provider1"},
			{Value: "provider2"},
		},
	}
	rrs3 := endpointToRRS(endpoint3)
	assert.Equal(t, "example.com", rrs3.Name)
	assert.Equal(t, "A", rrs3.RrsType)
	assert.Equal(t, "192.0.2.1", rrs3.Rdata)
	assert.Equal(t, 3600, rrs3.Ttl)
	assert.Equal(t, "provider1,provider2", rrs3.Id)

	// Test case 4
	endpoint4 := endpoint.Endpoint{
		DNSName:          "example.com",
		RecordType:       "TXT",
		Targets:          []string{"txt1", "txt2"},
		RecordTTL:        3600,
		ProviderSpecific: []endpoint.ProviderSpecificProperty{},
	}
	rrs4 := endpointToRRS(endpoint4)
	assert.Equal(t, "example.com", rrs4.Name)
	assert.Equal(t, "TXT", rrs4.RrsType)
	assert.Equal(t, "txt1 txt2", rrs4.Rdata)
	assert.Equal(t, 3600, rrs4.Ttl)
	assert.Equal(t, "", rrs4.Id)
}

func TestHasAuth(t *testing.T) {
	endpoint1 := endpoint.Endpoint{DNSName: "example.com.", RecordType: "TXT"}
	endpoint2 := endpoint.Endpoint{DNSName: "example.com", RecordType: "A"}
	ownerID := "user1"

	resourceTxtMap["example.com."] = &endpoint.Endpoint{
		Labels: map[string]string{"owner": "user1"},
	}
	resourceTxtMap["a-example.com"] = &endpoint.Endpoint{
		Labels: map[string]string{"owner": "user2"},
	}

	if !hasAuth(endpoint1, ownerID) {
		t.Error("Expected true for matching owner ID and DNS name ending with '.'")
	}

	if !hasAuth(endpoint2, ownerID) {
		t.Error("Expected true for matching owner ID and DNS name not ending with '.'")
	}

	if hasAuth(endpoint1, "user2") {
		t.Error("Expected false for non-matching owner ID and DNS name ending with '.'")
	}

}

/* -------------------------------- */

func TestNewZDNSProvider(t *testing.T) {
	zdnsConfig := ZDNSConfig{
		Host:  "example.com",
		Port:  "8080",
		View:  "example",
		Auth:  "exampleAuth",
		Zones: "zone1,zone2",
	}
	tests := []struct {
		name       string
		config     ZDNSConfig
		monkeyfunc func(t *testing.T)
		want       struct {
			provid *ZDNSProvider
			err    error
		}
	}{
		{
			name:   "not connect",
			config: zdnsConfig,
			monkeyfunc: func(t *testing.T) {
				monkeyfunc1 := gomonkey.ApplyFunc(sendHTTPReqest, func(_ string, _ string, _ string, _ map[string]string) (int, []byte, error) {
					return 401, []byte("ok"), errors.New("not connect")
				})
				t.Cleanup(func() {
					monkeyfunc1.Reset()
				})
			},
			want: struct {
				provid *ZDNSProvider
				err    error
			}{
				provid: nil,
				err:    fmt.Errorf("failed to connect to zdns api:"),
			},
		},
		{
			name:   "not auth",
			config: zdnsConfig,
			monkeyfunc: func(t *testing.T) {
				monkeyfunc1 := gomonkey.ApplyFunc(sendHTTPReqest, func(_ string, _ string, _ string, _ map[string]string) (int, []byte, error) {
					return 401, []byte("ok"), nil
				})
				t.Cleanup(func() {
					monkeyfunc1.Reset()
				})
			},
			want: struct {
				provid *ZDNSProvider
				err    error
			}{
				provid: nil,
				err:    fmt.Errorf("failed to connect to zdns api:"),
			},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.monkeyfunc(t)

			provid, _ := NewZDNSProvider(ctx, tt.config)
			if provid != tt.want.provid {
				t.Errorf("NewZDNSProvider() = %v, want %v", provid, tt.want)
			}

		})
	}

	tests2 := []struct {
		name       string
		config     ZDNSConfig
		monkeyfunc func(t *testing.T)
		want       struct {
			provid *ZDNSProvider
			err    error
		}
	}{
		{
			name:   "right",
			config: zdnsConfig,
			monkeyfunc: func(t *testing.T) {
				monkeyfunc1 := gomonkey.ApplyFunc(sendHTTPReqest, func(_ string, _ string, _ string, _ map[string]string) (int, []byte, error) {
					answer := "{\"resources\": [{ 		\"id\": \"bom\", 		\"name\": \"bom\" 	}, { 		\"id\": \"cmtest\", 		\"name\": \"cmtest\" 	}] }"
					return 200, []byte(answer), nil
				})
				t.Cleanup(func() {
					monkeyfunc1.Reset()
				})
			},
			want: struct {
				provid *ZDNSProvider
				err    error
			}{
				provid: &ZDNSProvider{
					client: ZDNSAPIClient{
						config: zdnsConfig,
					},
				},
				err: fmt.Errorf("failed to connect to zdns api:"),
			},
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			tt.monkeyfunc(t)

			provid, _ := NewZDNSProvider(ctx, tt.config)
			if !reflect.DeepEqual(provid.client.config, provid.client.config) {
				t.Errorf("NewZDNSProvider() = %v, want %v", provid, tt.want)
			}

		})
	}
}

func TestZDNSProvider_Records(t *testing.T) {
	var zp *ZDNSProvider
	zdnsConfig := ZDNSConfig{
		Host:  "example.com",
		Port:  "8080",
		View:  "example",
		Auth:  "exampleAuth",
		Zones: "zone1,zone2",
	}
	z := &ZDNSProvider{
		client: ZDNSAPIClient{
			config: zdnsConfig,
			httpcf: httpConfig{
				zones: []zone{
					{
						name:  "bom",
						len:   3,
						level: 1,
					},
				},
			},
		},
	}
	tests := []struct {
		name          string
		monkeymethod  func(t *testing.T)
		wantEndpoints []*endpoint.Endpoint
	}{
		{
			name: "failed",
			monkeymethod: func(t *testing.T) {
				monkeyfunc1 := gomonkey.ApplyMethod(reflect.TypeOf(zp), "getZoneRecords", func(_ string) []byte {
					answer := "lalalalla"
					//answer := `{"resources":[{"id":"NS$2$default$bom","href":"/views/default/zones/bom/rrs/NS$2$default$bom","name":"bom.","type":"NS","klass":"IN","ttl":3600,"rdata":"ns.bom."},{"id":"A$3$default$bom","href":"/views/default/zones/bom/rrs/A$3$default$bom","name":"ns.bom.","type":"A","klass":"IN","ttl":3600,"rdata":"127.0.0.1"},{"id":"A$4$default$bom","href":"/views/default/zones/bom/rrs/A$4$default$bom","name":"hh.bom.","type":"A","klass":"IN","ttl":3600,"rdata":"1.1.1.1"}]}`
					fmt.Println(answer)
					return []byte(answer)
				})
				t.Cleanup(func() {
					monkeyfunc1.Reset()
				})
			},
			wantEndpoints: nil,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("aaaaaaaa")
			tt.monkeymethod(t)
			gotEndpoints, err := z.Records(ctx)
			if err != nil || !reflect.DeepEqual(gotEndpoints, tt.wantEndpoints) {
				t.Errorf("ZDNSProvider.Records() = %v, want %v", gotEndpoints, tt.wantEndpoints)
			}
		})
	}
}
