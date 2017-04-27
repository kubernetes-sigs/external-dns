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
)

func TestValidateFlags(t *testing.T) {
	cfg := newValidConfig(t)
	if err := ValidateConfig(cfg); err != nil {
		t.Errorf("valid config should be valid: %s", err)
	}

	cfg = newValidConfig(t)
	cfg.LogFormat = "test"
	if err := ValidateConfig(cfg); err == nil {
		t.Errorf("unsupported log format should fail: %s", cfg.LogFormat)
	}

	cfg = newValidConfig(t)
	cfg.LogFormat = ""
	if err := ValidateConfig(cfg); err == nil {
		t.Errorf("unsupported log format should fail: %s", cfg.LogFormat)
	}

	for _, format := range []string{"text", "json"} {
		cfg = newValidConfig(t)
		cfg.LogFormat = format
		if err := ValidateConfig(cfg); err != nil {
			t.Errorf("supported log format: %s should not fail", format)
		}
	}

	cfg = newValidConfig(t)
	cfg.Sources = []string{}
	if err := ValidateConfig(cfg); err == nil {
		t.Error("missing at least one source should fail")
	}

	cfg = newValidConfig(t)
	cfg.Provider = ""
	if err := ValidateConfig(cfg); err == nil {
		t.Error("missing provider should fail")
	}
}

func newValidConfig(t *testing.T) *externaldns.Config {
	cfg := externaldns.NewConfig()

	cfg.LogFormat = "json"
	cfg.Sources = []string{"test-source"}
	cfg.Provider = "test-provider"

	if err := ValidateConfig(cfg); err != nil {
		t.Fatalf("newValidConfig should return valid config")
	}

	return cfg
}
