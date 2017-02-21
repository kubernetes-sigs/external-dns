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

package config

import (
	"testing"
)

func TestValidateFlags(t *testing.T) {
	cfg := NewConfig()
	cfg.LogFormat = "test"
	if err := cfg.Validate(); err == nil {
		t.Errorf("unsupported log format should fail: %s", cfg.LogFormat)
	}
	cfg.LogFormat = ""
	if err := cfg.Validate(); err == nil {
		t.Errorf("unsupported log format should fail: %s", cfg.LogFormat)
	}
	for _, format := range []string{"text", "json"} {
		cfg.LogFormat = format
		if err := cfg.Validate(); err != nil {
			t.Errorf("supported log format: %s should not fail", format)
		}
	}
}
