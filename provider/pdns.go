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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"sort"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"crypto/tls"
	"crypto/x509"
	pgo "github.com/ffledgling/pdns-go"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"io/ioutil"
	"net"
)

type pdnsChangeType string

const (
	apiBase = "/api/v1"

	// Unless we use something like pdnsproxy (discontinued upsteam), this value will _always_ be localhost
	defaultServerID = "localhost"
	defaultTTL      = 300

	// PdnsDelete and PdnsReplace are effectively an enum for "pgo.RrSet.changetype"
	// TODO: Can we somehow get this from the pgo swagger client library itself?

	// PdnsDelete : PowerDNS changetype used for deleting rrsets
	// ref: https://doc.powerdns.com/authoritative/http-api/zone.html#rrset (see "changetype")
	PdnsDelete pdnsChangeType = "DELETE"
	// PdnsReplace : PowerDNS changetype for creating, updating and patching rrsets
	PdnsReplace pdnsChangeType = "REPLACE"
	// Number of times to retry failed PDNS requests
	retryLimit = 3
	// time in milliseconds
	retryAfterTime = 250 * time.Millisecond
)

// PDNSConfig is comprised of the fields necessary to create a new PDNSProvider
type PDNSConfig struct {
	DomainFilter DomainFilter
	DryRun       bool
	Server       string
	APIKey       string
	TLSConfig    TLSConfig
}

// TLSConfig is comprised of the TLS-related fields necessary to create a new PDNSProvider
type TLSConfig struct {
	TLSEnabled            bool
	CAFilePath            string
	ClientCertFilePath    string
	ClientCertKeyFilePath string
}

func (tlsConfig *TLSConfig) setHTTPClient(pdnsClientConfig *pgo.Configuration) error {
	if !tlsConfig.TLSEnabled {
		if tlsConfig.CAFilePath != "" {
			return errors.New("certificate authority file path was specified, but TLS was not enabled")
		}
		if tlsConfig.ClientCertFilePath != "" {
			return errors.New("client certificate file path was specified, but TLS was not enabled")
		}
		if tlsConfig.ClientCertKeyFilePath != "" {
			return errors.New("client certificate key file path was specified, but TLS was not enabled")
		}
		return nil
	}

	log.Debug("Configuring TLS for PDNS Provider.")
	if tlsConfig.CAFilePath == "" {
		return errors.New("certificate authority file path must be specified if TLS is enabled")
	}
	if tlsConfig.ClientCertFilePath == "" && tlsConfig.ClientCertKeyFilePath != "" ||
		tlsConfig.ClientCertFilePath != "" && tlsConfig.ClientCertKeyFilePath == "" {
		return errors.New("client certificate and client certificate key should be specified together if at all")
	}

	certificateAuthority, err := loadCertificateAuthority(tlsConfig.CAFilePath)
	if err != nil {
		return err
	}

	certificate, err := loadCertificate(tlsConfig.ClientCertFilePath, tlsConfig.ClientCertKeyFilePath)
	if err != nil {
		return err
	}

	// Timeouts taken from net.http.DefaultTransport
	transporter := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: certificate,
			RootCAs:      certificateAuthority,
		},
	}
	pdnsClientConfig.HTTPClient = &http.Client{
		Transport: transporter,
	}

	return nil
}

func loadCertificateAuthority(certificateAuthorityFilePath string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()

	pem, err := ioutil.ReadFile(certificateAuthorityFilePath)
	if err != nil {
		return nil, err
	}

	ok := pool.AppendCertsFromPEM(pem)
	if !ok {
		return nil, errors.New("error appending certificate to pool")
	}

	return pool, nil
}

func loadCertificate(certificateFilePath string, certificateKeyFilePath string) ([]tls.Certificate, error) {
	if certificateFilePath == "" || certificateKeyFilePath == "" {
		return []tls.Certificate{}, nil
	}
	certificate, err := tls.LoadX509KeyPair(certificateFilePath, certificateKeyFilePath)
	if err != nil {
		return nil, err
	}
	return []tls.Certificate{certificate}, nil
}

// Function for debug printing
func stringifyHTTPResponseBody(r *http.Response) (body string) {

	if r == nil {
		return ""
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body = buf.String()
	return body

}

// PDNSAPIProvider : Interface used and extended by the PDNSAPIClient struct as
// well as mock APIClients used in testing
type PDNSAPIProvider interface {
	ListZones() ([]pgo.Zone, *http.Response, error)
	ListZone(zoneID string) (pgo.Zone, *http.Response, error)
	PatchZone(zoneID string, zoneStruct pgo.Zone) (*http.Response, error)
}

// PDNSAPIClient : Struct that encapsulates all the PowerDNS specific implementation details
type PDNSAPIClient struct {
	dryRun  bool
	authCtx context.Context
	client  *pgo.APIClient
}

// ListZones : Method returns all enabled zones from PowerDNS
// ref: https://doc.powerdns.com/authoritative/http-api/zone.html#get--servers-server_id-zones
func (c *PDNSAPIClient) ListZones() (zones []pgo.Zone, resp *http.Response, err error) {
	for i := 0; i < retryLimit; i++ {
		zones, resp, err = c.client.ZonesApi.ListZones(c.authCtx, defaultServerID)
		if err != nil {
			log.Debugf("Unable to fetch zones %v", err)
			log.Debugf("Retrying ListZones() ... %d", i)
			time.Sleep(retryAfterTime * (1 << uint(i)))
			continue

		}
		return zones, resp, err
	}

	log.Errorf("Unable to fetch zones. %v", err)
	return zones, resp, err

}

// ListZone : Method returns the details of a specific zone from PowerDNS
// ref: https://doc.powerdns.com/authoritative/http-api/zone.html#get--servers-server_id-zones-zone_id
func (c *PDNSAPIClient) ListZone(zoneID string) (zone pgo.Zone, resp *http.Response, err error) {
	for i := 0; i < retryLimit; i++ {
		zone, resp, err = c.client.ZonesApi.ListZone(c.authCtx, defaultServerID, zoneID)
		if err != nil {
			log.Debugf("Unable to fetch zone %v", err)
			log.Debugf("Retrying ListZone() ... %d", i)
			time.Sleep(retryAfterTime * (1 << uint(i)))
			continue

		}
		return zone, resp, err
	}

	log.Errorf("Unable to list zone. %v", err)
	return zone, resp, err

}

// PatchZone : Method used to update the contents of a particular zone from PowerDNS
// ref: https://doc.powerdns.com/authoritative/http-api/zone.html#patch--servers-server_id-zones-zone_id
func (c *PDNSAPIClient) PatchZone(zoneID string, zoneStruct pgo.Zone) (resp *http.Response, err error) {
	for i := 0; i < retryLimit; i++ {
		resp, err = c.client.ZonesApi.PatchZone(c.authCtx, defaultServerID, zoneID, zoneStruct)
		if err != nil {
			log.Debugf("Unable to patch zone %v", err)
			log.Debugf("Retrying PatchZone() ... %d", i)
			time.Sleep(retryAfterTime * (1 << uint(i)))
			continue

		}
		return resp, err
	}

	log.Errorf("Unable to patch zone. %v", err)
	return resp, err
}

// PDNSProvider is an implementation of the Provider interface for PowerDNS
type PDNSProvider struct {
	client PDNSAPIProvider
}

// NewPDNSProvider initializes a new PowerDNS based Provider.
func NewPDNSProvider(config PDNSConfig) (*PDNSProvider, error) {

	// Do some input validation

	if config.APIKey == "" {
		return nil, errors.New("Missing API Key for PDNS. Specify using --pdns-api-key=")
	}

	// The default for when no --domain-filter is passed is [""], instead of [], so we check accordingly.
	if len(config.DomainFilter.filters) != 1 && config.DomainFilter.filters[0] != "" {
		return nil, errors.New("PDNS Provider does not support domain filter")
	}
	// We do not support dry running, exit safely instead of surprising the user
	// TODO: Add Dry Run support
	if config.DryRun {
		return nil, errors.New("PDNS Provider does not currently support dry-run")
	}

	if config.Server == "localhost" {
		log.Warnf("PDNS Server is set to localhost, this may not be what you want. Specify using --pdns-server=")
	}

	pdnsClientConfig := pgo.NewConfiguration()
	pdnsClientConfig.Host = config.Server
	pdnsClientConfig.BasePath = config.Server + apiBase
	if err := config.TLSConfig.setHTTPClient(pdnsClientConfig); err != nil {
		return nil, err
	}

	provider := &PDNSProvider{
		client: &PDNSAPIClient{
			dryRun:  config.DryRun,
			authCtx: context.WithValue(context.TODO(), pgo.ContextAPIKey, pgo.APIKey{Key: config.APIKey}),
			client:  pgo.NewAPIClient(pdnsClientConfig),
		},
	}

	return provider, nil
}

func (p *PDNSProvider) convertRRSetToEndpoints(rr pgo.RrSet) (endpoints []*endpoint.Endpoint, _ error) {
	endpoints = []*endpoint.Endpoint{}

	for _, record := range rr.Records {
		// If a record is "Disabled", it's not supposed to be "visible"
		if !record.Disabled {
			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(rr.Name, rr.Type_, endpoint.TTL(rr.Ttl), record.Content))
		}
	}

	return endpoints, nil
}

// ConvertEndpointsToZones marshals endpoints into pdns compatible Zone structs
func (p *PDNSProvider) ConvertEndpointsToZones(eps []*endpoint.Endpoint, changetype pdnsChangeType) (zonelist []pgo.Zone, _ error) {

	zonelist = []pgo.Zone{}
	endpoints := make([]*endpoint.Endpoint, len(eps))
	copy(endpoints, eps)

	// Sort the endpoints array so we have deterministic inserts
	sort.SliceStable(endpoints,
		func(i, j int) bool {
			// We only care about sorting endpoints with the same dnsname
			if endpoints[i].DNSName == endpoints[j].DNSName {
				return endpoints[i].RecordType < endpoints[j].RecordType
			}
			return endpoints[i].DNSName < endpoints[j].DNSName
		})

	zones, _, err := p.client.ListZones()
	if err != nil {
		return nil, err
	}

	// Sort the zone by length of the name in descending order, we use this
	// property later to ensure we add a record to the longest matching zone

	sort.SliceStable(zones, func(i, j int) bool { return len(zones[i].Name) > len(zones[j].Name) })

	// NOTE: Complexity of this loop is O(Zones*Endpoints).
	// A possibly faster implementation would be a search of the reversed
	// DNSName in a trie of Zone names, which should be O(Endpoints), but at this point it's not
	// necessary.
	for _, zone := range zones {
		zone.Rrsets = []pgo.RrSet{}
		for i := 0; i < len(endpoints); {
			ep := endpoints[i]
			dnsname := ensureTrailingDot(ep.DNSName)
			if strings.HasSuffix(dnsname, zone.Name) {
				// The assumption here is that there will only ever be one target
				// per (ep.DNSName, ep.RecordType) tuple, which holds true for
				// external-dns v5.0.0-alpha onwards
				records := []pgo.Record{}
				for _, t := range ep.Targets {
					if "CNAME" == ep.RecordType {
						t = ensureTrailingDot(t)
					}

					records = append(records, pgo.Record{Content: t})
				}
				rrset := pgo.RrSet{
					Name:       dnsname,
					Type_:      ep.RecordType,
					Records:    records,
					Changetype: string(changetype),
				}

				// DELETEs explicitly forbid a TTL, therefore only PATCHes need the TTL
				if changetype == PdnsReplace {
					if int64(ep.RecordTTL) > int64(math.MaxInt32) {
						return nil, errors.New("Value of record TTL overflows, limited to int32")
					}
					if ep.RecordTTL == 0 {
						// No TTL was sepecified for the record, we use the default
						rrset.Ttl = int32(defaultTTL)
					} else {
						rrset.Ttl = int32(ep.RecordTTL)
					}
				}

				zone.Rrsets = append(zone.Rrsets, rrset)

				// "pop" endpoint if it's matched
				endpoints = append(endpoints[0:i], endpoints[i+1:]...)
			} else {
				// If we didn't pop anything, we move to the next item in the list
				i++
			}

		}

		if len(zone.Rrsets) > 0 {
			zonelist = append(zonelist, zone)
		}

	}

	// If we still have some endpoints left, it means we couldn't find a matching zone for them
	// We warn instead of hard fail here because we don't want a misconfig to cause everything to go down
	if len(endpoints) > 0 {
		log.Warnf("No matching zones were found for the following endpoints: %+v", endpoints)
	}

	log.Debugf("Zone List generated from Endpoints: %+v", zonelist)

	return zonelist, nil
}

// mutateRecords takes a list of endpoints and creates, replaces or deletes them based on the changetype
func (p *PDNSProvider) mutateRecords(endpoints []*endpoint.Endpoint, changetype pdnsChangeType) error {
	zonelist, err := p.ConvertEndpointsToZones(endpoints, changetype)
	if err != nil {
		return err
	}
	for _, zone := range zonelist {
		jso, err := json.Marshal(zone)
		if err != nil {
			log.Errorf("JSON Marshal for zone struct failed!")
		} else {
			log.Debugf("Struct for PatchZone:\n%s", string(jso))
		}

		resp, err := p.client.PatchZone(zone.Id, zone)
		if err != nil {
			log.Debugf("PDNS API response: %s", stringifyHTTPResponseBody(resp))
			return err
		}

	}
	return nil
}

// Records returns all DNS records controlled by the configured PDNS server (for all zones)
func (p *PDNSProvider) Records() (endpoints []*endpoint.Endpoint, _ error) {

	zones, _, err := p.client.ListZones()
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		z, _, err := p.client.ListZone(zone.Id)
		if err != nil {
			log.Warnf("Unable to fetch Records")
			return nil, err
		}

		for _, rr := range z.Rrsets {
			e, err := p.convertRRSetToEndpoints(rr)
			if err != nil {
				return nil, err
			}
			endpoints = append(endpoints, e...)
		}
	}

	log.Debugf("Records fetched:\n%+v", endpoints)
	return endpoints, nil
}

// ApplyChanges takes a list of changes (endpoints) and updates the PDNS server
// by sending the correct HTTP PATCH requests to a matching zone
func (p *PDNSProvider) ApplyChanges(changes *plan.Changes) error {

	startTime := time.Now()

	// Create
	for _, change := range changes.Create {
		log.Debugf("CREATE: %+v", change)
	}
	// We only attempt to mutate records if there are any to mutate.  A
	// call to mutate records with an empty list of endpoints is still a
	// valid call and a no-op, but we might as well not make the call to
	// prevent unnecessary logging
	if len(changes.Create) > 0 {
		// "Replacing" non-existant records creates them
		err := p.mutateRecords(changes.Create, PdnsReplace)
		if err != nil {
			return err
		}
	}

	// Update
	for _, change := range changes.UpdateOld {
		// Since PDNS "Patches", we don't need to specify the "old"
		// record. The Update New change type will automatically take
		// care of replacing the old RRSet with the new one We simply
		// leave this logging here for information
		log.Debugf("UPDATE-OLD (ignored): %+v", change)
	}

	for _, change := range changes.UpdateNew {
		log.Debugf("UPDATE-NEW: %+v", change)
	}
	if len(changes.UpdateNew) > 0 {
		err := p.mutateRecords(changes.UpdateNew, PdnsReplace)
		if err != nil {
			return err
		}
	}

	// Delete
	for _, change := range changes.Delete {
		log.Debugf("DELETE: %+v", change)
	}
	if len(changes.Delete) > 0 {
		err := p.mutateRecords(changes.Delete, PdnsDelete)
		if err != nil {
			return err
		}
	}

	log.Debugf("Changes pushed out to PowerDNS in %s\n", time.Since(startTime))
	return nil
}
