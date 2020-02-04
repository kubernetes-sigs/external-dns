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
	"time"

	"sigs.k8s.io/external-dns/endpoint"
)

// multiSource is a Source that merges the endpoints of its nested Sources.
type multiSource struct {
	children []Source
}

// Endpoints collects endpoints of all nested Sources and returns them in a single slice.
func (ms *multiSource) Endpoints() ([]*endpoint.Endpoint, error) {
	result := []*endpoint.Endpoint{}

	for _, s := range ms.children {
		endpoints, err := s.Endpoints()
		if err != nil {
			return nil, err
		}

		result = append(result, endpoints...)
	}

	return result, nil
}

func (ms *multiSource) AddEventHandler(handler func() error, stopChan <-chan struct{}, minInterval time.Duration) {
	for _, s := range ms.children {
		s.AddEventHandler(handler, stopChan, minInterval)
	}
}

// NewMultiSource creates a new multiSource.
func NewMultiSource(children []Source) Source {
	return &multiSource{children: children}
}
