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
	"strconv"
	"strings"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/crn"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"
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

const (
	// cisCreate is a ChangeAction enum value
	cisCreate = "CREATE"
	// cisDelete is a ChangeAction enum value
	cisDelete = "DELETE"
	// cisUpdate is a ChangeAction enum value
	cisUpdate = "UPDATE"
	// defaultCISRecordTTL 1 = automatic
	defaultCISRecordTTL = 1
)

type RecordsClient interface {
	ListAllDnsRecordsWithContext(ctx context.Context, listAllDnsRecordsOptions *dnsrecordsv1.ListAllDnsRecordsOptions) (result *dnsrecordsv1.ListDnsrecordsResp, response *core.DetailedResponse, err error)
	CreateDnsRecordWithContext(ctx context.Context, createDnsRecordOptions *dnsrecordsv1.CreateDnsRecordOptions) (result *dnsrecordsv1.DnsrecordResp, response *core.DetailedResponse, err error)
	DeleteDnsRecordWithContext(ctx context.Context, deleteDnsRecordOptions *dnsrecordsv1.DeleteDnsRecordOptions) (result *dnsrecordsv1.DeleteDnsrecordResp, response *core.DetailedResponse, err error)
	UpdateDnsRecordWithContext(ctx context.Context, updateDnsRecordOptions *dnsrecordsv1.UpdateDnsRecordOptions) (result *dnsrecordsv1.DnsrecordResp, response *core.DetailedResponse, err error)
}

type recordService struct {
	service *dnsrecordsv1.DnsRecordsV1
}

func (r recordService) ListAllDnsRecordsWithContext(ctx context.Context, listAllDnsRecordsOptions *dnsrecordsv1.ListAllDnsRecordsOptions) (result *dnsrecordsv1.ListDnsrecordsResp, response *core.DetailedResponse, err error) {
	return r.service.ListAllDnsRecordsWithContext(ctx, listAllDnsRecordsOptions)
}

func (r recordService) CreateDnsRecordWithContext(ctx context.Context, createDnsRecordOptions *dnsrecordsv1.CreateDnsRecordOptions) (result *dnsrecordsv1.DnsrecordResp, response *core.DetailedResponse, err error) {
	return r.service.CreateDnsRecordWithContext(ctx, createDnsRecordOptions)
}

func (r recordService) DeleteDnsRecordWithContext(ctx context.Context, deleteDnsRecordOptions *dnsrecordsv1.DeleteDnsRecordOptions) (result *dnsrecordsv1.DeleteDnsrecordResp, response *core.DetailedResponse, err error) {
	return r.service.DeleteDnsRecordWithContext(ctx, deleteDnsRecordOptions)
}

func (r recordService) UpdateDnsRecordWithContext(ctx context.Context, updateDnsRecordOptions *dnsrecordsv1.UpdateDnsRecordOptions) (result *dnsrecordsv1.DnsrecordResp, response *core.DetailedResponse, err error) {
	return r.service.UpdateDnsRecordWithContext(ctx, updateDnsRecordOptions)
}

// IBMCloudProvider is an implementation of Provider for IBM Cloud DNS.
type IBMCloudProvider struct {
	provider.BaseProvider
	recordsClient RecordsClient
	// only consider hosted zones managing domains ending in this suffix
	domainFilter     endpoint.DomainFilter
	zoneIDFilter     provider.ZoneIDFilter
	proxiedByDefault bool
	DryRun           bool
}

type ibmcloudConfig struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
	APIKey   string `json:"apiKey" yaml:"apiKey"`
	CRN      string `json:"instanceCrn" yaml:"instanceCrn"`
	DomainID string `json:"domainID" yaml:"domainID"`
	IAMURL   string `json:"iamUrl" yaml:"iamUrl"`
	VPCID    string `json:"vpcId" yaml:"vpcId"`
}

// cisChange differentiates between ChangActions
type cisChange struct {
	Action         string
	ResourceRecord dnsrecordsv1.DnsrecordDetails
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

	crn, err := crn.Parse(cfg.CRN)
	if !strings.Contains(crn.ServiceName, "internet-svcs") || err != nil {
		return nil, fmt.Errorf("IBM Cloud CIS instance crn is not provided or invalid'%s': %v", cfg.CRN, err)
	}
	if cfg.DomainID == "" {
		return nil, fmt.Errorf("IBM Cloud CIS Domain ID is not provided or invalid'%s': %v", cfg.DomainID, err)
	}
	return cfg, nil
}

// NewIBMCloudProvider creates a new IBMCloud provider.
//
// Returns the provider or an error if a provider could not be created.
func NewIBMCloudProvider(configFile string, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, proxiedByDefault bool, dryRun bool) (*IBMCloudProvider, error) {
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

	recordsClient, err := dnsrecordsv1.NewDnsRecordsV1(&dnsrecordsv1.DnsRecordsV1Options{
		Authenticator:  authenticator,
		Crn:            &cfg.CRN,
		ZoneIdentifier: &cfg.DomainID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize ibmcloud records client: %v", err)
	}
	if cfg.Endpoint != "" {
		recordsClient.SetServiceURL(cfg.Endpoint)
	}

	logFields := log.Fields{
		"instanceCrn": cfg.CRN,
		"domainID":    cfg.DomainID,
		"proxied":     strconv.FormatBool(proxiedByDefault),
	}
	log.WithFields(logFields).Info("Initializing ibmcloud.")
	provider := &IBMCloudProvider{
		recordsClient:    recordsClient,
		domainFilter:     domainFilter,
		zoneIDFilter:     zoneIDFilter,
		proxiedByDefault: proxiedByDefault,
		DryRun:           dryRun,
	}
	return provider, nil
}

// Records gets the current records.
//
// Returns the current records or an error if the operation failed.
func (p *IBMCloudProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	records, _, err := p.recordsClient.ListAllDnsRecordsWithContext(ctx, &dnsrecordsv1.ListAllDnsRecordsOptions{})
	if err != nil {
		return nil, err
	}

	return groupByNameAndType(records.Result), nil
}

/*
// Zones returns the list of hosted zones.
func (p *IBMCloudProvider) Zones(ctx context.Context) ([]zonesv1.ZoneDetails, error) {
	var result []zonesv1.ZoneDetails
	if len(p.zoneIDFilter.ZoneIDs) > 0 && p.zoneIDFilter.ZoneIDs[0] != "" {
		log.Debugln("zoneIDFilter configured. only looking up zone IDs defined")
		for _, zoneID := range p.zoneIDFilter.ZoneIDs {
			log.Debugf("looking up zone %s", zoneID)
			getZonesOptions := new(zonesv1.GetZoneOptions)
			getZonesOptions.SetZoneIdentifier(zoneID)
			zoneresp, _, err := p.zonesClient.GetZoneWithContext(ctx, getZonesOptions)
			if err != nil {
				log.Errorf("zone %s lookup failed, %v", zoneID, err)
				continue
			}
			log.WithFields(log.Fields{
				"zoneName": zoneresp.Result.Name,
				"zoneID":   zoneresp.Result.ID,
			}).Debugln("adding zone for consideration")
			result = append(result, *zoneresp.Result)
		}
		return result, nil
	}

	log.Debugln("no zoneIDFilter configured, looking at all zones")

	zonesResponse, _, err := p.zonesClient.ListZonesWithContext(ctx, &zonesv1.ListZonesOptions{})
	if err != nil {
		return nil, err
	}

	for _, zone := range zonesResponse.Result {
		if !p.domainFilter.Match(zone.Name) {
			log.Debugf("zone %s not in domain filter", zone.Name)
			continue
		}
		result = append(result, zone)
	}

	return result, nil
}

*/
// ApplyChanges applies a given set of changes in a given zone.
func (p *IBMCloudProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	cisChanges := []*cisChange{}

	for _, endpoint := range changes.Create {
		for _, target := range endpoint.Targets {
			cisChanges = append(cisChanges, p.newCISChange(cisCreate, endpoint, target))
		}
	}

	for i, desired := range changes.UpdateNew {
		current := changes.UpdateOld[i]

		add, remove, leave := provider.Difference(current.Targets, desired.Targets)

		for _, a := range add {
			cisChanges = append(cisChanges, p.newCISChange(cisCreate, desired, a))
		}

		for _, a := range leave {
			cisChanges = append(cisChanges, p.newCISChange(cisUpdate, desired, a))
		}

		for _, a := range remove {
			cisChanges = append(cisChanges, p.newCISChange(cisDelete, current, a))
		}
	}

	for _, endpoint := range changes.Delete {
		for _, target := range endpoint.Targets {
			cisChanges = append(cisChanges, p.newCISChange(cisDelete, endpoint, target))
		}
	}

	return p.submitChanges(ctx, cisChanges)
}

func (p *IBMCloudProvider) PropertyValuesEqual(name string, previous string, current string) bool {
	if name == source.CisProxiedKey {
		return plan.CompareBoolean(p.proxiedByDefault, name, previous, current)
	}

	return p.BaseProvider.PropertyValuesEqual(name, previous, current)
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (p *IBMCloudProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	adjustedEndpoints := []*endpoint.Endpoint{}
	for _, e := range endpoints {
		if shouldBeProxied(e, p.proxiedByDefault) {
			e.RecordTTL = 0
		}
		adjustedEndpoints = append(adjustedEndpoints, e)
	}
	return adjustedEndpoints
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *IBMCloudProvider) submitChanges(ctx context.Context, changes []*cisChange) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	records, _, err := p.recordsClient.ListAllDnsRecordsWithContext(ctx, &dnsrecordsv1.ListAllDnsRecordsOptions{})
	if err != nil {
		return fmt.Errorf("could not fetch records from zone, %v", err)
	}
	for _, change := range changes {
		logFields := log.Fields{
			"record": *change.ResourceRecord.Name,
			"type":   *change.ResourceRecord.Type,
			"ttl":    *change.ResourceRecord.TTL,
			"action": change.Action,
		}

		log.WithFields(logFields).Info("Changing record.")

		if p.DryRun {
			continue
		}

		if change.Action == cisUpdate {
			recordID := p.getRecordID(records.Result, change.ResourceRecord)
			if recordID == "" {
				log.WithFields(logFields).Errorf("failed to find previous record: %v", change.ResourceRecord)
				continue
			}
			p.updateRecord(ctx, recordID, change)
		} else if change.Action == cisDelete {
			recordID := p.getRecordID(records.Result, change.ResourceRecord)
			if recordID == "" {
				log.WithFields(logFields).Errorf("failed to find previous record: %v", change.ResourceRecord)
				continue
			}
			p.deleteRecord(ctx, recordID)
		} else if change.Action == cisCreate {
			p.createRecord(ctx, change)
		}
	}

	return nil
}

func groupByNameAndType(records []dnsrecordsv1.DnsrecordDetails) []*endpoint.Endpoint {
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

		endpoints = append(endpoints,
			endpoint.NewEndpointWithTTL(
				*records[0].Name,
				*records[0].Type,
				endpoint.TTL(*records[0].TTL),
				targets...).WithProviderSpecific(source.CisProxiedKey, strconv.FormatBool(*records[0].Proxied)),
		)
	}
	return endpoints
}

func (p *IBMCloudProvider) getRecordID(records []dnsrecordsv1.DnsrecordDetails, record dnsrecordsv1.DnsrecordDetails) string {
	for _, zoneRecord := range records {
		if *zoneRecord.Name == *record.Name && *zoneRecord.Type == *record.Type && *zoneRecord.Content == *record.Content {
			return *zoneRecord.ID
		}
	}
	return ""
}

func (p *IBMCloudProvider) newCISChange(action string, endpoint *endpoint.Endpoint, target string) *cisChange {
	ttl := defaultCISRecordTTL
	proxied := shouldBeProxied(endpoint, p.proxiedByDefault)

	if endpoint.RecordTTL.IsConfigured() {
		ttl = int(endpoint.RecordTTL)
	}

	return &cisChange{
		Action: action,
		ResourceRecord: dnsrecordsv1.DnsrecordDetails{
			Name:    core.StringPtr(endpoint.DNSName),
			TTL:     core.Int64Ptr(int64(ttl)),
			Proxied: core.BoolPtr(proxied),
			Type:    core.StringPtr(endpoint.RecordType),
			Content: core.StringPtr(target),
		},
	}
}

func (p *IBMCloudProvider) createRecord(ctx context.Context, change *cisChange) {
	createDnsRecordOptions := &dnsrecordsv1.CreateDnsRecordOptions{
		Name:    change.ResourceRecord.Name,
		Type:    change.ResourceRecord.Type,
		TTL:     change.ResourceRecord.TTL,
		Content: change.ResourceRecord.Content,
	}
	_, _, err := p.recordsClient.CreateDnsRecordWithContext(ctx, createDnsRecordOptions)
	if err != nil {
		log.Errorf("failed to create %s type record named %s: %v", *change.ResourceRecord.Type, *change.ResourceRecord.Name, err)
	}
}

func (p *IBMCloudProvider) updateRecord(ctx context.Context, recordID string, change *cisChange) {
	updateDnsRecordOptions := &dnsrecordsv1.UpdateDnsRecordOptions{
		DnsrecordIdentifier: &recordID,
		Name:                change.ResourceRecord.Name,
		Type:                change.ResourceRecord.Type,
		TTL:                 change.ResourceRecord.TTL,
		Content:             change.ResourceRecord.Content,
		Proxied:             change.ResourceRecord.Proxied,
	}
	_, _, err := p.recordsClient.UpdateDnsRecordWithContext(ctx, updateDnsRecordOptions)
	if err != nil {
		log.Errorf("failed to update %s type record named %s: %v", *change.ResourceRecord.Type, *change.ResourceRecord.Name, err)
	}
}

func (p *IBMCloudProvider) deleteRecord(ctx context.Context, recordID string) {
	deleteDnsRecordOptions := &dnsrecordsv1.DeleteDnsRecordOptions{
		DnsrecordIdentifier: &recordID,
	}
	_, _, err := p.recordsClient.DeleteDnsRecordWithContext(ctx, deleteDnsRecordOptions)
	if err != nil {
		log.Errorf("failed to delete record %s: %v", recordID, err)
	}
}

func shouldBeProxied(endpoint *endpoint.Endpoint, proxiedByDefault bool) bool {
	proxied := proxiedByDefault

	for _, v := range endpoint.ProviderSpecific {
		if v.Name == source.CloudflareProxiedKey {
			b, err := strconv.ParseBool(v.Value)
			if err != nil {
				log.Errorf("Failed to parse annotation [%s]: %v", source.CloudflareProxiedKey, err)
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
