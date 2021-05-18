/*
Copyright 2020 The Kubernetes Authors.

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

package genericapi

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var (
	// ErrRecordToMutateNotFound when ApplyChange has to update/delete and didn't found the record in the existing zone (Change with no record ID)
	ErrRecordToMutateNotFound = errors.New("record to mutate not found in current zone")
	// ErrNoDryRun No dry run support for the moment
	ErrNoDryRun = errors.New("dry run not supported")
)

type httpClient interface {
	Patch(string, interface{}, interface{}) error
	Post(string, interface{}, interface{}) error
	Put(string, interface{}, interface{}) error
	Get(string, interface{}) error
	Delete(string, interface{}) error
}

type GDProvider struct {
	provider.BaseProvider

	domainFilter endpoint.DomainFilter
	client       httpClient
	ttl          int64
	DryRun       bool
}

// NewGoDaddyProvider initializes a new GoDaddy DNS based Provider.
func NewGoDaddyProvider(ctx context.Context, domainFilter endpoint.DomainFilter, ttl int64, apiKey, apiSecret string, useOTE, dryRun bool) (*GDProvider, error) {
	client, err := NewClient(useOTE, apiKey, apiSecret)

	if err != nil {
		return nil, err
	}

	// TODO: Add Dry Run support
	if dryRun {
		return nil, ErrNoDryRun
	}

	return &GDProvider{
		client:       client,
		domainFilter: domainFilter,
		ttl:          600,
		DryRun:       dryRun,
	}, nil
}

// Records returns the list of records in all relevant zones.
func (p *GDProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	// TODO: Move to config
	apiEndpoint := "/v1/regions/test-region/domains"

	//var endpoints []*endpoint.Endpoint
	endpoints := []*endpoint.Endpoint{}

	if err := p.client.Get(apiEndpoint, &endpoints); err != nil {
		return nil, err
	}

	log.Infof("GenericHTTP: %d endpoints have been found", len(endpoints))

	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *GDProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {


	// TODO: Move to config
	apiEndpoint := "/v1/regions/test-region/domains"

	var apiResponse string

	if err := p.client.Post(apiEndpoint, changes, apiResponse); err != nil {
		log.Error("There was a problem contacting to the API endpoint while trying to update the DNS records")
	}

	return nil
}

func maxOf(vars ...int64) int64 {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
