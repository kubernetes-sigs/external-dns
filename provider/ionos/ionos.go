package ionos

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	sdk "github.com/ionos-developer/dns-sdk-go"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// IonosProvider implements the DNS provider for IONOS DNS.
type IonosProvider struct {
	provider.BaseProvider
	Client IonosDnsService

	DomainFilter endpoint.DomainFilter
	DryRun       bool
}

type IonosDnsService interface {
	GetZones(ctx context.Context) ([]sdk.Zone, error)
	GetZone(ctx context.Context, zoneId string) (*sdk.CustomerZone, error)
	CreateRecords(ctx context.Context, zoneId string, records []sdk.Record) error
	DeleteRecord(ctx context.Context, zoneId string, recordId string) error
}

type IonosDnsClient struct {
	client *sdk.APIClient
}

func (c IonosDnsClient) GetZones(ctx context.Context) ([]sdk.Zone, error) {
	zones, _, err := c.client.ZonesApi.GetZones(ctx).Execute()
	if err != nil {
		return nil, err
	}

	return zones, err
}

func (c IonosDnsClient) GetZone(ctx context.Context, zoneId string) (*sdk.CustomerZone, error) {
	zoneInfo, _, err := c.client.ZonesApi.GetZone(ctx, zoneId).Execute()
	return zoneInfo, err
}

func (c IonosDnsClient) CreateRecords(ctx context.Context, zoneId string, records []sdk.Record) error {
	_, _, err := c.client.RecordsApi.CreateRecords(ctx, zoneId).Record(records).Execute()
	return err
}

func (c IonosDnsClient) DeleteRecord(ctx context.Context, zoneId string, recordId string) error {
	_, err := c.client.RecordsApi.DeleteRecord(ctx, zoneId, recordId).Execute()
	return err
}

// NewIonosProvider creates a new IONOS DNS provider.
func NewIonosProvider(domainFilter endpoint.DomainFilter, dryRun bool) (*IonosProvider, error) {
	const apiKeyEnvVar = "IONOS_API_KEY"
	apiKey := os.Getenv(apiKeyEnvVar)
	if len(apiKey) == 0 {
		return nil, fmt.Errorf("failed to initialize ionos provider: %s not present", apiKeyEnvVar)
	}

	configuration := sdk.NewConfiguration()
	if url, ok := os.LookupEnv("IONOS_API_URL"); ok {
		configuration.Servers[0].URL = url
	}
	configuration.AddDefaultHeader(getEnv("IONOS_AUTH_HEADER", "X-API-Key"), apiKey)
	configuration.UserAgent = fmt.Sprintf(
		"external-dns os %s arch %s",
		runtime.GOOS, runtime.GOARCH)
	if os.Getenv("IONOS_DEBUG") != "" {
		configuration.Debug = true
	}

	client := sdk.NewAPIClient(configuration)

	provider := &IonosProvider{
		Client:       IonosDnsClient{client: client},
		DomainFilter: domainFilter,
		DryRun:       dryRun,
	}

	return provider, nil
}

// Records returns the list of resource records in all zones.
func (p *IonosProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.getZones(ctx)
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint

	for zoneId := range zones {
		zoneInfo, err := p.Client.GetZone(ctx, zoneId)
		if err != nil {
			log.Warnf("Failed to fetch zoneId %v: %v", zoneId, err)
			continue
		}

		recordSets := map[string]*endpoint.Endpoint{}
		for _, r := range zoneInfo.Records {
			key := *r.Name + "/" + getType(r) + "/" + strconv.Itoa(int(*r.Ttl))
			if rrset, ok := recordSets[key]; ok {
				rrset.Targets = append(rrset.Targets, *r.Content)
			} else {
				recordSets[key] = recordToEndpoint(r)
			}
		}

		for _, endpoint := range recordSets {
			endpoints = append(endpoints, endpoint)
		}
	}

	log.Debugf("Records() result")
	for _, e := range endpoints {
		log.Debugf("   %v %v", e, e.Labels)
	}
	return endpoints, nil
}

// ApplyChanges applies a given set of changes.
func (p *IonosProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	zones, err := p.getZones(ctx)
	if err != nil {
		return err
	}

	toCreate := make([]*endpoint.Endpoint, len(changes.Create))
	copy(toCreate, changes.Create)

	toDelete := make([]*endpoint.Endpoint, len(changes.Delete))
	copy(toDelete, changes.Delete)

	for i, updateOldEndpoint := range changes.UpdateOld {
		if !sameEndpoints(*updateOldEndpoint, *changes.UpdateNew[i]) {
			toDelete = append(toDelete, updateOldEndpoint)
			toCreate = append(toCreate, changes.UpdateNew[i])
		}
	}

	zonesToDeleteFrom := p.fetchZonesToDeleteFrom(ctx, toDelete, zones)

	for _, endpoint := range toDelete {
		zoneId := getHostZoneID(endpoint.DNSName, zones)
		if zoneId == "" {
			log.Warnf("No zone to delete %v from", endpoint)
			continue
		}

		if zone, ok := zonesToDeleteFrom[zoneId]; ok {
			p.deleteEndpoint(ctx, endpoint, zone)
		} else {
			log.Warnf("No zone to delete %v from", endpoint)
		}
	}

	for _, endpoint := range toCreate {
		p.createEndpoint(ctx, endpoint, zones)
	}

	return nil
}

// fetchZonesToDeleteFrom fetches all the zones that will be performed deletions upon.
func (p *IonosProvider) fetchZonesToDeleteFrom(ctx context.Context, toDelete []*endpoint.Endpoint, zones map[string]string) map[string]*sdk.CustomerZone {
	zonesIdsToDeleteFrom := map[string]bool{}
	for _, endpoint := range toDelete {
		zoneId := getHostZoneID(endpoint.DNSName, zones)
		if zoneId != "" {
			zonesIdsToDeleteFrom[zoneId] = true
		}
	}

	zonesToDeleteFrom := map[string]*sdk.CustomerZone{}
	for zoneId := range zonesIdsToDeleteFrom {
		zone, err := p.Client.GetZone(ctx, zoneId)
		if err == nil {
			zonesToDeleteFrom[zoneId] = zone
		}
	}

	return zonesToDeleteFrom
}

// deleteEndpoint deletes all resource records for the endpoint through the IONOS DNS API.
func (p *IonosProvider) deleteEndpoint(ctx context.Context, e *endpoint.Endpoint, zone *sdk.CustomerZone) {
	log.Infof("Delete endpoint %v", e)
	if p.DryRun {
		return
	}

	for _, target := range e.Targets {
		recordId := ""
		for _, record := range zone.Records {
			if *record.Name == e.DNSName && getType(record) == e.RecordType && *record.Content == target {
				recordId = *record.Id
				break
			}
		}

		if recordId == "" {
			log.Warnf("Record %v %v %v not found in zone", e.DNSName, e.RecordType, target)
			continue
		}

		if p.Client.DeleteRecord(ctx, *zone.Id, recordId) != nil {
			log.Warnf("Failed to delete record %v %v %v", e.DNSName, e.RecordType, target)
		}
	}
}

// createEndpoint creates the record set for the endpoint using the IONOS DNS API.
func (p *IonosProvider) createEndpoint(ctx context.Context, e *endpoint.Endpoint, zones map[string]string) {
	log.Infof("Create endpoint %v", e)
	if p.DryRun {
		return
	}

	zoneId := getHostZoneID(e.DNSName, zones)
	if zoneId == "" {
		log.Warnf("No zone to create %v into", e)
		return
	}

	records := endpointToRecords(e)
	if p.Client.CreateRecords(ctx, zoneId, records) != nil {
		log.Warnf("Failed to create record for %v", e)
	}
}

// endpointToRecords converts an endpoint to a slice of records.
func endpointToRecords(endpoint *endpoint.Endpoint) []sdk.Record {
	records := make([]sdk.Record, 0)

	for _, target := range endpoint.Targets {
		record := sdk.NewRecord()

		record.SetName(endpoint.DNSName)
		record.SetType(sdk.RecordTypes(endpoint.RecordType))
		record.SetContent(target)

		ttl := int32(endpoint.RecordTTL)
		if ttl != 0 {
			record.SetTtl(ttl)
		}

		records = append(records, *record)
	}

	return records
}

// recordToEndpoint converts a record to an endpoint.
func recordToEndpoint(r sdk.RecordResponse) *endpoint.Endpoint {
	return endpoint.NewEndpointWithTTL(*r.Name, getType(r), endpoint.TTL(*r.Ttl), *r.Content)
}

// getEnv returns the value of an environment variable, returning the fallback if the variable is not set.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// getZones returns a ZoneID -> ZoneName mapping for zones that match domain filter.
func (p IonosProvider) getZones(ctx context.Context) (map[string]string, error) {
	zones, err := p.Client.GetZones(ctx)
	if err != nil {
		return nil, err
	}

	result := map[string]string{}

	for _, zone := range zones {
		if p.DomainFilter.Match(*zone.Name) {
			result[*zone.Id] = *zone.Name
		}
	}

	return result, nil
}

// getHostZoneID finds the best suitable DNS zone for the hostname.
func getHostZoneID(hostname string, zones map[string]string) string {
	longestZoneLength := 0
	resultID := ""

	for zoneID, zoneName := range zones {
		if !strings.HasSuffix(hostname, zoneName) {
			continue
		}
		ln := len(zoneName)
		if ln > longestZoneLength {
			resultID = zoneID
			longestZoneLength = ln
		}
	}

	return resultID
}

// getType returns the record type as string.
func getType(record sdk.RecordResponse) string {
	return string(*record.Type)
}

// sameEndpoints returns if the two endpoints have the same values.
func sameEndpoints(a endpoint.Endpoint, b endpoint.Endpoint) bool {
	return a.DNSName == b.DNSName && a.RecordType == b.RecordType && a.RecordTTL == b.RecordTTL && a.Targets.Same(b.Targets)
}
