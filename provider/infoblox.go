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

package provider

import (
	"os"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	ibclient "github.com/khrisrichardson/infoblox-go-client"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// InfobloxConfig clarifies the method signature
type InfobloxConfig struct {
	DomainFilter DomainFilter
	Host         string
	Port         int
	Username     string
	Password     string
	Version      string
	SSLVerify    bool
	DryRun       bool
}

// InfobloxProvider implements the DNS provider for Infoblox.
type InfobloxProvider struct {
	client       ibclient.IBConnector
	domainFilter DomainFilter
	dryRun       bool
}

type ibObjRes struct {
	obj ibclient.IBObject
	res interface{}
}

// NewInfobloxProvider creates a new Infoblox provider.
func NewInfobloxProvider(infobloxConfig InfobloxConfig) (*InfobloxProvider, error) {
	hostConfig := ibclient.HostConfig{
		Host:     infobloxConfig.Host,
		Port:     strconv.Itoa(infobloxConfig.Port),
		Username: infobloxConfig.Username,
		Password: infobloxConfig.Password,
		Version:  infobloxConfig.Version,
	}

	httpPoolConnections := lookupEnvAtoi("EXTERNAL_DNS_INFOBLOX_HTTP_POOL_CONNECTIONS", 10)
	httpRequestTimeout := lookupEnvAtoi("EXTERNAL_DNS_INFOBLOX_HTTP_REQUEST_TIMEOUT", 60)

	transportConfig := ibclient.NewTransportConfig(
		strconv.FormatBool(infobloxConfig.SSLVerify),
		httpRequestTimeout,
		httpPoolConnections,
	)

	requestBuilder := &ibclient.WapiRequestBuilder{}
	requestor := &ibclient.WapiHttpRequestor{}

	client, err := ibclient.NewConnector(hostConfig, transportConfig, requestBuilder, requestor)

	if err != nil {
		return nil, err
	}

	provider := &InfobloxProvider{
		client:       client,
		domainFilter: infobloxConfig.DomainFilter,
		dryRun:       infobloxConfig.DryRun,
	}

	return provider, nil
}

// Records gets the current records.
func (p *InfobloxProvider) Records() (endpoints []*endpoint.Endpoint, err error) {
	zones, err := p.zones()
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		var resA []ibclient.RecordA
		objA := ibclient.NewRecordA(
			ibclient.RecordA{
				Zone: zone.Fqdn,
			},
		)
		err = p.client.GetObject(objA, "", &resA)
		if err != nil {
			return nil, err
		}
		for _, res := range resA {
			endpoints = append(endpoints, endpoint.NewEndpoint(res.Name, res.Ipv4Addr, endpoint.RecordTypeA))
		}

		// Include Host records since they should be treated synonymously with A records
		var resH []ibclient.RecordHost
		objH := ibclient.NewRecordHost(
			ibclient.RecordHost{
				Zone: zone.Fqdn,
			},
		)
		err = p.client.GetObject(objH, "", &resH)
		if err != nil {
			return nil, err
		}
		for _, res := range resH {
			for _, ip := range res.Ipv4Addrs {
				endpoints = append(endpoints, endpoint.NewEndpoint(res.Name, ip.Ipv4Addr, endpoint.RecordTypeA))
			}
		}

		var resC []ibclient.RecordCNAME
		objC := ibclient.NewRecordCNAME(
			ibclient.RecordCNAME{
				Zone: zone.Fqdn,
			},
		)
		err = p.client.GetObject(objC, "", &resC)
		if err != nil {
			return nil, err
		}
		for _, res := range resC {
			endpoints = append(endpoints, endpoint.NewEndpoint(res.Name, res.Canonical, endpoint.RecordTypeCNAME))
		}

		var resT []ibclient.RecordTXT
		objT := ibclient.NewRecordTXT(
			ibclient.RecordTXT{
				Zone: zone.Fqdn,
			},
		)
		err = p.client.GetObject(objT, "", &resT)
		if err != nil {
			return nil, err
		}
		for _, res := range resT {
			// The Infoblox API strips enclosing double quotes from TXT records lacking whitespace.
			// Unhandled, the missing double quotes would break the extractOwnerID method of the registry package.
			if _, err := strconv.Unquote(res.Text); err != nil {
				res.Text = strconv.Quote(res.Text)
			}
			endpoints = append(endpoints, endpoint.NewEndpoint(res.Name, res.Text, endpoint.RecordTypeTXT))
		}
	}
	return endpoints, nil
}

// ApplyChanges applies the given changes.
func (p *InfobloxProvider) ApplyChanges(changes *plan.Changes) error {
	zones, err := p.zones()
	if err != nil {
		return err
	}

	created, deleted, updated := p.mapChanges(zones, changes)
	p.createRecords(created)
	p.deleteRecords(deleted)
	p.updateRecords(updated)
	return nil
}

func (p *InfobloxProvider) zones() ([]ibclient.ZoneAuth, error) {
	var res, result []ibclient.ZoneAuth
	obj := ibclient.NewZoneAuth(ibclient.ZoneAuth{})
	err := p.client.GetObject(obj, "", &res)

	if err != nil {
		return nil, err
	}

	for _, zone := range res {
		if p.domainFilter.Match(zone.Fqdn) {
			result = append(result, zone)
		}
	}

	return result, nil
}

type infobloxChangeMap map[string][]*endpoint.Endpoint

func (p *InfobloxProvider) mapChanges(zones []ibclient.ZoneAuth, changes *plan.Changes) (infobloxChangeMap, infobloxChangeMap, infobloxChangeMap) {
	created := infobloxChangeMap{}
	deleted := infobloxChangeMap{}
	updated := infobloxChangeMap{}

	mapChange := func(changeMap infobloxChangeMap, change *endpoint.Endpoint) {
		zone := p.findZone(zones, change.DNSName)
		if zone == nil {
			logrus.Infof("Ignoring changes to '%s' because a suitable Infoblox DNS zone was not found.", change.DNSName)
			return
		}
		// Ensure the record type is suitable
		changeMap[zone.Fqdn] = append(changeMap[zone.Fqdn], change)
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

	return created, deleted, updated
}

func (p *InfobloxProvider) findZone(zones []ibclient.ZoneAuth, name string) *ibclient.ZoneAuth {
	var result *ibclient.ZoneAuth

	// Go through every zone looking for the longest name (i.e. most specific) as a matching suffix
	for idx := range zones {
		zone := &zones[idx]
		if strings.HasSuffix(name, zone.Fqdn) {
			if result == nil || len(zone.Fqdn) > len(result.Fqdn) {
				result = zone
			}
		}
	}
	return result
}

func (p *InfobloxProvider) mapObjRes(ep *endpoint.Endpoint, getObject bool) (objRes ibObjRes, err error) {
	var obj ibclient.IBObject
	switch ep.RecordType {
	case endpoint.RecordTypeA:
		var res []ibclient.RecordA
		obj = ibclient.NewRecordA(
			ibclient.RecordA{
				Name:     ep.DNSName,
				Ipv4Addr: ep.Target,
			},
		)
		if getObject {
			err = p.client.GetObject(obj, "", &res)
			if err != nil {
				return
			}
		}
		objRes = ibObjRes{
			obj: obj,
			res: &res,
		}
	case endpoint.RecordTypeCNAME:
		var res []ibclient.RecordCNAME
		obj = ibclient.NewRecordCNAME(
			ibclient.RecordCNAME{
				Name:      ep.DNSName,
				Canonical: ep.Target,
			},
		)
		if getObject {
			err = p.client.GetObject(obj, "", &res)
			if err != nil {
				return
			}
		}
		objRes = ibObjRes{
			obj: obj,
			res: &res,
		}
	case endpoint.RecordTypeTXT:
		var res []ibclient.RecordTXT
		// The Infoblox API strips enclosing double quotes from TXT records lacking whitespace.
		// Here we reconcile that fact by making this state match that reality.
		if target, err := strconv.Unquote(ep.Target); err == nil && !strings.Contains(ep.Target, " ") {
			ep.Target = target
		}
		obj = ibclient.NewRecordTXT(
			ibclient.RecordTXT{
				Name: ep.DNSName,
				Text: ep.Target,
			},
		)
		if getObject {
			err = p.client.GetObject(obj, "", &res)
			if err != nil {
				return
			}
		}
		objRes = ibObjRes{
			obj: obj,
			res: &res,
		}
	}
	return
}

func (p *InfobloxProvider) createRecords(created infobloxChangeMap) {
	for zone, endpoints := range created {
		for _, ep := range endpoints {
			if p.dryRun {
				logrus.Infof(
					"Would create %s record named '%s' to '%s' for Infoblox DNS zone '%s'.",
					ep.RecordType,
					ep.DNSName,
					ep.Target,
					zone,
				)
				continue
			}

			logrus.Infof(
				"Creating %s record named '%s' to '%s' for Infoblox DNS zone '%s'.",
				ep.RecordType,
				ep.DNSName,
				ep.Target,
				zone,
			)

			mapObjRes, err := p.mapObjRes(ep, false)
			if err != nil {
				logrus.Errorf(
					"Failed to retrieve %s record named '%s' to '%s' for DNS zone '%s': %v",
					ep.RecordType,
					ep.DNSName,
					ep.Target,
					zone,
					err,
				)
				continue
			}
			_, err = p.client.CreateObject(mapObjRes.obj)
			if err != nil {
				logrus.Errorf(
					"Failed to create %s record named '%s' to '%s' for DNS zone '%s': %v",
					ep.RecordType,
					ep.DNSName,
					ep.Target,
					zone,
					err,
				)
			}
		}
	}
}

func (p *InfobloxProvider) deleteRecords(deleted infobloxChangeMap) {
	// Delete records first
	for zone, endpoints := range deleted {
		for _, ep := range endpoints {
			if p.dryRun {
				logrus.Infof("Would delete %s record named '%s' for Infoblox DNS zone '%s'.", ep.RecordType, ep.DNSName, zone)
			} else {
				logrus.Infof("Deleting %s record named '%s' for Infoblox DNS zone '%s'.", ep.RecordType, ep.DNSName, zone)
				mapObjRes, err := p.mapObjRes(ep, true)
				if err != nil {
					logrus.Errorf(
						"Failed to retrieve %s record named '%s' to '%s' for DNS zone '%s': %v",
						ep.RecordType,
						ep.DNSName,
						ep.Target,
						zone,
						err,
					)
					continue
				}
				switch ep.RecordType {
				case endpoint.RecordTypeA:
					for _, record := range *mapObjRes.res.(*[]ibclient.RecordA) {
						_, err = p.client.DeleteObject(record.Ref)
					}
				case endpoint.RecordTypeCNAME:
					for _, record := range *mapObjRes.res.(*[]ibclient.RecordCNAME) {
						_, err = p.client.DeleteObject(record.Ref)
					}
				case endpoint.RecordTypeTXT:
					for _, record := range *mapObjRes.res.(*[]ibclient.RecordTXT) {
						_, err = p.client.DeleteObject(record.Ref)
					}
				}
				if err != nil {
					logrus.Errorf(
						"Failed to delete %s record named '%s' for Infoblox DNS zone '%s': %v",
						ep.RecordType,
						ep.DNSName,
						zone,
						err,
					)
				}
			}
		}
	}
}

func (p *InfobloxProvider) updateRecords(updated infobloxChangeMap) {
	for zone, endpoints := range updated {
		for _, ep := range endpoints {
			if p.dryRun {
				logrus.Infof(
					"Would update %s record named '%s' to '%s' for Infoblox DNS zone '%s'.",
					ep.RecordType,
					ep.DNSName,
					ep.Target,
					zone,
				)
				continue
			}

			logrus.Infof(
				"Updating %s record named '%s' to '%s' for Infoblox DNS zone '%s'.",
				ep.RecordType,
				ep.DNSName,
				ep.Target,
				zone,
			)

			mapObjRes, err := p.mapObjRes(ep, true)
			if err != nil {
				logrus.Errorf(
					"Failed to retrieve %s record named '%s' to '%s' for DNS zone '%s': %v",
					ep.RecordType,
					ep.DNSName,
					ep.Target,
					zone,
					err,
				)
				continue
			}
			switch ep.RecordType {
			case endpoint.RecordTypeA:
				for _, record := range *mapObjRes.res.(*[]ibclient.RecordA) {
					_, err = p.client.UpdateObject(mapObjRes.obj, record.Ref)
				}
			case endpoint.RecordTypeCNAME:
				for _, record := range *mapObjRes.res.(*[]ibclient.RecordCNAME) {
					_, err = p.client.UpdateObject(mapObjRes.obj, record.Ref)
				}
			case endpoint.RecordTypeTXT:
				for _, record := range *mapObjRes.res.(*[]ibclient.RecordTXT) {
					_, err = p.client.UpdateObject(mapObjRes.obj, record.Ref)
				}
			}
			if err != nil {
				logrus.Errorf(
					"Failed to update %s record named '%s' to '%s' for DNS zone '%s': %v",
					ep.RecordType,
					ep.DNSName,
					ep.Target,
					zone,
					err,
				)
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
