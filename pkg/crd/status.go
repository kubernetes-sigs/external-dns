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
	"fmt"
	"unicode/utf8"

	log "github.com/sirupsen/logrus"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/plan"
)

const dnsEndpointKind = "DNSEndpoint"

// StatusWriter reports each sync's outcome on the DNSEndpoints behind it.
type StatusWriter struct {
	clients *CRDClients
	dryRun  bool
}

// NewStatusWriter returns a StatusWriter using the crd source's clients. Under
// dryRun nothing reaches the provider, so Ready is never Programmed.
func NewStatusWriter(clients *CRDClients, dryRun bool) *StatusWriter {
	return &StatusWriter{clients: clients, dryRun: dryRun}
}

// ReportStatus sets the Ready condition of every DNSEndpoint behind the sync.
func (w *StatusWriter) ReportStatus(ctx context.Context, objects []plan.PlannedObject, applyErr error) {
	for _, obj := range objects {
		if obj.Ref == nil || obj.Ref.Kind() != dnsEndpointKind {
			continue
		}

		key := client.ObjectKey{Namespace: obj.Ref.Namespace(), Name: obj.Ref.Name()}
		dnsEndpoint := &apiv1alpha1.DNSEndpoint{}
		if err := w.clients.Reader().Get(ctx, key, dnsEndpoint); err != nil {
			// Deleted since the plan.
			if !apierrors.IsNotFound(err) {
				log.Warnf("Could not get DNSEndpoint %s: %v", key, err)
			}
			continue
		}

		err := applyErr
		if !obj.Changed {
			// Not in the failed batch: its records are already in place.
			err = nil
		}

		condition := readyCondition(obj.Endpoints, err, w.dryRun)
		condition.ObservedGeneration = dnsEndpoint.Generation
		planned := int32(obj.Endpoints) // #nosec G115 -- bounded by spec.endpoints MaxItems=1000

		UpdateStatus(ctx, w.clients.Writer(), dnsEndpoint, func(status *apiv1alpha1.DNSEndpointStatus) {
			status.Endpoints = planned
			meta.SetStatusCondition(&status.Conditions, condition)
		})
	}
}

// readyCondition describes what became of the endpoints an object contributed.
func readyCondition(planned int, applyErr error, dryRun bool) metav1.Condition {
	condition := metav1.Condition{Type: apiv1alpha1.ReadyCondition}

	switch {
	case planned == 0:
		condition.Status = metav1.ConditionFalse
		condition.Reason = apiv1alpha1.FilteredReason
		condition.Message = "No endpoint reached the DNS provider: --domain-filter, the managed record types or another owner's records excluded all of them"
	case applyErr != nil:
		condition.Status = metav1.ConditionFalse
		condition.Reason = apiv1alpha1.FailedReason
		condition.Message = TruncateConditionMessage(fmt.Sprintf("Provider rejected the batch: %v", applyErr))
	case dryRun:
		condition.Status = metav1.ConditionUnknown
		condition.Reason = apiv1alpha1.DryRunReason
		condition.Message = fmt.Sprintf("%d endpoint(s) planned; --dry-run kept them from the DNS provider", planned)
	default:
		condition.Status = metav1.ConditionTrue
		condition.Reason = apiv1alpha1.ProgrammedReason
		condition.Message = fmt.Sprintf("%d endpoint(s) applied to the DNS provider", planned)
	}

	return condition
}

// UpdateStatus writes the mutated status only if it changed. Accepted and Ready
// are written from a cache that may lag; a stale copy would drop the other
// condition, but its resourceVersion gets it rejected, and only then do we re-read.
func UpdateStatus(ctx context.Context, writer client.Client, dnsEndpoint *apiv1alpha1.DNSEndpoint, mutate func(*apiv1alpha1.DNSEndpointStatus)) {
	updated := dnsEndpoint.DeepCopy()
	mutate(&updated.Status)
	if apiequality.Semantic.DeepEqual(&dnsEndpoint.Status, &updated.Status) {
		return
	}

	err := writer.Status().Update(ctx, updated)
	if apierrors.IsConflict(err) {
		key := client.ObjectKeyFromObject(dnsEndpoint)
		err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
			latest := &apiv1alpha1.DNSEndpoint{}
			if err := writer.Get(ctx, key, latest); err != nil {
				return err
			}
			mutate(&latest.Status)
			return writer.Status().Update(ctx, latest)
		})
	}
	if err != nil {
		log.Warnf("Could not update status of [%s/%s/%s]: %v",
			"dnsendpoint", dnsEndpoint.Namespace, dnsEndpoint.Name, err)
	}
}

// TruncateConditionMessage fits the API server's 32768-character limit on
// Condition.Message, cutting on a rune boundary since it quotes user input.
func TruncateConditionMessage(msg string) string {
	const maxConditionMessageLength = 32768
	if utf8.RuneCountInString(msg) <= maxConditionMessageLength {
		return msg
	}

	runes := []rune(msg)

	return string(runes[:maxConditionMessageLength-3]) + "..."
}
