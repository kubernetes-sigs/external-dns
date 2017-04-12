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

	"github.com/spf13/pflag"

	"k8s.io/client-go/pkg/api/v1"
)

var (
	defaultMetricsAddress = ":7979"
	defaultLogFormat      = "text"
)

// Config is a project-wide configuration
type Config struct {
	InCluster      bool
	KubeConfig     string
	Namespace      string
	Zone           string
	Sources        []string
	Provider       string
	GoogleProject  string
	Policy         string
	Compatibility  bool
	MetricsAddress string
	Interval       time.Duration
	Once           bool
	DryRun         bool
	Debug          bool
	LogFormat      string
	Version        bool
	Registry       string
	RecordOwnerID  string
	TXTPrefix      string
}

// NewConfig returns new Config object
func NewConfig() *Config {
	return &Config{}
}

// ParseFlags adds and parses flags from command line
func (cfg *Config) ParseFlags(args []string) error {
	flags := pflag.NewFlagSet("external-dns", pflag.ContinueOnError)
	flags.BoolVar(&cfg.InCluster, "in-cluster", false, "whether to use in-cluster config")
	flags.StringVar(&cfg.KubeConfig, "kubeconfig", "", "path to a local kubeconfig file")
	flags.StringVar(&cfg.Namespace, "namespace", v1.NamespaceAll, "the namespace to look for endpoints; all namespaces by default")
	flags.StringVar(&cfg.Zone, "zone", "", "the ID of the hosted zone to target")
	flags.StringArrayVar(&cfg.Sources, "source", nil, "the sources to gather endpoints: [service, ingress], e.g. --source service --source ingress")
	flags.StringVar(&cfg.Provider, "provider", "", "the DNS provider to materialize the records in: <aws|google>")
	flags.StringVar(&cfg.GoogleProject, "google-project", "", "gcloud project to target")
	flags.StringVar(&cfg.Policy, "policy", "sync", "the policy to use: <sync|upsert-only>")
	flags.BoolVar(&cfg.Compatibility, "compatibility", false, "enable to process annotation semantics from legacy implementations")
	flags.StringVar(&cfg.MetricsAddress, "metrics-address", defaultMetricsAddress, "address to expose metrics on")
	flags.StringVar(&cfg.LogFormat, "log-format", defaultLogFormat, "log format output: <text|json>")
	flags.DurationVar(&cfg.Interval, "interval", time.Minute, "interval between synchronizations")
	flags.BoolVar(&cfg.Once, "once", false, "run once and exit")
	flags.BoolVar(&cfg.DryRun, "dry-run", true, "run without updating DNS provider")
	flags.BoolVar(&cfg.Debug, "debug", false, "debug mode")
	flags.BoolVar(&cfg.Version, "version", false, "display the version")
	flags.StringVar(&cfg.Registry, "registry", "noop", "type of registry for ownership: <noop|txt>")
	flags.StringVar(&cfg.RecordOwnerID, "record-owner-id", "", "id for keeping track of the managed records")
	flags.StringVar(&cfg.TXTPrefix, "txt-prefix", "", `prefix assigned to DNS name of the associated TXT record; e.g. for --txt-prefix=abc_ [CNAME example.org] <-> [TXT abc_example.org]`)
	return flags.Parse(args)
}
