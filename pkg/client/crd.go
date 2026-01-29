/*
Copyright 2026 The Kubernetes Authors.

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

package kubeclient

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
)

// NewCRDClientForAPIVersionKind return rest client for the given apiVersion and kind of the CRD
// TODO: consider clientcmd.NewDefaultClientConfigLoadingRules() with clientcmd.NewNonInteractiveDeferredLoadingClientConfig
func NewCRDClientForAPIVersionKind(
	client kubernetes.Interface,
	kubeConfig, apiServerURL, apiVersion, kind string) (*rest.RESTClient, error) {
	if kubeConfig == "" {
		if _, err := os.Stat(clientcmd.RecommendedHomeFile); err == nil {
			kubeConfig = clientcmd.RecommendedHomeFile
		}
	}

	config, err := clientcmd.BuildConfigFromFlags(apiServerURL, kubeConfig)
	if err != nil {
		return nil, err
	}

	groupVersion, err := schema.ParseGroupVersion(apiVersion)
	if err != nil {
		return nil, err
	}
	apiResourceList, err := client.Discovery().ServerResourcesForGroupVersion(groupVersion.String())
	if err != nil {
		return nil, fmt.Errorf("error listing resources in GroupVersion %q: %w", groupVersion.String(), err)
	}

	var crdAPIResource *metav1.APIResource
	for _, apiResource := range apiResourceList.APIResources {
		if apiResource.Kind == kind {
			crdAPIResource = &apiResource
			break
		}
	}
	if crdAPIResource == nil {
		return nil, fmt.Errorf("unable to find Resource Kind %q in GroupVersion %q", kind, apiVersion)
	}

	scheme := runtime.NewScheme()
	_ = apiv1alpha1.AddToScheme(scheme)

	config.GroupVersion = &groupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}

	crdClient, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		return nil, err
	}
	return crdClient, nil
}
