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

package crd

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"maps"
	"regexp"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	crcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	kubeclient "sigs.k8s.io/external-dns/pkg/client"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
)

// CRDRegistry implements registry interface with ownership implemented via associated custom resource records (DNSRecord)
type CRDRegistry struct {
	// crReader serves the bulk Records() list from a controller-runtime cache
	// (informer-backed), keeping the per-reconcile list off the API server.
	crReader client.Reader
	// crWriter performs create/update/delete against the API server. It also
	// serves every get-before-write read (getDNSRecord), so a record written
	// this reconcile is visible to a read later in the same reconcile — the
	// informer cache behind crReader may still lag and must not back writes.
	crWriter client.Client

	namespace string
	provider  provider.Provider
	ownerID   string // refers to the owner id of the current instance
}

func New(cfg *externaldns.Config, p provider.Provider) (registry.Registry, error) {
	return NewCRDRegistry(p, cfg.KubeConfig, cfg.APIServerURL, cfg.Namespace, cfg.TXTOwnerID, cfg.RequestTimeout)
}

// NewCRDRegistry returns new CRDRegistry object backed by a controller-runtime
// cache (reads) and client (writes).
func NewCRDRegistry(provider provider.Provider, kubeConfig, apiServerURL, namespace, ownerID string, apiServerTimeout time.Duration) (*CRDRegistry, error) {
	if ownerID == "" {
		return nil, errors.New("owner id cannot be empty")
	}

	if namespace == "" {
		log.Info("Registry: namespace not specified, using `default`")
		namespace = "default"
	}

	// Build the REST config from the shared client package: this registry may run
	// against a remote, shared cluster, so it does not reuse the source clients.
	restConfig, err := kubeclient.InstrumentedRESTConfig(kubeConfig, apiServerURL, apiServerTimeout, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("unable to build rest config: %w", err)
	}

	opts, err := buildCacheOptions(namespace)
	if err != nil {
		return nil, err
	}

	c, err := crcache.New(restConfig, opts)
	if err != nil {
		return nil, fmt.Errorf("unable to create cache: %w", err)
	}

	// crWriter is used for writes and for every get-before-write read; the
	// cache serves only the bulk Records() list.
	crWriter, err := client.New(restConfig, client.Options{Scheme: opts.Scheme})
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	// The cache lives for the whole process lifetime, mirroring the registry.
	ctx := context.Background()
	// GetInformer registers the DNSRecord watch that backs crReader; the returned
	// informer is unused because reads are always served from the synced cache.
	if _, err := c.GetInformer(ctx, &apiv1alpha1.DNSRecord{}); err != nil {
		return nil, fmt.Errorf("unable to get informer: %w", err)
	}
	// Bound the initial sync so a missing RBAC or unreachable API server fails
	// fast at startup instead of hanging forever. The cache itself keeps running
	// under ctx; only the wait is bounded.
	syncTimeout := apiServerTimeout
	if syncTimeout <= 0 {
		syncTimeout = 30 * time.Second
	}
	syncCtx, cancel := context.WithTimeout(ctx, syncTimeout)
	defer cancel()
	if err := startAndSync(ctx, syncCtx, c); err != nil {
		return nil, err
	}

	return &CRDRegistry{
		crReader:  c,
		crWriter:  crWriter,
		namespace: namespace,
		provider:  provider,
		ownerID:   ownerID,
	}, nil
}

func (cr *CRDRegistry) GetDomainFilter() endpoint.DomainFilterInterface {
	return cr.provider.GetDomainFilter()
}

func (cr *CRDRegistry) OwnerID() string {
	return cr.ownerID
}

// Records returns the current records from the registry
func (cr *CRDRegistry) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var records apiv1alpha1.DNSRecordList
	if err := cr.crReader.List(ctx, &records,
		client.InNamespace(cr.namespace),
		client.MatchingLabels{apiv1alpha1.RecordOwnerLabel: cr.ownerID},
	); err != nil {
		return []*endpoint.Endpoint{}, err
	}

	endpoints := make([]*endpoint.Endpoint, 0, len(records.Items))
	for i := range records.Items {
		// Only records confirmed applied to the provider represent current
		// state. An un-programmed record — one left Accepted by a provider
		// failure — is skipped so the plan re-applies it on the next reconcile
		// instead of mistaking it for a record that already exists.
		if !meta.IsStatusConditionTrue(records.Items[i].Status.Conditions, apiv1alpha1.ReadyCondition) {
			continue
		}
		endpoints = append(endpoints, &records.Items[i].Spec.Endpoint)
	}
	return endpoints, nil
}

// ApplyChanges updates dns provider with the changes and creates/updates/delete a DNSRecord accordingly.
//
// Intent is recorded first: the DNSRecord objects for created and updated
// endpoints are written and marked Accepted before the provider is called. Their
// Programmed condition is then set from the provider outcome — True on success,
// False with reason Failed on error. Records() treats only Programmed records as
// current state, so an object left Accepted by a provider failure is re-applied
// on the next reconcile rather than mistaken for an existing record — the same
// safety the apply-first ordering gave, while making the full lifecycle
// (Accepted, Programmed, Failed) visible on the object.
//
// The provider reports a single batch error and cannot attribute it to
// individual records, so on failure every record in the batch is marked Failed.
// Provider applies are idempotent, so records that did get applied are corrected
// to Programmed on the next reconcile.
func (cr *CRDRegistry) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	filteredChanges := &plan.Changes{
		Create:    changes.Create,
		UpdateNew: endpoint.FilterEndpointsByOwnerID(cr.ownerID, changes.UpdateNew),
		UpdateOld: endpoint.FilterEndpointsByOwnerID(cr.ownerID, changes.UpdateOld),
		Delete:    endpoint.FilterEndpointsByOwnerID(cr.ownerID, changes.Delete),
	}

	// Stamp the owner label before applying so the provider sees owner-labeled
	// endpoints, matching the TXT registry behavior.
	for _, r := range filteredChanges.Create {
		r.WithLabel(endpoint.OwnerLabelKey, cr.ownerID)
	}

	// Write the DNSRecord objects and collect them so their Ready condition can be
	// set from the provider outcome below.
	applied := make([]*apiv1alpha1.DNSRecord, 0, len(filteredChanges.Create)+len(filteredChanges.UpdateNew))

	for _, r := range filteredChanges.Create {
		dnsrecord, err := cr.ensureDNSRecord(ctx, r)
		if err != nil {
			return err
		}
		applied = append(applied, dnsrecord)
	}

	// Update existing DNS records to reflect the newest change.
	for i, e := range filteredChanges.UpdateNew {
		old := filteredChanges.UpdateOld[i]
		dnsrecord, err := cr.getDNSRecord(ctx, old)
		if err != nil {
			return fmt.Errorf("unable to get DNSRecord of %s: %w", old, err)
		}
		if dnsrecord == nil {
			continue
		}
		dnsrecord.Spec.Endpoint = *e
		if err := cr.crWriter.Update(ctx, dnsrecord); err != nil {
			return fmt.Errorf("unable to update DNSRecord %s in %s: %w", dnsrecord.Name, dnsrecord.Namespace, err)
		}
		applied = append(applied, dnsrecord)
	}

	// Record intent before calling the provider: each collected record is marked
	// Ready=False/Accepted until the provider confirms it.
	for _, dnsrecord := range applied {
		cr.setStatus(ctx, dnsrecord, apiv1alpha1.AcceptedReason, "Endpoint accepted by external-dns; not yet programmed")
	}

	if err := cr.provider.ApplyChanges(ctx, filteredChanges); err != nil {
		// The provider reports a single batch error and cannot attribute it to
		// individual records, so every record in the batch is marked Failed; the
		// apply is idempotent, so records that were in fact applied are corrected
		// to Programmed on the next reconcile.
		for _, dnsrecord := range applied {
			cr.setStatus(ctx, dnsrecord, apiv1alpha1.FailedReason, fmt.Sprintf("Provider rejected the batch: %v", err))
		}
		return fmt.Errorf("provider cannot apply changes: %w", err)
	}

	// Provider accepted the changes; mark the records Programmed.
	for _, dnsrecord := range applied {
		cr.setStatus(ctx, dnsrecord, apiv1alpha1.ProgrammedReason, "Endpoint applied to the DNS provider")
	}

	// Deletes are reconciled last: the DNS record is gone from the provider, so
	// drop its DNSRecord too.
	for _, r := range filteredChanges.Delete {
		dnsrecord, err := cr.getDNSRecord(ctx, r)
		if err != nil {
			return fmt.Errorf("unable to get DNSRecord of %s: %w", r, err)
		}
		if dnsrecord == nil {
			continue
		}
		if err := cr.crWriter.Delete(ctx, dnsrecord); err != nil {
			// Ignore not found as it's a benign error, the record isn't present and it's the end goal here, to remove
			// all records. All other errors should surface back to the user.
			if !k8sErrors.IsNotFound(err) {
				return fmt.Errorf("unable to delete DNSRecord %s in %s: %w", dnsrecord.Name, cr.namespace, err)
			}
		}
	}

	return cr.adjustLabelsFromProvider(ctx, applied)
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (cr *CRDRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return cr.provider.AdjustEndpoints(endpoints)
}

// getDNSRecord retrieves the DNSRecord backing the given endpoint by its
// deterministic object name. It reads through crWriter (live API), not the
// cache, so a record written earlier in the same reconcile is seen here. It
// returns (nil, nil) when no such record exists.
func (cr *CRDRegistry) getDNSRecord(ctx context.Context, record *endpoint.Endpoint) (*apiv1alpha1.DNSRecord, error) {
	var dnsrecord apiv1alpha1.DNSRecord
	key := client.ObjectKey{Namespace: cr.namespace, Name: recordObjectName(record)}
	if err := cr.crWriter.Get(ctx, key, &dnsrecord); err != nil {
		if k8sErrors.IsNotFound(err) {
			// A missing DNSRecord is an expected, benign result; callers treat a
			// nil record as "not found".
			return nil, nil //nolint:nilnil // intentional not-found sentinel
		}
		return nil, err
	}
	return &dnsrecord, nil
}

// nameInvalidChars matches every character that is not allowed in the readable
// part of a DNSRecord object name (anything outside lowercase RFC 1123 alphanumerics).
var nameInvalidChars = regexp.MustCompile(`[^a-z0-9]+`)

// recordObjectName builds a deterministic, RFC 1123 compliant object name for a
// DNSRecord. The dashed DNS name is kept as a readable prefix so records remain
// discoverable with `kubectl get dnsrecords`, while a short hash of the full
// identity (DNS name, record type, set identifier) guarantees uniqueness for
// records that would otherwise collide after sanitization — e.g. distinct set
// identifiers (weighted/latency policies), `sub.example.io` vs `sub-example.io`,
// or names truncated to the length limit.
func recordObjectName(record *endpoint.Endpoint) string {
	// Canonicalize the identity so equivalent records (case, trailing dot) map to
	// the same object name on both the write and the provider-readback path.
	dnsName := strings.TrimSuffix(strings.ToLower(record.DNSName), ".")
	recordType := strings.ToUpper(record.RecordType)
	hash := sha256.Sum256([]byte(strings.Join([]string{dnsName, recordType, record.SetIdentifier}, "/")))
	suffix := fmt.Sprintf("-%x", hash[:4]) // 8 hex characters

	base := strings.Trim(nameInvalidChars.ReplaceAllString(dnsName, "-"), "-")
	// Keep the whole name within the 253 character RFC 1123 subdomain limit.
	if maxBase := 253 - len(suffix); len(base) > maxBase {
		base = strings.Trim(base[:maxBase], "-")
	}
	if base == "" {
		base = "record"
	}
	return base + suffix
}

// ensureDNSRecord creates the DNSRecord backing the endpoint, or returns the
// existing one — refreshed to the desired endpoint — when a previous reconcile
// already created it (e.g. a provider failure left it Accepted but not
// Programmed). The returned object always reflects the desired endpoint.
func (cr *CRDRegistry) ensureDNSRecord(ctx context.Context, record *endpoint.Endpoint) (*apiv1alpha1.DNSRecord, error) {
	dnsrecord := &apiv1alpha1.DNSRecord{
		ObjectMeta: metav1.ObjectMeta{
			Name:      recordObjectName(record),
			Namespace: cr.namespace,
			Labels: map[string]string{
				apiv1alpha1.RecordOwnerLabel: cr.OwnerID(),
			},
		},
		Spec: apiv1alpha1.DNSRecordSpec{
			Endpoint: *record,
		},
	}

	if err := cr.crWriter.Create(ctx, dnsrecord); err != nil {
		if !k8sErrors.IsAlreadyExists(err) {
			return nil, fmt.Errorf("unable to create DNSRecord %s in %s: %w", dnsrecord.Name, dnsrecord.Namespace, err)
		}
		// A previous reconcile already created the object; fetch it and refresh
		// its spec so the status we set below applies to the current endpoint.
		existing, err := cr.getDNSRecord(ctx, record)
		if err != nil {
			return nil, fmt.Errorf("unable to get DNSRecord %s in %s: %w", dnsrecord.Name, cr.namespace, err)
		}
		if existing == nil {
			return nil, fmt.Errorf("DNSRecord %s in %s reported as existing but was not found", dnsrecord.Name, cr.namespace)
		}
		existing.Spec.Endpoint = *record
		if err := cr.crWriter.Update(ctx, existing); err != nil {
			return nil, fmt.Errorf("unable to update DNSRecord %s in %s: %w", existing.Name, cr.namespace, err)
		}
		return existing, nil
	}
	return dnsrecord, nil
}

// setStatus sets the Ready condition on the DNSRecord with the given reason and
// message, and persists the status subresource. Ready is True only for
// ProgrammedReason (the endpoint is live in the provider); every other reason
// leaves it False. Status is best-effort observability: a failure to write it is
// logged but never fails reconciliation, since the DNS record itself is already
// applied (or its failure already surfaced through the apply error).
func (cr *CRDRegistry) setStatus(ctx context.Context, dnsrecord *apiv1alpha1.DNSRecord, reason, message string) {
	status := metav1.ConditionFalse
	if reason == apiv1alpha1.ProgrammedReason {
		status = metav1.ConditionTrue
	}
	meta.SetStatusCondition(&dnsrecord.Status.Conditions, metav1.Condition{
		Type:    apiv1alpha1.ReadyCondition,
		Status:  status,
		Reason:  reason,
		Message: message,
	})
	if err := cr.crWriter.Status().Update(ctx, dnsrecord); err != nil {
		log.Warnf("unable to update status of DNSRecord %s in %s: %v", dnsrecord.Name, cr.namespace, err)
	}
}

// adjustLabelsFromProvider reconciles the labels of the records applied this
// reconcile with the labels the provider ended up storing (some providers, e.g.
// coredns, rewrite them). Only the just-applied records are considered: records
// untouched this round were not changed by the provider either. It is a no-op
// when nothing was created or updated, avoiding an extra provider read on
// delete-only or empty reconciles. It should be called after applyChanges.
func (cr *CRDRegistry) adjustLabelsFromProvider(ctx context.Context, applied []*apiv1alpha1.DNSRecord) error {
	if len(applied) == 0 {
		return nil
	}

	records, err := cr.provider.Records(ctx)
	if err != nil {
		return fmt.Errorf("unable to get records from provider: %w", err)
	}

	// Index the provider records by the deterministic DNSRecord object name so
	// each applied record is matched in memory, without a per-record API read.
	byName := make(map[string]*endpoint.Endpoint, len(records))
	for _, record := range records {
		byName[recordObjectName(record)] = record
	}

	for _, dnsrecord := range applied {
		record, ok := byName[dnsrecord.Name]
		if !ok {
			log.Debugf("no provider record matched DNSRecord %s; skipping label adjust", dnsrecord.Name)
			continue
		}
		if !maps.Equal(dnsrecord.Spec.Endpoint.Labels, record.Labels) {
			log.Debug("update DNSRecord with modified labels from provider")
			if err := cr.updateDNSRecordWithEndpointLabels(ctx, dnsrecord, record); err != nil {
				return err
			}
		}
	}

	return nil
}

func (cr *CRDRegistry) updateDNSRecordWithEndpointLabels(ctx context.Context, dnsrecord *apiv1alpha1.DNSRecord, record *endpoint.Endpoint) error {
	dnsrecord.MergeProviderLabels(record.Labels, cr.ownerID)
	if err := cr.crWriter.Update(ctx, dnsrecord); err != nil {
		return fmt.Errorf("unable to update DNSRecord %s: %w", dnsrecord.Name, err)
	}
	return nil
}

// buildCacheOptions constructs the controller-runtime cache options scoped to
// the given namespace, with the DNSRecord type registered in the scheme.
func buildCacheOptions(namespace string) (crcache.Options, error) {
	scheme := runtime.NewScheme()
	if err := apiv1alpha1.AddToScheme(scheme); err != nil {
		return crcache.Options{}, err
	}
	// metav1.AddToGroupVersion registers ListOptions (and other meta types) so
	// that watch/list requests for this group can be encoded.
	metav1.AddToGroupVersion(scheme, apiv1alpha1.GroupVersion)

	return crcache.Options{
		Scheme: scheme,
		ByObject: map[client.Object]crcache.ByObject{
			&apiv1alpha1.DNSRecord{}: {
				Namespaces: map[string]crcache.Config{namespace: {}},
			},
		},
	}, nil
}

// startAndSync starts the cache under startCtx (process lifetime) and waits for
// the initial sync under syncCtx (bounded), returning an error if the cache
// fails to start or does not sync before syncCtx expires.
func startAndSync(startCtx, syncCtx context.Context, c crcache.Cache) error {
	errCh := make(chan error, 1)
	go func() { errCh <- c.Start(startCtx) }()
	if !c.WaitForCacheSync(syncCtx) {
		select {
		case err := <-errCh:
			if err != nil {
				return fmt.Errorf("cache failed to sync: %w", err)
			}
			return fmt.Errorf("cache failed to sync")
		default:
			return fmt.Errorf("cache failed to sync: %w", syncCtx.Err())
		}
	}
	return nil
}
