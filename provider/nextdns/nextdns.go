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

func NewNextDNSProvider(cfg NextDNSConfig) (*NextDnsProvider, error) {
	client, _ := api.New(api.WithAPIKey(cfg.NextDNSAPIKey))

	return &NextDnsProvider{
		api:          api.NewRewritesService(client),
		profileId:    cfg.NextDNSProfileId,
		dryRun:       cfg.DryRun,
		domainFilter: cfg.DomainFilter,
	}, nil
}

func (p *NextDnsProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {

	for _, change := range changes.Delete {
		id, exists := change.GetProviderSpecificProperty("id")
		switch change.RecordType {
		case "A":
			fallthrough
		case "CNAME":
			if exists {
				if p.dryRun {
					log.Infof("DELETE[%s]: %s -> %s", id, change.DNSName, change.Targets[0])
				} else {
					p.api.Delete(ctx, &api.DeleteRewritesRequest{
						ProfileID: p.profileId,
						ID:        id.Value,
					})
				}
			}
		}
	}

	for _, change := range changes.UpdateOld {
		id, exists := change.GetProviderSpecificProperty("id")
		switch change.RecordType {
		case "A":
			fallthrough
		case "CNAME":
			if exists {
				if p.dryRun {
					log.Infof("DELETE[%s]: %s -> %s", id, change.DNSName, change.Targets[0])
				} else {
					p.api.Delete(ctx, &api.DeleteRewritesRequest{
						ProfileID: p.profileId,
						ID:        id.Value,
					})
				}
			}
		}
	}

	for _, change := range changes.UpdateNew {
		switch change.RecordType {
		case "A":
			fallthrough
		case "CNAME":
			if p.dryRun {
				log.Infof("CREATE: %s -> %s", change.DNSName, change.Targets[0])
			} else {
				p.api.Create(ctx, &api.CreateRewritesRequest{
					ProfileID: p.profileId,
					Rewrites: &api.Rewrites{
						Name:    change.DNSName,
						Content: change.Targets[0],
					}})
			}
		}
	}

	for _, change := range changes.Create {
		switch change.RecordType {
		case "A":
			fallthrough
		case "CNAME":
			if p.dryRun {
				log.Infof("CREATE: %s -> %s", change.DNSName, change.Targets[0])
			} else {
				p.api.Create(ctx, &api.CreateRewritesRequest{
					ProfileID: p.profileId,
					Rewrites: &api.Rewrites{
						Name:    change.DNSName,
						Content: change.Targets[0],
					}})
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
		log.Infof("Found rewrite [%s]: %s -> %s", rewrite.ID, rewrite.Name, rewrite.Content)
	}

	return out, nil
}
