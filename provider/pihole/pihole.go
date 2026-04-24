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
	"slices"

	"github.com/google/go-cmp/cmp"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// ErrNoPiholeServer is returned when there is no Pihole server configured
// in the environment.
var ErrNoPiholeServer = errors.New("no pihole server found in the environment or flags")

const (
	warningMsg = "Pi-hole v5 API support is deprecated. Set --pihole-api-version=\"6\" to use the Pi-hole v6 API. The v5 API will be removed in a future release."
)

// PiholeProvider is an implementation of Provider for Pi-hole Local DNS.
type PiholeProvider struct {
	provider.BaseProvider
	api        piholeAPI
	apiVersion string
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
	DomainFilter *endpoint.DomainFilter
	// Do nothing and log what would have changed to stdout.
	DryRun bool
	// PiHole API version =<5 or >=6, default is 5
	APIVersion string
}

// Helper struct for de-duping DNS entry updates.
type piholeEntryKey struct {
	Target     string
	RecordType string
}

// New creates a Pi-hole provider from the given configuration.
func New(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	return newProvider(
		PiholeConfig{
			Server:                cfg.PiholeServer,
			Password:              cfg.PiholePassword,
			TLSInsecureSkipVerify: cfg.PiholeTLSInsecureSkipVerify,
			DomainFilter:          domainFilter,
			DryRun:                cfg.DryRun,
			APIVersion:            cfg.PiholeApiVersion,
		},
	)
}

// newProvider initializes a new Pi-hole Local DNS based Provider.
func newProvider(cfg PiholeConfig) (*PiholeProvider, error) {
	var api piholeAPI
	var err error
	switch cfg.APIVersion {
	case "6":
		api, err = newPiholeClientV6(cfg)
	default:
		log.Warn(warningMsg)
		api, err = newPiholeClient(cfg)
	}
	if err != nil {
		return nil, err
	}
	return &PiholeProvider{api: api, apiVersion: cfg.APIVersion}, nil
}

// Records implements Provider, populating a slice of endpoints from
// Pi-Hole local DNS.
func (p *PiholeProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	aRecords, err := p.api.listRecords(ctx, endpoint.RecordTypeA)
	if err != nil {
		return nil, err
	}
	aaaaRecords, err := p.api.listRecords(ctx, endpoint.RecordTypeAAAA)
	if err != nil {
		return nil, err
	}
	cnameRecords, err := p.api.listRecords(ctx, endpoint.RecordTypeCNAME)
	if err != nil {
		return nil, err
	}
	aRecords = append(aRecords, aaaaRecords...)
	return append(aRecords, cnameRecords...), nil
}

// ApplyChanges implements Provider, syncing desired state with the Pi-hole server Local DNS.
func (p *PiholeProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if err := p.applyDeletes(ctx, changes.Delete); err != nil {
		return err
	}

	updateNew := p.buildUpdateMap(changes.UpdateNew)

	if err := p.applyUpdateOld(ctx, changes.UpdateOld, updateNew); err != nil {
		return err
	}

	return p.applyCreates(ctx, changes.Create, updateNew)
}

// applyDeletes performs pure deletes from the plan.
func (p *PiholeProvider) applyDeletes(ctx context.Context, deletes []*endpoint.Endpoint) error {
	for _, ep := range deletes {
		if err := p.api.deleteRecord(ctx, ep); err != nil {
			return err
		}
	}
	return nil
}

// buildUpdateMap collapses UpdateNew endpoints by (DNSName, RecordType). For
// the v6 API, endpoints that share a key have their Targets merged and
// deduplicated so a single createRecord call carries all desired targets.
func (p *PiholeProvider) buildUpdateMap(updateNew []*endpoint.Endpoint) map[piholeEntryKey]*endpoint.Endpoint {
	m := make(map[piholeEntryKey]*endpoint.Endpoint, len(updateNew))
	for _, ep := range updateNew {
		key := piholeEntryKey{ep.DNSName, ep.RecordType}
		if p.apiVersion == "6" {
			if existing, ok := m[key]; ok {
				existing.Targets = append(existing.Targets, ep.Targets...)
				slices.Sort(existing.Targets)
				existing.Targets = slices.Compact(existing.Targets)
				ep = existing
			}
		}
		m[key] = ep
	}
	return m
}

// applyUpdateOld walks the old side of in-place updates. For each old entry
// whose paired new entry is unchanged, the update is dropped from updateNew
// (nothing to do). Otherwise the old record is deleted so the new record can
// be created in the subsequent phase.
func (p *PiholeProvider) applyUpdateOld(ctx context.Context, updateOld []*endpoint.Endpoint, updateNew map[piholeEntryKey]*endpoint.Endpoint) error {
	for _, ep := range updateOld {
		key := piholeEntryKey{ep.DNSName, ep.RecordType}
		newRecord, ok := updateNew[key]
		if !ok {
			continue
		}
		if p.updateIsNoOp(ep, newRecord) {
			delete(updateNew, key)
			continue
		}
		if err := p.api.deleteRecord(ctx, ep); err != nil {
			return err
		}
	}
	return nil
}

// updateIsNoOp reports whether an update is actually a change. For v6 all
// targets must match; for older APIs only the first target is compared, which
// matches the historical single-target semantics of the v5 local-DNS endpoint.
func (p *PiholeProvider) updateIsNoOp(oldEP, newEP *endpoint.Endpoint) bool {
	if p.apiVersion == "6" {
		return cmp.Diff(oldEP.Targets, newEP.Targets) == ""
	}
	return newEP.Targets[0] == oldEP.Targets[0]
}

// applyCreates runs pure creates followed by the (possibly pruned) updates.
func (p *PiholeProvider) applyCreates(ctx context.Context, creates []*endpoint.Endpoint, updateNew map[piholeEntryKey]*endpoint.Endpoint) error {
	for _, ep := range creates {
		if err := p.api.createRecord(ctx, ep); err != nil {
			return err
		}
	}
	for _, ep := range updateNew {
		if err := p.api.createRecord(ctx, ep); err != nil {
			return err
		}
	}
	return nil
}
