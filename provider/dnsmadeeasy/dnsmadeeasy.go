package dnsmadeeasy

import (
	"context"
	"fmt"
	dme "github.com/DNSMadeEasy/dme-go-client/client"
	"github.com/DNSMadeEasy/dme-go-client/models"
	log "github.com/sirupsen/logrus"
	"os"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"strconv"
	"strings"
)

const dnsmadeeasyDefaultTTL = 3600 // Default TTL of 1 hour if not set (dnsmadeeasy's default)

// dnsmadeeasyZoneServiceInterface is an interface that contains all necessary zone services from dnsmadeeasy
type dmeZoneServiceInterface interface {
	GetZoneByName(ctx context.Context, zoneName string) (response *Zone, _ error)
	ListZones(ctx context.Context, page int) (response *dmeListZoneResponse, _ error)
	ListRecords(ctx context.Context, zoneID string, page int) (response *dmeListRecordResponse, _ error)
	CreateRecord(ctx context.Context, zoneID string, recordAttributes Record) (*Record, error)
	DeleteRecord(ctx context.Context, zoneID string, recordID string) error
	UpdateRecord(ctx context.Context, zoneID string, recordID string, recordAttributes Record) (*Record, error)
}

type dmeService struct {
	client dme.Client
}

type dmePaginatedResponse struct {
	Page         int
	TotalPages   int
	TotalRecords int
}

type Zone struct {
	models.DomainAttribute
	Id string
}

type Record struct {
	models.ManagedDNSRecordActions
	Id   string
	Zone *Zone
}

// Wrap each of the response types with a struct
type dmeListZoneResponse struct {
	dmePaginatedResponse
	Zones []Zone
}

type dmeListRecordResponse struct {
	dmePaginatedResponse
	Records []Record
}

func (z dmeService) ListZones(_ context.Context, page int) (response *dmeListZoneResponse, _ error) {
	response = &dmeListZoneResponse{
		Zones: []Zone{},
	}

	resp, err := z.client.GetbyId(fmt.Sprintf("dns/managed?page=%d", page))
	if err != nil {
		return nil, err
	}
	pageCount, err := strconv.Atoi(StripQuotes(resp.S("totalPages").String()))
	if err != nil {
		return nil, err
	}
	response.TotalPages = pageCount

	recordCount, err := strconv.Atoi(StripQuotes(resp.S("totalRecords").String()))
	if err != nil {
		return nil, err
	}
	response.TotalRecords = recordCount
	response.Page = page

	count, _ := resp.ArrayCount("data")
	for i := 0; i < count; i++ {
		tempItem, _ := resp.ArrayElement(i, "data")
		zoneName := StripQuotes(tempItem.S("name").String())
		zoneId := StripQuotes(tempItem.S("id").String())

		domainAttr := &Zone{}
		domainAttr.Name = zoneName
		domainAttr.FolderID = StripQuotes(tempItem.S("folderId").String())
		domainAttr.GtdEnabled = StripQuotes(tempItem.S("gtdEnabled").String())
		domainAttr.Created = StripQuotes(tempItem.S("created").String())
		domainAttr.Updated = StripQuotes(tempItem.S("updated").String())
		domainAttr.SOAID = StripQuotes(tempItem.S("soaId").String())
		domainAttr.TemplateID = StripQuotes(tempItem.S("templateId").String())
		domainAttr.TransferAClID = StripQuotes(tempItem.S("transferAclId").String())
		domainAttr.VanityID = StripQuotes(tempItem.S("vanityId").String())
		domainAttr.Id = zoneId
		response.Zones = append(response.Zones, *domainAttr)
	}

	return response, nil
}

func (z dmeService) GetZoneByName(ctx context.Context, zoneName string) (response *Zone, _ error) {
	page := 1
	for {
		tmpZones, err := z.ListZones(ctx, page)
		for _, zone := range tmpZones.Zones {
			if zone.Name == zoneName {
				return &zone, nil
			}
		}

		page++
		if page > tmpZones.Page {
			break
		}

		tmpZones, err = z.ListZones(ctx, page)
		if err != nil {
			return nil, err
		}
	}
	return nil, fmt.Errorf("No Zone found with name '%s'", zoneName)
}

func (z dmeService) ListRecords(_ context.Context, zoneID string, page int) (response *dmeListRecordResponse, _ error) {
	resp, err := z.client.GetbyId(fmt.Sprintf("dns/managed/%s/records?page=%d", zoneID, page))
	if err != nil {
		return nil, err
	}
	response = &dmeListRecordResponse{
		Records: []Record{},
	}
	pageCount, err := strconv.Atoi(StripQuotes(resp.S("totalPages").String()))
	if err != nil {
		return nil, err
	}
	response.TotalPages = pageCount

	recordCount, err := strconv.Atoi(StripQuotes(resp.S("totalRecords").String()))
	if err != nil {
		return nil, err
	}
	response.TotalRecords = recordCount
	response.Page = page

	count, _ := resp.ArrayCount("data")
	for i := 0; i < count; i++ {
		tempItem, _ := resp.ArrayElement(i, "data")
		recordName := StripQuotes(tempItem.S("name").String())
		recordId := StripQuotes(tempItem.S("id").String())

		recordAttr := &Record{}
		recordAttr.Name = recordName
		recordAttr.Id = recordId
		recordAttr.Value = StripQuotes(tempItem.S("value").String())
		recordAttr.Type = StripQuotes(tempItem.S("type").String())
		recordAttr.Ttl = StripQuotes(tempItem.S("ttl").String())
		response.Records = append(response.Records, *recordAttr)
	}

	return response, nil
}

func (z dmeService) CreateRecord(_ context.Context, zoneID string, recordAttributes Record) (response *Record, _ error) {
	response = &Record{
		Zone: &Zone{
			Id: zoneID,
		},
	}
	resp, err := z.client.Save(&recordAttributes, fmt.Sprintf("dns/managed/%s/records/", zoneID))
	if err != nil {
		return nil, err
	}

	response.Name = StripQuotes(resp.S("name").String())
	response.Id = StripQuotes(resp.S("id").String())
	response.Value = StripQuotes(resp.S("value").String())
	response.Type = StripQuotes(resp.S("type").String())
	response.Ttl = StripQuotes(resp.S("ttl").String())

	return response, nil
}

func (z dmeService) DeleteRecord(_ context.Context, zoneID string, recordID string) error {
	return z.client.Delete(fmt.Sprintf("dns/managed/%s/records/%s", zoneID, recordID))
}

func (z dmeService) UpdateRecord(_ context.Context, zoneID string, recordID string, recordAttributes Record) (response *Record, _ error) {
	response = &Record{
		Zone: &Zone{
			Id: zoneID,
		},
	}
	resp, err := z.client.Update(&recordAttributes, fmt.Sprintf("dns/managed/%s/records/%s", zoneID, recordID))
	if err != nil {
		return nil, err
	}

	response.Name = StripQuotes(resp.S("name").String())
	response.Id = StripQuotes(resp.S("id").String())
	response.Value = StripQuotes(resp.S("value").String())
	response.Type = StripQuotes(resp.S("type").String())
	response.Ttl = StripQuotes(resp.S("ttl").String())

	return response, nil
}

type dmeProvider struct {
	provider.BaseProvider
	service      dmeZoneServiceInterface
	apiKey       string
	secretKey    string
	domainFilter endpoint.DomainFilter
	zoneIDFilter provider.ZoneIDFilter
	dryRun       bool
}

type dmeChange struct {
	Action            string
	ResourceRecordSet Record
}

const (
	dmeCreate = "CREATE"
	dmeDelete = "DELETE"
	dmeUpdate = "UPDATE"
)

// NewDmeProvider initializes a new dnsmadeeasy based provider
func NewDmeProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool) (provider.Provider, error) {
	apiKey := os.Getenv("dme_apikey")
	if len(apiKey) == 0 {
		return nil, fmt.Errorf("no dme_apikey provided")
	}
	secretKey := os.Getenv("dme_secretkey")
	if len(secretKey) == 0 {
		return nil, fmt.Errorf("no dme_secretkey provided")
	}
	insecure := false
	if os.Getenv("dme_insecure") == "true" {
		insecure = true
	}

	client := dme.GetClient(apiKey, secretKey, dme.Insecure(insecure))

	provider := &dmeProvider{
		service:      dmeService{client: *client},
		apiKey:       apiKey,
		secretKey:    secretKey,
		domainFilter: domainFilter,
		zoneIDFilter: zoneIDFilter,
		dryRun:       dryRun,
	}

	return provider, nil
}

// Returns a list of filtered Zones
func (p *dmeProvider) Zones(ctx context.Context) (zones map[string]Zone, _ error) {
	page := 1
	zones = make(map[string]Zone)

	for {
		tmpZones, err := p.service.ListZones(ctx, page)
		for _, zone := range tmpZones.Zones {
			if !p.domainFilter.Match(zone.Name) {
				continue
			}

			if !p.zoneIDFilter.Match(zone.Id) {
				continue
			}
			zones[zone.Id] = zone
		}

		page++
		if page > tmpZones.Page {
			break
		}

		tmpZones, err = p.service.ListZones(ctx, page)
		if err != nil {
			return nil, err
		}
	}

	return zones, nil
}

// Records returns a list of endpoints in a given zone
func (p *dmeProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, _ error) {

	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {

		page := 1

		for {
			records, err := p.service.ListRecords(ctx, zone.Id, page)
			if err != nil {
				return nil, err
			}

			for _, record := range records.Records {
				switch record.Type {
				case "A", "CNAME", "TXT":
					break
				default:
					continue
				}
				// Apex records have an empty string for their name.
				// Consider this when creating the endpoint dnsName
				dnsName := fmt.Sprintf("%s.%s", record.Name, zone.Name)
				if record.Name == "" {
					dnsName = zone.Name
				}

				rTtl, err := strconv.Atoi(record.Ttl)
				if err != nil {
					endpoints = append(endpoints, endpoint.NewEndpointWithTTL(
						dnsName,
						record.Type,
						endpoint.TTL(rTtl),
						record.Value,
					))
				} else {
					endpoints = append(endpoints, endpoint.NewEndpoint(
						dnsName,
						record.Type,
						record.Value,
					))
				}
			}

			page++
			if page > records.Page {
				break
			}

			records, err = p.service.ListRecords(ctx, zone.Id, page)
			if err != nil {
				return nil, err
			}
		}
	}

	return endpoints, nil
}

// newDmeChange initializes a new change to dns records
func newDmeChange(action string, e *endpoint.Endpoint) *dmeChange {
	ttl := 1800
	if e.RecordTTL.IsConfigured() {
		ttl = int(e.RecordTTL)
	}

	change := &dmeChange{
		Action: action,
		ResourceRecordSet: Record{
			ManagedDNSRecordActions: models.ManagedDNSRecordActions{
				Name:  e.DNSName,
				Type:  e.RecordType,
				Value: e.Targets[0],
				Ttl:   strconv.Itoa(ttl),
			},
		},
	}
	return change
}

// newdnsmadeeasyChanges returns a slice of changes based on given action and record
func newDmeChanges(action string, endpoints []*endpoint.Endpoint) []*dmeChange {
	changes := make([]*dmeChange, 0, len(endpoints))
	for _, e := range endpoints {
		changes = append(changes, newDmeChange(action, e))
	}
	return changes
}

// submitChanges takes a zone and a collection of changes and makes all changes from the collection
func (p *dmeProvider) submitChanges(ctx context.Context, changes []*dmeChange) error {
	if len(changes) == 0 {
		log.Infof("All records are already up to date")
		return nil
	}
	zones, err := p.Zones(ctx)
	if err != nil {
		return err
	}
	for _, change := range changes {
		zoneId, zone := dnsmadeeasySuitableZone(change.ResourceRecordSet.Name, zones)
		if zone == nil {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", change.ResourceRecordSet.Name)
			continue
		}

		log.Infof("Changing records: %s %v in zone: %s", change.Action, change.ResourceRecordSet, zone.Name)

		if change.ResourceRecordSet.Name == zone.Name {
			change.ResourceRecordSet.Name = "" // Apex records have an empty name
		} else {
			change.ResourceRecordSet.Name = strings.TrimSuffix(change.ResourceRecordSet.Name, fmt.Sprintf(".%s", zone.Name))
		}

		recordAttributes := Record{
			ManagedDNSRecordActions: models.ManagedDNSRecordActions{
				Name:  change.ResourceRecordSet.Name,
				Type:  change.ResourceRecordSet.Type,
				Value: change.ResourceRecordSet.Value,
				Ttl:   change.ResourceRecordSet.Ttl,
			},
		}

		if !p.dryRun {
			switch change.Action {
			case dmeCreate:
				_, err := p.service.CreateRecord(ctx, zoneId, recordAttributes)
				if err != nil {
					return err
				}
			case dmeDelete:
				recordID, err := p.GetRecordID(ctx, zoneId, recordAttributes.Name)
				if err != nil {
					return err
				}
				err = p.service.DeleteRecord(ctx, zoneId, recordID)
				if err != nil {
					return err
				}
			case dmeUpdate:
				recordID, err := p.GetRecordID(ctx, zoneId, recordAttributes.Name)
				if err != nil {
					return err
				}
				_, err = p.service.UpdateRecord(ctx, zoneId, recordID, recordAttributes)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// GetRecordID returns the record ID for a given record name and zone.
func (p *dmeProvider) GetRecordID(ctx context.Context, zoneId string, recordName string) (recordID string, err error) {
	page := 1

	for {
		records, err := p.service.ListRecords(ctx, zoneId, page)
		if err != nil {
			return "", err
		}

		for _, record := range records.Records {
			switch record.Type {
			case "A", "CNAME", "TXT":
				break
			default:
				continue
			}

			if record.Name == recordName {
				return record.Id, nil
			}
		}

		page++
		if page > records.Page {
			break
		}

		records, err = p.service.ListRecords(ctx, zoneId, page)
		if err != nil {
			return "", err
		}
	}
	return "", fmt.Errorf("no record id found")
}

// dnsmadeeasySuitableZone returns the most suitable zone for a given hostname and a set of zones.
func dnsmadeeasySuitableZone(hostname string, zones map[string]Zone) (zoneId string, zoneDetails *Zone) {
	for zid, z := range zones {
		if strings.HasSuffix(hostname, z.Name) {
			if zoneDetails == nil || len(z.Name) > len(zoneDetails.Name) {
				newZ := z
				zoneDetails = &newZ
				zoneId = zid
			}
		}
	}
	return zoneId, zoneDetails
}

// CreateRecords creates records for a given slice of endpoints
func (p *dmeProvider) CreateRecords(ctx context.Context, endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(ctx, newDmeChanges(dmeCreate, endpoints))
}

// DeleteRecords deletes records for a given slice of endpoints
func (p *dmeProvider) DeleteRecords(ctx context.Context, endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(ctx, newDmeChanges(dmeDelete, endpoints))
}

// UpdateRecords updates records for a given slice of endpoints
func (p *dmeProvider) UpdateRecords(ctx context.Context, endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(ctx, newDmeChanges(dmeUpdate, endpoints))
}

// ApplyChanges applies a given set of changes
func (p *dmeProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	combinedChanges := make([]*dmeChange, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newDmeChanges(dmeCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newDmeChanges(dmeUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newDmeChanges(dmeDelete, changes.Delete)...)

	return p.submitChanges(ctx, combinedChanges)
}

func StripQuotes(word string) string {
	if strings.HasPrefix(word, "\"") && strings.HasSuffix(word, "\"") {
		return strings.TrimPrefix(strings.TrimSuffix(word, "\""), "\"")
	} else if word == "{}" {
		word = ""
		return word
	}
	return word
}
