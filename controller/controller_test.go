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

package controller

import (
	"errors"
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// mockSource returns mock endpoints.
type mockSource struct{}

// Endpoints returns a single desired test endpoint
func (s *mockSource) Endpoints() ([]endpoint.Endpoint, error) {
	endpoints := []endpoint.Endpoint{
		{
			DNSName: "test-record",
			Target:  "1.2.3.4",
		},
	}

	return endpoints, nil
}

// mockDNSProvider returns no current endpoints and validates that the applied
// list of endpoints is correct.
type mockDNSProvider struct{}

// Records returns an empty list of current endpoints.
func (p *mockDNSProvider) Records(zone string) ([]endpoint.Endpoint, error) {
	return []endpoint.Endpoint{}, nil
}

// ApplyChanges validates that the passed in changes satisfy a specifc assumtion.
func (p *mockDNSProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	if zone != "test-zone" {
		return errors.New("zone is incorrect")
	}

	if len(changes.Create) != 1 {
		return errors.New("number of created records is wrong")
	}

	create := changes.Create[0]

	if create.DNSName != "test-record" || create.Target != "1.2.3.4" {
		return errors.New("created record is wrong")
	}

	return nil
}

// TestRunOnce tests that RunOnce correctly orchestrates the different components.
func TestRunOnce(t *testing.T) {
	ctrl := &Controller{
		Zone: "test-zone",

		Source:      &mockSource{},
		DNSProvider: &mockDNSProvider{},
	}

	err := ctrl.RunOnce()

	if err != nil {
		t.Fatal(err)
	}
}
