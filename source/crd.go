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
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"sigs.k8s.io/external-dns/endpoint"
)

// crdSource is an implementation of Source that provides endpoints by listing
// specified CRD and fetching Endpoints embedded in Spec.
type crdSource struct {
	crdClient   rest.Interface
	namespace   string
	crdResource string
	codec       runtime.ParameterCodec
}

func addKnownTypes(scheme *runtime.Scheme, groupVersion schema.GroupVersion) error {
	scheme.AddKnownTypes(groupVersion,
		&endpoint.DNSEndpoint{},
		&endpoint.DNSEndpointList{},
	)
	metav1.AddToGroupVersion(scheme, groupVersion)
	return nil
}

// NewCRDClientForAPIVersionKind return rest client for the given apiVersion and kind of the CRD
func NewCRDClientForAPIVersionKind(client kubernetes.Interface, kubeConfig, kubeMaster, apiVersion, kind string) (*rest.RESTClient, *runtime.Scheme, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	config, err := clientcmd.BuildConfigFromFlags(kubeMaster, kubeConfig)
	if err != nil {
		return nil, nil, err
	}

	groupVersion, err := schema.ParseGroupVersion(apiVersion)
	if err != nil {
		return nil, nil, err
	}
	apiResourceList, err := client.Discovery().ServerResourcesForGroupVersion(groupVersion.String())
	if err != nil {
		return nil, nil, fmt.Errorf("error listing resources in GroupVersion %q: %s", groupVersion.String(), err)
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
	addKnownTypes(scheme, groupVersion)

	config.ContentConfig.GroupVersion = &groupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}

	crdClient, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		return nil, nil, err
	}
	return crdClient, scheme, nil
}

// NewCRDSource creates a new crdSource with the given config.
func NewCRDSource(crdClient rest.Interface, namespace, kind string, scheme *runtime.Scheme) (Source, error) {
	return &crdSource{
		crdResource: strings.ToLower(kind) + "s",
		namespace:   namespace,
		crdClient:   crdClient,
		codec:       runtime.NewParameterCodec(scheme),
	}, nil
}

func (cs *crdSource) AddEventHandler(handler func() error, stopChan <-chan struct{}, minInterval time.Duration) {
}

// Endpoints returns endpoint objects.
func (cs *crdSource) Endpoints() ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	result, err := cs.List(&metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, dnsEndpoint := range result.Items {
		// Make sure that all endpoints have targets for A or CNAME type
		crdEndpoints := []*endpoint.Endpoint{}
		for _, ep := range dnsEndpoint.Spec.Endpoints {
			if (ep.RecordType == "CNAME" || ep.RecordType == "A" || ep.RecordType == "AAAA") && len(ep.Targets) < 1 {
				log.Warnf("Endpoint %s with DNSName %s has an empty list of targets", dnsEndpoint.ObjectMeta.Name, ep.DNSName)
				continue
			}

			illegalTarget := false
			for _, target := range ep.Targets {
				if strings.HasSuffix(target, ".") {
					illegalTarget = true
					break
				}
			}
			if illegalTarget {
				log.Warnf("Endpoint %s with DNSName %s has an illegal target. The subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com')", dnsEndpoint.ObjectMeta.Name, ep.DNSName)
				continue
			}

			if ep.Labels == nil {
				ep.Labels = endpoint.NewLabels()
			}

			crdEndpoints = append(crdEndpoints, ep)
		}

		cs.setResourceLabel(&dnsEndpoint, crdEndpoints)
		endpoints = append(endpoints, crdEndpoints...)

		if dnsEndpoint.Status.ObservedGeneration == dnsEndpoint.Generation {
			continue
		}

		dnsEndpoint.Status.ObservedGeneration = dnsEndpoint.Generation
		// Update the ObservedGeneration
		_, err = cs.UpdateStatus(&dnsEndpoint)
		if err != nil {
			log.Warnf("Could not update ObservedGeneration of the CRD: %v", err)
		}
	}

	return endpoints, nil
}

func (cs *crdSource) setResourceLabel(crd *endpoint.DNSEndpoint, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("crd/%s/%s", crd.ObjectMeta.Namespace, crd.ObjectMeta.Name)
	}
}

func (cs *crdSource) List(opts *metav1.ListOptions) (result *endpoint.DNSEndpointList, err error) {
	result = &endpoint.DNSEndpointList{}
	err = cs.crdClient.Get().
		Namespace(cs.namespace).
		Resource(cs.crdResource).
		VersionedParams(opts, cs.codec).
		Do().
		Into(result)
	return
}

func (cs *crdSource) UpdateStatus(dnsEndpoint *endpoint.DNSEndpoint) (result *endpoint.DNSEndpoint, err error) {
	result = &endpoint.DNSEndpoint{}
	err = cs.crdClient.Put().
		Namespace(dnsEndpoint.Namespace).
		Resource(cs.crdResource).
		Name(dnsEndpoint.Name).
		SubResource("status").
		Body(dnsEndpoint).
		Do().
		Into(result)
	return
}
