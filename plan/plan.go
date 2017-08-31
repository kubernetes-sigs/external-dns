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

// RecordKey is used to group records
type RecordKey struct {
	RecordType string
	DNSName    string
}

// Plan can convert a list of desired and current records to a series of create,
// update and delete actions.
type Plan struct {
	Current map[RecordKey][]string

	Desired map[RecordKey][]string

	Labels map[RecordKey]map[string]string
	// Policies under which the desired changes are calculated
	Policies []Policy
	// List of changes necessary to move towards desired state
	// Populated after calling Calculate()
	Changes *Changes
}

// NewPlan returns a plan with current and desired records grouped by RecordKey
func NewPlan(current, desired []*endpoint.Endpoint, policies []Policy) *Plan {
	p := &Plan{
		Policies: policies,
		Current:  make(map[RecordKey][]string),
		Desired:  make(map[RecordKey][]string),
		Labels:   make(map[RecordKey]map[string]string),
	}

	// aggregate desired endpoint target values
	for _, ep := range desired {
		key := RecordKey{
			RecordType: ep.RecordType,
			DNSName:    ep.DNSName,
		}
		p.Desired[key] = append(p.Desired[key], ep.Targets...)
	}

	// aggregate current endpoint target values
	for _, ep := range current {
		key := RecordKey{
			RecordType: ep.RecordType,
			DNSName:    ep.DNSName,
		}
		p.Labels[key] = ep.Labels
		p.Current[key] = append(p.Current[key], ep.Targets...)
	}

	return p
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
	for key, desired := range p.Desired {
		if _, exists := p.Current[key]; !exists {
			changes.Create = append(changes.Create, &endpoint.Endpoint{
				DNSName:    key.DNSName,
				RecordType: key.RecordType,
				Targets:    desired,
			})
		} else if endpoint.TargetSliceEquals(p.Current[key], desired) {
			log.Debugf("Skipping EndpointSet %s -> (%+v) because targets have not changed", key, desired)
		} else {
			changes.UpdateOld = append(changes.UpdateOld, &endpoint.Endpoint{
				DNSName:    key.DNSName,
				RecordType: key.RecordType,
				Targets:    p.Current[key],
				Labels:     p.Labels[key],
			})

			changes.UpdateNew = append(changes.UpdateNew, &endpoint.Endpoint{
				DNSName:    key.DNSName,
				RecordType: key.RecordType,
				Targets:    desired,
				Labels:     p.Labels[key],
			})
		}
	}

	for key, current := range p.Current {
		if _, exists := p.Desired[key]; !exists {
			changes.Delete = append(changes.Delete, &endpoint.Endpoint{
				DNSName:    key.DNSName,
				RecordType: key.RecordType,
				Targets:    current,
				Labels:     p.Labels[key],
			})
		}
	}

	for _, pol := range p.Policies {
		changes = pol.Apply(changes)
	}

	return &Plan{
		Current: p.Current,
		Desired: p.Desired,
		Changes: changes,
	}
}
