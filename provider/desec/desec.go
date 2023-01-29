/*
Copyright 2022 The Kubernetes Authors.

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

package desec

import (
	"context"
	"fmt"
	"strings"

	"github.com/nrdcg/desec"
	nc "github.com/nrdcg/desec"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// DesecProvider is an implementation of Provider for Desec DNS.
type DesecProvider struct {
	provider.BaseProvider
	client       *nc.Client
	domainFilter endpoint.DomainFilter
	dryRun       bool
}

// DesecChange includes the changesets that need to be applied to the Desec CCP API
type DesecChange struct {
	Create    *[]nc.RRSet
	UpdateNew *[]nc.RRSet
	UpdateOld *[]nc.RRSet
	Delete    *[]nc.RRSet
}

// NewDesecProvider creates a new provider including the Desec CCP API client
func NewDesecProvider(domainFilter endpoint.DomainFilter, apiKey string, dryRun bool) (*DesecProvider, error) {
	if !domainFilter.IsConfigured() {
		return nil, fmt.Errorf("Desec provider requires at least one configured domain in the domainFilter")
	}

	if apiKey == "" {
		return nil, fmt.Errorf("Desec provider requires an API Key")
	}

	client := desec.New(apiKey, desec.NewDefaultClientOptions())

	return &DesecProvider{
		client:       client,
		domainFilter: domainFilter,
		dryRun:       dryRun,
	}, nil
}

// Records delivers the list of Endpoint records for all zones.
func (p *DesecProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints := make([]*endpoint.Endpoint, 0)

	if p.dryRun {
		log.Debugf("dry run - skipping login")
	} else {

		for _, domain := range p.domainFilter.Filters {
			records, err := p.client.Records.GetAll(ctx, domain, nil)
			if err != nil {
				return nil, fmt.Errorf("unable to query DNS Records for domain '%v': %v", domain, err)
			}

			for _, record := range records {
				ep := endpoint.NewEndpointWithTTL(record.Name, record.Type, endpoint.TTL(record.TTL), record.Records...)
				endpoints = append(endpoints, ep)
			}
		}
	}
	log.Debugf("Endpoints collected: %v", endpoints)
	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *DesecProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if !changes.HasChanges() {
		log.Debugf("no changes detected - nothing to do")
		return nil
	}

	perZoneChanges := map[string]*plan.Changes{}

	for _, zoneName := range p.domainFilter.Filters {
		log.Debugf("zone detected - %s", zoneName)

		perZoneChanges[zoneName] = &plan.Changes{}
	}

	for _, ep := range changes.Create {
		zoneName := endpointZoneName(ep, p.domainFilter.Filters)
		if zoneName == "" {
			log.Debugf("create - ignoring change since %s did not match any zone", ep)
			continue
		}
		log.Debugf("planning Create %v in %s", ep, zoneName)

		perZoneChanges[zoneName].Create = append(perZoneChanges[zoneName].Create, ep)
	}

	for _, ep := range changes.UpdateOld {
		zoneName := endpointZoneName(ep, p.domainFilter.Filters)
		if zoneName == "" {
			log.Debugf("updateOld - ignoring change since %v did not match any zone", ep)
			continue
		}
		log.Debugf("planning UpdateOld %v in %s", ep, zoneName)

		perZoneChanges[zoneName].UpdateOld = append(perZoneChanges[zoneName].UpdateOld, ep)
	}

	for _, ep := range changes.UpdateNew {
		zoneName := endpointZoneName(ep, p.domainFilter.Filters)
		if zoneName == "" {
			log.Debugf("updateNew - ignoring change since %v did not match any zone", ep)
			continue
		}
		log.Debugf("planning UpdateNew %v in %s", ep, zoneName)
		perZoneChanges[zoneName].UpdateNew = append(perZoneChanges[zoneName].UpdateNew, ep)
	}

	for _, ep := range changes.Delete {
		zoneName := endpointZoneName(ep, p.domainFilter.Filters)
		if zoneName == "" {
			log.Debugf("ignoring change since %v did not match any zone", ep)
			continue
		}
		log.Debugf("planning Delete %v in %s", ep, zoneName)
		perZoneChanges[zoneName].Delete = append(perZoneChanges[zoneName].Delete, ep)
	}

	if p.dryRun {
		log.Infof("dry run - not applying changes")
		return nil
	}

	// Assemble changes per zone and prepare it for the Desec API client
	for zoneName, c := range perZoneChanges {

		change := &DesecChange{
			Create:    convertToDesecRecord(c.Create, zoneName),
			UpdateNew: convertToDesecRecord(c.UpdateNew, zoneName),
			UpdateOld: convertToDesecRecord(c.UpdateOld, zoneName),
			Delete:    convertToDesecRecord(c.Delete, zoneName),
		}

		for _, cr := range *change.UpdateOld {
			_, err := p.client.Records.Update(ctx, cr.Domain, cr.SubName, cr.Type, cr)
			if err != nil {
				return err
			}
		}

		for _, cr := range *change.Delete {
			err := p.client.Records.Delete(ctx, cr.Domain, cr.SubName, cr.Type)
			if err != nil {
				return err
			}
		}

		for _, cr := range *change.Create {
			_, err := p.client.Records.Create(ctx, cr)
			if err != nil {
				return err
			}
		}

		for _, cr := range *change.UpdateNew {
			_, err := p.client.Records.Update(ctx, cr.Domain, cr.SubName, cr.Type, cr)
			if err != nil {
				return err
			}
		}

	}

	log.Debugf("update completed")

	return nil
}

func addTailingDotToRecordName(recordName string) string {
	if recordName == "" {
		return recordName
	}

	if !strings.HasSuffix(recordName, ".") {
		recordName = recordName + "."
	}

	return recordName
}

// convertToDesecRecord transforms a list of endpoints into a list of Desec DNS Records
// returns a pointer to a list of DNS Records
func convertToDesecRecord(endpoints []*endpoint.Endpoint, zoneName string) *[]nc.RRSet {

	records := make([]nc.RRSet, len(endpoints))

	for i, ep := range endpoints {
		ttl := 3600
		if ep.RecordTTL.IsConfigured() {
			ttl = int(ep.RecordTTL)
		}

		records[i] = nc.RRSet{
			Name:    addTailingDotToRecordName(ep.DNSName),
			Domain:  zoneName,
			SubName: getSubnameDependingOnZonename(ep, zoneName),
			Type:    ep.RecordType,
			Records: ep.Targets,
			TTL:     ttl,
		}

		fmt.Println(records[i])
	}
	return &records
}

func getSubnameDependingOnZonename(ep *endpoint.Endpoint, zoneName string) (subName string) {
	if ep.DNSName != zoneName {
		return extractSubname(ep.DNSName)
	}

	return ""
}

// extract the subname of fqdn. Example: foo.bar.org will return foo
func extractSubname(DNSName string) (subname string) {
	splitSubname := strings.Split(DNSName, ".")
	sn := []string{}

	for i, part := range splitSubname {
		if i == len(splitSubname)-1 {
			continue
		}
		if i == len(splitSubname)-2 {
			continue
		}
		sn = append(sn, part)
	}

	return strings.Join(sn, ".")
}

// endpointZoneName determines zoneName for endpoint by taking longest suffix zoneName match in endpoint DNSName
// returns empty string if no match found
func endpointZoneName(endpoint *endpoint.Endpoint, zones []string) (zone string) {
	var matchZoneName string = ""
	for _, zoneName := range zones {
		if strings.HasSuffix(endpoint.DNSName, zoneName) && len(zoneName) > len(matchZoneName) {
			matchZoneName = zoneName
		}
	}
	return matchZoneName
}
