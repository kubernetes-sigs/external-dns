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

package provider

import (
	"context"
	"net"
	"strings"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

// Provider defines the interface DNS providers should implement.
type Provider interface {
	Records(ctx context.Context) ([]*endpoint.Endpoint, error)
	ApplyChanges(ctx context.Context, changes *plan.Changes) error
}

type contextKey struct {
	name string
}

func (k *contextKey) String() string { return "provider context value " + k.name }

// RecordsContextKey is a context key. It can be used during ApplyChanges
// to access previously cached records. The associated value will be of
// type []*endpoint.Endpoint.
var RecordsContextKey = &contextKey{"records"}

// ensureTrailingDot ensures that the hostname receives a trailing dot if it hasn't already.
func ensureTrailingDot(hostname string) string {
	if net.ParseIP(hostname) != nil {
		return hostname
	}

	return strings.TrimSuffix(hostname, ".") + "."
}
