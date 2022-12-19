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

package pihole

import (
	"context"
	"errors"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// ErrNoPiholeServer is returned when there is no Pihole server configured
// in the environment.
var ErrNoPiholeServer = errors.New("no pihole server found in the environment or flags")

// PiholeProvider is an implementation of Provider for Pi-hole Local DNS.
type PiholeProvider struct {
	provider.BaseProvider
	api piholeAPI
}

// PiholeConfig is used for configuring a PiholeProvider.
type PiholeConfig struct {
	// The root URL of the Pi-hole server.
	Server string
	// An optional password if the server is protected.
	Password string
	// Disable verification of TLS certificates.
	TLSInsecureSkipVerify bool
	// A filter to apply when looking up and applying records.
	DomainFilter endpoint.DomainFilter
	// Do nothing and log what would have changed to stdout.
	DryRun bool
}

// NewPiholeProvider initializes a new Pi-hole Local DNS based Provider.
func NewPiholeProvider(cfg PiholeConfig) (*PiholeProvider, error) {
	api, err := newPiholeClient(cfg)
	if err != nil {
		return nil, err
	}
	return &PiholeProvider{api: api}, nil
}

// Records implements Provider, populating a slice of endpoints from
// Pi-Hole local DNS.
func (p *PiholeProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	aRecords, err := p.api.listRecords(ctx, endpoint.RecordTypeA)
	if err != nil {
		return nil, err
	}
	cnameRecords, err := p.api.listRecords(ctx, endpoint.RecordTypeCNAME)
	if err != nil {
		return nil, err
	}
	return append(aRecords, cnameRecords...), nil
}

// ApplyChanges implements Provider, syncing desired state with the Pi-hole server Local DNS.
func (p *PiholeProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// Handle deletions first - there are no endpoints for updating in place.
	for _, ep := range changes.Delete {
		if err := p.api.deleteRecord(ctx, ep); err != nil {
			return err
		}
	}
	for _, ep := range changes.UpdateOld {
		if err := p.api.deleteRecord(ctx, ep); err != nil {
			return err
		}
	}

	// Handle desired state
	for _, ep := range changes.Create {
		if err := p.api.createRecord(ctx, ep); err != nil {
			return err
		}
	}
	for _, ep := range changes.UpdateNew {
		if err := p.api.createRecord(ctx, ep); err != nil {
			return err
		}
	}

	return nil
}
