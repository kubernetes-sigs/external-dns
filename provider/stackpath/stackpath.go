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
	"io"
	"net/http"
	"os"
	"strings"
	"time"

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
	zoneIDFilter provider.ZoneIDFilter
	stackID      string
	dryRun       bool
	testing      bool
}

type StackPathConfig struct {
	Context      context.Context
	DomainFilter endpoint.DomainFilter
	ZoneIDFilter provider.ZoneIDFilter
	DryRun       bool
	Testing      bool
}

func NewStackPathProvider(config StackPathConfig) (*StackPathProvider, error) {

	log.Info("Creating StackPath provider")

	clientID, ok := os.LookupEnv("STACKPATH_CLIENT_ID")
	if !ok {
		return nil, fmt.Errorf("STACKPATH_CLIENT_ID environment variable is not set")
	}

	clientSecret, ok := os.LookupEnv("STACKPATH_CLIENT_SECRET")
	if !ok {
		return nil, fmt.Errorf("STACKPATH_CLIENT_SECRET environment variable is not set")
	}

	stackID, ok := os.LookupEnv("STACKPATH_STACK_ID")
	if !ok {
		return nil, fmt.Errorf("STACKPATH_STACK_ID environment variable is not set")
	}

	oauthSource := oauth2.NewTokenSource(clientID, clientSecret, oauth2.HTTPClientOption(http.DefaultClient))
	_, err := oauthSource.Token()
	if err != nil {
		return nil, err
	} else {
		log.Info("Successfully authenticated with StackPath")
	}

	authorizedContext := context.WithValue(config.Context, dns.ContextOAuth2, oauthSource)

	clientConfig := dns.NewConfiguration()

	client := dns.NewAPIClient(clientConfig)

	provider := &StackPathProvider{
		client:       client,
		context:      authorizedContext,
		domainFilter: config.DomainFilter,
		zoneIDFilter: config.ZoneIDFilter,
		stackID:      stackID,
		dryRun:       config.DryRun,
		testing:      config.Testing,
	}

	return provider, nil
}

//Base Provider Functions

func (p *StackPathProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {

	log.Info("Getting records from StackPath")

	var endpoints []*endpoint.Endpoint

	zones, err := p.zones()
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {

		recordsResponse, _, err := p.getZoneRecords(zone.GetId())
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

	merged := mergeEndpointsByNameType(endpoints)
	out := "Found:"
	for _, e := range merged {
		out = out + " [" + e.DNSName + " " + e.RecordType + " " + string(rune(len(e.Targets))) + "]"
	}
	log.Infof(out)

	return merged, nil
}

func (p *StackPathProvider) StackPathStyleRecords() ([]dns.ZoneZoneRecord, error) {

	log.Info("Getting records from StackPath")

	var records []dns.ZoneZoneRecord

	zones, err := p.zones()
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {

		recordsResponse, _, err := p.getZoneRecords(zone.GetId())
		if err != nil || (p.testing && p.dryRun) {
			return nil, err
		}

		records = append(records, recordsResponse.GetRecords()...)

	}

	out := "Found:"
	for _, e := range records {
		out = out + " [" + e.GetName() + " " + e.GetType() + " " + e.GetData() + " " + fmt.Sprint(e.GetTtl()) + "]"
	}
	log.Infof(out)

	return records, nil
}

func (p *StackPathProvider) getZoneRecords(zoneID string) (dns.ZoneGetZoneRecordsResponse, *http.Response, error) {

	if p.testing && p.dryRun {
		return testGetZoneZoneRecordsResponse, nil, fmt.Errorf("testing")
	} else if p.testing {
		return testGetZoneZoneRecordsResponse, nil, nil
	}

	return p.client.ResourceRecordsApi.GetZoneRecords(p.context, p.stackID, zoneID).Execute()
}

func (p *StackPathProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {

	zs, err := p.zones()
	if err != nil {
		return err
	}
	zones := &zs

	zoneIDNameMap := provider.ZoneIDName{}
	for _, zone := range zs {
		zoneIDNameMap.Add(zone.GetId(), zone.GetDomain())
	}

	records, err := p.StackPathStyleRecords()
	if err != nil {
		return err
	}

	err = p.create(changes.Create, zones, &zoneIDNameMap)
	if err != nil {
		return err
	}

	err = p.delete(changes.Delete, zones, &zoneIDNameMap, &records)
	if err != nil {
		return err
	}

	err = p.update(changes.UpdateOld, changes.UpdateNew, zones, &zoneIDNameMap, &records)
	if err != nil {
		return err
	}

	return nil
}

func (p *StackPathProvider) create(endpoints []*endpoint.Endpoint, zones *[]dns.ZoneZone, zoneIDNameMap *provider.ZoneIDName) error {

	createsByZoneID := endpointsByZoneID(*zoneIDNameMap, endpoints)

	for zoneID, endpoints := range createsByZoneID {
		log.Infof("Creating %d records in zone %s (ID:%s)", len(endpoints), (*zoneIDNameMap)[zoneID], zoneID)
		domain := (*zoneIDNameMap)[zoneID]
		for _, endpoint := range endpoints {
			for _, target := range endpoint.Targets {

				if !p.dryRun {
					err := p.createTarget(zoneID, domain, endpoint, target)
					if err != nil {
						return err
					}
				} else {
					log.Infof("Would have created record: %s %s %s %s", endpoint.DNSName, endpoint.RecordType, target, fmt.Sprint(endpoint.RecordTTL))
				}

			}
		}
	}

	return nil
}

func (p *StackPathProvider) createTarget(zoneID string, domain string, ep *endpoint.Endpoint, target string) error {

	msg := dns.NewZoneUpdateZoneRecordMessage()
	name := strings.TrimSuffix(ep.DNSName, domain)
	name = strings.TrimSuffix(name, ".")
	if name == "" || name == "@" {
		name = "@"
	}

	if !ep.RecordTTL.IsConfigured() {
		ep.RecordTTL = endpoint.TTL(120)
	}

	msg.SetName(name)
	msg.SetType(dns.ZoneRecordType(ep.RecordType))
	msg.SetTtl(int32(ep.RecordTTL))
	msg.SetData(target)

	log.Infof("Creating record " + name + "." + domain + " " + ep.RecordType + " " + target + " " + fmt.Sprint(ep.RecordTTL))

	a, r, err := p.createCall(zoneID, domain, ep, target, msg)

	s := &http.Response{}

	if err != nil {
		log.Infof(err.Error())
		if r == s {
			r.Body.Close()
			b, _ := io.ReadAll(r.Body)
			log.Infof(string(b))
		}
		return err
	}

	b := dns.ZoneCreateZoneRecordResponse{}

	if a != b {
		log.Infof("Created record " + *a.Record.Name + "." + domain + " (ID:" + *a.Record.Id + ")")
	}

	return nil
}

func (p *StackPathProvider) createCall(zoneID string, domain string, endpoint *endpoint.Endpoint, target string, msg *dns.ZoneUpdateZoneRecordMessage) (dns.ZoneCreateZoneRecordResponse, *http.Response, error) {
	if p.testing && p.dryRun {
		return dns.ZoneCreateZoneRecordResponse{}, nil, fmt.Errorf("testing")
	} else if p.testing {
		return dns.ZoneCreateZoneRecordResponse{}, nil, nil
	} else {
		return p.client.ResourceRecordsApi.CreateZoneRecord(p.context, p.stackID, zoneID).ZoneUpdateZoneRecordMessage(*msg).Execute()
	}
}

func (p *StackPathProvider) delete(endpoints []*endpoint.Endpoint, zones *[]dns.ZoneZone, zoneIDNameMap *provider.ZoneIDName, records *[]dns.ZoneZoneRecord) error {
	log.Infof("Deleting %s record(s)", fmt.Sprint(len(endpoints)))

	deleteByZoneID := endpointsByZoneID(*zoneIDNameMap, endpoints)

	for zoneID, endpoints := range deleteByZoneID {
		for _, endpoint := range endpoints {
			for _, target := range endpoint.Targets {
				if !p.dryRun {
					domain := (*zoneIDNameMap)[zoneID]
					recordID, err := recordFromTarget(endpoint, target, records, domain)
					if err != nil {
						return err
					}
					p.deleteTarget(zoneID, recordID)
				} else {
					log.Infof("Would have deleted record: %s %s %s %s", endpoint.DNSName, endpoint.RecordType, target, fmt.Sprint(endpoint.RecordTTL))
				}
			}
		}
	}

	return nil
}

func (p *StackPathProvider) deleteTarget(zone string, record string) error {
	resp, err := p.deleteCall(zone, record)

	s := &http.Response{}

	if err != nil {
		log.Infof(err.Error())
		if resp == s {
			resp.Body.Close()
			b, _ := io.ReadAll(resp.Body)
			log.Infof(string(b))
		}
		return err
	}

	log.Infof("Deleted record " + record)

	return nil
}

func (p *StackPathProvider) deleteCall(zone string, record string) (*http.Response, error) {
	if p.testing && p.dryRun {
		return nil, fmt.Errorf("testing")
	} else if p.testing {
		return nil, nil
	} else {
		return p.client.ResourceRecordsApi.DeleteZoneRecord(p.context, p.stackID, zone, record).Execute()
	}
}

func (p *StackPathProvider) update(old []*endpoint.Endpoint, new []*endpoint.Endpoint, zones *[]dns.ZoneZone, zoneIDNameMap *provider.ZoneIDName, records *[]dns.ZoneZoneRecord) error {

	err := p.create(new, zones, zoneIDNameMap)
	if err != nil {
		return err
	}

	err = p.delete(old, zones, zoneIDNameMap, records)
	if err != nil {
		return err
	}

	return nil
}

func (p *StackPathProvider) zones() ([]dns.ZoneZone, error) {

	zoneResponse, _, err := p.getZones()
	if err != nil {
		return nil, err
	}

	zones := zoneResponse.GetZones()
	filteredZones := []dns.ZoneZone{}

	for _, zone := range zones {
		if p.zoneIDFilter.Match(zone.GetId()) && p.domainFilter.Match(zone.GetDomain()) {
			filteredZones = append(filteredZones, zone)
			log.Debugf("Matched zone " + zone.GetId())
		} else {
			log.Debugf("Filtered zone " + zone.GetId())
		}
	}

	return filteredZones, nil
}

func (p *StackPathProvider) getZones() (dns.ZoneGetZonesResponse, *http.Response, error) {
	if p.testing && !p.dryRun {
		return testGetZoneRecords, nil, nil
	} else if p.testing {
		return testGetZoneRecords, nil, fmt.Errorf("testing")
	} else {
		return p.client.ZonesApi.GetZones(p.context, p.stackID).Execute()
	}
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
		ttl := endpoints[0].RecordTTL

		targets := make([]string, len(endpoints))
		for i, e := range endpoints {
			targets[i] = e.Targets[0]
		}

		e := endpoint.NewEndpoint(dnsName, recordType, targets...)
		e.RecordTTL = ttl
		result = append(result, e)
	}

	return result
}

//From pkg/digitalocean/provider.go
func endpointsByZoneID(zoneNameIDMapper provider.ZoneIDName, endpoints []*endpoint.Endpoint) map[string][]*endpoint.Endpoint {
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

func recordFromTarget(endpoint *endpoint.Endpoint, target string, records *[]dns.ZoneZoneRecord, domain string) (string, error) {

	var name string

	if endpoint.DNSName == "" {
		name = "@"
	} else {
		name = strings.TrimSuffix(endpoint.DNSName, "."+domain)
	}

	for _, record := range *records {
		if record.GetName() == name && record.GetType() == endpoint.RecordType && record.GetData() == strings.Trim(target, "\\\"") {
			return *record.Id, nil
		}
	}

	return "", fmt.Errorf("record not found")
}

var (
	testStackID     = "TEST_STACK_ID"
	testAccountID   = "TEST_ACCOUNT_ID"
	testNameservers = []string{"ns1.example.com", "ns2.example.com"}

	testZoneID                           = []string{"TEST_ZONE_ID1", ""}
	testZoneDomain                       = []string{"one.com", "two.com"}
	testZoneVersion                      = []string{"TEST_ZONE_VERSION1", ""}
	testZoneLabels                       = make(map[string]string)
	testZoneDisabled                     = []bool{false, false}
	testZoneStatus    dns.ZoneZoneStatus = "ACTIVE"
	testGetZonesZones                    = []dns.ZoneZone{
		{
			StackId:     &testStackID,
			AccountId:   &testAccountID,
			Id:          &testZoneID[0],
			Domain:      &testZoneDomain[0],
			Version:     &testZoneVersion[0],
			Labels:      &testZoneLabels,
			Created:     &time.Time{},
			Updated:     &time.Time{},
			Nameservers: &testNameservers,
			Verified:    &time.Time{},
			Status:      &testZoneStatus,
			Disabled:    &testZoneDisabled[0],
		},
	}

	badZone = dns.ZoneZone{
		StackId:     &testStackID,
		AccountId:   &testAccountID,
		Id:          &testZoneID[1],
		Domain:      &testZoneDomain[1],
		Version:     &testZoneVersion[1],
		Labels:      &testZoneLabels,
		Created:     &time.Time{},
		Updated:     &time.Time{},
		Nameservers: &testNameservers,
		Verified:    &time.Time{},
		Status:      &testZoneStatus,
		Disabled:    &testZoneDisabled[1],
	}

	testGetZonesTotalCount      = "1"
	testGetZonesHasPreviousPage = false
	testGetZonesHasNextPage     = false
	testGetZonesEndCursor       = "0"
	testGetZoneRecords          = dns.ZoneGetZonesResponse{
		PageInfo: &dns.PaginationPageInfo{
			TotalCount:      &testGetZonesTotalCount,
			HasPreviousPage: &testGetZonesHasPreviousPage,
			HasNextPage:     &testGetZonesHasNextPage,
			EndCursor:       &testGetZonesEndCursor,
		},
		Zones: &testGetZonesZones,
	}

	testZoneZoneRecordID     = []string{"TEST_ZONE_ZONE_RECORD_ID1", "TEST_ZONE_ZONE_RECORD_ID2", "TEST_ZONE_ZONE_RECORD_ID3"}
	testZoneZoneRecordZoneID = []string{"TEST_ZONE_ID1", "TEST_ZONE_ID1", "TEST_ZONE_ID1"}
	testZoneZoneRecordName   = []string{"www", "www", "@"}
	testZoneZoneRecordType   = []string{"A", "A", "CNAME"}
	testZoneZoneRecordClass  = []string{"TEST_ZONE_ZONE_RECORD_CLASS1", "TEST_ZONE_ZONE_RECORD_CLASS2", "TEST_ZONE_ZONE_RECORD_CLASS3"}
	testZoneZoneRecordTTL    = []int32{int32(60), int32(60), int32(180)}
	testZoneZoneRecordData   = []string{"1.1.1.1", "2.2.2.2", "testing.com"}
	testZoneZoneRecordWeight = []int32{int32(1), int32(2), int32(3)}
	testZoneZoneRecordLabels = make(map[string]string)

	testGetZoneZoneRecords = []dns.ZoneZoneRecord{
		{
			Id:      &testZoneZoneRecordID[0],
			ZoneId:  &testZoneZoneRecordZoneID[0],
			Name:    &testZoneZoneRecordName[0],
			Type:    &testZoneZoneRecordType[0],
			Class:   &testZoneZoneRecordClass[0],
			Ttl:     &testZoneZoneRecordTTL[0],
			Data:    &testZoneZoneRecordData[0],
			Weight:  &testZoneZoneRecordWeight[0],
			Labels:  &testZoneZoneRecordLabels,
			Created: &time.Time{},
			Updated: &time.Time{},
		},
		{
			Id:      &testZoneZoneRecordID[1],
			ZoneId:  &testZoneZoneRecordZoneID[1],
			Name:    &testZoneZoneRecordName[1],
			Type:    &testZoneZoneRecordType[1],
			Class:   &testZoneZoneRecordClass[1],
			Ttl:     &testZoneZoneRecordTTL[1],
			Data:    &testZoneZoneRecordData[1],
			Weight:  &testZoneZoneRecordWeight[1],
			Labels:  &testZoneZoneRecordLabels,
			Created: &time.Time{},
			Updated: &time.Time{},
		},
		{
			Id:      &testZoneZoneRecordID[2],
			ZoneId:  &testZoneZoneRecordZoneID[2],
			Name:    &testZoneZoneRecordName[2],
			Type:    &testZoneZoneRecordType[2],
			Class:   &testZoneZoneRecordClass[2],
			Ttl:     &testZoneZoneRecordTTL[2],
			Data:    &testZoneZoneRecordData[2],
			Weight:  &testZoneZoneRecordWeight[2],
			Labels:  &testZoneZoneRecordLabels,
			Created: &time.Time{},
			Updated: &time.Time{},
		},
	}

	testGetZoneZoneRecordsTotalCount      = "3"
	testGetZoneZoneRecordsHasPreviousPage = false
	testGetZoneZoneRecordsHasNextPage     = false
	testGetZoneZoneRecordsEndCursor       = "2"
	testGetZoneZoneRecordsResponse        = dns.ZoneGetZoneRecordsResponse{
		PageInfo: &dns.PaginationPageInfo{
			TotalCount:      &testGetZoneZoneRecordsTotalCount,
			HasPreviousPage: &testGetZoneZoneRecordsHasPreviousPage,
			HasNextPage:     &testGetZoneZoneRecordsHasNextPage,
			EndCursor:       &testGetZoneZoneRecordsEndCursor,
		},
		Records: &testGetZoneZoneRecords,
	}

	allTestEndpoints = []*endpoint.Endpoint{
		{
			DNSName:          "www.one.com",
			Targets:          endpoint.Targets{"1.1.1.1"},
			RecordType:       endpoint.RecordTypeA,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(60),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
		{
			DNSName:          "www.one.com",
			Targets:          endpoint.Targets{"2.2.2.2"},
			RecordType:       endpoint.RecordTypeA,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(60),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
		{
			DNSName:          "@.two.com",
			Targets:          endpoint.Targets{"testing.com"},
			RecordType:       endpoint.RecordTypeCNAME,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(180),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
		{
			DNSName:          "three.one.com",
			Targets:          endpoint.Targets{"3.3.3.3"},
			RecordType:       endpoint.RecordTypeA,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(180),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
		{
			DNSName:          "four.one.com",
			Targets:          endpoint.Targets{"4.4.4.4"},
			RecordType:       endpoint.RecordTypeA,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(180),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
		{
			DNSName:          "five.one.com",
			Targets:          endpoint.Targets{"5.5.5.5"},
			RecordType:       endpoint.RecordTypeA,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(180),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
		{
			DNSName:          "six.one.com",
			Targets:          endpoint.Targets{"6.6.6.6"},
			RecordType:       endpoint.RecordTypeA,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(180),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
		{
			DNSName:          "seven.one.com",
			Targets:          endpoint.Targets{"7.7.7.7"},
			RecordType:       endpoint.RecordTypeA,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(180),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
		{
			DNSName:          "eight.one.com",
			Targets:          endpoint.Targets{"8.8.8.8"},
			RecordType:       endpoint.RecordTypeA,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(180),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
	}

	threeTestEndpoints = allTestEndpoints[0:3]

	testMergedEndpoints = []*endpoint.Endpoint{
		allTestEndpoints[2],
		{
			DNSName:          "www.one.com",
			Targets:          endpoint.Targets{"1.1.1.1", "2.2.2.2"},
			RecordType:       endpoint.RecordTypeA,
			SetIdentifier:    "",
			RecordTTL:        endpoint.TTL(60),
			Labels:           testZoneZoneRecordLabels,
			ProviderSpecific: nil,
		},
	}

	testChanges = &plan.Changes{
		Create:    allTestEndpoints[3:5],
		UpdateOld: allTestEndpoints[5:7],
		UpdateNew: allTestEndpoints[7:9],
		Delete:    allTestEndpoints[0:3],
	}
)
