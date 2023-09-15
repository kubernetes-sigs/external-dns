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

//nolint:staticcheck // Required due to the current dependency on a deprecated version of azure-sdk-for-go
package azure

import (
	"context"
	"fmt"
	"strings"

	azcoreruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	privatedns "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// PrivateZonesClient is an interface of privatedns.PrivateZoneClient that can be stubbed for testing.
type PrivateZonesClient interface {
	NewListByResourceGroupPager(resourceGroupName string, options *privatedns.PrivateZonesClientListByResourceGroupOptions) *azcoreruntime.Pager[privatedns.PrivateZonesClientListByResourceGroupResponse]
}

// PrivateRecordSetsClient is an interface of privatedns.RecordSetsClient that can be stubbed for testing.
type PrivateRecordSetsClient interface {
	NewListPager(resourceGroupName string, privateZoneName string, options *privatedns.RecordSetsClientListOptions) *azcoreruntime.Pager[privatedns.RecordSetsClientListResponse]
	Delete(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, options *privatedns.RecordSetsClientDeleteOptions) (privatedns.RecordSetsClientDeleteResponse, error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, recordType privatedns.RecordType, relativeRecordSetName string, parameters privatedns.RecordSet, options *privatedns.RecordSetsClientCreateOrUpdateOptions) (privatedns.RecordSetsClientCreateOrUpdateResponse, error)
}

// AzurePrivateDNSProvider implements the DNS provider for Microsoft's Azure Private DNS service
type AzurePrivateDNSProvider struct {
	provider.BaseProvider
	domainFilter                 endpoint.DomainFilter
	zoneIDFilter                 provider.ZoneIDFilter
	dryRun                       bool
	resourceGroup                string
	userAssignedIdentityClientID string
	zonesClient                  PrivateZonesClient
	recordSetsClient             PrivateRecordSetsClient
}

// NewAzurePrivateDNSProvider creates a new Azure Private DNS provider.
//
// Returns the provider or an error if a provider could not be created.
func NewAzurePrivateDNSProvider(configFile string, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, resourceGroup, userAssignedIdentityClientID string, dryRun bool) (*AzurePrivateDNSProvider, error) {
	cfg, err := getConfig(configFile, resourceGroup, userAssignedIdentityClientID)
	if err != nil {
		return nil, fmt.Errorf("failed to read Azure config file '%s': %v", configFile, err)
	}
	cred, err := getCredentials(*cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to get credentials: %w", err)
	}
	zonesClient, err := privatedns.NewPrivateZonesClient(cfg.SubscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}
	recordSetsClient, err := privatedns.NewRecordSetsClient(cfg.SubscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}
	return &AzurePrivateDNSProvider{
		domainFilter:                 domainFilter,
		zoneIDFilter:                 zoneIDFilter,
		dryRun:                       dryRun,
		resourceGroup:                cfg.ResourceGroup,
		userAssignedIdentityClientID: cfg.UserAssignedIdentityID,
		zonesClient:                  zonesClient,
		recordSetsClient:             recordSetsClient,
	}, nil
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
		pager := p.recordSetsClient.NewListPager(p.resourceGroup, *zone.Name, &privatedns.RecordSetsClientListOptions{Top: nil})
		for pager.More() {
			nextResult, err := pager.NextPage(ctx)
			if err != nil {
				return nil, err
			}

			for _, recordSet := range nextResult.Value {
				var recordType string
				if recordSet.Type == nil {
					log.Debugf("Skipping invalid record set with missing type.")
					continue
				}
				recordType = strings.TrimPrefix(*recordSet.Type, "Microsoft.Network/privateDnsZones/")

				var name string
				if recordSet.Name == nil {
					log.Debugf("Skipping invalid record set with missing name.")
					continue
				}
				name = formatAzureDNSName(*recordSet.Name, *zone.Name)

				targets := extractAzurePrivateDNSTargets(recordSet)
				if len(targets) == 0 {
					log.Debugf("Failed to extract targets for '%s' with type '%s'.", name, recordType)
					continue
				}

				var ttl endpoint.TTL
				if recordSet.Properties.TTL != nil {
					ttl = endpoint.TTL(*recordSet.Properties.TTL)
				}

				ep := endpoint.NewEndpointWithTTL(name, recordType, ttl, targets...)
				log.Debugf(
					"Found %s record for '%s' with target '%s'.",
					ep.RecordType,
					ep.DNSName,
					ep.Targets,
				)
				endpoints = append(endpoints, ep)
			}
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

	pager := p.zonesClient.NewListByResourceGroupPager(p.resourceGroup, &privatedns.PrivateZonesClientListByResourceGroupOptions{Top: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, zone := range nextResult.Value {
			log.Debugf("Validating Zone: %v", *zone.Name)

			if zone.Name != nil && p.domainFilter.Match(*zone.Name) && p.zoneIDFilter.Match(*zone.ID) {
				zones = append(zones, *zone)
			}
		}
	}

	log.Debugf("Found %d Azure Private DNS zone(s).", len(zones))
	return zones, nil
}

type azurePrivateDNSChangeMap map[string][]*endpoint.Endpoint

func (p *AzurePrivateDNSProvider) mapChanges(zones []privatedns.PrivateZone, changes *plan.Changes) (azurePrivateDNSChangeMap, azurePrivateDNSChangeMap) {
	ignored := map[string]bool{}
	deleted := azurePrivateDNSChangeMap{}
	updated := azurePrivateDNSChangeMap{}
	zoneNameIDMapper := provider.ZoneIDName{}
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
		for _, ep := range endpoints {
			name := p.recordSetNameForZone(zone, ep)
			if p.dryRun {
				log.Infof("Would delete %s record named '%s' for Azure Private DNS zone '%s'.", ep.RecordType, name, zone)
			} else {
				log.Infof("Deleting %s record named '%s' for Azure Private DNS zone '%s'.", ep.RecordType, name, zone)
				if _, err := p.recordSetsClient.Delete(ctx, p.resourceGroup, zone, privatedns.RecordType(ep.RecordType), name, nil); err != nil {
					log.Errorf(
						"Failed to delete %s record named '%s' for Azure Private DNS zone '%s': %v",
						ep.RecordType,
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
		for _, ep := range endpoints {
			name := p.recordSetNameForZone(zone, ep)
			if p.dryRun {
				log.Infof(
					"Would update %s record named '%s' to '%s' for Azure Private DNS zone '%s'.",
					ep.RecordType,
					name,
					ep.Targets,
					zone,
				)
				continue
			}

			log.Infof(
				"Updating %s record named '%s' to '%s' for Azure Private DNS zone '%s'.",
				ep.RecordType,
				name,
				ep.Targets,
				zone,
			)

			recordSet, err := p.newRecordSet(ep)
			if err == nil {
				_, err = p.recordSetsClient.CreateOrUpdate(
					ctx,
					p.resourceGroup,
					zone,
					privatedns.RecordType(ep.RecordType),
					name,
					recordSet,
					nil,
				)
			}
			if err != nil {
				log.Errorf(
					"Failed to update %s record named '%s' to '%s' for Azure Private DNS zone '%s': %v",
					ep.RecordType,
					name,
					ep.Targets,
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
	case privatedns.RecordTypeA:
		aRecords := make([]*privatedns.ARecord, len(endpoint.Targets))
		for i, target := range endpoint.Targets {
			aRecords[i] = &privatedns.ARecord{
				IPv4Address: to.Ptr(target),
			}
		}
		return privatedns.RecordSet{
			Properties: &privatedns.RecordSetProperties{
				TTL:      to.Ptr(ttl),
				ARecords: aRecords,
			},
		}, nil
	case privatedns.RecordTypeCNAME:
		return privatedns.RecordSet{
			Properties: &privatedns.RecordSetProperties{
				TTL: to.Ptr(ttl),
				CnameRecord: &privatedns.CnameRecord{
					Cname: to.Ptr(endpoint.Targets[0]),
				},
			},
		}, nil
	case privatedns.RecordTypeMX:
		mxRecords := make([]*privatedns.MxRecord, len(endpoint.Targets))
		for i, target := range endpoint.Targets {
			mxRecord, err := parseMxTarget[privatedns.MxRecord](target)
			if err != nil {
				return privatedns.RecordSet{}, err
			}
			mxRecords[i] = &mxRecord
		}
		return privatedns.RecordSet{
			Properties: &privatedns.RecordSetProperties{
				TTL:       to.Ptr(ttl),
				MxRecords: mxRecords,
			},
		}, nil
	case privatedns.RecordTypeTXT:
		return privatedns.RecordSet{
			Properties: &privatedns.RecordSetProperties{
				TTL: to.Ptr(ttl),
				TxtRecords: []*privatedns.TxtRecord{
					{
						Value: []*string{
							&endpoint.Targets[0],
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
	properties := recordSet.Properties
	if properties == nil {
		return []string{}
	}

	// Check for A records
	aRecords := properties.ARecords
	if len(aRecords) > 0 && (aRecords)[0].IPv4Address != nil {
		targets := make([]string, len(aRecords))
		for i, aRecord := range aRecords {
			targets[i] = *aRecord.IPv4Address
		}
		return targets
	}

	// Check for CNAME records
	cnameRecord := properties.CnameRecord
	if cnameRecord != nil && cnameRecord.Cname != nil {
		return []string{*cnameRecord.Cname}
	}

	// Check for MX records
	mxRecords := properties.MxRecords
	if len(mxRecords) > 0 && (mxRecords)[0].Exchange != nil {
		targets := make([]string, len(mxRecords))
		for i, mxRecord := range mxRecords {
			targets[i] = fmt.Sprintf("%d %s", *mxRecord.Preference, *mxRecord.Exchange)
		}
		return targets
	}

	// Check for TXT records
	txtRecords := properties.TxtRecords
	if len(txtRecords) > 0 && (txtRecords)[0].Value != nil {
		values := (txtRecords)[0].Value
		if len(values) > 0 {
			return []string{*(values)[0]}
		}
	}
	return []string{}
}
