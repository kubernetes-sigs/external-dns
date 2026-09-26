/*
Copyright 2026 The Kubernetes Authors.

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

// Package crd holds types and bookkeeping shared between the crd source and the
// controller for CRD-backed sources (currently DNSEndpoint).
package crd

import (
	"context"

	log "github.com/sirupsen/logrus"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const dnsEndpointKind = "DNSEndpoint"

type refCounts struct {
	created, updated int
}

// SyncStatus stands in for the eventual status-condition write, deferred until
// kubernetes-sigs/external-dns#6609 settles that schema.
//
// changes.Delete carries no RefObjects (it comes from registry state, not source
// state — see plan/plan.go's Delete construction) and is not attributed here.
func SyncStatus(ctx context.Context, cl client.Client, changes *plan.Changes) {
	if cl == nil || changes == nil {
		return
	}
	if len(changes.Create) == 0 && len(changes.UpdateNew) == 0 {
		return
	}

	counts := make(map[types.NamespacedName]*refCounts)
	tally(counts, changes.Create, func(c *refCounts) *int { return &c.created })
	tally(counts, changes.UpdateNew, func(c *refCounts) *int { return &c.updated })

	for key, c := range counts {
		var dnsEndpoint apiv1alpha1.DNSEndpoint
		if err := cl.Get(ctx, key, &dnsEndpoint); err != nil {
			if !apierrors.IsNotFound(err) {
				log.Warnf("Could not get DNSEndpoint %s: %v", key, err)
			}
			continue
		}
		log.Debugf("DNSEndpoint %s: %d created, %d updated (generation=%d, observedGeneration=%d)",
			key, c.created, c.updated, dnsEndpoint.Generation, dnsEndpoint.Status.ObservedGeneration)
	}
}

func tally(counts map[types.NamespacedName]*refCounts, eps []*endpoint.Endpoint, field func(*refCounts) *int) {
	for _, ep := range eps {
		for _, ref := range ep.RefObjects() {
			if ref.Kind() != dnsEndpointKind {
				continue
			}
			key := types.NamespacedName{Namespace: ref.Namespace(), Name: ref.Name()}
			c, ok := counts[key]
			if !ok {
				c = &refCounts{}
				counts[key] = c
			}
			*field(c)++
		}
	}
}
