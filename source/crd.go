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
	"unicode/utf8"

	log "github.com/sirupsen/logrus"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
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
	emitter  events.EventEmitter // events.Discard unless --events-emit is set
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

		if reason := rejectionReason(ep); reason != "" {
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

// rejectionReason returns a user-facing explanation of why external-dns refuses
// to plan ep, or "" when it is acceptable.
//
// It has to cover everything the pipeline would drop later: dedupSource re-runs
// CheckEndpoint and discards what fails with only a log line, and an endpoint
// discarded there never reaches the plan, so ReportStatus never sees its object.
// Rejecting it here keeps Accepted authoritative.
func rejectionReason(ep *endpoint.Endpoint) string {
	if reason := illegalTargetReason(ep); reason != "" {
		return reason
	}

	return rfcViolationReason(ep)
}

// illegalTargetReason returns a user-facing explanation of the first target on ep
// whose trailing dot does not match what the record type expects, or "" when they
// all do.
func illegalTargetReason(ep *endpoint.Endpoint) string {
	for _, target := range ep.Targets {
		switch ep.RecordType {
		case endpoint.RecordTypeTXT, endpoint.RecordTypeMX:
			continue // no format constraint on targets
		case endpoint.RecordTypeCNAME:
			continue // RFC 1035 §5.1: trailing dot denotes an absolute FQDN in zone file notation; both forms are valid
		case endpoint.RecordTypeSRV:
			// RFC 2782 wants an absolute host, so here the trailing dot is
			// mandatory rather than illegal — rejecting it below would loop users
			// between the two errors (#6357). rfcViolationReason checks the whole
			// grammar, dot included.
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

// rfcViolationReason turns an Endpoint.CheckEndpoint failure into a message that
// names the grammar the record type expects. It returns "" when ep passes.
func rfcViolationReason(ep *endpoint.Endpoint) string {
	if ep.CheckEndpoint() {
		return ""
	}

	switch ep.RecordType {
	case endpoint.RecordTypeA, endpoint.RecordTypeAAAA:
		family := 4
		if ep.RecordType == endpoint.RecordTypeAAAA {
			family = 6
		}
		return fmt.Sprintf("targets of a %s record must be IPv%d addresses, unless the endpoint sets the %q provider-specific property for a provider-native alias",
			ep.RecordType, family, endpoint.ProviderSpecificAlias)
	case endpoint.RecordTypeMX:
		return `MX targets must be "<preference> <host>", e.g. "10 mail.example.com"`
	case endpoint.RecordTypeSRV:
		return `SRV targets must be "<priority> <weight> <port> <host>" with an absolute host, e.g. "10 5 5060 sip.example.com."`
	case endpoint.RecordTypePTR:
		return "a PTR record needs a dnsName under .in-addr.arpa or .ip6.arpa and at least one non-empty target"
	}

	return fmt.Sprintf("a %s record does not support the %q provider-specific property", ep.RecordType, endpoint.ProviderSpecificAlias)
}

// reportAccepted records the source-level verdict on a DNSEndpoint: whether
// external-dns understood its spec, and why it refused any endpoint. The plan
// outcome lands later, in ReportStatus. It runs on every reconcile but only
// issues an API write when the computed status differs from what is stored.
func (cs *crdSource) reportAccepted(ctx context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint, accepted int, rejections []string) {
	condition := metav1.Condition{
		Type:               apiv1alpha1.AcceptedCondition,
		Status:             metav1.ConditionTrue,
		Reason:             apiv1alpha1.AcceptedReason,
		Message:            fmt.Sprintf("%d endpoint(s) accepted", accepted),
		ObservedGeneration: dnsEndpoint.Generation,
	}

	if len(rejections) > 0 {
		condition.Status = metav1.ConditionFalse
		condition.Reason = apiv1alpha1.InvalidReason
		condition.Message = truncateConditionMessage(strings.Join(rejections, "; "))
		// Events carry a timestamped name, so nothing collapses them into a series.
		// external-dns re-lists on a timer: emitting unconditionally would create
		// one Event per sync interval, forever, for a spec nobody touched.
		if verdictChanged(dnsEndpoint.Status.Conditions, condition) {
			cs.emit(dnsEndpoint, condition.Message, events.ActionRejected, events.RecordInvalid)
		}
	}

	cs.updateStatus(ctx, dnsEndpoint, func(status *apiv1alpha1.DNSEndpointStatus) {
		status.ObservedGeneration = dnsEndpoint.Generation
		meta.SetStatusCondition(&status.Conditions, condition)

		if accepted > 0 {
			return
		}
		// Nothing reached the plan, so ReportStatus will never be called for this
		// object. Left alone, the count and Ready would keep advertising the last
		// reconcile that did produce endpoints.
		status.Endpoints = 0
		if len(rejections) == 0 {
			// An empty spec: nothing to be ready about.
			meta.RemoveStatusCondition(&status.Conditions, apiv1alpha1.ReadyCondition)
			return
		}
		meta.SetStatusCondition(&status.Conditions, metav1.Condition{
			Type:               apiv1alpha1.ReadyCondition,
			Status:             metav1.ConditionFalse,
			Reason:             apiv1alpha1.InvalidReason,
			Message:            "No endpoint reached the DNS provider: every endpoint in spec was rejected",
			ObservedGeneration: dnsEndpoint.Generation,
		})
	})
}

// verdictChanged reports whether condition says anything stored does not already say.
func verdictChanged(stored []metav1.Condition, condition metav1.Condition) bool {
	current := meta.FindStatusCondition(stored, condition.Type)

	return current == nil || current.Status != condition.Status || current.Reason != condition.Reason || current.Message != condition.Message
}

// ReportStatus implements StatusReporter. The controller calls it once per sync
// with every object that contributed a desired endpoint, so a DNSEndpoint learns
// how many of its endpoints external-dns planned and how the provider answered.
func (cs *crdSource) ReportStatus(ctx context.Context, objects []PlannedObject, applyErr error) {
	for _, obj := range objects {
		if obj.Ref == nil || obj.Ref.Source() != types.CRD {
			continue
		}

		dnsEndpoint := &apiv1alpha1.DNSEndpoint{}
		key := client.ObjectKey{Namespace: obj.Ref.Namespace(), Name: obj.Ref.Name()}
		if err := cs.crReader.Get(ctx, key, dnsEndpoint); err != nil {
			// The object may have been deleted between the plan and the apply.
			log.Debugf("Could not read dnsendpoint %s to report sync status: %v", key, err)
			continue
		}

		condition := readyCondition(obj.Endpoints, applyErr)
		condition.ObservedGeneration = dnsEndpoint.Generation
		planned := int32(obj.Endpoints) // #nosec G115 -- bounded by spec.endpoints MaxItems=1000

		cs.updateStatus(ctx, dnsEndpoint, func(status *apiv1alpha1.DNSEndpointStatus) {
			status.Endpoints = planned
			meta.SetStatusCondition(&status.Conditions, condition)
		})
	}
}

// readyCondition describes what became of the endpoints an object contributed:
// excluded by the filters before reaching the provider, rejected by it, or
// programmed.
func readyCondition(planned int, applyErr error) metav1.Condition {
	condition := metav1.Condition{Type: apiv1alpha1.ReadyCondition}

	switch {
	case planned == 0:
		condition.Status = metav1.ConditionFalse
		condition.Reason = apiv1alpha1.FilteredReason
		condition.Message = "No endpoint reached the DNS provider: --domain-filter or the managed record types excluded all of them"
	case applyErr != nil:
		condition.Status = metav1.ConditionFalse
		condition.Reason = apiv1alpha1.FailedReason
		condition.Message = truncateConditionMessage(fmt.Sprintf("Provider rejected the batch: %v", applyErr))
	default:
		condition.Status = metav1.ConditionTrue
		condition.Reason = apiv1alpha1.ProgrammedReason
		condition.Message = fmt.Sprintf("%d endpoint(s) applied to the DNS provider", planned)
	}

	return condition
}

// updateStatus applies mutate to the object's status and writes it back only if
// that produced a change, so an unchanging DNSEndpoint costs no API writes.
//
// Accepted and Ready are written at different points of a sync, and the read path
// is a cache that may lag. Pushing a stale copy would drop whichever condition the
// other writer just set — but it also carries a stale resourceVersion, so the API
// server rejects it and only then is a re-read worth its round trip.
func (cs *crdSource) updateStatus(ctx context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint, mutate func(*apiv1alpha1.DNSEndpointStatus)) {
	updated := dnsEndpoint.DeepCopy()
	mutate(&updated.Status)
	if apiequality.Semantic.DeepEqual(&dnsEndpoint.Status, &updated.Status) {
		return
	}

	err := cs.crWriter.Status().Update(ctx, updated)
	if apierrors.IsConflict(err) {
		key := client.ObjectKeyFromObject(dnsEndpoint)
		err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
			latest := &apiv1alpha1.DNSEndpoint{}
			if err := cs.crWriter.Get(ctx, key, latest); err != nil {
				return err
			}
			mutate(&latest.Status)
			return cs.crWriter.Status().Update(ctx, latest)
		})
	}
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
// API server enforces on Condition.Message. It cuts on a rune boundary: the
// message quotes user-supplied names, so slicing bytes could store mojibake.
func truncateConditionMessage(msg string) string {
	const maxConditionMessageLength = 32768
	if utf8.RuneCountInString(msg) <= maxConditionMessageLength {
		return msg
	}

	runes := []rune(msg)

	return string(runes[:maxConditionMessageLength-3]) + "..."
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
