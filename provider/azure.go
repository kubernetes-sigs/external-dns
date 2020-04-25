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
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"

	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const (
	azureRecordTTL = 300
)

type config struct {
	Cloud                       string `json:"cloud" yaml:"cloud"`
	TenantID                    string `json:"tenantId" yaml:"tenantId"`
	SubscriptionID              string `json:"subscriptionId" yaml:"subscriptionId"`
	ResourceGroup               string `json:"resourceGroup" yaml:"resourceGroup"`
	Location                    string `json:"location" yaml:"location"`
	ClientID                    string `json:"aadClientId" yaml:"aadClientId"`
	ClientSecret                string `json:"aadClientSecret" yaml:"aadClientSecret"`
	UseManagedIdentityExtension bool   `json:"useManagedIdentityExtension" yaml:"useManagedIdentityExtension"`
	UserAssignedIdentityID      string `json:"userAssignedIdentityID" yaml:"userAssignedIdentityID"`
}

// ZonesClient is an interface of dns.ZoneClient that can be stubbed for testing.
type ZonesClient interface {
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result dns.ZoneListResultIterator, err error)
}

// RecordSetsClient is an interface of dns.RecordSetsClient that can be stubbed for testing.
type RecordSetsClient interface {
	ListAllByDNSZoneComplete(ctx context.Context, resourceGroupName string, zoneName string, top *int32, recordSetNameSuffix string) (result dns.RecordSetListResultIterator, err error)
	Delete(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, ifMatch string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, parameters dns.RecordSet, ifMatch string, ifNoneMatch string) (result dns.RecordSet, err error)
}

// AzureProvider implements the DNS provider for Microsoft's Azure cloud platform.
type AzureProvider struct {
	domainFilter                 endpoint.DomainFilter
	zoneIDFilter                 ZoneIDFilter
	dryRun                       bool
	resourceGroup                string
	userAssignedIdentityClientID string
	zonesClient                  ZonesClient
	recordSetsClient             RecordSetsClient
}

// NewAzureProvider creates a new Azure provider.
//
// Returns the provider or an error if a provider could not be created.
func NewAzureProvider(configFile string, domainFilter endpoint.DomainFilter, zoneIDFilter ZoneIDFilter, resourceGroup string, userAssignedIdentityClientID string, dryRun bool) (*AzureProvider, error) {
	contents, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read Azure config file '%s': %v", configFile, err)
	}
	cfg := config{}
	err = yaml.Unmarshal(contents, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read Azure config file '%s': %v", configFile, err)
	}

	// If a resource group was given, override what was present in the config file
	if resourceGroup != "" {
		cfg.ResourceGroup = resourceGroup
	}
	// If userAssignedIdentityClientID is provided explicitly, override existing one in config file
	if userAssignedIdentityClientID != "" {
		cfg.UserAssignedIdentityID = userAssignedIdentityClientID
	}

	var environment azure.Environment
	if cfg.Cloud == "" {
		environment = azure.PublicCloud
	} else {
		environment, err = azure.EnvironmentFromName(cfg.Cloud)
		if err != nil {
			return nil, fmt.Errorf("invalid cloud value '%s': %v", cfg.Cloud, err)
		}
	}

	token, err := getAccessToken(cfg, environment)
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %v", err)
	}

	zonesClient := dns.NewZonesClientWithBaseURI(environment.ResourceManagerEndpoint, cfg.SubscriptionID)
	zonesClient.Authorizer = autorest.NewBearerAuthorizer(token)
	recordSetsClient := dns.NewRecordSetsClientWithBaseURI(environment.ResourceManagerEndpoint, cfg.SubscriptionID)
	recordSetsClient.Authorizer = autorest.NewBearerAuthorizer(token)

	provider := &AzureProvider{
		domainFilter:                 domainFilter,
		zoneIDFilter:                 zoneIDFilter,
		dryRun:                       dryRun,
		resourceGroup:                cfg.ResourceGroup,
		userAssignedIdentityClientID: cfg.UserAssignedIdentityID,
		zonesClient:                  zonesClient,
		recordSetsClient:             recordSetsClient,
	}
	return provider, nil
}

// getAccessToken retrieves Azure API access token.
func getAccessToken(cfg config, environment azure.Environment) (*adal.ServicePrincipalToken, error) {
	// Try to retrieve token with service principal credentials.
	// Try to use service principal first, some AKS clusters are in an intermediate state that `UseManagedIdentityExtension` is `true`
	// and service principal exists. In this case, we still want to use service principal to authenticate.
	if len(cfg.ClientID) > 0 &&
		len(cfg.ClientSecret) > 0 &&
		// due to some historical reason, for pure MSI cluster,
		// they will use "msi" as placeholder in azure.json.
		// In this case, we shouldn't try to use SPN to authenticate.
		!strings.EqualFold(cfg.ClientID, "msi") &&
		!strings.EqualFold(cfg.ClientSecret, "msi") {
		log.Info("Using client_id+client_secret to retrieve access token for Azure API.")
		oauthConfig, err := adal.NewOAuthConfig(environment.ActiveDirectoryEndpoint, cfg.TenantID)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve OAuth config: %v", err)
		}

		token, err := adal.NewServicePrincipalToken(*oauthConfig, cfg.ClientID, cfg.ClientSecret, environment.ResourceManagerEndpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to create service principal token: %v", err)
		}
		return token, nil
	}

	// Try to retrieve token with MSI.
	if cfg.UseManagedIdentityExtension {
		log.Info("Using managed identity extension to retrieve access token for Azure API.")
		msiEndpoint, err := adal.GetMSIVMEndpoint()
		if err != nil {
			return nil, fmt.Errorf("failed to get the managed service identity endpoint: %v", err)
		}

		if cfg.UserAssignedIdentityID != "" {
			log.Infof("Resolving to user assigned identity, client id is %s.", cfg.UserAssignedIdentityID)
			token, err := adal.NewServicePrincipalTokenFromMSIWithUserAssignedID(msiEndpoint, environment.ServiceManagementEndpoint, cfg.UserAssignedIdentityID)
			if err != nil {
				return nil, fmt.Errorf("failed to create the managed service identity token: %v", err)
			}
			return token, nil
		}

		log.Info("Resolving to system assigned identity.")
		token, err := adal.NewServicePrincipalTokenFromMSI(msiEndpoint, environment.ServiceManagementEndpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to create the managed service identity token: %v", err)
		}
		return token, nil
	}

	return nil, fmt.Errorf("no credentials provided for Azure API")
}

// Records gets the current records.
//
// Returns the current records or an error if the operation failed.
func (p *AzureProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.zones(ctx)
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		err := p.iterateRecords(ctx, *zone.Name, func(recordSet dns.RecordSet) bool {
			if recordSet.Name == nil || recordSet.Type == nil {
				log.Error("Skipping invalid record set with nil name or type.")
				return true
			}
			recordType := strings.TrimPrefix(*recordSet.Type, "Microsoft.Network/dnszones/")
			if !supportedRecordType(recordType) {
				return true
			}
			name := formatAzureDNSName(*recordSet.Name, *zone.Name)
			targets := extractAzureTargets(&recordSet)
			if len(targets) == 0 {
				log.Errorf("Failed to extract targets for '%s' with type '%s'.", name, recordType)
				return true
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
			return true
		})
		if err != nil {
			return nil, err
		}
	}
	return endpoints, nil
}

// ApplyChanges applies the given changes.
//
// Returns nil if the operation was successful or an error if the operation failed.
func (p *AzureProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	zones, err := p.zones(ctx)
	if err != nil {
		return err
	}

	deleted, updated := p.mapChanges(zones, changes)
	p.deleteRecords(ctx, deleted)
	p.updateRecords(ctx, updated)
	return nil
}

func (p *AzureProvider) zones(ctx context.Context) ([]dns.Zone, error) {
	log.Debugf("Retrieving Azure DNS zones for resource group: %s.", p.resourceGroup)

	var zones []dns.Zone

	zonesIterator, err := p.zonesClient.ListByResourceGroupComplete(ctx, p.resourceGroup, nil)
	if err != nil {
		return nil, err
	}

	for zonesIterator.NotDone() {
		zone := zonesIterator.Value()

		if zone.Name != nil && p.domainFilter.Match(*zone.Name) && p.zoneIDFilter.Match(*zone.ID) {
			zones = append(zones, zone)
		}

		err := zonesIterator.NextWithContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	log.Debugf("Found %d Azure DNS zone(s).", len(zones))
	return zones, nil
}

func (p *AzureProvider) iterateRecords(ctx context.Context, zoneName string, callback func(dns.RecordSet) bool) error {
	log.Debugf("Retrieving Azure DNS records for zone '%s'.", zoneName)

	recordSetsIterator, err := p.recordSetsClient.ListAllByDNSZoneComplete(ctx, p.resourceGroup, zoneName, nil, "")
	if err != nil {
		return err
	}

	for recordSetsIterator.NotDone() {
		if !callback(recordSetsIterator.Value()) {
			return nil
		}

		err := recordSetsIterator.NextWithContext(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

type azureChangeMap map[string][]*endpoint.Endpoint

func (p *AzureProvider) mapChanges(zones []dns.Zone, changes *plan.Changes) (azureChangeMap, azureChangeMap) {
	ignored := map[string]bool{}
	deleted := azureChangeMap{}
	updated := azureChangeMap{}
	zoneNameIDMapper := zoneIDName{}
	for _, z := range zones {
		if z.Name != nil {
			zoneNameIDMapper.Add(*z.Name, *z.Name)
		}
	}
	mapChange := func(changeMap azureChangeMap, change *endpoint.Endpoint) {
		zone, _ := zoneNameIDMapper.FindZone(change.DNSName)
		if zone == "" {
			if _, ok := ignored[change.DNSName]; !ok {
				ignored[change.DNSName] = true
				log.Infof("Ignoring changes to '%s' because a suitable Azure DNS zone was not found.", change.DNSName)
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

func (p *AzureProvider) deleteRecords(ctx context.Context, deleted azureChangeMap) {
	// Delete records first
	for zone, endpoints := range deleted {
		for _, endpoint := range endpoints {
			name := p.recordSetNameForZone(zone, endpoint)
			if p.dryRun {
				log.Infof("Would delete %s record named '%s' for Azure DNS zone '%s'.", endpoint.RecordType, name, zone)
			} else {
				log.Infof("Deleting %s record named '%s' for Azure DNS zone '%s'.", endpoint.RecordType, name, zone)
				if _, err := p.recordSetsClient.Delete(ctx, p.resourceGroup, zone, name, dns.RecordType(endpoint.RecordType), ""); err != nil {
					log.Errorf(
						"Failed to delete %s record named '%s' for Azure DNS zone '%s': %v",
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

func (p *AzureProvider) updateRecords(ctx context.Context, updated azureChangeMap) {
	for zone, endpoints := range updated {
		for _, endpoint := range endpoints {
			name := p.recordSetNameForZone(zone, endpoint)
			if p.dryRun {
				log.Infof(
					"Would update %s record named '%s' to '%s' for Azure DNS zone '%s'.",
					endpoint.RecordType,
					name,
					endpoint.Targets,
					zone,
				)
				continue
			}

			log.Infof(
				"Updating %s record named '%s' to '%s' for Azure DNS zone '%s'.",
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
					name,
					dns.RecordType(endpoint.RecordType),
					recordSet,
					"",
					"",
				)
			}
			if err != nil {
				log.Errorf(
					"Failed to update %s record named '%s' to '%s' for DNS zone '%s': %v",
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

func (p *AzureProvider) recordSetNameForZone(zone string, endpoint *endpoint.Endpoint) string {
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

func (p *AzureProvider) newRecordSet(endpoint *endpoint.Endpoint) (dns.RecordSet, error) {
	var ttl int64 = azureRecordTTL
	if endpoint.RecordTTL.IsConfigured() {
		ttl = int64(endpoint.RecordTTL)
	}
	switch dns.RecordType(endpoint.RecordType) {
	case dns.A:
		aRecords := make([]dns.ARecord, len(endpoint.Targets))
		for i, target := range endpoint.Targets {
			aRecords[i] = dns.ARecord{
				Ipv4Address: to.StringPtr(target),
			}
		}
		return dns.RecordSet{
			RecordSetProperties: &dns.RecordSetProperties{
				TTL:      to.Int64Ptr(ttl),
				ARecords: &aRecords,
			},
		}, nil
	case dns.CNAME:
		return dns.RecordSet{
			RecordSetProperties: &dns.RecordSetProperties{
				TTL: to.Int64Ptr(ttl),
				CnameRecord: &dns.CnameRecord{
					Cname: to.StringPtr(endpoint.Targets[0]),
				},
			},
		}, nil
	case dns.TXT:
		return dns.RecordSet{
			RecordSetProperties: &dns.RecordSetProperties{
				TTL: to.Int64Ptr(ttl),
				TxtRecords: &[]dns.TxtRecord{
					{
						Value: &[]string{
							endpoint.Targets[0],
						},
					},
				},
			},
		}, nil
	}
	return dns.RecordSet{}, fmt.Errorf("unsupported record type '%s'", endpoint.RecordType)
}

// Helper function (shared with test code)
func formatAzureDNSName(recordName, zoneName string) string {
	if recordName == "@" {
		return zoneName
	}
	return fmt.Sprintf("%s.%s", recordName, zoneName)
}

// Helper function (shared with text code)
func extractAzureTargets(recordSet *dns.RecordSet) []string {
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
