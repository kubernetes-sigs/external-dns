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

func NewYandexProvider(ctx context.Context, cfg *YandexConfig) (*YandexProvider, error) {
	creds, err := cfg.credentials()
	if err != nil {
		return nil, err
	}

	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: creds,
	})

	if err != nil {
		return nil, err
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
			ep := p.convert(record)

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

	for _, zone := range zones {
		// todo: map changes per zone to batch
		//       create -> merges
		//       delete -> deletes
		//       update -> replacements
		log.Debugf(zone.Name)
	}

	return nil
}

func (p *YandexProvider) zones(ctx context.Context) ([]*dnsInt.DnsZone, error) {
	log.Debugf("Retrieving Yandex DNS zones for folder: %s.", p.FolderId)

	iterator := p.client.DnsZoneIterator(ctx, &dnsInt.ListDnsZonesRequest{
		FolderId: p.FolderId,
	})
	zones, err := iterator.TakeAll()
	if err != nil {
		return nil, err
	}

	log.Debugf("Found %d Yandex DNS zone(s).", len(zones))
	return zones, nil
}

func (p *YandexProvider) records(ctx context.Context, zoneName, zoneId string) ([]*dnsInt.RecordSet, error) {
	log.Debugf("Retrieving Yandex DNS records for zone '%s'.", zoneName)

	iterator := p.client.DnsZoneRecordSetsIterator(ctx, &dnsInt.ListDnsZoneRecordSetsRequest{
		DnsZoneId: zoneId,
	})
	records, err := iterator.TakeAll()
	if err != nil {
		return nil, err
	}

	log.Debugf("Found %d Yandex DNS records for zone '%s'.", len(records), zoneName)
	return records, nil
}

func (p *YandexProvider) convert(record *dnsInt.RecordSet) *endpoint.Endpoint {
	if record == nil {
		log.Errorf("Skipping invalid record set with nil definition")
		return nil
	}

	recordType := record.GetType()
	recordName := record.GetName()

	if !provider.SupportedRecordType(recordType) {
		return nil
	}

	if len(p.ZoneNameFilter.Filters) > 0 && !p.DomainFilter.Match(recordName) {
		log.Debugf("Skipping return of record %s because it was filtered out by the specified --domain-filter", recordName)
		return nil
	}

	if record.Data == nil || len(record.Data) == 0 {
		log.Debugf("Skipping return of record %s because it with empty targets", recordName)
		return nil
	}

	return endpoint.NewEndpointWithTTL(
		recordName,
		recordType,
		endpoint.TTL(record.GetTtl()),
		record.Data...,
	)
}
