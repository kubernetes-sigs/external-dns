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
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
)

// PropertyComparator is used in Plan for comparing the previous and current custom annotations.
type PropertyComparator func(name string, previous string, current string) bool

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
	// DomainFilter matches DNS names
	DomainFilter endpoint.MatchAllDomainFilters
	// ManagedRecords are DNS record types that will be considered for management.
	ManagedRecords []string
	// ExcludeRecords are DNS record types that will be excluded from management.
	ExcludeRecords []string
	// OwnerID of records to manage
	OwnerID string
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

// planKey is a key for a row in `planTable`.
type planKey struct {
	dnsName       string
	setIdentifier string
}

// planTable is a supplementary struct for Plan
// each row correspond to a planKey -> (current records + all desired records)
//
//	planTable (-> = target)
//	--------------------------------------------------------------
//	DNSName | Current record       | Desired Records             |
//	--------------------------------------------------------------
//	foo.com | [->1.1.1.1 ]         | [->1.1.1.1]                 |  = no action
//	--------------------------------------------------------------
//	bar.com |                      | [->191.1.1.1, ->190.1.1.1]  |  = create (bar.com [-> 190.1.1.1])
//	--------------------------------------------------------------
//	dog.com | [->1.1.1.2]          |                             |  = delete (dog.com [-> 1.1.1.2])
//	--------------------------------------------------------------
//	cat.com | [->::1, ->1.1.1.3]   | [->1.1.1.3]                 |  = update old (cat.com [-> ::1, -> 1.1.1.3]) new (cat.com [-> 1.1.1.3])
//	--------------------------------------------------------------
//	big.com | [->1.1.1.4]          | [->ing.elb.com]             |  = update old (big.com [-> 1.1.1.4]) new (big.com [-> ing.elb.com])
//	--------------------------------------------------------------
//	"=", i.e. result of calculation relies on supplied ConflictResolver
type planTable struct {
	rows     map[planKey]*planTableRow
	resolver ConflictResolver
}

func newPlanTable() planTable { // TODO: make resolver configurable
	return planTable{map[planKey]*planTableRow{}, PerResource{}}
}

// planTableRow represents a set of current and desired domain resource records.
type planTableRow struct {
	// current corresponds to the records currently occupying dns name on the dns provider. More than one record may
	// be represented here: for example A and AAAA. If the current domain record is a CNAME, no other record types
	// are allowed per [RFC 1034 3.6.2]
	//
	// [RFC 1034 3.6.2]: https://datatracker.ietf.org/doc/html/rfc1034#autoid-15
	current []*endpoint.Endpoint
	// candidates corresponds to the list of records which would like to have this dnsName.
	candidates []*endpoint.Endpoint
	// records is a grouping of current and candidates by record type, for example A, AAAA, CNAME.
	records map[string]*domainEndpoints
}

// domainEndpoints is a grouping of current, which are existing records from the registry, and candidates,
// which are desired records from the source. All records in this grouping have the same record type.
type domainEndpoints struct {
	// current corresponds to existing record from the registry. Maybe nil if no current record of the type exists.
	current *endpoint.Endpoint
	// candidates corresponds to the list of records which would like to have this dnsName.
	candidates []*endpoint.Endpoint
}

func (t planTableRow) String() string {
	return fmt.Sprintf("planTableRow{current=%v, candidates=%v}", t.current, t.candidates)
}

func (t planTable) addCurrent(e *endpoint.Endpoint) {
	key := t.newPlanKey(e)
	t.rows[key].current = append(t.rows[key].current, e)
	t.rows[key].records[e.RecordType].current = e
}

func (t planTable) addCandidate(e *endpoint.Endpoint) {
	key := t.newPlanKey(e)
	t.rows[key].candidates = append(t.rows[key].candidates, e)
	t.rows[key].records[e.RecordType].candidates = append(t.rows[key].records[e.RecordType].candidates, e)
}

func (t *planTable) newPlanKey(e *endpoint.Endpoint) planKey {
	key := planKey{
		dnsName:       normalizeDNSName(e.DNSName),
		setIdentifier: e.SetIdentifier,
	}

	if _, ok := t.rows[key]; !ok {
		t.rows[key] = &planTableRow{
			records: make(map[string]*domainEndpoints),
		}
	}

	if _, ok := t.rows[key].records[e.RecordType]; !ok {
		t.rows[key].records[e.RecordType] = &domainEndpoints{}
	}

	return key
}

func (c *Changes) HasChanges() bool {
	if len(c.Create) > 0 || len(c.Delete) > 0 {
		return true
	}
	return !cmp.Equal(c.UpdateNew, c.UpdateOld)
}

// Calculate computes the actions needed to move current state towards desired
// state. It then passes those changes to the current policy for further
// processing. It returns a copy of Plan with the changes populated.
func (p *Plan) Calculate() *Plan {
	t := newPlanTable()

	if p.DomainFilter == nil {
		p.DomainFilter = endpoint.MatchAllDomainFilters(nil)
	}

	for _, current := range filterRecordsForPlan(p.Current, p.DomainFilter, p.ManagedRecords, p.ExcludeRecords) {
		t.addCurrent(current)
	}
	for _, desired := range filterRecordsForPlan(p.Desired, p.DomainFilter, p.ManagedRecords, p.ExcludeRecords) {
		t.addCandidate(desired)
	}

	changes := &Changes{}

	for key, row := range t.rows {
		// dns name not taken
		if len(row.current) == 0 {
			recordsByType := t.resolver.ResolveRecordTypes(key, row)
			for _, records := range recordsByType {
				if len(records.candidates) > 0 {
					changes.Create = append(changes.Create, t.resolver.ResolveCreate(records.candidates))
				}
			}
		}

		// dns name released or possibly owned by a different external dns
		if len(row.current) > 0 && len(row.candidates) == 0 {
			changes.Delete = append(changes.Delete, row.current...)
		}

		// dns name is taken
		if len(row.current) > 0 && len(row.candidates) > 0 {
			creates := []*endpoint.Endpoint{}

			// apply changes for each record type
			recordsByType := t.resolver.ResolveRecordTypes(key, row)
			for _, records := range recordsByType {
				// record type not desired
				if records.current != nil && len(records.candidates) == 0 {
					changes.Delete = append(changes.Delete, records.current)
				}

				// new record type desired
				if records.current == nil && len(records.candidates) > 0 {
					update := t.resolver.ResolveCreate(records.candidates)
					// creates are evaluated after all domain records have been processed to
					// validate that this external dns has ownership claim on the domain before
					// adding the records to planned changes.
					creates = append(creates, update)
				}

				// update existing record
				if records.current != nil && len(records.candidates) > 0 {
					update := t.resolver.ResolveUpdate(records.current, records.candidates)

					if shouldUpdateTTL(update, records.current) || targetChanged(update, records.current) || p.shouldUpdateProviderSpecific(update, records.current) {
						inheritOwner(records.current, update)
						changes.UpdateNew = append(changes.UpdateNew, update)
						changes.UpdateOld = append(changes.UpdateOld, records.current)
					}
				}
			}

			if len(creates) > 0 {
				// only add creates if the external dns has ownership claim on the domain
				ownersMatch := true
				for _, current := range row.current {
					if p.OwnerID != "" && !current.IsOwnedBy(p.OwnerID) {
						ownersMatch = false
					}
				}

				if ownersMatch {
					changes.Create = append(changes.Create, creates...)
				}
			}
		}
	}

	for _, pol := range p.Policies {
		changes = pol.Apply(changes)
	}

	// filter out updates this external dns does not have ownership claim over
	if p.OwnerID != "" {
		changes.Delete = endpoint.FilterEndpointsByOwnerID(p.OwnerID, changes.Delete)
		changes.Delete = endpoint.RemoveDuplicates(changes.Delete)
		changes.UpdateOld = endpoint.FilterEndpointsByOwnerID(p.OwnerID, changes.UpdateOld)
		changes.UpdateNew = endpoint.FilterEndpointsByOwnerID(p.OwnerID, changes.UpdateNew)
	}

	plan := &Plan{
		Current:        p.Current,
		Desired:        p.Desired,
		Changes:        changes,
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
	}

	return plan
}

func inheritOwner(from, to *endpoint.Endpoint) {
	if to.Labels == nil {
		to.Labels = map[string]string{}
	}
	if from.Labels == nil {
		from.Labels = map[string]string{}
	}
	to.Labels[endpoint.OwnerLabelKey] = from.Labels[endpoint.OwnerLabelKey]
}

func targetChanged(desired, current *endpoint.Endpoint) bool {
	return !desired.Targets.Same(current.Targets)
}

func shouldUpdateTTL(desired, current *endpoint.Endpoint) bool {
	if !desired.RecordTTL.IsConfigured() {
		return false
	}
	return desired.RecordTTL != current.RecordTTL
}

func (p *Plan) shouldUpdateProviderSpecific(desired, current *endpoint.Endpoint) bool {
	desiredProperties := map[string]endpoint.ProviderSpecificProperty{}

	for _, d := range desired.ProviderSpecific {
		desiredProperties[d.Name] = d
	}
	for _, c := range current.ProviderSpecific {
		if d, ok := desiredProperties[c.Name]; ok {
			if c.Value != d.Value {
				return true
			}
			delete(desiredProperties, c.Name)
		} else {
			return true
		}
	}

	return len(desiredProperties) > 0
}

// filterRecordsForPlan removes records that are not relevant to the planner.
// Currently this just removes TXT records to prevent them from being
// deleted erroneously by the planner (only the TXT registry should do this.)
//
// Per RFC 1034, CNAME records conflict with all other records - it is the
// only record with this property. The behavior of the planner may need to be
// made more sophisticated to codify this.
func filterRecordsForPlan(records []*endpoint.Endpoint, domainFilter endpoint.MatchAllDomainFilters, managedRecords, excludeRecords []string) []*endpoint.Endpoint {
	filtered := []*endpoint.Endpoint{}

	for _, record := range records {
		// Ignore records that do not match the domain filter provided
		if !domainFilter.Match(record.DNSName) {
			log.Debugf("ignoring record %s that does not match domain filter", record.DNSName)
			continue
		}
		if IsManagedRecord(record.RecordType, managedRecords, excludeRecords) {
			filtered = append(filtered, record)
		}
	}

	return filtered
}

// normalizeDNSName converts a DNS name to a canonical form, so that we can use string equality
// it: removes space, converts to lower case, ensures there is a trailing dot
func normalizeDNSName(dnsName string) string {
	s := strings.TrimSpace(strings.ToLower(dnsName))
	if !strings.HasSuffix(s, ".") {
		s += "."
	}
	return s
}

func IsManagedRecord(record string, managedRecords, excludeRecords []string) bool {
	for _, r := range excludeRecords {
		if record == r {
			return false
		}
	}
	for _, r := range managedRecords {
		if record == r {
			return true
		}
	}
	return false
}
