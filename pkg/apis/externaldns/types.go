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
	version               = "unknown"
	defaultMetricsAddress = ":7979"
	defaultLogFormat      = "text"
)

// Config is a project-wide configuration
type Config struct {
	InCluster      bool
	KubeConfig     string
	Namespace      string
	Domain         string
	Sources        []string
	Provider       string
	GoogleProject  string
	Policy         string
	Compatibility  string
	MetricsAddress string
	Interval       time.Duration
	Once           bool
	DryRun         bool
	Debug          bool
	LogFormat      string
	Registry       string
	RecordOwnerID  string
	TXTPrefix      string
	FqdnTemplate   string
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

	app.Flag("in-cluster", "Retrieve target cluster configuration from the environment").BoolVar(&cfg.InCluster)
	app.Flag("kubeconfig", "Retrieve target cluster configuration from a Kubernetes configuration file").StringVar(&cfg.KubeConfig)
	app.Flag("namespace", "Limit sources of endpoints to a specific namespace; defaults to all").StringVar(&cfg.Namespace)
	app.Flag("domain", "Limit possible target zones by a domain suffix").StringVar(&cfg.Domain)
	app.Flag("source", "The resource types that are queried for endpoints").StringsVar(&cfg.Sources)
	app.Flag("provider", "The DNS provider where the DNS records will be created").StringVar(&cfg.Provider)
	app.Flag("google-project", "When using the Google provider, specify the Google project").StringVar(&cfg.GoogleProject)
	app.Flag("policy", "Modify how DNS records are sychronized between sources and providers; possible choices: <sync|upsert-only> default: sync").Default("sync").StringVar(&cfg.Policy)
	app.Flag("compatibility", "Process annotation semantics from legacy implementations; possible choices: <mate|molecule>, default: none").StringVar(&cfg.Compatibility)
	app.Flag("metrics-address", "Specify were to serve the metrics and health check endpoint").Default(defaultMetricsAddress).StringVar(&cfg.MetricsAddress)
	app.Flag("log-format", "The format in which log messages are printed").Default(defaultLogFormat).StringVar(&cfg.LogFormat)
	app.Flag("interval", "The interval between two consecutive synchronizations in duration format").Default("1m").DurationVar(&cfg.Interval)
	app.Flag("once", "When enabled, exits the synchronization loop after the first iteration").BoolVar(&cfg.Once)
	app.Flag("dry-run", "When enabled, prints DNS record changes rather than actually performing them").BoolVar(&cfg.DryRun)
	app.Flag("debug", "When enabled, increases the logging output for debugging purposes").BoolVar(&cfg.Debug)
	app.Flag("registry", "The registry implementation to use to keep track of DNS record ownership").Default("txt").StringVar(&cfg.Registry)
	app.Flag("record-owner-id", "When using the TXT registry, a name that identifies this instance of ExternalDNS").Default("default").StringVar(&cfg.RecordOwnerID)
	app.Flag("txt-prefix", "When using the TXT registry, a custom string that's prefixed to each ownership DNS record").StringVar(&cfg.TXTPrefix)
	app.Flag("fqdn-template", "A templated string that's used to generate DNS names from sources that don't define a hostname themselves").StringVar(&cfg.FqdnTemplate)

	_, err := app.Parse(args)
	if err != nil {
		return err
	}

	return nil
}
