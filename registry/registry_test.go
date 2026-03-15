/*
Copyright 2025 The Kubernetes Authors.

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

package registry

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/provider"
	fakeprovider "sigs.k8s.io/external-dns/provider/fakes"
	"sigs.k8s.io/external-dns/provider/inmemory"
	"sigs.k8s.io/external-dns/registry/awssd"
	"sigs.k8s.io/external-dns/registry/dynamodb"
	"sigs.k8s.io/external-dns/registry/noop"
	"sigs.k8s.io/external-dns/registry/txt"
)

var (
	_ Registry = &awssd.AWSSDRegistry{}
	_ Registry = &dynamodb.DynamoDBRegistry{}
	_ Registry = &noop.NoopRegistry{}
	_ Registry = &txt.TXTRegistry{}
)

func TestSelectRegistry(t *testing.T) {
	tests := []struct {
		name     string
		cfg      *externaldns.Config
		provider provider.Provider
		wantErr  bool
		wantType string
	}{
		{
			name: "dynamoDB registry",
			cfg: &externaldns.Config{
				Registry:               DYNAMODB,
				AWSDynamoDBRegion:      "us-west-2",
				AWSDynamoDBTable:       "test-table",
				TXTOwnerID:             "owner-id",
				TXTWildcardReplacement: "wildcard",
				ManagedDNSRecordTypes:  []string{"A", "CNAME"},
				ExcludeDNSRecordTypes:  []string{"TXT"},
				TXTCacheInterval:       60,
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  false,
			wantType: "DynamoDBRegistry",
		},
		{
			name: "noop registry",
			cfg: &externaldns.Config{
				Registry: NOOP,
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  false,
			wantType: "NoopRegistry",
		},
		{
			name: "TXT registry",
			cfg: &externaldns.Config{
				Registry:               TXT,
				TXTPrefix:              "prefix",
				TXTOwnerID:             "owner-id",
				TXTCacheInterval:       60,
				TXTWildcardReplacement: "wildcard",
				ManagedDNSRecordTypes:  []string{"A", "CNAME"},
				ExcludeDNSRecordTypes:  []string{"TXT"},
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  false,
			wantType: "TXTRegistry",
		},
		{
			name: "aws-sd registry",
			cfg: &externaldns.Config{
				Registry:   AWSSD,
				TXTOwnerID: "owner-id",
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  false,
			wantType: "AWSSDRegistry",
		},
		{
			name: "unknown registry",
			cfg: &externaldns.Config{
				Registry: "unknown",
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  true,
			wantType: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg, err := SelectRegistry(tt.cfg, tt.provider)

			if tt.wantErr {
				require.Nil(t, reg)
				require.Error(t, err)
			} else {
				assert.NotNil(t, reg)
				require.NoError(t, err)
				assert.Contains(t, reflect.TypeOf(reg).String(), tt.wantType)
			}
		})
	}
}

func TestSelectRegistryUnknown(t *testing.T) {
	cfg := externaldns.NewConfig()
	cfg.Registry = "nope"

	reg, err := SelectRegistry(cfg, inmemory.NewInMemoryProvider())
	require.Error(t, err)
	require.Nil(t, reg)
}
