/*
Copyright 2025 The Kubernetes Authors.

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

package utils

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestCoreServiceAddFuncEventHandler(t *testing.T) {
	tests := []struct {
		name           string
		obj            interface{}
		expectedLogMsg string
		logLevel       log.Level
	}{
		{
			name: "valid service object",
			obj: &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name:      "test-service",
				},
			},
			expectedLogMsg: "event handler added for 'service/v1' in 'namespace:default' with 'name:test-service'.",
			logLevel:       log.DebugLevel,
		},
		{
			name:           "invalid object type",
			obj:            "invalid object",
			expectedLogMsg: "event handler not added. want 'service/v1' got 'string'",
			logLevel:       log.DebugLevel,
		},
		{
			name:           "on info level no logging",
			obj:            "invalid object",
			expectedLogMsg: "",
			logLevel:       log.InfoLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := testutils.LogsToBuffer(tt.logLevel, t)
			CoreServiceAddFuncEventHandler(tt.obj)
			if tt.expectedLogMsg != "" {
				assert.Contains(t, string(buf.Bytes()), tt.expectedLogMsg)
			} else {
				assert.Empty(t, string(buf.Bytes()))
			}
		})
	}
}
