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

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/privatedns/mgmt/privatedns"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

// PrivateZonesClient is an interface of privatedns.PrivateZoneClient that can be stubbed for testing.
type PrivateZonesClient interface {
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result privatedns.PrivateZoneListResultIterator, err error)
}

// PrivateRecordSetsClient is an interface of privatedns.RecordSetsClient that can be stubbed for testing.
type PrivateRecordSetsClient interface {
	ListComplete(ctx context.Context, resourceGroupName string, zoneName string, top *int32, recordSetNameSuffix string) (result privatedns.RecordSetListResultIterator, err error)
	Delete(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, ifMatch string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, parameters privatedns.RecordSet, ifMatch string, ifNoneMatch string) (result privatedns.RecordSet, err error)
}

// AzurePrivateDNSProvider implements the DNS provider for Microsoft's Azure Private DNS service
type AzurePrivateDNSProvider struct {
	domainFilter     endpoint.DomainFilter
	zoneIDFilter     ZoneIDFilter
	dryRun           bool
	subscriptionID   string
	resourceGroup    string
	zonesClient      PrivateZonesClient
	recordSetsClient PrivateRecordSetsClient
}

// NewAzurePrivateDNSProvider creates a new Azure Private DNS provider.
//
// Returns the provider or an error if a provider could not be created.
func NewAzurePrivateDNSProvider(domainFilter endpoint.DomainFilter, zoneIDFilter ZoneIDFilter, resourceGroup string, subscriptionID string, dryRun bool) (*AzurePrivateDNSProvider, error) {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return nil, err
	}

	zonesClient := privatedns.NewPrivateZonesClient(subscriptionID)
	zonesClient.Authorizer = authorizer
	recordSetsClient := privatedns.NewRecordSetsClient(subscriptionID)
	recordSetsClient.Authorizer = authorizer

	provider := &AzurePrivateDNSProvider{
		domainFilter:     domainFilter,
		zoneIDFilter:     zoneIDFilter,
		dryRun:           dryRun,
		subscriptionID:   subscriptionID,
		resourceGroup:    resourceGroup,
		zonesClient:      zonesClient,
		recordSetsClient: recordSetsClient,
	}
	return provider, nil
}

// Records gets the current records.
//
// Returns the current records or an error if the operation failed.
func (p *AzurePrivateDNSProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.zones(ctx)
	if err != nil {
		return nil, err
	}

	log.Debugf("Retrieving Azure Private DNS Records for resource group '%s'", p.resourceGroup)

	for _, zone := range zones {
		err := p.iterateRecords(ctx, *zone.Name, func(recordSet privatedns.RecordSet) {
			var recordType string
			if recordSet.Type == nil {
				log.Debugf("Skipping invalid record set with missing type.")
				return
			}
			recordType = strings.TrimPrefix(*recordSet.Type, "Microsoft.Network/privateDnsZones/")

			var name string
			if recordSet.Name == nil {
				log.Debugf("Skipping invalid record set with missing name.")
				return
			}
			name = formatAzureDNSName(*recordSet.Name, *zone.Name)

			targets := extractAzurePrivateDNSTargets(&recordSet)
			if len(targets) == 0 {
				log.Debugf("Failed to extract targets for '%s' with type '%s'.", name, recordType)
				return
			}

			var ttl endpoint.TTL
			if recordSet.TTL != nil {
				ttl = endpoint.TTL(*recordSet.TTL)
			}

			ep := endpoint.NewEndpointWithTTL(name, recordType, ttl, targets...)
			log.Debugf(
				"Found %s record for '%s' with target '%s'.",
				ep.RecordType,
				ep.DNSName,
				ep.Targets,
			)
			endpoints = append(endpoints, ep)
		})
		if err != nil {
			return nil, err
		}
	}

	log.Debugf("Returning %d Azure Private DNS Records for resource group '%s'", len(endpoints), p.resourceGroup)

	return endpoints, nil
}

// ApplyChanges applies the given changes.
//
// Returns nil if the operation was successful or an error if the operation failed.
func (p *AzurePrivateDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	log.Debugf("Received %d changes to process", len(changes.Create)+len(changes.Delete)+len(changes.UpdateNew)+len(changes.UpdateOld))

	zones, err := p.zones(ctx)
	if err != nil {
		return err
	}

	deleted, updated := p.mapChanges(zones, changes)
	p.deleteRecords(ctx, deleted)
	p.updateRecords(ctx, updated)
	return nil
}

func (p *AzurePrivateDNSProvider) zones(ctx context.Context) ([]privatedns.PrivateZone, error) {
	log.Debugf("Retrieving Azure Private DNS zones for Resource Group '%s'", p.resourceGroup)

	var zones []privatedns.PrivateZone

	i, err := p.zonesClient.ListByResourceGroupComplete(ctx, p.resourceGroup, nil)
	if err != nil {
		return nil, err
	}

	for i.NotDone() {
		zone := i.Value()
		log.Debugf("Validating Zone: %v", *zone.Name)

		if zone.Name != nil && p.domainFilter.Match(*zone.Name) && p.zoneIDFilter.Match(*zone.ID) {
			zones = append(zones, zone)
		}

		err := i.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	log.Debugf("Found %d Azure Private DNS zone(s).", len(zones))
	return zones, nil
}

func (p *AzurePrivateDNSProvider) iterateRecords(ctx context.Context, zoneName string, callback func(privatedns.RecordSet)) error {
	log.Debugf("Retrieving Azure Private DNS Records for zone '%s'.", zoneName)

	i, err := p.recordSetsClient.ListComplete(ctx, p.resourceGroup, zoneName, nil, "")
	if err != nil {
		return err
	}

	for i.NotDone() {
		callback(i.Value())

		err := i.NextWithContext(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

type azurePrivateDNSChangeMap map[string][]*endpoint.Endpoint

func (p *AzurePrivateDNSProvider) mapChanges(zones []privatedns.PrivateZone, changes *plan.Changes) (azurePrivateDNSChangeMap, azurePrivateDNSChangeMap) {
	ignored := map[string]bool{}
	deleted := azurePrivateDNSChangeMap{}
	updated := azurePrivateDNSChangeMap{}
	zoneNameIDMapper := zoneIDName{}
	for _, z := range zones {
		if z.Name != nil {
			zoneNameIDMapper.Add(*z.Name, *z.Name)
		}
	}
	mapChange := func(changeMap azurePrivateDNSChangeMap, change *endpoint.Endpoint) {
		zone, _ := zoneNameIDMapper.FindZone(change.DNSName)
		if zone == "" {
			if _, ok := ignored[change.DNSName]; !ok {
				ignored[change.DNSName] = true
				log.Infof("Ignoring changes to '%s' because a suitable Azure Private DNS zone was not found.", change.DNSName)
			}
			return
		}
		// Ensure the record type is suitable
		changeMap[zone] = append(changeMap[zone], change)
	}

	for _, change := range changes.Delete {
		mapChange(deleted, change)
	}

	for _, change := range changes.UpdateOld {
		mapChange(deleted, change)
	}

	for _, change := range changes.Create {
		mapChange(updated, change)
	}

	for _, change := range changes.UpdateNew {
		mapChange(updated, change)
	}
	return deleted, updated
}

func (p *AzurePrivateDNSProvider) deleteRecords(ctx context.Context, deleted azurePrivateDNSChangeMap) {
	log.Debugf("Records to be deleted: %d", len(deleted))
	// Delete records first
	for zone, endpoints := range deleted {
		for _, endpoint := range endpoints {
			name := p.recordSetNameForZone(zone, endpoint)
			if p.dryRun {
				log.Infof("Would delete %s record named '%s' for Azure Private DNS zone '%s'.", endpoint.RecordType, name, zone)
			} else {
				log.Infof("Deleting %s record named '%s' for Azure Private DNS zone '%s'.", endpoint.RecordType, name, zone)
				if _, err := p.recordSetsClient.Delete(ctx, p.resourceGroup, zone, privatedns.RecordType(endpoint.RecordType), name, ""); err != nil {
					log.Errorf(
						"Failed to delete %s record named '%s' for Azure Private DNS zone '%s': %v",
						endpoint.RecordType,
						name,
						zone,
						err,
					)
				}
			}
		}
	}
}

func (p *AzurePrivateDNSProvider) updateRecords(ctx context.Context, updated azurePrivateDNSChangeMap) {
	log.Debugf("Records to be updated: %d", len(updated))
	for zone, endpoints := range updated {
		for _, endpoint := range endpoints {
			name := p.recordSetNameForZone(zone, endpoint)
			if p.dryRun {
				log.Infof(
					"Would update %s record named '%s' to '%s' for Azure Private DNS zone '%s'.",
					endpoint.RecordType,
					name,
					endpoint.Targets,
					zone,
				)
				continue
			}

			log.Infof(
				"Updating %s record named '%s' to '%s' for Azure Private DNS zone '%s'.",
				endpoint.RecordType,
				name,
				endpoint.Targets,
				zone,
			)

			recordSet, err := p.newRecordSet(endpoint)
			if err == nil {
				_, err = p.recordSetsClient.CreateOrUpdate(
					ctx,
					p.resourceGroup,
					zone,
					privatedns.RecordType(endpoint.RecordType),
					name,
					recordSet,
					"",
					"",
				)
			}
			if err != nil {
				log.Errorf(
					"Failed to update %s record named '%s' to '%s' for Azure Private DNS zone '%s': %v",
					endpoint.RecordType,
					name,
					endpoint.Targets,
					zone,
					err,
				)
			}
		}
	}
}

func (p *AzurePrivateDNSProvider) recordSetNameForZone(zone string, endpoint *endpoint.Endpoint) string {
	// Remove the zone from the record set
	name := endpoint.DNSName
	name = name[:len(name)-len(zone)]
	name = strings.TrimSuffix(name, ".")

	// For root, use @
	if name == "" {
		return "@"
	}
	return name
}

func (p *AzurePrivateDNSProvider) newRecordSet(endpoint *endpoint.Endpoint) (privatedns.RecordSet, error) {
	var ttl int64 = azureRecordTTL
	if endpoint.RecordTTL.IsConfigured() {
		ttl = int64(endpoint.RecordTTL)
	}
	switch privatedns.RecordType(endpoint.RecordType) {
	case privatedns.A:
		aRecords := make([]privatedns.ARecord, len(endpoint.Targets))
		for i, target := range endpoint.Targets {
			aRecords[i] = privatedns.ARecord{
				Ipv4Address: to.StringPtr(target),
			}
		}
		return privatedns.RecordSet{
			RecordSetProperties: &privatedns.RecordSetProperties{
				TTL:      to.Int64Ptr(ttl),
				ARecords: &aRecords,
			},
		}, nil
	case privatedns.CNAME:
		return privatedns.RecordSet{
			RecordSetProperties: &privatedns.RecordSetProperties{
				TTL: to.Int64Ptr(ttl),
				CnameRecord: &privatedns.CnameRecord{
					Cname: to.StringPtr(endpoint.Targets[0]),
				},
			},
		}, nil
	case privatedns.TXT:
		return privatedns.RecordSet{
			RecordSetProperties: &privatedns.RecordSetProperties{
				TTL: to.Int64Ptr(ttl),
				TxtRecords: &[]privatedns.TxtRecord{
					{
						Value: &[]string{
							endpoint.Targets[0],
						},
					},
				},
			},
		}, nil
	}
	return privatedns.RecordSet{}, fmt.Errorf("unsupported record type '%s'", endpoint.RecordType)
}

// Helper function (shared with test code)
func extractAzurePrivateDNSTargets(recordSet *privatedns.RecordSet) []string {
	properties := recordSet.RecordSetProperties
	if properties == nil {
		return []string{}
	}

	// Check for A records
	aRecords := properties.ARecords
	if aRecords != nil && len(*aRecords) > 0 && (*aRecords)[0].Ipv4Address != nil {
		targets := make([]string, len(*aRecords))
		for i, aRecord := range *aRecords {
			targets[i] = *aRecord.Ipv4Address
		}
		return targets
	}

	// Check for CNAME records
	cnameRecord := properties.CnameRecord
	if cnameRecord != nil && cnameRecord.Cname != nil {
		return []string{*cnameRecord.Cname}
	}

	// Check for TXT records
	txtRecords := properties.TxtRecords
	if txtRecords != nil && len(*txtRecords) > 0 && (*txtRecords)[0].Value != nil {
		values := (*txtRecords)[0].Value
		if values != nil && len(*values) > 0 {
			return []string{(*values)[0]}
		}
	}
	return []string{}
}
