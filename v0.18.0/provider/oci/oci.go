/*
Copyright 2018 The Kubernetes Authors.

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

package oci

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"github.com/oracle/oci-go-sdk/v65/dns"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const defaultTTL = 300

// OCIAuthConfig holds connection parameters for the OCI API.
type OCIAuthConfig struct {
	Region               string `yaml:"region"`
	TenancyID            string `yaml:"tenancy"`
	UserID               string `yaml:"user"`
	PrivateKey           string `yaml:"key"`
	Fingerprint          string `yaml:"fingerprint"`
	Passphrase           string `yaml:"passphrase"`
	UseInstancePrincipal bool   `yaml:"useInstancePrincipal"`
	UseWorkloadIdentity  bool   `yaml:"useWorkloadIdentity"`
}

// OCIConfig holds the configuration for the OCI Provider.
type OCIConfig struct {
	Auth              OCIAuthConfig `yaml:"auth"`
	CompartmentID     string        `yaml:"compartment"`
	ZoneCacheDuration time.Duration
}

// OCIProvider is an implementation of Provider for Oracle Cloud Infrastructure
// (OCI) DNS.
type OCIProvider struct {
	provider.BaseProvider
	client ociDNSClient
	cfg    OCIConfig

	domainFilter *endpoint.DomainFilter
	zoneIDFilter provider.ZoneIDFilter
	zoneScope    string
	zoneCache    *zoneCache
	dryRun       bool
}

// ociDNSClient is the subset of the OCI DNS API required by the OCI Provider.
type ociDNSClient interface {
	ListZones(ctx context.Context, request dns.ListZonesRequest) (response dns.ListZonesResponse, err error)
	GetZoneRecords(ctx context.Context, request dns.GetZoneRecordsRequest) (response dns.GetZoneRecordsResponse, err error)
	PatchZoneRecords(ctx context.Context, request dns.PatchZoneRecordsRequest) (response dns.PatchZoneRecordsResponse, err error)
}

// LoadOCIConfig reads and parses the OCI ExternalDNS config file at the given
// path.
func LoadOCIConfig(path string) (*OCIConfig, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading OCI config file %q: %w", path, err)
	}

	cfg := OCIConfig{}
	if err := yaml.Unmarshal(contents, &cfg); err != nil {
		return nil, fmt.Errorf("parsing OCI config file %q: %w", path, err)
	}
	return &cfg, nil
}

// NewOCIProvider initializes a new OCI DNS based Provider.
func NewOCIProvider(cfg OCIConfig, domainFilter *endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, zoneScope string, dryRun bool) (*OCIProvider, error) {
	var client ociDNSClient
	var err error
	var configProvider common.ConfigurationProvider
	if cfg.Auth.UseInstancePrincipal && cfg.Auth.UseWorkloadIdentity {
		return nil, errors.New("only one of 'useInstancePrincipal' and 'useWorkloadIdentity' may be enabled for Oracle authentication")
	}
	if cfg.Auth.UseWorkloadIdentity {
		// OCI SDK requires specific, dynamic environment variables for workload identity.
		if err := os.Setenv(auth.ResourcePrincipalVersionEnvVar, auth.ResourcePrincipalVersion2_2); err != nil {
			return nil, fmt.Errorf("unable to set OCI SDK environment variable: %s: %w", auth.ResourcePrincipalVersionEnvVar, err)
		}
		if err := os.Setenv(auth.ResourcePrincipalRegionEnvVar, cfg.Auth.Region); err != nil {
			return nil, fmt.Errorf("unable to set OCI SDK environment variable: %s: %w", auth.ResourcePrincipalRegionEnvVar, err)
		}
		configProvider, err = auth.OkeWorkloadIdentityConfigurationProvider()
		if err != nil {
			return nil, fmt.Errorf("error creating OCI workload identity config provider: %w", err)
		}
	} else if cfg.Auth.UseInstancePrincipal {
		configProvider, err = auth.InstancePrincipalConfigurationProvider()
		if err != nil {
			return nil, fmt.Errorf("error creating OCI instance principal config provider: %w", err)
		}
	} else {
		configProvider = common.NewRawConfigurationProvider(
			cfg.Auth.TenancyID,
			cfg.Auth.UserID,
			cfg.Auth.Region,
			cfg.Auth.Fingerprint,
			cfg.Auth.PrivateKey,
			&cfg.Auth.Passphrase,
		)
	}

	client, err = dns.NewDnsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, fmt.Errorf("initializing OCI DNS API client: %w", err)
	}

	return &OCIProvider{
		client:       client,
		cfg:          cfg,
		domainFilter: domainFilter,
		zoneIDFilter: zoneIDFilter,
		zoneScope:    zoneScope,
		zoneCache: &zoneCache{
			duration: cfg.ZoneCacheDuration,
		},
		dryRun: dryRun,
	}, nil
}

func (p *OCIProvider) zones(ctx context.Context) (map[string]dns.ZoneSummary, error) {
	if !p.zoneCache.Expired() {
		log.Debug("Using cached zones list")
		return p.zoneCache.zones, nil
	}
	zones := make(map[string]dns.ZoneSummary)
	scopes := []dns.GetZoneScopeEnum{dns.GetZoneScopeEnum(p.zoneScope)}
	// If the zone scope is empty, list all zones types.
	if p.zoneScope == "" {
		scopes = dns.GetGetZoneScopeEnumValues()
	}
	log.Debugf("Matching zones against domain filters: %v", p.domainFilter.Filters)
	for _, scope := range scopes {
		if err := p.addPaginatedZones(ctx, zones, scope); err != nil {
			return nil, err
		}
	}
	if len(zones) == 0 {
		log.Warnf("No zones in compartment %q match domain filters %v", p.cfg.CompartmentID, p.domainFilter)
	}
	p.zoneCache.Reset(zones)
	return zones, nil
}

// Merge Endpoints with the same Name and Type into a single endpoint with multiple Targets.
func mergeEndpointsMultiTargets(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	endpointsByNameType := map[string][]*endpoint.Endpoint{}

	for _, ep := range endpoints {
		key := fmt.Sprintf("%s-%s", ep.DNSName, ep.RecordType)
		endpointsByNameType[key] = append(endpointsByNameType[key], ep)
	}

	// If there were no merges, return endpoints.
	if len(endpointsByNameType) == len(endpoints) {
		return endpoints
	}

	// Otherwise, create a new list of endpoints with the consolidated targets.
	var mergedEndpoints []*endpoint.Endpoint
	for _, ep := range endpointsByNameType {
		dnsName := ep[0].DNSName
		recordType := ep[0].RecordType
		recordTTL := ep[0].RecordTTL

		targets := make([]string, len(ep))
		for i, e := range ep {
			targets[i] = e.Targets[0]
		}

		e := endpoint.NewEndpointWithTTL(dnsName, recordType, recordTTL, targets...)
		mergedEndpoints = append(mergedEndpoints, e)
	}

	return mergedEndpoints
}

func (p *OCIProvider) addPaginatedZones(ctx context.Context, zones map[string]dns.ZoneSummary, scope dns.GetZoneScopeEnum) error {
	var page *string
	// Loop until we have listed all zones.
	for {
		resp, err := p.client.ListZones(ctx, dns.ListZonesRequest{
			CompartmentId: &p.cfg.CompartmentID,
			ZoneType:      dns.ListZonesZoneTypePrimary,
			Scope:         dns.ListZonesScopeEnum(scope),
			Page:          page,
		})
		if err != nil {
			return provider.NewSoftError(fmt.Errorf("listing zones in %s: %w", p.cfg.CompartmentID, err))
		}
		for _, zone := range resp.Items {
			if p.domainFilter.Match(*zone.Name) && p.zoneIDFilter.Match(*zone.Id) {
				zones[*zone.Id] = zone
				log.Debugf("Matched %q (%q)", *zone.Name, *zone.Id)
			} else {
				log.Debugf("Filtered %q (%q)", *zone.Name, *zone.Id)
			}
		}
		if page = resp.OpcNextPage; resp.OpcNextPage == nil {
			break
		}
	}
	return nil
}

func (p *OCIProvider) newFilteredRecordOperations(endpoints []*endpoint.Endpoint, opType dns.RecordOperationOperationEnum) []dns.RecordOperation {
	var ops []dns.RecordOperation
	for _, ep := range endpoints {
		if ep == nil {
			continue
		}
		if p.domainFilter.Match(ep.DNSName) {
			for _, t := range ep.Targets {
				singleTargetEp := &endpoint.Endpoint{
					DNSName:          ep.DNSName,
					Targets:          []string{t},
					RecordType:       ep.RecordType,
					RecordTTL:        ep.RecordTTL,
					Labels:           ep.Labels,
					ProviderSpecific: ep.ProviderSpecific,
				}
				ops = append(ops, newRecordOperation(singleTargetEp, opType))
			}
		}
	}
	return ops
}

// Records returns the list of records in a given hosted zone.
func (p *OCIProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.zones(ctx)
	if err != nil {
		return nil, provider.NewSoftError(fmt.Errorf("getting zones: %w", err))
	}

	var endpoints []*endpoint.Endpoint
	for _, zone := range zones {
		var page *string
		for {
			resp, err := p.client.GetZoneRecords(ctx, dns.GetZoneRecordsRequest{
				ZoneNameOrId:  zone.Id,
				Page:          page,
				CompartmentId: &p.cfg.CompartmentID,
			})
			if err != nil {
				return nil, provider.NewSoftError(fmt.Errorf("getting records for zone %q: %w", *zone.Id, err))
			}

			for _, record := range resp.Items {
				if !provider.SupportedRecordType(*record.Rtype) {
					continue
				}
				endpoints = append(endpoints,
					endpoint.NewEndpointWithTTL(
						*record.Domain,
						*record.Rtype,
						endpoint.TTL(*record.Ttl),
						*record.Rdata,
					),
				)
			}

			if page = resp.OpcNextPage; resp.OpcNextPage == nil {
				break
			}
		}
	}

	endpoints = mergeEndpointsMultiTargets(endpoints)

	return endpoints, nil
}

// ApplyChanges applies a given set of changes to a given zone.
func (p *OCIProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	log.Debugf("Processing changes: %+v", changes)

	var ops []dns.RecordOperation
	ops = append(ops, p.newFilteredRecordOperations(changes.Create, dns.RecordOperationOperationAdd)...)

	ops = append(ops, p.newFilteredRecordOperations(changes.UpdateNew, dns.RecordOperationOperationAdd)...)
	ops = append(ops, p.newFilteredRecordOperations(changes.UpdateOld, dns.RecordOperationOperationRemove)...)

	ops = append(ops, p.newFilteredRecordOperations(changes.Delete, dns.RecordOperationOperationRemove)...)

	if len(ops) == 0 {
		log.Info("All records are already up to date")
		return nil
	}

	zones, err := p.zones(ctx)
	if err != nil {
		return provider.NewSoftError(fmt.Errorf("fetching zones: %w", err))
	}

	// Separate into per-zone change sets to be passed to OCI API.
	opsByZone := operationsByZone(zones, ops)
	for zoneID, ops := range opsByZone {
		log.Infof("Change zone: %q", zoneID)
		for _, op := range ops {
			log.Info(op)
		}
	}

	if p.dryRun {
		return nil
	}

	for zoneID, ops := range opsByZone {
		if _, err := p.client.PatchZoneRecords(ctx, dns.PatchZoneRecordsRequest{
			CompartmentId:           &p.cfg.CompartmentID,
			ZoneNameOrId:            &zoneID,
			PatchZoneRecordsDetails: dns.PatchZoneRecordsDetails{Items: ops},
		}); err != nil {
			return provider.NewSoftError(err)
		}
	}

	return nil
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (p *OCIProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	var adjustedEndpoints []*endpoint.Endpoint
	for _, e := range endpoints {
		// OCI DNS does not support the set-identifier attribute, so we remove it to avoid plan failure
		if e.SetIdentifier != "" {
			log.Warnf("Adjusting endpont: %v. Ignoring unsupported annotation 'set-identifier': %s", *e, e.SetIdentifier)
			e.SetIdentifier = ""
		}
		adjustedEndpoints = append(adjustedEndpoints, e)
	}
	return adjustedEndpoints, nil
}

// newRecordOperation returns a RecordOperation based on a given endpoint.
func newRecordOperation(ep *endpoint.Endpoint, opType dns.RecordOperationOperationEnum) dns.RecordOperation {
	targets := make([]string, len(ep.Targets))
	copy(targets, ep.Targets)
	if ep.RecordType == endpoint.RecordTypeCNAME {
		targets[0] = provider.EnsureTrailingDot(targets[0])
	}
	rdata := strings.Join(targets, " ")

	ttl := defaultTTL
	if ep.RecordTTL.IsConfigured() {
		ttl = int(ep.RecordTTL)
	}

	return dns.RecordOperation{
		Domain:    &ep.DNSName,
		Rdata:     &rdata,
		Ttl:       &ttl,
		Rtype:     &ep.RecordType,
		Operation: opType,
	}
}

// operationsByZone segments a slice of RecordOperations by their zone.
func operationsByZone(zones map[string]dns.ZoneSummary, ops []dns.RecordOperation) map[string][]dns.RecordOperation {
	changes := make(map[string][]dns.RecordOperation)

	zoneNameIDMapper := provider.ZoneIDName{}
	for _, z := range zones {
		zoneNameIDMapper.Add(*z.Id, *z.Name)
		changes[*z.Id] = []dns.RecordOperation{}
	}

	for _, op := range ops {
		if zoneID, _ := zoneNameIDMapper.FindZone(*op.Domain); zoneID != "" {
			changes[zoneID] = append(changes[zoneID], op)
		} else {
			log.Warnf("No matching zone for record operation %s", op)
		}
	}

	// Remove zones that don't have any changes.
	for zone, ops := range changes {
		if len(ops) == 0 {
			delete(changes, zone)
		}
	}

	return changes
}
