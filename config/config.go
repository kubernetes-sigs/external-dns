package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	defaultHealthPort = "9090"
)

// Config is a project-wide configuration
type Config struct {
	HealthPort string
	Debug      bool
	LogFormat  string
}

// NewConfig returns new Configuration object
func NewConfig() *Config {
	return &Config{}
}

// ParseFlags adds and parses flags from command line
func (cfg *Config) ParseFlags() {
	flags := pflag.NewFlagSet("", pflag.ExitOnError)
	flags.StringVar(&cfg.HealthPort, "health-port", defaultHealthPort, "health port to listen on")
	flags.StringVar(&cfg.LogFormat, "log-format", "text", "log format output. options: [\"text\", \"json\"]")
	flags.BoolVar(&cfg.Debug, "debug", false, "debug mode")
	flags.Parse(os.Args)
}

// Validate custom validation for flags aside from flag library provided
func (cfg *Config) Validate() error {
	if cfg.LogFormat != "text" && cfg.LogFormat != "json" {
		return fmt.Errorf("unsupported log format: %s", cfg.LogFormat)
	}
	return nil
}
