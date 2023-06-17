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

package nextdns

import (
	"context"
	"errors"

	api "github.com/amalucelli/nextdns-go/nextdns"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type NextDnsProvider struct {
	provider.BaseProvider
	api          api.RewritesService
	profileId    string
	dryRun       bool
	domainFilter endpoint.DomainFilter
}

type NextDNSConfig struct {
	DryRun           bool
	NextDNSAPIKey    string
	NextDNSProfileId string
	DomainFilter     endpoint.DomainFilter
}

type nextDNSEntryKey struct {
	Target     string
	RecordType string
}

func NewNextDNSProvider(cfg NextDNSConfig) (*NextDnsProvider, error) {
	if cfg.NextDNSAPIKey == "" {
		return nil, errors.New("no nextdns api key provided")
	}
	if cfg.NextDNSProfileId == "" {
		return nil, errors.New("no nextdns profile id provided")
	}

	client, _ := api.New(api.WithAPIKey(cfg.NextDNSAPIKey))

	return &NextDnsProvider{
		api:          api.NewRewritesService(client),
		profileId:    cfg.NextDNSProfileId,
		dryRun:       cfg.DryRun,
		domainFilter: cfg.DomainFilter,
	}, nil
}

func (p *NextDnsProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {

	for _, ep := range changes.Delete {
		if err := deleteRecord(ep, p, ctx); err != nil {
			return err
		}
	}

	// Handle updated state - there are no endpoints for updating in place.
	updateNew := make(map[nextDNSEntryKey]*endpoint.Endpoint)
	for _, ep := range changes.UpdateNew {
		key := nextDNSEntryKey{ep.DNSName, ep.RecordType}
		updateNew[key] = ep
	}

	for _, ep := range changes.UpdateOld {
		// Check if this existing entry has an exact match for an updated entry and skip it if so.
		key := nextDNSEntryKey{ep.DNSName, ep.RecordType}
		if newRecord := updateNew[key]; newRecord != nil {
			if newRecord.Targets[0] == ep.Targets[0] {
				delete(updateNew, key)
				continue
			}
		}
		if err := deleteRecord(ep, p, ctx); err != nil {
			return err
		}
	}

	for _, ep := range changes.Create {
		if err := createRecord(ep, p, ctx); err != nil {
			return err
		}
	}

	for _, ep := range updateNew {
		if err := createRecord(ep, p, ctx); err != nil {
			return err
		}
	}

	return nil
}

func deleteRecord(ep *endpoint.Endpoint, p *NextDnsProvider, ctx context.Context) error {
	id, exists := ep.GetProviderSpecificProperty("id")
	switch ep.RecordType {
	case "A":
		fallthrough
	case "CNAME":
		if exists {
			if p.dryRun {
				log.Infof("DELETE[%s]: %s -> %s", id, ep.DNSName, ep.Targets[0])
			} else {
				log.Debugf("DELETE[%s]: %s -> %s", id, ep.DNSName, ep.Targets[0])
				if err := p.api.Delete(ctx, &api.DeleteRewritesRequest{
					ProfileID: p.profileId,
					ID:        id,
				}); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func createRecord(ep *endpoint.Endpoint, p *NextDnsProvider, ctx context.Context) error {
	switch ep.RecordType {
	case "A":
		fallthrough
	case "CNAME":
		if p.dryRun {
			log.Infof("CREATE: %s -> %s", ep.DNSName, ep.Targets[0])
		} else {
			log.Debugf("CREATE: %s -> %s", ep.DNSName, ep.Targets[0])
			if _, err := p.api.Create(ctx, &api.CreateRewritesRequest{
				ProfileID: p.profileId,
				Rewrites: &api.Rewrites{
					Name:    ep.DNSName,
					Content: ep.Targets[0],
				}},
			); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *NextDnsProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	rewrites, err := p.api.List(ctx, &api.ListRewritesRequest{ProfileID: p.profileId})

	if err != nil {
		return nil, err
	}

	out := make([]*endpoint.Endpoint, 0)
	for _, rewrite := range rewrites {
		endpoint := endpoint.NewEndpoint(rewrite.Name, rewrite.Type, rewrite.Content).WithProviderSpecific("id", rewrite.ID)
		out = append(out, endpoint)
		log.Debugf("Found rewrite [%s]: %s -> %s", rewrite.ID, rewrite.Name, rewrite.Content)
	}

	return out, nil
}
