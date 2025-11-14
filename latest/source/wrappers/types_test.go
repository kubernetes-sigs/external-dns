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

package wrappers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildSourceWithWrappers(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *Config
		asserts func(*testing.T, *Config)
	}{
		{
			name: "configuration with target filter wrapper",
			cfg: NewConfig(
				WithTargetNetFilter([]string{"10.0.0.0/8"}),
			),
			asserts: func(t *testing.T, cfg *Config) {
				assert.True(t, cfg.isSourceWrapperInstrumented("target-filter"))
			},
		},
		{
			name: "configuration with nat64 networks",
			cfg: NewConfig(
				WithNAT64Networks([]string{"2001:db8::/96"}),
			),
			asserts: func(t *testing.T, cfg *Config) {
				assert.True(t, cfg.isSourceWrapperInstrumented("nat64"))
			},
		},
		{
			name: "default configuration",
			cfg:  NewConfig(),
			asserts: func(t *testing.T, cfg *Config) {
				assert.True(t, cfg.isSourceWrapperInstrumented("dedup"))
				assert.False(t, cfg.isSourceWrapperInstrumented("nat64"))
				assert.False(t, cfg.isSourceWrapperInstrumented("target-filter"))
			},
		},
		{
			name: "with TTL and NAT64",
			cfg: NewConfig(
				WithMinTTL(300),
				WithNAT64Networks([]string{"2001:db8::/96"}),
			),
			asserts: func(t *testing.T, cfg *Config) {
				assert.True(t, cfg.isSourceWrapperInstrumented("dedup"))
				assert.True(t, cfg.isSourceWrapperInstrumented("nat64"))
				assert.True(t, cfg.isSourceWrapperInstrumented("post-processor"))
				assert.False(t, cfg.isSourceWrapperInstrumented("target-filter"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := WrapSources(nil, tt.cfg)
			require.NoError(t, err)
			tt.asserts(t, tt.cfg)
		})
	}
}

func TestWrapSources_NAT64Error(t *testing.T) {
	cfg := NewConfig(WithNAT64Networks([]string{"badnet"}))
	src, err := WrapSources(nil, cfg)
	assert.Nil(t, src)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create NAT64 source wrapper")
}

func TestWithDefaultTargets(t *testing.T) {
	cfg := &Config{}
	opt := WithDefaultTargets([]string{"1.2.3.4"})
	opt(cfg)
	assert.Equal(t, []string{"1.2.3.4"}, cfg.defaultTargets)
}

func TestWithForceDefaultTargets(t *testing.T) {
	cfg := &Config{}
	opt := WithForceDefaultTargets(true)
	opt(cfg)
	assert.True(t, cfg.forceDefaultTargets)
}

func TestWithNAT64Networks(t *testing.T) {
	cfg := &Config{}
	opt := WithNAT64Networks([]string{"2001:db8::/96"})
	opt(cfg)
	assert.Equal(t, []string{"2001:db8::/96"}, cfg.nat64Networks)
}

func TestWithTargetNetFilter(t *testing.T) {
	cfg := &Config{}
	opt := WithTargetNetFilter([]string{"10.0.0.0/8"})
	opt(cfg)
	assert.Equal(t, []string{"10.0.0.0/8"}, cfg.targetNetFilter)
}

func TestWithExcludeTargetNets(t *testing.T) {
	cfg := &Config{}
	opt := WithExcludeTargetNets([]string{"192.168.0.0/16"})
	opt(cfg)
	assert.Equal(t, []string{"192.168.0.0/16"}, cfg.excludeTargetNets)
}

func TestWithMinTTL(t *testing.T) {
	cfg := &Config{}
	opt := WithMinTTL(300 * time.Second)
	opt(cfg)
	assert.Equal(t, 300*time.Second, cfg.minTTL)
}

func TestAddSourceWrapperAndIsSourceWrapperInstrumented(t *testing.T) {
	cfg := &Config{}
	assert.False(t, cfg.isSourceWrapperInstrumented("dedup"))
	cfg.addSourceWrapper("dedup")
	assert.True(t, cfg.isSourceWrapperInstrumented("dedup"))
	cfg.addSourceWrapper("nat64")
	assert.True(t, cfg.isSourceWrapperInstrumented("nat64"))
	assert.False(t, cfg.isSourceWrapperInstrumented("target-filter"))
}
