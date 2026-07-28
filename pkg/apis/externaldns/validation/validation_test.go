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
	cfg.FQDNTemplate = []string{}
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
	cfg.AnnotationPrefix = "external-dns.kubernetes.io/"
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
	cfg.FQDNTemplate = []string{}

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
			AnnotationPrefix:        "external-dns.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.kubernetes.io/",
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
			AnnotationPrefix:        "external-dns.kubernetes.io/",
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

func TestValidateBadAzureConfig(t *testing.T) {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "azure"
	cfg.AnnotationPrefix = "external-dns.kubernetes.io/"
	// AzureConfigFile is empty

	err := ValidateConfig(cfg)

	assert.Error(t, err)
}

func TestValidateGoodAzureConfig(t *testing.T) {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "azure"
	cfg.AnnotationPrefix = "external-dns.kubernetes.io/"
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

func TestValidateHostnameConfig(t *testing.T) {
	tests := []struct {
		name             string
		ignoreAnnotation bool
		fqdnTemplate     []string
		wantErr          bool
	}{
		{"not ignoring annotations", false, nil, false},
		{"ignoring annotations with template", true, []string{"{{.Name}}"}, false},
		{"ignoring annotations without template", true, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &externaldns.Config{IgnoreHostnameAnnotation: tt.ignoreAnnotation, FQDNTemplate: tt.fqdnTemplate}
			if tt.wantErr {
				require.NotEmpty(t, validateHostnameConfig(cfg))
			} else {
				require.Empty(t, validateHostnameConfig(cfg))
			}
		})
	}
}

func TestValidateTXTRegistryConfig(t *testing.T) {
	tests := []struct {
		name    string
		prefix  string
		suffix  string
		wantErr bool
	}{
		{"neither set", "", "", false},
		{"prefix only", "pre-", "", false},
		{"suffix only", "", "-suf", false},
		{"both set", "pre-", "-suf", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &externaldns.Config{TXTPrefix: tt.prefix, TXTSuffix: tt.suffix}
			if tt.wantErr {
				require.NotEmpty(t, validateTXTRegistryConfig(cfg))
			} else {
				require.Empty(t, validateTXTRegistryConfig(cfg))
			}
		})
	}
}

func TestValidateLabelSelectors(t *testing.T) {
	tests := []struct {
		name             string
		labelFilter      string
		annotationFilter string
		wantErr          bool
	}{
		{"both empty", "", "", false},
		{"valid label filter", "foo=bar", "", false},
		{"invalid label filter", "#invalid-selector", "", true},
		{"invalid annotation filter", "", "kubernetes.io/gateway.name in (a b)", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &externaldns.Config{LabelFilter: tt.labelFilter, AnnotationFilter: tt.annotationFilter}
			if tt.wantErr {
				require.NotEmpty(t, validateLabelSelectors(cfg))
			} else {
				require.Empty(t, validateLabelSelectors(cfg))
			}
		})
	}
}

func TestValidateAnnotationPrefix(t *testing.T) {
	tests := []struct {
		name    string
		prefix  string
		wantErr bool
	}{
		{"valid", "external-dns.kubernetes.io/", false},
		{"empty", "", true},
		{"missing trailing slash", "custom.io", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &externaldns.Config{AnnotationPrefix: tt.prefix}
			if tt.wantErr {
				require.NotEmpty(t, validateAnnotationPrefix(cfg))
			} else {
				require.Empty(t, validateAnnotationPrefix(cfg))
			}
		})
	}
}

func TestValidateKubeAPILimits(t *testing.T) {
	tests := []struct {
		name    string
		qps     int
		burst   int
		wantErr bool
	}{
		{"valid", 5, 10, false},
		{"zero qps", 0, 10, true},
		{"negative qps", -1, 10, true},
		{"zero burst", 5, 0, true},
		{"negative burst", 5, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &externaldns.Config{KubeAPIQPS: tt.qps, KubeAPIBurst: tt.burst}
			if tt.wantErr {
				require.NotEmpty(t, validateKubeAPILimits(cfg))
			} else {
				require.Empty(t, validateKubeAPILimits(cfg))
			}
		})
	}
}

func TestValidatePTRConfig(t *testing.T) {
	t.Run("create-ptr disabled", func(t *testing.T) {
		cfg := newValidConfig(t)
		cfg.CreatePTR = false
		require.Empty(t, validatePTRConfig(cfg))
	})
	t.Run("create-ptr without PTR managed", func(t *testing.T) {
		cfg := newValidConfig(t)
		cfg.CreatePTR = true
		require.NotEmpty(t, validatePTRConfig(cfg))
	})
	t.Run("create-ptr with PTR managed", func(t *testing.T) {
		cfg := newValidConfig(t)
		cfg.CreatePTR = true
		cfg.ManagedDNSRecordTypes = append(cfg.ManagedDNSRecordTypes, "PTR")
		require.Empty(t, validatePTRConfig(cfg))
	})
}
