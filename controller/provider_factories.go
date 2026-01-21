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

package controller

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	sd "github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	log "github.com/sirupsen/logrus"

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

type providerFactory func(context.Context, *externaldns.Config, *endpoint.DomainFilter) (provider.Provider, error)

var providerFactories = map[string]providerFactory{
	"akamai": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return akamai.NewAkamaiProvider(
			akamai.AkamaiConfig{
				DomainFilter:          domainFilter,
				ZoneIDFilter:          provider.NewZoneIDFilter(cfg.ZoneIDFilter),
				ServiceConsumerDomain: cfg.AkamaiServiceConsumerDomain,
				ClientToken:           cfg.AkamaiClientToken,
				ClientSecret:          cfg.AkamaiClientSecret,
				AccessToken:           cfg.AkamaiAccessToken,
				EdgercPath:            cfg.AkamaiEdgercPath,
				EdgercSection:         cfg.AkamaiEdgercSection,
				DryRun:                cfg.DryRun,
			}, nil)
	},
	"alibabacloud": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return alibabacloud.NewAlibabaCloudProvider(cfg.AlibabaCloudConfigFile, domainFilter, provider.NewZoneIDFilter(cfg.ZoneIDFilter), cfg.AlibabaCloudZoneType, cfg.DryRun)
	},
	"aws": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		configs := aws.CreateV2Configs(cfg)
		clients := make(map[string]aws.Route53API, len(configs))
		for profile, config := range configs {
			clients[profile] = route53.NewFromConfig(config)
		}

		return aws.NewAWSProvider(
			aws.AWSConfig{
				DomainFilter:          domainFilter,
				ZoneIDFilter:          provider.NewZoneIDFilter(cfg.ZoneIDFilter),
				ZoneTypeFilter:        provider.NewZoneTypeFilter(cfg.AWSZoneType),
				ZoneTagFilter:         provider.NewZoneTagFilter(cfg.AWSZoneTagFilter),
				ZoneMatchParent:       cfg.AWSZoneMatchParent,
				BatchChangeSize:       cfg.AWSBatchChangeSize,
				BatchChangeSizeBytes:  cfg.AWSBatchChangeSizeBytes,
				BatchChangeSizeValues: cfg.AWSBatchChangeSizeValues,
				BatchChangeInterval:   cfg.AWSBatchChangeInterval,
				EvaluateTargetHealth:  cfg.AWSEvaluateTargetHealth,
				PreferCNAME:           cfg.AWSPreferCNAME,
				DryRun:                cfg.DryRun,
				ZoneCacheDuration:     cfg.AWSZoneCacheDuration,
			},
			clients,
		)
	},
	"aws-sd": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		// Check that only compatible Registry is used with AWS-SD
		if cfg.Registry != "noop" && cfg.Registry != "aws-sd" {
			log.Infof("Registry \"%s\" cannot be used with AWS Cloud Map. Switching to \"aws-sd\".", cfg.Registry)
			cfg.Registry = "aws-sd"
		}
		return awssd.NewAWSSDProvider(domainFilter, cfg.AWSZoneType, cfg.DryRun, cfg.AWSSDServiceCleanup, cfg.TXTOwnerID, cfg.AWSSDCreateTag, sd.NewFromConfig(aws.CreateDefaultV2Config(cfg)))
	},
	"azure-dns": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return azure.NewAzureProvider(cfg.AzureConfigFile, domainFilter, endpoint.NewDomainFilter(cfg.ZoneNameFilter), provider.NewZoneIDFilter(cfg.ZoneIDFilter), cfg.AzureSubscriptionID, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.AzureActiveDirectoryAuthorityHost, cfg.AzureZonesCacheDuration, cfg.AzureMaxRetriesCount, cfg.DryRun)
	},
	"azure": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return azure.NewAzureProvider(cfg.AzureConfigFile, domainFilter, endpoint.NewDomainFilter(cfg.ZoneNameFilter), provider.NewZoneIDFilter(cfg.ZoneIDFilter), cfg.AzureSubscriptionID, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.AzureActiveDirectoryAuthorityHost, cfg.AzureZonesCacheDuration, cfg.AzureMaxRetriesCount, cfg.DryRun)
	},
	"azure-private-dns": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return azure.NewAzurePrivateDNSProvider(cfg.AzureConfigFile, domainFilter, endpoint.NewDomainFilter(cfg.ZoneNameFilter), provider.NewZoneIDFilter(cfg.ZoneIDFilter), cfg.AzureSubscriptionID, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.AzureActiveDirectoryAuthorityHost, cfg.AzureZonesCacheDuration, cfg.AzureMaxRetriesCount, cfg.DryRun)
	},
	"civo": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return civo.NewCivoProvider(domainFilter, cfg.DryRun)
	},
	"cloudflare": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return cloudflare.NewCloudFlareProvider(
			domainFilter,
			provider.NewZoneIDFilter(cfg.ZoneIDFilter),
			cfg.CloudflareProxied,
			cfg.DryRun,
			cloudflare.RegionalServicesConfig{
				Enabled:   cfg.CloudflareRegionalServices,
				RegionKey: cfg.CloudflareRegionKey,
			},
			cloudflare.CustomHostnamesConfig{
				Enabled:              cfg.CloudflareCustomHostnames,
				MinTLSVersion:        cfg.CloudflareCustomHostnamesMinTLSVersion,
				CertificateAuthority: cfg.CloudflareCustomHostnamesCertificateAuthority,
			},
			cloudflare.DNSRecordsConfig{
				PerPage: cfg.CloudflareDNSRecordsPerPage,
				Comment: cfg.CloudflareDNSRecordsComment,
			})
	},
	"google": func(ctx context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return google.NewGoogleProvider(ctx, cfg.GoogleProject, domainFilter, provider.NewZoneIDFilter(cfg.ZoneIDFilter), cfg.GoogleBatchChangeSize, cfg.GoogleBatchChangeInterval, cfg.GoogleZoneVisibility, cfg.DryRun)
	},
	"digitalocean": func(ctx context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return digitalocean.NewDigitalOceanProvider(ctx, domainFilter, cfg.DryRun, cfg.DigitalOceanAPIPageSize)
	},
	"ovh": func(ctx context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return ovh.NewOVHProvider(ctx, domainFilter, cfg.OVHEndpoint, cfg.OVHApiRateLimit, cfg.OVHEnableCNAMERelative, cfg.DryRun)
	},
	"linode": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return linode.NewLinodeProvider(domainFilter, cfg.DryRun)
	},
	"dnsimple": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return dnsimple.NewDnsimpleProvider(domainFilter, provider.NewZoneIDFilter(cfg.ZoneIDFilter), cfg.DryRun)
	},
	"coredns": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return coredns.NewCoreDNSProvider(domainFilter, cfg.CoreDNSPrefix, cfg.TXTOwnerID, cfg.CoreDNSStrictlyOwned, cfg.DryRun)
	},
	"skydns": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return coredns.NewCoreDNSProvider(domainFilter, cfg.CoreDNSPrefix, cfg.TXTOwnerID, cfg.CoreDNSStrictlyOwned, cfg.DryRun)
	},
	"exoscale": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return exoscale.NewExoscaleProvider(
			cfg.ExoscaleAPIEnvironment,
			cfg.ExoscaleAPIZone,
			cfg.ExoscaleAPIKey,
			cfg.ExoscaleAPISecret,
			cfg.DryRun,
			exoscale.ExoscaleWithDomain(domainFilter),
			exoscale.ExoscaleWithLogging(),
		)
	},
	"inmemory": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return inmemory.NewInMemoryProvider(inmemory.InMemoryInitZones(cfg.InMemoryZones), inmemory.InMemoryWithDomain(domainFilter), inmemory.InMemoryWithLogging()), nil
	},
	"pdns": func(ctx context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return pdns.NewPDNSProvider(
			ctx,
			pdns.PDNSConfig{
				DomainFilter: domainFilter,
				DryRun:       cfg.DryRun,
				Server:       cfg.PDNSServer,
				ServerID:     cfg.PDNSServerID,
				APIKey:       cfg.PDNSAPIKey,
				TLSConfig: pdns.TLSConfig{
					SkipTLSVerify:         cfg.PDNSSkipTLSVerify,
					CAFilePath:            cfg.TLSCA,
					ClientCertFilePath:    cfg.TLSClientCert,
					ClientCertKeyFilePath: cfg.TLSClientCertKey,
				},
			},
		)
	},
	"oci": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		var config *oci.OCIConfig
		var err error
		// if the instance-principals flag was set, and a compartment OCID was provided, then ignore the
		// OCI config file, and provide a config that uses instance principal authentication.
		if cfg.OCIAuthInstancePrincipal {
			if len(cfg.OCICompartmentOCID) == 0 {
				return nil, fmt.Errorf("instance principal authentication requested, but no compartment OCID provided")
			}
			authConfig := oci.OCIAuthConfig{UseInstancePrincipal: true}
			config = &oci.OCIConfig{Auth: authConfig, CompartmentID: cfg.OCICompartmentOCID}
		} else {
			if config, err = oci.LoadOCIConfig(cfg.OCIConfigFile); err != nil {
				return nil, err
			}
		}
		config.ZoneCacheDuration = cfg.OCIZoneCacheDuration
		return oci.NewOCIProvider(*config, domainFilter, provider.NewZoneIDFilter(cfg.ZoneIDFilter), cfg.OCIZoneScope, cfg.DryRun)
	},
	"rfc2136": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		tlsConfig := rfc2136.TLSConfig{
			UseTLS:                cfg.RFC2136UseTLS,
			SkipTLSVerify:         cfg.RFC2136SkipTLSVerify,
			CAFilePath:            cfg.TLSCA,
			ClientCertFilePath:    cfg.TLSClientCert,
			ClientCertKeyFilePath: cfg.TLSClientCertKey,
		}
		return rfc2136.NewRfc2136Provider(cfg.RFC2136Host, cfg.RFC2136Port, cfg.RFC2136Zone, cfg.RFC2136Insecure, cfg.RFC2136TSIGKeyName, cfg.RFC2136TSIGSecret, cfg.RFC2136TSIGSecretAlg, cfg.RFC2136TAXFR, domainFilter, cfg.DryRun, cfg.RFC2136MinTTL, cfg.RFC2136CreatePTR, cfg.RFC2136GSSTSIG, cfg.RFC2136KerberosUsername, cfg.RFC2136KerberosPassword, cfg.RFC2136KerberosRealm, cfg.RFC2136BatchChangeSize, tlsConfig, cfg.RFC2136LoadBalancingStrategy, nil)
	},
	"ns1": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return ns1.NewNS1Provider(
			ns1.NS1Config{
				DomainFilter:  domainFilter,
				ZoneIDFilter:  provider.NewZoneIDFilter(cfg.ZoneIDFilter),
				NS1Endpoint:   cfg.NS1Endpoint,
				NS1IgnoreSSL:  cfg.NS1IgnoreSSL,
				DryRun:        cfg.DryRun,
				MinTTLSeconds: cfg.NS1MinTTLSeconds,
			},
		)
	},
	"transip": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return transip.NewTransIPProvider(cfg.TransIPAccountName, cfg.TransIPPrivateKeyFile, domainFilter, cfg.DryRun)
	},
	"scaleway": func(ctx context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return scaleway.NewScalewayProvider(ctx, domainFilter, cfg.DryRun)
	},
	"godaddy": func(ctx context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return godaddy.NewGoDaddyProvider(ctx, domainFilter, cfg.GoDaddyTTL, cfg.GoDaddyAPIKey, cfg.GoDaddySecretKey, cfg.GoDaddyOTE, cfg.DryRun)
	},
	"gandi": func(ctx context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return gandi.NewGandiProvider(ctx, domainFilter, cfg.DryRun)
	},
	"pihole": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return pihole.NewPiholeProvider(
			pihole.PiholeConfig{
				Server:                cfg.PiholeServer,
				Password:              cfg.PiholePassword,
				TLSInsecureSkipVerify: cfg.PiholeTLSInsecureSkipVerify,
				DomainFilter:          domainFilter,
				DryRun:                cfg.DryRun,
				APIVersion:            cfg.PiholeApiVersion,
			},
		)
	},
	"plural": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return plural.NewPluralProvider(cfg.PluralCluster, cfg.PluralProvider)
	},
	"webhook": func(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
		return webhook.NewWebhookProvider(cfg.WebhookProviderURL)
	},
}
