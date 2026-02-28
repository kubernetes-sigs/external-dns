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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	tools "k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/informers"
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
// +externaldns:source:events=false
type crdSource struct {
	crClient client.Client
	informer tools.SharedIndexInformer
	indexer  tools.Indexer
}

// NewCRDSource creates a new crdSource with the given config.
func NewCRDSource(
	ctx context.Context,
	restConfig *rest.Config,
	cfg *Config,
) (Source, error) {
	scheme := runtime.NewScheme()
	if err := apiv1alpha1.AddToScheme(scheme); err != nil {
		return nil, err
	}

	// Build a REST client scoped to the externaldns.k8s.io/v1alpha1 API group.
	// We use a raw client-go SharedIndexInformer rather than the controller-runtime
	// cache to avoid the multiNamespaceInformer wrapper that controller-runtime
	// inserts when namespace-scoping is applied. That wrapper does not implement
	// SharedIndexInformer, breaking GetIndexer(). A client-go informer with a
	// namespace-scoped ListWatch is RBAC-compatible with a Role (not ClusterRole)
	// when --namespace is set, and falls back to a cluster-wide watch otherwise.
	gv := apiv1alpha1.GroupVersion
	rCfg := *restConfig
	rCfg.APIPath = "/apis"
	rCfg.GroupVersion = &gv
	rCfg.NegotiatedSerializer = serializer.NewCodecFactory(scheme).WithoutConversion()

	restClient, err := rest.RESTClientFor(&rCfg)
	if err != nil {
		return nil, err
	}

	// NewFilteredListWatchFromClient pushes the label selector to the API server
	// so only matching DNSEndpoint objects enter the informer store.
	lw := tools.NewFilteredListWatchFromClient(
		restClient,
		"dnsendpoints",
		cfg.Namespace, // empty string = all namespaces
		func(opts *metav1.ListOptions) {
			if cfg.LabelFilter != nil && !cfg.LabelFilter.Empty() {
				opts.LabelSelector = cfg.LabelFilter.String()
			}
		},
	)

	inf := tools.NewSharedIndexInformer(lw, &apiv1alpha1.DNSEndpoint{}, 0, tools.Indexers{})

	// crClient is used only for status writes; reads come from the indexer.
	crClient, err := client.New(restConfig, client.Options{Scheme: scheme})
	if err != nil {
		return nil, err
	}

	cs, err := newCRDSource(inf, crClient, cfg.AnnotationFilter, cfg.LabelFilter, cfg.Namespace)
	if err != nil {
		return nil, err
	}

	go inf.Run(ctx.Done())
	if !tools.WaitForCacheSync(ctx.Done(), inf.HasSynced) {
		return nil, fmt.Errorf("cache failed to sync")
	}

	return cs, nil
}

// newCRDSource wires an informer and client into a crdSource with indexer setup.
func newCRDSource(
	inf tools.SharedIndexInformer,
	crClient client.Client,
	annotationFilter string,
	labelFilter labels.Selector,
	namespace string) (*crdSource, error) {
	if err := inf.AddIndexers(informers.IndexerWithOptions[*apiv1alpha1.DNSEndpoint](
		informers.IndexSelectorWithAnnotationFilter(annotationFilter),
		informers.IndexSelectorWithLabelSelector(labelFilter),
		informers.IndexSelectorWithNamespace(namespace))); err != nil {
		return nil, err
	}

	// TODO: add event handlers to update the CRD for debugging purposes

	return &crdSource{
		crClient: crClient,
		informer: inf,
		indexer:  inf.GetIndexer(),
	}, nil
}

func (cs *crdSource) AddEventHandler(_ context.Context, handler func()) {
	if cs.informer != nil {
		log.Debug("crdSource: adding event handler")
		// Right now there is no way to remove event handler from informer, see:
		// https://github.com/kubernetes/kubernetes/issues/79610
		_, _ = cs.informer.AddEventHandler(eventHandlerFunc(handler))
	}
}

// Endpoints returns endpoint objects.
func (cs *crdSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	indexKeys := cs.indexer.ListIndexFuncValues(informers.IndexWithSelectors)
	if len(indexKeys) == 0 {
		return nil, nil
	}

	endpoints := make([]*endpoint.Endpoint, 0, len(indexKeys))

	for _, key := range indexKeys {
		el, err := informers.GetByKey[*apiv1alpha1.DNSEndpoint](cs.indexer, key)
		// el may be nil (with no error) if the object was deleted between
		// ListIndexFuncValues and GetByKey (TOCTOU race on the store).
		if err != nil || el == nil {
			log.Debugf("Failed to get DNSEndpoint for key %s: %v", key, err)
			continue
		}

		// Compute resource label once per DNSEndpoint
		resourceLabel := fmt.Sprintf("crd/%s/%s", el.Namespace, el.Name)
		fmt.Println("resourcelabel", resourceLabel)

		for _, ep := range el.Spec.Endpoints {
			if (ep.RecordType == endpoint.RecordTypeCNAME || ep.RecordType == endpoint.RecordTypeA || ep.RecordType == endpoint.RecordTypeAAAA) && len(ep.Targets) < 1 {
				log.Debugf("Endpoint %s with DNSName %s has an empty list of targets, allowing it to pass through for default-targets processing", el.Name, ep.DNSName)
			}
			illegalTarget := false
			for _, target := range ep.Targets {
				switch ep.RecordType {
				case endpoint.RecordTypeTXT, endpoint.RecordTypeMX:
					continue // TXT records allow arbitrary text, skip validation; MX records can have trailing dot but it's not required, skip validation
				case endpoint.RecordTypeCNAME:
					continue // RFC 1035 §5.1: trailing dot denotes an absolute FQDN in zone file notation; both forms are valid
				}

				hasDot := strings.HasSuffix(target, ".")

				switch ep.RecordType {
				case endpoint.RecordTypeNAPTR:
					illegalTarget = !hasDot // Must have trailing dot
				default:
					illegalTarget = hasDot // Must NOT have trailing dot
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
				log.Warnf("Endpoint %s/%s with DNSName %s has an illegal target format.", el.Namespace, el.Name, ep.DNSName)
				continue
			}

			ep.WithLabel(endpoint.ResourceLabelKey, resourceLabel)
			endpoints = append(endpoints, ep)
		}

		if el.Status.ObservedGeneration == el.Generation {
			continue
		}

		el.Status.ObservedGeneration = el.Generation
		if err := cs.crClient.Status().Update(ctx, el); err != nil {
			log.Warnf("Could not update ObservedGeneration of the CRD: %v", err)
		}
	}

	return MergeEndpoints(endpoints), nil
}
