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

package ibmcloud

import (
	"context"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/crn"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"
	"github.com/IBM/networking-go-sdk/dnssvcsv1"
	"github.com/IBM/networking-go-sdk/zonesv1"
	"gopkg.in/yaml.v2"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/source"
)

var proxyTypeNotSupported = map[string]bool{
	"LOC": true,
	"MX":  true,
	"NS":  true,
	"SPF": true,
	"TXT": true,
	"SRV": true,
}

var privateTypeSupported = map[string]bool{
	"A":     true,
	"CNAME": true,
	"TXT":   true,
}

const (
	// recordCreate is a ChangeAction enum value
	recordCreate = "CREATE"
	// recordDelete is a ChangeAction enum value
	recordDelete = "DELETE"
	// recordUpdate is a ChangeAction enum value
	recordUpdate = "UPDATE"
	// defaultPublicRecordTTL 1 = automatic
	defaultPublicRecordTTL = 1

	proxyFilter             = "ibmcloud-proxied"
	vpcFilter               = "ibmcloud-vpc"
	zoneStatePendingNetwork = "PENDING_NETWORK_ADD"
	zoneStateActive         = "ACTIVE"
)

// Source shadow the interface source.Source. used primarily for unit testing.
type Source interface {
	Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error)
	AddEventHandler(context.Context, func())
}

// ibmcloudClient is a minimal implementation of DNS API that we actually use, used primarily for unit testing.
type ibmcloudClient interface {
	ListAllDDNSRecordsWithContext(ctx context.Context, listAllDNSRecordsOptions *dnsrecordsv1.ListAllDnsRecordsOptions) (result *dnsrecordsv1.ListDnsrecordsResp, response *core.DetailedResponse, err error)
	CreateDNSRecordWithContext(ctx context.Context, createDNSRecordOptions *dnsrecordsv1.CreateDnsRecordOptions) (result *dnsrecordsv1.DnsrecordResp, response *core.DetailedResponse, err error)
	DeleteDNSRecordWithContext(ctx context.Context, deleteDNSRecordOptions *dnsrecordsv1.DeleteDnsRecordOptions) (result *dnsrecordsv1.DeleteDnsrecordResp, response *core.DetailedResponse, err error)
	UpdateDNSRecordWithContext(ctx context.Context, updateDNSRecordOptions *dnsrecordsv1.UpdateDnsRecordOptions) (result *dnsrecordsv1.DnsrecordResp, response *core.DetailedResponse, err error)
	ListDnszonesWithContext(ctx context.Context, listDnszonesOptions *dnssvcsv1.ListDnszonesOptions) (result *dnssvcsv1.ListDnszones, response *core.DetailedResponse, err error)
	GetDnszoneWithContext(ctx context.Context, getDnszoneOptions *dnssvcsv1.GetDnszoneOptions) (result *dnssvcsv1.Dnszone, response *core.DetailedResponse, err error)
	CreatePermittedNetworkWithContext(ctx context.Context, createPermittedNetworkOptions *dnssvcsv1.CreatePermittedNetworkOptions) (result *dnssvcsv1.PermittedNetwork, response *core.DetailedResponse, err error)
	ListResourceRecordsWithContext(ctx context.Context, listResourceRecordsOptions *dnssvcsv1.ListResourceRecordsOptions) (result *dnssvcsv1.ListResourceRecords, response *core.DetailedResponse, err error)
	CreateResourceRecordWithContext(ctx context.Context, createResourceRecordOptions *dnssvcsv1.CreateResourceRecordOptions) (result *dnssvcsv1.ResourceRecord, response *core.DetailedResponse, err error)
	DeleteResourceRecordWithContext(ctx context.Context, deleteResourceRecordOptions *dnssvcsv1.DeleteResourceRecordOptions) (response *core.DetailedResponse, err error)
	UpdateResourceRecordWithContext(ctx context.Context, updateResourceRecordOptions *dnssvcsv1.UpdateResourceRecordOptions) (result *dnssvcsv1.ResourceRecord, response *core.DetailedResponse, err error)
	NewResourceRecordInputRdataRdataARecord(ip string) (model *dnssvcsv1.ResourceRecordInputRdataRdataARecord, err error)
	NewResourceRecordInputRdataRdataCnameRecord(cname string) (model *dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord, err error)
	NewResourceRecordInputRdataRdataTxtRecord(text string) (model *dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord, err error)
	NewResourceRecordUpdateInputRdataRdataARecord(ip string) (model *dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord, err error)
	NewResourceRecordUpdateInputRdataRdataCnameRecord(cname string) (model *dnssvcsv1.ResourceRecordUpdateInputRdataRdataCnameRecord, err error)
	NewResourceRecordUpdateInputRdataRdataTxtRecord(text string) (model *dnssvcsv1.ResourceRecordUpdateInputRdataRdataTxtRecord, err error)
}

type ibmcloudService struct {
	publicZonesService   *zonesv1.ZonesV1
	publicRecordsService *dnsrecordsv1.DnsRecordsV1
	privateDNSService    *dnssvcsv1.DnsSvcsV1
}

func (i ibmcloudService) ListAllDDNSRecordsWithContext(ctx context.Context, listAllDNSRecordsOptions *dnsrecordsv1.ListAllDnsRecordsOptions) (result *dnsrecordsv1.ListDnsrecordsResp, response *core.DetailedResponse, err error) {
	return i.publicRecordsService.ListAllDnsRecordsWithContext(ctx, listAllDNSRecordsOptions)
}

func (i ibmcloudService) CreateDNSRecordWithContext(ctx context.Context, createDNSRecordOptions *dnsrecordsv1.CreateDnsRecordOptions) (result *dnsrecordsv1.DnsrecordResp, response *core.DetailedResponse, err error) {
	return i.publicRecordsService.CreateDnsRecordWithContext(ctx, createDNSRecordOptions)
}

func (i ibmcloudService) DeleteDNSRecordWithContext(ctx context.Context, deleteDNSRecordOptions *dnsrecordsv1.DeleteDnsRecordOptions) (result *dnsrecordsv1.DeleteDnsrecordResp, response *core.DetailedResponse, err error) {
	return i.publicRecordsService.DeleteDnsRecordWithContext(ctx, deleteDNSRecordOptions)
}

func (i ibmcloudService) UpdateDNSRecordWithContext(ctx context.Context, updateDNSRecordOptions *dnsrecordsv1.UpdateDnsRecordOptions) (result *dnsrecordsv1.DnsrecordResp, response *core.DetailedResponse, err error) {
	return i.publicRecordsService.UpdateDnsRecordWithContext(ctx, updateDNSRecordOptions)
}

func (i ibmcloudService) ListDnszonesWithContext(ctx context.Context, listDnszonesOptions *dnssvcsv1.ListDnszonesOptions) (result *dnssvcsv1.ListDnszones, response *core.DetailedResponse, err error) {
	return i.privateDNSService.ListDnszonesWithContext(ctx, listDnszonesOptions)
}

func (i ibmcloudService) GetDnszoneWithContext(ctx context.Context, getDnszoneOptions *dnssvcsv1.GetDnszoneOptions) (result *dnssvcsv1.Dnszone, response *core.DetailedResponse, err error) {
	return i.privateDNSService.GetDnszoneWithContext(ctx, getDnszoneOptions)
}

func (i ibmcloudService) CreatePermittedNetworkWithContext(ctx context.Context, createPermittedNetworkOptions *dnssvcsv1.CreatePermittedNetworkOptions) (result *dnssvcsv1.PermittedNetwork, response *core.DetailedResponse, err error) {
	return i.privateDNSService.CreatePermittedNetworkWithContext(ctx, createPermittedNetworkOptions)
}

func (i ibmcloudService) ListResourceRecordsWithContext(ctx context.Context, listResourceRecordsOptions *dnssvcsv1.ListResourceRecordsOptions) (result *dnssvcsv1.ListResourceRecords, response *core.DetailedResponse, err error) {
	return i.privateDNSService.ListResourceRecordsWithContext(ctx, listResourceRecordsOptions)
}

func (i ibmcloudService) CreateResourceRecordWithContext(ctx context.Context, createResourceRecordOptions *dnssvcsv1.CreateResourceRecordOptions) (result *dnssvcsv1.ResourceRecord, response *core.DetailedResponse, err error) {
	return i.privateDNSService.CreateResourceRecordWithContext(ctx, createResourceRecordOptions)
}

func (i ibmcloudService) DeleteResourceRecordWithContext(ctx context.Context, deleteResourceRecordOptions *dnssvcsv1.DeleteResourceRecordOptions) (response *core.DetailedResponse, err error) {
	return i.privateDNSService.DeleteResourceRecordWithContext(ctx, deleteResourceRecordOptions)
}

func (i ibmcloudService) UpdateResourceRecordWithContext(ctx context.Context, updateResourceRecordOptions *dnssvcsv1.UpdateResourceRecordOptions) (result *dnssvcsv1.ResourceRecord, response *core.DetailedResponse, err error) {
	return i.privateDNSService.UpdateResourceRecordWithContext(ctx, updateResourceRecordOptions)
}

func (i ibmcloudService) NewResourceRecordInputRdataRdataARecord(ip string) (model *dnssvcsv1.ResourceRecordInputRdataRdataARecord, err error) {
	return i.privateDNSService.NewResourceRecordInputRdataRdataARecord(ip)
}

func (i ibmcloudService) NewResourceRecordInputRdataRdataCnameRecord(cname string) (model *dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord, err error) {
	return i.privateDNSService.NewResourceRecordInputRdataRdataCnameRecord(cname)
}

func (i ibmcloudService) NewResourceRecordInputRdataRdataTxtRecord(text string) (model *dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord, err error) {
	return i.privateDNSService.NewResourceRecordInputRdataRdataTxtRecord(text)
}

func (i ibmcloudService) NewResourceRecordUpdateInputRdataRdataARecord(ip string) (model *dnssvcsv1.ResourceRecordUpdateInputRdataRdataARecord, err error) {
	return i.privateDNSService.NewResourceRecordUpdateInputRdataRdataARecord(ip)
}

func (i ibmcloudService) NewResourceRecordUpdateInputRdataRdataCnameRecord(cname string) (model *dnssvcsv1.ResourceRecordUpdateInputRdataRdataCnameRecord, err error) {
	return i.privateDNSService.NewResourceRecordUpdateInputRdataRdataCnameRecord(cname)
}

func (i ibmcloudService) NewResourceRecordUpdateInputRdataRdataTxtRecord(text string) (model *dnssvcsv1.ResourceRecordUpdateInputRdataRdataTxtRecord, err error) {
	return i.privateDNSService.NewResourceRecordUpdateInputRdataRdataTxtRecord(text)
}

// IBMCloudProvider is an implementation of Provider for IBM Cloud DNS.
type IBMCloudProvider struct {
	provider.BaseProvider
	source Source
	Client ibmcloudClient
	// only consider hosted zones managing domains ending in this suffix
	domainFilter     endpoint.DomainFilter
	zoneIDFilter     provider.ZoneIDFilter
	instanceID       string
	privateZone      bool
	proxiedByDefault bool
	DryRun           bool
}

type ibmcloudConfig struct {
	Endpoint   string `json:"endpoint" yaml:"endpoint"`
	APIKey     string `json:"apiKey" yaml:"apiKey"`
	CRN        string `json:"instanceCrn" yaml:"instanceCrn"`
	IAMURL     string `json:"iamUrl" yaml:"iamUrl"`
	InstanceID string `json:"-" yaml:"-"`
}

// ibmcloudChange differentiates between ChangActions
type ibmcloudChange struct {
	Action                string
	PublicResourceRecord  dnsrecordsv1.DnsrecordDetails
	PrivateResourceRecord dnssvcsv1.ResourceRecord
}

func getConfig(configFile string) (*ibmcloudConfig, error) {
	contents, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read IBM Cloud config file '%s': %v", configFile, err)
	}
	cfg := &ibmcloudConfig{}
	err = yaml.Unmarshal(contents, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read IBM Cloud config file '%s': %v", configFile, err)
	}

	return cfg, nil
}

func (c *ibmcloudConfig) Validate(authenticator core.Authenticator, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter) (ibmcloudService, bool, error) {
	var service ibmcloudService
	isPrivate := false
	log.Debugf("filters: %v, %v", domainFilter.Filters, zoneIDFilter.ZoneIDs)
	if domainFilter.Filters[0] == "" && zoneIDFilter.ZoneIDs[0] == "" {
		return service, isPrivate, fmt.Errorf("at lease one of filters: 'domain-filter', 'zone-id-filter' needed")
	}

	crn, err := crn.Parse(c.CRN)
	if err != nil {
		return service, isPrivate, err
	}
	log.Infof("IBM Cloud Service: %s", crn.ServiceName)
	c.InstanceID = crn.ServiceInstance

	switch {
	case strings.Contains(crn.ServiceName, "internet-svcs"):
		if len(domainFilter.Filters) > 1 || len(zoneIDFilter.ZoneIDs) > 1 {
			return service, isPrivate, fmt.Errorf("for public zone, only one domain id filter or domain name filter allowed")
		}
		var zoneID string
		// Public DNS service
		service.publicZonesService, err = zonesv1.NewZonesV1(&zonesv1.ZonesV1Options{
			Authenticator: authenticator,
			Crn:           core.StringPtr(c.CRN),
		})
		if err != nil {
			return service, isPrivate, fmt.Errorf("failed to initialize ibmcloud public zones client: %v", err)
		}
		if c.Endpoint != "" {
			service.publicZonesService.SetServiceURL(c.Endpoint)
		}

		zonesResp, _, err := service.publicZonesService.ListZones(&zonesv1.ListZonesOptions{})
		if err != nil {
			return service, isPrivate, fmt.Errorf("failed to list ibmcloud public zones: %v", err)
		}
		for _, zone := range zonesResp.Result {
			log.Debugf("zoneName: %s, zoneID: %s", *zone.Name, *zone.ID)
			if len(domainFilter.Filters[0]) != 0 && domainFilter.Match(*zone.Name) {
				log.Debugf("zone %s found.", *zone.ID)
				zoneID = *zone.ID
				break
			}
			if len(zoneIDFilter.ZoneIDs[0]) != 0 && zoneIDFilter.Match(*zone.ID) {
				log.Debugf("zone %s found.", *zone.ID)
				zoneID = *zone.ID
				break
			}
		}
		if len(zoneID) == 0 {
			return service, isPrivate, fmt.Errorf("no matched zone found")
		}

		service.publicRecordsService, err = dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
			Authenticator:  authenticator,
			Crn:            core.StringPtr(c.CRN),
			ZoneIdentifier: core.StringPtr(zoneID),
		})
		if err != nil {
			return service, isPrivate, fmt.Errorf("failed to initialize ibmcloud public records client: %v", err)
		}
		if c.Endpoint != "" {
			service.publicRecordsService.SetServiceURL(c.Endpoint)
		}
	case strings.Contains(crn.ServiceName, "dns-svcs"):
		isPrivate = true
		// Private DNS service
		service.privateDNSService, err = dnssvcsv1.NewDnsSvcsV1(&dnssvcsv1.DnsSvcsV1Options{
			Authenticator: authenticator,
		})
		if err != nil {
			return service, isPrivate, fmt.Errorf("failed to initialize ibmcloud private records client: %v", err)
		}
		if c.Endpoint != "" {
			service.privateDNSService.SetServiceURL(c.Endpoint)
		}
	default:
		return service, isPrivate, fmt.Errorf("IBM Cloud instance crn is not provided or invalid dns crn : %s", c.CRN)
	}

	return service, isPrivate, nil
}

// NewIBMCloudProvider creates a new IBMCloud provider.
//
// Returns the provider or an error if a provider could not be created.
func NewIBMCloudProvider(configFile string, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, source source.Source, proxiedByDefault bool, dryRun bool) (*IBMCloudProvider, error) {
	cfg, err := getConfig(configFile)
	if err != nil {
		return nil, err
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: cfg.APIKey,
	}
	if cfg.IAMURL != "" {
		authenticator = &core.IamAuthenticator{
			ApiKey: cfg.APIKey,
			URL:    cfg.IAMURL,
		}
	}

	client, isPrivate, err := cfg.Validate(authenticator, domainFilter, zoneIDFilter)
	if err != nil {
		return nil, err
	}

	provider := &IBMCloudProvider{
		Client:           client,
		source:           source,
		domainFilter:     domainFilter,
		zoneIDFilter:     zoneIDFilter,
		instanceID:       cfg.InstanceID,
		privateZone:      isPrivate,
		proxiedByDefault: proxiedByDefault,
		DryRun:           dryRun,
	}
	return provider, nil
}

// Records gets the current records.
//
// Returns the current records or an error if the operation failed.
func (p *IBMCloudProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, err error) {
	if p.privateZone {
		endpoints, err = p.privateRecords(ctx)
	} else {
		endpoints, err = p.publicRecords(ctx)
	}
	return endpoints, err
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *IBMCloudProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	log.Debugln("applying change...")
	ibmcloudChanges := []*ibmcloudChange{}
	for _, endpoint := range changes.Create {
		for _, target := range endpoint.Targets {
			ibmcloudChanges = append(ibmcloudChanges, p.newIBMCloudChange(recordCreate, endpoint, target))
		}
	}

	for i, desired := range changes.UpdateNew {
		current := changes.UpdateOld[i]

		add, remove, leave := provider.Difference(current.Targets, desired.Targets)

		log.Debugf("add: %v, remove: %v, leave: %v", add, remove, leave)
		for _, a := range add {
			ibmcloudChanges = append(ibmcloudChanges, p.newIBMCloudChange(recordCreate, desired, a))
		}

		for _, a := range leave {
			ibmcloudChanges = append(ibmcloudChanges, p.newIBMCloudChange(recordUpdate, desired, a))
		}

		for _, a := range remove {
			ibmcloudChanges = append(ibmcloudChanges, p.newIBMCloudChange(recordDelete, current, a))
		}
	}

	for _, endpoint := range changes.Delete {
		for _, target := range endpoint.Targets {
			ibmcloudChanges = append(ibmcloudChanges, p.newIBMCloudChange(recordDelete, endpoint, target))
		}
	}

	return p.submitChanges(ctx, ibmcloudChanges)
}

func (p *IBMCloudProvider) PropertyValuesEqual(name string, previous string, current string) bool {
	if name == proxyFilter {
		return plan.CompareBoolean(p.proxiedByDefault, name, previous, current)
	}

	return p.BaseProvider.PropertyValuesEqual(name, previous, current)
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (p *IBMCloudProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	adjustedEndpoints := []*endpoint.Endpoint{}
	for _, e := range endpoints {
		log.Debugf("adjusting endpont: %v", *e)
		if shouldBeProxied(e, p.proxiedByDefault) {
			e.RecordTTL = 0
		}

		adjustedEndpoints = append(adjustedEndpoints, e)
	}
	return adjustedEndpoints
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *IBMCloudProvider) submitChanges(ctx context.Context, changes []*ibmcloudChange) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	log.Debugln("submmiting change...")
	if p.privateZone {
		return p.submitChangesForPrivateDNS(ctx, changes)
	}
	return p.submitChangesForPublicDNS(ctx, changes)
}

// submitChangesForPublicDNS takes a zone and a collection of Changes and sends them as a single transaction on public dns.
func (p *IBMCloudProvider) submitChangesForPublicDNS(ctx context.Context, changes []*ibmcloudChange) error {
	records, err := p.listAllPublicRecords(ctx)
	if err != nil {
		return err
	}

	for _, change := range changes {
		logFields := log.Fields{
			"record": *change.PublicResourceRecord.Name,
			"type":   *change.PublicResourceRecord.Type,
			"ttl":    *change.PublicResourceRecord.TTL,
			"action": change.Action,
		}

		if p.DryRun {
			continue
		}

		log.WithFields(logFields).Info("Changing record.")

		if change.Action == recordUpdate {
			recordID := p.getPublicRecordID(records, change.PublicResourceRecord)
			if recordID == "" {
				log.WithFields(logFields).Errorf("failed to find previous record: %v", *change.PublicResourceRecord.Name)
				continue
			}
			p.updateRecord(ctx, "", recordID, change)
		} else if change.Action == recordDelete {
			recordID := p.getPublicRecordID(records, change.PublicResourceRecord)
			if recordID == "" {
				log.WithFields(logFields).Errorf("failed to find previous record: %v", *change.PublicResourceRecord.Name)
				continue
			}
			p.deleteRecord(ctx, "", recordID)
		} else if change.Action == recordCreate {
			p.createRecord(ctx, "", change)
		}
	}

	return nil
}

// submitChangesForPrivateDNS takes a zone and a collection of Changes and sends them as a single transaction on private dns.
func (p *IBMCloudProvider) submitChangesForPrivateDNS(ctx context.Context, changes []*ibmcloudChange) error {
	zones, err := p.privateZones(ctx)
	if err != nil {
		return err
	}
	// separate into per-zone change sets to be passed to the API.
	changesByPrivateZone := p.changesByPrivateZone(ctx, zones, changes)

	for zoneID, changes := range changesByPrivateZone {
		records, err := p.listAllPrivateRecords(ctx, zoneID)
		if err != nil {
			return err
		}

		for _, change := range changes {
			logFields := log.Fields{
				"record": *change.PrivateResourceRecord.Name,
				"type":   *change.PrivateResourceRecord.Type,
				"ttl":    *change.PrivateResourceRecord.TTL,
				"action": change.Action,
			}

			log.WithFields(logFields).Info("Changing record.")

			if p.DryRun {
				continue
			}

			if change.Action == recordUpdate {
				recordID := p.getPrivateRecordID(records, change.PrivateResourceRecord)
				if recordID == "" {
					log.WithFields(logFields).Errorf("failed to find previous record: %v", change.PrivateResourceRecord)
					continue
				}
				p.updateRecord(ctx, zoneID, recordID, change)
			} else if change.Action == recordDelete {
				recordID := p.getPrivateRecordID(records, change.PrivateResourceRecord)
				if recordID == "" {
					log.WithFields(logFields).Errorf("failed to find previous record: %v", change.PrivateResourceRecord)
					continue
				}
				p.deleteRecord(ctx, zoneID, recordID)
			} else if change.Action == recordCreate {
				p.createRecord(ctx, zoneID, change)
			}
		}
	}

	return nil
}

// privateZones return zones in private dns
func (p *IBMCloudProvider) privateZones(ctx context.Context) ([]dnssvcsv1.Dnszone, error) {
	result := []dnssvcsv1.Dnszone{}
	// if there is a zoneIDfilter configured
	// && if the filter isn't just a blank string (used in tests)
	if len(p.zoneIDFilter.ZoneIDs) > 0 && p.zoneIDFilter.ZoneIDs[0] != "" {
		log.Debugln("zoneIDFilter configured. only looking up zone IDs defined")
		for _, zoneID := range p.zoneIDFilter.ZoneIDs {
			log.Debugf("looking up zone %s", zoneID)
			detailResponse, _, err := p.Client.GetDnszoneWithContext(ctx, &dnssvcsv1.GetDnszoneOptions{
				InstanceID: core.StringPtr(p.instanceID),
				DnszoneID:  core.StringPtr(zoneID),
			})
			if err != nil {
				log.Errorf("zone %s lookup failed, %v", zoneID, err)
				continue
			}
			log.WithFields(log.Fields{
				"zoneName": *detailResponse.Name,
				"zoneID":   *detailResponse.ID,
			}).Debugln("adding zone for consideration")
			result = append(result, *detailResponse)
		}
		return result, nil
	}

	log.Debugln("no zoneIDFilter configured, looking at all zones")

	zonesResponse, _, err := p.Client.ListDnszonesWithContext(ctx, &dnssvcsv1.ListDnszonesOptions{
		InstanceID: core.StringPtr(p.instanceID),
	})
	if err != nil {
		return nil, err
	}

	for _, zone := range zonesResponse.Dnszones {
		if !p.domainFilter.Match(*zone.Name) {
			log.Debugf("zone %s not in domain filter", *zone.Name)
			continue
		}
		result = append(result, zone)
	}

	return result, nil
}

// activePrivateZone active zone with new records add if not active
func (p *IBMCloudProvider) activePrivateZone(ctx context.Context, zoneID, vpc string) {
	permittedNetworkVpc := &dnssvcsv1.PermittedNetworkVpc{
		VpcCrn: core.StringPtr(vpc),
	}
	createPermittedNetworkOptions := &dnssvcsv1.CreatePermittedNetworkOptions{
		InstanceID:       core.StringPtr(p.instanceID),
		DnszoneID:        core.StringPtr(zoneID),
		PermittedNetwork: permittedNetworkVpc,
		Type:             core.StringPtr("vpc"),
	}
	_, _, err := p.Client.CreatePermittedNetworkWithContext(ctx, createPermittedNetworkOptions)
	if err != nil {
		log.Errorf("failed to active zone %s in VPC %s with error: %v", zoneID, vpc, err)
	}
}

// changesByPrivateZone separates a multi-zone change into a single change per zone.
func (p *IBMCloudProvider) changesByPrivateZone(ctx context.Context, zones []dnssvcsv1.Dnszone, changeSet []*ibmcloudChange) map[string][]*ibmcloudChange {
	changes := make(map[string][]*ibmcloudChange)
	zoneNameIDMapper := provider.ZoneIDName{}
	for _, z := range zones {
		zoneNameIDMapper.Add(*z.ID, *z.Name)
		changes[*z.ID] = []*ibmcloudChange{}
	}

	for _, c := range changeSet {
		zoneID, _ := zoneNameIDMapper.FindZone(*c.PrivateResourceRecord.Name)
		if zoneID == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", *c.PrivateResourceRecord.Name)
			continue
		}
		changes[zoneID] = append(changes[zoneID], c)
	}

	return changes
}

func (p *IBMCloudProvider) publicRecords(ctx context.Context) ([]*endpoint.Endpoint, error) {
	log.Debugf("Listing records on public zone")
	dnsRecords, err := p.listAllPublicRecords(ctx)
	if err != nil {
		return nil, err
	}
	return p.groupPublicRecords(dnsRecords), nil
}

func (p *IBMCloudProvider) listAllPublicRecords(ctx context.Context) ([]dnsrecordsv1.DnsrecordDetails, error) {
	var dnsRecords []dnsrecordsv1.DnsrecordDetails
	page := 1
GETRECORDS:
	listAllDNSRecordsOptions := &dnsrecordsv1.ListAllDnsRecordsOptions{
		Page: core.Int64Ptr(int64(page)),
	}
	records, _, err := p.Client.ListAllDDNSRecordsWithContext(ctx, listAllDNSRecordsOptions)
	if err != nil {
		return dnsRecords, err
	}
	dnsRecords = append(dnsRecords, records.Result...)
	// Loop if more records exist
	if *records.ResultInfo.TotalCount > int64(page*100) {
		page = page + 1
		log.Debugf("More than one pages records found, page: %d", page)
		goto GETRECORDS
	}
	return dnsRecords, nil
}

func (p *IBMCloudProvider) groupPublicRecords(records []dnsrecordsv1.DnsrecordDetails) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}

	// group supported records by name and type
	groups := map[string][]dnsrecordsv1.DnsrecordDetails{}

	for _, r := range records {
		if !provider.SupportedRecordType(*r.Type) {
			continue
		}

		groupBy := *r.Name + *r.Type
		if _, ok := groups[groupBy]; !ok {
			groups[groupBy] = []dnsrecordsv1.DnsrecordDetails{}
		}

		groups[groupBy] = append(groups[groupBy], r)
	}

	// create single endpoint with all the targets for each name/type
	for _, records := range groups {
		targets := make([]string, len(records))
		for i, record := range records {
			targets[i] = *record.Content
		}

		ep := endpoint.NewEndpointWithTTL(
			*records[0].Name,
			*records[0].Type,
			endpoint.TTL(*records[0].TTL),
			targets...).WithProviderSpecific(proxyFilter, strconv.FormatBool(*records[0].Proxied))

		log.Debugf(
			"Found %s record for '%s' with target '%s'.",
			ep.RecordType,
			ep.DNSName,
			ep.Targets,
		)

		endpoints = append(endpoints, ep)
	}
	return endpoints
}

func (p *IBMCloudProvider) privateRecords(ctx context.Context) ([]*endpoint.Endpoint, error) {
	log.Debugf("Listing records on private zone")
	var vpc string
	zones, err := p.privateZones(ctx)
	if err != nil {
		return nil, err
	}
	sources, err := p.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}
	// Filter VPC annoation for private zone active
	for _, source := range sources {
		vpc = checkVPCAnnotation(source)
		if len(vpc) > 0 {
			log.Debugf("VPC found: %s", vpc)
			break
		}
	}

	endpoints := []*endpoint.Endpoint{}
	for _, zone := range zones {
		if len(vpc) > 0 && *zone.State == zoneStatePendingNetwork {
			log.Debugf("active zone: %s", *zone.ID)
			p.activePrivateZone(ctx, *zone.ID, vpc)
		}

		dnsRecords, err := p.listAllPrivateRecords(ctx, *zone.ID)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, p.groupPrivateRecords(dnsRecords)...)
	}

	return endpoints, nil
}

func (p *IBMCloudProvider) listAllPrivateRecords(ctx context.Context, zoneID string) ([]dnssvcsv1.ResourceRecord, error) {
	var dnsRecords []dnssvcsv1.ResourceRecord
	offset := 0
GETRECORDS:
	listResourceRecordsOptions := &dnssvcsv1.ListResourceRecordsOptions{
		InstanceID: core.StringPtr(p.instanceID),
		DnszoneID:  core.StringPtr(zoneID),
		Offset:     core.Int64Ptr(int64(offset)),
	}
	records, _, err := p.Client.ListResourceRecordsWithContext(ctx, listResourceRecordsOptions)
	if err != nil {
		return dnsRecords, err
	}
	oRecords := records.ResourceRecords
	dnsRecords = append(dnsRecords, oRecords...)
	// Loop if more records exist
	if int64(offset+1) < *records.TotalCount && int64(offset+200) < *records.TotalCount {
		offset = offset + 200
		log.Debugf("More than one pages records found, page: %d", offset/200+1)
		goto GETRECORDS
	}
	return dnsRecords, nil
}

func (p *IBMCloudProvider) groupPrivateRecords(records []dnssvcsv1.ResourceRecord) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}
	// group supported records by name and type
	groups := map[string][]dnssvcsv1.ResourceRecord{}
	for _, r := range records {
		if !provider.SupportedRecordType(*r.Type) || !privateTypeSupported[*r.Type] {
			continue
		}
		rname := *r.Name
		rtype := *r.Type
		groupBy := rname + rtype
		if _, ok := groups[groupBy]; !ok {
			groups[groupBy] = []dnssvcsv1.ResourceRecord{}
		}

		groups[groupBy] = append(groups[groupBy], r)
	}

	// create single endpoint with all the targets for each name/type
	for _, records := range groups {
		targets := make([]string, len(records))
		for i, record := range records {
			data := record.Rdata.(map[string]interface{})
			log.Debugf("record data: %v", data)
			switch *record.Type {
			case "A":
				if !isNil(data["ip"]) {
					targets[i] = data["ip"].(string)
				}
			case "CNAME":
				if !isNil(data["cname"]) {
					targets[i] = data["cname"].(string)
				}
			case "TXT":
				if !isNil(data["text"]) {
					targets[i] = data["text"].(string)
				}
				log.Debugf("text record data: %v", targets[i])
			}
		}

		ep := endpoint.NewEndpointWithTTL(
			*records[0].Name,
			*records[0].Type,
			endpoint.TTL(*records[0].TTL), targets...)

		log.Debugf(
			"Found %s record for '%s' with target '%s'.",
			ep.RecordType,
			ep.DNSName,
			ep.Targets,
		)

		endpoints = append(endpoints, ep)
	}
	return endpoints
}

func (p *IBMCloudProvider) getPublicRecordID(records []dnsrecordsv1.DnsrecordDetails, record dnsrecordsv1.DnsrecordDetails) string {
	for _, zoneRecord := range records {
		if *zoneRecord.Name == *record.Name && *zoneRecord.Type == *record.Type && *zoneRecord.Content == *record.Content {
			return *zoneRecord.ID
		}
	}
	return ""
}

func (p *IBMCloudProvider) getPrivateRecordID(records []dnssvcsv1.ResourceRecord, record dnssvcsv1.ResourceRecord) string {
	for _, zoneRecord := range records {
		if *zoneRecord.Name == *record.Name && *zoneRecord.Type == *record.Type {
			return *zoneRecord.ID
		}
	}
	return ""
}

func (p *IBMCloudProvider) newIBMCloudChange(action string, endpoint *endpoint.Endpoint, target string) *ibmcloudChange {
	ttl := defaultPublicRecordTTL
	proxied := shouldBeProxied(endpoint, p.proxiedByDefault)

	if endpoint.RecordTTL.IsConfigured() {
		ttl = int(endpoint.RecordTTL)
	}

	if p.privateZone {
		var rData interface{}
		switch endpoint.RecordType {
		case "A":
			rData = &dnssvcsv1.ResourceRecordInputRdataRdataARecord{
				Ip: core.StringPtr(target),
			}
		case "CNAME":
			rData = &dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord{
				Cname: core.StringPtr(target),
			}
		case "TXT":
			rData = &dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord{
				Text: core.StringPtr(target),
			}
		}
		return &ibmcloudChange{
			Action: action,
			PrivateResourceRecord: dnssvcsv1.ResourceRecord{
				Name:  core.StringPtr(endpoint.DNSName),
				TTL:   core.Int64Ptr(int64(ttl)),
				Type:  core.StringPtr(endpoint.RecordType),
				Rdata: rData,
			},
		}
	}

	return &ibmcloudChange{
		Action: action,
		PublicResourceRecord: dnsrecordsv1.DnsrecordDetails{
			Name:    core.StringPtr(endpoint.DNSName),
			TTL:     core.Int64Ptr(int64(ttl)),
			Proxied: core.BoolPtr(proxied),
			Type:    core.StringPtr(endpoint.RecordType),
			Content: core.StringPtr(target),
		},
	}
}

func (p *IBMCloudProvider) createRecord(ctx context.Context, zoneID string, change *ibmcloudChange) {
	if p.privateZone {
		createResourceRecordOptions := &dnssvcsv1.CreateResourceRecordOptions{
			InstanceID: core.StringPtr(p.instanceID),
			DnszoneID:  core.StringPtr(zoneID),
			Name:       change.PrivateResourceRecord.Name,
			Type:       change.PrivateResourceRecord.Type,
			TTL:        change.PrivateResourceRecord.TTL,
		}
		switch *change.PrivateResourceRecord.Type {
		case "A":
			data, _ := change.PrivateResourceRecord.Rdata.(*dnssvcsv1.ResourceRecordInputRdataRdataARecord)
			aData, _ := p.Client.NewResourceRecordInputRdataRdataARecord(*data.Ip)
			createResourceRecordOptions.SetRdata(aData)
		case "CNAME":
			data, _ := change.PrivateResourceRecord.Rdata.(*dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord)
			cnameData, _ := p.Client.NewResourceRecordInputRdataRdataCnameRecord(*data.Cname)
			createResourceRecordOptions.SetRdata(cnameData)
		case "TXT":
			data, _ := change.PrivateResourceRecord.Rdata.(*dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord)
			txtData, _ := p.Client.NewResourceRecordInputRdataRdataTxtRecord(*data.Text)
			createResourceRecordOptions.SetRdata(txtData)
		}
		_, _, err := p.Client.CreateResourceRecordWithContext(ctx, createResourceRecordOptions)
		if err != nil {
			log.Errorf("failed to create %s type record named %s: %v", *change.PrivateResourceRecord.Type, *change.PrivateResourceRecord.Name, err)
		}
	} else {
		createDNSRecordOptions := &dnsrecordsv1.CreateDnsRecordOptions{
			Name:    change.PublicResourceRecord.Name,
			Type:    change.PublicResourceRecord.Type,
			TTL:     change.PublicResourceRecord.TTL,
			Content: change.PublicResourceRecord.Content,
		}
		_, _, err := p.Client.CreateDNSRecordWithContext(ctx, createDNSRecordOptions)
		if err != nil {
			log.Errorf("failed to create %s type record named %s: %v", *change.PublicResourceRecord.Type, *change.PublicResourceRecord.Name, err)
		}
	}
}

func (p *IBMCloudProvider) updateRecord(ctx context.Context, zoneID, recordID string, change *ibmcloudChange) {
	if p.privateZone {
		updateResourceRecordOptions := &dnssvcsv1.UpdateResourceRecordOptions{
			InstanceID: core.StringPtr(p.instanceID),
			DnszoneID:  core.StringPtr(zoneID),
			RecordID:   core.StringPtr(recordID),
			Name:       change.PrivateResourceRecord.Name,
			TTL:        change.PrivateResourceRecord.TTL,
		}
		switch *change.PrivateResourceRecord.Type {
		case "A":
			data, _ := change.PrivateResourceRecord.Rdata.(*dnssvcsv1.ResourceRecordInputRdataRdataARecord)
			aData, _ := p.Client.NewResourceRecordUpdateInputRdataRdataARecord(*data.Ip)
			updateResourceRecordOptions.SetRdata(aData)
		case "CNAME":
			data, _ := change.PrivateResourceRecord.Rdata.(*dnssvcsv1.ResourceRecordInputRdataRdataCnameRecord)
			cnameData, _ := p.Client.NewResourceRecordUpdateInputRdataRdataCnameRecord(*data.Cname)
			updateResourceRecordOptions.SetRdata(cnameData)
		case "TXT":
			data, _ := change.PrivateResourceRecord.Rdata.(*dnssvcsv1.ResourceRecordInputRdataRdataTxtRecord)
			txtData, _ := p.Client.NewResourceRecordUpdateInputRdataRdataTxtRecord(*data.Text)
			updateResourceRecordOptions.SetRdata(txtData)
		}
		_, _, err := p.Client.UpdateResourceRecordWithContext(ctx, updateResourceRecordOptions)
		if err != nil {
			log.Errorf("failed to update %s type record named %s: %v", *change.PublicResourceRecord.Type, *change.PublicResourceRecord.Name, err)
		}
	} else {
		updateDNSRecordOptions := &dnsrecordsv1.UpdateDnsRecordOptions{
			DnsrecordIdentifier: &recordID,
			Name:                change.PublicResourceRecord.Name,
			Type:                change.PublicResourceRecord.Type,
			TTL:                 change.PublicResourceRecord.TTL,
			Content:             change.PublicResourceRecord.Content,
			Proxied:             change.PublicResourceRecord.Proxied,
		}
		_, _, err := p.Client.UpdateDNSRecordWithContext(ctx, updateDNSRecordOptions)
		if err != nil {
			log.Errorf("failed to update %s type record named %s: %v", *change.PublicResourceRecord.Type, *change.PublicResourceRecord.Name, err)
		}
	}
}

func (p *IBMCloudProvider) deleteRecord(ctx context.Context, zoneID, recordID string) {
	if p.privateZone {
		deleteResourceRecordOptions := &dnssvcsv1.DeleteResourceRecordOptions{
			InstanceID: core.StringPtr(p.instanceID),
			DnszoneID:  core.StringPtr(zoneID),
			RecordID:   core.StringPtr(recordID),
		}
		_, err := p.Client.DeleteResourceRecordWithContext(ctx, deleteResourceRecordOptions)
		if err != nil {
			log.Errorf("failed to delete record %s: %v", recordID, err)
		}
	} else {
		deleteDNSRecordOptions := &dnsrecordsv1.DeleteDnsRecordOptions{
			DnsrecordIdentifier: &recordID,
		}
		_, _, err := p.Client.DeleteDNSRecordWithContext(ctx, deleteDNSRecordOptions)
		if err != nil {
			log.Errorf("failed to delete record %s: %v", recordID, err)
		}
	}
}

func shouldBeProxied(endpoint *endpoint.Endpoint, proxiedByDefault bool) bool {
	proxied := proxiedByDefault

	for _, v := range endpoint.ProviderSpecific {
		if v.Name == proxyFilter {
			b, err := strconv.ParseBool(v.Value)
			if err != nil {
				log.Errorf("Failed to parse annotation [%s]: %v", proxyFilter, err)
			} else {
				proxied = b
			}
			break
		}
	}

	if proxyTypeNotSupported[endpoint.RecordType] || strings.Contains(endpoint.DNSName, "*") {
		proxied = false
	}
	return proxied
}

func checkVPCAnnotation(endpoint *endpoint.Endpoint) string {
	var vpc string
	for _, v := range endpoint.ProviderSpecific {
		if v.Name == vpcFilter {
			vpcCrn, err := crn.Parse(v.Value)
			if vpcCrn.ResourceType != "vpc" || err != nil {
				log.Errorf("Failed to parse vpc [%s]: %v", v.Value, err)
			} else {
				vpc = v.Value
			}
			break
		}
	}
	return vpc
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
