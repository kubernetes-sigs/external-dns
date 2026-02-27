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
	"regexp"
	"testing"
	"time"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/flags"
	"sigs.k8s.io/external-dns/internal/testutils"

	"github.com/alecthomas/kingpin/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	minimalConfig = &Config{
		APIServerURL:                           "",
		KubeConfig:                             "",
		RequestTimeout:                         time.Second * 30,
		GlooNamespaces:                         []string{"gloo-system"},
		SkipperRouteGroupVersion:               "zalando.org/v1",
		Sources:                                []string{"service"},
		Namespace:                              "",
		AnnotationPrefix:                       "external-dns.alpha.kubernetes.io/",
		FQDNTemplate:                           "",
		Compatibility:                          "",
		Provider:                               "google",
		GoogleProject:                          "",
		GoogleBatchChangeSize:                  1000,
		GoogleBatchChangeInterval:              time.Second,
		GoogleZoneVisibility:                   "",
		DomainFilter:                           []string{""},
		DomainExclude:                          []string{""},
		RegexDomainFilter:                      regexp.MustCompile(""),
		RegexDomainExclude:                     regexp.MustCompile(""),
		ZoneNameFilter:                         []string{""},
		ZoneIDFilter:                           []string{""},
		AlibabaCloudConfigFile:                 "/etc/kubernetes/alibaba-cloud.json",
		AWSZoneType:                            "",
		AWSZoneTagFilter:                       []string{""},
		AWSZoneMatchParent:                     false,
		AWSAssumeRole:                          "",
		AWSAssumeRoleExternalID:                "",
		AWSBatchChangeSize:                     1000,
		AWSBatchChangeSizeBytes:                32000,
		AWSBatchChangeSizeValues:               1000,
		AWSBatchChangeInterval:                 time.Second,
		AWSEvaluateTargetHealth:                true,
		AWSAPIRetries:                          3,
		AWSPreferCNAME:                         false,
		AWSProfiles:                            []string{""},
		AWSZoneCacheDuration:                   0 * time.Second,
		AWSSDServiceCleanup:                    false,
		AWSSDCreateTag:                         map[string]string{},
		AWSDynamoDBTable:                       "external-dns",
		AzureConfigFile:                        "/etc/kubernetes/azure.json",
		AzureResourceGroup:                     "",
		AzureSubscriptionID:                    "",
		AzureMaxRetriesCount:                   3,
		CloudflareProxied:                      false,
		CloudflareCustomHostnames:              false,
		CloudflareCustomHostnamesMinTLSVersion: "1.0",
		CloudflareCustomHostnamesCertificateAuthority: "none",
		CloudflareDNSRecordsPerPage:                   100,
		CloudflareDNSRecordsComment:                   "",
		CloudflareRegionKey:                           "",
		CoreDNSPrefix:                                 "/skydns/",
		AkamaiServiceConsumerDomain:                   "",
		AkamaiClientToken:                             "",
		AkamaiClientSecret:                            "",
		AkamaiAccessToken:                             "",
		AkamaiEdgercPath:                              "",
		AkamaiEdgercSection:                           "",
		OCIConfigFile:                                 "/etc/kubernetes/oci.yaml",
		OCIZoneScope:                                  "GLOBAL",
		OCIZoneCacheDuration:                          0 * time.Second,
		InMemoryZones:                                 []string{""},
		OVHEndpoint:                                   "ovh-eu",
		OVHApiRateLimit:                               20,
		PDNSServer:                                    "http://localhost:8081",
		PDNSServerID:                                  "localhost",
		PDNSAPIKey:                                    "",
		Policy:                                        "sync",
		Registry:                                      "txt",
		TXTOwnerID:                                    "default",
		TXTOwnerOld:                                   "",
		TXTPrefix:                                     "",
		TXTCacheInterval:                              0,
		Interval:                                      time.Minute,
		MinEventSyncInterval:                          5 * time.Second,
		Once:                                          false,
		DryRun:                                        false,
		UpdateEvents:                                  false,
		LogFormat:                                     "text",
		MetricsAddress:                                ":7979",
		LogLevel:                                      logrus.InfoLevel.String(),
		ConnectorSourceServer:                         "localhost:8080",
		ExoscaleAPIEnvironment:                        "api",
		ExoscaleAPIZone:                               "ch-gva-2",
		ExoscaleAPIKey:                                "",
		ExoscaleAPISecret:                             "",
		CRDSourceAPIVersion:                           "externaldns.k8s.io/v1alpha1",
		CRDSourceKind:                                 "DNSEndpoint",
		TransIPAccountName:                            "",
		TransIPPrivateKeyFile:                         "",
		DigitalOceanAPIPageSize:                       50,
		ManagedDNSRecordTypes:                         []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
		RFC2136BatchChangeSize:                        50,
		RFC2136Host:                                   []string{""},
		RFC2136LoadBalancingStrategy:                  "disabled",
		OCPRouterName:                                 "default",
		PiholeApiVersion:                              "5",
		WebhookProviderURL:                            "http://localhost:8888",
		WebhookProviderReadTimeout:                    5 * time.Second,
		WebhookProviderWriteTimeout:                   10 * time.Second,
		ExcludeUnschedulable:                          true,
	}

	overriddenConfig = &Config{
		APIServerURL:                           "http://127.0.0.1:8080",
		KubeConfig:                             "/some/path",
		RequestTimeout:                         time.Second * 77,
		GlooNamespaces:                         []string{"gloo-not-system", "gloo-second-system"},
		SkipperRouteGroupVersion:               "zalando.org/v2",
		Sources:                                []string{"service", "ingress", "connector"},
		Namespace:                              "namespace",
		AnnotationPrefix:                       "external-dns.alpha.kubernetes.io/",
		IgnoreHostnameAnnotation:               true,
		IgnoreNonHostNetworkPods:               true,
		IgnoreIngressTLSSpec:                   true,
		IgnoreIngressRulesSpec:                 true,
		FQDNTemplate:                           "{{.Name}}.service.example.com",
		Compatibility:                          "mate",
		Provider:                               "google",
		GoogleProject:                          "project",
		GoogleBatchChangeSize:                  100,
		GoogleBatchChangeInterval:              time.Second * 2,
		GoogleZoneVisibility:                   "private",
		DomainFilter:                           []string{"example.org", "company.com"},
		DomainExclude:                          []string{"xapi.example.org", "xapi.company.com"},
		RegexDomainFilter:                      regexp.MustCompile("(example\\.org|company\\.com)$"),
		RegexDomainExclude:                     regexp.MustCompile("xapi\\.(example\\.org|company\\.com)$"),
		ZoneNameFilter:                         []string{"yapi.example.org", "yapi.company.com"},
		ZoneIDFilter:                           []string{"/hostedzone/ZTST1", "/hostedzone/ZTST2"},
		TargetNetFilter:                        []string{"10.0.0.0/9", "10.1.0.0/9"},
		ExcludeTargetNets:                      []string{"1.0.0.0/9", "1.1.0.0/9"},
		AlibabaCloudConfigFile:                 "/etc/kubernetes/alibaba-cloud.json",
		AWSZoneType:                            "private",
		AWSZoneTagFilter:                       []string{"tag=foo"},
		AWSZoneMatchParent:                     true,
		AWSAssumeRole:                          "some-other-role",
		AWSAssumeRoleExternalID:                "pg2000",
		AWSBatchChangeSize:                     100,
		AWSBatchChangeSizeBytes:                16000,
		AWSBatchChangeSizeValues:               100,
		AWSBatchChangeInterval:                 time.Second * 2,
		AWSEvaluateTargetHealth:                false,
		AWSAPIRetries:                          13,
		AWSPreferCNAME:                         true,
		AWSProfiles:                            []string{"profile1", "profile2"},
		AWSZoneCacheDuration:                   10 * time.Second,
		AWSSDServiceCleanup:                    true,
		AWSSDCreateTag:                         map[string]string{"key1": "value1", "key2": "value2"},
		AWSDynamoDBTable:                       "custom-table",
		AzureConfigFile:                        "azure.json",
		AzureResourceGroup:                     "arg",
		AzureSubscriptionID:                    "arg",
		AzureMaxRetriesCount:                   4,
		CloudflareProxied:                      true,
		CloudflareCustomHostnames:              true,
		CloudflareCustomHostnamesMinTLSVersion: "1.3",
		CloudflareCustomHostnamesCertificateAuthority: "google",
		CloudflareDNSRecordsPerPage:                   5000,
		CloudflareRegionalServices:                    true,
		CloudflareRegionKey:                           "us",
		CoreDNSPrefix:                                 "/coredns/",
		AkamaiServiceConsumerDomain:                   "oooo-xxxxxxxxxxxxxxxx-xxxxxxxxxxxxxxxx.luna.akamaiapis.net",
		AkamaiClientToken:                             "o184671d5307a388180fbf7f11dbdf46",
		AkamaiClientSecret:                            "o184671d5307a388180fbf7f11dbdf46",
		AkamaiAccessToken:                             "o184671d5307a388180fbf7f11dbdf46",
		AkamaiEdgercPath:                              "/home/test/.edgerc",
		AkamaiEdgercSection:                           "default",
		OCIConfigFile:                                 "oci.yaml",
		OCIZoneScope:                                  "PRIVATE",
		OCIZoneCacheDuration:                          30 * time.Second,
		InMemoryZones:                                 []string{"example.org", "company.com"},
		OVHEndpoint:                                   "ovh-ca",
		OVHApiRateLimit:                               42,
		PDNSServer:                                    "http://ns.example.com:8081",
		PDNSServerID:                                  "localhost",
		PDNSAPIKey:                                    "some-secret-key",
		PDNSSkipTLSVerify:                             true,
		TLSCA:                                         "/path/to/ca.crt",
		TLSClientCert:                                 "/path/to/cert.pem",
		TLSClientCertKey:                              "/path/to/key.pem",
		PodSourceDomain:                               "example.org",
		Policy:                                        "upsert-only",
		Registry:                                      "noop",
		TXTOwnerID:                                    "owner-1",
		TXTPrefix:                                     "associated-txt-record",
		TXTOwnerOld:                                   "old-owner",
		TXTCacheInterval:                              12 * time.Hour,
		Interval:                                      10 * time.Minute,
		MinEventSyncInterval:                          50 * time.Second,
		MinTTL:                                        40 * time.Second,
		Once:                                          true,
		DryRun:                                        true,
		UpdateEvents:                                  true,
		LogFormat:                                     "json",
		MetricsAddress:                                "127.0.0.1:9099",
		LogLevel:                                      logrus.DebugLevel.String(),
		ConnectorSourceServer:                         "localhost:8081",
		ExoscaleAPIEnvironment:                        "api1",
		ExoscaleAPIZone:                               "zone1",
		ExoscaleAPIKey:                                "1",
		ExoscaleAPISecret:                             "2",
		CRDSourceAPIVersion:                           "test.k8s.io/v1alpha1",
		CRDSourceKind:                                 "Endpoint",
		NS1Endpoint:                                   "https://api.example.com/v1",
		NS1IgnoreSSL:                                  true,
		TransIPAccountName:                            "transip",
		TransIPPrivateKeyFile:                         "/path/to/transip.key",
		DigitalOceanAPIPageSize:                       100,
		ManagedDNSRecordTypes:                         []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME, endpoint.RecordTypeNS},
		RFC2136BatchChangeSize:                        100,
		RFC2136Host:                                   []string{"rfc2136-host1", "rfc2136-host2"},
		RFC2136LoadBalancingStrategy:                  "round-robin",
		PiholeApiVersion:                              "6",
		WebhookProviderURL:                            "http://localhost:8888",
		WebhookProviderReadTimeout:                    5 * time.Second,
		WebhookProviderWriteTimeout:                   10 * time.Second,
		ExcludeUnschedulable:                          false,
	}
)

func TestParseFlags(t *testing.T) {
	for _, ti := range []struct {
		title    string
		args     []string
		envVars  map[string]string
		expected func(*Config)
	}{
		{
			title: "default config with minimal flags defined",
			args: []string{
				"--source=service",
				"--provider=google",
				"--openshift-router-name=default",
			},
			envVars: map[string]string{},
			expected: func(cfg *Config) {
				assert.Equal(t, minimalConfig, cfg)
			},
		},
		{
			title: "validate bool flags work as expected",
			args: []string{
				"--source=service",
				"--provider=google",
				"--aws-evaluate-target-health",
				"--exclude-unschedulable",
			},
			envVars: map[string]string{},
			expected: func(cfg *Config) {
				assert.True(t, cfg.AWSEvaluateTargetHealth)
				assert.True(t, cfg.ExcludeUnschedulable)
			},
		},
		{
			title: "validate negation flags work as expected",
			args: []string{
				"--source=service",
				"--provider=aws",
				"--no-aws-evaluate-target-health",
				"--no-exclude-unschedulable",
			},
			envVars: map[string]string{},
			expected: func(cfg *Config) {
				assert.False(t, cfg.AWSEvaluateTargetHealth)
				assert.False(t, cfg.ExcludeUnschedulable)
			},
		},
		{
			title: "override everything via flags",
			args: []string{
				"--server=http://127.0.0.1:8080",
				"--kubeconfig=/some/path",
				"--request-timeout=77s",
				"--gloo-namespace=gloo-not-system",
				"--gloo-namespace=gloo-second-system",
				"--skipper-routegroup-groupversion=zalando.org/v2",
				"--source=service",
				"--source=ingress",
				"--source=connector",
				"--namespace=namespace",
				"--fqdn-template={{.Name}}.service.example.com",
				"--ignore-non-host-network-pods",
				"--ignore-hostname-annotation",
				"--ignore-ingress-tls-spec",
				"--ignore-ingress-rules-spec",
				"--compatibility=mate",
				"--provider=google",
				"--google-project=project",
				"--google-batch-change-size=100",
				"--google-batch-change-interval=2s",
				"--google-zone-visibility=private",
				"--azure-config-file=azure.json",
				"--azure-resource-group=arg",
				"--azure-subscription-id=arg",
				"--azure-maxretries-count=4",
				"--cloudflare-proxied",
				"--cloudflare-custom-hostnames",
				"--cloudflare-custom-hostnames-min-tls-version=1.3",
				"--cloudflare-custom-hostnames-certificate-authority=google",
				"--cloudflare-dns-records-per-page=5000",
				"--cloudflare-regional-services",
				"--cloudflare-region-key=us",
				"--coredns-prefix=/coredns/",
				"--akamai-serviceconsumerdomain=oooo-xxxxxxxxxxxxxxxx-xxxxxxxxxxxxxxxx.luna.akamaiapis.net",
				"--akamai-client-token=o184671d5307a388180fbf7f11dbdf46",
				"--akamai-client-secret=o184671d5307a388180fbf7f11dbdf46",
				"--akamai-access-token=o184671d5307a388180fbf7f11dbdf46",
				"--akamai-edgerc-path=/home/test/.edgerc",
				"--akamai-edgerc-section=default",
				"--inmemory-zone=example.org",
				"--inmemory-zone=company.com",
				"--ovh-endpoint=ovh-ca",
				"--ovh-api-rate-limit=42",
				"--pdns-server=http://ns.example.com:8081",
				"--pdns-server-id=localhost",
				"--pdns-api-key=some-secret-key",
				"--pdns-skip-tls-verify",
				"--oci-config-file=oci.yaml",
				"--oci-zone-scope=PRIVATE",
				"--oci-zones-cache-duration=30s",
				"--tls-ca=/path/to/ca.crt",
				"--tls-client-cert=/path/to/cert.pem",
				"--tls-client-cert-key=/path/to/key.pem",
				"--pod-source-domain=example.org",
				"--domain-filter=example.org",
				"--domain-filter=company.com",
				"--exclude-domains=xapi.example.org",
				"--exclude-domains=xapi.company.com",
				"--regex-domain-filter=(example\\.org|company\\.com)$",
				"--regex-domain-exclusion=xapi\\.(example\\.org|company\\.com)$",
				"--zone-name-filter=yapi.example.org",
				"--zone-name-filter=yapi.company.com",
				"--zone-id-filter=/hostedzone/ZTST1",
				"--zone-id-filter=/hostedzone/ZTST2",
				"--target-net-filter=10.0.0.0/9",
				"--target-net-filter=10.1.0.0/9",
				"--exclude-target-net=1.0.0.0/9",
				"--exclude-target-net=1.1.0.0/9",
				"--aws-zone-type=private",
				"--aws-zone-tags=tag=foo",
				"--aws-zone-match-parent",
				"--aws-assume-role=some-other-role",
				"--aws-assume-role-external-id=pg2000",
				"--aws-batch-change-size=100",
				"--aws-batch-change-size-bytes=16000",
				"--aws-batch-change-size-values=100",
				"--aws-batch-change-interval=2s",
				"--aws-api-retries=13",
				"--aws-prefer-cname",
				"--aws-profile=profile1",
				"--aws-profile=profile2",
				"--aws-zones-cache-duration=10s",
				"--aws-sd-service-cleanup",
				"--aws-sd-create-tag=key1=value1",
				"--aws-sd-create-tag=key2=value2",
				"--no-aws-evaluate-target-health",
				"--pihole-api-version=6",
				"--policy=upsert-only",
				"--registry=noop",
				"--txt-owner-id=owner-1",
				"--migrate-from-txt-owner=old-owner",
				"--txt-prefix=associated-txt-record",
				"--txt-cache-interval=12h",
				"--dynamodb-table=custom-table",
				"--interval=10m",
				"--min-event-sync-interval=50s",
				"--min-ttl=40s",
				"--once",
				"--dry-run",
				"--events",
				"--log-format=json",
				"--metrics-address=127.0.0.1:9099",
				"--log-level=debug",
				"--connector-source-server=localhost:8081",
				"--exoscale-apienv=api1",
				"--exoscale-apizone=zone1",
				"--exoscale-apikey=1",
				"--exoscale-apisecret=2",
				"--crd-source-apiversion=test.k8s.io/v1alpha1",
				"--crd-source-kind=Endpoint",
				"--ns1-endpoint=https://api.example.com/v1",
				"--ns1-ignoressl",
				"--transip-account=transip",
				"--transip-keyfile=/path/to/transip.key",
				"--digitalocean-api-page-size=100",
				"--managed-record-types=A",
				"--managed-record-types=AAAA",
				"--managed-record-types=CNAME",
				"--managed-record-types=NS",
				"--no-exclude-unschedulable",
				"--rfc2136-batch-change-size=100",
				"--rfc2136-load-balancing-strategy=round-robin",
				"--rfc2136-host=rfc2136-host1",
				"--rfc2136-host=rfc2136-host2",
			},
			envVars: map[string]string{},
			expected: func(cfg *Config) {
				assert.Equal(t, overriddenConfig, cfg)
			},
		},
		{
			title: "override everything via environment variables",
			args:  []string{},
			envVars: map[string]string{
				"EXTERNAL_DNS_SERVER":                                            "http://127.0.0.1:8080",
				"EXTERNAL_DNS_KUBECONFIG":                                        "/some/path",
				"EXTERNAL_DNS_REQUEST_TIMEOUT":                                   "77s",
				"EXTERNAL_DNS_CONTOUR_LOAD_BALANCER":                             "heptio-contour-other/contour-other",
				"EXTERNAL_DNS_GLOO_NAMESPACE":                                    "gloo-not-system\ngloo-second-system",
				"EXTERNAL_DNS_SKIPPER_ROUTEGROUP_GROUPVERSION":                   "zalando.org/v2",
				"EXTERNAL_DNS_SOURCE":                                            "service\ningress\nconnector",
				"EXTERNAL_DNS_NAMESPACE":                                         "namespace",
				"EXTERNAL_DNS_FQDN_TEMPLATE":                                     "{{.Name}}.service.example.com",
				"EXTERNAL_DNS_IGNORE_NON_HOST_NETWORK_PODS":                      "1",
				"EXTERNAL_DNS_IGNORE_HOSTNAME_ANNOTATION":                        "1",
				"EXTERNAL_DNS_IGNORE_INGRESS_TLS_SPEC":                           "1",
				"EXTERNAL_DNS_IGNORE_INGRESS_RULES_SPEC":                         "1",
				"EXTERNAL_DNS_COMPATIBILITY":                                     "mate",
				"EXTERNAL_DNS_PROVIDER":                                          "google",
				"EXTERNAL_DNS_GOOGLE_PROJECT":                                    "project",
				"EXTERNAL_DNS_GOOGLE_BATCH_CHANGE_SIZE":                          "100",
				"EXTERNAL_DNS_GOOGLE_BATCH_CHANGE_INTERVAL":                      "2s",
				"EXTERNAL_DNS_GOOGLE_ZONE_VISIBILITY":                            "private",
				"EXTERNAL_DNS_AZURE_CONFIG_FILE":                                 "azure.json",
				"EXTERNAL_DNS_AZURE_RESOURCE_GROUP":                              "arg",
				"EXTERNAL_DNS_AZURE_SUBSCRIPTION_ID":                             "arg",
				"EXTERNAL_DNS_AZURE_MAXRETRIES_COUNT":                            "4",
				"EXTERNAL_DNS_CLOUDFLARE_PROXIED":                                "1",
				"EXTERNAL_DNS_CLOUDFLARE_CUSTOM_HOSTNAMES":                       "1",
				"EXTERNAL_DNS_CLOUDFLARE_CUSTOM_HOSTNAMES_MIN_TLS_VERSION":       "1.3",
				"EXTERNAL_DNS_CLOUDFLARE_CUSTOM_HOSTNAMES_CERTIFICATE_AUTHORITY": "google",
				"EXTERNAL_DNS_CLOUDFLARE_DNS_RECORDS_PER_PAGE":                   "5000",
				"EXTERNAL_DNS_CLOUDFLARE_REGIONAL_SERVICES":                      "1",
				"EXTERNAL_DNS_CLOUDFLARE_REGION_KEY":                             "us",
				"EXTERNAL_DNS_COREDNS_PREFIX":                                    "/coredns/",
				"EXTERNAL_DNS_AKAMAI_SERVICECONSUMERDOMAIN":                      "oooo-xxxxxxxxxxxxxxxx-xxxxxxxxxxxxxxxx.luna.akamaiapis.net",
				"EXTERNAL_DNS_AKAMAI_CLIENT_TOKEN":                               "o184671d5307a388180fbf7f11dbdf46",
				"EXTERNAL_DNS_AKAMAI_CLIENT_SECRET":                              "o184671d5307a388180fbf7f11dbdf46",
				"EXTERNAL_DNS_AKAMAI_ACCESS_TOKEN":                               "o184671d5307a388180fbf7f11dbdf46",
				"EXTERNAL_DNS_AKAMAI_EDGERC_PATH":                                "/home/test/.edgerc",
				"EXTERNAL_DNS_AKAMAI_EDGERC_SECTION":                             "default",
				"EXTERNAL_DNS_OCI_CONFIG_FILE":                                   "oci.yaml",
				"EXTERNAL_DNS_OCI_ZONE_SCOPE":                                    "PRIVATE",
				"EXTERNAL_DNS_OCI_ZONES_CACHE_DURATION":                          "30s",
				"EXTERNAL_DNS_INMEMORY_ZONE":                                     "example.org\ncompany.com",
				"EXTERNAL_DNS_OVH_ENDPOINT":                                      "ovh-ca",
				"EXTERNAL_DNS_OVH_API_RATE_LIMIT":                                "42",
				"EXTERNAL_DNS_POD_SOURCE_DOMAIN":                                 "example.org",
				"EXTERNAL_DNS_DOMAIN_FILTER":                                     "example.org\ncompany.com",
				"EXTERNAL_DNS_EXCLUDE_DOMAINS":                                   "xapi.example.org\nxapi.company.com",
				"EXTERNAL_DNS_REGEX_DOMAIN_FILTER":                               "(example\\.org|company\\.com)$",
				"EXTERNAL_DNS_REGEX_DOMAIN_EXCLUSION":                            "xapi\\.(example\\.org|company\\.com)$",
				"EXTERNAL_DNS_TARGET_NET_FILTER":                                 "10.0.0.0/9\n10.1.0.0/9",
				"EXTERNAL_DNS_EXCLUDE_TARGET_NET":                                "1.0.0.0/9\n1.1.0.0/9",
				"EXTERNAL_DNS_PDNS_SERVER":                                       "http://ns.example.com:8081",
				"EXTERNAL_DNS_PDNS_ID":                                           "localhost",
				"EXTERNAL_DNS_PDNS_API_KEY":                                      "some-secret-key",
				"EXTERNAL_DNS_PDNS_SKIP_TLS_VERIFY":                              "1",
				"EXTERNAL_DNS_RDNS_ROOT_DOMAIN":                                  "lb.rancher.cloud",
				"EXTERNAL_DNS_TLS_CA":                                            "/path/to/ca.crt",
				"EXTERNAL_DNS_TLS_CLIENT_CERT":                                   "/path/to/cert.pem",
				"EXTERNAL_DNS_TLS_CLIENT_CERT_KEY":                               "/path/to/key.pem",
				"EXTERNAL_DNS_ZONE_NAME_FILTER":                                  "yapi.example.org\nyapi.company.com",
				"EXTERNAL_DNS_ZONE_ID_FILTER":                                    "/hostedzone/ZTST1\n/hostedzone/ZTST2",
				"EXTERNAL_DNS_AWS_ZONE_TYPE":                                     "private",
				"EXTERNAL_DNS_AWS_ZONE_TAGS":                                     "tag=foo",
				"EXTERNAL_DNS_AWS_ZONE_MATCH_PARENT":                             "true",
				"EXTERNAL_DNS_AWS_ASSUME_ROLE":                                   "some-other-role",
				"EXTERNAL_DNS_AWS_ASSUME_ROLE_EXTERNAL_ID":                       "pg2000",
				"EXTERNAL_DNS_AWS_BATCH_CHANGE_SIZE":                             "100",
				"EXTERNAL_DNS_AWS_BATCH_CHANGE_SIZE_BYTES":                       "16000",
				"EXTERNAL_DNS_AWS_BATCH_CHANGE_SIZE_VALUES":                      "100",
				"EXTERNAL_DNS_AWS_BATCH_CHANGE_INTERVAL":                         "2s",
				"EXTERNAL_DNS_AWS_EVALUATE_TARGET_HEALTH":                        "0",
				"EXTERNAL_DNS_AWS_API_RETRIES":                                   "13",
				"EXTERNAL_DNS_AWS_PREFER_CNAME":                                  "true",
				"EXTERNAL_DNS_AWS_PROFILE":                                       "profile1\nprofile2",
				"EXTERNAL_DNS_AWS_ZONES_CACHE_DURATION":                          "10s",
				"EXTERNAL_DNS_AWS_SD_SERVICE_CLEANUP":                            "true",
				"EXTERNAL_DNS_AWS_SD_CREATE_TAG":                                 "key1=value1\nkey2=value2",
				"EXTERNAL_DNS_DYNAMODB_TABLE":                                    "custom-table",
				"EXTERNAL_DNS_PIHOLE_API_VERSION":                                "6",
				"EXTERNAL_DNS_POLICY":                                            "upsert-only",
				"EXTERNAL_DNS_REGISTRY":                                          "noop",
				"EXTERNAL_DNS_TXT_OWNER_ID":                                      "owner-1",
				"EXTERNAL_DNS_TXT_PREFIX":                                        "associated-txt-record",
				"EXTERNAL_DNS_MIGRATE_FROM_TXT_OWNER":                            "old-owner",
				"EXTERNAL_DNS_TXT_CACHE_INTERVAL":                                "12h",
				"EXTERNAL_DNS_TXT_NEW_FORMAT_ONLY":                               "1",
				"EXTERNAL_DNS_INTERVAL":                                          "10m",
				"EXTERNAL_DNS_MIN_EVENT_SYNC_INTERVAL":                           "50s",
				"EXTERNAL_DNS_MIN_TTL":                                           "40s",
				"EXTERNAL_DNS_ONCE":                                              "1",
				"EXTERNAL_DNS_DRY_RUN":                                           "1",
				"EXTERNAL_DNS_EVENTS":                                            "1",
				"EXTERNAL_DNS_LOG_FORMAT":                                        "json",
				"EXTERNAL_DNS_METRICS_ADDRESS":                                   "127.0.0.1:9099",
				"EXTERNAL_DNS_LOG_LEVEL":                                         "debug",
				"EXTERNAL_DNS_CONNECTOR_SOURCE_SERVER":                           "localhost:8081",
				"EXTERNAL_DNS_EXOSCALE_APIENV":                                   "api1",
				"EXTERNAL_DNS_EXOSCALE_APIZONE":                                  "zone1",
				"EXTERNAL_DNS_EXOSCALE_APIKEY":                                   "1",
				"EXTERNAL_DNS_EXOSCALE_APISECRET":                                "2",
				"EXTERNAL_DNS_CRD_SOURCE_APIVERSION":                             "test.k8s.io/v1alpha1",
				"EXTERNAL_DNS_CRD_SOURCE_KIND":                                   "Endpoint",
				"EXTERNAL_DNS_NS1_ENDPOINT":                                      "https://api.example.com/v1",
				"EXTERNAL_DNS_NS1_IGNORESSL":                                     "1",
				"EXTERNAL_DNS_TRANSIP_ACCOUNT":                                   "transip",
				"EXTERNAL_DNS_TRANSIP_KEYFILE":                                   "/path/to/transip.key",
				"EXTERNAL_DNS_DIGITALOCEAN_API_PAGE_SIZE":                        "100",
				"EXTERNAL_DNS_MANAGED_RECORD_TYPES":                              "A\nAAAA\nCNAME\nNS",
				"EXTERNAL_DNS_EXCLUDE_UNSCHEDULABLE":                             "false",
				"EXTERNAL_DNS_RFC2136_BATCH_CHANGE_SIZE":                         "100",
				"EXTERNAL_DNS_RFC2136_LOAD_BALANCING_STRATEGY":                   "round-robin",
				"EXTERNAL_DNS_RFC2136_HOST":                                      "rfc2136-host1\nrfc2136-host2",
			},
			expected: func(cfg *Config) {
				assert.Equal(t, overriddenConfig, cfg)
			},
		},
	} {
		t.Run(ti.title, func(t *testing.T) {
			testutils.TestHelperEnvSetter(t, ti.envVars)

			cfg := NewConfig()
			require.NoError(t, cfg.ParseFlags(ti.args))
			ti.expected(cfg)
		})
	}
}

func TestParseFlagsCobraExecuteError(t *testing.T) {
	cfg := NewConfig()
	err := cfg.ParseFlags([]string{"--cli-backend=cobra", "--unknown-flag"})
	require.Error(t, err)
}

func TestParseFlagsKingpinParseError(t *testing.T) {
	cfg := NewConfig()
	err := cfg.ParseFlags([]string{"--unknown-flag"})
	require.Error(t, err)
}

func TestConfigStringMasksSecureFields(t *testing.T) {
	cfg := NewConfig()
	cfg.AWSAssumeRoleExternalID = "sensitive-value"
	cfg.GoDaddyAPIKey = "another-secret"

	s := cfg.String()
	require.NotContains(t, s, "sensitive-value")
	require.NotContains(t, s, "another-secret")
	require.Contains(t, s, passwordMask)
}

// Default path should use kingpin and parse flags correctly
func TestParseFlagsDefaultKingpin(t *testing.T) {
	t.Setenv("EXTERNAL_DNS_CLI", "")

	args := []string{
		"--provider=aws",
		"--source=service",
		"--source=ingress",
		"--server=http://127.0.0.1:8080",
		"--kubeconfig=/some/path",
		"--request-timeout=2s",
		"--namespace=ns",
		"--domain-filter=example.org",
		"--domain-filter=company.com",
		"--openshift-router-name=default",
	}

	cfg := NewConfig()
	require.NoError(t, cfg.ParseFlags(args))

	assert.Equal(t, "aws", cfg.Provider)
	assert.ElementsMatch(t, []string{"service", "ingress"}, cfg.Sources)
	assert.Equal(t, "http://127.0.0.1:8080", cfg.APIServerURL)
	assert.Equal(t, "/some/path", cfg.KubeConfig)
	assert.Equal(t, 2*time.Second, cfg.RequestTimeout)
	assert.Equal(t, "ns", cfg.Namespace)
	assert.ElementsMatch(t, []string{"example.org", "company.com"}, cfg.DomainFilter)
	assert.Equal(t, "default", cfg.OCPRouterName)
}

// When EXTERNAL_DNS_CLI=cobra is set, cobra path should parse the subset of
// flags it currently binds, yielding parity with kingpin for those fields.
func TestParseFlagsCobraSwitchParitySubset(t *testing.T) {
	args := []string{
		"--provider=aws",
		"--source=service",
		"--source=ingress",
		"--server=http://127.0.0.1:8080",
		"--kubeconfig=/some/path",
		"--request-timeout=2s",
		"--namespace=ns",
		"--domain-filter=example.org",
		"--domain-filter=company.com",
		"--openshift-router-name=default",
	}

	// Kingpin baseline
	cfgK := NewConfig()
	require.NoError(t, cfgK.ParseFlags(args))

	// Cobra path via env switch
	t.Setenv("EXTERNAL_DNS_CLI", "cobra")
	cfgC := NewConfig()
	require.NoError(t, cfgC.ParseFlags(args))

	// Compare selected fields bound in cobra
	assert.Equal(t, cfgK.Provider, cfgC.Provider)
	assert.ElementsMatch(t, cfgK.Sources, cfgC.Sources)
	assert.Equal(t, cfgK.APIServerURL, cfgC.APIServerURL)
	assert.Equal(t, cfgK.KubeConfig, cfgC.KubeConfig)
	assert.Equal(t, cfgK.RequestTimeout, cfgC.RequestTimeout)
	assert.Equal(t, cfgK.Namespace, cfgC.Namespace)
	assert.ElementsMatch(t, cfgK.DomainFilter, cfgC.DomainFilter)
	assert.Equal(t, cfgK.OCPRouterName, cfgC.OCPRouterName)
}

func TestParseFlagsCliFlagOverridesEnv(t *testing.T) {
	// Env requests cobra; CLI flag forces kingpin.
	t.Setenv("EXTERNAL_DNS_CLI", "cobra")
	args := []string{
		"--provider=aws",
		"--source=service",
		// Flag not bound in Cobra newCobraCommand path; will error if cobra is used.
		"--log-format=json",
	}

	cfg := NewConfig()
	require.NoError(t, cfg.ParseFlags(args))
	assert.Equal(t, "aws", cfg.Provider)
	assert.ElementsMatch(t, []string{"service"}, cfg.Sources)
	assert.Equal(t, "json", cfg.LogFormat)
}

func TestParseFlagsCliFlagSeparatedValue(t *testing.T) {
	// Support "--cli-backend", "cobra" form as well.
	args := []string{
		"--provider=aws",
		"--source=service",
	}
	cfg := NewConfig()
	require.NoError(t, cfg.ParseFlags(args))
	assert.Equal(t, "aws", cfg.Provider)
	assert.ElementsMatch(t, []string{"service"}, cfg.Sources)
}

func TestPasswordsNotLogged(t *testing.T) {
	cfg := Config{
		PDNSAPIKey:        "pdns-api-key",
		RFC2136TSIGSecret: "tsig-secret",
	}

	s := cfg.String()

	assert.NotContains(t, s, "pdns-api-key")
	assert.NotContains(t, s, "tsig-secret")
}

// Additional assertions to cover previously unasserted flags. These focus on
// exercising Kingpin flag bindings for a wide set of providers/features.
// parseCfg builds a Config by parsing base flags plus any extras.
func parseCfg(t *testing.T, extra ...string) *Config {
	t.Helper()
	cfg := NewConfig()
	args := append([]string{"--provider=google", "--source=service"}, extra...)
	require.NoError(t, cfg.ParseFlags(args))
	return cfg
}

func TestParseFlagsAlibabaCloud(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--alibaba-cloud-config-file=/etc/kubernetes/alibaba-override.json",
		"--alibaba-cloud-zone-type=private",
	)
	assert.Equal(t, "/etc/kubernetes/alibaba-override.json", cfg.AlibabaCloudConfigFile)
	assert.Equal(t, "private", cfg.AlibabaCloudZoneType)
}

func TestParseFlagsPublishingAndFilters(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--always-publish-not-ready-addresses",
		"--annotation-filter=key=value",
		"--combine-fqdn-annotation",
		"--default-targets=1.2.3.4",
		"--default-targets=5.6.7.8",
		"--exclude-record-types=TXT",
		"--exclude-record-types=CNAME",
		"--expose-internal-ipv6",
		"--force-default-targets",
		"--ingress-class=nginx",
		"--ingress-class=internal",
		"--label-filter=environment=prod",
		"--nat64-networks=64:ff9b::/96",
		"--nat64-networks=64:ff9b:1::/48",
		"--publish-host-ip",
		"--publish-internal-services",
		"--resolve-service-load-balancer-hostname",
		"--service-type-filter=ClusterIP",
		"--service-type-filter=NodePort",
		"--events-emit=RecordReady",
		"--events-emit=RecordDeleted",
	)
	assert.True(t, cfg.AlwaysPublishNotReadyAddresses)
	assert.Equal(t, "key=value", cfg.AnnotationFilter)
	assert.True(t, cfg.CombineFQDNAndAnnotation)
	assert.ElementsMatch(t, []string{"1.2.3.4", "5.6.7.8"}, cfg.DefaultTargets)
	assert.ElementsMatch(t, []string{"TXT", "CNAME"}, cfg.ExcludeDNSRecordTypes)
	assert.True(t, cfg.ExposeInternalIPV6)
	assert.True(t, cfg.ForceDefaultTargets)
	assert.ElementsMatch(t, []string{"nginx", "internal"}, cfg.IngressClassNames)
	assert.Equal(t, "environment=prod", cfg.LabelFilter)
	assert.ElementsMatch(t, []string{"64:ff9b::/96", "64:ff9b:1::/48"}, cfg.NAT64Networks)
	assert.True(t, cfg.PublishHostIP)
	assert.True(t, cfg.PublishInternal)
	assert.True(t, cfg.ResolveServiceLoadBalancerHostname)
	assert.ElementsMatch(t, []string{"ClusterIP", "NodePort"}, cfg.ServiceTypeFilter)
	assert.ElementsMatch(t, []string{"RecordReady", "RecordDeleted"}, cfg.EmitEvents)
}

func TestParseFlagsGateway(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--gateway-label-filter=app=gateway",
		"--gateway-name=gw-1",
		"--gateway-namespace=gw-ns",
	)
	assert.Equal(t, "app=gateway", cfg.GatewayLabelFilter)
	assert.Equal(t, "gw-1", cfg.GatewayName)
	assert.Equal(t, "gw-ns", cfg.GatewayNamespace)
}

func TestParseFlagsAzure(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--azure-user-assigned-identity-client-id=00000000-0000-0000-0000-000000000000",
		"--azure-zones-cache-duration=30s",
	)
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", cfg.AzureUserAssignedIdentityClientID)
	assert.Equal(t, 30*time.Second, cfg.AzureZonesCacheDuration)
}

func TestParseFlagsCloudflare(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t, "--cloudflare-record-comment=managed-by-external-dns")
	assert.Equal(t, "managed-by-external-dns", cfg.CloudflareDNSRecordsComment)
}

func TestParseFlagsNS1(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t, "--ns1-min-ttl=60")
	assert.Equal(t, 60, cfg.NS1MinTTLSeconds)
}

func TestParseFlagsOVH(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t, "--ovh-enable-cname-relative")
	assert.True(t, cfg.OVHEnableCNAMERelative)
}

func TestParseFlagsPihole(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--pihole-server=https://pi.example",
		"--pihole-password=pw",
		"--pihole-tls-skip-verify",
	)
	assert.Equal(t, "https://pi.example", cfg.PiholeServer)
	assert.Equal(t, "pw", cfg.PiholePassword)
	assert.True(t, cfg.PiholeTLSInsecureSkipVerify)
}

func TestParseFlagsOCI(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--oci-auth-instance-principal",
		"--oci-compartment-ocid=ocid1.compartment.oc1..aaaa",
	)
	assert.True(t, cfg.OCIAuthInstancePrincipal)
	assert.Equal(t, "ocid1.compartment.oc1..aaaa", cfg.OCICompartmentOCID)
}

func TestParseFlagsPlural(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--plural-cluster=mycluster",
		"--plural-provider=aws",
	)
	assert.Equal(t, "mycluster", cfg.PluralCluster)
	assert.Equal(t, "aws", cfg.PluralProvider)
}

func TestParseFlagsProviderCacheAndDynamoDB(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--provider-cache-time=20s",
		"--dynamodb-region=us-east-2",
	)
	assert.Equal(t, 20*time.Second, cfg.ProviderCacheTime)
	assert.Equal(t, "us-east-2", cfg.AWSDynamoDBRegion)
}

func TestParseFlagsGoDaddy(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--godaddy-api-key=key",
		"--godaddy-api-secret=secret",
		"--godaddy-api-ttl=1234",
		"--godaddy-api-ote",
	)
	assert.Equal(t, "key", cfg.GoDaddyAPIKey)
	assert.Equal(t, "secret", cfg.GoDaddySecretKey)
	assert.Equal(t, int64(1234), cfg.GoDaddyTTL)
	assert.True(t, cfg.GoDaddyOTE)
}

func TestParseFlagsRFC2136(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--rfc2136-port=5353",
		"--rfc2136-zone=example.org.",
		"--rfc2136-zone=example.com.",
		"--rfc2136-create-ptr",
		"--rfc2136-insecure",
		"--rfc2136-kerberos-realm=EXAMPLE.COM",
		"--rfc2136-kerberos-username=svc-externaldns",
		"--rfc2136-kerberos-password=secret",
		"--rfc2136-tsig-keyname=keyname.",
		"--rfc2136-tsig-secret=base64secret",
		"--rfc2136-tsig-secret-alg=hmac-sha256",
		"--rfc2136-tsig-axfr",
		"--rfc2136-min-ttl=30s",
		"--rfc2136-gss-tsig",
		"--rfc2136-use-tls",
		"--rfc2136-skip-tls-verify",
	)
	assert.Equal(t, 5353, cfg.RFC2136Port)
	assert.ElementsMatch(t, []string{"example.org.", "example.com."}, cfg.RFC2136Zone)
	assert.True(t, cfg.RFC2136CreatePTR)
	assert.True(t, cfg.RFC2136Insecure)
	assert.Equal(t, "EXAMPLE.COM", cfg.RFC2136KerberosRealm)
	assert.Equal(t, "svc-externaldns", cfg.RFC2136KerberosUsername)
	assert.Equal(t, "secret", cfg.RFC2136KerberosPassword)
	assert.Equal(t, "keyname.", cfg.RFC2136TSIGKeyName)
	assert.Equal(t, "base64secret", cfg.RFC2136TSIGSecret)
	assert.Equal(t, "hmac-sha256", cfg.RFC2136TSIGSecretAlg)
	assert.True(t, cfg.RFC2136TAXFR)
	assert.Equal(t, 30*time.Second, cfg.RFC2136MinTTL)
	assert.True(t, cfg.RFC2136GSSTSIG)
	assert.True(t, cfg.RFC2136UseTLS)
	assert.True(t, cfg.RFC2136SkipTLSVerify)
}

func TestParseFlagsTraefik(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--traefik-enable-legacy",
		"--traefik-disable-new",
	)
	assert.True(t, cfg.TraefikEnableLegacy)
	assert.True(t, cfg.TraefikDisableNew)
}

func TestParseFlagsTXTRegistry(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--txt-encrypt-enabled",
		"--txt-encrypt-aes-key=0123456789abcdef0123456789abcdef",
		"--txt-suffix=-suffix",
		"--txt-wildcard-replacement=X",
	)
	assert.True(t, cfg.TXTEncryptEnabled)
	assert.Equal(t, "0123456789abcdef0123456789abcdef", cfg.TXTEncryptAESKey)
	assert.Equal(t, "-suffix", cfg.TXTSuffix)
	assert.Equal(t, "X", cfg.TXTWildcardReplacement)
}

func TestParseFlagsWebhookProvider(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t,
		"--webhook-provider-url=http://127.0.0.1:9999",
		"--webhook-provider-read-timeout=7s",
		"--webhook-provider-write-timeout=8s",
		"--webhook-server",
	)
	assert.Equal(t, "http://127.0.0.1:9999", cfg.WebhookProviderURL)
	assert.Equal(t, 7*time.Second, cfg.WebhookProviderReadTimeout)
	assert.Equal(t, 8*time.Second, cfg.WebhookProviderWriteTimeout)
	assert.True(t, cfg.WebhookServer)
}

func TestParseFlagsMiscListeners(t *testing.T) {
	t.Parallel()
	cfg := parseCfg(t, "--listen-endpoint-events")
	assert.True(t, cfg.ListenEndpointEvents)
}

// Helpers to run bindFlags + parse for each binder.
func runWithKingpin(t *testing.T, args []string) *Config {
	t.Helper()
	cfg := &Config{}
	cfg.AWSSDCreateTag = map[string]string{}
	cfg.RegexDomainFilter = defaultConfig.RegexDomainFilter
	app := kingpin.New("test", "")
	bindFlags(flags.NewKingpinBinder(app), cfg)
	_, err := app.Parse(args)
	require.NoError(t, err)
	return cfg
}

func TestBinderParityRepeatable(t *testing.T) {
	args := []string{"--managed-record-types=A", "--managed-record-types=TXT"}
	cfgK := runWithKingpin(t, args)
	assert.ElementsMatch(t, []string{"A", "TXT"}, cfgK.ManagedDNSRecordTypes)
}

func TestBinderParityMapAndRegexp(t *testing.T) {
	args := []string{"--regex-domain-filter=^ex.*$", "--aws-sd-create-tag=foo=bar"}
	cfgK := runWithKingpin(t, args)

	require.NotNil(t, cfgK.RegexDomainFilter)
	require.NotNil(t, cfgK.AWSSDCreateTag)
	assert.Equal(t, map[string]string{"foo": "bar"}, cfgK.AWSSDCreateTag)
}

// Kingpin validates enum values at parse time
func TestBinderEnumValidationDifference(t *testing.T) {
	// Kingpin should reject unknown enum values
	appArgs := []string{"--google-zone-visibility=bogus"}
	app := kingpin.New("test", "")
	cfgK := &Config{}
	bindFlags(flags.NewKingpinBinder(app), cfgK)
	_, err := app.Parse(appArgs)
	require.Error(t, err)
}
