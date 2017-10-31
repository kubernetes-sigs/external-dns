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
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
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
	ownerMap map[string]string // map DNSName => ownerID
	recCount map[string]int    // map DNSName => number of records
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

	im.ownerMap = make(map[string]string)
	im.recCount = make(map[string]int)

	for _, record := range records {
		if record.RecordType != endpoint.RecordTypeTXT {
			endpoints = append(endpoints, record)
			im.recCount[record.DNSName]++
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
		im.ownerMap[endpointDNSName] = ownerID
	}

	log.Debugf("after scanning: ownerMap: %s, recCount: %s", im.ownerMap, im.recCount)

	return endpoints, err
}

// ApplyChanges updates dns provider with the changes
// for each created/deleted record it will also take into account TXT records for creation/deletion
func (im *TXTRegistry) ApplyChanges(changes *plan.Changes) error {
	if im.ownerMap == nil {
		// Records() never called yet or error during last ApplyChanges() call => rescan
		_, err := im.Records()
		if err != nil {
			return err
		}
	}

	filteredChanges := &plan.Changes{
		Create:    []*endpoint.Endpoint{},
		UpdateNew: im.filterOwnedRecords(changes.UpdateNew),
		UpdateOld: im.filterOwnedRecords(changes.UpdateOld),
		Delete:    im.filterOwnedRecords(changes.Delete),
	}

	for _, r := range changes.Create {
		rOwner := im.ownerMap[r.DNSName]

		if rOwner == "" || rOwner == im.ownerID {
			filteredChanges.Create = append(filteredChanges.Create, r)

			im.ownerMap[r.DNSName] = im.ownerID
			im.recCount[r.DNSName]++

			if 1 == im.recCount[r.DNSName] {
				txt := endpoint.NewEndpoint(im.mapper.toTXTName(r.DNSName), im.getTXTLabel(), endpoint.RecordTypeTXT)
				filteredChanges.Create = append(filteredChanges.Create, txt)
			}
		}
	}

	for _, r := range filteredChanges.Delete {
		if im.recCount[r.DNSName]--; im.recCount[r.DNSName] <= 0 {
			im.recCount[r.DNSName] = 0
			delete(im.ownerMap, r.DNSName)
			txt := endpoint.NewEndpoint(im.mapper.toTXTName(r.DNSName), im.getTXTLabel(), endpoint.RecordTypeTXT)
			filteredChanges.Delete = append(filteredChanges.Delete, txt)
		}
	}

	log.Debugf("before provider.ApplyChanges: ownerMap: %s, recCount: %s", im.ownerMap, im.recCount)

	err := im.provider.ApplyChanges(filteredChanges)
	if err != nil {
		// error occured in the provider => we don't know which records were stored => force re-scanning on the next call
		im.ownerMap = nil
	}
	return err
}

func (im *TXTRegistry) filterOwnedRecords(eps []*endpoint.Endpoint) []*endpoint.Endpoint {
	filtered := []*endpoint.Endpoint{}
	for _, ep := range eps {
		if endpointOwner, ok := im.ownerMap[ep.DNSName]; !ok || endpointOwner != im.ownerID {
			log.Debugf(`Skipping endpoint %v because owner id does not match, found: "%s", required: "%s"`, ep, endpointOwner, im.ownerID)
			continue
		}
		filtered = append(filtered, ep)
	}
	return filtered
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
