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

package genericapi

import (
	"context"
	"encoding/json"
	"sort"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type mockGoDaddyClient struct {
	mock.Mock
	currentTest *testing.T
}

func newMockGoDaddyClient(t *testing.T) *mockGoDaddyClient {
	return &mockGoDaddyClient{
		currentTest: t,
	}
}

var (
	zoneNameExampleOrg string = "example.org"
	zoneNameExampleNet string = "example.net"
)

func (c *mockGoDaddyClient) Post(endpoint string, input interface{}, output interface{}) error {
	log.Infof("POST: %s - %v", endpoint, input)
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockGoDaddyClient) Patch(endpoint string, input interface{}, output interface{}) error {
	log.Infof("PATCH: %s - %v", endpoint, input)
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockGoDaddyClient) Put(endpoint string, input interface{}, output interface{}) error {
	log.Infof("PUT: %s - %v", endpoint, input)
	stub := c.Called(endpoint, input)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockGoDaddyClient) Get(endpoint string, output interface{}) error {
	log.Infof("GET: %s", endpoint)
	stub := c.Called(endpoint)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func (c *mockGoDaddyClient) Delete(endpoint string, output interface{}) error {
	log.Infof("DELETE: %s", endpoint)
	stub := c.Called(endpoint)
	data, _ := json.Marshal(stub.Get(0))
	json.Unmarshal(data, output)
	return stub.Error(1)
}

func TestGoDaddyRecords(t *testing.T) {
	assert := assert.New(t)
	client := newMockGoDaddyClient(t)
	provider := &GDProvider{
		client: client,
	}

	client.On("Get", "/v1/regions/test-region/domains").Return([]endpoint.Endpoint{
		{
			DNSName: "agent",
			Targets: endpoint.Targets{"192.168.0.138", "192.168.1.138"},
			RecordType: "A",
			SetIdentifier: "agent.test-region.cratedb.net",
			RecordTTL: endpoint.TTL(100),
		},
	}, nil).Once()

	endpoints, err := provider.Records(context.TODO())
	assert.NoError(err)

	// Little fix for multi targets endpoint
	for _, endpoint := range endpoints {
		sort.Strings(endpoint.Targets)
	}

	assert.ElementsMatch(endpoints, []*endpoint.Endpoint{
		{
			DNSName:    "agent",
			Targets: []string{
				"192.168.0.138",
				"192.168.1.138",
			},
			RecordType: "A",
			SetIdentifier: "agent.test-region.cratedb.net",
			RecordTTL:  100,
		},
	})

	client.AssertExpectations(t)
}

func TestGoDaddyChange(t *testing.T) {
	assert := assert.New(t)
	client := newMockGoDaddyClient(t)
	provider := &GDProvider{
		client: client,
	}

	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName: "to-be-created",
				Targets: []string{"192.168.0.138", "192.168.1.138"},
				RecordType: "A",
				SetIdentifier: "to-be-created.test-region.cratedb.net",
				RecordTTL: endpoint.TTL(100),
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName: "to-be-deleted",
				Targets: []string{"192.168.0.138", "192.168.1.138"},
				RecordType: "A",
				SetIdentifier: "to-be-deleted.test-region.cratedb.net",
				RecordTTL: endpoint.TTL(100),
			},
		},
		UpdateOld: []*endpoint.Endpoint{
			{
				DNSName: "to-be-updated",
				Targets: []string{"192.168.0.138", "192.168.1.138"},
				RecordType: "A",
				SetIdentifier: "to-be-updated.test-region.cratedb.net",
				RecordTTL: endpoint.TTL(100),
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName: "to-be-updated",
				Targets: []string{"192.168.0.230", "192.168.1.231"},
				RecordType: "A",
				SetIdentifier: "to-be-updated.test-region.cratedb.net",
				RecordTTL: endpoint.TTL(444),
			},
		},
	}

	client.On("Post", "/v1/regions/test-region/domains", &changes).Return("OK", nil).Once()

	assert.NoError(provider.ApplyChanges(context.TODO(), &changes))

	client.AssertExpectations(t)
}
