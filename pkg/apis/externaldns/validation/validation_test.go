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

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"

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

func TestValidateBadIgnoreHostnameAnnotationsConfig(t *testing.T) {
	cfg := externaldns.NewConfig()
	cfg.IgnoreHostnameAnnotation = true
	cfg.FQDNTemplate = ""

	assert.Error(t, ValidateConfig(cfg))
}

func TestValidateBadRfc2136Config(t *testing.T) {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "rfc2136"
	cfg.RFC2136MinTTL = -1
	cfg.RFC2136BatchChangeSize = 50

	err := ValidateConfig(cfg)

	assert.NotNil(t, err)
}

func TestValidateBadRfc2136Batch(t *testing.T) {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "rfc2136"
	cfg.RFC2136MinTTL = 3600
	cfg.RFC2136BatchChangeSize = 0

	err := ValidateConfig(cfg)

	assert.NotNil(t, err)
}

func TestValidateGoodRfc2136Config(t *testing.T) {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "rfc2136"
	cfg.RFC2136MinTTL = 3600
	cfg.RFC2136BatchChangeSize = 50

	err := ValidateConfig(cfg)

	assert.Nil(t, err)
}

func TestValidateBadRfc2136GssTsigConfig(t *testing.T) {
	invalidRfc2136GssTsigConfigs := []*externaldns.Config{
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			RFC2136GSSTSIG:          true,
			RFC2136KerberosRealm:    "test-realm",
			RFC2136KerberosUsername: "test-user",
			RFC2136KerberosPassword: "",
			RFC2136MinTTL:           3600,
			RFC2136BatchChangeSize:  50,
		},
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			RFC2136GSSTSIG:          true,
			RFC2136KerberosRealm:    "test-realm",
			RFC2136KerberosUsername: "",
			RFC2136KerberosPassword: "test-pass",
			RFC2136MinTTL:           3600,
			RFC2136BatchChangeSize:  50,
		},
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			RFC2136GSSTSIG:          true,
			RFC2136Insecure:         true,
			RFC2136KerberosRealm:    "test-realm",
			RFC2136KerberosUsername: "test-user",
			RFC2136KerberosPassword: "test-pass",
			RFC2136MinTTL:           3600,
			RFC2136BatchChangeSize:  50,
		},
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			RFC2136GSSTSIG:          true,
			RFC2136KerberosRealm:    "",
			RFC2136KerberosUsername: "test-user",
			RFC2136KerberosPassword: "",
			RFC2136MinTTL:           3600,
			RFC2136BatchChangeSize:  50,
		},
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			RFC2136GSSTSIG:          true,
			RFC2136KerberosRealm:    "",
			RFC2136KerberosUsername: "",
			RFC2136KerberosPassword: "test-pass",
			RFC2136MinTTL:           3600,
			RFC2136BatchChangeSize:  50,
		},
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			RFC2136GSSTSIG:          true,
			RFC2136Insecure:         true,
			RFC2136KerberosRealm:    "",
			RFC2136KerberosUsername: "test-user",
			RFC2136KerberosPassword: "test-pass",
			RFC2136MinTTL:           3600,
			RFC2136BatchChangeSize:  50,
		},
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			RFC2136GSSTSIG:          true,
			RFC2136KerberosRealm:    "",
			RFC2136KerberosUsername: "test-user",
			RFC2136KerberosPassword: "test-pass",
			RFC2136MinTTL:           3600,
			RFC2136BatchChangeSize:  50,
		},
	}

	for _, cfg := range invalidRfc2136GssTsigConfigs {
		err := ValidateConfig(cfg)

		assert.NotNil(t, err)
	}
}

func TestValidateGoodRfc2136GssTsigConfig(t *testing.T) {
	validRfc2136GssTsigConfigs := []*externaldns.Config{
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			RFC2136GSSTSIG:          true,
			RFC2136Insecure:         false,
			RFC2136KerberosRealm:    "test-realm",
			RFC2136KerberosUsername: "test-user",
			RFC2136KerberosPassword: "test-pass",
			RFC2136MinTTL:           3600,
			RFC2136BatchChangeSize:  50,
		},
	}

	for _, cfg := range validRfc2136GssTsigConfigs {
		err := ValidateConfig(cfg)

		assert.Nil(t, err)
	}
}
