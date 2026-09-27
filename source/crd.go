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
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	crcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/crd"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/annotations"
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
	crReader         client.Reader
	crWriter         client.Client // status writes
	informer         crcache.Informer
	listOpts         []client.ListOption
	annotationFilter labels.Selector
	emitter          events.EventEmitter // events.Discard unless --events-emit is set
	// defaultTargets is set when --default-targets can fill an endpoint without
	// targets; otherwise dedupSource drops it and it must be rejected here.
	defaultTargets bool
}

// NewCRDSource creates a new crdSource backed by a controller-runtime cache.
// It builds the scheme, cache, and status-write client from restConfig and cfg.
func NewCRDSource(ctx context.Context, restConfig *rest.Config, cfg *Config) (Source, error) {
	namespace := cfg.Namespace()
	opts, err := buildCacheOptions(namespace, cfg.LabelFilter)
	if err != nil {
		return nil, err
	}

	crReader, err := crcache.New(restConfig, opts)
	if err != nil {
		return nil, err
	}

	// crWriter is used exclusively for status writes; reads come from the cache.
	crWriter, err := client.New(restConfig, client.Options{Scheme: opts.Scheme})
	if err != nil {
		return nil, err
	}

	cs, err := newCrdSource(ctx, crReader, crWriter, namespace, cfg.LabelFilter, cfg.AnnotationFilter, cfg.EventEmitter)
	if err != nil {
		return nil, err
	}
	cs.defaultTargets = len(cfg.DefaultTargets) > 0

	return cs, nil
}

func (cs *crdSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("crd: adding event handler")
	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	_, _ = cs.informer.AddEventHandler(eventHandlerFunc(handler))
}

// Endpoints returns endpoint objects for all DNSEndpoint resources visible to this
// source. The cache scopes namespace and labels;
// annotation filtering and target-format validation happen here.
func (cs *crdSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	list := &apiv1alpha1.DNSEndpointList{}
	if err := cs.crReader.List(ctx, list, cs.listOpts...); err != nil {
		return nil, err
	}

	items := make([]*apiv1alpha1.DNSEndpoint, 0, len(list.Items))
	for i := range list.Items {
		items = append(items, &list.Items[i])
	}
	filtered := annotations.Filter(items, cs.annotationFilter)

	endpoints := make([]*endpoint.Endpoint, 0, len(filtered))
	for _, dnsEndpoint := range filtered {
		crdEndpoints, rejections := validateEndpoints(dnsEndpoint, cs.defaultTargets)

		endpoint.AttachRefObject(crdEndpoints, events.NewObjectReference(dnsEndpoint, types.CRD))
		endpoints = append(endpoints, crdEndpoints...)

		cs.reportAccepted(ctx, dnsEndpoint, len(crdEndpoints), rejections)
	}

	return endpoint.MergeEndpoints(endpoints), nil
}

// validateEndpoints splits spec into endpoints for the plan and one rejection per
// dropped endpoint. Rejections surface in Accepted and RecordInvalid, so they are
// written for `kubectl describe`, not for a log line.
func validateEndpoints(dnsEndpoint *apiv1alpha1.DNSEndpoint, defaultTargets bool) ([]*endpoint.Endpoint, []string) {
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

		if reason := rejectionReason(ep, defaultTargets); reason != "" {
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

// rejectionReason explains why ep cannot be planned, or returns "".
//
// It must cover everything dedupSource would later drop with only a log line
// (CheckEndpoint), or Accepted would lie. Empty targets pass with defaultTargets:
// the multi-source fills them before dedupSource runs.
func rejectionReason(ep *endpoint.Endpoint, defaultTargets bool) string {
	if len(ep.Targets) == 0 {
		if defaultTargets {
			return ""
		}
		return "no targets: set targets, or start external-dns with --default-targets"
	}
	if reason := illegalTargetReason(ep); reason != "" {
		return reason
	}

	return rfcViolationReason(ep)
}

// illegalTargetReason explains the first target whose trailing dot is wrong for
// the record type, or returns "".
func illegalTargetReason(ep *endpoint.Endpoint) string {
	for key, target := range ep.Targets {
		// CNAME/DNAME targets are domain names where a trailing dot is
		// valid (RFC 1035 §5.1 absolute FQDN), so accept both dotted and
		// bare forms.
		if endpoint.RequiresTrailingDot(ep.RecordType) {
			continue
		}
		switch ep.RecordType {
		case endpoint.RecordTypeTXT:
			continue // no format constraint on targets
		case endpoint.RecordTypeMX:
			// normalized, else it diffs against the provider's rendering
			ep.Targets[key] = endpoint.NormalizeMXTarget(target)
			continue
		case endpoint.RecordTypeSRV:
			// RFC 2782 requires the dot; rfcViolationReason checks it (#6357).
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

// reportAccepted records whether external-dns understood the spec, and why it
// refused any endpoint.
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
		condition.Message = crd.TruncateConditionMessage(strings.Join(rejections, "; "))
		// Events are never collapsed into a series: emitting on every sync would
		// create one Event per interval, forever, for an untouched spec.
		if verdictChanged(dnsEndpoint.Status.Conditions, condition) {
			cs.emit(dnsEndpoint, condition.Message, events.ActionRejected, events.RecordInvalid)
		}
	}

	crd.UpdateStatus(ctx, cs.crWriter, dnsEndpoint, func(status *apiv1alpha1.DNSEndpointStatus) {
		status.ObservedGeneration = dnsEndpoint.Generation
		meta.SetStatusCondition(&status.Conditions, condition)

		if accepted > 0 {
			return
		}
		// The status writer never sees this object; don't leave a stale Ready.
		status.Endpoints = 0
		if len(rejections) == 0 {
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

// emit sends a Kubernetes event on the DNSEndpoint. It is a no-op unless the
// matching reason was enabled with --events-emit.
func (cs *crdSource) emit(dnsEndpoint *apiv1alpha1.DNSEndpoint, msg string, action events.Action, reason events.Reason) {
	if cs.emitter == nil {
		return
	}
	ref := events.NewObjectReference(dnsEndpoint, types.CRD)
	cs.emitter.Add(events.NewWarningEvent(ref, msg, action, reason))
}

// newCrdSource wires a cache and writer into a running crdSource.
func newCrdSource(
	ctx context.Context,
	c crcache.Cache,
	crWriter client.Client,
	namespace string,
	labelSelector, annotationFilter labels.Selector,
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
		crReader:         c,
		crWriter:         crWriter,
		informer:         inf,
		listOpts:         listOpts,
		annotationFilter: annotationFilter,
		emitter:          emitter,
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
//
// No annotation filter here: dropping an object from the transform empties the whole
// cache since client-go 1.36 (#6728). crdSource.Endpoints filters instead.
func buildCacheOptions(namespace string, labelFilter labels.Selector) (crcache.Options, error) {
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
