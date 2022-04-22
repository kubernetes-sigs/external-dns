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
// TODO: Remove studdering
// TODO: Make API calls more consistent (eg error handling on HTTP response codes)
// TODO: zone-id-filter does not seem to work with our provider

package bluecat

import (
	"context"
	"encoding/json"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	api "sigs.k8s.io/external-dns/provider/bluecat/gateway"
)

// BluecatProvider implements the DNS provider for Bluecat DNS
type BluecatProvider struct {
	provider.BaseProvider
	domainFilter     endpoint.DomainFilter
	zoneIDFilter     provider.ZoneIDFilter
	dryRun           bool
	RootZone         string
	DNSConfiguration string
	DNSServerName    string
	DNSDeployType    string
	View             string
	gatewayClient    api.GatewayClient
	TxtPrefix        string
	TxtSuffix        string
}

type bluecatRecordSet struct {
	obj interface{}
	res interface{}
}

// NewBluecatProvider creates a new Bluecat provider.
//
// Returns a pointer to the provider or an error if a provider could not be created.
func NewBluecatProvider(configFile, dnsConfiguration, dnsServerName, dnsDeployType, dnsView, gatewayHost, rootZone, txtPrefix, txtSuffix string, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun, skipTLSVerify bool) (*BluecatProvider, error) {
	cfg := api.BluecatConfig{}
	contents, err := os.ReadFile(configFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cfg = api.BluecatConfig{
				GatewayHost:      gatewayHost,
				DNSConfiguration: dnsConfiguration,
				DNSServerName:    dnsServerName,
				DNSDeployType:    dnsDeployType,
				View:             dnsView,
				RootZone:         rootZone,
				SkipTLSVerify:    skipTLSVerify,
				GatewayUsername:  "",
				GatewayPassword:  "",
			}
		} else {
			return nil, errors.Wrapf(err, "failed to read Bluecat config file %v", configFile)
		}
	} else {
		err = json.Unmarshal(contents, &cfg)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse Bluecat JSON config file %v", configFile)
		}
	}

	if !api.IsValidDNSDeployType(cfg.DNSDeployType) {
		return nil, errors.Errorf("%v is not a valid deployment type", cfg.DNSDeployType)
	}

	token, cookie, err := api.GetBluecatGatewayToken(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get API token from Bluecat Gateway")
	}
	gatewayClient := api.NewGatewayClientConfig(cookie, token, cfg.GatewayHost, cfg.DNSConfiguration, cfg.View, cfg.RootZone, cfg.DNSServerName, cfg.SkipTLSVerify)

	provider := &BluecatProvider{
		domainFilter:     domainFilter,
		zoneIDFilter:     zoneIDFilter,
		dryRun:           dryRun,
		gatewayClient:    gatewayClient,
		DNSConfiguration: cfg.DNSConfiguration,
		DNSServerName:    cfg.DNSServerName,
		DNSDeployType:    cfg.DNSDeployType,
		View:             cfg.View,
		RootZone:         cfg.RootZone,
		TxtPrefix:        txtPrefix,
		TxtSuffix:        txtSuffix,
	}
	return provider, nil
}

// Records fetches Host, CNAME, and TXT records from bluecat gateway
func (p *BluecatProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, err error) {
	zones, err := p.zones()
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch zones")
	}

	// Parsing Text records first, so we can get the owner from them.
	for _, zone := range zones {
		log.Debugf("fetching records from zone '%s'", zone)

		var resT []api.BluecatTXTRecord
		err = p.gatewayClient.GetTXTRecords(zone, &resT)
		if err != nil {
			return nil, errors.Wrapf(err, "could not fetch TXT records for zone: %v", zone)
		}
		for _, rec := range resT {
			tempEndpoint := endpoint.NewEndpoint(rec.Name, endpoint.RecordTypeTXT, rec.Properties)
			tempEndpoint.Labels[endpoint.OwnerLabelKey], err = extractOwnerfromTXTRecord(rec.Properties)
			if err != nil {
				log.Debugf("External DNS Owner %s", err)
			}
			endpoints = append(endpoints, tempEndpoint)
		}

		var resH []api.BluecatHostRecord
		err = p.gatewayClient.GetHostRecords(zone, &resH)
		if err != nil {
			return nil, errors.Wrapf(err, "could not fetch host records for zone: %v", zone)
		}
		var ep *endpoint.Endpoint
		for _, rec := range resH {
			propMap := api.SplitProperties(rec.Properties)
			ips := strings.Split(propMap["addresses"], ",")
			for _, ip := range ips {
				if _, ok := propMap["ttl"]; ok {
					ttl, err := strconv.Atoi(propMap["ttl"])
					if err != nil {
						return nil, errors.Wrapf(err, "could not parse ttl '%d' as int for host record %v", ttl, rec.Name)
					}
					ep = endpoint.NewEndpointWithTTL(propMap["absoluteName"], endpoint.RecordTypeA, endpoint.TTL(ttl), ip)
				} else {
					ep = endpoint.NewEndpoint(propMap["absoluteName"], endpoint.RecordTypeA, ip)
				}
				for _, txtRec := range resT {
					if strings.Compare(p.TxtPrefix+rec.Name+p.TxtSuffix, txtRec.Name) == 0 {
						ep.Labels[endpoint.OwnerLabelKey], err = extractOwnerfromTXTRecord(txtRec.Properties)
						if err != nil {
							log.Debugf("External DNS Owner %s", err)
						}
					}
				}
				endpoints = append(endpoints, ep)
			}
		}

		var resC []api.BluecatCNAMERecord
		err = p.gatewayClient.GetCNAMERecords(zone, &resC)
		if err != nil {
			return nil, errors.Wrapf(err, "could not fetch CNAME records for zone: %v", zone)
		}

		for _, rec := range resC {
			propMap := api.SplitProperties(rec.Properties)
			if _, ok := propMap["ttl"]; ok {
				ttl, err := strconv.Atoi(propMap["ttl"])
				if err != nil {
					return nil, errors.Wrapf(err, "could not parse ttl '%d' as int for CNAME record %v", ttl, rec.Name)
				}
				ep = endpoint.NewEndpointWithTTL(propMap["absoluteName"], endpoint.RecordTypeCNAME, endpoint.TTL(ttl), propMap["linkedRecordName"])
			} else {
				ep = endpoint.NewEndpoint(propMap["absoluteName"], endpoint.RecordTypeCNAME, propMap["linkedRecordName"])
			}
			for _, txtRec := range resT {
				if strings.Compare(p.TxtPrefix+rec.Name+p.TxtSuffix, txtRec.Name) == 0 {
					ep.Labels[endpoint.OwnerLabelKey], err = extractOwnerfromTXTRecord(txtRec.Properties)
					if err != nil {
						log.Debugf("External DNS Owner %s", err)
					}
				}
			}
			endpoints = append(endpoints, ep)
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

	if p.DNSServerName != "" {
		if p.dryRun {
			log.Debug("Not executing deploy because this is running in dry-run mode")
		} else {
			switch p.DNSDeployType {
			case "full-deploy":
				err := p.gatewayClient.ServerFullDeploy()
				if err != nil {
					return err
				}
			case "no-deploy":
				log.Debug("Not executing deploy because DNSDeployType is set to 'no-deploy'")
			}
		}
	} else {
		log.Debug("Not executing deploy because server name was not provided")
	}

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

	zonelist, err := p.gatewayClient.GetBluecatZones(p.RootZone)
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

		zoneProps := api.SplitProperties(zone.Properties)

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
				err = p.gatewayClient.CreateHostRecord(zone, recordSet.obj.(*api.BluecatCreateHostRecordRequest))
			case endpoint.RecordTypeCNAME:
				err = p.gatewayClient.CreateCNAMERecord(zone, recordSet.obj.(*api.BluecatCreateCNAMERecordRequest))
			case endpoint.RecordTypeTXT:
				err = p.gatewayClient.CreateTXTRecord(zone, recordSet.obj.(*api.BluecatCreateTXTRecordRequest))
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
					for _, record := range *recordSet.res.(*[]api.BluecatHostRecord) {
						err = p.gatewayClient.DeleteHostRecord(record.Name, zone)
					}
				case endpoint.RecordTypeCNAME:
					for _, record := range *recordSet.res.(*[]api.BluecatCNAMERecord) {
						err = p.gatewayClient.DeleteCNAMERecord(record.Name, zone)
					}
				case endpoint.RecordTypeTXT:
					for _, record := range *recordSet.res.(*[]api.BluecatTXTRecord) {
						err = p.gatewayClient.DeleteTXTRecord(record.Name, zone)
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

func (p *BluecatProvider) recordSet(ep *endpoint.Endpoint, getObject bool) (bluecatRecordSet, error) {
	recordSet := bluecatRecordSet{}
	switch ep.RecordType {
	case endpoint.RecordTypeA:
		var res []api.BluecatHostRecord
		obj := api.BluecatCreateHostRecordRequest{
			AbsoluteName: ep.DNSName,
			IP4Address:   ep.Targets[0],
			TTL:          int(ep.RecordTTL),
			Properties:   "",
		}
		if getObject {
			var record api.BluecatHostRecord
			err := p.gatewayClient.GetHostRecord(ep.DNSName, &record)
			if err != nil {
				return bluecatRecordSet{}, err
			}
			res = append(res, record)
		}
		recordSet = bluecatRecordSet{
			obj: &obj,
			res: &res,
		}
	case endpoint.RecordTypeCNAME:
		var res []api.BluecatCNAMERecord
		obj := api.BluecatCreateCNAMERecordRequest{
			AbsoluteName: ep.DNSName,
			LinkedRecord: ep.Targets[0],
			TTL:          int(ep.RecordTTL),
			Properties:   "",
		}
		if getObject {
			var record api.BluecatCNAMERecord
			err := p.gatewayClient.GetCNAMERecord(ep.DNSName, &record)
			if err != nil {
				return bluecatRecordSet{}, err
			}
			res = append(res, record)
		}
		recordSet = bluecatRecordSet{
			obj: &obj,
			res: &res,
		}
	case endpoint.RecordTypeTXT:
		var res []api.BluecatTXTRecord
		// TODO: Allow setting TTL
		// This is not implemented in the Bluecat Gateway
		obj := api.BluecatCreateTXTRecordRequest{
			AbsoluteName: ep.DNSName,
			Text:         ep.Targets[0],
		}
		if getObject {
			var record api.BluecatTXTRecord
			err := p.gatewayClient.GetTXTRecord(ep.DNSName, &record)
			if err != nil {
				return bluecatRecordSet{}, err
			}
			res = append(res, record)
		}
		recordSet = bluecatRecordSet{
			obj: &obj,
			res: &res,
		}
	}
	return recordSet, nil
}

// extractOwnerFromTXTRecord takes a single text property string and returns the owner after parsing the owner string.
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
