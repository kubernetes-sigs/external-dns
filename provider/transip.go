package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/transip/gotransip/v6"
	transipdomain "github.com/transip/gotransip/v6/domain"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const (
	// 60 seconds is the current minimal TTL for TransIP and will replace unconfigured
	// TTL's for Endpoints
	transipMinimalValidTTL = 60
)

// TransIPProvider is an implementation of Provider for TransIP.
type TransIPProvider struct {
	repo         transipdomain.Repository
	domainFilter endpoint.DomainFilter
	dryRun       bool
}

// TransIPZone holds the domain and its dns dnsEntries as the Domain object doesn't contain the dnsEntries itself
type TransIPZone struct {
	domain     transipdomain.Domain
	dnsEntries []transipdomain.DNSEntry
}

// NewTransIPProvider initializes a new TransIP Provider.
func NewTransIPProvider(accountName, privateKeyFile string, domainFilter endpoint.DomainFilter, dryRun bool) (*TransIPProvider, error) {
	// check given arguments
	if accountName == "" {
		return nil, errors.New("required --transip-account not set")
	}

	if privateKeyFile == "" {
		return nil, errors.New("required --transip-keyfile not set")
	}

	var apiMode gotransip.APIMode
	if dryRun {
		apiMode = gotransip.APIModeReadOnly
	} else {
		apiMode = gotransip.APIModeReadWrite
	}

	// create new TransIP API repo
	c, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:    accountName,
		PrivateKeyPath: privateKeyFile,
		Mode:           apiMode,
	})
	if err != nil {
		return nil, fmt.Errorf("could not setup TransIP API repo: %s", err.Error())
	}

	repo := transipdomain.Repository{Client: c}

	// return tipCloud struct
	return &TransIPProvider{
		repo:         repo,
		domainFilter: domainFilter,
		dryRun:       dryRun,
	}, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *TransIPProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// build zonefinder with all our zones so we can use FindZone
	// and a mapping of zones and their domain name
	zones, err := p.repo.GetAll()
	if err != nil {
		return err
	}

	zoneNameMapper := zoneIDName{}
	zonesByName := make(map[string]TransIPZone)
	updatedZones := make(map[string]bool)
	for _, zone := range zones {
		// TransIP API doesn't expose the zone when fetching all domains. Instead we have to fetch them for each
		// domain individually.

		entries, err := p.repo.GetDNSEntries(zone.Name)
		if err != nil {
			return err
		}

		transipZone := TransIPZone{
			domain:     zone,
			dnsEntries: entries,
		}

		// TransIP API doesn't expose a unique identifier for zones, other than than
		// the domain name itself.
		zoneNameMapper.Add(zone.Name, zone.Name)
		zonesByName[zone.Name] = transipZone
	}

	// first see if we need to delete anything
	for _, ep := range changes.Delete {
		log.WithFields(log.Fields{"record": ep.DNSName, "type": ep.RecordType}).Info("endpoint has to go")

		zone, err := p.zoneForZoneName(ep.DNSName, zoneNameMapper, zonesByName)
		if err != nil {
			log.Errorf("could not find zone for %s: %s", ep.DNSName, err.Error())
			continue
		}

		log.Debugf("removing records for %s", zone.domain.Name)

		// remove current records from DNS entry set
		entries := p.removeEndpointFromEntries(ep, zone)

		// update zone in zone map
		zone.dnsEntries = entries
		zonesByName[zone.domain.Name] = zone
		// flag zone for updating
		updatedZones[zone.domain.Name] = true
	}

	for _, ep := range changes.Create {
		log.WithFields(log.Fields{"record": ep.DNSName, "type": ep.RecordType}).Info("endpoint is missing")

		zone, err := p.zoneForZoneName(ep.DNSName, zoneNameMapper, zonesByName)
		if err != nil {
			log.Errorf("could not find zone for %s: %s", ep.DNSName, err.Error())
			continue
		}

		log.Debugf("creating records for %s", zone.domain.Name)

		// add new dnsEntries to set
		zone.dnsEntries = p.addEndpointToEntries(ep, zone, zone.dnsEntries)

		// update zone in zone map
		zonesByName[zone.domain.Name] = zone
		// flag zone for updating
		updatedZones[zone.domain.Name] = true
		log.WithFields(log.Fields{"zone": zone.domain.Name}).Debug("flagging for update")
	}

	for _, ep := range changes.UpdateNew {
		log.WithFields(log.Fields{"record": ep.DNSName, "type": ep.RecordType}).Debug("needs updating")

		zone, err := p.zoneForZoneName(ep.DNSName, zoneNameMapper, zonesByName)
		if err != nil {
			log.WithFields(log.Fields{"record": ep.DNSName}).Warn(err.Error())
			continue
		}

		// updating the records is basically finding all matching records according
		// to the name and the type, removing them from the set and add the new
		// records
		log.WithFields(log.Fields{
			"zone":       zone.domain.Name,
			"dnsname":    ep.DNSName,
			"recordtype": ep.RecordType,
		}).Debug("removing matching dnsEntries")

		// remove current records from DNS entry set
		entries := p.removeEndpointFromEntries(ep, zone)

		// add new dnsEntries to set
		entries = p.addEndpointToEntries(ep, zone, entries)

		// check to see if actually anything changed in the DNSEntry set
		if p.dnsEntriesAreEqual(entries, zone.dnsEntries) {
			log.WithFields(log.Fields{"zone": zone.domain.Name}).Debug("not updating identical dnsEntries")
			continue
		}

		// update zone in zone map
		zone.dnsEntries = entries
		zonesByName[zone.domain.Name] = zone
		// flag zone for updating
		updatedZones[zone.domain.Name] = true

		log.WithFields(log.Fields{"zone": zone.domain.Name}).Debug("flagging for update")
	}

	// go over all updated zones and set new DNSEntry set
	for uz := range updatedZones {
		zone, ok := zonesByName[uz]
		if !ok {
			log.WithFields(log.Fields{"zone": uz}).Debug("updated zone no longer found")
			continue
		}

		if p.dryRun {
			log.WithFields(log.Fields{"zone": zone.domain.Name}).Info("not updating in dry-run mode")
			continue
		}

		log.WithFields(log.Fields{"zone": zone.domain.Name}).Info("updating DNS dnsEntries")
		if err := p.repo.ReplaceDNSEntries(zone.domain.Name, zone.dnsEntries); err != nil {
			log.WithFields(log.Fields{"zone": zone.domain.Name, "error": err.Error()}).Warn("failed to update")
		}
	}

	return nil
}

// fetchZones returns a list of all domains within the account
func (p *TransIPProvider) fetchZones() ([]TransIPZone, error) {
	domains, err := p.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var zones []TransIPZone
	for _, d := range domains {
		if !p.domainFilter.Match(d.Name) {
			continue
		}

		entries, err := p.repo.GetDNSEntries(d.Name)
		if err != nil {
			return nil, err
		}

		z := TransIPZone{
			domain:     d,
			dnsEntries: entries,
		}
		zones = append(zones, z)
	}

	return zones, nil
}

// Zones returns the list of hosted zones.
func (p *TransIPProvider) Zones() ([]TransIPZone, error) {
	zones, err := p.fetchZones()
	if err != nil {
		return nil, err
	}

	return zones, nil
}

// Records returns the list of records in a given zone.
func (p *TransIPProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.Zones()
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint
	var name string
	// go over all zones and their DNS dnsEntries and create endpoints for them
	for _, zone := range zones {
		for _, r := range zone.dnsEntries {
			if !supportedRecordType(r.Type) {
				continue
			}

			name = p.endpointNameForRecord(r, zone.domain)
			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(name, string(r.Type), endpoint.TTL(r.Expire), r.Content))
		}
	}

	return endpoints, nil
}

// endpointNameForRecord returns "www.example.org" for DNSEntry with Name "www" and
// Doman with Name "example.org"
func (p *TransIPProvider) endpointNameForRecord(r transipdomain.DNSEntry, d transipdomain.Domain) string {
	// root name is identified by "@" and should be translated to domain name for
	// the endpoint entry.
	if r.Name == "@" {
		return d.Name
	}

	return fmt.Sprintf("%s.%s", r.Name, d.Name)
}

// recordNameForEndpoint returns "www" for Endpoint with DNSName "www.example.org"
// and Domain with Name "example.org"
func (p *TransIPProvider) recordNameForEndpoint(ep *endpoint.Endpoint, d transipdomain.Domain) string {
	// root name is identified by "@" and should be translated to domain name for
	// the endpoint entry.
	if ep.DNSName == d.Name {
		return "@"
	}

	return strings.TrimSuffix(ep.DNSName, "."+d.Name)
}

// getMinimalValidTTL returns max between given Endpoint's RecordTTL and
// transipMinimalValidTTL
func (p *TransIPProvider) getMinimalValidTTL(ep *endpoint.Endpoint) int64 {
	// TTL cannot be lower than transipMinimalValidTTL
	if ep.RecordTTL < transipMinimalValidTTL {
		return transipMinimalValidTTL
	}

	return int64(ep.RecordTTL)
}

// dnsEntriesAreEqual compares the dnsEntries in 2 sets and returns true if the
// content of the dnsEntries is equal
func (p *TransIPProvider) dnsEntriesAreEqual(a, b []transipdomain.DNSEntry) bool {
	if len(a) != len(b) {
		return false
	}

	match := 0
	for _, aa := range a {
		for _, bb := range b {
			if aa.Content != bb.Content {
				continue
			}

			if aa.Name != bb.Name {
				continue
			}

			if aa.Expire != bb.Expire {
				continue
			}

			if aa.Type != bb.Type {
				continue
			}

			match++
		}
	}

	return len(a) == match
}

// removeEndpointFromEntries removes DNS dnsEntries from zone's set that match the
// type and name from given endpoint and returns the resulting DNS entry set
func (p *TransIPProvider) removeEndpointFromEntries(ep *endpoint.Endpoint, zone TransIPZone) []transipdomain.DNSEntry {
	// create new entry set
	entries := make([]transipdomain.DNSEntry, 0)
	// go over each DNS entry to see if it is a match
	for _, e := range zone.dnsEntries {
		// if we have match, don't copy it to the new entry set
		if p.endpointNameForRecord(e, zone.domain) == ep.DNSName && string(e.Type) == ep.RecordType {
			log.WithFields(log.Fields{
				"name":    e.Name,
				"content": e.Content,
				"type":    e.Type,
			}).Debug("found match")
			continue
		}

		entries = append(entries, e)
	}

	return entries
}

// addEndpointToEntries creates DNS dnsEntries for given endpoint and returns
// resulting DNS entry set
func (p *TransIPProvider) addEndpointToEntries(ep *endpoint.Endpoint, zone TransIPZone, entries []transipdomain.DNSEntry) []transipdomain.DNSEntry {
	ttl := p.getMinimalValidTTL(ep)
	for _, target := range ep.Targets {
		log.WithFields(log.Fields{
			"zone":       zone.domain.Name,
			"dnsname":    ep.DNSName,
			"recordtype": ep.RecordType,
			"ttl":        ttl,
			"target":     target,
		}).Debugf("adding new record")
		entries = append(entries, transipdomain.DNSEntry{
			Name:    p.recordNameForEndpoint(ep, zone.domain),
			Expire:  int(ttl),
			Type:    ep.RecordType,
			Content: target,
		})
	}

	return entries
}

// zoneForZoneName returns the zone mapped to given name or error if zone could
// not be found
func (p *TransIPProvider) zoneForZoneName(name string, m zoneIDName, z map[string]TransIPZone) (TransIPZone, error) {
	_, zoneName := m.FindZone(name)
	if zoneName == "" {
		return TransIPZone{}, fmt.Errorf("could not find zoneName for %s", name)
	}

	zone, ok := z[zoneName]
	if !ok {
		return zone, fmt.Errorf("could not find zone for %s", zoneName)
	}

	return zone, nil
}
