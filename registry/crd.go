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

package registry

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/crds"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"

	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
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

// CRDRegistry implements registry interface with ownership implemented via associated custom resource records (DSNEntry)
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
		&crds.DNSEntry{},
		&crds.DNSEntryList{},
	)
	metav1.AddToGroupVersion(scheme, groupVersion)

	config.ContentConfig.GroupVersion = &groupVersion
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
		if apiResource.Kind == "DNSEntry" {
			crdAPIResource = &apiResource
			break
		}
	}
	if crdAPIResource == nil {
		return nil, fmt.Errorf("unable to find Resource Kind %q in GroupVersion %q", "DNSEntry", apiVersion)
	}
	return &crdclient{scheme: scheme, resource: crdAPIResource, codec: runtime.NewParameterCodec(scheme), Interface: crdClient}, nil
}

// NewCRDRegistry returns new CRDRegistry object
func NewCRDRegistry(provider provider.Provider, crdClient CRDClient, ownerID string, cacheInterval time.Duration, namespace string) (*CRDRegistry, error) {
	if ownerID == "" {
		return nil, errors.New("owner id cannot be empty")
	}

	if namespace == "" {
		log.Info("Registry: No namespace specified, using `default`")
		namespace = "default"
	}

	return &CRDRegistry{
		client:        crdClient,
		namespace:     namespace,
		provider:      provider,
		ownerID:       ownerID,
		cacheInterval: cacheInterval,
	}, nil
}

func (im *CRDRegistry) GetDomainFilter() endpoint.DomainFilter {
	return im.provider.GetDomainFilter()
}

func (im *CRDRegistry) OwnerID() string {
	return im.ownerID
}

// Records returns the current records from the registry
func (im *CRDRegistry) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	// If we have the zones cached AND we have refreshed the cache since the
	// last given interval, then just use the cached results.
	if im.recordsCache != nil && time.Since(im.recordsCacheRefreshTime) < im.cacheInterval {
		log.Debug("Using cached records.")
		return im.recordsCache, nil
	}

	records, err := im.provider.Records(ctx)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}

	for _, record := range records {
		// AWS Alias records have "new" format encoded as type "cname"
		if isAlias, found := record.GetProviderSpecificProperty("alias"); found && isAlias == "true" && record.RecordType == endpoint.RecordTypeA {
			record.RecordType = endpoint.RecordTypeCNAME
		}

		endpoints = append(endpoints, record)
	}

	var entries crds.DNSEntryList
	for more := true; more; more = entries.Continue != "" {
		opts := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s", crds.RegistryOwnerLabel, im.ownerID),
		}

		if entries.Continue != "" {
			opts.Continue = entries.Continue
		}

		// Populate the labels for each record with the RegistryEntry matching.
		err = im.client.Get().Namespace(im.namespace).Params(&opts).Do(ctx).Into(&entries)
		if err != nil {
			return nil, err
		}

		for _, entry := range entries.Items {
			for _, endpoint := range endpoints {
				if entry.IsEndpoint(endpoint) {
					endpoint.Labels = entry.EndpointLabels()
				}
			}
		}
	}

	// Update the cache.
	if im.cacheInterval > 0 {
		im.recordsCache = endpoints
		im.recordsCacheRefreshTime = time.Now()
	}

	return endpoints, nil
}

// ApplyChanges updates dns provider with the changes and creates/updates/delete a DNSEntry
// custom resource as the registry entry.
func (im *CRDRegistry) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	filteredChanges := &plan.Changes{
		Create:    changes.Create,
		UpdateNew: endpoint.FilterEndpointsByOwnerID(im.ownerID, changes.UpdateNew),
		UpdateOld: endpoint.FilterEndpointsByOwnerID(im.ownerID, changes.UpdateOld),
		Delete:    endpoint.FilterEndpointsByOwnerID(im.ownerID, changes.Delete),
	}

	for _, r := range filteredChanges.Create {
		entry := crds.DNSEntry{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s", im.OwnerID(), r.SetIdentifier),
				Namespace: im.namespace,
				Labels: map[string]string{
					crds.RegistryOwnerLabel:      im.OwnerID(),
					crds.RegistryRecordNameLabel: r.DNSName,
					crds.RegistryRecordTypeLabel: r.RecordType,
					crds.RegistryIdentifierLabel: r.SetIdentifier,
				},
			},
			Spec: crds.DNSEntrySpec{
				Endpoint: *r,
			},
		}

		result := im.client.Post().Namespace(im.namespace).Body(&entry).Do(ctx)
		if err := result.Error(); err != nil {
			// It could be possible that a record already exists if a previous apply change happened
			// and there was an error while creating those records through the provider. For that reason,
			// this error is ignored, all others will be surfaced back to the user
			if !k8sErrors.IsAlreadyExists(err) {
				return err
			}
		}

		if im.cacheInterval > 0 {
			im.addToCache(r)
		}
	}

	for _, r := range filteredChanges.Delete {
		var entries crds.DNSEntryList
		opts := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s,%s=%s", crds.RegistryIdentifierLabel, r.SetIdentifier, crds.RegistryOwnerLabel, im.ownerID),
		}

		err := im.client.Get().Namespace(im.namespace).Params(&opts).Do(ctx).Into(&entries)
		if err != nil {
			return err
		}

		// While this is a list, it is expected that this call will return 0 or 1 entries.
		for _, e := range entries.Items {
			result := im.client.Delete().Namespace(im.namespace).Name(e.Name).Do(ctx)
			if err := result.Error(); err != nil {
				// Ignore not found as it's a benign error, the entry record isn't present and it's the end goal here, to remove
				// all entries. All other errors should surface back to the user.
				if !k8sErrors.IsNotFound(err) {
					return err
				}
			}
		}

		if im.cacheInterval > 0 {
			im.removeFromCache(r)
		}
	}

	// Update existing DNS entries to reflect the newest change.
	for i, e := range filteredChanges.UpdateNew {
		old := filteredChanges.UpdateOld[i]

		var entries crds.DNSEntryList
		opts := metav1.ListOptions{
			LabelSelector: fmt.Sprintf("%s=%s,%s=%s", crds.RegistryIdentifierLabel, old.SetIdentifier, crds.RegistryOwnerLabel, im.ownerID),
		}

		err := im.client.Get().Namespace(im.namespace).Params(&opts).Do(ctx).Into(&entries)
		if err != nil {
			return err
		}

		for _, entry := range entries.Items {
			entry.Spec.Endpoint = *e
			result := im.client.Put().Namespace(im.namespace).Name(entry.Name).Body(&entry).Do(ctx)
			if err := result.Error(); err != nil {
				return err
			}
		}

		if im.cacheInterval > 0 {
			im.addToCache(e)
		}

		if im.cacheInterval > 0 {
			im.removeFromCache(old)
		}
	}

	return im.provider.ApplyChanges(ctx, filteredChanges)
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (im *CRDRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return im.provider.AdjustEndpoints(endpoints)
}

func (im *CRDRegistry) addToCache(ep *endpoint.Endpoint) {
	if im.recordsCache != nil {
		im.recordsCache = append(im.recordsCache, ep)
	}
}

func (im *CRDRegistry) removeFromCache(ep *endpoint.Endpoint) {
	if im.recordsCache == nil || ep == nil {
		return
	}

	for i, e := range im.recordsCache {
		if e.DNSName == ep.DNSName && e.RecordType == ep.RecordType && e.SetIdentifier == ep.SetIdentifier && e.Targets.Same(ep.Targets) {
			// We found a match delete the endpoint from the cache.
			im.recordsCache = append(im.recordsCache[:i], im.recordsCache[i+1:]...)
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
	var real *rest.Request
	switch r.method {
	case "POST":
		real = r.client.Interface.Post()
	case "PUT":
		real = r.client.Interface.Put()
	case "DELETE":
		real = r.client.Interface.Delete()
	default:
		real = r.client.Interface.Get()
	}

	real = real.Namespace(r.namespace).Resource(r.resource.Name)
	if r.name != "" {
		real = real.Name(r.name)
	}

	if r.params != nil {
		real = real.VersionedParams(r.params, r.client.codec)
	}

	if r.body != nil {
		real = real.Body(r.body)
	}

	result := real.Do(ctx)
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
