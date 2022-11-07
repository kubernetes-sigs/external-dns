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

package externaldns

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/source"
)

type ExternalDNSProviderConfig struct {
	// CaCrt should contain the value of the "ca.crt" field in a K8s ServiceAccount Secret
	CaCrt string
	// Namespace should contain the value of the "namespace" field in a K8s ServiceAccount Secret
	Namespace string
	// Token should contain the value of the "token" field in a K8s ServiceAccount Secret
	Token string

	// CaCrtPath should contain a filepath to the "ca.crt" field in a K8s ServiceAccount Secret. If set, will overwrite CaCrt
	CaCrtPath string
	// NamespacePath should contain a filepath to the "namespace" field in a K8s ServiceAccount Secret. If set, will overwrite Namespace
	NamespacePath string
	// TokenPath should contain a filepath to the "token" field in a K8s ServiceAccount Secret. If set, will overwrite Token
	TokenPath string

	// CrName is the name of the DNSEndpoint CR in the target cluster, it will reside in the Namespace namespace
	CrName string

	// KubernetesHost is the hostname of the target kubernetes cluster
	KubernetesHost string

	// KubernetesPort is the port (as string) of the target kubernetes cluster
	KubernetesPort string

	// CrdAPIVersion is the APIVersion of the DNSEndpoint CRD we use
	CrdAPIVersion string
	// CrdKind is the Kind of the DNSEndpoint CRD we use
	CrdKind string
}

// ExternalDNSProvider - dns provider that uses the DNSEndpoint CRD in a (remote)
// kubernetes cluster, e.g. a management cluster.
// The design is simple. There is a single DNSEndpoint CR in the target cluster,
// and it contains all the desired endpoints. Records() fetches this CR, returns
// the endpoint list, and ApplyChanges() updates (or creates) the CR.
type ExternalDNSProvider struct {
	provider.BaseProvider

	dryRun      bool
	crdClient   *rest.RESTClient
	scheme      *runtime.Scheme
	cfg         *ExternalDNSProviderConfig
	crdResource string
}

type ExternalDNSInterface interface {
	GetCurrentDNSEndpointCr(ctx context.Context) (*endpoint.DNSEndpoint, error)
}

var (
	Cp        = invokeGetCurrentDNSEndpointCr
	readFile  = os.ReadFile
	crdClient = source.NewCRDClientForAPIVersionKindWithConfig
	makeReq   = makeRequest
)

func invokeGetCurrentDNSEndpointCr(ctx context.Context, e ExternalDNSInterface) (*endpoint.DNSEndpoint, error) {
	return e.GetCurrentDNSEndpointCr(ctx)
}

// getCurrentDNSEndpointCr fetches the current DNSEndpoint CR in the target cluster, without error handling
func (p *ExternalDNSProvider) GetCurrentDNSEndpointCr(ctx context.Context) (*endpoint.DNSEndpoint, error) {
	epCr := &endpoint.DNSEndpoint{}

	err := p.crdClient.Get().
		Namespace(p.cfg.Namespace).
		Name(p.cfg.CrName).
		Resource(p.crdResource).
		VersionedParams(&metav1.GetOptions{}, runtime.NewParameterCodec(p.scheme)).
		Do(ctx).
		Into(epCr)
	return epCr, err
}

func makeRequest(ctx context.Context, p *ExternalDNSProvider, epCr *endpoint.DNSEndpoint, req *rest.Request) error {
	err := req.
		Namespace(p.cfg.Namespace).
		Name(epCr.Name).
		Resource(p.crdResource).
		VersionedParams(&metav1.UpdateOptions{}, runtime.NewParameterCodec(p.scheme)).
		Body(epCr).
		Do(ctx).
		Error()
	return err
}

// Records returns the list of endpoints
func (p *ExternalDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	epCr, err := Cp(ctx, p)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return []*endpoint.Endpoint{}, nil
		}
		return nil, err
	}

	return epCr.Spec.Endpoints, nil
}

// ApplyChanges simply passes the request to the DNSEndpoint CR in the target cluster, or creates it
// create record - record should not exist
// update/delete record - record should exist
// create/update/delete lists should not have overlapping records
func (p *ExternalDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	create := false
	//epCr, err := p.getCurrentDnsEndpointCr(ctx)
	epCr, err := Cp(ctx, p)
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return fmt.Errorf("ApplyChanges: failed to get DNSEndpoint CR from target cluster: %s", err)
		}
		create = true
		epCr = &endpoint.DNSEndpoint{
			ObjectMeta: metav1.ObjectMeta{
				Name:      p.cfg.CrName,
				Namespace: p.cfg.Namespace,
			},
			Spec: endpoint.DNSEndpointSpec{
				Endpoints: []*endpoint.Endpoint{},
			},
		}
	}

	epCr.Spec.Endpoints = append(epCr.Spec.Endpoints, changes.Create...)

	toBeRemovedEpsSet := map[string]struct{}{}
	for _, removeEp := range changes.UpdateOld {
		toBeRemovedEpsSet[removeEp.String()] = struct{}{}
	}
	for _, removeEp := range changes.Delete {
		toBeRemovedEpsSet[removeEp.String()] = struct{}{}
	}

	keepEndpoints := make([]*endpoint.Endpoint, 0)
	for _, existEp := range epCr.Spec.Endpoints {
		existEpStr := existEp.String()
		_, remove := toBeRemovedEpsSet[existEpStr]
		if !remove {
			keepEndpoints = append(keepEndpoints, existEp)
		}
	}
	epCr.Spec.Endpoints = keepEndpoints

	epCr.Spec.Endpoints = append(epCr.Spec.Endpoints, changes.UpdateNew...)

	if p.dryRun {
		log.Infof("DryRun: ApplyChanges: Created/Updated DNSEndpoint CR in target cluster")
		return nil
	}

	req := p.crdClient.Post()
	if !create {
		req = p.crdClient.Put()
	}
	/*err = req.
	Namespace(p.cfg.Namespace).
	Name(epCr.Name).
	Resource(p.crdResource).
	VersionedParams(&metav1.UpdateOptions{}, runtime.NewParameterCodec(p.scheme)).
	Body(epCr).
	Do(ctx).
	Error()*/
	err = makeReq(ctx, p, epCr, req)
	if err != nil {
		return fmt.Errorf("ApplyChanges: failed to create/update DNSEndpoint CR in target cluster: %s", err)
	}
	log.Infof("ApplyChanges: Created/Updated DNSEndpoint CR in target cluster")

	return nil
}

// NewExternalDNS returns ExternalDNSProvider DNS provider interface implementation
func NewExternalDNS(cfg *ExternalDNSProviderConfig, dryRun bool) (*ExternalDNSProvider, error) {
	if cfg.NamespacePath != "" {
		nsB, err := readFile(cfg.NamespacePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read namespace from file %s: %s", cfg.NamespacePath, err)
		}
		cfg.Namespace = string(nsB)
	}
	if cfg.CaCrtPath != "" {
		caCrtB, err := readFile(cfg.CaCrtPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA cert from file %s: %s", cfg.CaCrtPath, err)
		}
		cfg.CaCrt = string(caCrtB)
	}
	if cfg.TokenPath != "" {
		tokenB, err := readFile(cfg.TokenPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read token from file %s: %s", cfg.TokenPath, err)
		}
		cfg.Token = string(tokenB)
	}

	restConfig := &rest.Config{
		Host: "https://" + net.JoinHostPort(cfg.KubernetesHost, cfg.KubernetesPort),
		TLSClientConfig: rest.TLSClientConfig{
			CAData: []byte(cfg.CaCrt),
		},
		BearerToken: cfg.Token,
	}
	k8sClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s client: %s", err)
	}
	crdClient, scheme, err := crdClient(k8sClient, restConfig, cfg.CrdAPIVersion, cfg.CrdKind)
	if err != nil {
		return nil, fmt.Errorf("failed to get CRD k8s client: %s", err)
	}
	provider := &ExternalDNSProvider{
		dryRun:      dryRun,
		crdResource: strings.ToLower(cfg.CrdKind) + "s",
		crdClient:   crdClient,
		scheme:      scheme,
		cfg:         cfg,
	}

	return provider, nil
}
