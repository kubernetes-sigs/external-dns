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

package provider

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/Azure/azure-sdk-for-go/arm/dns"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// azureProvider is an implementation of Provider for Azure DNS.
type azureProvider struct {
	// Azure resource group containing the DNS zones
	resourceGroupName string
	// Enabled dry-run will print any modifying actions rather than execute them.
	dryRun bool
	// A client for managing resource record sets
	recordSetsClient dns.RecordSetsClient
	// A client for managing hosted zones
	zonesClient dns.ZonesClient
}

type azureEndpoint struct {
   endpoint.Endpoint
}

// NewAzureProvider initializes a new Azure DNS based Provider.
func NewAzureProvider(subscriptionID string, resourceGroupName string, dryRun bool) (Provider, error) {
	recordsClient := dns.NewRecordSetsClient(subscriptionID)
	zonesClient := dns.NewZonesClient(subscriptionID)

	provider := &azureProvider{
		resourceGroupName: resourceGroupName,
		recordSetsClient: recordsClient,
		zonesClient: zonesClient,
		dryRun: dryRun,
	}

	return provider, nil
}

func (e *azureEndpoint) getRecordSetProperties() (*dns.RecordSetProperties, error) {
	switch e.RecordType {
	case "CNAME":
		return &dns.RecordSetProperties{
			CnameRecord: &dns.CnameRecord{
				Cname: &e.Target,
			},
		}, nil
	case "A":
		return &dns.RecordSetProperties{
			ARecords: &[]dns.ARecord{
				dns.ARecord{
					Ipv4Address: &e.Target,
				},
			},
		}, nil
	case "TXT":
		v := strings.Fields(e.Target)
		return &dns.RecordSetProperties{
			TxtRecords: &[]dns.TxtRecord{
				dns.TxtRecord{
					Value: &v,
				},
			},
		}, nil
	}
	return nil, fmt.Errorf("Invalid record type. Only CNAME, A and TXT are currently supported.")
}

func (e *azureEndpoint) toRecordSet() (dns.RecordSet, error) {
	p, err := e.getRecordSetProperties()
	if err != nil {
		return dns.RecordSet{}, err
	}
	return dns.RecordSet{
		Name: &e.DNSName,
		Type: &e.RecordType,
		RecordSetProperties: p,
	}, nil
}

func RecordSetToEndpoints(r dns.RecordSet) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	switch *r.Type {
	case "CNAME":
		endpoints = append(endpoints, endpoint.NewEndpoint(*r.Name, *r.RecordSetProperties.CnameRecord.Cname, *r.Type))
	case "A":
		for _, ar := range *r.RecordSetProperties.ARecords {
			endpoints = append(endpoints, endpoint.NewEndpoint(*r.Name, *ar.Ipv4Address, *r.Type))
		}
	case "TXT":
		for _, tr := range *r.RecordSetProperties.TxtRecords {
			endpoints = append(endpoints, endpoint.NewEndpoint(*r.Name, strings.Join(*tr.Value, " "), *r.Type))
		}
	default:
		return nil, fmt.Errorf("Invalid record type. Only CNAME, A and TXT are currently supported.")
	}

	return endpoints, nil
}

func RecordSetListToEndpoints(records []dns.RecordSet) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	for _, r := range records {
		moreEndpoints, err := RecordSetToEndpoints(r)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, moreEndpoints...)
	}
	return endpoints, nil
}


func (p *azureProvider) List(zone string) ([]dns.RecordSet, error) {
	var records []dns.RecordSet

	// Get first 100 records
	res, err := p.recordSetsClient.ListByDNSZone(p.resourceGroupName, zone, nil)
	if err != nil {
		return nil, err
	}
	records = append(records, *res.Value...)
	// Get following records by 100
	for err != nil {
		res, err = p.recordSetsClient.ListByDNSZone(p.resourceGroupName, zone, nil)
		if err == nil {
			records = append(records, *res.Value...)
		}
	}

	return records, nil
}

func (p *azureProvider) Create(zone string, e *endpoint.Endpoint) error {
	a := azureEndpoint(*e)
	r, err :=  a.toRecordSet()
	if err != nil {
		return err
	}
	_, err = p.recordSetsClient.CreateOrUpdate(p.resourceGroupName, zone, e.DNSName, dns.RecordType(e.RecordType), r, "", "")
	if err != nil {
		return err
	}
	log.Infof("Created record: %s %s", zone, e.DNSName)
	return nil
}

func (p *azureProvider) Update(zone string, oe, ne *endpoint.Endpoint) error {
	r, err := p.recordSetsClient.Get(p.resourceGroupName, zone, oe.DNSName, dns.RecordType(oe.RecordType))
	if err != nil {
		return err
	}

	props, err := azureEndpoint(oe).getRecordSetProperties()
	if err != nil {
		return err
	}

	r.Name = &ne.DNSName
	r.Type = &ne.RecordType
	r.RecordSetProperties = props

	_, err = p.recordSetsClient.CreateOrUpdate(p.resourceGroupName, zone, ne.DNSName, dns.RecordType(ne.RecordType), r, "", "")
	if err != nil {
		return err
	}
	log.Infof("Updated record: %s %s", zone, oe.DNSName)
	return nil
}

func (p *azureProvider) Delete(zone string, e *endpoint.Endpoint) error {
	_, err := p.recordSetsClient.Delete(p.resourceGroupName, zone, e.DNSName, dns.RecordType(e.RecordType), "")
	if err != nil {
		return err
	}
	log.Infof("Deleted record: %s %s", zone, e.DNSName)
	return nil
}

// Records returns the list of records in the relevant zone.
func (p *azureProvider) Records(zone string) ([]*endpoint.Endpoint, error) {
	records, err := p.List(zone)
	if err != nil {
		return nil, err
	}
	return RecordSetListToEndpoints(records)
}

func (p *azureProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	var err error

	// Create new DNS records
	for _, e := range changes.Create {
		err = p.Create(zone, e)
		if err != nil {
			return err
		}
	}

	// Update existing DNS records
	for i, ne := range changes.UpdateNew {
		err := p.Update(zone, changes.UpdateOld[i], ne)
		if err != nil {
			return err
		}
	}

	// Delete existing DNS records
	for _, e := range changes.Delete {
		err := p.Delete(zone, e)
		if err != nil {
			return err
		}
	}

	return nil
}
