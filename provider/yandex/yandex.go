package yandex

import (
	"context"
	"errors"
	"strings"

	log "github.com/sirupsen/logrus"
	dnsInt "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"
	op "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	ycsdk "github.com/yandex-cloud/go-sdk"
	"github.com/yandex-cloud/go-sdk/gen/dns"
	"github.com/yandex-cloud/go-sdk/iamkey"
	"google.golang.org/grpc"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	YandexAuthorizationTypeInstanceServiceAccount = "instance-service-account"
	YandexAuthorizationTypeOAuthToken             = "iam-token"
	YandexAuthorizationTypeKey                    = "iam-key-file"
)

type YandexConfig struct {
	DomainFilter            endpoint.DomainFilter
	ZoneNameFilter          endpoint.DomainFilter
	ZoneIdFilter            provider.ZoneIDFilter
	DryRun                  bool
	FolderId                string
	AuthorizationType       string
	AuthorizationOAuthToken string
	AuthorizationKeyFile    string
}

type YandexProvider struct {
	provider.BaseProvider

	DomainFilter   endpoint.DomainFilter
	ZoneNameFilter endpoint.DomainFilter
	ZoneIdFilter   provider.ZoneIDFilter
	DryRun         bool
	FolderId       string

	client DNSZoneClient
}

type upsertBatch struct {
	ZoneId   string
	ZoneName string
	Deletes  []*dnsInt.RecordSet
	Updates  []*dnsInt.RecordSet
	Creates  []*dnsInt.RecordSet
}

type upsertBatchMap map[string]*upsertBatch

type DNSZoneClient interface {
	DnsZoneIterator(ctx context.Context,
		req *dnsInt.ListDnsZonesRequest,
		opts ...grpc.CallOption,
	) *dns.DnsZoneIterator

	DnsZoneRecordSetsIterator(ctx context.Context,
		req *dnsInt.ListDnsZoneRecordSetsRequest,
		opts ...grpc.CallOption,
	) *dns.DnsZoneRecordSetsIterator

	UpsertRecordSets(ctx context.Context,
		in *dnsInt.UpsertRecordSetsRequest,
		opts ...grpc.CallOption,
	) (*op.Operation, error)
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
	if cfg.FolderId == "" {
		return nil, errors.New("empty folderId specified")
	}

	return &YandexProvider{
		DomainFilter:   cfg.DomainFilter,
		ZoneNameFilter: cfg.ZoneNameFilter,
		ZoneIdFilter:   cfg.ZoneIdFilter,
		DryRun:         cfg.DryRun,
		FolderId:       cfg.FolderId,

		client: sdk.DNS().DnsZone(),
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

	zoneId := provider.ZoneIDName{}
	for _, z := range zones {
		zoneId.Add(z.Id, strings.TrimSuffix(z.Zone, "."))
	}

	batchMap := make(upsertBatchMap)

	for _, change := range changes.Delete {
		zoneId, zoneName := zoneId.FindZone(change.DNSName)
		if zoneId != "" && zoneName != "" {
			batchMap.GetOrCreate(zoneId, zoneName).AddDeleted(change)
		}
	}

	for _, change := range changes.Create {
		zoneId, zoneName := zoneId.FindZone(change.DNSName)
		if zoneId != "" && zoneName != "" {
			batchMap.GetOrCreate(zoneId, zoneName).AddCreated(change)
		}
	}

	for _, change := range changes.UpdateNew {
		zoneId, zoneName := zoneId.FindZone(change.DNSName)
		if zoneId != "" && zoneName != "" {
			batchMap.GetOrCreate(zoneId, zoneName).AddUpdated(change)
		}
	}

	for _, batch := range batchMap {
		if p.DryRun {
			log.Infof("Would perform be batch update for zone: %s\n"+
				"Records to delete: %s\n"+
				"Records to create: %s\n"+
				"Records to update: %s\n",
				batch.ZoneName,
				toString(batch.Deletes),
				toString(batch.Creates),
				toString(batch.Updates),
			)
			continue
		}

		if err := p.upsertRecords(ctx, batch); err != nil {
			log.Errorf("Failed to execute upsert operation: %v", err)
		}
	}

	return nil
}

func (p *YandexProvider) zones(ctx context.Context) ([]*dnsInt.DnsZone, error) {
	log.Debugf("Retrieving Yandex DNS zones for folder: %s.", p.FolderId)

	iterator := p.client.DnsZoneIterator(ctx, &dnsInt.ListDnsZonesRequest{
		FolderId: p.FolderId,
	})

	zones := make([]*dnsInt.DnsZone, 0)

	for iterator.Next() {
		zone := iterator.Value()

		if !p.DomainFilter.Match(zone.Zone) || !p.ZoneIdFilter.Match(zone.Id) {
			log.Debugf("Skipping zone '%s' because of Domain And ZoneId filters", zone.Zone)
			continue
		}

		if !p.ZoneNameFilter.Match(zone.Zone) {
			log.Debugf("Skipping zone '%s' because of ZoneName filter", zone.Zone)
			continue
		}

		zones = append(zones, zone)
	}

	log.Debugf("Found %d Yandex DNS zone(s).", len(zones))
	return zones, nil
}

func (p *YandexProvider) records(ctx context.Context, zoneName, zoneId string) ([]*dnsInt.RecordSet, error) {
	log.Debugf("Retrieving Yandex DNS records for zone '%s'.", zoneName)

	iterator := p.client.DnsZoneRecordSetsIterator(ctx, &dnsInt.ListDnsZoneRecordSetsRequest{
		DnsZoneId: zoneId,
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
			DnsZoneId:    batch.ZoneId,
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

func (b *upsertBatch) AddDeleted(ep *endpoint.Endpoint) {
	b.Deletes = append(b.Deletes, toRecordSet(ep))
}

func (b *upsertBatch) AddCreated(ep *endpoint.Endpoint) {
	b.Creates = append(b.Creates, toRecordSet(ep))
}

func (b *upsertBatch) AddUpdated(ep *endpoint.Endpoint) {
	b.Updates = append(b.Updates, toRecordSet(ep))
}

func (m upsertBatchMap) GetOrCreate(zoneId, zoneName string) *upsertBatch {
	batch, ok := m[zoneId]

	if !ok {
		batch = &upsertBatch{
			ZoneId:   zoneId,
			ZoneName: zoneName,
			Creates:  make([]*dnsInt.RecordSet, 0),
			Deletes:  make([]*dnsInt.RecordSet, 0),
			Updates:  make([]*dnsInt.RecordSet, 0),
		}
		m[zoneId] = batch
	}

	return batch
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

	return &dnsInt.RecordSet{
		Name: ep.DNSName + ".",
		Type: ep.RecordType,
		Ttl:  int64(ep.RecordTTL),
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
