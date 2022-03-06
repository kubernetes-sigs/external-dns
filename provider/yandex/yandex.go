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
	"fmt"
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
		log.WithFields(log.Fields{
			"deletes": toString(batch.Deletes),
			"creates": toString(batch.Creates),
			"updates": toString(batch.Updates),
			"zone":    batch.ZoneName,
			"zoneID":  batch.ZoneID,
		}).Info("Would execute upsert operation")

		if p.DryRun {
			continue
		}

		if err := p.upsertRecords(ctx, batch); err != nil {
			log.WithFields(log.Fields{
				"error":  err,
				"zone":   batch.ZoneName,
				"zoneID": batch.ZoneID,
			}).Error("Failed to execute upsert operation")
		}
	}

	return nil
}

func (p *YandexProvider) zones(ctx context.Context) ([]*dnsInt.DnsZone, error) {
	log.WithFields(log.Fields{"folder": p.FolderID}).Debug("Retrieving Yandex DNS zones for folder")

	iterator := p.client.ZoneIterator(ctx, &dnsInt.ListDnsZonesRequest{
		FolderId: p.FolderID,
	})

	zones := make([]*dnsInt.DnsZone, 0)

	for iterator.Next() {
		zone := iterator.Value()

		if !p.DomainFilter.Match(zone.Zone) || !p.ZoneIDFilter.Match(zone.Id) {
			log.WithFields(log.Fields{
				"zone":   zone.Zone,
				"zoneID": zone.Id,
			}).Debug("Skipping zone because of Domain and ZoneId filters")
			continue
		}

		if !p.ZoneNameFilter.Match(zone.Zone) {
			log.WithFields(log.Fields{
				"zone":   zone.Zone,
				"zoneID": zone.Id,
			}).Debug("Skipping zone because of ZoneName filter")
			continue
		}

		zones = append(zones, zone)
	}

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"folder": p.FolderID,
		"zones":  len(zones),
	}).Debug("Found Yandex DNS zone(s) for folder")
	return zones, nil
}

func (p *YandexProvider) records(ctx context.Context, zoneName, zoneID string) ([]*dnsInt.RecordSet, error) {
	log.WithFields(log.Fields{
		"zone":   zoneName,
		"zoneID": zoneID,
	}).Debug("Retrieving Yandex DNS records for zone")

	iterator := p.client.RecordSetIterator(ctx, &dnsInt.ListDnsZoneRecordSetsRequest{
		DnsZoneId: zoneID,
	})

	records := make([]*dnsInt.RecordSet, 0)

	for iterator.Next() {
		record := iterator.Value()

		if record == nil {
			log.Debug("Skipping invalid nil record")
			continue
		}

		if !provider.SupportedRecordType(record.Type) {
			log.WithFields(log.Fields{
				"record":     record.Name,
				"recordType": record.Type,
			}).Debug("Skipping record because of not supported type")
			continue
		}

		if len(p.ZoneNameFilter.Filters) > 0 && !p.DomainFilter.Match(record.Name) {
			log.WithFields(log.Fields{
				"record":     record.Name,
				"recordType": record.Type,
			}).Debug("Skipping return of record because it was filtered out by the specified --domain-filter")
			continue
		}

		if record.Data == nil || len(record.Data) == 0 {
			log.WithFields(log.Fields{
				"record":     record.Name,
				"recordType": record.Type,
			}).Debug("Skipping return of record because it with empty targets")
			continue
		}

		records = append(records, record)
	}

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{
		"records": len(records),
		"zone":    zoneName,
		"zoneID":  zoneID,
	}).Debug("Found Yandex DNS record(s) for zone")
	return records, nil
}

func (p *YandexProvider) upsertRecords(ctx context.Context, batch *upsertBatch) error {
	log.WithFields(log.Fields{
		"zone":         batch.ZoneName,
		"zoneID":       batch.ZoneID,
		"createsCount": len(batch.Creates),
		"updatesCount": len(batch.Updates),
		"deletesCount": len(batch.Deletes),
	}).Info("Perform upsert operation for zone")

	_, err := p.client.UpsertRecordSets(ctx,
		&dnsInt.UpsertRecordSetsRequest{
			DnsZoneId:    batch.ZoneID,
			Deletions:    batch.Deletes,
			Replacements: batch.Updates,
			Merges:       batch.Creates,
		},
	)

	if err != nil {
		log.WithFields(log.Fields{
			"zone":   batch.ZoneName,
			"zoneID": batch.ZoneID,
		}).Error("Failed to perform upsert operation for zone")
		return err
	}
	return nil
}

func toEndpoint(record *dnsInt.RecordSet) *endpoint.Endpoint {
	if record == nil {
		log.Error("Skipping invalid record set with nil definition")
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
		log.Error("Skipping invalid endpoint with nil definition")
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
	var message strings.Builder

	for _, record := range records {
		message.WriteString(fmt.Sprintf("%s (%s); ", record.Name, record.Type))
	}

	return message.String()
}
