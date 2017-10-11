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
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/sirupsen/logrus"
)

var (
	version = "unknown"
)

// Config is a project-wide configuration
type Config struct {
	Master             string
	KubeConfig         string
	Sources            []string
	Namespace          string
	FQDNTemplate       string
	Compatibility      string
	PublishInternal    bool
	Provider           string
	GoogleProject      string
	DomainFilter       []string
	AWSZoneType        string
	AzureConfigFile    string
	AzureResourceGroup string
	CloudflareProxied  bool
	Policy             string
	Registry           string
	TXTOwnerID         string
	TXTPrefix          string
	Interval           time.Duration
	Once               bool
	DryRun             bool
	LogFormat          string
	MetricsAddress     string
	LogLevel           string
}

var defaultConfig = &Config{
	Master:             "",
	KubeConfig:         "",
	Sources:            nil,
	Namespace:          "",
	FQDNTemplate:       "",
	Compatibility:      "",
	PublishInternal:    false,
	Provider:           "",
	GoogleProject:      "",
	DomainFilter:       []string{},
	AWSZoneType:        "",
	AzureConfigFile:    "/etc/kubernetes/azure.json",
	AzureResourceGroup: "",
	CloudflareProxied:  false,
	Policy:             "sync",
	Registry:           "txt",
	TXTOwnerID:         "default",
	TXTPrefix:          "",
	Interval:           time.Minute,
	Once:               false,
	DryRun:             false,
	LogFormat:          "text",
	MetricsAddress:     ":7979",
	LogLevel:           logrus.InfoLevel.String(),
}

// NewConfig returns new Config object
func NewConfig() *Config {
	return &Config{}
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
	app.Version(version)
	app.DefaultEnvars()

	// Flags related to Kubernetes
	app.Flag("master", "The Kubernetes API server to connect to (default: auto-detect)").Default(defaultConfig.Master).StringVar(&cfg.Master)
	app.Flag("kubeconfig", "Retrieve target cluster configuration from a Kubernetes configuration file (default: auto-detect)").Default(defaultConfig.KubeConfig).StringVar(&cfg.KubeConfig)

	// Flags related to processing sources
	app.Flag("source", "The resource types that are queried for endpoints; specify multiple times for multiple sources (required, options: service, ingress, fake)").Required().PlaceHolder("source").EnumsVar(&cfg.Sources, "service", "ingress", "fake")
	app.Flag("namespace", "Limit sources of endpoints to a specific namespace (default: all namespaces)").Default(defaultConfig.Namespace).StringVar(&cfg.Namespace)
	app.Flag("fqdn-template", "A templated string that's used to generate DNS names from sources that don't define a hostname themselves, or to add a hostname suffix when paired with the fake source (optional)").Default(defaultConfig.FQDNTemplate).StringVar(&cfg.FQDNTemplate)
	app.Flag("compatibility", "Process annotation semantics from legacy implementations (optional, options: mate, molecule)").Default(defaultConfig.Compatibility).EnumVar(&cfg.Compatibility, "", "mate", "molecule")
	app.Flag("publish-internal-services", "Allow external-dns to publish DNS records for ClusterIP services (optional)").BoolVar(&cfg.PublishInternal)

	// Flags related to providers
	app.Flag("provider", "The DNS provider where the DNS records will be created (required, options: aws, google, azure, cloudflare, digitalocean, dnsimple, inmemory)").Required().PlaceHolder("provider").EnumVar(&cfg.Provider, "aws", "google", "azure", "cloudflare", "digitalocean", "dnsimple", "inmemory")
	app.Flag("domain-filter", "Limit possible target zones by a domain suffix; specify multiple times for multiple domains (optional)").Default("").StringsVar(&cfg.DomainFilter)
	app.Flag("google-project", "When using the Google provider, specify the Google project (required when --provider=google)").Default(defaultConfig.GoogleProject).StringVar(&cfg.GoogleProject)
	app.Flag("aws-zone-type", "When using the AWS provider, filter for zones of this type (optional, options: public, private)").Default(defaultConfig.AWSZoneType).EnumVar(&cfg.AWSZoneType, "", "public", "private")
	app.Flag("azure-config-file", "When using the Azure provider, specify the Azure configuration file (required when --provider=azure").Default(defaultConfig.AzureConfigFile).StringVar(&cfg.AzureConfigFile)
	app.Flag("azure-resource-group", "When using the Azure provider, override the Azure resource group to use (optional)").Default(defaultConfig.AzureResourceGroup).StringVar(&cfg.AzureResourceGroup)
	app.Flag("cloudflare-proxied", "When using the Cloudflare provider, specify if the proxy mode must be enabled (default: disabled)").BoolVar(&cfg.CloudflareProxied)

	// Flags related to policies
	app.Flag("policy", "Modify how DNS records are sychronized between sources and providers (default: sync, options: sync, upsert-only)").Default(defaultConfig.Policy).EnumVar(&cfg.Policy, "sync", "upsert-only")

	// Flags related to the registry
	app.Flag("registry", "The registry implementation to use to keep track of DNS record ownership (default: txt, options: txt, noop)").Default(defaultConfig.Registry).EnumVar(&cfg.Registry, "txt", "noop")
	app.Flag("txt-owner-id", "When using the TXT registry, a name that identifies this instance of ExternalDNS (default: default)").Default(defaultConfig.TXTOwnerID).StringVar(&cfg.TXTOwnerID)
	app.Flag("txt-prefix", "When using the TXT registry, a custom string that's prefixed to each ownership DNS record (optional)").Default(defaultConfig.TXTPrefix).StringVar(&cfg.TXTPrefix)

	// Flags related to the main control loop
	app.Flag("interval", "The interval between two consecutive synchronizations in duration format (default: 1m)").Default(defaultConfig.Interval.String()).DurationVar(&cfg.Interval)
	app.Flag("once", "When enabled, exits the synchronization loop after the first iteration (default: disabled)").BoolVar(&cfg.Once)
	app.Flag("dry-run", "When enabled, prints DNS record changes rather than actually performing them (default: disabled)").BoolVar(&cfg.DryRun)

	// Miscellaneous flags
	app.Flag("log-format", "The format in which log messages are printed (default: text, options: text, json)").Default(defaultConfig.LogFormat).EnumVar(&cfg.LogFormat, "text", "json")
	app.Flag("metrics-address", "Specify where to serve the metrics and health check endpoint (default: :7979)").Default(defaultConfig.MetricsAddress).StringVar(&cfg.MetricsAddress)
	app.Flag("log-level", "Set the level of logging. (default: info, options: panic, debug, info, warn, error, fatal").Default(defaultConfig.LogLevel).EnumVar(&cfg.LogLevel, allLogLevelsAsStrings()...)

	_, err := app.Parse(args)
	if err != nil {
		return err
	}

	return nil
}
