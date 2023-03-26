/*
Copyright 2023 The Kubernetes Authors.

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
	"bytes"
	"context"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/jsonpath"

	"sigs.k8s.io/external-dns/endpoint"
)

// unstructuredSource is an implementation of Source that provides endpoints by
// listing a specified Kubernetes API and fetching endpoints based on the
// source's configured property paths.
type unstructuredSource struct {
	dynamicKubeClient  dynamic.Interface
	kubeClient         kubernetes.Interface
	targetJsonPath     *jsonpath.JSONPath
	hostnameJsonPath   *jsonpath.JSONPath
	namespace          string
	namespacedResource bool
	informer           informers.GenericInformer
	labelSelector      labels.Selector
	annotationSelector labels.Selector
}

// NewUnstructuredSource creates a new unstructuredSource with the given config.
func NewUnstructuredSource(
	ctx context.Context,
	dynamicKubeClient dynamic.Interface,
	kubeClient kubernetes.Interface,
	namespace, apiVersion, kind,
	targetJsonPath, hostnameJsonPath string,
	labelSelector labels.Selector, annotationFilter string) (Source, error) {
	var (
		targetJsonPathObj   *jsonpath.JSONPath
		hostnameJsonPathObj *jsonpath.JSONPath
	)
	if targetJsonPath != "" {
		targetJsonPathObj = jsonpath.New("p")
		if err := targetJsonPathObj.Parse(targetJsonPath); err != nil {
			return nil, fmt.Errorf("failed to parse targetJsonPath %q as JSONPath expression while initializing unstructured source: %w", targetJsonPath, err)
		}
	}
	if hostnameJsonPath != "" {
		hostnameJsonPathObj = jsonpath.New("p")
		if err := hostnameJsonPathObj.Parse(hostnameJsonPath); err != nil {
			return nil, fmt.Errorf("failed to parse hostnameJsonPath %q as JSONPath expression while initializing unstructured source: %w", hostnameJsonPath, err)
		}
	}

	// Build a selector to filter resources by annotations.
	annotationProtoSelector, err := metav1.ParseToLabelSelector(annotationFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to parse annotationFilter %q while initializing unstructured source: %w", annotationFilter, err)
	}
	annotationSelector, err := metav1.LabelSelectorAsSelector(annotationProtoSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to create annotationSelector while initializing unstructured source: %w", err)
	}

	groupVersion, err := schema.ParseGroupVersion(apiVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to parse apiVersion %q while initializing unstructured source: %w", apiVersion, err)
	}
	groupVersionResource := groupVersion.WithResource(strings.ToLower(kind) + "s")

	// Determine if the specified API group and resource exists and
	// whether or not the API is namespace-scoped or cluster-scoped.
	var apiResource *metav1.APIResource
	apiResourceList, err := kubeClient.Discovery().ServerResourcesForGroupVersion(apiVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to list resources for apiVersion %q while initializing unstructured source: %s", apiVersion, err)
	}
	for _, ar := range apiResourceList.APIResources {
		if ar.Kind == kind {
			apiResource = &ar
			break
		}
	}
	if apiResource == nil {
		return nil, fmt.Errorf("failed to find kind %q in apiVersion %q while initializing unstructured source", kind, apiVersion)
	}

	if apiResource.Namespaced {
		log.Infof("Unstructured source configured for namespace-scoped resource with kind %q in apiVersion %q in namespace %q", kind, apiVersion, namespace)
	} else {
		log.Infof("Unstructured source configured for cluster-scoped resource with kind %q in apiVersion %q", kind, apiVersion)
	}

	// Use the shared informer to listen for add/update/delete for the
	// resource in the specified namespace (if it's a namespaced resource).
	// If the resource is cluster-scoped, the informer is provided with an
	// empty namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(
		dynamicKubeClient,
		0,
		namespace,
		nil)
	informer := informerFactory.ForResource(groupVersionResource)

	// Add default resource event handlers to properly initialize informer.
	informer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)

	informerFactory.Start(ctx.Done())

	if err := waitForDynamicCacheSync(context.Background(), informerFactory); err != nil {
		return nil, fmt.Errorf(
			"failed to wait for dynamic client cache to sync while initializing unstructured source: %w", err)
	}

	return &unstructuredSource{
		kubeClient:         kubeClient,
		dynamicKubeClient:  dynamicKubeClient,
		informer:           informer,
		targetJsonPath:     targetJsonPathObj,
		hostnameJsonPath:   hostnameJsonPathObj,
		namespace:          namespace,
		namespacedResource: apiResource.Namespaced,
		labelSelector:      labelSelector,
		annotationSelector: annotationSelector,
	}, nil
}

func (us *unstructuredSource) AddEventHandler(ctx context.Context, handler func()) {
}

// Endpoints returns endpoint objects.
func (us *unstructuredSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var (
		endpoints []*endpoint.Endpoint
		listFn    func(labels.Selector) ([]runtime.Object, error)
	)

	// List the resources using the label selector.
	lister := us.informer.Lister()
	if us.namespacedResource {
		listFn = lister.ByNamespace(us.namespace).List
	} else {
		listFn = lister.List
	}
	list, err := listFn(us.labelSelector)
	if err != nil {
		if us.namespacedResource {
			return nil, fmt.Errorf(
				"failed to list resources in namespace %q in unstructured source: %w", us.namespace, err)
		}
		return nil, fmt.Errorf(
			"failed to list cluster resources in unstructured source: %w", err)
	}

	for _, obj := range list {
		item, ok := obj.(*unstructured.Unstructured)
		if !ok {
			return nil, fmt.Errorf(
				"failed to assert %[1]T %[1]v is *unstructured.Unstructured", obj)
		}

		// If there is an annotation selector then skip any resource
		// that does not match the annotation selector.
		if !us.annotationSelector.Empty() {
			if !us.annotationSelector.Matches(labels.Set(item.GetAnnotations())) {
				continue
			}
		}

		hostnames, err := us.getHostnames(item)
		if err != nil {
			log.Warnf("failed to get hostname(s) for resource %q: %s", item.GetName(), err)
			continue
		}

		targets, err := us.getTargets(item)
		if err != nil {
			log.Warnf("failed to get target(s) for resource %q that has hostnames %s: %s", item.GetName(), hostnames, err)
			continue
		}

		ttl, err := us.getTTL(item)
		if err != nil {
			log.Warnf("failed to get TTL for resource %q that has hostnames %s and targets %s: %s", item.GetName(), hostnames, targets, err)
			continue
		}

		recordType := us.getRecordType(item)

		log.Infof("resource=%q, hostnames=%s, targets=%s, ttl=%v, recordType=%q\n", item.GetName(), hostnames, targets, ttl, recordType)

		// Create an endpoint for each host name.
		for i := range hostnames {
			endpoints = append(
				endpoints, &endpoint.Endpoint{
					Labels:     us.getEndpointLabels(item),
					DNSName:    hostnames[i],
					Targets:    endpoint.NewTargets(targets...),
					RecordType: recordType,
					RecordTTL:  ttl,
				})
		}
	}

	return endpoints, nil
}

func (us *unstructuredSource) getEndpointLabels(item *unstructured.Unstructured) map[string]string {
	var val string
	if us.namespacedResource {
		val = fmt.Sprintf(
			"unstructured/%s/%s",
			item.GetNamespace(),
			item.GetName())
	} else {
		val = fmt.Sprintf(
			"unstructured/%s",
			item.GetName())
	}
	return map[string]string{
		endpoint.ResourceLabelKey: val,
	}
}

func (us *unstructuredSource) getTTL(item *unstructured.Unstructured) (endpoint.TTL, error) {
	ttl, err := getTTLFromAnnotations(item.GetAnnotations())
	if err != nil {
		return 0, err
	}
	return ttl, nil
}

func (us *unstructuredSource) getRecordType(item *unstructured.Unstructured) string {
	recordType := getRecordTypeFromAnnotations(item.GetAnnotations())
	if recordType == "" {
		recordType = endpoint.RecordTypeA
	}
	return recordType
}

func (us *unstructuredSource) getTargets(item *unstructured.Unstructured) ([]string, error) {
	var targets []string
	if us.targetJsonPath == nil {
		targets = getTargetsFromTargetAnnotation(item.GetAnnotations())
	} else {
		var err error
		targets, err = us.executeJsonPath(us.targetJsonPath, item.Object)
		if err != nil {
			return nil, fmt.Errorf("failed to execute JSONPath while getting targets: %w", err)
		}
	}
	if len(targets) == 0 {
		return nil, fmt.Errorf("no targets")
	}
	return targets, nil
}

func (us *unstructuredSource) getHostnames(item *unstructured.Unstructured) ([]string, error) {
	var hostnames []string
	if us.hostnameJsonPath == nil {
		hostnames = getHostnamesFromAnnotations(item.GetAnnotations())
	} else {
		var err error
		hostnames, err = us.executeJsonPath(us.hostnameJsonPath, item.Object)
		if err != nil {
			return nil, fmt.Errorf("failed to execute JSONPath while getting hostnames: %w", err)
		}
	}
	if len(hostnames) == 0 {
		return nil, fmt.Errorf("no hostnames")
	}
	return hostnames, nil
}

func (us *unstructuredSource) executeJsonPath(jp *jsonpath.JSONPath, data interface{}) ([]string, error) {
	var w bytes.Buffer
	if err := jp.Execute(&w, data); err != nil {
		return nil, err
	}
	s := w.String()

	// Trim any surrounding brackets that decorate JSONPath expressions
	// that evaluate to a list.
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")

	return splitByWhitespaceAndComma(s), nil
}
