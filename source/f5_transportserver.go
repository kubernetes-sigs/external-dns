/*
Copyright 2022 The Kubernetes Authors.

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

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"

	f5 "github.com/F5Networks/k8s-bigip-ctlr/v2/config/apis/cis/v1"

	"sigs.k8s.io/external-dns/source/informers"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

var f5TransportServerGVR = schema.GroupVersionResource{
	Group:    "cis.f5.com",
	Version:  "v1",
	Resource: "transportservers",
}

// transportServerSource is an implementation of Source for F5 TransportServer objects.
type f5TransportServerSource struct {
	dynamicKubeClient       dynamic.Interface
	transportServerInformer kubeinformers.GenericInformer
	kubeClient              kubernetes.Interface
	annotationFilter        string
	namespace               string
	unstructuredConverter   *unstructuredConverter
}

func NewF5TransportServerSource(
	ctx context.Context,
	dynamicKubeClient dynamic.Interface,
	kubeClient kubernetes.Interface,
	namespace string,
	annotationFilter string,
) (Source, error) {
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicKubeClient, 0, namespace, nil)
	transportServerInformer := informerFactory.ForResource(f5TransportServerGVR)

	transportServerInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForDynamicCacheSync(context.Background(), informerFactory); err != nil {
		return nil, err
	}

	uc, err := newTSUnstructuredConverter()
	if err != nil {
		return nil, fmt.Errorf("failed to setup unstructured converter: %w", err)
	}

	return &f5TransportServerSource{
		dynamicKubeClient:       dynamicKubeClient,
		transportServerInformer: transportServerInformer,
		kubeClient:              kubeClient,
		namespace:               namespace,
		annotationFilter:        annotationFilter,
		unstructuredConverter:   uc,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all TransportServers in the source's namespace(s).
func (ts *f5TransportServerSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	transportServerObjects, err := ts.transportServerInformer.Lister().ByNamespace(ts.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var transportServers []*f5.TransportServer
	for _, tsObj := range transportServerObjects {
		unstructuredHost, ok := tsObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert")
		}

		transportServer := &f5.TransportServer{}
		err := ts.unstructuredConverter.scheme.Convert(unstructuredHost, transportServer, nil)
		if err != nil {
			return nil, err
		}
		transportServers = append(transportServers, transportServer)
	}

	transportServers, err = ts.filterByAnnotations(transportServers)
	if err != nil {
		return nil, fmt.Errorf("failed to filter TransportServers: %w", err)
	}

	endpoints, err := ts.endpointsFromTransportServers(transportServers)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

func (ts *f5TransportServerSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("Adding event handler for TransportServer")

	ts.transportServerInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}

// endpointsFromTransportServers extracts the endpoints from a slice of TransportServers
func (ts *f5TransportServerSource) endpointsFromTransportServers(transportServers []*f5.TransportServer) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	for _, transportServer := range transportServers {
		if !hasValidTransportServerIP(transportServer) {
			log.Warnf("F5 TransportServer %s/%s is missing a valid IP address, skipping endpoint creation.",
				transportServer.Namespace, transportServer.Name)
			continue
		}

		resource := fmt.Sprintf("f5-transportserver/%s/%s", transportServer.Namespace, transportServer.Name)

		ttl := annotations.TTLFromAnnotations(transportServer.Annotations, resource)

		targets := annotations.TargetsFromTargetAnnotation(transportServer.Annotations)
		if len(targets) == 0 && transportServer.Spec.VirtualServerAddress != "" {
			targets = append(targets, transportServer.Spec.VirtualServerAddress)
		}
		if len(targets) == 0 && transportServer.Status.VSAddress != "" {
			targets = append(targets, transportServer.Status.VSAddress)
		}

		endpoints = append(endpoints, endpointsForHostname(transportServer.Spec.Host, targets, ttl, nil, "", resource)...)
	}

	return endpoints, nil
}

// newUnstructuredConverter returns a new unstructuredConverter initialized
func newTSUnstructuredConverter() (*unstructuredConverter, error) {
	uc := &unstructuredConverter{
		scheme: runtime.NewScheme(),
	}

	// Add the core types we need
	uc.scheme.AddKnownTypes(f5TransportServerGVR.GroupVersion(), &f5.TransportServer{}, &f5.TransportServerList{})
	if err := scheme.AddToScheme(uc.scheme); err != nil {
		return nil, err
	}

	return uc, nil
}

// filterByAnnotations filters a list of TransportServers by a given annotation selector.
func (ts *f5TransportServerSource) filterByAnnotations(transportServers []*f5.TransportServer) ([]*f5.TransportServer, error) {
	labelSelector, err := metav1.ParseToLabelSelector(ts.annotationFilter)
	if err != nil {
		return nil, err
	}

	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}

	// empty filter returns original list
	if selector.Empty() {
		return transportServers, nil
	}

	filteredList := []*f5.TransportServer{}

	for _, ts := range transportServers {
		// include TransportServer if its annotations match the selector
		if selector.Matches(labels.Set(ts.Annotations)) {
			filteredList = append(filteredList, ts)
		}
	}

	return filteredList, nil
}

func hasValidTransportServerIP(vs *f5.TransportServer) bool {
	normalizedAddress := strings.ToLower(vs.Status.VSAddress)
	return normalizedAddress != "none" && normalizedAddress != ""
}
