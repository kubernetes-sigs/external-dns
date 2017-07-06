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
	// Policies under which the desired changes are calculated
	Policies []Policy
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
	usedDesired := map[int]bool{}
	usedCurrent := map[int]bool{}

	for c, current := range p.Current {
		for d, desired := range p.Desired {
			if current.DNSName == desired.DNSName && current.Target == desired.Target &&
				current.SuitableType() == desired.SuitableType() {
				usedDesired[d] = true
				usedCurrent[c] = true
				log.Debugf("Skipping endpoint %v because target has not changed", desired)
			}
		}
	}

	for c, current := range p.Current {
		if _, found := usedCurrent[c]; found {
			continue
		}
		for d, desired := range p.Desired {
			if _, found := usedDesired[d]; found {
				continue
			}
			if current.DNSName == desired.DNSName && current.SuitableType() == desired.SuitableType() {
				changes.UpdateOld = append(changes.UpdateOld, current)
				changes.UpdateNew = append(changes.UpdateNew, desired)
				if desired.Labels != nil && desired.Labels[endpoint.OwnerLabelKey] == "" {
					desired.Labels[endpoint.OwnerLabelKey] = current.Labels[endpoint.OwnerLabelKey]
				}
				desired.RecordType = current.RecordType
				usedDesired[d] = true
				usedCurrent[c] = true
			}
		}
	}
	for d, desired := range p.Desired {
		if _, found := usedDesired[d]; !found {
			changes.Create = append(changes.Create, desired)
		}
	}
	for c, current := range p.Current {
		if _, found := usedCurrent[c]; !found {
			changes.Delete = append(changes.Delete, current)
		}
	}

	// Apply policies to list of changes.
	for _, pol := range p.Policies {
		changes = pol.Apply(changes)
	}

	plan := &Plan{
		Current: p.Current,
		Desired: p.Desired,
		Changes: changes,
	}

	return plan
}
