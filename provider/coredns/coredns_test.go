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

package coredns

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.etcd.io/etcd/api/v3/mvccpb"
	etcdcv3 "go.etcd.io/etcd/client/v3"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/plan"

	"github.com/stretchr/testify/require"
)

const defaultCoreDNSPrefix = "/skydns/"

type fakeETCDClient struct {
	services map[string]Service
}

func (c fakeETCDClient) GetServices(prefix string) ([]*Service, error) {
	var result []*Service
	for key, value := range c.services {
		if strings.HasPrefix(key, prefix) {
			valueCopy := value
			valueCopy.Key = key
			result = append(result, &valueCopy)
		}
	}
	return result, nil
}

func (c fakeETCDClient) SaveService(service *Service) error {
	c.services[service.Key] = *service
	return nil
}

func (c fakeETCDClient) DeleteService(key string) error {
	delete(c.services, key)
	return nil
}

type MockEtcdKV struct {
	etcdcv3.KV
	mock.Mock
}

func (m *MockEtcdKV) Put(ctx context.Context, key, input string, _ ...etcdcv3.OpOption) (*etcdcv3.PutResponse, error) {
	args := m.Called(ctx, key, input)
	return args.Get(0).(*etcdcv3.PutResponse), args.Error(1)
}

func (m *MockEtcdKV) Get(ctx context.Context, key string, _ ...etcdcv3.OpOption) (*etcdcv3.GetResponse, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(*etcdcv3.GetResponse), args.Error(1)
}

func (m *MockEtcdKV) Delete(ctx context.Context, key string, opts ...etcdcv3.OpOption) (*etcdcv3.DeleteResponse, error) {
	args := m.Called(ctx, key, opts[0])
	return args.Get(0).(*etcdcv3.DeleteResponse), args.Error(1)
}

func TestETCDConfig(t *testing.T) {
	var tests = []struct {
		name  string
		input map[string]string
		want  *etcdcv3.Config
	}{
		{
			"default config",
			map[string]string{},
			&etcdcv3.Config{Endpoints: []string{"http://localhost:2379"}},
		},
		{
			"config with ETCD_URLS",
			map[string]string{"ETCD_URLS": "http://example.com:2379"},
			&etcdcv3.Config{Endpoints: []string{"http://example.com:2379"}},
		},
		{
			"config with ETCD_USERNAME and ETCD_PASSWORD",
			map[string]string{"ETCD_USERNAME": "root", "ETCD_PASSWORD": "test"},
			&etcdcv3.Config{
				Endpoints: []string{"http://localhost:2379"},
				Username:  "root",
				Password:  "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.TestHelperEnvSetter(t, tt.input)
			cfg, _ := getETCDConfig()
			if !reflect.DeepEqual(cfg, tt.want) {
				t.Errorf("unexpected config. Got %v, want %v", cfg, tt.want)
			}
		})
	}
}

func TestEtcdHttpsProtocol(t *testing.T) {
	envs := map[string]string{
		"ETCD_URLS": "https://example.com:2379",
	}
	testutils.TestHelperEnvSetter(t, envs)

	cfg, err := getETCDConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
}

func TestEtcdHttpsIncorrectConfigError(t *testing.T) {
	envs := map[string]string{
		"ETCD_URLS":     "https://example.com:2379",
		"ETCD_KEY_FILE": "incorrect-path-to-etcd-tls-key",
	}
	testutils.TestHelperEnvSetter(t, envs)

	_, err := getETCDConfig()
	assert.Errorf(t, err, "Error creating TLS config: either both cert and key or none must be provided")
}

func TestEtcdUnsupportedProtocolError(t *testing.T) {
	envs := map[string]string{
		"ETCD_URLS": "jdbc:ftp:RemoteHost=MyFTPServer",
	}
	testutils.TestHelperEnvSetter(t, envs)

	_, err := getETCDConfig()
	assert.Errorf(t, err, "etcd URLs must start with either http:// or https://")
}

func TestAServiceTranslation(t *testing.T) {
	expectedTarget := "1.2.3.4"
	expectedDNSName := "example.com"
	expectedRecordType := endpoint.RecordTypeA

	client := fakeETCDClient{
		map[string]Service{
			"/skydns/com/example": {Host: expectedTarget},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
		txtOwnerID:    "default",
	}
	endpoints, err := provider.Records(context.Background())
	require.NoError(t, err)
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

func TestCNAMEServiceTranslation(t *testing.T) {
	expectedTarget := "example.net"
	expectedDNSName := "example.com"
	expectedRecordType := endpoint.RecordTypeCNAME

	client := fakeETCDClient{
		map[string]Service{
			"/skydns/com/example": {Host: expectedTarget},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
		txtOwnerID:    "default",
	}
	endpoints, err := provider.Records(context.Background())
	require.NoError(t, err)
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

func TestTXTServiceTranslation(t *testing.T) {
	expectedTarget := "string"
	expectedDNSName := "example.com"
	expectedRecordType := endpoint.RecordTypeTXT

	client := fakeETCDClient{
		map[string]Service{
			"/skydns/com/example": {Text: expectedTarget},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
		txtOwnerID:    "default",
	}
	endpoints, err := provider.Records(context.Background())
	require.NoError(t, err)
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

func TestAWithTXTServiceTranslation(t *testing.T) {
	expectedTargets := map[string]string{
		endpoint.RecordTypeA:   "1.2.3.4",
		endpoint.RecordTypeTXT: "string",
	}
	expectedDNSName := "example.com"

	client := fakeETCDClient{
		map[string]Service{
			"/skydns/com/example": {Host: "1.2.3.4", Text: "string"},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
		txtOwnerID:    "default",
	}
	endpoints, err := provider.Records(context.Background())
	require.NoError(t, err)
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

func TestCNAMEWithTXTServiceTranslation(t *testing.T) {
	expectedTargets := map[string]string{
		endpoint.RecordTypeCNAME: "example.net",
		endpoint.RecordTypeTXT:   "string",
	}
	expectedDNSName := "example.com"

	client := fakeETCDClient{
		map[string]Service{
			"/skydns/com/example": {Host: "example.net", Text: "string"},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
		txtOwnerID:    "default",
	}
	endpoints, err := provider.Records(context.Background())
	require.NoError(t, err)
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

func TestCoreDNSApplyChanges(t *testing.T) {
	client := fakeETCDClient{
		map[string]Service{},
	}
	coredns := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
	}

	changes1 := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeA, "5.5.5.5"),
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeTXT, "string1"),
			endpoint.NewEndpoint("domain2.local", endpoint.RecordTypeCNAME, "site.local"),
		},
	}
	err := coredns.ApplyChanges(context.Background(), changes1)
	require.NoError(t, err)

	expectedServices1 := map[string][]*Service{
		"/skydns/local/domain1": {{Host: "5.5.5.5"}, {Text: "string1"}},
		"/skydns/local/domain2": {{Host: "site.local"}},
	}
	validateServices(client.services, expectedServices1, t, 1)

	changes2 := &plan.Changes{
		UpdateNew: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain1.local", "A", "6.6.6.6"),
		},
	}
	records, _ := coredns.Records(context.Background())
	for _, ep := range records {
		if ep.DNSName == "domain1.local" {
			changes2.UpdateOld = append(changes2.UpdateOld, ep)
		}
	}
	err = applyServiceChanges(coredns, changes2)
	require.NoError(t, err)

	expectedServices2 := map[string][]*Service{
		"/skydns/local/domain1": {{Host: "6.6.6.6"}, {Text: "string1"}},
		"/skydns/local/domain2": {{Host: "site.local"}},
	}
	validateServices(client.services, expectedServices2, t, 2)

	changes3 := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeA, "6.6.6.6"),
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeTXT, "string1"),
		},
	}

	err = applyServiceChanges(coredns, changes3)
	require.NoError(t, err)

	expectedServices3 := map[string][]*Service{
		"/skydns/local/domain2": {{Host: "site.local"}},
	}
	validateServices(client.services, expectedServices3, t, 3)

	// Test for multiple A records for the same FQDN
	changes4 := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeA, "5.5.5.5"),
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeA, "6.6.6.6"),
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeA, "7.7.7.7"),
		},
	}
	for _, ep := range changes4.Create {
		ep.Labels = map[string]string{
			"5.5.5.5": "pfx1",
			"6.6.6.6": "pfx2",
			"7.7.7.7": "pfx3",
		}
	}
	err = coredns.ApplyChanges(context.Background(), changes4)
	require.NoError(t, err)

	expectedServices4 := map[string][]*Service{
		"/skydns/local/domain2": {{Host: "site.local"}},
		"/skydns/local/domain1": {{Host: "5.5.5.5"}, {Host: "6.6.6.6"}, {Host: "7.7.7.7"}},
	}
	validateServices(client.services, expectedServices4, t, 4)
}

func TestCoreDNSApplyChanges_DomainDoNotMatch(t *testing.T) {
	client := fakeETCDClient{
		map[string]Service{},
	}
	coredns := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
		domainFilter:  endpoint.NewDomainFilter([]string{"example.local"}),
	}

	changes1 := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeA, "5.5.5.5"),
			endpoint.NewEndpoint("example.local", endpoint.RecordTypeTXT, "string1"),
			endpoint.NewEndpoint("domain2.local", endpoint.RecordTypeCNAME, "site.local"),
		},
	}
	hook := testutils.LogsUnderTestWithLogLevel(log.DebugLevel, t)
	err := coredns.ApplyChanges(context.Background(), changes1)
	require.NoError(t, err)

	testutils.TestHelperLogContains("Skipping record \"domain1.local\" due to domain filter", hook, t)
	testutils.TestHelperLogContains("Skipping record \"domain2.local\" due to domain filter", hook, t)
}

func applyServiceChanges(provider coreDNSProvider, changes *plan.Changes) error {
	ctx := context.Background()
	records, _ := provider.Records(ctx)
	for _, col := range [][]*endpoint.Endpoint{changes.Create, changes.UpdateNew, changes.Delete} {
		for _, record := range col {
			for _, existingRecord := range records {
				if existingRecord.DNSName == record.DNSName && existingRecord.RecordType == record.RecordType {
					mergeLabels(record, existingRecord.Labels)
				}
			}
		}
	}
	return provider.ApplyChanges(ctx, changes)
}

func validateServices(services map[string]Service, expectedServices map[string][]*Service, t *testing.T, step int) {
	t.Helper()
	for key, value := range services {
		keyParts := strings.Split(key, "/")
		expectedKey := strings.Join(keyParts[:len(keyParts)-value.TargetStrip], "/")
		if value.TargetStrip == 0 {
			expectedKey = strings.Join(keyParts[:len(keyParts)-1], "/")
		}
		expectedServiceEntries := expectedServices[expectedKey]
		if expectedServiceEntries == nil {
			//This is a TXT record for ownership tracking, so we can ignore it
			if value.Text != "" && strings.Contains(value.Text, "heritage=external-dns") {
				continue
			}
			t.Errorf("unexpected service %s", key)
			continue
		}
		found := false
		for i, expectedServiceEntry := range expectedServiceEntries {
			if value.Host == expectedServiceEntry.Host && value.Text == expectedServiceEntry.Text {
				expectedServiceEntries = append(expectedServiceEntries[:i], expectedServiceEntries[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			t.Errorf("unexpected service %s: %s on step %d", key, value.Host, step)
		}
		if len(expectedServiceEntries) == 0 {
			delete(expectedServices, expectedKey)
		} else {
			expectedServices[expectedKey] = expectedServiceEntries
		}
	}
	if len(expectedServices) != 0 {
		t.Errorf("unmatched expected services: %+v on step %d", expectedServices, step)
	}
}

// mergeLabels adds keys to labels if not defined for the endpoint
func mergeLabels(e *endpoint.Endpoint, labels map[string]string) {
	for k, v := range labels {
		if e.Labels[k] == "" {
			e.Labels[k] = v
		}
	}
}

func TestGetServices_Success(t *testing.T) {
	svc := Service{Host: "example.com", Port: 80, Priority: 1, Weight: 10, Text: "hello"}
	value, err := json.Marshal(svc)
	require.NoError(t, err)
	mockKV := new(MockEtcdKV)
	mockKV.On("Get", mock.Anything, "/prefix").Return(&etcdcv3.GetResponse{
		Kvs: []*mvccpb.KeyValue{
			{
				Key:   []byte("/prefix/1"),
				Value: value,
			},
		},
	}, nil)

	c := etcdClient{
		client: &etcdcv3.Client{
			KV: mockKV,
		},
		ctx: context.TODO(),
	}

	result, err := c.GetServices("/prefix")
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "example.com", result[0].Host)
}

func TestGetServices_Duplicate(t *testing.T) {
	mockKV := new(MockEtcdKV)
	c := etcdClient{
		client: &etcdcv3.Client{
			KV: mockKV,
		},
		ctx: context.TODO(),
	}

	svc := Service{Host: "example.com", Port: 80, Priority: 1, Weight: 10, Text: "hello"}
	value, err := json.Marshal(svc)
	require.NoError(t, err)

	mockKV.On("Get", mock.Anything, "/prefix").Return(&etcdcv3.GetResponse{
		Kvs: []*mvccpb.KeyValue{
			{
				Key:   []byte("/prefix/1"),
				Value: value,
			},
			{
				Key:   []byte("/prefix/1"),
				Value: value,
			},
		},
	}, nil)

	result, err := c.GetServices("/prefix")
	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestGetServices_Multiple(t *testing.T) {
	mockKV := new(MockEtcdKV)
	c := etcdClient{
		client: &etcdcv3.Client{
			KV: mockKV,
		},
		ctx: context.TODO(),
	}

	svc := Service{Host: "example.com", Port: 80, Priority: 1, Weight: 10, Text: "hello"}
	value, err := json.Marshal(svc)
	require.NoError(t, err)
	svc2 := Service{Host: "example.com", Port: 80, Priority: 0, Weight: 10, Text: "hello"}
	value2, err := json.Marshal(svc2)
	require.NoError(t, err)

	mockKV.On("Get", mock.Anything, "/prefix").Return(&etcdcv3.GetResponse{
		Kvs: []*mvccpb.KeyValue{
			{
				Key:   []byte("/prefix/1"),
				Value: value,
			},
			{
				Key:   []byte("/prefix/2"),
				Value: value2,
			},
		},
	}, nil)

	result, err := c.GetServices("/prefix")
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, priority, result[1].Priority)
}

func TestGetServices_UnmarshalError(t *testing.T) {
	mockKV := new(MockEtcdKV)
	c := etcdClient{
		client: &etcdcv3.Client{
			KV: mockKV,
		},
		ctx: context.TODO(),
	}

	mockKV.On("Get", mock.Anything, "/prefix").Return(&etcdcv3.GetResponse{
		Kvs: []*mvccpb.KeyValue{
			{
				Key:   []byte("/prefix/1"),
				Value: []byte("invalid-json"),
			},
			{
				Key:   []byte("/prefix/1"),
				Value: []byte("invalid-json"),
			},
		},
	}, nil)

	_, err := c.GetServices("/prefix")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "/prefix/1")
}

func TestGetServices_GetError(t *testing.T) {
	mockKV := new(MockEtcdKV)
	c := etcdClient{
		client: &etcdcv3.Client{
			KV: mockKV,
		},
		ctx: context.TODO(),
	}

	mockKV.On("Get", mock.Anything, "/prefix").Return(&etcdcv3.GetResponse{}, errors.New("etcd failure"))

	_, err := c.GetServices("/prefix")
	assert.Error(t, err)
	assert.EqualError(t, err, "etcd failure")
}

func TestDeleteService(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		mockErr error
		wantErr bool
	}{
		{
			name: "successful deletion",
			key:  "/skydns/local/test",
		},
		{
			name:    "etcd error",
			key:     "/skydns/local/test",
			mockErr: errors.New("etcd failure"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockKV := new(MockEtcdKV)
			mockKV.On("Delete", mock.Anything, mock.Anything, mock.AnythingOfType("clientv3.OpOption")).
				Return(&etcdcv3.DeleteResponse{}, tt.mockErr)

			c := etcdClient{
				client: &etcdcv3.Client{
					KV: mockKV,
				},
				ctx: context.Background(),
			}

			err := c.DeleteService(tt.key)

			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, tt.mockErr, err)
			} else {
				require.NoError(t, err)
			}
			mockKV.AssertExpectations(t)
		})
	}
}

func TestSaveService(t *testing.T) {
	type testCase struct {
		name       string
		service    *Service
		mockPutErr error
		wantErr    bool
	}
	tests := []testCase{
		{
			name: "success",
			service: &Service{
				Host:     "example.com",
				Port:     80,
				Priority: 1,
				Weight:   10,
				Text:     "hello",
				Key:      "/prefix/1",
			},
		},
		{
			name: "etcd put error",
			service: &Service{
				Host: "example.com",
				Key:  "/prefix/2",
			},
			mockPutErr: errors.New("etcd failure"),
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockKV := new(MockEtcdKV)
			value, err := json.Marshal(&tt.service)
			require.NoError(t, err)
			mockKV.On("Put", mock.Anything, tt.service.Key, string(value)).
				Return(&etcdcv3.PutResponse{}, tt.mockPutErr)

			c := etcdClient{
				client: &etcdcv3.Client{
					KV: mockKV,
				},
				ctx: context.TODO(),
			}

			err = c.SaveService(tt.service)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockKV.AssertExpectations(t)
		})
	}
}

func TestNewCoreDNSProvider(t *testing.T) {
	tests := []struct {
		name    string
		envs    map[string]string
		wantErr bool
		errMsg  string
	}{
		{
			name: "default config",
			envs: map[string]string{},
		},
		{
			name: "config with ETCD_URLS",
			envs: map[string]string{"ETCD_URLS": "http://example.com:2379"},
		},
		{
			name:    "config with unsupported protocol",
			envs:    map[string]string{"ETCD_URLS": "ftp://example.com:20"},
			wantErr: true,
			errMsg:  "etcd URLs must start with either http:// or https://",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.TestHelperEnvSetter(t, tt.envs)

			provider, err := NewCoreDNSProvider(&endpoint.DomainFilter{}, "/prefix/", "test-owner-id", false)
			if tt.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, provider)
			}
		})
	}
}

func TestFindEp(t *testing.T) {
	tests := []struct {
		name     string
		slice    []*endpoint.Endpoint
		dnsName  string
		want     *endpoint.Endpoint
		wantBool bool
	}{
		{
			name: "found",
			slice: []*endpoint.Endpoint{
				{DNSName: "foo.example.com"},
				{DNSName: "bar.example.com"},
			},
			dnsName:  "bar.example.com",
			want:     &endpoint.Endpoint{DNSName: "bar.example.com"},
			wantBool: true,
		},
		{
			name: "not found",
			slice: []*endpoint.Endpoint{
				{DNSName: "foo.example.com"},
			},
			dnsName:  "baz.example.com",
			want:     nil,
			wantBool: false,
		},
		{
			name:     "empty slice",
			slice:    []*endpoint.Endpoint{},
			dnsName:  "foo.example.com",
			want:     nil,
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := findEp(tt.slice, tt.dnsName)
			assert.Equal(t, tt.wantBool, ok)
			if ok {
				assert.Equal(t, tt.dnsName, got.DNSName)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestCoreDNSProvider_updateTXTRecords_WithEdpoints(t *testing.T) {
	provider := coreDNSProvider{coreDNSPrefix: "/prefix/", txtOwnerID: "default"}
	dnsName := "foo.example.com"

	group := []*endpoint.Endpoint{
		{
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"txt-value"},
			Labels:     map[string]string{randomPrefixLabel: "pfx"},
			RecordTTL:  60,
		},
		{
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"txt-value-2"},
			Labels:     map[string]string{randomPrefixLabel: ""},
			RecordTTL:  60,
		},
	}

	services := provider.updateTXTRecords(dnsName, group, []*Service{})
	assert.Len(t, services, 2)
	expectedTexts := map[string]bool{"txt-value": false, "txt-value-2": false}
	for _, service := range services {
		if _, exists := expectedTexts[service.Text]; exists {
			expectedTexts[service.Text] = true
		}
	}
	for text, found := range expectedTexts {
		if !found {
			t.Errorf("Expected TXT value %q not found in services", text)
		}
	}
}

func TestCoreDNSProvider_updateTXTRecords_ClearsExtraText(t *testing.T) {
	provider := coreDNSProvider{coreDNSPrefix: "/prefix/", txtOwnerID: "default"}
	dnsName := "foo.example.com"

	group := []*endpoint.Endpoint{
		{
			RecordType: endpoint.RecordTypeTXT,
			Targets:    endpoint.Targets{"txt-value"},
			Labels:     map[string]string{randomPrefixLabel: "pfx"},
			RecordTTL:  60,
		},
	}

	var services []*Service
	services = append(services, &Service{Key: "/prefix/1", Text: "should-be-txt-value"})
	services = append(services, &Service{Key: "/prefix/2", Text: "should-be-empty"})
	services = append(services, &Service{Key: "/prefix/3", Text: "should-be-empty"})

	services = provider.updateTXTRecords(dnsName, group, services)
	assert.Len(t, services, 4)

	assert.Equal(t, "", services[0].Text)
	assert.Equal(t, "", services[1].Text)
	assert.Equal(t, "", services[2].Text)
	assert.Equal(t, "txt-value", services[3].Text)
}

func TestCoreDNSProviderMultiTXT(t *testing.T) {
	client := fakeETCDClient{
		services: make(map[string]Service),
	}
	provider := coreDNSProvider{
		client:        client,
		dryRun:        false,
		coreDNSPrefix: "/skydns/dev/test/",
		domainFilter:  nil,
		txtOwnerID:    "default",
	}

	// Test multi-target TXT record creation
	desired := []*endpoint.Endpoint{
		{
			DNSName:    "example.test.dev",
			RecordType: endpoint.RecordTypeTXT,
			RecordTTL:  30,
			Targets:    []string{"v=1", "key=value", "third=string"},
			Labels:     map[string]string{},
		},
	}

	changes := &plan.Changes{
		Create: desired,
	}

	err := provider.ApplyChanges(context.Background(), changes)
	if err != nil {
		t.Fatalf("ApplyChanges failed: %v", err)
	}

	// Verify three separate etcd keys were created
	if len(client.services) != 3 {
		t.Errorf("Expected 3 services, got %d", len(client.services))
		for k, v := range client.services {
			t.Logf("Service key: %s, text: %s", k, v.Text)
		}
	}

	// Verify each target has its own service
	expectedTexts := map[string]bool{"v=1": false, "key=value": false, "third=string": false}
	for _, service := range client.services {
		if _, exists := expectedTexts[service.Text]; exists {
			expectedTexts[service.Text] = true
		}
	}

	for text, found := range expectedTexts {
		if !found {
			t.Errorf("Expected TXT value %q not found in services", text)
		}
	}
}

func TestTXTRecordsHaveOwnerLabel(t *testing.T) {
	client := fakeETCDClient{
		map[string]Service{
			"/skydns/com/example": {Text: "test-value"},
		},
	}
	provider := coreDNSProvider{
		client:        client,
		coreDNSPrefix: defaultCoreDNSPrefix,
		txtOwnerID:    "test-owner-id",
	}

	endpoints, err := provider.Records(context.Background())
	require.NoError(t, err)

	// Find the TXT endpoint
	var txtEndpoint *endpoint.Endpoint
	for _, ep := range endpoints {
		if ep.RecordType == endpoint.RecordTypeTXT {
			txtEndpoint = ep
			break
		}
	}

	require.NotNil(t, txtEndpoint, "TXT endpoint should exist")
	assert.Equal(t, "test-owner-id", txtEndpoint.Labels[endpoint.OwnerLabelKey],
		"TXT endpoint should have owner label set to configured txtOwnerID")
}

func TestCoreDNSProviderMultiTXTCleanup(t *testing.T) {
	client := fakeETCDClient{
		services: make(map[string]Service),
	}
	provider := coreDNSProvider{
		client:        client,
		dryRun:        false,
		coreDNSPrefix: "/skydns/dev/test/",
		domainFilter:  nil,
		txtOwnerID:    "default",
	}

	// Create multi-target TXT record
	desired := []*endpoint.Endpoint{
		{
			DNSName:    "peer.test.dev",
			RecordType: endpoint.RecordTypeTXT,
			RecordTTL:  30,
			Targets:    []string{"v=1;id=node1", "additional-data", "third-value"},
			Labels:     map[string]string{},
		},
	}

	changes := &plan.Changes{
		Create: desired,
	}

	// Apply the creation
	err := provider.ApplyChanges(context.Background(), changes)
	require.NoError(t, err)

	// Verify all three TXT services were created
	assert.Equal(t, 3, len(client.services), "Expected 3 TXT services to be created")

	// Verify all expected targets exist
	expectedTexts := map[string]bool{"v=1;id=node1": false, "additional-data": false, "third-value": false}
	for _, service := range client.services {
		if _, exists := expectedTexts[service.Text]; exists {
			expectedTexts[service.Text] = true
		}
	}
	for text, found := range expectedTexts {
		assert.True(t, found, "Expected TXT value %q should exist before deletion", text)
	}

	// Now delete the multi-target TXT record
	deleteChanges := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "peer.test.dev",
				RecordType: endpoint.RecordTypeTXT,
				RecordTTL:  30,
				Targets:    []string{"v=1;id=node1", "additional-data", "third-value"},
				Labels:     map[string]string{}, // Note: labels are empty, simulating real deletion scenario
			},
		},
	}

	// Apply the deletion
	err = provider.ApplyChanges(context.Background(), deleteChanges)
	require.NoError(t, err)

	// Verify ALL TXT services were deleted
	remainingTXTServices := 0
	for _, service := range client.services {
		if service.Text != "" {
			remainingTXTServices++
			t.Errorf("Found remaining TXT service with text: %s", service.Text)
		}
	}
	assert.Equal(t, 0, remainingTXTServices, "All TXT services should be deleted")
}

func TestCoreDNSProviderPartialTXTCleanup(t *testing.T) {
	client := fakeETCDClient{
		services: make(map[string]Service),
	}
	provider := coreDNSProvider{
		client:        client,
		dryRun:        false,
		coreDNSPrefix: "/skydns/dev/test/",
		domainFilter:  nil,
		txtOwnerID:    "default",
	}

	// Create multi-target TXT record
	desired := []*endpoint.Endpoint{
		{
			DNSName:    "peer.test.dev",
			RecordType: endpoint.RecordTypeTXT,
			RecordTTL:  30,
			Targets:    []string{"keep-this", "delete-this", "also-delete"},
			Labels:     map[string]string{},
		},
	}

	changes := &plan.Changes{
		Create: desired,
	}

	// Apply the creation
	err := provider.ApplyChanges(context.Background(), changes)
	require.NoError(t, err)

	// Verify all three TXT services were created
	assert.Equal(t, 3, len(client.services), "Expected 3 TXT services to be created")

	// Now delete only some targets
	partialDeleteChanges := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "peer.test.dev",
				RecordType: endpoint.RecordTypeTXT,
				RecordTTL:  30,
				Targets:    []string{"delete-this", "also-delete"}, // Only deleting 2 of 3 targets
				Labels:     map[string]string{},
			},
		},
	}

	// Apply the partial deletion
	err = provider.ApplyChanges(context.Background(), partialDeleteChanges)
	require.NoError(t, err)

	// Verify only the specified targets were deleted, "keep-this" should remain
	remainingTexts := make([]string, 0)
	for _, service := range client.services {
		if service.Text != "" {
			remainingTexts = append(remainingTexts, service.Text)
		}
	}

	assert.Equal(t, 1, len(remainingTexts), "Should have exactly 1 remaining TXT service")
	assert.Equal(t, "keep-this", remainingTexts[0], "Only 'keep-this' should remain")
}

func TestCoreDNSProviderPeerScenarioTXTCleanup(t *testing.T) {
	client := fakeETCDClient{
		services: make(map[string]Service),
	}
	provider := coreDNSProvider{
		client:        client,
		dryRun:        false,
		coreDNSPrefix: "/skydns/",
		domainFilter:  nil,
		txtOwnerID:    "default",
	}

	// Create the exact scenario from the logs - peer1.kaleido.dev with two TXT targets
	desired := []*endpoint.Endpoint{
		{
			DNSName:    "peer1.example.dev",
			RecordType: endpoint.RecordTypeTXT,
			RecordTTL:  5,
			Targets:    []string{"v=1;signature=aBx3d5..", "additional-txt-value"},
			Labels:     map[string]string{},
		},
	}

	changes := &plan.Changes{
		Create: desired,
	}

	// Apply the creation
	err := provider.ApplyChanges(context.Background(), changes)
	require.NoError(t, err)

	// Verify both TXT services were created
	assert.Equal(t, 2, len(client.services), "Expected 2 TXT services to be created")

	// Log what was created for debugging
	for key, service := range client.services {
		t.Logf("Created service: key=%s, text=%q", key, service.Text)
	}

	// Verify both expected targets exist
	expectedTexts := map[string]bool{
		"v=1;signature=aBx3d5..": false,
		"additional-txt-value":   false,
	}
	for _, service := range client.services {
		if _, exists := expectedTexts[service.Text]; exists {
			expectedTexts[service.Text] = true
		}
	}
	for text, found := range expectedTexts {
		assert.True(t, found, "Expected TXT value %q should exist before deletion", text)
	}

	// Now delete the TXT record (simulating the exact deletion scenario)
	deleteChanges := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "peer1.example.dev",
				RecordType: endpoint.RecordTypeTXT,
				RecordTTL:  5,
				Targets:    []string{"v=1;signature=aBx3d5..", "additional-txt-value"},
				Labels:     map[string]string{}, // Empty labels simulating deletion flow
			},
		},
	}

	// Apply the deletion
	err = provider.ApplyChanges(context.Background(), deleteChanges)
	require.NoError(t, err)

	// Verify ALL TXT services were deleted
	remainingTXTServices := 0
	for key, service := range client.services {
		if service.Text != "" {
			remainingTXTServices++
			t.Errorf("Found remaining TXT service: key=%s, text=%q", key, service.Text)
		}
	}
	assert.Equal(t, 0, remainingTXTServices, "All TXT services should be deleted but found %d remaining", remainingTXTServices)
}

func TestCoreDNSProviderSpecialCharactersTXTCleanup(t *testing.T) {
	client := fakeETCDClient{
		services: make(map[string]Service),
	}
	provider := coreDNSProvider{
		client:        client,
		dryRun:        false,
		coreDNSPrefix: "/skydns/",
		domainFilter:  nil,
		txtOwnerID:    "default",
	}

	// Test with special characters that might cause encoding issues
	complexText := "v=1;signature=aBx3d5..SomeHashHere123"
	desired := []*endpoint.Endpoint{
		{
			DNSName:    "test.example.com",
			RecordType: endpoint.RecordTypeTXT,
			RecordTTL:  30,
			Targets:    []string{complexText, "simple-value"},
			Labels:     map[string]string{},
		},
	}

	changes := &plan.Changes{Create: desired}
	err := provider.ApplyChanges(context.Background(), changes)
	require.NoError(t, err)

	// Verify creation
	assert.Equal(t, 2, len(client.services), "Expected 2 TXT services to be created")

	foundComplex := false
	foundSimple := false
	for _, service := range client.services {
		if service.Text == complexText {
			foundComplex = true
		}
		if service.Text == "simple-value" {
			foundSimple = true
		}
	}
	assert.True(t, foundComplex, "Complex text should be found")
	assert.True(t, foundSimple, "Simple text should be found")

	// Now delete with the exact same text
	deleteChanges := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "test.example.com",
				RecordType: endpoint.RecordTypeTXT,
				RecordTTL:  30,
				Targets:    []string{complexText, "simple-value"},
				Labels:     map[string]string{},
			},
		},
	}

	err = provider.ApplyChanges(context.Background(), deleteChanges)
	require.NoError(t, err)

	// Verify ALL services were deleted
	remainingTXTServices := 0
	for key, service := range client.services {
		if service.Text != "" {
			remainingTXTServices++
			t.Errorf("Found remaining TXT service: key=%s, text=%q", key, service.Text)
		}
	}
	assert.Equal(t, 0, remainingTXTServices, "All TXT services should be deleted")
}

// TestCoreDNSProviderTXTRecordOrdering tests that TXT targets maintain user-defined order
func TestCoreDNSProviderTXTRecordOrdering(t *testing.T) {
	provider := coreDNSProvider{
		coreDNSPrefix: defaultCoreDNSPrefix,
		txtOwnerID:    "default",
	}

	// Test ordered targets
	orderedTargets := []string{
		"first-target",
		"second-target",
		"third-target",
	}

	testEndpoint := &endpoint.Endpoint{
		DNSName:    "test.example.com",
		RecordType: endpoint.RecordTypeTXT,
		Targets:    orderedTargets,
		RecordTTL:  300,
		Labels:     map[string]string{},
	}

	// First creation - should get ordered prefixes
	services := provider.updateTXTRecords("test.example.com", []*endpoint.Endpoint{testEndpoint}, []*Service{})

	require.Len(t, services, 3, "should create 3 services")

	// Verify ordering by checking prefixes
	servicesByText := make(map[string]*Service)
	for _, svc := range services {
		servicesByText[svc.Text] = svc
	}

	// Extract prefix from key and verify ordering
	firstKey := servicesByText["first-target"].Key
	secondKey := servicesByText["second-target"].Key
	thirdKey := servicesByText["third-target"].Key

	// Keys should be lexicographically ordered due to index prefixes
	require.True(t, firstKey < secondKey, "first target should have lexicographically smaller key")
	require.True(t, secondKey < thirdKey, "second target should have lexicographically smaller key than third")

	// Verify prefixes start with correct indices
	assert.True(t, strings.Contains(firstKey, "/0-"), "first target should have prefix starting with 0-")
	assert.True(t, strings.Contains(secondKey, "/1-"), "second target should have prefix starting with 1-")
	assert.True(t, strings.Contains(thirdKey, "/2-"), "third target should have prefix starting with 2-")
}

// TestCoreDNSProviderTXTRecordReordering tests reordering when targets change
func TestCoreDNSProviderTXTRecordReordering(t *testing.T) {
	provider := coreDNSProvider{
		coreDNSPrefix: defaultCoreDNSPrefix,
		txtOwnerID:    "default",
	}

	// Initial targets
	initialTargets := []string{"a", "b", "c"}
	testEndpoint := &endpoint.Endpoint{
		DNSName:    "test.example.com",
		RecordType: endpoint.RecordTypeTXT,
		Targets:    initialTargets,
		RecordTTL:  300,
		Labels:     map[string]string{},
	}

	// Create initial services
	services := provider.updateTXTRecords("test.example.com", []*endpoint.Endpoint{testEndpoint}, []*Service{})
	require.Len(t, services, 3)

	// Store the labels that were created
	initialLabels := make(map[string]string)
	for k, v := range testEndpoint.Labels {
		initialLabels[k] = v
	}

	// Now change the order: insert "x" in the middle, remove "b"
	reorderedTargets := []string{"a", "x", "c"}
	testEndpoint.Targets = reorderedTargets
	testEndpoint.Labels = initialLabels // Start with previous labels

	// Debug: Check what labels were actually created
	t.Logf("Initial labels: %v", initialLabels)
	t.Logf("Reordered targets: %v", reorderedTargets)

	// Check if reordering is needed (should be true because "x" is new and "c" moved position)
	needsReorder := provider.checkIfReorderNeeded(reorderedTargets, testEndpoint.Labels)
	t.Logf("Reorder needed: %v", needsReorder)
	require.True(t, needsReorder, "should detect that reordering is needed")

	// Update with reordered targets
	newServices := provider.updateTXTRecords("test.example.com", []*endpoint.Endpoint{testEndpoint}, []*Service{})

	// Should have services for new targets
	require.Len(t, newServices, 3, "should have 3 services after reordering")

	// Verify new ordering
	servicesByText := make(map[string]*Service)
	for _, svc := range newServices {
		servicesByText[svc.Text] = svc
	}

	aKey := servicesByText["a"].Key
	xKey := servicesByText["x"].Key
	cKey := servicesByText["c"].Key

	// Verify lexicographic ordering
	require.True(t, aKey < xKey, "a should come before x")
	require.True(t, xKey < cKey, "x should come before c")

	// Verify correct index prefixes
	assert.True(t, strings.Contains(aKey, "/0-"), "a should have prefix 0-")
	assert.True(t, strings.Contains(xKey, "/1-"), "x should have prefix 1-")
	assert.True(t, strings.Contains(cKey, "/2-"), "c should have prefix 2-")
}

// TestCoreDNSProviderTXTRecordOrderingEdgeCases tests complex reordering scenarios
func TestCoreDNSProviderTXTRecordOrderingEdgeCases(t *testing.T) {
	provider := coreDNSProvider{
		coreDNSPrefix: defaultCoreDNSPrefix,
		txtOwnerID:    "default",
	}

	// Step 1: Start with 5 targets
	targets := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
	testEndpoint := &endpoint.Endpoint{
		DNSName:    "test.example.com",
		RecordType: endpoint.RecordTypeTXT,
		Targets:    targets,
		RecordTTL:  300,
		Labels:     map[string]string{},
	}

	services := provider.updateTXTRecords("test.example.com", []*endpoint.Endpoint{testEndpoint}, []*Service{})
	require.Len(t, services, 5, "should create 5 initial services")

	// Verify initial ordering
	servicesByText := make(map[string]*Service)
	for _, svc := range services {
		servicesByText[svc.Text] = svc
	}

	assert.True(t, strings.Contains(servicesByText["alpha"].Key, "/0-"), "alpha should have prefix 0-")
	assert.True(t, strings.Contains(servicesByText["beta"].Key, "/1-"), "beta should have prefix 1-")
	assert.True(t, strings.Contains(servicesByText["gamma"].Key, "/2-"), "gamma should have prefix 2-")
	assert.True(t, strings.Contains(servicesByText["delta"].Key, "/3-"), "delta should have prefix 3-")
	assert.True(t, strings.Contains(servicesByText["epsilon"].Key, "/4-"), "epsilon should have prefix 4-")

	// Step 2: Replace the 3rd target (gamma → charlie)
	t.Logf("Step 2: Replace 3rd target (gamma → charlie)")
	targets = []string{"alpha", "beta", "charlie", "delta", "epsilon"}
	testEndpoint.Targets = targets

	needsReorder := provider.checkIfReorderNeeded(targets, testEndpoint.Labels)
	assert.True(t, needsReorder, "should need reorder when replacing target")

	services = provider.updateTXTRecords("test.example.com", []*endpoint.Endpoint{testEndpoint}, []*Service{})
	require.Len(t, services, 5, "should have 5 services after replacement")

	servicesByText = make(map[string]*Service)
	for _, svc := range services {
		servicesByText[svc.Text] = svc
	}

	assert.True(t, strings.Contains(servicesByText["alpha"].Key, "/0-"), "alpha should stay at prefix 0-")
	assert.True(t, strings.Contains(servicesByText["beta"].Key, "/1-"), "beta should stay at prefix 1-")
	assert.True(t, strings.Contains(servicesByText["charlie"].Key, "/2-"), "charlie should have prefix 2-")
	assert.True(t, strings.Contains(servicesByText["delta"].Key, "/3-"), "delta should stay at prefix 3-")
	assert.True(t, strings.Contains(servicesByText["epsilon"].Key, "/4-"), "epsilon should stay at prefix 4-")
	assert.Nil(t, servicesByText["gamma"], "gamma should no longer exist")

	// Step 3: Remove the 4th target (delta)
	t.Logf("Step 3: Remove 4th target (delta)")
	targets = []string{"alpha", "beta", "charlie", "epsilon"}
	testEndpoint.Targets = targets

	needsReorder = provider.checkIfReorderNeeded(targets, testEndpoint.Labels)
	assert.True(t, needsReorder, "should need reorder when removing middle target")

	services = provider.updateTXTRecords("test.example.com", []*endpoint.Endpoint{testEndpoint}, []*Service{})
	require.Len(t, services, 4, "should have 4 services after removal")

	servicesByText = make(map[string]*Service)
	for _, svc := range services {
		servicesByText[svc.Text] = svc
	}

	assert.True(t, strings.Contains(servicesByText["alpha"].Key, "/0-"), "alpha should stay at prefix 0-")
	assert.True(t, strings.Contains(servicesByText["beta"].Key, "/1-"), "beta should stay at prefix 1-")
	assert.True(t, strings.Contains(servicesByText["charlie"].Key, "/2-"), "charlie should stay at prefix 2-")
	assert.True(t, strings.Contains(servicesByText["epsilon"].Key, "/3-"), "epsilon should move to prefix 3-")
	assert.Nil(t, servicesByText["delta"], "delta should no longer exist")

	// Step 4: Swap 1st and 2nd targets (alpha ↔ beta)
	t.Logf("Step 4: Swap 1st and 2nd targets (alpha ↔ beta)")
	targets = []string{"beta", "alpha", "charlie", "epsilon"}
	testEndpoint.Targets = targets

	needsReorder = provider.checkIfReorderNeeded(targets, testEndpoint.Labels)
	assert.True(t, needsReorder, "should need reorder when swapping targets")

	services = provider.updateTXTRecords("test.example.com", []*endpoint.Endpoint{testEndpoint}, []*Service{})
	require.Len(t, services, 4, "should have 4 services after swap")

	servicesByText = make(map[string]*Service)
	for _, svc := range services {
		servicesByText[svc.Text] = svc
	}

	assert.True(t, strings.Contains(servicesByText["beta"].Key, "/0-"), "beta should move to prefix 0-")
	assert.True(t, strings.Contains(servicesByText["alpha"].Key, "/1-"), "alpha should move to prefix 1-")
	assert.True(t, strings.Contains(servicesByText["charlie"].Key, "/2-"), "charlie should stay at prefix 2-")
	assert.True(t, strings.Contains(servicesByText["epsilon"].Key, "/3-"), "epsilon should stay at prefix 3-")

	// Step 5: Remove all targets
	t.Logf("Step 5: Remove all targets")
	targets = []string{}
	testEndpoint.Targets = targets

	services = provider.updateTXTRecords("test.example.com", []*endpoint.Endpoint{testEndpoint}, []*Service{})

	// Should return empty services for an endpoint with no TXT targets
	txtServiceCount := 0
	for _, svc := range services {
		if svc.Text != "" {
			txtServiceCount++
		}
	}
	assert.Equal(t, 0, txtServiceCount, "should have no TXT services after removing all targets")
}
