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

package akamai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v13/pkg/dns"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v13/pkg/edgegrid"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v13/pkg/session"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/blueprint"
)

const (
	// Default Record TTL
	defaultTTL = 600
	maxUint    = ^uint(0)
	maxInt     = int(maxUint >> 1)
	// zoneListConcurrency bounds the parallel recordset listings issued by Records.
	// Edge DNS enforces per-account API quotas, so this stays low on purpose.
	zoneListConcurrency = 4
)

// AkamaiDNSService is a proxy interface of the Akamai edgegrid dns package that can be stubbed for testing.
// It is a subset of dns.DNS, so a real client satisfies it as is.
type AkamaiDNSService interface {
	ListZones(ctx context.Context, params dns.ListZonesRequest) (*dns.ZoneListResponse, error)
	GetRecordSets(ctx context.Context, params dns.GetRecordSetsRequest) (*dns.GetRecordSetsResponse, error)
	CreateRecordSets(ctx context.Context, params dns.CreateRecordSetsRequest) error
	UpdateRecord(ctx context.Context, params dns.UpdateRecordRequest) error
	DeleteRecord(ctx context.Context, params dns.DeleteRecordRequest) error
}

type AkamaiConfig struct {
	DomainFilter          *endpoint.DomainFilter
	ZoneIDFilter          provider.ZoneIDFilter
	ServiceConsumerDomain string
	ClientToken           string
	ClientSecret          string
	AccessToken           string
	EdgercPath            string
	EdgercSection         string
	MaxBody               int
	AccountKey            string
	ZoneCacheDuration     time.Duration
	DryRun                bool
}

// AkamaiProvider implements the DNS provider for Akamai.
type AkamaiProvider struct {
	provider.BaseProvider
	// Edgedns zones to filter on
	domainFilter *endpoint.DomainFilter
	// Contract Ids to filter on
	zoneIDFilter provider.ZoneIDFilter
	dryRun       bool
	// Defines client. Allows for mocking.
	client AkamaiDNSService
	// Caches the filtered zone list between Records and ApplyChanges.
	zonesCache *blueprint.ZoneCache[[]akamaiZone]
}

type akamaiZones struct {
	Zones []akamaiZone `json:"zones"`
}

type akamaiZone struct {
	ContractID string `json:"contractId"`
	Zone       string `json:"zone"`
}

// New creates an Akamai provider from the given configuration.
func New(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	return newProvider(
		AkamaiConfig{
			DomainFilter:          domainFilter,
			ZoneIDFilter:          provider.NewZoneIDFilter(cfg.ZoneIDFilter),
			ServiceConsumerDomain: cfg.AkamaiServiceConsumerDomain,
			ClientToken:           cfg.AkamaiClientToken,
			ClientSecret:          cfg.AkamaiClientSecret,
			AccessToken:           cfg.AkamaiAccessToken,
			EdgercPath:            cfg.AkamaiEdgercPath,
			EdgercSection:         cfg.AkamaiEdgercSection,
			ZoneCacheDuration:     cfg.AkamaiZoneCacheDuration,
			DryRun:                cfg.DryRun,
		}, nil)
}

// newProvider initializes a new Akamai DNS based Provider.
func newProvider(akamaiConfig AkamaiConfig, akaService AkamaiDNSService) (provider.Provider, error) {
	edgeGridConfig, err := buildEdgegridConfig(akamaiConfig)
	if err != nil {
		return &AkamaiProvider{}, err // return an empty provider for backward compatibility
	}

	p := &AkamaiProvider{
		domainFilter: akamaiConfig.DomainFilter,
		zoneIDFilter: akamaiConfig.ZoneIDFilter,
		dryRun:       akamaiConfig.DryRun,
		zonesCache:   blueprint.NewZoneCache[[]akamaiZone](akamaiConfig.ZoneCacheDuration),
	}

	if akaService != nil {
		log.Debugf("Using STUB")
		p.client = akaService
		return p, nil
	}

	sess, err := session.New(
		session.WithSigner(edgeGridConfig),
		session.WithUserAgent(externaldns.UserAgent()),
		session.WithRequestLimit(edgeGridConfig.RequestLimit),
	)
	if err != nil {
		log.Errorf("Edgegrid session init failed")
		return &AkamaiProvider{}, err
	}
	p.client = dns.Client(sess)

	return p, nil
}

// buildEdgegridConfig resolves the Edgegrid credentials, preferring an explicit external-dns
// configuration and falling back to the .edgerc file and the AKAMAI_* environment.
func buildEdgegridConfig(akamaiConfig AkamaiConfig) (*edgegrid.Config, error) {
	// environment overrides edgerc file but config needs to be complete
	if akamaiConfig.ServiceConsumerDomain == "" || akamaiConfig.ClientToken == "" || akamaiConfig.ClientSecret == "" || akamaiConfig.AccessToken == "" {
		// Kubernetes config incomplete or non existent. Can't mix and match.
		// Look for Akamai environment or .edgerc creds
		cfg, err := edgegrid.New(
			edgegrid.WithFile(akamaiConfig.EdgercPath),
			edgegrid.WithSection(akamaiConfig.EdgercSection),
			edgegrid.WithEnv(true),
		)
		if err != nil {
			log.Errorf("Edgegrid Init Failed")
			return nil, err
		}
		cfg.HeaderToSign = append(cfg.HeaderToSign, "X-External-DNS")
		return cfg, nil
	}

	// Use external-dns config
	cfg := &edgegrid.Config{
		Host:         akamaiConfig.ServiceConsumerDomain,
		ClientToken:  akamaiConfig.ClientToken,
		ClientSecret: akamaiConfig.ClientSecret,
		AccessToken:  akamaiConfig.AccessToken,
		MaxBody:      edgegrid.MaxBodySize,
		HeaderToSign: []string{
			"X-External-DNS",
		},
		Debug: false,
	}
	// Check for edgegrid overrides
	if envval, ok := os.LookupEnv("AKAMAI_MAX_BODY"); ok {
		if i, err := strconv.Atoi(envval); err == nil {
			cfg.MaxBody = i
			log.Debugf("Edgegrid maxbody set to %s", envval)
		}
	}
	if envval, ok := os.LookupEnv("AKAMAI_ACCOUNT_KEY"); ok {
		cfg.AccountKey = envval
		log.Debugf("Edgegrid applying account key %s", envval)
	}
	if envval, ok := os.LookupEnv("AKAMAI_DEBUG"); ok {
		if dbgval, err := strconv.ParseBool(envval); err == nil {
			cfg.Debug = dbgval
			log.Debugf("Edgegrid debug set to %s", envval)
		}
	}

	return cfg, nil
}

// Fetch zones using Edgegrid DNS v2 API
func (p AkamaiProvider) fetchZones(ctx context.Context) (akamaiZones, error) {
	if !p.zonesCache.Expired() {
		cached := p.zonesCache.Get()
		log.Debugf("Fetched '%d' zones from cache", len(cached))
		return akamaiZones{Zones: cached}, nil
	}

	filteredZones := akamaiZones{Zones: make([]akamaiZone, 0)}
	req := dns.ListZonesRequest{Types: "primary", ShowAll: true}
	// filter based on contractIds
	if len(p.zoneIDFilter.ZoneIDs) > 0 {
		req.ContractIDs = strings.Join(p.zoneIDFilter.ZoneIDs, ",")
	}
	resp, err := p.client.ListZones(ctx, req) // retrieve all primary zones filtered by contract ids
	if err != nil {
		log.Errorf("Failed to fetch zones from Akamai")
		return filteredZones, err
	}

	for _, zone := range resp.Zones {
		if p.domainFilter.Match(zone.Zone) {
			filteredZones.Zones = append(filteredZones.Zones, akamaiZone{ContractID: zone.ContractID, Zone: zone.Zone})
			log.Debugf("Fetched zone: '%s' (ZoneID: %s)", zone.Zone, zone.ContractID)
		}
	}
	lenFilteredZones := len(filteredZones.Zones)
	if lenFilteredZones == 0 {
		log.Warnf("No zones could be fetched")
	} else {
		log.Debugf("Fetched '%d' zones from Akamai", lenFilteredZones)
	}
	p.zonesCache.Reset(filteredZones.Zones)

	return filteredZones, nil
}

// Records returns the list of records in a given zone.
func (p AkamaiProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.fetchZones(ctx)
	if err != nil {
		log.Warnf("Failed to identify target zones! Error: %s", err.Error())
		return nil, err
	}

	// One listing per zone, run concurrently but capped: the round trips dominate, the
	// local decoding does not. Results are indexed by zone so the output stays ordered.
	perZone := make([][]*endpoint.Endpoint, len(zones.Zones))
	group, groupCtx := errgroup.WithContext(ctx)
	group.SetLimit(zoneListConcurrency)
	for i, zone := range zones.Zones {
		group.Go(func() error {
			eps, err := p.zoneEndpoints(groupCtx, zone.Zone)
			if err != nil {
				// A zone that cannot be listed is skipped, as it was before: one unreadable
				// zone must not fail the whole sync.
				log.Errorf("Recordsets retrieval for zone: '%s' failed! %s", zone.Zone, err.Error())
				return nil
			}
			perZone[i] = eps
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint
	for _, eps := range perZone {
		endpoints = append(endpoints, eps...)
	}

	lenEndpoints := len(endpoints)
	if lenEndpoints == 0 {
		log.Warnf("No endpoints could be fetched")
	} else {
		log.Debugf("Fetched '%d' endpoints from Akamai", lenEndpoints)
		log.Debugf("Endpoints [%v]", endpoints)
	}

	return endpoints, nil
}

// zoneEndpoints lists a single zone and converts its supported recordsets to endpoints.
func (p AkamaiProvider) zoneEndpoints(ctx context.Context, zone string) ([]*endpoint.Endpoint, error) {
	recordsets, err := p.client.GetRecordSets(ctx, dns.GetRecordSetsRequest{
		Zone:      zone,
		QueryArgs: &dns.RecordSetQueryArgs{ShowAll: true},
	})
	if err != nil {
		return nil, err
	}
	if len(recordsets.RecordSets) == 0 {
		log.Warnf("Zone %s contains no recordsets", zone)
		return nil, nil
	}

	var endpoints []*endpoint.Endpoint
	for _, recordset := range recordsets.RecordSets {
		if !provider.SupportedRecordType(recordset.Type) {
			log.Debugf("Skipping endpoint DNSName: '%s' RecordType: '%s'. Record type not supported.", recordset.Name, recordset.Type)
			continue
		}
		if !p.domainFilter.Match(recordset.Name) {
			log.Debugf("Skipping endpoint. Record name %s doesn't match containing zone %s.", recordset.Name, zone)
			continue
		}
		endpoints = append(endpoints, endpoint.NewEndpointWithTTL(recordset.Name,
			recordset.Type,
			endpoint.TTL(recordset.TTL),
			trimTxtRdata(recordset.Rdata, recordset.Type)...))
		log.Debugf("Fetched endpoint DNSName: '%s' RecordType: '%s' Rdata: '%s')", recordset.Name, recordset.Type, recordset.Rdata)
	}

	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p AkamaiProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// Nothing to do means nothing to fetch: in steady state this is the common path and
	// it must not cost a full zone listing.
	if !changes.HasChanges() {
		log.Debug("No changes to apply")
		return nil
	}

	zoneNameIDMapper := provider.ZoneIDName{}
	zones, err := p.fetchZones(ctx)
	if err != nil {
		log.Errorf("Failed to fetch zones from Akamai")
		return err
	}

	for _, z := range zones.Zones {
		zoneNameIDMapper[z.Zone] = z.Zone
	}
	log.Debugf("Processing zones: [%v]", zoneNameIDMapper)

	// Create recordsets
	log.Debugf("Create Changes requested [%v]", changes.Create)
	if err := p.createRecordsets(ctx, zoneNameIDMapper, changes.Create); err != nil {
		return err
	}
	// Delete recordsets
	log.Debugf("Delete Changes requested [%v]", changes.Delete)
	if err := p.deleteRecordsets(ctx, zoneNameIDMapper, changes.Delete); err != nil {
		return err
	}
	// Update recordsets
	log.Debugf("Update Changes requested [%v]", changes.UpdateNew)
	if err := p.updateNewRecordsets(ctx, zoneNameIDMapper, changes.UpdateNew); err != nil {
		return err
	}
	// Check that all old endpoints were accounted for
	revRecs := changes.Delete
	revRecs = append(revRecs, changes.UpdateNew...)
	for _, rec := range changes.UpdateOld {
		found := false
		for _, r := range revRecs {
			if rec.DNSName == r.DNSName {
				found = true
				break
			}
		}
		if !found {
			log.Warnf("UpdateOld endpoint '%s' is not accounted for in UpdateNew|Delete endpoint list", rec.DNSName)
		}
	}

	return nil
}

// Create DNS Recordset
func newAkamaiRecordset(dnsName, recordType string, ttl int, targets []string) dns.RecordSet {
	return dns.RecordSet{
		Name:  strings.TrimSuffix(dnsName, "."),
		Rdata: targets,
		Type:  recordType,
		TTL:   ttl,
	}
}

// cleanTargets preps recordset rdata if necessary for EdgeDNS
func cleanTargets(rtype string, targets ...string) []string {
	log.Debugf("Targets to clean: [%v]", targets)
	switch rtype {
	case "CNAME", "SRV":
		for idx, target := range targets {
			targets[idx] = strings.TrimSuffix(target, ".")
		}
	case "TXT":
		for idx, target := range targets {
			log.Debugf("TXT data to clean: [%s]", target)
			// need to embed text data in quotes. Make sure not piling on
			target = strings.Trim(target, "\"")
			// bug in DNS API with embedded quotes.
			if strings.Contains(target, "owner") && strings.Contains(target, "\"") {
				target = strings.ReplaceAll(target, "\"", "`")
			}
			targets[idx] = "\"" + target + "\""
		}
	}
	log.Debugf("Clean targets: [%v]", targets)

	return targets
}

// trimTxtRdata removes surrounding quotes for received TXT rdata
func trimTxtRdata(rdata []string, rtype string) []string {
	if rtype == "TXT" {
		for idx, d := range rdata {
			if strings.Contains(d, "`") {
				rdata[idx] = strings.ReplaceAll(d, "`", "\"")
			}
		}
	}
	log.Debugf("Trimmed data: [%v]", rdata)

	return rdata
}

func ttlAsInt(src endpoint.TTL) int {
	var temp any = int64(src)
	temp64 := temp.(int64)
	var ttl = defaultTTL
	if temp64 > 0 && temp64 <= int64(maxInt) {
		ttl = int(temp64)
	}

	return ttl
}

// Create Endpoint Recordsets
func (p AkamaiProvider) createRecordsets(ctx context.Context, zoneNameIDMapper provider.ZoneIDName, endpoints []*endpoint.Endpoint) error {
	if len(endpoints) == 0 {
		log.Info("No endpoints to create")
		return nil
	}

	endpointsByZone := edgeChangesByZone(zoneNameIDMapper, endpoints)

	// create all recordsets by zone
	for zone, endpoints := range endpointsByZone {
		recordsets := &dns.RecordSets{RecordSets: make([]dns.RecordSet, 0, len(endpoints))}
		for _, endpoint := range endpoints {
			newrec := newAkamaiRecordset(endpoint.DNSName,
				endpoint.RecordType,
				ttlAsInt(endpoint.RecordTTL),
				cleanTargets(endpoint.RecordType, endpoint.Targets...))
			logfields := log.Fields{
				"record": newrec.Name,
				"type":   newrec.Type,
				"ttl":    newrec.TTL,
				"target": fmt.Sprintf("%v", newrec.Rdata),
				"zone":   zone,
			}
			log.WithFields(logfields).Info("Creating recordsets")
			recordsets.RecordSets = append(recordsets.RecordSets, newrec)
		}

		if p.dryRun {
			continue
		}
		// Create recordsets all at once
		err := p.client.CreateRecordSets(ctx, dns.CreateRecordSetsRequest{
			RecordSets: recordsets,
			Zone:       zone,
			RecLock:    []bool{true},
		})
		if err != nil {
			log.Errorf("Failed to create endpoints for DNS zone %s. Error: %s", zone, err.Error())
			return err
		}
	}

	return nil
}

func (p AkamaiProvider) deleteRecordsets(ctx context.Context, zoneNameIDMapper provider.ZoneIDName, endpoints []*endpoint.Endpoint) error {
	for _, endpoint := range endpoints {
		zoneName, _ := zoneNameIDMapper.FindZone(endpoint.DNSName)
		if zoneName == "" {
			log.Debugf("Skipping Akamai Edge DNS endpoint deletion: '%s' type: '%s', it does not match against Domain filters", endpoint.DNSName, endpoint.RecordType)
			continue
		}
		log.Infof("Akamai Edge DNS recordset deletion- Zone: '%s', DNSName: '%s', RecordType: '%s', Targets: '%+v'", zoneName, endpoint.DNSName, endpoint.RecordType, endpoint.Targets)

		if p.dryRun {
			continue
		}

		// DELETE is already addressed by name and type, so the record does not have to be
		// read back first. A missing record is the outcome we wanted anyway.
		err := p.client.DeleteRecord(ctx, dns.DeleteRecordRequest{
			Zone:       zoneName,
			Name:       strings.TrimSuffix(endpoint.DNSName, "."),
			RecordType: endpoint.RecordType,
			RecLock:    []bool{true},
		})
		if err != nil {
			if isNotFound(err) {
				log.Infof("Endpoint deletion. Record doesn't exist. Name: %s, Type: %s", endpoint.DNSName, endpoint.RecordType)
				continue
			}
			log.Errorf("edge dns recordset deletion failed. error: %s", err.Error())
			return err
		}
	}

	return nil
}

// Update endpoint recordsets
func (p AkamaiProvider) updateNewRecordsets(ctx context.Context, zoneNameIDMapper provider.ZoneIDName, endpoints []*endpoint.Endpoint) error {
	for _, endpoint := range endpoints {
		zoneName, _ := zoneNameIDMapper.FindZone(endpoint.DNSName)
		if zoneName == "" {
			log.Debugf("Skipping Akamai Edge DNS endpoint update: '%s' type: '%s', it does not match against Domain filters", endpoint.DNSName, endpoint.RecordType)
			continue
		}
		log.Infof("Akamai Edge DNS recordset update - Zone: '%s', DNSName: '%s', RecordType: '%s', Targets: '%+v'", zoneName, endpoint.DNSName, endpoint.RecordType, endpoint.Targets)

		if p.dryRun {
			continue
		}

		// PUT replaces the recordset addressed by name and type, and the desired TTL and
		// rdata both come from the endpoint, so reading the record back first buys nothing.
		ttl := ttlAsInt(endpoint.RecordTTL)
		err := p.client.UpdateRecord(ctx, dns.UpdateRecordRequest{
			Record: &dns.RecordBody{
				Name:       strings.TrimSuffix(endpoint.DNSName, "."),
				RecordType: endpoint.RecordType,
				TTL:        &ttl,
				Target:     cleanTargets(endpoint.RecordType, endpoint.Targets...),
			},
			Zone:    zoneName,
			RecLock: []bool{true},
		})
		if err != nil {
			log.Errorf("Akamai Edge DNS recordset update failed. Error: %s", err.Error())
			return err
		}
	}

	return nil
}

// isNotFound reports whether the Edge DNS API answered with a 404.
func isNotFound(err error) bool {
	apiErr := &dns.Error{}
	return errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound
}

// edgeChangesByZone separates a multi-zone change into a single change per zone.
func edgeChangesByZone(zoneMap provider.ZoneIDName, endpoints []*endpoint.Endpoint) map[string][]*endpoint.Endpoint {
	createsByZone := make(map[string][]*endpoint.Endpoint, len(zoneMap))
	for _, z := range zoneMap {
		createsByZone[z] = make([]*endpoint.Endpoint, 0)
	}
	for _, ep := range endpoints {
		zone, _ := zoneMap.FindZone(ep.DNSName)
		if zone != "" {
			createsByZone[zone] = append(createsByZone[zone], ep)
			continue
		}
		log.Debugf("Skipping Akamai Edge DNS creation of endpoint: '%s' type: '%s', it does not match against Domain filters", ep.DNSName, ep.RecordType)
	}

	return createsByZone
}
