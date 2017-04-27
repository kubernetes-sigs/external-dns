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
	minimalConfig = &Config{
		Master:         "",
		KubeConfig:     "",
		Sources:        []string{"service"},
		Namespace:      "",
		FqdnTemplate:   "",
		Compatibility:  "",
		Provider:       "google",
		GoogleProject:  "",
		DomainFilter:   "",
		Policy:         "sync",
		Registry:       "txt",
		TXTOwnerID:     "default",
		TXTPrefix:      "",
		Interval:       time.Minute,
		Once:           false,
		DryRun:         false,
		LogFormat:      "text",
		MetricsAddress: ":7979",
		Debug:          false,
	}

	overriddenConfig = &Config{
		Master:         "http://127.0.0.1:8080",
		KubeConfig:     "/some/path",
		Sources:        []string{"service", "ingress"},
		Namespace:      "namespace",
		FqdnTemplate:   "{{.Name}}.service.example.com",
		Compatibility:  "mate",
		Provider:       "google",
		GoogleProject:  "project",
		DomainFilter:   "example.org.",
		Policy:         "upsert-only",
		Registry:       "noop",
		TXTOwnerID:     "owner-1",
		TXTPrefix:      "associated-txt-record",
		Interval:       10 * time.Minute,
		Once:           true,
		DryRun:         true,
		LogFormat:      "json",
		MetricsAddress: "127.0.0.1:9099",
		Debug:          true,
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
			title: "default config with minimal flags defined",
			args: []string{
				"--source=service",
				"--provider=google",
			},
			envVars:  map[string]string{},
			expected: minimalConfig,
		},
		{
			title: "override everything via flags",
			args: []string{
				"--master=http://127.0.0.1:8080",
				"--kubeconfig=/some/path",
				"--source=service",
				"--source=ingress",
				"--namespace=namespace",
				"--fqdn-template={{.Name}}.service.example.com",
				"--compatibility=mate",
				"--provider=google",
				"--google-project=project",
				"--domain-filter=example.org.",
				"--policy=upsert-only",
				"--registry=noop",
				"--txt-owner-id=owner-1",
				"--txt-prefix=associated-txt-record",
				"--interval=10m",
				"--once",
				"--dry-run",
				"--log-format=json",
				"--metrics-address=127.0.0.1:9099",
				"--debug",
			},
			envVars:  map[string]string{},
			expected: overriddenConfig,
		},
		{
			title: "override everything via environment variables",
			args:  []string{},
			envVars: map[string]string{
				"EXTERNAL_DNS_MASTER":          "http://127.0.0.1:8080",
				"EXTERNAL_DNS_KUBECONFIG":      "/some/path",
				"EXTERNAL_DNS_SOURCE":          "service\ningress",
				"EXTERNAL_DNS_NAMESPACE":       "namespace",
				"EXTERNAL_DNS_FQDN_TEMPLATE":   "{{.Name}}.service.example.com",
				"EXTERNAL_DNS_COMPATIBILITY":   "mate",
				"EXTERNAL_DNS_PROVIDER":        "google",
				"EXTERNAL_DNS_GOOGLE_PROJECT":  "project",
				"EXTERNAL_DNS_DOMAIN_FILTER":   "example.org.",
				"EXTERNAL_DNS_POLICY":          "upsert-only",
				"EXTERNAL_DNS_REGISTRY":        "noop",
				"EXTERNAL_DNS_TXT_OWNER_ID":    "owner-1",
				"EXTERNAL_DNS_TXT_PREFIX":      "associated-txt-record",
				"EXTERNAL_DNS_INTERVAL":        "10m",
				"EXTERNAL_DNS_ONCE":            "1",
				"EXTERNAL_DNS_DRY_RUN":         "1",
				"EXTERNAL_DNS_LOG_FORMAT":      "json",
				"EXTERNAL_DNS_METRICS_ADDRESS": "127.0.0.1:9099",
				"EXTERNAL_DNS_DEBUG":           "1",
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
