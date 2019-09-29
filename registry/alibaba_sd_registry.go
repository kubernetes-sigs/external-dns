/*
Copyright 2018 The Kubernetes Authors.

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
	"context"
	"errors"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
	"strings"
)

// AWSSDRegistry implements registry interface with ownership information associated via the Description field of SD Service
type AlibabaCloudSDRegistry struct {
	provider provider.Provider
	ownerID  string
	prefix   string
}

// NewAWSSDRegistry returns implementation of registry for AWS SD
func NewAlibabaCloudSDRegistry(provider provider.Provider, prefix string, ownerID string) (*AlibabaCloudSDRegistry, error) {
	if ownerID == "" {
		return nil, errors.New("owner id cannot be empty")
	}
	return &AlibabaCloudSDRegistry{
		provider: provider,
		ownerID:  ownerID,
		prefix:   prefix,
	}, nil
}

// Records calls AWS SD API and expects AWS SD provider to provider Owner/Resource information as a serialized
// value in the AWSSDDescriptionLabel value in the Labels map
func (adr *AlibabaCloudSDRegistry) Records() ([]*endpoint.Endpoint, error) {
	records, err := adr.provider.Records()
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
		// We simply assume that TXT records for the registry will always have only one target.
		labels, err := endpoint.NewLabelsFromString(record.Targets[0])
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
		endpointDNSName := adr.toEndpointName(record.DNSName)
		labelMap[endpointDNSName] = labels
	}
	for _, ep := range endpoints {
		if ep.Labels == nil {
			ep.Labels = endpoint.NewLabels()
		}
		txtKey := adr.toTxtKey(ep)
		if labels, ok := labelMap[txtKey]; ok {
			for k, v := range labels {
				ep.Labels[k] = v
			}
		}
	}
	return endpoints, nil
}

func (adr *AlibabaCloudSDRegistry) toTxtKey(ep *endpoint.Endpoint) string {
	return strings.ReplaceAll(ep.Targets[0], ".", "_") + "-" + ep.DNSName
}

// ApplyChanges filters out records not owned the External-DNS, additionally it adds the required label
// inserted in the AWS SD instance as a CreateID field
func (adr *AlibabaCloudSDRegistry) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	filteredChanges := &plan.Changes{
		Create:    changes.Create,
		UpdateNew: filterOwnedRecords(adr.ownerID, changes.UpdateNew),
		UpdateOld: filterOwnedRecords(adr.ownerID, changes.UpdateOld),
		Delete:    filterOwnedRecords(adr.ownerID, changes.Delete),
	}

	for _, r := range filteredChanges.Create {
		if r.Labels == nil {
			r.Labels = make(map[string]string)
		}
		r.Labels[endpoint.OwnerLabelKey] = adr.ownerID
		txt := endpoint.NewEndpoint(adr.toName(r.Targets[0], r.DNSName), endpoint.RecordTypeTXT, r.Labels.Serialize(true))
		filteredChanges.Create = append(filteredChanges.Create, txt)
	}

	for _, r := range filteredChanges.Delete {
		txt := endpoint.NewEndpoint(adr.toName(r.Targets[0], r.DNSName), endpoint.RecordTypeTXT, r.Labels.Serialize(true))
		// when we delete TXT records for which value has changed (due to new label) this would still work because
		// !!! TXT record value is uniquely generated from the Labels of the endpoint. Hence old TXT record can be uniquely reconstructed
		filteredChanges.Delete = append(filteredChanges.Delete, txt)
	}

	// make sure TXT records are consistently updated as well
	for _, r := range filteredChanges.UpdateOld {
		txt := endpoint.NewEndpoint(adr.toName(r.Targets[0], r.DNSName), endpoint.RecordTypeTXT, r.Labels.Serialize(true))
		// when we updateOld TXT records for which value has changed (due to new label) this would still work because
		// !!! TXT record value is uniquely generated from the Labels of the endpoint. Hence old TXT record can be uniquely reconstructed
		filteredChanges.UpdateOld = append(filteredChanges.UpdateOld, txt)
	}

	// make sure TXT records are consistently updated as well
	for _, r := range filteredChanges.UpdateNew {
		txt := endpoint.NewEndpoint(adr.toName(r.Targets[0], r.DNSName), endpoint.RecordTypeTXT, r.Labels.Serialize(true))
		filteredChanges.UpdateNew = append(filteredChanges.UpdateNew, txt)
	}
	return adr.provider.ApplyChanges(ctx, filteredChanges)
}

func (sdr *AlibabaCloudSDRegistry) toName(target string, endpointDNSName string) string {
	return sdr.prefix + "-" + strings.ReplaceAll(target, ".", "_") + "-" + endpointDNSName
}

func (adr *AlibabaCloudSDRegistry) toEndpointName(txtDNSName string) string {
	lowerDNSName := strings.ToLower(txtDNSName)
	if strings.HasPrefix(lowerDNSName, adr.prefix) {
		return strings.TrimPrefix(lowerDNSName, adr.prefix+"-")
	}
	return ""
}
