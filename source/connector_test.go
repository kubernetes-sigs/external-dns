/*
Copyright 2018 The Kubernetes Authors.

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

package source

import (
	"encoding/gob"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"sigs.k8s.io/external-dns/endpoint"
)

type ConnectorSuite struct {
	suite.Suite
}

func (suite *ConnectorSuite) SetupTest() {

}

func startServerToServeTargets(t *testing.T, server string, endpoints []*endpoint.Endpoint) {
	ln, err := net.Listen("tcp", server)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			ln.Close()
			return
		}
		enc := gob.NewEncoder(conn)
		enc.Encode(endpoints)
		ln.Close()
	}()
	t.Logf("Server listening on %s", server)
}

func TestConnectorSource(t *testing.T) {
	suite.Run(t, new(ConnectorSuite))
	t.Run("Interface", testConnectorSourceImplementsSource)
	t.Run("Endpoints", testConnectorSourceEndpoints)
}

// testConnectorSourceImplementsSource tests that connectorSource is a valid Source.
func testConnectorSourceImplementsSource(t *testing.T) {
	assert.Implements(t, (*Source)(nil), new(connectorSource))
}

// testConnectorSourceEndpoints tests that NewConnectorSource doesn't return an error.
func testConnectorSourceEndpoints(t *testing.T) {
	for _, ti := range []struct {
		title               string
		serverListenAddress string
		serverAddress       string
		expected            []*endpoint.Endpoint
		expectError         bool
	}{
		{
			title:               "invalid remote server",
			serverListenAddress: "",
			serverAddress:       "localhost:8091",
			expectError:         true,
		},
		{
			title:               "valid remote server with no endpoints",
			serverListenAddress: "127.0.0.1:8080",
			serverAddress:       "127.0.0.1:8080",
			expectError:         false,
		},
		{
			title:               "valid remote server",
			serverListenAddress: "127.0.0.1:8081",
			serverAddress:       "127.0.0.1:8081",
			expected: []*endpoint.Endpoint{
				{DNSName: "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectError: false,
		},
		{
			title:               "valid remote server with multiple endpoints",
			serverListenAddress: "127.0.0.1:8082",
			serverAddress:       "127.0.0.1:8082",
			expected: []*endpoint.Endpoint{
				{DNSName: "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
				{DNSName: "xyz.example.org",
					Targets:    endpoint.Targets{"abc.example.org"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  180,
				},
			},
			expectError: false,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			if ti.serverListenAddress != "" {
				startServerToServeTargets(t, ti.serverListenAddress, ti.expected)
			}
			cs, _ := NewConnectorSource(ti.serverAddress)

			endpoints, err := cs.Endpoints()
			if ti.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Validate returned endpoints against expected endpoints.
			validateEndpoints(t, endpoints, ti.expected)
		})
	}
}
