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

package efficientip

import (
	"context"
	"crypto/tls"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	eip "github.com/efficientip-labs/solidserver-go-client/sdsclient"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type EfficientIPConfig struct {
	DomainFilter endpoint.DomainFilter
	ZoneIDFilter provider.ZoneIDFilter
	DryRun       bool
	Host         string
	Port         int
	Username     string
	Password     string
	SSlVerify    bool
}

type EfficientipClient interface {
	ZonesList() ([]*ZoneAuth, error)
	RecordAdd(rr *endpoint.Endpoint) error
	RecordDelete(rr *endpoint.Endpoint) error
	RecordList(Zone ZoneAuth) (endpoints []*endpoint.Endpoint, _ error)
}

func NewEfficientipAPI(ctx context.Context, config *eip.Configuration) EfficientIPAPI {
	return EfficientIPAPI{
		client:  eip.NewAPIClient(config),
		context: ctx,
	}
}

type EfficientIPAPI struct {
	client  *eip.APIClient
	context context.Context
}

func (e *EfficientIPAPI) ZonesList() ([]*ZoneAuth, error) {
	zones, _, err := e.client.DnsApi.DnsZoneList(e.context).Execute()

	if err.Error() != "" && (!zones.HasSuccess() || !zones.GetSuccess()) {
		return nil, err
	}
	var result []*ZoneAuth
	for _, zone := range zones.GetData() {
		result = append(result, NewZoneAuth(zone))
	}
	return result, nil
}

func (e *EfficientIPAPI) RecordAdd(rr *endpoint.Endpoint) error {
	for _, value := range rr.Targets {
		log.Infof("Creating %s record named '%s' to '%s' for Efficientip",
			rr.RecordType,
			rr.DNSName,
			value,
		)

		ttl := int32(rr.RecordTTL)
		_, _, err := e.client.DnsApi.DnsRrAdd(e.context).DnsRrAddInput(eip.DnsRrAddInput{
			RrName:   &rr.DNSName,
			RrType:   &rr.RecordType,
			RrTtl:    &ttl,
			RrValue1: &value,
		}).Execute()

		if err.Error() != "" {
			log.Errorf("Creation of the RR %v %v  [%v]-> %v : failed!", rr.RecordType, rr.DNSName, ttl, value)
		}
	}
	return nil
}

func (e *EfficientIPAPI) RecordDelete(rr *endpoint.Endpoint) error {
	for _, value := range rr.Targets {
		log.Infof("Deleting %s record named '%s' to '%s' for Efficientip",
			rr.RecordType,
			rr.DNSName,
			value,
		)

		_, _, err := e.client.DnsApi.DnsRrDelete(e.context).RrName(rr.DNSName).RrType(rr.RecordType).RrValue1(value).Execute()
		if err.Error() != "" {
			log.Errorf("Deletion of the RR %v %v -> %v : failed!", rr.RecordType, rr.DNSName, value)
		}
	}
	return nil
}

func (e *EfficientIPAPI) RecordList(Zone ZoneAuth) (endpoints []*endpoint.Endpoint, _ error) {
	records, _, err := e.client.DnsApi.DnsRrList(e.context).Where("zone_id=" + Zone.ID).Orderby("rr_full_name").Execute()
	if err.Error() != "" && (!records.HasSuccess() || !records.GetSuccess()) {
		log.Errorf("Failed to get RRs for zone [%s]", Zone.Name)
		return nil, err
	}

	Host := make(map[string]*endpoint.Endpoint)
	for _, rr := range records.GetData() {
		ttl, _ := strconv.Atoi(rr.GetRrTtl())

		switch rr.GetRrType() {
		case "A":
			log.Debugf("Found A Record : %s -> %s", rr.GetRrFullName(), rr.GetRrAllValue())
			if h, found := Host[rr.GetRrFullName()+":"+rr.GetRrType()]; found {
				h.Targets = append(h.Targets, rr.GetRrAllValue())
			} else {
				Host[rr.GetRrFullName()+":"+rr.GetRrType()] = endpoint.NewEndpointWithTTL(rr.GetRrFullName(), endpoint.RecordTypeA, endpoint.TTL(ttl), rr.GetRrAllValue())
			}
		case "TXT":
			log.Debugf("Found TXT Record : %s -> %s", rr.GetRrFullName(), rr.GetRrAllValue())
			tmp := endpoint.NewEndpointWithTTL(rr.GetRrFullName(), endpoint.RecordTypeTXT, endpoint.TTL(ttl), rr.GetRrAllValue())
			endpoints = append(endpoints, tmp)
		default:
			log.Debugf("Found %s Record : %s -> %s", rr.GetRrType(), rr.GetRrFullName(), rr.GetRrAllValue())
			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(rr.GetRrFullName(), rr.GetRrType(), endpoint.TTL(ttl), rr.GetRrAllValue()))
		}
	}
	for _, rr := range Host {
		endpoints = append(endpoints, rr)
	}
	return endpoints, nil
}

type EfficientIPProvider struct {
	provider.BaseProvider
	domainFilter endpoint.DomainFilter
	zoneIDFilter provider.ZoneIDFilter
	dryRun       bool
	client       EfficientipClient
	context      context.Context
}

func NewEfficientIPProvider(config EfficientIPConfig) (*EfficientIPProvider, error) {
	clientConfig := eip.NewConfiguration()
	if !config.SSlVerify {
		customTransport := http.DefaultTransport.(*http.Transport).Clone()
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		clientConfig.HTTPClient = &http.Client{Transport: customTransport}
	}

	//client := eip.NewAPIClient(clientConfig)
	ctx := context.WithValue(context.Background(), eip.ContextBasicAuth, eip.BasicAuth{
		UserName: config.Username,
		Password: config.Password,
	})
	ctx = context.WithValue(ctx, eip.ContextServerVariables, map[string]string{
		"host": 	     	 config.Host,
		"port":                  strconv.Itoa(config.Port),
	})
	client := NewEfficientipAPI(ctx, clientConfig)

	eipProvider := &EfficientIPProvider{
		domainFilter: config.DomainFilter,
		zoneIDFilter: config.ZoneIDFilter,
		dryRun:       config.DryRun,
		client:       &client,
		context:      ctx,
	}
	return eipProvider, nil
}

type ZoneAuth struct {
	Name string
	Type string
	ID   string
}

func NewZoneAuth(zone eip.DnsZoneDataData) *ZoneAuth {
	return &ZoneAuth{
		Name: zone.GetZoneName(),
		Type: zone.GetZoneType(),
		ID:   zone.GetZoneId(),
	}
}

func (p *EfficientIPProvider) Zones(_ context.Context) ([]*ZoneAuth, error) {
	var result []*ZoneAuth

	zones, err := p.client.ZonesList()

	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		if !p.domainFilter.Match(zone.Name) {
			log.Debugf("Ignore zone [%s] by domainFilter", zone.Name)
			continue
		}
		if !p.zoneIDFilter.Match(zone.ID) {
			log.Debugf("Ignore zone [%s][%s] by zoneIDFilter", zone.Name, zone.ID)
			continue
		}
		result = append(result, zone)
	}
	return result, nil
}

// Records gets the current records.
func (p *EfficientIPProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, _ error) {
	log.Debug("Get Record list from EfficientIP SOLIDserver")

	zones, err := p.Zones(ctx)
	if err != nil {
		log.Errorf("Failed to get Zone list from EfficientIP SOLIDserver")
		return nil, err
	}

	for _, zone := range zones {
		records, err := p.client.RecordList(*zone)

		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, records...)
	}

	return endpoints, nil
}

func (p *EfficientIPProvider) DeleteChanges(_ context.Context, changes *endpoint.Endpoint) error {
	if p.dryRun {
		for _, value := range changes.Targets {
			log.Infof("Would delete %s record named '%s' to '%s' for Efficientip",
				changes.RecordType,
				changes.DNSName,
				value,
			)
		}
		return nil
	}
	_ = p.client.RecordDelete(changes)
	return nil
}

func (p *EfficientIPProvider) CreateChanges(_ context.Context, changes *endpoint.Endpoint) error {
	if p.dryRun {
		for _, value := range changes.Targets {
			log.Infof("Would create %s record named '%s' to '%s' for Efficientip",
				changes.RecordType,
				changes.DNSName,
				value,
			)
		}
		return nil
	}
	_ = p.client.RecordAdd(changes)
	return nil
}

// ApplyChanges applies the given changes.
func (p *EfficientIPProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	for _, change := range changes.Delete {
		err := p.DeleteChanges(ctx, change)
		if err != nil {
			return err
		}
	}
	for _, change := range changes.UpdateOld {
		err := p.DeleteChanges(ctx, change)
		if err != nil {
			return err
		}
	}
	for _, change := range changes.UpdateNew {
		err := p.CreateChanges(ctx, change)
		if err != nil {
			return err
		}
	}
	for _, change := range changes.Create {
		err := p.CreateChanges(ctx, change)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *EfficientIPProvider) PropertyValuesEqual(name string, previous string, current string) bool {
	return p.BaseProvider.PropertyValuesEqual(name, previous, current)
}

func (p *EfficientIPProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	return endpoints
}
