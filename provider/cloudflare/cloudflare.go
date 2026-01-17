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

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/addressing"
	"github.com/cloudflare/cloudflare-go/v5/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v5/dns"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/zones"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/publicsuffix"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/source/annotations"
)

type changeAction int

const (
	// Environment variable names for CloudFlare authentication
	cfAPIEmailEnvKey = "CF_API_EMAIL"
	cfAPIKeyEnvKey   = "CF_API_KEY"
	cfAPITokenEnvKey = "CF_API_TOKEN"

	// cloudFlareCreate is a ChangeAction enum value
	cloudFlareCreate changeAction = iota
	// cloudFlareDelete is a ChangeAction enum value
	cloudFlareDelete
	// cloudFlareUpdate is a ChangeAction enum value
	cloudFlareUpdate
	// defaultTTL 1 = automatic
	defaultTTL = 1

	// Cloudflare tier limitations https://developers.cloudflare.com/dns/manage-dns-records/reference/record-attributes/#availability
	freeZoneMaxCommentLength = 100
	paidZoneMaxCommentLength = 500
)

var changeActionNames = map[changeAction]string{
	cloudFlareCreate: "CREATE",
	cloudFlareDelete: "DELETE",
	cloudFlareUpdate: "UPDATE",
}

func (action changeAction) String() string {
	return changeActionNames[action]
}

type DNSRecordIndex struct {
	Name    string
	Type    string
	Content string
}

type DNSRecordsMap map[DNSRecordIndex]dns.RecordResponse

var recordTypeProxyNotSupported = map[string]bool{
	"LOC": true,
	"MX":  true,
	"NS":  true,
	"SPF": true,
	"TXT": true,
	"SRV": true,
}

// cloudFlareDNS is the subset of the CloudFlare API that we actually use.  Add methods as required. Signatures must match exactly.
type cloudFlareDNS interface {
	ZoneIDByName(zoneName string) (string, error)
	ListZones(ctx context.Context, params zones.ZoneListParams) autoPager[zones.Zone]
	GetZone(ctx context.Context, zoneID string) (*zones.Zone, error)
	ListDNSRecords(ctx context.Context, params dns.RecordListParams) autoPager[dns.RecordResponse]
	CreateDNSRecord(ctx context.Context, params dns.RecordNewParams) (*dns.RecordResponse, error)
	DeleteDNSRecord(ctx context.Context, recordID string, params dns.RecordDeleteParams) error
	UpdateDNSRecord(ctx context.Context, recordID string, params dns.RecordUpdateParams) (*dns.RecordResponse, error)
	ListDataLocalizationRegionalHostnames(ctx context.Context, params addressing.RegionalHostnameListParams) autoPager[addressing.RegionalHostnameListResponse]
	CreateDataLocalizationRegionalHostname(ctx context.Context, params addressing.RegionalHostnameNewParams) error
	UpdateDataLocalizationRegionalHostname(ctx context.Context, hostname string, params addressing.RegionalHostnameEditParams) error
	DeleteDataLocalizationRegionalHostname(ctx context.Context, hostname string, params addressing.RegionalHostnameDeleteParams) error
	CustomHostnames(ctx context.Context, zoneID string) autoPager[custom_hostnames.CustomHostnameListResponse]
	DeleteCustomHostname(ctx context.Context, customHostnameID string, params custom_hostnames.CustomHostnameDeleteParams) error
	CreateCustomHostname(ctx context.Context, zoneID string, ch customHostname) error
}

type zoneService struct {
	service *cloudflare.Client
}

func (z zoneService) ZoneIDByName(zoneName string) (string, error) {
	// Use v4 API to find zone by name
	params := zones.ZoneListParams{
		Name: cloudflare.F(zoneName),
	}

	iter := z.service.Zones.ListAutoPaging(context.Background(), params)
	for zone := range autoPagerIterator(iter) {
		if zone.Name == zoneName {
			return zone.ID, nil
		}
	}

	if err := iter.Err(); err != nil {
		return "", fmt.Errorf("failed to list zones from CloudFlare API: %w", err)
	}

	return "", fmt.Errorf("zone %q not found in CloudFlare account - verify the zone exists and API credentials have access to it", zoneName)
}

func (z zoneService) CreateDNSRecord(ctx context.Context, params dns.RecordNewParams) (*dns.RecordResponse, error) {
	return z.service.DNS.Records.New(ctx, params)
}

func (z zoneService) ListDNSRecords(ctx context.Context, params dns.RecordListParams) autoPager[dns.RecordResponse] {
	return z.service.DNS.Records.ListAutoPaging(ctx, params)
}

func (z zoneService) UpdateDNSRecord(ctx context.Context, recordID string, params dns.RecordUpdateParams) (*dns.RecordResponse, error) {
	return z.service.DNS.Records.Update(ctx, recordID, params)
}

func (z zoneService) DeleteDNSRecord(ctx context.Context, recordID string, params dns.RecordDeleteParams) error {
	_, err := z.service.DNS.Records.Delete(ctx, recordID, params)
	return err
}

func (z zoneService) ListZones(ctx context.Context, params zones.ZoneListParams) autoPager[zones.Zone] {
	return z.service.Zones.ListAutoPaging(ctx, params)
}

func (z zoneService) GetZone(ctx context.Context, zoneID string) (*zones.Zone, error) {
	return z.service.Zones.Get(ctx, zones.ZoneGetParams{ZoneID: cloudflare.F(zoneID)})
}

// listZonesV4Params returns the appropriate Zone List Params for v4 API
func listZonesV4Params() zones.ZoneListParams {
	return zones.ZoneListParams{}
}

type DNSRecordsConfig struct {
	PerPage int
	Comment string
}

func (c *DNSRecordsConfig) trimAndValidateComment(dnsName, comment string, paidZone func(string) bool) string {
	if len(comment) <= freeZoneMaxCommentLength {
		return comment
	}

	maxLength := freeZoneMaxCommentLength
	if paidZone(dnsName) {
		maxLength = paidZoneMaxCommentLength
	}

	if len(comment) > maxLength {
		log.Warnf("DNS record comment is invalid. Trimming comment of %s. To avoid endless syncs, please set it to less than %d chars.", dnsName, maxLength)
		return comment[:maxLength]
	}

	return comment
}

func (p *CloudFlareProvider) ZoneHasPaidPlan(hostname string) bool {
	zone, err := publicsuffix.EffectiveTLDPlusOne(hostname)
	if err != nil {
		log.Errorf("Failed to get effective TLD+1 for hostname %s %v", hostname, err)
		return false
	}
	zoneID, err := p.Client.ZoneIDByName(zone)
	if err != nil {
		log.Errorf("Failed to get zone %s by name %v", zone, err)
		return false
	}

	zoneDetails, err := p.Client.GetZone(context.Background(), zoneID)
	if err != nil {
		log.Errorf("Failed to get zone %s details %v", zone, err)
		return false
	}

	return zoneDetails.Plan.IsSubscribed //nolint:staticcheck // SA1019: Plan.IsSubscribed is deprecated but no replacement available yet
}

// CloudFlareProvider is an implementation of Provider for CloudFlare DNS.
type CloudFlareProvider struct {
	provider.BaseProvider
	Client cloudFlareDNS
	// only consider hosted zones managing domains ending in this suffix
	domainFilter           *endpoint.DomainFilter
	zoneIDFilter           provider.ZoneIDFilter
	proxiedByDefault       bool
	DryRun                 bool
	CustomHostnamesConfig  CustomHostnamesConfig
	DNSRecordsConfig       DNSRecordsConfig
	RegionalServicesConfig RegionalServicesConfig
}

// cloudFlareChange differentiates between ChangeActions
type cloudFlareChange struct {
	Action              changeAction
	ResourceRecord      dns.RecordResponse
	RegionalHostname    regionalHostname
	CustomHostnames     map[string]customHostname
	CustomHostnamesPrev []string
}

// updateDNSRecordParam is a function that returns the appropriate Record Param based on the cloudFlareChange passed in
func getUpdateDNSRecordParam(zoneID string, cfc cloudFlareChange) dns.RecordUpdateParams {
	return dns.RecordUpdateParams{
		ZoneID: cloudflare.F(zoneID),
		Body: dns.RecordUpdateParamsBody{
			Name:     cloudflare.F(cfc.ResourceRecord.Name),
			TTL:      cloudflare.F(cfc.ResourceRecord.TTL),
			Proxied:  cloudflare.F(cfc.ResourceRecord.Proxied),
			Type:     cloudflare.F(dns.RecordUpdateParamsBodyType(cfc.ResourceRecord.Type)),
			Content:  cloudflare.F(cfc.ResourceRecord.Content),
			Priority: cloudflare.F(cfc.ResourceRecord.Priority),
			Comment:  cloudflare.F(cfc.ResourceRecord.Comment),
			Tags:     cloudflare.F(cfc.ResourceRecord.Tags),
		},
	}
}

// getCreateDNSRecordParam is a function that returns the appropriate Record Param based on the cloudFlareChange passed in
func getCreateDNSRecordParam(zoneID string, cfc *cloudFlareChange) dns.RecordNewParams {
	return dns.RecordNewParams{
		ZoneID: cloudflare.F(zoneID),
		Body: dns.RecordNewParamsBody{
			Name:     cloudflare.F(cfc.ResourceRecord.Name),
			TTL:      cloudflare.F(cfc.ResourceRecord.TTL),
			Proxied:  cloudflare.F(cfc.ResourceRecord.Proxied),
			Type:     cloudflare.F(dns.RecordNewParamsBodyType(cfc.ResourceRecord.Type)),
			Content:  cloudflare.F(cfc.ResourceRecord.Content),
			Priority: cloudflare.F(cfc.ResourceRecord.Priority),
			Comment:  cloudflare.F(cfc.ResourceRecord.Comment),
			Tags:     cloudflare.F(cfc.ResourceRecord.Tags),
		},
	}
}

func convertCloudflareError(err error) error {
	// Handle CloudFlare v5 SDK errors according to the documentation:
	// https://github.com/cloudflare/cloudflare-go?tab=readme-ov-file#errors
	var apierr *cloudflare.Error
	if errors.As(err, &apierr) {
		// Rate limit errors (429) and server errors (5xx) should be treated as soft errors
		// so that external-dns will retry them later
		if apierr.StatusCode == http.StatusTooManyRequests || apierr.StatusCode >= http.StatusInternalServerError {
			return provider.NewSoftError(err)
		}
		// For other structured API errors (4xx), return the error unchanged
		// Note: We must NOT call err.Error() on v5 cloudflare.Error types with nil internal fields
		return err
	}

	// Also check for rate limit indicators in error message strings as a fallback.
	// The v5 SDK's retry logic and error wrapping can hide the structured error type,
	// so we need string matching to catch rate limits in wrapped errors like:
	// "exceeded available rate limit retries" from the SDK's auto-retry mechanism.
	errMsg := strings.ToLower(err.Error())
	if strings.Contains(errMsg, "rate limit") ||
		strings.Contains(errMsg, "429") ||
		strings.Contains(errMsg, "exceeded available rate limit retries") ||
		strings.Contains(errMsg, "too many requests") {
		return provider.NewSoftError(err)
	}

	return err
}

// NewCloudFlareProvider initializes a new CloudFlare DNS based Provider.
func NewCloudFlareProvider(
	domainFilter *endpoint.DomainFilter,
	zoneIDFilter provider.ZoneIDFilter,
	proxiedByDefault bool,
	dryRun bool,
	regionalServicesConfig RegionalServicesConfig,
	customHostnamesConfig CustomHostnamesConfig,
	dnsRecordsConfig DNSRecordsConfig,
) (*CloudFlareProvider, error) {
	// initialize via chosen auth method and returns new API object

	var client *cloudflare.Client

	token := os.Getenv(cfAPITokenEnvKey)
	if token != "" {
		if trimed, ok := strings.CutPrefix(token, "file:"); ok {
			tokenBytes, err := os.ReadFile(trimed)
			if err != nil {
				return nil, fmt.Errorf("failed to read %s from file: %w", cfAPITokenEnvKey, err)
			}
			token = strings.TrimSpace(string(tokenBytes))
		}
		client = cloudflare.NewClient(
			option.WithAPIToken(token),
		)
	} else {
		apiKey := os.Getenv(cfAPIKeyEnvKey)
		apiEmail := os.Getenv(cfAPIEmailEnvKey)
		if apiKey == "" || apiEmail == "" {
			return nil, fmt.Errorf("cloudflare credentials are not configured: set either %s or both %s and %s environment variables", cfAPITokenEnvKey, cfAPIKeyEnvKey, cfAPIEmailEnvKey)
		}
		client = cloudflare.NewClient(
			option.WithAPIKey(apiKey),
			option.WithAPIEmail(apiEmail),
		)
	}

	if regionalServicesConfig.RegionKey != "" {
		regionalServicesConfig.Enabled = true
	}

	return &CloudFlareProvider{
		Client:                 zoneService{client},
		domainFilter:           domainFilter,
		zoneIDFilter:           zoneIDFilter,
		proxiedByDefault:       proxiedByDefault,
		CustomHostnamesConfig:  customHostnamesConfig,
		DryRun:                 dryRun,
		RegionalServicesConfig: regionalServicesConfig,
		DNSRecordsConfig:       dnsRecordsConfig,
	}, nil
}

// Zones returns the list of hosted zones.
func (p *CloudFlareProvider) Zones(ctx context.Context) ([]zones.Zone, error) {
	var result []zones.Zone

	// if there is a zoneIDfilter configured
	// && if the filter isn't just a blank string (used in tests)
	if len(p.zoneIDFilter.ZoneIDs) > 0 && p.zoneIDFilter.ZoneIDs[0] != "" {
		log.Debugln("zoneIDFilter configured. only looking up zone IDs defined")
		for _, zoneID := range p.zoneIDFilter.ZoneIDs {
			log.Debugf("looking up zone %q", zoneID)
			detailResponse, err := p.Client.GetZone(ctx, zoneID)
			if err != nil {
				log.Errorf("zone %q lookup failed, %v", zoneID, err)
				return result, convertCloudflareError(err)
			}
			log.WithFields(log.Fields{
				"zoneName": detailResponse.Name,
				"zoneID":   detailResponse.ID,
			}).Debugln("adding zone for consideration")
			result = append(result, *detailResponse)
		}
		return result, nil
	}

	log.Debugln("no zoneIDFilter configured, looking at all zones")

	params := listZonesV4Params()
	iter := p.Client.ListZones(ctx, params)
	for zone := range autoPagerIterator(iter) {
		if !p.domainFilter.Match(zone.Name) {
			log.Debugf("zone %q not in domain filter", zone.Name)
			continue
		}
		log.WithFields(log.Fields{
			"zoneName": zone.Name,
			"zoneID":   zone.ID,
		}).Debugln("adding zone for consideration")
		result = append(result, zone)
	}
	if iter.Err() != nil {
		return nil, convertCloudflareError(iter.Err())
	}

	return result, nil
}

// Records returns the list of records.
func (p *CloudFlareProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint
	for _, zone := range zones {
		records, err := p.getDNSRecordsMap(ctx, zone.ID)
		if err != nil {
			return nil, err
		}

		// nil if custom hostnames are not enabled
		chs, chErr := p.listCustomHostnamesWithPagination(ctx, zone.ID)
		if chErr != nil {
			return nil, chErr
		}

		// As CloudFlare does not support "sets" of targets, but instead returns
		// a single entry for each name/type/target, we have to group by name
		// and record to allow the planner to calculate the correct plan. See #992.
		zoneEndpoints := p.groupByNameAndTypeWithCustomHostnames(records, chs)

		if err := p.addEnpointsProviderSpecificRegionKeyProperty(ctx, zone.ID, zoneEndpoints); err != nil {
			return nil, err
		}

		endpoints = append(endpoints, zoneEndpoints...)
	}

	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *CloudFlareProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	var cloudflareChanges []*cloudFlareChange

	// if custom hostnames are enabled, deleting first allows to avoid conflicts with the new ones
	if p.CustomHostnamesConfig.Enabled {
		for _, e := range changes.Delete {
			for _, target := range e.Targets {
				change, err := p.newCloudFlareChange(cloudFlareDelete, e, target, nil)
				if err != nil {
					log.Errorf("failed to create cloudflare change: %v", err)
					continue
				}
				cloudflareChanges = append(cloudflareChanges, change)
			}
		}
	}

	for _, e := range changes.Create {
		for _, target := range e.Targets {
			change, err := p.newCloudFlareChange(cloudFlareCreate, e, target, nil)
			if err != nil {
				log.Errorf("failed to create cloudflare change: %v", err)
				continue
			}
			cloudflareChanges = append(cloudflareChanges, change)
		}
	}

	for i, desired := range changes.UpdateNew {
		current := changes.UpdateOld[i]

		add, remove, leave := provider.Difference(current.Targets, desired.Targets)

		for _, a := range remove {
			change, err := p.newCloudFlareChange(cloudFlareDelete, current, a, current)
			if err != nil {
				log.Errorf("failed to create cloudflare change: %v", err)
				continue
			}
			cloudflareChanges = append(cloudflareChanges, change)
		}

		for _, a := range add {
			change, err := p.newCloudFlareChange(cloudFlareCreate, desired, a, current)
			if err != nil {
				log.Errorf("failed to create cloudflare change: %v", err)
				continue
			}
			cloudflareChanges = append(cloudflareChanges, change)
		}

		for _, a := range leave {
			change, err := p.newCloudFlareChange(cloudFlareUpdate, desired, a, current)
			if err != nil {
				log.Errorf("failed to create cloudflare change: %v", err)
				continue
			}
			cloudflareChanges = append(cloudflareChanges, change)
		}
	}

	// TODO: consider deleting before creating even if custom hostnames are not in use
	if !p.CustomHostnamesConfig.Enabled {
		for _, e := range changes.Delete {
			for _, target := range e.Targets {
				change, err := p.newCloudFlareChange(cloudFlareDelete, e, target, nil)
				if err != nil {
					log.Errorf("failed to create cloudflare change: %v", err)
					continue
				}
				cloudflareChanges = append(cloudflareChanges, change)
			}
		}
	}

	return p.submitChanges(ctx, cloudflareChanges)
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *CloudFlareProvider) submitChanges(ctx context.Context, changes []*cloudFlareChange) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		log.Info("All records are already up to date")
		return nil
	}

	zones, err := p.Zones(ctx)
	if err != nil {
		return err
	}
	// separate into per-zone change sets to be passed to the API.
	changesByZone := p.changesByZone(zones, changes)

	var failedZones []string
	for zoneID, zoneChanges := range changesByZone {
		var failedChange bool

		for _, change := range zoneChanges {
			logFields := log.Fields{
				"record": change.ResourceRecord.Name,
				"type":   change.ResourceRecord.Type,
				"ttl":    change.ResourceRecord.TTL,
				"action": change.Action.String(),
				"zone":   zoneID,
			}

			log.WithFields(logFields).Info("Changing record.")

			if p.DryRun {
				continue
			}

			records, err := p.getDNSRecordsMap(ctx, zoneID)
			if err != nil {
				return fmt.Errorf("could not fetch records from zone, %w", err)
			}
			chs, chErr := p.listCustomHostnamesWithPagination(ctx, zoneID)
			if chErr != nil {
				return fmt.Errorf("could not fetch custom hostnames from zone, %w", chErr)
			}
			switch change.Action {
			case cloudFlareUpdate:
				if !p.submitCustomHostnameChanges(ctx, zoneID, change, chs, logFields) {
					failedChange = true
				}
				recordID := p.getRecordID(records, change.ResourceRecord)
				if recordID == "" {
					log.WithFields(logFields).Errorf("failed to find previous record: %v", change.ResourceRecord)
					continue
				}
				recordParam := getUpdateDNSRecordParam(zoneID, *change)
				_, err := p.Client.UpdateDNSRecord(ctx, recordID, recordParam)
				if err != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to update record: %v", err)
				}
			case cloudFlareDelete:
				recordID := p.getRecordID(records, change.ResourceRecord)
				if recordID == "" {
					log.WithFields(logFields).Errorf("failed to find previous record: %v", change.ResourceRecord)
					continue
				}
				err := p.Client.DeleteDNSRecord(ctx, recordID, dns.RecordDeleteParams{ZoneID: cloudflare.F(zoneID)})
				if err != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to delete record: %v", err)
				}
				if !p.submitCustomHostnameChanges(ctx, zoneID, change, chs, logFields) {
					failedChange = true
				}
			case cloudFlareCreate:
				recordParam := getCreateDNSRecordParam(zoneID, change)
				_, err := p.Client.CreateDNSRecord(ctx, recordParam)
				if err != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to create record: %v", err)
				}
				if !p.submitCustomHostnameChanges(ctx, zoneID, change, chs, logFields) {
					failedChange = true
				}
			}
		}

		if p.RegionalServicesConfig.Enabled {
			desiredRegionalHostnames, err := desiredRegionalHostnames(zoneChanges)
			if err != nil {
				return fmt.Errorf("failed to build desired regional hostnames: %w", err)
			}
			if len(desiredRegionalHostnames) > 0 {
				regionalHostnames, err := p.listDataLocalisationRegionalHostnames(ctx, zoneID)
				if err != nil {
					return fmt.Errorf("could not fetch regional hostnames from zone, %w", err)
				}
				regionalHostnamesChanges := regionalHostnamesChanges(desiredRegionalHostnames, regionalHostnames)
				if !p.submitRegionalHostnameChanges(ctx, zoneID, regionalHostnamesChanges) {
					failedChange = true
				}
			}
		}

		if failedChange {
			failedZones = append(failedZones, zoneID)
		}
	}

	if len(failedZones) > 0 {
		return provider.NewSoftErrorf("failed to submit all changes for the following zones: %q", failedZones)
	}

	return nil
}

// parseTagsAnnotation is the single helper method to handle tags from the annotation string.
// It splits the string, cleans up whitespace, and sorts the tags to create a canonical representation.
func parseTagsAnnotation(tagString string) []string {
	tags := strings.Split(tagString, ",")
	cleanedTags := make([]string, 0, len(tags))
	for _, tag := range tags {
		trimmed := strings.TrimSpace(tag)
		if trimmed != "" {
			cleanedTags = append(cleanedTags, trimmed)
		}
	}
	sort.Strings(cleanedTags)
	return cleanedTags
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (p *CloudFlareProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	var adjustedEndpoints []*endpoint.Endpoint
	for _, e := range endpoints {
		proxied := shouldBeProxied(e, p.proxiedByDefault)
		if proxied {
			e.RecordTTL = 0
		}
		e.SetProviderSpecificProperty(annotations.CloudflareProxiedKey, strconv.FormatBool(proxied))

		if p.CustomHostnamesConfig.Enabled {
			// sort custom hostnames in annotation to properly detect changes
			if customHostnames := getEndpointCustomHostnames(e); len(customHostnames) > 1 {
				sort.Strings(customHostnames)
				e.SetProviderSpecificProperty(annotations.CloudflareCustomHostnameKey, strings.Join(customHostnames, ","))
			}
		} else {
			// ignore custom hostnames annotations if not enabled
			e.DeleteProviderSpecificProperty(annotations.CloudflareCustomHostnameKey)
		}

		if val, ok := e.GetProviderSpecificProperty(annotations.CloudflareTagsKey); ok {
			sortedTags := parseTagsAnnotation(val)
			e.SetProviderSpecificProperty(annotations.CloudflareTagsKey, strings.Join(sortedTags, ","))
		}

		p.adjustEndpointProviderSpecificRegionKeyProperty(e)

		if p.DNSRecordsConfig.Comment != "" {
			if _, found := e.GetProviderSpecificProperty(annotations.CloudflareRecordCommentKey); !found {
				e.SetProviderSpecificProperty(annotations.CloudflareRecordCommentKey, p.DNSRecordsConfig.Comment)
			}
		}

		adjustedEndpoints = append(adjustedEndpoints, e)
	}
	return adjustedEndpoints, nil
}

// changesByZone separates a multi-zone change into a single change per zone.
func (p *CloudFlareProvider) changesByZone(zones []zones.Zone, changeSet []*cloudFlareChange) map[string][]*cloudFlareChange {
	changes := make(map[string][]*cloudFlareChange)
	zoneNameIDMapper := provider.ZoneIDName{}

	for _, z := range zones {
		zoneNameIDMapper.Add(z.ID, z.Name)
		changes[z.ID] = []*cloudFlareChange{}
	}

	for _, c := range changeSet {
		zoneID, _ := zoneNameIDMapper.FindZone(c.ResourceRecord.Name)
		if zoneID == "" {
			log.Debugf("Skipping record %q because no hosted zone matching record DNS Name was detected", c.ResourceRecord.Name)
			continue
		}
		changes[zoneID] = append(changes[zoneID], c)
	}

	return changes
}

func (p *CloudFlareProvider) getRecordID(records DNSRecordsMap, record dns.RecordResponse) string {
	if zoneRecord, ok := records[DNSRecordIndex{Name: record.Name, Type: string(record.Type), Content: record.Content}]; ok {
		return zoneRecord.ID
	}
	return ""
}

func (p *CloudFlareProvider) newCloudFlareChange(action changeAction, ep *endpoint.Endpoint, target string, current *endpoint.Endpoint) (*cloudFlareChange, error) {
	ttl := dns.TTL(defaultTTL)
	proxied := shouldBeProxied(ep, p.proxiedByDefault)

	if ep.RecordTTL.IsConfigured() {
		ttl = dns.TTL(ep.RecordTTL)
	}

	prevCustomHostnames := []string{}
	newCustomHostnames := map[string]customHostname{}
	if p.CustomHostnamesConfig.Enabled {
		if current != nil {
			prevCustomHostnames = getEndpointCustomHostnames(current)
		}
		for _, v := range getEndpointCustomHostnames(ep) {
			newCustomHostnames[v] = p.newCustomHostname(v, ep.DNSName)
		}
	}

	// Load comment from program flag
	comment := p.DNSRecordsConfig.Comment
	if val, ok := ep.GetProviderSpecificProperty(annotations.CloudflareRecordCommentKey); ok {
		// Replace comment with Ingress annotation
		comment = val
	}

	var tags []string
	if val, ok := ep.GetProviderSpecificProperty(annotations.CloudflareTagsKey); ok {
		tags = parseTagsAnnotation(val)
	}

	if len(comment) > freeZoneMaxCommentLength {
		comment = p.DNSRecordsConfig.trimAndValidateComment(ep.DNSName, comment, p.ZoneHasPaidPlan)
	}

	var priority float64
	if ep.RecordType == "MX" {
		mxRecord, err := endpoint.NewMXRecord(target)
		if err != nil {
			return &cloudFlareChange{}, fmt.Errorf("failed to parse MX record target %q: %w", target, err)
		} else {
			priority = float64(*mxRecord.GetPriority())
			target = *mxRecord.GetHost()
		}
	}

	return &cloudFlareChange{
		Action: action,
		ResourceRecord: dns.RecordResponse{
			Name:     ep.DNSName,
			TTL:      ttl,
			Proxied:  proxied,
			Type:     dns.RecordResponseType(ep.RecordType),
			Content:  target,
			Comment:  comment,
			Tags:     tags,
			Priority: priority,
		},
		RegionalHostname:    p.regionalHostname(ep),
		CustomHostnamesPrev: prevCustomHostnames,
		CustomHostnames:     newCustomHostnames,
	}, nil
}

func newDNSRecordIndex(r dns.RecordResponse) DNSRecordIndex {
	return DNSRecordIndex{Name: r.Name, Type: string(r.Type), Content: r.Content}
}

// getDNSRecordsMap retrieves all DNS records for a given zone and returns them as a DNSRecordsMap.
func (p *CloudFlareProvider) getDNSRecordsMap(ctx context.Context, zoneID string) (DNSRecordsMap, error) {
	// for faster getRecordID lookup
	recordsMap := make(DNSRecordsMap)
	params := dns.RecordListParams{ZoneID: cloudflare.F(zoneID)}
	if p.DNSRecordsConfig.PerPage > 0 {
		params.PerPage = cloudflare.F(float64(p.DNSRecordsConfig.PerPage))
	}
	iter := p.Client.ListDNSRecords(ctx, params)
	for record := range autoPagerIterator(iter) {
		recordsMap[newDNSRecordIndex(record)] = record
	}
	if iter.Err() != nil {
		return nil, convertCloudflareError(iter.Err())
	}
	return recordsMap, nil
}

func shouldBeProxied(ep *endpoint.Endpoint, proxiedByDefault bool) bool {
	proxied := proxiedByDefault

	for _, v := range ep.ProviderSpecific {
		if v.Name == annotations.CloudflareProxiedKey {
			b, err := strconv.ParseBool(v.Value)
			if err != nil {
				log.Errorf("Failed to parse annotation [%q]: %v", annotations.CloudflareProxiedKey, err)
			} else {
				proxied = b
			}
			break
		}
	}

	if recordTypeProxyNotSupported[ep.RecordType] {
		proxied = false
	}
	return proxied
}

func getEndpointCustomHostnames(ep *endpoint.Endpoint) []string {
	for _, v := range ep.ProviderSpecific {
		if v.Name == annotations.CloudflareCustomHostnameKey {
			customHostnames := strings.Split(v.Value, ",")
			return customHostnames
		}
	}
	return []string{}
}

func (p *CloudFlareProvider) groupByNameAndTypeWithCustomHostnames(records DNSRecordsMap, chs customHostnamesMap) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	// group supported records by name and type
	groups := map[string][]dns.RecordResponse{}

	for _, r := range records {
		if !p.SupportedAdditionalRecordTypes(string(r.Type)) {
			continue
		}

		groupBy := r.Name + string(r.Type)
		if _, ok := groups[groupBy]; !ok {
			groups[groupBy] = []dns.RecordResponse{}
		}

		groups[groupBy] = append(groups[groupBy], r)
	}

	// map custom origin to custom hostname, custom origin should match to a dns record
	customHostnames := map[string][]string{}

	for _, c := range chs {
		customHostnames[c.customOriginServer] = append(customHostnames[c.customOriginServer], c.hostname)
	}

	// create a single endpoint with all the targets for each name/type
	for _, records := range groups {
		if len(records) == 0 {
			return endpoints
		}
		targets := make([]string, len(records))
		for i, record := range records {
			if records[i].Type == "MX" {
				targets[i] = fmt.Sprintf("%v %v", record.Priority, record.Content)
			} else {
				targets[i] = record.Content
			}
		}
		e := endpoint.NewEndpointWithTTL(
			records[0].Name,
			string(records[0].Type),
			endpoint.TTL(records[0].TTL),
			targets...)
		proxied := records[0].Proxied
		if e == nil {
			continue
		}
		e = e.WithProviderSpecific(annotations.CloudflareProxiedKey, strconv.FormatBool(proxied))
		// noop (customHostnames is empty) if custom hostnames feature is not in use
		if customHostnames, ok := customHostnames[records[0].Name]; ok {
			sort.Strings(customHostnames)
			e = e.WithProviderSpecific(annotations.CloudflareCustomHostnameKey, strings.Join(customHostnames, ","))
		}

		if records[0].Comment != "" {
			e = e.WithProviderSpecific(annotations.CloudflareRecordCommentKey, records[0].Comment)
		}

		if records[0].Tags != nil {
			if tags, ok := records[0].Tags.([]string); ok && len(tags) > 0 {
				sort.Strings(tags)
				e = e.WithProviderSpecific(annotations.CloudflareTagsKey, strings.Join(tags, ","))
			}
		}

		endpoints = append(endpoints, e)
	}
	return endpoints
}

// SupportedRecordType returns true if the record type is supported by the provider
func (p *CloudFlareProvider) SupportedAdditionalRecordTypes(recordType string) bool {
	switch recordType {
	case endpoint.RecordTypeMX:
		return true
	default:
		return provider.SupportedRecordType(recordType)
	}
}
