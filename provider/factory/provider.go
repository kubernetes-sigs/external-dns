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

package factory

import (
	"context"
	"fmt"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/akamai"
	"sigs.k8s.io/external-dns/provider/alibabacloud"
	"sigs.k8s.io/external-dns/provider/aws"
	"sigs.k8s.io/external-dns/provider/awssd"
	"sigs.k8s.io/external-dns/provider/azure"
	"sigs.k8s.io/external-dns/provider/civo"
	"sigs.k8s.io/external-dns/provider/cloudflare"
	"sigs.k8s.io/external-dns/provider/coredns"
	"sigs.k8s.io/external-dns/provider/digitalocean"
	"sigs.k8s.io/external-dns/provider/dnsimple"
	"sigs.k8s.io/external-dns/provider/exoscale"
	"sigs.k8s.io/external-dns/provider/gandi"
	"sigs.k8s.io/external-dns/provider/godaddy"
	"sigs.k8s.io/external-dns/provider/google"
	"sigs.k8s.io/external-dns/provider/inmemory"
	"sigs.k8s.io/external-dns/provider/linode"
	"sigs.k8s.io/external-dns/provider/ns1"
	"sigs.k8s.io/external-dns/provider/oci"
	"sigs.k8s.io/external-dns/provider/ovh"
	"sigs.k8s.io/external-dns/provider/pdns"
	"sigs.k8s.io/external-dns/provider/pihole"
	"sigs.k8s.io/external-dns/provider/plural"
	"sigs.k8s.io/external-dns/provider/rfc2136"
	"sigs.k8s.io/external-dns/provider/scaleway"
	"sigs.k8s.io/external-dns/provider/transip"
	"sigs.k8s.io/external-dns/provider/webhook"
)

// ProviderConstructor is a function that creates a provider from configuration.
type ProviderConstructor func(
	ctx context.Context,
	cfg *externaldns.Config,
	domainFilter *endpoint.DomainFilter,
) (provider.Provider, error)

// Select creates a provider based on the given configuration.
func Select(
	ctx context.Context,
	cfg *externaldns.Config,
	domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	constructor, ok := providers(cfg.Provider)
	if !ok {
		return nil, fmt.Errorf("unknown dns provider: %s", cfg.Provider)
	}
	p, err := constructor(ctx, cfg, domainFilter)
	if err != nil {
		return nil, err
	}
	if p != nil && cfg.ProviderCacheTime > 0 {
		return provider.NewCachedProvider(p, cfg.ProviderCacheTime), nil
	}
	return p, nil
}

// providers looks up the constructor for the named provider.
func providers(selector string) (ProviderConstructor, bool) {
	m := map[string]ProviderConstructor{
		externaldns.ProviderAkamai:       akamai.New,
		externaldns.ProviderAlibabaCloud: alibabacloud.New,
		externaldns.ProviderAWS:          aws.New,
		externaldns.ProviderAWSSD:        awssd.New,
		externaldns.ProviderAzure:        azure.New,
		externaldns.ProviderAzureDNS:     azure.New,
		externaldns.ProviderAzurePrivate: azure.NewPrivate,
		externaldns.ProviderCivo:         civo.New,
		externaldns.ProviderCloudflare:   cloudflare.New,
		externaldns.ProviderCoreDNS:      coredns.New,
		externaldns.ProviderSkyDNS:       coredns.New,
		externaldns.ProviderDigitalOcean: digitalocean.New,
		externaldns.ProviderDNSimple:     dnsimple.New,
		externaldns.ProviderExoscale:     exoscale.New,
		externaldns.ProviderGandi:        gandi.New,
		externaldns.ProviderGoDaddy:      godaddy.New,
		externaldns.ProviderGoogle:       google.New,
		externaldns.ProviderInMemory:     inmemory.New,
		externaldns.ProviderLinode:       linode.New,
		externaldns.ProviderNS1:          ns1.New,
		externaldns.ProviderOCI:          oci.New,
		externaldns.ProviderOVH:          ovh.New,
		externaldns.ProviderPDNS:         pdns.New,
		externaldns.ProviderPihole:       pihole.New,
		externaldns.ProviderPlural:       plural.New,
		externaldns.ProviderRFC2136:      rfc2136.New,
		externaldns.ProviderScaleway:     scaleway.New,
		externaldns.ProviderTransip:      transip.New,
		externaldns.ProviderWebhook:      webhook.New,
	}
	c, ok := m[selector]
	return c, ok
}
