/*
Copyright 2019 The Kubernetes Authors.

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
	"time"

	"sigs.k8s.io/external-dns/endpoint"
)

// emptySource is a Source that returns no endpoints.
type emptySource struct{}

func (e *emptySource) AddEventHandler(handler func() error, stopChan <-chan struct{}, minInterval time.Duration) {
}

// Endpoints collects endpoints of all nested Sources and returns them in a single slice.
func (e *emptySource) Endpoints() ([]*endpoint.Endpoint, error) {
	return []*endpoint.Endpoint{}, nil
}

// NewEmptySource creates a new emptySource.
func NewEmptySource() Source {
	return &emptySource{}
}
