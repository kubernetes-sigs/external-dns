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
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	ibclient "github.com/infobloxopen/infoblox-go-client"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// InfobloxProvider implements the DNS provider for Infoblox.
type InfobloxProvider struct {
	client       ibclient.IBConnector
	domainFilter DomainFilter
	dryRun       bool
}

// NewInfobloxProvider creates a new Infoblox provider.
func NewInfobloxProvider(
	domainFilter DomainFilter,
	host string,
	port int,
	username string,
	password string,
	version string,
	sslVerify bool,
	dryRun bool) (*InfobloxProvider, error) {
	hostConfig := ibclient.HostConfig{
		Host:     host,
		Port:     strconv.Itoa(port),
		Username: username,
		Password: password,
		Version:  version,
	}

	httpPoolConnections, _ := strconv.ParseInt(getEnv("EXTERNAL_DNS_INFOBLOX_HTTP_POOL_CONNECTIONS", "10"), 0, 0)
	httpRequestTimeout, _ := strconv.ParseInt(getEnv("EXTERNAL_DNS_INFOBLOX_HTTP_REQUEST_TIMEOUT", "60"), 0, 0)

	transportConfig := ibclient.NewTransportConfig(
		strconv.FormatBool(sslVerify),
		int(httpRequestTimeout),
		int(httpPoolConnections),
	)

	requestBuilder := &ibclient.WapiRequestBuilder{}
	requestor := &ibclient.WapiHttpRequestor{}

	client, err := ibclient.NewConnector(hostConfig, transportConfig, requestBuilder, requestor)

	if err != nil {
		return nil, err
	}

	provider := &InfobloxProvider{
		client:       client,
		domainFilter: domainFilter,
		dryRun:       dryRun,
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
	ignored := map[string]bool{}
	created := infobloxChangeMap{}
	deleted := infobloxChangeMap{}
	updated := infobloxChangeMap{}

	mapChange := func(changeMap infobloxChangeMap, change *endpoint.Endpoint) {
		zone := p.findZone(zones, change.DNSName)
		if zone == nil {
			if _, ok := ignored[change.DNSName]; !ok {
				ignored[change.DNSName] = true
				logrus.Infof("Ignoring changes to '%s' because a suitable Infoblox DNS zone was not found.", change.DNSName)
			}
			return
		}
		// Ensure the record type is suitable
		changeMap[zone.Fqdn] = append(changeMap[zone.Fqdn], change)
	}

	for _, change := range changes.Create {
		mapChange(created, change)
	}
	for _, change := range changes.Delete {
		mapChange(deleted, change)
	}
	for _, change := range changes.UpdateNew {
		mapChange(updated, change)
	}
	for _, change := range changes.UpdateOld {
		mapChange(deleted, change)
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

func (p *InfobloxProvider) newObject(e *endpoint.Endpoint, getRef bool) (obj ibclient.IBObject, ref string, err error) {
	switch e.RecordType {
	case endpoint.RecordTypeA:
		var records []ibclient.RecordA
		obj = ibclient.NewRecordA(
			ibclient.RecordA{
				Name:     e.DNSName,
				Ipv4Addr: e.Target,
			},
		)
		if getRef {
			p.client.GetObject(obj, "", &records)
			if len(records) == 1 {
				ref = records[0].Ref
			}
		}
		return obj, ref, nil
	case endpoint.RecordTypeCNAME:
		var records []ibclient.RecordCNAME
		obj = ibclient.NewRecordCNAME(
			ibclient.RecordCNAME{
				Name:      e.DNSName,
				Canonical: e.Target,
			},
		)
		if getRef {
			p.client.GetObject(obj, "", &records)
			if len(records) == 1 {
				ref = records[0].Ref
			}
		}
		return obj, ref, nil
	case endpoint.RecordTypeTXT:
		var records []ibclient.RecordTXT
		if target, err := strconv.Unquote(e.Target); err == nil && !strings.Contains(e.Target, " ") {
			e.Target = target
		}
		obj = ibclient.NewRecordTXT(
			ibclient.RecordTXT{
				Name: e.DNSName,
				Text: e.Target,
			},
		)
		if getRef {
			p.client.GetObject(obj, "", &records)
			if len(records) == 1 {
				ref = records[0].Ref
			}
		}
		return obj, ref, nil
	default:
		return nil, "", fmt.Errorf("unsupported record type '%s'", e.RecordType)
	}
}

func (p *InfobloxProvider) createRecords(created infobloxChangeMap) {
	for zone, endpoints := range created {
		for _, endpoint := range endpoints {
			if p.dryRun {
				logrus.Infof(
					"Would create %s record named '%s' to '%s' for Infoblox DNS zone '%s'.",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Target,
					zone,
				)
				continue
			}

			logrus.Infof(
				"Creating %s record named '%s' to '%s' for Infoblox DNS zone '%s'.",
				endpoint.RecordType,
				endpoint.DNSName,
				endpoint.Target,
				zone,
			)

			obj, _, err := p.newObject(endpoint, false)
			if err == nil {
				_, err = p.client.CreateObject(obj)
			}
			if err != nil {
				logrus.Errorf(
					"Failed to create %s record named '%s' to '%s' for DNS zone '%s': %v",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Target,
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
		for _, endpoint := range endpoints {
			if p.dryRun {
				logrus.Infof("Would delete %s record named '%s' for Infoblox DNS zone '%s'.", endpoint.RecordType, endpoint.DNSName, zone)
			} else {
				logrus.Infof("Deleting %s record named '%s' for Infoblox DNS zone '%s'.", endpoint.RecordType, endpoint.DNSName, zone)
				_, ref, err := p.newObject(endpoint, true)
				if err == nil {
					if _, err = p.client.DeleteObject(ref); err != nil {
						logrus.Errorf(
							"Failed to delete %s record named '%s' for Infoblox DNS zone '%s': %v",
							endpoint.RecordType,
							endpoint.DNSName,
							zone,
							err,
						)
					}
				}
			}
		}
	}
}

func (p *InfobloxProvider) updateRecords(updated infobloxChangeMap) {
	for zone, endpoints := range updated {
		for _, endpoint := range endpoints {
			if p.dryRun {
				logrus.Infof(
					"Would update %s record named '%s' to '%s' for Infoblox DNS zone '%s'.",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Target,
					zone,
				)
				continue
			}

			logrus.Infof(
				"Updating %s record named '%s' to '%s' for Infoblox DNS zone '%s'.",
				endpoint.RecordType,
				endpoint.DNSName,
				endpoint.Target,
				zone,
			)

			obj, ref, err := p.newObject(endpoint, true)
			if err == nil {
				_, err = p.client.UpdateObject(obj, ref)
			}
			if err != nil {
				logrus.Errorf(
					"Failed to update %s record named '%s' to '%s' for DNS zone '%s': %v",
					endpoint.RecordType,
					endpoint.DNSName,
					endpoint.Target,
					zone,
					err,
				)
			}
		}
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
