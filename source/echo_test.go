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
	"testing"

	"sigs.k8s.io/external-dns/endpoint"
)

func TestEchoSourceReturnGivenSources(t *testing.T) {
	startEndpoints := []*endpoint.Endpoint{{
		DNSName:    "foo.bar.com",
		RecordType: "A",
		Targets:    endpoint.Targets{"1.2.3.4"},
		RecordTTL:  endpoint.TTL(300),
		Labels:     endpoint.Labels{},
	}}
	e := NewEchoSource(startEndpoints)

	endpoints, err := e.Endpoints(context.Background())
	if err != nil {
		t.Errorf("Expected no error but got %s", err.Error())
	}

	for i, endpoint := range endpoints {
		if endpoint != startEndpoints[i] {
			t.Errorf("Expected %s but got %s", startEndpoints[i], endpoint)
		}
	}
}
