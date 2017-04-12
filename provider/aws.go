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
	"net"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// Route53API is the subset of the AWS Route53 API that we actually use.  Add methods as required. Signatures must match exactly.
// mostly taken from: https://github.com/kubernetes/kubernetes/blob/853167624edb6bc0cfdcdfb88e746e178f5db36c/federation/pkg/dnsprovider/providers/aws/route53/stubs/route53api.go
type Route53API interface {
	ListResourceRecordSetsPages(input *route53.ListResourceRecordSetsInput, fn func(resp *route53.ListResourceRecordSetsOutput, lastPage bool) (shouldContinue bool)) error
	ChangeResourceRecordSets(*route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)
	CreateHostedZone(*route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)
	ListHostedZonesPages(input *route53.ListHostedZonesInput, fn func(resp *route53.ListHostedZonesOutput, lastPage bool) (shouldContinue bool)) error
}

// AWSProvider is an implementation of Provider for AWS Route53.
type AWSProvider struct {
	Client Route53API
	DryRun bool
	// only consider hosted zones managing domains ending in this suffix
	Domain string
}

// NewAWSProvider initializes a new AWS Route53 based Provider.
func NewAWSProvider(domain string, dryRun bool) (Provider, error) {
	config := aws.NewConfig()

	session, err := session.NewSessionWithOptions(session.Options{
		Config:            *config,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	provider := &AWSProvider{
		Client: route53.New(session),
		Domain: domain,
		DryRun: dryRun,
	}

	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *AWSProvider) Zones() (map[string]*route53.HostedZone, error) {
	zones := make(map[string]*route53.HostedZone)

	f := func(resp *route53.ListHostedZonesOutput, lastPage bool) (shouldContinue bool) {
		for _, zone := range resp.HostedZones {
			if strings.HasSuffix(aws.StringValue(zone.Name), p.Domain) {
				zones[aws.StringValue(zone.Id)] = zone
			}
		}

		return true
	}

	err := p.Client.ListHostedZonesPages(&route53.ListHostedZonesInput{}, f)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

// Records returns the list of records in a given hosted zone.
func (p *AWSProvider) Records(_ string) (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.Zones()
	if err != nil {
		return nil, err
	}

	f := func(resp *route53.ListResourceRecordSetsOutput, lastPage bool) (shouldContinue bool) {
		for _, r := range resp.ResourceRecordSets {
			// TODO(linki, ownership): Remove once ownership system is in place.
			// See: https://github.com/kubernetes-incubator/external-dns/pull/122/files/74e2c3d3e237411e619aefc5aab694742001cdec#r109863370
			switch aws.StringValue(r.Type) {
			case route53.RRTypeA, route53.RRTypeCname, route53.RRTypeTxt:
				break
			default:
				continue
			}

			for _, rr := range r.ResourceRecords {
				endpoints = append(endpoints, endpoint.NewEndpoint(aws.StringValue(r.Name), aws.StringValue(rr.Value), aws.StringValue(r.Type)))
			}
		}

		return true
	}

	for _, z := range zones {
		params := &route53.ListResourceRecordSetsInput{
			HostedZoneId: z.Id,
		}

		if err := p.Client.ListResourceRecordSetsPages(params, f); err != nil {
			return nil, err
		}
	}

	return endpoints, nil
}

// CreateRecords creates a given set of DNS records in the given hosted zone.
func (p *AWSProvider) CreateRecords(endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(newChanges(route53.ChangeActionCreate, endpoints))
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *AWSProvider) UpdateRecords(endpoints, _ []*endpoint.Endpoint) error {
	return p.submitChanges(newChanges(route53.ChangeActionUpsert, endpoints))
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *AWSProvider) DeleteRecords(endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(newChanges(route53.ChangeActionDelete, endpoints))
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *AWSProvider) ApplyChanges(_ string, changes *plan.Changes) error {
	combinedChanges := make([]*route53.Change, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newChanges(route53.ChangeActionCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newChanges(route53.ChangeActionUpsert, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newChanges(route53.ChangeActionDelete, changes.Delete)...)

	return p.submitChanges(combinedChanges)
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *AWSProvider) submitChanges(changes []*route53.Change) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	if p.DryRun {
		for _, change := range changes {
			log.Infof("Changing records: %s %s", aws.StringValue(change.Action), change.String())
		}

		return nil
	}

	zones, err := p.Zones()
	if err != nil {
		return err
	}

	// separate into per-zone change sets to be passed to the API.
	changesByZone := changesByZone(zones, changes)

	for z, cs := range changesByZone {
		params := &route53.ChangeResourceRecordSetsInput{
			HostedZoneId: aws.String(z),
			ChangeBatch: &route53.ChangeBatch{
				Changes: cs,
			},
		}

		if _, err := p.Client.ChangeResourceRecordSets(params); err != nil {
			return err
		}
	}

	return nil
}

// changesByZone separates a multi-zone change into a single change per zone.
func changesByZone(zones map[string]*route53.HostedZone, changeSet []*route53.Change) map[string][]*route53.Change {
	changes := make(map[string][]*route53.Change)

	for _, z := range zones {
		changes[aws.StringValue(z.Id)] = []*route53.Change{}

		for _, c := range changeSet {
			if strings.HasSuffix(ensureTrailingDot(aws.StringValue(c.ResourceRecordSet.Name)), aws.StringValue(z.Name)) {
				changes[aws.StringValue(z.Id)] = append(changes[aws.StringValue(z.Id)], c)
			}
		}
	}

	// separating a change could lead to empty sub changes, remove them here.
	for zone, change := range changes {
		if len(change) == 0 {
			delete(changes, zone)
		}
	}

	return changes
}

// newChanges returns a collection of Changes based on the given records and action.
func newChanges(action string, endpoints []*endpoint.Endpoint) []*route53.Change {
	changes := make([]*route53.Change, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, newChange(action, endpoint))
	}

	return changes
}

// newChange returns a Change of the given record by the given action, e.g.
// action=ChangeActionCreate returns a change for creation of the record and
// action=ChangeActionDelete returns a change for deletion of the record.
func newChange(action string, endpoint *endpoint.Endpoint) *route53.Change {
	change := &route53.Change{
		Action: aws.String(action),
		ResourceRecordSet: &route53.ResourceRecordSet{
			Name: aws.String(endpoint.DNSName),
			ResourceRecords: []*route53.ResourceRecord{
				{
					Value: aws.String(endpoint.Target),
				},
			},
			TTL:  aws.Int64(300),
			Type: aws.String(suitableType(endpoint)),
		},
	}

	return change
}

// ensureTrailingDot ensures that the hostname receives a trailing dot if it hasn't already.
func ensureTrailingDot(hostname string) string {
	if net.ParseIP(hostname) != nil {
		return hostname
	}

	return strings.TrimSuffix(hostname, ".") + "."
}
