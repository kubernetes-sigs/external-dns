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
	"slices"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/idna"
)

// Plan can convert a list of desired and current records to a series of create,
// update and delete actions.
type Plan struct {
	// List of current records
	// Records that already exist in the DNS provider (e.g., Route53, Cloudflare, etc.). These are fetched from the provider's registry.
	Current []*endpoint.Endpoint
	// List of desired records
	// Records that should exist based on Kubernetes resources (Ingress, Service, etc.). These are computed from the source.
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
	// Old owner ID we migrate from
	OldOwnerID string
}

// Changes holds lists of actions to be executed by dns providers
type Changes struct {
	// Records that need to be created
	Create []*endpoint.Endpoint `json:"create,omitempty"`
	// Records that need to be updated (current data)
	UpdateOld []*endpoint.Endpoint `json:"updateOld,omitempty"`
	// Records that need to be updated (desired data)
	UpdateNew []*endpoint.Endpoint `json:"updateNew,omitempty"`
	// Records that need to be deleted
	Delete []*endpoint.Endpoint `json:"delete,omitempty"`
	// SuppressedDelete lists records whose deletion was held back by policy
	// (upsert-only / create-only).
	//
	// After Plan.Calculate() it contains ONLY records owned by this
	// instance — use SuppressedDeleteTotal for the pre-partition count
	// (owned + foreign). During Policy.Apply this field transiently holds
	// the full list as an intermediate between the policy chain and the
	// ownership partition; external consumers must read it only from the
	// Changes returned by Plan.Calculate().
	//
	// json:"-" keeps the provider ApplyChanges ABI stable; observability-only.
	SuppressedDelete []*endpoint.Endpoint `json:"-"`
	// SuppressedDeleteTotal is the count of all records held back by policy
	// for this reconcile (owned + foreign), post-dedup.
	SuppressedDeleteTotal int `json:"-"`
	// SuppressedUpdateOld lists the pre-change state of records whose
	// update was held back by policy (create-only). Used only to emit a
	// per-record debug log in Plan.Calculate with the same schema as the
	// deletion log; no metric is computed because suppressed updates are
	// not a safety-net signal in the sense that suppressed deletions are.
	// Transient: populated by the policy chain and consumed in-place.
	SuppressedUpdateOld []*endpoint.Endpoint `json:"-"`
	// SuppressedUpdateTotal is the count of THIS instance's records whose
	// update was held back by policy for this reconcile. Foreign-owned
	// records are excluded because the ownership filter would have
	// discarded their updates regardless of policy — counting them
	// would make the controller's no-op log falsely blame policy in
	// shared-zone deployments where only another instance's records
	// drifted. Survives Plan.Calculate so the controller's no-op log
	// can distinguish a reconcile where create-only dropped the user's
	// own updates from a true no-op.
	SuppressedUpdateTotal int `json:"-"`
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

func (t *planTable) addCurrent(e *endpoint.Endpoint) {
	key := t.newPlanKey(e)
	t.rows[key].current = append(t.rows[key].current, e)
	t.rows[key].records[e.RecordType].current = e
}

func (t *planTable) addCandidate(e *endpoint.Endpoint) {
	key := t.newPlanKey(e)
	row := t.rows[key]
	row.candidates = append(row.candidates, e)
	row.records[e.RecordType].candidates = append(row.records[e.RecordType].candidates, e)
}

func (t *planTable) newPlanKey(e *endpoint.Endpoint) planKey {
	key := planKey{
		dnsName:       idna.NormalizeDNSName(e.DNSName),
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
	return !cmp.Equal(c.UpdateNew, c.UpdateOld, cmpopts.IgnoreUnexported(endpoint.Endpoint{}))
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

	if p.OwnerID != "" {
		registryOwnerMismatchPerSync.Gauge.Reset()
	}
	changes := p.calculateChanges(t)

	// Return a minimal plan with only the fields relevant to callers.
	// ManagedRecords is reset to the canonical defaults (A/AAAA/CNAME) —
	// this is intentional: it restores the default managed set regardless
	// of what was passed in, preventing callers that chain off Calculate()
	// from accidentally inheriting a non-default managed record configuration.
	// See: https://github.com/kubernetes-sigs/external-dns/pull/1915
	plan := &Plan{
		Current:        p.Current,
		Desired:        p.Desired,
		Changes:        changes,
		ManagedRecords: []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
	}

	return plan
}

func (p *Plan) calculateChanges(t planTable) *Changes {
	changes := &Changes{}

	for key, row := range t.rows {
		switch {
		// dns name not taken
		case len(row.current) == 0:
			recordsByType := t.resolver.ResolveRecordTypes(key, row)
			for _, records := range recordsByType {
				if len(records.candidates) > 0 {
					changes.Create = append(changes.Create, t.resolver.ResolveCreate(records.candidates))
				}
			}

		// dns name released or possibly owned by a different external dns
		case len(row.candidates) == 0:
			changes.Delete = append(changes.Delete, row.current...)

		// dns name is taken
		case len(row.candidates) > 0:
			p.appendTakenDNSNameChanges(t, changes, key, row)
		}
	}

	for _, pol := range p.Policies {
		changes = pol.Apply(changes)
	}

	if len(changes.SuppressedDelete) > 0 {
		changes.SuppressedDelete, changes.SuppressedDeleteTotal = p.partitionSuppressed(changes.SuppressedDelete)
	}

	// Log suppressed updates (create-only drops UpdateOld/UpdateNew wholesale;
	// this is the only operator-facing visibility, as there is no metric).
	// The log fires BEFORE the ownership filter below, same as suppressed
	// deletions, so `owned` distinguishes records this instance manages from
	// those belonging to other external-dns instances in a shared zone. The
	// field is cleared unconditionally (not gated on debug level) so the
	// returned Changes does not leak transient policy state to downstream
	// consumers, and log-level changes don't alter the public contract.
	if len(changes.SuppressedUpdateOld) > 0 {
		// Count only owned suppressions — foreign-owned updates would
		// have been dropped by the ownership filter below regardless of
		// policy, so counting them would falsely blame the policy in
		// the controller's no-op log for a shared-zone reconcile where
		// only another instance's records drifted.
		for _, ep := range changes.SuppressedUpdateOld {
			if p.ownsEndpoint(ep) {
				changes.SuppressedUpdateTotal++
			}
		}
		if log.IsLevelEnabled(log.DebugLevel) {
			policyField := p.policyChainString()
			for _, ep := range changes.SuppressedUpdateOld {
				log.WithFields(log.Fields{
					"record":  ep.DNSName,
					"type":    ep.RecordType,
					"targets": strings.Join(ep.Targets, ","),
					"owned":   p.ownsEndpoint(ep),
					"policy":  policyField,
				}).Debug("skipping update of record due to policy")
			}
		}
		changes.SuppressedUpdateOld = nil
	}

	// filter out changes this external dns does not have ownership claim over
	if p.OwnerID != "" {
		changes.Delete = endpoint.FilterEndpointsByOwnerID(p.OwnerID, changes.Delete)
		changes.Delete = endpoint.RemoveDuplicates(changes.Delete)
		changes.UpdateOld = endpoint.FilterEndpointsByOwnerID(p.OwnerID, changes.UpdateOld)
		changes.UpdateNew = endpoint.FilterEndpointsByOwnerID(p.OwnerID, changes.UpdateNew)
	}

	return changes
}

// policyChainString renders the comma-joined registry keys of p.Policies
// for use as a structured log field. An empty chain is rendered as "none"
// so the field stays grep-friendly even in pathological configurations.
func (p *Plan) policyChainString() string {
	names := make([]string, 0, len(p.Policies))
	for _, pol := range p.Policies {
		names = append(names, policyName(pol))
	}
	s := strings.Join(names, ",")
	if s == "" {
		return "none"
	}
	return s
}

// ownsEndpoint reports whether ep carries this instance's owner label.
// When OwnerID is unset (ownership disabled) every endpoint is considered
// owned so suppressed deletions are not reported as orphans — this
// short-circuit DIVERGES from endpoint.IsOwnedBy, which returns false for
// an empty ownerID. For non-empty OwnerID the check delegates to
// endpoint.IsOwnedBy so the ownership predicate stays consistent with
// the rest of the codebase.
func (p *Plan) ownsEndpoint(ep *endpoint.Endpoint) bool {
	if p.OwnerID == "" {
		return true
	}
	return ep.IsOwnedBy(p.OwnerID)
}

// partitionSuppressed splits suppressed into owned / foreign buckets,
// deduplicates within each, and (when debug logging is enabled) emits
// one structured "skipping deletion of record due to policy" entry per
// input record — the log fires BEFORE dedup, so pathological providers
// can still inflate log volume; the `!!! info "Debug-log volume…"`
// admonition in operational-best-practices.md is the operator-facing
// mitigation. It returns the owned bucket (nil when empty so the
// zero-value contract on Changes.SuppressedDelete holds) and the total
// count across both buckets post-dedup.
//
// Partitioning BEFORE dedup is load-bearing: endpoint.RemoveDuplicates
// keys only on (DNS name, type, set identifier) and ignores labels, so
// a mixed-owner duplicate for the same key would otherwise let the
// foreign twin evict the owned record. Dedup within each bucket keeps
// pathological providers from inflating the metric.
func (p *Plan) partitionSuppressed(suppressed []*endpoint.Endpoint) ([]*endpoint.Endpoint, int) {
	logSuppressions := log.IsLevelEnabled(log.DebugLevel)
	var policyField string
	if logSuppressions {
		policyField = p.policyChainString()
	}

	owned := make([]*endpoint.Endpoint, 0, len(suppressed))
	// foreign is left nil and grown by append: when OwnerID is unset
	// every record is considered owned and this stays nil; when OwnerID
	// is set the foreign count is typically a minority of suppressed
	// and pre-sizing to len(suppressed) would over-allocate.
	var foreign []*endpoint.Endpoint
	for _, ep := range suppressed {
		isOwned := p.ownsEndpoint(ep)
		if logSuppressions {
			log.WithFields(log.Fields{
				"record":  ep.DNSName,
				"type":    ep.RecordType,
				"targets": strings.Join(ep.Targets, ","),
				"owned":   isOwned,
				"policy":  policyField,
			}).Debug("skipping deletion of record due to policy")
		}
		if isOwned {
			owned = append(owned, ep)
		} else {
			foreign = append(foreign, ep)
		}
	}
	if len(owned) > 0 {
		owned = endpoint.RemoveDuplicates(owned)
	}
	if len(foreign) > 0 {
		foreign = endpoint.RemoveDuplicates(foreign)
	}
	total := len(owned) + len(foreign)
	if len(owned) == 0 {
		return nil, total
	}
	return owned, total
}

func (p *Plan) appendTakenDNSNameChanges(
	t planTable,
	changes *Changes,
	key planKey,
	row *planTableRow) {
	// apply changes for each record type
	rowChanges := p.calculatePlanTableRowChanges(t, key, row)
	changes.Delete = append(changes.Delete, rowChanges.Delete...)
	changes.UpdateNew = append(changes.UpdateNew, rowChanges.UpdateNew...)
	changes.UpdateOld = append(changes.UpdateOld, rowChanges.UpdateOld...)
	if len(rowChanges.Create) == 0 {
		return
	}

	// only add creates if the external dns has ownership claim on the domain
	ownersMatch := true
	if p.OwnerID != "" {
		for _, current := range row.current {
			if !current.IsOwnedBy(p.OwnerID) {
				ownersMatch = false
				recordOwnerMismatch(p.OwnerID, current)
				if log.IsLevelEnabled(log.DebugLevel) {
					log.Debugf(`Skipping endpoint %v because owner id does not match for one or more items to create, found: "%s", required: "%s"`, current, current.Labels[endpoint.OwnerLabelKey], p.OwnerID)
				}
			}
		}
	}

	if ownersMatch {
		changes.Create = append(changes.Create, rowChanges.Create...)
	}
}

func (p *Plan) calculatePlanTableRowChanges(t planTable, key planKey, row *planTableRow) *Changes {
	changes := &Changes{}

	recordsByType := t.resolver.ResolveRecordTypes(key, row)
	for _, records := range recordsByType {
		switch {
		// record type not desired
		case records.current != nil && len(records.candidates) == 0:
			changes.Delete = append(changes.Delete, records.current)

		// new record type desired
		case records.current == nil && len(records.candidates) > 0:
			update := t.resolver.ResolveCreate(records.candidates)
			// creates are evaluated after all domain records have been processed to
			// validate that this external dns has ownership claim on the domain before
			// adding the records to planned changes.
			changes.Create = append(changes.Create, update)

		// update existing record
		case records.current != nil && len(records.candidates) > 0:
			p.appendEndpointUpdates(t, changes, records.current, records.candidates)
		}
	}

	return changes
}

func (p *Plan) appendEndpointUpdates(t planTable, changes *Changes, current *endpoint.Endpoint, candidates []*endpoint.Endpoint) {
	update := t.resolver.ResolveUpdate(current, candidates)

	if shouldUpdateTTL(update, current) || targetChanged(update, current) ||
		p.providerSpecificChanged(update, current) || p.isOldOwnerIDSetAndDifferent(current) {
		inheritOwner(current, update)
		changes.UpdateNew = append(changes.UpdateNew, update)
		changes.UpdateOld = append(changes.UpdateOld, current)
	}
}

func (p *Plan) isOldOwnerIDSetAndDifferent(current *endpoint.Endpoint) bool {
	return p.OldOwnerID != "" && current.Labels[endpoint.OwnerLabelKey] != p.OldOwnerID
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

func (p *Plan) providerSpecificChanged(desired, current *endpoint.Endpoint) bool {
	desiredProperties := make(map[string]endpoint.ProviderSpecificProperty, len(desired.ProviderSpecific))

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
// Currently, this just removes TXT records to prevent them from being
// deleted erroneously by the planner (only the TXT registry should do this.)
//
// Per RFC 1034, CNAME records conflict with all other records - it is the
// only record with this property. The behavior of the planner may need to be
// made more sophisticated to codify this.
func filterRecordsForPlan(records []*endpoint.Endpoint, domainFilter endpoint.MatchAllDomainFilters, managedRecords, excludeRecords []string) []*endpoint.Endpoint {
	filtered := make([]*endpoint.Endpoint, 0, len(records))

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

func IsManagedRecord(record string, managedRecords, excludeRecords []string) bool {
	if slices.Contains(excludeRecords, record) {
		return false
	}
	return slices.Contains(managedRecords, record)
}
