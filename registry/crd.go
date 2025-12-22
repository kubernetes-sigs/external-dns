/*
Copyright 2025 The Kubernetes Authors.

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

package registry

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/source"
)

type CRDConfig struct {
	KubeConfig   string
	APIServerURL string
	APIVersion   string
	Kind         string
}

// The CRD interfaces are built as k8s' rest.Interface doesn't have proper support for testing
// These interfaces exists so the runtime will use the rest.Interface but gives an
// option for writing tests without building a complete k8s client.
type CRDClient interface {
	Get() CRDRequest
	List() CRDRequest
	Put() CRDRequest
	Post() CRDRequest
	Delete() CRDRequest
}

type CRDRequest interface {
	Name(string) CRDRequest
	Namespace(string) CRDRequest
	Body(interface{}) CRDRequest
	Params(runtime.Object) CRDRequest
	Do(context.Context) CRDResult
}

type CRDResult interface {
	Error() error
	Into(runtime.Object) error
}

// CRDRegistry implements registry interface with ownership implemented via associated custom resource records (DNSRecord)
type CRDRegistry struct {
	client    CRDClient
	namespace string
	provider  provider.Provider
	ownerID   string // refers to the owner id of the current instance

	// cache the records in memory and update on an interval instead.
	recordsCache            []*endpoint.Endpoint
	recordsCacheRefreshTime time.Time
	cacheInterval           time.Duration
}

// NewCRDClientForAPIVersionKind return rest client for the given apiVersion and kind of the CRD
func NewCRDClientForAPIVersionKind(client kubernetes.Interface, kubeConfig, apiServerURL, apiVersion string) (CRDClient, error) {
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

	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(groupVersion,
		&apiv1alpha1.DNSRecord{},
		&apiv1alpha1.DNSRecordList{},
	)
	metav1.AddToGroupVersion(scheme, groupVersion)

	config.GroupVersion = &groupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}

	crdClient, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		return nil, err
	}

	apiResourceList, err := client.Discovery().ServerResourcesForGroupVersion(groupVersion.String())
	if err != nil {
		return nil, fmt.Errorf("error listing resources in GroupVersion %q: %w", groupVersion.String(), err)
	}

	var crdAPIResource *metav1.APIResource
	for _, apiResource := range apiResourceList.APIResources {
		if apiResource.Kind == "DNSRecord" {
			crdAPIResource = &apiResource
			break
		}
	}
	if crdAPIResource == nil {
		return nil, fmt.Errorf("unable to find Resource Kind %q in GroupVersion %q", "DNSRecord", apiVersion)
	}
	return &crdclient{scheme: scheme, resource: crdAPIResource, codec: runtime.NewParameterCodec(scheme), Interface: crdClient}, nil
}

// NewCRDRegistry returns new CRDRegistry object
func NewCRDRegistry(provider provider.Provider, kubeConfig, apiServerURL, apiVersion, namespace, ownerID string, cacheInterval, apiServerTimeOut time.Duration) (*CRDRegistry, error) {
	var err error
	var k8sClient kubernetes.Interface

	if ownerID == "" {
		return nil, errors.New("owner id cannot be empty")
	}

	if namespace == "" {
		log.Info("Registry: namespace not specified, using `default`")
		namespace = "default"
	}

	// new Singleton because the user may want to store this registry on a
	// remote (and shared) cluster between multiple external-dns instances
	clientGenerator := &source.SingletonClientGenerator{
		KubeConfig:   kubeConfig,
		APIServerURL: apiServerURL,
		// If update events are enabled, disable timeout.
		RequestTimeout: func() time.Duration {
			return apiServerTimeOut
		}(),
	}

	k8sClient, err = clientGenerator.KubeClient()
	if err != nil {
		return nil, fmt.Errorf("unable to create kubeclient: %w", err)
	}

	crdClient, err := NewCRDClientForAPIVersionKind(k8sClient, kubeConfig, apiServerURL, apiVersion)
	if err != nil {
		return nil, fmt.Errorf("unable to create crdclient: %w", err)
	}

	return &CRDRegistry{
		client:        crdClient,
		namespace:     namespace,
		provider:      provider,
		ownerID:       ownerID,
		cacheInterval: cacheInterval,
	}, nil
}

func (cr *CRDRegistry) GetDomainFilter() endpoint.DomainFilterInterface {
	return cr.provider.GetDomainFilter()
}

func (cr *CRDRegistry) OwnerID() string {
	return cr.ownerID
}

// Records returns the current records from the registry
func (cr *CRDRegistry) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	// If we have the zones cached AND we have refreshed the cache since the
	// last given interval, then just use the cached results.
	if cr.recordsCache != nil && time.Since(cr.recordsCacheRefreshTime) < cr.cacheInterval {
		log.Debug("Using cached records.")
		return cr.recordsCache, nil
	}

	endpoints := []*endpoint.Endpoint{}

	var records apiv1alpha1.DNSRecordList
	for more := true; more; more = records.Continue != "" {
		opts := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s", apiv1alpha1.RecordOwnerLabel, cr.ownerID),
		}

		err := cr.client.Get().Namespace(cr.namespace).Params(&opts).Do(ctx).Into(&records)
		if err != nil {
			return []*endpoint.Endpoint{}, err
		}

		for _, record := range records.Items {
			endpoints = append(endpoints, &record.Spec.Endpoint)
		}
	}

	// Update the cache.
	if cr.cacheInterval > 0 {
		cr.recordsCache = endpoints
		cr.recordsCacheRefreshTime = time.Now()
	}
	return endpoints, nil
}

// ApplyChanges updates dns provider with the changes and creates/updates/delete a DNSRecord accordingly.
func (cr *CRDRegistry) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	filteredChanges := &plan.Changes{
		Create:    changes.Create,
		UpdateNew: endpoint.FilterEndpointsByOwnerID(cr.ownerID, changes.UpdateNew),
		UpdateOld: endpoint.FilterEndpointsByOwnerID(cr.ownerID, changes.UpdateOld),
		Delete:    endpoint.FilterEndpointsByOwnerID(cr.ownerID, changes.Delete),
	}

	for _, r := range filteredChanges.Create {
		dnsname := strings.ReplaceAll(r.DNSName, ".", "-")
		r.Labels[endpoint.OwnerLabelKey] = cr.ownerID
		record := apiv1alpha1.DNSRecord{
			ObjectMeta: metav1.ObjectMeta{
				Name:      strings.ToLower(fmt.Sprintf("%s-%s", dnsname, r.RecordType)),
				Namespace: cr.namespace,
				Labels: map[string]string{
					apiv1alpha1.RecordOwnerLabel: cr.OwnerID(),
					apiv1alpha1.RecordNameLabel:  r.DNSName,
					apiv1alpha1.RecordTypeLabel:  r.RecordType,
					apiv1alpha1.RecordKeyLabel:   r.Key().String(),
				},
			},
			Spec: apiv1alpha1.DNSRecordSpec{
				Endpoint: *r,
			},
		}

		result := cr.client.Post().Namespace(cr.namespace).Body(&record).Do(ctx)
		if err := result.Error(); err != nil {
			// It could be possible that a record already exists if a previous apply change happened
			// and there was an error while creating those records through the provider. For that reason,
			// this error is ignored, all others will be surfaced back to the user
			if !k8sErrors.IsAlreadyExists(err) {
				return err
			}
		}

		if cr.cacheInterval > 0 {
			cr.addToCache(r)
		}
	}

	for _, r := range filteredChanges.Delete {
		var records apiv1alpha1.DNSRecordList
		opts := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s,%s=%s", apiv1alpha1.RecordKeyLabel, r.Key().String(), apiv1alpha1.RecordOwnerLabel, cr.ownerID),
		}

		err := cr.client.Get().Namespace(cr.namespace).Params(&opts).Do(ctx).Into(&records)
		if err != nil {
			return err
		}

		// While this is a list, it is expected that this call will return 0 or 1 records.
		for _, e := range records.Items {
			result := cr.client.Delete().Namespace(cr.namespace).Name(e.Name).Do(ctx)
			if err := result.Error(); err != nil {
				// Ignore not found as it's a benign error, the record isn't present and it's the end goal here, to remove
				// all records. All other errors should surface back to the user.
				if !k8sErrors.IsNotFound(err) {
					return err
				}
			}
		}

		if cr.cacheInterval > 0 {
			cr.removeFromCache(r)
		}
	}

	// Update existing DNS records to reflect the newest change.
	for i, e := range filteredChanges.UpdateNew {
		old := filteredChanges.UpdateOld[i]

		var records apiv1alpha1.DNSRecordList
		opts := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s,%s=%s", apiv1alpha1.RecordKeyLabel, old.Key().String(), apiv1alpha1.RecordOwnerLabel, cr.ownerID),
		}

		err := cr.client.Get().Namespace(cr.namespace).Params(&opts).Do(ctx).Into(&records)
		if err != nil {
			return err
		}

		for _, record := range records.Items {
			record.Spec.Endpoint = *e
			result := cr.client.Put().Namespace(cr.namespace).Name(record.Name).Body(&record).Do(ctx)
			if err := result.Error(); err != nil {
				return err
			}
		}

		if cr.cacheInterval > 0 {
			cr.addToCache(e)
		}

		if cr.cacheInterval > 0 {
			cr.removeFromCache(old)
		}
	}

	return cr.provider.ApplyChanges(ctx, filteredChanges)
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (cr *CRDRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return cr.provider.AdjustEndpoints(endpoints)
}

func (cr *CRDRegistry) addToCache(ep *endpoint.Endpoint) {
	if cr.recordsCache != nil {
		cr.recordsCache = append(cr.recordsCache, ep)
	}
}

func (cr *CRDRegistry) removeFromCache(ep *endpoint.Endpoint) {
	if cr.recordsCache == nil || ep == nil {
		return
	}

	for i, e := range cr.recordsCache {
		if e.DNSName == ep.DNSName && e.RecordType == ep.RecordType && e.SetIdentifier == ep.SetIdentifier && e.Targets.Same(ep.Targets) {
			// We found a match delete the endpoint from the cache.
			cr.recordsCache = append(cr.recordsCache[:i], cr.recordsCache[i+1:]...)
			return
		}
	}
}

type crdclient struct {
	scheme   *runtime.Scheme
	resource *metav1.APIResource
	codec    runtime.ParameterCodec
	rest.Interface
}

func (c crdclient) Get() CRDRequest {
	return &crdrequest{client: &c, method: "GET", resource: c.resource}
}

func (c crdclient) List() CRDRequest {
	return &crdrequest{client: &c, method: "LIST", resource: c.resource}
}

func (c crdclient) Post() CRDRequest {
	return &crdrequest{client: &c, method: "POST", resource: c.resource}
}

func (c crdclient) Put() CRDRequest {
	return &crdrequest{client: &c, method: "PUT", resource: c.resource}
}

func (c crdclient) Delete() CRDRequest {
	return &crdrequest{client: &c, method: "DELETE", resource: c.resource}
}

type crdrequest struct {
	client   *crdclient
	resource *metav1.APIResource

	method    string
	namespace string
	name      string
	params    runtime.Object
	body      interface{}
}

func (r *crdrequest) Name(name string) CRDRequest {
	r.name = name
	return r
}

func (r *crdrequest) Namespace(namespace string) CRDRequest {
	r.namespace = namespace
	return r
}

func (r *crdrequest) Params(obj runtime.Object) CRDRequest {
	r.params = obj
	return r
}

func (r *crdrequest) Body(obj interface{}) CRDRequest {
	r.body = obj
	return r
}

func (r *crdrequest) Do(ctx context.Context) CRDResult {
	var req *rest.Request
	switch r.method {
	case "POST":
		req = r.client.Interface.Post()
	case "PUT":
		req = r.client.Interface.Put()
	case "DELETE":
		req = r.client.Interface.Delete()
	default:
		req = r.client.Interface.Get()
	}

	req = req.Namespace(r.namespace).Resource(r.resource.Name)
	if r.name != "" {
		req = req.Name(r.name)
	}

	if r.params != nil {
		req = req.VersionedParams(r.params, r.client.codec)
	}

	if r.body != nil {
		req = req.Body(r.body)
	}

	result := req.Do(ctx)
	return &crdresult{result}
}

type crdresult struct {
	rest.Result
}

func (r crdresult) Error() error {
	return r.Result.Error()
}

func (r crdresult) Into(obj runtime.Object) error {
	return r.Result.Into(obj)
}
