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
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tools "k8s.io/client-go/tools/cache"

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
	informer cache.Informer
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

	// Build cache options with label selector and optional namespace filter
	byObject := cache.ByObject{Label: cfg.LabelFilter}
	if cfg.Namespace != "" {
		byObject.Namespaces = map[string]cache.Config{cfg.Namespace: {}}
	}

	crCache, err := cache.New(restConfig, cache.Options{
		Scheme:           scheme,
		DefaultTransform: cache.TransformStripManagedFields(),
		ByObject:         map[client.Object]cache.ByObject{&apiv1alpha1.DNSEndpoint{}: byObject},
	})
	if err != nil {
		return nil, err
	}

	crClient, err := client.New(restConfig, client.Options{
		Scheme: scheme,
		Cache:  &client.CacheOptions{Reader: crCache},
	})
	if err != nil {
		return nil, err
	}

	if err := informers.StartAndWaitForCacheSync(ctx, crCache); err != nil {
		return nil, err
	}

	inf, err := crCache.GetInformer(ctx, &apiv1alpha1.DNSEndpoint{})
	if err != nil {
		return nil, err
	}
	// controller-runtime's cache.GetInformer always returns a SharedIndexInformer.
	return newCRDSource(inf.(tools.SharedIndexInformer), crClient, cfg.AnnotationFilter, cfg.LabelFilter, cfg.Namespace)
}

// newCRDSource wires a SharedIndexInformer and client into a crdSource.
// It is called by NewCRDSource (production) and the test helper so both share
// the same indexer setup and struct construction.
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
		if err != nil {
			log.Debugf("Failed to get DNSEndpoint for key %s: %v", key, err)
			continue
		}

		// Compute resource label once per DNSEndpoint
		resourceLabel := fmt.Sprintf("crd/%s/%s", el.Namespace, el.Name)

		for _, ep := range el.Spec.Endpoints {
			if (ep.RecordType == endpoint.RecordTypeCNAME || ep.RecordType == endpoint.RecordTypeA || ep.RecordType == endpoint.RecordTypeAAAA) && len(ep.Targets) < 1 {
				log.Debugf("Endpoint %s with DNSName %s has an empty list of targets, allowing it to pass through for default-targets processing", el.Name, ep.DNSName)
			}
			illegalTarget := false
			for _, target := range ep.Targets {
				hasDot := strings.HasSuffix(target, ".")

				switch ep.RecordType {
				case endpoint.RecordTypeTXT, endpoint.RecordTypeMX:
					continue // TXT records allow arbitrary text, skip validation; MX records can have trailing dot but it's not required, skip validation
				case endpoint.RecordTypeNAPTR:
					illegalTarget = !hasDot // Must have trailing dot
				default:
					illegalTarget = hasDot // Must NOT have trailing dot
				}

				if illegalTarget {
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
		// Update the ObservedGeneration
		if err := cs.crClient.Status().Update(ctx, el); err != nil {
			log.Warnf("Could not update ObservedGeneration of the CRD: %v", err)
		}
	}

	return MergeEndpoints(endpoints), nil
}
