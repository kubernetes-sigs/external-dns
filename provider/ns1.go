package provider

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	api "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

const (
	// ns1Create is a ChangeAction enum value
	ns1Create = "CREATE"
	// ns1Delete is a ChangeAction enum value
	ns1Delete = "DELETE"
	// ns1Update is a ChangeAction enum value
	ns1Update = "UPDATE"
	// ns1DefaultTTL is the default ttl for ttls that are not set
	ns1DefaultTTL = 10
)

// NS1Config passes cli args to the NS1Provider
type NS1Config struct {
	DomainFilter DomainFilter
	ZoneIDFilter ZoneIDFilter
	DryRun       bool
}

// NS1Provider is the NS1 provider
type NS1Provider struct {
	client       *api.Client
	domainFilter DomainFilter
	zoneIDFilter ZoneIDFilter
	dryRun       bool
}

// NewNS1Provider creates a new NS1 Provider
func NewNS1Provider(config NS1Config) (*NS1Provider, error) {
	return newNS1ProviderWithHTTPClient(config, http.DefaultClient)
}

func newNS1ProviderWithHTTPClient(config NS1Config, client *http.Client) (*NS1Provider, error) {
	token, ok := os.LookupEnv("NS1_APIKEY")
	if !ok {
		return nil, fmt.Errorf("NS1_APIKEY environment variable is not set")
	}

	apiClient := api.NewClient(client, api.SetAPIKey(token))

	provider := &NS1Provider{
		client:       apiClient,
		domainFilter: config.DomainFilter,
		zoneIDFilter: config.ZoneIDFilter,
	}
	return provider, nil
}

func (p *NS1Provider) matchEither(id string) bool {
	return p.domainFilter.Match(id) || p.zoneIDFilter.Match(id)
}

// Records returns the endpoints this provider knows about
func (p *NS1Provider) Records() ([]*endpoint.Endpoint, error) {
	zones, err := p.zonesFiltered()
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint

	for _, zone := range zones {

		// TODO handle Header Codes
		zoneData, _, err := p.client.Zones.Get(zone.String())
		if err != nil {
			return nil, err
		}

		for _, record := range zoneData.Records {
			if supportedRecordType(record.Type) {
				endpoints = append(endpoints, endpoint.NewEndpointWithTTL(
					record.Domain,
					record.Type,
					endpoint.TTL(record.TTL),
					record.ShortAns...,
				),
				)
			}
		}
	}

	return endpoints, nil
}

func ns1BuildRecord(zoneName string, change *ns1Change) *dns.Record {
	record := dns.NewRecord(zoneName, change.Endpoint.DNSName, change.Endpoint.RecordType)
	for _, v := range change.Endpoint.Targets {
		record.AddAnswer(dns.NewAnswer(strings.Split(v, " ")))
	}
	// set detault ttl
	var ttl = ns1DefaultTTL
	if change.Endpoint.RecordTTL.IsConfigured() {
		ttl = int(change.Endpoint.RecordTTL)
	}
	record.TTL = ttl

	return record
}

func (p *NS1Provider) ns1SubmitChanges(changes []*ns1Change) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	zones, err := p.zonesFiltered()
	if err != nil {
		return err
	}

	// separate into per-zone change sets to be passed to the API.
	changesByZone := ns1ChangesByZone(zones, changes)
	for zoneName, changes := range changesByZone {
		for _, change := range changes {
			record := ns1BuildRecord(zoneName, change)
			logFields := log.Fields{
				"record": record.Domain,
				"type":   record.Type,
				"ttl":    record.TTL,
				"action": change.Action,
				"zone":   zoneName,
			}

			log.WithFields(logFields).Info("Changing record.")

			if p.dryRun {
				continue
			}

			switch change.Action {
			case ns1Create:
				_, err := p.client.Records.Create(record)
				if err != nil {
					return err
				}
			case ns1Delete:
				_, err := p.client.Records.Delete(zoneName, record.Domain, record.Type)
				if err != nil {
					return err
				}
			case ns1Update:
				_, err := p.client.Records.Update(record)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Zones returns the list of hosted zones.
func (p *NS1Provider) zonesFiltered() ([]*dns.Zone, error) {
	// TODO handle Header Codes
	zones, _, err := p.client.Zones.List()
	if err != nil {
		return nil, err
	}

	toReturn := []*dns.Zone{}

	for _, z := range zones {
		if !p.matchEither(z.Zone) && !p.matchEither(z.ID) {
			continue
		}
		toReturn = append(toReturn, z)
	}

	return toReturn, nil
}

// ns1Change differentiates between ChangActions
type ns1Change struct {
	Action   string
	Endpoint *endpoint.Endpoint
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *NS1Provider) ApplyChanges(changes *plan.Changes) error {
	combinedChanges := make([]*ns1Change, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newNS1Changes(ns1Create, changes.Create)...)
	combinedChanges = append(combinedChanges, newNS1Changes(ns1Update, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newNS1Changes(ns1Delete, changes.Delete)...)

	return p.ns1SubmitChanges(combinedChanges)
}

// newNS1Changes returns a collection of Changes based on the given records and action.
func newNS1Changes(action string, endpoints []*endpoint.Endpoint) []*ns1Change {
	changes := make([]*ns1Change, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, &ns1Change{
			Action:   action,
			Endpoint: endpoint,
		},
		)
	}

	return changes
}

// ns1ChangesByZone separates a multi-zone change into a single change per zone.
func ns1ChangesByZone(zones []*dns.Zone, changeSets []*ns1Change) map[string][]*ns1Change {
	changes := make(map[string][]*ns1Change)
	zoneNameIDMapper := zoneIDName{}
	for _, z := range zones {
		zoneNameIDMapper.Add(z.Zone, z.Zone)
		changes[z.Zone] = []*ns1Change{}
	}

	for _, c := range changeSets {
		zone, _ := zoneNameIDMapper.FindZone(c.Endpoint.DNSName)
		if zone == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected ", c.Endpoint.DNSName)
			continue
		}
		changes[zone] = append(changes[zone], c)
	}

	return changes
}
