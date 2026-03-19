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
	"context"

	"sigs.k8s.io/external-dns/endpoint"
)

// emptySource is a Source that returns no endpoints.
//
// +externaldns:source:name=empty
// +externaldns:source:category=Testing
// +externaldns:source:description=Returns no endpoints (used for testing or as a placeholder)
// +externaldns:source:resources=None
// +externaldns:source:filters=
// +externaldns:source:namespace=
// +externaldns:source:fqdn-template=false
// +externaldns:source:provider-specific=false
type emptySource struct{}

func (e *emptySource) AddEventHandler(_ context.Context, _ func()) {
}

// Endpoints collects endpoints of all nested Sources and returns them in a single slice.
func (e *emptySource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	return []*endpoint.Endpoint{}, nil
}

// NewEmptySource creates a new emptySource.
func NewEmptySource() Source {
	return &emptySource{}
}
