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

package adguardhome

import (
	"context"
	"errors"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

var (
	// ErrNoServer is returned when there is no AdGuard Home server configured in the environment.
	ErrNoServer = errors.New("no adguard home server found in the environment or flags")

	// ErrInvalidResponse is returned when the server sends an unexpected response.
	ErrInvalidResponse = errors.New("invalid response from adguard home")
)

// AdGuardHomeProvider is an implementation of AdGuardHomeProvider for AdGuard Home Local DNS.
type AdGuardHomeProvider struct {
	provider.BaseProvider
	// api AdGuard Home API client
	api api
}

// AdGuardHomeConfig is used for configuring a AdGuardHomeProvider.
type AdGuardHomeConfig struct {
	// Server The root URL of the AdGuard Home server.
	Server string
	// Username An optional username if the server is protected.
	Username string
	// Password An optional password if the server is protected.
	Password string
	// TLSInsecureSkipVerify Disable verification of TLS certificates.
	TLSInsecureSkipVerify bool
	// DomainFilter A filter to apply when looking up and applying records.
	DomainFilter endpoint.DomainFilter
	// DryRun Do nothing and log what would have changed to stdout.
	DryRun bool
}

// rewriteModel Model for AdGuard Home rewrite requests and responses.
type rewriteModel struct {
	// Domain DNS name to rewrite
	Domain string `json:"domain"`
	// Answer IP or hostname to respond with
	Answer string `json:"answer"`
}

// entryKey Helper struct for de-duping DNS entry updates.
type entryKey struct {
	Target     string
	RecordType string
}

// NewAdGuardHomeProvider initializes a new AdGuard Home Local DNS-based AdGuardHomeProvider.
func NewAdGuardHomeProvider(cfg AdGuardHomeConfig) (*AdGuardHomeProvider, error) {
	api, err := newClient(cfg)
	if err != nil {
		return nil, err
	}
	return &AdGuardHomeProvider{api: api}, nil
}

// Records implements AdGuardHomeProvider, populating a slice of endpoints from AdGuard Home local DNS.
func (p *AdGuardHomeProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return p.api.listRecords(ctx)
}

// ApplyChanges implements AdGuardHomeProvider, syncing desired state with the AdGuard Home server Local DNS.
func (p *AdGuardHomeProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// Handle pure deletes.
	for _, ep := range changes.Delete {
		if err := p.api.deleteRecord(ctx, ep); err != nil {
			return err
		}
	}

	// Handle updated state - there are no endpoints for updating in place.
	updateNew := make(map[entryKey]*endpoint.Endpoint)
	for _, ep := range changes.UpdateNew {
		key := entryKey{ep.DNSName, ep.RecordType}
		updateNew[key] = ep
	}

	for _, ep := range changes.UpdateOld {
		// Check if this existing entry has an exact match for an updated entry and skip it if so.
		key := entryKey{ep.DNSName, ep.RecordType}
		if newRecord := updateNew[key]; newRecord != nil {
			// AdGuard Home only has a single target; no need to compare other fields.
			if newRecord.Targets[0] == ep.Targets[0] {
				delete(updateNew, key)
				continue
			}
		}
		if err := p.api.deleteRecord(ctx, ep); err != nil {
			return err
		}
	}

	// Handle pure creates.
	for _, ep := range changes.Create {
		if err := p.api.createRecord(ctx, ep); err != nil {
			return err
		}
	}

	// Create updated entries
	for _, ep := range updateNew {
		if err := p.api.createRecord(ctx, ep); err != nil {
			return err
		}
	}

	return nil
}
