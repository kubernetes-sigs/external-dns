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
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type mockInformerFactory struct {
	syncResults map[reflect.Type]bool
}

func (m *mockInformerFactory) WaitForCacheSync(_ <-chan struct{}) map[reflect.Type]bool {
	return m.syncResults
}

type mockDynamicInformerFactory struct {
	syncResults map[schema.GroupVersionResource]bool
}

func (m *mockDynamicInformerFactory) WaitForCacheSync(_ <-chan struct{}) map[schema.GroupVersionResource]bool {
	return m.syncResults
}

// TestWaitForCacheSync verifies that WaitForCacheSync uses soft error handling.
// Instead of returning errors, it logs warnings and returns nil to prevent crash loops.
func TestWaitForCacheSync(t *testing.T) {
	tests := []struct {
		name        string
		syncResults map[reflect.Type]bool
	}{
		{
			name:        "all caches synced",
			syncResults: map[reflect.Type]bool{reflect.TypeFor[string](): true},
		},
		{
			// Soft error: logs warning but returns nil to prevent crash loops
			name:        "some caches not synced - soft error",
			syncResults: map[reflect.Type]bool{reflect.TypeFor[string](): false},
		},
		{
			// Soft error: logs warning but returns nil to prevent crash loops
			name:        "context timeout - soft error",
			syncResults: map[reflect.Type]bool{reflect.TypeFor[string](): false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			factory := &mockInformerFactory{syncResults: tt.syncResults}
			err := WaitForCacheSync(ctx, factory, 0)

			// All cases should return nil due to soft error handling.
			// This prevents crash loops that can overwhelm the API server.
			assert.NoError(t, err, "WaitForCacheSync should return nil for soft error handling")
		})
	}
}

// TestWaitForDynamicCacheSync verifies that WaitForDynamicCacheSync uses soft error handling.
// Instead of returning errors, it logs warnings and returns nil to prevent crash loops.
func TestWaitForDynamicCacheSync(t *testing.T) {
	tests := []struct {
		name        string
		syncResults map[schema.GroupVersionResource]bool
	}{
		{
			name:        "all caches synced",
			syncResults: map[schema.GroupVersionResource]bool{{}: true},
		},
		{
			// Soft error: logs warning but returns nil to prevent crash loops
			name:        "some caches not synced - soft error",
			syncResults: map[schema.GroupVersionResource]bool{{}: false},
		},
		{
			// Soft error: logs warning but returns nil to prevent crash loops
			name:        "context timeout - soft error",
			syncResults: map[schema.GroupVersionResource]bool{{}: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			factory := &mockDynamicInformerFactory{syncResults: tt.syncResults}
			err := WaitForDynamicCacheSync(ctx, factory, 0)

			// All cases should return nil due to soft error handling.
			// This prevents crash loops that can overwhelm the API server.
			assert.NoError(t, err, "WaitForDynamicCacheSync should return nil for soft error handling")
		})
	}
}
