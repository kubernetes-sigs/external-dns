/*
Copyright 2017 The Kubernetes Authors.

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
	"bytes"
	"context"
	"fmt"
	"net/netip"
	"reflect"
	"strings"
	"text/template"
	"time"
	"unicode"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/utils"
)

const (
	controllerAnnotationKey       = annotations.ControllerKey
	hostnameAnnotationKey         = annotations.HostnameKey
	accessAnnotationKey           = annotations.AccessKey
	endpointsTypeAnnotationKey    = annotations.EndpointsTypeKey
	targetAnnotationKey           = annotations.TargetKey
	ttlAnnotationKey              = annotations.TtlKey
	aliasAnnotationKey            = annotations.AliasKey
	ingressHostnameSourceKey      = annotations.IngressHostnameSourceKey
	controllerAnnotationValue     = annotations.ControllerValue
	internalHostnameAnnotationKey = annotations.InternalHostnameKey
)

const (
	EndpointsTypeNodeExternalIP = "NodeExternalIP"
	EndpointsTypeHostIP         = "HostIP"
)

// Provider-specific annotations
const (
	// CloudflareProxiedKey The annotation used for determining if traffic will go through Cloudflare
	CloudflareProxiedKey        = annotations.CloudflareProxiedKey
	CloudflareCustomHostnameKey = annotations.CloudflareCustomHostnameKey

	SetIdentifierKey = annotations.SetIdentifierKey
)

// Source defines the interface Endpoint sources should implement.
type Source interface {
	Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error)
	// AddEventHandler adds an event handler that should be triggered if something in source changes
	AddEventHandler(context.Context, func())
}

func getTTLFromAnnotations(input map[string]string, resource string) endpoint.TTL {
	return annotations.TTLFromAnnotations(input, resource)
}

type kubeObject interface {
	runtime.Object
	metav1.Object
}

func execTemplate(tmpl *template.Template, obj kubeObject) (hostnames []string, err error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, obj); err != nil {
		kind := obj.GetObjectKind().GroupVersionKind().Kind
		return nil, fmt.Errorf("failed to apply template on %s %s/%s: %w", kind, obj.GetNamespace(), obj.GetName(), err)
	}
	for _, name := range strings.Split(buf.String(), ",") {
		name = strings.TrimFunc(name, unicode.IsSpace)
		name = strings.TrimSuffix(name, ".")
		hostnames = append(hostnames, name)
	}
	return hostnames, nil
}

func parseTemplate(fqdnTemplate string) (tmpl *template.Template, err error) {
	if fqdnTemplate == "" {
		return nil, nil
	}
	funcs := template.FuncMap{
		"trimPrefix": strings.TrimPrefix,
	}
	return template.New("endpoint").Funcs(funcs).Parse(fqdnTemplate)
}

func getHostnamesFromAnnotations(annotations map[string]string) []string {
	hostnameAnnotation, ok := annotations[hostnameAnnotationKey]
	if !ok {
		return nil
	}
	return splitHostnameAnnotation(hostnameAnnotation)
}

func getAccessFromAnnotations(annotations map[string]string) string {
	return annotations[accessAnnotationKey]
}

func getEndpointsTypeFromAnnotations(annotations map[string]string) string {
	return annotations[endpointsTypeAnnotationKey]
}

func getInternalHostnamesFromAnnotations(annotations map[string]string) []string {
	internalHostnameAnnotation, ok := annotations[internalHostnameAnnotationKey]
	if !ok {
		return nil
	}
	return splitHostnameAnnotation(internalHostnameAnnotation)
}

func splitHostnameAnnotation(annotation string) []string {
	return strings.Split(strings.Replace(annotation, " ", "", -1), ",")
}

func getProviderSpecificAnnotations(input map[string]string) (endpoint.ProviderSpecific, string) {
	return annotations.ProviderSpecificAnnotations(input)
}

// getTargetsFromTargetAnnotation gets endpoints from optional "target" annotation.
// Returns empty endpoints array if none are found.
func getTargetsFromTargetAnnotation(input map[string]string) endpoint.Targets {
	var targets endpoint.Targets

	// Get the desired hostname of the ingress from the annotation.
	targetAnnotation, ok := input[targetAnnotationKey]
	if ok && targetAnnotation != "" {
		// splits the hostname annotation and removes the trailing periods
		targetsList := strings.Split(strings.Replace(targetAnnotation, " ", "", -1), ",")
		for _, targetHostname := range targetsList {
			targetHostname = strings.TrimSuffix(targetHostname, ".")
			targets = append(targets, targetHostname)
		}
	}
	return targets
}

// suitableType returns the DNS resource record type suitable for the target.
// In this case type A/AAAA for IPs and type CNAME for everything else.
func suitableType(target string) string {
	netIP, err := netip.ParseAddr(target)
	if err == nil && netIP.Is4() {
		return endpoint.RecordTypeA
	} else if err == nil && netIP.Is6() {
		return endpoint.RecordTypeAAAA
	}
	return endpoint.RecordTypeCNAME
}

// endpointsForHostname returns the endpoint objects for each host-target combination.
func endpointsForHostname(hostname string, targets endpoint.Targets, ttl endpoint.TTL, providerSpecific endpoint.ProviderSpecific, setIdentifier string, resource string) []*endpoint.Endpoint {
	return utils.EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)
}

func getLabelSelector(annotationFilter string) (labels.Selector, error) {
	labelSelector, err := metav1.ParseToLabelSelector(annotationFilter)
	if err != nil {
		return nil, err
	}
	return metav1.LabelSelectorAsSelector(labelSelector)
}

func matchLabelSelector(selector labels.Selector, srcAnnotations map[string]string) bool {
	annotations := labels.Set(srcAnnotations)
	return selector.Matches(annotations)
}

type eventHandlerFunc func()

func (fn eventHandlerFunc) OnAdd(obj interface{}, isInInitialList bool) { fn() }
func (fn eventHandlerFunc) OnUpdate(oldObj, newObj interface{})         { fn() }
func (fn eventHandlerFunc) OnDelete(obj interface{})                    { fn() }

type informerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
}

func waitForCacheSync(ctx context.Context, factory informerFactory) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	for typ, done := range factory.WaitForCacheSync(ctx.Done()) {
		if !done {
			select {
			case <-ctx.Done():
				return fmt.Errorf("failed to sync %v: %v", typ, ctx.Err())
			default:
				return fmt.Errorf("failed to sync %v", typ)
			}
		}
	}
	return nil
}

type dynamicInformerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[schema.GroupVersionResource]bool
}

func waitForDynamicCacheSync(ctx context.Context, factory dynamicInformerFactory) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	for typ, done := range factory.WaitForCacheSync(ctx.Done()) {
		if !done {
			select {
			case <-ctx.Done():
				return fmt.Errorf("failed to sync %v: %v", typ, ctx.Err())
			default:
				return fmt.Errorf("failed to sync %v", typ)
			}
		}
	}
	return nil
}
