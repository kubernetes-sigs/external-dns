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

package cloudns

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	cloudns "github.com/ppmathis/cloudns-go"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type ClouDNSProvider struct {
	provider.BaseProvider
	client       *cloudns.Client
	context      context.Context
	domainFilter endpoint.DomainFilter
	zoneIDFilter provider.ZoneIDFilter
	ownerID      string
	dryRun       bool
	testing      bool
}

type ClouDNSConfig struct {
	Context      context.Context
	DomainFilter endpoint.DomainFilter
	ZoneIDFilter provider.ZoneIDFilter
	OwnerID      string
	DryRun       bool
	Testing      bool
}

func NewClouDNSProvider(config ClouDNSConfig) (*ClouDNSProvider, error) {

	var client *cloudns.Client

	log.Info("Creating ClouDNS Provider")

	loginType, ok := os.LookupEnv("CLOUDNS_LOGIN_TYPE")
	if !ok {
		return nil, fmt.Errorf("CLOUDNS_LOGIN_TYPE is not set")
	}
	if loginType != "user-id" && loginType != "sub-user" && loginType != "sub-user-name" {
		return nil, fmt.Errorf("CLOUDNS_LOGIN_TYPE is not valid")
	}

	userPassword, ok := os.LookupEnv("CLOUDNS_USER_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("CLOUDNS_USER_PASSWORD is not set")
	}

	switch loginType {
	case "user-id":
		log.Info("Using user-id login type")

		userIDString, ok := os.LookupEnv("CLOUDNS_USER_ID")
		if !ok {
			return nil, fmt.Errorf("CLOUDNS_USER_ID is not set")
		}

		userIDInt, error := strconv.Atoi(userIDString)
		if error != nil {
			return nil, fmt.Errorf("CLOUDNS_USER_ID is not a valid integer")
		}

		c, error := cloudns.New(
			cloudns.AuthUserID(userIDInt, userPassword),
		)
		if error != nil {
			return nil, fmt.Errorf("error creating ClouDNS client: %s", error)
		}

		client = c
		log.Info("Authenticated with ClouDNS using user-id login type")

	case "sub-user":
		log.Info("Using sub-user login type")

		subUYserIDString, ok := os.LookupEnv("CLOUDNS_SUB_USER_ID")
		if !ok {
			return nil, fmt.Errorf("CLOUDNS_SUB_USER_ID is not set")
		}

		subUserIDInt, error := strconv.Atoi(subUYserIDString)
		if error != nil {
			return nil, fmt.Errorf("CLOUDNS_SUB_USER_ID is not a valid integer")
		}

		c, error := cloudns.New(
			cloudns.AuthSubUserID(subUserIDInt, userPassword),
		)
		if error != nil {
			return nil, fmt.Errorf("error creating ClouDNS client: %s", error)
		}

		client = c
		log.Info("Authenticated with ClouDNS using sub-user login type")

	case "sub-user-name":
		log.Info("Using sub-user-name login type")

		subUserName, ok := os.LookupEnv("CLOUDNS_SUB_USER_NAME")
		if !ok {
			return nil, fmt.Errorf("CLOUDNS_SUB_USER_NAME is not set")
		}

		c, error := cloudns.New(
			cloudns.AuthSubUserName(subUserName, userPassword),
		)
		if error != nil {
			return nil, fmt.Errorf("error creating ClouDNS client: %s", error)
		}

		client = c
		log.Info("Authenticated with ClouDNS using sub-user-name login type")
	}

	provider := &ClouDNSProvider{
		client:       client,
		context:      config.Context,
		domainFilter: config.DomainFilter,
		zoneIDFilter: config.ZoneIDFilter,
		ownerID:      config.OwnerID,
		dryRun:       config.DryRun,
		testing:      config.Testing,
	}

	return provider, nil
}

func (p *ClouDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	log.Info("Getting Records from ClouDNS")

	var endpoints []*endpoint.Endpoint

	zones, err := p.client.Zones.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting zones: %s", err)
	}

	for _, zone := range zones {

		records, err := p.client.Records.List(ctx, zone.Name)
		if err != nil {
			return nil, fmt.Errorf("error getting records: %s", err)
		}

		for _, record := range records {
			if provider.SupportedRecordType(string(record.RecordType)) {
				name := ""

				if record.Host == "" || record.Host == "@" {
					name = zone.Name
				} else {
					name = record.Host + "." + zone.Name
				}

				if record.RecordType == cloudns.RecordTypeTXT {
					if record.Host == "adash" {
						name = "a-" + zone.Name
					}
				}

				endpoints = append(endpoints, endpoint.NewEndpointWithTTL(
					name,
					string(record.RecordType),
					endpoint.TTL(record.TTL),
					record.Record,
				))
			}
		}
	}

	merged := mergeEndpointsByNameType(endpoints)

	out := "Found:"
	for _, e := range merged {
		if e.RecordType != endpoint.RecordTypeTXT {
			out = out + " [" + e.DNSName + " " + e.RecordType + " " + e.Targets[0] + " " + fmt.Sprint(e.RecordTTL) + "]"
		}
	}
	log.Infof(out)

	return merged, nil
}

func (p *ClouDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	infoString := "Creating " + fmt.Sprint(len(changes.Create)) + " Record(s), Updating " + fmt.Sprint(len(changes.UpdateNew)) + " Record(s), Deleting " + fmt.Sprint(len(changes.Delete)) + " Record(s)"

	if len(changes.Create) == 0 && len(changes.Delete) == 0 && len(changes.UpdateNew) == 0 && len(changes.UpdateOld) == 0 {
		log.Info("No Changes")
		return nil
	} else if p.dryRun {
		log.Info("DRY RUN: " + infoString)
	} else {
		log.Info(infoString)
	}

	zones, err := p.client.Zones.List(ctx)
	if err != nil {
		return err
	}

	err = p.createRecords(ctx, zones, changes.Create)
	if err != nil {
		return err
	}

	err = p.deleteRecords(ctx, zones, changes.Delete)
	if err != nil {
		return err
	}

	err = p.updateRecords(ctx, zones, changes.UpdateNew)
	if err != nil {
		return err
	}

	return nil
}

func (p *ClouDNSProvider) createRecords(ctx context.Context, zones []cloudns.Zone, endpoints []*endpoint.Endpoint) error {
	for _, ep := range endpoints {

		dnsParts := strings.Split(ep.DNSName, ".")
		partLength := len(dnsParts)

		if ep.RecordType == "TXT" {
			if !p.dryRun {
				if partLength == 2 && dnsParts[0][0:2] == "a-" {
					_, err := p.client.Records.Create(ctx, dnsParts[0][2:]+"."+dnsParts[1], cloudns.Record{
						Host:       "adash",
						Record:     ep.Targets[0],
						RecordType: cloudns.RecordType("TXT"),
						TTL:        60,
					})
					if err != nil {
						return err
					}
				} else {
					hostName := ""

					zoneName := dnsParts[partLength-2] + "." + dnsParts[partLength-1]
					i := strings.LastIndex(ep.DNSName, "."+zoneName)
					if i != -1 {
						hostName = ep.DNSName[:i] + strings.Replace(ep.DNSName[i:], "."+zoneName, "", 1)
					}

					_, err := p.client.Records.Create(ctx, zoneName, cloudns.Record{
						Host:       hostName,
						Record:     ep.Targets[0],
						RecordType: cloudns.RecordType("TXT"),
						TTL:        60,
					})
					if err != nil {
						return err
					}
				}
				log.Infof("CREATE %s %s %s %s", ep.DNSName, ep.RecordType, ep.Targets[0], fmt.Sprint(ep.RecordTTL))
			} else {
				log.Infof("DRY RUN: CREATE %s %s %s %s", ep.DNSName, ep.RecordType, ep.Targets[0], fmt.Sprint(ep.RecordTTL))
			}
		}

		if !isValidTTL(strconv.Itoa(int(ep.RecordTTL))) && !(ep.RecordType == "TXT") {
			return fmt.Errorf("invalid TTL %d for %s - must be one of '60', '300', '900', '1800', '3600', '21600', '43200', '86400', '172800', '259200', '604800', '1209600', '2592000'", ep.RecordTTL, ep.DNSName)
		}

		if partLength == 2 && !(ep.RecordType == "TXT") {
			for _, target := range ep.Targets {
				if !p.dryRun {
					_, err := p.client.Records.Create(ctx, ep.DNSName, cloudns.Record{
						Host:       "",
						Record:     target,
						RecordType: cloudns.RecordType(ep.RecordType),
						TTL:        int(ep.RecordTTL),
					})
					if err != nil {
						return err
					}

					log.Infof("CREATE %s %s %s %s", ep.DNSName, ep.RecordType, target, fmt.Sprint(ep.RecordTTL))
				} else {
					log.Infof("DRY RUN: CREATE %s %s %s %s", ep.DNSName, ep.RecordType, target, fmt.Sprint(ep.RecordTTL))
				}
			}
		} else if partLength > 2 && !(ep.RecordType == "TXT") {

			zoneName := dnsParts[partLength-2] + "." + dnsParts[partLength-1]
			i := strings.LastIndex(ep.DNSName, "."+zoneName)
			hostName := ep.DNSName[:i] + strings.Replace(ep.DNSName[i:], "."+zoneName, "", 1)

			log.Info(hostName)

			for _, target := range ep.Targets {
				if !p.dryRun {
					_, err := p.client.Records.Create(ctx, zoneName, cloudns.Record{
						Host:       hostName,
						Record:     target,
						RecordType: cloudns.RecordType(ep.RecordType),
						TTL:        int(ep.RecordTTL),
					})
					if err != nil {
						return err
					}

					log.Infof("CREATE %s %s %s %s", ep.DNSName, ep.RecordType, target, fmt.Sprint(ep.RecordTTL))
				} else {
					log.Infof("DRY RUN: CREATE %s %s %s %s", ep.DNSName, ep.RecordType, target, fmt.Sprint(ep.RecordTTL))
				}
			}
		}
	}

	return nil
}

func (p *ClouDNSProvider) deleteRecords(ctx context.Context, zones []cloudns.Zone, endpoints []*endpoint.Endpoint) error {

	for _, endpoint := range endpoints {
		log.Info("Deleting record: ", endpoint.DNSName, " ", endpoint.RecordType, " ", endpoint.Targets[0], " ", fmt.Sprint(endpoint.RecordTTL))

		dnsParts := strings.Split(endpoint.DNSName, ".")
		partLength := len(dnsParts)

		epZoneName := ""
		if partLength == 2 {
			epZoneName = endpoint.DNSName
			log.Info("1")
		} else if partLength > 2 {
			epZoneName = dnsParts[partLength-2] + "." + dnsParts[partLength-1]
			log.Info("2")
		}

		if endpoint.RecordType == "TXT" {
			log.Info("3")
			if epZoneName[0:2] == "a-" {
				log.Info("4")
				recordID, err := p.recordFromTarget(ctx, endpoint, endpoint.Targets[0], epZoneName, "adash")
				if err != nil {
					return err
				}

				if !p.dryRun {
					_, err := p.client.Records.Delete(ctx, epZoneName, recordID)
					if err != nil {
						return err
					}
					log.Infof("DELETE %s %s %s", endpoint.DNSName, endpoint.RecordType, endpoint.Targets[0])
				} else {
					log.Infof("DRY RUN: DELETE %s %s %s", endpoint.DNSName, endpoint.RecordType, endpoint.Targets[0])
				}
			} else {
				log.Info("10")
				epHostName := ""
				i := strings.LastIndex(endpoint.DNSName, "."+epZoneName)
				if i != -1 {
					epHostName = endpoint.DNSName[:i] + strings.Replace(endpoint.DNSName[i:], "."+epZoneName, "", 1)
					log.Info("11")
				} else {
					epHostName = strings.ReplaceAll(endpoint.DNSName, "."+epZoneName, "")
					log.Info("12")
				}

				recordID, err := p.recordFromTarget(ctx, endpoint, endpoint.Targets[0], epZoneName, epHostName)
				if err != nil {
					return err
				}

				if !p.dryRun {
					_, err := p.client.Records.Delete(ctx, epZoneName, recordID)
					if err != nil {
						return err
					}
					log.Infof("DELETE %s %s %s", endpoint.DNSName, endpoint.RecordType, endpoint.Targets[0])
				} else {
					log.Infof("DRY RUN: DELETE %s %s %s", endpoint.DNSName, endpoint.RecordType, endpoint.Targets[0])
				}

			}
		} else {
			log.Info("5")
			epHostName := ""
			i := strings.LastIndex(endpoint.DNSName, "."+epZoneName)
			if i != -1 {
				epHostName = endpoint.DNSName[:i] + strings.Replace(endpoint.DNSName[i:], "."+epZoneName, "", 1)
			} else {
				epHostName = strings.ReplaceAll(endpoint.DNSName, "."+epZoneName, "")
			}

			for _, target := range endpoint.Targets {

				recordID, err := p.recordFromTarget(ctx, endpoint, target, epZoneName, epHostName)
				if err != nil {
					return err
				}

				if !p.dryRun {
					_, err := p.client.Records.Delete(ctx, epZoneName, recordID)
					if err != nil {
						return err
					}
					log.Infof("DELETE %s %s %s", endpoint.DNSName, endpoint.RecordType, endpoint.Targets[0])
				} else {
					log.Infof("DRY RUN: DELETE %s %s %s", endpoint.DNSName, endpoint.RecordType, target)
				}

			}
		}

	}
	return nil
}

func (p *ClouDNSProvider) updateRecords(ctx context.Context, zones []cloudns.Zone, endpoints []*endpoint.Endpoint) error {
	for _, endpoint := range endpoints {
		log.Infof("Creating Record: %s %s %s", endpoint.DNSName, endpoint.RecordType, endpoint.Targets[0])
	}
	return nil
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
		for i, ep := range endpoints {
			targets[i] = ep.Targets[0]
		}

		e := endpoint.NewEndpoint(dnsName, recordType, targets...)
		e.RecordTTL = ttl
		result = append(result, e)
	}

	return result
}

func isValidTTL(ttl string) bool {
	validTTLs := []string{"60", "300", "900", "1800", "3600", "21600", "43200", "86400", "172800", "259200", "604800", "1209600", "2592000"}

	for _, validTTL := range validTTLs {
		if ttl == validTTL {
			return true
		}
	}

	return false
}

func (p *ClouDNSProvider) recordFromTarget(ctx context.Context, endpoint *endpoint.Endpoint, target string, epZoneName string, epHostName string) (int, error) {

	log.Info("epZoneName: " + epZoneName)
	log.Info("epHostName: " + epHostName)
	log.Info("target: " + target)
	log.Info("recordType: " + endpoint.RecordType)
	log.Info("dnsName: " + endpoint.DNSName)
	log.Info("recordTTL: " + strconv.Itoa(int(endpoint.RecordTTL)))

	if epZoneName == epHostName {
		epHostName = ""
	}

	recordsByZone, err := p.zoneRecordMap(ctx)
	if err != nil {
		return 0, err
	}

	if records, ok := recordsByZone[epZoneName]; ok {

		for _, record := range records {

			if record.RecordType == cloudns.RecordTypeTXT {
				if record.Host == "adash" {
					record.Host = ""
				}
			}

			log.Infof("record.Host: %s, epHostName: %s, record.Record: %s, target: %s", record.Host, epHostName, record.Record, target)

			if record.Host == epHostName && record.Record == target {
				return record.ID, nil
			}

		}

	}

	return 0, fmt.Errorf("record not found")
}

func (p *ClouDNSProvider) zoneRecordMap(ctx context.Context) (map[string][]cloudns.Record, error) {

	zoneRecordMap := make(map[string][]cloudns.Record)

	zones, err := p.client.Zones.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		zoneRecordMap[zone.Name] = []cloudns.Record{}

		recordMap, err := p.client.Records.List(ctx, zone.Name)
		if err != nil {
			return nil, err
		}

		for _, record := range recordMap {
			zoneRecordMap[zone.Name] = append(zoneRecordMap[zone.Name], record)
		}
	}

	return zoneRecordMap, nil
}
