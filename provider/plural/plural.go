/*
Copyright 2022 The Kubernetes Authors.

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

package plural

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	CreateAction = "c"
	DeleteAction = "d"
)

type PluralProvider struct {
	provider.BaseProvider
	Client Client
}

type RecordChange struct {
	Action string
	Record *DnsRecord
}

func NewPluralProvider(cluster, provider string) (*PluralProvider, error) {
	token := os.Getenv("PLURAL_ACCESS_TOKEN")
	endpoint := os.Getenv("PLURAL_ENDPOINT")

	if token == "" {
		return nil, fmt.Errorf("No plural access token provided, you must set the PLURAL_ACCESS_TOKEN env var")
	}

	config := &Config{
		Token:    token,
		Endpoint: endpoint,
		Cluster:  cluster,
		Provider: provider,
	}

	client, err := NewClient(config)
	if err != nil {
		return nil, err
	}

	prov := &PluralProvider{
		Client: client,
	}

	return prov, nil
}

func (p *PluralProvider) Records(_ context.Context) (endpoints []*endpoint.Endpoint, err error) {
	records, err := p.Client.DnsRecords()
	if err != nil {
		return
	}

	endpoints = make([]*endpoint.Endpoint, len(records))
	for i, record := range records {
		endpoints[i] = endpoint.NewEndpoint(record.Name, record.Type, record.Records...)
	}
	return
}

func (p *PluralProvider) PropertyValuesEqual(name, previous, current string) bool {
	return p.BaseProvider.PropertyValuesEqual(name, previous, current)
}

func (p *PluralProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	return endpoints
}

func (p *PluralProvider) ApplyChanges(_ context.Context, diffs *plan.Changes) error {
	var changes []*RecordChange
	for _, endpoint := range diffs.Create {
		changes = append(changes, makeChange(CreateAction, endpoint.Targets, endpoint))
	}

	for _, desired := range diffs.UpdateNew {
		changes = append(changes, makeChange(CreateAction, desired.Targets, desired))
	}

	for _, deleted := range diffs.Delete {
		changes = append(changes, makeChange(DeleteAction, []string{}, deleted))
	}

	return p.applyChanges(changes)
}

func makeChange(change string, target []string, endpoint *endpoint.Endpoint) *RecordChange {
	return &RecordChange{
		Action: change,
		Record: &DnsRecord{
			Name:    endpoint.DNSName,
			Type:    endpoint.RecordType,
			Records: target,
		},
	}
}

func (p *PluralProvider) applyChanges(changes []*RecordChange) error {
	for _, change := range changes {
		logFields := log.Fields{
			"name":   change.Record.Name,
			"type":   change.Record.Type,
			"action": change.Action,
		}
		log.WithFields(logFields).Info("Changing record.")

		if change.Action == CreateAction {
			_, err := p.Client.CreateRecord(change.Record)
			if err != nil {
				return err
			}
		}
		if change.Action == DeleteAction {
			if err := p.Client.DeleteRecord(change.Record.Name, change.Record.Type); err != nil {
				return err
			}
		}
	}

	return nil
}
