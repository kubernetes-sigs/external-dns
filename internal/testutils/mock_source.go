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

package testutils

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"

	"sigs.k8s.io/external-dns/endpoint"
)

// MockSource returns mock endpoints.
type MockSource struct {
	mock.Mock
}

// Endpoints returns the desired mock endpoints.
func (m *MockSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	args := m.Called()

	endpoints := args.Get(0)
	if endpoints == nil {
		return nil, args.Error(1)
	}

	return endpoints.([]*endpoint.Endpoint), args.Error(1)
}

// AddEventHandler adds an event handler that should be triggered if something in source changes
func (m *MockSource) AddEventHandler(ctx context.Context, handler func()) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				handler()
			}
		}
	}()
}
