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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/compute/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"sigs.k8s.io/external-dns/endpoint"
)

// computeAddressSource is an implementation of Source that provides endpoints by listing
// ComputeAddresses and fetching an Address embedded in each Spec.
type computeAddressSource struct {
	crdClient        rest.Interface
	namespace        string
	crdResource      string
	codec            runtime.ParameterCodec
	annotationFilter string
	labelFilter      string
}

// NewComputeAddressSource creates a new computeAddressSource with the given config.
func NewComputeAddressSource(crdClient rest.Interface, namespace, annotationFilter, labelFilter string, scheme *runtime.Scheme) (Source, error) {
	return &computeAddressSource{
		crdResource:      "computeaddresses",
		namespace:        namespace,
		annotationFilter: annotationFilter,
		labelFilter:      labelFilter,
		crdClient:        crdClient,
		codec:            runtime.NewParameterCodec(scheme),
	}, nil
}

// NewCRDClientForComputeAddress return rest client for ComputeAddress CRDs
func NewCRDClientForComputeAddress(client kubernetes.Interface, kubeConfig string, apiServerURL string) (*rest.RESTClient, *runtime.Scheme, error) {
	group := v1beta1.ComputeAddressGVK.Group
	version := v1beta1.ComputeAddressGVK.Version
	apiVersion := group + "/" + version
	kind := v1beta1.ComputeAddressGVK.Kind

	return NewCRDClientForAPIVersionKind(client, kubeConfig, apiServerURL, apiVersion, kind)
}

func (cas *computeAddressSource) AddEventHandler(ctx context.Context, handler func()) {
}

// Endpoints returns endpoint objects.
func (cas *computeAddressSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints := []*endpoint.Endpoint{}

	var (
		computeAddressList *v1beta1.ComputeAddressList
		err                error
	)

	if cas.labelFilter != "" {
		computeAddressList, err = cas.List(ctx, &metav1.ListOptions{LabelSelector: cas.labelFilter})
	} else {
		computeAddressList, err = cas.List(ctx, &metav1.ListOptions{})
	}
	if err != nil {
		return nil, err
	}

	computeAddressList, err = cas.filterByAnnotations(computeAddressList)

	if err != nil {
		return nil, err
	}

	for _, computeAddress := range computeAddressList.Items {
		ttl, _ := getTTLFromAnnotations(computeAddress.ObjectMeta.Annotations)

		for _, hostname := range getHostnamesFromAnnotations(computeAddress.ObjectMeta.Annotations) {
			ep := endpoint.NewEndpointWithTTL(hostname, endpoint.RecordTypeA, ttl, *computeAddress.Spec.Address)

			if ep.Labels == nil {
				ep.Labels = endpoint.NewLabels()
			}

			cas.setResourceLabel(&computeAddress, endpoints)
			endpoints = append(endpoints, ep)
		}
	}

	return endpoints, nil
}

func (cas *computeAddressSource) setResourceLabel(ca *v1beta1.ComputeAddress, endpoints []*endpoint.Endpoint) {
	for _, ep := range endpoints {
		ep.Labels[endpoint.ResourceLabelKey] = fmt.Sprintf("computeaddress/%s/%s", ca.ObjectMeta.Namespace, ca.ObjectMeta.Name)
	}
}

func (cas *computeAddressSource) List(ctx context.Context, opts *metav1.ListOptions) (result *v1beta1.ComputeAddressList, err error) {
	result = &v1beta1.ComputeAddressList{}

	err = cas.crdClient.Get().
		Namespace(cas.namespace).
		Resource(cas.crdResource).
		VersionedParams(opts, cas.codec).
		Do(ctx).
		Into(result)

	return
}

// filterByAnnotations filters a ComputeAddressList by a given annotation selector.
func (cas *computeAddressSource) filterByAnnotations(computeAddressList *v1beta1.ComputeAddressList) (*v1beta1.ComputeAddressList, error) {
	labelSelector, err := metav1.ParseToLabelSelector(cas.annotationFilter)
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}

	// empty filter returns original list
	if selector.Empty() {
		return computeAddressList, nil
	}

	filteredList := v1beta1.ComputeAddressList{}

	for _, computeAddress := range computeAddressList.Items {
		// convert the ComputeAddress' annotations to an equivalent label selector
		annotations := labels.Set(computeAddress.ObjectMeta.Annotations)

		// include ComputeAddress if its annotations match the selector
		if selector.Matches(annotations) {
			filteredList.Items = append(filteredList.Items, computeAddress)
		}
	}

	return &filteredList, nil
}
