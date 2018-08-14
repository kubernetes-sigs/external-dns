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
	"strconv"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/sirupsen/logrus"
)

const (
	passwordMask = "******"
)

var (
	// Version is the current version of the app, generated at build time
	Version = "unknown"
)

// Config is a project-wide configuration
type Config struct {
	Master                   string
	KubeConfig               string
	Sources                  []string
	Namespace                string
	AnnotationFilter         string
	FQDNTemplate             string
	CombineFQDNAndAnnotation bool
	Compatibility            string
	PublishInternal          bool
	PublishHostIP            bool
	ConnectorSourceServer    string
	Provider                 string
	GoogleProject            string
	DomainFilter             []string
	ZoneIDFilter             []string
	AWSZoneType              string
	AWSAssumeRole            string
	AWSMaxChangeCount        int
	AWSEvaluateTargetHealth  bool
	AzureConfigFile          string
	AzureResourceGroup       string
	CloudflareProxied        bool
	InfobloxGridHost         string
	InfobloxWapiPort         int
	InfobloxWapiUsername     string
	InfobloxWapiPassword     string
	InfobloxWapiVersion      string
	InfobloxSSLVerify        bool
	DynCustomerName          string
	DynUsername              string
	DynPassword              string
	DynMinTTLSeconds         int
	OCIConfigFile            string
	InMemoryZones            []string
	PDNSServer               string
	PDNSAPIKey               string
	PDNSTLSEnabled           bool
	TLSCA                    string
	TLSClientCert            string
	TLSClientCertKey         string
	Policy                   string
	Registry                 string
	TXTOwnerID               string
	TXTPrefix                string
	Interval                 time.Duration
	Once                     bool
	DryRun                   bool
	LogFormat                string
	MetricsAddress           string
	LogLevel                 string
	TXTCacheInterval         time.Duration
	ExoscaleEndpoint         string
	ExoscaleAPIKey           string
	ExoscaleAPISecret        string
}

var defaultConfig = &Config{
	Master:                   "",
	KubeConfig:               "",
	Sources:                  nil,
	Namespace:                "",
	AnnotationFilter:         "",
	FQDNTemplate:             "",
	CombineFQDNAndAnnotation: false,
	Compatibility:            "",
	PublishInternal:          false,
	PublishHostIP:            false,
	ConnectorSourceServer:    "localhost:8080",
	Provider:                 "",
	GoogleProject:            "",
	DomainFilter:             []string{},
	AWSZoneType:              "",
	AWSAssumeRole:            "",
	AWSMaxChangeCount:        4000,
	AWSEvaluateTargetHealth:  true,
	AzureConfigFile:          "/etc/kubernetes/azure.json",
	AzureResourceGroup:       "",
	CloudflareProxied:        false,
	InfobloxGridHost:         "",
	InfobloxWapiPort:         443,
	InfobloxWapiUsername:     "admin",
	InfobloxWapiPassword:     "",
	InfobloxWapiVersion:      "2.3.1",
	InfobloxSSLVerify:        true,
	OCIConfigFile:            "/etc/kubernetes/oci.yaml",
	InMemoryZones:            []string{},
	PDNSServer:               "http://localhost:8081",
	PDNSAPIKey:               "",
	PDNSTLSEnabled:           false,
	TLSCA:                    "",
	TLSClientCert:            "",
	TLSClientCertKey:         "",
	Policy:                   "sync",
	Registry:                 "txt",
	TXTOwnerID:               "default",
	TXTPrefix:                "",
	TXTCacheInterval:         0,
	Interval:                 time.Minute,
	Once:                     false,
	DryRun:                   false,
	LogFormat:                "text",
	MetricsAddress:           ":7979",
	LogLevel:                 logrus.InfoLevel.String(),
	ExoscaleEndpoint:         "https://api.exoscale.ch/dns",
	ExoscaleAPIKey:           "",
	ExoscaleAPISecret:        "",
}

// NewConfig returns new Config object
func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) String() string {
	// prevent logging of sensitive information
	temp := *cfg
	if temp.DynPassword != "" {
		temp.DynPassword = passwordMask
	}
	if temp.InfobloxWapiPassword != "" {
		temp.InfobloxWapiPassword = passwordMask
	}
	if temp.PDNSAPIKey != "" {
		temp.PDNSAPIKey = ""
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
	app := kingpin.New("external-dns", "ExternalDNS synchronizes exposed Kubernetes Services and Ingresses with DNS providers.\n\nNote that all flags may be replaced with env vars - `--flag` -> `EXTERNAL_DNS_FLAG=1` or `--flag value` -> `EXTERNAL_DNS_FLAG=value`")
	app.Version(Version)
	app.DefaultEnvars()

	// Flags related to Kubernetes
	app.Flag("master", "The Kubernetes API server to connect to (default: auto-detect)").Default(defaultConfig.Master).StringVar(&cfg.Master)
	app.Flag("kubeconfig", "Retrieve target cluster configuration from a Kubernetes configuration file (default: auto-detect)").Default(defaultConfig.KubeConfig).StringVar(&cfg.KubeConfig)

	// Flags related to processing sources
	app.Flag("source", "The resource types that are queried for endpoints; specify multiple times for multiple sources (required, options: service, ingress, fake, connector)").Required().PlaceHolder("source").EnumsVar(&cfg.Sources, "service", "ingress", "fake", "connector")
	app.Flag("namespace", "Limit sources of endpoints to a specific namespace (default: all namespaces)").Default(defaultConfig.Namespace).StringVar(&cfg.Namespace)
	app.Flag("annotation-filter", "Filter sources managed by external-dns via annotation using label selector semantics (default: all sources)").Default(defaultConfig.AnnotationFilter).StringVar(&cfg.AnnotationFilter)
	app.Flag("fqdn-template", "A templated string that's used to generate DNS names from sources that don't define a hostname themselves, or to add a hostname suffix when paired with the fake source (optional). Accepts comma separated list for multiple global FQDN.").Default(defaultConfig.FQDNTemplate).StringVar(&cfg.FQDNTemplate)
	app.Flag("combine-fqdn-annotation", "Combine FQDN template and Annotations instead of overwriting").BoolVar(&cfg.CombineFQDNAndAnnotation)
	app.Flag("compatibility", "Process annotation semantics from legacy implementations (optional, options: mate, molecule)").Default(defaultConfig.Compatibility).EnumVar(&cfg.Compatibility, "", "mate", "molecule")
	app.Flag("publish-internal-services", "Allow external-dns to publish DNS records for ClusterIP services (optional)").BoolVar(&cfg.PublishInternal)
	app.Flag("publish-host-ip", "Allow external-dns to publish host-ip for headless services (optional)").BoolVar(&cfg.PublishHostIP)
	app.Flag("connector-source-server", "The server to connect for connector source, valid only when using connector source").Default(defaultConfig.ConnectorSourceServer).StringVar(&cfg.ConnectorSourceServer)

	// Flags related to providers
	app.Flag("provider", "The DNS provider where the DNS records will be created (required, options: aws, aws-sd, google, azure, cloudflare, digitalocean, dnsimple, infoblox, dyn, designate, coredns, skydns, inmemory, pdns, oci, exoscale, linode)").Required().PlaceHolder("provider").EnumVar(&cfg.Provider, "aws", "aws-sd", "google", "azure", "cloudflare", "digitalocean", "dnsimple", "infoblox", "dyn", "designate", "coredns", "skydns", "inmemory", "pdns", "oci", "exoscale", "linode")
	app.Flag("domain-filter", "Limit possible target zones by a domain suffix; specify multiple times for multiple domains (optional)").Default("").StringsVar(&cfg.DomainFilter)
	app.Flag("zone-id-filter", "Filter target zones by hosted zone id; specify multiple times for multiple zones (optional)").Default("").StringsVar(&cfg.ZoneIDFilter)
	app.Flag("google-project", "When using the Google provider, current project is auto-detected, when running on GCP. Specify other project with this. Must be specified when running outside GCP.").Default(defaultConfig.GoogleProject).StringVar(&cfg.GoogleProject)
	app.Flag("aws-zone-type", "When using the AWS provider, filter for zones of this type (optional, options: public, private)").Default(defaultConfig.AWSZoneType).EnumVar(&cfg.AWSZoneType, "", "public", "private")
	app.Flag("aws-assume-role", "When using the AWS provider, assume this IAM role. Useful for hosted zones in another AWS account. Specify the full ARN, e.g. `arn:aws:iam::123455567:role/external-dns` (optional)").Default(defaultConfig.AWSAssumeRole).StringVar(&cfg.AWSAssumeRole)
	app.Flag("aws-max-change-count", "When using the AWS provider, set the maximum number of changes that will be applied.").Default(strconv.Itoa(defaultConfig.AWSMaxChangeCount)).IntVar(&cfg.AWSMaxChangeCount)
	app.Flag("aws-evaluate-target-health", "When using the AWS provider, set whether to evaluate the health of a DNS target (default: enabled, disable with --no-aws-evaluate-target-health)").Default(strconv.FormatBool(defaultConfig.AWSEvaluateTargetHealth)).BoolVar(&cfg.AWSEvaluateTargetHealth)
	app.Flag("azure-config-file", "When using the Azure provider, specify the Azure configuration file (required when --provider=azure").Default(defaultConfig.AzureConfigFile).StringVar(&cfg.AzureConfigFile)
	app.Flag("azure-resource-group", "When using the Azure provider, override the Azure resource group to use (optional)").Default(defaultConfig.AzureResourceGroup).StringVar(&cfg.AzureResourceGroup)
	app.Flag("cloudflare-proxied", "When using the Cloudflare provider, specify if the proxy mode must be enabled (default: disabled)").BoolVar(&cfg.CloudflareProxied)
	app.Flag("infoblox-grid-host", "When using the Infoblox provider, specify the Grid Manager host (required when --provider=infoblox)").Default(defaultConfig.InfobloxGridHost).StringVar(&cfg.InfobloxGridHost)
	app.Flag("infoblox-wapi-port", "When using the Infoblox provider, specify the WAPI port (default: 443)").Default(strconv.Itoa(defaultConfig.InfobloxWapiPort)).IntVar(&cfg.InfobloxWapiPort)
	app.Flag("infoblox-wapi-username", "When using the Infoblox provider, specify the WAPI username (default: admin)").Default(defaultConfig.InfobloxWapiUsername).StringVar(&cfg.InfobloxWapiUsername)
	app.Flag("infoblox-wapi-password", "When using the Infoblox provider, specify the WAPI password (required when --provider=infoblox)").Default(defaultConfig.InfobloxWapiPassword).StringVar(&cfg.InfobloxWapiPassword)
	app.Flag("infoblox-wapi-version", "When using the Infoblox provider, specify the WAPI version (default: 2.3.1)").Default(defaultConfig.InfobloxWapiVersion).StringVar(&cfg.InfobloxWapiVersion)
	app.Flag("infoblox-ssl-verify", "When using the Infoblox provider, specify whether to verify the SSL certificate (default: true, disable with --no-infoblox-ssl-verify)").Default(strconv.FormatBool(defaultConfig.InfobloxSSLVerify)).BoolVar(&cfg.InfobloxSSLVerify)
	app.Flag("dyn-customer-name", "When using the Dyn provider, specify the Customer Name").Default("").StringVar(&cfg.DynCustomerName)
	app.Flag("dyn-username", "When using the Dyn provider, specify the Username").Default("").StringVar(&cfg.DynUsername)
	app.Flag("dyn-password", "When using the Dyn provider, specify the pasword").Default("").StringVar(&cfg.DynPassword)
	app.Flag("dyn-min-ttl", "Minimal TTL (in seconds) for records. This value will be used if the provided TTL for a service/ingress is lower than this.").IntVar(&cfg.DynMinTTLSeconds)
	app.Flag("oci-config-file", "When using the OCI provider, specify the OCI configuration file (required when --provider=oci").Default(defaultConfig.OCIConfigFile).StringVar(&cfg.OCIConfigFile)

	app.Flag("inmemory-zone", "Provide a list of pre-configured zones for the inmemory provider; specify multiple times for multiple zones (optional)").Default("").StringsVar(&cfg.InMemoryZones)
	app.Flag("pdns-server", "When using the PowerDNS/PDNS provider, specify the URL to the pdns server (required when --provider=pdns)").Default(defaultConfig.PDNSServer).StringVar(&cfg.PDNSServer)
	app.Flag("pdns-api-key", "When using the PowerDNS/PDNS provider, specify the API key to use to authorize requests (required when --provider=pdns)").Default(defaultConfig.PDNSAPIKey).StringVar(&cfg.PDNSAPIKey)
	app.Flag("pdns-tls-enabled", "When using the PowerDNS/PDNS provider, specify whether to use TLS (default: false, requires --tls-ca, optionally specify --tls-client-cert and --tls-client-cert-key)").Default(strconv.FormatBool(defaultConfig.PDNSTLSEnabled)).BoolVar(&cfg.PDNSTLSEnabled)

	// Flags related to TLS communication
	app.Flag("tls-ca", "When using TLS communication, the path to the certificate authority to verify server communications (optionally specify --tls-client-cert for two-way TLS)").Default(defaultConfig.TLSCA).StringVar(&cfg.TLSCA)
	app.Flag("tls-client-cert", "When using TLS communication, the path to the certificate to present as a client (not required for TLS)").Default(defaultConfig.TLSClientCert).StringVar(&cfg.TLSClientCert)
	app.Flag("tls-client-cert-key", "When using TLS communication, the path to the certificate key to use with the client certificate (not required for TLS)").Default(defaultConfig.TLSClientCertKey).StringVar(&cfg.TLSClientCertKey)

	app.Flag("exoscale-endpoint", "Provide the endpoint for the Exoscale provider").Default(defaultConfig.ExoscaleEndpoint).StringVar(&cfg.ExoscaleEndpoint)
	app.Flag("exoscale-apikey", "Provide your API Key for the Exoscale provider").Default(defaultConfig.ExoscaleAPIKey).StringVar(&cfg.ExoscaleAPIKey)
	app.Flag("exoscale-apisecret", "Provide your API Secret for the Exoscale provider").Default(defaultConfig.ExoscaleAPISecret).StringVar(&cfg.ExoscaleAPISecret)

	// Flags related to policies
	app.Flag("policy", "Modify how DNS records are sychronized between sources and providers (default: sync, options: sync, upsert-only)").Default(defaultConfig.Policy).EnumVar(&cfg.Policy, "sync", "upsert-only")

	// Flags related to the registry
	app.Flag("registry", "The registry implementation to use to keep track of DNS record ownership (default: txt, options: txt, noop, aws-sd)").Default(defaultConfig.Registry).EnumVar(&cfg.Registry, "txt", "noop", "aws-sd")
	app.Flag("txt-owner-id", "When using the TXT registry, a name that identifies this instance of ExternalDNS (default: default)").Default(defaultConfig.TXTOwnerID).StringVar(&cfg.TXTOwnerID)
	app.Flag("txt-prefix", "When using the TXT registry, a custom string that's prefixed to each ownership DNS record (optional)").Default(defaultConfig.TXTPrefix).StringVar(&cfg.TXTPrefix)

	// Flags related to the main control loop
	app.Flag("txt-cache-interval", "The interval between cache synchronizations in duration format (default: disabled)").Default(defaultConfig.TXTCacheInterval.String()).DurationVar(&cfg.TXTCacheInterval)
	app.Flag("interval", "The interval between two consecutive synchronizations in duration format (default: 1m)").Default(defaultConfig.Interval.String()).DurationVar(&cfg.Interval)
	app.Flag("once", "When enabled, exits the synchronization loop after the first iteration (default: disabled)").BoolVar(&cfg.Once)
	app.Flag("dry-run", "When enabled, prints DNS record changes rather than actually performing them (default: disabled)").BoolVar(&cfg.DryRun)

	// Miscellaneous flags
	app.Flag("log-format", "The format in which log messages are printed (default: text, options: text, json)").Default(defaultConfig.LogFormat).EnumVar(&cfg.LogFormat, "text", "json")
	app.Flag("metrics-address", "Specify where to serve the metrics and health check endpoint (default: :7979)").Default(defaultConfig.MetricsAddress).StringVar(&cfg.MetricsAddress)
	app.Flag("log-level", "Set the level of logging. (default: info, options: panic, debug, info, warning, error, fatal").Default(defaultConfig.LogLevel).EnumVar(&cfg.LogLevel, allLogLevelsAsStrings()...)

	_, err := app.Parse(args)
	if err != nil {
		return err
	}

	return nil
}
