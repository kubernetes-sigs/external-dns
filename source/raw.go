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
	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// rawSource
type rawSource struct {
	endpoints []*endpoint.Endpoint
}

// NewRawSource creates a new rawSource with the given config.
func NewRawSource(endpoints []*endpoint.Endpoint) Source {
	return &rawSource{
		endpoints: endpoints,
	}
}

// Endpoints returns endpoint objects for each txt record.
func (sc *rawSource) Endpoints() ([]*endpoint.Endpoint, error) {
	return sc.endpoints, nil
}
