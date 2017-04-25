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

package plan

import (
	log "github.com/Sirupsen/logrus"
	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// Plan can convert a list of desired and current records to a series of create,
// update and delete actions.
type Plan struct {
	// List of current records
	Current []*endpoint.Endpoint
	// List of desired records
	Desired []*endpoint.Endpoint
	// Policy under which the desired changes are calculated
	Policy Policy
	// List of changes necessary to move towards desired state
	// Populated after calling Calculate()
	Changes *Changes
}

// Changes holds lists of actions to be executed by dns providers
type Changes struct {
	// Records that need to be created
	Create []*endpoint.Endpoint
	// Records that need to be updated (current data)
	UpdateOld []*endpoint.Endpoint
	// Records that need to be updated (desired data)
	UpdateNew []*endpoint.Endpoint
	// Records that need to be deleted
	Delete []*endpoint.Endpoint
}

// Calculate computes the actions needed to move current state towards desired
// state. It then passes those changes to the current policy for further
// processing. It returns a copy of Plan with the changes populated.
func (p *Plan) Calculate() *Plan {
	changes := &Changes{}

	// Ensure all desired records exist. For each desired record make sure it's
	// either created or updated.
	for _, desired := range p.Desired {
		// Get the matching current record if it exists.
		current, exists := recordExists(desired, p.Current)

		// If there's no current record create desired record.
		if !exists {
			changes.Create = append(changes.Create, desired)
			continue
		}

		// If there already is a record update it if it changed.
		if desired.Target != current.Target {
			changes.UpdateOld = append(changes.UpdateOld, current)

			desired.RecordType = current.RecordType // inherit the type from the dns provider
			desired.MergeLabels(current.Labels)     // inherit the labels from the dns provider, including Owner ID
			changes.UpdateNew = append(changes.UpdateNew, desired)
			continue
		}

		log.Debugf("Skipping endpoint %v because target has not changed", desired)
	}

	// Ensure all undesired records are removed. Each current record that cannot
	// be found in the list of desired records is removed.
	for _, current := range p.Current {
		if _, exists := recordExists(current, p.Desired); !exists {
			changes.Delete = append(changes.Delete, current)
		}
	}

	// Apply policy to list of changes.
	changes = p.Policy.Apply(changes)

	plan := &Plan{
		Current: p.Current,
		Desired: p.Desired,
		Changes: changes,
	}

	return plan
}

// recordExists checks whether a record can be found in a list of records.
func recordExists(needle *endpoint.Endpoint, haystack []*endpoint.Endpoint) (*endpoint.Endpoint, bool) {
	for _, record := range haystack {
		if record.DNSName == needle.DNSName {
			return record, true
		}
	}

	return nil, false
}
