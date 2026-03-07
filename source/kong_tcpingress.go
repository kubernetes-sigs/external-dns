/*
Copyright 2021 The Kubernetes Authors.

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

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
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

	"sigs.k8s.io/external-dns/source/types"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
)

var kongGroupdVersionResource = schema.GroupVersionResource{
	Group:    "configuration.konghq.com",
	Version:  "v1beta1",
	Resource: "tcpingresses",
}

// kongTCPIngressSource is an implementation of Source for Kong TCPIngress objects.
//
// +externaldns:source:name=kong-tcpingress
// +externaldns:source:category=Ingress Controllers
// +externaldns:source:description=Creates DNS entries from Kong TCPIngress resources
// +externaldns:source:resources=TCPIngress.configuration.konghq.com
// +externaldns:source:filters=annotation
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=false
// +externaldns:source:provider-specific=true
type kongTCPIngressSource struct {
	annotationFilter         string
	ignoreHostnameAnnotation bool
	dynamicKubeClient        dynamic.Interface
	kongTCPIngressInformer   kubeinformers.GenericInformer
	kubeClient               kubernetes.Interface
	namespace                string
	unstructuredConverter    *unstructuredConverter
}

// NewKongTCPIngressSource creates a new kongTCPIngressSource with the given config.
func NewKongTCPIngressSource(
	ctx context.Context,
	dynamicKubeClient dynamic.Interface, kubeClient kubernetes.Interface,
	namespace, annotationFilter string, ignoreHostnameAnnotation bool,
) (Source, error) {
	// Use shared informer to listen for add/update/delete of Host in the specified namespace.
	// Set resync period to 0, to prevent processing when nothing has changed.
	informerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicKubeClient, 0, namespace, nil)
	kongTCPIngressInformer := informerFactory.ForResource(kongGroupdVersionResource)

	// Add default resource event handlers to properly initialize informer.
	_, _ = kongTCPIngressInformer.Informer().AddEventHandler(informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForDynamicCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	uc, err := newKongUnstructuredConverter()
	if err != nil {
		return nil, fmt.Errorf("failed to setup Unstructured Converter: %w", err)
	}

	return &kongTCPIngressSource{
		annotationFilter:         annotationFilter,
		ignoreHostnameAnnotation: ignoreHostnameAnnotation,
		dynamicKubeClient:        dynamicKubeClient,
		kongTCPIngressInformer:   kongTCPIngressInformer,
		kubeClient:               kubeClient,
		namespace:                namespace,
		unstructuredConverter:    uc,
	}, nil
}

// Endpoints returns endpoint objects for each host-target combination that should be processed.
// Retrieves all TCPIngresses in the source's namespace(s).
func (sc *kongTCPIngressSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	tis, err := sc.kongTCPIngressInformer.Lister().ByNamespace(sc.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	var tcpIngresses []*TCPIngress
	for _, tcpIngressObj := range tis {
		unstructuredHost, ok := tcpIngressObj.(*unstructured.Unstructured)
		if !ok {
			return nil, errors.New("could not convert")
		}

		tcpIngress := &TCPIngress{}
		err := sc.unstructuredConverter.scheme.Convert(unstructuredHost, tcpIngress, nil)
		if err != nil {
			return nil, err
		}
		tcpIngresses = append(tcpIngresses, tcpIngress)
	}

	tcpIngresses, err = annotations.Filter(tcpIngresses, sc.annotationFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to filter TCPIngresses: %w", err)
	}

	var endpoints []*endpoint.Endpoint
	for _, tcpIngress := range tcpIngresses {
		targets := annotations.TargetsFromTargetAnnotation(tcpIngress.Annotations)
		if len(targets) == 0 {
			for _, lb := range tcpIngress.Status.LoadBalancer.Ingress {
				if lb.IP != "" {
					targets = append(targets, lb.IP)
				}
				if lb.Hostname != "" {
					targets = append(targets, lb.Hostname)
				}
			}
		}

		fullname := fmt.Sprintf("%s/%s", tcpIngress.Namespace, tcpIngress.Name)

		ingressEndpoints := sc.endpointsFromTCPIngress(tcpIngress, targets)
		if endpoint.HasNoEmptyEndpoints(ingressEndpoints, types.KongTCPIngress, tcpIngress) {
			continue
		}

		log.Debugf("Endpoints generated from TCPIngress: %s: %v", fullname, ingressEndpoints)
		endpoints = append(endpoints, ingressEndpoints...)
	}

	return MergeEndpoints(endpoints), nil
}

// endpointsFromTCPIngress extracts the endpoints from a TCPIngress object
func (sc *kongTCPIngressSource) endpointsFromTCPIngress(tcpIngress *TCPIngress, targets endpoint.Targets) []*endpoint.Endpoint {
	var endpoints []*endpoint.Endpoint

	resource := fmt.Sprintf("tcpingress/%s/%s", tcpIngress.Namespace, tcpIngress.Name)

	ttl := annotations.TTLFromAnnotations(tcpIngress.Annotations, resource)

	providerSpecific, setIdentifier := annotations.ProviderSpecificAnnotations(tcpIngress.Annotations)

	if !sc.ignoreHostnameAnnotation {
		hostnameList := annotations.HostnamesFromAnnotations(tcpIngress.Annotations)
		for _, hostname := range hostnameList {
			endpoints = append(endpoints, EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)...)
		}
	}

	if tcpIngress.Spec.Rules != nil {
		for _, rule := range tcpIngress.Spec.Rules {
			if rule.Host != "" {
				endpoints = append(endpoints, EndpointsForHostname(rule.Host, targets, ttl, providerSpecific, setIdentifier, resource)...)
			}
		}
	}

	return endpoints
}

func (sc *kongTCPIngressSource) AddEventHandler(_ context.Context, handler func()) {
	log.Debug("Adding event handler for TCPIngress")

	// Right now there is no way to remove event handler from informer, see:
	// https://github.com/kubernetes/kubernetes/issues/79610
	_, _ = sc.kongTCPIngressInformer.Informer().AddEventHandler(eventHandlerFunc(handler))
}

// newUnstructuredConverter returns a new unstructuredConverter initialized
func newKongUnstructuredConverter() (*unstructuredConverter, error) {
	uc := &unstructuredConverter{
		scheme: runtime.NewScheme(),
	}

	// Add the core types we need
	uc.scheme.AddKnownTypes(kongGroupdVersionResource.GroupVersion(), &TCPIngress{}, &TCPIngressList{})
	if err := scheme.AddToScheme(uc.scheme); err != nil {
		return nil, err
	}

	return uc, nil
}

// Kong types based on https://github.com/Kong/kubernetes-ingress-controller/blob/v1.2.0/pkg/apis/configuration/v1beta1/types.go to facilitate testing
// When trying to import them from the Kong repo as a dependency it required upgrading the k8s.io/client-go and k8s.io/apimachinery which seemed
// cause several changes in how the mock clients were working that resulted in a bunch of failures in other tests
// If that is dealt with at some point the below can be removed and replaced with an actual import
type TCPIngress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   tcpIngressSpec   `json:"spec"`
	Status tcpIngressStatus `json:"status"`
}

type TCPIngressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []TCPIngress `json:"items"`
}

type tcpIngressSpec struct {
	Rules []tcpIngressRule `json:"rules,omitempty"`
	TLS   []tcpIngressTLS  `json:"tls,omitempty"`
}

type tcpIngressTLS struct {
	Hosts      []string `json:"hosts,omitempty"`
	SecretName string   `json:"secretName,omitempty"`
}

type tcpIngressStatus struct {
	LoadBalancer corev1.LoadBalancerStatus `json:"loadBalancer"`
}

type tcpIngressRule struct {
	Host    string            `json:"host,omitempty"`
	Port    int               `json:"port,omitempty"`
	Backend tcpIngressBackend `json:"backend"`
}

type tcpIngressBackend struct {
	ServiceName string `json:"serviceName"`
	ServicePort int    `json:"servicePort"`
}

func (in *tcpIngressBackend) DeepCopyInto(out *tcpIngressBackend) {
	*out = *in
}

func (in *tcpIngressBackend) DeepCopy() *tcpIngressBackend {
	if in == nil {
		return nil
	}
	out := new(tcpIngressBackend)
	in.DeepCopyInto(out)
	return out
}

func (in *tcpIngressRule) DeepCopyInto(out *tcpIngressRule) {
	*out = *in
	out.Backend = in.Backend
}

func (in *tcpIngressRule) DeepCopy() *tcpIngressRule {
	if in == nil {
		return nil
	}
	out := new(tcpIngressRule)
	in.DeepCopyInto(out)
	return out
}

func (in *tcpIngressSpec) DeepCopyInto(out *tcpIngressSpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]tcpIngressRule, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]tcpIngressTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *tcpIngressSpec) DeepCopy() *tcpIngressSpec {
	if in == nil {
		return nil
	}
	out := new(tcpIngressSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *tcpIngressStatus) DeepCopyInto(out *tcpIngressStatus) {
	*out = *in
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
}

func (in *tcpIngressStatus) DeepCopy() *tcpIngressStatus {
	if in == nil {
		return nil
	}
	out := new(tcpIngressStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *tcpIngressTLS) DeepCopyInto(out *tcpIngressTLS) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

func (in *tcpIngressTLS) DeepCopy() *tcpIngressTLS {
	if in == nil {
		return nil
	}
	out := new(tcpIngressTLS)
	in.DeepCopyInto(out)
	return out
}

func (in *TCPIngress) DeepCopyInto(out *TCPIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

func (in *TCPIngress) DeepCopy() *TCPIngress {
	if in == nil {
		return nil
	}
	out := new(TCPIngress)
	in.DeepCopyInto(out)
	return out
}

func (in *TCPIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *TCPIngressList) DeepCopyInto(out *TCPIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TCPIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *TCPIngressList) DeepCopy() *TCPIngressList {
	if in == nil {
		return nil
	}
	out := new(TCPIngressList)
	in.DeepCopyInto(out)
	return out
}

func (in *TCPIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
