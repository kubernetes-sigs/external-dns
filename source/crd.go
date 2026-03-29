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
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	crcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
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
	crReader client.Reader
	crWriter client.Client // status writes
	informer crcache.Informer
	listOpts []client.ListOption
}

// NewCRDSource creates a new crdSource backed by a controller-runtime cache.
// It builds the scheme, cache, and status-write client from restConfig and cfg.
func NewCRDSource(ctx context.Context, restConfig *rest.Config, cfg *Config) (Source, error) {
	annotationSelector, err := annotations.ParseFilter(cfg.AnnotationFilter)
	if err != nil {
		return nil, err
	}
	opts, err := buildCacheOptions(cfg.Namespace, cfg.LabelFilter, annotationSelector)
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

	return newCrdSource(ctx, c, crWriter, cfg.Namespace, cfg.LabelFilter)
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
		var crdEndpoints []*endpoint.Endpoint
		for _, ep := range dnsEndpoint.Spec.Endpoints {
			if ep == nil {
				log.Debugf(
					"Skipping nil endpoint in DNSEndpoint %s/%s at spec.endpoints",
					dnsEndpoint.Namespace,
					dnsEndpoint.Name,
				)
				continue
			}

			if (ep.RecordType == endpoint.RecordTypeCNAME || ep.RecordType == endpoint.RecordTypeA || ep.RecordType == endpoint.RecordTypeAAAA) && len(ep.Targets) < 1 {
				log.Debugf("Endpoint %s with DNSName %s has an empty list of targets, allowing it to pass through for default-targets processing", dnsEndpoint.Name, ep.DNSName)
			}
			illegalTarget := false
			for _, target := range ep.Targets {
				switch ep.RecordType {
				case endpoint.RecordTypeTXT, endpoint.RecordTypeMX:
					continue // no format constraint on targets
				case endpoint.RecordTypeCNAME:
					continue // RFC 1035 §5.1: trailing dot denotes an absolute FQDN in zone file notation; both forms are valid
				}

				hasDot := strings.HasSuffix(target, ".")

				switch ep.RecordType {
				case endpoint.RecordTypeNAPTR:
					illegalTarget = !hasDot
				default:
					illegalTarget = hasDot
				}

				if illegalTarget {
					fixed := target + "."
					if ep.RecordType != endpoint.RecordTypeNAPTR {
						fixed = strings.TrimSuffix(target, ".")
					}
					log.Warnf("Endpoint %s/%s with DNSName %s has an illegal target %q for %s record — use %q not %q.",
						dnsEndpoint.Namespace, dnsEndpoint.Name, ep.DNSName, target, ep.RecordType, fixed, target)
					break
				}
			}
			if illegalTarget {
				continue
			}

			ep.WithLabel(endpoint.ResourceLabelKey, fmt.Sprintf("crd/%s/%s", dnsEndpoint.Namespace, dnsEndpoint.Name))
			crdEndpoints = append(crdEndpoints, ep)
		}

		endpoint.AttachRefObject(crdEndpoints, events.NewObjectReference(dnsEndpoint, types.CRD))
		endpoints = append(endpoints, crdEndpoints...)

		if dnsEndpoint.Status.ObservedGeneration == dnsEndpoint.Generation {
			continue
		}

		dnsEndpoint.Status.ObservedGeneration = dnsEndpoint.Generation
		if err := cs.crWriter.Status().Update(ctx, dnsEndpoint); err != nil {
			log.Warnf("Could not update ObservedGeneration of [%s/%s/%s]: %v",
				"dnsendpoint", dnsEndpoint.Namespace, dnsEndpoint.Name, err)
		}
	}

	return MergeEndpoints(endpoints), nil
}

// newCrdSource wires a cache and writer into a running crdSource.
func newCrdSource(
	ctx context.Context,
	c crcache.Cache,
	crWriter client.Client,
	namespace string,
	labelSelector labels.Selector) (*crdSource, error) {
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
