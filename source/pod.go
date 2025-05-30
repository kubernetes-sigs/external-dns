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

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
)

type podSource struct {
	client                   kubernetes.Interface
	namespace                string
	podInformer              coreinformers.PodInformer
	nodeInformer             coreinformers.NodeInformer
	compatibility            string
	ignoreNonHostNetworkPods bool
	podSourceDomain          string
}

// NewPodSource creates a new podSource with the given config.
func NewPodSource(ctx context.Context, kubeClient kubernetes.Interface, namespace string, compatibility string, ignoreNonHostNetworkPods bool, podSourceDomain string) (Source, error) {
	informerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(kubeClient, 0, kubeinformers.WithNamespace(namespace))
	podInformer := informerFactory.Core().V1().Pods()
	nodeInformer := informerFactory.Core().V1().Nodes()

	podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)
	nodeInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
			},
		},
	)

	informerFactory.Start(ctx.Done())

	// wait for the local cache to be populated.
	if err := informers.WaitForCacheSync(context.Background(), informerFactory); err != nil {
		return nil, err
	}

	return &podSource{
		client:                   kubeClient,
		podInformer:              podInformer,
		nodeInformer:             nodeInformer,
		namespace:                namespace,
		compatibility:            compatibility,
		ignoreNonHostNetworkPods: ignoreNonHostNetworkPods,
		podSourceDomain:          podSourceDomain,
	}, nil
}

func (*podSource) AddEventHandler(ctx context.Context, handler func()) {
}

func (ps *podSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	pods, err := ps.podInformer.Lister().Pods(ps.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	endpointMap := make(map[endpoint.EndpointKey][]string)
	for _, pod := range pods {
		if ps.ignoreNonHostNetworkPods && !pod.Spec.HostNetwork {
			log.Debugf("skipping pod %s. hostNetwork=false", pod.Name)
			continue
		}

		targets := annotations.TargetsFromTargetAnnotation(pod.Annotations)

		if domainAnnotation, ok := pod.Annotations[internalHostnameAnnotationKey]; ok {
			domainList := annotations.SplitHostnameAnnotation(domainAnnotation)
			for _, domain := range domainList {
				if len(targets) == 0 {
					addToEndpointMap(endpointMap, domain, suitableType(pod.Status.PodIP), pod.Status.PodIP)
				} else {
					for _, target := range targets {
						addToEndpointMap(endpointMap, domain, suitableType(target), target)
					}
				}
			}
		}

		if domainAnnotation, ok := pod.Annotations[hostnameAnnotationKey]; ok {
			domainList := annotations.SplitHostnameAnnotation(domainAnnotation)
			for _, domain := range domainList {
				if len(targets) == 0 {
					node, _ := ps.nodeInformer.Lister().Get(pod.Spec.NodeName)
					for _, address := range node.Status.Addresses {
						recordType := suitableType(address.Address)
						// IPv6 addresses are labeled as NodeInternalIP despite being usable externally as well.
						if address.Type == corev1.NodeExternalIP || (address.Type == corev1.NodeInternalIP && recordType == endpoint.RecordTypeAAAA) {
							addToEndpointMap(endpointMap, domain, recordType, address.Address)
						}
					}
				} else {
					for _, target := range targets {
						addToEndpointMap(endpointMap, domain, suitableType(target), target)
					}
				}
			}
		}

		if ps.compatibility == "kops-dns-controller" {
			if domainAnnotation, ok := pod.Annotations[kopsDNSControllerInternalHostnameAnnotationKey]; ok {
				domainList := annotations.SplitHostnameAnnotation(domainAnnotation)
				for _, domain := range domainList {
					addToEndpointMap(endpointMap, domain, suitableType(pod.Status.PodIP), pod.Status.PodIP)
				}
			}

			if domainAnnotation, ok := pod.Annotations[kopsDNSControllerHostnameAnnotationKey]; ok {
				domainList := annotations.SplitHostnameAnnotation(domainAnnotation)
				for _, domain := range domainList {
					node, _ := ps.nodeInformer.Lister().Get(pod.Spec.NodeName)
					for _, address := range node.Status.Addresses {
						recordType := suitableType(address.Address)
						// IPv6 addresses are labeled as NodeInternalIP despite being usable externally as well.
						if address.Type == corev1.NodeExternalIP || (address.Type == corev1.NodeInternalIP && recordType == endpoint.RecordTypeAAAA) {
							addToEndpointMap(endpointMap, domain, recordType, address.Address)
						}
					}
				}
			}
		}

		if ps.podSourceDomain != "" {
			domain := pod.Name + "." + ps.podSourceDomain
			if len(targets) == 0 {
				addToEndpointMap(endpointMap, domain, suitableType(pod.Status.PodIP), pod.Status.PodIP)
			}
			for _, target := range targets {
				addToEndpointMap(endpointMap, domain, suitableType(target), target)
			}
		}
	}
	var endpoints []*endpoint.Endpoint
	for key, targets := range endpointMap {
		endpoints = append(endpoints, endpoint.NewEndpoint(key.DNSName, key.RecordType, targets...))
	}
	return endpoints, nil
}

func addToEndpointMap(endpointMap map[endpoint.EndpointKey][]string, domain string, recordType string, address string) {
	key := endpoint.EndpointKey{
		DNSName:    domain,
		RecordType: recordType,
	}
	if _, ok := endpointMap[key]; !ok {
		endpointMap[key] = []string{}
	}
	endpointMap[key] = append(endpointMap[key], address)
}
