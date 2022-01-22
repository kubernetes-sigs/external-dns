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

	"sigs.k8s.io/external-dns/endpoint"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubeinformers "k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

type podSource struct {
	client        kubernetes.Interface
	namespace     string
	podInformer   coreinformers.PodInformer
	nodeInformer  coreinformers.NodeInformer
	compatibility string
}

// NewPodSource creates a new podSource with the given config.
func NewPodSource(ctx context.Context, kubeClient kubernetes.Interface, namespace string, compatibility string) (Source, error) {
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
	if err := waitForCacheSync(context.Background(), informerFactory); err != nil {
		return nil, err
	}

	return &podSource{
		client:        kubeClient,
		podInformer:   podInformer,
		nodeInformer:  nodeInformer,
		namespace:     namespace,
		compatibility: compatibility,
	}, nil
}

func (*podSource) AddEventHandler(ctx context.Context, handler func()) {

}

func (ps *podSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	pods, err := ps.podInformer.Lister().Pods(ps.namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	domains := make(map[string][]string)
	for _, pod := range pods {
		if !pod.Spec.HostNetwork {
			log.Debugf("skipping pod %s. hostNetwork=false", pod.Name)
			continue
		}

		if domain, ok := pod.Annotations[internalHostnameAnnotationKey]; ok {
			if _, ok := domains[domain]; !ok {
				domains[domain] = []string{}
			}
			domains[domain] = append(domains[domain], pod.Status.PodIP)
		}

		if domain, ok := pod.Annotations[hostnameAnnotationKey]; ok {
			if _, ok := domains[domain]; !ok {
				domains[domain] = []string{}
			}

			node, _ := ps.nodeInformer.Lister().Get(pod.Spec.NodeName)
			for _, address := range node.Status.Addresses {
				if address.Type == corev1.NodeExternalIP {
					domains[domain] = append(domains[domain], address.Address)
				}
			}
		}

		if ps.compatibility == "kops-dns-controller" {
			if domain, ok := pod.Annotations[kopsDNSControllerInternalHostnameAnnotationKey]; ok {
				if _, ok := domains[domain]; !ok {
					domains[domain] = []string{}
				}
				domains[domain] = append(domains[domain], pod.Status.PodIP)
			}

			if domain, ok := pod.Annotations[kopsDNSControllerHostnameAnnotationKey]; ok {
				if _, ok := domains[domain]; !ok {
					domains[domain] = []string{}
				}

				node, _ := ps.nodeInformer.Lister().Get(pod.Spec.NodeName)
				for _, address := range node.Status.Addresses {
					if address.Type == corev1.NodeExternalIP {
						domains[domain] = append(domains[domain], address.Address)
					}
				}
			}
		}
	}
	endpoints := []*endpoint.Endpoint{}
	for domain, targets := range domains {
		endpoints = append(endpoints, endpoint.NewEndpoint(domain, endpoint.RecordTypeA, targets...))
	}
	return endpoints, nil
}
