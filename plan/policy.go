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

// Policy allows to apply different rules to a set of changes.
type Policy interface {
	Apply(changes *Changes) *Changes
}

// Policies is a registry of available policies.
var Policies = map[string]Policy{
	"sync":        &SyncPolicy{},
	"upsert-only": &UpsertOnlyPolicy{},
	"create-only": &CreateOnlyPolicy{},
	"first-half":  &FirstHalfChangesPolicy{},
	"last-half":   &LastHalfChangesPolicy{},
}

// SyncPolicy allows for full synchronization of DNS records.
type SyncPolicy struct{}

// Apply applies the sync policy which returns the set of changes as is.
func (p *SyncPolicy) Apply(changes *Changes) *Changes {
	return changes
}

// UpsertOnlyPolicy allows everything but deleting DNS records.
type UpsertOnlyPolicy struct{}

// Apply applies the upsert-only policy which strips out any deletions.
func (p *UpsertOnlyPolicy) Apply(changes *Changes) *Changes {
	return &Changes{
		Create:    changes.Create,
		UpdateOld: changes.UpdateOld,
		UpdateNew: changes.UpdateNew,
	}
}

// CreateOnlyPolicy allows only creating DNS records.
type CreateOnlyPolicy struct{}

// Apply applies the create-only policy which strips out updates and deletions.
func (p *CreateOnlyPolicy) Apply(changes *Changes) *Changes {
	return &Changes{
		Create: changes.Create,
	}
}

// FirstHalfChangesPolicy allows limiting amount of records modified.
type FirstHalfChangesPolicy struct{}

// Apply applies the first half changes policy which limits change list to its first half.
func (p *FirstHalfChangesPolicy) Apply(changes *Changes) *Changes {
	if len(changes.Create) > 0 {
		halfChanges := len(changes.Create) / 2
		return &Changes{
			Create: changes.Create[:halfChanges],
		}
	}

	if len(changes.UpdateOld) > 0 && len(changes.UpdateNew) > 0 {
		halfChanges := len(changes.UpdateNew) / 2
		return &Changes{
			UpdateOld: changes.UpdateOld[:halfChanges],
			UpdateNew: changes.UpdateNew[:halfChanges],
		}
	}

	if len(changes.Delete) > 0 {
		halfChanges := len(changes.Delete) / 2
		return &Changes{
			Delete: changes.Delete[:halfChanges],
		}
	}

	return &Changes{}
}

// LastHalfChangesPolicy allows limiting amount of records modified.
type LastHalfChangesPolicy struct{}

// Apply applies the last half changes policy which limits change list to its last half.
func (p *LastHalfChangesPolicy) Apply(changes *Changes) *Changes {
	if len(changes.Create) > 0 {
		halfChanges := len(changes.Create) / 2
		return &Changes{
			Create: changes.Create[halfChanges:],
		}
	}

	if len(changes.UpdateOld) > 0 && len(changes.UpdateNew) > 0 {
		halfChanges := len(changes.UpdateNew) / 2
		return &Changes{
			UpdateOld: changes.UpdateOld[halfChanges:],
			UpdateNew: changes.UpdateNew[halfChanges:],
		}
	}

	if len(changes.Delete) > 0 {
		halfChanges := len(changes.Delete) / 2
		return &Changes{
			Delete: changes.Delete[halfChanges:],
		}
	}

	return &Changes{}
}
