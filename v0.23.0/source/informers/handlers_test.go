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

package informers

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestDefaultEventHandler_AddFunc(t *testing.T) {
	tests := []struct {
		name     string
		obj      any
		expected bool
	}{
		{
			name:     "calls handler for unstructured object",
			obj:      &unstructured.Unstructured{},
			expected: true,
		},
		{
			name:     "does not call handler for unknown object",
			obj:      "not-unstructured",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called := false
			handler := DefaultEventHandler(func() { called = true })
			handler.OnAdd(tt.obj, true)
			if called != tt.expected {
				t.Errorf("handler called = %v, want %v", called, tt.expected)
			}
		})
	}
}
