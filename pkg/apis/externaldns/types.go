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

import "github.com/spf13/pflag"

var (
	defaultHealthPort = "9090"
)

// Config is a project-wide configuration
type Config struct {
	InCluster     bool
	KubeConfig    string
	GoogleProject string
	GoogleZone    string
	HealthPort    string
	DryRun        bool
	Debug         bool
	LogFormat     string
}

// NewConfig returns new Config object
func NewConfig() *Config {
	return &Config{}
}

// ParseFlags adds and parses flags from command line
func (cfg *Config) ParseFlags(args []string) error {
	flags := pflag.NewFlagSet("", pflag.ContinueOnError)
	flags.BoolVar(&cfg.InCluster, "in-cluster", false, "whether to use in-cluster config")
	flags.StringVar(&cfg.KubeConfig, "kubeconfig", "", "path to a local kubeconfig file")
	flags.StringVar(&cfg.GoogleProject, "google-project", "", "gcloud project to target")
	flags.StringVar(&cfg.GoogleZone, "google-zone", "", "gcloud dns hosted zone to target")
	flags.StringVar(&cfg.HealthPort, "health-port", defaultHealthPort, "health port to listen on")
	flags.StringVar(&cfg.LogFormat, "log-format", "text", "log format output. options: [\"text\", \"json\"]")
	flags.BoolVar(&cfg.DryRun, "dry-run", true, "dry-run mode")
	flags.BoolVar(&cfg.Debug, "debug", false, "debug mode")
	return flags.Parse(args)
}
