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

package dnsprovider

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// Route53API is the subset of the AWS Route53 API that we actually use.  Add methods as required. Signatures must match exactly.
// mostly taken from: https://github.com/kubernetes/kubernetes/blob/853167624edb6bc0cfdcdfb88e746e178f5db36c/federation/pkg/dnsprovider/providers/aws/route53/stubs/route53api.go
type Route53API interface {
	ListResourceRecordSetsPages(input *route53.ListResourceRecordSetsInput, fn func(resp *route53.ListResourceRecordSetsOutput, lastPage bool) (shouldContinue bool)) error
	ChangeResourceRecordSets(*route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)
	ListHostedZonesPages(input *route53.ListHostedZonesInput, fn func(resp *route53.ListHostedZonesOutput, lastPage bool) (shouldContinue bool)) error
	ListHostedZonesByName(input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error)
	CreateHostedZone(*route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)
	DeleteHostedZone(*route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error)
}

// AWSProvider is an implementation of DNSProvider for AWS Route53.
type AWSProvider struct {
	Client Route53API
	DryRun bool
}

// Zones returns the list of hosted zones.
func (p *AWSProvider) Zones() ([]string, error) {
	zones := []string{}

	f := func(resp *route53.ListHostedZonesOutput, lastPage bool) (shouldContinue bool) {
		for _, zone := range resp.HostedZones {
			zones = append(zones, *zone.Name)
		}

		return true
	}

	err := p.Client.ListHostedZonesPages(&route53.ListHostedZonesInput{}, f)
	if err != nil {
		return zones, err
	}

	return zones, nil
}

// Zone returns a single zone given a DNS name.
func (p *AWSProvider) Zone(dnsName string) (*route53.HostedZone, error) {
	params := &route53.ListHostedZonesByNameInput{
		DNSName: aws.String(dnsName),
	}

	resp, err := p.Client.ListHostedZonesByName(params)
	if err != nil {
		return nil, err
	}

	if len(resp.HostedZones) != 1 {
		return nil, fmt.Errorf("not exactly one hosted zone found by name, got %d", len(resp.HostedZones))
	}

	return resp.HostedZones[0], nil
}

// CreateZone creates a hosted zone given a name.
func (p *AWSProvider) CreateZone(name string) (*route53.HostedZone, error) {
	params := &route53.CreateHostedZoneInput{
		CallerReference: aws.String(uuid.New().String()),
		Name:            aws.String(name),
	}

	resp, err := p.Client.CreateHostedZone(params)
	if err != nil {
		return nil, err
	}

	return resp.HostedZone, nil
}

// DeleteZone deletes a hosted zone given a name.
func (p *AWSProvider) DeleteZone(name string) error {
	params := &route53.DeleteHostedZoneInput{
		Id: aws.String(name),
	}

	_, err := p.Client.DeleteHostedZone(params)
	if err != nil {
		return err
	}

	return nil
}

// Records returns the list of records in a given hosted zone.
func (p *AWSProvider) Records(zone string) ([]endpoint.Endpoint, error) {
	hostedZone, err := p.Zone(zone)
	if err != nil {
		return nil, err
	}

	params := &route53.ListResourceRecordSetsInput{
		HostedZoneId: hostedZone.Id,
	}

	endpoints := []endpoint.Endpoint{}

	f := func(resp *route53.ListResourceRecordSetsOutput, lastPage bool) (shouldContinue bool) {
		for _, r := range resp.ResourceRecordSets {
			if *r.Type != route53.RRTypeA {
				continue
			}

			for _, rr := range r.ResourceRecords {
				endpoint := endpoint.Endpoint{
					DNSName: *r.Name,
					Target:  *rr.Value,
				}

				endpoints = append(endpoints, endpoint)
			}
		}

		return true
	}

	err = p.Client.ListResourceRecordSetsPages(params, f)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

// CreateRecords creates a given set of DNS records in the given hosted zone.
func (p *AWSProvider) CreateRecords(zone string, records []endpoint.Endpoint) error {
	hostedZone, err := p.Zone(zone)
	if err != nil {
		return err
	}

	changes := []*route53.Change{}

	for _, record := range records {
		change := &route53.Change{
			Action: aws.String(route53.ChangeActionCreate),
			ResourceRecordSet: &route53.ResourceRecordSet{
				Name: aws.String(record.DNSName),
				ResourceRecords: []*route53.ResourceRecord{
					{
						Value: aws.String(record.Target),
					},
				},
				TTL:  aws.Int64(300),
				Type: aws.String(route53.RRTypeA),
			},
		}

		changes = append(changes, change)
	}

	params := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: hostedZone.Id,
		ChangeBatch: &route53.ChangeBatch{
			Changes: changes,
		},
	}

	if p.DryRun {
		log.Infof("Creating records: %#v", params.ChangeBatch.Changes)
		return nil
	}

	_, err = p.Client.ChangeResourceRecordSets(params)
	if err != nil {
		return err
	}

	return nil
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *AWSProvider) UpdateRecords(zone string, newRecords, _ []endpoint.Endpoint) error {
	hostedZone, err := p.Zone(zone)
	if err != nil {
		return err
	}

	changes := []*route53.Change{}

	for _, record := range newRecords {
		change := &route53.Change{
			Action: aws.String(route53.ChangeActionUpsert),
			ResourceRecordSet: &route53.ResourceRecordSet{
				Name: aws.String(record.DNSName),
				ResourceRecords: []*route53.ResourceRecord{
					{
						Value: aws.String(record.Target),
					},
				},
				TTL:  aws.Int64(300),
				Type: aws.String(route53.RRTypeA),
			},
		}

		changes = append(changes, change)
	}

	params := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: hostedZone.Id,
		ChangeBatch: &route53.ChangeBatch{
			Changes: changes,
		},
	}

	if p.DryRun {
		log.Infof("Updating records: %#v", params.ChangeBatch.Changes)
		return nil
	}

	_, err = p.Client.ChangeResourceRecordSets(params)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *AWSProvider) DeleteRecords(zone string, records []endpoint.Endpoint) error {
	hostedZone, err := p.Zone(zone)
	if err != nil {
		return err
	}

	changes := []*route53.Change{}

	for _, record := range records {
		change := &route53.Change{
			Action: aws.String(route53.ChangeActionDelete),
			ResourceRecordSet: &route53.ResourceRecordSet{
				Name: aws.String(record.DNSName),
				ResourceRecords: []*route53.ResourceRecord{
					{
						Value: aws.String(record.Target),
					},
				},
				TTL:  aws.Int64(300),
				Type: aws.String(route53.RRTypeA),
			},
		}

		changes = append(changes, change)
	}

	params := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: hostedZone.Id,
		ChangeBatch: &route53.ChangeBatch{
			Changes: changes,
		},
	}

	if p.DryRun {
		log.Infof("Deleting records: %#v", params.ChangeBatch.Changes)
		return nil
	}

	_, err = p.Client.ChangeResourceRecordSets(params)
	if err != nil {
		return err
	}

	return nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *AWSProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	err := p.CreateRecords(zone, changes.Create)
	if err != nil {
		return err
	}

	err = p.UpdateRecords(zone, changes.UpdateNew, changes.UpdateOld)
	if err != nil {
		return err
	}

	err = p.DeleteRecords(zone, changes.Delete)
	if err != nil {
		return err
	}

	return nil
}
