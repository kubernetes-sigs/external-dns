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

package provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	udnssdk "github.com/ultradns/ultradns-sdk-go"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const (
	ultradnsDefaultTTL = 198
	ultradnsCreate     = "CREATE"
	ultradnsDelete     = "DELETE"
	ultradnsUpdate     = "UPDATE"
	sbPoolPriority     = 1
	sbPoolOrder        = "ROUND_ROBIN"
	rdPoolOrder        = "ROUND_ROBIN"
)

// global variables
var sbPoolRunProbes = true
var sbPoolActOnProbes = true
var ultradnsPoolType = "rdpool"

//Setting custom headers for ultradns api calls
var customHeader = []udnssdk.CustomHeader{
	udnssdk.CustomHeader {
		Key:   "UltraClient",
		Value: "kube-client",
	},
}

type UltraDNSProvider struct {
	client udnssdk.Client

	domainFilter endpoint.DomainFilter
	DryRun       bool
	AccountName  string
}

type UltraDNSChanges struct {
	Action string

	ResourceRecordSetUltraDNS udnssdk.RRSet
}

// NewUltraDNSProvider initializes a new UltraDNS DNS based provider
func NewUltraDNSProvider(domainFilter endpoint.DomainFilter, dryRun bool) (*UltraDNSProvider, error) {
	Username, ok := os.LookupEnv("ULTRADNS_USERNAME")
	udnssdk.SetCustomHeader = customHeader
	if !ok {
		return nil, fmt.Errorf("no username found")
	}

	Base64Password, ok := os.LookupEnv("ULTRADNS_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("no password found")
	}

	// Base64 Standard Decoding
	Password, err := base64.StdEncoding.DecodeString(Base64Password)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return nil, err
	}

	BaseURL, ok := os.LookupEnv("ULTRADNS_BASEURL")
	if !ok {
		return nil, fmt.Errorf("no baseurl found")
	}
	AccountName, ok := os.LookupEnv("ULTRADNS_ACCOUNTNAME")
	if !ok {
		AccountName = ""
	}

	probeValue, ok := os.LookupEnv("ULTRADNS_ENABLE_PROBING")
	if ok {
		if (probeValue != "true") && (probeValue != "false") {
			return nil, fmt.Errorf("please set proper probe value, the values can be either true or false")
		} else {
			sbPoolRunProbes, _ = strconv.ParseBool(probeValue)
		}
	}

	actOnProbeValue, ok := os.LookupEnv("ULTRADNS_ENABLE_ACTONPROBE")
	if ok {
		if (actOnProbeValue != "true") && (actOnProbeValue != "false") {
			return nil, fmt.Errorf("please set proper act on probe value, the values can be either true or false")
		} else {
			sbPoolActOnProbes, _ = strconv.ParseBool(actOnProbeValue)
		}
	}

	poolValue, ok := os.LookupEnv("ULTRADNS_POOL_TYPE")
	if ok {
		if (poolValue != "sbpool") && (poolValue != "rdpool") {
			return nil, fmt.Errorf(" please set proper ULTRADNS_POOL_TYPE, supported types are sbpool or rdpool")
		}
		ultradnsPoolType = poolValue
	}

	client, err := udnssdk.NewClient(Username, string(Password), BaseURL)
	if err != nil {

		return nil, fmt.Errorf("Connection cannot be established")
	}

	provider := &UltraDNSProvider{
		client:       *client,
		domainFilter: domainFilter,
		DryRun:       dryRun,
		AccountName:  AccountName,
	}

	return provider, nil
}

// Zones returns list of hosted zones
func (p *UltraDNSProvider) Zones(ctx context.Context) ([]udnssdk.Zone, error) {
	zoneKey := &udnssdk.ZoneKey{}
	if p.AccountName != "" {
		zoneKey = &udnssdk.ZoneKey{
			Zone:        "",
			AccountName: p.AccountName,
		}
	}

	zones, err := p.fetchZones(ctx, zoneKey)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

func (p *UltraDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		log.Infof("zones : %v", zone)
		rrsetType := ""
		ownerName := ""
		rrsetKey := udnssdk.RRSetKey{
			Zone: zone.Properties.Name,
			Type: rrsetType,
			Name: ownerName,
		}

		if zone.Properties.ResourceRecordCount != 0 {
			records, err := p.fetchRecords(ctx, rrsetKey)
			if err != nil {
				return nil, err
			}

			for _, r := range records {
				recordTypeArray := strings.Fields(r.RRType)
				if supportedRecordType(recordTypeArray[0]) {
					log.Infof("owner name %s", r.OwnerName)
					name := fmt.Sprintf("%s", r.OwnerName)

					// root name is identified by the empty string and should be
					// translated to zone name for the endpoint entry.
					if r.OwnerName == "" {
						name = zone.Properties.Name
					}

					endPointTTL := endpoint.NewEndpointWithTTL(name, recordTypeArray[0], endpoint.TTL(r.TTL), r.RData...)
					endpoints = append(endpoints, endPointTTL)
				}
			}
		}

	}
	log.Infof("endpoints %v", endpoints)
	return endpoints, nil
}

func (p *UltraDNSProvider) fetchRecords(ctx context.Context, k udnssdk.RRSetKey) ([]udnssdk.RRSet, error) {
	// TODO: Sane Configuration for timeouts / retries
	maxerrs := 5
	waittime := 5 * time.Second

	rrsets := []udnssdk.RRSet{}
	errcnt := 0
	offset := 0
	limit := 1000

	for {
		reqRrsets, ri, res, err := p.client.RRSets.SelectWithOffsetWithLimit(k, offset, limit)
		if err != nil {
			if res != nil && res.StatusCode >= 500 {
				errcnt = errcnt + 1
				if errcnt < maxerrs {
					time.Sleep(waittime)
					continue
				}
			}
			return rrsets, err
		}

		for _, rrset := range reqRrsets {
			rrsets = append(rrsets, rrset)
		}
		if ri.ReturnedCount+ri.Offset >= ri.TotalCount {
			return rrsets, nil
		}
		offset = ri.ReturnedCount + ri.Offset
		continue
	}
}

func (p *UltraDNSProvider) fetchZones(ctx context.Context, zoneKey *udnssdk.ZoneKey) ([]udnssdk.Zone, error) {
	// Select will list the zone rrsets, paginating through all available results
	// TODO: Sane Configuration for timeouts / retries
	maxerrs := 5
	waittime := 5 * time.Second

	zones := []udnssdk.Zone{}

	errcnt := 0
	offset := 0
	limit := 1000

	for {
		reqZones, ri, res, err := p.client.Zone.SelectWithOffsetWithLimit(zoneKey, offset, limit)
		if err != nil {
			if res != nil && res.StatusCode >= 500 {
				errcnt = errcnt + 1
				if errcnt < maxerrs {
					time.Sleep(waittime)
					continue
				}
			}
			return zones, err
		}

		for _, zone := range reqZones {

			if p.domainFilter.IsConfigured() {
				if p.domainFilter.Match(zone.Properties.Name) {
					zones = append(zones, zone)
				}
			} else {
				zones = append(zones, zone)
			}
		}
		if ri.ReturnedCount+ri.Offset >= ri.TotalCount {
			return zones, nil
		}
		offset = ri.ReturnedCount + ri.Offset
		continue
	}
}

func (p *UltraDNSProvider) submitChanges(ctx context.Context, changes []*UltraDNSChanges) error {
	cnameownerName := "cname"
	txtownerName := "txt"
	if len(changes) == 0 {
		log.Infof("All records are already up to date")
		return nil
	}

	zones, err := p.Zones(ctx)
	if err != nil {
		return err
	}
	zoneChanges := seperateChangeByZone(zones, changes)

	for zoneName, changes := range zoneChanges {

		for _, change := range changes {

			if change.ResourceRecordSetUltraDNS.RRType == "CNAME" {
				cnameownerName = change.ResourceRecordSetUltraDNS.OwnerName
			} else if change.ResourceRecordSetUltraDNS.RRType == "TXT" {
				txtownerName = change.ResourceRecordSetUltraDNS.OwnerName
			}

			if cnameownerName == txtownerName {
				rrsetKey := udnssdk.RRSetKey{
					Zone: zoneName,
					Type: endpoint.RecordTypeCNAME,
					Name: change.ResourceRecordSetUltraDNS.OwnerName,
				}
				err := p.getSpecificRecord(ctx, rrsetKey)
				if err != nil {
					return err
				}
				if p.DryRun != true {
					_, err = p.client.RRSets.Delete(rrsetKey)
					if err != nil {
						return err
					}
				}
				return fmt.Errorf("The CNAME and TXT Record name cannot be same please recreate external-dns with - --txt-prefix=")
			}
			rrsetKey := udnssdk.RRSetKey{
				Zone: zoneName,
				Type: change.ResourceRecordSetUltraDNS.RRType,
				Name: change.ResourceRecordSetUltraDNS.OwnerName,
			}
			record := udnssdk.RRSet{}
			if ((change.ResourceRecordSetUltraDNS.RRType == "A" || change.ResourceRecordSetUltraDNS.RRType == "AAAA" ) && (len(change.ResourceRecordSetUltraDNS.RData) >= 2)) {
				if ultradnsPoolType == "sbpool" && change.ResourceRecordSetUltraDNS.RRType == "A" {
					sbPoolObject, _ := p.newSBPoolObjectCreation(ctx, change)
					record = udnssdk.RRSet{
						RRType:    change.ResourceRecordSetUltraDNS.RRType,
						OwnerName: change.ResourceRecordSetUltraDNS.OwnerName,
						RData:     change.ResourceRecordSetUltraDNS.RData,
						TTL:       change.ResourceRecordSetUltraDNS.TTL,
						Profile:   sbPoolObject.RawProfile(),
					}
				} else if ultradnsPoolType == "rdpool" {
					rdPoolObject, _ := p.newRDPoolObjectCreation(ctx, change)
					record = udnssdk.RRSet{
						RRType:    change.ResourceRecordSetUltraDNS.RRType,
						OwnerName: change.ResourceRecordSetUltraDNS.OwnerName,
						RData:     change.ResourceRecordSetUltraDNS.RData,
						TTL:       change.ResourceRecordSetUltraDNS.TTL,
						Profile:   rdPoolObject.RawProfile(),
					}
				}else{
					return fmt.Errorf("We do not support Multiple target AAAA records in SB Pool please contact to Neustar for further details")
				}
			}else {
				record = udnssdk.RRSet{
					RRType:    change.ResourceRecordSetUltraDNS.RRType,
					OwnerName: change.ResourceRecordSetUltraDNS.OwnerName,
					RData:     change.ResourceRecordSetUltraDNS.RData,
					TTL:       change.ResourceRecordSetUltraDNS.TTL,
				}
			}

			log.WithFields(log.Fields{
				"record":  record.OwnerName,
				"type":    record.RRType,
				"ttl":     record.TTL,
				"action":  change.Action,
				"zone":    zoneName,
				"profile": record.Profile,
			}).Info("Changing record.")

			switch change.Action {
			case ultradnsCreate:
				if p.DryRun != true {
					res, err := p.client.RRSets.Create(rrsetKey, record)
					_ = res
					if err != nil {
						return err
					}
				}

			case ultradnsDelete:
				err := p.getSpecificRecord(ctx, rrsetKey)
				if err != nil {
					return err
				}

				if p.DryRun != true {
					_, err = p.client.RRSets.Delete(rrsetKey)
					if err != nil {
						return err
					}
				}
			case ultradnsUpdate:
				err := p.getSpecificRecord(ctx, rrsetKey)
				if err != nil {
					return err
				}

				if p.DryRun != true {
					_, err = p.client.RRSets.Update(rrsetKey, record)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (p *UltraDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	combinedChanges := make([]*UltraDNSChanges, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))
	log.Infof("value of changes %v,%v,%v", changes.Create, changes.UpdateNew, changes.Delete)
	combinedChanges = append(combinedChanges, newUltraDNSChanges(ultradnsCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newUltraDNSChanges(ultradnsUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newUltraDNSChanges(ultradnsDelete, changes.Delete)...)

	return p.submitChanges(ctx, combinedChanges)
}

func newUltraDNSChanges(action string, endpoints []*endpoint.Endpoint) []*UltraDNSChanges {
	changes := make([]*UltraDNSChanges, 0, len(endpoints))
	ttl := ultradnsDefaultTTL
	for _, e := range endpoints {

		if e.RecordTTL.IsConfigured() {
			ttl = int(e.RecordTTL)
		}

		// Adding suffix dot to the record name
		recordName := fmt.Sprintf("%s.", e.DNSName)
		change := &UltraDNSChanges{
			Action: action,
			ResourceRecordSetUltraDNS: udnssdk.RRSet{
				RRType:    e.RecordType,
				OwnerName: recordName,
				RData:     e.Targets,
				TTL:       ttl,
			},
		}
		changes = append(changes, change)
	}
	return changes
}

func seperateChangeByZone(zones []udnssdk.Zone, changes []*UltraDNSChanges) map[string][]*UltraDNSChanges {
	change := make(map[string][]*UltraDNSChanges)
	zoneNameID := zoneIDName{}
	for _, z := range zones {
		zoneNameID.Add(z.Properties.Name, z.Properties.Name)
		change[z.Properties.Name] = []*UltraDNSChanges{}
	}

	for _, c := range changes {
		zone, _ := zoneNameID.FindZone(c.ResourceRecordSetUltraDNS.OwnerName)
		if zone == "" {
			log.Infof("Skipping record %s because no hosted zone matching record DNS Name was detected", c.ResourceRecordSetUltraDNS.OwnerName)
			continue
		}
		change[zone] = append(change[zone], c)

	}
	return change
}

func (p *UltraDNSProvider) getSpecificRecord(ctx context.Context, rrsetKey udnssdk.RRSetKey) (err error) {
	_, err = p.client.RRSets.Select(rrsetKey)
	if err != nil {
		return fmt.Errorf("no record was found for %v", rrsetKey)
	} else {
		return nil
	}
}

// Creation of SBPoolObject
func (p *UltraDNSProvider) newSBPoolObjectCreation(ctx context.Context, change *UltraDNSChanges) (sbPool udnssdk.SBPoolProfile, err error) {

	sbpoolRDataList := []udnssdk.SBRDataInfo{}
	for _, _ = range change.ResourceRecordSetUltraDNS.RData {

		rrdataInfo := udnssdk.SBRDataInfo{
			RunProbes: sbPoolRunProbes,
			Priority:  sbPoolPriority,
			State:     "NORMAL",
			Threshold: 1,
			Weight:    nil,
		}
		sbpoolRDataList = append(sbpoolRDataList, rrdataInfo)
	}
	sbPoolObject := udnssdk.SBPoolProfile{
		Context:     udnssdk.SBPoolSchema,
		Order:       sbPoolOrder,
		Description: change.ResourceRecordSetUltraDNS.OwnerName,
		MaxActive:   len(change.ResourceRecordSetUltraDNS.RData),
		MaxServed:   len(change.ResourceRecordSetUltraDNS.RData),
		RDataInfo:   sbpoolRDataList,
		RunProbes:   sbPoolRunProbes,
		ActOnProbes: sbPoolActOnProbes,
	}
	return sbPoolObject, nil
}

//Creation of RDPoolObject
func (p *UltraDNSProvider) newRDPoolObjectCreation(ctx context.Context, change *UltraDNSChanges) (rdPool udnssdk.RDPoolProfile, err error) {

	rdPoolObject := udnssdk.RDPoolProfile{
		Context:     udnssdk.RDPoolSchema,
		Order:       rdPoolOrder,
		Description: change.ResourceRecordSetUltraDNS.OwnerName,
	}
	return rdPoolObject, nil
}
