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

package infoblox

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/StackExchange/dnscontrol/pkg/transform"
	ibclient "github.com/infobloxopen/infoblox-go-client/v2"
	"github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	// provider specific key to track if PTR-record was already created or not for A-records
	providerSpecificInfobloxPtrRecord = "infoblox-ptr-record-exists"
)

// Regular expression for domain name validation. In compliance with https://datatracker.ietf.org/doc/html/rfc1035
const dnsLabel = "[a-z]+(([a-z0-9-]*[a-z0-9]+))?"

var domainRegexpTemplate = fmt.Sprintf("^(?i)%s(\\.%s)*\\.?$", dnsLabel, dnsLabel)
var domainRegexp = regexp.MustCompile(domainRegexpTemplate)

func isDomainValid(name string) bool {
	return domainRegexp.MatchString(name)
}

func isNotFoundError(err error) bool {
	_, ok := err.(*ibclient.NotFoundError)
	return ok
}

// StartupConfig clarifies the method signature
type StartupConfig struct {
	DomainFilter  endpoint.DomainFilter
	ZoneIDFilter  provider.ZoneIDFilter
	Host          string
	Port          int
	Username      string
	Password      string
	Version       string
	SSLVerify     bool
	ClientCert    string
	ClientKey     string
	HostRO        string
	PortRO        int
	UsernameRO    string
	PasswordRO    string
	VersionRO     string
	SSLVerifyRO   bool
	ClientCertRO  string
	ClientKeyRO   string
	DryRun        bool
	View          string
	MaxResults    int
	FQDNRegEx     string
	CreatePTR     bool
	CacheDuration int
}

// ProviderConfig implements the DNS provider for Infoblox.
type ProviderConfig struct {
	provider.BaseProvider
	clientRW      ibclient.IBConnector
	clientRO      ibclient.IBConnector
	domainFilter  endpoint.DomainFilter
	zoneIDFilter  provider.ZoneIDFilter
	view          string
	dryRun        bool
	fqdnRegEx     *regexp.Regexp
	createPTR     bool
	cacheDuration int
}

type infobloxRecordSet struct {
	obj ibclient.IBObject
	res interface{}
}

// ExtendedRequestBuilder implements a HttpRequestBuilder which sets
// additional query parameter on all get requests
type ExtendedRequestBuilder struct {
	fqdnRegEx  string
	maxResults int
	ibclient.WapiRequestBuilder
}

// NewExtendedRequestBuilder returns a ExtendedRequestBuilder which adds
// _max_results query parameter to all GET requests
func NewExtendedRequestBuilder(maxResults int, fqdnRegEx string) *ExtendedRequestBuilder {
	return &ExtendedRequestBuilder{
		fqdnRegEx:  fqdnRegEx,
		maxResults: maxResults,
	}
}

// BuildRequest prepares the api request. it uses BuildRequest of
// WapiRequestBuilder and then add the _max_requests parameter
func (mrb *ExtendedRequestBuilder) BuildRequest(t ibclient.RequestType, obj ibclient.IBObject, ref string, queryParams *ibclient.QueryParams) (req *http.Request, err error) {
	req, err = mrb.WapiRequestBuilder.BuildRequest(t, obj, ref, queryParams)
	if req.Method == "GET" {
		query := req.URL.Query()
		if mrb.maxResults > 0 {
			query.Set("_max_results", strconv.Itoa(mrb.maxResults))
		}
		_, ok := obj.(*ibclient.ZoneAuth)
		if ok && t == ibclient.GET && mrb.fqdnRegEx != "" {
			query.Set("fqdn~", mrb.fqdnRegEx)
		}
		req.URL.RawQuery = query.Encode()
	}
	return
}

// NewInfobloxProvider creates a new Infoblox provider.
func NewInfobloxProvider(ibStartupCfg StartupConfig) (*ProviderConfig, error) {
	var (
		authCfgRW, authCfgRO ibclient.AuthConfig
		authMethodCertRW     bool
		authMethodCertRO     bool
		err                  error
		requestBuilderRO     ibclient.HttpRequestBuilder
	)

	if strings.TrimSpace(ibStartupCfg.View) == "" {
		return nil, fmt.Errorf("non-empty DNS view's name is required")
	}

	if (ibStartupCfg.ClientCert == "") != (ibStartupCfg.ClientKey == "") {
		return nil, fmt.Errorf("to use certificate-based authentication you must specify BOTH certificate's and private key's file paths")
	}

	roEndpointInUse := true
	if (ibStartupCfg.HostRO == "" || ibStartupCfg.HostRO == ibStartupCfg.Host) &&
		ibStartupCfg.PortRO == ibStartupCfg.Port {

		roEndpointInUse = false // no separate read-only endpoint is used
	}

	hostCfgRW := ibclient.HostConfig{
		Host:    ibStartupCfg.Host,
		Port:    strconv.Itoa(ibStartupCfg.Port),
		Version: ibStartupCfg.Version,
	}

	hostCfgRO := ibclient.HostConfig{
		Host:    ibStartupCfg.HostRO,
		Port:    strconv.Itoa(ibStartupCfg.PortRO),
		Version: ibStartupCfg.VersionRO,
	}

	if ibStartupCfg.ClientCert != "" {
		if authCfgRW.ClientCert, err = ioutil.ReadFile(ibStartupCfg.ClientCert); err != nil {
			return nil, err
		}

		if authCfgRW.ClientKey, err = ioutil.ReadFile(ibStartupCfg.ClientKey); err != nil {
			return nil, err
		}

		authMethodCertRW = true
	} else {
		if ibStartupCfg.Username == "" || ibStartupCfg.Password == "" {
			return nil, fmt.Errorf("either username AND password or certificate AND key MUST be specified")
		}
		authCfgRW.Username = ibStartupCfg.Username
		authCfgRW.Password = ibStartupCfg.Password
	}

	if roEndpointInUse {
		if ibStartupCfg.UsernameRO == "" &&
			ibStartupCfg.PasswordRO == "" &&
			ibStartupCfg.ClientKeyRO == "" &&
			ibStartupCfg.ClientCertRO == "" {

			authCfgRO = authCfgRW
		} else {
			if (ibStartupCfg.ClientCertRO == "") != (ibStartupCfg.ClientKeyRO == "") {
				return nil, fmt.Errorf("to use certificate-based authentication you must specify BOTH certificate's and private key's file paths")
			}
			if ibStartupCfg.ClientCertRO != "" {
				if authCfgRO.ClientCert, err = ioutil.ReadFile(ibStartupCfg.ClientCertRO); err != nil {
					return nil, err
				}

				if authCfgRO.ClientKey, err = ioutil.ReadFile(ibStartupCfg.ClientKeyRO); err != nil {
					return nil, err
				}

				authMethodCertRO = true
			} else {
				if ibStartupCfg.UsernameRO == "" || ibStartupCfg.PasswordRO == "" {
					return nil, fmt.Errorf("either username AND password or certificate AND key MUST be specified")
				}
				authCfgRO.Username = ibStartupCfg.UsernameRO
				authCfgRO.Password = ibStartupCfg.PasswordRO
			}
		}
	} else {
		authCfgRO = authCfgRW
		hostCfgRO = hostCfgRW
	}

	httpPoolConnections := lookupEnvAtoi("EXTERNAL_DNS_INFOBLOX_HTTP_POOL_CONNECTIONS", 10)
	httpRequestTimeout := lookupEnvAtoi("EXTERNAL_DNS_INFOBLOX_HTTP_REQUEST_TIMEOUT", 60)

	transportConfigRW := ibclient.NewTransportConfig(
		strconv.FormatBool(ibStartupCfg.SSLVerify),
		httpRequestTimeout,
		httpPoolConnections,
	)

	transportConfigRO := ibclient.NewTransportConfig(
		strconv.FormatBool(ibStartupCfg.SSLVerifyRO),
		httpRequestTimeout,
		httpPoolConnections,
	)

	if ibStartupCfg.MaxResults != 0 || ibStartupCfg.FQDNRegEx != "" {
		// use our own HttpRequestBuilder which sets _max_results parameter on GET requests
		requestBuilderRO = NewExtendedRequestBuilder(ibStartupCfg.MaxResults, ibStartupCfg.FQDNRegEx)
	} else {
		// Use the default HttpRequestBuilder of the Infoblox client
		// It will be initialized later, in ibclient.NewConnector().
		requestBuilderRO = &ibclient.WapiRequestBuilder{}
	}

	clientRO, err := ibclient.NewConnector(
		hostCfgRO, authCfgRO, transportConfigRO, requestBuilderRO, &ibclient.WapiHttpRequestor{})
	if err != nil {
		return nil, err
	}

	clientRW, err := ibclient.NewConnector(
		hostCfgRW, authCfgRW, transportConfigRW, &ibclient.WapiRequestBuilder{}, &ibclient.WapiHttpRequestor{})
	if err != nil {
		return nil, err
	}

	providerCfg := &ProviderConfig{
		clientRO:      clientRO,
		clientRW:      clientRW,
		domainFilter:  ibStartupCfg.DomainFilter,
		zoneIDFilter:  ibStartupCfg.ZoneIDFilter,
		dryRun:        ibStartupCfg.DryRun,
		view:          ibStartupCfg.View,
		fqdnRegEx:     regexp.MustCompile(ibStartupCfg.FQDNRegEx),
		createPTR:     ibStartupCfg.CreatePTR,
		cacheDuration: ibStartupCfg.CacheDuration,
	}

	if authMethodCertRW {
		logrus.Infof("client-certificate authentication method is used for connecting to Infoblox NIOS server (read-write endpoint).")
	} else {
		logrus.Infof("password-based authentication method is used for connecting to Infoblox NIOS server (read-write endpoint).")
	}

	if hostCfgRO.Host != "" {
		if authMethodCertRO {
			logrus.Infof("client-certificate authentication method is used for connecting to Infoblox NIOS server (read-only endpoint).")
		} else {
			logrus.Infof("password-based authentication method is used for connecting to Infoblox NIOS server (read-only endpoint).")
		}
	}

	logrus.Infof("endpoint for 'read' requests: %s", net.JoinHostPort(hostCfgRO.Host, hostCfgRO.Port))
	logrus.Infof("endpoint for 'write' requests: %s", net.JoinHostPort(hostCfgRW.Host, hostCfgRW.Port))

	return providerCfg, nil
}

// Records gets the current records.
func (p *ProviderConfig) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, err error) {
	zones, err := p.zones()
	if err != nil {
		return nil, fmt.Errorf("could not fetch zones: %s", err)
	}

	for _, zone := range zones {
		logrus.Debugf("fetching records from zone '%s'", zone.Fqdn)

		// Query params for the majority of records are the same
		queryParams := ibclient.NewQueryParams(
			false,
			map[string]string{
				"view": p.view,
				"zone": zone.Fqdn,
			},
		)

		var resA []ibclient.RecordA
		objA := ibclient.NewEmptyRecordA()
		err = p.clientRO.GetObject(objA, "", queryParams, &resA)
		if err != nil && !isNotFoundError(err) {
			return nil, fmt.Errorf("could not fetch A-records from zone '%s': %s", zone.Fqdn, err)
		}
		for _, res := range resA {
			// Check if endpoint already exists and add to existing endpoint if it does
			foundExisting := false
			for _, ep := range endpoints {
				if ep.DNSName == res.Name && ep.RecordType == endpoint.RecordTypeA {
					foundExisting = true
					duplicateTarget := false

					for _, t := range ep.Targets {
						if t == res.Ipv4Addr {
							duplicateTarget = true
							break
						}
					}

					if duplicateTarget {
						logrus.Debugf("A duplicate target '%s' found for existing A-record '%s'", res.Ipv4Addr, ep.DNSName)
					} else {
						logrus.Debugf("Adding target '%s' to existing A-record '%s'", res.Ipv4Addr, res.Name)
						ep.Targets = append(ep.Targets, res.Ipv4Addr)
					}
					break
				}
			}
			if !foundExisting {
				newEndpoint := endpoint.NewEndpoint(res.Name, endpoint.RecordTypeA, res.Ipv4Addr)
				if p.createPTR {
					newEndpoint.WithProviderSpecific(providerSpecificInfobloxPtrRecord, "true")
				}
				endpoints = append(endpoints, newEndpoint)
			}
		}
		// sort targets so that they are always in same order, as Infoblox NIOS might return them in different order
		for _, ep := range endpoints {
			sort.Sort(ep.Targets)
		}

		// Include Host records since they should be treated synonymously with A-records
		var resH []ibclient.HostRecord
		objH := ibclient.NewEmptyHostRecord()
		err = p.clientRO.GetObject(objH, "", queryParams, &resH)
		if err != nil && !isNotFoundError(err) {
			return nil, fmt.Errorf("could not fetch host records from zone '%s': %s", zone.Fqdn, err)
		}
		for _, res := range resH {
			for _, ip := range res.Ipv4Addrs {
				logrus.Debugf("Record='%s' A(H):'%s'", res.Name, ip.Ipv4Addr)

				// host record is an abstraction in Infoblox that combines A and PTR-records
				// for any host record we already should have a PTR-record in Infoblox, so mark it as created
				newEndpoint := endpoint.NewEndpoint(res.Name, endpoint.RecordTypeA, ip.Ipv4Addr)
				if p.createPTR {
					newEndpoint.WithProviderSpecific(providerSpecificInfobloxPtrRecord, "true")
				}
				endpoints = append(endpoints, newEndpoint)
			}
		}

		var resC []ibclient.RecordCNAME
		objC := ibclient.NewEmptyRecordCNAME()
		err = p.clientRO.GetObject(objC, "", queryParams, &resC)
		if err != nil && !isNotFoundError(err) {
			return nil, fmt.Errorf("could not fetch CNAME-records from zone '%s': %s", zone.Fqdn, err)
		}
		for _, res := range resC {
			logrus.Debugf("Record='%s' CNAME:'%s'", res.Name, res.Canonical)
			endpoints = append(endpoints, endpoint.NewEndpoint(res.Name, endpoint.RecordTypeCNAME, res.Canonical))
		}

		if p.createPTR {
			// Infoblox NIOS doesn't accept reverse zone's fqdn, and instead expects .in-addr.arpa zone
			// so convert our zone fqdn (if it is a correct cidr block) into in-addr.arpa address and pass that to Infoblox NIOS
			// example: 10.196.38.0/24 becomes 38.196.10.in-addr.arpa
			arpaZone, err := transform.ReverseDomainName(zone.Fqdn)
			if err == nil {
				var resPtrStatic, resPtrDynamic []ibclient.RecordPTR
				objP := ibclient.NewEmptyRecordPTR()
				qp := ibclient.NewQueryParams(false, map[string]string{
					"view":    p.view,
					"zone":    arpaZone,
					"creator": "STATIC",
				})
				err = p.clientRO.GetObject(objP, "", qp, &resPtrStatic)
				if err != nil && !isNotFoundError(err) {
					return nil, fmt.Errorf("could not fetch PTR-records from zone '%s': %s", zone.Fqdn, err.Error())
				}
				qp = ibclient.NewQueryParams(false, map[string]string{
					"view":    p.view,
					"zone":    arpaZone,
					"creator": "DYNAMIC",
				})
				err = p.clientRO.GetObject(objP, "", qp, &resPtrDynamic)
				if err != nil && !isNotFoundError(err) {
					return nil, fmt.Errorf("could not fetch PTR-records from zone '%s': %s", zone.Fqdn, err.Error())
				}
				for _, res := range resPtrStatic {
					endpoints = append(endpoints, endpoint.NewEndpoint(res.PtrdName, endpoint.RecordTypePTR, res.Ipv4Addr))
				}
				for _, res := range resPtrDynamic {
					endpoints = append(endpoints, endpoint.NewEndpoint(res.PtrdName, endpoint.RecordTypePTR, res.Ipv4Addr))
				}
			}
		}

		var resT []ibclient.RecordTXT
		objT := ibclient.NewEmptyRecordTXT()
		err = p.clientRO.GetObject(objT, "", queryParams, &resT)
		if err != nil && !isNotFoundError(err) {
			return nil, fmt.Errorf("could not fetch TXT-records from zone '%s': %s", zone.Fqdn, err)
		}
		for _, res := range resT {
			// The Infoblox API strips enclosing double quotes from TXT-records lacking whitespace.
			// Unhandled, the missing double quotes would break the extractOwnerID method of the registry package.
			if _, err := strconv.Unquote(res.Text); err != nil {
				res.Text = strconv.Quote(res.Text)
			}

			foundExisting := false

			for _, ep := range endpoints {
				if ep.DNSName == res.Name && ep.RecordType == endpoint.RecordTypeTXT {
					foundExisting = true
					duplicateTarget := false

					for _, t := range ep.Targets {
						if t == res.Text {
							duplicateTarget = true
							break
						}
					}

					if duplicateTarget {
						logrus.Debugf("A duplicate target '%s' found for existing TXT-record '%s'", res.Text, ep.DNSName)
					} else {
						logrus.Debugf("Adding target '%s' to existing TXT-record '%s'", res.Text, res.Name)
						ep.Targets = append(ep.Targets, res.Text)
					}
					break
				}
			}
			if !foundExisting {
				logrus.Debugf("Record='%s' TXT:'%s'", res.Name, res.Text)
				newEndpoint := endpoint.NewEndpoint(res.Name, endpoint.RecordTypeTXT, res.Text)
				endpoints = append(endpoints, newEndpoint)
			}
		}
	}

	// update A-records that have PTR-record created for them already
	if p.createPTR {
		// save all PTR-records into map for a quick look up
		ptrRecordsMap := make(map[string]bool)
		for _, ptrRecord := range endpoints {
			if ptrRecord.RecordType != endpoint.RecordTypePTR {
				continue
			}
			ptrRecordsMap[ptrRecord.DNSName] = true
		}

		for i := range endpoints {
			if endpoints[i].RecordType != endpoint.RecordTypeA {
				continue
			}
			// if PTR-record already exists for A-record, then mark it as such
			if ptrRecordsMap[endpoints[i].DNSName] {
				found := false
				for j := range endpoints[i].ProviderSpecific {
					if endpoints[i].ProviderSpecific[j].Name == providerSpecificInfobloxPtrRecord {
						endpoints[i].ProviderSpecific[j].Value = "true"
						found = true
					}
				}
				if !found {
					endpoints[i].WithProviderSpecific(providerSpecificInfobloxPtrRecord, "true")
				}
			}
		}
	}
	logrus.Debugf("fetched %d records from Infoblox NIOS", len(endpoints))
	return endpoints, nil
}

func (p *ProviderConfig) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	// Update user specified TTL (0 == disabled)
	for i := range endpoints {
		endpoints[i].RecordTTL = endpoint.TTL(p.cacheDuration)
	}

	if !p.createPTR {
		return endpoints
	}

	// for all A-records, we want to create PTR-records
	// so add provider specific property to track if the record was created or not
	for i := range endpoints {
		if endpoints[i].RecordType == endpoint.RecordTypeA {
			found := false
			for j := range endpoints[i].ProviderSpecific {
				if endpoints[i].ProviderSpecific[j].Name == providerSpecificInfobloxPtrRecord {
					endpoints[i].ProviderSpecific[j].Value = "true"
					found = true
				}
			}
			if !found {
				endpoints[i].WithProviderSpecific(providerSpecificInfobloxPtrRecord, "true")
			}
		}
	}

	return endpoints
}

// ApplyChanges applies the given changes.
func (p *ProviderConfig) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	zones, err := p.zones()
	if err != nil {
		return err
	}

	created, deleted := p.mapChanges(zones, changes)
	p.deleteRecords(deleted)
	p.createRecords(created)
	return nil
}

func (p *ProviderConfig) zones() ([]ibclient.ZoneAuth, error) {
	var res, result []ibclient.ZoneAuth
	obj := ibclient.NewZoneAuth(ibclient.ZoneAuth{})
	qp := ibclient.NewQueryParams(
		false, map[string]string{"view": p.view})
	err := p.clientRO.GetObject(obj, "", qp, &res)
	if err != nil && !isNotFoundError(err) {
		return nil, err
	}

	for _, zone := range res {
		if p.fqdnRegEx != nil && !p.fqdnRegEx.MatchString(zone.Fqdn) {
			continue
		}

		if !p.domainFilter.Match(zone.Fqdn) {
			continue
		}

		if !p.zoneIDFilter.Match(zone.Ref) {
			continue
		}

		result = append(result, zone)
	}

	return result, nil
}

type infobloxChangeMap map[string][]*endpoint.Endpoint

func (p *ProviderConfig) mapChanges(zones []ibclient.ZoneAuth, changes *plan.Changes) (infobloxChangeMap, infobloxChangeMap) {
	created := infobloxChangeMap{}
	deleted := infobloxChangeMap{}

	mapChange := func(changeMap infobloxChangeMap, change *endpoint.Endpoint) {
		zone := p.findZone(zones, change.DNSName)
		if zone == nil {
			logrus.Debugf(
				"Ignoring changes to '%s' because a suitable Infoblox DNS zone was not found; "+
					"check the zone’s existence at the DNS server and domain filters in the ExternalDNS configuration",
				change.DNSName)
			return
		}
		// Ensure the record type is suitable
		changeMap[zone.Fqdn] = append(changeMap[zone.Fqdn], change)

		if p.createPTR && change.RecordType == endpoint.RecordTypeA {
			reverseZone := p.findReverseZone(zones, change.Targets[0])
			if reverseZone == nil {
				logrus.Debugf("Ignoring changes to '%s' because a suitable Infoblox DNS reverse zone was not found; "+
					"check the zone’s existence at the DNS server and domain filters in the ExternalDNS configuration",
					change.Targets[0])
				return
			}
			changecopy := *change
			changecopy.RecordType = endpoint.RecordTypePTR
			changeMap[reverseZone.Fqdn] = append(changeMap[reverseZone.Fqdn], &changecopy)
		}
	}

	for _, change := range changes.Delete {
		mapChange(deleted, change)
	}
	for _, change := range changes.UpdateOld {
		mapChange(deleted, change)
	}
	for _, change := range changes.Create {
		mapChange(created, change)
	}
	for _, change := range changes.UpdateNew {
		mapChange(created, change)
	}

	return created, deleted
}

func isDirectZoneRecord(zoneFqdn, recFqdn string) bool {
	parts := strings.SplitN(recFqdn, ".", 2)
	if parts == nil || len(parts) != 2 {
		return false
	}

	return parts[1] == zoneFqdn
}

func (p *ProviderConfig) findZone(zones []ibclient.ZoneAuth, name string) *ibclient.ZoneAuth {
	var result *ibclient.ZoneAuth

	// Go through every zone looking for the longest name (i.e. most specific) as a matching suffix
	for idx := range zones {
		zone := &zones[idx]
		if isDirectZoneRecord(zone.Fqdn, name) || strings.EqualFold(name, zone.Fqdn) {
			if result == nil || len(zone.Fqdn) > len(result.Fqdn) {
				result = zone
			}
		}
	}
	return result
}

func (p *ProviderConfig) findReverseZone(zones []ibclient.ZoneAuth, name string) *ibclient.ZoneAuth {
	ip := net.ParseIP(name)
	networks := map[int]*ibclient.ZoneAuth{}
	maxMask := 0

	for i, zone := range zones {
		_, rZoneNet, err := net.ParseCIDR(zone.Fqdn)
		if err != nil {
			if !isDomainValid(zone.Fqdn) {
				logrus.Debugf("FQDN '%s' is neither a valid network address nor a valid domain name", zone.Fqdn)
			}
		} else {
			if rZoneNet.Contains(ip) {
				_, mask := rZoneNet.Mask.Size()
				networks[mask] = &zones[i]
				if mask > maxMask {
					maxMask = mask
				}
			}
		}
	}
	return networks[maxMask]
}

func (p *ProviderConfig) recordSet(ep *endpoint.Endpoint, getObject bool, targetIndex int) (recordSet infobloxRecordSet, err error) {
	switch ep.RecordType {
	case endpoint.RecordTypeA:
		var res []ibclient.RecordA
		obj := ibclient.NewEmptyRecordA()
		obj.Name = ep.DNSName
		obj.Ipv4Addr = ep.Targets[targetIndex]
		obj.View = p.view
		sf := map[string]string{
			"view":     obj.View,
			"name":     obj.Name,
			"ipv4addr": obj.Ipv4Addr,
		}
		if getObject {
			err = p.clientRO.GetObject(
				obj, "", ibclient.NewQueryParams(false, sf), &res)
			if err != nil && !isNotFoundError(err) {
				return
			}
		}
		recordSet = infobloxRecordSet{
			obj: obj,
			res: &res,
		}
	case endpoint.RecordTypePTR:
		var res []ibclient.RecordPTR
		obj := ibclient.NewEmptyRecordPTR()
		obj.PtrdName = ep.DNSName
		obj.Ipv4Addr = ep.Targets[targetIndex]
		obj.View = p.view
		sf := map[string]string{
			"view":     obj.View,
			"ptrdname": obj.PtrdName,
			"ipv4addr": obj.Ipv4Addr,
		}
		if getObject {
			var staticPtrs, dynamicPtrs []ibclient.RecordPTR
			sf["creator"] = "STATIC"
			err = p.clientRO.GetObject(
				obj, "", ibclient.NewQueryParams(false, sf), &staticPtrs)
			if err != nil && !isNotFoundError(err) {
				return
			}
			sf["creator"] = "DYNAMIC"
			err = p.clientRO.GetObject(
				obj, "", ibclient.NewQueryParams(false, sf), &dynamicPtrs)
			if err != nil && !isNotFoundError(err) {
				return
			}
			res = append(staticPtrs, dynamicPtrs...)
		}
		recordSet = infobloxRecordSet{
			obj: obj,
			res: &res,
		}
	case endpoint.RecordTypeCNAME:
		var res []ibclient.RecordCNAME
		obj := ibclient.NewEmptyRecordCNAME()
		obj.Name = ep.DNSName
		obj.Canonical = ep.Targets[0]
		obj.View = p.view
		sf := map[string]string{
			"view":      obj.View,
			"name":      obj.Name,
			"canonical": obj.Canonical,
		}
		if getObject {
			err = p.clientRO.GetObject(
				obj, "", ibclient.NewQueryParams(false, sf), &res)
			if err != nil && !isNotFoundError(err) {
				return
			}
		}
		recordSet = infobloxRecordSet{
			obj: obj,
			res: &res,
		}
	case endpoint.RecordTypeTXT:
		var res []ibclient.RecordTXT
		// The Infoblox API strips enclosing double quotes from TXT-records lacking whitespace.
		// Here we reconcile that fact by making this state match that reality.
		if target, err2 := strconv.Unquote(ep.Targets[0]); err2 == nil && !strings.Contains(ep.Targets[0], " ") {
			ep.Targets = endpoint.Targets{target}
		}
		obj := ibclient.NewRecordTXT(
			p.view,
			"", ep.DNSName, ep.Targets[0], 0, false, "", nil)
		sf := map[string]string{
			"view": obj.View,
			"name": obj.Name,
			"text": obj.Text,
		}
		if getObject {
			err = p.clientRO.GetObject(
				obj, "", ibclient.NewQueryParams(false, sf), &res)
			if err != nil && !isNotFoundError(err) {
				return
			}
		}
		recordSet = infobloxRecordSet{
			obj: obj,
			res: &res,
		}
	}
	return
}

func (p *ProviderConfig) createRecords(created infobloxChangeMap) {
	expTypeA := ibclient.NewEmptyRecordA().ObjectType()
	expTypeTXT := ibclient.NewEmptyRecordTXT().ObjectType()
	expTypePTR := ibclient.NewEmptyRecordPTR().ObjectType()
	expTypeCNAME := ibclient.NewEmptyRecordCNAME().ObjectType()

	for zone, endpoints := range created {
		for _, ep := range endpoints {
			for targetIndex := range ep.Targets {
				if p.dryRun {
					logrus.Infof(

						"Would create %s-record named '%s' to '%s' for Infoblox DNS zone '%s'.",
						ep.RecordType,
						ep.DNSName,
						ep.Targets[targetIndex],
						zone,
					)
					continue
				}

				logrus.Infof(
					"Creating %s-record named '%s' to '%s' for Infoblox DNS zone '%s'.",
					ep.RecordType,
					ep.DNSName,
					ep.Targets[targetIndex],
					zone,
				)

				recordSet, err := p.recordSet(ep, false, targetIndex)
				if err != nil && !isNotFoundError(err) {
					logrus.Errorf(
						"Failed to retrieve %s-record named '%s' to '%s' for DNS zone '%s': %s",
						ep.RecordType,
						ep.DNSName,
						ep.Targets[targetIndex],
						zone,
						err,
					)
					continue
				}
				actType := recordSet.obj.ObjectType()
				switch actType {
				case expTypeA:
					recA := recordSet.obj.(*ibclient.RecordA)
					if !isDomainValid(recA.Name) {
						logrus.Errorf("cannot create an A-record with name '%s', which is invalid", recA.Name)
						continue
					}
				case expTypeTXT:
					recTXT := recordSet.obj.(*ibclient.RecordTXT)
					if !isDomainValid(recTXT.Name) {
						logrus.Errorf("cannot create a TXT-record with name '%s', which is invalid", recTXT.Name)
						continue
					}
				case expTypePTR:
					recPTR := recordSet.obj.(*ibclient.RecordPTR)
					if recPTR.Name != "" && !isDomainValid(recPTR.Name) {
						logrus.Errorf("cannot create a PTR-record with name '%s', which is invalid", recPTR.Name)
						continue
					}
					if !isDomainValid(recPTR.PtrdName) {
						logrus.Errorf("cannot create a PTR-record with ptrdname '%s', which is invalid", recPTR.PtrdName)
						continue
					}
				case expTypeCNAME:
					recCNAME := recordSet.obj.(*ibclient.RecordCNAME)
					if !isDomainValid(recCNAME.Name) {
						logrus.Errorf("cannot create a CNAME-record with name '%s', which is invalid", recCNAME.Name)
						continue
					}
					if !isDomainValid(recCNAME.Canonical) {
						logrus.Errorf("cannot create a CNAME-record with canonical name '%s', which is invalid", recCNAME.Canonical)
						continue
					}
				}

				_, err = p.clientRW.CreateObject(recordSet.obj)
				if err != nil {
					logrus.Errorf(
						"Failed to create %s-record named '%s' to '%s' for DNS zone '%s': %v",
						ep.RecordType,
						ep.DNSName,
						ep.Targets[targetIndex],
						zone,
						err,
					)
				}
			}
		}
	}
}

func (p *ProviderConfig) deleteRecords(deleted infobloxChangeMap) {
	for zone, endpoints := range deleted {
		for _, ep := range endpoints {
			for targetIndex := range ep.Targets {
				recordSet, err := p.recordSet(ep, true, targetIndex)
				if err != nil && !isNotFoundError(err) {
					logrus.Errorf(
						"Failed to retrieve %s-record named '%s' to '%s' for DNS zone '%s': %s",
						ep.RecordType,
						ep.DNSName,
						ep.Targets[targetIndex],
						zone,
						err,
					)
					continue
				}
				switch ep.RecordType {
				case endpoint.RecordTypeA:
					for _, record := range *recordSet.res.(*[]ibclient.RecordA) {
						if p.dryRun {
							logrus.Infof("Would delete A-record named '%s' to '%s' from Infoblox DNS zone '%s'.", record.Name, record.Ipv4Addr, record.Zone)
						} else {
							logrus.Infof("Deleting A-record named '%s' to '%s' from Infoblox DNS zone '%s'.", record.Name, record.Ipv4Addr, record.Zone)
							_, err = p.clientRW.DeleteObject(record.Ref)
						}
					}
				case endpoint.RecordTypePTR:
					for _, record := range *recordSet.res.(*[]ibclient.RecordPTR) {
						if p.dryRun {
							logrus.Infof("Would delete PTR-record named '%s' to '%s' from Infoblox DNS zone '%s'.", record.PtrdName, record.Ipv4Addr, record.Zone)
						} else {
							logrus.Infof("Deleting PTR-record named '%s' to '%s' from Infoblox DNS zone '%s'.", record.PtrdName, record.Ipv4Addr, record.Zone)
							_, err = p.clientRW.DeleteObject(record.Ref)
						}
					}
				case endpoint.RecordTypeCNAME:
					for _, record := range *recordSet.res.(*[]ibclient.RecordCNAME) {
						if p.dryRun {
							logrus.Infof("Would delete CNAME-record named '%s' to '%s' from Infoblox DNS zone '%s'.", record.Name, record.Canonical, record.Zone)
						} else {
							logrus.Infof("Deleting CNAME-record named '%s' to '%s' from Infoblox DNS zone '%s'.", record.Name, record.Canonical, record.Zone)
							_, err = p.clientRW.DeleteObject(record.Ref)
						}
					}
				case endpoint.RecordTypeTXT:
					for _, record := range *recordSet.res.(*[]ibclient.RecordTXT) {
						if p.dryRun {
							logrus.Infof("Would delete TXT-record named '%s' to '%s' from Infoblox DNS zone '%s'.", record.Name, record.Text, record.Zone)
						} else {
							logrus.Infof("Deleting TXT-record named '%s' to '%s' from Infoblox DNS zone '%s'.", record.Name, record.Text, record.Zone)
							_, err = p.clientRW.DeleteObject(record.Ref)
						}
					}
				}
				if err != nil && !isNotFoundError(err) {
					logrus.Errorf(
						"Failed to delete %s-record named '%s' to '%s' from Infoblox DNS zone '%s': %s",
						ep.RecordType,
						ep.DNSName,
						ep.Targets[targetIndex],
						zone,
						err,
					)
				}
			}
		}
	}
}

func lookupEnvAtoi(key string, fallback int) (i int) {
	val, ok := os.LookupEnv(key)
	if !ok {
		i = fallback
		return
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		i = fallback
		return
	}
	return
}
