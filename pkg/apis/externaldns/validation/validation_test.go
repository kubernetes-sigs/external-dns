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

package validation

import (
	"testing"

	"github.com/kubernetes-incubator/external-dns/pkg/apis/externaldns"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateFlags(t *testing.T) {
	cfg := newValidConfig(t)
	assert.NoError(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.LogFormat = "test"
	assert.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.LogFormat = ""
	assert.Error(t, ValidateConfig(cfg))

	for _, format := range []string{"text", "json"} {
		cfg = newValidConfig(t)
		cfg.LogFormat = format
		assert.NoError(t, ValidateConfig(cfg))
	}

	cfg = newValidConfig(t)
	cfg.Sources = []string{}
	assert.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.Provider = ""
	assert.Error(t, ValidateConfig(cfg))
}

func newValidConfig(t *testing.T) *externaldns.Config {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "test-provider"

	require.NoError(t, ValidateConfig(cfg))

	return cfg
}
