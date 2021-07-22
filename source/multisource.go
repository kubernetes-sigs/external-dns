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

	"sigs.k8s.io/external-dns/endpoint"
)

// multiSource is a Source that merges the endpoints of its nested Sources.
type multiSource struct {
	children       []Source
	defaultTargets []string
}

// Endpoints collects endpoints of all nested Sources and returns them in a single slice.
func (ms *multiSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	result := []*endpoint.Endpoint{}

	for _, s := range ms.children {
		endpoints, err := s.Endpoints(ctx)
		if err != nil {
			return nil, err
		}
		if len(ms.defaultTargets) > 0 {
			for i := range endpoints {
				endpoints[i].Targets = ms.defaultTargets
			}
		}
		result = append(result, endpoints...)
	}

	return result, nil
}

func (ms *multiSource) AddEventHandler(ctx context.Context, handler func()) {
	for _, s := range ms.children {
		s.AddEventHandler(ctx, handler)
	}
}

// NewMultiSource creates a new multiSource.
func NewMultiSource(children []Source, defaultTargets []string) Source {
	return &multiSource{children: children, defaultTargets: defaultTargets}
}
