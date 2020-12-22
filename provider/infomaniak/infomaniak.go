/*
Copyright 2020 The Kubernetes Authors.
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

package infomaniak

import (
	"context"
	"errors"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const APITokenVariable = "INFOMANIAK_API_TOKEN"

type InfomaniakProvider struct {
	provider.BaseProvider
	API          *InfomaniakAPI
	domainFilter endpoint.DomainFilter
	DryRun       bool
}

func NewInfomaniakProvider(ctx context.Context, domainFilter endpoint.DomainFilter, dryRun bool) (*InfomaniakProvider, error) {
	token, ok := os.LookupEnv(APITokenVariable)
	if !ok {
		return nil, errors.New("environment variable " + APITokenVariable + " missing")
	}

	api := NewInfomaniakAPI(token)

	provider := &InfomaniakProvider{
		API:          api,
		domainFilter: domainFilter,
		DryRun:       dryRun,
	}
	return provider, nil
}

func (p *InfomaniakProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.API.ListDomains()
	if err != nil {
		return nil, err
	}
	endpoints := []*endpoint.Endpoint{}
	for _, zone := range *zones {
		if !p.domainFilter.Match(zone.CustomerName) {
			continue
		}

		records, err := p.API.GetRecords(&zone)
		if err != nil {
			return nil, err
		}

		for _, r := range *records {
			if provider.SupportedRecordType(r.Type) {
				name := r.Source + "." + zone.CustomerName

				if r.Source == "." {
					name = zone.CustomerName
				}

				endpoints = append(endpoints, endpoint.NewEndpoint(name, r.Type, r.Target))
			}
		}
	}

	for i, endpoint := range endpoints {
		log.Infof("Endpoint %d: %s / %s / %s", i, endpoint.DNSName, endpoint.RecordType, endpoint.Targets)
	}
	return endpoints, nil
}

// findMatchingZone find longest matching domain from list
func findMatchingZone(pdomains *[]InfomaniakDNSDomain, record string) (*InfomaniakDNSDomain, string, error) {
	domains := *pdomains
	sort.Slice(domains, func(i, j int) bool { return len(domains[i].CustomerName) > len(domains[j].CustomerName) })
	for _, domain := range domains {
		if strings.HasSuffix(record, domain.CustomerName) {
			return &domain, strings.TrimSuffix(record, "."+domain.CustomerName), nil
		}
	}
	return nil, "", errors.New("not found")
}

func (p *InfomaniakProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if len(changes.Create) == 0 && len(changes.UpdateNew) == 0 && len(changes.Delete) == 0 {
		log.Infof("All records are already up to date")
		return nil
	}

	var zones []InfomaniakDNSDomain

	allZones, err := p.API.ListDomains()
	if err != nil {
		return err
	}
	for _, zone := range *allZones {
		if p.domainFilter.Match(zone.CustomerName) {
			zones = append(zones, zone)
		}
	}

	log.Infof("Records to create:")
	for _, endpoint := range changes.Create {
		zone, source, err := findMatchingZone(&zones, endpoint.DNSName)
		if err != nil {
			return err
		}
		log.Infof("(%s) %s IN %s:", zone.CustomerName, source, endpoint.RecordType)
		for i, target := range endpoint.Targets {
			log.Infof("\t - %d : %s", i, target)
			target = strings.Trim(target, "\"")
			err := p.API.EnsureDNSRecord(zone, source, target, endpoint.RecordType, uint64(endpoint.RecordTTL))
			if err != nil {
				return err
			}
		}
	}

	log.Infof("Records to delete:")
	for _, endpoint := range changes.Delete {
		zone, source, err := findMatchingZone(&zones, endpoint.DNSName)
		if err != nil {
			return err
		}
		log.Infof("(%s) %s IN %s:", zone.CustomerName, source, endpoint.RecordType)
		for i, target := range endpoint.Targets {
			log.Infof("\t - %d : %s", i, target)
			target = strings.Trim(target, "\"")
			err := p.API.RemoveDNSRecord(zone, source, target, endpoint.RecordType)
			if err != nil {
				return err
			}
		}
	}

	log.Infof("Records to update:")
	for iz, endpoint := range changes.UpdateOld {
		zone, source, err := findMatchingZone(&zones, endpoint.DNSName)
		if err != nil {
			return err
		}
		newEndpoint := changes.UpdateNew[iz]
		log.Infof("(%s) %s IN %s:", zone.CustomerName, source, endpoint.RecordType)
		for i, target := range endpoint.Targets {
			newTarget := newEndpoint.Targets[i]
			if target == newTarget && endpoint.RecordTTL == newEndpoint.RecordTTL {
				log.Infof("\t[skip] - %s (%d)", target, endpoint.RecordTTL)
				continue
			}
			err := p.API.ModifyDNSRecord(zone, source, target, newTarget, endpoint.RecordType, uint64(newEndpoint.RecordTTL))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
