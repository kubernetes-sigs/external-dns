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
	"context"
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

func startServerToServeTargets(t *testing.T, endpoints []*endpoint.Endpoint) net.Listener {
	ln, err := net.Listen("tcp", "localhost:0")
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
	t.Logf("Server listening on %s", ln.Addr().String())
	return ln
}

func TestConnectorSource(t *testing.T) {
	t.Parallel()

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
		title       string
		server      bool
		expected    []*endpoint.Endpoint
		expectError bool
	}{
		{
			title:       "invalid remote server",
			server:      false,
			expectError: true,
		},
		{
			title:       "valid remote server with no endpoints",
			server:      true,
			expectError: false,
		},
		{
			title:  "valid remote server",
			server: true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
			},
			expectError: false,
		},
		{
			title:  "valid remote server with multiple endpoints",
			server: true,
			expected: []*endpoint.Endpoint{
				{
					DNSName:    "abc.example.org",
					Targets:    endpoint.Targets{"1.2.3.4"},
					RecordType: endpoint.RecordTypeA,
					RecordTTL:  180,
				},
				{
					DNSName:    "xyz.example.org",
					Targets:    endpoint.Targets{"abc.example.org"},
					RecordType: endpoint.RecordTypeCNAME,
					RecordTTL:  180,
				},
			},
			expectError: false,
		},
	} {
		ti := ti
		t.Run(ti.title, func(t *testing.T) {
			t.Parallel()

			addr := "localhost:9999"
			if ti.server {
				ln := startServerToServeTargets(t, ti.expected)
				defer ln.Close()
				addr = ln.Addr().String()
			}
			cs, _ := NewConnectorSource(addr)

			endpoints, err := cs.Endpoints(context.Background())
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
