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

package designate

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
	"github.com/gophercloud/gophercloud/pagination"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/tlsutils"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	// ID of the RecordSet from which endpoint was created
	designateRecordSetID = "designate-recordset-id"
	// Zone ID of the RecordSet
	designateZoneID = "designate-record-id"

	// Initial records values of the RecordSet. This label is required in order not to loose records that haven't
	// changed where there are several targets per domain and only some of them changed.
	// Values are joined by zero-byte to in order to get a single string
	designateOriginalRecords = "designate-original-records"
)

// interface between provider and OpenStack DNS API
type designateClientInterface interface {
	// ForEachZone calls handler for each zone managed by the Designate
	ForEachZone(handler func(zone *zones.Zone) error) error

	// ForEachRecordSet calls handler for each recordset in the given DNS zone
	ForEachRecordSet(zoneID string, handler func(recordSet *recordsets.RecordSet) error) error

	// CreateRecordSet creates recordset in the given DNS zone
	CreateRecordSet(zoneID string, opts recordsets.CreateOpts) (string, error)

	// UpdateRecordSet updates recordset in the given DNS zone
	UpdateRecordSet(zoneID, recordSetID string, opts recordsets.UpdateOpts) error

	// DeleteRecordSet deletes recordset in the given DNS zone
	DeleteRecordSet(zoneID, recordSetID string) error
}

// implementation of the designateClientInterface
type designateClient struct {
	serviceClient *gophercloud.ServiceClient
}

// factory function for the designateClientInterface
func newDesignateClient() (designateClientInterface, error) {
	serviceClient, err := createDesignateServiceClient()
	if err != nil {
		return nil, err
	}
	return &designateClient{serviceClient}, nil
}

// copies environment variables to new names without overwriting existing values
func remapEnv(mapping map[string]string) {
	for k, v := range mapping {
		currentVal := os.Getenv(k)
		newVal := os.Getenv(v)
		if currentVal == "" && newVal != "" {
			os.Setenv(k, newVal)
		}
	}
}

// returns OpenStack Keystone authentication settings by obtaining values from standard environment variables.
// also fixes incompatibilities between gophercloud implementation and *-stackrc files that can be downloaded
// from OpenStack dashboard in latest versions
func getAuthSettings() (gophercloud.AuthOptions, error) {
	remapEnv(map[string]string{
		"OS_TENANT_NAME": "OS_PROJECT_NAME",
		"OS_TENANT_ID":   "OS_PROJECT_ID",
		"OS_DOMAIN_NAME": "OS_USER_DOMAIN_NAME",
		"OS_DOMAIN_ID":   "OS_USER_DOMAIN_ID",
	})

	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return gophercloud.AuthOptions{}, err
	}
	opts.AllowReauth = true
	if !strings.HasSuffix(opts.IdentityEndpoint, "/") {
		opts.IdentityEndpoint += "/"
	}
	if !strings.HasSuffix(opts.IdentityEndpoint, "/v2.0/") && !strings.HasSuffix(opts.IdentityEndpoint, "/v3/") {
		opts.IdentityEndpoint += "v2.0/"
	}
	return opts, nil
}

// authenticate in OpenStack and obtain Designate service endpoint
func createDesignateServiceClient() (*gophercloud.ServiceClient, error) {
	opts, err := getAuthSettings()
	if err != nil {
		return nil, err
	}
	log.Infof("Using OpenStack Keystone at %s", opts.IdentityEndpoint)
	authProvider, err := openstack.NewClient(opts.IdentityEndpoint)
	if err != nil {
		return nil, err
	}

	tlsConfig, err := tlsutils.CreateTLSConfig("OPENSTACK")
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}
	authProvider.HTTPClient.Transport = transport

	if err = openstack.Authenticate(authProvider, opts); err != nil {
		return nil, err
	}

	eo := gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	}

	client, err := openstack.NewDNSV2(authProvider, eo)
	if err != nil {
		return nil, err
	}
	log.Infof("Found OpenStack Designate service at %s", client.Endpoint)
	return client, nil
}

// ForEachZone calls handler for each zone managed by the Designate
func (c designateClient) ForEachZone(handler func(zone *zones.Zone) error) error {
	pager := zones.List(c.serviceClient, zones.ListOpts{})
	return pager.EachPage(
		func(page pagination.Page) (bool, error) {
			list, err := zones.ExtractZones(page)
			if err != nil {
				return false, err
			}
			for _, zone := range list {
				err := handler(&zone)
				if err != nil {
					return false, err
				}
			}
			return true, nil
		},
	)
}

// ForEachRecordSet calls handler for each recordset in the given DNS zone
func (c designateClient) ForEachRecordSet(zoneID string, handler func(recordSet *recordsets.RecordSet) error) error {
	pager := recordsets.ListByZone(c.serviceClient, zoneID, recordsets.ListOpts{})
	return pager.EachPage(
		func(page pagination.Page) (bool, error) {
			list, err := recordsets.ExtractRecordSets(page)
			if err != nil {
				return false, err
			}
			for _, recordSet := range list {
				err := handler(&recordSet)
				if err != nil {
					return false, err
				}
			}
			return true, nil
		},
	)
}

// CreateRecordSet creates recordset in the given DNS zone
func (c designateClient) CreateRecordSet(zoneID string, opts recordsets.CreateOpts) (string, error) {
	r, err := recordsets.Create(c.serviceClient, zoneID, opts).Extract()
	if err != nil {
		return "", err
	}
	return r.ID, nil
}

// UpdateRecordSet updates recordset in the given DNS zone
func (c designateClient) UpdateRecordSet(zoneID, recordSetID string, opts recordsets.UpdateOpts) error {
	_, err := recordsets.Update(c.serviceClient, zoneID, recordSetID, opts).Extract()
	return err
}

// DeleteRecordSet deletes recordset in the given DNS zone
func (c designateClient) DeleteRecordSet(zoneID, recordSetID string) error {
	return recordsets.Delete(c.serviceClient, zoneID, recordSetID).ExtractErr()
}

// designate provider type
type designateProvider struct {
	provider.BaseProvider
	client designateClientInterface

	// only consider hosted zones managing domains ending in this suffix
	domainFilter endpoint.DomainFilter
	dryRun       bool
}

// NewDesignateProvider is a factory function for OpenStack designate providers
func NewDesignateProvider(domainFilter endpoint.DomainFilter, dryRun bool) (provider.Provider, error) {
	client, err := newDesignateClient()
	if err != nil {
		return nil, err
	}
	return &designateProvider{
		client:       client,
		domainFilter: domainFilter,
		dryRun:       dryRun,
	}, nil
}

// converts domain names to FQDN
func canonicalizeDomainNames(domains []string) []string {
	var cDomains []string
	for _, d := range domains {
		if !strings.HasSuffix(d, ".") {
			d += "."
			cDomains = append(cDomains, strings.ToLower(d))
		}
	}
	return cDomains
}

// converts domain name to FQDN
func canonicalizeDomainName(d string) string {
	if !strings.HasSuffix(d, ".") {
		d += "."
	}
	return strings.ToLower(d)
}

// returns ZoneID -> ZoneName mapping for zones that are managed by the Designate and match domain filter
func (p designateProvider) getZones() (map[string]string, error) {
	result := map[string]string{}

	err := p.client.ForEachZone(
		func(zone *zones.Zone) error {
			if zone.Type != "" && strings.ToUpper(zone.Type) != "PRIMARY" || zone.Status != "ACTIVE" {
				return nil
			}

			zoneName := canonicalizeDomainName(zone.Name)
			if !p.domainFilter.Match(zoneName) {
				return nil
			}
			result[zone.ID] = zoneName
			return nil
		},
	)

	return result, err
}

// finds best suitable DNS zone for the hostname
func (p designateProvider) getHostZoneID(hostname string, managedZones map[string]string) (string, error) {
	longestZoneLength := 0
	resultID := ""

	for zoneID, zoneName := range managedZones {
		if !strings.HasSuffix(hostname, zoneName) {
			continue
		}
		ln := len(zoneName)
		if ln > longestZoneLength {
			resultID = zoneID
			longestZoneLength = ln
		}
	}

	return resultID, nil
}

// Records returns the list of records.
func (p designateProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var result []*endpoint.Endpoint
	managedZones, err := p.getZones()
	if err != nil {
		return nil, err
	}
	recordSets, err := p.records(managedZones)
	if err != nil {
		return nil, err
	}
	for _, recordSet := range recordSets {
		for _, record := range recordSet.Records {
			ep := endpoint.NewEndpoint(recordSet.Name, recordSet.Type, record)
			ep.Labels[designateRecordSetID] = recordSet.ID
			ep.Labels[designateZoneID] = recordSet.ZoneID
			ep.Labels[designateOriginalRecords] = strings.Join(recordSet.Records, "\000")
			result = append(result, ep)
		}
	}

	return result, nil
}

func (p designateProvider) records(zones map[string]string) ([]*recordsets.RecordSet, error) {
	var records []*recordsets.RecordSet
	for zone := range zones {
		err := p.client.ForEachRecordSet(zone, func(recordSet *recordsets.RecordSet) error {
			if recordSet.Type != endpoint.RecordTypeA && recordSet.Type != endpoint.RecordTypeTXT && recordSet.Type != endpoint.RecordTypeCNAME {
				return nil
			}
			rs := *recordSet
			records = append(records, &rs)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return records, nil
}

// temporary structure to hold recordset parameters so that we could aggregate endpoints into recordsets
type recordSet struct {
	dnsName     string
	recordType  string
	zoneID      string
	recordSetID string
	names       map[string]bool
}

// adds endpoint into recordset aggregation, loading original values from endpoint labels first
func addEndpoint(ep *endpoint.Endpoint, recordSets map[string]*recordSet, delete bool) {
	key := fmt.Sprintf("%s/%s", ep.DNSName, ep.RecordType)
	rs := recordSets[key]
	if rs == nil {
		rs = &recordSet{
			dnsName:    canonicalizeDomainName(ep.DNSName),
			recordType: ep.RecordType,
			names:      make(map[string]bool),
		}
	}
	if rs.zoneID == "" {
		rs.zoneID = ep.Labels[designateZoneID]
	}
	if rs.recordSetID == "" {
		rs.recordSetID = ep.Labels[designateRecordSetID]
	}
	for _, rec := range strings.Split(ep.Labels[designateOriginalRecords], "\000") {
		if _, ok := rs.names[rec]; !ok && rec != "" {
			rs.names[rec] = true
		}
	}
	targets := ep.Targets
	if ep.RecordType == endpoint.RecordTypeCNAME {
		targets = canonicalizeDomainNames(targets)
	}
	for _, t := range targets {
		rs.names[t] = !delete
	}
	recordSets[key] = rs
}

// ApplyChanges applies a given set of changes in a given zone.
func (p designateProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	managedZones, err := p.getZones()
	if err != nil {
		return err
	}
	recordSets := map[string]*recordSet{}
	for _, ep := range changes.Create {
		addEndpoint(ep, recordSets, false)
	}
	var knownRecordSets []*recordsets.RecordSet
	if len(changes.Delete) > 0 || len(changes.UpdateNew) > 0 {
		// The why... TXTRegistry.Records on purpose filters TXT records.
		// The planner on the other hand, "creates" a TXT Endpoint if needed, for instance in the
		// Delete case in changes.Delete.
		// But because generated and added to the list of endpoints
		// (and not coming from p.Records), there is no ID associated with it.
		// upsertRecordSet depends on IDs for the recordSet and the zone, primarily because this
		// is how designate works. So we fill missing IDs in here.
		var zoneErr error
		knownRecordSets, zoneErr = p.records(managedZones)
		if zoneErr != nil {
			return fmt.Errorf("failed to fetch knownRecordSets: %w", zoneErr)
		}
	}
	for _, ep := range changes.UpdateNew {
		err := enrichEndpoint(ep, knownRecordSets)
		if err != nil {
			// This error likely only happens if the corresponding entry is already gone.
			log.Info(err)
			continue
		}
		addEndpoint(ep, recordSets, false)
	}
	for _, ep := range changes.UpdateOld {
		addEndpoint(ep, recordSets, true)
	}
	for _, ep := range changes.Delete {
		err := enrichEndpoint(ep, knownRecordSets)
		if err != nil {
			// This error likely only happens if the corresponding entry is already gone.
			log.Info(err)
			continue
		}
		addEndpoint(ep, recordSets, true)
	}
	for _, rs := range recordSets {
		if err2 := p.upsertRecordSet(rs, managedZones); err == nil {
			err = err2
		}
	}
	return err
}

// enrichEndpoint adds designateZoneID and designateRecordSetID to an endpoint based on records
// in designate.
func enrichEndpoint(ep *endpoint.Endpoint, recordSets []*recordsets.RecordSet) error {
	if _, ok := ep.Labels[designateRecordSetID]; !ok {
		for _, r := range recordSets {
			n := canonicalizeDomainName(ep.DNSName)
			if r.Name == n && r.Type == ep.RecordType {
				ep.Labels[designateZoneID] = r.ZoneID
				ep.Labels[designateRecordSetID] = r.ID
				return nil
			}
		}
	}
	if _, ok := ep.Labels[designateZoneID]; !ok {
		return fmt.Errorf("unable to enrich Endpoint because of missing ZoneID: %v", *ep)
	}
	if _, ok := ep.Labels[designateRecordSetID]; !ok {
		return fmt.Errorf("unable to enrich Endpoint because of missing RecordSetID: %v", *ep)
	}
	return nil
}

// apply recordset changes by inserting/updating/deleting recordsets
func (p designateProvider) upsertRecordSet(rs *recordSet, managedZones map[string]string) error {
	if rs.zoneID == "" {
		var err error
		rs.zoneID, err = p.getHostZoneID(rs.dnsName, managedZones)
		if err != nil {
			return err
		}
		if rs.zoneID == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", rs.dnsName)
			return nil
		}
	}
	var records []string
	for rec, v := range rs.names {
		if v {
			records = append(records, rec)
		}
	}
	if rs.recordSetID == "" && records == nil {
		return nil
	}
	if rs.recordSetID == "" {
		opts := recordsets.CreateOpts{
			Name:    rs.dnsName,
			Type:    rs.recordType,
			Records: records,
		}
		log.Infof("Creating records: %s/%s: %s", rs.dnsName, rs.recordType, strings.Join(records, ","))
		if p.dryRun {
			return nil
		}
		_, err := p.client.CreateRecordSet(rs.zoneID, opts)
		return err
	} else if len(records) == 0 {
		log.Infof("Deleting records for %s/%s", rs.dnsName, rs.recordType)
		if p.dryRun {
			return nil
		}
		return p.client.DeleteRecordSet(rs.zoneID, rs.recordSetID)
	} else {
		ttl := 0
		opts := recordsets.UpdateOpts{
			Records: records,
			TTL:     &ttl,
		}
		log.Infof("Updating records: %s/%s: %s", rs.dnsName, rs.recordType, strings.Join(records, ","))
		if p.dryRun {
			return nil
		}
		return p.client.UpdateRecordSet(rs.zoneID, rs.recordSetID, opts)
	}
}
