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

package google

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"cloud.google.com/go/compute/metadata"
	"github.com/linki/instrumented_http"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	dns "google.golang.org/api/dns/v1"
	googleapi "google.golang.org/api/googleapi"
	"google.golang.org/api/option"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	googleRecordTTL = 300
)

type managedZonesCreateCallInterface interface {
	Do(opts ...googleapi.CallOption) (*dns.ManagedZone, error)
}

type managedZonesListCallInterface interface {
	Pages(ctx context.Context, f func(*dns.ManagedZonesListResponse) error) error
}

type managedZonesServiceInterface interface {
	Create(managedZone *dns.ManagedZone) managedZonesCreateCallInterface
	List() managedZonesListCallInterface
}

type resourceRecordSetsListCallInterface interface {
	Pages(ctx context.Context, f func(*dns.ResourceRecordSetsListResponse) error) error
}

type resourceRecordSetsClientInterface interface {
	List(managedZone string) resourceRecordSetsListCallInterface
}

type changesCreateCallInterface interface {
	Do(opts ...googleapi.CallOption) (*dns.Change, error)
}

type changesServiceInterface interface {
	Create(managedZone string, change *dns.Change) changesCreateCallInterface
}

type resourceRecordSetsService struct {
	project string
	service *dns.ResourceRecordSetsService
}

func (r resourceRecordSetsService) List(managedZone string) resourceRecordSetsListCallInterface {
	return r.service.List(r.project, managedZone)
}

type managedZonesService struct {
	project string
	service *dns.ManagedZonesService
}

func (m managedZonesService) Create(managedZone *dns.ManagedZone) managedZonesCreateCallInterface {
	return m.service.Create(m.project, managedZone)
}

func (m managedZonesService) List() managedZonesListCallInterface {
	return m.service.List(m.project)
}

type changesService struct {
	project string
	service *dns.ChangesService
}

func (c changesService) Create(managedZone string, change *dns.Change) changesCreateCallInterface {
	return c.service.Create(c.project, managedZone, change)
}

// GoogleProvider is an implementation of Provider for Google CloudDNS.
type GoogleProvider struct {
	provider.BaseProvider
	// Enabled dry-run will print any modifying actions rather than execute them.
	dryRun bool
	// Max batch size to submit to Google Cloud DNS per transaction.
	batchChangeSize int
	// Interval between batch updates.
	batchChangeInterval time.Duration
	// only consider hosted zones managing domains ending in this suffix
	domainFilter endpoint.DomainFilter
	// filter for zones based on visibility
	zoneTypeFilter provider.ZoneTypeFilter
	// only consider hosted zones ending with this zone id
	zoneIDFilter provider.ZoneIDFilter
	// A client for managing resource record sets
	resourceRecordSetsClient resourceRecordSetsClientInterface
	// A client for managing hosted zones
	managedZonesClient managedZonesServiceInterface
	// A client for managing change sets
	changesClient changesServiceInterface
	// The context parameter to be passed for gcloud API calls.
	ctx context.Context
}

// NewGoogleProvider initializes a new Google CloudDNS based Provider.
func NewGoogleProvider(ctx context.Context, project string, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, batchChangeSize int, batchChangeInterval time.Duration, zoneVisibility string, dryRun bool) (*GoogleProvider, error) {
	gcloud, err := google.DefaultClient(ctx, dns.NdevClouddnsReadwriteScope)
	if err != nil {
		return nil, err
	}

	gcloud = instrumented_http.NewClient(gcloud, &instrumented_http.Callbacks{
		PathProcessor: func(path string) string {
			parts := strings.Split(path, "/")
			return parts[len(parts)-1]
		},
	})

	dnsClient, err := dns.NewService(ctx, option.WithHTTPClient(gcloud))
	if err != nil {
		return nil, err
	}

	if project == "" {
		mProject, mErr := metadata.ProjectIDWithContext(ctx)
		if mErr != nil {
			return nil, fmt.Errorf("failed to auto-detect the project id: %w", mErr)
		}
		log.Infof("Google project auto-detected: %s", mProject)
		project = mProject
	}

	zoneTypeFilter := provider.NewZoneTypeFilter(zoneVisibility)

	provider := &GoogleProvider{
		dryRun:              dryRun,
		batchChangeSize:     batchChangeSize,
		batchChangeInterval: batchChangeInterval,
		domainFilter:        domainFilter,
		zoneTypeFilter:      zoneTypeFilter,
		zoneIDFilter:        zoneIDFilter,
		managedZonesClient: managedZonesService{
			project: project,
			service: dnsClient.ManagedZones,
		},
		resourceRecordSetsClient: resourceRecordSetsService{
			project: project,
			service: dnsClient.ResourceRecordSets,
		},
		changesClient: changesService{
			project: project,
			service: dnsClient.Changes,
		},
		ctx: ctx,
	}

	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *GoogleProvider) Zones(ctx context.Context) (map[string]*dns.ManagedZone, error) {
	zones := make(map[string]*dns.ManagedZone)

	f := func(resp *dns.ManagedZonesListResponse) error {
		for _, zone := range resp.ManagedZones {
			if zone.PeeringConfig == nil {
				if p.domainFilter.Match(zone.DnsName) && p.zoneTypeFilter.Match(zone.Visibility) && (p.zoneIDFilter.Match(fmt.Sprintf("%v", zone.Id)) || p.zoneIDFilter.Match(fmt.Sprintf("%v", zone.Name))) {
					zones[zone.Name] = zone
					log.Debugf("Matched %s (zone: %s) (visibility: %s)", zone.DnsName, zone.Name, zone.Visibility)
				} else {
					log.Debugf("Filtered %s (zone: %s) (visibility: %s)", zone.DnsName, zone.Name, zone.Visibility)
				}
			} else {
				log.Debugf("Filtered peering zone %s (zone: %s) (visibility: %s)", zone.DnsName, zone.Name, zone.Visibility)
			}
		}

		return nil
	}

	log.Debugf("Matching zones against domain filters: %v", p.domainFilter)
	if err := p.managedZonesClient.List().Pages(ctx, f); err != nil {
		return nil, provider.NewSoftError(fmt.Errorf("failed to list zones: %w", err))
	}

	if len(zones) == 0 {
		log.Warnf("No zones match domain filters: %v", p.domainFilter)
	}

	for _, zone := range zones {
		log.Debugf("Considering zone: %s (domain: %s)", zone.Name, zone.DnsName)
	}

	return zones, nil
}

// Records returns the list of records in all relevant zones.
func (p *GoogleProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	f := func(resp *dns.ResourceRecordSetsListResponse) error {
		for _, r := range resp.Rrsets {
			if !p.SupportedRecordType(r.Type) {
				continue
			}
			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(r.Name, r.Type, endpoint.TTL(r.Ttl), r.Rrdatas...))
		}

		return nil
	}

	for _, z := range zones {
		if err := p.resourceRecordSetsClient.List(z.Name).Pages(ctx, f); err != nil {
			return nil, provider.NewSoftError(fmt.Errorf("failed to list records in zone %s: %w", z.Name, err))
		}
	}

	return endpoints, nil
}

// ApplyChanges applies a given set of changes.
func (p *GoogleProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	zones, err := p.Zones(ctx)
	if err != nil {
		return err
	}
	zoneMap := provider.ZoneIDName{}
	for _, z := range zones {
		zoneMap.Add(z.Name, z.DnsName)
	}
	zoneBatches := map[string][]*dns.Change{}
	for rrSetChange := range changes.All() {
		if zone, _ := zoneMap.FindZone(string(rrSetChange.Name)); zone != "" {
			change := p.newChange(rrSetChange)
			changeSize := len(change.Additions) + len(change.Deletions)
			if changeSize == 0 {
				continue
			}
			if _, ok := zoneBatches[zone]; !ok {
				zoneBatches[zone] = []*dns.Change{{}}
			}
			batch := zoneBatches[zone][len(zoneBatches[zone])-1]
			if p.batchChangeSize > 0 && len(batch.Additions)+len(batch.Deletions)+changeSize > p.batchChangeSize {
				batch = &dns.Change{}
				zoneBatches[zone] = append(zoneBatches[zone], batch)
			}
			batch.Additions = append(batch.Additions, change.Additions...)
			batch.Deletions = append(batch.Deletions, change.Deletions...)
		}
	}

	for zone, batches := range zoneBatches {
		for index, batch := range batches {
			log.Infof("Change zone: %v batch #%d", zone, index)
			for _, record := range batch.Deletions {
				log.Infof("Del records: %s %s %s %d", record.Name, record.Type, record.Rrdatas, record.Ttl)
			}
			for _, record := range batch.Additions {
				log.Infof("Add records: %s %s %s %d", record.Name, record.Type, record.Rrdatas, record.Ttl)
			}
			if p.dryRun {
				continue
			}
			if index > 0 {
				time.Sleep(p.batchChangeInterval)
			}
			if _, err := p.changesClient.Create(zone, batch).Do(); err != nil {
				return provider.NewSoftError(fmt.Errorf("failed to create changes: %w", err))
			}
		}
	}

	return nil
}

// SupportedRecordType returns true if the record type is supported by the provider
func (p *GoogleProvider) SupportedRecordType(recordType string) bool {
	switch recordType {
	case "MX":
		return true
	default:
		return provider.SupportedRecordType(recordType)
	}
}

// newChange returns a DNS change based upon the given resource record set change.
func (p *GoogleProvider) newChange(rrSetChange *plan.RRSetChange) *dns.Change {
	change := dns.Change{}
	for index, endpoints := range [][]*endpoint.Endpoint{rrSetChange.Delete, rrSetChange.Create} {
		for _, ep := range endpoints {
			record := dns.ResourceRecordSet{
				Name: provider.EnsureTrailingDot(ep.DNSName),
				Ttl:  googleRecordTTL,
				Type: ep.RecordType,
			}
			if ep.RecordTTL.IsConfigured() {
				record.Ttl = int64(ep.RecordTTL)
			}
			// TODO(linki): works around appending a trailing dot to TXT records. I think
			// we should go back to storing DNS names with a trailing dot internally. This
			// way we can use it has is here and trim it off if it exists when necessary.
			switch record.Type {
			case endpoint.RecordTypeCNAME:
				record.Rrdatas = []string{provider.EnsureTrailingDot(ep.Targets[0])}
			case endpoint.RecordTypeMX:
				fallthrough
			case endpoint.RecordTypeNS:
				fallthrough
			case endpoint.RecordTypeSRV:
				record.Rrdatas = make([]string, len(ep.Targets))
				for i, target := range ep.Targets {
					record.Rrdatas[i] = provider.EnsureTrailingDot(target)
				}
			default:
				record.Rrdatas = slices.Clone(ep.Targets)
			}
			switch index {
			case 0:
				change.Deletions = append(change.Deletions, &record)
			case 1:
				change.Additions = append(change.Additions, &record)
			}
		}
	}
	return &change
}
