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
	"slices"

	"sigs.k8s.io/external-dns/endpoint"
)

// Policy allows to apply different rules to a set of changes.
type Policy interface {
	Apply(changes *Changes) *Changes
}

// policyName returns the stable registry key for p when it implements
// an optional `Name() string` method. Built-in policies do; third-party
// Policy implementations are not required to and fall back to an
// `unknown(<type>)` string so adding them does not break this module's
// ABI yet still leaves the operator with enough signal in debug logs to
// identify the concrete implementation. The type rendering follows Go's
// `%T` verb, which reflects the receiver: a value-type third-party
// policy renders as `unknown(plan.myPolicy)` while a pointer registration
// renders as `unknown(*plan.myPolicy)`.
func policyName(p Policy) string {
	if n, ok := p.(interface{ Name() string }); ok {
		return n.Name()
	}
	return fmt.Sprintf("unknown(%T)", p)
}

// concatSuppressed joins an intermediate SuppressedDelete carried by a
// previous policy in the chain with the Delete list the current policy
// is stripping. The explicit nil guard documents intent: when no
// suppression has occurred we want Changes.SuppressedDelete to remain
// nil (not an empty non-nil slice) so tests that use reflect.DeepEqual
// against nil keep passing and callers can rely on
// `len(SuppressedDelete) == 0` as the single zero-value check.
func concatSuppressed(prev, deletes []*endpoint.Endpoint) []*endpoint.Endpoint {
	if len(prev) == 0 && len(deletes) == 0 {
		return nil
	}
	return slices.Concat(prev, deletes)
}

// Policies is a registry of available policies, keyed by name.
var Policies = map[string]Policy{
	"sync":        &SyncPolicy{},
	"upsert-only": &UpsertOnlyPolicy{},
	"create-only": &CreateOnlyPolicy{},
}

// SyncPolicy allows for full synchronization of DNS records.
type SyncPolicy struct{}

// Name returns the registry key for this policy.
func (p *SyncPolicy) Name() string { return "sync" }

// Apply is a pass-through: sync allows all changes without restriction.
func (p *SyncPolicy) Apply(changes *Changes) *Changes {
	return changes
}

// UpsertOnlyPolicy allows everything but deleting DNS records.
type UpsertOnlyPolicy struct{}

// Name returns the registry key for this policy.
func (p *UpsertOnlyPolicy) Name() string { return "upsert-only" }

// Apply applies the upsert-only policy which strips out any deletions.
func (p *UpsertOnlyPolicy) Apply(changes *Changes) *Changes {
	return &Changes{
		Create:           changes.Create,
		UpdateOld:        changes.UpdateOld,
		UpdateNew:        changes.UpdateNew,
		SuppressedDelete: concatSuppressed(changes.SuppressedDelete, changes.Delete),
	}
}

// CreateOnlyPolicy allows only creating DNS records.
type CreateOnlyPolicy struct{}

// Name returns the registry key for this policy.
func (p *CreateOnlyPolicy) Name() string { return "create-only" }

// Apply applies the create-only policy which strips out updates and deletions.
// Dropped updates are stashed on Changes.SuppressedUpdateOld so Plan.Calculate
// can emit a per-record debug log with the same `owned` schema as suppressed
// deletions. No public metric is computed — the log is the only observability
// signal for the update path, by design (see operational-best-practices.md).
func (p *CreateOnlyPolicy) Apply(changes *Changes) *Changes {
	return &Changes{
		Create:              changes.Create,
		SuppressedDelete:    concatSuppressed(changes.SuppressedDelete, changes.Delete),
		SuppressedUpdateOld: changes.UpdateOld,
	}
}
