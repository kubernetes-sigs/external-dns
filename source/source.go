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
<<<<<<< HEAD
<<<<<<< HEAD
	// CloudflareProxiedKey The annotation used for determining if traffic will go through Cloudflare
	CloudflareProxiedKey        = annotations.CloudflareProxiedKey
	CloudflareCustomHostnameKey = annotations.CloudflareCustomHostnameKey
=======
	// The annotation used for determining if traffic will go through Cloudflare
	CloudflareProxiedKey        = "external-dns.alpha.kubernetes.io/cloudflare-proxied"
	CloudflareCustomHostnameKey = "external-dns.alpha.kubernetes.io/cloudflare-custom-hostname"
	CloudflareRegionKey         = "external-dns.alpha.kubernetes.io/cloudflare-region-key"
>>>>>>> c3f0cd66 (fix cloudflare regional hostnames)
=======
	// CloudflareProxiedKey The annotation used for determining if traffic will go through Cloudflare
	CloudflareProxiedKey        = annotations.CloudflareProxiedKey
	CloudflareCustomHostnameKey = annotations.CloudflareCustomHostnameKey
>>>>>>> 12ae6fea (chore(source): code cleanup)

	SetIdentifierKey = annotations.SetIdentifierKey
)

// Source defines the interface Endpoint sources should implement.
type Source interface {
	Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error)
	// AddEventHandler adds an event handler that should be triggered if something in source changes
	AddEventHandler(context.Context, func())
}

<<<<<<< HEAD
func getTTLFromAnnotations(input map[string]string, resource string) endpoint.TTL {
	return annotations.TTLFromAnnotations(input, resource)
<<<<<<< HEAD
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

func getHostnamesFromAnnotations(input map[string]string) []string {
	return annotations.HostnamesFromAnnotations(input)
}

func getAccessFromAnnotations(input map[string]string) string {
	return input[accessAnnotationKey]
}

func getEndpointsTypeFromAnnotations(input map[string]string) string {
	return input[endpointsTypeAnnotationKey]
}

func getInternalHostnamesFromAnnotations(input map[string]string) []string {
	return annotations.InternalHostnamesFromAnnotations(input)
}

<<<<<<< HEAD
func getProviderSpecificAnnotations(input map[string]string) (endpoint.ProviderSpecific, string) {
	return annotations.ProviderSpecificAnnotations(input)
=======
func splitHostnameAnnotation(annotation string) []string {
	return strings.Split(strings.ReplaceAll(annotation, " ", ""), ",")
}

func getAliasFromAnnotations(annotations map[string]string) bool {
	aliasAnnotation, exists := annotations[aliasAnnotationKey]
	return exists && aliasAnnotation == "true"
}

func getProviderSpecificAnnotations(annotations map[string]string) (endpoint.ProviderSpecific, string) {
	providerSpecificAnnotations := endpoint.ProviderSpecific{}

	if v, exists := annotations[CloudflareProxiedKey]; exists {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  CloudflareProxiedKey,
			Value: v,
		})
	}
	if v, exists := annotations[CloudflareCustomHostnameKey]; exists {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  CloudflareCustomHostnameKey,
			Value: v,
		})
	}
	if v, exists := annotations[CloudflareRegionKey]; exists {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  CloudflareRegionKey,
			Value: v,
		})
	}
	if getAliasFromAnnotations(annotations) {
		providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
			Name:  "alias",
			Value: "true",
		})
	}
	setIdentifier := ""
	for k, v := range annotations {
		if k == SetIdentifierKey {
			setIdentifier = v
		} else if strings.HasPrefix(k, "external-dns.alpha.kubernetes.io/aws-") {
			attr := strings.TrimPrefix(k, "external-dns.alpha.kubernetes.io/aws-")
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("aws/%s", attr),
				Value: v,
			})
		} else if strings.HasPrefix(k, "external-dns.alpha.kubernetes.io/scw-") {
			attr := strings.TrimPrefix(k, "external-dns.alpha.kubernetes.io/scw-")
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("scw/%s", attr),
				Value: v,
			})
		} else if strings.HasPrefix(k, "external-dns.alpha.kubernetes.io/ibmcloud-") {
			attr := strings.TrimPrefix(k, "external-dns.alpha.kubernetes.io/ibmcloud-")
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("ibmcloud-%s", attr),
				Value: v,
			})
		} else if strings.HasPrefix(k, "external-dns.alpha.kubernetes.io/webhook-") {
			// Support for wildcard annotations for webhook providers
			attr := strings.TrimPrefix(k, "external-dns.alpha.kubernetes.io/webhook-")
			providerSpecificAnnotations = append(providerSpecificAnnotations, endpoint.ProviderSpecificProperty{
				Name:  fmt.Sprintf("webhook/%s", attr),
				Value: v,
			})
		}
	}
	return providerSpecificAnnotations, setIdentifier
}

// getTargetsFromTargetAnnotation gets endpoints from optional "target" annotation.
// Returns empty endpoints array if none are found.
func getTargetsFromTargetAnnotation(annotations map[string]string) endpoint.Targets {
	var targets endpoint.Targets

	// Get the desired hostname of the ingress from the annotation.
	targetAnnotation, exists := annotations[targetAnnotationKey]
	if exists && targetAnnotation != "" {
		// splits the hostname annotation and removes the trailing periods
		targetsList := strings.Split(strings.ReplaceAll(targetAnnotation, " ", ""), ",")
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
>>>>>>> c3f0cd66 (fix cloudflare regional hostnames)
}

// endpointsForHostname returns the endpoint objects for each host-target combination.
func endpointsForHostname(hostname string, targets endpoint.Targets, ttl endpoint.TTL, providerSpecific endpoint.ProviderSpecific, setIdentifier string, resource string) []*endpoint.Endpoint {
	return EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)
}

func getLabelSelector(annotationFilter string) (labels.Selector, error) {
	labelSelector, err := metav1.ParseToLabelSelector(annotationFilter)
	if err != nil {
		return nil, err
	}
	return metav1.LabelSelectorAsSelector(labelSelector)
}

func matchLabelSelector(selector labels.Selector, srcAnnotations map[string]string) bool {
	return selector.Matches(labels.Set(srcAnnotations))
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
=======
}

=======
>>>>>>> 49bc71af (chore(source): code cleanup)
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

func getAccessFromAnnotations(input map[string]string) string {
	return input[accessAnnotationKey]
}

func getEndpointsTypeFromAnnotations(input map[string]string) string {
	return input[endpointsTypeAnnotationKey]
}

// endpointsForHostname returns the endpoint objects for each host-target combination.
func endpointsForHostname(hostname string, targets endpoint.Targets, ttl endpoint.TTL, providerSpecific endpoint.ProviderSpecific, setIdentifier string, resource string) []*endpoint.Endpoint {
	return EndpointsForHostname(hostname, targets, ttl, providerSpecific, setIdentifier, resource)
}

func getLabelSelector(annotationFilter string) (labels.Selector, error) {
	labelSelector, err := metav1.ParseToLabelSelector(annotationFilter)
	if err != nil {
		return nil, err
	}
	return metav1.LabelSelectorAsSelector(labelSelector)
}

func matchLabelSelector(selector labels.Selector, srcAnnotations map[string]string) bool {
	return selector.Matches(labels.Set(srcAnnotations))
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
>>>>>>> 12ae6fea (chore(source): code cleanup)
	return nil
}
