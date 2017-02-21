package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	defaultHealthPort = "9090"
)

// Config is a project-wide configuration
type config struct {
	healthPort string
	debug      bool
	logFormat  string
}

// NewConfig returns new configuration object
func newConfig() *config {
	return &config{}
}

// ParseFlags adds and parses flags from command line
func (cfg *config) parseFlags() {
	flags := pflag.NewFlagSet("", pflag.ExitOnError)
	flags.StringVar(&cfg.healthPort, "health-port", defaultHealthPort, "health port to listen on")
	flags.StringVar(&cfg.logFormat, "log-format", "text", "log format output")
	flags.BoolVar(&cfg.debug, "debug", false, "debug mode")
	flags.Parse(os.Args)
}

// ValidateFlags custom validation for flags aside from pflag provided
func (cfg *config) validateFlags() error {
	if cfg.logFormat != "text" && cfg.logFormat != "json" {
		return fmt.Errorf("unsupported log format: %s", cfg.logFormat)
	}
	return nil
}
