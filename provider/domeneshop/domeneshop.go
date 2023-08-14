/*
Copyright 2017 The Kubernetes Authors.

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

package domeneshop

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"strings"
)

// DomeneshopProvider is an implementation of Provider for Domeneshop's DNS.
type DomeneshopProvider struct {
	provider.BaseProvider
	Client       DomeneshopClient
	domainFilter endpoint.DomainFilter
	DryRun       bool
}

// NewDomeneshopProvider initializes a new Domeneshop DNS based Provider.
func NewDomeneshopProvider(ctx context.Context, domainFilter endpoint.DomainFilter, dryRun bool, appVersion string) (*DomeneshopProvider, error) {
	apiToken, ok := os.LookupEnv("DOMENESHOP_API_TOKEN")
	if !ok {
		return nil, fmt.Errorf("no api token found")
	}

	apiSecret, ok := os.LookupEnv("DOMENESHOP_API_SECRET")
	if !ok {
		return nil, fmt.Errorf("no api secret found")
	}

	client := NewClient(apiToken, apiSecret, appVersion)

	provider := &DomeneshopProvider{
		Client:       client,
		domainFilter: domainFilter,
		DryRun:       dryRun,
	}
	return provider, nil
}

// Records returns the list of records from the domains.
func (p *DomeneshopProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	domains, err := p.Client.ListDomains(ctx)
	if err != nil {
		return nil, err
	}
	domains = filterDomains(domains, p.domainFilter)

	for _, domain := range domains {
		records, err := p.Client.ListDNSRecords(ctx, domain, "", "")
		if err != nil {
			return nil, err
		}

		for _, record := range records {
			dnsName := record.Host + "." + domain.Name

			// root name is identified by @ and should be
			// translated to zone name for the endpoint entry.
			if record.Host == "@" {
				dnsName = domain.Name
			}

			e := endpoint.NewEndpointWithTTL(dnsName, record.Type, endpoint.TTL(record.TTL), record.Data)

			endpoints = append(endpoints, e)
		}
	}

	log.WithFields(log.Fields{
		"endpoint_count": len(endpoints),
	}).Info("Found endpoints.")

	return mergeEndpointsByNameType(endpoints), nil
}

// ApplyChanges applies a given set of changes to the domains.
func (p *DomeneshopProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	domains, err := p.Client.ListDomains(ctx)
	if err != nil {
		return err
	}
	domains = filterDomains(domains, p.domainFilter)

	for _, domain := range domains {
		for _, change := range changes.Create {
			if strings.HasSuffix(change.DNSName, domain.Name) {
				log.WithFields(log.Fields{
					"endpoint": change,
				}).Info("Creating endpoint")

				for _, target := range change.Targets {
					record := recordFromEndpointTarget(domain, change, target)

					log.WithFields(log.Fields{
						"record": record,
					}).Info("Creating record")

					if p.DryRun {
						continue
					}

					err := p.Client.AddDNSRecord(ctx, domain, record)
					if err != nil {
						var httpErr *HttpError
						if errors.As(err, &httpErr) {
							if httpErr.Response.StatusCode == 409 && httpErr.ErrorBody.Code == "record:collision" {
								// collision is ok, usually because a record exists without a corresponding registry entry
								log.WithFields(log.Fields{
									"endpoint": change,
								}).Info("Ignoring collision when trying to create record, will be handled in a later update.")
							} else {
								return err
							}
						} else {
							return err
						}
					}
				}
			}
		}
		for _, change := range changes.UpdateNew {
			if strings.HasSuffix(change.DNSName, domain.Name) {
				log.WithFields(log.Fields{
					"endpoint": change,
				}).Info("Updating endpoint")

				// Trim the domain off the name if present.
				adjustedName := strings.TrimSuffix(change.DNSName, "."+domain.Name)

				// Record at the root should be defined as @ instead of the full domain name.
				if adjustedName == domain.Name {
					adjustedName = "@"
				}

				// find record to delete by lookup
				records, err := p.Client.ListDNSRecords(ctx, domain, adjustedName, change.RecordType)
				if err != nil {
					return err
				}

				unmatchedRecords := make(map[int]*DNSRecord)
				for _, record := range records {
					unmatchedRecords[record.ID] = record
				}

			TARGET:
				for _, target := range change.Targets {
					for _, record := range records {
						// first, find a record with the same target
						if record.Data == target {
							record.Host = adjustedName
							record.Data = target
							record.TTL = getTTLFromEndpoint(change)

							log.WithFields(log.Fields{
								"record": record,
							}).Info("Updating record")

							delete(unmatchedRecords, record.ID)

							if p.DryRun {
								continue
							}

							if err := p.Client.UpdateDNSRecord(ctx, domain, record); err != nil {
								return err
							}
							continue TARGET
						}
					}
					// didn't find a matching record to update, should add new record
					record := recordFromEndpointTarget(domain, change, target)

					log.WithFields(log.Fields{
						"record": record,
					}).Info("Creating record")

					if p.DryRun {
						continue
					}

					if err := p.Client.AddDNSRecord(ctx, domain, record); err != nil {
						return err
					}
				}

				// remove records for removed targets
				for _, record := range unmatchedRecords {
					log.WithFields(log.Fields{
						"record": record,
					}).Info("Deleting record")

					if p.DryRun {
						continue
					}

					if err := p.Client.DeleteDNSRecord(ctx, domain, record); err != nil {
						return err
					}
				}
			}
		}

		for _, change := range changes.Delete {
			if strings.HasSuffix(change.DNSName, domain.Name) {
				log.WithFields(log.Fields{
					"endpoint": change,
				}).Info("Deleting endpoint")

				// Trim the domain off the name if present.
				adjustedName := strings.TrimSuffix(change.DNSName, "."+domain.Name)

				// Record at the root should be defined as @ instead of the full domain name.
				if adjustedName == domain.Name {
					adjustedName = "@"
				}

				// find record to delete by lookup?
				records, err := p.Client.ListDNSRecords(ctx, domain, adjustedName, change.RecordType)
				if err != nil {
					return err
				}

				for _, record := range records {
					if len(change.Targets) > 0 {
						for _, target := range change.Targets {
							if record.Data == target {
								log.WithFields(log.Fields{
									"record": record,
								}).Info("Deleting record")

								if p.DryRun {
									continue
								}

								if err := p.Client.DeleteDNSRecord(ctx, domain, record); err != nil {
									return err
								}
							}
						}
						continue
					}

					log.WithFields(log.Fields{
						"record": record,
					}).Info("Deleting record")

					if p.DryRun {
						continue
					}

					if err := p.Client.DeleteDNSRecord(ctx, domain, record); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func filterDomains(domains []*Domain, domainFilter endpoint.DomainFilter) []*Domain {
	if !domainFilter.IsConfigured() {
		return domains
	}
	var result []*Domain

	for _, domain := range domains {
		if domainFilter.Match(domain.Name) {
			result = append(result, domain)
		}
	}

	return result
}

func recordFromEndpointTarget(domain *Domain, e *endpoint.Endpoint, target string) *DNSRecord {
	record := &DNSRecord{
		Type: e.RecordType,
		Host: strings.TrimSuffix(e.DNSName, "."+domain.Name),
		Data: target,
		TTL:  getTTLFromEndpoint(e),
	}

	if record.Host == domain.Name {
		record.Host = "@"
	}

	return record
}

func getTTLFromEndpoint(ep *endpoint.Endpoint) int {
	if ep.RecordTTL.IsConfigured() {
		return int(ep.RecordTTL)
	}
	return defaultTTL
}

// Merge Endpoints with the same Name and Type into a single endpoint with multiple Targets.
func mergeEndpointsByNameType(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	endpointsByNameType := map[string][]*endpoint.Endpoint{}

	for _, e := range endpoints {
		key := fmt.Sprintf("%s-%s", e.DNSName, e.RecordType)
		endpointsByNameType[key] = append(endpointsByNameType[key], e)
	}

	// If no merge occurred, just return the existing endpoints.
	if len(endpointsByNameType) == len(endpoints) {
		return endpoints
	}

	// Otherwise, construct a new list of endpoints with the endpoints merged.
	var result []*endpoint.Endpoint
	for _, endpoints := range endpointsByNameType {
		dnsName := endpoints[0].DNSName
		recordType := endpoints[0].RecordType

		targets := make([]string, len(endpoints))
		for i, e := range endpoints {
			targets[i] = e.Targets[0]
		}

		e := endpoint.NewEndpoint(dnsName, recordType, targets...)
		result = append(result, e)
	}

	return result
}
