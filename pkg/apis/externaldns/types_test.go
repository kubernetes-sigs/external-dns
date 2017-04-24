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
	"reflect"
	"testing"
	"time"
)

func TestParseFlags(t *testing.T) {
	for _, ti := range []struct {
		title       string
		args        [][]string
		expectError bool
		expected    *Config
	}{
		{
			title: "set in-cluster true",
			args:  [][]string{{"--in-cluster", ""}},
			expected: &Config{
				InCluster:      true,
				KubeConfig:     "",
				Namespace:      "",
				Zone:           "",
				Domain:         "example.org.",
				Sources:        nil,
				Provider:       "",
				GoogleProject:  "",
				Policy:         "sync",
				Compatibility:  "",
				MetricsAddress: defaultMetricsAddress,
				Interval:       time.Minute,
				Once:           false,
				DryRun:         true,
				Debug:          false,
				LogFormat:      defaultLogFormat,
				Version:        false,
				Registry:       "noop",
				RecordOwnerID:  "",
				TXTPrefix:      "",
				FqdnTemplate:   "",
			},
		},
		{
			title: "all default",
			args:  [][]string{},
			expected: &Config{
				InCluster:      false,
				KubeConfig:     "",
				Namespace:      "",
				Zone:           "",
				Domain:         "example.org.",
				Sources:        nil,
				Provider:       "",
				GoogleProject:  "",
				Policy:         "sync",
				Compatibility:  "",
				MetricsAddress: defaultMetricsAddress,
				Interval:       time.Minute,
				Once:           false,
				DryRun:         true,
				Debug:          false,
				LogFormat:      defaultLogFormat,
				Version:        false,
				Registry:       "noop",
				RecordOwnerID:  "",
				TXTPrefix:      "",
				FqdnTemplate:   "",
			},
		},
		{
			title: "set string var",
			args:  [][]string{{"--kubeconfig", "myhome"}},
			expected: &Config{
				InCluster:      false,
				KubeConfig:     "myhome",
				Namespace:      "",
				Zone:           "",
				Domain:         "example.org.",
				Sources:        nil,
				Provider:       "",
				GoogleProject:  "",
				Policy:         "sync",
				Compatibility:  "",
				MetricsAddress: defaultMetricsAddress,
				Interval:       time.Minute,
				Once:           false,
				DryRun:         true,
				Debug:          false,
				LogFormat:      defaultLogFormat,
				Version:        false,
				Registry:       "noop",
				RecordOwnerID:  "",
				TXTPrefix:      "",
				FqdnTemplate:   "",
			},
		},
		{
			title:       "unexpected flag",
			args:        [][]string{{"--random", "myhome"}},
			expectError: true,
		},
		{
			title: "override default",
			args:  [][]string{{"--log-format", "json"}},
			expected: &Config{
				InCluster:      false,
				KubeConfig:     "",
				Namespace:      "",
				Zone:           "",
				Domain:         "example.org.",
				Sources:        nil,
				Provider:       "",
				GoogleProject:  "",
				Policy:         "sync",
				Compatibility:  "",
				MetricsAddress: defaultMetricsAddress,
				Interval:       time.Minute,
				Once:           false,
				DryRun:         true,
				Debug:          false,
				LogFormat:      "json",
				Version:        false,
				Registry:       "noop",
				RecordOwnerID:  "",
				TXTPrefix:      "",
				FqdnTemplate:   "",
			},
		},
		{
			title: "set everything",
			args: [][]string{{"--in-cluster",
				"--log-format", "yaml",
				"--kubeconfig", "/some/path",
				"--namespace", "namespace",
				"--zone", "zone",
				"--domain", "kubernetes.io.",
				"--source", "source",
				"--provider", "provider",
				"--google-project", "project",
				"--policy", "upsert-only",
				"--compatibility=mate",
				"--metrics-address", "127.0.0.1:9099",
				"--interval", "10m",
				"--once",
				"--dry-run=false",
				"--debug",
				"--registry=txt",
				"--record-owner-id=owner-1",
				"--txt-prefix=associated-txt-record",
				"--fqdn-template={{.Name}}.service.example.com",
				"--version"}},
			expected: &Config{
				InCluster:      true,
				KubeConfig:     "/some/path",
				Namespace:      "namespace",
				Zone:           "zone",
				Domain:         "kubernetes.io.",
				Sources:        []string{"source"},
				Provider:       "provider",
				GoogleProject:  "project",
				Policy:         "upsert-only",
				Compatibility:  "mate",
				MetricsAddress: "127.0.0.1:9099",
				Interval:       10 * time.Minute,
				Once:           true,
				DryRun:         false,
				Debug:          true,
				LogFormat:      "yaml",
				Version:        true,
				Registry:       "txt",
				RecordOwnerID:  "owner-1",
				TXTPrefix:      "associated-txt-record",
				FqdnTemplate:   "{{.Name}}.service.example.com",
			},
		},
		{
			title:       "--help trigger error",
			args:        [][]string{{"--help"}},
			expectError: true,
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			cfg := NewConfig()
			spaceArgs := []string{"external-dns"}
			for _, arg := range ti.args {
				spaceArgs = append(spaceArgs, arg...)
			}
			err := cfg.ParseFlags(spaceArgs)
			if !ti.expectError && err != nil {
				t.Errorf("unexpected parse flags fail for args %#v, error: %v", ti.args, err)
			}
			if ti.expectError && err == nil {
				t.Errorf("parse flags should fail for args %#v", ti.args)
			}
			if !ti.expectError {
				validateConfig(t, cfg, ti.expected)
			}
		})
	}
}

// helper functions

func validateConfig(t *testing.T, got, expected *Config) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("config is wrong")
	}
}
