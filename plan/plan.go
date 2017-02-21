package plan

// DNSRecord holds information about a DNS record.
type DNSRecord struct {
	// The hostname of the DNS record
	DNSName string
	// The target the DNS record points to
	Target string
}

// Plan can convert a list of desired and current records to a series of create,
// update and delete actions.
type Plan struct {
	// List of current records
	Current []DNSRecord
	// List of desired records
	Desired []DNSRecord

	// The following lists hold actions to take in orer to move current state to
	// desired state.

	// Records that need to be created
	Create []DNSRecord
	// Records that need to be updated (current data)
	UpdateOld []DNSRecord
	// Records that need to be updated (desired data)
	UpdateNew []DNSRecord
	// Records that need to be deleted
	Delete []DNSRecord
}

// Calculate computes the actions needed to move current state to desired state.
func (p *Plan) Calculate() *Plan {
	plan := &Plan{}

	// Ensure all desired records exist. For each desired record make sure it's
	// either created or updated.
	for _, desired := range p.Desired {
		// Get the matching current record if it exists.
		current, exists := recordExists(desired, p.Current)

		// If there's no current record create desired record.
		if !exists {
			plan.Create = append(plan.Create, desired)
			continue
		}

		// If there already is a record update it if it changed.
		if desired.Target != current.Target {
			plan.UpdateOld = append(plan.UpdateOld, current)
			plan.UpdateNew = append(plan.UpdateNew, desired)
		}
	}

	// Ensure all undesired records are removed. Each current record that cannot
	// be found in the list of desired records is removed.
	for _, current := range p.Current {
		if _, exists := recordExists(current, p.Desired); !exists {
			plan.Delete = append(plan.Delete, current)
		}
	}

	return plan
}

// recordExists checks whether a record can be found in a list of records.
func recordExists(needle DNSRecord, haystack []DNSRecord) (DNSRecord, bool) {
	for _, record := range haystack {
		if record.DNSName == needle.DNSName {
			return record, true
		}
	}

	return DNSRecord{}, false
}
