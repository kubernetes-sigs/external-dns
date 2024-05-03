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

package arvancloud

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/arvancloud/api"
	"sigs.k8s.io/external-dns/provider/arvancloud/dto"
)

type ArvanAdapter interface {
	GetDomains(ctx context.Context, perPage ...int) ([]dto.Zone, error)
	GetDnsRecords(ctx context.Context, zone string, perPage ...int) ([]dto.DnsRecord, error)
	CreateDnsRecord(ctx context.Context, zone string, record dto.DnsRecord) (dto.DnsRecord, error)
	UpdateDnsRecord(ctx context.Context, zone string, record dto.DnsRecord) (dto.DnsRecord, error)
	DeleteDnsRecord(ctx context.Context, zone, recordId string) error
}

type Provider struct {
	provider.BaseProvider
	domainFilter endpoint.DomainFilter
	client       ArvanAdapter
	zoneFilter   provider.ZoneIDFilter
	options      providerOptions
}

var _ provider.Provider = (*Provider)(nil)

var recordTypeProxyNotSupported = map[string]bool{
	"AAAA": true,
	"NS":   true,
	"MX":   true,
	"SRV":  true,
	"TXT":  true,
	"PTR":  true,
	"SPF":  true,
	"CAA":  true,
	"TLSA": true,
}

func NewArvanCloudProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, opts ...Option) (*Provider, error) {
	options := defaultOptions()
	options.dryRun = dryRun

	for _, o := range opts {
		o.apply(&options)
	}

	if err := optionValidation(options); err != nil {
		return nil, err
	}

	token, err := getToken()
	if err != nil {
		return nil, err
	}

	client, err := api.NewClientApi(token, options.apiVersion)
	if err != nil {
		return nil, err
	}

	arvanProvider := &Provider{
		client:       client,
		domainFilter: domainFilter,
		zoneFilter:   zoneIDFilter,
		options:      options,
	}

	return arvanProvider, nil
}

func (p *Provider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.client.GetDomains(ctx, p.options.domainPerPage)
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint
	for _, zone := range zones {
		if !p.domainFilter.Match(zone.Name) {
			log.Debugf("zone %s not in domain filter", zone.Name)
			continue
		}
		records, err := p.client.GetDnsRecords(ctx, zone.Name, 1)
		if err != nil {
			return nil, err
		}

		ge, err := groupByNameAndType(records)
		if err != nil {
			return nil, err
		}

		endpoints = append(endpoints, ge...)
	}

	return endpoints, nil
}

func (p *Provider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	var adjustedEndpoints []*endpoint.Endpoint
	for _, e := range endpoints {
		proxy := shouldBeProxy(e, p.options.enableCloudProxy)
		e.SetProviderSpecificProperty(providerSpecificCloudProxyKey, strconv.FormatBool(proxy))

		adjustedEndpoints = append(adjustedEndpoints, e)
	}

	return adjustedEndpoints, nil
}

func (p *Provider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if changes == nil {
		return nil
	}

	changeSet := make([]dto.DnsChange, 0, len(changes.Create)+len(changes.UpdateOld)+len(changes.UpdateNew)+len(changes.Delete))

	for _, ch := range changes.Create {
		changeSet = append(changeSet, p.applyDnsChange(dto.CreateDns, ch))
	}
	for _, desired := range changes.UpdateNew {
		changeSet = append(changeSet, p.applyDnsChange(dto.UpdateDns, desired))
	}
	for _, ch := range changes.Delete {
		changeSet = append(changeSet, p.applyDnsChange(dto.DeleteDns, ch))
	}

	return p.submitChanges(ctx, changeSet)
}

func (p *Provider) applyDnsChange(action dto.DnsChangeAction, endpointData *endpoint.Endpoint) dto.DnsChange {
	ttl := int64(_defaultDnsTtl)
	if endpointData.RecordTTL != 0 {
		ttl = int64(endpointData.RecordTTL)
	}
	return dto.DnsChange{
		Action: action,
		Record: dto.DnsRecord{
			Type:     dto.DnsType(endpointData.RecordType),
			Name:     endpointData.DNSName,
			TTL:      ttl,
			Cloud:    shouldBeProxy(endpointData, p.options.enableCloudProxy),
			Contents: endpointData.Targets,
		},
	}
}

func (p *Provider) submitChanges(ctx context.Context, changeSet []dto.DnsChange) error {
	if len(changeSet) == 0 {
		return nil
	}

	zones, err := p.client.GetDomains(ctx, p.options.domainPerPage)
	if err != nil {
		return err
	}

	cbz := changesByZone(zones, changeSet)

	for zoneName, changes := range cbz {
		records, err := p.client.GetDnsRecords(ctx, zoneName, p.options.dnsPerPage)
		if err != nil {
			return err
		}

		for _, change := range changes {
			recordID := getRecordID(records, change.Record)

			var actionType dto.DnsChangeAction
			if change.Action == dto.CreateDns && recordID == "" {
				actionType = dto.CreateDns
			} else if change.Action == dto.CreateDns && recordID != "" {
				actionType = dto.UpdateDns
			} else {
				actionType = change.Action
			}

			logFields := log.Fields{
				"record": change.Record.Name,
				"type":   change.Record.Type,
				"ttl":    change.Record.TTL,
				"action": actionType,
				"zone":   zoneName,
			}

			if p.options.dryRun {
				log.WithFields(logFields).Info("Changing record.")
				continue
			}

			switch actionType {
			case dto.CreateDns:
				_, err := p.client.CreateDnsRecord(ctx, zoneName, change.Record)
				if err != nil {
					log.WithFields(logFields).Errorf("failed to create record: %v", err)
				}
			case dto.UpdateDns:
				if recordID == "" {
					log.WithFields(logFields).Errorf("failed to find previous record: %v", change.Record)
					continue
				}
				change.Record.ID = recordID
				_, err := p.client.UpdateDnsRecord(ctx, zoneName, change.Record)
				if err != nil {
					log.WithFields(logFields).Errorf("failed to update record: %v", err)
				}
			case dto.DeleteDns:
				if recordID == "" {
					log.WithFields(logFields).Errorf("failed to find previous record: %v", change.Record)
					continue
				}
				err := p.client.DeleteDnsRecord(ctx, zoneName, recordID)
				if err != nil {
					log.WithFields(logFields).Errorf("failed to delete record: %v", err)
				}
			}

			log.WithFields(logFields).Info("Changing record.")
		}
	}

	return nil
}

func changesByZone(zones []dto.Zone, changeSet []dto.DnsChange) map[string][]dto.DnsChange {
	changes := make(map[string][]dto.DnsChange)
	zoneNameIDMapper := provider.ZoneIDName{}

	for _, z := range zones {
		zoneNameIDMapper.Add(z.ID, z.Name)
		changes[z.Name] = []dto.DnsChange{}
	}

	for _, c := range changeSet {
		_, zoneName := zoneNameIDMapper.FindZone(c.Record.Name)
		if zoneName == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", c.Record.Name)
			continue
		}
		c.Record.Zone = zoneName
		c.Record.Name = strings.TrimSuffix(c.Record.Name, "."+zoneName)
		changes[zoneName] = append(changes[zoneName], c)
	}

	return changes
}

func getRecordID(records []dto.DnsRecord, record dto.DnsRecord) string {
	for _, zoneRecord := range records {
		if zoneRecord.Name == record.Name && zoneRecord.Type == record.Type {
			return zoneRecord.ID
		}
	}
	return ""
}

func groupByNameAndType(records []dto.DnsRecord) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	groups := map[string][]dto.DnsRecord{}
	for _, r := range records {
		recordType := string(r.Type)
		if !provider.SupportedRecordType(recordType) {
			continue
		}

		groupBy := r.Name + "-" + recordType
		if _, ok := groups[groupBy]; !ok {
			groups[groupBy] = []dto.DnsRecord{}
		}

		groups[groupBy] = append(groups[groupBy], r)
	}

	for _, g := range groups {
		var targets []string

		for _, r := range g {
			targets = append(targets, r.Contents...)
		}

		var recordName string
		if g[0].Name == "@" {
			recordName = strings.TrimPrefix(g[0].Zone, ".")
		} else {
			recordName = strings.TrimPrefix(fmt.Sprintf("%s.%s", g[0].Name, g[0].Zone), ".")
		}

		ep := endpoint.NewEndpointWithTTL(
			recordName,
			string(g[0].Type),
			endpoint.TTL(g[0].TTL),
			targets...,
		)
		endpoints = append(endpoints, ep)
	}

	return endpoints, nil
}

func shouldBeProxy(endpoint *endpoint.Endpoint, proxyByDefault bool) bool {
	proxy := proxyByDefault
	if endpoint == nil {
		return proxy
	}

	for _, v := range endpoint.ProviderSpecific {
		if v.Name == providerSpecificCloudProxyKey {
			b, err := strconv.ParseBool(v.Value)
			if err != nil {
				log.Errorf("Failed to parse annotation [%s]: %v", providerSpecificCloudProxyKey, err)
			} else {
				proxy = b
			}
			break
		}
	}

	if recordTypeProxyNotSupported[endpoint.RecordType] {
		proxy = false
	}

	return proxy
}

func defaultOptions() (options providerOptions) {
	options.domainPerPage = 15
	options.dnsPerPage = 300
	options.enableCloudProxy = true
	options.apiVersion = "4.0"

	return options
}

func optionValidation(option providerOptions) error {
	if match, _ := regexp.MatchString("[0-9]+.[0-9]+", option.apiVersion); !match {
		return dto.NewApiVersionError()
	}

	return nil
}

func getToken() (string, error) {
	if os.Getenv(_apiTokenEnv) == "" {
		return "", dto.NewApiTokenRequireError()
	}

	var token string
	token = os.Getenv(_apiTokenEnv)
	if strings.HasPrefix(token, "file:") {
		tokenBytes, err := os.ReadFile(strings.TrimPrefix(token, "file:"))
		if err != nil {
			return "", dto.NewApiTokenFromFileError(err)
		}

		token = string(tokenBytes)
	}

	return token, nil
}
