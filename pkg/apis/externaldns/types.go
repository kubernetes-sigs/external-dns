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
	"strconv"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/endpoint"

	"github.com/alecthomas/kingpin/v2"
	"github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/source"
)

const (
	passwordMask = "******"
)

// Version is the current version of the app, generated at build time
var Version = "unknown"

// Config is a project-wide configuration
type Config struct {
	APIServerURL                       string
	KubeConfig                         string
	RequestTimeout                     time.Duration
	DefaultTargets                     []string
	GlooNamespaces                     []string
	SkipperRouteGroupVersion           string
	Sources                            []string
	Namespace                          string
	AnnotationFilter                   string
	LabelFilter                        string
	IngressClassNames                  []string
	FQDNTemplate                       string
	CombineFQDNAndAnnotation           bool
	IgnoreHostnameAnnotation           bool
	IgnoreNonHostNetworkPods           bool
	IgnoreIngressTLSSpec               bool
	IgnoreIngressRulesSpec             bool
	GatewayNamespace                   string
	GatewayLabelFilter                 string
	Compatibility                      string
	PodSourceDomain                    string
	PublishInternal                    bool
	PublishHostIP                      bool
	AlwaysPublishNotReadyAddresses     bool
	ConnectorSourceServer              string
	Provider                           string
	ProviderCacheTime                  time.Duration
	GoogleProject                      string
	GoogleBatchChangeSize              int
	GoogleBatchChangeInterval          time.Duration
	GoogleZoneVisibility               string
	DomainFilter                       []string
	ExcludeDomains                     []string
	RegexDomainFilter                  *regexp.Regexp
	RegexDomainExclusion               *regexp.Regexp
	ZoneNameFilter                     []string
	ZoneIDFilter                       []string
	TargetNetFilter                    []string
	ExcludeTargetNets                  []string
	AlibabaCloudConfigFile             string
	AlibabaCloudZoneType               string
	AWSZoneType                        string
	AWSZoneTagFilter                   []string
	AWSAssumeRole                      string
	AWSProfiles                        []string
	AWSAssumeRoleExternalID            string `secure:"yes"`
	AWSBatchChangeSize                 int
	AWSBatchChangeSizeBytes            int
	AWSBatchChangeSizeValues           int
	AWSBatchChangeInterval             time.Duration
	AWSEvaluateTargetHealth            bool
	AWSAPIRetries                      int
	AWSPreferCNAME                     bool
	AWSZoneCacheDuration               time.Duration
	AWSSDServiceCleanup                bool
	AWSSDCreateTag                     map[string]string
	AWSZoneMatchParent                 bool
	AWSDynamoDBRegion                  string
	AWSDynamoDBTable                   string
	AzureConfigFile                    string
	AzureResourceGroup                 string
	AzureSubscriptionID                string
	AzureUserAssignedIdentityClientID  string
	AzureActiveDirectoryAuthorityHost  string
	AzureZonesCacheDuration            time.Duration
	CloudflareProxied                  bool
	CloudflareDNSRecordsPerPage        int
	CloudflareRegionKey                string
	CoreDNSPrefix                      string
	AkamaiServiceConsumerDomain        string
	AkamaiClientToken                  string
	AkamaiClientSecret                 string
	AkamaiAccessToken                  string
	AkamaiEdgercPath                   string
	AkamaiEdgercSection                string
	OCIConfigFile                      string
	OCICompartmentOCID                 string
	OCIAuthInstancePrincipal           bool
	OCIZoneScope                       string
	OCIZoneCacheDuration               time.Duration
	InMemoryZones                      []string
	OVHEndpoint                        string
	OVHApiRateLimit                    int
	PDNSServer                         string
	PDNSServerID                       string
	PDNSAPIKey                         string `secure:"yes"`
	PDNSSkipTLSVerify                  bool
	TLSCA                              string
	TLSClientCert                      string
	TLSClientCertKey                   string
	Policy                             string
	Registry                           string
	TXTOwnerID                         string
	TXTPrefix                          string
	TXTSuffix                          string
	TXTEncryptEnabled                  bool
	TXTEncryptAESKey                   string `secure:"yes"`
	TXTNewFormatOnly                   bool
	Interval                           time.Duration
	MinEventSyncInterval               time.Duration
	Once                               bool
	DryRun                             bool
	UpdateEvents                       bool
	LogFormat                          string
	MetricsAddress                     string
	LogLevel                           string
	TXTCacheInterval                   time.Duration
	TXTWildcardReplacement             string
	ExoscaleEndpoint                   string
	ExoscaleAPIKey                     string `secure:"yes"`
	ExoscaleAPISecret                  string `secure:"yes"`
	ExoscaleAPIEnvironment             string
	ExoscaleAPIZone                    string
	CRDSourceAPIVersion                string
	CRDSourceKind                      string
	ServiceTypeFilter                  []string
	CFAPIEndpoint                      string
	CFUsername                         string
	CFPassword                         string
	ResolveServiceLoadBalancerHostname bool
	RFC2136Host                        []string
	RFC2136Port                        int
	RFC2136Zone                        []string
	RFC2136Insecure                    bool
	RFC2136GSSTSIG                     bool
	RFC2136CreatePTR                   bool
	RFC2136KerberosRealm               string
	RFC2136KerberosUsername            string
	RFC2136KerberosPassword            string `secure:"yes"`
	RFC2136TSIGKeyName                 string
	RFC2136TSIGSecret                  string `secure:"yes"`
	RFC2136TSIGSecretAlg               string
	RFC2136TAXFR                       bool
	RFC2136MinTTL                      time.Duration
	RFC2136LoadBalancingStrategy       string
	RFC2136BatchChangeSize             int
	RFC2136UseTLS                      bool
	RFC2136SkipTLSVerify               bool
	NS1Endpoint                        string
	NS1IgnoreSSL                       bool
	NS1MinTTLSeconds                   int
	TransIPAccountName                 string
	TransIPPrivateKeyFile              string
	DigitalOceanAPIPageSize            int
	ManagedDNSRecordTypes              []string
	ExcludeDNSRecordTypes              []string
	GoDaddyAPIKey                      string `secure:"yes"`
	GoDaddySecretKey                   string `secure:"yes"`
	GoDaddyTTL                         int64
	GoDaddyOTE                         bool
	OCPRouterName                      string
	IBMCloudProxied                    bool
	IBMCloudConfigFile                 string
	TencentCloudConfigFile             string
	TencentCloudZoneType               string
	PiholeServer                       string
	PiholePassword                     string `secure:"yes"`
	PiholeTLSInsecureSkipVerify        bool
	PluralCluster                      string
	PluralProvider                     string
	WebhookProviderURL                 string
	WebhookProviderReadTimeout         time.Duration
	WebhookProviderWriteTimeout        time.Duration
	WebhookServer                      bool
	TraefikDisableLegacy               bool
	TraefikDisableNew                  bool
	NAT64Networks                      []string
}

var defaultConfig = &Config{
	APIServerURL:                 "",
	KubeConfig:                   "",
	RequestTimeout:               time.Second * 30,
	DefaultTargets:               []string{},
	GlooNamespaces:               []string{"gloo-system"},
	SkipperRouteGroupVersion:     "zalando.org/v1",
	Sources:                      nil,
	Namespace:                    "",
	AnnotationFilter:             "",
	LabelFilter:                  labels.Everything().String(),
	IngressClassNames:            nil,
	FQDNTemplate:                 "",
	CombineFQDNAndAnnotation:     false,
	IgnoreHostnameAnnotation:     false,
	IgnoreIngressTLSSpec:         false,
	IgnoreIngressRulesSpec:       false,
	GatewayNamespace:             "",
	GatewayLabelFilter:           "",
	Compatibility:                "",
	PublishInternal:              false,
	PublishHostIP:                false,
	ConnectorSourceServer:        "localhost:8080",
	Provider:                     "",
	ProviderCacheTime:            0,
	GoogleProject:                "",
	GoogleBatchChangeSize:        1000,
	GoogleBatchChangeInterval:    time.Second,
	GoogleZoneVisibility:         "",
	DomainFilter:                 []string{},
	ZoneIDFilter:                 []string{},
	ExcludeDomains:               []string{},
	RegexDomainFilter:            regexp.MustCompile(""),
	RegexDomainExclusion:         regexp.MustCompile(""),
	TargetNetFilter:              []string{},
	ExcludeTargetNets:            []string{},
	AlibabaCloudConfigFile:       "/etc/kubernetes/alibaba-cloud.json",
	AWSZoneType:                  "",
	AWSZoneTagFilter:             []string{},
	AWSZoneMatchParent:           false,
	AWSAssumeRole:                "",
	AWSAssumeRoleExternalID:      "",
	AWSBatchChangeSize:           1000,
	AWSBatchChangeSizeBytes:      32000,
	AWSBatchChangeSizeValues:     1000,
	AWSBatchChangeInterval:       time.Second,
	AWSEvaluateTargetHealth:      true,
	AWSAPIRetries:                3,
	AWSPreferCNAME:               false,
	AWSZoneCacheDuration:         0 * time.Second,
	AWSSDServiceCleanup:          false,
	AWSSDCreateTag:               map[string]string{},
	AWSDynamoDBRegion:            "",
	AWSDynamoDBTable:             "external-dns",
	AzureConfigFile:              "/etc/kubernetes/azure.json",
	AzureResourceGroup:           "",
	AzureSubscriptionID:          "",
	AzureZonesCacheDuration:      0 * time.Second,
	CloudflareProxied:            false,
	CloudflareDNSRecordsPerPage:  100,
	CloudflareRegionKey:          "earth",
	CoreDNSPrefix:                "/skydns/",
	AkamaiServiceConsumerDomain:  "",
	AkamaiClientToken:            "",
	AkamaiClientSecret:           "",
	AkamaiAccessToken:            "",
	AkamaiEdgercSection:          "",
	AkamaiEdgercPath:             "",
	OCIConfigFile:                "/etc/kubernetes/oci.yaml",
	OCIZoneScope:                 "GLOBAL",
	OCIZoneCacheDuration:         0 * time.Second,
	InMemoryZones:                []string{},
	OVHEndpoint:                  "ovh-eu",
	OVHApiRateLimit:              20,
	PDNSServer:                   "http://localhost:8081",
	PDNSServerID:                 "localhost",
	PDNSAPIKey:                   "",
	PDNSSkipTLSVerify:            false,
	PodSourceDomain:              "",
	TLSCA:                        "",
	TLSClientCert:                "",
	TLSClientCertKey:             "",
	Policy:                       "sync",
	Registry:                     "txt",
	TXTOwnerID:                   "default",
	TXTPrefix:                    "",
	TXTSuffix:                    "",
	TXTCacheInterval:             0,
	TXTWildcardReplacement:       "",
	MinEventSyncInterval:         5 * time.Second,
	TXTEncryptEnabled:            false,
	TXTEncryptAESKey:             "",
	TXTNewFormatOnly:             false,
	Interval:                     time.Minute,
	Once:                         false,
	DryRun:                       false,
	UpdateEvents:                 false,
	LogFormat:                    "text",
	MetricsAddress:               ":7979",
	LogLevel:                     logrus.InfoLevel.String(),
	ExoscaleAPIEnvironment:       "api",
	ExoscaleAPIZone:              "ch-gva-2",
	ExoscaleAPIKey:               "",
	ExoscaleAPISecret:            "",
	CRDSourceAPIVersion:          "externaldns.k8s.io/v1alpha1",
	CRDSourceKind:                "DNSEndpoint",
	ServiceTypeFilter:            []string{},
	CFAPIEndpoint:                "",
	CFUsername:                   "",
	CFPassword:                   "",
	RFC2136Host:                  []string{""},
	RFC2136Port:                  0,
	RFC2136Zone:                  []string{},
	RFC2136Insecure:              false,
	RFC2136GSSTSIG:               false,
	RFC2136KerberosRealm:         "",
	RFC2136KerberosUsername:      "",
	RFC2136KerberosPassword:      "",
	RFC2136TSIGKeyName:           "",
	RFC2136TSIGSecret:            "",
	RFC2136TSIGSecretAlg:         "",
	RFC2136TAXFR:                 true,
	RFC2136MinTTL:                0,
	RFC2136BatchChangeSize:       50,
	RFC2136UseTLS:                false,
	RFC2136LoadBalancingStrategy: "disabled",
	RFC2136SkipTLSVerify:         false,
	NS1Endpoint:                  "",
	NS1IgnoreSSL:                 false,
	TransIPAccountName:           "",
	TransIPPrivateKeyFile:        "",
	DigitalOceanAPIPageSize:      50,
	ManagedDNSRecordTypes:        []string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME},
	ExcludeDNSRecordTypes:        []string{},
	GoDaddyAPIKey:                "",
	GoDaddySecretKey:             "",
	GoDaddyTTL:                   600,
	GoDaddyOTE:                   false,
	IBMCloudProxied:              false,
	IBMCloudConfigFile:           "/etc/kubernetes/ibmcloud.json",
	TencentCloudConfigFile:       "/etc/kubernetes/tencent-cloud.json",
	TencentCloudZoneType:         "",
	PiholeServer:                 "",
	PiholePassword:               "",
	PiholeTLSInsecureSkipVerify:  false,
	PluralCluster:                "",
	PluralProvider:               "",
	WebhookProviderURL:           "http://localhost:8888",
	WebhookProviderReadTimeout:   5 * time.Second,
	WebhookProviderWriteTimeout:  10 * time.Second,
	WebhookServer:                false,
	TraefikDisableLegacy:         false,
	TraefikDisableNew:            false,
	NAT64Networks:                []string{},
}

// NewConfig returns new Config object
func NewConfig() *Config {
	return &Config{
		AWSSDCreateTag: map[string]string{},
	}
}

func (cfg *Config) String() string {
	// prevent logging of sensitive information
	temp := *cfg

	t := reflect.TypeOf(temp)
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
	app := App(cfg)

	_, err := app.Parse(args)
	if err != nil {
		return err
	}

	return nil
}

func App(cfg *Config) *kingpin.Application {
	app := kingpin.New("external-dns", "ExternalDNS synchronizes exposed Kubernetes Services and Ingresses with DNS providers.\n\nNote that all flags may be replaced with env vars - `--flag` -> `EXTERNAL_DNS_FLAG=1` or `--flag value` -> `EXTERNAL_DNS_FLAG=value`")
	app.Version(Version)
	app.DefaultEnvars()

	// Flags related to Kubernetes
	app.Flag("server", "The Kubernetes API server to connect to (default: auto-detect)").Default(defaultConfig.APIServerURL).StringVar(&cfg.APIServerURL)
	app.Flag("kubeconfig", "Retrieve target cluster configuration from a Kubernetes configuration file (default: auto-detect)").Default(defaultConfig.KubeConfig).StringVar(&cfg.KubeConfig)
	app.Flag("request-timeout", "Request timeout when calling Kubernetes APIs. 0s means no timeout").Default(defaultConfig.RequestTimeout.String()).DurationVar(&cfg.RequestTimeout)
	app.Flag("resolve-service-load-balancer-hostname", "Resolve the hostname of LoadBalancer-type Service object to IP addresses in order to create DNS A/AAAA records instead of CNAMEs").BoolVar(&cfg.ResolveServiceLoadBalancerHostname)

	// Flags related to cloud foundry
	app.Flag("cf-api-endpoint", "The fully-qualified domain name of the cloud foundry instance you are targeting").Default(defaultConfig.CFAPIEndpoint).StringVar(&cfg.CFAPIEndpoint)
	app.Flag("cf-username", "The username to log into the cloud foundry API").Default(defaultConfig.CFUsername).StringVar(&cfg.CFUsername)
	app.Flag("cf-password", "The password to log into the cloud foundry API").Default(defaultConfig.CFPassword).StringVar(&cfg.CFPassword)

	// Flags related to Gloo
	app.Flag("gloo-namespace", "The Gloo Proxy namespace; specify multiple times for multiple namespaces. (default: gloo-system)").Default("gloo-system").StringsVar(&cfg.GlooNamespaces)

	// Flags related to Skipper RouteGroup
	app.Flag("skipper-routegroup-groupversion", "The resource version for skipper routegroup").Default(source.DefaultRoutegroupVersion).StringVar(&cfg.SkipperRouteGroupVersion)

	// Flags related to processing source
	app.Flag("source", "The resource types that are queried for endpoints; specify multiple times for multiple sources (required, options: service, ingress, node, pod, fake, connector, gateway-httproute, gateway-grpcroute, gateway-tlsroute, gateway-tcproute, gateway-udproute, istio-gateway, istio-virtualservice, cloudfoundry, contour-httpproxy, gloo-proxy, crd, empty, skipper-routegroup, openshift-route, ambassador-host, kong-tcpingress, f5-virtualserver, f5-transportserver, traefik-proxy)").Required().PlaceHolder("source").EnumsVar(&cfg.Sources, "service", "ingress", "node", "pod", "gateway-httproute", "gateway-grpcroute", "gateway-tlsroute", "gateway-tcproute", "gateway-udproute", "istio-gateway", "istio-virtualservice", "cloudfoundry", "contour-httpproxy", "gloo-proxy", "fake", "connector", "crd", "empty", "skipper-routegroup", "openshift-route", "ambassador-host", "kong-tcpingress", "f5-virtualserver", "f5-transportserver", "traefik-proxy")
	app.Flag("openshift-router-name", "if source is openshift-route then you can pass the ingress controller name. Based on this name external-dns will select the respective router from the route status and map that routerCanonicalHostname to the route host while creating a CNAME record.").StringVar(&cfg.OCPRouterName)
	app.Flag("namespace", "Limit resources queried for endpoints to a specific namespace (default: all namespaces)").Default(defaultConfig.Namespace).StringVar(&cfg.Namespace)
	app.Flag("annotation-filter", "Filter resources queried for endpoints by annotation, using label selector semantics").Default(defaultConfig.AnnotationFilter).StringVar(&cfg.AnnotationFilter)
	app.Flag("label-filter", "Filter resources queried for endpoints by label selector; currently supported by source types crd, gateway-httproute, gateway-grpcroute, gateway-tlsroute, gateway-tcproute, gateway-udproute, ingress, node, openshift-route, service and ambassador-host").Default(defaultConfig.LabelFilter).StringVar(&cfg.LabelFilter)
	app.Flag("ingress-class", "Require an Ingress to have this class name (defaults to any class; specify multiple times to allow more than one class)").StringsVar(&cfg.IngressClassNames)
	app.Flag("fqdn-template", "A templated string that's used to generate DNS names from sources that don't define a hostname themselves, or to add a hostname suffix when paired with the fake source (optional). Accepts comma separated list for multiple global FQDN.").Default(defaultConfig.FQDNTemplate).StringVar(&cfg.FQDNTemplate)
	app.Flag("combine-fqdn-annotation", "Combine FQDN template and Annotations instead of overwriting").BoolVar(&cfg.CombineFQDNAndAnnotation)
	app.Flag("ignore-hostname-annotation", "Ignore hostname annotation when generating DNS names, valid only when --fqdn-template is set (default: false)").BoolVar(&cfg.IgnoreHostnameAnnotation)
	app.Flag("ignore-non-host-network-pods", "Ignore pods not running on host network when using pod source (default: true)").BoolVar(&cfg.IgnoreNonHostNetworkPods)
	app.Flag("ignore-ingress-tls-spec", "Ignore the spec.tls section in Ingress resources (default: false)").BoolVar(&cfg.IgnoreIngressTLSSpec)
	app.Flag("gateway-namespace", "Limit Gateways of Route endpoints to a specific namespace (default: all namespaces)").StringVar(&cfg.GatewayNamespace)
	app.Flag("gateway-label-filter", "Filter Gateways of Route endpoints via label selector (default: all gateways)").StringVar(&cfg.GatewayLabelFilter)
	app.Flag("compatibility", "Process annotation semantics from legacy implementations (optional, options: mate, molecule, kops-dns-controller)").Default(defaultConfig.Compatibility).EnumVar(&cfg.Compatibility, "", "mate", "molecule", "kops-dns-controller")
	app.Flag("ignore-ingress-rules-spec", "Ignore the spec.rules section in Ingress resources (default: false)").BoolVar(&cfg.IgnoreIngressRulesSpec)
	app.Flag("pod-source-domain", "Domain to use for pods records (optional)").Default(defaultConfig.PodSourceDomain).StringVar(&cfg.PodSourceDomain)
	app.Flag("publish-internal-services", "Allow external-dns to publish DNS records for ClusterIP services (optional)").BoolVar(&cfg.PublishInternal)
	app.Flag("publish-host-ip", "Allow external-dns to publish host-ip for headless services (optional)").BoolVar(&cfg.PublishHostIP)
	app.Flag("always-publish-not-ready-addresses", "Always publish also not ready addresses for headless services (optional)").BoolVar(&cfg.AlwaysPublishNotReadyAddresses)
	app.Flag("connector-source-server", "The server to connect for connector source, valid only when using connector source").Default(defaultConfig.ConnectorSourceServer).StringVar(&cfg.ConnectorSourceServer)
	app.Flag("crd-source-apiversion", "API version of the CRD for crd source, e.g. `externaldns.k8s.io/v1alpha1`, valid only when using crd source").Default(defaultConfig.CRDSourceAPIVersion).StringVar(&cfg.CRDSourceAPIVersion)
	app.Flag("crd-source-kind", "Kind of the CRD for the crd source in API group and version specified by crd-source-apiversion").Default(defaultConfig.CRDSourceKind).StringVar(&cfg.CRDSourceKind)
	app.Flag("service-type-filter", "The service types to take care about (default: all, expected: ClusterIP, NodePort, LoadBalancer or ExternalName)").StringsVar(&cfg.ServiceTypeFilter)
	app.Flag("managed-record-types", "Record types to manage; specify multiple times to include many; (default: A, AAAA, CNAME) (supported records: A, AAAA, CNAME, NS, SRV, TXT)").Default("A", "AAAA", "CNAME").StringsVar(&cfg.ManagedDNSRecordTypes)
	app.Flag("exclude-record-types", "Record types to exclude from management; specify multiple times to exclude many; (optional)").Default().StringsVar(&cfg.ExcludeDNSRecordTypes)
	app.Flag("default-targets", "Set globally default host/IP that will apply as a target instead of source addresses. Specify multiple times for multiple targets (optional)").StringsVar(&cfg.DefaultTargets)
	app.Flag("target-net-filter", "Limit possible targets by a net filter; specify multiple times for multiple possible nets (optional)").StringsVar(&cfg.TargetNetFilter)
	app.Flag("exclude-target-net", "Exclude target nets (optional)").StringsVar(&cfg.ExcludeTargetNets)
	app.Flag("traefik-disable-legacy", "Disable listeners on Resources under the traefik.containo.us API Group").Default(strconv.FormatBool(defaultConfig.TraefikDisableLegacy)).BoolVar(&cfg.TraefikDisableLegacy)
	app.Flag("traefik-disable-new", "Disable listeners on Resources under the traefik.io API Group").Default(strconv.FormatBool(defaultConfig.TraefikDisableNew)).BoolVar(&cfg.TraefikDisableNew)
	app.Flag("nat64-networks", "Adding an A record for each AAAA record in NAT64-enabled networks; specify multiple times for multiple possible nets (optional)").StringsVar(&cfg.NAT64Networks)

	// Flags related to providers
	providers := []string{"akamai", "alibabacloud", "aws", "aws-sd", "azure", "azure-dns", "azure-private-dns", "civo", "cloudflare", "coredns", "designate", "digitalocean", "dnsimple", "exoscale", "gandi", "godaddy", "google", "ibmcloud", "inmemory", "linode", "ns1", "oci", "ovh", "pdns", "pihole", "plural", "rfc2136", "scaleway", "skydns", "tencentcloud", "transip", "ultradns", "webhook"}
	app.Flag("provider", "The DNS provider where the DNS records will be created (required, options: "+strings.Join(providers, ", ")+")").Required().PlaceHolder("provider").EnumVar(&cfg.Provider, providers...)
	app.Flag("provider-cache-time", "The time to cache the DNS provider record list requests.").Default(defaultConfig.ProviderCacheTime.String()).DurationVar(&cfg.ProviderCacheTime)
	app.Flag("domain-filter", "Limit possible target zones by a domain suffix; specify multiple times for multiple domains (optional)").Default("").StringsVar(&cfg.DomainFilter)
	app.Flag("exclude-domains", "Exclude subdomains (optional)").Default("").StringsVar(&cfg.ExcludeDomains)
	app.Flag("regex-domain-filter", "Limit possible domains and target zones by a Regex filter; Overrides domain-filter (optional)").Default(defaultConfig.RegexDomainFilter.String()).RegexpVar(&cfg.RegexDomainFilter)
	app.Flag("regex-domain-exclusion", "Regex filter that excludes domains and target zones matched by regex-domain-filter (optional); Require 'regex-domain-filter' ").Default(defaultConfig.RegexDomainExclusion.String()).RegexpVar(&cfg.RegexDomainExclusion)
	app.Flag("zone-name-filter", "Filter target zones by zone domain (For now, only AzureDNS provider is using this flag); specify multiple times for multiple zones (optional)").Default("").StringsVar(&cfg.ZoneNameFilter)
	app.Flag("zone-id-filter", "Filter target zones by hosted zone id; specify multiple times for multiple zones (optional)").Default("").StringsVar(&cfg.ZoneIDFilter)
	app.Flag("google-project", "When using the Google provider, current project is auto-detected, when running on GCP. Specify other project with this. Must be specified when running outside GCP.").Default(defaultConfig.GoogleProject).StringVar(&cfg.GoogleProject)
	app.Flag("google-batch-change-size", "When using the Google provider, set the maximum number of changes that will be applied in each batch.").Default(strconv.Itoa(defaultConfig.GoogleBatchChangeSize)).IntVar(&cfg.GoogleBatchChangeSize)
	app.Flag("google-batch-change-interval", "When using the Google provider, set the interval between batch changes.").Default(defaultConfig.GoogleBatchChangeInterval.String()).DurationVar(&cfg.GoogleBatchChangeInterval)
	app.Flag("google-zone-visibility", "When using the Google provider, filter for zones with this visibility (optional, options: public, private)").Default(defaultConfig.GoogleZoneVisibility).EnumVar(&cfg.GoogleZoneVisibility, "", "public", "private")
	app.Flag("alibaba-cloud-config-file", "When using the Alibaba Cloud provider, specify the Alibaba Cloud configuration file (required when --provider=alibabacloud)").Default(defaultConfig.AlibabaCloudConfigFile).StringVar(&cfg.AlibabaCloudConfigFile)
	app.Flag("alibaba-cloud-zone-type", "When using the Alibaba Cloud provider, filter for zones of this type (optional, options: public, private)").Default(defaultConfig.AlibabaCloudZoneType).EnumVar(&cfg.AlibabaCloudZoneType, "", "public", "private")
	app.Flag("aws-zone-type", "When using the AWS provider, filter for zones of this type (optional, options: public, private)").Default(defaultConfig.AWSZoneType).EnumVar(&cfg.AWSZoneType, "", "public", "private")
	app.Flag("aws-zone-tags", "When using the AWS provider, filter for zones with these tags").Default("").StringsVar(&cfg.AWSZoneTagFilter)
	app.Flag("aws-profile", "When using the AWS provider, name of the profile to use").Default("").StringsVar(&cfg.AWSProfiles)
	app.Flag("aws-assume-role", "When using the AWS API, assume this IAM role. Useful for hosted zones in another AWS account. Specify the full ARN, e.g. `arn:aws:iam::123455567:role/external-dns` (optional)").Default(defaultConfig.AWSAssumeRole).StringVar(&cfg.AWSAssumeRole)
	app.Flag("aws-assume-role-external-id", "When using the AWS API and assuming a role then specify this external ID` (optional)").Default(defaultConfig.AWSAssumeRoleExternalID).StringVar(&cfg.AWSAssumeRoleExternalID)
	app.Flag("aws-batch-change-size", "When using the AWS provider, set the maximum number of changes that will be applied in each batch.").Default(strconv.Itoa(defaultConfig.AWSBatchChangeSize)).IntVar(&cfg.AWSBatchChangeSize)
	app.Flag("aws-batch-change-size-bytes", "When using the AWS provider, set the maximum byte size that will be applied in each batch.").Default(strconv.Itoa(defaultConfig.AWSBatchChangeSizeBytes)).IntVar(&cfg.AWSBatchChangeSizeBytes)
	app.Flag("aws-batch-change-size-values", "When using the AWS provider, set the maximum total record values that will be applied in each batch.").Default(strconv.Itoa(defaultConfig.AWSBatchChangeSizeValues)).IntVar(&cfg.AWSBatchChangeSizeValues)
	app.Flag("aws-batch-change-interval", "When using the AWS provider, set the interval between batch changes.").Default(defaultConfig.AWSBatchChangeInterval.String()).DurationVar(&cfg.AWSBatchChangeInterval)
	app.Flag("aws-evaluate-target-health", "When using the AWS provider, set whether to evaluate the health of a DNS target (default: enabled, disable with --no-aws-evaluate-target-health)").Default(strconv.FormatBool(defaultConfig.AWSEvaluateTargetHealth)).BoolVar(&cfg.AWSEvaluateTargetHealth)
	app.Flag("aws-api-retries", "When using the AWS API, set the maximum number of retries before giving up.").Default(strconv.Itoa(defaultConfig.AWSAPIRetries)).IntVar(&cfg.AWSAPIRetries)
	app.Flag("aws-prefer-cname", "When using the AWS provider, prefer using CNAME instead of ALIAS (default: disabled)").BoolVar(&cfg.AWSPreferCNAME)
	app.Flag("aws-zones-cache-duration", "When using the AWS provider, set the zones list cache TTL (0s to disable).").Default(defaultConfig.AWSZoneCacheDuration.String()).DurationVar(&cfg.AWSZoneCacheDuration)
	app.Flag("aws-zone-match-parent", "Expand limit possible target by sub-domains (default: disabled)").BoolVar(&cfg.AWSZoneMatchParent)
	app.Flag("aws-sd-service-cleanup", "When using the AWS CloudMap provider, delete empty Services without endpoints (default: disabled)").BoolVar(&cfg.AWSSDServiceCleanup)
	app.Flag("aws-sd-create-tag", "When using the AWS CloudMap provider, add tag to created services. The flag can be used multiple times").StringMapVar(&cfg.AWSSDCreateTag)
	app.Flag("azure-config-file", "When using the Azure provider, specify the Azure configuration file (required when --provider=azure)").Default(defaultConfig.AzureConfigFile).StringVar(&cfg.AzureConfigFile)
	app.Flag("azure-resource-group", "When using the Azure provider, override the Azure resource group to use (optional)").Default(defaultConfig.AzureResourceGroup).StringVar(&cfg.AzureResourceGroup)
	app.Flag("azure-subscription-id", "When using the Azure provider, override the Azure subscription to use (optional)").Default(defaultConfig.AzureSubscriptionID).StringVar(&cfg.AzureSubscriptionID)
	app.Flag("azure-user-assigned-identity-client-id", "When using the Azure provider, override the client id of user assigned identity in config file (optional)").Default("").StringVar(&cfg.AzureUserAssignedIdentityClientID)
	app.Flag("azure-zones-cache-duration", "When using the Azure provider, set the zones list cache TTL (0s to disable).").Default(defaultConfig.AzureZonesCacheDuration.String()).DurationVar(&cfg.AzureZonesCacheDuration)
	app.Flag("tencent-cloud-config-file", "When using the Tencent Cloud provider, specify the Tencent Cloud configuration file (required when --provider=tencentcloud)").Default(defaultConfig.TencentCloudConfigFile).StringVar(&cfg.TencentCloudConfigFile)
	app.Flag("tencent-cloud-zone-type", "When using the Tencent Cloud provider, filter for zones with visibility (optional, options: public, private)").Default(defaultConfig.TencentCloudZoneType).EnumVar(&cfg.TencentCloudZoneType, "", "public", "private")

	app.Flag("cloudflare-proxied", "When using the Cloudflare provider, specify if the proxy mode must be enabled (default: disabled)").BoolVar(&cfg.CloudflareProxied)
	app.Flag("cloudflare-dns-records-per-page", "When using the Cloudflare provider, specify how many DNS records listed per page, max possible 5,000 (default: 100)").Default(strconv.Itoa(defaultConfig.CloudflareDNSRecordsPerPage)).IntVar(&cfg.CloudflareDNSRecordsPerPage)
	app.Flag("cloudflare-region-key", "When using the Cloudflare provider, specify the region (default: earth)").StringVar(&cfg.CloudflareRegionKey)
	app.Flag("coredns-prefix", "When using the CoreDNS provider, specify the prefix name").Default(defaultConfig.CoreDNSPrefix).StringVar(&cfg.CoreDNSPrefix)
	app.Flag("akamai-serviceconsumerdomain", "When using the Akamai provider, specify the base URL (required when --provider=akamai and edgerc-path not specified)").Default(defaultConfig.AkamaiServiceConsumerDomain).StringVar(&cfg.AkamaiServiceConsumerDomain)
	app.Flag("akamai-client-token", "When using the Akamai provider, specify the client token (required when --provider=akamai and edgerc-path not specified)").Default(defaultConfig.AkamaiClientToken).StringVar(&cfg.AkamaiClientToken)
	app.Flag("akamai-client-secret", "When using the Akamai provider, specify the client secret (required when --provider=akamai and edgerc-path not specified)").Default(defaultConfig.AkamaiClientSecret).StringVar(&cfg.AkamaiClientSecret)
	app.Flag("akamai-access-token", "When using the Akamai provider, specify the access token (required when --provider=akamai and edgerc-path not specified)").Default(defaultConfig.AkamaiAccessToken).StringVar(&cfg.AkamaiAccessToken)
	app.Flag("akamai-edgerc-path", "When using the Akamai provider, specify the .edgerc file path. Path must be reachable form invocation environment. (required when --provider=akamai and *-token, secret serviceconsumerdomain not specified)").Default(defaultConfig.AkamaiEdgercPath).StringVar(&cfg.AkamaiEdgercPath)
	app.Flag("akamai-edgerc-section", "When using the Akamai provider, specify the .edgerc file path (Optional when edgerc-path is specified)").Default(defaultConfig.AkamaiEdgercSection).StringVar(&cfg.AkamaiEdgercSection)
	app.Flag("oci-config-file", "When using the OCI provider, specify the OCI configuration file (required when --provider=oci").Default(defaultConfig.OCIConfigFile).StringVar(&cfg.OCIConfigFile)
	app.Flag("oci-compartment-ocid", "When using the OCI provider, specify the OCID of the OCI compartment containing all managed zones and records.  Required when using OCI IAM instance principal authentication.").StringVar(&cfg.OCICompartmentOCID)
	app.Flag("oci-zone-scope", "When using OCI provider, filter for zones with this scope (optional, options: GLOBAL, PRIVATE). Defaults to GLOBAL, setting to empty value will target both.").Default(defaultConfig.OCIZoneScope).EnumVar(&cfg.OCIZoneScope, "", "GLOBAL", "PRIVATE")
	app.Flag("oci-auth-instance-principal", "When using the OCI provider, specify whether OCI IAM instance principal authentication should be used (instead of key-based auth via the OCI config file).").Default(strconv.FormatBool(defaultConfig.OCIAuthInstancePrincipal)).BoolVar(&cfg.OCIAuthInstancePrincipal)
	app.Flag("oci-zones-cache-duration", "When using the OCI provider, set the zones list cache TTL (0s to disable).").Default(defaultConfig.OCIZoneCacheDuration.String()).DurationVar(&cfg.OCIZoneCacheDuration)
	app.Flag("inmemory-zone", "Provide a list of pre-configured zones for the inmemory provider; specify multiple times for multiple zones (optional)").Default("").StringsVar(&cfg.InMemoryZones)
	app.Flag("ovh-endpoint", "When using the OVH provider, specify the endpoint (default: ovh-eu)").Default(defaultConfig.OVHEndpoint).StringVar(&cfg.OVHEndpoint)
	app.Flag("ovh-api-rate-limit", "When using the OVH provider, specify the API request rate limit, X operations by seconds (default: 20)").Default(strconv.Itoa(defaultConfig.OVHApiRateLimit)).IntVar(&cfg.OVHApiRateLimit)
	app.Flag("pdns-server", "When using the PowerDNS/PDNS provider, specify the URL to the pdns server (required when --provider=pdns)").Default(defaultConfig.PDNSServer).StringVar(&cfg.PDNSServer)
	app.Flag("pdns-server-id", "When using the PowerDNS/PDNS provider, specify the id of the server to retrieve. Should be `localhost` except when the server is behind a proxy (optional when --provider=pdns) (default: localhost)").Default(defaultConfig.PDNSServerID).StringVar(&cfg.PDNSServerID)
	app.Flag("pdns-api-key", "When using the PowerDNS/PDNS provider, specify the API key to use to authorize requests (required when --provider=pdns)").Default(defaultConfig.PDNSAPIKey).StringVar(&cfg.PDNSAPIKey)
	app.Flag("pdns-skip-tls-verify", "When using the PowerDNS/PDNS provider, disable verification of any TLS certificates (optional when --provider=pdns) (default: false)").Default(strconv.FormatBool(defaultConfig.PDNSSkipTLSVerify)).BoolVar(&cfg.PDNSSkipTLSVerify)
	app.Flag("ns1-endpoint", "When using the NS1 provider, specify the URL of the API endpoint to target (default: https://api.nsone.net/v1/)").Default(defaultConfig.NS1Endpoint).StringVar(&cfg.NS1Endpoint)
	app.Flag("ns1-ignoressl", "When using the NS1 provider, specify whether to verify the SSL certificate (default: false)").Default(strconv.FormatBool(defaultConfig.NS1IgnoreSSL)).BoolVar(&cfg.NS1IgnoreSSL)
	app.Flag("ns1-min-ttl", "Minimal TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is lower than this.").IntVar(&cfg.NS1MinTTLSeconds)
	app.Flag("digitalocean-api-page-size", "Configure the page size used when querying the DigitalOcean API.").Default(strconv.Itoa(defaultConfig.DigitalOceanAPIPageSize)).IntVar(&cfg.DigitalOceanAPIPageSize)
	app.Flag("ibmcloud-config-file", "When using the IBM Cloud provider, specify the IBM Cloud configuration file (required when --provider=ibmcloud").Default(defaultConfig.IBMCloudConfigFile).StringVar(&cfg.IBMCloudConfigFile)
	app.Flag("ibmcloud-proxied", "When using the IBM provider, specify if the proxy mode must be enabled (default: disabled)").BoolVar(&cfg.IBMCloudProxied)
	// GoDaddy flags
	app.Flag("godaddy-api-key", "When using the GoDaddy provider, specify the API Key (required when --provider=godaddy)").Default(defaultConfig.GoDaddyAPIKey).StringVar(&cfg.GoDaddyAPIKey)
	app.Flag("godaddy-api-secret", "When using the GoDaddy provider, specify the API secret (required when --provider=godaddy)").Default(defaultConfig.GoDaddySecretKey).StringVar(&cfg.GoDaddySecretKey)
	app.Flag("godaddy-api-ttl", "TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is not provided.").Int64Var(&cfg.GoDaddyTTL)
	app.Flag("godaddy-api-ote", "When using the GoDaddy provider, use OTE api (optional, default: false, when --provider=godaddy)").BoolVar(&cfg.GoDaddyOTE)

	// Flags related to TLS communication
	app.Flag("tls-ca", "When using TLS communication, the path to the certificate authority to verify server communications (optionally specify --tls-client-cert for two-way TLS)").Default(defaultConfig.TLSCA).StringVar(&cfg.TLSCA)
	app.Flag("tls-client-cert", "When using TLS communication, the path to the certificate to present as a client (not required for TLS)").Default(defaultConfig.TLSClientCert).StringVar(&cfg.TLSClientCert)
	app.Flag("tls-client-cert-key", "When using TLS communication, the path to the certificate key to use with the client certificate (not required for TLS)").Default(defaultConfig.TLSClientCertKey).StringVar(&cfg.TLSClientCertKey)

	// Flags related to Exoscale provider
	app.Flag("exoscale-apienv", "When using Exoscale provider, specify the API environment (optional)").Default(defaultConfig.ExoscaleAPIEnvironment).StringVar(&cfg.ExoscaleAPIEnvironment)
	app.Flag("exoscale-apizone", "When using Exoscale provider, specify the API Zone (optional)").Default(defaultConfig.ExoscaleAPIZone).StringVar(&cfg.ExoscaleAPIZone)
	app.Flag("exoscale-apikey", "Provide your API Key for the Exoscale provider").Default(defaultConfig.ExoscaleAPIKey).StringVar(&cfg.ExoscaleAPIKey)
	app.Flag("exoscale-apisecret", "Provide your API Secret for the Exoscale provider").Default(defaultConfig.ExoscaleAPISecret).StringVar(&cfg.ExoscaleAPISecret)

	// Flags related to RFC2136 provider
	app.Flag("rfc2136-host", "When using the RFC2136 provider, specify the host of the DNS server (optionally specify multiple times when when using --rfc2136-load-balancing-strategy)").Default(defaultConfig.RFC2136Host[0]).StringsVar(&cfg.RFC2136Host)
	app.Flag("rfc2136-port", "When using the RFC2136 provider, specify the port of the DNS server").Default(strconv.Itoa(defaultConfig.RFC2136Port)).IntVar(&cfg.RFC2136Port)
	app.Flag("rfc2136-zone", "When using the RFC2136 provider, specify zone entries of the DNS server to use").StringsVar(&cfg.RFC2136Zone)
	app.Flag("rfc2136-create-ptr", "When using the RFC2136 provider, enable PTR management").Default(strconv.FormatBool(defaultConfig.RFC2136CreatePTR)).BoolVar(&cfg.RFC2136CreatePTR)
	app.Flag("rfc2136-insecure", "When using the RFC2136 provider, specify whether to attach TSIG or not (default: false, requires --rfc2136-tsig-keyname and rfc2136-tsig-secret)").Default(strconv.FormatBool(defaultConfig.RFC2136Insecure)).BoolVar(&cfg.RFC2136Insecure)
	app.Flag("rfc2136-tsig-keyname", "When using the RFC2136 provider, specify the TSIG key to attached to DNS messages (required when --rfc2136-insecure=false)").Default(defaultConfig.RFC2136TSIGKeyName).StringVar(&cfg.RFC2136TSIGKeyName)
	app.Flag("rfc2136-tsig-secret", "When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)").Default(defaultConfig.RFC2136TSIGSecret).StringVar(&cfg.RFC2136TSIGSecret)
	app.Flag("rfc2136-tsig-secret-alg", "When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)").Default(defaultConfig.RFC2136TSIGSecretAlg).StringVar(&cfg.RFC2136TSIGSecretAlg)
	app.Flag("rfc2136-tsig-axfr", "When using the RFC2136 provider, specify the TSIG (base64) value to attached to DNS messages (required when --rfc2136-insecure=false)").BoolVar(&cfg.RFC2136TAXFR)
	app.Flag("rfc2136-min-ttl", "When using the RFC2136 provider, specify minimal TTL (in duration format) for records. This value will be used if the provided TTL for a service/ingress is lower than this").Default(defaultConfig.RFC2136MinTTL.String()).DurationVar(&cfg.RFC2136MinTTL)
	app.Flag("rfc2136-gss-tsig", "When using the RFC2136 provider, specify whether to use secure updates with GSS-TSIG using Kerberos (default: false, requires --rfc2136-kerberos-realm, --rfc2136-kerberos-username, and rfc2136-kerberos-password)").Default(strconv.FormatBool(defaultConfig.RFC2136GSSTSIG)).BoolVar(&cfg.RFC2136GSSTSIG)
	app.Flag("rfc2136-kerberos-username", "When using the RFC2136 provider with GSS-TSIG, specify the username of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)").Default(defaultConfig.RFC2136KerberosUsername).StringVar(&cfg.RFC2136KerberosUsername)
	app.Flag("rfc2136-kerberos-password", "When using the RFC2136 provider with GSS-TSIG, specify the password of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)").Default(defaultConfig.RFC2136KerberosPassword).StringVar(&cfg.RFC2136KerberosPassword)
	app.Flag("rfc2136-kerberos-realm", "When using the RFC2136 provider with GSS-TSIG, specify the realm of the user with permissions to update DNS records (required when --rfc2136-gss-tsig=true)").Default(defaultConfig.RFC2136KerberosRealm).StringVar(&cfg.RFC2136KerberosRealm)
	app.Flag("rfc2136-batch-change-size", "When using the RFC2136 provider, set the maximum number of changes that will be applied in each batch.").Default(strconv.Itoa(defaultConfig.RFC2136BatchChangeSize)).IntVar(&cfg.RFC2136BatchChangeSize)
	app.Flag("rfc2136-use-tls", "When using the RFC2136 provider, communicate with name server over tls").BoolVar(&cfg.RFC2136UseTLS)
	app.Flag("rfc2136-skip-tls-verify", "When using TLS with the RFC2136 provider, disable verification of any TLS certificates").BoolVar(&cfg.RFC2136SkipTLSVerify)
	app.Flag("rfc2136-load-balancing-strategy", "When using the RFC2136 provider, specify the load balancing strategy (default: disabled, options: random, round-robin, disabled)").Default(defaultConfig.RFC2136LoadBalancingStrategy).EnumVar(&cfg.RFC2136LoadBalancingStrategy, "random", "round-robin", "disabled")

	// Flags related to TransIP provider
	app.Flag("transip-account", "When using the TransIP provider, specify the account name (required when --provider=transip)").Default(defaultConfig.TransIPAccountName).StringVar(&cfg.TransIPAccountName)
	app.Flag("transip-keyfile", "When using the TransIP provider, specify the path to the private key file (required when --provider=transip)").Default(defaultConfig.TransIPPrivateKeyFile).StringVar(&cfg.TransIPPrivateKeyFile)

	// Flags related to Pihole provider
	app.Flag("pihole-server", "When using the Pihole provider, the base URL of the Pihole web server (required when --provider=pihole)").Default(defaultConfig.PiholeServer).StringVar(&cfg.PiholeServer)
	app.Flag("pihole-password", "When using the Pihole provider, the password to the server if it is protected").Default(defaultConfig.PiholePassword).StringVar(&cfg.PiholePassword)
	app.Flag("pihole-tls-skip-verify", "When using the Pihole provider, disable verification of any TLS certificates").BoolVar(&cfg.PiholeTLSInsecureSkipVerify)

	// Flags related to the Plural provider
	app.Flag("plural-cluster", "When using the plural provider, specify the cluster name you're running with").Default(defaultConfig.PluralCluster).StringVar(&cfg.PluralCluster)
	app.Flag("plural-provider", "When using the plural provider, specify the provider name you're running with").Default(defaultConfig.PluralProvider).StringVar(&cfg.PluralProvider)

	// Flags related to policies
	app.Flag("policy", "Modify how DNS records are synchronized between sources and providers (default: sync, options: sync, upsert-only, create-only)").Default(defaultConfig.Policy).EnumVar(&cfg.Policy, "sync", "upsert-only", "create-only")

	// Flags related to the registry
	app.Flag("registry", "The registry implementation to use to keep track of DNS record ownership (default: txt, options: txt, noop, dynamodb, aws-sd)").Default(defaultConfig.Registry).EnumVar(&cfg.Registry, "txt", "noop", "dynamodb", "aws-sd")
	app.Flag("txt-owner-id", "When using the TXT or DynamoDB registry, a name that identifies this instance of ExternalDNS (default: default)").Default(defaultConfig.TXTOwnerID).StringVar(&cfg.TXTOwnerID)
	app.Flag("txt-prefix", "When using the TXT registry, a custom string that's prefixed to each ownership DNS record (optional). Could contain record type template like '%{record_type}-prefix-'. Mutual exclusive with txt-suffix!").Default(defaultConfig.TXTPrefix).StringVar(&cfg.TXTPrefix)
	app.Flag("txt-suffix", "When using the TXT registry, a custom string that's suffixed to the host portion of each ownership DNS record (optional). Could contain record type template like '-%{record_type}-suffix'. Mutual exclusive with txt-prefix!").Default(defaultConfig.TXTSuffix).StringVar(&cfg.TXTSuffix)
	app.Flag("txt-wildcard-replacement", "When using the TXT registry, a custom string that's used instead of an asterisk for TXT records corresponding to wildcard DNS records (optional)").Default(defaultConfig.TXTWildcardReplacement).StringVar(&cfg.TXTWildcardReplacement)
	app.Flag("txt-encrypt-enabled", "When using the TXT registry, set if TXT records should be encrypted before stored (default: disabled)").BoolVar(&cfg.TXTEncryptEnabled)
	app.Flag("txt-encrypt-aes-key", "When using the TXT registry, set TXT record decryption and encryption 32 byte aes key (required when --txt-encrypt=true)").Default(defaultConfig.TXTEncryptAESKey).StringVar(&cfg.TXTEncryptAESKey)
	app.Flag("txt-new-format-only", "When using the TXT registry, only use new format records which include record type information (e.g., prefix: 'a-'). Reduces number of TXT records (default: disabled)").BoolVar(&cfg.TXTNewFormatOnly)
	app.Flag("dynamodb-region", "When using the DynamoDB registry, the AWS region of the DynamoDB table (optional)").Default(cfg.AWSDynamoDBRegion).StringVar(&cfg.AWSDynamoDBRegion)
	app.Flag("dynamodb-table", "When using the DynamoDB registry, the name of the DynamoDB table (default: \"external-dns\")").Default(defaultConfig.AWSDynamoDBTable).StringVar(&cfg.AWSDynamoDBTable)

	// Flags related to the main control loop
	app.Flag("txt-cache-interval", "The interval between cache synchronizations in duration format (default: disabled)").Default(defaultConfig.TXTCacheInterval.String()).DurationVar(&cfg.TXTCacheInterval)
	app.Flag("interval", "The interval between two consecutive synchronizations in duration format (default: 1m)").Default(defaultConfig.Interval.String()).DurationVar(&cfg.Interval)
	app.Flag("min-event-sync-interval", "The minimum interval between two consecutive synchronizations triggered from kubernetes events in duration format (default: 5s)").Default(defaultConfig.MinEventSyncInterval.String()).DurationVar(&cfg.MinEventSyncInterval)
	app.Flag("once", "When enabled, exits the synchronization loop after the first iteration (default: disabled)").BoolVar(&cfg.Once)
	app.Flag("dry-run", "When enabled, prints DNS record changes rather than actually performing them (default: disabled)").BoolVar(&cfg.DryRun)
	app.Flag("events", "When enabled, in addition to running every interval, the reconciliation loop will get triggered when supported sources change (default: disabled)").BoolVar(&cfg.UpdateEvents)

	// Miscellaneous flags
	app.Flag("log-format", "The format in which log messages are printed (default: text, options: text, json)").Default(defaultConfig.LogFormat).EnumVar(&cfg.LogFormat, "text", "json")
	app.Flag("metrics-address", "Specify where to serve the metrics and health check endpoint (default: :7979)").Default(defaultConfig.MetricsAddress).StringVar(&cfg.MetricsAddress)
	app.Flag("log-level", "Set the level of logging. (default: info, options: panic, debug, info, warning, error, fatal)").Default(defaultConfig.LogLevel).EnumVar(&cfg.LogLevel, allLogLevelsAsStrings()...)

	// Webhook provider
	app.Flag("webhook-provider-url", "The URL of the remote endpoint to call for the webhook provider (default: http://localhost:8888)").Default(defaultConfig.WebhookProviderURL).StringVar(&cfg.WebhookProviderURL)
	app.Flag("webhook-provider-read-timeout", "The read timeout for the webhook provider in duration format (default: 5s)").Default(defaultConfig.WebhookProviderReadTimeout.String()).DurationVar(&cfg.WebhookProviderReadTimeout)
	app.Flag("webhook-provider-write-timeout", "The write timeout for the webhook provider in duration format (default: 10s)").Default(defaultConfig.WebhookProviderWriteTimeout.String()).DurationVar(&cfg.WebhookProviderWriteTimeout)

	app.Flag("webhook-server", "When enabled, runs as a webhook server instead of a controller. (default: false).").BoolVar(&cfg.WebhookServer)

	return app
}
