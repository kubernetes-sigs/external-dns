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

package scaleway

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	domain "github.com/scaleway/scaleway-sdk-go/api/domain/v2beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/blueprint"
)

const (
	defaultTTL              uint32 = 300
	scalewayDefaultPriority uint32 = 0
	scalewayPriorityKey     string = "scw/priority"
	// zonesCacheDuration is the TTL of the zone list. Zones is called several
	// times per reconciliation loop (Records, AdjustEndpoints and ApplyChanges),
	// this keeps them down to a single API call while staying well below the
	// default sync interval so new zones are still picked up quickly.
	zonesCacheDuration = 30 * time.Second
)

// ScalewayProvider implements the DNS provider for Scaleway DNS
type ScalewayProvider struct {
	provider.BaseProvider
	domainAPI DomainAPI
	dryRun    bool
	// only consider hosted zones managing domains ending in this suffix
	domainFilter *endpoint.DomainFilter
	// zones from the last listing, shared by Records, AdjustEndpoints and ApplyChanges
	zonesCache *blueprint.ZoneCache[[]*domain.DNSZone]
}

// ScalewayChange differentiates between ChangActions
type ScalewayChange struct {
	Action string
	Record []domain.Record
}

// New creates a Scaleway provider from the given configuration.
func New(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	return newProvider(domainFilter, cfg.DryRun)
}

// newProvider initializes a new Scaleway DNS provider
func newProvider(domainFilter *endpoint.DomainFilter, dryRun bool) (*ScalewayProvider, error) {
	var err error
	defaultPageSize := uint64(1000)
	if envPageSize, ok := os.LookupEnv("SCW_DEFAULT_PAGE_SIZE"); ok {
		defaultPageSize, err = strconv.ParseUint(envPageSize, 10, 32)
		if err != nil {
			log.Infof("Ignoring default page size %s, defaulting to 1000", envPageSize)
			defaultPageSize = 1000
		}
	}

	p := &scw.Profile{}
	c, err := scw.LoadConfig()
	if err != nil {
		log.Warnf("Cannot load config: %v", err)
	} else {
		p, err = c.GetActiveProfile()
		if err != nil {
			log.Warnf("Cannot get active profile: %v", err)
		}
	}

	scwClient, err := scw.NewClient(
		scw.WithProfile(p),
		scw.WithEnv(),
		scw.WithUserAgent(externaldns.UserAgent()),
		scw.WithDefaultPageSize(uint32(defaultPageSize)),
	)
	if err != nil {
		return nil, err
	}

	if _, ok := scwClient.GetAccessKey(); !ok {
		return nil, fmt.Errorf("access key no set")
	}

	if _, ok := scwClient.GetSecretKey(); !ok {
		return nil, fmt.Errorf("secret key no set")
	}

	domainAPI := domain.NewAPI(scwClient)

	return &ScalewayProvider{
		domainAPI:    domainAPI,
		dryRun:       dryRun,
		domainFilter: domainFilter,
		zonesCache:   blueprint.NewZoneCache[[]*domain.DNSZone](zonesCacheDuration),
	}, nil
}

// AdjustEndpoints is used to normalize the endpoints
func (p *ScalewayProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	apexNames := p.apexNames(context.Background())
	eps := make([]*endpoint.Endpoint, len(endpoints))
	for i := range endpoints {
		eps[i] = endpoints[i]
		if !eps[i].RecordTTL.IsConfigured() {
			eps[i].RecordTTL = endpoint.TTL(defaultTTL)
		}
		if _, ok := eps[i].GetProviderSpecificProperty(scalewayPriorityKey); !ok {
			eps[i] = eps[i].WithProviderSpecific(scalewayPriorityKey, fmt.Sprintf("%d", scalewayDefaultPriority))
		}
		adjustAliasProperty(eps[i], apexNames)
	}
	return eps, nil
}

// apexNames returns the apex name of every zone handled by the provider.
func (p *ScalewayProvider) apexNames(ctx context.Context) map[string]struct{} {
	zones, err := p.Zones(ctx)
	if err != nil {
		log.Errorf("Failed listing zones, CNAME endpoints at a zone apex are treated as regular CNAME records: %v", err)
		return nil
	}
	names := make(map[string]struct{}, len(zones))
	for _, zone := range zones {
		names[getCompleteZoneName(zone)] = struct{}{}
	}
	return names
}

// adjustAliasProperty normalizes the alias property: "true" on CNAME endpoints
// stored as ALIAS records (zone apex or alias annotation), absent otherwise.
func adjustAliasProperty(ep *endpoint.Endpoint, apexNames map[string]struct{}) {
	_, atApex := apexNames[ep.DNSName]
	if isAliasCNAME(ep, atApex) {
		ep.WithAliasProperty(endpoint.AliasTrue)
	} else {
		ep.DeleteProviderSpecificProperty(endpoint.ProviderSpecificAlias)
	}
}

// isAliasCNAME reports whether a CNAME endpoint has to be stored as a Scaleway
// ALIAS record: either it sits at the zone apex, where Scaleway rejects CNAME
// records, or it opts in through the alias property.
func isAliasCNAME(ep *endpoint.Endpoint, atApex bool) bool {
	return ep.RecordType == endpoint.RecordTypeCNAME &&
		(atApex || ep.GetAliasProperty() == endpoint.AliasTrue)
}

// Zones returns the list of hosted zones.
func (p *ScalewayProvider) Zones(ctx context.Context) ([]*domain.DNSZone, error) {
	if !p.zonesCache.Expired() {
		return p.zonesCache.Get(), nil
	}

	res := []*domain.DNSZone{}

	dnsZones, err := p.domainAPI.ListDNSZones(&domain.ListDNSZonesRequest{}, scw.WithAllPages(), scw.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	for _, dnsZone := range dnsZones.DNSZones {
		if p.domainFilter.Match(getCompleteZoneName(dnsZone)) {
			res = append(res, dnsZone)
		}
	}

	p.zonesCache.Reset(res)

	return res, nil
}

// Records returns the list of records in a given zone.
func (p *ScalewayProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints := map[string]*endpoint.Endpoint{}
	dnsZones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	for _, zone := range dnsZones {
		recordsResp, err := p.domainAPI.ListDNSZoneRecords(&domain.ListDNSZoneRecordsRequest{
			DNSZone: getCompleteZoneName(zone),
		}, scw.WithAllPages())
		if err != nil {
			return nil, err
		}

		for _, record := range recordsResp.Records {
			name := record.Name + "."

			// trim any leading or ending dot
			fullRecordName := strings.Trim(name+getCompleteZoneName(zone), ".")

			recordType := record.Type.String()
			// ExternalDNS writes ALIAS records for CNAME endpoints, read them back as CNAME
			isAliasRecord := record.Type == domain.RecordTypeALIAS
			if isAliasRecord {
				recordType = endpoint.RecordTypeCNAME
			}

			if !provider.SupportedRecordType(recordType) {
				log.Infof("Skipping record %s because type %s is not supported", fullRecordName, recordType)
				continue
			}

			// in external DNS, same endpoint have the same ttl and same priority
			// it's not the case in Scaleway DNS. It should never happen, but if
			// the record is modified without going through ExternalDNS, we could have
			// different priorities of ttls for a same name.
			// In this case, we juste take the first one.
			// The key uses the Scaleway record type so that an ALIAS and a CNAME
			// record sharing a name never merge into a single endpoint.
			mapKey := record.Type.String() + "/" + fullRecordName
			if existingEndpoint, ok := endpoints[mapKey]; ok {
				existingEndpoint.Targets = append(existingEndpoint.Targets, record.Data)
				log.Infof("Appending target %s to record %s, using TTL and priority of target %s", record.Data, fullRecordName, existingEndpoint.Targets[0])
			} else {
				ep := endpoint.NewEndpointWithTTL(fullRecordName, recordType, endpoint.TTL(record.TTL), record.Data)
				ep = ep.WithProviderSpecific(scalewayPriorityKey, fmt.Sprintf("%d", record.Priority))
				if isAliasRecord {
					ep = ep.WithAliasProperty(endpoint.AliasTrue)
				}
				endpoints[mapKey] = ep
			}
		}
	}
	returnedEndpoints := []*endpoint.Endpoint{}
	for _, ep := range endpoints {
		returnedEndpoints = append(returnedEndpoints, ep)
	}

	return returnedEndpoints, nil
}

// ApplyChanges applies a set of changes in a zone.
func (p *ScalewayProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	requests, err := p.generateApplyRequests(ctx, changes)
	if err != nil {
		return err
	}
	for _, req := range requests {
		logChanges(req)
		if p.dryRun {
			log.Info("Running in dry run mode")
			continue
		}
		_, err := p.domainAPI.UpdateDNSZoneRecords(req, scw.WithContext(ctx))
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *ScalewayProvider) generateApplyRequests(ctx context.Context, changes *plan.Changes) ([]*domain.UpdateDNSZoneRecordsRequest, error) {
	returnedRequests := []*domain.UpdateDNSZoneRecordsRequest{}
	recordsToAdd := map[string]*domain.RecordChangeAdd{}
	recordsToDelete := map[string][]*domain.RecordChange{}

	dnsZones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	zoneNameMapper := provider.ZoneIDName{}
	for _, zone := range dnsZones {
		zoneName := getCompleteZoneName(zone)
		zoneNameMapper.Add(zoneName, zoneName)
		recordsToAdd[zoneName] = &domain.RecordChangeAdd{
			Records: []*domain.Record{},
		}
		recordsToDelete[zoneName] = []*domain.RecordChange{}
	}

	log.Debugf("Following records present in updateOld")
	for _, c := range changes.UpdateOld {
		zone, _ := zoneNameMapper.FindZone(c.DNSName)
		if zone == "" {
			log.Infof("Ignore record %s since it's not handled by ExternalDNS", c.DNSName)
			continue
		}
		recordsToDelete[zone] = append(recordsToDelete[zone], endpointToScalewayRecordsChangeDelete(zone, c)...)
		log.Debugf("%s", c.String())
	}

	log.Debugf("Following records present in delete")
	for _, c := range changes.Delete {
		zone, _ := zoneNameMapper.FindZone(c.DNSName)
		if zone == "" {
			log.Infof("Ignore record %s since it's not handled by ExternalDNS", c.DNSName)
			continue
		}
		recordsToDelete[zone] = append(recordsToDelete[zone], endpointToScalewayRecordsChangeDelete(zone, c)...)
		log.Debugf("%s", c.String())
	}

	log.Debugf("Following records present in create")
	for _, c := range changes.Create {
		zone, _ := zoneNameMapper.FindZone(c.DNSName)
		if zone == "" {
			log.Infof("Ignore record %s since it's not handled by ExternalDNS", c.DNSName)
			continue
		}
		recordsToAdd[zone].Records = append(recordsToAdd[zone].Records, endpointToScalewayRecords(zone, c)...)
		log.Debugf("%s", c.String())
	}

	log.Debugf("Following records present in updateNew")
	for _, c := range changes.UpdateNew {
		zone, _ := zoneNameMapper.FindZone(c.DNSName)
		if zone == "" {
			log.Infof("Ignore record %s since it's not handled by ExternalDNS", c.DNSName)
			continue
		}
		recordsToAdd[zone].Records = append(recordsToAdd[zone].Records, endpointToScalewayRecords(zone, c)...)
		log.Debugf("%s", c.String())
	}

	for _, zone := range dnsZones {
		zoneName := getCompleteZoneName(zone)
		req := &domain.UpdateDNSZoneRecordsRequest{
			DNSZone: zoneName,
			Changes: recordsToDelete[zoneName],
		}
		req.Changes = append(req.Changes, &domain.RecordChange{
			Add: recordsToAdd[zoneName],
		})
		// ignore sending empty update requests
		if len(req.Changes) == 1 && len(req.Changes[0].Add.Records) == 0 {
			continue
		}
		returnedRequests = append(returnedRequests, req)
	}

	return returnedRequests, nil
}

func getCompleteZoneName(zone *domain.DNSZone) string {
	subdomain := zone.Subdomain + "."
	if zone.Subdomain == "" {
		subdomain = ""
	}
	return subdomain + zone.Domain
}

// scalewayRecordType returns the record type to write: ALIAS for CNAME
// endpoints at the zone apex (where Scaleway rejects CNAME) or with the
// alias property set, the endpoint's own type otherwise.
func scalewayRecordType(ep *endpoint.Endpoint, relativeRecordName string) domain.RecordType {
	if isAliasCNAME(ep, relativeRecordName == "") {
		return domain.RecordTypeALIAS
	}
	return domain.RecordType(ep.RecordType)
}

func endpointToScalewayRecords(zoneName string, ep *endpoint.Endpoint) []*domain.Record {
	// no annotation results in a TTL of 0, default to 300 for consistency with other providers
	ttl := defaultTTL
	if ep.RecordTTL.IsConfigured() {
		ttl = uint32(ep.RecordTTL)
	}
	priority := scalewayDefaultPriority
	if prop, ok := ep.GetProviderSpecificProperty(scalewayPriorityKey); ok {
		prio, err := strconv.ParseUint(prop, 10, 32)
		if err != nil {
			log.Errorf("Failed parsing value of %s: %s: %v; using priority of %d", scalewayPriorityKey, prop, err, scalewayDefaultPriority)
		} else {
			priority = uint32(prio)
		}
	}

	records := []*domain.Record{}

	recordName := strings.Trim(strings.TrimSuffix(ep.DNSName, zoneName), ". ")
	recordType := scalewayRecordType(ep, recordName)

	for _, target := range ep.Targets {
		finalTargetName := target
		if endpoint.RequiresTrailingDot(ep.RecordType) {
			finalTargetName = provider.EnsureTrailingDot(target)
		}

		records = append(records, &domain.Record{
			Data:     finalTargetName,
			Name:     recordName,
			Priority: priority,
			TTL:      ttl,
			Type:     recordType,
		})
	}

	return records
}

func endpointToScalewayRecordsChangeDelete(zoneName string, ep *endpoint.Endpoint) []*domain.RecordChange {
	records := []*domain.RecordChange{}

	recordName := strings.Trim(strings.TrimSuffix(ep.DNSName, zoneName), ". ")
	recordType := scalewayRecordType(ep, recordName)

	for _, target := range ep.Targets {
		finalTargetName := target
		if endpoint.RequiresTrailingDot(ep.RecordType) {
			finalTargetName = provider.EnsureTrailingDot(target)
		}

		records = append(records, &domain.RecordChange{
			Delete: &domain.RecordChangeDelete{
				IDFields: &domain.RecordIdentifier{
					Data: &finalTargetName,
					Name: recordName,
					Type: recordType,
				},
			},
		})
	}

	return records
}

func logChanges(req *domain.UpdateDNSZoneRecordsRequest) {
	if !log.IsLevelEnabled(log.InfoLevel) {
		return
	}
	log.Infof("Updating zone %s", req.DNSZone)
	for _, change := range req.Changes {
		if change.Add != nil {
			for _, add := range change.Add.Records {
				name := add.Name + "."
				if add.Name == "" {
					name = ""
				}

				logFields := log.Fields{
					"record":   name + req.DNSZone,
					"type":     add.Type.String(),
					"ttl":      add.TTL,
					"priority": add.Priority,
					"data":     add.Data,
				}
				log.WithFields(logFields).Info("Adding record")
			}
		} else if change.Delete != nil {
			name := change.Delete.IDFields.Name + "."
			if change.Delete.IDFields.Name == "" {
				name = ""
			}

			logFields := log.Fields{
				"record": name + req.DNSZone,
				"type":   change.Delete.IDFields.Type.String(),
				"data":   *change.Delete.IDFields.Data,
			}

			log.WithFields(logFields).Info("Deleting record")
		}
	}
}
