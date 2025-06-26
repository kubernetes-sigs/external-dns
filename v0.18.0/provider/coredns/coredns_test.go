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
		"/skydns/local/domain1": {{Host: "5.5.5.5", Text: "string1"}},
		"/skydns/local/domain2": {{Host: "site.local"}},
	}
	validateServices(client.services, expectedServices1, t, 1)

	changes2 := &plan.Changes{
		Create: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain3.local", endpoint.RecordTypeA, "7.7.7.7"),
		},
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
		"/skydns/local/domain1": {{Host: "6.6.6.6", Text: "string1"}},
		"/skydns/local/domain2": {{Host: "site.local"}},
		"/skydns/local/domain3": {{Host: "7.7.7.7"}},
	}
	validateServices(client.services, expectedServices2, t, 2)

	changes3 := &plan.Changes{
		Delete: []*endpoint.Endpoint{
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeA, "6.6.6.6"),
			endpoint.NewEndpoint("domain1.local", endpoint.RecordTypeTXT, "string"),
			endpoint.NewEndpoint("domain3.local", endpoint.RecordTypeA, "7.7.7.7"),
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
		expectedServiceEntries := expectedServices[expectedKey]
		if expectedServiceEntries == nil {
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

			provider, err := NewCoreDNSProvider(&endpoint.DomainFilter{}, "/prefix/", false)
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
	provider := coreDNSProvider{coreDNSPrefix: "/prefix/"}
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
	assert.Equal(t, "txt-value", services[0].Text)
	assert.Equal(t, "txt-value-2", services[1].Text)
}

func TestCoreDNSProvider_updateTXTRecords_ClearsExtraText(t *testing.T) {
	provider := coreDNSProvider{coreDNSPrefix: "/prefix/"}
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
	assert.Len(t, services, 3)

	assert.Equal(t, "txt-value", services[0].Text)
	assert.Empty(t, services[1].Text)
}
