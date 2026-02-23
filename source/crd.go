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
	"os"
	"strings"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/source/annotations"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
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
	crdClient        rest.Interface
	namespace        string
	crdResource      string
	codec            runtime.ParameterCodec
	annotationFilter string
	labelSelector    labels.Selector
	informer         cache.SharedInformer
}

// NewCRDClientForAPIVersionKind return rest client for the given apiVersion and kind of the CRD
func NewCRDClientForAPIVersionKind(client kubernetes.Interface, kubeConfig, apiServerURL, apiVersion, kind string) (*rest.RESTClient, *runtime.Scheme, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	config, err := clientcmd.BuildConfigFromFlags(apiServerURL, kubeConfig)
	if err != nil {
		return nil, nil, err
	}

	groupVersion, err := schema.ParseGroupVersion(apiVersion)
	if err != nil {
		return nil, nil, err
	}
	apiResourceList, err := client.Discovery().ServerResourcesForGroupVersion(groupVersion.String())
	if err != nil {
		return nil, nil, fmt.Errorf("error listing resources in GroupVersion %q: %w", groupVersion.String(), err)
	}

	var crdAPIResource *metav1.APIResource
	for _, apiResource := range apiResourceList.APIResources {
		if apiResource.Kind == kind {
			crdAPIResource = &apiResource
			break
		}
	}
	if crdAPIResource == nil {
		return nil, nil, fmt.Errorf("unable to find Resource Kind %q in GroupVersion %q", kind, apiVersion)
	}

	scheme := runtime.NewScheme()
	_ = apiv1alpha1.AddToScheme(scheme)

	config.GroupVersion = &groupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}

	crdClient, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		return nil, nil, err
	}
	return crdClient, scheme, nil
}

// NewCRDSource creates a new crdSource with the given config.
func NewCRDSource(
	crdClient rest.Interface,
	namespace, kind, annotationFilter string,
	labelSelector labels.Selector,
	scheme *runtime.Scheme,
	startInformer bool) (Source, error) {
	sourceCrd := crdSource{
		crdResource:      strings.ToLower(kind) + "s",
		namespace:        namespace,
		annotationFilter: annotationFilter,
		labelSelector:    labelSelector,
		crdClient:        crdClient,
		codec:            runtime.NewParameterCodec(scheme),
	}
	if startInformer {
		// external-dns already runs its sync-handler periodically (controlled by `--interval` flag) to ensure any
		// missed or dropped events are handled. specify resync period 0 to avoid unnecessary sync handler invocations.
		sourceCrd.informer = cache.NewSharedInformer(
			&cache.ListWatch{
				ListWithContextFunc: func(ctx context.Context, lo metav1.ListOptions) (runtime.Object, error) {
					return sourceCrd.List(ctx, &lo)
				},
				WatchFuncWithContext: func(ctx context.Context, lo metav1.ListOptions) (watch.Interface, error) {
					return sourceCrd.watch(ctx, &lo)
				},
			},
			&apiv1alpha1.DNSEndpoint{},
			0)
		go sourceCrd.informer.Run(wait.NeverStop)
	}
	return &sourceCrd, nil
}

func (cs *crdSource) AddEventHandler(_ context.Context, handler func()) {
	if cs.informer != nil {
		log.Debug("Adding event handler for CRD")
		// Right now there is no way to remove event handler from informer, see:
		// https://github.com/kubernetes/kubernetes/issues/79610
		_, _ = cs.informer.AddEventHandler(
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(_ any) {
					handler()
				},
				UpdateFunc: func(_ any, _ any) {
					handler()
				},
				DeleteFunc: func(_ any) {
					handler()
				},
			},
		)
	}
}

// Endpoints returns endpoint objects.
func (cs *crdSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	var (
		result *apiv1alpha1.DNSEndpointList
		err    error
	)

	result, err = cs.List(ctx, &metav1.ListOptions{LabelSelector: cs.labelSelector.String()})
	if err != nil {
		return nil, err
	}

	itemPtrs := make([]*apiv1alpha1.DNSEndpoint, len(result.Items))
	for i := range result.Items {
		itemPtrs[i] = &result.Items[i]
	}

	filtered, err := annotations.Filter(itemPtrs, cs.annotationFilter)
	if err != nil {
		return nil, err
	}

	for _, dnsEndpoint := range filtered {
		var crdEndpoints []*endpoint.Endpoint
		for _, ep := range dnsEndpoint.Spec.Endpoints {
			if (ep.RecordType == endpoint.RecordTypeCNAME || ep.RecordType == endpoint.RecordTypeA || ep.RecordType == endpoint.RecordTypeAAAA) && len(ep.Targets) < 1 {
				log.Debugf("Endpoint %s with DNSName %s has an empty list of targets, allowing it to pass through for default-targets processing", dnsEndpoint.Name, ep.DNSName)
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
				continue
			}

			ep.WithLabel(endpoint.ResourceLabelKey, fmt.Sprintf("crd/%s/%s", dnsEndpoint.Namespace, dnsEndpoint.Name))

			crdEndpoints = append(crdEndpoints, ep)
		}

		endpoints = append(endpoints, crdEndpoints...)

		if dnsEndpoint.Status.ObservedGeneration == dnsEndpoint.Generation {
			continue
		}

		dnsEndpoint.Status.ObservedGeneration = dnsEndpoint.Generation
		// Update the ObservedGeneration
		_, err = cs.UpdateStatus(ctx, dnsEndpoint)
		if err != nil {
			log.Warnf("Could not update ObservedGeneration of the CRD: %v", err)
		}
	}

	return MergeEndpoints(endpoints), nil
}

func (cs *crdSource) watch(ctx context.Context, opts *metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return cs.crdClient.Get().
		Namespace(cs.namespace).
		Resource(cs.crdResource).
		VersionedParams(opts, cs.codec).
		Watch(ctx)
}

func (cs *crdSource) List(ctx context.Context, opts *metav1.ListOptions) (*apiv1alpha1.DNSEndpointList, error) {
	result := &apiv1alpha1.DNSEndpointList{}
	return result, cs.crdClient.Get().
		Namespace(cs.namespace).
		Resource(cs.crdResource).
		VersionedParams(opts, cs.codec).
		Do(ctx).
		Into(result)
}

func (cs *crdSource) UpdateStatus(ctx context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint) (*apiv1alpha1.DNSEndpoint, error) {
	result := &apiv1alpha1.DNSEndpoint{}
	return result, cs.crdClient.Put().
		Namespace(dnsEndpoint.Namespace).
		Resource(cs.crdResource).
		Name(dnsEndpoint.Name).
		SubResource("status").
		Body(dnsEndpoint).
		Do(ctx).
		Into(result)
}
