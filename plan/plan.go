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

		// Check for AWS Route 53 specific policies
		if current.Policy.HasAWSRoute53Policy() {
			// If there is a record with no policy but the desired has a policy, do not update
			if !desired.Policy.HasAWSRoute53Policy() {
				continue
			}
			// If there's no current record with the desired policy, create it.
			current, exists = awsRecordExistsWithSetIdentifier(desired, p.Current)
			if !exists {
				log.Debugf("Endpoint %s does not exist with a set identifier. Will create it", desired.DNSName)
				desired.Labels[endpoint.TxtOwnedLabelKey] = "true"
				changes.Create = append(changes.Create, desired)
				continue
			}
		} else {
			// If there is a record with no policy but the desired has a policy, do not update as this is
			// not supported by AWS
			if desired.Policy.HasAWSRoute53Policy() {
				log.Errorf("Cannot add record %s with a policy because there is an existing record "+
					"by the same name that does not have a policy", desired.DNSName)
				continue
			}
		}

		// If there already is a record update it if it changed.
		if desired.Target != current.Target {
			changes.UpdateOld = append(changes.UpdateOld, current)

			desired.RecordType = current.RecordType // inherit the type from the dns provider
			desired.MergeLabels(current.Labels)     // inherit the labels from the dns provider, including Owner ID
			changes.UpdateNew = append(changes.UpdateNew, desired)
			continue
		}

		// If there is a policy and it doesn't match, update it.
		if desired.Policy != current.Policy {
			changes.UpdateOld = append(changes.UpdateOld, current)
			desired.Policy = current.Policy
		}
		log.Debugf("Skipping endpoint %v because it has not changed", desired)
	}

	// Ensure all undesired records are removed. Each current record that cannot
	// be found in the list of desired records is removed.
	// TODO (jswoods): consider moving this out of plan and into its own type
	referenceCounter := make(map[string]map[string][]*endpoint.Endpoint)
	for _, current := range p.Current {
		if referenceCounter[current.DNSName] == nil {
			referenceCounter[current.DNSName] = map[string][]*endpoint.Endpoint{}
		}
		if _, exists := recordExists(current, p.Desired); exists {
			if current.Policy.HasAWSRoute53Policy() {

				// Do not delete if any existing records do not have AWS policies
				if !awsRecordHasSetIdentifier(p.Desired) {
					continue
				}
				if _, SetIdentifierExists := awsRecordExistsWithSetIdentifier(current, p.Desired); !SetIdentifierExists {

					log.Debugf("Adding endpoint %v with weight policy %v to reference counter",
						current, current.Policy.AWSRoute53.Weight)

					referenceCounter[current.DNSName]["delete"] =
						append(referenceCounter[current.DNSName]["delete"], current)
				} else {
					referenceCounter[current.DNSName]["keep"] = append(referenceCounter[current.DNSName]["keep"], current)
				}
			}
		} else {
			referenceCounter[current.DNSName]["delete"] = append(referenceCounter[current.DNSName]["delete"], current)
		}
	}

	for _, endpointsList := range referenceCounter {
		for _, ep := range endpointsList["delete"] {
			if ep.Policy.HasAWSRoute53Policy() {
				if len(endpointsList["keep"]) != 0 {
					ep.Labels[endpoint.TxtOwnedLabelKey] = "true"
				}
			}
			changes.Delete = append(changes.Delete, ep)
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

// recordExists checks whether a record can be found in a list of records.
func recordExists(needle *endpoint.Endpoint, haystack []*endpoint.Endpoint) (*endpoint.Endpoint, bool) {
	for _, record := range haystack {
		if record.DNSName == needle.DNSName {
			return record, true
		}
	}

	return nil, false
}

// awsRecordExistsWithSetIdentifier checks whether a record with a Route53 policy exists in a list of records
func awsRecordExistsWithSetIdentifier(needle *endpoint.Endpoint, haystack []*endpoint.Endpoint) (*endpoint.Endpoint, bool) {
	for _, record := range haystack {
		if record.DNSName == needle.DNSName && record.Policy.HasAWSRoute53Policy() &&
			needle.Policy.HasAWSRoute53Policy() {
			if record.Policy.AWSRoute53.SetIdentifier == needle.Policy.AWSRoute53.SetIdentifier {
				return record, true
			}
		}
	}

	return nil, false
}

func awsRecordHasSetIdentifier(records []*endpoint.Endpoint) bool {
	for _, record := range records {
		if record.Policy.HasAWSRoute53Policy() {
			return true
		}
	}
	return false
}
