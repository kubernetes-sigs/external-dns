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

import "github.com/kubernetes-incubator/external-dns/endpoint"

// multiSource is a Source that merges the endpoints of its nested Sources.
type multiSource struct {
	children []Source
}

// struct to allow for quickly creating distinct set of endpoints to return
type endpointKey struct {
	DNSName string
	TxtType bool
}

// Endpoints collects endpoints of all nested Sources and returns them in a single slice.
func (ms *multiSource) Endpoints() ([]*endpoint.Endpoint, error) {
	uniqueEndpoints := make(map[endpointKey]*endpoint.Endpoint)

	for _, s := range ms.children {
		endpoints, err := s.Endpoints()
		if err != nil {
			return nil, err
		}

		for _, endpoint := range endpoints {
			uniqueEndpoints[newEndpointKey(endpoint)] = endpoint
		}
	}

	result := []*endpoint.Endpoint{}
	for _, e := range uniqueEndpoints {
		result = append(result, e)
	}

	return result, nil
}

func newEndpointKey(endpoint *endpoint.Endpoint) endpointKey {
	return endpointKey{
		DNSName: endpoint.DNSName,
		TxtType: endpoint.RecordType == "TXT",
	}
}

// NewMultiSource creates a new multiSource.
func NewMultiSource(children []Source) Source {
	return &multiSource{children: children}
}
