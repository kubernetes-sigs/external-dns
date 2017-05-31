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
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/labels"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/provider"
)

const (
	heritageLabel  = "heritage"
	heritageValue  = "external-dns"
	labelKeyPrefix = "external-dns/"
	txtRecordType  = "TXT"
)

// TXTRegistry stores labels as DNS TXT records in the wrapped Provider.
type TXTRegistry struct {
	provider provider.Provider
}

// NewTXTRegistry returns a new DNS-based registry for the passed in Provider.
func NewTXTRegistry(prov provider.Provider) *TXTRegistry {
	return &TXTRegistry{provider: prov}
}

// Records reads DNS records from the nested provider and translates any TXT records containing
// label information into labels on the returned Endpoint objects.
func (p *TXTRegistry) Records() ([]*endpoint.Endpoint, error) {
	// Read all records from the DNS provider.
	records, err := p.provider.Records()
	if err != nil {
		return nil, err
	}

	// A temporary store that holds all found labels.
	labels := map[string]map[string]string{}

	// Go through all TXT records and temporary collect and store label information.
	for _, r := range records {
		// Only consider records of type TXT.
		if r.RecordType != txtRecordType {
			continue
		}

		// Parse the value of the TXT record into a label set (a key-value map).
		labelMap := parseLabels(r.Target)

		// If the TXT record indicates a heritage of ExternalDNS, then attach the remaining labels
		// to the endpoint object. (More specific: store them to be attached in the following section.)
		// If not, skip this record.
		if labelMap[heritageLabel] != heritageValue {
			continue
		}

		// Initialize a temporary label set for the endpoint to be attached later.
		labels[r.DNSName] = map[string]string{}
		for key, value := range labelMap {
			// Any label that is prefixed by ExternalDNS' label prefix is processed. The prefix is
			// trimmed before attaching it to the endpoint object.
			//
			//   e.g.: TXT: "heritage=external-dns,external-dns/foo=bar" => labels["foo"] = "bar"
			if strings.HasPrefix(key, labelKeyPrefix) {
				labels[r.DNSName][key[len(labelKeyPrefix):]] = value
			}
		}
	}

	// Initialize the final list of endpoints to return.
	result := []*endpoint.Endpoint{}

	// Go through all non-TXT records and attach the previously collected label sets.
	for _, r := range records {
		// Skip all TXT records.
		if r.RecordType == txtRecordType {
			continue
		}

		// Attach the label set and add the record to the final list.
		r.Labels = labels[r.DNSName]
		result = append(result, r)
	}

	// Return the endpoints with their labels attached.
	return result, nil
}

// ApplyChanges creates, updates and removes a co-located TXT record for each desired endpoint
// holding information about its labels.
func (p *TXTRegistry) ApplyChanges(changes *plan.Changes) error {
	for _, ep := range changes.Create {
		// Don't create a TXT record if we don't have any labels to remember.
		if len(ep.Labels) == 0 {
			continue
		}

		// For each desired label, prefix the key with ExternalDNS' namespace.
		labelMap := map[string]string{}
		for key, value := range ep.Labels {
			labelMap[labelKeyPrefix+key] = value
		}

		// Append the TXT record to the original creation list. We also add a special label called
		// heritage to indicate that this TXT record belongs to ExternalDNS.
		changes.Create = append(changes.Create, &endpoint.Endpoint{
			DNSName:    ep.DNSName,
			Target:     fmt.Sprintf("%s=%s,%s", heritageLabel, heritageValue, formatLabels(labelMap)),
			RecordType: txtRecordType,
		})
	}

	// Forward the modified change set to the underlying provider.
	return p.provider.ApplyChanges(changes)
}

func parseLabels(labelStr string) map[string]string {
	labelMap, _ := labels.ConvertSelectorToLabelsMap(labelStr)
	return labelMap
}

func formatLabels(labelMap map[string]string) string {
	return labels.FormatLabels(labelMap)
}
