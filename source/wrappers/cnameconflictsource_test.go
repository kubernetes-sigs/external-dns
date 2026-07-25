/*
Copyright 2026 The Kubernetes Authors.

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

package wrappers

import (
	"errors"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/testutils"
	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
	"sigs.k8s.io/external-dns/source"
)

// Validates that cnameConflictSource is a Source
var _ source.Source = &cnameConflictSource{}

func TestCNAMEConflictSourceEndpoints(t *testing.T) {
	for _, tc := range []struct {
		title           string
		domainFilter    endpoint.DomainFilterInterface
		endpoints       []*endpoint.Endpoint
		wantWarnings    []string
		unwantedRecords []string
	}{
		{
			title:        "conflict matching the domain filter warns",
			domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "b.elb.com"),
			},
			wantWarnings: []string{
				"Only one CNAME per name",
				"api.example.com CNAME a.elb.com",
				"api.example.com CNAME b.elb.com",
			},
		},
		{
			title:        "conflict outside the domain filter does not warn",
			domainFilter: endpoint.NewDomainFilter([]string{"managed.example.internal"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "b.elb.com"),
			},
			unwantedRecords: []string{"Only one CNAME per name"},
		},
		{
			title:        "conflict warns when no domain filter is configured",
			domainFilter: nil,
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "b.elb.com"),
			},
			wantWarnings: []string{"Only one CNAME per name"},
		},
		{
			title:        "conflict warns with an empty domain filter",
			domainFilter: endpoint.NewDomainFilterWithOptions(),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "b.elb.com"),
			},
			wantWarnings: []string{"Only one CNAME per name"},
		},
		{
			title:        "identical CNAMEs do not warn",
			domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
			},
			unwantedRecords: []string{"Only one CNAME per name"},
		},
		{
			title:        "same DNS name with different set identifiers does not warn",
			domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "a.elb.com").WithSetIdentifier("weight-1"),
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "b.elb.com").WithSetIdentifier("weight-2"),
			},
			unwantedRecords: []string{"Only one CNAME per name"},
		},
		{
			title:        "non-CNAME records with the same DNS name do not warn",
			domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
			endpoints: []*endpoint.Endpoint{
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeA, "1.1.1.1"),
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeA, "2.2.2.2"),
			},
			unwantedRecords: []string{"Only one CNAME per name"},
		},
		{
			title:        "nil endpoints and CNAMEs without targets are skipped",
			domainFilter: endpoint.NewDomainFilter([]string{"example.com"}),
			endpoints: []*endpoint.Endpoint{
				nil,
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME),
				endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
			},
			unwantedRecords: []string{"Only one CNAME per name"},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)

			src := NewCNAMEConflictSource(testutils.NewMockSource(tc.endpoints...), tc.domainFilter)

			endpoints, err := src.Endpoints(t.Context())
			require.NoError(t, err)
			assert.Equal(t, tc.endpoints, endpoints, "endpoints must be returned unmodified")

			for _, msg := range tc.wantWarnings {
				logtest.TestHelperLogContainsWithLogLevel(msg, log.WarnLevel, hook, t)
			}
			for _, msg := range tc.unwantedRecords {
				logtest.TestHelperLogNotContains(msg, hook, t)
			}
		})
	}
}

func TestCNAMEConflictSourceOutOfScopeConflictLoggedAtDebug(t *testing.T) {
	hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

	src := NewCNAMEConflictSource(
		testutils.NewMockSource(
			endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "a.elb.com"),
			endpoint.NewEndpoint("api.example.com", endpoint.RecordTypeCNAME, "b.elb.com"),
		),
		endpoint.NewDomainFilter([]string{"managed.example.internal"}),
	)

	_, err := src.Endpoints(t.Context())
	require.NoError(t, err)

	logtest.TestHelperLogContainsWithLogLevel("Skipping CNAME conflict warning for api.example.com", log.DebugLevel, hook, t)
}

func TestCNAMEConflictSourceEndpointsPropagatesError(t *testing.T) {
	mockSource := new(testutils.MockSource)
	mockSource.On("Endpoints").Return(nil, errors.New("inner source failed"))

	src := NewCNAMEConflictSource(mockSource, nil)

	endpoints, err := src.Endpoints(t.Context())
	assert.Nil(t, endpoints)
	require.EqualError(t, err, "inner source failed")
}

func TestCNAMEConflictSourceAddEventHandler(t *testing.T) {
	mockSource := new(testutils.MockSource)
	mockSource.On("AddEventHandler", mock.Anything).Return()

	src := NewCNAMEConflictSource(mockSource, nil)
	src.AddEventHandler(t.Context(), func() {})

	mockSource.AssertExpectations(t)
}
