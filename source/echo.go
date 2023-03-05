/*
Copyright 2023 The Kubernetes Authors.

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

	"sigs.k8s.io/external-dns/endpoint"
)

// echoSource is a Source that returns the endpoints passed in on creation.
type echoSource struct {
	endpoints []*endpoint.Endpoint
}

func (e *echoSource) AddEventHandler(ctx context.Context, handler func()) {
}

// Endpoints returns all of the endpoints passed in on creation
func (e *echoSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return e.endpoints, nil
}

// NewEchoSource creates a new echoSource.
func NewEchoSource(endpoints []*endpoint.Endpoint) Source {
	return &echoSource{endpoints: endpoints}
}
