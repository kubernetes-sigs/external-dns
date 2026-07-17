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

	f5 "github.com/F5Networks/k8s-bigip-ctlr/v2/config/apis/cis/v1"
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

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
	"sigs.k8s.io/external-dns/source/types"
)

var f5VirtualServerGVR = schema.GroupVersionResource{
	Group:    "cis.f5.com",
	Version:  "v1",
	Resource: "virtualservers",
}

// virtualServerSource is an implementation of Source for F5 VirtualServer objects.
//
// +externaldns:source:name=f5-virtualserver
// +externaldns:source:category=Load Balancers
// +externaldns:source:description=Creates DNS entries from F5 VirtualServer resources
// +externaldns:source:resources=VirtualServer.cis.f5.com
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=false
type f5VirtualServerSource struct {
	dynamicKubeClient     dynamic.Interface
	virtualServerInformer kubeinformers.GenericInformer
	kubeClient            kubernetes.Interface
	annotationFilter      labels.Selector
	labelSelector         labels.Selector
	namespace             string
	templateEngine        template.Engine
	unstructuredConverter *unstructuredConverter
}

func NewF5VirtualServerSource(
	ctx context.Context,
	dynamicKubeClient dynamic.Interface,
	kubeClient kubernetes.Interface,
	cfg *Config,
) (Source, error) {
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicKubeClient, 0, cfg.Namespace, nil)
	virtualServerInformer := informerFactory.ForResource(f5VirtualServerGVR)

	informers.MustSetTransform(virtualServerInformer.Informer(), informers.TransformerWithOptions[*unstructured.Unstructured](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
	))

	informers.MustAddEventHandler(virtualServerInformer.Informer(), informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForDynamicCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	uc, err := newVSUnstructuredConverter()
	if err != nil {
		return nil, fmt.Errorf("failed to setup unstructured converter: %w", err)
	}

	return &f5VirtualServerSource{
		dynamicKubeClient:     dynamicKubeClient,
		virtualServerInformer: virtualServerInformer,
		kubeClient:            kubeClient,
		namespace:             cfg.Namespace,
		annotationFilter:      cfg.AnnotationFilter,
		labelSelector:         cfg.LabelFilter,
		templateEngine:        cfg.TemplateEngine,
		unstructuredConverter: uc,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all VirtualServers in the source's namespace(s).
func (vs *f5VirtualServerSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	virtualServerObjects, err := vs.virtualServerInformer.Lister().ByNamespace(vs.namespace).List(vs.labelSelector)
	if err != nil {
		return nil, err
	}

	var virtualServers []*f5.VirtualServer
	for _, vsObj := range virtualServerObjects {
		unstructuredHost, ok := vsObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert")
		}

		virtualServer := &f5.VirtualServer{}
		err := vs.unstructuredConverter.scheme.Convert(unstructuredHost, virtualServer, nil)
		if err != nil {
			return nil, err
		}
		virtualServer.TypeMeta = metav1.TypeMeta{
			Kind:       "VirtualServer",
			APIVersion: f5VirtualServerGVR.GroupVersion().String(),
		}
		virtualServers = append(virtualServers, virtualServer)
	}

	virtualServers = annotations.Filter(virtualServers, vs.annotationFilter)

	endpoints, err := vs.endpointsFromVirtualServers(virtualServers)
	if err != nil {
		return nil, err
	}

	return endpoint.MergeEndpoints(endpoints), nil
}

func (vs *f5VirtualServerSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handler for VirtualServer")

	informers.MustAddEventHandler(vs.virtualServerInformer.Informer(), eventHandlerFunc(handler))
}

// endpointsFromVirtualServers extracts the endpoints from a slice of VirtualServers.
func (vs *f5VirtualServerSource) endpointsFromVirtualServers(virtualServers []*f5.VirtualServer) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint

	for _, virtualServer := range virtualServers {
		var vsEndpoints []*endpoint.Endpoint

		if hasValidVirtualServerIP(virtualServer) {
			resource := fmt.Sprintf("f5-virtualserver/%s/%s", virtualServer.Namespace, virtualServer.Name)

			ttl := annotations.TTLFromAnnotations(virtualServer.Annotations, resource)

			targets := annotations.TargetsFromTargetAnnotation(virtualServer.Annotations)
			if len(targets) == 0 && virtualServer.Spec.VirtualServerAddress != "" {
				targets = append(targets, virtualServer.Spec.VirtualServerAddress)
			}
			if len(targets) == 0 && virtualServer.Status.VSAddress != "" {
				targets = append(targets, virtualServer.Status.VSAddress)
			}

			vsEndpoints = append(vsEndpoints, endpoint.EndpointsForHostname(virtualServer.Spec.Host, targets, ttl, nil, "", resource)...)
			for _, alias := range virtualServer.Spec.HostAliases {
				if alias != "" {
					vsEndpoints = append(vsEndpoints, endpoint.EndpointsForHostname(alias, targets, ttl, nil, "", resource)...)
				}
			}
		}

		var err error
		vsEndpoints, err = vs.templateEngine.ApplyTemplates(vsEndpoints, virtualServer)
		if err != nil {
			return nil, err
		}

		if len(vsEndpoints) == 0 {
			log.Warnf("F5 VirtualServer %s/%s is missing a valid IP address, skipping endpoint creation.",
				virtualServer.Namespace, virtualServer.Name)
		}

		endpoint.AttachRefObject(vsEndpoints, events.NewObjectReference(virtualServer, types.F5VirtualServer))

		endpoints = append(endpoints, vsEndpoints...)
	}

	return endpoints, nil
}

// newUnstructuredConverter returns a new unstructuredConverter initialized
func newVSUnstructuredConverter() (*unstructuredConverter, error) {
	uc := &unstructuredConverter{
		scheme: runtime.NewScheme(),
	}

	// Add the core types we need
	uc.scheme.AddKnownTypes(f5VirtualServerGVR.GroupVersion(), &f5.VirtualServer{}, &f5.VirtualServerList{})
	if err := scheme.AddToScheme(uc.scheme); err != nil {
		return nil, err
	}

	return uc, nil
}

func hasValidVirtualServerIP(vs *f5.VirtualServer) bool {
	normalizedAddress := strings.ToLower(vs.Status.VSAddress)
	return normalizedAddress != "none" && normalizedAddress != ""
}
