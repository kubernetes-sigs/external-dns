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

package registry

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	crcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/registry"
	"sigs.k8s.io/external-dns/source"
)

// CRDRegistry implements registry interface with ownership implemented via associated custom resource records (DNSRecord)
type CRDRegistry struct {
	// crReader serves reads from a controller-runtime cache (informer-backed).
	crReader client.Reader
	// crWriter performs create/update/delete against the API server.
	crWriter client.Client
	// informer warms the DNSRecord watch backing crReader; the registry does
	// not subscribe to events, reads are always served from the synced cache.
	informer crcache.Informer

	namespace string
	provider  provider.Provider
	ownerID   string // refers to the owner id of the current instance
}

func New(cfg *externaldns.Config, p provider.Provider) (registry.Registry, error) {
	return NewCRDRegistry(p, cfg.KubeConfig, cfg.APIServerURL, cfg.Namespace, cfg.TXTOwnerID, cfg.RequestTimeout)
}

// NewCRDRegistry returns new CRDRegistry object backed by a controller-runtime
// cache (reads) and client (writes).
func NewCRDRegistry(provider provider.Provider, kubeConfig, apiServerURL, namespace, ownerID string, apiServerTimeout time.Duration) (*CRDRegistry, error) {
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
		KubeConfig:     kubeConfig,
		APIServerURL:   apiServerURL,
		RequestTimeout: apiServerTimeout,
	}

	restConfig, err := clientGenerator.RESTConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to build rest config: %w", err)
	}

	opts, err := buildCacheOptions(namespace)
	if err != nil {
		return nil, err
	}

	c, err := crcache.New(restConfig, opts)
	if err != nil {
		return nil, fmt.Errorf("unable to create cache: %w", err)
	}

	// crWriter is used exclusively for writes; reads come from the cache.
	crWriter, err := client.New(restConfig, client.Options{Scheme: opts.Scheme})
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	// The cache lives for the whole process lifetime, mirroring the registry.
	ctx := context.Background()
	inf, err := c.GetInformer(ctx, &apiv1alpha1.DNSRecord{})
	if err != nil {
		return nil, fmt.Errorf("unable to get informer: %w", err)
	}
	if err := startAndSync(ctx, c); err != nil {
		return nil, err
	}

	return &CRDRegistry{
		crReader:  c,
		crWriter:  crWriter,
		informer:  inf,
		namespace: namespace,
		provider:  provider,
		ownerID:   ownerID,
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
	var records apiv1alpha1.DNSRecordList
	if err := cr.crReader.List(ctx, &records,
		client.InNamespace(cr.namespace),
		client.MatchingLabels{apiv1alpha1.RecordOwnerLabel: cr.ownerID},
	); err != nil {
		return []*endpoint.Endpoint{}, err
	}

	endpoints := make([]*endpoint.Endpoint, 0, len(records.Items))
	for i := range records.Items {
		endpoints = append(endpoints, &records.Items[i].Spec.Endpoint)
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
		if r.Labels == nil {
			r.Labels = endpoint.NewLabels()
		}
		r.Labels[endpoint.OwnerLabelKey] = cr.ownerID
		if err := cr.createDNSRecord(ctx, r); err != nil {
			return err
		}
	}

	for _, r := range filteredChanges.Delete {
		records, err := cr.getDNSRecords(ctx, r)
		if err != nil {
			return fmt.Errorf("unable to get DNSRecord of %s: %w", r, err)
		}

		// While this is a list, it is expected that this call will return 0 or 1 records.
		for i := range records.Items {
			e := &records.Items[i]
			if err := cr.crWriter.Delete(ctx, e); err != nil {
				// Ignore not found as it's a benign error, the record isn't present and it's the end goal here, to remove
				// all records. All other errors should surface back to the user.
				if !k8sErrors.IsNotFound(err) {
					return fmt.Errorf("unable to delete DNSRecord %s in %s: %w", e.Name, cr.namespace, err)
				}
			}
		}
	}

	// Update existing DNS records to reflect the newest change.
	for i, e := range filteredChanges.UpdateNew {
		old := filteredChanges.UpdateOld[i]
		records, err := cr.getDNSRecords(ctx, old)
		if err != nil {
			return fmt.Errorf("unable to get DNSRecord of %s: %w", old, err)
		}

		// While this is a list, it is expected that this call will return 0 or 1 records.
		for j := range records.Items {
			record := &records.Items[j]
			record.Spec.Endpoint = *e
			if err := cr.crWriter.Update(ctx, record); err != nil {
				return fmt.Errorf("unable to update DNSRecord %s in %s: %w", record.Name, record.Namespace, err)
			}
		}
	}

	err := cr.provider.ApplyChanges(ctx, filteredChanges)
	if err != nil {
		return fmt.Errorf("provider cannot apply changes: %w", err)
	}
	return cr.adjustLabelsFromProvider(ctx)
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider
func (cr *CRDRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return cr.provider.AdjustEndpoints(endpoints)
}

// getDNSRecords retrieve k8s DNSRecords resources from k8s api
func (cr *CRDRegistry) getDNSRecords(ctx context.Context, record *endpoint.Endpoint) (*apiv1alpha1.DNSRecordList, error) {
	var records apiv1alpha1.DNSRecordList
	err := cr.crReader.List(ctx, &records,
		client.InNamespace(cr.namespace),
		client.MatchingLabels{
			apiv1alpha1.RecordOwnerLabel:         cr.ownerID,
			apiv1alpha1.RecordNameLabel:          record.DNSName,
			apiv1alpha1.RecordTypeLabel:          record.RecordType,
			apiv1alpha1.RecordSetIdentifierLabel: record.SetIdentifier,
		},
	)
	return &records, err
}

// createDNSRecord create a new DNSRecord with k8s API
func (cr *CRDRegistry) createDNSRecord(ctx context.Context, record *endpoint.Endpoint) error {
	// name has to follow rfc 1123
	dnsname := strings.ReplaceAll(record.DNSName, ".", "-")
	dnsrecord := &apiv1alpha1.DNSRecord{
		ObjectMeta: metav1.ObjectMeta{
			Name:      strings.ToLower(fmt.Sprintf("%s-%s", dnsname, record.RecordType)),
			Namespace: cr.namespace,
			Labels: map[string]string{
				apiv1alpha1.RecordOwnerLabel:         cr.OwnerID(),
				apiv1alpha1.RecordNameLabel:          record.DNSName,
				apiv1alpha1.RecordTypeLabel:          record.RecordType,
				apiv1alpha1.RecordSetIdentifierLabel: record.SetIdentifier,
			},
		},
		Spec: apiv1alpha1.DNSRecordSpec{
			Endpoint: *record,
		},
	}

	if err := cr.crWriter.Create(ctx, dnsrecord); err != nil {
		// It could be possible that a record already exists if a previous apply change happened
		// and there was an error while creating those records through the provider. For that reason,
		// this error is ignored, all others will be surfaced back to the user
		if !k8sErrors.IsAlreadyExists(err) {
			return fmt.Errorf("unable to create DNSRecord %s in %s: %w", dnsrecord.Name, dnsrecord.Namespace, err)
		}
	}
	return nil
}

// adjustLabelsFromProvider ensures labels in CRD registry are accurate
// It should be called after applyChanges
func (cr *CRDRegistry) adjustLabelsFromProvider(ctx context.Context) error {
	records, err := cr.provider.Records(ctx)
	if err != nil {
		return fmt.Errorf("unable to get records from provider: %w", err)
	}

	for _, record := range records {
		dnsrecords, err := cr.getDNSRecords(ctx, record)
		if err != nil {
			return fmt.Errorf("unable to list DNSRecord for %s in %s: %w", record.DNSName, cr.namespace, err)
		}

		for i := range dnsrecords.Items {
			dnsrecord := &dnsrecords.Items[i]
			if !maps.Equal(dnsrecord.Spec.Endpoint.Labels, record.Labels) {
				log.Debug("update DNSRecord with modified labels from provider")
				if err := cr.updateDNSRecordWithEndpointLabels(ctx, dnsrecord, record); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (cr *CRDRegistry) updateDNSRecordWithEndpointLabels(ctx context.Context, dnsrecord *apiv1alpha1.DNSRecord, record *endpoint.Endpoint) error {
	// safety net on Resource & Owner labels
	resource := dnsrecord.Spec.Endpoint.Labels[endpoint.ResourceLabelKey]
	dnsrecord.Spec.Endpoint.Labels = record.Labels
	dnsrecord.Spec.Endpoint.Labels[endpoint.OwnerLabelKey] = cr.ownerID
	if resource != "" {
		dnsrecord.Spec.Endpoint.Labels[endpoint.ResourceLabelKey] = resource
	}

	if err := cr.crWriter.Update(ctx, dnsrecord); err != nil {
		return fmt.Errorf("unable to update DNSRecord %s: %w", dnsrecord.Name, err)
	}
	return nil
}

// buildCacheOptions constructs the controller-runtime cache options scoped to
// the given namespace, with the DNSRecord type registered in the scheme.
func buildCacheOptions(namespace string) (crcache.Options, error) {
	scheme := runtime.NewScheme()
	if err := apiv1alpha1.AddToScheme(scheme); err != nil {
		return crcache.Options{}, err
	}
	// metav1.AddToGroupVersion registers ListOptions (and other meta types) so
	// that watch/list requests for this group can be encoded.
	metav1.AddToGroupVersion(scheme, apiv1alpha1.GroupVersion)

	return crcache.Options{
		Scheme: scheme,
		ByObject: map[client.Object]crcache.ByObject{
			&apiv1alpha1.DNSRecord{}: {
				Namespaces: map[string]crcache.Config{
					namespace: {}, // "" == NamespaceAll
				},
			},
		},
	}, nil
}

// startAndSync starts the cache in a goroutine and waits for it to sync.
// Returns an error if the cache fails to start or sync.
func startAndSync(ctx context.Context, c crcache.Cache) error {
	errCh := make(chan error, 1)
	go func() { errCh <- c.Start(ctx) }()
	if !c.WaitForCacheSync(ctx) {
		select {
		case err := <-errCh:
			if err != nil {
				return fmt.Errorf("cache failed to sync: %w", err)
			}
			return fmt.Errorf("cache failed to sync")
		case <-ctx.Done():
			return fmt.Errorf("cache failed to sync: %w", ctx.Err())
		}
	}
	return nil
}
