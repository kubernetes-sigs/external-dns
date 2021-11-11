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

package provider

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/provider/akamai"
	"sigs.k8s.io/external-dns/provider/alibabacloud"
	"sigs.k8s.io/external-dns/provider/aws"
	"sigs.k8s.io/external-dns/provider/awssd"
	"sigs.k8s.io/external-dns/provider/azure"
	"sigs.k8s.io/external-dns/provider/bluecat"
	"sigs.k8s.io/external-dns/provider/cloudflare"
	"sigs.k8s.io/external-dns/provider/coredns"
	"sigs.k8s.io/external-dns/provider/designate"
	"sigs.k8s.io/external-dns/provider/digitalocean"
	"sigs.k8s.io/external-dns/provider/dnsimple"
	"sigs.k8s.io/external-dns/provider/dyn"
	"sigs.k8s.io/external-dns/provider/exoscale"
	"sigs.k8s.io/external-dns/provider/gandi"
	"sigs.k8s.io/external-dns/provider/godaddy"
	"sigs.k8s.io/external-dns/provider/google"
	"sigs.k8s.io/external-dns/provider/hetzner"
	"sigs.k8s.io/external-dns/provider/infoblox"
	"sigs.k8s.io/external-dns/provider/inmemory"
	"sigs.k8s.io/external-dns/provider/linode"
	"sigs.k8s.io/external-dns/provider/ns1"
	"sigs.k8s.io/external-dns/provider/oci"
	"sigs.k8s.io/external-dns/provider/ovh"
	"sigs.k8s.io/external-dns/provider/pdns"
	"sigs.k8s.io/external-dns/provider/rcode0"
	"sigs.k8s.io/external-dns/provider/rdns"
	"sigs.k8s.io/external-dns/provider/rfc2136"
	"sigs.k8s.io/external-dns/provider/scaleway"
	"sigs.k8s.io/external-dns/provider/transip"
	"sigs.k8s.io/external-dns/provider/ultradns"
	"sigs.k8s.io/external-dns/provider/vinyldns"
	"sigs.k8s.io/external-dns/provider/vultr"
)

// NewBuilder returns list of all providers constructors
func NewBuilder(ctx context.Context, domainFilter endpoint.DomainFilter, cfg *externaldns.Config) Builder {
	zoneNameFilter := endpoint.NewDomainFilter(cfg.ZoneNameFilter)
	zoneIDFilter := NewZoneIDFilter(cfg.ZoneIDFilter)
	zoneTypeFilter := NewZoneTypeFilter(cfg.AWSZoneType)
	zoneTagFilter := NewZoneTagFilter(cfg.AWSZoneTagFilter)

	return Builder{
		"akamai": func() (Provider, error) {
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
		},
		"alibabacloud": func() (Provider, error) {
			return alibabacloud.NewAlibabaCloudProvider(cfg.AlibabaCloudConfigFile, domainFilter, zoneIDFilter, cfg.AlibabaCloudZoneType, cfg.DryRun)
		},
		"aws": func() (Provider, error) {
			return aws.NewAWSProvider(
				aws.AWSConfig{
					DomainFilter:         domainFilter,
					ZoneIDFilter:         zoneIDFilter,
					ZoneTypeFilter:       zoneTypeFilter,
					ZoneTagFilter:        zoneTagFilter,
					BatchChangeSize:      cfg.AWSBatchChangeSize,
					BatchChangeInterval:  cfg.AWSBatchChangeInterval,
					EvaluateTargetHealth: cfg.AWSEvaluateTargetHealth,
					AssumeRole:           cfg.AWSAssumeRole,
					APIRetries:           cfg.AWSAPIRetries,
					PreferCNAME:          cfg.AWSPreferCNAME,
					DryRun:               cfg.DryRun,
					ZoneCacheDuration:    cfg.AWSZoneCacheDuration,
				},
			)
		},
		"aws-sd": func() (Provider, error) {
			// Check that only compatible Registry is used with AWS-SD
			if cfg.Registry != "noop" && cfg.Registry != "aws-sd" {
				log.Infof("Registry \"%s\" cannot be used with AWS Cloud Map. Switching to \"aws-sd\".", cfg.Registry)
				cfg.Registry = "aws-sd"
			}
			return awssd.NewAWSSDProvider(domainFilter, cfg.AWSZoneType, cfg.AWSAssumeRole, cfg.DryRun)
		},
		"azure-dns": func() (Provider, error) {
			return azure.NewAzureProvider(cfg.AzureConfigFile, domainFilter, zoneNameFilter, zoneIDFilter, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.DryRun)
		},
		"azure": func() (Provider, error) {
			return azure.NewAzureProvider(cfg.AzureConfigFile, domainFilter, zoneNameFilter, zoneIDFilter, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.DryRun)
		},
		"azure-private-dns": func() (Provider, error) {
			return azure.NewAzurePrivateDNSProvider(cfg.AzureConfigFile, domainFilter, zoneIDFilter, cfg.AzureResourceGroup, cfg.AzureUserAssignedIdentityClientID, cfg.DryRun)
		},
		"bluecat": func() (Provider, error) {
			return bluecat.NewBluecatProvider(cfg.BluecatConfigFile, domainFilter, zoneIDFilter, cfg.DryRun)
		},
		"vinyldns": func() (Provider, error) {
			return vinyldns.NewVinylDNSProvider(domainFilter, zoneIDFilter, cfg.DryRun)
		},
		"vultr": func() (Provider, error) {
			return vultr.NewVultrProvider(ctx, domainFilter, cfg.DryRun)
		},
		"ultradns": func() (Provider, error) {
			return ultradns.NewUltraDNSProvider(domainFilter, cfg.DryRun)
		},
		"cloudflare": func() (Provider, error) {
			return cloudflare.NewCloudFlareProvider(domainFilter, zoneIDFilter, cfg.CloudflareZonesPerPage, cfg.CloudflareProxied, cfg.DryRun)
		},
		"rcodezero": func() (Provider, error) {
			return rcode0.NewRcodeZeroProvider(domainFilter, cfg.DryRun, cfg.RcodezeroTXTEncrypt)
		},
		"google": func() (Provider, error) {
			return google.NewGoogleProvider(ctx, cfg.GoogleProject, domainFilter, zoneIDFilter, cfg.GoogleBatchChangeSize, cfg.GoogleBatchChangeInterval, cfg.GoogleZoneVisibility, cfg.DryRun)
		},
		"digitalocean": func() (Provider, error) {
			return digitalocean.NewDigitalOceanProvider(ctx, domainFilter, cfg.DryRun, cfg.DigitalOceanAPIPageSize)
		},
		"hetzner": func() (Provider, error) {
			return hetzner.NewHetznerProvider(ctx, domainFilter, cfg.DryRun)
		},
		"ovh": func() (Provider, error) {
			return ovh.NewOVHProvider(ctx, domainFilter, cfg.OVHEndpoint, cfg.OVHApiRateLimit, cfg.DryRun)
		},
		"linode": func() (Provider, error) {
			return linode.NewLinodeProvider(domainFilter, cfg.DryRun, externaldns.Version)
		},
		"dnsimple": func() (Provider, error) {
			return dnsimple.NewDnsimpleProvider(domainFilter, zoneIDFilter, cfg.DryRun)
		},
		"infoblox": func() (Provider, error) {
			return infoblox.NewInfobloxProvider(
				infoblox.InfobloxConfig{
					DomainFilter: domainFilter,
					ZoneIDFilter: zoneIDFilter,
					Host:         cfg.InfobloxGridHost,
					Port:         cfg.InfobloxWapiPort,
					Username:     cfg.InfobloxWapiUsername,
					Password:     cfg.InfobloxWapiPassword,
					Version:      cfg.InfobloxWapiVersion,
					SSLVerify:    cfg.InfobloxSSLVerify,
					View:         cfg.InfobloxView,
					MaxResults:   cfg.InfobloxMaxResults,
					DryRun:       cfg.DryRun,
					FQDNRexEx:    cfg.InfobloxFQDNRegEx,
					CreatePTR:    cfg.InfobloxCreatePTR,
				},
			)
		},
		"dyn": func() (Provider, error) {
			return dyn.NewDynProvider(
				dyn.DynConfig{
					DomainFilter:  domainFilter,
					ZoneIDFilter:  zoneIDFilter,
					DryRun:        cfg.DryRun,
					CustomerName:  cfg.DynCustomerName,
					Username:      cfg.DynUsername,
					Password:      cfg.DynPassword,
					MinTTLSeconds: cfg.DynMinTTLSeconds,
					AppVersion:    externaldns.Version,
				},
			)
		},
		"coredns": func() (Provider, error) {
			return coredns.NewCoreDNSProvider(domainFilter, cfg.CoreDNSPrefix, cfg.DryRun)
		},
		"skydns": func() (Provider, error) {
			return coredns.NewCoreDNSProvider(domainFilter, cfg.CoreDNSPrefix, cfg.DryRun)
		},
		"rdns": func() (Provider, error) {
			return rdns.NewRDNSProvider(
				rdns.RDNSConfig{
					DomainFilter: domainFilter,
					DryRun:       cfg.DryRun,
				},
			)
		},
		"exoscale": func() (Provider, error) {
			return exoscale.NewExoscaleProvider(cfg.ExoscaleEndpoint, cfg.ExoscaleAPIKey, cfg.ExoscaleAPISecret, cfg.DryRun, exoscale.ExoscaleWithDomain(domainFilter), exoscale.ExoscaleWithLogging()), nil
		},
		"inmemory": func() (Provider, error) {
			return inmemory.NewInMemoryProvider(inmemory.InMemoryInitZones(cfg.InMemoryZones), inmemory.InMemoryWithDomain(domainFilter), inmemory.InMemoryWithLogging()), nil
		},
		"designate": func() (Provider, error) {
			return designate.NewDesignateProvider(domainFilter, cfg.DryRun)
		},
		"pdns": func() (Provider, error) {
			return pdns.NewPDNSProvider(
				ctx,
				pdns.PDNSConfig{
					DomainFilter: domainFilter,
					DryRun:       cfg.DryRun,
					Server:       cfg.PDNSServer,
					APIKey:       cfg.PDNSAPIKey,
					TLSConfig: pdns.TLSConfig{
						TLSEnabled:            cfg.PDNSTLSEnabled,
						CAFilePath:            cfg.TLSCA,
						ClientCertFilePath:    cfg.TLSClientCert,
						ClientCertKeyFilePath: cfg.TLSClientCertKey,
					},
				},
			)
		},
		"oci": func() (Provider, error) {
			var config *oci.OCIConfig
			config, err := oci.LoadOCIConfig(cfg.OCIConfigFile)
			if err == nil {
				return oci.NewOCIProvider(*config, domainFilter, zoneIDFilter, cfg.DryRun)
			}
			return nil, fmt.Errorf("oci config: %w", err)
		},
		"rfc2136": func() (Provider, error) {
			return rfc2136.NewRfc2136Provider(cfg.RFC2136Host, cfg.RFC2136Port, cfg.RFC2136Zone, cfg.RFC2136Insecure, cfg.RFC2136TSIGKeyName, cfg.RFC2136TSIGSecret, cfg.RFC2136TSIGSecretAlg, cfg.RFC2136TAXFR, domainFilter, cfg.DryRun, cfg.RFC2136MinTTL, cfg.RFC2136GSSTSIG, cfg.RFC2136KerberosUsername, cfg.RFC2136KerberosPassword, cfg.RFC2136KerberosRealm, cfg.RFC2136BatchChangeSize, nil)
		},
		"ns1": func() (Provider, error) {
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
		},
		"transip": func() (Provider, error) {
			return transip.NewTransIPProvider(cfg.TransIPAccountName, cfg.TransIPPrivateKeyFile, domainFilter, cfg.DryRun)
		},
		"scaleway": func() (Provider, error) {
			return scaleway.NewScalewayProvider(ctx, domainFilter, cfg.DryRun)
		},
		"godaddy": func() (Provider, error) {
			return godaddy.NewGoDaddyProvider(ctx, domainFilter, cfg.GoDaddyTTL, cfg.GoDaddyAPIKey, cfg.GoDaddySecretKey, cfg.GoDaddyOTE, cfg.DryRun)
		},
		"gandi": func() (Provider, error) {
			return gandi.NewGandiProvider(ctx, domainFilter, cfg.DryRun)
		},
	}
}

// Builder as single point for list of providers
type Builder map[string]func() (Provider, error)

// Names returns all providers codes
func (pb Builder) Names() []string {
	names := make([]string, 0, len(pb))
	for name := range pb {
		names = append(names, name)
	}
	return names
}

// Build returns concrete provider instance or error
func (pb Builder) Build(provider string) (Provider, error) {
	builder, ok := pb[provider]
	if !ok {
		return nil, fmt.Errorf("unknown dns provider: %s", provider)
	}
	return builder()
}
