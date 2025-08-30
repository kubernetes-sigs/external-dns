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
	"maps"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	cloudflarev0 "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/addressing"
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

// for faster getCustomHostname() lookup
type CustomHostnameIndex struct {
	Hostname string
}

type CustomHostnamesMap map[CustomHostnameIndex]cloudflarev0.CustomHostname

var recordTypeProxyNotSupported = map[string]bool{
	"LOC": true,
	"MX":  true,
	"NS":  true,
	"SPF": true,
	"TXT": true,
	"SRV": true,
}

type CustomHostnamesConfig struct {
	Enabled              bool
	MinTLSVersion        string
	CertificateAuthority string
}

var recordTypeCustomHostnameSupported = map[string]bool{
	"A":     true,
	"CNAME": true,
}

// cloudFlareDNS is the subset of the CloudFlare API that we actually use.  Add methods as required. Signatures must match exactly.
type cloudFlareDNS interface {
	ZoneIDByName(zoneName string) (string, error)
	ListZones(ctx context.Context, params zones.ZoneListParams) autoPager[zones.Zone]
	GetZone(ctx context.Context, zoneID string) (*zones.Zone, error)
	ListDNSRecords(ctx context.Context, rc *cloudflarev0.ResourceContainer, rp cloudflarev0.ListDNSRecordsParams) ([]dns.RecordResponse, *cloudflarev0.ResultInfo, error)
	CreateDNSRecord(ctx context.Context, params dns.RecordNewParams) (*dns.RecordResponse, error)
	DeleteDNSRecord(ctx context.Context, recordID string, params dns.RecordDeleteParams) error
	UpdateDNSRecord(ctx context.Context, rc *cloudflarev0.ResourceContainer, rp cloudflarev0.UpdateDNSRecordParams) error
	ListDataLocalizationRegionalHostnames(ctx context.Context, params addressing.RegionalHostnameListParams) autoPager[addressing.RegionalHostnameListResponse]
	CreateDataLocalizationRegionalHostname(ctx context.Context, params addressing.RegionalHostnameNewParams) error
	UpdateDataLocalizationRegionalHostname(ctx context.Context, hostname string, params addressing.RegionalHostnameEditParams) error
	DeleteDataLocalizationRegionalHostname(ctx context.Context, hostname string, params addressing.RegionalHostnameDeleteParams) error
	CustomHostnames(ctx context.Context, zoneID string, page int, filter cloudflarev0.CustomHostname) ([]cloudflarev0.CustomHostname, cloudflarev0.ResultInfo, error)
	DeleteCustomHostname(ctx context.Context, zoneID string, customHostnameID string) error
	CreateCustomHostname(ctx context.Context, zoneID string, ch cloudflarev0.CustomHostname) (*cloudflarev0.CustomHostnameResponse, error)
}

type zoneService struct {
	serviceV0 *cloudflarev0.API
	service   *cloudflare.Client
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

func (z zoneService) ListDNSRecords(ctx context.Context, rc *cloudflarev0.ResourceContainer, rp cloudflarev0.ListDNSRecordsParams) ([]dns.RecordResponse, *cloudflarev0.ResultInfo, error) {
	records, info, err := z.serviceV0.ListDNSRecords(ctx, rc, rp)
	convertedRecords := make([]dns.RecordResponse, 0, len(records))
	for _, record := range records {
		convertedRecords = append(convertedRecords, dnsRecordResponseFromLegacyDNSRecord(record))
	}
	return convertedRecords, info, err
}

func (z zoneService) UpdateDNSRecord(ctx context.Context, rc *cloudflarev0.ResourceContainer, rp cloudflarev0.UpdateDNSRecordParams) error {
	_, err := z.serviceV0.UpdateDNSRecord(ctx, rc, rp)
	return err
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

func (z zoneService) CustomHostnames(ctx context.Context, zoneID string, page int, filter cloudflarev0.CustomHostname) ([]cloudflarev0.CustomHostname, cloudflarev0.ResultInfo, error) {
	return z.serviceV0.CustomHostnames(ctx, zoneID, page, filter)
}

func (z zoneService) DeleteCustomHostname(ctx context.Context, zoneID string, customHostnameID string) error {
	return z.serviceV0.DeleteCustomHostname(ctx, zoneID, customHostnameID)
}

func (z zoneService) CreateCustomHostname(ctx context.Context, zoneID string, ch cloudflarev0.CustomHostname) (*cloudflarev0.CustomHostnameResponse, error) {
	return z.serviceV0.CreateCustomHostname(ctx, zoneID, ch)
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
	CustomHostnames     map[string]cloudflarev0.CustomHostname
	CustomHostnamesPrev []string
}

// RecordParamsTypes is a typeset of the possible Record Params that can be passed to cloudflare-go library
type RecordParamsTypes interface {
	cloudflarev0.UpdateDNSRecordParams | cloudflarev0.CreateDNSRecordParams
}

// updateDNSRecordParam is a function that returns the appropriate Record Param based on the cloudFlareChange passed in
func updateDNSRecordParam(cfc cloudFlareChange) cloudflarev0.UpdateDNSRecordParams {
	priority := uint16(cfc.ResourceRecord.Priority)

	params := cloudflarev0.UpdateDNSRecordParams{
		Name:     cfc.ResourceRecord.Name,
		TTL:      int(cfc.ResourceRecord.TTL),
		Proxied:  &cfc.ResourceRecord.Proxied,
		Type:     string(cfc.ResourceRecord.Type),
		Content:  cfc.ResourceRecord.Content,
		Priority: &priority,
		Comment:  cloudflarev0.StringPtr(cfc.ResourceRecord.Comment),
	}

	return params
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
		},
	}
}

// getDeleteDNSRecordParam is a function that returns the appropriate Record Delete Param based on the zoneID passed in
func getDeleteDNSRecordParam(zoneID string) dns.RecordDeleteParams {
	return dns.RecordDeleteParams{
		ZoneID: cloudflare.F(zoneID),
	}
}

func convertCloudflareError(err error) error {
	var apiErr *cloudflarev0.Error
	if errors.As(err, &apiErr) {
		if apiErr.ClientRateLimited() || apiErr.StatusCode >= http.StatusInternalServerError {
			// Handle rate limit error as a soft error
			return provider.NewSoftError(err)
		}
	}
	// This is a workaround because Cloudflare library does not return a specific error type for rate limit exceeded.
	// See https://github.com/cloudflare/cloudflare-go/issues/4155 and https://github.com/kubernetes-sigs/external-dns/pull/5524
	// This workaround can be removed once Cloudflare library returns a specific error type.
	if strings.Contains(err.Error(), "exceeded available rate limit retries") {
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
	var (
		config   *cloudflarev0.API
		configV4 *cloudflare.Client
		err      error
	)
	if os.Getenv("CF_API_TOKEN") != "" {
		token := os.Getenv("CF_API_TOKEN")
		if strings.HasPrefix(token, "file:") {
			tokenBytes, err := os.ReadFile(strings.TrimPrefix(token, "file:"))
			if err != nil {
				return nil, fmt.Errorf("failed to read CF_API_TOKEN from file: %w", err)
			}
			token = strings.TrimSpace(string(tokenBytes))
		}
		config, err = cloudflarev0.NewWithAPIToken(token)
		configV4 = cloudflare.NewClient(
			option.WithAPIToken(token),
		)
	} else {
		config, err = cloudflarev0.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
		configV4 = cloudflare.NewClient(
			option.WithAPIKey(os.Getenv("CF_API_KEY")),
			option.WithAPIEmail(os.Getenv("CF_API_EMAIL")),
		)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloudflare provider: %w", err)
	}

	if regionalServicesConfig.RegionKey != "" {
		regionalServicesConfig.Enabled = true
	}

	return &CloudFlareProvider{
		Client:                 zoneService{config, configV4},
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
		records, err := p.listDNSRecordsWithAutoPagination(ctx, zone.ID)
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

// submitCustomHostnameChanges implements Custom Hostname functionality for the Change, returns false if it fails
func (p *CloudFlareProvider) submitCustomHostnameChanges(ctx context.Context, zoneID string, change *cloudFlareChange, chs CustomHostnamesMap, logFields log.Fields) bool {
	failedChange := false
	// return early if disabled
	if !p.CustomHostnamesConfig.Enabled {
		return true
	}

	switch change.Action {
	case cloudFlareUpdate:
		if recordTypeCustomHostnameSupported[string(change.ResourceRecord.Type)] {
			add, remove, _ := provider.Difference(change.CustomHostnamesPrev, slices.Collect(maps.Keys(change.CustomHostnames)))

			for _, changeCH := range remove {
				if prevCh, err := getCustomHostname(chs, changeCH); err == nil {
					prevChID := prevCh.ID
					if prevChID != "" {
						log.WithFields(logFields).Infof("Removing previous custom hostname %q/%q", prevChID, changeCH)
						chErr := p.Client.DeleteCustomHostname(ctx, zoneID, prevChID)
						if chErr != nil {
							failedChange = true
							log.WithFields(logFields).Errorf("failed to remove previous custom hostname %q/%q: %v", prevChID, changeCH, chErr)
						}
					}
				}
			}
			for _, changeCH := range add {
				log.WithFields(logFields).Infof("Adding custom hostname %q", changeCH)
				_, chErr := p.Client.CreateCustomHostname(ctx, zoneID, change.CustomHostnames[changeCH])
				if chErr != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to add custom hostname %q: %v", changeCH, chErr)
				}
			}
		}
	case cloudFlareDelete:
		for _, changeCH := range change.CustomHostnames {
			if recordTypeCustomHostnameSupported[string(change.ResourceRecord.Type)] && changeCH.Hostname != "" {
				log.WithFields(logFields).Infof("Deleting custom hostname %q", changeCH.Hostname)
				if ch, err := getCustomHostname(chs, changeCH.Hostname); err == nil {
					chID := ch.ID
					chErr := p.Client.DeleteCustomHostname(ctx, zoneID, chID)
					if chErr != nil {
						failedChange = true
						log.WithFields(logFields).Errorf("failed to delete custom hostname %q/%q: %v", chID, changeCH.Hostname, chErr)
					}
				} else {
					log.WithFields(logFields).Warnf("failed to delete custom hostname %q: %v", changeCH.Hostname, err)
				}
			}
		}
	case cloudFlareCreate:
		for _, changeCH := range change.CustomHostnames {
			if recordTypeCustomHostnameSupported[string(change.ResourceRecord.Type)] && changeCH.Hostname != "" {
				log.WithFields(logFields).Infof("Creating custom hostname %q", changeCH.Hostname)
				if ch, err := getCustomHostname(chs, changeCH.Hostname); err == nil {
					if changeCH.CustomOriginServer == ch.CustomOriginServer {
						log.WithFields(logFields).Warnf("custom hostname %q already exists with the same origin %q, continue", changeCH.Hostname, ch.CustomOriginServer)
					} else {
						failedChange = true
						log.WithFields(logFields).Errorf("failed to create custom hostname, %q already exists with origin %q", changeCH.Hostname, ch.CustomOriginServer)
					}
				} else {
					_, chErr := p.Client.CreateCustomHostname(ctx, zoneID, changeCH)
					if chErr != nil {
						failedChange = true
						log.WithFields(logFields).Errorf("failed to create custom hostname %q: %v", changeCH.Hostname, chErr)
					}
				}
			}
		}
	}
	return !failedChange
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
		resourceContainer := cloudflarev0.ZoneIdentifier(zoneID)

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

			records, err := p.listDNSRecordsWithAutoPagination(ctx, zoneID)
			if err != nil {
				return fmt.Errorf("could not fetch records from zone, %w", err)
			}
			chs, chErr := p.listCustomHostnamesWithPagination(ctx, zoneID)
			if chErr != nil {
				return fmt.Errorf("could not fetch custom hostnames from zone, %w", chErr)
			}
			if change.Action == cloudFlareUpdate {
				if !p.submitCustomHostnameChanges(ctx, zoneID, change, chs, logFields) {
					failedChange = true
				}
				recordID := p.getRecordID(records, change.ResourceRecord)
				if recordID == "" {
					log.WithFields(logFields).Errorf("failed to find previous record: %v", change.ResourceRecord)
					continue
				}
				recordParam := updateDNSRecordParam(*change)
				recordParam.ID = recordID
				err := p.Client.UpdateDNSRecord(ctx, resourceContainer, recordParam)
				if err != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to update record: %v", err)
				}
			} else if change.Action == cloudFlareDelete {
				recordID := p.getRecordID(records, change.ResourceRecord)
				if recordID == "" {
					log.WithFields(logFields).Errorf("failed to find previous record: %v", change.ResourceRecord)
					continue
				}
				recordParam := getDeleteDNSRecordParam(zoneID)
				err := p.Client.DeleteDNSRecord(ctx, recordID, recordParam)
				if err != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to delete record: %v", err)
				}
				if !p.submitCustomHostnameChanges(ctx, zoneID, change, chs, logFields) {
					failedChange = true
				}
			} else if change.Action == cloudFlareCreate {
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
		return fmt.Errorf("failed to submit all changes for the following zones: %q", failedZones)
	}

	return nil
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

		if p.RegionalServicesConfig.Enabled {
			// Add default region key if not set
			if _, ok := e.GetProviderSpecificProperty(annotations.CloudflareRegionKey); !ok {
				e.SetProviderSpecificProperty(annotations.CloudflareRegionKey, p.RegionalServicesConfig.RegionKey)
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

func getCustomHostname(chs CustomHostnamesMap, chName string) (cloudflarev0.CustomHostname, error) {
	if chName == "" {
		return cloudflarev0.CustomHostname{}, fmt.Errorf("failed to get custom hostname: %q is empty", chName)
	}
	if ch, ok := chs[CustomHostnameIndex{Hostname: chName}]; ok {
		return ch, nil
	}
	return cloudflarev0.CustomHostname{}, fmt.Errorf("failed to get custom hostname: %q not found", chName)
}

func (p *CloudFlareProvider) newCustomHostname(customHostname string, origin string) cloudflarev0.CustomHostname {
	return cloudflarev0.CustomHostname{
		Hostname:           customHostname,
		CustomOriginServer: origin,
		SSL:                getCustomHostnamesSSLOptions(p.CustomHostnamesConfig),
	}
}

func (p *CloudFlareProvider) newCloudFlareChange(action changeAction, ep *endpoint.Endpoint, target string, current *endpoint.Endpoint) (*cloudFlareChange, error) {
	ttl := dns.TTL(defaultTTL)
	proxied := shouldBeProxied(ep, p.proxiedByDefault)

	if ep.RecordTTL.IsConfigured() {
		ttl = dns.TTL(ep.RecordTTL)
	}

	prevCustomHostnames := []string{}
	newCustomHostnames := map[string]cloudflarev0.CustomHostname{}
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

// listDNSRecordsWithAutoPagination performs automatic pagination of results on requests to cloudflare.ListDNSRecords with custom per_page values
func (p *CloudFlareProvider) listDNSRecordsWithAutoPagination(ctx context.Context, zoneID string) (DNSRecordsMap, error) {
	// for faster getRecordID lookup
	records := make(DNSRecordsMap)
	resultInfo := cloudflarev0.ResultInfo{PerPage: p.DNSRecordsConfig.PerPage, Page: 1}
	params := cloudflarev0.ListDNSRecordsParams{ResultInfo: resultInfo}
	for {
		pageRecords, resultInfo, err := p.Client.ListDNSRecords(ctx, cloudflarev0.ZoneIdentifier(zoneID), params)
		if err != nil {
			return nil, convertCloudflareError(err)
		}

		for _, r := range pageRecords {
			records[newDNSRecordIndex(r)] = r
		}
		params.ResultInfo = resultInfo.Next()
		if params.Done() {
			break
		}
	}
	return records, nil
}

func newCustomHostnameIndex(ch cloudflarev0.CustomHostname) CustomHostnameIndex {
	return CustomHostnameIndex{Hostname: ch.Hostname}
}

// listCustomHostnamesWithPagination performs automatic pagination of results on requests to cloudflare.CustomHostnames
func (p *CloudFlareProvider) listCustomHostnamesWithPagination(ctx context.Context, zoneID string) (CustomHostnamesMap, error) {
	if !p.CustomHostnamesConfig.Enabled {
		return nil, nil
	}
	chs := make(CustomHostnamesMap)
	resultInfo := cloudflarev0.ResultInfo{Page: 1}
	for {
		pageCustomHostnameListResponse, result, err := p.Client.CustomHostnames(ctx, zoneID, resultInfo.Page, cloudflarev0.CustomHostname{})
		if err != nil {
			convertedError := convertCloudflareError(err)
			if !errors.Is(convertedError, provider.SoftError) {
				log.Errorf("zone %q failed to fetch custom hostnames. Please check if \"Cloudflare for SaaS\" is enabled and API key permissions, %v", zoneID, err)
			}
			return nil, convertedError
		}
		for _, ch := range pageCustomHostnameListResponse {
			chs[newCustomHostnameIndex(ch)] = ch
		}
		resultInfo = result.Next()
		if resultInfo.Done() {
			break
		}
	}
	return chs, nil
}

func getCustomHostnamesSSLOptions(customHostnamesConfig CustomHostnamesConfig) *cloudflarev0.CustomHostnameSSL {
	ssl := &cloudflarev0.CustomHostnameSSL{
		Type:         "dv",
		Method:       "http",
		BundleMethod: "ubiquitous",
		Settings: cloudflarev0.CustomHostnameSSLSettings{
			MinTLSVersion: customHostnamesConfig.MinTLSVersion,
		},
	}
	// Set CertificateAuthority if provided
	// We're not able to set it at all (even with a blank) if you're not on an enterprise plan
	if customHostnamesConfig.CertificateAuthority != "none" {
		ssl.CertificateAuthority = customHostnamesConfig.CertificateAuthority
	}
	return ssl
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

func (p *CloudFlareProvider) groupByNameAndTypeWithCustomHostnames(records DNSRecordsMap, chs CustomHostnamesMap) []*endpoint.Endpoint {
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
		customHostnames[c.CustomOriginServer] = append(customHostnames[c.CustomOriginServer], c.Hostname)
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

func dnsRecordResponseFromLegacyDNSRecord(record cloudflarev0.DNSRecord) dns.RecordResponse {
	var priority float64
	if record.Priority != nil {
		priority = float64(*record.Priority)
	}

	return dns.RecordResponse{
		CreatedOn:  record.CreatedOn,
		ModifiedOn: record.ModifiedOn,
		Type:       dns.RecordResponseType(record.Type),
		Name:       record.Name,
		Content:    record.Content,
		Meta:       record.Meta,
		Data:       record.Data,
		ID:         record.ID,
		Priority:   priority,
		TTL:        dns.TTL(record.TTL),
		Proxied:    record.Proxied != nil && *record.Proxied,
		Proxiable:  record.Proxiable,
		Comment:    record.Comment,
		Tags:       record.Tags,
	}
}
