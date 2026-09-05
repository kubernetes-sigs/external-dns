/*
Copyright 2026 The Kubernetes Authors.
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

package schema

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"

	logtest "sigs.k8s.io/external-dns/internal/testutils/log"
)

func withRegistry(t *testing.T, entries ...AnnotationSpec) {
	t.Helper()
	orig := Registry
	Registry = entries
	t.Cleanup(func() { Registry = orig })
}

func makePod(annotations map[string]string) *corev1.Pod {
	p := &corev1.Pod{}
	p.SetName("test-pod")
	p.SetNamespace("default")
	p.SetAnnotations(annotations)
	return p
}

func TestIsValid(t *testing.T) {
	const testKey = "external-dns.kubernetes.io/test-annotation"
	const otherKey = "external-dns.kubernetes.io/other-annotation"

	strictCfg := Config{
		Validators:    []Validator{ValidateOneOf("valid-value")},
		WarnMessage:   "falls back to default",
		StrictMessage: "excluded entirely",
	}
	warnOnlyCfg := Config{
		Validators:  []Validator{ValidateOneOf("valid-value")},
		WarnMessage: "falls back to default",
	}

	tests := []struct {
		name               string
		registry           []AnnotationSpec
		podAnnotations     map[string]string
		source             string
		mode               Mode
		wantValid          bool
		wantLogContains    []string
		wantLogNotContains []string
	}{
		{
			name:               "annotation absent is valid with no logs",
			registry:           []AnnotationSpec{{Key: testKey, Config: strictCfg}},
			source:             "pod",
			mode:               ModeWarn,
			wantValid:          true,
			wantLogNotContains: []string{testKey},
		},
		{
			name:               "valid value is valid with no logs",
			registry:           []AnnotationSpec{{Key: testKey, Config: strictCfg}},
			podAnnotations:     map[string]string{testKey: "valid-value"},
			source:             "pod",
			mode:               ModeStrict,
			wantValid:          true,
			wantLogNotContains: []string{testKey},
		},
		{
			name:            "invalid value in warn mode is valid but warns",
			registry:        []AnnotationSpec{{Key: testKey, Config: strictCfg}},
			podAnnotations:  map[string]string{testKey: "invalid-value"},
			source:          "pod",
			mode:            ModeWarn,
			wantValid:       true,
			wantLogContains: []string{"falls back to default"},
		},
		{
			name:            "invalid value in strict mode with StrictMessage is invalid and warns",
			registry:        []AnnotationSpec{{Key: testKey, Config: strictCfg}},
			podAnnotations:  map[string]string{testKey: "invalid-value"},
			source:          "pod",
			mode:            ModeStrict,
			wantValid:       false,
			wantLogContains: []string{"Excluding", "excluded entirely"},
		},
		{
			name:            "invalid value in strict mode without StrictMessage falls back to warn",
			registry:        []AnnotationSpec{{Key: testKey, Config: warnOnlyCfg}},
			podAnnotations:  map[string]string{testKey: "invalid-value"},
			source:          "pod",
			mode:            ModeStrict,
			wantValid:       true,
			wantLogContains: []string{"falls back to default"},
		},
		{
			name: "value present but source not supported logs debug without affecting validity",
			registry: []AnnotationSpec{{
				Key:              testKey,
				SupportedSources: []string{"service"},
				Config:           strictCfg,
			}},
			podAnnotations:  map[string]string{testKey: "valid-value"},
			source:          "pod",
			mode:            ModeWarn,
			wantValid:       true,
			wantLogContains: []string{"not supported for source"},
		},
		{
			name: "unsupported source in strict mode with StrictMessage excludes the object",
			registry: []AnnotationSpec{{
				Key:              testKey,
				SupportedSources: []string{"service"},
				Config:           strictCfg,
			}},
			podAnnotations:  map[string]string{testKey: "valid-value"},
			source:          "pod",
			mode:            ModeStrict,
			wantValid:       false,
			wantLogContains: []string{"Excluding", "not supported for source", "excluded entirely"},
		},
		{
			name: "unsupported source in strict mode without StrictMessage does not exclude",
			registry: []AnnotationSpec{{
				Key:              testKey,
				SupportedSources: []string{"service"},
				Config:           warnOnlyCfg,
			}},
			podAnnotations:     map[string]string{testKey: "valid-value"},
			source:             "pod",
			mode:               ModeStrict,
			wantValid:          true,
			wantLogContains:    []string{"not supported for source"},
			wantLogNotContains: []string{"Excluding"},
		},
		{
			name:               "empty SupportedSources enforces no restriction",
			registry:           []AnnotationSpec{{Key: testKey, Config: strictCfg}},
			podAnnotations:     map[string]string{testKey: "valid-value"},
			source:             "any-source",
			mode:               ModeWarn,
			wantValid:          true,
			wantLogNotContains: []string{"not supported for source"},
		},
		{
			name: "multiple entries are all checked, not short-circuited",
			registry: []AnnotationSpec{
				{Key: testKey, Config: strictCfg},
				{Key: otherKey, Config: strictCfg},
			},
			podAnnotations: map[string]string{
				testKey:  "invalid-value",
				otherKey: "valid-value",
			},
			source:          "pod",
			mode:            ModeStrict,
			wantValid:       false,
			wantLogContains: []string{testKey},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			withRegistry(t, tt.registry...)
			hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

			valid := IsValid(makePod(tt.podAnnotations), tt.source, tt.mode)

			assert.Equal(t, tt.wantValid, valid)
			for _, s := range tt.wantLogContains {
				logtest.TestHelperLogContains(s, hook, t)
			}
			for _, s := range tt.wantLogNotContains {
				logtest.TestHelperLogNotContains(s, hook, t)
			}
		})
	}
}
