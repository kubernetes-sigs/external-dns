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
	Namespace      string
	Sources        []string
	FqdnTemplate   string
	Compatibility  string
	Provider       string
	GoogleProject  string
	Domain         string
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
	Namespace:      "",
	Sources:        nil,
	FqdnTemplate:   "",
	Compatibility:  "",
	Provider:       "",
	GoogleProject:  "",
	Domain:         "",
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
	app := kingpin.New("external-dns", "ExternalDNS synchronizes exposed Kubernetes Services and Ingresses with DNS providers.")
	app.Version(version)
	app.DefaultEnvars()

	// Flags related to Kubernetes
	app.Flag("master", "The Kubernetes API server to connect to; defaults to auto-detection from the environment").Default(defaultConfig.Master).StringVar(&cfg.Master)
	app.Flag("kubeconfig", "Retrieve target cluster configuration from a Kubernetes configuration file; defaults to auto-detection from the environment").Default(defaultConfig.KubeConfig).StringVar(&cfg.KubeConfig)
	app.Flag("namespace", "Limit sources of endpoints to a specific namespace; defaults to all").Default(defaultConfig.Namespace).StringVar(&cfg.Namespace)

	// Flags related to processing sources
	app.Flag("source", "The resource types that are queried for endpoints").EnumsVar(&cfg.Sources, "service", "ingress")
	app.Flag("fqdn-template", "A templated string that's used to generate DNS names from sources that don't define a hostname themselves").Default(defaultConfig.FqdnTemplate).StringVar(&cfg.FqdnTemplate)
	app.Flag("compatibility", "Process annotation semantics from legacy implementations; possible choices: <mate|molecule>; defaults to none").Default(defaultConfig.Compatibility).EnumVar(&cfg.Compatibility, "", "mate", "molecule")

	// Flags related to providers
	app.Flag("provider", "The DNS provider where the DNS records will be created").EnumVar(&cfg.Provider, "aws", "google")
	app.Flag("google-project", "When using the Google provider, specify the Google project").Default(defaultConfig.GoogleProject).StringVar(&cfg.GoogleProject)
	app.Flag("domain", "Limit possible target zones by a domain suffix").Default(defaultConfig.Domain).StringVar(&cfg.Domain)

	// Flags related to policies
	app.Flag("policy", "Modify how DNS records are sychronized between sources and providers; possible choices: <sync|upsert-only>; defaults to sync").Default(defaultConfig.Policy).EnumVar(&cfg.Policy, "sync", "upsert-only")

	// Flags related to the registry
	app.Flag("registry", "The registry implementation to use to keep track of DNS record ownership; possible choices: <txt|noop>; defaults to txt").Default(defaultConfig.Registry).EnumVar(&cfg.Registry, "txt", "noop")
	app.Flag("txt-owner-id", "When using the TXT registry, a name that identifies this instance of ExternalDNS").Default(defaultConfig.TXTOwnerID).StringVar(&cfg.TXTOwnerID)
	app.Flag("txt-prefix", "When using the TXT registry, a custom string that's prefixed to each ownership DNS record").Default(defaultConfig.TXTPrefix).StringVar(&cfg.TXTPrefix)

	// Flags related to the main control loop
	app.Flag("interval", "The interval between two consecutive synchronizations in duration format").Default(defaultConfig.Interval.String()).DurationVar(&cfg.Interval)
	app.Flag("once", "When enabled, exits the synchronization loop after the first iteration").BoolVar(&cfg.Once)
	app.Flag("dry-run", "When enabled, prints DNS record changes rather than actually performing them").BoolVar(&cfg.DryRun)

	// Miscellaneous flags
	app.Flag("log-format", "The format in which log messages are printed; possible choices: <text|json>; defaults to txt").Default(defaultConfig.LogFormat).EnumVar(&cfg.LogFormat, "text", "json")
	app.Flag("metrics-address", "Specify were to serve the metrics and health check endpoint").Default(defaultConfig.MetricsAddress).StringVar(&cfg.MetricsAddress)
	app.Flag("debug", "When enabled, increases the logging output for debugging purposes").BoolVar(&cfg.Debug)

	_, err := app.Parse(args)
	if err != nil {
		return err
	}

	return nil
}
