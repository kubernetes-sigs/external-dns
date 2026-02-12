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

package source

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"
	networkv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

var globalAcceleratorGVR = schema.GroupVersionResource{
	Group:    "aga.k8s.aws",
	Version:  "v1beta1",
	Resource: "globalaccelerators",
}

func ingressTargets(ctx context.Context, dynamicClient dynamic.Interface, ing *networkv1.Ingress) endpoint.Targets {
	targets := annotations.TargetsFromTargetAnnotation(ing.Annotations)
	if len(targets) > 0 {
		return targets
	}

	gaTargets, ok := targetsFromGlobalAccelerator(ctx, dynamicClient, ing)
	if ok {
		return gaTargets
	}

	return targetsFromIngressStatus(ing.Status)
}

func targetsFromGlobalAccelerator(ctx context.Context, dynamicClient dynamic.Interface, ing *networkv1.Ingress) (endpoint.Targets, bool) {
	if dynamicClient == nil {
		return nil, false
	}

	acceleratorRef, ok := ing.Annotations[annotations.GlobalAcceleratorKey]
	if !ok {
		return nil, false
	}
	acceleratorRef = strings.TrimSpace(acceleratorRef)
	if acceleratorRef == "" {
		return nil, false
	}

	namespace, name, err := ParseIngress(acceleratorRef)
	if err != nil {
		log.Warnf("Ingress %s/%s has invalid %s annotation: %v", ing.Namespace, ing.Name, annotations.GlobalAcceleratorKey, err)
		return nil, false
	}
	if namespace == "" {
		namespace = ing.Namespace
	}

	accelerator, err := dynamicClient.Resource(globalAcceleratorGVR).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		log.Warnf("Failed to fetch Global Accelerator %s/%s for ingress %s/%s: %v", namespace, name, ing.Namespace, ing.Name, err)
		return nil, false
	}

	dnsName, found, err := unstructured.NestedString(accelerator.Object, "status", "dnsName")
	if err != nil {
		log.Warnf("Failed reading Global Accelerator %s/%s status.dnsName for ingress %s/%s: %v", namespace, name, ing.Namespace, ing.Name, err)
		return nil, false
	}
	dnsName = strings.TrimSpace(dnsName)
	if !found || dnsName == "" {
		return nil, false
	}

	return endpoint.Targets{strings.TrimSuffix(dnsName, ".")}, true
}
