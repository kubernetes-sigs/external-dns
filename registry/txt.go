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

	"fmt"
	"regexp"
	"strings"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
	"github.com/kubernetes-incubator/external-dns/source"
)

var (
	txtLabelRegex  = regexp.MustCompile("^\"heritage=external-dns,external-dns/owner=(.+)\"")
	txtLabelFormat = "\"heritage=external-dns,external-dns/owner=%s\""
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

	endpoints := make([]*endpoint.Endpoint, 0)

	ownerMap := map[string]string{}

	for _, record := range records {
		if record.RecordType != endpoint.RecordTypeTXT {
			endpoints = append(endpoints, record)
			continue
		}
		ownerID := im.extractOwnerID(record.Target)
		if ownerID == "" {
			//case when value of txt record cannot be identified
			//record will not be removed as it will have empty owner
			endpoints = append(endpoints, record)
			continue
		}
		endpointDNSName := im.mapper.toEndpointName(record.DNSName)
		ownerMap[endpointDNSName] = ownerID
	}

	for _, ep := range endpoints {
		ep.Labels[endpoint.OwnerLabelKey] = ownerMap[ep.DNSName]
	}

	return endpoints, err
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

	txtChanges := map[string][]*endpoint.Endpoint{
		"create": {},
		"delete": {},
	}

	for _, r := range filteredChanges.Create {
		// If an endpoint already has a TXT record, do not add one
		if r.Labels[endpoint.TxtOwnedLabelKey] == "true" {
			continue
		}
		txt := endpoint.NewEndpoint(im.mapper.toTXTName(r.DNSName), im.getTXTLabel(), endpoint.RecordTypeTXT)
		txtChanges["create"] = append(txtChanges["create"], txt)
	}

	for _, r := range filteredChanges.Delete {
		// If there are multiple endpoints under a single TXT record, do not delete it
		if r.Labels[endpoint.TxtOwnedLabelKey] == "true" {
			continue
		}
		txt := endpoint.NewEndpoint(im.mapper.toTXTName(r.DNSName), im.getTXTLabel(), endpoint.RecordTypeTXT)
		txtChanges["delete"] = append(txtChanges["delete"], txt)

	}

	txtSourceCreate := source.NewDedupSource(source.NewRawSource(txtChanges["create"]))
	if endpoints, err := txtSourceCreate.Endpoints(); err == nil {
		filteredChanges.Create = append(filteredChanges.Create, endpoints...)
	}

	txtSourceDelete := source.NewDedupSource(source.NewRawSource(txtChanges["delete"]))
	if endpoints, err := txtSourceDelete.Endpoints(); err == nil {
		filteredChanges.Delete = append(filteredChanges.Delete, endpoints...)
	}

	return im.provider.ApplyChanges(filteredChanges)
}

/**
  TXT registry specific private methods
*/

func (im *TXTRegistry) getTXTLabel() string {
	return fmt.Sprintf(txtLabelFormat, im.ownerID)
}

func (im *TXTRegistry) extractOwnerID(txtLabel string) string {
	if matches := txtLabelRegex.FindStringSubmatch(txtLabel); len(matches) == 2 {
		return matches[1]
	}
	return ""
}

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
