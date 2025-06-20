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

func TestWaitForCacheSync(t *testing.T) {
	tests := []struct {
		name        string
		syncResults map[reflect.Type]bool
		expectError bool
		errorMsg    string
	}{
		{
			name:        "all caches synced",
			syncResults: map[reflect.Type]bool{reflect.TypeOf(""): true},
		},
		{
			name:        "some caches not synced",
			syncResults: map[reflect.Type]bool{reflect.TypeOf(""): false},
			expectError: true,
			errorMsg:    "failed to sync string with timeout 1m0s",
		},
		{
			name:        "context timeout",
			syncResults: map[reflect.Type]bool{reflect.TypeOf(""): false},
			expectError: true,
			errorMsg:    "failed to sync string with timeout 1m0s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			factory := &mockInformerFactory{syncResults: tt.syncResults}
			err := WaitForCacheSync(ctx, factory)

			if tt.expectError {
				assert.Error(t, err)
				assert.Errorf(t, err, tt.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestWaitForDynamicCacheSync(t *testing.T) {
	tests := []struct {
		name        string
		syncResults map[schema.GroupVersionResource]bool
		expectError bool
		errorMsg    string
	}{
		{
			name:        "all caches synced",
			syncResults: map[schema.GroupVersionResource]bool{{}: true},
		},
		{
			name:        "some caches not synced",
			syncResults: map[schema.GroupVersionResource]bool{{}: false},
			expectError: true,
			errorMsg:    "failed to sync string with timeout 1m0s",
		},
		{
			name:        "context timeout",
			syncResults: map[schema.GroupVersionResource]bool{{}: false},
			expectError: true,
			errorMsg:    "failed to sync string with timeout 1m0s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			factory := &mockDynamicInformerFactory{syncResults: tt.syncResults}
			err := WaitForDynamicCacheSync(ctx, factory)

			if tt.expectError {
				assert.Error(t, err)
				assert.Errorf(t, err, tt.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
