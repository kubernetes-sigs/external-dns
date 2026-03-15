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

// Package factory provides provider selection and construction from configuration.
package factory

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

// SelectProvider constructs the DNS provider specified by cfg.Provider and
// wraps it in a cache when cfg.ProviderCacheTime is set.
func SelectProvider(
	ctx context.Context,
	cfg *externaldns.Config,
	domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	p, err := buildProvider(ctx, cfg, domainFilter)
	if err != nil {
		return nil, err
	}
	if p != nil && cfg.ProviderCacheTime > 0 {
		return provider.NewCachedProvider(p, cfg.ProviderCacheTime), nil
	}
	return p, nil
}

func buildProvider(
	ctx context.Context,
	cfg *externaldns.Config,
	domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	zoneNameFilter := endpoint.NewDomainFilter(cfg.ZoneNameFilter)
	zoneIDFilter := provider.NewZoneIDFilter(cfg.ZoneIDFilter)

	switch cfg.Provider {
	case "akamai":
		return akamai.NewAkamaiProvider(
			akamai.AkamaiConfig{
				DomainFilter:          domainFilter,
				ZoneIDFilter:          zoneIDFilter,
				ServiceConsumerDomain: cfg.AkamaiServiceConsumerDomain,
				ClientToken:           cfg.AkamaiClientToken,
				ClientSecret:          cfg.AkamaiClientSecret,
				AccessToken:           cfg.AkamaiAccessToken,
				EdgercPath:            cfg.AkamaiEdgercPath,
				EdgercSection:         cfg.AkamaiEdgercSection,
				DryRun:                cfg.DryRun,
			}, nil)
	case "alibabacloud":
		return alibabacloud.NewAlibabaCloudProvider(cfg.AlibabaCloudConfigFile, domainFilter, zoneIDFilter, cfg.AlibabaCloudZoneType, cfg.DryRun)
	case "aws":
		configs := aws.CreateV2Configs(cfg)
		clients := make(map[string]aws.Route53API, len(configs))
		for profile, config := range configs {
			clients[profile] = route53.NewFromConfig(config)
		}
		return aws.NewAWSProvider(
			aws.AWSConfig{
				DomainFilter:          domainFilter,
				ZoneIDFilter:          zoneIDFilter,
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
	case "aws-sd":
		// Only compatible registries can be used with AWS Cloud Map.
		if cfg.Registry != "noop" && cfg.Registry != "aws-sd" {
			log.Infof("Registry %q cannot be used with AWS Cloud Map. Switching to \"aws-sd\".", cfg.Registry)
			cfg.Registry = "aws-sd"
		}
		return awssd.NewAWSSDProvider(domainFilter, cfg.AWSZoneType, cfg.DryRun, cfg.AWSSDServiceCleanup, cfg.TXTOwnerID, cfg.AWSSDCreateTag, sd.NewFromConfig(aws.CreateDefaultV2Config(cfg)))
	case "azure-dns", "azure":
		return azure.NewAzureProvider(cfg.AzureConfigFile, domainFilter, zoneNameFilter, zoneIDFilter, cfg.AzureSubscriptionID, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.AzureActiveDirectoryAuthorityHost, cfg.AzureZonesCacheDuration, cfg.AzureMaxRetriesCount, cfg.DryRun)
	case "azure-private-dns":
		return azure.NewAzurePrivateDNSProvider(cfg.AzureConfigFile, domainFilter, zoneNameFilter, zoneIDFilter, cfg.AzureSubscriptionID, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.AzureActiveDirectoryAuthorityHost, cfg.AzureZonesCacheDuration, cfg.AzureMaxRetriesCount, cfg.DryRun)
	case "civo":
		return civo.NewCivoProvider(domainFilter, cfg.DryRun)
	case "cloudflare":
		return cloudflare.NewCloudFlareProvider(
			domainFilter,
			zoneIDFilter,
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
				PerPage:             cfg.CloudflareDNSRecordsPerPage,
				Comment:             cfg.CloudflareDNSRecordsComment,
				BatchChangeSize:     cfg.BatchChangeSize,
				BatchChangeInterval: cfg.BatchChangeInterval,
			})
	case "coredns", "skydns":
		return coredns.NewCoreDNSProvider(domainFilter, cfg.CoreDNSPrefix, cfg.TXTOwnerID, cfg.CoreDNSStrictlyOwned, cfg.DryRun)
	case "digitalocean":
		return digitalocean.NewDigitalOceanProvider(ctx, domainFilter, cfg.DryRun, cfg.DigitalOceanAPIPageSize)
	case "dnsimple":
		return dnsimple.NewDnsimpleProvider(domainFilter, zoneIDFilter, cfg.DryRun)
	case "exoscale":
		return exoscale.NewExoscaleProvider(
			cfg.ExoscaleAPIEnvironment,
			cfg.ExoscaleAPIZone,
			cfg.ExoscaleAPIKey,
			cfg.ExoscaleAPISecret,
			cfg.DryRun,
			exoscale.ExoscaleWithDomain(domainFilter),
			exoscale.ExoscaleWithLogging(),
		)
	case "gandi":
		return gandi.NewGandiProvider(domainFilter, cfg.DryRun)
	case "godaddy":
		return godaddy.NewGoDaddyProvider(ctx, domainFilter, cfg.GoDaddyTTL, cfg.GoDaddyAPIKey, cfg.GoDaddySecretKey, cfg.GoDaddyOTE, cfg.DryRun)
	case "google":
		return google.NewGoogleProvider(ctx, cfg.GoogleProject, domainFilter, zoneIDFilter, cfg.GoogleBatchChangeSize, cfg.GoogleBatchChangeInterval, cfg.GoogleZoneVisibility, cfg.DryRun)
	case "inmemory":
		return inmemory.NewInMemoryProvider(inmemory.InMemoryInitZones(cfg.InMemoryZones), inmemory.InMemoryWithDomain(domainFilter), inmemory.InMemoryWithLogging()), nil
	case "linode":
		return linode.NewLinodeProvider(domainFilter, cfg.DryRun)
	case "ns1":
		return ns1.NewNS1Provider(
			ns1.NS1Config{
				DomainFilter:  domainFilter,
				ZoneIDFilter:  zoneIDFilter,
				NS1Endpoint:   cfg.NS1Endpoint,
				NS1IgnoreSSL:  cfg.NS1IgnoreSSL,
				DryRun:        cfg.DryRun,
				MinTTLSeconds: cfg.NS1MinTTLSeconds,
			},
		)
	case "oci":
		var config *oci.OCIConfig
		var err error
		if cfg.OCIAuthInstancePrincipal {
			if cfg.OCICompartmentOCID == "" {
				return nil, fmt.Errorf("instance principal authentication requested, but no compartment OCID provided")
			}
			config = &oci.OCIConfig{Auth: oci.OCIAuthConfig{UseInstancePrincipal: true}, CompartmentID: cfg.OCICompartmentOCID}
		} else {
			if config, err = oci.LoadOCIConfig(cfg.OCIConfigFile); err != nil {
				return nil, err
			}
		}
		config.ZoneCacheDuration = cfg.OCIZoneCacheDuration
		return oci.NewOCIProvider(*config, domainFilter, zoneIDFilter, cfg.OCIZoneScope, cfg.DryRun)
	case "ovh":
		return ovh.NewOVHProvider(domainFilter, cfg.OVHEndpoint, cfg.OVHApiRateLimit, cfg.OVHEnableCNAMERelative, cfg.DryRun)
	case "pdns":
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
	case "pihole":
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
	case "plural":
		return plural.NewPluralProvider(cfg.PluralCluster, cfg.PluralProvider)
	case "rfc2136":
		return rfc2136.NewRfc2136Provider(
			cfg.RFC2136Host,
			cfg.RFC2136Port,
			cfg.RFC2136Zone,
			cfg.RFC2136Insecure,
			cfg.RFC2136TSIGKeyName,
			cfg.RFC2136TSIGSecret,
			cfg.RFC2136TSIGSecretAlg,
			cfg.RFC2136TAXFR,
			domainFilter,
			cfg.DryRun,
			cfg.RFC2136MinTTL,
			cfg.RFC2136CreatePTR,
			cfg.RFC2136GSSTSIG,
			cfg.RFC2136KerberosUsername,
			cfg.RFC2136KerberosPassword,
			cfg.RFC2136KerberosRealm,
			cfg.RFC2136BatchChangeSize,
			rfc2136.TLSConfig{
				UseTLS:                cfg.RFC2136UseTLS,
				SkipTLSVerify:         cfg.RFC2136SkipTLSVerify,
				CAFilePath:            cfg.TLSCA,
				ClientCertFilePath:    cfg.TLSClientCert,
				ClientCertKeyFilePath: cfg.TLSClientCertKey,
			},
			cfg.RFC2136LoadBalancingStrategy,
			nil,
		)
	case "scaleway":
		return scaleway.NewScalewayProvider(domainFilter, cfg.DryRun)
	case "transip":
		return transip.NewTransIPProvider(cfg.TransIPAccountName, cfg.TransIPPrivateKeyFile, domainFilter, cfg.DryRun)
	case "webhook":
		return webhook.NewWebhookProvider(cfg.WebhookProviderURL)
	default:
		return nil, fmt.Errorf("unknown dns provider: %s", cfg.Provider)
	}
}
