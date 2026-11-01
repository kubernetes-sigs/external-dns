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
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	crcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/config"
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
// +externaldns:source:namespace=all,single,multiple
// +externaldns:source:fqdn-template=false
// +externaldns:source:events=true
// +externaldns:source:provider-specific=true
// +externaldns:source:validation=spec
type crdSource struct {
	crReader         client.Reader
	crWriter         client.Client // status writes
	informer         crcache.Informer
	listOpts         []client.ListOption
	annotationFilter labels.Selector
	// defaultTargets is set when --default-targets can fill an endpoint without
	// targets; otherwise dedupSource drops it and it must be rejected here.
	defaultTargets bool
	// forceDefaultTargets: the multi-source replaces every target, so spec
	// targets are not checked.
	forceDefaultTargets bool
}

// NewCRDSource creates a new crdSource backed by a controller-runtime cache.
// It builds the scheme, cache, and status-write client from restConfig and cfg.
func NewCRDSource(ctx context.Context, restConfig *rest.Config, cfg *config.Config) (Source, error) {
	opts, err := buildCacheOptions(cfg.Namespaces, cfg.LabelFilter)
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

	cs, err := newCrdSource(ctx, crReader, crWriter, cfg.Namespaces, cfg.LabelFilter, cfg.AnnotationFilter)
	if err != nil {
		return nil, err
	}
	cs.defaultTargets = len(cfg.DefaultTargets) > 0
	cs.forceDefaultTargets = cs.defaultTargets && cfg.ForceDefaultTargets

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
		crdEndpoints, rejections := cs.validateEndpoints(dnsEndpoint)

		endpoint.AttachRefObject(crdEndpoints, events.NewObjectReference(dnsEndpoint, types.CRD))
		endpoints = append(endpoints, crdEndpoints...)

		if err := cs.reportAccepted(ctx, dnsEndpoint, len(crdEndpoints), rejections); err != nil {
			return nil, err
		}
	}

	return endpoint.MergeEndpoints(endpoints), nil
}

// validateEndpoints splits spec into endpoints for the plan and one rejection per
// dropped endpoint. Rejections are shown to users in status conditions.
func (cs *crdSource) validateEndpoints(dnsEndpoint *apiv1alpha1.DNSEndpoint) ([]*endpoint.Endpoint, []string) {
	var (
		accepted   []*endpoint.Endpoint
		rejections []string
	)
	resourceLabel := fmt.Sprintf("crd/%s/%s", dnsEndpoint.Namespace, dnsEndpoint.Name)

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

		if ep.RecordType == endpoint.RecordTypeMX {
			// normalized, else it diffs against the provider's rendering
			for i, target := range ep.Targets {
				ep.Targets[i] = endpoint.NormalizeMXTarget(target)
			}
		}

		if reason := cs.rejectionReason(ep); reason != "" {
			log.Warnf("Endpoint %s/%s with DNSName %s rejected: %s",
				dnsEndpoint.Namespace, dnsEndpoint.Name, ep.DNSName, reason)
			rejections = append(rejections, fmt.Sprintf("spec.endpoints[%d] (%s %s): %s", idx, ep.RecordType, ep.DNSName, reason))
			continue
		}

		ep.WithLabel(endpoint.ResourceLabelKey, resourceLabel)
		accepted = append(accepted, ep)
	}

	return accepted, rejections
}

// rejectionReason explains why ep cannot be planned, or returns "".
//
// It must cover what later stages drop with only a log line (MergeEndpoints,
// CheckEndpoint), or Accepted would lie. Default targets apply in between.
func (cs *crdSource) rejectionReason(ep *endpoint.Endpoint) string {
	if len(ep.Targets) == 0 {
		switch {
		case ep.RecordType == endpoint.RecordTypeCNAME || ep.RecordType == endpoint.RecordTypeDNAME:
			// MergeEndpoints drops these before the multi-source can fill them.
			return fmt.Sprintf("no targets: a %s record needs one target", ep.RecordType)
		case cs.defaultTargets:
			return ""
		}
		return "no targets: set targets, or start external-dns with --default-targets"
	}
	if cs.forceDefaultTargets {
		return ""
	}
	if reason := illegalTargetReason(ep); reason != "" {
		return reason
	}

	return rfcViolationReason(ep)
}

// illegalTargetReason explains the first target whose trailing dot is wrong for
// the record type, or returns "".
func illegalTargetReason(ep *endpoint.Endpoint) string {
	for _, target := range ep.Targets {
		// CNAME/DNAME targets are domain names where a trailing dot is
		// valid (RFC 1035 §5.1 absolute FQDN), so accept both dotted and
		// bare forms.
		if endpoint.RequiresTrailingDot(ep.RecordType) {
			continue
		}
		switch ep.RecordType {
		case endpoint.RecordTypeTXT, endpoint.RecordTypeMX:
			continue // TXT is free-form; rfcViolationReason checks MX
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

// reportAccepted follows Gateway API: an object with some valid endpoints stays
// Accepted, and PartiallyInvalid lists the dropped ones.
func (cs *crdSource) reportAccepted(ctx context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint, accepted int, rejections []string) error {
	condition := metav1.Condition{
		Type:               apiv1alpha1.AcceptedCondition,
		Status:             metav1.ConditionTrue,
		Reason:             apiv1alpha1.AcceptedReason,
		Message:            fmt.Sprintf("%d endpoint(s) accepted", accepted),
		ObservedGeneration: dnsEndpoint.Generation,
	}
	var partial *metav1.Condition

	if len(rejections) > 0 {
		message := strings.Join(rejections, "; ")
		if accepted == 0 {
			condition.Status = metav1.ConditionFalse
			condition.Reason = apiv1alpha1.InvalidReason
			condition.Message = truncateConditionMessage(message)
		} else {
			partial = &metav1.Condition{
				Type:               apiv1alpha1.PartiallyInvalidCondition,
				Status:             metav1.ConditionTrue,
				Reason:             apiv1alpha1.InvalidReason,
				Message:            truncateConditionMessage("Dropped endpoint(s): " + message),
				ObservedGeneration: dnsEndpoint.Generation,
			}
		}
	}

	return cs.updateStatus(ctx, dnsEndpoint, func(status *apiv1alpha1.DNSEndpointStatus) {
		status.ObservedGeneration = dnsEndpoint.Generation
		meta.SetStatusCondition(&status.Conditions, condition)
		if partial != nil {
			meta.SetStatusCondition(&status.Conditions, *partial)
		} else {
			meta.RemoveStatusCondition(&status.Conditions, apiv1alpha1.PartiallyInvalidCondition)
		}
	})
}

// updateStatus writes the mutated status only if it changed. A conflict means
// the cached copy is stale: no retry, the next sync reads the fresh object.
// It fails when the API server prunes the conditions (CRD older than this
// binary): else every sync rewrites every object.
func (cs *crdSource) updateStatus(ctx context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint, mutate func(*apiv1alpha1.DNSEndpointStatus)) error {
	status := dnsEndpoint.Status.DeepCopy()
	mutate(status)
	if apiequality.Semantic.DeepEqual(&dnsEndpoint.Status, status) {
		return nil
	}

	updated := dnsEndpoint.DeepCopy()
	updated.Status = *status
	err := cs.crWriter.Status().Update(ctx, updated)
	switch {
	case apierrors.IsConflict(err):
		log.Debugf("Status of [%s/%s/%s] changed since last sync, will retry on next sync",
			"dnsendpoint", dnsEndpoint.Namespace, dnsEndpoint.Name)
	case err != nil:
		log.Warnf("Could not update status of [%s/%s/%s]: %v",
			"dnsendpoint", dnsEndpoint.Namespace, dnsEndpoint.Name, err)
	case meta.FindStatusCondition(updated.Status.Conditions, apiv1alpha1.AcceptedCondition) == nil:
		return errors.New("the installed DNSEndpoint CRD has no status.conditions: " +
			"apply config/crd/standard/dnsendpoints.externaldns.k8s.io.yaml from the release of this binary")
	}

	return nil
}

// truncateConditionMessage fits the API server's 32768-character limit on
// Condition.Message, cutting on a rune boundary since it quotes user input.
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
	namespaces []string,
	labelSelector, annotationFilter labels.Selector) (*crdSource, error) {
	inf, err := c.GetInformer(ctx, &apiv1alpha1.DNSEndpoint{})
	if err != nil {
		return nil, err
	}

	_, _ = inf.AddEventHandler(informers.DefaultEventHandler())

	// The cache is already scoped to the watched namespaces, so the list is restricted
	// further only when a single one is watched, as it was before multi-namespace support.
	var listOpts []client.ListOption
	if len(namespaces) == 1 {
		listOpts = append(listOpts, client.InNamespace(namespaces[0]))
	}
	if labelSelector != nil && !labelSelector.Empty() {
		listOpts = append(listOpts, client.MatchingLabelsSelector{Selector: labelSelector})
	}

	cs := &crdSource{
		crReader:         c,
		crWriter:         crWriter,
		informer:         inf,
		listOpts:         listOpts,
		annotationFilter: annotationFilter,
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

// buildCacheOptions constructs the controller-runtime cache options for the given
// namespaces and label selector. Extracted so the namespace/label scoping logic can be
// unit-tested without a running API server.
//
// No annotation filter here: dropping an object from the transform empties the whole
// cache since client-go 1.36 (#6728). crdSource.Endpoints filters instead.
func buildCacheOptions(namespaces []string, labelFilter labels.Selector) (crcache.Options, error) {
	scheme := runtime.NewScheme()
	if err := apiv1alpha1.AddToScheme(scheme); err != nil {
		return crcache.Options{}, err
	}

	if len(namespaces) == 0 {
		namespaces = []string{v1.NamespaceAll}
	}
	nsMap := make(map[string]crcache.Config, len(namespaces))
	for _, namespace := range namespaces {
		nsMap[namespace] = crcache.Config{} // "" == NamespaceAll
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
