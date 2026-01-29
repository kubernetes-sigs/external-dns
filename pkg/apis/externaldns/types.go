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
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/internal/flags"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"

	"github.com/alecthomas/kingpin/v2"
	"github.com/sirupsen/logrus"
)

const (
	passwordMask = "******"
)

// Config is a project-wide configuration
type Config struct {
	APIServerURL                                  string
	KubeConfig                                    string
	RequestTimeout                                time.Duration
	DefaultTargets                                []string
	GlooNamespaces                                []string
	SkipperRouteGroupVersion                      string
	Sources                                       []string
	Namespace                                     string
	AnnotationFilter                              string
	AnnotationPrefix                              string
	LabelFilter                                   string
	IngressClassNames                             []string
	FQDNTemplate                                  string
	CombineFQDNAndAnnotation                      bool
	IgnoreHostnameAnnotation                      bool
	IgnoreNonHostNetworkPods                      bool
	IgnoreIngressTLSSpec                          bool
	IgnoreIngressRulesSpec                        bool
	ListenEndpointEvents                          bool
	ExposeInternalIPV6                            bool
	GatewayName                                   string
	GatewayNamespace                              string
	GatewayLabelFilter                            string
	Compatibility                                 string
	PodSourceDomain                               string
	PublishInternal                               bool
	PublishHostIP                                 bool
	AlwaysPublishNotReadyAddresses                bool
	ConnectorSourceServer                         string
	Provider                                      string
	ProviderCacheTime                             time.Duration
	GoogleProject                                 string
	GoogleBatchChangeSize                         int
	GoogleBatchChangeInterval                     time.Duration
	GoogleZoneVisibility                          string
	DomainFilter                                  []string
	DomainExclude                                 []string
	RegexDomainFilter                             *regexp.Regexp
	RegexDomainExclude                            *regexp.Regexp
	ZoneNameFilter                                []string
	ZoneIDFilter                                  []string
	TargetNetFilter                               []string
	ExcludeTargetNets                             []string
	AlibabaCloudConfigFile                        string
	AlibabaCloudZoneType                          string
	AWSZoneType                                   string
	AWSZoneTagFilter                              []string
	AWSAssumeRole                                 string
	AWSProfiles                                   []string
	AWSAssumeRoleExternalID                       string `secure:"yes"`
	AWSBatchChangeSize                            int
	AWSBatchChangeSizeBytes                       int
	AWSBatchChangeSizeValues                      int
	AWSBatchChangeInterval                        time.Duration
	AWSEvaluateTargetHealth                       bool
	AWSAPIRetries                                 int
	AWSPreferCNAME                                bool
	AWSZoneCacheDuration                          time.Duration
	AWSSDServiceCleanup                           bool
	AWSSDCreateTag                                map[string]string
	AWSZoneMatchParent                            bool
	AWSDynamoDBRegion                             string
	AWSDynamoDBTable                              string
	AzureConfigFile                               string
	AzureResourceGroup                            string
	AzureSubscriptionID                           string
	AzureUserAssignedIdentityClientID             string
	AzureActiveDirectoryAuthorityHost             string
	AzureZonesCacheDuration                       time.Duration
	AzureMaxRetriesCount                          int
	CloudflareProxied                             bool
	CloudflareCustomHostnames                     bool
	CloudflareDNSRecordsPerPage                   int
	CloudflareDNSRecordsComment                   string
	CloudflareCustomHostnamesMinTLSVersion        string
	CloudflareCustomHostnamesCertificateAuthority string
	CloudflareRegionalServices                    bool
	CloudflareRegionKey                           string
	CoreDNSPrefix                                 string
	CoreDNSStrictlyOwned                          bool
	AkamaiServiceConsumerDomain                   string
	AkamaiClientToken                             string
	AkamaiClientSecret                            string
	AkamaiAccessToken                             string
	AkamaiEdgercPath                              string
	AkamaiEdgercSection                           string
	OCIConfigFile                                 string
	OCICompartmentOCID                            string
	OCIAuthInstancePrincipal                      bool
	OCIZoneScope                                  string
	OCIZoneCacheDuration                          time.Duration
	InMemoryZones                                 []string
	OVHEndpoint                                   string
	OVHApiRateLimit                               int
	OVHEnableCNAMERelative                        bool
	PDNSServer                                    string
	PDNSServerID                                  string
	PDNSAPIKey                                    string `secure:"yes"`
	PDNSSkipTLSVerify                             bool
	TLSCA                                         string
	TLSClientCert                                 string
	TLSClientCertKey                              string
	Policy                                        string
	Registry                                      string
	TXTOwnerID                                    string
	TXTOwnerOld                                   string
	TXTPrefix                                     string
	TXTSuffix                                     string
	TXTEncryptEnabled                             bool
	TXTEncryptAESKey                              string `secure:"yes"`
	Interval                                      time.Duration
	MinEventSyncInterval                          time.Duration
	MinTTL                                        time.Duration
	Once                                          bool
	DryRun                                        bool
	UpdateEvents                                  bool
	LogFormat                                     string
	MetricsAddress                                string
	LogLevel                                      string
	TXTCacheInterval                              time.Duration
	TXTWildcardReplacement                        string
	ExoscaleEndpoint                              string
	ExoscaleAPIKey                                string `secure:"yes"`
	ExoscaleAPISecret                             string `secure:"yes"`
	ExoscaleAPIEnvironment                        string
	ExoscaleAPIZone                               string
	CRDSourceAPIVersion                           string
	CRDSourceKind                                 string
	ServiceTypeFilter                             []string
	ResolveServiceLoadBalancerHostname            bool
	RFC2136Host                                   []string
	RFC2136Port                                   int
	RFC2136Zone                                   []string
	RFC2136Insecure                               bool
	RFC2136GSSTSIG                                bool
	RFC2136CreatePTR                              bool
	RFC2136KerberosRealm                          string
	RFC2136KerberosUsername                       string
	RFC2136KerberosPassword                       string `secure:"yes"`
	RFC2136TSIGKeyName                            string
	RFC2136TSIGSecret                             string `secure:"yes"`
	RFC2136TSIGSecretAlg                          string
	RFC2136TAXFR                                  bool
	RFC2136MinTTL                                 time.Duration
	RFC2136LoadBalancingStrategy                  string
	RFC2136BatchChangeSize                        int
	RFC2136UseTLS                                 bool
	RFC2136SkipTLSVerify                          bool
	NS1Endpoint                                   string
	NS1IgnoreSSL                                  bool
	NS1MinTTLSeconds                              int
	TransIPAccountName                            string
	TransIPPrivateKeyFile                         string
	DigitalOceanAPIPageSize                       int
	ManagedDNSRecordTypes                         []string
	ExcludeDNSRecordTypes                         []string
	GoDaddyAPIKey                                 string `secure:"yes"`
	GoDaddySecretKey                              string `secure:"yes"`
	GoDaddyTTL                                    int64
	GoDaddyOTE                                    bool
	OCPRouterName                                 string
	PiholeServer                                  string
	PiholePassword                                string `secure:"yes"`
	PiholeTLSInsecureSkipVerify                   bool
	PiholeApiVersion                              string
	PluralCluster                                 string
	PluralProvider                                string
	WebhookProviderURL                            string
	WebhookProviderReadTimeout                    time.Duration
	WebhookProviderWriteTimeout                   time.Duration
	WebhookServer                                 bool
	TraefikEnableLegacy                           bool
	TraefikDisableNew                             bool
	NAT64Networks                                 []string
	ExcludeUnschedulable                          bool
	EmitEvents                                    []string
	ForceDefaultTargets                           bool
	PreferAlias                                   bool
}

var defaultConfig = &Config{
	AkamaiAccessToken:           "",
	AkamaiClientSecret:          "",
	AkamaiClientToken:           "",
	AkamaiEdgercPath:            "",
	AkamaiEdgercSection:         "",
	AkamaiServiceConsumerDomain: "",
	AlibabaCloudConfigFile:      "/etc/kubernetes/alibaba-cloud.json",
	AnnotationFilter:            "",
	AnnotationPrefix:            annotations.DefaultAnnotationPrefix,
	APIServerURL:                "",
	AWSAPIRetries:               3,
	AWSAssumeRole:               "",
	AWSAssumeRoleExternalID:     "",
	AWSBatchChangeInterval:      time.Second,
	AWSBatchChangeSize:          1000,
	AWSBatchChangeSizeBytes:     32000,
	AWSBatchChangeSizeValues:    1000,
	AWSDynamoDBRegion:           "",
	AWSDynamoDBTable:            "external-dns",
	AWSEvaluateTargetHealth:     true,
	AWSPreferCNAME:              false,
	AWSSDCreateTag:              map[string]string{},
	AWSSDServiceCleanup:         false,
	AWSZoneCacheDuration:        0 * time.Second,
	AWSZoneMatchParent:          false,
	AWSZoneTagFilter:            []string{},
	AWSZoneType:                 "",
	AzureConfigFile:             "/etc/kubernetes/azure.json",
	AzureResourceGroup:          "",
	AzureSubscriptionID:         "",
	AzureZonesCacheDuration:     0 * time.Second,
	AzureMaxRetriesCount:        3,
	CloudflareCustomHostnamesCertificateAuthority: "none",
	CloudflareCustomHostnames:                     false,
	CloudflareCustomHostnamesMinTLSVersion:        "1.0",
	CloudflareDNSRecordsPerPage:                   100,
	CloudflareProxied:                             false,
	CloudflareRegionalServices:                    false,
	CloudflareRegionKey:                           "earth",

	CombineFQDNAndAnnotation:     false,
	Compatibility:                "",
	ConnectorSourceServer:        "localhost:8080",
	CoreDNSPrefix:                "/skydns/",
	CoreDNSStrictlyOwned:         false,
	CRDSourceAPIVersion:          "externaldns.k8s.io/v1alpha1",
	CRDSourceKind:                "DNSEndpoint",
	DefaultTargets:               []string{},
	DigitalOceanAPIPageSize:      50,
	DomainFilter:                 []string{},
	DryRun:                       false,
	ExcludeDNSRecordTypes:        []string{},
	DomainExclude:                []string{},
	ExcludeTargetNets:            []string{},
	EmitEvents:                   []string{},
	ExcludeUnschedulable:         true,
	ExoscaleAPIEnvironment:       "api",
	ExoscaleAPIKey:               "",
	ExoscaleAPISecret:            "",
	ExoscaleAPIZone:              "ch-gva-2",
	ExposeInternalIPV6:           false,
	FQDNTemplate:                 "",
	GatewayLabelFilter:           "",
	GatewayName:                  "",
	GatewayNamespace:             "",
	GlooNamespaces:               []string{"gloo-system"},
	GoDaddyAPIKey:                "",
	GoDaddyOTE:                   false,
	GoDaddySecretKey:             "",
	GoDaddyTTL:                   600,
	GoogleBatchChangeInterval:    time.Second,
	GoogleBatchChangeSize:        1000,
	GoogleProject:                "",
	GoogleZoneVisibility:         "",
	IgnoreHostnameAnnotation:     false,
	IgnoreIngressRulesSpec:       false,
	IgnoreIngressTLSSpec:         false,
	IngressClassNames:            nil,
	InMemoryZones:                []string{},
	Interval:                     time.Minute,
	KubeConfig:                   "",
	LabelFilter:                  labels.Everything().String(),
	LogFormat:                    "text",
	LogLevel:                     logrus.InfoLevel.String(),
	ManagedDNSRecordTypes:        []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
	MetricsAddress:               ":7979",
	MinEventSyncInterval:         5 * time.Second,
	MinTTL:                       0,
	Namespace:                    "",
	NAT64Networks:                []string{},
	NS1Endpoint:                  "",
	NS1IgnoreSSL:                 false,
	OCIConfigFile:                "/etc/kubernetes/oci.yaml",
	OCIZoneCacheDuration:         0 * time.Second,
	OCIZoneScope:                 "GLOBAL",
	Once:                         false,
	OVHApiRateLimit:              20,
	OVHEnableCNAMERelative:       false,
	OVHEndpoint:                  "ovh-eu",
	PDNSAPIKey:                   "",
	PDNSServer:                   "http://localhost:8081",
	PDNSServerID:                 "localhost",
	PDNSSkipTLSVerify:            false,
	PiholeApiVersion:             "5",
	PiholePassword:               "",
	PiholeServer:                 "",
	PiholeTLSInsecureSkipVerify:  false,
	PluralCluster:                "",
	PluralProvider:               "",
	PodSourceDomain:              "",
	Policy:                       "sync",
	Provider:                     "",
	ProviderCacheTime:            0,
	PublishHostIP:                false,
	PublishInternal:              false,
	RegexDomainExclude:           regexp.MustCompile(""),
	RegexDomainFilter:            regexp.MustCompile(""),
	Registry:                     "txt",
	RequestTimeout:               time.Second * 30,
	RFC2136BatchChangeSize:       50,
	RFC2136GSSTSIG:               false,
	RFC2136Host:                  []string{""},
	RFC2136Insecure:              false,
	RFC2136KerberosPassword:      "",
	RFC2136KerberosRealm:         "",
	RFC2136KerberosUsername:      "",
	RFC2136LoadBalancingStrategy: "disabled",
	RFC2136MinTTL:                0,
	RFC2136Port:                  0,
	RFC2136SkipTLSVerify:         false,
	RFC2136TAXFR:                 true,
	RFC2136TSIGKeyName:           "",
	RFC2136TSIGSecret:            "",
	RFC2136TSIGSecretAlg:         "",
	RFC2136UseTLS:                false,
	RFC2136Zone:                  []string{},
	ServiceTypeFilter:            []string{},
	SkipperRouteGroupVersion:     "zalando.org/v1",
	Sources:                      nil,
	TargetNetFilter:              []string{},
	TLSCA:                        "",
	TLSClientCert:                "",
	TLSClientCertKey:             "",
	TraefikEnableLegacy:          false,
	TraefikDisableNew:            false,
	TransIPAccountName:           "",
	TransIPPrivateKeyFile:        "",
	TXTCacheInterval:             0,
	TXTEncryptAESKey:             "",
	TXTEncryptEnabled:            false,
	TXTOwnerID:                   "default",
	TXTOwnerOld:                  "",
	TXTPrefix:                    "",
	TXTSuffix:                    "",
	TXTWildcardReplacement:       "",
	UpdateEvents:                 false,
	WebhookProviderReadTimeout:   5 * time.Second,
	WebhookProviderURL:           "http://localhost:8888",
	WebhookProviderWriteTimeout:  10 * time.Second,
	WebhookServer:                false,
	ZoneIDFilter:                 []string{},
	ForceDefaultTargets:          false,
	PreferAlias:                  false,
}

var providerNames = []string{
	"akamai",
	"alibabacloud",
	"aws",
	"aws-sd",
	"azure",
	"azure-dns",
	"azure-private-dns",
	"civo",
	"cloudflare",
	"coredns",
	"digitalocean",
	"dnsimple",
	"exoscale",
	"gandi",
	"godaddy",
	"google",
	"inmemory",
	"linode",
	"ns1",
	"oci",
	"ovh",
	"pdns",
	"pihole",
	"plural",
	"rfc2136",
	"scaleway",
	"skydns",
	"transip",
	"webhook",
}

var allowedSources = []string{
	"service",
	"ingress",
	"node",
	"pod",
	"gateway-httproute",
	"gateway-grpcroute",
	"gateway-tlsroute",
	"gateway-tcproute",
	"gateway-udproute",
	"istio-gateway",
	"istio-virtualservice",
	"contour-httpproxy",
	"gloo-proxy",
	"fake",
	"connector",
	"crd",
	"empty",
	"skipper-routegroup",
	"openshift-route",
	"ambassador-host",
	"kong-tcpingress",
	"f5-virtualserver",
	"f5-transportserver",
	"traefik-proxy",
}

// NewConfig returns new Config object
func NewConfig() *Config {
	return &Config{
		AnnotationPrefix: annotations.DefaultAnnotationPrefix,
		AWSSDCreateTag:   map[string]string{},
	}
}

func (cfg *Config) String() string {
	// prevent logging of sensitive information
	temp := *cfg

	t := reflect.TypeFor[Config]()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if val, ok := f.Tag.Lookup("secure"); ok && val == "yes" {
			if f.Type.Kind() != reflect.String {
				continue
			}
			v := reflect.ValueOf(&temp).Elem().Field(i)
			if v.String() != "" {
				v.SetString(passwordMask)
			}
		}
	}

	return fmt.Sprintf("%+v", temp)
}

// allLogLevelsAsStrings returns all logrus levels as a list of strings
func allLogLevelsAsStrings() []string {
	var levels []string
	for _, level := range logrus.AllLevels {
		levels = append(levels, level.String())
	}
	return levels
}

// ParseFlags adds and parses flags from command line
func (cfg *Config) ParseFlags(args []string) error {
	if _, err := App(cfg).Parse(args); err != nil {
		return err
	}
	return nil
}

func bindFlags(b flags.FlagBinder, cfg *Config) {
	// Flags related to Kubernetes
	b.StringVar("server", "The Kubernetes API server to connect to (default: auto-detect)", defaultConfig.APIServerURL, &cfg.APIServerURL)
	b.StringVar("kubeconfig", "Retrieve target cluster configuration from a Kubernetes configuration file (default: auto-detect)", defaultConfig.KubeConfig, &cfg.KubeConfig)
	b.DurationVar("request-timeout", "Request timeout when calling Kubernetes APIs. 0s means no timeout", defaultConfig.RequestTimeout, &cfg.RequestTimeout)
	b.BoolVar("resolve-service-load-balancer-hostname", "Resolve the hostname of LoadBalancer-type Service object to IP addresses in order to create DNS A/AAAA records instead of CNAMEs", false, &cfg.ResolveServiceLoadBalancerHostname)
	b.BoolVar("listen-endpoint-events", "Trigger a reconcile on changes to EndpointSlices, for Service source (default: false)", false, &cfg.ListenEndpointEvents)

	// Flags related to Gloo
	b.StringsVar("gloo-namespace", "The Gloo Proxy namespace; specify multiple times for multiple namespaces. (default: gloo-system)", []string{"gloo-system"}, &cfg.GlooNamespaces)

	// Flags related to Skipper RouteGroup
	b.StringVar("skipper-routegroup-groupversion", "The resource version for skipper routegroup", defaultConfig.SkipperRouteGroupVersion, &cfg.SkipperRouteGroupVersion)

	// Flags related to processing source
	b.BoolVar("always-publish-not-ready-addresses", "Always publish also not ready addresses for headless services (optional)", false, &cfg.AlwaysPublishNotReadyAddresses)
	b.StringVar("annotation-filter", "Filter resources queried for endpoints by annotation, using label selector semantics", defaultConfig.AnnotationFilter, &cfg.AnnotationFilter)
	b.StringVar("annotation-prefix", "Annotation prefix for external-dns annotations (default: external-dns.alpha.kubernetes.io/)", defaultConfig.AnnotationPrefix, &cfg.AnnotationPrefix)
	b.BoolVar("combine-fqdn-annotation", "Combine FQDN template and Annotations instead of overwriting (default: false)", false, &cfg.CombineFQDNAndAnnotation)
	b.EnumVar("compatibility", "Process annotation semantics from legacy implementations (optional, options: mate, molecule, kops-dns-controller)", defaultConfig.Compatibility, &cfg.Compatibility, "", "mate", "molecule", "kops-dns-controller")
	b.StringVar("connector-source-server", "The server to connect for connector source, valid only when using connector source", defaultConfig.ConnectorSourceServer, &cfg.ConnectorSourceServer)
	b.StringVar("crd-source-apiversion", "API version of the CRD for crd source, e.g. `externaldns.k8s.io/v1alpha1`, valid only when using crd source", defaultConfig.CRDSourceAPIVersion, &cfg.CRDSourceAPIVersion)
	b.StringVar("crd-source-kind", "Kind of the CRD for the crd source in API group and version specified by crd-source-apiversion", defaultConfig.CRDSourceKind, &cfg.CRDSourceKind)
	b.StringsVar("default-targets", "Set globally default host/IP that will apply as a target instead of source addresses. Specify multiple times for multiple targets (optional)", nil, &cfg.DefaultTargets)
	b.BoolVar("force-default-targets", "Force the application of --default-targets, overriding any targets provided by the source (DEPRECATED: This reverts to (improved) legacy behavior which allows empty CRD targets for migration to new state)", defaultConfig.ForceDefaultTargets, &cfg.ForceDefaultTargets)
	b.BoolVar("prefer-alias", "When enabled, CNAME records will have the alias annotation set, signaling providers that support ALIAS records to use them instead of CNAMEs. Supported by: PowerDNS, AWS (with --aws-prefer-cname disabled)", defaultConfig.PreferAlias, &cfg.PreferAlias)
	b.StringsVar("exclude-record-types", "Record types to exclude from management; specify multiple times to exclude many; (optional)", nil, &cfg.ExcludeDNSRecordTypes)
	b.StringsVar("exclude-target-net", "Exclude target nets (optional)", nil, &cfg.ExcludeTargetNets)
	b.BoolVar("exclude-unschedulable", "Exclude nodes that are considered unschedulable (default: true)", defaultConfig.ExcludeUnschedulable, &cfg.ExcludeUnschedulable)
	b.BoolVar("expose-internal-ipv6", "When using the node source, expose internal IPv6 addresses (optional, default: false)", false, &cfg.ExposeInternalIPV6)
	b.StringVar("fqdn-template", "A templated string that's used to generate DNS names from sources that don't define a hostname themselves, or to add a hostname suffix when paired with the fake source (optional). Accepts comma separated list for multiple global FQDN.", defaultConfig.FQDNTemplate, &cfg.FQDNTemplate)
	b.StringVar("gateway-label-filter", "Filter Gateways of Route endpoints via label selector (default: all gateways)", defaultConfig.GatewayLabelFilter, &cfg.GatewayLabelFilter)
	b.StringVar("gateway-name", "Limit Gateways of Route endpoints to a specific name (default: all names)", defaultConfig.GatewayName, &cfg.GatewayName)
	b.StringVar("gateway-namespace", "Limit Gateways of Route endpoints to a specific namespace (default: all namespaces)", defaultConfig.GatewayNamespace, &cfg.GatewayNamespace)
	b.BoolVar("ignore-hostname-annotation", "Ignore hostname annotation when generating DNS names, valid only when --fqdn-template is set (default: false)", false, &cfg.IgnoreHostnameAnnotation)
	b.BoolVar("ignore-ingress-rules-spec", "Ignore the spec.rules section in Ingress resources (default: false)", false, &cfg.IgnoreIngressRulesSpec)
	b.BoolVar("ignore-ingress-tls-spec", "Ignore the spec.tls section in Ingress resources (default: false)", false, &cfg.IgnoreIngressTLSSpec)
	b.BoolVar("ignore-non-host-network-pods", "Ignore pods not running on host network when using pod source (default: false)", false, &cfg.IgnoreNonHostNetworkPods)
	b.StringsVar("ingress-class", "Require an Ingress to have this class name; specify multiple times to allow more than one class (optional; defaults to any class)", nil, &cfg.IngressClassNames)
	b.StringVar("label-filter", "Filter resources queried for endpoints by label selector; currently supported by source types crd, gateway-httproute, gateway-grpcroute, gateway-tlsroute, gateway-tcproute, gateway-udproute, ingress, node, openshift-route, service and ambassador-host", defaultConfig.LabelFilter, &cfg.LabelFilter)
	managedRecordTypesHelp := fmt.Sprintf("Record types to manage; specify multiple times to include many; (default: %s) (supported records: A, AAAA, CNAME, NS, SRV, TXT)", strings.Join(defaultConfig.ManagedDNSRecordTypes, ","))
	b.StringsVar("managed-record-types", managedRecordTypesHelp, defaultConfig.ManagedDNSRecordTypes, &cfg.ManagedDNSRecordTypes)
	b.StringVar("namespace", "Limit resources queried for endpoints to a specific namespace (default: all namespaces)", defaultConfig.Namespace, &cfg.Namespace)
	b.StringsVar("nat64-networks", "Adding an A record for each AAAA record in NAT64-enabled networks; specify multiple times for multiple possible nets (optional)", nil, &cfg.NAT64Networks)
	b.StringVar("openshift-router-name", "if source is openshift-route then you can pass the ingress controller name. Based on this name external-dns will select the respective router from the route status and map that routerCanonicalHostname to the route host while creating a CNAME record.", defaultConfig.OCPRouterName, &cfg.OCPRouterName)
	b.StringVar("pod-source-domain", "Domain to use for pods records (optional)", defaultConfig.PodSourceDomain, &cfg.PodSourceDomain)
	b.BoolVar("publish-host-ip", "Allow external-dns to publish host-ip for headless services (optional)", false, &cfg.PublishHostIP)
	b.BoolVar("publish-internal-services", "Allow external-dns to publish DNS records for ClusterIP services (optional)", false, &cfg.PublishInternal)
	b.StringsVar("service-type-filter", "The service types to filter by. Specify multiple times for multiple filters to be applied. (optional, default: all, expected: ClusterIP, NodePort, LoadBalancer or ExternalName)", defaultConfig.ServiceTypeFilter, &cfg.ServiceTypeFilter)
	b.StringsVar("target-net-filter", "Limit possible targets by a net filter; specify multiple times for multiple possible nets (optional)", nil, &cfg.TargetNetFilter)
	b.BoolVar("traefik-enable-legacy", "Enable legacy listeners on Resources under the traefik.containo.us API Group", defaultConfig.TraefikEnableLegacy, &cfg.TraefikEnableLegacy)
	b.BoolVar("traefik-disable-new", "Disable listeners on Resources under the traefik.io API Group", defaultConfig.TraefikDisableNew, &cfg.TraefikDisableNew)

	b.StringsVar("events-emit", "Events that should be emitted. Specify multiple times for multiple events support (optional, default: none, expected: RecordReady, RecordDeleted, RecordError)", defaultConfig.EmitEvents, &cfg.EmitEvents)

	b.DurationVar("provider-cache-time", "The time to cache the DNS provider record list requests.", defaultConfig.ProviderCacheTime, &cfg.ProviderCacheTime)
	b.StringsVar("domain-filter", "Limit possible target zones by a domain suffix; specify multiple times for multiple domains (optional)", []string{""}, &cfg.DomainFilter)
	b.StringsVar("exclude-domains", "Exclude subdomains (optional)", []string{""}, &cfg.DomainExclude)
	b.RegexpVar("regex-domain-filter", "Limit possible domains and target zones by a Regex filter; Overrides domain-filter (optional)", defaultConfig.RegexDomainFilter, &cfg.RegexDomainFilter)
	b.RegexpVar("regex-domain-exclusion", "Regex filter that excludes domains and target zones matched by regex-domain-filter (optional)", defaultConfig.RegexDomainExclude, &cfg.RegexDomainExclude)
	b.StringsVar("zone-name-filter", "Filter target zones by zone domain (For now, only AzureDNS provider is using this flag); specify multiple times for multiple zones (optional)", []string{""}, &cfg.ZoneNameFilter)
	b.StringsVar("zone-id-filter", "Filter target zones by hosted zone id; specify multiple times for multiple zones (optional)", []string{""}, &cfg.ZoneIDFilter)
	b.StringVar("google-project", "When using the Google provider, current project is auto-detected, when running on GCP. Specify other project with this. Must be specified when running outside GCP.", defaultConfig.GoogleProject, &cfg.GoogleProject)
	b.IntVar("google-batch-change-size", "When using the Google provider, set the maximum number of changes that will be applied in each batch.", defaultConfig.GoogleBatchChangeSize, &cfg.GoogleBatchChangeSize)
	b.DurationVar("google-batch-change-interval", "When using the Google provider, set the interval between batch changes.", defaultConfig.GoogleBatchChangeInterval, &cfg.GoogleBatchChangeInterval)
	b.EnumVar("google-zone-visibility", "When using the Google provider, filter for zones with this visibility (optional, options: public, private)", defaultConfig.GoogleZoneVisibility, &cfg.GoogleZoneVisibility, "", "public", "private")
	b.StringVar("alibaba-cloud-config-file", "When using the Alibaba Cloud provider, specify the Alibaba Cloud configuration file (required when --provider=alibabacloud)", defaultConfig.AlibabaCloudConfigFile, &cfg.AlibabaCloudConfigFile)
	b.EnumVar("alibaba-cloud-zone-type", "When using the Alibaba Cloud provider, filter for zones of this type (optional, options: public, private)", defaultConfig.AlibabaCloudZoneType, &cfg.AlibabaCloudZoneType, "", "public", "private")
	b.EnumVar("aws-zone-type", "When using the AWS provider, filter for zones of this type (optional, default: any, options: public, private)", defaultConfig.AWSZoneType, &cfg.AWSZoneType, "", "public", "private")
	b.StringsVar("aws-zone-tags", "When using the AWS provider, filter for zones with these tags", []string{""}, &cfg.AWSZoneTagFilter)
	b.StringsVar("aws-profile", "When using the AWS provider, name of the profile to use", []string{""}, &cfg.AWSProfiles)
	b.StringVar("aws-assume-role", "When using the AWS API, assume this IAM role. Useful for hosted zones in another AWS account. Specify the full ARN, e.g. `arn:aws:iam::123455567:role/external-dns` (optional)", defaultConfig.AWSAssumeRole, &cfg.AWSAssumeRole)
	b.StringVar("aws-assume-role-external-id", "When using the AWS API and assuming a role then specify this external ID` (optional)", defaultConfig.AWSAssumeRoleExternalID, &cfg.AWSAssumeRoleExternalID)
	b.IntVar("aws-batch-change-size", "When using the AWS provider, set the maximum number of changes that will be applied in each batch.", defaultConfig.AWSBatchChangeSize, &cfg.AWSBatchChangeSize)
	b.IntVar("aws-batch-change-size-bytes", "When using the AWS provider, set the maximum byte size that will be applied in each batch.", defaultConfig.AWSBatchChangeSizeBytes, &cfg.AWSBatchChangeSizeBytes)
	b.IntVar("aws-batch-change-size-values", "When using the AWS provider, set the maximum total record values that will be applied in each batch.", defaultConfig.AWSBatchChangeSizeValues, &cfg.AWSBatchChangeSizeValues)
	b.DurationVar("aws-batch-change-interval", "When using the AWS provider, set the interval between batch changes.", defaultConfig.AWSBatchChangeInterval, &cfg.AWSBatchChangeInterval)
	b.BoolVar("aws-evaluate-target-health", "When using the AWS provider, set whether to evaluate the health of a DNS target (default: enabled, disable with --no-aws-evaluate-target-health)", defaultConfig.AWSEvaluateTargetHealth, &cfg.AWSEvaluateTargetHealth)
	b.IntVar("aws-api-retries", "When using the AWS API, set the maximum number of retries before giving up.", defaultConfig.AWSAPIRetries, &cfg.AWSAPIRetries)
	b.BoolVar("aws-prefer-cname", "When using the AWS provider, prefer using CNAME instead of ALIAS (default: disabled)", defaultConfig.AWSPreferCNAME, &cfg.AWSPreferCNAME)
	b.DurationVar("aws-zones-cache-duration", "When using the AWS provider, set the zones list cache TTL (0s to disable).", defaultConfig.AWSZoneCacheDuration, &cfg.AWSZoneCacheDuration)
	b.BoolVar("aws-zone-match-parent", "Expand limit possible target by sub-domains (default: disabled)", defaultConfig.AWSZoneMatchParent, &cfg.AWSZoneMatchParent)
	b.BoolVar("aws-sd-service-cleanup", "When using the AWS CloudMap provider, delete empty Services without endpoints (default: disabled)", defaultConfig.AWSSDServiceCleanup, &cfg.AWSSDServiceCleanup)
	b.StringMapVar("aws-sd-create-tag", "When using the AWS CloudMap provider, add tag to created services. The flag can be used multiple times", &cfg.AWSSDCreateTag)
	b.StringVar("azure-config-file", "When using the Azure provider, specify the Azure configuration file (required when --provider=azure)", defaultConfig.AzureConfigFile, &cfg.AzureConfigFile)
	b.StringVar("azure-resource-group", "When using the Azure provider, override the Azure resource group to use (optional)", defaultConfig.AzureResourceGroup, &cfg.AzureResourceGroup)
	b.StringVar("azure-subscription-id", "When using the Azure provider, override the Azure subscription to use (optional)", defaultConfig.AzureSubscriptionID, &cfg.AzureSubscriptionID)
	b.StringVar("azure-user-assigned-identity-client-id", "When using the Azure provider, override the client id of user assigned identity in config file (optional)", "", &cfg.AzureUserAssignedIdentityClientID)
	b.DurationVar("azure-zones-cache-duration", "When using the Azure provider, set the zones list cache TTL (0s to disable).", defaultConfig.AzureZonesCacheDuration, &cfg.AzureZonesCacheDuration)
	b.IntVar("azure-maxretries-count", "When using the Azure provider, set the number of retries for API calls (When less than 0, it disables retries). (optional)", defaultConfig.AzureMaxRetriesCount, &cfg.AzureMaxRetriesCount)

	b.BoolVar("cloudflare-proxied", "When using the Cloudflare provider, specify if the proxy mode must be enabled (default: disabled)", false, &cfg.CloudflareProxied)
	b.BoolVar("cloudflare-custom-hostnames", "When using the Cloudflare provider, specify if the Custom Hostnames feature will be used. Requires \"Cloudflare for SaaS\" enabled. (default: disabled)", false, &cfg.CloudflareCustomHostnames)
	b.EnumVar("cloudflare-custom-hostnames-min-tls-version", "When using the Cloudflare provider with the Custom Hostnames, specify which Minimum TLS Version will be used by default. (default: 1.0, options: 1.0, 1.1, 1.2, 1.3)", "1.0", &cfg.CloudflareCustomHostnamesMinTLSVersion, "1.0", "1.1", "1.2", "1.3")
	b.EnumVar("cloudflare-custom-hostnames-certificate-authority", "When using the Cloudflare provider with the Custom Hostnames, specify which Certificate Authority will be used. A value of none indicates no Certificate Authority will be sent to the Cloudflare API (default: none, options: google, ssl_com, lets_encrypt, none)", "none", &cfg.CloudflareCustomHostnamesCertificateAuthority, "google", "ssl_com", "lets_encrypt", "none")
	b.IntVar("cloudflare-dns-records-per-page", "When using the Cloudflare provider, specify how many DNS records listed per page, max possible 5,000 (default: 100)", defaultConfig.CloudflareDNSRecordsPerPage, &cfg.CloudflareDNSRecordsPerPage)
	b.BoolVar("cloudflare-regional-services", "When using the Cloudflare provider, specify if Regional Services feature will be used (default: disabled)", defaultConfig.CloudflareRegionalServices, &cfg.CloudflareRegionalServices)
	b.StringVar("cloudflare-region-key", "When using the Cloudflare provider, specify the default region for Regional Services. Any value other than an empty string will enable the Regional Services feature (optional)", "", &cfg.CloudflareRegionKey)
	b.StringVar("cloudflare-record-comment", "When using the Cloudflare provider, specify the comment for the DNS records (default: '')", "", &cfg.CloudflareDNSRecordsComment)

	b.StringVar("coredns-prefix", "When using the CoreDNS provider, specify the prefix name", defaultConfig.CoreDNSPrefix, &cfg.CoreDNSPrefix)
	b.BoolVar("coredns-strictly-owned", "When using the CoreDNS provider, store and filter strictly by txt-owner-id using an extra field inside of the etcd service (default: false)", defaultConfig.CoreDNSStrictlyOwned, &cfg.CoreDNSStrictlyOwned)
	b.StringVar("akamai-serviceconsumerdomain", "When using the Akamai provider, specify the base URL (required when --provider=akamai and edgerc-path not specified)", defaultConfig.AkamaiServiceConsumerDomain, &cfg.AkamaiServiceConsumerDomain)
	b.StringVar("akamai-client-token", "When using the Akamai provider, specify the client token (required when --provider=akamai and edgerc-path not specified)", defaultConfig.AkamaiClientToken, &cfg.AkamaiClientToken)
	b.StringVar("akamai-client-secret", "When using the Akamai provider, specify the client secret (required when --provider=akamai and edgerc-path not specified)", defaultConfig.AkamaiClientSecret, &cfg.AkamaiClientSecret)
	b.StringVar("akamai-access-token", "When using the Akamai provider, specify the access token (required when --provider=akamai and edgerc-path not specified)", defaultConfig.AkamaiAccessToken, &cfg.AkamaiAccessToken)
	b.StringVar("akamai-edgerc-path", "When using the Akamai provider, specify the .edgerc file path. Path must be reachable form invocation environment. (required when --provider=akamai and *-token, secret serviceconsumerdomain not specified)", defaultConfig.AkamaiEdgercPath, &cfg.AkamaiEdgercPath)
	b.StringVar("akamai-edgerc-section", "When using the Akamai provider, specify the .edgerc file path (Optional when edgerc-path is specified)", defaultConfig.AkamaiEdgercSection, &cfg.AkamaiEdgercSection)
	b.StringVar("oci-config-file", "When using the OCI provider, specify the OCI configuration file (required when --provider=oci", defaultConfig.OCIConfigFile, &cfg.OCIConfigFile)
	b.StringVar("oci-compartment-ocid", "When using the OCI provider, specify the OCID of the OCI compartment containing all managed zones and records.  Required when using OCI IAM instance principal authentication.", defaultConfig.OCICompartmentOCID, &cfg.OCICompartmentOCID)
	b.EnumVar("oci-zone-scope", "When using OCI provider, filter for zones with this scope (optional, options: GLOBAL, PRIVATE). Defaults to GLOBAL, setting to empty value will target both.", defaultConfig.OCIZoneScope, &cfg.OCIZoneScope, "", "GLOBAL", "PRIVATE")
	b.BoolVar("oci-auth-instance-principal", "When using the OCI provider, specify whether OCI IAM instance principal authentication should be used (instead of key-based auth via the OCI config file).", defaultConfig.OCIAuthInstancePrincipal, &cfg.OCIAuthInstancePrincipal)
	b.DurationVar("oci-zones-cache-duration", "When using the OCI provider, set the zones list cache TTL (0s to disable).", defaultConfig.OCIZoneCacheDuration, &cfg.OCIZoneCacheDuration)
	b.StringsVar("inmemory-zone", "Provide a list of pre-configured zones for the inmemory provider; specify multiple times for multiple zones (optional)", []string{""}, &cfg.InMemoryZones)
	b.StringVar("ovh-endpoint", "When using the OVH provider, specify the endpoint (default: ovh-eu)", defaultConfig.OVHEndpoint, &cfg.OVHEndpoint)
	b.IntVar("ovh-api-rate-limit", "When using the OVH provider, specify the API request rate limit, X operations by seconds (default: 20)", defaultConfig.OVHApiRateLimit, &cfg.OVHApiRateLimit)
	b.BoolVar("ovh-enable-cname-relative", "When using the OVH provider, specify if CNAME should be treated as relative on target without final dot (default: false)", defaultConfig.OVHEnableCNAMERelative, &cfg.OVHEnableCNAMERelative)
	b.StringVar("pdns-server", "When using the PowerDNS/PDNS provider, specify the URL to the pdns server (required when --provider=pdns)", defaultConfig.PDNSServer, &cfg.PDNSServer)
	b.StringVar("pdns-server-id", "When using the PowerDNS/PDNS provider, specify the id of the server to retrieve. Should be `localhost` except when the server is behind a proxy (optional when --provider=pdns) (default: localhost)", defaultConfig.PDNSServerID, &cfg.PDNSServerID)
	b.StringVar("pdns-api-key", "When using the PowerDNS/PDNS provider, specify the API key to use to authorize requests (required when --provider=pdns)", defaultConfig.PDNSAPIKey, &cfg.PDNSAPIKey)
	b.BoolVar("pdns-skip-tls-verify", "When using the PowerDNS/PDNS provider, disable verification of any TLS certificates (optional when --provider=pdns) (default: false)", defaultConfig.PDNSSkipTLSVerify, &cfg.PDNSSkipTLSVerify)
	b.StringVar("ns1-endpoint", "When using the NS1 provider, specify the URL of the API endpoint to target (default: https://api.nsone.net/v1/)", defaultConfig.NS1Endpoint, &cfg.NS1Endpoint)
	b.BoolVar("ns1-ignoressl", "When using the NS1 provider, specify whether to verify the SSL certificate (default: false)", defaultConfig.NS1IgnoreSSL, &cfg.NS1IgnoreSSL)
	b.IntVar("ns1-min-ttl", "Minimal TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is lower than this.", cfg.NS1MinTTLSeconds, &cfg.NS1MinTTLSeconds)
	b.IntVar("digitalocean-api-page-size", "Configure the page size used when querying the DigitalOcean API.", defaultConfig.DigitalOceanAPIPageSize, &cfg.DigitalOceanAPIPageSize)
	// GoDaddy flags
	b.StringVar("godaddy-api-key", "When using the GoDaddy provider, specify the API Key (required when --provider=godaddy)", defaultConfig.GoDaddyAPIKey, &cfg.GoDaddyAPIKey)
	b.StringVar("godaddy-api-secret", "When using the GoDaddy provider, specify the API secret (required when --provider=godaddy)", defaultConfig.GoDaddySecretKey, &cfg.GoDaddySecretKey)
	b.Int64Var("godaddy-api-ttl", "TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is not provided.", cfg.GoDaddyTTL, &cfg.GoDaddyTTL)
	b.BoolVar("godaddy-api-ote", "When using the GoDaddy provider, use OTE api (optional, default: false, when --provider=godaddy)", defaultConfig.GoDaddyOTE, &cfg.GoDaddyOTE)

	// Flags related to TLS communication
	b.StringVar("tls-ca", "When using TLS communication, the path to the certificate authority to verify server communications (optionally specify --tls-client-cert for two-way TLS)", defaultConfig.TLSCA, &cfg.TLSCA)
	b.StringVar("tls-client-cert", "When using TLS communication, the path to the certificate to present as a client (not required for TLS)", defaultConfig.TLSClientCert, &cfg.TLSClientCert)
	b.StringVar("tls-client-cert-key", "When using TLS communication, the path to the certificate key to use with the client certificate (not required for TLS)", defaultConfig.TLSClientCertKey, &cfg.TLSClientCertKey)

	// Flags related to Exoscale provider
	b.StringVar("exoscale-apienv", "When using Exoscale provider, specify the API environment (optional)", defaultConfig.ExoscaleAPIEnvironment, &cfg.ExoscaleAPIEnvironment)
	b.StringVar("exoscale-apizone", "When using Exoscale provider, specify the API Zone (optional)", defaultConfig.ExoscaleAPIZone, &cfg.ExoscaleAPIZone)
	b.StringVar("exoscale-apikey", "Provide your API Key for the Exoscale provider", defaultConfig.ExoscaleAPIKey, &cfg.ExoscaleAPIKey)
	b.StringVar("exoscale-apisecret", "Provide your API Secret for the Exoscale provider", defaultConfig.ExoscaleAPISecret, &cfg.ExoscaleAPISecret)

	// Flags related to RFC2136 provider
	b.StringsVar("rfc2136-host", "When using the RFC2136 provider, specify the host of the DNS server (optionally specify multiple times when using --rfc2136-load-balancing-strategy)", []string{defaultConfig.RFC2136Host[0]}, &cfg.RFC2136Host)
	b.IntVar("rfc2136-port", "When using the RFC2136 provider, specify the port of the DNS server", defaultConfig.RFC2136Port, &cfg.RFC2136Port)
	b.StringsVar("rfc2136-zone", "When using the RFC2136 provider, specify zone entry of the DNS server to use (can be specified multiple times)", nil, &cfg.RFC2136Zone)
	b.BoolVar("rfc2136-create-ptr", "When using the RFC2136 provider, enable PTR management", defaultConfig.RFC2136CreatePTR, &cfg.RFC2136CreatePTR)
	b.BoolVar("rfc2136-insecure", "When using the RFC2136 provider, specify whether to attach TSIG or not (default: false, requires --rfc2136-tsig-keyname and rfc2136-tsig-secret)", defaultConfig.RFC2136Insecure, &cfg.RFC2136Insecure)
	b.StringVar("rfc2136-tsig-keyname", "When using the RFC2136 provider, specify the TSIG key to attached to DNS messages (required when --rfc2136-insecure=false)", defaultConfig.RFC2136TSIGKeyName, &cfg.RFC2136TSIGKeyName)
	b.StringVar("rfc2136-tsig-secret", "When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)", defaultConfig.RFC2136TSIGSecret, &cfg.RFC2136TSIGSecret)
	b.StringVar("rfc2136-tsig-secret-alg", "When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)", defaultConfig.RFC2136TSIGSecretAlg, &cfg.RFC2136TSIGSecretAlg)
	b.BoolVar("rfc2136-tsig-axfr", "When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)", false, &cfg.RFC2136TAXFR)
	b.DurationVar("rfc2136-min-ttl", "When using the RFC2136 provider, specify minimal TTL (in duration format) for records. This value will be used if the provided TTL for a service/ingress is lower than this", defaultConfig.RFC2136MinTTL, &cfg.RFC2136MinTTL)
	b.BoolVar("rfc2136-gss-tsig", "When using the RFC2136 provider, specify whether to use secure updates with GSS-TSIG using Kerberos (default: false, requires --rfc2136-kerberos-realm, --rfc2136-kerberos-username, and rfc2136-kerberos-password)", defaultConfig.RFC2136GSSTSIG, &cfg.RFC2136GSSTSIG)
	b.StringVar("rfc2136-kerberos-username", "When using the RFC2136 provider with GSS-TSIG, specify the username of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)", defaultConfig.RFC2136KerberosUsername, &cfg.RFC2136KerberosUsername)
	b.StringVar("rfc2136-kerberos-password", "When using the RFC2136 provider with GSS-TSIG, specify the password of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)", defaultConfig.RFC2136KerberosPassword, &cfg.RFC2136KerberosPassword)
	b.StringVar("rfc2136-kerberos-realm", "When using the RFC2136 provider with GSS-TSIG, specify the realm of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)", defaultConfig.RFC2136KerberosRealm, &cfg.RFC2136KerberosRealm)
	b.IntVar("rfc2136-batch-change-size", "When using the RFC2136 provider, set the maximum number of changes that will be applied in each batch.", defaultConfig.RFC2136BatchChangeSize, &cfg.RFC2136BatchChangeSize)
	b.BoolVar("rfc2136-use-tls", "When using the RFC2136 provider, communicate with name server over tls", defaultConfig.RFC2136UseTLS, &cfg.RFC2136UseTLS)
	b.BoolVar("rfc2136-skip-tls-verify", "When using TLS with the RFC2136 provider, disable verification of any TLS certificates", defaultConfig.RFC2136SkipTLSVerify, &cfg.RFC2136SkipTLSVerify)
	b.EnumVar("rfc2136-load-balancing-strategy", "When using the RFC2136 provider, specify the load balancing strategy (default: disabled, options: random, round-robin, disabled)", defaultConfig.RFC2136LoadBalancingStrategy, &cfg.RFC2136LoadBalancingStrategy, "random", "round-robin", "disabled")

	// Flags related to TransIP provider
	b.StringVar("transip-account", "When using the TransIP provider, specify the account name (required when --provider=transip)", defaultConfig.TransIPAccountName, &cfg.TransIPAccountName)
	b.StringVar("transip-keyfile", "When using the TransIP provider, specify the path to the private key file (required when --provider=transip)", defaultConfig.TransIPPrivateKeyFile, &cfg.TransIPPrivateKeyFile)

	// Flags related to Pihole provider
	b.StringVar("pihole-server", "When using the Pihole provider, the base URL of the Pihole web server (required when --provider=pihole)", defaultConfig.PiholeServer, &cfg.PiholeServer)
	b.StringVar("pihole-password", "When using the Pihole provider, the password to the server if it is protected", defaultConfig.PiholePassword, &cfg.PiholePassword)
	b.BoolVar("pihole-tls-skip-verify", "When using the Pihole provider, disable verification of any TLS certificates", defaultConfig.PiholeTLSInsecureSkipVerify, &cfg.PiholeTLSInsecureSkipVerify)
	b.StringVar("pihole-api-version", "When using the Pihole provider, specify the pihole API version (default: 5, options: 5, 6)", defaultConfig.PiholeApiVersion, &cfg.PiholeApiVersion)

	// Flags related to the Plural provider
	b.StringVar("plural-cluster", "When using the plural provider, specify the cluster name you're running with", defaultConfig.PluralCluster, &cfg.PluralCluster)
	b.StringVar("plural-provider", "When using the plural provider, specify the provider name you're running with", defaultConfig.PluralProvider, &cfg.PluralProvider)

	// Flags related to policies
	b.EnumVar("policy", "Modify how DNS records are synchronized between sources and providers (default: sync, options: sync, upsert-only, create-only)", defaultConfig.Policy, &cfg.Policy, "sync", "upsert-only", "create-only")

	// Flags related to the registry
	b.EnumVar("registry", "The registry implementation to use to keep track of DNS record ownership (default: txt, options: txt, noop, dynamodb, aws-sd)", defaultConfig.Registry, &cfg.Registry, "txt", "noop", "dynamodb", "aws-sd")
	b.StringVar("txt-owner-id", "When using the TXT or DynamoDB registry, a name that identifies this instance of ExternalDNS (default: default)", defaultConfig.TXTOwnerID, &cfg.TXTOwnerID)
	b.StringVar("txt-prefix", "When using the TXT registry, a custom string that's prefixed to each ownership DNS record (optional). Could contain record type template like '%{record_type}-prefix-'. Mutual exclusive with txt-suffix!", defaultConfig.TXTPrefix, &cfg.TXTPrefix)
	b.StringVar("txt-suffix", "When using the TXT registry, a custom string that's suffixed to the host portion of each ownership DNS record (optional). Could contain record type template like '-%{record_type}-suffix'. Mutual exclusive with txt-prefix!", defaultConfig.TXTSuffix, &cfg.TXTSuffix)
	b.StringVar("txt-wildcard-replacement", "When using the TXT registry, a custom string that's used instead of an asterisk for TXT records corresponding to wildcard DNS records (optional)", defaultConfig.TXTWildcardReplacement, &cfg.TXTWildcardReplacement)
	b.BoolVar("txt-encrypt-enabled", "When using the TXT registry, set if TXT records should be encrypted before stored (default: disabled)", defaultConfig.TXTEncryptEnabled, &cfg.TXTEncryptEnabled)
	b.StringVar("txt-encrypt-aes-key", "When using the TXT registry, set TXT record decryption and encryption 32 byte aes key (required when --txt-encrypt=true)", defaultConfig.TXTEncryptAESKey, &cfg.TXTEncryptAESKey)
	b.StringVar("migrate-from-txt-owner", "Old txt-owner-id that needs to be overwritten (default: default)", defaultConfig.TXTOwnerOld, &cfg.TXTOwnerOld)
	b.StringVar("dynamodb-region", "When using the DynamoDB registry, the AWS region of the DynamoDB table (optional)", cfg.AWSDynamoDBRegion, &cfg.AWSDynamoDBRegion)
	b.StringVar("dynamodb-table", "When using the DynamoDB registry, the name of the DynamoDB table (default: \"external-dns\")", defaultConfig.AWSDynamoDBTable, &cfg.AWSDynamoDBTable)

	// Flags related to the main control loop
	b.DurationVar("txt-cache-interval", "The interval between cache synchronizations in duration format (default: disabled)", defaultConfig.TXTCacheInterval, &cfg.TXTCacheInterval)
	b.DurationVar("interval", "The interval between two consecutive synchronizations in duration format (default: 1m)", defaultConfig.Interval, &cfg.Interval)
	b.DurationVar("min-event-sync-interval", "The minimum interval between two consecutive synchronizations triggered from kubernetes events in duration format (default: 5s)", defaultConfig.MinEventSyncInterval, &cfg.MinEventSyncInterval)
	b.BoolVar("once", "When enabled, exits the synchronization loop after the first iteration (default: disabled)", defaultConfig.Once, &cfg.Once)
	b.BoolVar("dry-run", "When enabled, prints DNS record changes rather than actually performing them (default: disabled)", defaultConfig.DryRun, &cfg.DryRun)
	b.BoolVar("events", "When enabled, in addition to running every interval, the reconciliation loop will get triggered when supported sources change (default: disabled)", defaultConfig.UpdateEvents, &cfg.UpdateEvents)
	b.DurationVar("min-ttl", "Configure global TTL for records in duration format. This value is used when the TTL for a source is not set or set to 0. (optional; examples: 1m12s, 72s, 72)", defaultConfig.MinTTL, &cfg.MinTTL)

	// Miscellaneous flags
	b.EnumVar("log-format", "The format in which log messages are printed (default: text, options: text, json)", defaultConfig.LogFormat, &cfg.LogFormat, "text", "json")
	b.StringVar("metrics-address", "Specify where to serve the metrics and health check endpoint (default: :7979)", defaultConfig.MetricsAddress, &cfg.MetricsAddress)
	b.EnumVar("log-level", "Set the level of logging. (default: info, options: panic, debug, info, warning, error, fatal)", defaultConfig.LogLevel, &cfg.LogLevel, allLogLevelsAsStrings()...)

	// Webhook provider
	b.StringVar("webhook-provider-url", "The URL of the remote endpoint to call for the webhook provider (default: http://localhost:8888)", defaultConfig.WebhookProviderURL, &cfg.WebhookProviderURL)
	b.DurationVar("webhook-provider-read-timeout", "The read timeout for the webhook provider in duration format (default: 5s)", defaultConfig.WebhookProviderReadTimeout, &cfg.WebhookProviderReadTimeout)
	b.DurationVar("webhook-provider-write-timeout", "The write timeout for the webhook provider in duration format (default: 10s)", defaultConfig.WebhookProviderWriteTimeout, &cfg.WebhookProviderWriteTimeout)
	b.BoolVar("webhook-server", "When enabled, runs as a webhook server instead of a controller. (default: false).", defaultConfig.WebhookServer, &cfg.WebhookServer)
}

func App(cfg *Config) *kingpin.Application {
	app := kingpin.New("external-dns", "ExternalDNS synchronizes exposed Kubernetes Services and Ingresses with DNS providers.\n\nNote that all flags may be replaced with env vars - `--flag` -> `EXTERNAL_DNS_FLAG=1` or `--flag value` -> `EXTERNAL_DNS_FLAG=value`")
	app.Version(Version)
	app.DefaultEnvars()

	bindFlags(flags.NewKingpinBinder(app), cfg)

	// Kingpin-only semantics: preserve Required/PlaceHolder and enum validation
	// that Kingpin provided before the flags were migrated into the binder.
	providerHelp := "The DNS provider where the DNS records will be created (required, options: " + strings.Join(providerNames, ", ") + ")"
	app.Flag("provider", providerHelp).Required().PlaceHolder("provider").EnumVar(&cfg.Provider, providerNames...)

	// Reintroduce source enum/required validation in Kingpin to match previous behavior.
	sourceHelp := "The resource types that are queried for endpoints; specify multiple times for multiple sources (required, options: " + strings.Join(allowedSources, ", ") + ")"
	app.Flag("source", sourceHelp).Required().PlaceHolder("source").EnumsVar(&cfg.Sources, allowedSources...)

	return app
}
