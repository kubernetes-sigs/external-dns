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

// TODO:
// support
// - set-identifier for endpoints created
// - set resource aka fmt.Sprintf("pod/%s/%s", pod.Namespace, pod.Name)
package source

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
	"sigs.k8s.io/external-dns/source/types"
)

// podSource is an implementation of Source for Kubernetes Pod objects.
//
// +externaldns:source:name=pod
// +externaldns:source:category=Kubernetes Core
// +externaldns:source:description=Creates DNS entries based on Kubernetes Pod resources
// +externaldns:source:resources=Pod
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=false
// +externaldns:source:events=true
type podSource struct {
	client         kubernetes.Interface
	namespace      string
	templateEngine template.Engine

	podInformer              coreinformers.PodInformer
	nodeInformer             coreinformers.NodeInformer
	compatibility            string
	ignoreNonHostNetworkPods bool
	podSourceDomain          string
}

// NewPodSource creates a new podSource with the given config.
func NewPodSource(
	ctx context.Context,
	kubeClient kubernetes.Interface,
	cfg *Config,
) (Source, error) {
	namespace := cfg.Namespace
	annotationFilter := cfg.AnnotationFilter
	labelSelector := cfg.LabelFilter

	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(kubeClient, 0, kubeinformers.WithNamespace(namespace))
	podInformer := informerFactory.Core().V1().Pods()
	nodeInformer := informerFactory.Core().V1().Nodes()

	informers.MustAddIndexers(podInformer.Informer(), informers.IndexerWithOptions[*v1.Pod](
		informers.IndexSelectorWithAnnotationFilter(annotationFilter),
		informers.IndexSelectorWithLabelSelector(labelSelector),
	))
	informers.MustSetTransform(podInformer.Informer(), informers.TransformerWithOptions[*v1.Pod](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
		informers.TransformRemoveStatusConditions(),
	))
	informers.MustSetTransform(nodeInformer.Informer(), informers.TransformerWithOptions[*v1.Node](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
		informers.TransformRemoveStatusConditions(),
	))

	informers.MustAddEventHandler(podInformer.Informer(), informers.DefaultEventHandler())
	informers.MustAddEventHandler(nodeInformer.Informer(), informers.DefaultEventHandler())

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForCacheSync(ctx, informerFactory); err != nil {
		return nil, err
	}

	return &podSource{
		client:                   kubeClient,
		podInformer:              podInformer,
		nodeInformer:             nodeInformer,
		namespace:                namespace,
		compatibility:            cfg.Compatibility,
		ignoreNonHostNetworkPods: cfg.IgnoreNonHostNetworkPods,
		podSourceDomain:          cfg.PodSourceDomain,
		templateEngine:           cfg.TemplateEngine,
	}, nil
}

func (ps *podSource) AddEventHandler(_ context.Context, handler func()) {
	informers.MustAddEventHandler(ps.podInformer.Informer(), eventHandlerFunc(handler))
}

func (ps *podSource) Endpoints(_ context.Context) ([]*endpoint.Endpoint, error) {
	indexKeys := ps.podInformer.Informer().GetIndexer().ListIndexFuncValues(informers.IndexWithSelectors)

	endpoints := make([]*endpoint.Endpoint, 0)
	for _, key := range indexKeys {
		pod, err := informers.GetByKey[*v1.Pod](ps.podInformer.Informer().GetIndexer(), key)
		if err != nil {
			continue
		}

		podEndpoints := ps.endpointsFromPodAnnotations(pod)

		podEndpoints, err = ps.templateEngine.CombineWithEndpoints(
			podEndpoints,
			func() ([]*endpoint.Endpoint, error) { return ps.endpointsFromPodTemplate(pod) },
		)
		if err != nil {
			return nil, err
		}

		endpoint.AttachRefObject(podEndpoints, events.NewObjectReference(pod, types.Pod))

		endpoints = append(endpoints, podEndpoints...)
	}

	return MergeEndpoints(endpoints), nil
}

func (ps *podSource) endpointsFromPodAnnotations(pod *v1.Pod) []*endpoint.Endpoint {
	endpointMap := make(map[endpoint.EndpointKey][]string)
	ps.addPodEndpointsToEndpointMap(endpointMap, pod)

	var endpoints []*endpoint.Endpoint
	for key, targets := range endpointMap {
		endpoints = append(endpoints, endpoint.NewEndpointWithTTL(key.DNSName, key.RecordType, key.RecordTTL, targets...))
	}
	return endpoints
}

func (ps *podSource) endpointsFromPodTemplate(pod *v1.Pod) ([]*endpoint.Endpoint, error) {
	hostsMap, err := ps.hostsFromTemplate(pod)
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint
	for key, targets := range hostsMap {
		endpoints = append(endpoints, endpoint.NewEndpointWithTTL(key.DNSName, key.RecordType, key.RecordTTL, targets...))
	}
	return endpoints, nil
}

func (ps *podSource) addPodEndpointsToEndpointMap(endpointMap map[endpoint.EndpointKey][]string, pod *v1.Pod) {
	if ps.ignoreNonHostNetworkPods && !pod.Spec.HostNetwork {
		log.Debugf("skipping pod %s. hostNetwork=false", pod.Name)
		return
	}

	targets := annotations.TargetsFromTargetAnnotation(pod.Annotations)

	ps.addInternalHostnameAnnotationEndpoints(endpointMap, pod, targets)
	ps.addHostnameAnnotationEndpoints(endpointMap, pod, targets)
	ps.addKopsDNSControllerEndpoints(endpointMap, pod)
	ps.addPodSourceDomainEndpoints(endpointMap, pod, targets)
}

func (ps *podSource) addInternalHostnameAnnotationEndpoints(endpointMap map[endpoint.EndpointKey][]string, pod *v1.Pod, targets []string) {
	if domainAnnotation, ok := pod.Annotations[annotations.InternalHostnameKey]; ok {
		domainList := annotations.SplitHostnameAnnotation(domainAnnotation)
		for _, domain := range domainList {
			if len(targets) == 0 {
				addToEndpointMap(endpointMap, pod, domain, endpoint.SuitableType(pod.Status.PodIP), pod.Status.PodIP)
			} else {
				addTargetsToEndpointMap(endpointMap, pod, targets, domain)
			}
		}
	}
}

func (ps *podSource) addHostnameAnnotationEndpoints(endpointMap map[endpoint.EndpointKey][]string, pod *v1.Pod, targets []string) {
	if domainAnnotation, ok := pod.Annotations[annotations.HostnameKey]; ok {
		domainList := annotations.SplitHostnameAnnotation(domainAnnotation)
		if len(targets) == 0 {
			ps.addPodNodeEndpointsToEndpointMap(endpointMap, pod, domainList)
		} else {
			addTargetsToEndpointMap(endpointMap, pod, targets, domainList...)
		}
	}
}

func (ps *podSource) addKopsDNSControllerEndpoints(endpointMap map[endpoint.EndpointKey][]string, pod *v1.Pod) {
	if ps.compatibility == "kops-dns-controller" {
		if domainAnnotation, ok := pod.Annotations[kopsDNSControllerInternalHostnameAnnotationKey]; ok {
			domainList := annotations.SplitHostnameAnnotation(domainAnnotation)
			for _, domain := range domainList {
				addToEndpointMap(endpointMap, pod, domain, endpoint.SuitableType(pod.Status.PodIP), pod.Status.PodIP)
			}
		}

		if domainAnnotation, ok := pod.Annotations[kopsDNSControllerHostnameAnnotationKey]; ok {
			domainList := annotations.SplitHostnameAnnotation(domainAnnotation)
			ps.addPodNodeEndpointsToEndpointMap(endpointMap, pod, domainList)
		}
	}
}

func (ps *podSource) addPodSourceDomainEndpoints(endpointMap map[endpoint.EndpointKey][]string, pod *v1.Pod, targets []string) {
	if ps.podSourceDomain != "" {
		domain := pod.Name + "." + ps.podSourceDomain
		if len(targets) == 0 {
			addToEndpointMap(endpointMap, pod, domain, endpoint.SuitableType(pod.Status.PodIP), pod.Status.PodIP)
		}
		addTargetsToEndpointMap(endpointMap, pod, targets, domain)
	}
}

func (ps *podSource) addPodNodeEndpointsToEndpointMap(endpointMap map[endpoint.EndpointKey][]string, pod *v1.Pod, domainList []string) {
	node, err := ps.nodeInformer.Lister().Get(pod.Spec.NodeName)
	if err != nil {
		log.Debugf("Get node[%s] of pod[%s] error: %v; ignoring", pod.Spec.NodeName, pod.GetName(), err)
		return
	}
	for _, domain := range domainList {
		for _, address := range node.Status.Addresses {
			recordType := endpoint.SuitableType(address.Address)
			// IPv6 addresses are labeled as NodeInternalIP despite being usable externally as well.
			if address.Type == v1.NodeExternalIP || (address.Type == v1.NodeInternalIP && recordType == endpoint.RecordTypeAAAA) {
				addToEndpointMap(endpointMap, pod, domain, recordType, address.Address)
			}
		}
	}
}

func (ps *podSource) hostsFromTemplate(pod *v1.Pod) (map[endpoint.EndpointKey][]string, error) {
	hosts, err := ps.templateEngine.ExecFQDN(pod)
	if err != nil {
		return nil, err
	}

	result := make(map[endpoint.EndpointKey][]string)
	for _, target := range hosts {
		for _, address := range pod.Status.PodIPs {
			if address.IP == "" {
				log.Debugf("skipping pod %q. PodIP is empty with phase %q", pod.Name, pod.Status.Phase)
				continue
			}
			key := endpoint.EndpointKey{
				DNSName:    target,
				RecordType: endpoint.SuitableType(address.IP),
				RecordTTL:  annotations.TTLFromAnnotations(pod.Annotations, fmt.Sprintf("pod/%s", pod.Name)),
			}
			result[key] = append(result[key], address.IP)
		}
	}

	return result, nil
}

func addTargetsToEndpointMap(endpointMap map[endpoint.EndpointKey][]string, pod *v1.Pod, targets []string, domainList ...string) {
	for _, domain := range domainList {
		for _, target := range targets {
			addToEndpointMap(endpointMap, pod, domain, endpoint.SuitableType(target), target)
		}
	}
}

func addToEndpointMap(endpointMap map[endpoint.EndpointKey][]string, pod *v1.Pod, domain string, recordType string, address string) {
	key := endpoint.EndpointKey{
		DNSName:    domain,
		RecordType: recordType,
		RecordTTL:  annotations.TTLFromAnnotations(pod.Annotations, fmt.Sprintf("pod/%s", pod.Name)),
	}
	if _, ok := endpointMap[key]; !ok {
		endpointMap[key] = []string{}
	}
	endpointMap[key] = append(endpointMap[key], address)
}
