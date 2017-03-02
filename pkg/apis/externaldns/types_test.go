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
	"fmt"
	"reflect"
	"testing"
)

func TestParseFlags(t *testing.T) {
	for _, ti := range []struct {
		title       string
		args        []arg
		expectError bool
		expected    *Config
	}{
		{
			title: "set in-cluster true",
			args: []arg{
				arg{
					flag: "--in-cluster",
					val:  "",
				},
			},
			expected: &Config{
				InCluster:     true,
				KubeConfig:    "",
				GoogleProject: "",
				GoogleZone:    "",
				HealthPort:    defaultHealthPort,
				DryRun:        true,
				Debug:         false,
				LogFormat:     defaultLogFormat,
			},
		},
		{
			title: "all default",
			args:  []arg{},
			expected: &Config{
				InCluster:     false,
				KubeConfig:    "",
				GoogleProject: "",
				GoogleZone:    "",
				HealthPort:    defaultHealthPort,
				DryRun:        true,
				Debug:         false,
				LogFormat:     defaultLogFormat,
			},
		},
		{
			title: "set string var",
			args: []arg{
				arg{
					flag: "--kubeconfig",
					val:  "myhome",
				},
			},
			expected: &Config{
				InCluster:     false,
				KubeConfig:    "myhome",
				GoogleProject: "",
				GoogleZone:    "",
				HealthPort:    defaultHealthPort,
				DryRun:        true,
				Debug:         false,
				LogFormat:     defaultLogFormat,
			},
		},
		{
			title: "unexpected flag",
			args: []arg{
				arg{
					flag: "--random",
					val:  "myhome",
				},
			},
			expectError: true,
		},
		{
			title: "override default",
			args: []arg{
				arg{
					flag: "--log-format",
					val:  "json",
				},
			},
			expected: &Config{
				InCluster:     false,
				KubeConfig:    "myhome",
				GoogleProject: "",
				GoogleZone:    "",
				HealthPort:    defaultHealthPort,
				DryRun:        true,
				Debug:         false,
				LogFormat:     "json",
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			cfg := NewConfig()
			spaceArgs := []string{"external-dns"}
			for _, arg := range ti.args {
				spaceArgs = append(spaceArgs, arg.Slice()...)
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
		t.Errorf("Config is wrong")
	}
}

type arg struct {
	flag string
	val  string
}

func (a *arg) Slice() []string {
	return []string{a.flag, a.val}
}

func (a *arg) String() string {
	return fmt.Sprintf("%s=%s", a.flag, a.val)
}
