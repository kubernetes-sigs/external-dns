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
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/types"
)

func makePod(annotations map[string]string) *corev1.Pod {
	p := &corev1.Pod{}
	p.SetName("test-pod")
	p.SetNamespace("default")
	p.SetAnnotations(annotations)
	return p
}

func TestIsValid(t *testing.T) {
	const testKeySuffix = "test-annotation"
	const testKey = "external-dns.kubernetes.io/" + testKeySuffix
	const otherKeySuffix = "other-annotation"
	const otherKey = "external-dns.kubernetes.io/" + otherKeySuffix

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
			registry:           []AnnotationSpec{{KeySuffix: testKeySuffix, Config: strictCfg}},
			source:             "pod",
			mode:               ModeWarn,
			wantValid:          true,
			wantLogNotContains: []string{testKey},
		},
		{
			name:               "valid value is valid with no logs",
			registry:           []AnnotationSpec{{KeySuffix: testKeySuffix, Config: strictCfg}},
			podAnnotations:     map[string]string{testKey: "valid-value"},
			source:             "pod",
			mode:               ModeStrict,
			wantValid:          true,
			wantLogNotContains: []string{testKey},
		},
		{
			name:            "invalid value in warn mode is valid but warns",
			registry:        []AnnotationSpec{{KeySuffix: testKeySuffix, Config: strictCfg}},
			podAnnotations:  map[string]string{testKey: "invalid-value"},
			source:          "pod",
			mode:            ModeWarn,
			wantValid:       true,
			wantLogContains: []string{"falls back to default"},
		},
		{
			name:            "invalid value in strict mode with StrictMessage is invalid and warns",
			registry:        []AnnotationSpec{{KeySuffix: testKeySuffix, Config: strictCfg}},
			podAnnotations:  map[string]string{testKey: "invalid-value"},
			source:          "pod",
			mode:            ModeStrict,
			wantValid:       false,
			wantLogContains: []string{"Excluding", "excluded entirely"},
		},
		{
			name:            "invalid value in strict mode without StrictMessage falls back to warn",
			registry:        []AnnotationSpec{{KeySuffix: testKeySuffix, Config: warnOnlyCfg}},
			podAnnotations:  map[string]string{testKey: "invalid-value"},
			source:          "pod",
			mode:            ModeStrict,
			wantValid:       true,
			wantLogContains: []string{"falls back to default"},
		},
		{
			name: "value present but source not supported logs debug without affecting validity",
			registry: []AnnotationSpec{{
				KeySuffix:        testKeySuffix,
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
				KeySuffix:        testKeySuffix,
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
				KeySuffix:        testKeySuffix,
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
			registry:           []AnnotationSpec{{KeySuffix: testKeySuffix, Config: strictCfg}},
			podAnnotations:     map[string]string{testKey: "valid-value"},
			source:             "any-source",
			mode:               ModeWarn,
			wantValid:          true,
			wantLogNotContains: []string{"not supported for source"},
		},
		{
			name: "multiple entries are all checked, not short-circuited",
			registry: []AnnotationSpec{
				{KeySuffix: testKeySuffix, Config: strictCfg},
				{KeySuffix: otherKeySuffix, Config: strictCfg},
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
			s := NewWithSpecs(tt.registry, Options{Mode: tt.mode}).WithSource(tt.source)
			hook := logtest.LogsUnderTestWithLogLevel(log.DebugLevel, t)

			valid := s.IsValid(makePod(tt.podAnnotations))

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

func TestRegistryFollowsCustomAnnotationPrefix(t *testing.T) {
	annotations.SetAnnotationPrefix("custom.io/")
	t.Cleanup(func() { annotations.SetAnnotationPrefix(annotations.DefaultAnnotationPrefix) })

	s := New(Options{Mode: ModeStrict}).WithSource("pod")
	assert.False(t, s.IsValid(makePod(map[string]string{annotations.HostnameKey: "bad..name"})))
	assert.True(t, s.IsValid(makePod(map[string]string{annotations.DefaultAnnotationPrefix + "hostname": "bad..name"})))
}

func TestIgnoreHostnameAnnotation(t *testing.T) {
	tests := []struct {
		name        string
		source      string
		annotations map[string]string
		opts        Options
		wantValid   bool
	}{
		{
			name:        "flag off: strict mode still drops",
			source:      "service",
			annotations: map[string]string{annotations.HostnameKey: "bad!.example.org"},
			opts:        Options{Mode: ModeStrict},
			wantValid:   false,
		},
		{
			name:   "flag on, service: hostname and internal-hostname both ignored",
			source: "service",
			annotations: map[string]string{
				annotations.HostnameKey:         "bad!.example.org",
				annotations.InternalHostnameKey: "bad!.internal.example.org",
			},
			opts:      Options{Mode: ModeStrict, IgnoreHostnameAnnotation: true},
			wantValid: true,
		},
		{
			name:        "flag on, ingress: hostname ignored",
			source:      "ingress",
			annotations: map[string]string{annotations.HostnameKey: "bad!.example.org"},
			opts:        Options{Mode: ModeStrict, IgnoreHostnameAnnotation: true},
			wantValid:   true,
		},
		{
			name:        "flag on, ingress: internal-hostname NOT ignored (ingress doesn't honor the flag for it)",
			source:      "ingress",
			annotations: map[string]string{annotations.InternalHostnameKey: "bad!.internal.example.org"},
			opts:        Options{Mode: ModeStrict, IgnoreHostnameAnnotation: true},
			wantValid:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.opts).WithSource(tt.source)
			assert.Equal(t, tt.wantValid, s.IsValid(makePod(tt.annotations)))
		})
	}
}

func makeService(annotations map[string]string) *corev1.Service {
	s := &corev1.Service{}
	s.SetName("test-service")
	s.SetNamespace("default")
	s.SetAnnotations(annotations)
	return s
}

// An invalid nodeport-filter falls back to publishing every node, so even strict mode must keep the
// service: excluding it would delete its records under --policy=sync.
// Not parallel: the log hook is global.
func TestNodePortFilterAnnotationNeverExcludesService(t *testing.T) {
	hook := logtest.LogsUnderTestWithLogLevel(log.WarnLevel, t)
	svc := makeService(map[string]string{annotations.NodePortFilterKey: "env in prod"})
	for _, mode := range []Mode{ModeWarn, ModeStrict} {
		assert.True(t, New(Options{Mode: mode}).WithSource(types.Service).IsValid(svc), mode)
	}
	// Unregistered keys are skipped silently, so the warning shows the annotation is validated.
	logtest.TestHelperLogContains(`"env in prod" is not a valid external-dns.kubernetes.io/nodeport-filter value`, hook, t)
}
