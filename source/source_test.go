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

package source

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockInformerFactory struct {
	syncResults map[reflect.Type]bool
}

func (m *mockInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	return m.syncResults
}

func TestWaitForCacheSync(t *testing.T) {
	tests := []struct {
		name        string
		syncResults map[reflect.Type]bool
		expectError bool
	}{
		{
			name:        "all caches synced",
			syncResults: map[reflect.Type]bool{reflect.TypeOf(""): true},
			expectError: false,
		},
		{
			name:        "some caches not synced",
			syncResults: map[reflect.Type]bool{reflect.TypeOf(""): false},
			expectError: true,
		},
		{
			name:        "context timeout",
			syncResults: map[reflect.Type]bool{reflect.TypeOf(""): false},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
			defer cancel()

			factory := &mockInformerFactory{syncResults: tt.syncResults}
			err := waitForCacheSync(ctx, factory)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
