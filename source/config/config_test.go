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

package config

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlfake "sigs.k8s.io/controller-runtime/pkg/client/fake"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/pkg/crd"
)

func TestNewSourceConfig(t *testing.T) {
	tests := []struct {
		name           string
		cfg            *externaldns.Config
		wantConfigured bool
		wantCombining  bool
		wantErr        bool
		errContains    string
	}{
		{
			name: "no templates configured",
			cfg:  &externaldns.Config{},
		},
		{
			name: "fqdn template only",
			cfg: &externaldns.Config{
				FQDNTemplate: []string{"{{.Name}}.example.com"},
			},
			wantConfigured: true,
		},
		{
			name: "fqdn template with combine",
			cfg: &externaldns.Config{
				FQDNTemplate:             []string{"{{.Name}}.example.com"},
				CombineFQDNAndAnnotation: true,
			},
			wantConfigured: true,
			wantCombining:  true,
		},
		{
			name: "all three templates configured",
			cfg: &externaldns.Config{
				FQDNTemplate:             []string{"{{.Name}}.example.com"},
				TargetTemplate:           []string{"{{.Name}}.targets.example.com"},
				FQDNTargetTemplate:       []string{"{{.Name}}.example.com:{{.Name}}.targets.example.com"},
				CombineFQDNAndAnnotation: true,
			},
			wantConfigured: true,
			wantCombining:  true,
		},
		{
			name: "multiple fqdn templates",
			cfg: &externaldns.Config{
				FQDNTemplate: []string{"{{.Name}}.a.example.com", "{{.Name}}.b.example.com"},
			},
			wantConfigured: true,
		},
		{
			name:        "invalid fqdn template",
			cfg:         &externaldns.Config{FQDNTemplate: []string{"{{.Name"}},
			wantErr:     true,
			errContains: `--fqdn-template[0]`,
		},
		{
			name:        "invalid target template",
			cfg:         &externaldns.Config{TargetTemplate: []string{"{{.Status.LoadBalancer.Ingress"}},
			wantErr:     true,
			errContains: `--target-template[0]`,
		},
		{
			name:        "invalid fqdn-target template",
			cfg:         &externaldns.Config{FQDNTargetTemplate: []string{"{{.Name}}.example.com:{{.Status"}},
			wantErr:     true,
			errContains: `--fqdn-target-template[0]`,
		},
		{
			name: "duplicate define block in fqdn templates",
			cfg: &externaldns.Config{
				FQDNTemplate: []string{
					`{{ define "zone" }}example.com{{ end }}{{.Name}}.{{ template "zone" }}`,
					`{{ define "zone" }}other.com{{ end }}{{.Name}}.{{ template "zone" }}`,
				},
			},
			wantErr:     true,
			errContains: `--fqdn-template[1]`,
		},
		{
			name:        "invalid label filter",
			cfg:         &externaldns.Config{LabelFilter: "#invalid-selector"},
			wantErr:     true,
			errContains: `Invalid value: "#invalid-selector"`,
		},
		{
			name:        "invalid annotation filter",
			cfg:         &externaldns.Config{AnnotationFilter: "kubernetes.io/gateway.name in (a b)"},
			wantErr:     true,
			errContains: `couldn't parse the selector string`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSourceConfig(tt.cfg)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					require.ErrorContains(t, err, tt.errContains)
				}
				return
			}
			require.NoError(t, err)
			tmpl := got.TemplateEngine
			assert.Equal(t, tt.wantConfigured, tmpl.IsConfigured(), "IsConfigured")
			assert.Equal(t, tt.wantCombining, tmpl.Combining(), "Combining")
		})
	}
}

func TestNewSourceConfig_PropagatesKubeAPIRateLimits(t *testing.T) {
	cfg := &externaldns.Config{
		KubeAPIQPS:   20,
		KubeAPIBurst: 40,
	}
	got, err := NewSourceConfig(cfg)
	require.NoError(t, err)
	assert.Equal(t, 20, got.KubeAPIQPS)
	assert.Equal(t, 40, got.KubeAPIBurst)
}

func TestConfig_Namespace(t *testing.T) {
	tests := []struct {
		name       string
		namespaces []string
		want       string
	}{
		{"no namespaces means all", nil, v1.NamespaceAll},
		{"empty slice means all", []string{}, v1.NamespaceAll},
		{"single namespace", []string{"team-a"}, "team-a"},
		{"multiple returns first", []string{"team-a", "team-b"}, "team-a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{Namespaces: tt.namespaces}
			assert.Equal(t, tt.want, cfg.Namespace())
		})
	}
}

func TestConfig_EnsureClientGen_InitsOnceAndCaches(t *testing.T) {
	cfg := &Config{}
	var calls atomic.Int32
	sentinel := new(struct{})
	init := func() any {
		calls.Add(1)
		return sentinel
	}

	got1 := cfg.EnsureClientGen(init)
	got2 := cfg.EnsureClientGen(func() any {
		t.Fatal("init called twice")
		return nil
	})

	assert.Same(t, sentinel, got1)
	assert.Same(t, got1, got2)
	assert.Equal(t, int32(1), calls.Load())
}

func TestConfig_EnsureClientGen_HonoursSetClientGen(t *testing.T) {
	injected := new(struct{})
	cfg := &Config{}
	cfg.SetClientGen(injected)

	got := cfg.EnsureClientGen(func() any {
		t.Fatal("init should be skipped when a generator is already set")
		return nil
	})
	assert.Same(t, injected, got)
}

func TestConfig_EnsureClientGen_ConcurrentCallersSeeSameValue(t *testing.T) {
	cfg := &Config{}
	var calls atomic.Int32
	init := func() any {
		calls.Add(1)
		return new(struct{})
	}

	const n = 32
	results := make([]any, n)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := range n {
		go func() {
			defer wg.Done()
			results[i] = cfg.EnsureClientGen(init)
		}()
	}
	wg.Wait()

	assert.Equal(t, int32(1), calls.Load(), "init must only run once")
	for i := 1; i < n; i++ {
		assert.Same(t, results[0], results[i])
	}
}

func TestConfig_CRDClients(t *testing.T) {
	t.Run("empty by default", func(t *testing.T) {
		assert.Nil(t, (&Config{}).CRDClients())
	})

	t.Run("SetCRDClients stores the value", func(t *testing.T) {
		cc := crd.NewCRDClients(fakeCRDClient(t), fakeCRDClient(t))
		cfg := &Config{}
		cfg.SetCRDClients(cc)
		assert.Same(t, cc, cfg.CRDClients())
	})

	t.Run("WithCRDClients applies via NewSourceConfig", func(t *testing.T) {
		cc := crd.NewCRDClients(fakeCRDClient(t), fakeCRDClient(t))
		cfg, err := NewSourceConfig(&externaldns.Config{}, WithCRDClients(cc))
		require.NoError(t, err)
		assert.Same(t, cc, cfg.CRDClients())
	})
}

func fakeCRDClient(t *testing.T) ctrlclient.Client {
	t.Helper()
	return ctrlfake.NewClientBuilder().Build()
}
