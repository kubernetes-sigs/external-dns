/*
Copyright 2022 The Kubernetes Authors.

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

package yandex

import (
	"context"
	"errors"
	"strings"

	log "github.com/sirupsen/logrus"
	dnsInt "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/iamkey"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	YandexAuthorizationTypeInstanceServiceAccount = "instance-service-account"
	YandexAuthorizationTypeOAuthToken             = "iam-token"
	YandexAuthorizationTypeKey                    = "iam-key-file"

	YandexDNSRecordSetDefaultTTL = int64(300)
)

type YandexConfig struct {
	DomainFilter            endpoint.DomainFilter
	ZoneNameFilter          endpoint.DomainFilter
	ZoneIDFilter            provider.ZoneIDFilter
	DryRun                  bool
	FolderID                string
	AuthorizationType       string
	AuthorizationOAuthToken string
	AuthorizationKeyFile    string
}

type YandexProvider struct {
	provider.BaseProvider

	DomainFilter   endpoint.DomainFilter
	ZoneNameFilter endpoint.DomainFilter
	ZoneIDFilter   provider.ZoneIDFilter
	DryRun         bool
	FolderID       string

	client DNSClient
}

func NewYandexProvider(ctx context.Context, cfg *YandexConfig) (*YandexProvider, error) {
	creds, err := cfg.ResolveCredentials()
	if err != nil {
		return nil, err
	}

	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: creds,
	})
	if err != nil {
		return nil, err
	}
	if cfg.FolderID == "" {
		return nil, errors.New("empty folderId specified")
	}

	return &YandexProvider{
		DomainFilter:   cfg.DomainFilter,
		ZoneNameFilter: cfg.ZoneNameFilter,
		ZoneIDFilter:   cfg.ZoneIDFilter,
		DryRun:         cfg.DryRun,
		FolderID:       cfg.FolderID,

		client: &DNSZoneClientAdapter{sdk.DNS().DnsZone()},
	}, nil
}

func (cfg *YandexConfig) ResolveCredentials() (ycsdk.Credentials, error) {
	auth := strings.TrimSpace(cfg.AuthorizationType)

	switch auth {
	case YandexAuthorizationTypeInstanceServiceAccount:
		return ycsdk.InstanceServiceAccount(), nil
	case YandexAuthorizationTypeOAuthToken:
		return ycsdk.OAuthToken(cfg.AuthorizationOAuthToken), nil
	case YandexAuthorizationTypeKey:
		key, err := iamkey.ReadFromJSONFile(cfg.AuthorizationKeyFile)
		if err != nil {
			return nil, err
		}
		return ycsdk.ServiceAccountKey(key)
	default:
		return nil, errors.New("unsupported authorization type")
	}
}

func (p *YandexProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.zones(ctx)
	if err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0)

	for _, zone := range zones {
		records, err := p.records(ctx, zone.Zone, zone.Id)

		if err != nil {
			return nil, err
		}

		for _, record := range records {
			ep := toEndpoint(record)

			if ep == nil {
				continue
			}

			endpoints = append(endpoints, ep)
		}
	}

	return endpoints, nil
}

func (p *YandexProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	zones, err := p.zones(ctx)
	if err != nil {
		return err
	}

	zoneIDMapper := provider.ZoneIDName{}
	for _, z := range zones {
		zoneIDMapper.Add(z.Id, strings.TrimSuffix(z.Zone, "."))
	}

	batchMap := make(upsertBatchMap)
	batchMap.ApplyChanges(zoneIDMapper, changes.Delete, func(batch *upsertBatch, rs *dnsInt.RecordSet) {
		batch.Deletes = append(batch.Deletes, rs)
	})
	batchMap.ApplyChanges(zoneIDMapper, changes.Create, func(batch *upsertBatch, rs *dnsInt.RecordSet) {
		batch.Creates = append(batch.Creates, rs)
	})
	batchMap.ApplyChanges(zoneIDMapper, changes.UpdateNew, func(batch *upsertBatch, rs *dnsInt.RecordSet) {
		batch.Updates = append(batch.Updates, rs)
	})

	for _, batch := range batchMap {
		log.Infof("Would perform be batch update for zone: %s\n"+
			"Records to delete: %s\n"+
			"Records to create: %s\n"+
			"Records to update: %s\n",
			batch.ZoneName,
			toString(batch.Deletes),
			toString(batch.Creates),
			toString(batch.Updates),
		)

		if p.DryRun {
			continue
		}

		if err := p.upsertRecords(ctx, batch); err != nil {
			log.Errorf("Failed to execute upsert operation: %v", err)
		}
	}

	return nil
}

func (p *YandexProvider) zones(ctx context.Context) ([]*dnsInt.DnsZone, error) {
	log.Debugf("Retrieving Yandex DNS zones for folder: %s.", p.FolderID)

	iterator := p.client.ZoneIterator(ctx, &dnsInt.ListDnsZonesRequest{
		FolderId: p.FolderID,
	})

	zones := make([]*dnsInt.DnsZone, 0)

	for iterator.Next() {
		zone := iterator.Value()

		if !p.DomainFilter.Match(zone.Zone) || !p.ZoneIDFilter.Match(zone.Id) {
			log.Debugf("Skipping zone '%s' because of Domain And ZoneId filters", zone.Zone)
			continue
		}

		if !p.ZoneNameFilter.Match(zone.Zone) {
			log.Debugf("Skipping zone '%s' because of ZoneName filter", zone.Zone)
			continue
		}

		zones = append(zones, zone)
	}

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	log.Debugf("Found %d Yandex DNS zone(s).", len(zones))
	return zones, nil
}

func (p *YandexProvider) records(ctx context.Context, zoneName, zoneID string) ([]*dnsInt.RecordSet, error) {
	log.Debugf("Retrieving Yandex DNS records for zone '%s'.", zoneName)

	iterator := p.client.RecordSetIterator(ctx, &dnsInt.ListDnsZoneRecordSetsRequest{
		DnsZoneId: zoneID,
	})

	records := make([]*dnsInt.RecordSet, 0)

	for iterator.Next() {
		record := iterator.Value()

		if record == nil {
			log.Debugf("Skipping invalid nil record")
			continue
		}

		if !provider.SupportedRecordType(record.Type) {
			log.Debugf("Skipping record because of not supported type")
			continue
		}

		if len(p.ZoneNameFilter.Filters) > 0 && !p.DomainFilter.Match(record.Name) {
			log.Debugf("Skipping return of record %s because it was filtered out by the specified --domain-filter", record.Name)
			continue
		}

		if record.Data == nil || len(record.Data) == 0 {
			log.Debugf("Skipping return of record %s (%s) because it with empty targets", record.Name, record.Type)
			continue
		}

		records = append(records, record)
	}

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	log.Debugf("Found %d Yandex DNS records for zone '%s'.", len(records), zoneName)
	return records, nil
}

func (p *YandexProvider) upsertRecords(ctx context.Context, batch *upsertBatch) error {
	log.Infof("Perform upsert operation for zone '%s'. Deletes: %d, Updates: %d, Creates: %d",
		batch.ZoneName,
		len(batch.Deletes),
		len(batch.Updates),
		len(batch.Creates),
	)

	_, err := p.client.UpsertRecordSets(ctx,
		&dnsInt.UpsertRecordSetsRequest{
			DnsZoneId:    batch.ZoneID,
			Deletions:    batch.Deletes,
			Replacements: batch.Updates,
			Merges:       batch.Creates,
		},
	)

	if err != nil {
		log.Errorf("Failed to perform upsert operation for zone '%s'", batch.ZoneName)
		return err
	}
	return nil
}

func toEndpoint(record *dnsInt.RecordSet) *endpoint.Endpoint {
	if record == nil {
		log.Errorf("Skipping invalid record set with nil definition")
		return nil
	}

	return endpoint.NewEndpointWithTTL(
		record.GetName(),
		record.GetType(),
		endpoint.TTL(record.GetTtl()),
		record.Data...,
	)
}

func toRecordSet(ep *endpoint.Endpoint) *dnsInt.RecordSet {
	if ep == nil {
		log.Errorf("Skipping invalid endpoint with nil definition")
		return nil
	}

	recordTTL := YandexDNSRecordSetDefaultTTL
	if ep.RecordTTL.IsConfigured() {
		recordTTL = int64(ep.RecordTTL)
	}

	return &dnsInt.RecordSet{
		Name: ep.DNSName + ".",
		Type: ep.RecordType,
		Ttl:  recordTTL,
		Data: ep.Targets,
	}
}

func toString(records []*dnsInt.RecordSet) string {
	message := ""

	for _, record := range records {
		message += record.Name
		message += ","
	}

	return message
}
