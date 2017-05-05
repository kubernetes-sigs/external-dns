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
	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// Policy allows to apply different rules to a set of changes.
type Policy interface {
	Apply(changes *Changes) *Changes
}

// Policies is a registry of available policies.
var Policies = map[string]Policy{
	"sync":        &SyncPolicy{},
	"upsert-only": &UpsertOnlyPolicy{},
}

type OwnershipPolicy struct {
	Owner string
}

func (p *OwnershipPolicy) Apply(changes *Changes) *Changes {
	changes.Create = filterByOwner(p.Owner, changes.Create)
	changes.UpdateOld = filterByOwner(p.Owner, changes.UpdateOld)
	changes.UpdateNew = filterByOwner(p.Owner, changes.UpdateNew)
	changes.Delete = filterByOwner(p.Owner, changes.Delete)

	return changes
}

func filterByOwner(owner string, endpoints []*endpoint.Endpoint) []*endpoint.Endpoint {
	result := []*endpoint.Endpoint{}

	for _, ep := range endpoints {
		if ep.Labels["owner"] == owner {
			result = append(result, ep)
		}
	}

	return result
}

// SyncPolicy allows for full synchronization of DNS records.
type SyncPolicy struct{}

// Apply applies the sync policy which returns the set of changes as is.
func (p *SyncPolicy) Apply(changes *Changes) *Changes {
	return changes
}

// UpsertOnlyPolicy allows evrything but deleting DNS records.
type UpsertOnlyPolicy struct{}

// Apply applies the upsert-only policy which strips out any deletions.
func (p *UpsertOnlyPolicy) Apply(changes *Changes) *Changes {
	return &Changes{
		Create:    changes.Create,
		UpdateOld: changes.UpdateOld,
		UpdateNew: changes.UpdateNew,
	}
}
