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
)

var (
	version = "unknown"
)

// Config is a project-wide configuration
type Config struct {
	Master         string
	KubeConfig     string
	Sources        []string
	Namespace      string
	FqdnTemplate   string
	Compatibility  string
	Provider       string
	GoogleProject  string
	DomainFilter   string
	Policy         string
	Registry       string
	TXTOwnerID     string
	TXTPrefix      string
	Interval       time.Duration
	Once           bool
	DryRun         bool
	LogFormat      string
	MetricsAddress string
	Debug          bool
}

var defaultConfig = &Config{
	Master:         "",
	KubeConfig:     "",
	Sources:        nil,
	Namespace:      "",
	FqdnTemplate:   "",
	Compatibility:  "",
	Provider:       "",
	GoogleProject:  "",
	DomainFilter:   "",
	Policy:         "sync",
	Registry:       "txt",
	TXTOwnerID:     "default",
	TXTPrefix:      "",
	Interval:       time.Minute,
	Once:           false,
	DryRun:         false,
	LogFormat:      "text",
	MetricsAddress: ":7979",
	Debug:          false,
}

// NewConfig returns new Config object
func NewConfig() *Config {
	return &Config{}
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
	app.Flag("source", "The resource types that are queried for endpoints; specify multiple times for multiple sources (required, options: service, ingress)").Required().PlaceHolder("source").EnumsVar(&cfg.Sources, "service", "ingress")
	app.Flag("namespace", "Limit sources of endpoints to a specific namespace (default: all namespaces)").Default(defaultConfig.Namespace).StringVar(&cfg.Namespace)
	app.Flag("fqdn-template", "A templated string that's used to generate DNS names from sources that don't define a hostname themselves (optional)").Default(defaultConfig.FqdnTemplate).StringVar(&cfg.FqdnTemplate)
	app.Flag("compatibility", "Process annotation semantics from legacy implementations (optional, options: mate, molecule)").Default(defaultConfig.Compatibility).EnumVar(&cfg.Compatibility, "", "mate", "molecule")

	// Flags related to providers
	app.Flag("provider", "The DNS provider where the DNS records will be created (required, options: aws, google)").Required().PlaceHolder("provider").EnumVar(&cfg.Provider, "aws", "google")
	app.Flag("google-project", "When using the Google provider, specify the Google project (required when --provider=google)").Default(defaultConfig.GoogleProject).StringVar(&cfg.GoogleProject)
	app.Flag("domain-filter", "Limit possible target zones by a domain suffix (optional)").Default(defaultConfig.DomainFilter).StringVar(&cfg.DomainFilter)

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
	app.Flag("metrics-address", "Specify were to serve the metrics and health check endpoint (default: :7979)").Default(defaultConfig.MetricsAddress).StringVar(&cfg.MetricsAddress)
	app.Flag("debug", "When enabled, increases the logging output for debugging purposes (default: disabled)").BoolVar(&cfg.Debug)

	_, err := app.Parse(args)
	if err != nil {
		return err
	}

	return nil
}
