/*
Copyright 2021 The Kubernetes Authors.
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

package gandi

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/go-gandi/go-gandi"
	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/livedns"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	gandiCreate          = "CREATE"
	gandiDelete          = "DELETE"
	gandiUpdate          = "UPDATE"
	gandiTTL             = 600
	gandiLiveDNSProvider = "livedns"
)

type GandiChanges struct {
	Action   string
	ZoneName string
	Record   livedns.DomainRecord
}

type GandiProvider struct {
	provider.BaseProvider
	LiveDNSClient LiveDNSClientAdapter
	DomainClient  DomainClientAdapter
	domainFilter  endpoint.DomainFilter
	DryRun        bool
}

func NewGandiProvider(ctx context.Context, domainFilter endpoint.DomainFilter, dryRun bool) (*GandiProvider, error) {
	key, ok_key := os.LookupEnv("GANDI_KEY")
	pat, ok_pat := os.LookupEnv("GANDI_PAT")
	if !(ok_key || ok_pat) {
		return nil, errors.New("no environment variable GANDI_KEY or GANDI_PAT provided")
	}
	if ok_key {
		log.Warning("Usage of GANDI_KEY (API Key) is deprecated. Please consider creating a Personal Access Token (PAT) instead, see https://api.gandi.net/docs/authentication/")
	}
	sharingID, _ := os.LookupEnv("GANDI_SHARING_ID")

	g := config.Config{
		APIKey:              key,
		PersonalAccessToken: pat,
		SharingID:           sharingID,
		Debug:               false,
		// dry-run doesn't work but it won't hurt passing the flag
		DryRun: dryRun,
	}

	liveDNSClient := gandi.NewLiveDNSClient(g)
	domainClient := gandi.NewDomainClient(g)

	gandiProvider := &GandiProvider{
		LiveDNSClient: NewLiveDNSClient(liveDNSClient),
		DomainClient:  NewDomainClient(domainClient),
		domainFilter:  domainFilter,
		DryRun:        dryRun,
	}
	return gandiProvider, nil
}

func (p *GandiProvider) Zones() (zones []string, err error) {
	availableDomains, err := p.DomainClient.ListDomains()
	if err != nil {
		return nil, err
	}
	zones = []string{}
	for _, domain := range availableDomains {
		if !p.domainFilter.Match(domain.FQDN) {
			log.Debugf("Excluding domain %s by domain-filter", domain.FQDN)
			continue
		}

		if domain.NameServer.Current != gandiLiveDNSProvider {
			log.Debugf("Excluding domain %s, not configured for livedns", domain.FQDN)
			continue
		}

		zones = append(zones, domain.FQDN)
	}
	return zones, nil
}

func (p *GandiProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	liveDNSZones, err := p.Zones()
	if err != nil {
		return nil, err
	}
	endpoints := []*endpoint.Endpoint{}
	for _, zone := range liveDNSZones {
		records, err := p.LiveDNSClient.GetDomainRecords(zone)
		if err != nil {
			return nil, err
		}

		for _, r := range records {
			if provider.SupportedRecordType(r.RrsetType) {
				name := r.RrsetName + "." + zone

				if r.RrsetName == "@" {
					name = zone
				}

				for _, v := range r.RrsetValues {
					log.WithFields(log.Fields{
						"record": r.RrsetName,
						"type":   r.RrsetType,
						"value":  v,
						"ttl":    r.RrsetTTL,
						"zone":   zone,
					}).Debug("Returning endpoint record")

					endpoints = append(
						endpoints,
						endpoint.NewEndpointWithTTL(name, r.RrsetType, endpoint.TTL(r.RrsetTTL), v),
					)
				}
			}
		}
	}

	return endpoints, nil
}

func (p *GandiProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	combinedChanges := make([]*GandiChanges, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, p.newGandiChanges(gandiCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, p.newGandiChanges(gandiUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, p.newGandiChanges(gandiDelete, changes.Delete)...)

	return p.submitChanges(ctx, combinedChanges)
}

func (p *GandiProvider) submitChanges(ctx context.Context, changes []*GandiChanges) error {
	if len(changes) == 0 {
		log.Infof("All records are already up to date")
		return nil
	}

	liveDNSDomains, err := p.Zones()
	if err != nil {
		return err
	}

	zoneChanges := p.groupAndFilterByZone(liveDNSDomains, changes)

	for _, changes := range zoneChanges {
		for _, change := range changes {
			if change.Record.RrsetType == endpoint.RecordTypeCNAME && !strings.HasSuffix(change.Record.RrsetValues[0], ".") {
				change.Record.RrsetValues[0] += "."
			}

			// Prepare record name
			if change.Record.RrsetName == change.ZoneName {
				log.WithFields(log.Fields{
					"record": change.Record.RrsetName,
					"type":   change.Record.RrsetType,
					"value":  change.Record.RrsetValues[0],
					"ttl":    change.Record.RrsetTTL,
					"action": change.Action,
					"zone":   change.ZoneName,
				}).Debugf("Converting record name: %s to apex domain (@)", change.Record.RrsetName)

				change.Record.RrsetName = "@"
			} else {
				change.Record.RrsetName = strings.TrimSuffix(
					change.Record.RrsetName,
					"."+change.ZoneName,
				)
			}

			log.WithFields(log.Fields{
				"record": change.Record.RrsetName,
				"type":   change.Record.RrsetType,
				"value":  change.Record.RrsetValues[0],
				"ttl":    change.Record.RrsetTTL,
				"action": change.Action,
				"zone":   change.ZoneName,
			}).Info("Changing record")

			if !p.DryRun {
				switch change.Action {
				case gandiCreate:
					answer, err := p.LiveDNSClient.CreateDomainRecord(
						change.ZoneName,
						change.Record.RrsetName,
						change.Record.RrsetType,
						change.Record.RrsetTTL,
						change.Record.RrsetValues,
					)
					if err != nil {
						log.WithFields(log.Fields{
							"Code":    answer.Code,
							"Message": answer.Message,
							"Cause":   answer.Cause,
							"Errors":  answer.Errors,
						}).Warning("Create problem")
						return err
					}
				case gandiDelete:
					err := p.LiveDNSClient.DeleteDomainRecord(change.ZoneName, change.Record.RrsetName, change.Record.RrsetType)
					if err != nil {
						log.Warning("Delete problem")
						return err
					}
				case gandiUpdate:
					answer, err := p.LiveDNSClient.UpdateDomainRecordByNameAndType(
						change.ZoneName,
						change.Record.RrsetName,
						change.Record.RrsetType,
						change.Record.RrsetTTL,
						change.Record.RrsetValues,
					)
					if err != nil {
						log.WithFields(log.Fields{
							"Code":    answer.Code,
							"Message": answer.Message,
							"Cause":   answer.Cause,
							"Errors":  answer.Errors,
						}).Warning("Update problem")
						return err
					}
				}
			}
		}
	}

	return nil
}

func (p *GandiProvider) newGandiChanges(action string, endpoints []*endpoint.Endpoint) []*GandiChanges {
	changes := make([]*GandiChanges, 0, len(endpoints))
	ttl := gandiTTL
	for _, e := range endpoints {
		if e.RecordTTL.IsConfigured() {
			ttl = int(e.RecordTTL)
		}
		change := &GandiChanges{
			Action: action,
			Record: livedns.DomainRecord{
				RrsetType:   e.RecordType,
				RrsetName:   e.DNSName,
				RrsetValues: e.Targets,
				RrsetTTL:    ttl,
			},
		}
		changes = append(changes, change)
	}
	return changes
}

func (p *GandiProvider) groupAndFilterByZone(zones []string, changes []*GandiChanges) map[string][]*GandiChanges {
	change := make(map[string][]*GandiChanges)
	zoneNameID := provider.ZoneIDName{}

	for _, z := range zones {
		zoneNameID.Add(z, z)
		change[z] = []*GandiChanges{}
	}

	for _, c := range changes {
		zoneID, zoneName := zoneNameID.FindZone(c.Record.RrsetName)
		if zoneName == "" {
			log.Debugf("Skipping record %s because no hosted domain matching record DNS Name was detected", c.Record.RrsetName)
			continue
		}
		c.ZoneName = zoneName
		change[zoneID] = append(change[zoneID], c)
	}
	return change
}
