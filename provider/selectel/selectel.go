/*
Copyright 2023 The Kubernetes Authors.

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

package selectel

import (
	"context"
	"fmt"
	"os"
	"strconv"

	v1 "github.com/selectel/domains-go/pkg/v1"
	"github.com/selectel/domains-go/pkg/v1/domain"
	"github.com/selectel/domains-go/pkg/v1/record"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	SelectelApiEndpoint = "https://api.selectel.ru/domains/v1"
	SelectelTTL         = 3600
)

// SelectelProvider is an implementation of Provider for Selectel DNS.
type SelectelProvider struct {
	provider.BaseProvider
	client v1.ServiceClient

	domainFilter endpoint.DomainFilter
	DryRun       bool
}

func NewSelectelProvider(_ context.Context, domainFilter endpoint.DomainFilter, dryRun bool) (*SelectelProvider, error) {
	apiKey, apiKeyOk := os.LookupEnv("SELECTEL_API_KEY")
	if !apiKeyOk {
		return nil, fmt.Errorf("no token found")
	}

	apiEndpoint, apiEndpointOk := os.LookupEnv("SELECTEL_API_ENDPOINT")
	if !apiEndpointOk {
		apiEndpoint = SelectelApiEndpoint
	}

	client := v1.NewDomainsClientV1(apiKey, apiEndpoint)

	p := &SelectelProvider{
		client:       *client,
		domainFilter: domainFilter,
		DryRun:       dryRun,
	}

	return p, nil
}

func ttl(ep *endpoint.Endpoint) int {
	if ep.RecordTTL.IsConfigured() {
		return int(ep.RecordTTL)
	} else {
		return SelectelTTL
	}
}

func (p *SelectelProvider) ZoneMapping(ctx context.Context) provider.ZoneIDName {
	zoneNameIDMapper := provider.ZoneIDName{}
	domainsList, err := p.Domains(ctx)
	if err != nil {
		log.Errorf("Error when listing domains: '%s'", err)
		return nil
	}
	for _, d := range domainsList {
		zoneNameIDMapper.Add(strconv.Itoa(d.ID), d.Name)
	}
	return zoneNameIDMapper
}

func zone(zoneMapper provider.ZoneIDName, record string) int {
	z, _ := zoneMapper.FindZone(record)
	if z == "" {
		return -1
	}
	zoneId, _ := strconv.Atoi(z)
	return zoneId
}

func (p *SelectelProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if !changes.HasChanges() {
		// bailing out early if there is nothing to do
		return nil
	}

	zoneNameIDMapper := p.ZoneMapping(ctx)

	for _, e := range changes.Create {
		for _, t := range e.Targets {
			zoneId := zone(zoneNameIDMapper, e.DNSName)

			log.WithField("name", e.DNSName).
				WithField("recordType", e.RecordType).
				WithField("zoneId", zoneId).
				WithField("target", t).
				WithField("ttl", ttl(e)).
				Debug("CreateRecord")

			if p.DryRun {
				continue
			}
			_, _, err := record.Create(ctx, &p.client, zoneId, &record.CreateOpts{
				Name:    e.DNSName,
				Type:    record.Type(e.RecordType),
				Content: t,
				TTL:     ttl(e),
			})

			if err != nil {
				log.Errorf("Error when creating endpoint: '%s'", err)
				return err
			}
		}
	}

	for _, e := range changes.Delete {
		for _, target := range e.Targets {
			zoneId := zone(zoneNameIDMapper, e.DNSName)

			records, _, err := record.ListByDomainID(ctx, &p.client, zoneId)
			if err != nil {
				return err
			}
			for _, rec := range records {
				if rec.Name == e.DNSName && rec.Content == target {
					log.WithField("name", e.DNSName).
						WithField("recordType", e.RecordType).
						WithField("zoneId", zoneId).
						WithField("target", target).
						Debug("DeleteRecord")
					if p.DryRun {
						continue
					}
					_, err := record.Delete(ctx, &p.client, zoneId, rec.ID)
					if err != nil {
						return err
					}
				}
			}
		}

	}

	for i := range changes.UpdateOld {
		old := changes.UpdateOld[i]
		new := changes.UpdateNew[i]

		if len(old.Targets) != len(new.Targets) {
			return fmt.Errorf("updating records with differently sized targets list is not supported")
		}

		zoneId := zone(zoneNameIDMapper, old.DNSName)

		for j, target := range old.Targets {

			rr, _, err := record.ListByDomainID(ctx, &p.client, zoneId)
			if err != nil {
				return err
			}
			for _, q := range rr {
				if q.Name == old.DNSName && q.Content == target {
					log.WithField("old:name", old.DNSName).
						WithField("old:recordType", old.RecordType).
						WithField("old:target", old.Targets[j]).
						WithField("old:ttl", old.RecordTTL).
						WithField("new:name", new.DNSName).
						WithField("new:recordType", new.RecordType).
						WithField("new:target", new.Targets[j]).
						WithField("new:ttl", ttl(new)).
						WithField("zoneId", zoneId).
						Debug("UpdateRecord")

					if p.DryRun {
						continue
					}

					_, _, err := record.Update(ctx, &p.client, zoneId, q.ID, &record.UpdateOpts{
						Name:    new.DNSName,
						Type:    record.Type(new.RecordType),
						Content: new.Targets[j],
						TTL:     ttl(new),
					})

					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (p *SelectelProvider) Domains(ctx context.Context) ([]*domain.View, error) {
	var filteredDomains []*domain.View

	domains, _, err := domain.List(ctx, &p.client)
	if err != nil {
		return nil, err
	}

	for _, d := range domains {
		if p.domainFilter.Match(d.Name) {
			filteredDomains = append(filteredDomains, d)
		} else {
			log.Debugf("Skipping domain %s because it was filtered out by the specified --domain-filter", d.Name)
		}
	}

	return filteredDomains, nil
}

func (p *SelectelProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	domains, err := p.Domains(ctx)
	if err != nil {
		return nil, err
	}

	for _, d := range domains {
		records, _, recordsErr := record.ListByDomainName(
			ctx,
			&p.client,
			d.Name,
		)

		if recordsErr != nil {
			continue
		}

		for _, r := range records {
			if provider.SupportedRecordType(string(r.Type)) {
				var txt string
				if r.Value != "" {
					txt = r.Value
				} else {
					txt = r.Content
				}
				endpoints = append(endpoints, endpoint.NewEndpointWithTTL(
					r.Name,
					string(r.Type),
					endpoint.TTL(r.TTL),
					txt,
				))
			}
		}

	}
	endpoints = mergeEndpointsByNameType(endpoints)
	return endpoints, nil
}

// Merge Endpoints with the same Name and Type into a single endpoint with multiple Targets.
func mergeEndpointsByNameType(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	endpointsByNameType := map[string][]*endpoint.Endpoint{}

	for _, e := range endpoints {
		key := fmt.Sprintf("%s-%s", e.DNSName, e.RecordType)
		endpointsByNameType[key] = append(endpointsByNameType[key], e)
	}

	// If no merge occurred, just return the existing endpoints.
	if len(endpointsByNameType) == len(endpoints) {
		return endpoints
	}

	// Otherwise, construct a new list of endpoints with the endpoints merged.
	var result []*endpoint.Endpoint
	for _, endpoints := range endpointsByNameType {
		dnsName := endpoints[0].DNSName
		recordType := endpoints[0].RecordType

		targets := make([]string, len(endpoints))
		for i, e := range endpoints {
			targets[i] = e.Targets[0]
		}

		e := endpoint.NewEndpoint(dnsName, recordType, targets...)
		result = append(result, e)
	}

	return result
}
