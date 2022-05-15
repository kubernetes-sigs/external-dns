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

package stackpath

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"

	"github.com/wmarchesi123/stackpath-go/pkg/dns"
	"github.com/wmarchesi123/stackpath-go/pkg/oauth2"
)

type StackPathProvider struct {
	provider.BaseProvider
	client       *dns.APIClient
	context      context.Context
	domainFilter endpoint.DomainFilter
	zoneIdFilter provider.ZoneIDFilter
	stackId      string
	dryRun       bool
}

type StackPathConfig struct {
	Context      context.Context
	DomainFilter endpoint.DomainFilter
	ZoneIDFilter provider.ZoneIDFilter
	DryRun       bool
}

func NewStackPathProvider(config StackPathConfig) (*StackPathProvider, error) {
	clientId, ok := os.LookupEnv("STACKPATH_CLIENT_ID")
	if !ok {
		return nil, fmt.Errorf("STACKPATH_CLIENT_ID environment variable is not set")
	}

	clientSecret, ok := os.LookupEnv("STACKPATH_CLIENT_SECRET")
	if !ok {
		return nil, fmt.Errorf("STACKPATH_CLIENT_SECRET environment variable is not set")
	}

	stackId, ok := os.LookupEnv("STACKPATH_STACK_ID")
	if !ok {
		return nil, fmt.Errorf("STACKPATH_STACK_ID environment variable is not set")
	}

	oauthSource := oauth2.NewTokenSource(clientId, clientSecret, oauth2.HTTPClientOption(http.DefaultClient))
	_, err := oauthSource.Token()
	if err != nil {
		return nil, fmt.Errorf("STACKPATH_CLIENT_ID or STACKPATH_CLIENT_SECRET environment variable(s) are invalid")
	}

	authorizedContext := context.WithValue(config.Context, dns.ContextOAuth2, oauthSource)

	clientConfig := dns.NewConfiguration()

	client := dns.NewAPIClient(clientConfig)

	provider := &StackPathProvider{
		client:       client,
		context:      authorizedContext,
		domainFilter: config.DomainFilter,
		zoneIdFilter: config.ZoneIDFilter,
		stackId:      stackId,
		dryRun:       config.DryRun,
	}

	return provider, nil
}

//Base Provider Functions

func (p *StackPathProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	zones, err := p.zones()
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {

		recordsResponse, _, err := p.client.ResourceRecordsApi.GetZoneRecords(p.context, p.stackId, zone.GetId()).Execute()
		if err != nil {
			return nil, err
		}

		records := recordsResponse.GetRecords()

		for _, record := range records {
			if provider.SupportedRecordType(record.GetType()) {
				endpoints = append(endpoints, endpoint.NewEndpointWithTTL(
					record.GetName()+"."+zone.GetDomain(),
					record.GetType(),
					endpoint.TTL(record.GetTtl()),
					record.GetData(),
				),
				)
			}
		}
	}

	// Log endpoints
	log.WithFields(log.Fields{
		"endpoints": endpoints,
	}).Debug("Endpoints fetched from StackPath API")

	return mergeEndpointsByNameType(endpoints), nil
}

func (p *StackPathProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {

	zs, err := p.zones()
	if err != nil {
		return err
	}
	zones := &zs

	zoneIdNameMap := provider.ZoneIDName{}
	for _, zone := range zs {
		zoneIdNameMap.Add(zone.GetId(), zone.GetDomain())
	}

	err = p.create(changes.Create, zones, &zoneIdNameMap)
	if err != nil {
		return err
	}

	err = p.delete(changes.Delete, zones, &zoneIdNameMap)
	if err != nil {
		return err
	}

	err = p.update(changes.UpdateOld, changes.UpdateNew, zones, &zoneIdNameMap)
	if err != nil {
		return err
	}

	return nil
}

func (p *StackPathProvider) create(endpoints []*endpoint.Endpoint, zones *[]dns.ZoneZone, zoneIdNameMap *provider.ZoneIDName) error {
	createsByZoneId := endpointsByZoneId(*zoneIdNameMap, endpoints)

	for zoneId, endpoints := range createsByZoneId {
		domain := (*zoneIdNameMap)[zoneId]
		for _, endpoint := range endpoints {
			for _, target := range endpoint.Targets {

				p.createTarget(zoneId, domain, endpoint, target)

			}
		}
	}

	return nil
}

func (p *StackPathProvider) createTarget(zoneId string, domain string, endpoint *endpoint.Endpoint, target string) error {

	msg := dns.NewZoneUpdateZoneRecordMessage()
	name := strings.TrimSuffix(endpoint.DNSName, domain)
	if name == "" {
		name = "@"
	}
	msg.SetName(endpoint.DNSName)
	msg.SetType(dns.ZoneRecordType(endpoint.RecordType))
	msg.SetTtl(int32(endpoint.RecordTTL))
	msg.SetData(target)

	_, _, err := p.client.ResourceRecordsApi.CreateZoneRecord(p.context, p.stackId, zoneId).ZoneUpdateZoneRecordMessage(*msg).Execute()

	return err
}

func (p *StackPathProvider) delete(endpoints []*endpoint.Endpoint, zones *[]dns.ZoneZone, zoneIdNameMap *provider.ZoneIDName) error {
	return nil
}

func (p *StackPathProvider) update(old []*endpoint.Endpoint, new []*endpoint.Endpoint, zones *[]dns.ZoneZone, zoneIdNameMap *provider.ZoneIDName) error {
	return nil
}

func (p *StackPathProvider) zones() ([]dns.ZoneZone, error) {
	zoneResponse, _, err := p.client.ZonesApi.GetZones(p.context, p.stackId).Execute()
	if err != nil {
		return nil, err
	}

	zones := zoneResponse.GetZones()
	filteredZones := []dns.ZoneZone{}

	for _, zone := range zones {
		if p.zoneIdFilter.Match(zone.GetId()) && p.domainFilter.Match(zone.GetDomain()) {
			filteredZones = append(filteredZones, zone)
			log.Debugf("Matched zone " + zone.GetId())
		} else {
			log.Debugf("Filtered zone " + zone.GetId())
		}
	}

	return filteredZones, nil
}

// Merge Endpoints with the same Name and Type into a single endpoint with
// multiple Targets. From pkg/digitalocean/provider.go
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

//From pkg/digitalocean/provider.go
func endpointsByZoneId(zoneNameIDMapper provider.ZoneIDName, endpoints []*endpoint.Endpoint) map[string][]*endpoint.Endpoint {
	endpointsByZone := make(map[string][]*endpoint.Endpoint)

	for _, ep := range endpoints {
		zoneID, _ := zoneNameIDMapper.FindZone(ep.DNSName)
		if zoneID == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", ep.DNSName)
			continue
		}
		endpointsByZone[zoneID] = append(endpointsByZone[zoneID], ep)
	}

	return endpointsByZone
}
