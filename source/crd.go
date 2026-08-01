/*
Copyright 2018 The Kubernetes Authors.

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

package source

import (
	"context"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
	crcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/types"
)

// crdSource is an implementation of Source that provides endpoints by listing
// specified CRD and fetching Endpoints embedded in Spec.
//
// +externaldns:source:name=crd
// +externaldns:source:category=ExternalDNS
// +externaldns:source:description=Creates DNS entries from DNSEndpoint CRD resources
// +externaldns:source:resources=DNSEndpoint.externaldns.k8s.io
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=false
// +externaldns:source:events=true
// +externaldns:source:provider-specific=true
type crdSource struct {
	crReader client.Reader
	crWriter client.Client // status writes
	informer crcache.Informer
	listOpts []client.ListOption
	emitter  events.EventEmitter // nil unless --events-emit is set
}

// NewCRDSource creates a new crdSource backed by a controller-runtime cache.
// It builds the scheme, cache, and status-write client from restConfig and cfg.
func NewCRDSource(ctx context.Context, restConfig *rest.Config, cfg *Config) (Source, error) {
	opts, err := buildCacheOptions(cfg.Namespace, cfg.LabelFilter, cfg.AnnotationFilter)
	if err != nil {
		return nil, err
	}

	c, err := crcache.New(restConfig, opts)
	if err != nil {
		return nil, err
	}

	// crWriter is used exclusively for status writes; reads come from the cache.
	crWriter, err := client.New(restConfig, client.Options{Scheme: opts.Scheme})
	if err != nil {
		return nil, err
	}

	return newCrdSource(ctx, c, crWriter, cfg.Namespace, cfg.LabelFilter, cfg.EventEmitter)
}

func (cs *crdSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("crd: adding event handler")
	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	_, _ = cs.informer.AddEventHandler(eventHandlerFunc(handler))
}

// Endpoints returns endpoint objects for all DNSEndpoint resources visible to
// this source. Namespace, label, and annotation filtering are handled at the
// cache level via buildCacheOptions; target-format validation is applied here.
func (cs *crdSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	list := &apiv1alpha1.DNSEndpointList{}
	if err := cs.crReader.List(ctx, list, cs.listOpts...); err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0, len(list.Items))
	for i := range list.Items {
		dnsEndpoint := &list.Items[i]

		crdEndpoints, rejections := validateEndpoints(dnsEndpoint)

		endpoint.AttachRefObject(crdEndpoints, events.NewObjectReference(dnsEndpoint, types.CRD))
		endpoints = append(endpoints, crdEndpoints...)

		cs.reportAccepted(ctx, dnsEndpoint, len(crdEndpoints), rejections)
	}

	return endpoint.MergeEndpoints(endpoints), nil
}

// validateEndpoints splits the endpoints declared in spec into the ones
// external-dns will hand to the plan and a rejection message per dropped
// endpoint. Rejections are what status.conditions[Accepted] and the
// RecordInvalid event report, so the messages are written for a user reading
// `kubectl describe dnsendpoint`, not for a log line.
func validateEndpoints(dnsEndpoint *apiv1alpha1.DNSEndpoint) ([]*endpoint.Endpoint, []string) {
	var (
		accepted   []*endpoint.Endpoint
		rejections []string
	)

	for idx, ep := range dnsEndpoint.Spec.Endpoints {
		if ep == nil {
			log.Debugf(
				"Skipping nil endpoint in DNSEndpoint %s/%s at spec.endpoints",
				dnsEndpoint.Namespace,
				dnsEndpoint.Name,
			)
			rejections = append(rejections, fmt.Sprintf("spec.endpoints[%d]: entry is null", idx))
			continue
		}

		if (ep.RecordType == endpoint.RecordTypeCNAME || ep.RecordType == endpoint.RecordTypeA || ep.RecordType == endpoint.RecordTypeAAAA) && len(ep.Targets) < 1 {
			log.Debugf("Endpoint %s with DNSName %s has an empty list of targets, allowing it to pass through for default-targets processing", dnsEndpoint.Name, ep.DNSName)
		}

		if reason := illegalTargetReason(ep); reason != "" {
			log.Warnf("Endpoint %s/%s with DNSName %s rejected: %s",
				dnsEndpoint.Namespace, dnsEndpoint.Name, ep.DNSName, reason)
			rejections = append(rejections, fmt.Sprintf("spec.endpoints[%d] (%s %s): %s", idx, ep.RecordType, ep.DNSName, reason))
			continue
		}

		ep.WithLabel(endpoint.ResourceLabelKey, fmt.Sprintf("crd/%s/%s", dnsEndpoint.Namespace, dnsEndpoint.Name))
		accepted = append(accepted, ep)
	}

	return accepted, rejections
}

// illegalTargetReason returns a user-facing explanation of the first malformed
// target on ep, or "" when every target is acceptable for its record type.
func illegalTargetReason(ep *endpoint.Endpoint) string {
	for _, target := range ep.Targets {
		switch ep.RecordType {
		case endpoint.RecordTypeTXT, endpoint.RecordTypeMX:
			continue // no format constraint on targets
		case endpoint.RecordTypeCNAME:
			continue // RFC 1035 §5.1: trailing dot denotes an absolute FQDN in zone file notation; both forms are valid
		case endpoint.RecordTypeSRV:
			// SRV targets are "<prio> <weight> <port> <host>"; RFC 2782
			// requires the host to be an absolute FQDN and
			// Endpoint.ValidateSRVRecord enforces the trailing dot.
			// Reject-on-trailing-dot (the default branch below) would
			// loop users between this warning and ValidateSRVRecord's
			// "does not end with a dot" error (#6357).
			continue
		}

		hasDot := strings.HasSuffix(target, ".")

		if ep.RecordType == endpoint.RecordTypeNAPTR {
			if !hasDot {
				return fmt.Sprintf("target %q must be absolute for a NAPTR record — use %q", target, target+".")
			}
			continue
		}

		if hasDot {
			return fmt.Sprintf("target %q must not end with a dot for a %s record — use %q", target, ep.RecordType, strings.TrimSuffix(target, "."))
		}
	}

	return ""
}

// reportAccepted records the source-level verdict on a DNSEndpoint: how many of
// its endpoints entered the plan, and why the others did not. It runs on every
// reconcile but only issues an API write when the computed status differs from
// what is already stored.
func (cs *crdSource) reportAccepted(ctx context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint, accepted int, rejections []string) {
	condition := metav1.Condition{
		Type:               apiv1alpha1.AcceptedCondition,
		Status:             metav1.ConditionTrue,
		Reason:             apiv1alpha1.AcceptedReason,
		Message:            fmt.Sprintf("%d endpoint(s) accepted into the external-dns plan", accepted),
		ObservedGeneration: dnsEndpoint.Generation,
	}

	if len(rejections) > 0 {
		condition.Status = metav1.ConditionFalse
		condition.Reason = apiv1alpha1.InvalidReason
		condition.Message = truncateConditionMessage(strings.Join(rejections, "; "))
		cs.emit(dnsEndpoint, condition.Message, events.ActionRejected, events.RecordInvalid)
	}

	cs.updateStatus(ctx, dnsEndpoint, func(status *apiv1alpha1.DNSEndpointStatus) {
		status.ObservedGeneration = dnsEndpoint.Generation
		status.Endpoints = int32(accepted) // #nosec G115 -- bounded by spec.endpoints MaxItems=1000
		meta.SetStatusCondition(&status.Conditions, condition)
	})
}

// ReportStatus implements StatusReporter. The controller calls it after
// ApplyChanges with every object reference that contributed to the plan, so a
// DNSEndpoint learns whether the provider actually programmed its records.
func (cs *crdSource) ReportStatus(ctx context.Context, refs []*events.ObjectReference, applyErr error) {
	condition := metav1.Condition{
		Type:    apiv1alpha1.ReadyCondition,
		Status:  metav1.ConditionTrue,
		Reason:  apiv1alpha1.ProgrammedReason,
		Message: "Records applied to the DNS provider",
	}
	if applyErr != nil {
		condition.Status = metav1.ConditionFalse
		condition.Reason = apiv1alpha1.FailedReason
		condition.Message = truncateConditionMessage(fmt.Sprintf("Provider rejected the batch: %v", applyErr))
	}

	for _, ref := range refs {
		if ref == nil || ref.Source() != types.CRD {
			continue
		}

		dnsEndpoint := &apiv1alpha1.DNSEndpoint{}
		key := client.ObjectKey{Namespace: ref.Namespace(), Name: ref.Name()}
		if err := cs.crReader.Get(ctx, key, dnsEndpoint); err != nil {
			// The object may have been deleted between the plan and the apply.
			log.Debugf("Could not read dnsendpoint %s to report sync status: %v", key, err)
			continue
		}

		condition.ObservedGeneration = dnsEndpoint.Generation
		cs.updateStatus(ctx, dnsEndpoint, func(status *apiv1alpha1.DNSEndpointStatus) {
			meta.SetStatusCondition(&status.Conditions, condition)
		})
	}
}

// updateStatus applies mutate to the object's status and writes it back only if
// that produced a change. meta.SetStatusCondition leaves LastTransitionTime
// alone when a condition's status is unchanged, so an unchanging DNSEndpoint
// costs no API writes across reconciles.
//
// The Accepted and Ready conditions are written at different points of a sync
// and the read path is a cache that may lag the last write, so the update is a
// read-modify-write against the API server: mutating a stale copy would drop
// whichever condition the other writer had just set.
func (cs *crdSource) updateStatus(ctx context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint, mutate func(*apiv1alpha1.DNSEndpointStatus)) {
	probe := dnsEndpoint.Status.DeepCopy()
	mutate(probe)
	if apiequality.Semantic.DeepEqual(&dnsEndpoint.Status, probe) {
		return
	}

	key := client.ObjectKeyFromObject(dnsEndpoint)
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		latest := &apiv1alpha1.DNSEndpoint{}
		if err := cs.crWriter.Get(ctx, key, latest); err != nil {
			return err
		}
		mutate(&latest.Status)
		return cs.crWriter.Status().Update(ctx, latest)
	})
	if err != nil {
		log.Warnf("Could not update status of [%s/%s/%s]: %v",
			"dnsendpoint", dnsEndpoint.Namespace, dnsEndpoint.Name, err)
	}
}

// emit sends a Kubernetes event on the DNSEndpoint. It is a no-op unless the
// matching reason was enabled with --events-emit.
func (cs *crdSource) emit(dnsEndpoint *apiv1alpha1.DNSEndpoint, msg string, action events.Action, reason events.Reason) {
	if cs.emitter == nil {
		return
	}
	ref := events.NewObjectReference(dnsEndpoint, types.CRD)
	cs.emitter.Add(events.NewWarningEvent(ref, msg, action, reason))
}

// truncateConditionMessage keeps a message within the 32768-character limit the
// API server enforces on Condition.Message.
func truncateConditionMessage(msg string) string {
	const maxConditionMessageLength = 32768
	if len(msg) <= maxConditionMessageLength {
		return msg
	}
	return msg[:maxConditionMessageLength-3] + "..."
}

// newCrdSource wires a cache and writer into a running crdSource.
func newCrdSource(
	ctx context.Context,
	c crcache.Cache,
	crWriter client.Client,
	namespace string,
	labelSelector labels.Selector,
	emitter events.EventEmitter) (*crdSource, error) {
	inf, err := c.GetInformer(ctx, &apiv1alpha1.DNSEndpoint{})
	if err != nil {
		return nil, err
	}

	_, _ = inf.AddEventHandler(informers.DefaultEventHandler())

	listOpts := []client.ListOption{client.InNamespace(namespace)}
	if labelSelector != nil && !labelSelector.Empty() {
		listOpts = append(listOpts, client.MatchingLabelsSelector{Selector: labelSelector})
	}

	cs := &crdSource{
		crReader: c,
		crWriter: crWriter,
		informer: inf,
		listOpts: listOpts,
		emitter:  emitter,
	}

	if err := startAndSync(ctx, c); err != nil {
		return nil, err
	}

	return cs, nil
}

// startAndSync starts the cache in a goroutine and waits for it to sync.
// Returns an error if the cache fails to start or sync.
func startAndSync(ctx context.Context, c crcache.Cache) error {
	errCh := make(chan error, 1)
	go func() { errCh <- c.Start(ctx) }()
	if !c.WaitForCacheSync(ctx) {
		select {
		case err := <-errCh:
			if err != nil {
				return fmt.Errorf("cache failed to sync: %w", err)
			}
			return fmt.Errorf("cache failed to sync")
		case <-ctx.Done():
			return fmt.Errorf("cache failed to sync: %w", ctx.Err())
		}
	}
	return nil
}

// buildCacheOptions constructs the controller-runtime cache options for the
// given namespace and label selector. Extracted so the namespace/label scoping
// logic can be unit-tested without a running API server.
func buildCacheOptions(namespace string, labelFilter, annotationSelector labels.Selector) (crcache.Options, error) {
	scheme := runtime.NewScheme()
	if err := apiv1alpha1.AddToScheme(scheme); err != nil {
		return crcache.Options{}, err
	}
	// metav1.AddToGroupVersion registers ListOptions (and other meta types) under
	// apiv1alpha1.GroupVersion so that runtime.NewParameterCodec can encode them
	// as URL parameters when building watch requests for this group.
	metav1.AddToGroupVersion(scheme, apiv1alpha1.GroupVersion)

	nsMap := map[string]crcache.Config{
		namespace: {}, // "" == NamespaceAll
	}
	byObj := crcache.ByObject{
		Namespaces: nsMap,
		Transform: informers.TransformerWithOptions[*apiv1alpha1.DNSEndpoint](
			informers.TransformRemoveManagedFields(),
			informers.TransformRemoveLastAppliedConfig(),
			informers.TransformRequireAnnotation(annotationSelector),
		),
	}
	if labelFilter != nil && !labelFilter.Empty() {
		byObj.Label = labelFilter
	}
	return crcache.Options{
		Scheme: scheme,
		ByObject: map[client.Object]crcache.ByObject{
			&apiv1alpha1.DNSEndpoint{}: byObj,
		},
	}, nil
}
