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

	cfg = newValidConfig(t)
	cfg.ServicePublishIPsType = ""
	assert.NoError(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.ServicePublishIPsType = "hello"
	assert.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.ServicePublishIPsType = "private"
	assert.NoError(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.ServicePublishIPsType = "public"
	assert.NoError(t, ValidateConfig(cfg))
}

func newValidConfig(t *testing.T) *externaldns.Config {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "test-provider"

	require.NoError(t, ValidateConfig(cfg))

	return cfg
}

func addRequiredFieldsForDyn(cfg *externaldns.Config) {
	cfg.LogFormat = "json"
	cfg.Sources = []string{"ingress"}
	cfg.Provider = "dyn"
}

func TestValidateBadDynConfig(t *testing.T) {
	badConfigs := []*externaldns.Config{
		{},
		{
			// only username
			DynUsername: "test",
		},
		{
			// only customer name
			DynCustomerName: "test",
		},
		{
			// negative timeout
			DynUsername:      "test",
			DynCustomerName:  "test",
			DynMinTTLSeconds: -1,
		},
	}

	for _, cfg := range badConfigs {
		addRequiredFieldsForDyn(cfg)
		err := ValidateConfig(cfg)
		assert.NotNil(t, err, "Configuration %+v should NOT have passed validation", cfg)
	}
}

func TestValidateGoodDynConfig(t *testing.T) {
	goodConfigs := []*externaldns.Config{
		{
			DynUsername:      "test",
			DynCustomerName:  "test",
			DynMinTTLSeconds: 600,
		},
		{
			DynUsername:      "test",
			DynCustomerName:  "test",
			DynMinTTLSeconds: 0,
		},
	}

	for _, cfg := range goodConfigs {
		addRequiredFieldsForDyn(cfg)
		err := ValidateConfig(cfg)
		assert.Nil(t, err, "Configuration should be valid, got this error instead", err)
	}
}
