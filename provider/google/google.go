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
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
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

	providerSpecificRoutingPolicy     = "google-routing-policy"
	providerSpecificRoutingPolicyNone = ""
	providerSpecificRoutingPolicyGeo  = "geo"
	providerSpecificRoutingPolicyWrr  = "wrr"
	providerSpecificLocation          = "google-location"
	providerSpecificWeight            = "google-weight"
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

type resourceRecordSetsGetCallInterface interface {
	Do(opts ...googleapi.CallOption) (*dns.ResourceRecordSet, error)
}

type resourceRecordSetsListCallInterface interface {
	Pages(ctx context.Context, f func(*dns.ResourceRecordSetsListResponse) error) error
}

type resourceRecordSetsClientInterface interface {
	Get(managedZone string, name string, type_ string) resourceRecordSetsGetCallInterface
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

func (r resourceRecordSetsService) Get(managedZone string, name string, type_ string) resourceRecordSetsGetCallInterface {
	return r.service.Get(r.project, managedZone, name, type_)
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
	// The default location to use
	location string
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
func NewGoogleProvider(ctx context.Context, project string, location string, domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, batchChangeSize int, batchChangeInterval time.Duration, zoneVisibility string, dryRun bool) (*GoogleProvider, error) {
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

	if location == "" {
		if zone, err := metadata.ZoneWithContext(ctx); err != nil {
			parts := strings.Split(zone, "-")
			location = strings.Join(parts[:len(parts)-1], "-")
			log.Infof("Google location auto-detected: %s", location)
		}
	}

	zoneTypeFilter := provider.NewZoneTypeFilter(zoneVisibility)

	provider := &GoogleProvider{
		dryRun:              dryRun,
		batchChangeSize:     batchChangeSize,
		batchChangeInterval: batchChangeInterval,
		domainFilter:        domainFilter,
		zoneTypeFilter:      zoneTypeFilter,
		zoneIDFilter:        zoneIDFilter,
		location:            location,
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
			if r.RoutingPolicy != nil {
				if r.RoutingPolicy.Geo != nil {
					for _, item := range r.RoutingPolicy.Geo.Items {
						ep := endpoint.NewEndpointWithTTL(r.Name, r.Type, endpoint.TTL(r.Ttl), item.Rrdatas...)
						ep.WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyGeo)
						ep.WithProviderSpecific(providerSpecificLocation, item.Location)
						endpoints = append(endpoints, ep.WithSetIdentifier(item.Location))
					}
				}
				if r.RoutingPolicy.Wrr != nil {
					for index, item := range r.RoutingPolicy.Wrr.Items {
						ep := endpoint.NewEndpointWithTTL(r.Name, r.Type, endpoint.TTL(r.Ttl), item.Rrdatas...)
						ep.WithProviderSpecific(providerSpecificRoutingPolicy, providerSpecificRoutingPolicyWrr)
						ep.WithProviderSpecific(providerSpecificWeight, strconv.FormatFloat(item.Weight, 'g', 2, 64))
						endpoints = append(endpoints, ep.WithSetIdentifier(strconv.FormatInt(int64(index), 10)))
					}
				}
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

// AdjustEndpoints augments Endpoints generated by various sources to be equivalent to those returned by Records
func (p *GoogleProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	for _, ep := range endpoints {
		if routingPolicy, ok := ep.GetProviderSpecificProperty(providerSpecificRoutingPolicy); ok {
			switch routingPolicy {
			case providerSpecificRoutingPolicyGeo:
				location, ok := ep.GetProviderSpecificProperty(providerSpecificLocation)
				if !ok && p.location != "" {
					ep.WithProviderSpecific(providerSpecificLocation, p.location)
					location = p.location
				}
				ep.WithSetIdentifier(location)
			case providerSpecificRoutingPolicyWrr:
				weight, ok := ep.GetProviderSpecificProperty(providerSpecificWeight)
				if !ok {
					weight = "100"
				}
				if weight, err := strconv.ParseFloat(weight, 64); err == nil {
					ep.WithProviderSpecific(providerSpecificWeight, strconv.FormatFloat(weight, 'g', 2, 64))
				}
				if index, err := strconv.ParseInt(ep.SetIdentifier, 10, 64); err != nil || index < 0 {
					resource := ep.Labels[endpoint.ResourceLabelKey]
					log.Warnf("Endpoint generated from '%s' has 'wrr' routing policy with non-integer set identifier '%s'", resource, ep.SetIdentifier)
				}
			}
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
			change := p.newChange(rrSetChange, zone)
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
func (p *GoogleProvider) newChange(rrSetChange *plan.RRSetChange, zone string) *dns.Change {
	change := dns.Change{}
	current, err := p.resourceRecordSetsClient.Get(zone, string(rrSetChange.Name), string(rrSetChange.Type)).Do()
	if err != nil {
		if err, ok := err.(*googleapi.Error); !ok || err.Code != http.StatusNotFound || len(rrSetChange.Delete) > 0 {
			log.Errorf("Error obtaining resource record set %s %s: %v", rrSetChange.Name, rrSetChange.Type, err)
			return &dns.Change{}
		}
		current = &dns.ResourceRecordSet{}
	} else {
		change.Deletions = []*dns.ResourceRecordSet{current}
	}
	desired := &dns.ResourceRecordSet{}
	data, err := current.MarshalJSON()
	if err == nil {
		err = json.Unmarshal(data, desired)
	}
	if err != nil {
		log.Errorf("Error processing resource record set %s %s: %v", rrSetChange.Name, rrSetChange.Type, err)
		return &dns.Change{}
	}
	for _, ep := range rrSetChange.Delete {
		routingPolicy, _ := ep.GetProviderSpecificProperty(providerSpecificRoutingPolicy)
		switch routingPolicy {
		case providerSpecificRoutingPolicyNone:
			if len(desired.Rrdatas) == len(ep.Targets) {
				match := true
				for _, data := range rrSetDatas(ep) {
					if !slices.Contains(desired.Rrdatas, data) {
						match = false
						break
					}
				}
				if match {
					desired = &dns.ResourceRecordSet{}
				}
			}
		case providerSpecificRoutingPolicyGeo:
			if desired.RoutingPolicy == nil || desired.RoutingPolicy.Geo == nil {
				continue
			}
			if location, ok := ep.GetProviderSpecificProperty(providerSpecificLocation); ok {
				filter := func(item *dns.RRSetRoutingPolicyGeoPolicyGeoPolicyItem) bool {
					if item.Location == location && len(item.Rrdatas) == len(ep.Targets) {
						for _, data := range rrSetDatas(ep) {
							if !slices.Contains(item.Rrdatas, data) {
								return false
							}
						}
						return true
					}
					return false
				}
				index := slices.IndexFunc(desired.RoutingPolicy.Geo.Items, filter)
				for index != -1 {
					desired.RoutingPolicy.Geo.Items = slices.Delete(
						desired.RoutingPolicy.Geo.Items, index, index+1,
					)
					index = slices.IndexFunc(desired.RoutingPolicy.Geo.Items, filter)
				}
				if len(desired.RoutingPolicy.Geo.Items) == 0 {
					desired = &dns.ResourceRecordSet{}
				}
			}
		case providerSpecificRoutingPolicyWrr:
			if desired.RoutingPolicy == nil || desired.RoutingPolicy.Wrr == nil {
				continue
			}
			index, err := strconv.ParseInt(ep.SetIdentifier, 10, 64)
			length := int64(len(desired.RoutingPolicy.Wrr.Items))
			weight, ok := ep.GetProviderSpecificProperty(providerSpecificWeight)
			if ok && err == nil && index >= 0 && index < length {
				weight, err := strconv.ParseFloat(weight, 64)
				if err != nil {
					continue
				}
				item := desired.RoutingPolicy.Wrr.Items[index]
				if item.Weight != weight || len(item.Rrdatas) != len(ep.Targets) {
					continue
				}
				match := true
				for _, data := range rrSetDatas(ep) {
					if !slices.Contains(item.Rrdatas, data) {
						match = false
						break
					}
				}
				if match {
					if index+1 < length {
						desired.RoutingPolicy.Wrr.Items[index] = rrSetRoutingPolicyWrrItemPlaceholder(desired.Type)
					} else {
						desired.RoutingPolicy.Wrr.Items = desired.RoutingPolicy.Wrr.Items[:index]
					}
				}
				filter := func(item *dns.RRSetRoutingPolicyWrrPolicyWrrPolicyItem) bool {
					return item.Weight != 0
				}
				index := slices.IndexFunc(desired.RoutingPolicy.Wrr.Items, filter)
				if index == -1 {
					desired = &dns.ResourceRecordSet{}
				}
			}
		}
	}
	for _, ep := range rrSetChange.Create {
		desired.Name = string(rrSetChange.Name)
		desired.Type = string(rrSetChange.Type)
		if ep.RecordTTL.IsConfigured() {
			desired.Ttl = int64(ep.RecordTTL)
		}
		desired.Rrdatas = rrSetDatas(ep)
		routingPolicy, _ := ep.GetProviderSpecificProperty(providerSpecificRoutingPolicy)
		switch routingPolicy {
		case providerSpecificRoutingPolicyGeo:
			if location, ok := ep.GetProviderSpecificProperty(providerSpecificLocation); ok {
				if desired.RoutingPolicy == nil {
					desired.RoutingPolicy = &dns.RRSetRoutingPolicy{}
				}
				if desired.RoutingPolicy.Geo == nil {
					desired.RoutingPolicy.Geo = &dns.RRSetRoutingPolicyGeoPolicy{}
				}
				index := -1
				for i, item := range desired.RoutingPolicy.Geo.Items {
					if item.Location == location {
						index = i
						break
					}
				}
				if index == -1 {
					index = len(desired.RoutingPolicy.Geo.Items)
					desired.RoutingPolicy.Geo.Items = append(desired.RoutingPolicy.Geo.Items, nil)
				}
				desired.RoutingPolicy.Geo.Items[index] = &dns.RRSetRoutingPolicyGeoPolicyGeoPolicyItem{
					Location: location,
					Rrdatas:  desired.Rrdatas,
				}
				desired.Rrdatas = nil
			}
		case providerSpecificRoutingPolicyWrr:
			index, err := strconv.ParseInt(ep.SetIdentifier, 10, 64)
			weight, ok := ep.GetProviderSpecificProperty(providerSpecificWeight)
			if ok && err == nil && index >= 0 {
				weight, err := strconv.ParseFloat(weight, 64)
				if err != nil {
					continue
				}
				if desired.RoutingPolicy == nil {
					desired.RoutingPolicy = &dns.RRSetRoutingPolicy{}
				}
				if desired.RoutingPolicy.Wrr == nil {
					desired.RoutingPolicy.Wrr = &dns.RRSetRoutingPolicyWrrPolicy{}
				}
				length := int64(len(desired.RoutingPolicy.Wrr.Items))
				if index >= length {
					desired.RoutingPolicy.Wrr.Items = slices.Grow(desired.RoutingPolicy.Wrr.Items, int(index+1-length))[:index+1]
					for i := length; i < index; i++ {
						desired.RoutingPolicy.Wrr.Items[i] = rrSetRoutingPolicyWrrItemPlaceholder(desired.Type)
					}
				}
				desired.RoutingPolicy.Wrr.Items[index] = &dns.RRSetRoutingPolicyWrrPolicyWrrPolicyItem{
					Weight:  weight,
					Rrdatas: desired.Rrdatas,
				}
				desired.Rrdatas = nil
			}
		}
	}
	if desired.Name != "" {
		if desired.Ttl == 0 {
			desired.Ttl = googleRecordTTL
		}
		change.Additions = []*dns.ResourceRecordSet{desired}
	}
	return &change
}

// Return Resource Record Set data for given endpoint
func rrSetDatas(ep *endpoint.Endpoint) []string {
	// TODO(linki): works around appending a trailing dot to TXT records. I think
	// we should go back to storing DNS names with a trailing dot internally. This
	// way we can use it has is here and trim it off if it exists when necessary.
	switch ep.RecordType {
	case endpoint.RecordTypeCNAME:
		return []string{provider.EnsureTrailingDot(ep.Targets[0])}
	case endpoint.RecordTypeMX:
		fallthrough
	case endpoint.RecordTypeNS:
		fallthrough
	case endpoint.RecordTypeSRV:
		rrdatas := make([]string, len(ep.Targets))
		for i, target := range ep.Targets {
			rrdatas[i] = provider.EnsureTrailingDot(target)
		}
		return rrdatas
	default:
		return slices.Clone(ep.Targets)
	}
}

// Return a Weighted Round Robin routing policy item placeholder for given resource record type
func rrSetRoutingPolicyWrrItemPlaceholder(type_ string) *dns.RRSetRoutingPolicyWrrPolicyWrrPolicyItem {
	var rrdatas []string
	switch type_ {
	case "A":
		rrdatas = []string{"0.0.0.0"}
	case "AAAA":
		rrdatas = []string{"::"}
	default:
		rrdatas = []string{"."}
	}
	return &dns.RRSetRoutingPolicyWrrPolicyWrrPolicyItem{
		Rrdatas: rrdatas,
	}
}
