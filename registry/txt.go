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

package registry

import (
	"errors"

	"strings"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
)

// TXTRegistry implements registry interface with ownership implemented via associated TXT records
type TXTRegistry struct {
	provider provider.Provider
	ownerID  string //refers to the owner id of the current instance
	mapper   nameMapper
}

// NewTXTRegistry returns new TXTRegistry object
func NewTXTRegistry(provider provider.Provider, txtPrefix, ownerID string) (*TXTRegistry, error) {
	if ownerID == "" {
		return nil, errors.New("owner id cannot be empty")
	}

	mapper := newPrefixNameMapper(txtPrefix)

	return &TXTRegistry{
		provider: provider,
		ownerID:  ownerID,
		mapper:   mapper,
	}, nil
}

// Records returns the current records from the registry excluding TXT Records
// If TXT records was created previously to indicate ownership its corresponding value
// will be added to the endpoints Labels map
func (im *TXTRegistry) Records() ([]*endpoint.Endpoint, error) {
	records, err := im.provider.Records()
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	labelMap := map[string]endpoint.Labels{}

	for _, record := range records {
		if record.RecordType != endpoint.RecordTypeTXT {
			endpoints = append(endpoints, record)
			continue
		}
		labels, err := endpoint.NewLabelsFromString(record.Target)
		if err == endpoint.ErrInvalidHeritage {
			//if no heritage is found or it is invalid
			//case when value of txt record cannot be identified
			//record will not be removed as it will have empty owner
			endpoints = append(endpoints, record)
			continue
		}
		if err != nil {
			return nil, err
		}
		endpointDNSName := im.mapper.toEndpointName(record.DNSName)
		labelMap[endpointDNSName] = labels
	}

	for _, ep := range endpoints {
		if labels, ok := labelMap[ep.DNSName]; ok {
			ep.Labels = labels
		} else {
			//this indicates that owner could not be identified, as there is no corresponding TXT record
			ep.Labels = endpoint.NewLabels()
		}
	}

	return endpoints, nil
}

// ApplyChanges updates dns provider with the changes
// for each created/deleted record it will also take into account TXT records for creation/deletion
func (im *TXTRegistry) ApplyChanges(changes *plan.Changes) error {
	filteredChanges := &plan.Changes{
		Create:    changes.Create,
		UpdateNew: filterOwnedRecords(im.ownerID, changes.UpdateNew),
		UpdateOld: filterOwnedRecords(im.ownerID, changes.UpdateOld),
		Delete:    filterOwnedRecords(im.ownerID, changes.Delete),
	}
	for _, r := range filteredChanges.Create {
		r.Labels[endpoint.OwnerLabelKey] = im.ownerID
		txt := endpoint.NewEndpoint(im.mapper.toTXTName(r.DNSName), r.Labels.Serialize(true), endpoint.RecordTypeTXT)
		filteredChanges.Create = append(filteredChanges.Create, txt)
	}

	for _, r := range filteredChanges.Delete {
		txt := endpoint.NewEndpoint(im.mapper.toTXTName(r.DNSName), r.Labels.Serialize(true), endpoint.RecordTypeTXT)

		// when we delete TXT records for which value has changed (due to new label) this would still work because
		// !!! TXT record value is uniquely generated from the Labels of the endpoint. Hence old TXT record can be uniquely reconstructed
		filteredChanges.Delete = append(filteredChanges.Delete, txt)
	}

	// make sure TXT records are consistently updated as well
	for _, r := range filteredChanges.UpdateNew {
		txt := endpoint.NewEndpoint(im.mapper.toTXTName(r.DNSName), r.Labels.Serialize(true), endpoint.RecordTypeTXT)
		filteredChanges.UpdateNew = append(filteredChanges.UpdateNew, txt)
	}
	// make sure TXT records are consistently updated as well
	for _, r := range filteredChanges.UpdateOld {
		txt := endpoint.NewEndpoint(im.mapper.toTXTName(r.DNSName), r.Labels.Serialize(true), endpoint.RecordTypeTXT)
		// when we updateOld TXT records for which value has changed (due to new label) this would still work because
		// !!! TXT record value is uniquely generated from the Labels of the endpoint. Hence old TXT record can be uniquely reconstructed
		filteredChanges.UpdateOld = append(filteredChanges.UpdateOld, txt)
	}

	return im.provider.ApplyChanges(filteredChanges)
}

/**
  TXT registry specific private methods
*/

/**
  nameMapper defines interface which maps the dns name defined for the source
  to the dns name which TXT record will be created with
*/

type nameMapper interface {
	toEndpointName(string) string
	toTXTName(string) string
}

type prefixNameMapper struct {
	prefix string
}

var _ nameMapper = prefixNameMapper{}

func newPrefixNameMapper(prefix string) prefixNameMapper {
	return prefixNameMapper{prefix: prefix}
}

func (pr prefixNameMapper) toEndpointName(txtDNSName string) string {
	if strings.HasPrefix(txtDNSName, pr.prefix) {
		return strings.TrimPrefix(txtDNSName, pr.prefix)
	}
	return ""
}

func (pr prefixNameMapper) toTXTName(endpointDNSName string) string {
	return pr.prefix + endpointDNSName
}
