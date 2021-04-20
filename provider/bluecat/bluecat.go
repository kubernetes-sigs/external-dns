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

// TODO: Ensure we have proper error handling/logging for API calls to Bluecat. getBluecatGatewayToken has a good example of this

package bluecat

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type bluecatConfig struct {
	GatewayHost      string `json:"gatewayHost"`
	GatewayUsername  string `json:"gatewayUsername"`
	GatewayPassword  string `json:"gatewayPassword"`
	DNSConfiguration string `json:"dnsConfiguration"`
	View             string `json:"dnsView"`
	RootZone         string `json:"rootZone"`
	SkipTLSVerify    bool   `json:"skipTLSVerify"`
}

// BluecatProvider implements the DNS provider for Bluecat DNS
type BluecatProvider struct {
	provider.BaseProvider
	domainFilter     endpoint.DomainFilter
	zoneIDFilter     provider.ZoneIDFilter
	dryRun           bool
	RootZone         string
	DNSConfiguration string
	View             string
	gatewayClient    GatewayClient
}

type GatewayClient interface {
	getBluecatZones(zoneName string) ([]BluecatZone, error)
	getHostRecords(zone string, records *[]BluecatHostRecord) error
	getCNAMERecords(zone string, records *[]BluecatCNAMERecord) error
	getHostRecord(name string, record *BluecatHostRecord) error
	getCNAMERecord(name string, record *BluecatCNAMERecord) error
	createHostRecord(zone string, req *bluecatCreateHostRecordRequest) (res interface{}, err error)
	createCNAMERecord(zone string, req *bluecatCreateCNAMERecordRequest) (res interface{}, err error)
	deleteHostRecord(name string) (err error)
	deleteCNAMERecord(name string) (err error)
	buildHTTPRequest(method, url string, body io.Reader) (*http.Request, error)
	getTXTRecords(zone string, records *[]BluecatTXTRecord) error
	getTXTRecord(name string, record *BluecatTXTRecord) error
	createTXTRecord(zone string, req *bluecatCreateTXTRecordRequest) (res interface{}, err error)
	deleteTXTRecord(name string) error
}

// GatewayClientConfig defines new client on bluecat gateway
type GatewayClientConfig struct {
	Cookie           http.Cookie
	Token            string
	Host             string
	DNSConfiguration string
	View             string
	RootZone         string
	SkipTLSVerify    bool
}

// BluecatZone defines a zone to hold records
type BluecatZone struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
	Type       string `json:"type"`
}

// BluecatHostRecord defines dns Host record
type BluecatHostRecord struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
	Type       string `json:"type"`
}

// BluecatCNAMERecord defines dns CNAME record
type BluecatCNAMERecord struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Properties string `json:"properties"`
	Type       string `json:"type"`
}

// BluecatTXTRecord defines dns TXT record
type BluecatTXTRecord struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Text string `json:"text"`
}

type bluecatRecordSet struct {
	obj interface{}
	res interface{}
}

type bluecatCreateHostRecordRequest struct {
	AbsoluteName string `json:"absolute_name"`
	IP4Address   string `json:"ip4_address"`
	TTL          int    `json:"ttl"`
	Properties   string `json:"properties"`
}

type bluecatCreateCNAMERecordRequest struct {
	AbsoluteName string `json:"absolute_name"`
	LinkedRecord string `json:"linked_record"`
	TTL          int    `json:"ttl"`
	Properties   string `json:"properties"`
}

type bluecatCreateTXTRecordRequest struct {
	AbsoluteName string `json:"absolute_name"`
	Text         string `json:"txt"`
}

// NewBluecatProvider creates a new Bluecat provider.
//
// Returns a pointer to the provider or an error if a provider could not be created.
func NewBluecatProvider(configFile string, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool) (*BluecatProvider, error) {
	contents, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read Bluecat config file %v", configFile)
	}

	cfg := bluecatConfig{}
	err = json.Unmarshal(contents, &cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read Bluecat config file %v", configFile)
	}

	token, cookie, err := getBluecatGatewayToken(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get API token from Bluecat Gateway")
	}
	gatewayClient := NewGatewayClient(cookie, token, cfg.GatewayHost, cfg.DNSConfiguration, cfg.View, cfg.RootZone, cfg.SkipTLSVerify)

	provider := &BluecatProvider{
		domainFilter:     domainFilter,
		zoneIDFilter:     zoneIDFilter,
		dryRun:           dryRun,
		gatewayClient:    gatewayClient,
		DNSConfiguration: cfg.DNSConfiguration,
		View:             cfg.View,
		RootZone:         cfg.RootZone,
	}
	return provider, nil
}

// NewGatewayClient creates and returns a new Bluecat gateway client
func NewGatewayClient(cookie http.Cookie, token, gatewayHost, dnsConfiguration, view, rootZone string, skipTLSVerify bool) GatewayClientConfig {
	// Right now the Bluecat gateway doesn't seem to have a way to get the root zone from the API. If the user
	// doesn't provide one via the config file we'll assume it's 'com'
	if rootZone == "" {
		rootZone = "com"
	}
	return GatewayClientConfig{
		Cookie:           cookie,
		Token:            token,
		Host:             gatewayHost,
		DNSConfiguration: dnsConfiguration,
		View:             view,
		RootZone:         rootZone,
		SkipTLSVerify:    skipTLSVerify,
	}
}

// Records fetches Host, CNAME, and TXT records from bluecat gateway
func (p *BluecatProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, err error) {
	zones, err := p.zones()
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch zones")
	}

	for _, zone := range zones {
		log.Debugf("fetching records from zone '%s'", zone)
		var resH []BluecatHostRecord
		err = p.gatewayClient.getHostRecords(zone, &resH)
		if err != nil {
			return nil, errors.Wrapf(err, "could not fetch host records for zone: %v", zone)
		}
		for _, rec := range resH {
			propMap := splitProperties(rec.Properties)
			ips := strings.Split(propMap["addresses"], ",")
			if _, ok := propMap["ttl"]; ok {
				ttl, err := strconv.Atoi(propMap["ttl"])
				if err != nil {
					return nil, errors.Wrapf(err, "could not parse ttl '%d' as int for host record %v", ttl, rec.Name)
				}

				for _, ip := range ips {
					ep := endpoint.NewEndpointWithTTL(propMap["absoluteName"], endpoint.RecordTypeA, endpoint.TTL(ttl), ip)
					endpoints = append(endpoints, ep)
				}
			} else {
				for _, ip := range ips {
					ep := endpoint.NewEndpoint(propMap["absoluteName"], endpoint.RecordTypeA, ip)
					endpoints = append(endpoints, ep)
				}
			}
		}

		var resC []BluecatCNAMERecord
		err = p.gatewayClient.getCNAMERecords(zone, &resC)
		if err != nil {
			return nil, errors.Wrapf(err, "could not fetch CNAME records for zone: %v", zone)
		}
		for _, rec := range resC {
			propMap := splitProperties(rec.Properties)
			if _, ok := propMap["ttl"]; ok {
				ttl, err := strconv.Atoi(propMap["ttl"])
				if err != nil {
					return nil, errors.Wrapf(err, "could not parse ttl '%d' as int for CNAME record %v", ttl, rec.Name)
				}
				endpoints = append(endpoints, endpoint.NewEndpointWithTTL(propMap["absoluteName"], endpoint.RecordTypeCNAME, endpoint.TTL(ttl), propMap["linkedRecordName"]))
			} else {
				endpoints = append(endpoints, endpoint.NewEndpoint(propMap["absoluteName"], endpoint.RecordTypeCNAME, propMap["linkedRecordName"]))
			}
		}

		var resT []BluecatTXTRecord
		err = p.gatewayClient.getTXTRecords(zone, &resT)
		if err != nil {
			return nil, errors.Wrapf(err, "could not fetch TXT records for zone: %v", zone)
		}
		for _, rec := range resT {
			endpoints = append(endpoints, endpoint.NewEndpoint(rec.Name, endpoint.RecordTypeTXT, rec.Text))
		}
	}

	log.Debugf("fetched %d records from Bluecat", len(endpoints))
	return endpoints, nil
}

// ApplyChanges updates necessary zones and replaces old records with new ones
//
// Returns nil upon success and err is there is an error
func (p *BluecatProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
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

	// TODO: add bluecat deploy API call here

	return nil
}

type bluecatChangeMap map[string][]*endpoint.Endpoint

func (p *BluecatProvider) mapChanges(zones []string, changes *plan.Changes) (bluecatChangeMap, bluecatChangeMap) {
	created := bluecatChangeMap{}
	deleted := bluecatChangeMap{}

	mapChange := func(changeMap bluecatChangeMap, change *endpoint.Endpoint) {
		zone := p.findZone(zones, change.DNSName)
		if zone == "" {
			log.Debugf("ignoring changes to '%s' because a suitable Bluecat DNS zone was not found", change.DNSName)
			return
		}
		changeMap[zone] = append(changeMap[zone], change)
	}

	for _, change := range changes.Delete {
		mapChange(deleted, change)
	}
	for _, change := range changes.UpdateOld {
		mapChange(deleted, change)
	}
	for _, change := range changes.Create {
		mapChange(created, change)
	}
	for _, change := range changes.UpdateNew {
		mapChange(created, change)
	}

	return created, deleted
}

// findZone finds the most specific matching zone for a given record 'name' from a list of all zones
func (p *BluecatProvider) findZone(zones []string, name string) string {
	var result string

	for _, zone := range zones {
		if strings.HasSuffix(name, "."+zone) {
			if result == "" || len(zone) > len(result) {
				result = zone
			}
		} else if strings.EqualFold(name, zone) {
			if result == "" || len(zone) > len(result) {
				result = zone
			}
		}
	}

	return result
}

func (p *BluecatProvider) zones() ([]string, error) {
	log.Debugf("retrieving Bluecat zones for configuration: %s, view: %s", p.DNSConfiguration, p.View)
	var zones []string

	zonelist, err := p.gatewayClient.getBluecatZones(p.RootZone)
	if err != nil {
		return nil, err
	}

	for _, zone := range zonelist {
		if !p.domainFilter.Match(zone.Name) {
			continue
		}

		// TODO: match to absoluteName(string) not Id(int)
		if !p.zoneIDFilter.Match(strconv.Itoa(zone.ID)) {
			continue
		}

		zoneProps := splitProperties(zone.Properties)

		zones = append(zones, zoneProps["absoluteName"])
	}
	log.Debugf("found %d zones", len(zones))
	return zones, nil
}

func (p *BluecatProvider) createRecords(created bluecatChangeMap) {
	for zone, endpoints := range created {
		for _, ep := range endpoints {
			if p.dryRun {
				log.Infof("would create %s record named '%s' to '%s' for Bluecat DNS zone '%s'.",
					ep.RecordType,
					ep.DNSName,
					ep.Targets,
					zone,
				)
				continue
			}

			log.Infof("creating %s record named '%s' to '%s' for Bluecat DNS zone '%s'.",
				ep.RecordType,
				ep.DNSName,
				ep.Targets,
				zone,
			)

			recordSet, err := p.recordSet(ep, false)
			if err != nil {
				log.Errorf(
					"Failed to retrieve %s record named '%s' to '%s' for Bluecat DNS zone '%s': %v",
					ep.RecordType,
					ep.DNSName,
					ep.Targets,
					zone,
					err,
				)
				continue
			}
			var response interface{}
			switch ep.RecordType {
			case endpoint.RecordTypeA:
				response, err = p.gatewayClient.createHostRecord(zone, recordSet.obj.(*bluecatCreateHostRecordRequest))
			case endpoint.RecordTypeCNAME:
				response, err = p.gatewayClient.createCNAMERecord(zone, recordSet.obj.(*bluecatCreateCNAMERecordRequest))
			case endpoint.RecordTypeTXT:
				response, err = p.gatewayClient.createTXTRecord(zone, recordSet.obj.(*bluecatCreateTXTRecordRequest))
			}
			log.Debugf("Response from create: %v", response)
			if err != nil {
				log.Errorf(
					"Failed to create %s record named '%s' to '%s' for Bluecat DNS zone '%s': %v",
					ep.RecordType,
					ep.DNSName,
					ep.Targets,
					zone,
					err,
				)
			}
		}
	}
}

func (p *BluecatProvider) deleteRecords(deleted bluecatChangeMap) {
	// run deletions first
	for zone, endpoints := range deleted {
		for _, ep := range endpoints {
			if p.dryRun {
				log.Infof("would delete %s record named '%s' for Bluecat DNS zone '%s'.",
					ep.RecordType,
					ep.DNSName,
					zone,
				)
				continue
			} else {
				log.Infof("deleting %s record named '%s' for Bluecat DNS zone '%s'.",
					ep.RecordType,
					ep.DNSName,
					zone,
				)

				recordSet, err := p.recordSet(ep, true)
				if err != nil {
					log.Errorf(
						"Failed to retrieve %s record named '%s' to '%s' for Bluecat DNS zone '%s': %v",
						ep.RecordType,
						ep.DNSName,
						ep.Targets,
						zone,
						err,
					)
					continue
				}

				switch ep.RecordType {
				case endpoint.RecordTypeA:
					for _, record := range *recordSet.res.(*[]BluecatHostRecord) {
						err = p.gatewayClient.deleteHostRecord(record.Name)
					}
				case endpoint.RecordTypeCNAME:
					for _, record := range *recordSet.res.(*[]BluecatCNAMERecord) {
						err = p.gatewayClient.deleteCNAMERecord(record.Name)
					}
				case endpoint.RecordTypeTXT:
					for _, record := range *recordSet.res.(*[]BluecatTXTRecord) {
						err = p.gatewayClient.deleteTXTRecord(record.Name)
					}
				}
				if err != nil {
					log.Errorf("Failed to delete %s record named '%s' for Bluecat DNS zone '%s': %v",
						ep.RecordType,
						ep.DNSName,
						zone,
						err)
				}
			}
		}
	}
}

func (p *BluecatProvider) recordSet(ep *endpoint.Endpoint, getObject bool) (recordSet bluecatRecordSet, err error) {
	switch ep.RecordType {
	case endpoint.RecordTypeA:
		var res []BluecatHostRecord
		obj := bluecatCreateHostRecordRequest{
			AbsoluteName: ep.DNSName,
			IP4Address:   ep.Targets[0],
			TTL:          int(ep.RecordTTL),
			Properties:   "",
		}
		if getObject {
			var record BluecatHostRecord
			err = p.gatewayClient.getHostRecord(ep.DNSName, &record)
			if err != nil {
				return
			}
			res = append(res, record)
		}
		recordSet = bluecatRecordSet{
			obj: &obj,
			res: &res,
		}
	case endpoint.RecordTypeCNAME:
		var res []BluecatCNAMERecord
		obj := bluecatCreateCNAMERecordRequest{
			AbsoluteName: ep.DNSName,
			LinkedRecord: ep.Targets[0],
			TTL:          int(ep.RecordTTL),
			Properties:   "",
		}
		if getObject {
			var record BluecatCNAMERecord
			err = p.gatewayClient.getCNAMERecord(ep.DNSName, &record)
			if err != nil {
				return
			}
			res = append(res, record)
		}
		recordSet = bluecatRecordSet{
			obj: &obj,
			res: &res,
		}
	case endpoint.RecordTypeTXT:
		var res []BluecatTXTRecord
		// TODO: Allow setting TTL
		// This is not implemented in the Bluecat Gateway
		obj := bluecatCreateTXTRecordRequest{
			AbsoluteName: ep.DNSName,
			Text:         ep.Targets[0],
		}
		if getObject {
			var record BluecatTXTRecord
			err = p.gatewayClient.getTXTRecord(ep.DNSName, &record)
			if err != nil {
				return
			}
			res = append(res, record)
		}
		recordSet = bluecatRecordSet{
			obj: &obj,
			res: &res,
		}
	}
	return
}

// getBluecatGatewayToken retrieves a Bluecat Gateway API token.
func getBluecatGatewayToken(cfg bluecatConfig) (string, http.Cookie, error) {
	body, err := json.Marshal(map[string]string{
		"username": cfg.GatewayUsername,
		"password": cfg.GatewayPassword,
	})
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "could not unmarshal credentials for bluecat gateway config")
	}

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.SkipTLSVerify},
		}}

	resp, err := c.Post(cfg.GatewayHost+"/rest_login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "error obtaining API token from bluecat gateway")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		details, _ := ioutil.ReadAll(resp.Body)
		return "", http.Cookie{}, errors.Errorf("got HTTP response code %v, detailed message: %v", resp.StatusCode, string(details))
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "error reading get_token response from bluecat gateway")
	}

	resJSON := map[string]string{}
	err = json.Unmarshal(res, &resJSON)
	if err != nil {
		return "", http.Cookie{}, errors.Wrap(err, "error unmarshaling json response (auth) from bluecat gateway")
	}

	// Example response: {"access_token": "BAMAuthToken: abc123"}
	// We only care about the actual token string - i.e. abc123
	// The gateway also creates a cookie as part of the response. This seems to be the actual auth mechanism, at least
	// for now.
	return strings.Split(resJSON["access_token"], " ")[1], *resp.Cookies()[0], nil
}

func (c GatewayClientConfig) getBluecatZones(zoneName string) ([]BluecatZone, error) {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}
	zonePath := expandZone(zoneName)
	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error retrieving zone(s) from gateway: %v, %v", url, zoneName)
	}

	defer resp.Body.Close()

	zones := []BluecatZone{}
	json.NewDecoder(resp.Body).Decode(&zones)

	// Bluecat Gateway only returns subzones one level deeper than the provided zone
	// so this recursion is needed to traverse subzones until none are returned
	for _, zone := range zones {
		zoneProps := splitProperties(zone.Properties)
		subZones, err := c.getBluecatZones(zoneProps["absoluteName"])
		if err != nil {
			return nil, errors.Wrapf(err, "error retrieving subzones from gateway: %v", zoneName)
		}
		zones = append(zones, subZones...)
	}

	return zones, nil
}

func (c GatewayClientConfig) getHostRecords(zone string, records *[]BluecatHostRecord) error {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	zonePath := expandZone(zone)

	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "host_records/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving record(s) from gateway: %v", zone)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(records)
	log.Debugf("Get Host Records Response: %v", records)

	return nil
}

func (c GatewayClientConfig) getCNAMERecords(zone string, records *[]BluecatCNAMERecord) error {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	zonePath := expandZone(zone)

	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "cname_records/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving record(s) from gateway: %v", zone)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(records)
	log.Debugf("Get CName Records Response: %v", records)

	return nil
}

func (c GatewayClientConfig) getTXTRecords(zone string, records *[]BluecatTXTRecord) error {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	zonePath := expandZone(zone)

	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "text_records/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}
	log.Debugf("Request: %v", req)

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving record(s) from gateway: %v", zone)
	}
	log.Debugf("Get Txt Records response: %v", resp)

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(records)
	log.Debugf("Get TXT Records Body: %v", records)

	return nil
}

func (c GatewayClientConfig) getHostRecord(name string, record *BluecatHostRecord) error {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"host_records/" + name + "/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrapf(err, "error building http request: %v", name)
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving record(s) from gateway: %v", name)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(record)
	log.Debugf("Get Host Record Response: %v", record)
	return nil
}

func (c GatewayClientConfig) getCNAMERecord(name string, record *BluecatCNAMERecord) error {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"cname_records/" + name + "/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrapf(err, "error building http request: %v", name)
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving record(s) from gateway: %v", name)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(record)
	log.Debugf("Get CName Record Response: %v", record)
	return nil
}

func (c GatewayClientConfig) getTXTRecord(name string, record *BluecatTXTRecord) error {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"text_records/" + name + "/"
	req, err := c.buildHTTPRequest("GET", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error retrieving record(s) from gateway: %v", name)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(record)
	log.Debugf("Get TXT Record Response: %v", record)

	return nil
}

func (c GatewayClientConfig) createHostRecord(zone string, req *bluecatCreateHostRecordRequest) (res interface{}, err error) {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "host_records/"
	body, _ := json.Marshal(req)
	hreq, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building http request")
	}
	hreq.Header.Add("Content-Type", "application/json")
	res, err = client.Do(hreq)

	return
}

func (c GatewayClientConfig) createCNAMERecord(zone string, req *bluecatCreateCNAMERecordRequest) (res interface{}, err error) {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "cname_records/"
	body, _ := json.Marshal(req)

	hreq, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building http request")
	}

	hreq.Header.Add("Content-Type", "application/json")
	res, err = client.Do(hreq)

	return
}

func (c GatewayClientConfig) createTXTRecord(zone string, req *bluecatCreateTXTRecordRequest) (interface{}, error) {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	zonePath := expandZone(zone)
	// Remove the trailing 'zones/'
	zonePath = strings.TrimSuffix(zonePath, "zones/")

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration + "/views/" + c.View + "/" + zonePath + "text_records/"
	body, _ := json.Marshal(req)
	hreq, err := c.buildHTTPRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	hreq.Header.Add("Content-Type", "application/json")
	res, err := client.Do(hreq)

	return res, err
}

func (c GatewayClientConfig) deleteHostRecord(name string) (err error) {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"host_records/" + name + "/"
	req, err := c.buildHTTPRequest("DELETE", url, nil)
	if err != nil {
		return errors.Wrapf(err, "error building http request: %v", name)
	}

	_, err = client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error deleting record(s) from gateway: %v", name)
	}

	return nil
}

func (c GatewayClientConfig) deleteCNAMERecord(name string) (err error) {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"cname_records/" + name + "/"
	req, err := c.buildHTTPRequest("DELETE", url, nil)
	if err != nil {
		return errors.Wrapf(err, "error building http request: %v", name)
	}

	_, err = client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error deleting record(s) from gateway: %v", name)
	}

	return nil
}

func (c GatewayClientConfig) deleteTXTRecord(name string) error {
	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.SkipTLSVerify},
	}
	client := &http.Client{
		Transport: transportCfg,
	}

	url := c.Host + "/api/v1/configurations/" + c.DNSConfiguration +
		"/views/" + c.View + "/" +
		"text_records/" + name + "/"

	req, err := c.buildHTTPRequest("DELETE", url, nil)
	if err != nil {
		return errors.Wrap(err, "error building http request")
	}

	_, err = client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "error deleting record(s) from gateway: %v", name)
	}

	return nil
}

//buildHTTPRequest builds a standard http Request and adds authentication headers required by Bluecat Gateway
func (c GatewayClientConfig) buildHTTPRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+c.Token)
	req.AddCookie(&c.Cookie)
	return req, err
}

//splitProperties is a helper function to break a '|' separated string into key/value pairs
// i.e. "foo=bar|baz=mop"
func splitProperties(props string) map[string]string {
	propMap := make(map[string]string)

	// remove trailing | character before we split
	props = strings.TrimSuffix(props, "|")

	splits := strings.Split(props, "|")
	for _, pair := range splits {
		items := strings.Split(pair, "=")
		propMap[items[0]] = items[1]
	}

	return propMap
}

//expandZone takes an absolute domain name such as 'example.com' and returns a zone hierarchy used by Bluecat Gateway,
//such as '/zones/com/zones/example/zones/'
func expandZone(zone string) string {
	ze := "zones/"
	parts := strings.Split(zone, ".")
	if len(parts) > 1 {
		last := len(parts) - 1
		for i := range parts {
			ze = ze + parts[last-i] + "/zones/"
		}
	} else {
		ze = ze + zone + "/zones/"
	}
	return ze
}
