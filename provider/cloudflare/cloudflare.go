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

	cloudflare "github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/annotations"
)

const (
	// cloudFlareCreate is a ChangeAction enum value
	cloudFlareCreate = "CREATE"
	// cloudFlareDelete is a ChangeAction enum value
	cloudFlareDelete = "DELETE"
	// cloudFlareUpdate is a ChangeAction enum value
	cloudFlareUpdate = "UPDATE"
	// defaultCloudFlareRecordTTL 1 = automatic
	defaultCloudFlareRecordTTL = 1
)

// We have to use pointers to bools now, as the upstream cloudflare-go library requires them
// see: https://github.com/cloudflare/cloudflare-go/pull/595

// proxyEnabled is a pointer to a bool true showing the record should be proxied through cloudflare
var proxyEnabled *bool = boolPtr(true)

// proxyDisabled is a pointer to a bool false showing the record should not be proxied through cloudflare
var proxyDisabled *bool = boolPtr(false)

// for faster getRecordID() lookup
type DNSRecordIndex struct {
	Name    string
	Type    string
	Content string
}

type DNSRecordsMap map[DNSRecordIndex]cloudflare.DNSRecord

// for faster getCustomHostname() lookup
type CustomHostnameIndex struct {
	Hostname string
}

type CustomHostnamesMap map[CustomHostnameIndex]cloudflare.CustomHostname

type DataLocalizationRegionalHostnameChange struct {
	Action string
	cloudflare.RegionalHostname
}

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

var recordTypeRegionalHostnameSupported = map[string]bool{
	"A":     true,
	"AAAA":  true,
	"CNAME": true,
}

// cloudFlareDNS is the subset of the CloudFlare API that we actually use.  Add methods as required. Signatures must match exactly.
type cloudFlareDNS interface {
	UserDetails(ctx context.Context) (cloudflare.User, error)
	ZoneIDByName(zoneName string) (string, error)
	ListZones(ctx context.Context, zoneID ...string) ([]cloudflare.Zone, error)
	ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error)
	ZoneDetails(ctx context.Context, zoneID string) (cloudflare.Zone, error)
	ListDNSRecords(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.ListDNSRecordsParams) ([]cloudflare.DNSRecord, *cloudflare.ResultInfo, error)
	CreateDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.CreateDNSRecordParams) (cloudflare.DNSRecord, error)
	DeleteDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, recordID string) error
	UpdateDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.UpdateDNSRecordParams) error
	CreateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.CreateDataLocalizationRegionalHostnameParams) error
	UpdateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.UpdateDataLocalizationRegionalHostnameParams) error
	DeleteDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, hostname string) error
	CustomHostnames(ctx context.Context, zoneID string, page int, filter cloudflare.CustomHostname) ([]cloudflare.CustomHostname, cloudflare.ResultInfo, error)
	DeleteCustomHostname(ctx context.Context, zoneID string, customHostnameID string) error
	CreateCustomHostname(ctx context.Context, zoneID string, ch cloudflare.CustomHostname) (*cloudflare.CustomHostnameResponse, error)
}

type zoneService struct {
	service *cloudflare.API
}

func (z zoneService) UserDetails(ctx context.Context) (cloudflare.User, error) {
	return z.service.UserDetails(ctx)
}

func (z zoneService) ListZones(ctx context.Context, zoneID ...string) ([]cloudflare.Zone, error) {
	return z.service.ListZones(ctx, zoneID...)
}

func (z zoneService) ZoneIDByName(zoneName string) (string, error) {
	return z.service.ZoneIDByName(zoneName)
}

func (z zoneService) CreateDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.CreateDNSRecordParams) (cloudflare.DNSRecord, error) {
	return z.service.CreateDNSRecord(ctx, rc, rp)
}

func (z zoneService) ListDNSRecords(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.ListDNSRecordsParams) ([]cloudflare.DNSRecord, *cloudflare.ResultInfo, error) {
	return z.service.ListDNSRecords(ctx, rc, rp)
}

func (z zoneService) UpdateDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.UpdateDNSRecordParams) error {
	_, err := z.service.UpdateDNSRecord(ctx, rc, rp)
	return err
}

func (z zoneService) CreateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.CreateDataLocalizationRegionalHostnameParams) error {
	_, err := z.service.CreateDataLocalizationRegionalHostname(ctx, rc, rp)
	return err
}

func (z zoneService) UpdateDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, rp cloudflare.UpdateDataLocalizationRegionalHostnameParams) error {
	_, err := z.service.UpdateDataLocalizationRegionalHostname(ctx, rc, rp)
	return err
}

func (z zoneService) DeleteDataLocalizationRegionalHostname(ctx context.Context, rc *cloudflare.ResourceContainer, hostname string) error {
	return z.service.DeleteDataLocalizationRegionalHostname(ctx, rc, hostname)
}

func (z zoneService) DeleteDNSRecord(ctx context.Context, rc *cloudflare.ResourceContainer, recordID string) error {
	return z.service.DeleteDNSRecord(ctx, rc, recordID)
}

func (z zoneService) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return z.service.ListZonesContext(ctx, opts...)
}

func (z zoneService) ZoneDetails(ctx context.Context, zoneID string) (cloudflare.Zone, error) {
	return z.service.ZoneDetails(ctx, zoneID)
}

func (z zoneService) CustomHostnames(ctx context.Context, zoneID string, page int, filter cloudflare.CustomHostname) ([]cloudflare.CustomHostname, cloudflare.ResultInfo, error) {
	return z.service.CustomHostnames(ctx, zoneID, page, filter)
}

func (z zoneService) DeleteCustomHostname(ctx context.Context, zoneID string, customHostnameID string) error {
	return z.service.DeleteCustomHostname(ctx, zoneID, customHostnameID)
}

func (z zoneService) CreateCustomHostname(ctx context.Context, zoneID string, ch cloudflare.CustomHostname) (*cloudflare.CustomHostnameResponse, error) {
	return z.service.CreateCustomHostname(ctx, zoneID, ch)
}

// CloudFlareProvider is an implementation of Provider for CloudFlare DNS.
type CloudFlareProvider struct {
	provider.BaseProvider
	Client cloudFlareDNS
	// only consider hosted zones managing domains ending in this suffix
	domainFilter          endpoint.DomainFilter
	zoneIDFilter          provider.ZoneIDFilter
	proxiedByDefault      bool
	CustomHostnamesConfig CustomHostnamesConfig
	DryRun                bool
	DNSRecordsPerPage     int
	RegionKey             string
}

// cloudFlareChange differentiates between ChangActions
type cloudFlareChange struct {
	Action              string
	ResourceRecord      cloudflare.DNSRecord
	RegionalHostname    cloudflare.RegionalHostname
	CustomHostnames     map[string]cloudflare.CustomHostname
	CustomHostnamesPrev []string
}

// RecordParamsTypes is a typeset of the possible Record Params that can be passed to cloudflare-go library
type RecordParamsTypes interface {
	cloudflare.UpdateDNSRecordParams | cloudflare.CreateDNSRecordParams
}

// updateDNSRecordParam is a function that returns the appropriate Record Param based on the cloudFlareChange passed in
func updateDNSRecordParam(cfc cloudFlareChange) cloudflare.UpdateDNSRecordParams {
	return cloudflare.UpdateDNSRecordParams{
		Name:    cfc.ResourceRecord.Name,
		TTL:     cfc.ResourceRecord.TTL,
		Proxied: cfc.ResourceRecord.Proxied,
		Type:    cfc.ResourceRecord.Type,
		Content: cfc.ResourceRecord.Content,
	}
}

// createDataLocalizationRegionalHostnameParams is a function that returns the appropriate RegionalHostname Param based on the cloudFlareChange passed in
func createDataLocalizationRegionalHostnameParams(rhc DataLocalizationRegionalHostnameChange) cloudflare.CreateDataLocalizationRegionalHostnameParams {
	return cloudflare.CreateDataLocalizationRegionalHostnameParams{
		Hostname:  rhc.Hostname,
		RegionKey: rhc.RegionKey,
	}
}

// updateDataLocalizationRegionalHostnameParams is a function that returns the appropriate RegionalHostname Param based on the cloudFlareChange passed in
func updateDataLocalizationRegionalHostnameParams(rhc DataLocalizationRegionalHostnameChange) cloudflare.UpdateDataLocalizationRegionalHostnameParams {
	return cloudflare.UpdateDataLocalizationRegionalHostnameParams{
		Hostname:  rhc.Hostname,
		RegionKey: rhc.RegionKey,
	}
}

// getCreateDNSRecordParam is a function that returns the appropriate Record Param based on the cloudFlareChange passed in
func getCreateDNSRecordParam(cfc cloudFlareChange) cloudflare.CreateDNSRecordParams {
	return cloudflare.CreateDNSRecordParams{
		Name:    cfc.ResourceRecord.Name,
		TTL:     cfc.ResourceRecord.TTL,
		Proxied: cfc.ResourceRecord.Proxied,
		Type:    cfc.ResourceRecord.Type,
		Content: cfc.ResourceRecord.Content,
	}
}

// NewCloudFlareProvider initializes a new CloudFlare DNS based Provider.
func NewCloudFlareProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, proxiedByDefault bool, dryRun bool, dnsRecordsPerPage int, regionKey string, customHostnamesConfig CustomHostnamesConfig) (*CloudFlareProvider, error) {
	// initialize via chosen auth method and returns new API object
	var (
		config *cloudflare.API
		err    error
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
		config, err = cloudflare.NewWithAPIToken(token)
	} else {
		config, err = cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloudflare provider: %w", err)
	}
	provider := &CloudFlareProvider{
		// Client: config,
		Client:                zoneService{config},
		domainFilter:          domainFilter,
		zoneIDFilter:          zoneIDFilter,
		proxiedByDefault:      proxiedByDefault,
		CustomHostnamesConfig: customHostnamesConfig,
		DryRun:                dryRun,
		DNSRecordsPerPage:     dnsRecordsPerPage,
		RegionKey:             regionKey,
	}
	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *CloudFlareProvider) Zones(ctx context.Context) ([]cloudflare.Zone, error) {
	result := []cloudflare.Zone{}

	// if there is a zoneIDfilter configured
	// && if the filter isn't just a blank string (used in tests)
	if len(p.zoneIDFilter.ZoneIDs) > 0 && p.zoneIDFilter.ZoneIDs[0] != "" {
		log.Debugln("zoneIDFilter configured. only looking up zone IDs defined")
		for _, zoneID := range p.zoneIDFilter.ZoneIDs {
			log.Debugf("looking up zone %q", zoneID)
			detailResponse, err := p.Client.ZoneDetails(ctx, zoneID)
			if err != nil {
				log.Errorf("zone %q lookup failed, %v", zoneID, err)
				return result, err
			}
			log.WithFields(log.Fields{
				"zoneName": detailResponse.Name,
				"zoneID":   detailResponse.ID,
			}).Debugln("adding zone for consideration")
			result = append(result, detailResponse)
		}
		return result, nil
	}

	log.Debugln("no zoneIDFilter configured, looking at all zones")

	zonesResponse, err := p.Client.ListZonesContext(ctx)
	if err != nil {
		var apiErr *cloudflare.Error
		if errors.As(err, &apiErr) {
			if apiErr.ClientRateLimited() || apiErr.StatusCode >= http.StatusInternalServerError {
				// Handle rate limit error as a soft error
				return nil, provider.NewSoftError(err)
			}
		}
		return nil, err
	}

	for _, zone := range zonesResponse.Result {
		if !p.domainFilter.Match(zone.Name) {
			log.Debugf("zone %q not in domain filter", zone.Name)
			continue
		}
		result = append(result, zone)
	}

	return result, nil
}

// Records returns the list of records.
func (p *CloudFlareProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}
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
		endpoints = append(endpoints, groupByNameAndTypeWithCustomHostnames(records, chs)...)
	}

	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *CloudFlareProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	cloudflareChanges := []*cloudFlareChange{}

	// if custom hostnames are enabled, deleting first allows to avoid conflicts with the new ones
	if p.CustomHostnamesConfig.Enabled {
		for _, e := range changes.Delete {
			for _, target := range e.Targets {
				cloudflareChanges = append(cloudflareChanges, p.newCloudFlareChange(cloudFlareDelete, e, target, nil))
			}
		}
	}

	for _, e := range changes.Create {
		for _, target := range e.Targets {
			cloudflareChanges = append(cloudflareChanges, p.newCloudFlareChange(cloudFlareCreate, e, target, nil))
		}
	}

	for i, desired := range changes.UpdateNew {
		current := changes.UpdateOld[i]

		add, remove, leave := provider.Difference(current.Targets, desired.Targets)

		for _, a := range remove {
			cloudflareChanges = append(cloudflareChanges, p.newCloudFlareChange(cloudFlareDelete, current, a, current))
		}

		for _, a := range add {
			cloudflareChanges = append(cloudflareChanges, p.newCloudFlareChange(cloudFlareCreate, desired, a, current))
		}

		for _, a := range leave {
			cloudflareChanges = append(cloudflareChanges, p.newCloudFlareChange(cloudFlareUpdate, desired, a, current))
		}
	}

	// TODO: consider deleting before creating even if custom hostnames are not in use
	if !p.CustomHostnamesConfig.Enabled {
		for _, e := range changes.Delete {
			for _, target := range e.Targets {
				cloudflareChanges = append(cloudflareChanges, p.newCloudFlareChange(cloudFlareDelete, e, target, nil))
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
		return !failedChange
	}

	switch change.Action {
	case cloudFlareUpdate:
		if recordTypeCustomHostnameSupported[change.ResourceRecord.Type] {
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
			if recordTypeCustomHostnameSupported[change.ResourceRecord.Type] && changeCH.Hostname != "" {
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
			if recordTypeCustomHostnameSupported[change.ResourceRecord.Type] && changeCH.Hostname != "" {
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

// submitDataLocalizationRegionalHostnameChanges applies a set of data localization regional hostname changes, returns false if it fails
func (p *CloudFlareProvider) submitDataLocalizationRegionalHostnameChanges(ctx context.Context, changes []DataLocalizationRegionalHostnameChange, resourceContainer *cloudflare.ResourceContainer) bool {
	failedChange := false

	for _, change := range changes {
		logFields := log.Fields{
			"hostname":   change.Hostname,
			"region_key": change.RegionKey,
			"action":     change.Action,
			"zone":       resourceContainer.Identifier,
		}
		log.WithFields(logFields).Info("Changing regional hostname")
		switch change.Action {
		case cloudFlareCreate:
			log.WithFields(logFields).Debug("Creating regional hostname")
			if p.DryRun {
				continue
			}
			regionalHostnameParam := createDataLocalizationRegionalHostnameParams(change)
			err := p.Client.CreateDataLocalizationRegionalHostname(ctx, resourceContainer, regionalHostnameParam)
			if err != nil {
				var apiErr *cloudflare.Error
				if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusConflict {
					log.WithFields(logFields).Debug("Regional hostname already exists, updating instead")
					params := updateDataLocalizationRegionalHostnameParams(change)
					err := p.Client.UpdateDataLocalizationRegionalHostname(ctx, resourceContainer, params)
					if err != nil {
						failedChange = true
						log.WithFields(logFields).Errorf("failed to update regional hostname: %v", err)
					}
					continue
				}
				failedChange = true
				log.WithFields(logFields).Errorf("failed to create regional hostname: %v", err)
			}
		case cloudFlareUpdate:
			log.WithFields(logFields).Debug("Updating regional hostname")
			if p.DryRun {
				continue
			}
			regionalHostnameParam := updateDataLocalizationRegionalHostnameParams(change)
			err := p.Client.UpdateDataLocalizationRegionalHostname(ctx, resourceContainer, regionalHostnameParam)
			if err != nil {
				var apiErr *cloudflare.Error
				if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
					log.WithFields(logFields).Debug("Regional hostname not does not exists, creating instead")
					params := createDataLocalizationRegionalHostnameParams(change)
					err := p.Client.CreateDataLocalizationRegionalHostname(ctx, resourceContainer, params)
					if err != nil {
						failedChange = true
						log.WithFields(logFields).Errorf("failed to create regional hostname: %v", err)
					}
					continue
				}
				failedChange = true
				log.WithFields(logFields).Errorf("failed to update regional hostname: %v", err)
			}
		case cloudFlareDelete:
			log.WithFields(logFields).Debug("Deleting regional hostname")
			if p.DryRun {
				continue
			}
			err := p.Client.DeleteDataLocalizationRegionalHostname(ctx, resourceContainer, change.Hostname)
			if err != nil {
				var apiErr *cloudflare.Error
				if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
					log.WithFields(logFields).Debug("Regional hostname does not exists, nothing to do")
					continue
				}
				failedChange = true
				log.WithFields(logFields).Errorf("failed to delete regional hostname: %v", err)
			}
		}
	}

	return !failedChange
}

// dataLocalizationRegionalHostnamesChanges processes a slice of cloudFlare changes and consolidates them
// into a list of data localization regional hostname changes.
// returns nil if no changes are needed
func dataLocalizationRegionalHostnamesChanges(changes []*cloudFlareChange) ([]DataLocalizationRegionalHostnameChange, error) {
	regionalHostnameChanges := make(map[string]DataLocalizationRegionalHostnameChange)
	for _, change := range changes {
		if change.RegionalHostname.Hostname == "" {
			continue
		}
		if change.RegionalHostname.RegionKey == "" {
			return nil, fmt.Errorf("region key is empty for regional hostname %q", change.RegionalHostname.Hostname)
		}
		regionalHostname, ok := regionalHostnameChanges[change.RegionalHostname.Hostname]
		switch change.Action {
		case cloudFlareCreate, cloudFlareUpdate:
			if !ok {
				regionalHostnameChanges[change.RegionalHostname.Hostname] = DataLocalizationRegionalHostnameChange{
					Action:           change.Action,
					RegionalHostname: change.RegionalHostname,
				}
				continue
			}
			if regionalHostname.RegionKey != change.RegionalHostname.RegionKey {
				return nil, fmt.Errorf("conflicting region keys for regional hostname %q: %q and %q", change.RegionalHostname.Hostname, regionalHostname.RegionKey, change.RegionalHostname.RegionKey)
			}
			if (change.Action == cloudFlareUpdate && regionalHostname.Action != cloudFlareUpdate) ||
				regionalHostname.Action == cloudFlareDelete {
				regionalHostnameChanges[change.RegionalHostname.Hostname] = DataLocalizationRegionalHostnameChange{
					Action:           cloudFlareUpdate,
					RegionalHostname: change.RegionalHostname,
				}
			}
		case cloudFlareDelete:
			if !ok {
				regionalHostnameChanges[change.RegionalHostname.Hostname] = DataLocalizationRegionalHostnameChange{
					Action:           cloudFlareDelete,
					RegionalHostname: change.RegionalHostname,
				}
				continue
			}
		}
	}
	return slices.Collect(maps.Values(regionalHostnameChanges)), nil
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
		resourceContainer := cloudflare.ZoneIdentifier(zoneID)

		for _, change := range zoneChanges {
			logFields := log.Fields{
				"record": change.ResourceRecord.Name,
				"type":   change.ResourceRecord.Type,
				"ttl":    change.ResourceRecord.TTL,
				"action": change.Action,
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
				return fmt.Errorf("could not fetch custom hostnames from zone, %v", chErr)
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
				err := p.Client.DeleteDNSRecord(ctx, resourceContainer, recordID)
				if err != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to delete record: %v", err)
				}
				if !p.submitCustomHostnameChanges(ctx, zoneID, change, chs, logFields) {
					failedChange = true
				}
			} else if change.Action == cloudFlareCreate {
				recordParam := getCreateDNSRecordParam(*change)
				_, err := p.Client.CreateDNSRecord(ctx, resourceContainer, recordParam)
				if err != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to create record: %v", err)
				}
				if !p.submitCustomHostnameChanges(ctx, zoneID, change, chs, logFields) {
					failedChange = true
				}
			}
		}

		if regionalHostnamesChanges, err := dataLocalizationRegionalHostnamesChanges(zoneChanges); err == nil {
			if !p.submitDataLocalizationRegionalHostnameChanges(ctx, regionalHostnamesChanges, resourceContainer) {
				failedChange = true
			}
		} else {
			logFields := log.Fields{
				"zone": zoneID,
			}
			log.WithFields(logFields).Errorf("failed to build data localization regional hostname changes: %v", err)
			failedChange = true
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
	adjustedEndpoints := []*endpoint.Endpoint{}
	for _, e := range endpoints {
		proxied := shouldBeProxied(e, p.proxiedByDefault)
		if proxied {
			e.RecordTTL = 0
		}
		e.SetProviderSpecificProperty(source.CloudflareProxiedKey, strconv.FormatBool(proxied))

		if p.CustomHostnamesConfig.Enabled {
			// sort custom hostnames in annotation to properly detect changes
			if customHostnames := getEndpointCustomHostnames(e); len(customHostnames) > 1 {
				sort.Strings(customHostnames)
				e.SetProviderSpecificProperty(source.CloudflareCustomHostnameKey, strings.Join(customHostnames, ","))
			}
		} else {
			// ignore custom hostnames annotations if not enabled
			e.DeleteProviderSpecificProperty(source.CloudflareCustomHostnameKey)
		}

		adjustedEndpoints = append(adjustedEndpoints, e)
	}
	return adjustedEndpoints, nil
}

// changesByZone separates a multi-zone change into a single change per zone.
func (p *CloudFlareProvider) changesByZone(zones []cloudflare.Zone, changeSet []*cloudFlareChange) map[string][]*cloudFlareChange {
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

func (p *CloudFlareProvider) getRecordID(records DNSRecordsMap, record cloudflare.DNSRecord) string {
	if zoneRecord, ok := records[DNSRecordIndex{Name: record.Name, Type: record.Type, Content: record.Content}]; ok {
		return zoneRecord.ID
	}
	return ""
}

func getCustomHostname(chs CustomHostnamesMap, chName string) (cloudflare.CustomHostname, error) {
	if chName == "" {
		return cloudflare.CustomHostname{}, fmt.Errorf("failed to get custom hostname: %q is empty", chName)
	}
	if ch, ok := chs[CustomHostnameIndex{Hostname: chName}]; ok {
		return ch, nil
	}
	return cloudflare.CustomHostname{}, fmt.Errorf("failed to get custom hostname: %q not found", chName)
}

func (p *CloudFlareProvider) newCustomHostname(customHostname string, origin string) cloudflare.CustomHostname {
	return cloudflare.CustomHostname{
		Hostname:           customHostname,
		CustomOriginServer: origin,
		SSL:                getCustomHostnamesSSLOptions(p.CustomHostnamesConfig),
	}
}

func (p *CloudFlareProvider) newCloudFlareChange(action string, ep *endpoint.Endpoint, target string, current *endpoint.Endpoint) *cloudFlareChange {
	ttl := defaultCloudFlareRecordTTL
	proxied := shouldBeProxied(ep, p.proxiedByDefault)

	if ep.RecordTTL.IsConfigured() {
		ttl = int(ep.RecordTTL)
	}

	prevCustomHostnames := []string{}
	newCustomHostnames := map[string]cloudflare.CustomHostname{}
	if p.CustomHostnamesConfig.Enabled {
		if current != nil {
			prevCustomHostnames = getEndpointCustomHostnames(current)
		}
		for _, v := range getEndpointCustomHostnames(ep) {
			newCustomHostnames[v] = p.newCustomHostname(v, ep.DNSName)
		}
	}
	regionalHostname := cloudflare.RegionalHostname{}
	if regionKey := getRegionKey(ep, p.RegionKey); regionKey != "" {
		regionalHostname = cloudflare.RegionalHostname{
			Hostname:  ep.DNSName,
			RegionKey: regionKey,
		}
	}
	return &cloudFlareChange{
		Action: action,
		ResourceRecord: cloudflare.DNSRecord{
			Name: ep.DNSName,
			TTL:  ttl,
			// We have to use pointers to bools now, as the upstream cloudflare-go library requires them
			// see: https://github.com/cloudflare/cloudflare-go/pull/595
			Proxied: &proxied,
			Type:    ep.RecordType,
			Content: target,
		},
		RegionalHostname:    regionalHostname,
		CustomHostnamesPrev: prevCustomHostnames,
		CustomHostnames:     newCustomHostnames,
	}
}

func newDNSRecordIndex(r cloudflare.DNSRecord) DNSRecordIndex {
	return DNSRecordIndex{Name: r.Name, Type: r.Type, Content: r.Content}
}

// listDNSRecordsWithAutoPagination performs automatic pagination of results on requests to cloudflare.ListDNSRecords with custom per_page values
func (p *CloudFlareProvider) listDNSRecordsWithAutoPagination(ctx context.Context, zoneID string) (DNSRecordsMap, error) {
	// for faster getRecordID lookup
	records := make(DNSRecordsMap)
	resultInfo := cloudflare.ResultInfo{PerPage: p.DNSRecordsPerPage, Page: 1}
	params := cloudflare.ListDNSRecordsParams{ResultInfo: resultInfo}
	for {
		pageRecords, resultInfo, err := p.Client.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			var apiErr *cloudflare.Error
			if errors.As(err, &apiErr) {
				if apiErr.ClientRateLimited() || apiErr.StatusCode >= http.StatusInternalServerError {
					// Handle rate limit error as a soft error
					return nil, provider.NewSoftError(err)
				}
			}
			return nil, err
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

func newCustomHostnameIndex(ch cloudflare.CustomHostname) CustomHostnameIndex {
	return CustomHostnameIndex{Hostname: ch.Hostname}
}

// listCustomHostnamesWithPagination performs automatic pagination of results on requests to cloudflare.CustomHostnames
func (p *CloudFlareProvider) listCustomHostnamesWithPagination(ctx context.Context, zoneID string) (CustomHostnamesMap, error) {
	if !p.CustomHostnamesConfig.Enabled {
		return nil, nil
	}
	chs := make(CustomHostnamesMap)
	resultInfo := cloudflare.ResultInfo{Page: 1}
	for {
		pageCustomHostnameListResponse, result, err := p.Client.CustomHostnames(ctx, zoneID, resultInfo.Page, cloudflare.CustomHostname{})
		if err != nil {
			var apiErr *cloudflare.Error
			if errors.As(err, &apiErr) {
				if apiErr.ClientRateLimited() || apiErr.StatusCode >= http.StatusInternalServerError {
					// Handle rate limit error as a soft error
					return nil, provider.NewSoftError(err)
				}
			}
			log.Errorf("zone %q failed to fetch custom hostnames. Please check if \"Cloudflare for SaaS\" is enabled and API key permissions, %v", zoneID, err)
			return nil, err
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

func getCustomHostnamesSSLOptions(customHostnamesConfig CustomHostnamesConfig) *cloudflare.CustomHostnameSSL {
	return &cloudflare.CustomHostnameSSL{
		Type:                 "dv",
		Method:               "http",
		CertificateAuthority: customHostnamesConfig.CertificateAuthority,
		BundleMethod:         "ubiquitous",
		Settings: cloudflare.CustomHostnameSSLSettings{
			MinTLSVersion: customHostnamesConfig.MinTLSVersion,
		},
	}
}

func shouldBeProxied(ep *endpoint.Endpoint, proxiedByDefault bool) bool {
	proxied := proxiedByDefault

	for _, v := range ep.ProviderSpecific {
		if v.Name == source.CloudflareProxiedKey {
			b, err := strconv.ParseBool(v.Value)
			if err != nil {
				log.Errorf("Failed to parse annotation [%q]: %v", source.CloudflareProxiedKey, err)
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

func getRegionKey(endpoint *endpoint.Endpoint, defaultRegionKey string) string {
	if !recordTypeRegionalHostnameSupported[endpoint.RecordType] {
		return ""
	}

	for _, v := range endpoint.ProviderSpecific {
		if v.Name == annotations.CloudflareRegionKey {
			return v.Value
		}
	}
	return defaultRegionKey
}

func getEndpointCustomHostnames(ep *endpoint.Endpoint) []string {
	for _, v := range ep.ProviderSpecific {
		if v.Name == source.CloudflareCustomHostnameKey {
			customHostnames := strings.Split(v.Value, ",")
			return customHostnames
		}
	}
	return []string{}
}

func groupByNameAndTypeWithCustomHostnames(records DNSRecordsMap, chs CustomHostnamesMap) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}

	// group supported records by name and type
	groups := map[string][]cloudflare.DNSRecord{}

	for _, r := range records {
		if !provider.SupportedRecordType(r.Type) {
			continue
		}

		groupBy := r.Name + r.Type
		if _, ok := groups[groupBy]; !ok {
			groups[groupBy] = []cloudflare.DNSRecord{}
		}

		groups[groupBy] = append(groups[groupBy], r)
	}

	// map custom origin to custom hostname, custom origin should match to a dns record
	customHostnames := map[string][]string{}

	for _, c := range chs {
		customHostnames[c.CustomOriginServer] = append(customHostnames[c.CustomOriginServer], c.Hostname)
	}

	// create single endpoint with all the targets for each name/type
	for _, records := range groups {
		if len(records) == 0 {
			return endpoints
		}
		targets := make([]string, len(records))
		for i, record := range records {
			targets[i] = record.Content
		}
		e := endpoint.NewEndpointWithTTL(
			records[0].Name,
			records[0].Type,
			endpoint.TTL(records[0].TTL),
			targets...)
		proxied := false
		if records[0].Proxied != nil {
			proxied = *records[0].Proxied
		}
		if e == nil {
			continue
		}
		e = e.WithProviderSpecific(source.CloudflareProxiedKey, strconv.FormatBool(proxied))
		// noop (customHostnames is empty) if custom hostnames feature is not in use
		if customHostnames, ok := customHostnames[records[0].Name]; ok {
			sort.Strings(customHostnames)
			e = e.WithProviderSpecific(source.CloudflareCustomHostnameKey, strings.Join(customHostnames, ","))
		}

		endpoints = append(endpoints, e)
	}

	return endpoints
}

// boolPtr is used as a helper function to return a pointer to a boolean
// Needed because some parameters require a pointer.
func boolPtr(b bool) *bool {
	return &b
}
