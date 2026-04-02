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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/internal/testutils"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/types"
)

func stubConfig(t *testing.T, extCfg *externaldns.Config) *source.Config {
	t.Helper()
	cfg, err := source.NewSourceConfig(extCfg, source.WithClientGenerator(testutils.StubClientGenerator{}))
	require.NoError(t, err)
	return cfg
}

func TestBuildWrappedSource(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *source.Config
		wantErr bool
	}{
		{
			name: "fake source with no extra wrappers",
			cfg:  stubConfig(t, &externaldns.Config{Sources: []string{types.Fake}}),
		},
		{
			name: "fake source with target filter wrapper",
			cfg: stubConfig(t, &externaldns.Config{
				Sources:         []string{types.Fake},
				TargetNetFilter: []string{"10.0.0.0/8"},
			}),
		},
		{
			name: "fake source with NAT64 networks",
			cfg: stubConfig(t, &externaldns.Config{
				Sources:       []string{types.Fake},
				NAT64Networks: []string{"2001:db8::/96"},
			}),
		},
		{
			name: "fake source with minTTL, provider, and preferAlias",
			cfg: stubConfig(t, &externaldns.Config{
				Sources:     []string{types.Fake},
				MinTTL:      300 * time.Second,
				Provider:    "aws",
				PreferAlias: true,
			}),
		},
		{
			name: "fake source with exclude target nets",
			cfg: stubConfig(t, &externaldns.Config{
				Sources:           []string{types.Fake},
				TargetNetFilter:   []string{"10.0.0.0/8"},
				ExcludeTargetNets: []string{"10.1.0.0/16"},
			}),
		},
		{
			name:    "unknown source returns error",
			cfg:     stubConfig(t, &externaldns.Config{Sources: []string{"does-not-exist"}}),
			wantErr: true,
		},
		{
			name: "invalid NAT64 network returns error",
			cfg: stubConfig(t, &externaldns.Config{
				Sources:       []string{types.Fake},
				NAT64Networks: []string{"not-a-cidr"},
			}),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src, err := Build(t.Context(), tt.cfg)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, src)
		})
	}
}
