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
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/plan"
)

const dnsEndpointKind = "DNSEndpoint"

// StatusWriter reports each sync's outcome on the DNSEndpoint objects behind it.
type StatusWriter struct {
	clients *CRDClients
}

// NewStatusWriter returns a StatusWriter using the crd source's clients.
func NewStatusWriter(clients *CRDClients) *StatusWriter {
	return &StatusWriter{clients: clients}
}

// ReportStatus only logs until DNSEndpoint has status conditions.
func (w *StatusWriter) ReportStatus(ctx context.Context, objects []plan.PlannedObject, applyErr error) {
	for _, obj := range objects {
		if obj.Ref == nil || obj.Ref.Kind() != dnsEndpointKind {
			continue
		}

		key := client.ObjectKey{Namespace: obj.Ref.Namespace(), Name: obj.Ref.Name()}
		var dnsEndpoint apiv1alpha1.DNSEndpoint
		if err := w.clients.Reader().Get(ctx, key, &dnsEndpoint); err != nil {
			// Deleted since the plan.
			if !apierrors.IsNotFound(err) {
				log.Warnf("Could not get DNSEndpoint %s: %v", key, err)
			}
			continue
		}

		log.Debugf("DNSEndpoint %s: %d endpoint(s) planned, apply error: %v (generation=%d, observedGeneration=%d)",
			key, obj.Endpoints, applyErr, dnsEndpoint.Generation, dnsEndpoint.Status.ObservedGeneration)
	}
}
