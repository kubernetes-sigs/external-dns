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
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"

	"github.com/Azure/azure-sdk-for-go/arm/dns"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
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
}

// ZonesClient is an interface of dns.ZoneClient that can be stubbed for testing.
type ZonesClient interface {
	ListByResourceGroup(resourceGroupName string, top *int32) (result dns.ZoneListResult, err error)
	ListByResourceGroupNextResults(lastResults dns.ZoneListResult) (result dns.ZoneListResult, err error)
}

// RecordsClient is an interface of dns.RecordClient that can be stubbed for testing.
type RecordsClient interface {
	ListByDNSZone(resourceGroupName string, zoneName string, top *int32) (result dns.RecordSetListResult, err error)
	ListByDNSZoneNextResults(list dns.RecordSetListResult) (result dns.RecordSetListResult, err error)
	Delete(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, ifMatch string) (result autorest.Response, err error)
	CreateOrUpdate(resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, parameters dns.RecordSet, ifMatch string, ifNoneMatch string) (result dns.RecordSet, err error)
}

// AzureProvider implements the DNS provider for Microsoft's Azure cloud platform.
type AzureProvider struct {
	domainFilter  DomainFilter
	zoneIDFilter  ZoneIDFilter
	dryRun        bool
	resourceGroup string
	zonesClient   ZonesClient
	recordsClient RecordsClient
}

// NewAzureProvider creates a new Azure provider.
//
// Returns the provider or an error if a provider could not be created.
func NewAzureProvider(configFile string, domainFilter DomainFilter, zoneIDFilter ZoneIDFilter, resourceGroup string, dryRun bool) (*AzureProvider, error) {
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
	recordsClient := dns.NewRecordSetsClientWithBaseURI(environment.ResourceManagerEndpoint, cfg.SubscriptionID)
	recordsClient.Authorizer = autorest.NewBearerAuthorizer(token)

	provider := &AzureProvider{
		domainFilter:  domainFilter,
		zoneIDFilter:  zoneIDFilter,
		dryRun:        dryRun,
		resourceGroup: cfg.ResourceGroup,
		zonesClient:   zonesClient,
		recordsClient: recordsClient,
	}
	return provider, nil
}

// getAccessToken retrieves Azure API access token.
func getAccessToken(cfg config, environment azure.Environment) (*adal.ServicePrincipalToken, error) {
	// Try to retrive token with MSI.
	if cfg.UseManagedIdentityExtension {
		log.Info("Using managed identity extension to retrieve access token for Azure API.")
		msiEndpoint, err := adal.GetMSIVMEndpoint()
		if err != nil {
			return nil, fmt.Errorf("failed to get the managed service identity endpoint: %v", err)
		}

		token, err := adal.NewServicePrincipalTokenFromMSI(msiEndpoint, environment.ServiceManagementEndpoint)
		if err != nil {
			return nil, fmt.Errorf("failed to create the managed service identity token: %v", err)
		}
		return token, nil
	}

	// Try to retrieve token with service principal credentials.
	if len(cfg.ClientID) > 0 && len(cfg.ClientSecret) > 0 {
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

	return nil, fmt.Errorf("no credentials provided for Azure API")
}

// Records gets the current records.
//
// Returns the current records or an error if the operation failed.
func (p *AzureProvider) Records() (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.zones()
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		err := p.iterateRecords(*zone.Name, func(recordSet dns.RecordSet) bool {
			if recordSet.Name == nil || recordSet.Type == nil {
				log.Error("Skipping invalid record set with nil name or type.")
				return true
			}
			recordType := strings.TrimLeft(*recordSet.Type, "Microsoft.Network/dnszones/")
			if !supportedRecordType(recordType) {
				return true
			}
			name := formatAzureDNSName(*recordSet.Name, *zone.Name)
			target := extractAzureTarget(&recordSet)
			if target == "" {
				log.Errorf("Failed to extract target for '%s' with type '%s'.", name, recordType)
				return true
			}
			var ttl endpoint.TTL
			if recordSet.TTL != nil {
				ttl = endpoint.TTL(*recordSet.TTL)
			}

			ep := endpoint.NewEndpointWithTTL(name, recordType, endpoint.TTL(ttl), target)
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
func (p *AzureProvider) ApplyChanges(changes *plan.Changes) error {
	zones, err := p.zones()
	if err != nil {
		return err
	}

	deleted, updated := p.mapChanges(zones, changes)
	p.deleteRecords(deleted)
	p.updateRecords(updated)
	return nil
}

func (p *AzureProvider) zones() ([]dns.Zone, error) {
	log.Debug("Retrieving Azure DNS zones.")

	var zones []dns.Zone
	list, err := p.zonesClient.ListByResourceGroup(p.resourceGroup, nil)
	if err != nil {
		return nil, err
	}

	for list.Value != nil && len(*list.Value) > 0 {
		for _, zone := range *list.Value {
			if zone.Name == nil {
				continue
			}

			if !p.domainFilter.Match(*zone.Name) {
				continue
			}

			if !p.zoneIDFilter.Match(*zone.ID) {
				continue
			}

			zones = append(zones, zone)
		}

		list, err = p.zonesClient.ListByResourceGroupNextResults(list)
		if err != nil {
			return nil, err
		}
	}
	log.Debugf("Found %d Azure DNS zone(s).", len(zones))
	return zones, nil
}

func (p *AzureProvider) iterateRecords(zoneName string, callback func(dns.RecordSet) bool) error {
	log.Debugf("Retrieving Azure DNS records for zone '%s'.", zoneName)

	list, err := p.recordsClient.ListByDNSZone(p.resourceGroup, zoneName, nil)
	if err != nil {
		return err
	}

	for list.Value != nil && len(*list.Value) > 0 {
		for _, recordSet := range *list.Value {
			if !callback(recordSet) {
				return nil
			}
		}

		list, err = p.recordsClient.ListByDNSZoneNextResults(list)
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

func (p *AzureProvider) deleteRecords(deleted azureChangeMap) {
	// Delete records first
	for zone, endpoints := range deleted {
		for _, endpoint := range endpoints {
			name := p.recordSetNameForZone(zone, endpoint)
			if p.dryRun {
				log.Infof("Would delete %s record named '%s' for Azure DNS zone '%s'.", endpoint.RecordType, name, zone)
			} else {
				log.Infof("Deleting %s record named '%s' for Azure DNS zone '%s'.", endpoint.RecordType, name, zone)
				if _, err := p.recordsClient.Delete(p.resourceGroup, zone, name, dns.RecordType(endpoint.RecordType), ""); err != nil {
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

func (p *AzureProvider) updateRecords(updated azureChangeMap) {
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
				_, err = p.recordsClient.CreateOrUpdate(
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
		return dns.RecordSet{
			RecordSetProperties: &dns.RecordSetProperties{
				TTL: to.Int64Ptr(ttl),
				ARecords: &[]dns.ARecord{
					{
						Ipv4Address: to.StringPtr(endpoint.Targets[0]),
					},
				},
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
func extractAzureTarget(recordSet *dns.RecordSet) string {
	properties := recordSet.RecordSetProperties
	if properties == nil {
		return ""
	}

	// Check for A records
	aRecords := properties.ARecords
	if aRecords != nil && len(*aRecords) > 0 && (*aRecords)[0].Ipv4Address != nil {
		return *(*aRecords)[0].Ipv4Address
	}

	// Check for CNAME records
	cnameRecord := properties.CnameRecord
	if cnameRecord != nil && cnameRecord.Cname != nil {
		return *cnameRecord.Cname
	}

	// Check for TXT records
	txtRecords := properties.TxtRecords
	if txtRecords != nil && len(*txtRecords) > 0 && (*txtRecords)[0].Value != nil {
		values := (*txtRecords)[0].Value
		if values != nil && len(*values) > 0 {
			return (*values)[0]
		}
	}
	return ""
}
