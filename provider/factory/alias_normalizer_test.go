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

package factory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type stubProvider struct {
	provider.BaseProvider
}

func (s *stubProvider) Records(_ context.Context) ([]*endpoint.Endpoint, error) { return nil, nil }
func (s *stubProvider) ApplyChanges(_ context.Context, _ *plan.Changes) error   { return nil }
func (s *stubProvider) AdjustEndpoints(eps []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return eps, nil
}

func TestAliasNormalizingMiddleware(t *testing.T) {
	tests := []struct {
		name        string
		recordType  string
		aliasIn     endpoint.AliasType
		aliasWanted endpoint.AliasType
	}{
		{
			name:        "A with alias=A normalized to true",
			recordType:  endpoint.RecordTypeA,
			aliasIn:     endpoint.AliasA,
			aliasWanted: endpoint.AliasTrue,
		},
		{
			name:        "AAAA with alias=AAAA normalized to true",
			recordType:  endpoint.RecordTypeAAAA,
			aliasIn:     endpoint.AliasAAAA,
			aliasWanted: endpoint.AliasTrue,
		},
		{
			name:        "A with alias=true unchanged",
			recordType:  endpoint.RecordTypeA,
			aliasIn:     endpoint.AliasTrue,
			aliasWanted: endpoint.AliasTrue,
		},
		{
			name:        "A with alias=AAAA not touched",
			recordType:  endpoint.RecordTypeA,
			aliasIn:     endpoint.AliasAAAA,
			aliasWanted: endpoint.AliasAAAA,
		},
		{
			name:        "AAAA with alias=A not touched",
			recordType:  endpoint.RecordTypeAAAA,
			aliasIn:     endpoint.AliasA,
			aliasWanted: endpoint.AliasA,
		},
		{
			name:        "A with alias=false unchanged",
			recordType:  endpoint.RecordTypeA,
			aliasIn:     endpoint.AliasFalse,
			aliasWanted: endpoint.AliasFalse,
		},
		{
			name:        "A with no alias property unchanged",
			recordType:  endpoint.RecordTypeA,
			aliasIn:     endpoint.AliasNone,
			aliasWanted: endpoint.AliasNone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ep := endpoint.NewEndpoint("example.com", tt.recordType, "target.example.com")
			if tt.aliasIn != endpoint.AliasNone {
				ep = ep.WithProviderSpecific(endpoint.ProviderSpecificAlias, string(tt.aliasIn))
			}

			p := newAliasNormalizingMiddleware(&stubProvider{})
			result, err := p.AdjustEndpoints([]*endpoint.Endpoint{ep})
			require.NoError(t, err)
			require.Len(t, result, 1)

			assert.Equal(t, tt.aliasWanted, result[0].GetAliasProperty())
		})
	}
}
