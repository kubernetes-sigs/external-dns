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

package externaldns

import (
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	defaultConfig = &Config{
		InCluster:      false,
		KubeConfig:     "",
		Namespace:      "",
		Domain:         "",
		Sources:        nil,
		Provider:       "",
		GoogleProject:  "",
		Policy:         "sync",
		Compatibility:  "",
		MetricsAddress: defaultMetricsAddress,
		Interval:       time.Minute,
		Once:           false,
		DryRun:         false,
		Debug:          false,
		LogFormat:      defaultLogFormat,
		Registry:       "txt",
		RecordOwnerID:  "default",
		TXTPrefix:      "",
		FqdnTemplate:   "",
	}
	overriddenConfig = &Config{
		InCluster:      true,
		KubeConfig:     "/some/path",
		Namespace:      "namespace",
		Domain:         "example.org.",
		Sources:        []string{"source1", "source2"},
		Provider:       "provider",
		GoogleProject:  "project",
		Policy:         "upsert-only",
		Compatibility:  "mate",
		MetricsAddress: "127.0.0.1:9099",
		Interval:       10 * time.Minute,
		Once:           true,
		DryRun:         true,
		Debug:          true,
		LogFormat:      "yaml",
		Registry:       "noop",
		RecordOwnerID:  "owner-1",
		TXTPrefix:      "associated-txt-record",
		FqdnTemplate:   "{{.Name}}.service.example.com",
	}
)

func TestParseFlags(t *testing.T) {
	for _, ti := range []struct {
		title    string
		args     []string
		envVars  map[string]string
		expected *Config
	}{
		{
			title:    "default config without anything defined",
			args:     []string{},
			envVars:  map[string]string{},
			expected: defaultConfig,
		},
		{
			title: "override everything via flags",
			args: []string{
				"--in-cluster",
				"--kubeconfig=/some/path",
				"--namespace=namespace",
				"--domain=example.org.",
				"--source=source1",
				"--source=source2",
				"--provider=provider",
				"--google-project=project",
				"--policy=upsert-only",
				"--compatibility=mate",
				"--metrics-address=127.0.0.1:9099",
				"--interval=10m",
				"--once",
				"--dry-run",
				"--debug",
				"--log-format=yaml",
				"--registry=noop",
				"--record-owner-id=owner-1",
				"--txt-prefix=associated-txt-record",
				"--fqdn-template={{.Name}}.service.example.com",
			},
			envVars:  map[string]string{},
			expected: overriddenConfig,
		},
		{
			title: "override everything via environment variables",
			args:  []string{},
			envVars: map[string]string{
				"EXTERNAL_DNS_IN_CLUSTER":      "1",
				"EXTERNAL_DNS_KUBECONFIG":      "/some/path",
				"EXTERNAL_DNS_NAMESPACE":       "namespace",
				"EXTERNAL_DNS_DOMAIN":          "example.org.",
				"EXTERNAL_DNS_SOURCE":          "source1\nsource2",
				"EXTERNAL_DNS_PROVIDER":        "provider",
				"EXTERNAL_DNS_GOOGLE_PROJECT":  "project",
				"EXTERNAL_DNS_POLICY":          "upsert-only",
				"EXTERNAL_DNS_COMPATIBILITY":   "mate",
				"EXTERNAL_DNS_METRICS_ADDRESS": "127.0.0.1:9099",
				"EXTERNAL_DNS_INTERVAL":        "10m",
				"EXTERNAL_DNS_ONCE":            "1",
				"EXTERNAL_DNS_DRY_RUN":         "1",
				"EXTERNAL_DNS_DEBUG":           "1",
				"EXTERNAL_DNS_LOG_FORMAT":      "yaml",
				"EXTERNAL_DNS_REGISTRY":        "noop",
				"EXTERNAL_DNS_RECORD_OWNER_ID": "owner-1",
				"EXTERNAL_DNS_TXT_PREFIX":      "associated-txt-record",
				"EXTERNAL_DNS_FQDN_TEMPLATE":   "{{.Name}}.service.example.com",
			},
			expected: overriddenConfig,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			originalEnv := setEnv(t, ti.envVars)
			defer func() {
				restoreEnv(t, originalEnv)
			}()

			cfg := NewConfig()

			if err := cfg.ParseFlags(ti.args); err != nil {
				t.Error(err)
			}

			validateConfig(t, cfg, ti.expected)
		})
	}
}

// helper functions

func validateConfig(t *testing.T, got, expected *Config) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("config is wrong")
	}
}

func setEnv(t *testing.T, env map[string]string) map[string]string {
	originalEnv := map[string]string{}

	for k, v := range env {
		originalEnv[k] = os.Getenv(k)

		if err := os.Setenv(k, v); err != nil {
			t.Fatal(err)
		}
	}

	return originalEnv
}

func restoreEnv(t *testing.T, originalEnv map[string]string) {
	for k, v := range originalEnv {
		if err := os.Setenv(k, v); err != nil {
			t.Fatal(err)
		}
	}
}
