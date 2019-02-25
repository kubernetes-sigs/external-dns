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
	"time"

	"github.com/stretchr/testify/mock"
	"sigs.k8s.io/external-dns/endpoint"
)

// MockSource returns mock endpoints.
type MockSource struct {
	mock.Mock
}

// Endpoints returns the desired mock endpoints.
func (m *MockSource) Endpoints() ([]*endpoint.Endpoint, error) {
	args := m.Called()

	endpoints := args.Get(0)
	if endpoints == nil {
		return nil, args.Error(1)
	}

	return endpoints.([]*endpoint.Endpoint), args.Error(1)
}

// AddEventHandler adds an event handler function that's called when sources that support such a thing have changed.
func (m *MockSource) AddEventHandler(handler func() error, stopChan <-chan struct{}, minInterval time.Duration) {
	// Execute callback handler no more than once per minInterval, until a message on stopChan is received.
	go func() {
		var lastCallbackTime time.Time
		for {
			select {
			case <-stopChan:
				return
			default:
				now := time.Now()
				if now.After(lastCallbackTime.Add(minInterval)) {
					handler()
					lastCallbackTime = time.Now()
				}
			}
		}
	}()
}
