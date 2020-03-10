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

package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/nesv/go-dynect/dynect"
	"github.com/sanyu/dynectsoap/dynectsoap"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const (
	// 10 minutes default timeout if not configured using flags
	dynDefaultTTL = 600

	// when rate limit is hit retry up to 5 times after sleep 1m between retries
	dynMaxRetriesOnErrRateLimited = 5

	// two consecutive bad logins happen at least this many seconds apart
	// While it is easy to get the username right, misconfiguring the password
	// can get account blocked. Exit(1) is not a good solution
	// as k8s will restart the pod and another login attempt will be made
	badLoginMinIntervalSeconds = 30 * 60

	// this prefix must be stripped from resource links before feeding them to dynect.Client.Do()
	restAPIPrefix = "/REST/"
)

func unixNow() int64 {
	return time.Now().Unix()
}

// DynConfig hold connection parameters to dyn.com and internal state
type DynConfig struct {
	DomainFilter  endpoint.DomainFilter
	ZoneIDFilter  ZoneIDFilter
	DryRun        bool
	CustomerName  string
	Username      string
	Password      string
	MinTTLSeconds int
	AppVersion    string
	DynVersion    string
}

// ZoneSnapshot stores a single recordset for a zone for a single serial
type ZoneSnapshot struct {
	serials   map[string]int
	endpoints map[string][]*endpoint.Endpoint
}

// GetRecordsForSerial retrieves from memory the last known recordset for the (zone, serial) tuple
func (snap *ZoneSnapshot) GetRecordsForSerial(zone string, serial int) []*endpoint.Endpoint {
	lastSerial, ok := snap.serials[zone]
	if !ok {
		// no mapping
		return nil
	}

	if lastSerial != serial {
		// outdated mapping
		return nil
	}

	endpoints, ok := snap.endpoints[zone]
	if !ok {
		// probably a bug
		return nil
	}

	return endpoints
}

// StoreRecordsForSerial associates a result set with a (zone, serial)
func (snap *ZoneSnapshot) StoreRecordsForSerial(zone string, serial int, records []*endpoint.Endpoint) {
	snap.serials[zone] = serial
	snap.endpoints[zone] = records
}

// DynProvider is the actual interface impl.
type dynProviderState struct {
	DynConfig
	LastLoginErrorTime int64

	ZoneSnapshot *ZoneSnapshot
}

// ZoneChange is missing from dynect: https://help.dyn.com/get-zone-changeset-api/
type ZoneChange struct {
	ID     int              `json:"id"`
	UserID int              `json:"user_id"`
	Zone   string           `json:"zone"`
	FQDN   string           `json:"FQDN"`
	Serial int              `json:"serial"`
	TTL    int              `json:"ttl"`
	Type   string           `json:"rdata_type"`
	RData  dynect.DataBlock `json:"rdata"`
}

// ZoneChangesResponse is missing from dynect: https://help.dyn.com/get-zone-changeset-api/
type ZoneChangesResponse struct {
	dynect.ResponseBlock
	Data []ZoneChange `json:"data"`
}

// ZonePublishRequest is missing from dynect but the notes field is a nice place to let
// external-dns report some internal info during commit
type ZonePublishRequest struct {
	Publish bool   `json:"publish"`
	Notes   string `json:"notes"`
}

// ZonePublishResponse holds the status after publish
type ZonePublishResponse struct {
	dynect.ResponseBlock
	Data map[string]interface{} `json:"data"`
}

// NewDynProvider initializes a new Dyn Provider.
func NewDynProvider(config DynConfig) (Provider, error) {
	return &dynProviderState{
		DynConfig: config,
		ZoneSnapshot: &ZoneSnapshot{
			endpoints: map[string][]*endpoint.Endpoint{},
			serials:   map[string]int{},
		},
	}, nil
}

// filterAndFixLinks removes from `links` all the records we don't care about
// and strops the /REST/ prefix
func filterAndFixLinks(links []string, filter endpoint.DomainFilter) []string {
	var result []string
	for _, link := range links {

		// link looks like /REST/CNAMERecord/acme.com/exchange.acme.com/349386875

		// strip /REST/
		link = strings.TrimPrefix(link, restAPIPrefix)

		// simply ignore all record types we don't care about
		if !strings.HasPrefix(link, endpoint.RecordTypeA) &&
			!strings.HasPrefix(link, endpoint.RecordTypeCNAME) &&
			!strings.HasPrefix(link, endpoint.RecordTypeTXT) {
			continue
		}

		// strip ID suffix
		domain := link[0:strings.LastIndexByte(link, '/')]
		// strip zone prefix
		domain = domain[strings.LastIndexByte(domain, '/')+1:]
		if filter.Match(domain) {
			result = append(result, link)
		}
	}

	return result
}

func fixMissingTTL(ttl endpoint.TTL, minTTLSeconds int) string {
	i := dynDefaultTTL
	if ttl.IsConfigured() {
		if int(ttl) < minTTLSeconds {
			i = minTTLSeconds
		} else {
			i = int(ttl)
		}
	}

	return strconv.Itoa(i)
}

// merge produces a singe list of records that can be used as a replacement.
// Dyn allows to replace all records with a single call
// Invariant: the result contains only elements from the updateNew parameter
func merge(updateOld, updateNew []*endpoint.Endpoint) []*endpoint.Endpoint {
	findMatch := func(template *endpoint.Endpoint) *endpoint.Endpoint {
		for _, new := range updateNew {
			if template.DNSName == new.DNSName &&
				template.RecordType == new.RecordType {
				return new
			}
		}
		return nil
	}

	var result []*endpoint.Endpoint
	for _, old := range updateOld {
		matchingNew := findMatch(old)
		if matchingNew == nil {
			// no match, shouldn't happen
			continue
		}

		if !matchingNew.Targets.Same(old.Targets) {
			// new target: always update, TTL will be overwritten too if necessary
			result = append(result, matchingNew)
			continue
		}

		if matchingNew.RecordTTL != 0 && matchingNew.RecordTTL != old.RecordTTL {
			// same target, but new non-zero TTL set in k8s, must update
			// probably would happen only if there is a bug in the code calling the provider
			result = append(result, matchingNew)
		}
	}

	return result
}

func apiRetryLoop(f func() error) error {
	var err error
	for i := 0; i < dynMaxRetriesOnErrRateLimited; i++ {
		err = f()
		if err == nil || err != dynect.ErrRateLimited {
			// success or not retryable error
			return err
		}

		// https://help.dyn.com/managed-dns-api-rate-limit/
		log.Debugf("Rate limit has been hit, sleeping for 1m (%d/%d)", i, dynMaxRetriesOnErrRateLimited)
		time.Sleep(1 * time.Minute)
	}

	return err
}

func (d *dynProviderState) allRecordsToEndpoints(records *dynectsoap.GetAllRecordsResponseType) []*endpoint.Endpoint {
	result := []*endpoint.Endpoint{}
	//Convert each record to an endpoint

	//Process A Records
	for _, rec := range records.Data.A_records {
		ep := &endpoint.Endpoint{
			DNSName:    rec.Fqdn,
			RecordTTL:  endpoint.TTL(rec.Ttl),
			RecordType: rec.Record_type,
			Targets:    endpoint.Targets{rec.Rdata.Address},
		}
		log.Debugf("A record: %v", *ep)
		result = append(result, ep)
	}

	//Process CNAME Records
	for _, rec := range records.Data.Cname_records {
		ep := &endpoint.Endpoint{
			DNSName:    rec.Fqdn,
			RecordTTL:  endpoint.TTL(rec.Ttl),
			RecordType: rec.Record_type,
			Targets:    endpoint.Targets{strings.TrimSuffix(rec.Rdata.Cname, ".")},
		}
		log.Debugf("CNAME record: %v", *ep)
		result = append(result, ep)
	}

	//Process TXT Records
	for _, rec := range records.Data.Txt_records {
		ep := &endpoint.Endpoint{
			DNSName:    rec.Fqdn,
			RecordTTL:  endpoint.TTL(rec.Ttl),
			RecordType: rec.Record_type,
			Targets:    endpoint.Targets{rec.Rdata.Txtdata},
		}
		log.Debugf("TXT record: %v", *ep)
		result = append(result, ep)
	}

	return result

}

func errorOrValue(err error, value interface{}) interface{} {
	if err == nil {
		return value
	}

	return err
}

// endpointToRecord puts the Target of an Endpoint in the correct field of DataBlock.
// See DataBlock comments for more info
func endpointToRecord(ep *endpoint.Endpoint) *dynect.DataBlock {
	result := dynect.DataBlock{}

	if ep.RecordType == endpoint.RecordTypeA {
		result.Address = ep.Targets[0]
	} else if ep.RecordType == endpoint.RecordTypeCNAME {
		result.CName = ep.Targets[0]
	} else if ep.RecordType == endpoint.RecordTypeTXT {
		result.TxtData = ep.Targets[0]
	}

	return &result
}

func (d *dynProviderState) fetchZoneSerial(client *dynect.Client, zone string) (int, error) {
	var resp dynect.ZoneResponse

	err := client.Do("GET", fmt.Sprintf("Zone/%s", zone), nil, &resp)

	if err != nil {
		return 0, err
	}

	return resp.Data.Serial, nil
}

//Use SOAP to fetch all records with a single call
func (d *dynProviderState) fetchAllRecordsInZone(zone string) (*dynectsoap.GetAllRecordsResponseType, error) {
	var err error
	client := dynectsoap.NewClient("https://api2.dynect.net/SOAP/")
	service := dynectsoap.NewDynect(client)

	sessionRequest := dynectsoap.SessionLoginRequestType{
		Customer_name:  d.CustomerName,
		User_name:      d.Username,
		Password:       d.Password,
		Fault_incompat: 0,
	}
	resp := dynectsoap.SessionLoginResponseType{}
	err = apiRetryLoop(func() error {
		return service.Do(&sessionRequest, &resp)
	})

	if err != nil {
		return nil, err
	}

	token := resp.Data.Token

	logoutRequest := dynectsoap.SessionLogoutRequestType{
		Token:          token,
		Fault_incompat: 0,
	}
	logoutResponse := dynectsoap.SessionLogoutResponseType{}
	defer service.Do(&logoutRequest, &logoutResponse)

	req := dynectsoap.GetAllRecordsRequestType{
		Token:          token,
		Zone:           zone,
		Fault_incompat: 0,
	}
	records := dynectsoap.GetAllRecordsResponseType{}
	err = apiRetryLoop(func() error {
		return service.Do(&req, &records)
	})

	if err != nil {
		return nil, err
	}
	log.Debugf("Got all Records, status is %s", records.Status)

	if strings.ToLower(records.Status) == "incomplete" {
		jobRequest := dynectsoap.GetJobRequestType{
			Token:          token,
			Job_id:         records.Job_id,
			Fault_incompat: 0,
		}

		jobResults := dynectsoap.GetJobResponseType{}
		err = apiRetryLoop(func() error {
			return service.GetJobRetry(&jobRequest, &jobResults)
		})
		if err != nil {
			return nil, err
		}
		return jobResults.Data.(*dynectsoap.GetAllRecordsResponseType), nil
	}

	return &records, nil

}

// buildLinkToRecord build a resource link. The symmetry of the dyn API is used to save
// switch-case boilerplate.
// Empty response means the endpoint is not mappable to a records link: either because the fqdn
// is not matched by the domainFilter or it is in the wrong zone
func (d *dynProviderState) buildLinkToRecord(ep *endpoint.Endpoint) string {
	if ep == nil {
		return ""
	}
	var matchingZone = ""
	for _, zone := range d.ZoneIDFilter.zoneIDs {
		if strings.HasSuffix(ep.DNSName, zone) {
			matchingZone = zone
			break
		}
	}

	if matchingZone == "" {
		// no matching zone, ignore
		return ""
	}

	if !d.DomainFilter.Match(ep.DNSName) {
		// no matching domain, ignore
		return ""
	}

	return fmt.Sprintf("%sRecord/%s/%s/", ep.RecordType, matchingZone, ep.DNSName)
}

// create a dynect client and performs login. You need to clean it up.
// This method also stores the DynAPI version.
// Don't user the dynect.Client.Login()
func (d *dynProviderState) login() (*dynect.Client, error) {
	if d.LastLoginErrorTime != 0 {
		secondsSinceLastError := unixNow() - d.LastLoginErrorTime
		if secondsSinceLastError < badLoginMinIntervalSeconds {
			return nil, fmt.Errorf("will not attempt an API call as the last login failure occurred just %ds ago", secondsSinceLastError)
		}
	}
	client := dynect.NewClient(d.CustomerName)

	var req = dynect.LoginBlock{
		Username:     d.Username,
		Password:     d.Password,
		CustomerName: d.CustomerName}

	var resp dynect.LoginResponse

	err := client.Do("POST", "Session", req, &resp)
	if err != nil {
		d.LastLoginErrorTime = unixNow()
		return nil, err
	}

	d.LastLoginErrorTime = 0
	client.Token = resp.Data.Token

	// this is the only change from the original
	d.DynVersion = resp.Data.Version
	return client, nil
}

// the zones we are allowed to touch. Currently only exact matches are considered, not all
// zones with the given suffix
func (d *dynProviderState) zones(client *dynect.Client) []string {
	return d.ZoneIDFilter.zoneIDs
}

func (d *dynProviderState) buildRecordRequest(ep *endpoint.Endpoint) (string, *dynect.RecordRequest) {
	link := d.buildLinkToRecord(ep)
	if link == "" {
		return "", nil
	}

	record := dynect.RecordRequest{
		TTL:   fixMissingTTL(ep.RecordTTL, d.MinTTLSeconds),
		RData: *endpointToRecord(ep),
	}
	return link, &record
}

// deleteRecord deletes all existing records (CNAME, TXT, A) for the given Endpoint.DNSName with 1 API call
func (d *dynProviderState) deleteRecord(client *dynect.Client, ep *endpoint.Endpoint) error {
	link := d.buildLinkToRecord(ep)
	if link == "" {
		return nil
	}

	response := dynect.RecordResponse{}

	err := apiRetryLoop(func() error {
		return client.Do("DELETE", link, nil, &response)
	})

	log.Debugf("Deleting record %s: %+v,", link, errorOrValue(err, &response))
	return err
}

// replaceRecord replaces all existing records pf the given type for the Endpoint.DNSName with 1 API call
func (d *dynProviderState) replaceRecord(client *dynect.Client, ep *endpoint.Endpoint) error {
	link, record := d.buildRecordRequest(ep)
	if link == "" {
		return nil
	}

	response := dynect.RecordResponse{}
	err := apiRetryLoop(func() error {
		return client.Do("PUT", link, record, &response)
	})

	log.Debugf("Replacing record %s: %+v,", link, errorOrValue(err, &response))
	return err
}

// createRecord creates a single record with 1 API call
func (d *dynProviderState) createRecord(client *dynect.Client, ep *endpoint.Endpoint) error {
	link, record := d.buildRecordRequest(ep)
	if link == "" {
		return nil
	}

	response := dynect.RecordResponse{}
	err := apiRetryLoop(func() error {
		return client.Do("POST", link, record, &response)
	})

	log.Debugf("Creating record %s: %+v,", link, errorOrValue(err, &response))
	return err
}

// commit commits all pending changes. It will always attempt to commit, if there are no
func (d *dynProviderState) commit(client *dynect.Client) error {
	errs := []error{}

	for _, zone := range d.zones(client) {
		// extra call if in debug mode to fetch pending changes
		if log.GetLevel() >= log.DebugLevel {
			response := ZoneChangesResponse{}
			err := client.Do("GET", fmt.Sprintf("ZoneChanges/%s/", zone), nil, &response)
			log.Debugf("Pending changes for zone %s: %+v", zone, errorOrValue(err, &response))
		}

		h, err := os.Hostname()
		if err != nil {
			h = "unknown-host"
		}
		notes := fmt.Sprintf("Change by external-dns@%s, DynAPI@%s, %s on %s",
			d.AppVersion,
			d.DynVersion,
			time.Now().Format(time.RFC3339),
			h,
		)

		zonePublish := ZonePublishRequest{
			Publish: true,
			Notes:   notes,
		}

		response := ZonePublishResponse{}

		// always retry the commit: don't waste the good work so far
		err = apiRetryLoop(func() error {
			return client.Do("PUT", fmt.Sprintf("Zone/%s/", zone), &zonePublish, &response)
		})
		log.Infof("Committing changes for zone %s: %+v", zone, errorOrValue(err, &response))
	}

	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return fmt.Errorf("Multiple errors committing: %+v", errs)
	}
}

// Records makes on average C + 2*Z  requests (Z = number of zones): 1 login + 1 fetchAllRecords
// A cache is used to avoid querying for every single record found. C is proportional to the number
// of expired/changed records
func (d *dynProviderState) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	client, err := d.login()
	if err != nil {
		return nil, err
	}
	defer client.Logout()

	log.Debugf("Using DynAPI@%s", d.DynVersion)

	var result []*endpoint.Endpoint

	zones := d.zones(client)
	log.Infof("Configured zones: %+v", zones)
	for _, zone := range zones {
		serial, err := d.fetchZoneSerial(client, zone)
		if err != nil {
			if strings.Contains(err.Error(), "404 Not Found") {
				log.Infof("Ignore zone %s as it does not exist", zone)
				continue
			}

			return nil, err
		}

		relevantRecords := d.ZoneSnapshot.GetRecordsForSerial(zone, serial)
		if relevantRecords != nil {
			log.Infof("Using %d cached records for zone %s@%d", len(relevantRecords), zone, serial)
			result = append(result, relevantRecords...)
			continue
		}

		//Fetch All Records
		records, err := d.fetchAllRecordsInZone(zone)
		if err != nil {
			return nil, err
		}
		relevantRecords = d.allRecordsToEndpoints(records)

		log.Debugf("Relevant records %+v", relevantRecords)

		d.ZoneSnapshot.StoreRecordsForSerial(zone, serial, relevantRecords)
		log.Infof("Stored %d records for %s@%d", len(relevantRecords), zone, serial)
		result = append(result, relevantRecords...)
	}

	return result, nil
}

// this method does C + 2*Z requests: C=total number of changes, Z = number of
// affected zones (1 login + 1 commit)
func (d *dynProviderState) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	log.Debugf("Processing chages: %+v", changes)

	if d.DryRun {
		log.Infof("Will NOT delete these records: %+v", changes.Delete)
		log.Infof("Will NOT create these records: %+v", changes.Create)
		log.Infof("Will NOT update these records: %+v", merge(changes.UpdateOld, changes.UpdateNew))
		return nil
	}

	client, err := d.login()
	if err != nil {
		return err
	}
	defer client.Logout()

	var errs []error

	needsCommit := false

	for _, ep := range changes.Delete {
		err := d.deleteRecord(client, ep)
		if err != nil {
			errs = append(errs, err)
		} else {
			needsCommit = true
		}
	}

	for _, ep := range changes.Create {
		err := d.createRecord(client, ep)
		if err != nil {
			errs = append(errs, err)
		} else {
			needsCommit = true
		}
	}

	updates := merge(changes.UpdateOld, changes.UpdateNew)
	log.Debugf("Updates after merging: %+v", updates)
	for _, ep := range updates {
		err := d.replaceRecord(client, ep)
		if err != nil {
			errs = append(errs, err)
		} else {
			needsCommit = true
		}
	}

	switch len(errs) {
	case 0:
	case 1:
		return errs[0]
	default:
		return fmt.Errorf("Multiple errors committing: %+v", errs)
	}

	if needsCommit {
		return d.commit(client)
	}

	return nil
}
