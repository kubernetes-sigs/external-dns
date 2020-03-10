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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	c "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type akamaiClient interface {
	NewRequest(config edgegrid.Config, method, path string, body io.Reader) (*http.Request, error)
	Do(config edgegrid.Config, req *http.Request) (*http.Response, error)
}

type akamaiOpenClient struct{}

func (*akamaiOpenClient) NewRequest(config edgegrid.Config, method, path string, body io.Reader) (*http.Request, error) {
	return c.NewRequest(config, method, path, body)
}

func (*akamaiOpenClient) Do(config edgegrid.Config, req *http.Request) (*http.Response, error) {
	return c.Do(config, req)
}

// AkamaiConfig clarifies the method signature
type AkamaiConfig struct {
	DomainFilter          endpoint.DomainFilter
	ZoneIDFilter          ZoneIDFilter
	ServiceConsumerDomain string
	ClientToken           string
	ClientSecret          string
	AccessToken           string
	DryRun                bool
}

// AkamaiProvider implements the DNS provider for Akamai.
type AkamaiProvider struct {
	domainFilter endpoint.DomainFilter
	zoneIDFilter ZoneIDFilter
	config       edgegrid.Config
	dryRun       bool
	client       akamaiClient
}

type akamaiZones struct {
	Zones []akamaiZone `json:"zones"`
}

type akamaiZone struct {
	ContractID string `json:"contractId"`
	Zone       string `json:"zone"`
}

type akamaiRecordsets struct {
	Recordsets []akamaiRecord `json:"recordsets"`
}

type akamaiRecord struct {
	Name  string        `json:"name"`
	Type  string        `json:"type"`
	TTL   int64         `json:"ttl"`
	Rdata []interface{} `json:"rdata"`
}

// NewAkamaiProvider initializes a new Akamai DNS based Provider.
func NewAkamaiProvider(akamaiConfig AkamaiConfig) *AkamaiProvider {
	edgeGridConfig := edgegrid.Config{
		Host:         akamaiConfig.ServiceConsumerDomain,
		ClientToken:  akamaiConfig.ClientToken,
		ClientSecret: akamaiConfig.ClientSecret,
		AccessToken:  akamaiConfig.AccessToken,
		MaxBody:      1024,
		HeaderToSign: []string{
			"X-External-DNS",
		},
		Debug: false,
	}

	provider := &AkamaiProvider{
		domainFilter: akamaiConfig.DomainFilter,
		zoneIDFilter: akamaiConfig.ZoneIDFilter,
		config:       edgeGridConfig,
		dryRun:       akamaiConfig.DryRun,
		client:       &akamaiOpenClient{},
	}
	return provider
}

func (p *AkamaiProvider) request(method, path string, body io.Reader) (*http.Response, error) {
	req, err := p.client.NewRequest(p.config, method, fmt.Sprintf("https://%s/%s", p.config.Host, path), body)
	if err != nil {
		log.Errorf("Akamai client failed to prepare the request")
		return nil, err
	}
	resp, err := p.client.Do(p.config, req)

	if err != nil {
		log.Errorf("Akamai client failed to do the request")
		return nil, err
	}
	if !c.IsSuccess(resp) {
		return nil, c.NewAPIError(resp)
	}

	return resp, err
}

//Look here for endpoint documentation -> https://developer.akamai.com/api/web_performance/fast_dns_zone_management/v2.html#getzones
func (p *AkamaiProvider) fetchZones() (zones akamaiZones, err error) {
	log.Debugf("Trying to fetch zones from Akamai")
	resp, err := p.request("GET", "config-dns/v2/zones?showAll=true&types=primary%2Csecondary", nil)
	if err != nil {
		log.Errorf("Failed to fetch zones from Akamai")
		return zones, err
	}

	err = json.NewDecoder(resp.Body).Decode(&zones)
	if err != nil {
		log.Errorf("Could not decode json response from Akamai on zone request")
		return zones, err
	}
	defer resp.Body.Close()

	filteredZones := akamaiZones{}
	for _, zone := range zones.Zones {
		if !p.zoneIDFilter.Match(zone.ContractID) {
			log.Debugf("Skipping zone: '%s' with ZoneID: '%s', it does not match against ZoneID filters", zone.Zone, zone.ContractID)
			continue
		}
		filteredZones.Zones = append(filteredZones.Zones, akamaiZone{ContractID: zone.ContractID, Zone: zone.Zone})
		log.Debugf("Fetched zone: '%s' (ZoneID: %s)", zone.Zone, zone.ContractID)
	}
	lenFilteredZones := len(filteredZones.Zones)
	if lenFilteredZones == 0 {
		log.Warnf("No zones could be fetched")
	} else {
		log.Debugf("Fetched '%d' zones from Akamai", lenFilteredZones)
	}

	return filteredZones, nil
}

//Look here for endpoint documentation -> https://developer.akamai.com/api/web_performance/fast_dns_zone_management/v2.html#getzonerecordsets
func (p *AkamaiProvider) fetchRecordSet(zone string) (recordSet akamaiRecordsets, err error) {
	log.Debugf("Trying to fetch endpoints for zone: '%s' from Akamai", zone)
	resp, err := p.request("GET", "config-dns/v2/zones/"+zone+"/recordsets?showAll=true&types=A%2CTXT%2CCNAME", nil)
	if err != nil {
		log.Errorf("Failed to fetch records from Akamai for zone: '%s'", zone)
		return recordSet, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&recordSet)
	if err != nil {
		log.Errorf("Could not decode json response from Akamai for zone: '%s' on request", zone)
		return recordSet, err
	}

	return recordSet, nil
}

//Records returns the list of records in a given zone.
func (p *AkamaiProvider) Records(context.Context) (endpoints []*endpoint.Endpoint, err error) {
	zones, err := p.fetchZones()
	if err != nil {
		log.Warnf("No zones to fetch endpoints from!")
		return endpoints, err
	}
	for _, zone := range zones.Zones {
		records, err := p.fetchRecordSet(zone.Zone)
		if err != nil {
			log.Warnf("No recordsets could be fetched for zone: '%s'!", zone.Zone)
			continue
		}

		for _, record := range records.Recordsets {
			rdata := make([]string, len(record.Rdata))

			for i, v := range record.Rdata {
				rdata[i] = v.(string)
			}

			if !p.domainFilter.Match(record.Name) {
				log.Debugf("Skipping endpoint DNSName: '%s' RecordType: '%s', it does not match against Domain filters", record.Name, record.Type)
				continue
			}

			endpoints = append(endpoints, endpoint.NewEndpoint(record.Name, record.Type, rdata...))
			log.Debugf("Fetched endpoint DNSName: '%s' RecordType: '%s' Rdata: '%s')", record.Name, record.Type, rdata)
		}
	}
	lenEndpoints := len(endpoints)
	if lenEndpoints == 0 {
		log.Warnf("No endpoints could be fetched")
	} else {
		log.Debugf("Fetched '%d' endpoints from Akamai", lenEndpoints)
	}

	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *AkamaiProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	zoneNameIDMapper := zoneIDName{}
	zones, err := p.fetchZones()
	if err != nil {
		log.Warnf("No zones to fetch endpoints from!")
		return nil
	}

	for _, z := range zones.Zones {
		zoneNameIDMapper[z.Zone] = z.Zone
	}

	_, cf := p.createRecords(zoneNameIDMapper, changes.Create)
	if !p.dryRun {
		if len(cf) > 0 {
			log.Warnf("Not all desired endpoints could be created, retrying next iteration")
			for _, f := range cf {
				log.Warnf("Not created was DNSName: '%s' RecordType: '%s'", f.DNSName, f.RecordType)
			}
		}
	}

	_, df := p.deleteRecords(zoneNameIDMapper, changes.Delete)
	if !p.dryRun {
		if len(df) > 0 {
			log.Warnf("Not all endpoints that require deletion could be deleted, retrying next iteration")
			for _, f := range df {
				log.Warnf("Not deleted was DNSName: '%s' RecordType: '%s'", f.DNSName, f.RecordType)
			}
		}
	}

	_, uf := p.updateNewRecords(zoneNameIDMapper, changes.UpdateNew)
	if !p.dryRun {
		if len(uf) > 0 {
			log.Warnf("Not all endpoints that require updating could be updated, retrying next iteration")
			for _, f := range uf {
				log.Warnf("Not updated was DNSName: '%s' RecordType: '%s'", f.DNSName, f.RecordType)
			}
		}
	}

	for _, uold := range changes.UpdateOld {
		if !p.dryRun {
			log.Debugf("UpdateOld (ignored) for DNSName: '%s' RecordType: '%s'", uold.DNSName, uold.RecordType)
		}
	}

	return nil
}

func (p *AkamaiProvider) newAkamaiRecord(dnsName, recordType string, targets ...string) *akamaiRecord {
	cleanTargets := make([]interface{}, len(targets))
	for idx, target := range targets {
		cleanTargets[idx] = strings.TrimSuffix(target, ".")
	}
	return &akamaiRecord{
		Name:  strings.TrimSuffix(dnsName, "."),
		Rdata: cleanTargets,
		Type:  recordType,
		TTL:   300,
	}
}

func (p *AkamaiProvider) createRecords(zoneNameIDMapper zoneIDName, endpoints []*endpoint.Endpoint) (created []*endpoint.Endpoint, failed []*endpoint.Endpoint) {
	for _, endpoint := range endpoints {

		if !p.domainFilter.Match(endpoint.DNSName) {
			log.Debugf("Skipping creation at Akamai of endpoint DNSName: '%s' RecordType: '%s', it does not match against Domain filters", endpoint.DNSName, endpoint.RecordType)
			continue
		}
		if zoneName, _ := zoneNameIDMapper.FindZone(endpoint.DNSName); zoneName != "" {
			akamaiRecord := p.newAkamaiRecord(endpoint.DNSName, endpoint.RecordType, endpoint.Targets...)
			body, _ := json.MarshalIndent(akamaiRecord, "", "  ")

			log.Infof("Create new Endpoint at Akamai FastDNS - Zone: '%s', DNSName: '%s', RecordType: '%s', Targets: '%+v'", zoneName, endpoint.DNSName, endpoint.RecordType, endpoint.Targets)

			if p.dryRun {
				continue
			}
			_, err := p.request("POST", "config-dns/v2/zones/"+zoneName+"/names/"+endpoint.DNSName+"/types/"+endpoint.RecordType, bytes.NewReader(body))
			if err != nil {
				log.Errorf("Failed to create Akamai endpoint DNSName: '%s' RecordType: '%s' for zone: '%s'", endpoint.DNSName, endpoint.RecordType, zoneName)
				failed = append(failed, endpoint)
				continue
			}
			created = append(created, endpoint)
		} else {
			log.Warnf("No matching zone for endpoint addition DNSName: '%s' RecordType: '%s'", endpoint.DNSName, endpoint.RecordType)
			failed = append(failed, endpoint)
		}
	}
	return created, failed
}

func (p *AkamaiProvider) deleteRecords(zoneNameIDMapper zoneIDName, endpoints []*endpoint.Endpoint) (deleted []*endpoint.Endpoint, failed []*endpoint.Endpoint) {
	for _, endpoint := range endpoints {

		if !p.domainFilter.Match(endpoint.DNSName) {
			log.Debugf("Skipping deletion at Akamai of endpoint: '%s' type: '%s', it does not match against Domain filters", endpoint.DNSName, endpoint.RecordType)
			continue
		}
		if zoneName, _ := zoneNameIDMapper.FindZone(endpoint.DNSName); zoneName != "" {
			log.Infof("Deletion at Akamai FastDNS - Zone: '%s', DNSName: '%s', RecordType: '%s', Targets: '%+v'", zoneName, endpoint.DNSName, endpoint.RecordType, endpoint.Targets)

			if p.dryRun {
				continue
			}

			_, err := p.request("DELETE", "config-dns/v2/zones/"+zoneName+"/names/"+endpoint.DNSName+"/types/"+endpoint.RecordType, nil)
			if err != nil {
				log.Errorf("Failed to delete Akamai endpoint DNSName: '%s' for zone: '%s'", endpoint.DNSName, zoneName)
				failed = append(failed, endpoint)
				continue
			}
			deleted = append(deleted, endpoint)
		} else {
			log.Warnf("No matching zone for endpoint deletion DNSName: '%s' RecordType: '%s'", endpoint.DNSName, endpoint.RecordType)
			failed = append(failed, endpoint)
		}
	}
	return deleted, failed
}

func (p *AkamaiProvider) updateNewRecords(zoneNameIDMapper zoneIDName, endpoints []*endpoint.Endpoint) (updated []*endpoint.Endpoint, failed []*endpoint.Endpoint) {
	for _, endpoint := range endpoints {

		if !p.domainFilter.Match(endpoint.DNSName) {
			log.Debugf("Skipping update at Akamai of endpoint DNSName: '%s' RecordType: '%s', it does not match against Domain filters", endpoint.DNSName, endpoint.RecordType)
			continue
		}
		if zoneName, _ := zoneNameIDMapper.FindZone(endpoint.DNSName); zoneName != "" {
			akamaiRecord := p.newAkamaiRecord(endpoint.DNSName, endpoint.RecordType, endpoint.Targets...)
			body, _ := json.MarshalIndent(akamaiRecord, "", "  ")

			log.Infof("Updating endpoint at Akamai FastDNS - Zone: '%s', DNSName: '%s', RecordType: '%s', Targets: '%+v'", zoneName, endpoint.DNSName, endpoint.RecordType, endpoint.Targets)

			if p.dryRun {
				continue
			}

			_, err := p.request("PUT", "config-dns/v2/zones/"+zoneName+"/names/"+endpoint.DNSName+"/types/"+endpoint.RecordType, bytes.NewReader(body))
			if err != nil {
				log.Errorf("Failed to update Akamai endpoint DNSName: '%s' for zone: '%s'", endpoint.DNSName, zoneName)
				failed = append(failed, endpoint)
				continue
			}
			updated = append(updated, endpoint)
		} else {
			log.Warnf("No matching zone for endpoint update DNSName: '%s' RecordType: '%s'", endpoint.DNSName, endpoint.RecordType)
			failed = append(failed, endpoint)
		}
	}
	return updated, failed
}
