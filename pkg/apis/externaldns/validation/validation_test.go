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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/rest"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

func TestValidateFlags(t *testing.T) {
	cfg := newValidConfig(t)
	require.NoError(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.LogFormat = "test"
	require.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.LogFormat = ""
	require.Error(t, ValidateConfig(cfg))

	for _, format := range []string{"text", "json"} {
		cfg = newValidConfig(t)
		cfg.LogFormat = format
		require.NoError(t, ValidateConfig(cfg))
	}

	cfg = newValidConfig(t)
	cfg.Sources = []string{}
	require.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.Provider = ""
	require.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.IgnoreHostnameAnnotation = true
	cfg.FQDNTemplate = ""
	require.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.TXTPrefix = "foo"
	cfg.TXTSuffix = "bar"
	require.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.LabelFilter = "foo"
	require.NoError(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.LabelFilter = "foo=bar"
	require.NoError(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.LabelFilter = "#invalid-selector"
	require.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.AnnotationFilter = "kubernetes.io/gateway.class in (alb, nginx)"
	require.NoError(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.AnnotationFilter = "kubernetes.io/gateway.name in (a b)"
	require.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.AnnotationPrefix = ""
	require.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.AnnotationPrefix = "custom.io"
	require.Error(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.AnnotationPrefix = "custom.io/"
	require.NoError(t, ValidateConfig(cfg))

	cfg = newValidConfig(t)
	cfg.AnnotationPrefix = "external-dns.alpha.kubernetes.io/"
	require.NoError(t, ValidateConfig(cfg))

	t.Run("kube-api-qps and kube-api-burst", func(t *testing.T) {
		for _, tc := range []struct {
			name    string
			qps     int
			burst   int
			wantErr bool
		}{
			{name: "positive QPS and burst", qps: 10, burst: 20, wantErr: false},
			{name: "zero QPS", qps: 0, burst: 10, wantErr: true},
			{name: "zero burst", qps: 5, burst: 0, wantErr: true},
			{name: "negative QPS", qps: -1, burst: 10, wantErr: true},
			{name: "negative burst", qps: 5, burst: -1, wantErr: true},
		} {
			t.Run(tc.name, func(t *testing.T) {
				cfg := newValidConfig(t)
				cfg.KubeAPIQPS = tc.qps
				cfg.KubeAPIBurst = tc.burst
				if tc.wantErr {
					require.Error(t, ValidateConfig(cfg))
				} else {
					require.NoError(t, ValidateConfig(cfg))
				}
			})
		}
	})
}

func newValidConfig(t *testing.T) *externaldns.Config {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "test-provider"
	cfg.KubeAPIQPS = int(rest.DefaultQPS)
	cfg.KubeAPIBurst = rest.DefaultBurst

	require.NoError(t, ValidateConfig(cfg))

	return cfg
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

	assert.Error(t, err)
}

func TestValidateBadRfc2136Batch(t *testing.T) {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "rfc2136"
	cfg.RFC2136MinTTL = 3600
	cfg.RFC2136BatchChangeSize = 0

	err := ValidateConfig(cfg)

	assert.Error(t, err)
}

func TestValidateGoodRfc2136Config(t *testing.T) {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "rfc2136"
	cfg.RFC2136MinTTL = 3600
	cfg.RFC2136BatchChangeSize = 50
	cfg.KubeAPIQPS = int(rest.DefaultQPS)
	cfg.KubeAPIBurst = rest.DefaultBurst

	err := ValidateConfig(cfg)

	assert.NoError(t, err)
}

func TestValidateBadRfc2136GssTsigConfig(t *testing.T) {
	invalidRfc2136GssTsigConfigs := []*externaldns.Config{
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			AnnotationPrefix:        "external-dns.alpha.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.alpha.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.alpha.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.alpha.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.alpha.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.alpha.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.alpha.kubernetes.io/",
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

		assert.Error(t, err)
	}
}

func TestValidateGoodRfc2136GssTsigConfig(t *testing.T) {
	validRfc2136GssTsigConfigs := []*externaldns.Config{
		{
			LogFormat:               "json",
			Sources:                 []string{"test-source"},
			Provider:                "rfc2136",
			AnnotationPrefix:        "external-dns.alpha.kubernetes.io/",
			RFC2136GSSTSIG:          true,
			RFC2136Insecure:         false,
			RFC2136KerberosRealm:    "test-realm",
			RFC2136KerberosUsername: "test-user",
			RFC2136KerberosPassword: "test-pass",
			RFC2136MinTTL:           3600,
			RFC2136BatchChangeSize:  50,
			KubeAPIQPS:              int(rest.DefaultQPS),
			KubeAPIBurst:            rest.DefaultBurst,
		},
	}

	for _, cfg := range validRfc2136GssTsigConfigs {
		err := ValidateConfig(cfg)

		assert.NoError(t, err)
	}
}

func TestValidateBadAkamaiConfig(t *testing.T) {
	invalidAkamaiConfigs := []*externaldns.Config{
		{
			LogFormat:          "json",
			Sources:            []string{"test-source"},
			Provider:           "akamai",
			AnnotationPrefix:   "external-dns.alpha.kubernetes.io/",
			AkamaiClientToken:  "test-token",
			AkamaiClientSecret: "test-secret",
			AkamaiAccessToken:  "test-access-token",
			AkamaiEdgercPath:   "/path/to/edgerc",
			// Missing AkamaiServiceConsumerDomain
		},
		{
			LogFormat:                   "json",
			Sources:                     []string{"test-source"},
			Provider:                    "akamai",
			AnnotationPrefix:            "external-dns.alpha.kubernetes.io/",
			AkamaiServiceConsumerDomain: "test-domain",
			AkamaiClientSecret:          "test-secret",
			AkamaiAccessToken:           "test-access-token",
			AkamaiEdgercPath:            "/path/to/edgerc",
			// Missing AkamaiClientToken
		},
		{
			LogFormat:                   "json",
			Sources:                     []string{"test-source"},
			Provider:                    "akamai",
			AnnotationPrefix:            "external-dns.alpha.kubernetes.io/",
			AkamaiServiceConsumerDomain: "test-domain",
			AkamaiClientToken:           "test-token",
			AkamaiAccessToken:           "test-access-token",
			AkamaiEdgercPath:            "/path/to/edgerc",
			// Missing AkamaiClientSecret
		},
		{
			LogFormat:                   "json",
			Sources:                     []string{"test-source"},
			Provider:                    "akamai",
			AnnotationPrefix:            "external-dns.alpha.kubernetes.io/",
			AkamaiServiceConsumerDomain: "test-domain",
			AkamaiClientToken:           "test-token",
			AkamaiClientSecret:          "test-secret",
			AkamaiEdgercPath:            "/path/to/edgerc",
			// Missing AkamaiAccessToken
		},
	}

	for _, cfg := range invalidAkamaiConfigs {
		err := ValidateConfig(cfg)
		assert.Error(t, err)
	}
}

func TestValidateGoodAkamaiConfig(t *testing.T) {
	validAkamaiConfigs := []*externaldns.Config{
		{
			LogFormat:                   "json",
			Sources:                     []string{"test-source"},
			Provider:                    "akamai",
			AnnotationPrefix:            "external-dns.alpha.kubernetes.io/",
			AkamaiServiceConsumerDomain: "test-domain",
			AkamaiClientToken:           "test-token",
			AkamaiClientSecret:          "test-secret",
			AkamaiAccessToken:           "test-access-token",
			AkamaiEdgercPath:            "/path/to/edgerc",
			KubeAPIQPS:                  int(rest.DefaultQPS),
			KubeAPIBurst:                rest.DefaultBurst,
		},
		{
			LogFormat:        "json",
			Sources:          []string{"test-source"},
			Provider:         "akamai",
			AnnotationPrefix: "external-dns.alpha.kubernetes.io/",
			KubeAPIQPS:       int(rest.DefaultQPS),
			KubeAPIBurst:     rest.DefaultBurst,
		},
	}

	for _, cfg := range validAkamaiConfigs {
		err := ValidateConfig(cfg)
		assert.NoError(t, err)
	}
}

func TestValidateBadAzureConfig(t *testing.T) {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "azure"
	cfg.AnnotationPrefix = "external-dns.alpha.kubernetes.io/"
	// AzureConfigFile is empty

	err := ValidateConfig(cfg)

	assert.Error(t, err)
}

func TestValidateGoodAzureConfig(t *testing.T) {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "azure"
	cfg.AnnotationPrefix = "external-dns.alpha.kubernetes.io/"
	cfg.AzureConfigFile = "/path/to/azure.json"
	cfg.KubeAPIQPS = int(rest.DefaultQPS)
	cfg.KubeAPIBurst = rest.DefaultBurst

	err := ValidateConfig(cfg)

	assert.NoError(t, err)
}

func TestValidateCreatePTRRequiresManagedRecordType(t *testing.T) {
	cfg := newValidConfig(t)
	cfg.CreatePTR = true
	// ManagedDNSRecordTypes defaults to [A, AAAA, CNAME] — no PTR

	err := ValidateConfig(cfg)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "--create-ptr requires PTR in --managed-record-types")
}

func TestValidateCreatePTRWithPTRManagedPasses(t *testing.T) {
	cfg := newValidConfig(t)
	cfg.CreatePTR = true
	cfg.ManagedDNSRecordTypes = append(cfg.ManagedDNSRecordTypes, "PTR")

	err := ValidateConfig(cfg)
	assert.NoError(t, err)
}
