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

package pv6connect

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type ProVisionConfig struct {
	ProVisionHost     string   `json:"provisionHost,omitempty"`
	ProVisionUsername string   `json:"provisionUsername,omitempty"`
	ProVisionPassword string   `json:"provisionPassword,omitempty"`
	ZoneIDs           []string `json:"zoneIDs,omitempty"`
	ProVisionPush     bool     `json:"provisionPush"`
	GetAllRecords     bool     `json:"getAllRecords"`
	SkipTLSVerify     bool     `json:"skipTLSVerify"`
}

// ProVisionProvider implements the DNS provider for 6connect ProVision
type ProVisionProvider struct {
	provider.BaseProvider
	domainFilter  endpoint.DomainFilter
	zoneIDFilter  provider.ZoneIDFilter
	dryRun        bool
	ZoneIDs       []string
	ProVisionPush bool
	GetAllRecords bool
	PVClient      PVClient
}

// PVClientConfig defines new client on ProVision
type PVClientConfig struct {
	Host          string
	Username      string
	Password      string
	ZoneIDs       []string
	SkipTLSVerify bool
}

type PVClient interface {
	getProVisionZones(zoneIDs []string) ([]PVResource, error)
	getProVisionRecords(zoneID string) ([]PVResource, error)
	buildHTTPRequest(method, url string, body io.Reader) (*http.Request, error)
	getProVisionSpecificRecord(ZoneID, RecordHost, RecordType, RecordValue string, OutputRes *PVResource) (bool, error)
	flagProVisionRecord(recordID string) (bool, error)
	createProVisionRecord(zoneID string, ep *endpoint.Endpoint) (bool, error)
	deleteProVisionRecord(recordID string) error
	pushProVisionZone(zoneID string) error
}

// PVResource defines a resource for both zones and records
type PVResource struct {
	ID    string                 `json:"id"`
	Name  string                 `json:"name"`
	Attrs map[string]interface{} `json:"attrs"`
}

func IsValidRecordType(recordType string) bool {
	switch recordType {
	case
		"A",
		"TXT",
		"CNAME":
		return true
	}
	return false
}

func SplitZoneIds(zoneIds string) []string {
	//TODO check if the zone id elements are integers
	var zone_ids_slice []string
	if zoneIds == "" {
		return zone_ids_slice
	}

	return strings.Split(zoneIds, ",")
}

func checkIfCfgApplied(cfg ProVisionConfig) bool {
	if cfg.ProVisionUsername == "" ||
		cfg.ProVisionHost == "" ||
		cfg.ProVisionPassword == "" ||
		len(cfg.ZoneIDs) == 0 {
		return false
	}

	return true
}

func fillCfgWithEnv(cfg *ProVisionConfig) error {

	if cfg.ProVisionUsername == "" {
		if v, ok := os.LookupEnv("PROVISION_USERNAME"); ok {
			cfg.ProVisionUsername = v
		} else {
			return errors.Errorf("JSON ProVisionUsername and PROVISION_USERNAME are empty")
		}
	}

	if cfg.ProVisionPassword == "" {
		if v, ok := os.LookupEnv("PROVISION_PASSWORD"); ok {
			cfg.ProVisionPassword = v
		} else {
			return errors.Errorf("JSON ProVisionPassword and PROVISION_PASSWORD are empty")
		}
	}

	if cfg.ProVisionHost == "" {
		if v, ok := os.LookupEnv("PROVISION_HOST"); ok {
			cfg.ProVisionHost = v
		} else {
			return errors.Errorf("JSON ProVisionHost and PROVISION_HOST are empty")
		}
	}

	if len(cfg.ZoneIDs) == 0 {
		if v, ok := os.LookupEnv("PROVISION_ZONEIDS"); ok {
			cfg.ZoneIDs = SplitZoneIds(v)
		} else {
			return errors.Errorf("JSON ZoneIDs and PROVISION_ZONEIDS are empty")
		}
	}

	if !cfg.GetAllRecords {
		if _, ok := os.LookupEnv("PROVISION_PUSH"); ok {
			cfg.ProVisionPush = true
		}
	}

	if !cfg.GetAllRecords {
		if _, ok := os.LookupEnv("PROVISION_GETALLRECORDS"); ok {
			cfg.GetAllRecords = true
		}
	}

	if !cfg.SkipTLSVerify {
		if _, ok := os.LookupEnv("PROVISION_SKIPTLSVERIFY"); ok {
			cfg.SkipTLSVerify = true
		}
	}

	return nil
}

// NewProVisionProvider creates a new 6connect ProVision provider.
//
// Returns a pointer to the provider or an error if a provider could not be created.
func NewProVisionProvider(cfg ProVisionConfig, configFile string, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool) (*ProVisionProvider, error) {
	if configFile != "" && !checkIfCfgApplied(cfg) {
		contents, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read ProVision config file %v", configFile)
		}

		err = json.Unmarshal(contents, &cfg)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read ProVision config file %v", configFile)
		}
	}

	err := fillCfgWithEnv(&cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "Missing ProVision Config Values : %v", err)
	}

	//log.Debugf("CONFIG : %v", cfg)

	pvClient := NewPVClient(cfg.ProVisionHost, cfg.ProVisionUsername, cfg.ProVisionPassword, cfg.ZoneIDs, cfg.SkipTLSVerify)

	provider := &ProVisionProvider{
		domainFilter:  domainFilter,
		zoneIDFilter:  zoneIDFilter,
		dryRun:        dryRun,
		ZoneIDs:       cfg.ZoneIDs,
		ProVisionPush: cfg.ProVisionPush,
		GetAllRecords: cfg.GetAllRecords,
		PVClient:      pvClient,
	}
	return provider, nil
}

// NewPVClient creates and returns a new ProVision client
func NewPVClient(hostname, username, password string, ZoneIDs []string, skipTLSVerify bool) PVClientConfig {

	return PVClientConfig{
		Host:          hostname,
		Username:      username,
		Password:      password,
		ZoneIDs:       ZoneIDs,
		SkipTLSVerify: skipTLSVerify,
	}
}

// Records fetches the Zone Records for ProVision
func (p *ProVisionProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, err error) {
	processed_zones := make(map[string]bool)
	zones, err := p.zones()
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch zones")
	}

	for _, zone := range zones {
		//as we may have more than 1 zone with the same name (views) we want to provide the records only from the first one
		if _, ok := processed_zones[zone.Name]; ok {
			continue
		}

		txt_record_owners := make(map[string]string)

		log.Debugf("fetching records from zone '%s'", zone.Name)
		processed_zones[zone.Name] = true

		var records []PVResource
		records, err = p.PVClient.getProVisionRecords(zone.ID)
		if err != nil {
			return nil, errors.Wrapf(err, "could not fetch  records for zone: %v", zone.Name)
		}

		//lets firstly process only the TXT records because of the owner
		for _, record := range records {
			var ep *endpoint.Endpoint
			record_type := record.Attrs["record_type"].(string)
			if record_type != "TXT" {
				continue
			}

			if !p.GetAllRecords {
				if _, ok := record.Attrs["k8s_externaldns"]; ok {
					k8s_externaldns := record.Attrs["k8s_externaldns"].(string)
					if k8s_externaldns == "0" {
						continue
					}

				} else {
					continue
				}
			}

			record_host := record.Attrs["record_host"].(string)
			record_value := record.Attrs["record_value"].(string)

			if _, ok := record.Attrs["record_ttl"]; ok {
				record_ttl_int := 0
				switch v := record.Attrs["record_ttl"].(type) {
				case int:
					record_ttl_int = v
				case string:
					tmpint, err := strconv.Atoi(v)
					if err != nil {
						return nil, errors.Wrapf(err, "could not parse ttl '%v' as int for host record %v", v, record_host)
					}
					record_ttl_int = tmpint
				}
				ep = endpoint.NewEndpointWithTTL(record_host, record_type, endpoint.TTL(record_ttl_int), record_value) //.WithProviderSpecific("RecordID", record.ID)
			} else {
				ep = endpoint.NewEndpoint(record_host, record_type, record_value) //.WithProviderSpecific("RecordID", record.ID)
			}

			txt_owner, err := extractOwnerfromTXTRecord(record_value)
			if err != nil {
				log.Debugf("External DNS Owner %s", err)
				ep.Labels[endpoint.OwnerLabelKey] = "default"
			} else {
				ep.Labels[endpoint.OwnerLabelKey] = txt_owner
				txt_record_owners[record_host] = txt_owner
			}

			endpoints = append(endpoints, ep)
		}

		for _, record := range records {
			var ep *endpoint.Endpoint
			record_type := record.Attrs["record_type"].(string)
			record_host := record.Attrs["record_host"].(string)
			record_value := record.Attrs["record_value"].(string)

			//lets skip the TXTs as they have been already processed
			if record_type == "TXT" {
				continue
			}

			if !p.GetAllRecords {
				if _, ok := record.Attrs["k8s_externaldns"]; ok {
					k8s_externaldns := record.Attrs["k8s_externaldns"].(string)
					if k8s_externaldns == "0" {
						continue
					}

				} else {
					continue
				}
			}

			if !IsValidRecordType(record_type) {
				continue
			}

			if err != nil {
				log.Debugf("External DNS Owner %s", err)
			}

			if _, ok := record.Attrs["record_ttl"]; ok {
				record_ttl_int := 0
				switch v := record.Attrs["record_ttl"].(type) {
				case int:
					record_ttl_int = v
				case string:
					tmpint, err := strconv.Atoi(v)
					if err != nil {
						return nil, errors.Wrapf(err, "could not parse ttl '%v' as int for host record %v", v, record_host)
					}
					record_ttl_int = tmpint
				}

				ep = endpoint.NewEndpointWithTTL(record_host, record_type, endpoint.TTL(record_ttl_int), record_value) //.WithProviderSpecific("RecordID", record.ID)
			} else {
				ep = endpoint.NewEndpoint(record_host, record_type, record_value) //.WithProviderSpecific("RecordID", record.ID)
			}

			if txt_owner, ok := txt_record_owners[record_host]; ok {
				ep.Labels[endpoint.OwnerLabelKey] = txt_owner
			} else {
				ep.Labels[endpoint.OwnerLabelKey] = "default"
			}

			endpoints = append(endpoints, ep)
		}

	}

	log.Debugf("fetched %d records from ProVision", len(endpoints))
	return endpoints, nil
}

// ApplyChanges updates necessary zones and replaces old records with new ones
//
// Returns nil upon success and err is there is an error
func (p *ProVisionProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	zones, err := p.zones()
	if err != nil {
		return err
	}
	log.Infof("zones is: %+v\n", zones)
	log.Infof("changes: %+v\n", changes)
	created, deleted := p.mapChanges(zones, changes)
	log.Infof("created: %+v\n", created)
	log.Infof("deleted: %+v\n", deleted)

	p.deleteRecords(deleted)
	p.createRecords(created)

	if p.ProVisionPush {
		//lets get hash of the zone ids for both created and deleted
		processed_zones := make(map[string]bool)
		for zoneID, _ := range deleted {
			processed_zones[zoneID] = true
		}
		for zoneID, _ := range created {
			processed_zones[zoneID] = true
		}

		//do the actual push
		for zoneID, _ := range processed_zones {
			log.Infof("Pushing zone ID '%s'.",
				zoneID,
			)
			if p.dryRun {
				continue
			}
			p.PVClient.pushProVisionZone(zoneID)
		}
	}

	return nil
}

type pvChangeMap map[string]map[string]*endpoint.Endpoint

func (p *ProVisionProvider) mapChanges(zones []PVResource, changes *plan.Changes) (pvChangeMap, pvChangeMap) {
	created := pvChangeMap{}
	deleted := pvChangeMap{}

	mapChange := func(created pvChangeMap, deleted pvChangeMap, is_delete bool, change *endpoint.Endpoint) {
		md5data := []byte(change.DNSName + change.RecordType + change.Targets[0])
		md5sum := md5.Sum(md5data)
		md5hash := hex.EncodeToString(md5sum[:])

		zoneIDs := p.findZone(zones, change.DNSName)
		if len(zoneIDs) == 0 {
			log.Debugf("ignoring changes to '%s' because a suitable ProVision DNS zone was not found", change.DNSName)
			return
		}

		for _, zoneID := range zoneIDs {
			if is_delete {
				if _, ok := created[zoneID]; ok {
					if _, ok := created[zoneID][md5hash]; ok {
						delete(created[zoneID], md5hash)
						return
					}
				}

				if _, ok := deleted[zoneID]; !ok {
					deleted[zoneID] = make(map[string]*endpoint.Endpoint)
				}

				deleted[zoneID][md5hash] = change
			} else {
				if _, ok := created[zoneID]; !ok {
					created[zoneID] = make(map[string]*endpoint.Endpoint)
				}

				created[zoneID][md5hash] = change
			}
		}
	}

	for _, change := range changes.Create {
		/*
			log.Infof("Name: %+v\n", change.DNSName)
			log.Infof("Targets: %+v\n", change.Targets)
			log.Infof("ProviderSpecific: %+v\n", change.ProviderSpecific)
			log.Infof("RecordType: %+v\n", change.RecordType)
			log.Infof("Labels: %+v\n", change.Labels)
		*/
		mapChange(created, deleted, false, change)
	}
	for _, change := range changes.UpdateNew {
		mapChange(created, deleted, false, change)
	}
	for _, change := range changes.Delete {
		mapChange(created, deleted, true, change)
	}
	for _, change := range changes.UpdateOld {
		mapChange(created, deleted, true, change)
	}

	return created, deleted
}

// findZone finds the most specific matching zone for a given record 'name' from a list of all zones
func (p *ProVisionProvider) findZone(zones []PVResource, name string) []string {
	var result []string

	if !strings.HasSuffix(name, ".") {
		name += "."
	}

	for _, zone := range zones {
		if strings.EqualFold(name, zone.Name) {
			result = append(result, zone.ID)
		} else if strings.HasSuffix(name, "."+zone.Name) {
			result = append(result, zone.ID)
		}
	}

	return result
}

func (p *ProVisionProvider) zones() ([]PVResource, error) {
	log.Debugf("retrieving ProVision zones")
	//var zones []PVResource

	zonelist, err := p.PVClient.getProVisionZones(p.ZoneIDs)
	if err != nil {
		return nil, err
	}

	log.Debugf("found %d zones", len(zonelist))
	return zonelist, nil

	/*
		for _, zone := range zonelist {

				if !p.domainFilter.Match(zone.Name) {
					continue
				}

				if !p.zoneIDFilter.Match(zone.ID) {
					continue
				}

			zones = append(zones, zone)
		}
	*/
}

func (p *ProVisionProvider) createRecords(created pvChangeMap) {
	for zoneID, endpoints := range created {
		for _, ep := range endpoints {
			log.Infof("creating %s record named '%s' to '%s' for ProVision DNS zone ID '%s'.",
				ep.RecordType,
				ep.DNSName,
				ep.Targets,
				zoneID,
			)
			if p.dryRun {
				continue
			}

			//lets check if the record already exists
			var record_res PVResource
			is_exist, err := p.PVClient.getProVisionSpecificRecord(zoneID, ep.DNSName, ep.RecordType, ep.Targets[0], &record_res)

			if err != nil {
				log.Errorf(
					"Failed to retrieve %s record named '%s' to '%s' for DNS zone id '%s': %v",
					ep.RecordType,
					ep.DNSName,
					ep.Targets,
					zoneID,
					err,
				)
				continue
			}
			//the record already exists so nothing to create
			if is_exist {
				log.Debugf("Record record already exist so nothing to create")
				if !p.GetAllRecords {
					log.Debugf("Record record already exist, updating the externaldns flag")
					p.PVClient.flagProVisionRecord(record_res.ID)
				}
				continue
			}

			//creating the record
			is_created, err := p.PVClient.createProVisionRecord(zoneID, ep)
			if err != nil {
				log.Errorf(
					"Failed to create %s record named '%s' to '%s' for ProVision DNS zone '%s': %v",
					ep.RecordType,
					ep.DNSName,
					ep.Targets,
					zoneID,
					err,
				)
				continue
			}

			if !is_created {
				log.Errorf("Failed to create %s record named '%s' to '%s' for ProVision DNS zone '%s'",
					ep.RecordType,
					ep.DNSName,
					ep.Targets,
					zoneID,
				)
			} else {

				log.Debugf("Record created %s for '%s' to '%s'", ep.RecordType, ep.DNSName, ep.Targets)
			}

		}
	}
}

func (p *ProVisionProvider) deleteRecords(deleted pvChangeMap) {
	// run deletions first
	for zoneID, endpoints := range deleted {
		for _, ep := range endpoints {
			log.Infof("deleting %s record named '%s' for ProVision DNS zone ID '%s'.",
				ep.RecordType,
				ep.DNSName,
				zoneID,
			)
			if p.dryRun {
				continue
			}
			//lets check if the record already exists
			var record_res PVResource
			is_exist, err := p.PVClient.getProVisionSpecificRecord(zoneID, ep.DNSName, ep.RecordType, ep.Targets[0], &record_res)

			//log.Debugf("IS EXIST %v %v %v", is_exist, record_res.ID, err)
			if err != nil {
				log.Errorf(
					"Failed to retrieve %s record named '%s' to '%s' for DNS zone id '%s': %v",
					ep.RecordType,
					ep.DNSName,
					ep.Targets,
					zoneID,
					err,
				)
				continue
			}
			//the record doesn't exist so nothing to remove
			if !is_exist {
				log.Debugf("Record doesn't exist so nothing to remove")
				continue
			}

			//log.Debugf("About to remove record id " + record_res.ID)
			err = p.PVClient.deleteProVisionRecord(record_res.ID)

			if err != nil {
				log.Errorf("Failed to delete %s record named '%s' for ProVision DNS zone ID '%s': %v",
					ep.RecordType,
					ep.DNSName,
					zoneID,
					err)
			}

		}
	}
}

func (c PVClientConfig) pushProVisionZone(zoneID string) error {
	client := newHTTPClient(c.SkipTLSVerify)

	url := "/api/v2/dns/zones/" + zoneID + "/push"
	log.Debugf("Executing Push on " + url)
	req, err := c.buildHTTPRequest("POST", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "error executing the request")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return errors.Errorf("status is " + resp.Status)
	}

	return nil
}

func (c PVClientConfig) deleteProVisionRecord(recordID string) error {
	client := newHTTPClient(c.SkipTLSVerify)

	url := "/api/v2/resources/" + recordID
	log.Debugf("Executing DELETE on " + url)
	req, err := c.buildHTTPRequest("DELETE", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "error executing the request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return errors.Errorf("status is " + resp.Status)
	}

	return nil
}

func (c PVClientConfig) flagProVisionRecord(recordID string) (bool, error) {
	client := newHTTPClient(c.SkipTLSVerify)
	body := []byte("{\"attrs\":{\"k8s_externaldns\":\"1\"}}")

	url := "/api/v2/resources/" + recordID
	req, err := c.buildHTTPRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return false, errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, errors.Wrap(err, "error executing the request")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		rbody, err := io.ReadAll(resp.Body)
		return false, errors.Wrap(err, string(rbody))
	}

	return true, nil
}

func (c PVClientConfig) createProVisionRecord(zoneID string, ep *endpoint.Endpoint) (bool, error) {
	client := newHTTPClient(c.SkipTLSVerify)

	RecordHost := ep.DNSName
	RecordType := ep.RecordType
	RecordValue := ep.Targets[0]
	RecordTTL := int(ep.RecordTTL)

	RecordValue = strings.Trim(RecordValue, "\"")
	RecordValue = strings.Trim(RecordValue, " ")

	if !strings.HasSuffix(RecordHost, ".") {
		RecordHost += "."
	}

	attrs := map[string]string{
		"record_host":     RecordHost,
		"record_type":     RecordType,
		"record_value":    RecordValue,
		"k8s_externaldns": "1",
	}

	if RecordTTL > 0 {
		attrs["record_ttl"] = strconv.Itoa(RecordTTL)
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":      "EXTERNALDNS Record",
		"type":      "dnsrecord",
		"parent_id": zoneID,
		"attrs":     attrs,
	})
	if err != nil {
		return false, errors.Wrap(err, "error with json body marshal")
	}

	url := "/api/v2/resources"
	req, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return false, errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, errors.Wrap(err, "error executing the request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		rbody, err := io.ReadAll(resp.Body)
		return false, errors.Wrap(err, string(rbody))
	}

	return true, nil
}

func (c PVClientConfig) getProVisionZones(zoneIDs []string) ([]PVResource, error) {
	client := newHTTPClient(c.SkipTLSVerify)

	body, err := json.Marshal(map[string]interface{}{
		"type":         "dnszone",
		"resource__in": zoneIDs,
	})
	//log.Debugf("BODY: %v", string(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building the json body")
	}

	url := "/api/v2/resources/query"
	//log.Debugf(url + " BODY : " + string(body))
	req, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error retrieving zone(s) from ProVision: %v", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.Wrapf(err, "error retrieving zone(s) from ProVision: %v", resp.Status)
	}

	zones := []PVResource{}
	json.NewDecoder(resp.Body).Decode(&zones)

	return zones, nil
}

func (c PVClientConfig) getProVisionRecords(zoneID string) ([]PVResource, error) {
	client := newHTTPClient(c.SkipTLSVerify)

	body, err := json.Marshal(map[string]interface{}{
		"type":      "dnsrecord",
		"parent_id": zoneID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error building the json body")
	}

	url := "/api/v2/resources/query"
	//log.Debugf(url + " BODY : " + string(body))
	req, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error retrieving records(s) from ProVision: %v, %v", url, zoneID)
	}
	defer resp.Body.Close()

	records := []PVResource{}
	json.NewDecoder(resp.Body).Decode(&records)

	return records, nil
}

func (c PVClientConfig) getProVisionSpecificRecord(ZoneID, RecordHost, RecordType, RecordValue string, OutputRes *PVResource) (bool, error) {
	client := newHTTPClient(c.SkipTLSVerify)

	//lets do some preprocessing
	if !strings.HasSuffix(RecordHost, ".") {
		RecordHost += "."
	}

	var filtersmap = []map[string]string{
		{
			"key":     "record_host",
			"value":   RecordHost,
			"compare": "=",
			"rule":    "AND",
		},
		{
			"key":     "record_type",
			"value":   RecordType,
			"compare": "=",
			"rule":    "AND",
		},
	}

	if RecordValue != "" {
		if RecordType == "TXT" {
			RecordValue = strings.Trim(RecordValue, "\"")
		}

		RecordValue = strings.Trim(RecordValue, " ")

		var valuemap = map[string]string{
			"key":     "record_value",
			"value":   RecordValue,
			"compare": "=",
			"rule":    "AND",
		}

		filtersmap = append(filtersmap, valuemap)
	}

	body, err := json.Marshal(map[string]interface{}{
		"type":      "dnsrecord",
		"parent_id": ZoneID,
		"attrs":     filtersmap,
	})
	if err != nil {
		return false, errors.Wrap(err, "error building the json body")
	}

	url := "/api/v2/resources/query"
	//log.Debugf(url + " BODY : " + string(body))
	req, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return false, errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, errors.Wrap(err, "error executing the request")
	}
	defer resp.Body.Close()

	rbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.Wrap(err, "can not read the body")
	}

	var z []PVResource

	json.Unmarshal(rbody, &z)
	//log.Debugf(string(rbody))
	//log.Debugln(z)
	if len(z) == 0 {
		return false, nil
	}

	OutputRes.ID = z[0].ID
	OutputRes.Name = z[0].Name
	OutputRes.Attrs = z[0].Attrs

	return true, nil
}

func newHTTPClient(skipTLSVerify bool) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipTLSVerify,
			},
		},
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c PVClientConfig) buildHTTPRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.Host+url, body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(c.Username, c.Password))

	if method != "GET" {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, err
}

//extractOwnerFromTXTRecord takes a single text property string and returns the owner after parsing theowner string.
func extractOwnerfromTXTRecord(propString string) (string, error) {
	if len(propString) == 0 {
		return "", errors.Errorf("External-DNS Owner not found")
	}
	re := regexp.MustCompile(`external-dns/owner=[^,]+`)
	match := re.FindStringSubmatch(propString)
	if len(match) == 0 {
		return "", errors.Errorf("External-DNS Owner not found, %s", propString)
	}
	return strings.Split(match[0], "=")[1], nil
}
