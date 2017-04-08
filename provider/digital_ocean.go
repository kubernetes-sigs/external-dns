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
	"context"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/oauth2"

	"github.com/digitalocean/godo"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// DigitalOceanProvider is an implementation of Provider for Digital Ocean's DNS.
type DigitalOceanProvider struct {
	Client godo.DomainsService
	DryRun bool
}

// TokenSource inherits AccessToken to initialize a new Digital Ocean client.
type TokenSource struct {
	AccessToken string
}

// DigitalOceanChange differentiates between ChangActions
type DigitalOceanChange struct {
	Action            string
	ResourceRecordSet godo.DomainRecord
}

const (
	// DigitalOceanCreate is a ChangeAction enum value
	DigitalOceanCreate = "CREATE"
	// DigitalOceanDelete is a ChangeAction enum value
	DigitalOceanDelete = "DELETE"
	// DigitalOceanUpdate is a ChangeAction enum value
	DigitalOceanUpdate = "UPDATE"
)

// Token returns oauth2 token struct in order to use Digital Ocean's API.
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// NewDigitalOceanProvider initializes a new DigitalOcean DNS based Provider.
func NewDigitalOceanProvider(dryRun bool) (Provider, error) {
	token := os.Getenv("DO_TOKEN")
	tokenSource := &TokenSource{
		AccessToken: token,
	}
	if len(token) == 0 {
		return nil, fmt.Errorf("No token found")
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)

	provider := &DigitalOceanProvider{
		Client: client.Domains,
		DryRun: dryRun,
	}
	return provider, nil
}

// Zones returns the list of zones.
func (p *DigitalOceanProvider) Zones() (*[]godo.Domain, error) {
	zones, _, err := p.Client.List(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return &zones, nil
}

// Zone returns a single zone given a DNS name.
func (p *DigitalOceanProvider) Zone(zoneName string) (*godo.Domain, error) {
	zone, _, err := p.Client.Get(context.TODO(), zoneName)
	if err != nil {
		return nil, err
	}
	return zone, nil
}

// CreateZone creates a hosted zone given a name.
func (p *DigitalOceanProvider) CreateZone(name, ip string) (*godo.Domain, error) {
	newDomain := &godo.DomainCreateRequest{
		Name:      name,
		IPAddress: ip,
	}
	zone, _, err := p.Client.Create(context.TODO(), newDomain)
	if err != nil {
		return nil, err
	}
	return zone, nil
}

// DeleteZone deletes a hosted zone given a name.
func (p *DigitalOceanProvider) DeleteZone(name string) (string, error) {
	_, err := p.Client.Delete(context.TODO(), name)
	if err != nil {
		return name, err
	}
	return name, nil
}

// Records returns the list of records in a given zone.
func (p *DigitalOceanProvider) Records(zone string) ([]*endpoint.Endpoint, error) {
	records, _, err := p.Client.Records(context.TODO(), zone, nil)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}
	for _, r := range records {
		endpoints = append(endpoints, endpoint.NewEndpoint(r.Name, r.Data, r.Type))
	}

	return endpoints, nil
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *DigitalOceanProvider) submitChanges(zone string, changes []*DigitalOceanChange) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}
	if p.DryRun {
		for _, change := range changes {
			log.Infof("Changing records: %s %s", change.Action, change.ResourceRecordSet.String())
		}

		return nil
	}
	for _, change := range changes {
		switch {
		case change.Action == DigitalOceanCreate:
			_, _, err := p.Client.CreateRecord(context.TODO(), zone,
				&godo.DomainRecordEditRequest{
					Data: change.ResourceRecordSet.Data,
					Name: change.ResourceRecordSet.Name,
					Type: change.ResourceRecordSet.Type,
				})
			if err != nil {
				return err
			}
		case change.Action == DigitalOceanDelete:
			recordID, err := p.getRecordID(zone, change.ResourceRecordSet)
			if err != nil {
				return err
			}
			_, err = p.Client.DeleteRecord(context.TODO(), zone, recordID)
			if err != nil {
				return err
			}
		case change.Action == DigitalOceanUpdate:
			recordID, err := p.getRecordID(zone, change.ResourceRecordSet)
			if err != nil {
				return err
			}
			_, _, err = p.Client.EditRecord(context.TODO(), zone, recordID,
				&godo.DomainRecordEditRequest{
					Data: change.ResourceRecordSet.Data,
					Name: change.ResourceRecordSet.Name,
					Type: change.ResourceRecordSet.Type,
				})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CreateRecords creates a given set of DNS records in the given hosted zone.
func (p *DigitalOceanProvider) CreateRecords(zone string, endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(zone, newDigitalOceanChanges(DigitalOceanCreate, endpoints))
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *DigitalOceanProvider) UpdateRecords(zone string, endpoints, _ []*endpoint.Endpoint) error {
	return p.submitChanges(zone, newDigitalOceanChanges(DigitalOceanUpdate, endpoints))
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *DigitalOceanProvider) DeleteRecords(zone string, endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(zone, newDigitalOceanChanges(DigitalOceanDelete, endpoints))
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *DigitalOceanProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	combinedChanges := make([]*DigitalOceanChange, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newDigitalOceanChanges(DigitalOceanCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newDigitalOceanChanges(DigitalOceanUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newDigitalOceanChanges(DigitalOceanDelete, changes.Delete)...)

	return p.submitChanges(zone, combinedChanges)
}

// newDigitalOceanChanges returns a collection of Changes based on the given records and action.
func newDigitalOceanChanges(action string, endpoints []*endpoint.Endpoint) []*DigitalOceanChange {
	changes := make([]*DigitalOceanChange, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, newDigitalOceanChange(action, endpoint))
	}

	return changes
}

func newDigitalOceanChange(action string, endpoint *endpoint.Endpoint) *DigitalOceanChange {
	change := &DigitalOceanChange{
		Action: action,
		ResourceRecordSet: godo.DomainRecord{
			Name: endpoint.DNSName,
			Type: suitableType(endpoint),
			Data: ensureTrailingDot(endpoint.Target),
		},
	}
	return change
}

// getRecordID returns the ID from a record.
// the ID is mandatory to update and delete records
func (p *DigitalOceanProvider) getRecordID(zone string, record godo.DomainRecord) (int, error) {
	zoneRecords, _, err := p.Client.Records(context.TODO(), zone, nil)
	if err != nil {
		return 0, err
	}
	for _, zoneRecord := range zoneRecords {
		if zoneRecord.Name == record.Name {
			return zoneRecord.ID, nil
		}
	}
	return 0, fmt.Errorf("No record id found")
}
